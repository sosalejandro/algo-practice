package simple_stack

import (
	"testing"

	"github.com/sosalejandro/data-structures/common"
)

type IntItem struct {
	value int
}

func (i IntItem) Equals(item common.Item[int]) bool {
	return i.value == item.Value()
}

func (i IntItem) Value() int {
	return i.value
}

func (i IntItem) IsEmpty() bool {
	return i.value == 0
}

func TestStack_Push(t *testing.T) {
	tests := []struct {
		name      string
		initial   []IntItem
		toPush    IntItem
		wantError bool
	}{
		{
			name:      "Push valid element",
			initial:   []IntItem{},
			toPush:    IntItem{value: 1},
			wantError: false,
		},
		{
			name:      "Push empty element",
			initial:   []IntItem{},
			toPush:    IntItem{value: 0},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := NewStack[int, IntItem]()
			for _, item := range tt.initial {
				stack.Push(item)
			}

			err := stack.Push(tt.toPush)
			if (err != nil) != tt.wantError {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

func TestStack_Pop(t *testing.T) {
	tests := []struct {
		name      string
		initial   []IntItem
		want      IntItem
		wantError bool
	}{
		{
			name:      "Pop from non-empty stack",
			initial:   []IntItem{{value: 1}},
			want:      IntItem{value: 1},
			wantError: false,
		},
		{
			name:      "Pop from empty stack",
			initial:   []IntItem{},
			want:      IntItem{},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := NewStack[int, IntItem]()
			for _, item := range tt.initial {
				stack.Push(item)
			}

			got, err := stack.Pop()
			if (err != nil) != tt.wantError {
				t.Errorf("unexpected error: %v", err)
			}
			if !got.Equals(tt.want) {
				t.Errorf("expected %v, got %v", tt.want.Value(), got.Value())
			}
		})
	}
}

func TestStack_Peek(t *testing.T) {
	tests := []struct {
		name    string
		initial []IntItem
		want    IntItem
	}{
		{
			name:    "Peek from non-empty stack",
			initial: []IntItem{{value: 1}},
			want:    IntItem{value: 1},
		},
		{
			name:    "Peek from empty stack",
			initial: []IntItem{},
			want:    IntItem{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := NewStack[int, IntItem]()
			for _, item := range tt.initial {
				stack.Push(item)
			}

			got := stack.Peek()
			if !got.Equals(tt.want) {
				t.Errorf("expected %v, got %v", tt.want.Value(), got.Value())
			}
		})
	}
}

func TestStack_IsEmpty(t *testing.T) {
	tests := []struct {
		name    string
		initial []IntItem
		want    bool
	}{
		{
			name:    "Empty stack",
			initial: []IntItem{},
			want:    true,
		},
		{
			name:    "Non-empty stack",
			initial: []IntItem{{value: 1}},
			want:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := NewStack[int, IntItem]()
			for _, item := range tt.initial {
				stack.Push(item)
			}

			got := stack.IsEmpty()
			if got != tt.want {
				t.Errorf("expected %v, got %v", tt.want, got)
			}
		})
	}
}

func TestStack_Size(t *testing.T) {
	tests := []struct {
		name    string
		initial []IntItem
		want    int
	}{
		{
			name:    "Empty stack",
			initial: []IntItem{},
			want:    0,
		},
		{
			name:    "Non-empty stack",
			initial: []IntItem{{value: 1}, {value: 2}},
			want:    2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := NewStack[int, IntItem]()
			for _, item := range tt.initial {
				stack.Push(item)
			}

			got := stack.Size()
			if got != tt.want {
				t.Errorf("expected %v, got %v", tt.want, got)
			}
		})
	}
}

func TestStack_Clear(t *testing.T) {
	tests := []struct {
		name    string
		initial []IntItem
	}{
		{
			name:    "Clear empty stack",
			initial: []IntItem{},
		},
		{
			name:    "Clear non-empty stack",
			initial: []IntItem{{value: 1}, {value: 2}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := NewStack[int, IntItem]()
			for _, item := range tt.initial {
				stack.Push(item)
			}

			stack.Clear()
			if !stack.IsEmpty() {
				t.Error("expected stack to be empty after clear")
			}
		})
	}
}

func TestStack_ToSlice(t *testing.T) {
	tests := []struct {
		name    string
		initial []IntItem
		want    []IntItem
	}{
		{
			name:    "Empty stack",
			initial: []IntItem{},
			want:    []IntItem{},
		},
		{
			name:    "Non-empty stack",
			initial: []IntItem{{value: 1}, {value: 2}},
			want:    []IntItem{{value: 1}, {value: 2}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := NewStack[int, IntItem]()
			for _, item := range tt.initial {
				stack.Push(item)
			}

			got := stack.ToSlice()
			if len(got) != len(tt.want) {
				t.Errorf("expected length %v, got %v", len(tt.want), len(got))
			}
			for i, item := range got {
				if !item.Equals(tt.want[i]) {
					t.Errorf("expected %v, got %v", tt.want[i].Value(), item.Value())
				}
			}
		})
	}
}

func TestStack_Contains(t *testing.T) {
	tests := []struct {
		name    string
		initial []IntItem
		toCheck IntItem
		want    bool
	}{
		{
			name:    "Contains element",
			initial: []IntItem{{value: 1}, {value: 2}},
			toCheck: IntItem{value: 1},
			want:    true,
		},
		{
			name:    "Does not contain element",
			initial: []IntItem{{value: 1}, {value: 2}},
			toCheck: IntItem{value: 3},
			want:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := NewStack[int, IntItem]()
			for _, item := range tt.initial {
				stack.Push(item)
			}

			got := stack.Contains(tt.toCheck)
			if got != tt.want {
				t.Errorf("expected %v, got %v", tt.want, got)
			}
		})
	}
}
