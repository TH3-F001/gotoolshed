package stack

import (
	"sync"
)

// A simple thread-safe stack implementation that can contain any type including non-comparable types
// The allowance of non-comparable types like slices and maps comes at the expense of being able to compare values
// within the struct itself. Such functionality is out of Stack's scope and can be handled by the caller
type Stack[T any] struct {
	elements []T
	maxSize  int
	mutex    sync.Mutex
}

// Creates a new stack instance that is contrained to maxSize. if maxSize is 0 or less, the stack is unbounded
func New[T any](maxSize int) *Stack[T] {
	if maxSize <= 0 {
		return &Stack[T]{
			elements: make([]T, 0),
			maxSize:  maxSize}
	}

	return &Stack[T]{
		elements: make([]T, 0, maxSize),
		maxSize:  maxSize}
}

// Adds a new value at the end of the stack
func (stk *Stack[T]) Push(val T) bool {
	stk.mutex.Lock()
	defer stk.mutex.Unlock()

	if len(stk.elements) >= stk.maxSize {
		return false
	}

	stk.elements = append(stk.elements, val)
	return true
}

// Removes the last element on the stack and returns it's value (Removes and returns stack[-1])
func (stk *Stack[T]) Pop() (T, bool) {
	stk.mutex.Lock()
	defer stk.mutex.Unlock()

	if len(stk.elements) == 0 {
		var zeroVal T
		return zeroVal, false
	}

	result := stk.elements[len(stk.elements)-1]
	stk.elements = stk.elements[:len(stk.elements)-1]
	return result, true
}

// Returns the value at the end of the stack without removing the last element (Returns stack[-1])
func (stk *Stack[T]) Peek() (T, bool) {
	stk.mutex.Lock()
	defer stk.mutex.Unlock()

	if len(stk.elements) == 0 {
		var zeroVal T
		return zeroVal, false
	}
	return stk.elements[len(stk.elements)-1], true
}

// Returns the current length of the stack
func (stk *Stack[T]) Size() int {
	stk.mutex.Lock()
	defer stk.mutex.Unlock()

	return len(stk.elements)
}

// Returns true if the length of the stack is equal to zero, else false
func (stk *Stack[T]) IsEmpty() bool {
	stk.mutex.Lock()
	defer stk.mutex.Unlock()

	return len(stk.elements) == 0
}

// Returns true if the length of the stack is equal to the stacks max size, else false
func (stk *Stack[T]) IsFull() bool {
	stk.mutex.Lock()
	defer stk.mutex.Unlock()

	return len(stk.elements) == stk.maxSize
}

// Clears the contents of the stack, making it's length zero
func (stk *Stack[T]) Clear() {
	stk.mutex.Lock()
	defer stk.mutex.Unlock()

	stk.elements = stk.elements[:0]
}

// returns a COPY of the stacks elements. This is in order to preserve immutability of the stack.
func (stk *Stack[T]) Elements() []T {
	stk.mutex.Lock()
	defer stk.mutex.Unlock()

	cpy := make([]T, len(stk.elements), stk.maxSize)
	copy(cpy, stk.elements)
	return cpy
}

// Runs a given function against each element in the stack starting from the end and working down to index 0
func (stk *Stack[T]) Traverse(fn func(T)) {
	stk.mutex.Lock()
	defer stk.mutex.Unlock()

	for i := len(stk.elements) - 1; i >= 0; i-- {
		fn(stk.elements[i])
	}
}
