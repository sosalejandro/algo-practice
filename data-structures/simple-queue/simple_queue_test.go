package simple_queue

import (
	"testing"

	"github.com/sosalejandro/algo-practice/data-structures/common"
)

type testStringItem struct {
	value string
}

func (t *testStringItem) Equals(other common.Item[string]) bool {
	return t.value == other.Value()
}

func (t *testStringItem) Value() string {
	return t.value
}

func (t *testStringItem) IsEmpty() bool {
	return t.value == ""
}

func TestQueue_Dequeue(t *testing.T) {
	tests := []struct {
		name        string
		initial     []*testStringItem
		expected    *testStringItem
		expectError bool
	}{
		{
			name:        "Dequeue from empty queue",
			initial:     []*testStringItem{},
			expected:    nil,
			expectError: true,
		},
		{
			name: "Dequeue from queue with one element",
			initial: []*testStringItem{
				{value: "element1"},
			},
			expected:    &testStringItem{value: "element1"},
			expectError: false,
		},
		{
			name: "Dequeue from queue with multiple elements",
			initial: []*testStringItem{
				{value: "element1"},
				{value: "element2"},
			},
			expected:    &testStringItem{value: "element1"},
			expectError: false,
		},
		{
			name: "Dequeue until queue is empty",
			initial: []*testStringItem{
				{value: "element1"},
				{value: "element2"},
			},
			expected:    &testStringItem{value: "element1"},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queue := NewQueue[string, *testStringItem]()
			for _, item := range tt.initial {
				queue.Enqueue(item)
			}

			result, err := queue.Dequeue()
			if (err != nil) != tt.expectError {
				t.Errorf("expected error: %v, got: %v", tt.expectError, err)
			}

			if !tt.expectError && !result.Equals(tt.expected) {
				t.Errorf("expected: %v, got: %v", tt.expected.Value(), result.Value())
			}

			if tt.name == "Dequeue until queue is empty" {
				_, err := queue.Dequeue()
				if err != nil {
					t.Errorf("expected no error, got: %v", err)
				}
				if !queue.IsEmpty() {
					t.Errorf("expected queue to be empty, but it wasn't")
				}
			}
		})
	}
}

func TestQueue_Enqueue(t *testing.T) {
	tests := []struct {
		name        string
		initial     []*testStringItem
		toEnqueue   *testStringItem
		expectError bool
	}{
		{
			name:        "Enqueue nil element",
			initial:     []*testStringItem{},
			toEnqueue:   &testStringItem{value: ""},
			expectError: true,
		},
		{
			name: "Enqueue valid element",
			initial: []*testStringItem{
				{value: "element1"},
			},
			toEnqueue:   &testStringItem{value: "element2"},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queue := NewQueue[string, *testStringItem]()
			for _, item := range tt.initial {
				queue.Enqueue(item)
			}

			err := queue.Enqueue(tt.toEnqueue)
			if (err != nil) != tt.expectError {
				t.Errorf("expected error: %v, got: %v", tt.expectError, err)
			}

			if !tt.expectError && !queue.Contains(tt.toEnqueue) {
				t.Errorf("expected queue to contain: %v", tt.toEnqueue.Value())
			}
		})
	}
}

func TestQueue_Peek(t *testing.T) {
	tests := []struct {
		name     string
		initial  []*testStringItem
		expected *testStringItem
	}{
		{
			name:     "Peek from empty queue",
			initial:  []*testStringItem{},
			expected: nil,
		},
		{
			name: "Peek from queue with elements",
			initial: []*testStringItem{
				{value: "element1"},
				{value: "element2"},
			},
			expected: &testStringItem{value: "element1"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queue := NewQueue[string, *testStringItem]()
			for _, item := range tt.initial {
				queue.Enqueue(item)
			}

			result := queue.Peek()
			if tt.expected == nil && result != nil {
				t.Errorf("expected: nil, got: %v", result.Value())
			} else if tt.expected != nil && !result.Equals(tt.expected) {
				t.Errorf("expected: %v, got: %v", tt.expected.Value(), result.Value())
			}
		})
	}
}

func TestQueue_IsEmpty(t *testing.T) {
	tests := []struct {
		name     string
		initial  []*testStringItem
		expected bool
	}{
		{
			name:     "Empty queue",
			initial:  []*testStringItem{},
			expected: true,
		},
		{
			name: "Non-empty queue",
			initial: []*testStringItem{
				{value: "element1"},
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queue := NewQueue[string, *testStringItem]()
			for _, item := range tt.initial {
				queue.Enqueue(item)
			}

			result := queue.IsEmpty()
			if result != tt.expected {
				t.Errorf("expected: %v, got: %v", tt.expected, result)
			}
		})
	}
}

func TestQueue_Size(t *testing.T) {
	tests := []struct {
		name     string
		initial  []*testStringItem
		expected int
	}{
		{
			name:     "Empty queue",
			initial:  []*testStringItem{},
			expected: 0,
		},
		{
			name: "Queue with elements",
			initial: []*testStringItem{
				{value: "element1"},
				{value: "element2"},
			},
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queue := NewQueue[string, *testStringItem]()
			for _, item := range tt.initial {
				queue.Enqueue(item)
			}

			result := queue.Size()
			if result != tt.expected {
				t.Errorf("expected: %v, got: %v", tt.expected, result)
			}
		})
	}
}

func TestQueue_Clear(t *testing.T) {
	tests := []struct {
		name    string
		initial []*testStringItem
	}{
		{
			name:    "Clear empty queue",
			initial: []*testStringItem{},
		},
		{
			name: "Clear non-empty queue",
			initial: []*testStringItem{
				{value: "element1"},
				{value: "element2"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queue := NewQueue[string, *testStringItem]()
			for _, item := range tt.initial {
				queue.Enqueue(item)
			}

			queue.Clear()
			if !queue.IsEmpty() {
				t.Errorf("expected queue to be empty after clear")
			}
		})
	}
}

func TestQueue_ToSlice(t *testing.T) {
	tests := []struct {
		name     string
		initial  []*testStringItem
		expected []*testStringItem
	}{
		{
			name:     "Empty queue",
			initial:  []*testStringItem{},
			expected: []*testStringItem{},
		},
		{
			name: "Queue with elements",
			initial: []*testStringItem{
				{value: "element1"},
				{value: "element2"},
			},
			expected: []*testStringItem{
				{value: "element1"},
				{value: "element2"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queue := NewQueue[string, *testStringItem]()
			for _, item := range tt.initial {
				queue.Enqueue(item)
			}

			result := queue.ToSlice()
			if len(result) != len(tt.expected) {
				t.Errorf("expected length: %v, got: %v", len(tt.expected), len(result))
			}

			for i, item := range result {
				if !item.Equals(tt.expected[i]) {
					t.Errorf("expected: %v, got: %v", tt.expected[i].Value(), item.Value())
				}
			}
		})
	}
}

func TestQueue_Contains(t *testing.T) {
	tests := []struct {
		name     string
		initial  []*testStringItem
		toCheck  *testStringItem
		expected bool
	}{
		{
			name:     "Empty queue",
			initial:  []*testStringItem{},
			toCheck:  &testStringItem{value: "element1"},
			expected: false,
		},
		{
			name: "Queue contains element",
			initial: []*testStringItem{
				{value: "element1"},
				{value: "element2"},
			},
			toCheck:  &testStringItem{value: "element1"},
			expected: true,
		},
		{
			name: "Queue does not contain element",
			initial: []*testStringItem{
				{value: "element1"},
				{value: "element2"},
			},
			toCheck:  &testStringItem{value: "element3"},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queue := NewQueue[string, *testStringItem]()
			for _, item := range tt.initial {
				queue.Enqueue(item)
			}

			result := queue.Contains(tt.toCheck)
			if result != tt.expected {
				t.Errorf("expected: %v, got: %v", tt.expected, result)
			}
		})
	}
}
