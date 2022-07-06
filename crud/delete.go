package crud

import (
	"context"
	"kanolibrary/mongo"
	"kanolibrary/util"
)

func DeleteOne(collection string, query map[string]interface{}) (string, error) {
	db, err := mongo.Connect()
	ctx := context.Background()
	util.ErrorChecker(err)

	_, err = db.Collection(collection).DeleteOne(ctx, query)

	util.ErrorChecker(err)

	return "Remove success!", err
}
