package main

import (
	"github.com/alewkinr/jb-space-oidc-fix/internal/app"
	"log"
)

func main() {
	a := app.MustNewApp()

	if runErr := a.Run(); runErr != nil {
		log.Fatalf("run app error, %v", runErr)
	}
}
