package main

import (
	"flag"
	"fmt"
)

const asciiArt = `

██████╗  █████╗ ██████╗ ███╗   ███╗██╗███╗   ██╗███████╗███████╗██████╗ 
██╔════╝ ██╔══██╗██╔══██╗████╗ ████║██║████╗  ██║╚══███╔╝██╔════╝██╔══██╗
██║  ███╗███████║██████╔╝██╔████╔██║██║██╔██╗ ██║  ███╔╝ █████╗  ██████╔╝
██║   ██║██╔══██║██╔══██╗██║╚██╔╝██║██║██║╚██╗██║ ███╔╝  ██╔══╝  ██╔══██╗
╚██████╔╝██║  ██║██║  ██║██║ ╚═╝ ██║██║██║ ╚████║███████╗███████╗██║  ██║
 ╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝     ╚═╝╚═╝╚═╝  ╚═══╝╚══════╝╚══════╝╚═╝  ╚═╝
                                                                         
`

var directoryFlag string
var sourceFlag string
var activityFlag string

func init() {
	flag.StringVar(&directoryFlag, "dir", ".", "Specifies the location of the file(s)")
	flag.StringVar(&sourceFlag, "source", "strava", "Specifies the source application (e.g., Strava)")
	flag.StringVar(&activityFlag, "activity", "peloton", "Specifies the activity (e.g. Peloton)")
}
func main() {
	flag.Parse()

	fmt.Println(asciiArt)

	fmt.Printf("Garminizer will now attempt to Garminize the '%s' files located at '%s' as activity '%s'\n",
		sourceFlag, directoryFlag, activityFlag)
	fmt.Println("==========")
}
