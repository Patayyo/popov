package store

import "go.mongodb.org/mongo-driver/bson/primitive"

/*type User struct {
	ID      int     `json:"ID"`
	User    string  `json:"User"`
	Cart    []Item  `json:"Cart"`
	CartSum float64 `json:"CartSum"`
	Role    string  `json:"Role"`
}*/

type Item struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name  string             `json:"Name" bson:"name,omitempty"`
	Price float64            `json:"Price" bson:"price,omitempty"`
	Game  string             `json:"Game" bson:"game,omitempty"`
}

type NewItemInput struct {
	Name     string  `json:"Name"`
	Price    float64 `json:"Price"`
	Quantity int     `json:"Quantity"`
}
