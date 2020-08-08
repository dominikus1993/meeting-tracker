package repository

import (
	"context"
	"server/domain/model"
	inframodel "server/infrastructure/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const MeetingCollectionName = "meetings"

type mongoMeetingRepository struct {
	mongo *mongo.Client
	db    *mongo.Database
}

func NewMeetingRepository(mongo *mongo.Client) *mongoMeetingRepository {
	return &mongoMeetingRepository{mongo: mongo, db: mongo.Database("Meetings")}
}

func (r *mongoMeetingRepository) FindById(id string, ctx context.Context) (*model.Meeting, error) {
	filter := bson.M{"_id": id}
	var mongoModel inframodel.MongoMeeting
	err := r.db.Collection(MeetingCollectionName).FindOne(ctx, filter).Decode(&mongoModel)
	if err != nil {
		return nil, err
	}
	return &model.Meeting{MeetingId: "dsdsaasd"}, nil
}

func (r *mongoMeetingRepository) Start(states *model.Meeting, ctx context.Context) (*model.Meeting, error) {
	filter := bson.M{"_id": id}
	var mongoModel inframodel.MongoMeeting
	err := r.db.Collection(MeetingCollectionName).FindOne(ctx, filter).Decode(&mongoModel)
	if err != nil {
		return nil, err
	}
	return &model.Meeting{MeetingId: "dsdsaasd"}, nil
}
