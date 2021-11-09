package value_vs_pointer_semantic_iteration

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Post struct {
	ID       int       `json:"id"`
	PostedAt time.Time `json:"posted_at"`
	Author   string    `json:"author"`
	Content  string    `json:"content"`
}

func encodePostsWithValueSemanticIteration(posts []Post, w io.Writer) error {
	e := json.NewEncoder(w)
	for i, post := range posts {
		if err := e.Encode(post); err != nil {
			return fmt.Errorf("failed encoding post #%d - %w", i, err)
		}
	}
	return nil
}

func encodePostsWithPointerSemanticIteration(posts []Post, w io.Writer) error {
	e := json.NewEncoder(w)
	for i := range posts {
		if err := e.Encode(posts[i]); err != nil {
			return fmt.Errorf("failed encoding post #%d - %w", i, err)
		}
	}
	return nil
}
