package stack

type Stack[T any] []T

func NewStack[T any]() Stack[T] {
	return make(Stack[T], 0)
}

func NewStackWithCapacity[T any](c int) Stack[T] {
	return make(Stack[T], 0, c)
}

// Push adds an element to the top of the stack
//
//goland:noinspection GoMixedReceiverTypes
func (s *Stack[T]) Push(x T) {
	*s = append(*s, x)
}

// PushMany adds elements to the top of the stack
//
//goland:noinspection GoMixedReceiverTypes
func (s *Stack[T]) PushMany(items ...T) {
	*s = append(*s, items...)
}

// Pop removes and returns the top element of the stack
//
//goland:noinspection GoMixedReceiverTypes
func (s *Stack[T]) Pop() T {
	x := s.Peek()
	s.Discard(1)
	return x
}

// Peek returns the top element of the stack without returning it
//
//goland:noinspection GoMixedReceiverTypes
func (s Stack[T]) Peek() T {
	return s[len(s)-1]
}

// Discard discards the top N elements of the stack
//
//goland:noinspection GoMixedReceiverTypes
func (s *Stack[T]) Discard(n int) {
	*s = (*s)[:len(*s)-n]
}

// Size returns the number of elements in the stack
//
//goland:noinspection GoMixedReceiverTypes
func (s Stack[T]) Size() int {
	return len(s)
}

// Get returns the element at the given index
//
//goland:noinspection GoMixedReceiverTypes
func (s Stack[T]) Get(i int) T {
	return s[i]
}

// IsEmpty returns true if the stack is empty
//
//goland:noinspection GoMixedReceiverTypes
func (s Stack[T]) IsEmpty() bool {
	return len(s) == 0
}

// Empty resets the stack
//
//goland:noinspection GoMixedReceiverTypes
func (s *Stack[T]) Empty() {
	*s = (*s)[:0]
}

// ToSlice converts the stack to a slice
//
//goland:noinspection GoMixedReceiverTypes
func (s Stack[T]) ToSlice() []T {
	return append(s[0:0], s...)
}
