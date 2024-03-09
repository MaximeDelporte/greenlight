package data

import (
	"time"
)

// `json:"-"`
// The - (hyphen) directive can be used when you never want a particular struct field to appear in the JSON output. This is useful for fields that contain internal system information that isn’t relevant to your users, or sensitive information that you don’t want to expose (like the hash of a password).

// `json:",omitempty"`
// In contrast the omitempty directive hides a field in the JSON output if and only if the struct field value is empty

// `json:"runtime,omitempty,string"`
// You can use string on individual struct fields to force the data to be represented as a string in the JSON output.

type Movie struct {
	ID        int64     `json:"id"`                       // Unique integer ID for the Movie
	CreatedAt time.Time `json:"-"`                        // When it was added to the DB
	Title     string    `json:"title"`                    // Movie title
	Year      int32     `json:"year,omitempty"`           // Movie release year
	Runtime   int32     `json:"runtime,omitempty,string"` // Movie runtime (in minutes)
	Genres    []string  `json:"genres,omitempty"`         // (romance, comedy, etc.)
	Version   int32     `json:"version"`
	// The version number starts at 1 and will be incremented each time the movie information is updated
}
