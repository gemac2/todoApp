package models

import (
	"encoding/json"
	"time"

	"github.com/gofrs/uuid"
)

type Todo struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Complete    bool      `json:"complete" db:"complete"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (t Todo) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}

// Todoes is not required by pop and may be deleted
type Todoes []Todo

// String is not required by pop and may be deleted
func (t Todoes) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}
