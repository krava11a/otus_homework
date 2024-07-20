package dialog

import "homework-dialogs/internal/models"

type DialogSender interface {
	Send(dialog models.Dialog) error
}

type DialogListener interface {
	List(from, to string) (dialogs []models.Dialog, err error)
}
