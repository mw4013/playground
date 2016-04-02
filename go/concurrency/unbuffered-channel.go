
package main

import "fmt"
import "time"

func main() {

    channel := make(chan string)

    go player("A", true, channel)
    go player("B", false, channel)

    var line string
    fmt.Scanln(&line)
}

// each player function runs on its own thread and use the channel to
// exchange a token back and forth
func player(name string, starts bool, channel chan string) {


    iHaveTheToken := starts

    //
    // we go in a loop and exchange the token
    //
    for {

        if iHaveTheToken {

            //
            // put it on the channel
            //

            fmt.Println(name + " sending the token to the channel ...")
            channel <- "."

        } else {

            //
            // wait to get the token
            //

            _ = <- channel
            fmt.Println(name + " got the token from the channel")
            fmt.Println()

            sleep()


        }

        iHaveTheToken = !iHaveTheToken
    }
}

func sleep() {

    time.Sleep(2 * time.Second)
}
