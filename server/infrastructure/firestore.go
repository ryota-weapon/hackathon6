package infrastructure

import (
	"context"
	"log"

	model "server/model"

	firebase "firebase.google.com/go"
	// "google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func RegisterUser(user model.User) error {
	ctx := context.Background()
	sa := option.WithCredentialsFile("./serviceAccountKey.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	_, _, err = client.Collection("users").Add(ctx, map[string]interface{}{
		"first":   user.Name,
		"last":    "Lovelace",
		"born":    1815,
		"user_id": user.Auth_id,
	})
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
		return err
	}

	defer client.Close()

	return err
}
