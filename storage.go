package main

// Data structure to hold a grocery list
type GroceryList struct {
	storeList map[string][]string
}

// Interface for storage
type Storage interface {
	// Add a store to the list
	addStore(store string)

	// Remove a store from the grocery list
	removeStore(store string)

	// add item to one of the stores
	addItem(store string, item string)

	// remove item from one of the stores
	removeItem(store string, item string)

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

func (s *InMemoryStorage) retrieve() *GroceryList {
	return &GroceryList{s.memstore}
}

func (s *InMemoryStorage) addStore(store string) {
	if _, ok := s.memstore[store]; !ok {
		s.memstore[store] = make([]string, 0)
	}
}

func (s *InMemoryStorage) removeStore(store string) {
	delete(s.memstore, store)
}

func (s *InMemoryStorage) addItem(store string, item string) {
}

func (s *InMemoryStorage) removeItem(store string, item string) {
}
