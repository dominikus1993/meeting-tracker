package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
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

func (r *mongoMeetingRepository) GetId() (*string, error) {
	panic("implement me")
}

func (r *mongoMeetingRepository) Start(states *model.Meeting, ctx context.Context) (*model.Meeting, error) {
	panic("implement me")
}

func (r *mongoMeetingRepository) Finish(meeting *model.Meeting, ctx context.Context) (*model.Meeting, error) {
	panic("implement me")
}

func (r *mongoMeetingRepository) GetAll(ctx context.Context) ([]*model.Meeting, error) {
	cur, err := r.db.Collection(MeetingCollectionName).Find(ctx, bson.D{{}})

	if err != nil {
		return nil, err
	}
	var results []*model.Meeting
	for cur.Next(ctx) {
		var result inframodel.MongoMeeting
		e := cur.Decode(&result)
		if e != nil {
			log.Println(e)
		}
		results = append(results, result.ToDomainMeeting())

	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(ctx)
	return results, nil
}

func (r *mongoMeetingRepository) GetById(id string, ctx context.Context) (*model.Meeting, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objId}
	var mongoModel inframodel.MongoMeeting
	err = r.db.Collection(MeetingCollectionName).FindOne(ctx, filter).Decode(&mongoModel)
	if err != nil {
		return nil, err
	}
	return mongoModel.ToDomainMeeting(), nil
}

func NewMeetingRepository(mongo *mongo.Client) *mongoMeetingRepository {
	return &mongoMeetingRepository{mongo: mongo, db: mongo.Database("Meetings")}
}

func (r *mongoMeetingRepository) FindById(id string, ctx context.Context) (*model.Meeting, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objId}
	var mongoModel inframodel.MongoMeeting
	err = r.db.Collection(MeetingCollectionName).FindOne(ctx, filter).Decode(&mongoModel)
	if err != nil {
		return nil, err
	}
	return mongoModel.ToDomainMeeting(), nil
}

// func (r *mongoMeetingRepository) Start(states *model.Meeting, ctx context.Context) (*model.Meeting, error) {
// }
