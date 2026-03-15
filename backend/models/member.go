package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Member struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name" json:"name"`
	SurName     string             `bson:"surname" json:"surname"`
	Address     string             `bson:"address" json:"address"`
	DateOfBirth string             `bson:"dateOfBirth" json:"dateOfBirth"`
	Age         int                `bson:"age" json:"age"`
	CreatedAt   time.Time          `bson:"createdAt" json:"createdAt"`
}

type MemberResponse struct {
	Name        string `bson:"name" json:"name"`
	SurName     string `bson:"surname" json:"surname"`
	Address     string `bson:"address" json:"address"`
	Age         int    `bson:"age" json:"age"`
	DateOfBirth string `bson:"dateOfBirth" json:"dateOfBirth"`
}

type PaginatedResponse struct {
	Data       []MemberResponse `json:"data"`
	Total      int64            `json:"total"`
	Page       int              `json:"page"`
	Limit      int              `json:"limit"`
	TotalPages int64            `json:"totalPages"`
}
