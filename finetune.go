package bfl

import (
	"fmt"
)

type FinetuneMode string

const (
	FinetuneModeGeneral   FinetuneMode = "general"
	FinetuneModeCharacter FinetuneMode = "character"
	FinetuneModeStyle     FinetuneMode = "style"
	FinetuneModeProduct   FinetuneMode = "product"
)

type FinetunePriority string

const (
	FinetunePrioritySpeed       FinetunePriority = "speed"
	FinetunePriorityQuality     FinetunePriority = "quality"
	FinetunePriorityHighResOnly FinetunePriority = "high_res_only"
)

type FinetuneType string

const (
	FinetuneTypeLora FinetuneType = "lora"
	FinetuneTypeFull FinetuneType = "full"
)

type LoraRank int

const (
	LoraRank16 LoraRank = 16
	LoraRank32 LoraRank = 32
)

// Task parameters for finetuning a flux model through the BFL API.
type FluxFinetune struct {
	// Base64-encoded ZIP file containing training images and, optionally, corresponding captions.
	FileData string `json:"file_data"`
	// Comment or name of the fine-tuned model. This will be added as a field to the finetune_details.
	FinetuneComment string `json:"finetune_comment"`
	// Trigger word for the fine-tuned model.
	// Default: TOK.
	TriggerWord string `json:"trigger_word"`
	// Mode for the fine-tuned model. Allowed values are 'general', 'character', 'style', 'product'. This will affect the caption behaviour. General will describe the image in full detail.
	Mode FinetuneMode `json:"mode"`
	// Number of iterations for fine-tuning.
	// Min: 100, Max: 1000, Default: 300.
	Iterations int `json:"iterations"`
	// Learning rate for fine-tuning. If not provided, defaults to 1e-5 for full fine-tuning and 1e-4 for lora fine-tuning.
	// Min: 0.000001, Max: 0.005.
	LearningRate float64 `json:"learning_rate,omitempty"`
	// Whether to enable captioning during fine-tuning.
	// Default: true.
	Captioning bool `json:"captioning"`
	// Priority of the fine-tuning process. 'speed' will prioritize iteration speed over quality, 'quality' will prioritize quality over speed.
	// Default: quality.
	Priority FinetunePriority `json:"priority"`
	// Type of fine-tuning. 'lora' is a standard LoRA Adapter, 'full' is a full fine-tuning mode, with a post hoc lora extraction.
	// Default: full.
	FinetuneType FinetuneType `json:"finetune_type"`
	// Rank of the fine-tuned model. 16 or 32. If finetune_type is 'full', this will be the rank of the extracted lora model.
	// Default: 32.
	LoraRank LoraRank `json:"lora_rank"`
	// URL to receive webhook notifications.
	// Min length: 1, Max length: 2083.
	WebhookURL string `json:"webhook_url,omitempty"`
	// Optional secret for webhook signature verification.
	WebhookSecret string `json:"webhook_secret,omitempty"`
}

// TODO: unimplemented
type FinetuneResult struct{}

// TODO: unimplemented
type FinetuneDetails struct{}

func (f *FluxFinetune) GetActionURL(baseURL string) string {
	return fmt.Sprintf("%s/v1/finetune", baseURL)
}
