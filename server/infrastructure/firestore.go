package infrastructure

import (
	"context"
	"log"

	model "server/model"

	"cloud.google.com/go/firestore"
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

	_, err = client.Collection("users").Doc(user.Id).Set(ctx, map[string]interface{}{
		"name":  user.Name,
		"email": user.Email,
		"icon":  user.Icon,
	})
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
		return err
	}

	defer client.Close()

	return err
}

func SaveMatch(match model.Match) error {
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

	_, _, err = client.Collection("matches").Add(ctx, map[string]interface{}{
		"team1_id":   match.Team1_id,
		"team2_id":   match.Team2_id,
		"start_time": match.Start_time,
	})
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
		return err
	}

	defer client.Close()

	return err
}

func ConnectToFirestore() (*firestore.Client, context.Context) {
	ctx := context.Background()
	sa := option.WithCredentialsFile("./serviceAccountKey.json")
	app, _ := firebase.NewApp(ctx, nil, sa)
	// if err != nil {
	// 	log.Fatalln(err)
	// 	return err
	// }

	client, _ := app.Firestore(ctx)
	// if err != nil {
	// 	log.Fatalln(err)
	// 	return err
	// }

	return client, ctx
}

func GetMatches() *firestore.DocumentIterator {
	ctx := context.Background()
	sa := option.WithCredentialsFile("./serviceAccountKey.json")
	app, _ := firebase.NewApp(ctx, nil, sa)

	client, _ := app.Firestore(ctx)

	matches := client.Collection("matches").Documents(ctx)
	defer client.Close()

	return matches
}
