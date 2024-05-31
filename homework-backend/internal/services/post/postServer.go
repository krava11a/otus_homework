package post

import (
	"context"
	"homework-backend/internal/models"
	"homework-backend/internal/proto"

	"google.golang.org/grpc"
)

type GrpcPost interface {
	PostCreate(post models.Post) (string, error)
	PostUpdate(post_id, text string) error
	PostDelete(post_id string) error
	PostGet(post_id string) (post models.Post, err error)
	PostFeed(user_id string, offset, limit uint32) (models.Posts, error)

	FriendSet(user_id, friend_id string) error
	FriendDelete(user_id, friend_id string) error
}

func Register(gRPCServer *grpc.Server, grpcPost GrpcPost) {
	proto.RegisterPostServiceServer(gRPCServer, &PostServer{grpcPost: grpcPost})

}

type PostServer struct {
	proto.UnimplementedPostServiceServer
	grpcPost GrpcPost
}

func (ps *PostServer) PostCreate(ctx context.Context, req *proto.PostCreateRequest) (*proto.PostCreateResponse, error) {
	nPost := models.Post{
		Id_user: req.UserId,
		Text:    req.Text,
	}

	post_uuid, err := ps.grpcPost.PostCreate(nPost)
	if err != nil {
		return &proto.PostCreateResponse{
			Status:  400,
			Message: err.Error(),
			PostId:  post_uuid,
		}, err
	}
	return &proto.PostCreateResponse{
		Status:  200,
		Message: "Все хорошо",
		PostId:  post_uuid,
	}, nil

}

func (ps *PostServer) PostUpdate(ctx context.Context, req *proto.PostUpdateRequest) (*proto.PostMainResponse, error) {

	err := ps.grpcPost.PostUpdate(req.PostId, req.Text)
	if err != nil {
		return &proto.PostMainResponse{
			Status:  500,
			Message: err.Error(),
		}, err
	}
	return &proto.PostMainResponse{
		Status:  200,
		Message: "Все хорошо",
	}, nil
}

func (ps *PostServer) PostDelete(ctx context.Context, req *proto.PostMainRequest) (*proto.PostMainResponse, error) {
	err := ps.grpcPost.PostDelete(req.PostId)
	if err != nil {
		return &proto.PostMainResponse{
			Status:  500,
			Message: err.Error(),
		}, err
	}
	return &proto.PostMainResponse{
		Status:  200,
		Message: "Все хорошо",
	}, nil
}

func (ps *PostServer) PostGet(ctx context.Context, req *proto.PostMainRequest) (*proto.PostGetResponse, error) {
	post, err := ps.grpcPost.PostGet(req.PostId)
	if err != nil {
		return &proto.PostGetResponse{
			Status:  500,
			Message: err.Error(),
			Post:    &proto.Post{},
		}, err
	}

	return &proto.PostGetResponse{
		Status:  200,
		Message: "Все хорошо",
		Post: &proto.Post{
			UserId: post.Id_user,
			PostId: post.Id_post,
			Text:   post.Text,
		},
	}, nil
}

func (ps *PostServer) PostFeed(ctx context.Context, req *proto.PostFeedRequest) (*proto.PostFeedResponse, error) {

	posts, err := ps.grpcPost.PostFeed(req.UserId, req.Offset, req.Limit)
	if err != nil {
		return &proto.PostFeedResponse{
			Status:  500,
			Message: err.Error(),
			Posts:   make([]*proto.Post, 0),
		}, err
	}

	postsProto := make([]*proto.Post, 0, req.Limit+1)
	for _, post := range posts.Posts {
		postsProto = append(postsProto, &proto.Post{
			UserId: post.Id_user,
			Text:   post.Text,
			PostId: post.Id_post,
		})
	}

	return &proto.PostFeedResponse{
		Status:  200,
		Message: "Все хорошо",
		Posts:   postsProto,
	}, err
}

func (ps *PostServer) FriendSet(ctx context.Context, req *proto.FriendRequest) (*proto.FriendResponse, error) {
	err := ps.grpcPost.FriendSet(req.UserId, req.FriendId)
	if err != nil {
		return &proto.FriendResponse{
			Status:  500,
			Message: err.Error(),
		}, err
	}
	return &proto.FriendResponse{
		Status:  200,
		Message: "Все хорошо",
	}, err
}

func (ps *PostServer) FriendDelete(ctx context.Context, req *proto.FriendRequest) (*proto.FriendResponse, error) {
	err := ps.grpcPost.FriendDelete(req.UserId, req.FriendId)
	if err != nil {
		return &proto.FriendResponse{
			Status:  500,
			Message: err.Error(),
		}, err
	}
	return &proto.FriendResponse{
		Status:  200,
		Message: "Все хорошо",
	}, err
}
