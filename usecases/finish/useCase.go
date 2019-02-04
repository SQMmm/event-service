package finish

import "fmt"

type UseCase interface {
	Finish(event string) error
}

type eventRepository interface {
	FinishEvent(string) error
}

type useCase struct {
	eventRepo eventRepository
}

func NewUseCase(eventRepo eventRepository) *useCase {
	return &useCase{eventRepo: eventRepo}
}

func (u *useCase) Finish(tp string) error {
	err := u.eventRepo.FinishEvent(tp)
	if err != nil {
		return fmt.Errorf("failed to finish event: %s", err)
	}

	return nil
}
