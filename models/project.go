package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Project struct {
	Title 			string 							`json:"title" bson:"title" binding:"required"`
	Description string 							`json:"description" bson:"description" binding:"required"`
	ImageUrl 		string 							`json:"image_url" bson:"image_url" binding:"required"`
	Tags 				[]string						`json:"tags" bson:"tags" binding:"required"`
	Stacks			[]string						`json:"stacks" bson:"stacks" binding:"required"`
}

type ResponseProject struct {
	ID primitive.ObjectID `json:"id" bson:"_id"`
	Title 			string 							`json:"title" bson:"title" binding:"required"`
	Description string 							`json:"description" bson:"description" binding:"required"`
	ImageUrl 		string 							`json:"image_url" bson:"image_url" binding:"required"`
	Tags 				[]string						`json:"tags" bson:"tags" binding:"required"`
	Stacks			[]string						`json:"stacks" bson:"stacks" binding:"required"`
}
