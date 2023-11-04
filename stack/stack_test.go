package stack

import "testing"

func (expected *Stack[T]) Equal(actual *Stack[T]) bool {
	return expected.ToString() == actual.ToString()
}

func TestPush(t *testing.T) {
	stack := CreateStack([]float64{})
	stack.Push(1.0)
	stack.Push(2.0)

	expected_stack := CreateStack([]float64{1.0, 2.0})
	if !expected_stack.Equal(stack) {
		t.Errorf("Result %v not equal to expected %v", stack.ToString(), expected_stack.ToString())
	}
}

func TestPop(t *testing.T) {
	stack := CreateStack([]int{1, 2})
	value, error := stack.Pop()
	if error != nil {
		t.Error("Pop did not return a value but should have")
	}

	expected_value := 2
	if value != expected_value {
		t.Errorf("Popped value %v not equal to expected %v", value, expected_value)
	}

    _, _ = stack.Pop()
	expected_stack := CreateStack([]int{})
	if !expected_stack.Equal(stack) {
		t.Errorf("Result %v not equal to expected %v", stack.ToString(), expected_stack.ToString())
	}

	value, error = stack.Pop()
	if error == nil {
		t.Error("Popping from empty stack returned a value but should not have")
	}
}
