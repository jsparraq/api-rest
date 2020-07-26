package repository

import (
	"context"
	"log"

	"../entity"
	"cloud.google.com/go/firestore"
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
	projectId      string = "pragmatic-reviews"
	collectionName string = "post"
)

func (*repo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Fail to create a Firestore Client: %v", err)
		return nil, err
	}
	
	//This line will execute when the function return any object
	defer client.Close()
	_, _, err = client.Collection(collectionName).Add(cxt, map[string]interface{}{
		"ID": post.ID,
		"Title": post.Title,
		"Text": post.Text
	})

	if err != nil {
		log.Fatalf("Fail adding a new post: %v", err)
		return nil, err
	}
}

// FindAll
func (*repo) FindAll() ([]entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Fail to create a Firestore Client: %v", err)
		return nil, err
	}
	
	//This line will execute when the function return any object
	defer client.Close()

	var post []entity.Post
	iterator := client.Collection(collectionName).Documents(ctx)

	for {
		doc, err := iterator.Next()
		if err != nil {
			log.Fatalf("Fail to iterate the list of posts: %v", err)
			return nil, err
		}
		post := entity.Post{
			ID:    doc.Data()["ID"].(int)
			Title: doc.Data()["Title"].(string)
			Text:  doc.Data()["Text"].(string)
		}
		posts = append(posts, post)
	}

	return posts, nil
}
