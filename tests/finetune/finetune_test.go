package bfl

import (
	"context"
	"encoding/base64"
	"io"
	"os"
	"testing"

	"github.com/Kodlak15/bfl-go/bfl"
)

func TestFinetune(t *testing.T) {
	key := os.Getenv("BFL_API_KEY")
	client := bfl.NewClient(key, "https://api.bfl.ai")
	zipFile, err := os.Open("../../assets/test-finetune-images.zip")
	if err != nil {
		t.Fatal(err.Error())
	}
	defer zipFile.Close()
	zipBytes, err := io.ReadAll(zipFile)
	if err != nil {
		t.Fatal(err.Error())
	}
	encodedZip := base64.StdEncoding.EncodeToString(zipBytes)
	task := &bfl.FluxFinetune{
		FileData:        encodedZip,
		FinetuneComment: "test finetune",
		TriggerWord:     "TOK",
		Mode:            bfl.FinetuneModeGeneral,
		Iterations:      100,
		LearningRate:    0.003,
		Captioning:      true,
		Priority:        bfl.FinetunePriorityQuality,
		FinetuneType:    bfl.FinetuneTypeFull,
		LoraRank:        32,
	}
	ar, err := client.AsyncRequest(context.Background(), task)
	if err != nil {
		t.Fatalf("Failed to create async request: %v", err)
	}
	resultResponse, err := bfl.Poll[*bfl.FinetuneResult, *bfl.FinetuneDetails](context.Background(), client, ar, true)
	if err != nil {
		t.Fatalf("Failed to poll result: %v", err)
	}
	t.Logf("Finetune result: %+v", resultResponse)
}
