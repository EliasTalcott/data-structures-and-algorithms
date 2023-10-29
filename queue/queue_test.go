package queue

import "testing"

func (expected *Queue) Equal(actual *Queue) bool {
	return expected.ToString() == actual.ToString()
}

func TestEnqueue(t *testing.T) {
	queue := CreateQueue([]int{})
	queue.Enqueue(1)
	queue.Enqueue(2)

	expected_queue := CreateQueue([]int{1, 2})
	if !expected_queue.Equal(queue) {
		t.Errorf("Result %v not equal to expected %v", queue.ToString(), expected_queue.ToString())
	}
}

func TestDequeue(t *testing.T) {
	queue := CreateQueue([]int{1})
	value, error := queue.Dequeue()
	if error != nil {
		t.Error("Dequeue did not return a value but should have")
	}

	expected_value := 1
	if value != expected_value {
		t.Errorf("Dequeued value %v not equal to expected %v", value, expected_value)
	}

	expected_queue := CreateQueue([]int{})
	if !expected_queue.Equal(queue) {
		t.Errorf("Result %v not equal to expected %v", queue.ToString(), expected_queue.ToString())
	}

	value, error = queue.Dequeue()
	if error == nil {
		t.Error("Dequeueing from empty queue returned a value but should not have")
	}
}
