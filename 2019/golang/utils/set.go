package utils

// Set is a map where keys are the elements of the set
type Set[T comparable] map[T]struct{}

// NewSet creates a new set
func NewSet[T comparable]() Set[T] {
	return make(Set[T])
}

// Add adds an element to the set
func (s Set[T]) Add(element T) {
	s[element] = struct{}{}
}

// Exists checks if an element exists in the set
func (s Set[T]) Exists(element T) bool {
	_, exists := s[element]
	return exists
}

// Remove removes an element from the set
func (s Set[T]) Remove(element T) {
	delete(s, element)
}
