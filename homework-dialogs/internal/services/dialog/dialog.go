package dialog

import (
	"fmt"
	"homework-dialogs/internal/models"

	"log/slog"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Dialog struct {
	log         *slog.Logger
	dlgSender   DialogSender
	dlgListener DialogListener
}

func New(log *slog.Logger, dialogSender DialogSender, dialogListener DialogListener) *Dialog {
	d := Dialog{
		log:         log,
		dlgSender:   dialogSender,
		dlgListener: dialogListener,
	}

	return &d
}

func (d *Dialog) Send(dialog models.Dialog, xid string) error {
	const op = "Dialog Send"

	log := d.log.With(
		slog.String("op", op),
		slog.String("dialog", fmt.Sprintf("dialog: from= %s,to= %s, text= %s", dialog.From, dialog.To, dialog.Text)),
		slog.String("X-Request-ID", xid),
	)
	log.Info("sending message Dialog")
	if dialog.From == "" {
		return status.Error(codes.InvalidArgument, "From is required")
	}
	if dialog.To == "" {
		return status.Error(codes.InvalidArgument, "To is required")
	}

	return d.dlgSender.Send(dialog, xid)
}

func (d *Dialog) List(from, to, xid string) (dialogs []models.Dialog, err error) {
	const op = "Dialog List"

	log := d.log.With(
		slog.String("op", op),
		slog.String("from", from),
		slog.String("to", to),
		slog.String("X-Request-ID", xid),
	)
	log.Info("sending message Dialog")
	if from == "" {
		return nil, status.Error(codes.InvalidArgument, "User_id from is required")
	}

	if to == "" {
		return nil, status.Error(codes.InvalidArgument, "User_id to is required")
	}

	return d.dlgListener.List(from, to, xid)
}
