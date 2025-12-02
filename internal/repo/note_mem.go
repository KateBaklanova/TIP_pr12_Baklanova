package repo

import (
	"fmt"
	"sync"
	"time"

	"Kate.com/notes-api/internal/core"
)

type NoteRepoMem struct {
	mu    sync.Mutex
	notes map[int64]*core.Note
	next  int64
}

func NewNoteRepoMem() *NoteRepoMem {
	return &NoteRepoMem{notes: make(map[int64]*core.Note)}
}

func (r *NoteRepoMem) Create(n core.Note) (int64, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.next++
	n.ID = r.next
	r.notes[n.ID] = &n
	return n.ID, nil
}

func (r *NoteRepoMem) GetByID(id int64) (*core.Note, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	note, exists := r.notes[id]
	if !exists {
		return nil, fmt.Errorf("note not found")
	}
	return note, nil
}

func (r *NoteRepoMem) GetAll() []*core.Note {
	r.mu.Lock()
	defer r.mu.Unlock()

	notes := make([]*core.Note, 0, len(r.notes))
	for _, note := range r.notes {
		notes = append(notes, note)
	}
	return notes
}

func (r *NoteRepoMem) Update(id int64, updates map[string]interface{}) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	note, exists := r.notes[id]
	if !exists {
		return fmt.Errorf("note not found")
	}

	if title, ok := updates["title"].(string); ok {
		note.Title = title
	}
	if content, ok := updates["content"].(string); ok {
		note.Content = content
	}

	now := time.Now()
	note.UpdatedAt = &now

	return nil
}

func (r *NoteRepoMem) Delete(id int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.notes[id]; !exists {
		return fmt.Errorf("note not found")
	}

	delete(r.notes, id)
	return nil
}
