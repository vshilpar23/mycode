/* RZFeeser | Alta3 Research
   Slices relationship to arrays     */

package main

import "fmt"

func main() {
    array := [5]int{1, 2, 3, 4, 5}   // this array will back the slice
    slice := array[:]                // [:] means "from start to the end of"

    //Modifying the slice (which is a pointer to the array)
    slice[1] = 7
    fmt.Println("Modifying Slice")
    fmt.Println("Array =", array)
    fmt.Println("Slice =", slice)

    //Modifying the array. Would reflect back in slice too
    array[1] = 2
    fmt.Println("Modifying Underlying Array")
    fmt.Println("Array =", array)
    fmt.Println("Slice =", slice)

    mySlice5 := []string{"a", "b", "c", "d"}

        fmt.Println("Contents of mySlice5:", mySlice5)
        fmt.Println("Length of mySlice5:",   len(mySlice5))    // returns length
        fmt.Println("Capacity of mySlice5:", cap(mySlice5))    // returns capacity

        mySlice6 := make([]bool, 2, 4)    // the length (window) is set to 2
                      // the capacity (size of array backing Slice) is 4

        fmt.Println("Contents of mySlice6:", mySlice6)
        fmt.Println("Length of mySlice6:",   len(mySlice6))    // returns length
        fmt.Println("Capacity of mySlice6:", cap(mySlice6))    // returns capacity



        // general use cases
  // append an element to a slice
        a := []int{1, 2}
        a = append(a, 3) // a == [1 2 3]
        fmt.Println(a)   // [123]

        // Concatenate two slices
        b := []int{1, 2}
        c := []int{11, 22}
        b = append(b, c...) // b == [1 2 11 22]
        fmt.Println(b)
        
        // The result does not depend on whether the arguments overlap
        // so we can concatenate a slice with itself
        z := []int{1, 2}
        z = append(z, z...) // z == [1 2 1 2]
        fmt.Println(z)
        

                // Copy from one slice to another
        var s = make([]int, 3)
        n := copy(s, []int{0, 1, 2, 3}) // n == 3, s == []int{0, 1, 2}
        fmt.Println(s)                  // [0 1 2]
        fmt.Println(n)                  // 3

        // Copy from a slice to itself
        d := []int{0, 1, 2}
        g := copy(d, d[1:])     // g == 2, d == []int{1, 2, 2}
        fmt.Println(d)          // [1 2 2]
        fmt.Println(g)          // 2


}

