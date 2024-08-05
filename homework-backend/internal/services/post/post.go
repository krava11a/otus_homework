package post

import (
	"errors"
	"fmt"
	"strconv"

	"homework-backend/internal/lib/logger/sl"
	"homework-backend/internal/models"
	"homework-backend/internal/storage"
	"log/slog"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Post struct {
	log         *slog.Logger
	pstCreater  PostCreater
	pstProvider PostProvider
	cache       Cache
	rqueue      RQueuePublisher
}

func New(
	log *slog.Logger,
	postCreater PostCreater,
	postProvider PostProvider,
	cache Cache,
	rqueue RQueuePublisher,
) *Post {

	p := Post{
		log:         log,
		pstCreater:  postCreater,
		pstProvider: postProvider,
		cache:       cache,
		rqueue:      rqueue,
	}
	go p.UpdateFeedAll()

	return &p
}

func (p *Post) PostCreate(post models.Post) (string, error) {

	const op = "PostCreate"

	log := p.log.With(
		slog.String("op", op),
		slog.String("user id", post.Id_user),
	)

	log.Info("creating post")

	if post.Id_user == "" {
		return "", status.Error(codes.InvalidArgument, "Id_user is required")
	}

	if post.Text == "" {
		return "", status.Error(codes.InvalidArgument, "Text is required")
	}

	uuid, err := p.pstCreater.PostCreate(post)
	if err != nil {
		log.Error("failed to create post", sl.Err(err))

		return "", fmt.Errorf("%s: %w", op, err)
	}
	go p.PublishTo(post.Id_user, post.Text)

	post.Id_post = uuid
	// posts := Posts{}

	// err = p.cache.Get(post.Id_user, &posts)
	// if err != nil {
	// 	log.Error("failed to create post", sl.Err(err), op)
	// }
	// if len(posts.Posts) < 1000 {
	// 	posts.Posts = append(posts.Posts, post)
	// }

	// postsFB, err := p.PostFeed(post.Id_user, 0, 1000)
	// if err != nil {
	// 	log.Error("failed to get posts", sl.Err(err), op)
	// }

	go p.UpdateFeed(post.Id_user)
	// err = p.cache.Set(post.Id_user, Posts{Posts: postsFB}, 0)
	// if err != nil {
	// 	log.Error("failed to create post", sl.Err(err), op)
	// }

	return uuid, nil
}

func (p *Post) PostUpdate(post_id, text string) error {

	const op = "Post.Update"
	log := p.log.With(
		slog.String("op", op),
		slog.String("post_id", post_id),
	)
	log.Info("updating post")

	if post_id == "" {
		return status.Error(codes.InvalidArgument, "post_id is required")
	}

	if text == "" {
		return status.Error(codes.InvalidArgument, "text is required")
	}

	err := p.pstCreater.PostUpdate(post_id, text)
	if err != nil {
		if errors.Is(err, storage.ErrPostNotFound) {
			log.Warn("post not found", sl.Err(err))

			return fmt.Errorf("%s: %w", op, err)
		}

		log.Error("failed to get post", sl.Err(err))

		return fmt.Errorf("%s: %w", op, err)
	}
	post, _ := p.pstProvider.PostGet(post_id)
	go p.UpdateFeed(post.Id_user)
	return nil
}

func (p *Post) PostDelete(post_id string) error {
	const op = "Post.Delete"
	log := p.log.With(
		slog.String("op", op),
		slog.String("post_id", post_id),
	)
	log.Info("deleting post by ID")

	if post_id == "" {
		return status.Error(codes.InvalidArgument, "post_id is required")
	}

	post, err := p.pstProvider.PostGet(post_id)
	if err != nil {
		log.Error("failed to delete post by id", sl.Err(err))

		return fmt.Errorf("%s: %w", op, err)
	}

	err = p.pstCreater.PostDelete(post_id)
	if err != nil {
		log.Error("failed to delete post by id", sl.Err(err))

		return fmt.Errorf("%s: %w", op, err)
	}

	go p.UpdateFeed(post.Id_user)

	return nil

}

func (p *Post) PostGet(post_id string) (post models.Post, err error) {
	const op = "Post.Get"
	log := p.log.With(
		slog.String("op", op),
		slog.String("post_id", post_id),
	)
	log.Info("attempting to SEARCH post by post_id")

	post, err = p.pstProvider.PostGet(post_id)
	if err != nil {
		if errors.Is(err, storage.ErrPostNotFound) {
			log.Warn("post not found", sl.Err(err))

			return models.Post{}, fmt.Errorf("%s: %w", op, err)
		}

		log.Error("failed to get post", sl.Err(err))

		return models.Post{}, fmt.Errorf("%s: %w", op, err)
	}
	return post, nil

}

func (p *Post) PostFeed(user_id string, offset, limit uint32) (models.Posts, error) {
	const op = "Post.Feed"
	log := p.log.With(
		slog.String("op", op),
		slog.String("user_id", user_id),
		slog.String("offset:", strconv.Itoa(int(offset))),
		slog.String("limit:", strconv.Itoa(int(limit))),
	)
	log.Info("attempting to feed posts by user_by")

	if user_id == "" {
		return models.Posts{}, status.Error(codes.InvalidArgument, "user_id is required")
	}
	if limit == 0 {
		limit = 1000
	}
	ps := models.Posts{}
	psCache := models.Posts{}
	err := p.cache.Get(user_id, &psCache)
	if err != nil {
		psDb, err := p.pstProvider.PostFeed(user_id, offset, limit)
		if err != nil {
			log.Error("failed to feed posts by user_id", sl.Err(err))

			return models.Posts{}, fmt.Errorf("%s: %w", op, err)
		}
		if len(psDb.Posts) == 0 {
			psDb, err = p.pstProvider.PostFeed(user_id, 0, 1000)
			if len(psDb.Posts) > 0 {
				go p.UpdateFeedAll()
			}
		}
		ps = psDb
	} else {
		if offset > uint32(len(psCache.Posts)) {
			ofs := int32(len(psCache.Posts)) - int32(limit)
			if ofs < 0 {
				offset = 0
				limit = uint32(len(psCache.Posts))
			}

		}
		for i := offset; i < offset+limit; i++ {
			ps.Posts = append(ps.Posts, psCache.Posts[i])
		}
	}

	return ps, nil

}

// update feed by author_id of post
func (p *Post) UpdateFeed(author_id string) error {
	const op = "Post.UpdateLenta"
	log := p.log.With(
		slog.String("op", op),
		slog.String("author_id", author_id),
	)
	log.Info("attempting to update feed posts by author_id")

	f_ids, err := p.getFriends(author_id)
	if err != nil {
		log.Error("failed to get friends by user_id", sl.Err(err))

		return fmt.Errorf("%s: %w", op, err)
	}
	for _, f_id := range f_ids {
		posts, err := p.pstProvider.PostFeed(f_id, 0, 1000)
		if err != nil {
			log.Error("failed to feed posts by user_id", sl.Err(err))

			return fmt.Errorf("%s: %w", op, err)
		}
		p.cache.Set(f_id, posts, 0)
	}

	return nil

}

func (p *Post) getFriends(author_id string) ([]string, error) {
	const op = "Post.UpdateLenta"

	if author_id == "" {
		return []string{}, status.Error(codes.InvalidArgument, "user_id is required")
	}

	f_ids, err := p.pstProvider.PostFriends(author_id)
	if err != nil {
		p.log.Error("failed to get friends by user_id", sl.Err(err))

		return []string{}, fmt.Errorf("%s: %w", op, err)
	}

	return f_ids, nil
}

func (p *Post) UpdateFeedAll() error {
	const op = "Post.UpdateFeedAll"
	log := p.log.With(
		slog.String("op", op),
	)
	log.Info("attempting to update ALL feed posts")

	f_ids, err := p.pstProvider.PostFriends("0")
	if err != nil {
		log.Error("failed to get friends by user_id", sl.Err(err))

		return fmt.Errorf("%s: %w", op, err)
	}

	for _, f_id := range f_ids {
		posts, err := p.pstProvider.PostFeed(f_id, 0, 1000)
		if err != nil {
			log.Error("failed to feed posts by user_id", sl.Err(err))

			return fmt.Errorf("%s: %w", op, err)
		}
		p.cache.Set(f_id, posts, 0)
	}

	return nil

}

func (p *Post) FriendSet(user_id, friend_id string) (err error) {
	const op = "Post.FriendSet"
	log := p.log.With(
		slog.String("op", op),
		slog.String("user_id", user_id),
		slog.String("friend_id", friend_id),
	)
	log.Info("attempting to Set friend for user by ID")

	if user_id == "" {
		return status.Error(codes.InvalidArgument, "user_id is required")
	}
	if friend_id == "" {
		return status.Error(codes.InvalidArgument, "friend_id is required")
	}

	err = p.pstCreater.FriendSet(user_id, friend_id)
	if err != nil {
		log.Error("failed to set friend for user by id", sl.Err(err))

		return fmt.Errorf("%s: %w", op, err)
	}

	go p.UpdateFeed(friend_id)
	return nil

}

func (p *Post) FriendDelete(user_id, friend_id string) (err error) {
	const op = "Post.FriendDelete"
	log := p.log.With(
		slog.String("op", op),
		slog.String("user_id", user_id),
		slog.String("friend_id", friend_id),
	)
	log.Info("attempting to delete friend for user by ID")

	if user_id == "" {
		return status.Error(codes.InvalidArgument, "user_id is required")
	}
	if friend_id == "" {
		return status.Error(codes.InvalidArgument, "friend_id is required")
	}

	err = p.pstCreater.FriendDelete(user_id, friend_id)
	if err != nil {
		log.Error("failed to delete friend for user by id", sl.Err(err))

		return fmt.Errorf("%s: %w", op, err)
	}
	go p.UpdateFeedAll()
	return nil

}

func (p *Post) PublishTo(name, message string) {
	const op = "Post.PublishTo"
	log := p.log.With(
		slog.String("op", op),
		slog.String("name", name),
		slog.String("message", message),
	)
	log.Info("attempting to publish to RQueue")

	f_ids, err := p.getFriends(name)
	if err != nil {
		log.Error("failed to publish messages to RQUEUE", sl.Err(err))

		return
	}

	for _, id_subscriber := range f_ids {
		err = p.rqueue.PublishTo(id_subscriber, message)
		if err != nil {
			log.Error("failed to publish messages to RQUEUE", sl.Err(err))

			return
		}
	}

}
