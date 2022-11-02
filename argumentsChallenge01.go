/* Alta3 Research
   counting  arguments from the command line*/

   package main

import (
    "fmt"
    "os"
)

func main() {
// os.Args provides access to raw CLI arguments
    argsLength := len(os.Args[1:])         // includes program name
    fmt.Println("Number of arguments %d: ",argsLength)
    
    for i,a := range os.Args[1:] {
        fmt.Printf("Arg %d is %s\n",i+1,a)
        }

}

