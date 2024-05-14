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

	listMonitoredComputersOutput, err := v.ListMonitoredComputers(ctx)
	if err != nil {
		log.Fatal("failed to list monitored computers with vanta sdk: %v", err)
	}

	for _, monitoredComputer := range listMonitoredComputersOutput.Results.Data {
		serial := "N/A"
		if monitoredComputer.SerialNumber != nil {
			serial = *monitoredComputer.SerialNumber
		}
		fmt.Printf(
			"%s: %s (OS=%s) ScreenLock=%s DiskEncryption=%s PasswordManager=%s AntivirusInstallation=%s\n",
			monitoredComputer.Owner.EmailAddress,
			serial,
			monitoredComputer.OperatingSystem.Type,
			monitoredComputer.ScreenLock.Outcome,
			monitoredComputer.DiskEncryption.Outcome,
			monitoredComputer.PasswordManager.Outcome,
			monitoredComputer.AntivirusInstallation.Outcome,
		)
	}
}
