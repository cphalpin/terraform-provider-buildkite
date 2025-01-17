package buildkite

import (
	"os"
	"testing"
)

func TestGetOrganizationIDMissing(t *testing.T) {
	slug := "doesnt match API key"

	config := &clientConfig{
		org:        slug,
		apiToken:   os.Getenv("BUILDKITE_API_TOKEN"),
		graphqlURL: defaultGraphqlEndpoint,
		restURL:    defaultRestEndpoint,
		userAgent:  "test-user-agent",
	}

	// NewClient calls GetOrganizationId so we can test the output
	client, err := NewClient(config)
	if err == nil {
		t.Fatalf("err: %s", err)
	}
	if client != nil {
		t.Fatalf("Nonexistent organization found")
	}
}
