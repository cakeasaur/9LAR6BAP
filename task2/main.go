package main

import (
    "encoding/json"
    "os"
)

func main() {
    var numbers []int
    json.NewDecoder(os.Stdin).Decode(&numbers)
    
    sum := 0
    for _, n := range numbers {
        sum += n
    }
    
    json.NewEncoder(os.Stdout).Encode(map[string]int{"sum": sum})
}