package services

import (
	"context"
	"math"
	"time"

	"backend/config"
	"backend/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const collectionName = "members"

// func calculateAge(dobStr string) int {
// 	dob, err := time.Parse("2006-01-02", dobStr) // yyyy-mm-dd
// 	if err != nil {
// 		return 0
// 	}
// 	now := time.Now()
// 	age := now.Year() - dob.Year()
// 	if now.YearDay() < dob.YearDay() {
// 		age--
// 	}
// 	return age
// }

func toResponse(m models.Member) models.MemberResponse {
	return models.MemberResponse{
		Name:        m.Name,
		SurName:     m.SurName,
		Address:     m.Address,
		DateOfBirth: m.DateOfBirth,
		Age:         m.Age,
		// Age:         calculateAge(m.DateOfBirth),
	}
}

func CreateMember(req *models.Member) (*models.MemberResponse, error) {
	collection := config.DB.Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	member := models.Member{
		ID:          primitive.NewObjectID(),
		Name:        req.Name,
		SurName:     req.SurName,
		Address:     req.Address,
		DateOfBirth: req.DateOfBirth,
		Age:         req.Age,
		CreatedAt:   time.Now(),
	}

	_, err := collection.InsertOne(ctx, member)
	if err != nil {
		return nil, err
	}

	resp := toResponse(member)
	return &resp, nil
}

func GetMembers(page, limit int) (*models.PaginatedResponse, error) {
	collection := config.DB.Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 10
	}

	skip := int64((page - 1) * limit)

	// Count total documents
	total, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	// Find with pagination, sorted by created_at ascending
	findOptions := options.Find()
	findOptions.SetSkip(skip)
	findOptions.SetLimit(int64(limit))
	findOptions.SetSort(bson.D{{Key: "createdAt", Value: 1}})

	cursor, err := collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var members []models.Member
	if err = cursor.All(ctx, &members); err != nil {
		return nil, err
	}

	var responses []models.MemberResponse
	for _, m := range members {
		responses = append(responses, toResponse(m))
	}

	if responses == nil {
		responses = []models.MemberResponse{}
	}

	totalPages := int64(math.Ceil(float64(total) / float64(limit)))

	return &models.PaginatedResponse{
		Data:       responses,
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
	}, nil
}
