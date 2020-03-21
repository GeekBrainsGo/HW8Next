package models

import (
	"context"
	"html/template"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Post stands for post object.
type Post struct {
	Mongo   `bson:"inline"`
	Title   string        `bson:"title" json:"title"`
	Author  string        `bson:"author" json:"author"`
	Content template.HTML `bson:"content" json:"content"`
}

const collection = "posts"

// Insert post to database.
func (p *Post) Insert(ctx context.Context, db *mongo.Database) error { // 2
	col := db.Collection(collection)
	res, err := col.InsertOne(ctx, p)
	if err != nil {
		return err
	}
	p.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

// Update updates post in database.
func (p *Post) Update(ctx context.Context, db *mongo.Database) (*Post, error) { // 1
	col := db.Collection(collection)
	_, err := col.ReplaceOne(ctx, bson.M{"_id": p.ID}, p)
	return p, err
}

// Delete deletes post from database.
func (p *Post) Delete(ctx context.Context, db *mongo.Database) (*Post, error) { // 1
	col := db.Collection(collection)
	_, err := col.DeleteOne(ctx, bson.M{"_id": p.ID})
	return p, err
}

// AllPosts return all posts from database.
func AllPosts(ctx context.Context, db *mongo.Database) ([]Post, error) { // 3
	col := db.Collection(collection)
	options := options.Find()
	options.SetLimit(20)
	options.SetSort(bson.D{{"_id", -1}})

	cur, err := col.Find(ctx, bson.M{}, options)
	if err != nil {
		return nil, err
	}

	posts := []Post{}
	if err := cur.All(ctx, &posts); err != nil {
		return nil, err
	}
	return posts, nil
}

// GetPost gets post by id.
func GetPost(ctx context.Context, db *mongo.Database, id string) (*Post, error) { // 3
	col := db.Collection(collection)
	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	res := col.FindOne(ctx, bson.M{"_id": docID})

	p := Post{}
	if err := res.Decode(&p); err != nil {
		return nil, err
	}
	return &p, nil
}
