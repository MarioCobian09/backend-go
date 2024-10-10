package config

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

func InitializeFirebaseApp() *firebase.App {
	opt := option.WithCredentialsFile("ruta")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Error al inicializar Firebase: %v", err)
	}
	return app
}

func getAuthClient(app *firebase.App) *auth.Client {
	client, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("Error al obtener el cliente: %v", err)
	}
	return client
}
