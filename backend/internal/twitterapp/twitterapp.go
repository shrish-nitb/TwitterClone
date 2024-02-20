package twitterapp

import (
	"log"
	"net/http"

	"github.com/codesmith-dev/twitter/internal/gen/api/apiconnect"
	"github.com/codesmith-dev/twitter/internal/gen/ent"
	"github.com/codesmith-dev/twitter/internal/services"
	"github.com/rs/cors"

	_ "github.com/lib/pq"
)

// Run runs the twitter app.
func Run() {
	db, err := ent.Open("postgres", "postgres://postgres:password@localhost:5432/DB02?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	corsHandler := cors.New(cors.Options{
		AllowedHeaders: []string{"Connect-Protocol-Version", "Content-Type"},
	})

	// Access-Control-Allow-Origin : http://localhost:3000
	// Access-Control-Allow-Credentials : true
	// Access-Control-Allow-Methods : GET, POST, OPTIONS
	// Access-Control-Allow-Headers : Origin, Content-Type, Accept

	userAPIPath, userHandler := apiconnect.NewUserServiceHandler(
		services.NewUserServiceHandler(db.User),
	)

	tweetAPIPath, tweetHandler := apiconnect.NewTweetServiceHandler(
		services.NewTweetServiceHandler(db.Tweet),
	)

	http.Handle(userAPIPath, corsHandler.Handler(userHandler))
	http.Handle(tweetAPIPath, corsHandler.Handler(tweetHandler))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err)
	}
}
