/* Alta3 Research | RZFeeser
   Simple Sorting */
   
   package main

   import (
	   "fmt"
	   "sort"
   )
   
   func main() {

    // create a slice of strings
    strs := []string{"r", "z", "f", "e", "e", "s", "e", "r"}

	fmt.Println("Before sort:",strs[3])
    
    // sorting is performed "in place" it doesn't RETURN anything
    sort.Strings(strs)
     // the slice is now sorted
	 fmt.Println("Strings:", strs)

	 fmt.Println("After sort:",strs[3])

	 // create a slice of integers
	 ints := []int{2, 4, 6, 0, 1}

	  // sorted low to high
	  sort.Ints(ints)
	  // display the results
    // results of sorted integers
    fmt.Println("Ints:   ", ints)
   }