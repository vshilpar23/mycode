/* Alta3 Research | RZFeeser
   Methods - defining         */

// Go program to illustrate the
// method with struct type receiver
package main
 
import "fmt"
 
// Author structure
type author struct {
    name      string
    branch    string
    books     int
    salary    int
}
 
// Method with a receiver
// of author type
func (a author) show() {

    fmt.Println("Author's Name: ", a.name)
    fmt.Println("Branch Name: ", a.branch)
    fmt.Println("Published articles: ", a.books)
    fmt.Println("Salary: ", a.salary)
}

// increase value of books by 1
func (a *author) bookup() {
    a.books++        // increment name by 1
}

// Main function
func main() {

    // Initializing the values
    // of the author structure
    res := author{
        name:      "RZFeeser",
        branch:    "Pennsylvania",
        books:     14,
        salary:    34000,
    }

 
     // run our new method
    res.bookup()   // increase books by 1

    // Calling the method
    res.show()

    fmt.Println("---------")

    res2 := author{
        name: "Shilpa",
        branch: "Dallas",
        books: 10,
        salary: 34000,
    }
    
    res2.show()
}
