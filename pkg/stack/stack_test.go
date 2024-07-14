package stack_test

import (
	"reflect"
	"testing"
)

import . "github.com/GabiBizdoc/collections/pkg/stack"

type stackItem int

func newStack(items ...stackItem) Stack[stackItem] {
	s := NewStack[stackItem]()
	s.PushMany(items...)
	return s
}

type args struct {
	x stackItem
}

type testCase[T stackItem] struct {
	name      string
	s         Stack[T]
	args      args
	wantItem  T
	wantStack Stack[T]
}

func TestStack_Discard(t *testing.T) {
	t.Parallel()

	type args struct {
		x int
	}

	type testCase[T stackItem] struct {
		name      string
		s         Stack[T]
		args      args
		wantItem  T
		wantStack Stack[T]
	}

	tests := []testCase[stackItem]{
		{
			name:      "Discard top 0 elements",
			s:         newStack(1, 2, 3, 4),
			args:      args{x: 0},
			wantStack: newStack(1, 2, 3, 4),
		},
		{
			name:      "Discard top 2 elements",
			s:         newStack(1, 2, 3, 4),
			args:      args{x: 2},
			wantStack: newStack(1, 2),
		},
		{
			name:      "Discard all elements",
			s:         newStack(1, 2, 3, 4),
			args:      args{x: 4},
			wantStack: newStack(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Discard(tt.args.x)
			if !reflect.DeepEqual(tt.s, tt.wantStack) {
				t.Errorf("Stack after Discard() = %v, want %v", tt.s, tt.wantStack)
			}
		})
	}
}

func TestStack_Empty(t *testing.T) {
	t.Parallel()

	tests := []testCase[stackItem]{
		{
			name:      "Empty a non-empty stack",
			s:         newStack(1, 2, 3),
			wantStack: newStack(),
		},
		{
			name:      "Empty an already empty stack",
			s:         newStack(),
			wantStack: newStack(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Empty()
			if !reflect.DeepEqual(tt.s, tt.wantStack) {
				t.Errorf("Stack after Empty() = %v, want %v", tt.s, tt.wantStack)
			}
		})
	}
}

func TestStack_Get(t *testing.T) {
	t.Parallel()

	type args struct {
		x int
	}
	type testCase[T stackItem] struct {
		name      string
		s         Stack[T]
		args      args
		want      T
		wantStack Stack[T]
	}

	tests := []testCase[stackItem]{
		{
			name:      "Get element at index 1",
			s:         newStack(1, 2, 3),
			args:      args{x: 1},
			want:      2,
			wantStack: newStack(1, 2, 3),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotItem := tt.s.Get(tt.args.x)
			if !reflect.DeepEqual(gotItem, tt.want) {
				t.Errorf("Get() = %v, want %v", gotItem, tt.want)
			}
			if !reflect.DeepEqual(tt.s, tt.wantStack) {
				t.Errorf("Stack after Get() = %v, want %v", tt.s, tt.wantStack)
			}
		})
	}
}

func TestStack_IsEmpty(t *testing.T) {
	t.Parallel()

	type testCase[T stackItem] struct {
		name string
		s    Stack[T]
		args args
		want bool
	}

	tests := []testCase[stackItem]{
		{
			name: "Non-empty stack",
			s:    newStack(1, 2, 3),
			want: false,
		},
		{
			name: "Empty stack",
			s:    newStack(),
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Peek(t *testing.T) {
	t.Parallel()

	tests := []testCase[stackItem]{
		{
			name:      "Peek stack",
			s:         newStack(1, 2, 3),
			wantItem:  3,
			wantStack: newStack(1, 2, 3),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotItem := tt.s.Peek()
			if !reflect.DeepEqual(gotItem, tt.wantItem) {
				t.Errorf("Peek() = %v, want %v", gotItem, tt.wantItem)
			}
			if !reflect.DeepEqual(tt.s, tt.wantStack) {
				t.Errorf("Stack after Peek() = %v, want %v", tt.s, tt.wantStack)
			}
		})
	}
}

func TestStack_Pop(t *testing.T) {
	t.Parallel()

	tests := []testCase[stackItem]{
		{
			name:      "Pop stack",
			s:         newStack(1, 2, 3),
			wantItem:  3,
			wantStack: newStack(1, 2),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotItem := tt.s.Pop()
			if !reflect.DeepEqual(gotItem, tt.wantItem) {
				t.Errorf("Pop() = %v, want %v", gotItem, tt.wantItem)
			}
			if !reflect.DeepEqual(tt.s, tt.wantStack) {
				t.Errorf("Stack after Pop() = %v, want %v", tt.s, tt.wantStack)
			}
		})
	}
}

func TestStack_Push(t *testing.T) {
	t.Parallel()

	tests := []testCase[stackItem]{
		{
			name:      "Push stack",
			s:         newStack(1, 2, 3),
			args:      args{x: 5},
			wantStack: newStack(1, 2, 3, 5),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Push(tt.args.x)
			if !reflect.DeepEqual(tt.s, tt.wantStack) {
				t.Errorf("Stack after Push() = %v, want %v", tt.s, tt.wantStack)
			}
		})
	}
}

func TestStack_Size(t *testing.T) {
	t.Parallel()

	type testCase[T stackItem] struct {
		name      string
		s         Stack[T]
		args      args
		want      int
		wantStack Stack[T]
	}

	tests := []testCase[stackItem]{
		{
			name:      "Size of stack",
			s:         newStack(1, 2, 2, 1),
			want:      4,
			wantStack: newStack(1, 2, 2, 1),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotItem := tt.s.Size()
			if !reflect.DeepEqual(gotItem, tt.want) {
				t.Errorf("Size() = %v, want %v", gotItem, tt.want)
			}
			if !reflect.DeepEqual(tt.s, tt.wantStack) {
				t.Errorf("Stack after Size() = %v, want %v", tt.s, tt.wantStack)
			}
		})
	}
}
