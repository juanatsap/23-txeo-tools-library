package models

import (
	"context"
	"database/sql"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Category model adapted for SQLite
type Category struct {
	ID         int            // SQLite uses int for primary keys by default
	Name       string         // The name of the category
	Count      int            // A count of how many times this category is used
	Deleted    bool           // Whether the category is marked as deleted
	Traduction sql.NullString // To handle cases where a translation might be optional or NULL
	Icon       string
}
type Categories []Category

var categories = Categories{
	{Name: "Catchups / Meetings", Icon: "📅"},
	{Name: "Implementation / Configuration tasks", Icon: "💪"},
	{Name: "Emails / Documentation", Icon: "📧"},
	{Name: "Slack / Teams Conversations", Icon: "💬"},
}

// Load
func LoadCategories(collection *mongo.Collection) (Categories, error) {
	var categories Categories
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.Background(), &categories); err != nil {
		return nil, err
	}
	return categories, nil
}

func (c Categories) GetCategories() Categories {
	return categories
}

// Print
func (c Categories) PrintCategories() {
	for _, category := range categories {
		fmt.Println(category.Name)
	}
}
