package usecases

import (
	"chat_receive_service/infra/repositories"
)

type ListenMessagesUC struct {
	repository repositories.MessageRepository
}

func (uc *ListenMessagesUC) ListenMessages() {

	uc.repository.ListenQueueMessages()

}
