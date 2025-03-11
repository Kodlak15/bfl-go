package generate

import (
	"os"
	"testing"

	"github.com/Kodlak15/bfl"
)

func TestGenerateDev(t *testing.T) {
	key := os.Getenv("BFL_API_KEY")
	client := bfl.NewBFL(key, "https://api.bfl.ai")
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
	url := task.GetActionURL(client.BaseURL)
	ar, err := bfl.AsyncRequest(client, url, task)
	if err != nil {
		t.Fatalf("Failed to create async request: %v", err)
	}
	resultResponse, err := bfl.Poll[*bfl.GenerateResult, *bfl.GenerateDetails](client, ar, true)
	if err != nil {
		t.Fatalf("Failed to poll result: %v", err)
	}
	sampleURL := resultResponse.Result.SampleURL
	t.Logf("Sample URL: %s", sampleURL)
}

func TestGeneratePro11UltraFinetuned(t *testing.T) {
	key := os.Getenv("BFL_API_KEY")
	finetuneID := os.Getenv("TEST_FINETUNE_ID")
	client := bfl.NewBFL(key, "https://api.bfl.ai")
	task := &bfl.FluxPro11UltraFinetunedGenerate{
		FinetuneID:       finetuneID,
		FinetuneStrength: 1.1,
		Prompt:           "TOK getting a tattoo of a dragon",
		PromptUpsampling: false,
		Seed:             42,
		AspectRatio:      "16:9",
		SafetyTolerance:  2,
		OutputFormat:     "jpeg",
		Raw:              false,
	}
	url := task.GetActionURL(client.BaseURL)
	ar, err := bfl.AsyncRequest(client, url, task)
	if err != nil {
		t.Fatalf("Failed to create async request: %v", err)
	}
	resultResponse, err := bfl.Poll[*bfl.GenerateResult, *bfl.GenerateDetails](client, ar, true)
	if err != nil {
		t.Fatalf("Failed to poll result: %v", err)
	}
	sampleURL := resultResponse.Result.SampleURL
	t.Logf("Sample URL: %s", sampleURL)
}
