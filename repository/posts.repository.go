package repository

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/jsparraq/api-rest/entity"
	"google.golang.org/api/iterator"
)

// PostRepository interface
type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type repo struct{}

// NewPostRepository function
func NewPostRepository() PostRepository {
	return &repo{}
}

const (
	projectID      string = "pragmatic-reviews-7fc06"
	collectionName string = "posts"
)

func (*repo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Fail to create a Firestore Client: %v", err)
		return nil, err
	}

	//This line will execute when the function return any object
	defer client.Close()
	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID":    post.ID,
		"Title": post.Title,
		"Text":  post.Text,
	})

	if err != nil {
		log.Fatalf("Fail adding a new post: %v", err)
		return nil, err
	}
	return post, nil
}

// FindAll
func (*repo) FindAll() ([]entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Fail to create a Firestore Client: %v", err)
		return nil, err
	}

	//This line will execute when the function return any object
	defer client.Close()

	var posts []entity.Post
	iter := client.Collection("posts").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Fail to iterate the list of posts: %v", err)
			return nil, err
		}
		post := entity.Post{
			ID:    doc.Data()["ID"].(int64),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}
		posts = append(posts, post)
	}

	return posts, nil
}
