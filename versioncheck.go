package main

import (
	"fmt"
	"github.com/dasJ/versioncheck/config"
	"github.com/dasJ/versioncheck/moduleRunner"
	"github.com/dasJ/versioncheck/notificator"
	"github.com/dasJ/versioncheck/verdb"
)

func main() {
	// Configuration
	cfg, err := config.Read()
	if err != nil {
		panic(err)
	}
	// Version database
	verDb, err := verdb.Read(cfg)
	if err != nil {
		panic(err)
	}
	defer verDb.Close()
	// Run modules
	res := moduleRunner.RunModules(cfg, verDb)
	// Update version database
	err = verDb.Write()
	if err != nil {
		panic(err)
	}
	// Notify
	notificator.Notify(cfg, res)
	// Output
	if len(res.Changed) != 0 {
		fmt.Println("Changed items:")
		for _, item := range res.Changed {
			fmt.Printf("	- %s\n", item.Name)
		}
		fmt.Println()
	}
	if len(res.Failed) != 0 {
		fmt.Println("Failed items:")
		for _, item := range res.Failed {
			fmt.Printf("	- %s\n", item)
		}
		fmt.Println()
	}
	fmt.Printf("Amount of changed versions: %d\n", len(res.Changed))
	fmt.Printf("Amount of failed versions:  %d\n", len(res.Failed))
}
