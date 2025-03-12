package bfl

import (
	"context"
	"fmt"
)

type GenerateResult struct {
	Prompt    string  `json:"prompt"`
	SampleURL string  `json:"sample"`
	Seed      int     `json:"seed"`
	StartTime float64 `json:"start_time"`
	EndTime   float64 `json:"end_time"`
	Duration  float64 `json:"duration"`
}

// TODO: unimplemented
type GenerateDetails struct{}

// An async task for generating an image.
type GenerateTask interface {
	AsyncTask
	GenerateTaskMarker()
}

// Submit an image generation task and poll for the result.
func Generate(ctx context.Context, c *Client, task GenerateTask) (*GenerateResult, error) {
	ar, err := c.AsyncRequest(ctx, task)
	if err != nil {
		return nil, err
	}
	result, err := Poll[*GenerateResult, *GenerateDetails](ctx, c, ar, true)
	if err != nil {
		return nil, err
	}
	return result.Result, nil
}

// Task parameters for generating an image with Flux Pro 1.1 through the BFL API.
type FluxPro11Generate struct {
	// Text prompt for image generation.
	Prompt string `json:"prompt,omitempty"`
	// Optional base64 encoded image to use with Flux Redux.
	ImagePrompt string `json:"image_prompt,omitempty"`
	// Width of the generated image in pixels. Must be a multiple of 32.
	// Min: 256, Max: 1440, Default: 1024.
	Width int `json:"width"`
	// Height of the generated image in pixels. Must be a multiple of 32.
	// Min: 256, Max: 1440, Default: 768.
	Height int `json:"height"`
	// Whether to perform upsampling on the prompt.
	// If active, automatically modifies the prompt for more creative generation.
	// Default: false.
	PromptUpsampling bool `json:"prompt_upsampling"`
	// Optional seed for reproducibility.
	Seed int `json:"seed,omitempty"`
	// Tolerance level for input and output moderation.
	// Between 0 and 6, 0 being most strict, 6 being least strict.
	// Min: 0, Max: 6, Default: 2.
	SafetyTolerance int `json:"safety_tolerance"`
	// Output format for the generated image.
	// Can be 'jpeg' or 'png'.
	// Default: jpeg.
	OutputFormat string `json:"output_format,omitempty"`
	// URL to receive webhook notifications.
	// Min length: 1, Max length: 2083.
	WebhookURL string `json:"webhook_url,omitempty"`
	// Optional secret for webhook signature verification.
	WebhookSecret string `json:"webhook_secret,omitempty"`
}

func (flx *FluxPro11Generate) GenerateTaskMarker() {}

func (flx *FluxPro11Generate) GetActionURL(baseURL string) string {
	return fmt.Sprintf("%s/v1/flux-pro-1.1", baseURL)
}

// Task parameters for generating an image with Flux Pro through the BFL API.
type FluxProGenerate struct {
	// Text prompt for image generation.
	Prompt string `json:"prompt,omitempty"`
	// Optional base64 encoded image to use as a prompt for generation.
	ImagePrompt string `json:"image_prompt,omitempty"`
	// Width of the generated image in pixels. Must be a multiple of 32.
	// Min: 256, Max: 1440, Default: 1024.
	Width int `json:"width"`
	// Height of the generated image in pixels. Must be a multiple of 32.
	// Min: 256, Max: 1440, Default: 768.
	Height int `json:"height"`
	// Number of steps for the image generation process.
	// Min: 1, Max: 50, Default: 40.
	Steps int `json:"steps,omitempty"`
	// Whether to perform upsampling on the prompt. If active, automatically modifies the prompt for more creative generation.
	// Default: false.
	PromptUpsampling bool `json:"prompt_upsampling"`
	// Optional seed for reproducibility.
	Seed int `json:"seed,omitempty"`
	// Guidance scale for image generation. High guidance scales improve prompt adherence at the cost of reduced realism.
	// Min: 1.5, Max: 5, Default: 2.5.
	Guidance float64 `json:"guidance,omitempty"`
	// Tolerance level for input and output moderation. Between 0 and 6, 0 being most strict, 6 being least strict.
	// Min: 0, Max: 6, Default: 2.
	SafetyTolerance int `json:"safety_tolerance"`
	// Interval parameter for guidance control.
	// Min: 1, Max: 4, Default: 2.
	Interval float64 `json:"interval,omitempty"`
	// Output format for the generated image. Can be 'jpeg' or 'png'.
	// Default: jpeg.
	OutputFormat string `json:"output_format,omitempty"`
	// URL to receive webhook notifications.
	// Min length: 1, Max length: 2083.
	WebhookURL string `json:"webhook_url,omitempty"`
	// Optional secret for webhook signature verification.
	WebhookSecret string `json:"webhook_secret,omitempty"`
}

func (flx *FluxProGenerate) GenerateTaskMarker() {}

func (flx *FluxProGenerate) GetActionURL(baseURL string) string {
	return fmt.Sprintf("%s/v1/flux-pro", baseURL)
}

// Task parameters for generating an image with Flux Dev through the BFL API.
type FluxDevGenerate struct {
	// Text prompt for image generation.
	Prompt string `json:"prompt"`
	// Optional base64 encoded image to use as a prompt for generation.
	ImagePrompt string `json:"image_prompt,omitempty"`
	// Width of the generated image in pixels. Must be a multiple of 32.
	// Min: 256, Max: 1440, Default: 1024.
	Width int `json:"width"`
	// Height of the generated image in pixels. Must be a multiple of 32.
	// Min: 256, Max: 1440, Default: 768.
	Height int `json:"height"`
	// Number of steps for the image generation process.
	// Min: 1, Max: 50, Default: 28.
	Steps int `json:"steps,omitempty"`
	// Whether to perform upsampling on the prompt.
	// If active, automatically modifies the prompt for more creative generation.
	// Default: false.
	PromptUpsampling bool `json:"prompt_upsampling"`
	// Optional seed for reproducibility.
	Seed int `json:"seed,omitempty"`
	// Guidance scale for image generation.
	// High guidance scales improve prompt adherence at the cost of reduced realism.
	// Min: 1.5, Max: 5, Default: 3.
	Guidance float64 `json:"guidance,omitempty"`
	// Tolerance level for input and output moderation.
	// Between 0 and 6, 0 being most strict, 6 being least strict.
	// Min: 0, Max: 6, Default: 2.
	SafetyTolerance int `json:"safety_tolerance"`
	// Output format for the generated image.
	// Can be 'jpeg' or 'png'.
	// Default: jpeg.
	OutputFormat string `json:"output_format,omitempty"`
	// URL to receive webhook notifications.
	// Min length: 1, Max length: 2083.
	WebhookURL string `json:"webhook_url,omitempty"`
	// Optional secret for webhook signature verification.
	WebhookSecret string `json:"webhook_secret,omitempty"`
}

func (flx *FluxDevGenerate) GenerateTaskMarker() {}

func (flx *FluxDevGenerate) GetActionURL(baseURL string) string {
	return fmt.Sprintf("%s/v1/flux-dev", baseURL)
}

// Task parameters for generating an image with Flux Pro 1.1 Ultra through the BFL API.
type FluxPro11UltraGenerate struct {
	// The prompt to use for image generation.
	Prompt string `json:"prompt,omitempty"`
	// Whether to perform upsampling on the prompt. If active, automatically modifies the prompt for more creative generation.
	// Default: false.
	PromptUpsampling bool `json:"prompt_upsampling"`
	// Optional seed for reproducibility. If not provided, a random seed will be used.
	Seed int `json:"seed,omitempty"`
	// Aspect ratio of the image between 21:9 and 9:21.
	// Default: 16:9.
	AspectRatio string `json:"aspect_ratio"`
	// Tolerance level for input and output moderation. Between 0 and 6, 0 being most strict, 6 being least strict.
	// Min: 0, Max: 6, Default: 2.
	SafetyTolerance int `json:"safety_tolerance"`
	// Output format for the generated image. Can be 'jpeg' or 'png'.
	// Default: jpeg.
	OutputFormat string `json:"output_format,omitempty"`
	// Generate less processed, more natural-looking images.
	// Default: false.
	Raw bool `json:"raw"`
	// Optional image to remix in base64 format.
	ImagePrompt string `json:"image_prompt,omitempty"`
	// Blend between the prompt and the image prompt.
	// Min: 0, Max: 1, Default: 0.1.
	ImagePromptStrength float64 `json:"image_prompt_strength"`
	// URL to receive webhook notifications.
	// Min length: 1, Max length: 2083.
	WebhookURL string `json:"webhook_url,omitempty"`
	// Optional secret for webhook signature verification.
	WebhookSecret string `json:"webhook_secret,omitempty"`
}

func (flx *FluxPro11UltraGenerate) GenerateTaskMarker() {}

func (flx *FluxPro11UltraGenerate) GetActionURL(baseURL string) string {
	return fmt.Sprintf("%s/v1/flux-pro-1.1-ultra", baseURL)
}

// Task parameters for generating an image with Flux Pro 1.0 Fill through the BFL API.
type FluxProFillGenerate struct {
	// A Base64-encoded string representing the image you wish to modify. Can contain alpha mask if desired.
	Image string `json:"image"`
	// A Base64-encoded string representing a mask for the areas you want to modify in the image. The mask should be the same dimensions as the image and in black and white. Black areas (0%) indicate no modification, while white areas (100%) specify areas for inpainting. Optional if you provide an alpha mask in the original image.
	Mask string `json:"mask,omitempty"`
	// The description of the changes you want to make. This text guides the inpainting process, allowing you to specify features, styles, or modifications for the masked area.
	Prompt string `json:"prompt,omitempty"`
	// Number of steps for the image generation process.
	// Min: 15, Max: 50, Default: 50.
	Steps int `json:"steps,omitempty"`
	// Whether to perform upsampling on the prompt. If active, automatically modifies the prompt for more creative generation.
	// Default: false.
	PromptUpsampling bool `json:"prompt_upsampling,omitempty"`
	// Optional seed for reproducibility.
	Seed int `json:"seed,omitempty"`
	// Guidance strength for the image generation process.
	// Min: 1.5, Max: 100, Default: 60.
	Guidance float64 `json:"guidance,omitempty"`
	// Output format for the generated image. Can be 'jpeg' or 'png'.
	// Default: jpeg.
	OutputFormat string `json:"output_format,omitempty"`
	// Tolerance level for input and output moderation. Between 0 and 6, 0 being most strict, 6 being least strict.
	// Min: 0, Max: 6, Default: 2.
	SafetyTolerance int `json:"safety_tolerance"`
	// URL to receive webhook notifications.
	// Min length: 1, Max length: 2083.
	WebhookURL string `json:"webhook_url,omitempty"`
	// Optional secret for webhook signature verification.
	WebhookSecret string `json:"webhook_secret,omitempty"`
}

func (flx *FluxProFillGenerate) GenerateTaskMarker() {}

func (flx *FluxProFillGenerate) GetActionURL(baseURL string) string {
	return fmt.Sprintf("%s/v1/flux-pro-1.0-fill", baseURL)
}

// Task parameters for generating an image with Flux Pro 1.0 Canny through the BFL API.
type FluxProCannyGenerate struct {
	// Text prompt for image generation.
	Prompt string `json:"prompt"`
	// Base64 encoded image to use as control input if no preprocessed image is provided.
	ControlImage string `json:"control_image,omitempty"`
	// Optional pre-processed image that will bypass the control preprocessing step.
	PreprocessedImage string `json:"preprocessed_image,omitempty"`
	// Low threshold for Canny edge detection.
	// Min: 0, Max: 500, Default: 50.
	CannyLowThreshold int `json:"canny_low_threshold,omitempty"`
	// High threshold for Canny edge detection.
	// Min: 0, Max: 500, Default: 200.
	CannyHighThreshold int `json:"canny_high_threshold,omitempty"`
	// Whether to perform upsampling on the prompt.
	// Default: false.
	PromptUpsampling bool `json:"prompt_upsampling,omitempty"`
	// Optional seed for reproducibility.
	Seed int `json:"seed,omitempty"`
	// Number of steps for the image generation process.
	// Min: 15, Max: 50, Default: 50.
	Steps int `json:"steps,omitempty"`
	// Output format for the generated image. Can be 'jpeg' or 'png'.
	// Default: jpeg.
	OutputFormat string `json:"output_format,omitempty"`
	// Guidance strength for the image generation process.
	// Min: 1, Max: 100, Default: 30.
	Guidance float64 `json:"guidance,omitempty"`
	// Tolerance level for input and output moderation. Between 0 and 6, 0 being most strict, 6 being least strict.
	// Min: 0, Max: 6, Default: 2.
	SafetyTolerance int `json:"safety_tolerance"`
	// URL to receive webhook notifications.
	// Min length: 1, Max length: 2083.
	WebhookURL string `json:"webhook_url,omitempty"`
	// Optional secret for webhook signature verification.
	WebhookSecret string `json:"webhook_secret,omitempty"`
}

func (flx *FluxProCannyGenerate) GenerateTaskMarker() {}

func (flx *FluxProCannyGenerate) GetActionURL(baseURL string) string {
	return fmt.Sprintf("%s/v1/flux-pro-1.0-canny", baseURL)
}

// Task parameters for generating an image with Flux Pro 1.0 Depth through the BFL API.
type FluxProDepthGenerate struct {
	// Text prompt for image generation.
	Prompt string `json:"prompt"`
	// Base64 encoded image to use as control input.
	ControlImage string `json:"control_image,omitempty"`
	// Optional pre-processed image that will bypass the control preprocessing step.
	PreprocessedImage string `json:"preprocessed_image,omitempty"`
	// Whether to perform upsampling on the prompt.
	// Default: false.
	PromptUpsampling bool `json:"prompt_upsampling,omitempty"`
	// Optional seed for reproducibility.
	Seed int `json:"seed,omitempty"`
	// Number of steps for the image generation process.
	// Min: 15, Max: 50, Default: 50.
	Steps int `json:"steps,omitempty"`
	// Output format for the generated image. Can be 'jpeg' or 'png'.
	// Default: jpeg.
	OutputFormat string `json:"output_format,omitempty"`
	// Guidance strength for the image generation process.
	// Min: 1, Max: 100, Default: 15.
	Guidance float64 `json:"guidance,omitempty"`
	// Tolerance level for input and output moderation. Between 0 and 6, 0 being most strict, 6 being least strict.
	// Min: 0, Max: 6, Default: 2.
	SafetyTolerance int `json:"safety_tolerance"`
	// URL to receive webhook notifications.
	// Min length: 1, Max length: 2083.
	WebhookURL string `json:"webhook_url,omitempty"`
	// Optional secret for webhook signature verification.
	WebhookSecret string `json:"webhook_secret,omitempty"`
}

func (flx *FluxProDepthGenerate) GenerateTaskMarker() {}

func (flx *FluxProDepthGenerate) GetActionURL(baseURL string) string {
	return fmt.Sprintf("%s/v1/flux-pro-1.0-depth", baseURL)
}

// Task parameters for generating an image with Flux Pro Finetuned through the BFL API.
type FluxProFinetunedGenerate struct {
	// ID of the fine-tuned model you want to use.
	FinetuneID string `json:"finetune_id"`
	// Strength of the fine-tuned model. 0.0 means no influence, 1.0 means full influence. Allowed values up to 2.0.
	// Min: 0, Max: 2, Default: 1.1.
	FinetuneStrength float64 `json:"finetune_strength"`
	// Number of steps for the fine-tuning process.
	// Min: 1, Max: 50, Default: 40.
	Steps int `json:"steps"`
	// Guidance scale for image generation. High guidance scales improve prompt adherence at the cost of reduced realism.
	// Min: 1.5, Max: 5, Default: 2.5.
	Guidance float64 `json:"guidance"`
	// Text prompt for image generation.
	Prompt string `json:"prompt"`
	// Optional base64 encoded image to use with Flux Redux.
	ImagePrompt string `json:"image_prompt,omitempty"`
	// Width of the generated image in pixels. Must be a multiple of 32.
	// Min: 256, Max: 1440, Default: 1024.
	Width int `json:"width"`
	// Height of the generated image in pixels. Must be a multiple of 32.
	// Min: 256, Max: 1440, Default: 768.
	Height int `json:"height"`
	// Whether to perform upsampling on the prompt. If active, automatically modifies the prompt for more creative generation.
	// Default: false.
	PromptUpsampling bool `json:"prompt_upsampling"`
	// Optional seed for reproducibility.
	Seed int `json:"seed,omitempty"`
	// Tolerance level for input and output moderation. Between 0 and 6, 0 being most strict, 6 being least strict.
	// Min: 0, Max: 6, Default: 2.
	SafetyTolerance int `json:"safety_tolerance"`
	// Output format for the generated image. Can be 'jpeg' or 'png'.
	// Default: jpeg.
	OutputFormat string `json:"output_format"`
	// URL to receive webhook notifications.
	// Min length: 1, Max length: 2083.
	WebhookURL string `json:"webhook_url,omitempty"`
	// Optional secret for webhook signature verification.
	WebhookSecret string `json:"webhook_secret,omitempty"`
}

func (flx *FluxProFinetunedGenerate) GenerateTaskMarker() {}

func (flx *FluxProFinetunedGenerate) GetActionURL(baseURL string) string {
	return fmt.Sprintf("%s/v1/flux-pro-finetuned", baseURL)
}

// Task parameters for generating an image with Flux Pro 1.0 Depth Finetuned through the BFL API.
type FluxProDepthFinetunedGenerate struct {
	// ID of the fine-tuned model you want to use.
	FinetuneID string `json:"finetune_id"`
	// Strength of the fine-tuned model. 0.0 means no influence, 1.0 means full influence. Allowed values up to 2.0.
	// Min: 0, Max: 2, Default: 1.1.
	FinetuneStrength float64 `json:"finetune_strength"`
	// Text prompt for image generation.
	Prompt string `json:"prompt"`
	// Base64 encoded image to use as control input.
	ControlImage string `json:"control_image"`
	// Whether to perform upsampling on the prompt.
	// Default: false.
	PromptUpsampling bool `json:"prompt_upsampling"`
	// Optional seed for reproducibility.
	Seed int `json:"seed,omitempty"`
	// Number of steps for the image generation process.
	// Min: 15, Max: 50, Default: 50.
	Steps int `json:"steps"`
	// Output format for the generated image. Can be 'jpeg' or 'png'.
	// Default: jpeg.
	OutputFormat string `json:"output_format,omitempty"`
	// Guidance strength for the image generation process.
	// Min: 1, Max: 100, Default: 15.
	Guidance float64 `json:"guidance"`
	// Tolerance level for input and output moderation. Between 0 and 6, 0 being most strict, 6 being least strict.
	// Min: 0, Max: 6, Default: 2.
	SafetyTolerance int `json:"safety_tolerance"`
	// URL to receive webhook notifications.
	// Min length: 1, Max length: 2083.
	WebhookURL string `json:"webhook_url,omitempty"`
	// Optional secret for webhook signature verification.
	WebhookSecret string `json:"webhook_secret,omitempty"`
}

func (flx *FluxProDepthFinetunedGenerate) GenerateTaskMarker() {}

func (flx *FluxProDepthFinetunedGenerate) GetActionURL(baseURL string) string {
	return fmt.Sprintf("%s/v1/flux-pro-1.0-depth-finetuned", baseURL)
}

// Task parameters for generating an image with Flux Pro 1.0 Canny Finetuned through the BFL API.
type FluxProCannyFinetunedGenerate struct {
	// ID of the fine-tuned model you want to use.
	FinetuneID string `json:"finetune_id"`
	// Strength of the fine-tuned model. 0.0 means no influence, 1.0 means full influence. Allowed values up to 2.0.
	// Min: 0, Max: 2, Default: 1.1.
	FinetuneStrength float64 `json:"finetune_strength"`
	// Text prompt for image generation.
	Prompt string `json:"prompt"`
	// Base64 encoded image to use as control input if no preprocessed image is provided.
	ControlImage string `json:"control_image,omitempty"`
	// Optional pre-processed image that will bypass the control preprocessing step.
	PreprocessedImage string `json:"preprocessed_image,omitempty"`
	// Low threshold for Canny edge detection.
	// Min: 0, Max: 500, Default: 50.
	CannyLowThreshold int `json:"canny_low_threshold,omitempty"`
	// High threshold for Canny edge detection.
	// Min: 0, Max: 500, Default: 200.
	CannyHighThreshold int `json:"canny_high_threshold,omitempty"`
	// Whether to perform upsampling on the prompt.
	// Default: false.
	PromptUpsampling bool `json:"prompt_upsampling,omitempty"`
	// Optional seed for reproducibility.
	Seed int `json:"seed,omitempty"`
	// Number of steps for the image generation process.
	// Min: 15, Max: 50, Default: 50.
	Steps int `json:"steps,omitempty"`
	// Output format for the generated image. Can be 'jpeg' or 'png'.
	// Default: jpeg.
	OutputFormat string `json:"output_format,omitempty"`
	// Guidance strength for the image generation process.
	// Min: 1, Max: 100, Default: 30.
	Guidance float64 `json:"guidance,omitempty"`
	// Tolerance level for input and output moderation. Between 0 and 6, 0 being most strict, 6 being least strict.
	// Min: 0, Max: 6, Default: 2.
	SafetyTolerance int `json:"safety_tolerance"`
	// URL to receive webhook notifications.
	// Min length: 1, Max length: 2083.
	WebhookURL string `json:"webhook_url,omitempty"`
	// Optional secret for webhook signature verification.
	WebhookSecret string `json:"webhook_secret,omitempty"`
}

func (flx *FluxProCannyFinetunedGenerate) GenerateTaskMarker() {}

func (flx *FluxProCannyFinetunedGenerate) GetActionURL(baseURL string) string {
	return fmt.Sprintf("%s/v1/flux-pro-1.0-canny-finetuned", baseURL)
}

// Task parameters for generating an image with Flux Pro 1.0 Fill Finetuned through the BFL API.
type FluxProFillFinetunedGenerate struct {
	// ID of the fine-tuned model you want to use.
	FinetuneID string `json:"finetune_id"`
	// Strength of the fine-tuned model. 0.0 means no influence, 1.0 means full influence. Allowed values up to 2.0.
	// Min: 0, Max: 2, Default: 1.1.
	FinetuneStrength float64 `json:"finetune_strength"`
	// A Base64-encoded string representing the image you wish to modify. Can contain alpha mask if desired.
	Image string `json:"image"`
	// A Base64-encoded string representing a mask for the areas you want to modify in the image. The mask should be the same dimensions as the image and in black and white. Black areas (0%) indicate no modification, while white areas (100%) specify areas for inpainting. Optional if you provide an alpha mask in the original image.
	Mask string `json:"mask,omitempty"`
	// The description of the changes you want to make. This text guides the inpainting process, allowing you to specify features, styles, or modifications for the masked area.
	Prompt string `json:"prompt,omitempty"`
	// Number of steps for the image generation process.
	// Min: 15, Max: 50, Default: 50.
	Steps int `json:"steps,omitempty"`
	// Whether to perform upsampling on the prompt. If active, automatically modifies the prompt for more creative generation.
	// Default: false.
	PromptUpsampling bool `json:"prompt_upsampling,omitempty"`
	// Optional seed for reproducibility.
	Seed int `json:"seed,omitempty"`
	// Guidance strength for the image generation process.
	// Min: 1.5, Max: 100, Default: 60.
	Guidance float64 `json:"guidance,omitempty"`
	// Output format for the generated image. Can be 'jpeg' or 'png'.
	// Default: jpeg.
	OutputFormat string `json:"output_format,omitempty"`
	// Tolerance level for input and output moderation. Between 0 and 6, 0 being most strict, 6 being least strict.
	// Min: 0, Max: 6, Default: 2.
	SafetyTolerance int `json:"safety_tolerance"`
	// URL to receive webhook notifications.
	// Min length: 1, Max length: 2083.
	WebhookURL string `json:"webhook_url,omitempty"`
	// Optional secret for webhook signature verification.
	WebhookSecret string `json:"webhook_secret,omitempty"`
}

func (flx *FluxProFillFinetunedGenerate) GenerateTaskMarker() {}

func (flx *FluxProFillFinetunedGenerate) GetActionURL(baseURL string) string {
	return fmt.Sprintf("%s/v1/flux-pro-1.0-fill-finetuned", baseURL)
}

// Task parameters for generating an image with Flux Pro 1.1 Ultra Finetuned through the BFL API.
type FluxPro11UltraFinetunedGenerate struct {
	// ID of the fine-tuned model you want to use.
	FinetuneID string `json:"finetune_id"`
	// Strength of the fine-tuned model. 0.0 means no influence, 1.0 means full influence. Allowed values up to 2.0.
	// Min: 0, Max: 2, Default: 1.1.
	FinetuneStrength float64 `json:"finetune_strength"`
	// The prompt to use for image generation.
	Prompt string `json:"prompt,omitempty"`
	// Whether to perform upsampling on the prompt. If active, automatically modifies the prompt for more creative generation.
	// Default: false.
	PromptUpsampling bool `json:"prompt_upsampling"`
	// Optional seed for reproducibility. If not provided, a random seed will be used.
	Seed int `json:"seed,omitempty"`
	// Aspect ratio of the image between 21:9 and 9:21.
	// Default: 16:9.
	AspectRatio string `json:"aspect_ratio"`
	// Tolerance level for input and output moderation. Between 0 and 6, 0 being most strict, 6 being least strict.
	// Min: 0, Max: 6, Default: 2.
	SafetyTolerance int `json:"safety_tolerance"`
	// Output format for the generated image. Can be 'jpeg' or 'png'.
	// Default: jpeg.
	OutputFormat string `json:"output_format,omitempty"`
	// Generate less processed, more natural-looking images.
	// Default: false.
	Raw bool `json:"raw"`
	// Optional image to remix in base64 format.
	ImagePrompt string `json:"image_prompt,omitempty"`
	// Blend between the prompt and the image prompt.
	// Min: 0, Max: 1, Default: 0.1.
	ImagePromptStrength float64 `json:"image_prompt_strength"`
	// URL to receive webhook notifications.
	// Min length: 1, Max length: 2083.
	WebhookURL string `json:"webhook_url,omitempty"`
	// Optional secret for webhook signature verification.
	WebhookSecret string `json:"webhook_secret,omitempty"`
}

func (flx *FluxPro11UltraFinetunedGenerate) GenerateTaskMarker() {}

func (flx *FluxPro11UltraFinetunedGenerate) GetActionURL(baseURL string) string {
	return fmt.Sprintf("%s/v1/flux-pro-1.1-ultra-finetuned", baseURL)
}
