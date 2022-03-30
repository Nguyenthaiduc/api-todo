package models

import "time"

type Todo struct {
	ID          *string   `json:"id,omitempty" bson:"_id,omitempty"`
	Title       *string   `json:"title"`
	Description *string   `json:"description"`
	Completed   *bool     `json:"completed"`
	CreateAt    time.Time `json:"createAt"`
	UpdateAt time.Time `json:"updateAt"`
}
