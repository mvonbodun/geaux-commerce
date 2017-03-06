package category

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Repo is a repository of categories
var Repo Repository

func init() {
	var err error
	var cred *mgo.Credential
	// Connect to the mongo database
	Repo, err = newMongoDB("mongodb://localhost", cred)
	if err != nil {
		log.Fatal(err)
	}
}

type mongoDB struct {
	conn *mgo.Session
	c    *mgo.Collection
}

// Ensure mongoDB conforms to the CategoryDatabase Interface
var _ Repository = &mongoDB{}

// newMongoDB creates a new CategoryDatabase backed by a given Mongo server,
// authenticated with given credentials.
func newMongoDB(addr string, cred *mgo.Credential) (Repository, error) {
	conn, err := mgo.Dial(addr)
	if err != nil {
		return nil, fmt.Errorf("mongo: could not dial: %v", err)
	}

	if cred != nil {
		if err := conn.Login(cred); err != nil {
			return nil, err
		}
	}

	return &mongoDB{
		conn: conn,
		c:    conn.DB("geaux-commerce").C("categories"),
	}, nil
}

// Close closes the database
func (db *mongoDB) Close() {
	db.conn.Close()
}

// Find retrieves a category by its ID.
func (db *mongoDB) Find(id int64) (*Category, error) {
	category := &Category{}
	if err := db.c.Find(bson.D{{Name: "id", Value: id}}).One(category); err != nil {
		return nil, ErrNotFound
	}
	return category, nil
}

var maxRand = big.NewInt(1<<63 - 1)

// randomID returns a positive number that fits within an int64.
func randomID() (int64, error) {
	// Get a random number within the range [0, 1<<63-1)
	n, err := rand.Int(rand.Reader, maxRand)
	if err != nil {
		return 0, err
	}
	// Don't assign 0.
	return n.Int64() + 1, nil
}

// Store saves a given category, assigning it a new ID.
func (db *mongoDB) Store(category *Category) (id int64, err error) {
	id, err = randomID()
	if err != nil {
		return 0, fmt.Errorf("mongodb: could not assign a new ID: %v", err)
	}
	category.ID = id
	if err := db.c.Insert(category); err != nil {
		return 0, fmt.Errorf("mongodb: could not add book: %v", err)
	}
	return id, nil
}
