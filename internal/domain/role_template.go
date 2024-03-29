package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type RoleTemplate struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	Avatar      string             `bson:"avatar"`
	Summary     string             `bson:"summary"`
	Description string             `bson:"description"`
	Example     string             `bson:"example"`
	StarCount   int64              `bson:"starCount"`
	CreatedBy   string             `bson:"createdBy"`
	CreatedTime int64              `bson:"createdTime"`
}

type RoleTemplateList []RoleTemplate

type ListCondition struct {
	Id       string
	Summary  string
	PageNum  int64
	PageSize int64
}
