/* Alta3 
 Map several programming languages to their file extension, 
 for example, the key Python should map to the value .py. 
 The key Golang should map to .go, the key Java should map to .java, 
 the key Ansible should map to .yml, and the key C++ should map to .cpp. 
 Once the map is constructed, display it on the screen.
Then remove the language C++, and add Julia and the key .jl. Display the new map after the change.*/

   package main

   import (
	   "fmt"
   )
   
   func main() {
	   
	   // make assigns a memory address (will not return nil)
	   
	     
	   programlang := map[string]string{
		   "Python": ".py",
		   "Golang": ".go",
		   "Java": ".java",
		   "C++": ".cpp",
		   "Julia": ".jl",
		   "Scala": ".scala",
	   }
	   // Returns values: map[abc:123 def:345 ghi:567 jkl:897]
	   fmt.Println("programlang:", programlang)
	   delete(programlang,"C++")
	   delete(programlang,"Julia")
programlang["Ansible"] = ".yml"
	   fmt.Println("new programlang:", programlang)
	   
   }
