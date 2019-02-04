package mongoDB

import (
	"fmt"
	"github.com/sqmmm/event-service/domain"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type repository struct {
	db *mgo.Database
}

func NewRepository(db *mgo.Database) *repository{
	return &repository{db: db}
}

func (r *repository) GetEvent(tp string) (*domain.Event, error) {
	event := domain.Event{}
	err := r.db.C("events").Find(bson.M{"type": tp}).One(&event)
	if err != nil {
		if err == mgo.ErrNotFound{
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("failed to get event: %s", err)
	}

	return &event, nil
}

func (r *repository) StartNewEvent(tp string) error {
	event := &domain.Event{ID: bson.NewObjectId(), Type: tp, State: 0}

	err := r.db.C("events").Insert(event)
	if err != nil {
		return fmt.Errorf("failed to add new event: %s" ,err)
	}

	return nil
}

func (r *repository) StartOldEvent(tp string) error {
	err := r.db.C("events").Update(bson.M{"type": tp}, bson.M{"$set": bson.M{"state": 0}})
	if err != nil {
		return fmt.Errorf("failed to add new event: %s" ,err)
	}

	return nil
}

func (r *repository) FinishEvent(tp string) error {
	err := r.db.C("events").Update(bson.M{"type": tp}, bson.M{"$set": bson.M{"state": 1}})
	if err != nil {
		return fmt.Errorf("failed to add new event: %s" ,err)
	}

	return nil
}