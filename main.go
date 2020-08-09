package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/andrewlader/go-garminizer/garminizer"
	"github.com/spf13/viper"
)

const asciiArt = `

██████╗  █████╗ ██████╗ ███╗   ███╗██╗███╗   ██╗███████╗███████╗██████╗ 
██╔════╝ ██╔══██╗██╔══██╗████╗ ████║██║████╗  ██║╚══███╔╝██╔════╝██╔══██╗
██║  ███╗███████║██████╔╝██╔████╔██║██║██╔██╗ ██║  ███╔╝ █████╗  ██████╔╝
██║   ██║██╔══██║██╔══██╗██║╚██╔╝██║██║██║╚██╗██║ ███╔╝  ██╔══╝  ██╔══██╗
╚██████╔╝██║  ██║██║  ██║██║ ╚═╝ ██║██║██║ ╚████║███████╗███████╗██║  ██║
 ╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝     ╚═╝╚═╝╚═╝  ╚═══╝╚══════╝╚══════╝╚═╝  ╚═╝
                                                                         
`

var sourceFlag string

func init() {
	flag.StringVar(&sourceFlag, "source", "peloton", "Specifies the source of the data (e.g. Peloton)")

	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("json")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // path to look for the config file in
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
}

func main() {
	flag.Parse()

	fmt.Println(asciiArt)

	fmt.Printf("Garminizer will now attempt to Garminize data from '%s'\n", sourceFlag)
	fmt.Println("==========")
	fmt.Println()

	peloton := &garminizer.Peloton{}
	err := peloton.Login()
	if err != nil {
		log.Printf("error logging into %s: %s", sourceFlag, err)
	}
}
