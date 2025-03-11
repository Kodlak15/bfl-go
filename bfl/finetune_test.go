package bfl

import (
	"encoding/base64"
	"io"
	"os"
	"testing"
)

func TestFinetune(t *testing.T) {
	key := os.Getenv("BFL_API_KEY")
	client := NewClient(key, "https://api.bfl.ai")
	zipFile, err := os.Open("../assets/test-finetune-images.zip")
	if err != nil {
		t.Fatal(err.Error())
	}
	defer zipFile.Close()
	zipBytes, err := io.ReadAll(zipFile)
	if err != nil {
		t.Fatal(err.Error())
	}
	encodedZip := base64.StdEncoding.EncodeToString(zipBytes)
	task := &FluxFinetune{
		FileData:        encodedZip,
		FinetuneComment: "test finetune",
		TriggerWord:     "TOK",
		Mode:            FinetuneModeGeneral,
		Iterations:      100,
		LearningRate:    0.003,
		Captioning:      true,
		Priority:        FinetunePriorityQuality,
		FinetuneType:    FinetuneTypeFull,
		LoraRank:        32,
	}
	ar, err := client.AsyncRequest(task.GetActionURL(client.BaseURL), task)
	if err != nil {
		t.Fatalf("Failed to create async request: %v", err)
	}
	resultResponse, err := Poll[*FinetuneResult, *FinetuneDetails](client, ar, true)
	if err != nil {
		t.Fatalf("Failed to poll result: %v", err)
	}
	t.Logf("Finetune result: %+v", resultResponse)
}
