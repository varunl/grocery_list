package main

// Data structure to hold a grocery list
type GroceryList struct {
	storeList map[string][]string
}

type Storage interface {
	// Store the contents of the grocery list in the storage.
	store(g GroceryList)

	// Retrieve the contents of the grocery list.
	retrieve() *GroceryList
}

// In memory version of the storage
type InMemoryStorage struct {
	memstore map[string][]string
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		memstore: make(map[string][]string),
	}
}

func (s *InMemoryStorage) store(g GroceryList) {
}

func (s *InMemoryStorage) retrieve() *GroceryList {
	return nil
}
