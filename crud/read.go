package crud

import (
	"context"
	"encoding/json"
	"fmt"
	"kanolibrary/models"
	"kanolibrary/mongo"
	"kanolibrary/util"

	"go.mongodb.org/mongo-driver/bson"
)

type books struct {
	Title, Author, Synopsis string
}

func FindAll(collection string) ([]byte, error) {

	ctx := context.Background()
	db, err := mongo.Connect()
	util.ErrorChecker(err)

	csr, err := db.Collection(collection).Find(ctx, bson.M{})
	util.ErrorChecker(err)
	defer csr.Close(ctx)

	result := make([]models.Books, 0)
	for csr.Next(ctx) {
		var row models.Books
		err := csr.Decode(&row)
		util.ErrorChecker(err)

		result = append(result, row)
	}
	data, err := json.Marshal(result)
	util.ErrorChecker(err)
	return data, err
}

func FindOne(collection string, query map[string]interface{}) ([]byte, error) {

	ctx := context.Background()
	db, err := mongo.Connect()
	util.ErrorChecker(err)

	csr, err := db.Collection(collection).Find(ctx, query)
	util.ErrorChecker(err)
	defer csr.Close(ctx)

	result := make([]models.Books, 0)
	for csr.Next(ctx) {
		var row models.Books
		err := csr.Decode(&row)
		util.ErrorChecker(err)

		result = append(result, row)
	}
	data, err := json.Marshal(result)
	util.ErrorChecker(err)
	return data, err

	var i = 0
	for {
		fmt.Println("Title      :", result[i].Title)
		fmt.Println("Author     :", result[i].Author)
		fmt.Println("Synopsis   :", result[i].Synopsis)
		fmt.Println("")
		i++
	}
}

// func Find(collection string, data interface{}) {
// 	ctx := context.Background()
// 	db, err := mongo.Connect()
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	csr, err := db.Collection("books").Find(ctx, bson.D{{}})
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}
// 	defer csr.Close(ctx)

// 	result := make([]books, 0)
// 	for csr.Next(ctx) {
// 		var row books
// 		err := csr.Decode(&row)
// 		if err != nil {
// 			log.Fatal(err.Error())
// 		}

// 		result = append(result, row)
// 	}

// 	var i = 0

// 	for {
// 		fmt.Println("Name  :", result[i].Name)
// 		fmt.Println("Author :", result[i].Author)
// 		i++
// 	}

// if len(result) > 0 {
// 	fmt.Println("Name  :", result[0].Name)
// 	fmt.Println("Author :", result[0].Author)
// }
// }
