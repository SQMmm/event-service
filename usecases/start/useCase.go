package start

import (
	"fmt"
	"github.com/sqmmm/event-service/domain"
)

type UseCase interface {
	Start(event string) error
}

type eventRepository interface {
	GetEvent(string) (*domain.Event, error)
	StartNewEvent(string) error
	StartOldEvent(string) error
}

type useCase struct {
	eventRepo eventRepository
}

func NewUseCase(eventRepo eventRepository) *useCase {
	return &useCase{eventRepo: eventRepo}
}

func (u *useCase) Start(tp string) error {
	event, err := u.eventRepo.GetEvent(tp)
	if err != nil {
		if err == domain.ErrNotFound {
			err = u.eventRepo.StartNewEvent(tp)
			if err != nil {
				return fmt.Errorf("failed to start new event: %s", err)
			}
			return nil
		}

		return fmt.Errorf("failed to check if event exists: %s", err)
	}

	if event.State == 1 {
		err = u.eventRepo.StartOldEvent(tp)
		if err != nil {
			return fmt.Errorf("failed to start old event: %s", err)
		}
	}

	return nil
}
