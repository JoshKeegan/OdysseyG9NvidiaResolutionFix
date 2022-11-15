package main

import (
	"fmt"
	"os"
)

var resolutions = []string{
	"2560x1439x8,16,32,64=1F;",
	"3440x1440x8,16,32,64=1F;",
}

func main() {
	fmt.Println("Odyssey G9 Nvidia Resolution Fix")
	fmt.Println("--------------------------------")

	for _, res := range resolutions {
		err := addResolution(res)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: unable to add resolution %v. Error %v\n", res, err)
			os.Exit(1)
		}
		fmt.Printf("Added resolution %v\n", res)
	}

	fmt.Println("Restart required. Close all running applications and press enter to restart . . .")
	fmt.Scanln()
	err := reboot()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error running restart: %v", err)
		os.Exit(2)
	}
}
