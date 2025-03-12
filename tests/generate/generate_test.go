package bfl

import (
	"context"
	"os"
	"testing"

	"github.com/Kodlak15/bfl-go/bfl"
)

func TestGenerateDev(t *testing.T) {
	key := os.Getenv("BFL_API_KEY")
	client := bfl.NewClient(key, "https://api.bfl.ai")
	task := &bfl.FluxDevGenerate{
		Prompt:           "A beautiful landscape with a river and mountains",
		ImagePrompt:      "",
		Width:            1024,
		Height:           768,
		Steps:            28,
		PromptUpsampling: false,
		Seed:             42,
		Guidance:         3,
		SafetyTolerance:  2,
		OutputFormat:     "jpeg",
	}
	result, err := bfl.Generate(context.Background(), client, task)
	if err != nil {
		t.Fatalf("Failed to generate: %v", err)
	}
	t.Log(result.SampleURL)
}
