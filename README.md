# Go Silent Map Key Errors

This example demonstrates a potential issue in Go where accessing a non-existent key in a map returns the zero value for the type instead of an error. This can lead to hard-to-debug issues if not handled carefully.

**File:** `bug.go`
The `bug.go` file contains the code that exhibits the silent map key access behavior.  The program accesses a non-existent key, leading to unexpected output. 

**File:** `bugSolution.go`
The `bugSolution.go` shows how to correctly handle potential missing keys, either by checking for the key's existence with the `map[key]` syntax or by using the `ok` idiom.