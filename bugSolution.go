```go
package main

import "fmt"

func main() {
    var m map[string]int
    //Correct way to handle potential missing keys
    value, ok := m["a"]
    if ok {
        fmt.Println("Value of a:", value)
    } else {
        fmt.Println("Key 'a' not found in map")
    }
    //Alternative by checking the length before accessing keys
    if len(m)>0{
        fmt.Println(m["a"])
    } else{
        fmt.Println("Map is empty")
    }
}
```