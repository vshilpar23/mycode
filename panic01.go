/* Alta3 Research | RZFeeser
   Calling the panic() function directly */
   
package main

import (
    "fmt"
)

func main() {

    // panic produces a quick exit
    panic("We have a problem")

    fmt.Println("You will not even see this line. The panic creates a fast fail.")
}

