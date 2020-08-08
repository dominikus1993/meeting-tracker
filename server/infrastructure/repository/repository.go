package repository

import "go.mongodb.org/mongo-driver/mongo"

type mongoMeetingRepository struct {
	redis *mongo.Client
}
