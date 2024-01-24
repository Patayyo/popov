package store_mongo

import (
	"context"
	"log"

	"github.com/gorepos/usercartv2/internal/store"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoStore struct {
	Store *mongo.Client
}

var db *mongo.Client

const (
	Database         = "usercart"
	ItemsCollection  = "items"
	UsersCollection  = "users"
	ConnectionString = "mongodb://root:example@mongo:27017/"
)

// NewStore creates child for the Store interface
/*
There are two functions NewStore and CreateConnection, NewStore()
creates child of the interface for Store, while CreateConnection creates low level store client
*/
func NewStore() (*MongoStore, error) {
	db, err := CreateConnection()
	if err != nil {
		return nil, err
	}

	ms := &MongoStore{Store: db}
	return ms, nil
}

// CreateConnection function creates low level connection to the mongo database
func CreateConnection() (*mongo.Client, error) {
	var err error
	db, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(ConnectionString))
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to mongodb!")
	return db, nil
}

func CloseConnection() {
	err := db.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Disconnected from mongodb!")
}

func (s *MongoStore) GetItems() ([]store.Item, error) {
	var items []store.Item
	collection := db.Database(Database).Collection(ItemsCollection)
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var item store.Item
		if err := cursor.Decode(&item); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	if cursor.Err() != nil {
		return nil, cursor.Err()
	}
	return items, nil

}

func (s *MongoStore) AddItem(item store.Item) error {
	collection := s.Store.Database(Database).Collection(ItemsCollection)

	/*existingItem, _ := s.GetItemByName(item.Name)
	if existingItem != nil {

		existingItem.Quantity += item.Quantity
		_, err := collection.UpdateOne(context.Background(), bson.M{"_id": existingItem.ID}, bson.M{"$set": existingItem})
		return err
	}*/

	_, err := collection.InsertOne(context.Background(), item)
	if err != nil {
		return err
	}
	return nil
}

func (s *MongoStore) GetItemByName(name string) (*store.Item, error) {
	var item store.Item
	collection := s.Store.Database(Database).Collection(ItemsCollection)

	err := collection.FindOne(context.Background(), bson.M{"name": name}).Decode(&item)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &item, nil
}

func (s *MongoStore) GetItemByID(id string) (*store.Item, error) {
	collection := s.Store.Database(Database).Collection(ItemsCollection)
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var item store.Item
	err = collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&item)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &item, nil
}

func (s *MongoStore) UpdateItem(id string, updatedItem store.Item) error {
	collection := s.Store.Database(Database).Collection(ItemsCollection)
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{
		"name":  updatedItem.Name,
		"price": updatedItem.Price,
	}}

	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (s *MongoStore) DeleteItem(id string) error {
	collection := s.Store.Database(Database).Collection(ItemsCollection)
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(context.Background(), bson.M{"_id": objectID})
	if err != nil {
		return err
	}

	return nil
}

/*func (s *MongoStore) GetUsers() ([]store.User, error) {
	collection := s.Store.Database(Database).Collection(UsersCollection)

	var users []store.User
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var user store.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if cursor.Err() != nil {
		return nil, cursor.Err()
	}
	return users, nil
}

func (s *MongoStore) GetUserByID(userID int) (*store.User, error) {
	collection := s.Store.Database(Database).Collection(UsersCollection)

	var user store.User
	err := collection.FindOne(context.Background(), bson.M{"id": userID}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		log.Println("Error fetching user by ID:", err)
		return nil, err
	}

	return &user, nil
}

func (s *MongoStore) AddUser(user store.User) error {
	collection := s.Store.Database(Database).Collection(UsersCollection)

	_, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}
	return nil
}

func (s *MongoStore) AddItemToUserCart(userID int, item store.Item) error {
	collection := s.Store.Database(Database).Collection(UsersCollection)

	userFilter := bson.M{"id": userID}
	update := bson.M{"$push": bson.M{"cart": item}}

	_, err := collection.UpdateOne(context.Background(), userFilter, update)
	if err != nil {
		return err
	}

	return nil
}

func (s *MongoStore) DeleteUser(userID int) error {
	collection := s.Store.Database(Database).Collection(UsersCollection)

	_, err := collection.DeleteOne(context.Background(), bson.M{"ID": userID})
	if err != nil {
		log.Printf("Error deleting user: %v", err)
		return err
	}

	log.Printf("User deleted successfully")
	return nil
}

func (s *MongoStore) AddUserRole(userID int, role string) error {
	collection := s.Store.Database(Database).Collection(UsersCollection)
	filter := bson.M{"id": userID}
	update := bson.M{"$set": bson.M{"role": role}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Println("Error updating user role:", err)
		return err
	}
	if result.ModifiedCount == 0 {
		log.Println("User not found or role not updated")
	} else {
		log.Printf("User with ID %d: role updated to %s", userID, role)
	}
	return nil
}*/
