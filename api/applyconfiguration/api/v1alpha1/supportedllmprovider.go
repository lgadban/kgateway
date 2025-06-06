// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// SupportedLLMProviderApplyConfiguration represents a declarative configuration of the SupportedLLMProvider type for use
// with apply.
type SupportedLLMProviderApplyConfiguration struct {
	OpenAI      *OpenAIConfigApplyConfiguration      `json:"openai,omitempty"`
	AzureOpenAI *AzureOpenAIConfigApplyConfiguration `json:"azureopenai,omitempty"`
	Anthropic   *AnthropicConfigApplyConfiguration   `json:"anthropic,omitempty"`
	Gemini      *GeminiConfigApplyConfiguration      `json:"gemini,omitempty"`
	VertexAI    *VertexAIConfigApplyConfiguration    `json:"vertexai,omitempty"`
}

// SupportedLLMProviderApplyConfiguration constructs a declarative configuration of the SupportedLLMProvider type for use with
// apply.
func SupportedLLMProvider() *SupportedLLMProviderApplyConfiguration {
	return &SupportedLLMProviderApplyConfiguration{}
}

// WithOpenAI sets the OpenAI field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OpenAI field is set to the value of the last call.
func (b *SupportedLLMProviderApplyConfiguration) WithOpenAI(value *OpenAIConfigApplyConfiguration) *SupportedLLMProviderApplyConfiguration {
	b.OpenAI = value
	return b
}

// WithAzureOpenAI sets the AzureOpenAI field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AzureOpenAI field is set to the value of the last call.
func (b *SupportedLLMProviderApplyConfiguration) WithAzureOpenAI(value *AzureOpenAIConfigApplyConfiguration) *SupportedLLMProviderApplyConfiguration {
	b.AzureOpenAI = value
	return b
}

// WithAnthropic sets the Anthropic field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Anthropic field is set to the value of the last call.
func (b *SupportedLLMProviderApplyConfiguration) WithAnthropic(value *AnthropicConfigApplyConfiguration) *SupportedLLMProviderApplyConfiguration {
	b.Anthropic = value
	return b
}

// WithGemini sets the Gemini field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Gemini field is set to the value of the last call.
func (b *SupportedLLMProviderApplyConfiguration) WithGemini(value *GeminiConfigApplyConfiguration) *SupportedLLMProviderApplyConfiguration {
	b.Gemini = value
	return b
}

// WithVertexAI sets the VertexAI field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the VertexAI field is set to the value of the last call.
func (b *SupportedLLMProviderApplyConfiguration) WithVertexAI(value *VertexAIConfigApplyConfiguration) *SupportedLLMProviderApplyConfiguration {
	b.VertexAI = value
	return b
}
