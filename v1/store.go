package v1

import (
	"sync"
	"context"
	"time"
	"fmt"
)

type ToDoStore struct {
	sync.Mutex
	items   map[int]*ToDoItem
	archive map[int]*ToDoItem
	overdue map[int]*ToDoItem
	counter int
}

func NewStore() *ToDoStore {
	return &ToDoStore{
		items:   make(map[int]*ToDoItem),
		archive: make(map[int]*ToDoItem),
		overdue: make(map[int]*ToDoItem),
	}
}

func (s *ToDoStore) Add(title string, due *time.Time) int {
	s.Lock()
	defer s.Unlock()
	s.counter++
	s.items[s.counter] = &ToDoItem{
		ID:        s.counter,
		Title:     title,
		Status:    NotStarted,
		CreatedAt: time.Now(),
		Due:       due,
	}
	return s.counter
}

// MoveToArchiveWithTimeout tries to move a task to the archive but aborts if it takes too long.
func (s *ToDoStore) MoveToArchiveWithTimeout(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	done := make(chan error)

	go func() {
		time.Sleep(1 * time.Second) // Simulate some slow work
		s.Lock()
		defer s.Unlock()
		item, ok := s.items[id]
		if !ok {
			done <- fmt.Errorf("item not found")
			return
		}
		s.archive[id] = item
		delete(s.items, id)
		done <- nil
	}()

	select {
	case <-ctx.Done():
		return fmt.Errorf("timeout exceeded")
	case err := <-done:
		return err
	}
}

func (s *ToDoStore) Get(id int) (*ToDoItem, bool) {
	s.Lock()
	defer s.Unlock()

	item, exists := s.items[id]
	return item, exists
}

func (s *ToDoStore) Remove(id int) {
	s.Lock()
	defer s.Unlock()
	delete(s.items, id)
}

func (s *ToDoStore) List(statusFilter string, scope string) []*ToDoItem {
	s.Lock()
	defer s.Unlock()

	var result []*ToDoItem

	// Choose which map to list from
	var source map[int]*ToDoItem
	switch scope {
	case "archive":
		source = s.archive
	case "overdue":
		source = s.overdue
	default:
		source = s.items
	}

	for _, item := range source {
		if statusFilter == "" || string(item.Status) == statusFilter {
			result = append(result, item)
		}
	}
	return result
}




// More methods: Remove, UpdateStatus, MoveToArchive, CheckOverdue
