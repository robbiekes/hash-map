package main

import (
	"fmt"
	"myhashtable/hashtable"
)

func main() {
	hash := hashtable.NewHashTable()
	hash.Set("a", "apple")
	pair := hash.Get("a")
	fmt.Println(pair)                     // shows key-value pair
	fmt.Println(hashtable.GetValue(pair)) // shows just a value
}
