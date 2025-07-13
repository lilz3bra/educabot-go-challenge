package providers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"educabot.com/bookshop/models"
)

type BooksProvider interface {
	GetBooks(ctx context.Context) ([]models.Book, error)
}

type APIClient struct {
	HTTPClient *http.Client
	BaseURL    string
}

// Constructor para el cliente
func NewBookAPIClient(baseURL string) *APIClient {
	return &APIClient{
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		BaseURL: baseURL,
	}
}

func (c *APIClient) GetBooks(ctx context.Context) ([]models.Book, error) {
	url := "https://6781684b85151f714b0aa5db.mockapi.io/api/v1/books" // TODO: Adaptar este path

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// TODO: Agregar Headers si son necesarios,.
	// req.Header.Set("Authorization", "Bearer API_KEY")

	// Ejecutar el request usando el cliente configurado.
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request to %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 status code: %d from %s", resp.StatusCode, url)
	}

	var book []models.Book
	if err := json.NewDecoder(resp.Body).Decode(&book); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	return book, nil
}
