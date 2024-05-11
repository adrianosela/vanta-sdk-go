package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/adrianosela/vanta-sdk-go"
)

func main() {
	ctx := context.Background()

	v, err := vanta.New(
		ctx,
		vanta.WithOAuthCredentials(os.Getenv("VANTA_OAUTH_CLIENT_ID"), os.Getenv("VANTA_OAUTH_CLIENT_SECRET")),
		vanta.WithScopes(vanta.ScopeAllRead, vanta.ScopeAllWrite),
	)
	if err != nil {
		log.Fatalf("failed to initialize vanta sdk: %v", err)
	}

	listPeopleOutput, err := v.ListPeople(ctx)
	if err != nil {
		log.Fatal("failed to list people with vanta sdk: %v", err)
	}

	for _, person := range listPeopleOutput.Results.Data {
		fmt.Println(person.EmailAddress)
	}
}
