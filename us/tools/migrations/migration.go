package main

import (
	"context"
	"log"

	"github.com/gorepos/usercartv2/internal/store"
	"github.com/gorepos/usercartv2/internal/store/store_mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	runMigration()
}

func runMigration() {
	log.Println("Run migration...")
	db, err := store_mongo.CreateConnection()
	if err != nil {
		panic(err)
	}
	defer store_mongo.CloseConnection()

	err = db.Database(store_mongo.Database).Drop(context.Background())
	if err != nil {
		panic(err)
	}

	collection := db.Database(store_mongo.Database).Collection(store_mongo.ItemsCollection)

	documents := []interface{}{
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "Azuki",
			Price: 6.18,
			Game:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "Chocola",
			Price: 5.86,
			Game:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "Cinnamon",
			Price: 7.62,
			Game:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "Coconut",
			Price: 5.15,
			Game:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "Maple",
			Price: 4.51,
			Game:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "Minaduki Family",
			Price: 5.15,
			Game:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "Shigure",
			Price: 4.43,
			Game:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "Vanilla",
			Price: 6.83,
			Game:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "Cuphead",
			Price: 2.43,
			Game:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "Mugman",
			Price: 2.10,
			Game:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "Elder Kettle",
			Price: 2.24,
			Game:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "Legendary Chalice",
			Price: 2.31,
			Game:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "The Devil",
			Price: 2.57,
			Game:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "King Dice",
			Price: 2.60,
			Game:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "Hilda Berg",
			Price: 2.91,
			Game:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "Cala Maria",
			Price: 2.24,
			Game:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "AAA",
			Price: 2.77,
			Game:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "Large Arachnoid",
			Price: 1.77,
			Game:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "Biomechanoid",
			Price: 2.27,
			Game:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "Gnaar",
			Price: 2.16,
			Game:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "Khnum",
			Price: 2.16,
			Game:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "Octanian",
			Price: 2.26,
			Game:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "Processed",
			Price: 2.10,
			Game:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "Draconian Pyromaniac",
			Price: 2.23,
			Game:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "Biker Sam",
			Price: 2.50,
			Game:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "Scrapjack",
			Price: 2.22,
			Game:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "Ugh Zan",
			Price: 2.30,
			Game:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "Sirian Werebull",
			Price: 2.27,
			Game:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "Beheaded Kamikaze",
			Price: 2.24,
			Game:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "Kleer Skeleton",
			Price: 2.13,
			Game:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "Lord Achriman",
			Price: 2.30,
			Game:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "Balkan",
			Price: 2,
			Game:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "Anarchist",
			Price: 2,
			Game:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "IDF",
			Price: 2,
			Game:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "FBI",
			Price: 1.72,
			Game:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "SWAT",
			Price: 1.73,
			Game:  "20",
		},
	}

	_, err = collection.InsertMany(context.Background(), documents)
	if err != nil {
		panic(err)
	}

	/*collectionUsers := db.Database(store_mongo.Database).Collection(store_mongo.UsersCollection)

	users := []interface{}{
		store.User{
			ID:      1,
			User:    "user1",
			Cart:    []store.Item{},
			CartSum: 0,
		},
		store.User{
			ID:      2,
			User:    "user2",
			Cart:    []store.Item{},
			CartSum: 0,
		},
	}

	_, err = collectionUsers.InsertMany(context.Background(), users)
	if err != nil {
		panic(err)
	}*/

	log.Println("Migration completed successfully.")
}
