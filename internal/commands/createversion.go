package commands

import (
	"fmt"
	"time"

	"github.com/urfave/cli"
)

func createVersion(c *cli.Context) {
	fmt.Printf("Creating version")
	time.Sleep(time.Second * 1)
	fmt.Printf(".")
	time.Sleep(time.Second * 1)
	fmt.Printf(".")
	time.Sleep(time.Second * 1)
	fmt.Printf(".")
	time.Sleep(time.Second * 1)
	fmt.Println("\nVersion created!")
	time.Sleep(time.Second * 1)

}
