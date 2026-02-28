package goutils

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"strconv"
)

// GetReader checks the CLI arguments provided to the program.
// If no argument is provided or if the first argument is "-", it returns os.Stdin.
// Otherwise, it attempts to open the file path provided in the first argument.
//
// The caller is responsible for calling .Close() on the returned io.ReadCloser.
func GetReader() (io.ReadCloser, error) {

	if len(os.Args) > 1 && os.Args[1] != "-" {
		file, err := os.Open(os.Args[1])
		if err != nil {
			return nil, err
		}
		return file, nil
	}
	return os.Stdin, nil
}

// check if a string is a palindrome
func Is_palindrome(input string) bool {

	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] {
			return false
		}
	}
	return true
}

// return the md5hash of a string
func Md5Hash(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	bs := h.Sum(nil)

	return fmt.Sprintf("%x", bs)
}

// Convert a string to int Panic if not possible
func ToInt(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return val
}


// ==============================================================
// Queue data type

type queue[T any] struct {
	array []T
}

// Add an element t to the queue
func (q *queue[T]) Enqueue (t T) {
	q.array = append(q.array, t)
}

// Remove and return the first element of queue. If queue is empty returns an err
func (q * queue[T]) Dequeue () (T, error) {
	var zero T
	if len(q.array) == 0 {
		return zero, fmt.Errorf("Queue is empty")
	}

	item := q.array[0]
	q.array = q.array[1:]
	return item, nil
}

// Return a reference to the element at the front of queue Q,
// without removing it; an error occurs if the queue is empty.
func (q * queue[T]) First () (T, error) {
	var zero T
	if q.IsEmpty() {
		return zero, fmt.Errorf("Queue is empty")
	}
	return q.array[0], nil
}

// Return True if queue Q does not contain any elements.
func(q * queue[T]) IsEmpty () bool {
	return len(q.array) == 0
}

// Return the number of elements in queue Q
func (q * queue[T]) Len() int {
	return len(q.array)
}

// NewQueue creates a new empty Queue for type T
//	Example Usage:
// 	q := NewQueue[int]()
//	q.Enqueue(5)
func NewQueue[T any]() *queue[T] {
    return &queue[T] {
        array: make([]T, 0),
    }
}