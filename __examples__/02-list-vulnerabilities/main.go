package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/adrianosela/vanta-sdk-go"
)

func main() {
	ctx := context.Background()

	searchString := ""
	if len(os.Args) > 1 {
		searchString = strings.Join(os.Args[1:], " ")
	}

	v, err := vanta.New(
		ctx,
		vanta.WithOAuthCredentials(os.Getenv("VANTA_OAUTH_CLIENT_ID"), os.Getenv("VANTA_OAUTH_CLIENT_SECRET")),
		vanta.WithScopes(vanta.ScopeAllRead, vanta.ScopeAllWrite),
	)
	if err != nil {
		log.Fatalf("failed to initialize vanta sdk: %v", err)
	}

	opts := []vanta.ListVulnerabilitiesOption{
		vanta.WithPageSize(100),
	}
	if searchString != "" {
		opts = append(opts, vanta.WithSearchQuery(searchString))
	}

	listVulnerabilitiesOutput, err := v.ListVulnerabilities(ctx, opts...)
	if err != nil {
		log.Fatalf("failed to list vulnerabilities with vanta sdk: %v", err)
	}

	for _, vuln := range listVulnerabilitiesOutput.Results.Data {
		description := vuln.Description
		if len(description) > 100 {
			description = fmt.Sprintf("%s...", description[:100])
		}
		fmt.Printf("%s %s %s\n", vuln.ID, vuln.Name, description)
	}
}
