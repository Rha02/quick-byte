package handlers

// Repository is a struct that holds all supported handlers
type Repository struct{}

// NewRepository creates a new handler repository
func NewRepository() *Repository {
	return &Repository{}
}
