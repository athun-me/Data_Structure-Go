package main

import "fmt"

// Define a struct for the key-value pair
type pair struct {
	key   string
	value int
}

// Define a struct for the hash table
type hashTable struct {
	buckets []pair
}

// Hash function to determine the index of the bucket

func hash(s string) int {
	sum := 0
	for _, c := range s {
		sum += int(c)
	}
	return sum % 10
}

// Function to insert a key-value pair into the hash table
func (ht *hashTable) insert(key string, value int) {
	// Get the index of the bucket using the hash function
	index := hash(key)

	fmt.Println(index)

	// Append the key-value pair to the end of the bucket
	ht.buckets = append(ht.buckets, pair{key, value})
}

// Function to retrieve the value associated with a key from the hash table
func (ht *hashTable) get(key string) (int, bool) {
	// Get the index of the bucket using the hash function
	index := hash(key)
	fmt.Println(index)


	// Traverse the bucket to find the value with the key
	for _, p := range ht.buckets {
		if p.key == key {
			return p.value, true
		}
	}

	// If the key is not found, return false
	return 0, false
}

func main() {
	// Initialize the hash table with an empty array
	ht := hashTable{make([]pair, 0)}

	// Insert some key-value pairs into the hash table
	ht.insert("apple", 2)
	ht.insert("banana", 3)
	ht.insert("orange", 5)

	// Retrieve the value associated with a key
	val, found := ht.get("apple")
	if found {
		fmt.Println("apple:", val)
	} else {
		fmt.Println("apple not found")
	}

	// Retrieve the value associated with a key that doesn't exist
	val, found = ht.get("grape")
	if found {
		fmt.Println("grape:", val)
	} else {
		fmt.Println("grape not found")
	}
}
