package main

import (
	"log"
	"net/http"

	messaging "my-firebase/messaging"

	"github.com/gorilla/mux"
	"golang.org/x/oauth2"
	"google.golang.org/api/option"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/follow", func(w http.ResponseWriter, r *http.Request) {
		ctx := oauth2.NoContext
		config := &Config{
			ProjectID:        "learn-firebase-d705d",
			ServiceAccountID: "115139974536431150872",
		}

		clientOption := option.WithCredentialsFile("creds.json")

		app, _ := NewApp(ctx, config, clientOption)

		appMessage, _ := app.Messaging(ctx)
		// messaging.SubscribeToTopic(ctx, []string{"test-token"}, "ini topic")

		message := &messaging.Message{
			Data: map[string]string{
				"foo": "bar",
			},
			Token: "eap2c2WVmzaG2bWH81dSC-:APA91bEozG0nnuHzwjAnmUbgW-zKQg8DcYJ6bQc3fF9ynajGntgbYiewvJIi2MZUHAm2WIXoYX_3zomLQtKr45hynK5o_GuC49d_FeH2jR4x8orMYO4zOX5U9Vnq3L0yMoEvqLLexJBl",
		}

		appMessage.Send(ctx, message)
	}).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", r))
}
