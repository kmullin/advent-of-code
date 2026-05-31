package aoc

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type Client struct {
	client *http.Client

	session  string
	cacheDir string
}

func NewClient(session, cacheDir string) *Client {
	return &Client{
		session:  session,
		cacheDir: cacheDir,
		client:   http.DefaultClient,
	}
}

func (c *Client) GetInput(year, day int) ([]byte, error) {
	path := filepath.Join(
		c.cacheDir,
		fmt.Sprintf("%d", year),
		fmt.Sprintf("%02d.txt", day),
	)

	// we hit the cache
	if b, err := os.ReadFile(path); err == nil {
		return b, nil
	}

	u := fmt.Sprintf("https://adventofcode.com/%v/day/%v/input", year, day)
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, fmt.Errorf("err creating request: %w", err)
	}

	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: c.session,
	})

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("err fetching %q: %w", u, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server returned %v", resp.StatusCode)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("err reading response body: %w", err)
	}

	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return nil, err
	}

	if err := os.WriteFile(path, b, 0o644); err != nil {
		return nil, err
	}
	return b, nil
}
