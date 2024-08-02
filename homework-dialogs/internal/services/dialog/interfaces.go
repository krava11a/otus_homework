package dialog

import "homework-dialogs/internal/models"

type DialogSender interface {
	Send(dialog models.Dialog, xid string) error
}

type DialogListener interface {
	List(from, to, xid string) (dialogs []models.Dialog, err error)
}
