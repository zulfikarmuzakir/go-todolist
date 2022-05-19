package repository

import (
	"go-todolist/config"
	"go-todolist/model"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoRepositoryImpl struct {
	Collection *mongo.Collection
}

func NewTodoRepository(database *mongo.Database) TodoRepository {
	return &TodoRepositoryImpl{
		Collection: database.Collection("todos"),
	}
}

func (repository *TodoRepositoryImpl) Insert(todo model.Todo) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	_, err := repository.Collection.InsertOne(ctx, bson.M{
		"_id":         todo.Id,
		"name":        todo.Name,
		"description": todo.Description,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func (repository *TodoRepositoryImpl) GetAll() (todos []model.Todo) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	cursor, err := repository.Collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err.Error())
	}

	var documents []bson.M
	err = cursor.All(ctx, &documents)
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, document := range documents {
		todos = append(todos, model.Todo{
			Id:          document["_id"].(string),
			Name:        document["name"].(string),
			Description: document["description"].(string),
		})
	}

	return todos
}

func (repository *TodoRepositoryImpl) Delete(param string) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	_, err := repository.Collection.DeleteOne(ctx, bson.M{"_id": param})
	if err != nil {
		log.Fatal(err.Error())
	}
}
