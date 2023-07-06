//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azopenai

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

// MarshalJSON implements the json.Marshaller interface for type azureCoreFoundationsError.
func (a azureCoreFoundationsError) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "code", a.Code)
	populate(objectMap, "details", a.Details)
	populate(objectMap, "innererror", a.Innererror)
	populate(objectMap, "message", a.Message)
	populate(objectMap, "target", a.Target)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type azureCoreFoundationsError.
func (a *azureCoreFoundationsError) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", a, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "code":
			err = unpopulate(val, "Code", &a.Code)
			delete(rawMsg, key)
		case "details":
			err = unpopulate(val, "Details", &a.Details)
			delete(rawMsg, key)
		case "innererror":
			err = unpopulate(val, "Innererror", &a.Innererror)
			delete(rawMsg, key)
		case "message":
			err = unpopulate(val, "Message", &a.Message)
			delete(rawMsg, key)
		case "target":
			err = unpopulate(val, "Target", &a.Target)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", a, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type azureCoreFoundationsErrorInnererror.
func (a azureCoreFoundationsErrorInnererror) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "code", a.Code)
	populate(objectMap, "innererror", a.Innererror)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type azureCoreFoundationsErrorInnererror.
func (a *azureCoreFoundationsErrorInnererror) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", a, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "code":
			err = unpopulate(val, "Code", &a.Code)
			delete(rawMsg, key)
		case "innererror":
			err = unpopulate(val, "Innererror", &a.Innererror)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", a, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type azureCoreFoundationsErrorResponse.
func (a azureCoreFoundationsErrorResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "error", a.Error)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type azureCoreFoundationsErrorResponse.
func (a *azureCoreFoundationsErrorResponse) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", a, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "error":
			err = unpopulate(val, "Error", &a.Error)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", a, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type azureCoreFoundationsErrorResponseError.
func (a azureCoreFoundationsErrorResponseError) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "code", a.Code)
	populate(objectMap, "details", a.Details)
	populate(objectMap, "innererror", a.Innererror)
	populate(objectMap, "message", a.Message)
	populate(objectMap, "target", a.Target)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type azureCoreFoundationsErrorResponseError.
func (a *azureCoreFoundationsErrorResponseError) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", a, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "code":
			err = unpopulate(val, "Code", &a.Code)
			delete(rawMsg, key)
		case "details":
			err = unpopulate(val, "Details", &a.Details)
			delete(rawMsg, key)
		case "innererror":
			err = unpopulate(val, "Innererror", &a.Innererror)
			delete(rawMsg, key)
		case "message":
			err = unpopulate(val, "Message", &a.Message)
			delete(rawMsg, key)
		case "target":
			err = unpopulate(val, "Target", &a.Target)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", a, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type azureCoreFoundationsInnerError.
func (a azureCoreFoundationsInnerError) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "code", a.Code)
	populate(objectMap, "innererror", a.Innererror)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type azureCoreFoundationsInnerError.
func (a *azureCoreFoundationsInnerError) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", a, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "code":
			err = unpopulate(val, "Code", &a.Code)
			delete(rawMsg, key)
		case "innererror":
			err = unpopulate(val, "Innererror", &a.Innererror)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", a, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type azureCoreFoundationsInnerErrorInnererror.
func (a azureCoreFoundationsInnerErrorInnererror) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "code", a.Code)
	populate(objectMap, "innererror", a.Innererror)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type azureCoreFoundationsInnerErrorInnererror.
func (a *azureCoreFoundationsInnerErrorInnererror) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", a, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "code":
			err = unpopulate(val, "Code", &a.Code)
			delete(rawMsg, key)
		case "innererror":
			err = unpopulate(val, "Innererror", &a.Innererror)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", a, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type batchImageGenerationOperationResponse.
func (b batchImageGenerationOperationResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "created", b.Created)
	populate(objectMap, "error", b.Error)
	populate(objectMap, "expires", b.Expires)
	populate(objectMap, "id", b.ID)
	populate(objectMap, "result", b.Result)
	populate(objectMap, "status", b.Status)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type batchImageGenerationOperationResponse.
func (b *batchImageGenerationOperationResponse) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", b, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "created":
			err = unpopulate(val, "Created", &b.Created)
			delete(rawMsg, key)
		case "error":
			err = unpopulate(val, "Error", &b.Error)
			delete(rawMsg, key)
		case "expires":
			err = unpopulate(val, "Expires", &b.Expires)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, "ID", &b.ID)
			delete(rawMsg, key)
		case "result":
			err = unpopulate(val, "Result", &b.Result)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, "Status", &b.Status)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", b, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ChatChoice.
func (c ChatChoice) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "delta", c.Delta)
	populate(objectMap, "finish_reason", c.FinishReason)
	populate(objectMap, "index", c.Index)
	populate(objectMap, "message", c.Message)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ChatChoice.
func (c *ChatChoice) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "delta":
			err = unpopulate(val, "Delta", &c.Delta)
			delete(rawMsg, key)
		case "finish_reason":
			err = unpopulate(val, "FinishReason", &c.FinishReason)
			delete(rawMsg, key)
		case "index":
			err = unpopulate(val, "Index", &c.Index)
			delete(rawMsg, key)
		case "message":
			err = unpopulate(val, "Message", &c.Message)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ChatChoiceDelta.
func (c ChatChoiceDelta) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "content", c.Content)
	populate(objectMap, "role", c.Role)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ChatChoiceDelta.
func (c *ChatChoiceDelta) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "content":
			err = unpopulate(val, "Content", &c.Content)
			delete(rawMsg, key)
		case "role":
			err = unpopulate(val, "Role", &c.Role)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ChatChoiceMessage.
func (c ChatChoiceMessage) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "content", c.Content)
	populate(objectMap, "role", c.Role)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ChatChoiceMessage.
func (c *ChatChoiceMessage) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "content":
			err = unpopulate(val, "Content", &c.Content)
			delete(rawMsg, key)
		case "role":
			err = unpopulate(val, "Role", &c.Role)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ChatCompletions.
func (c ChatCompletions) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "choices", c.Choices)
	populate(objectMap, "created", c.Created)
	populate(objectMap, "id", c.ID)
	populate(objectMap, "usage", c.Usage)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ChatCompletions.
func (c *ChatCompletions) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "choices":
			err = unpopulate(val, "Choices", &c.Choices)
			delete(rawMsg, key)
		case "created":
			err = unpopulate(val, "Created", &c.Created)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, "ID", &c.ID)
			delete(rawMsg, key)
		case "usage":
			err = unpopulate(val, "Usage", &c.Usage)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ChatCompletionsOptions.
func (c ChatCompletionsOptions) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "frequency_penalty", c.FrequencyPenalty)
	populate(objectMap, "logit_bias", c.LogitBias)
	populate(objectMap, "max_tokens", c.MaxTokens)
	populate(objectMap, "messages", c.Messages)
	populate(objectMap, "model", c.Model)
	populate(objectMap, "n", c.N)
	populate(objectMap, "presence_penalty", c.PresencePenalty)
	populate(objectMap, "stop", c.Stop)
	populate(objectMap, "temperature", c.Temperature)
	populate(objectMap, "top_p", c.TopP)
	populate(objectMap, "user", c.User)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ChatCompletionsOptions.
func (c *ChatCompletionsOptions) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "frequency_penalty":
			err = unpopulate(val, "FrequencyPenalty", &c.FrequencyPenalty)
			delete(rawMsg, key)
		case "logit_bias":
			err = unpopulate(val, "LogitBias", &c.LogitBias)
			delete(rawMsg, key)
		case "max_tokens":
			err = unpopulate(val, "MaxTokens", &c.MaxTokens)
			delete(rawMsg, key)
		case "messages":
			err = unpopulate(val, "Messages", &c.Messages)
			delete(rawMsg, key)
		case "model":
			err = unpopulate(val, "Model", &c.Model)
			delete(rawMsg, key)
		case "n":
			err = unpopulate(val, "N", &c.N)
			delete(rawMsg, key)
		case "presence_penalty":
			err = unpopulate(val, "PresencePenalty", &c.PresencePenalty)
			delete(rawMsg, key)
		case "stop":
			err = unpopulate(val, "Stop", &c.Stop)
			delete(rawMsg, key)
		case "temperature":
			err = unpopulate(val, "Temperature", &c.Temperature)
			delete(rawMsg, key)
		case "top_p":
			err = unpopulate(val, "TopP", &c.TopP)
			delete(rawMsg, key)
		case "user":
			err = unpopulate(val, "User", &c.User)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ChatMessage.
func (c ChatMessage) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "content", c.Content)
	populate(objectMap, "role", c.Role)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ChatMessage.
func (c *ChatMessage) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "content":
			err = unpopulate(val, "Content", &c.Content)
			delete(rawMsg, key)
		case "role":
			err = unpopulate(val, "Role", &c.Role)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Choice.
func (c Choice) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "finish_reason", c.FinishReason)
	populate(objectMap, "index", c.Index)
	populate(objectMap, "logprobs", c.LogProbs)
	populate(objectMap, "text", c.Text)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Choice.
func (c *Choice) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "finish_reason":
			err = unpopulate(val, "FinishReason", &c.FinishReason)
			delete(rawMsg, key)
		case "index":
			err = unpopulate(val, "Index", &c.Index)
			delete(rawMsg, key)
		case "logprobs":
			err = unpopulate(val, "LogProbs", &c.LogProbs)
			delete(rawMsg, key)
		case "text":
			err = unpopulate(val, "Text", &c.Text)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ChoiceLogProbs.
func (c ChoiceLogProbs) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "text_offset", c.TextOffset)
	populate(objectMap, "token_logprobs", c.TokenLogProbs)
	populate(objectMap, "tokens", c.Tokens)
	populate(objectMap, "top_logprobs", c.TopLogProbs)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ChoiceLogProbs.
func (c *ChoiceLogProbs) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "text_offset":
			err = unpopulate(val, "TextOffset", &c.TextOffset)
			delete(rawMsg, key)
		case "token_logprobs":
			err = unpopulate(val, "TokenLogProbs", &c.TokenLogProbs)
			delete(rawMsg, key)
		case "tokens":
			err = unpopulate(val, "Tokens", &c.Tokens)
			delete(rawMsg, key)
		case "top_logprobs":
			err = unpopulate(val, "TopLogProbs", &c.TopLogProbs)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Completions.
func (c Completions) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "choices", c.Choices)
	populate(objectMap, "created", c.Created)
	populate(objectMap, "id", c.ID)
	populate(objectMap, "usage", c.Usage)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Completions.
func (c *Completions) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "choices":
			err = unpopulate(val, "Choices", &c.Choices)
			delete(rawMsg, key)
		case "created":
			err = unpopulate(val, "Created", &c.Created)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, "ID", &c.ID)
			delete(rawMsg, key)
		case "usage":
			err = unpopulate(val, "Usage", &c.Usage)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type CompletionsLogProbabilityModel.
func (c CompletionsLogProbabilityModel) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "text_offset", c.TextOffset)
	populate(objectMap, "token_logprobs", c.TokenLogProbs)
	populate(objectMap, "tokens", c.Tokens)
	populate(objectMap, "top_logprobs", c.TopLogProbs)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type CompletionsLogProbabilityModel.
func (c *CompletionsLogProbabilityModel) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "text_offset":
			err = unpopulate(val, "TextOffset", &c.TextOffset)
			delete(rawMsg, key)
		case "token_logprobs":
			err = unpopulate(val, "TokenLogProbs", &c.TokenLogProbs)
			delete(rawMsg, key)
		case "tokens":
			err = unpopulate(val, "Tokens", &c.Tokens)
			delete(rawMsg, key)
		case "top_logprobs":
			err = unpopulate(val, "TopLogProbs", &c.TopLogProbs)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type CompletionsOptions.
func (c CompletionsOptions) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "best_of", c.BestOf)
	populate(objectMap, "echo", c.Echo)
	populate(objectMap, "frequency_penalty", c.FrequencyPenalty)
	populate(objectMap, "logit_bias", c.LogitBias)
	populate(objectMap, "logprobs", c.LogProbs)
	populate(objectMap, "max_tokens", c.MaxTokens)
	populate(objectMap, "model", c.Model)
	populate(objectMap, "n", c.N)
	populate(objectMap, "presence_penalty", c.PresencePenalty)
	populate(objectMap, "prompt", c.Prompt)
	populate(objectMap, "stop", c.Stop)
	populate(objectMap, "temperature", c.Temperature)
	populate(objectMap, "top_p", c.TopP)
	populate(objectMap, "user", c.User)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type CompletionsOptions.
func (c *CompletionsOptions) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "best_of":
			err = unpopulate(val, "BestOf", &c.BestOf)
			delete(rawMsg, key)
		case "echo":
			err = unpopulate(val, "Echo", &c.Echo)
			delete(rawMsg, key)
		case "frequency_penalty":
			err = unpopulate(val, "FrequencyPenalty", &c.FrequencyPenalty)
			delete(rawMsg, key)
		case "logit_bias":
			err = unpopulate(val, "LogitBias", &c.LogitBias)
			delete(rawMsg, key)
		case "logprobs":
			err = unpopulate(val, "LogProbs", &c.LogProbs)
			delete(rawMsg, key)
		case "max_tokens":
			err = unpopulate(val, "MaxTokens", &c.MaxTokens)
			delete(rawMsg, key)
		case "model":
			err = unpopulate(val, "Model", &c.Model)
			delete(rawMsg, key)
		case "n":
			err = unpopulate(val, "N", &c.N)
			delete(rawMsg, key)
		case "presence_penalty":
			err = unpopulate(val, "PresencePenalty", &c.PresencePenalty)
			delete(rawMsg, key)
		case "prompt":
			err = unpopulate(val, "Prompt", &c.Prompt)
			delete(rawMsg, key)
		case "stop":
			err = unpopulate(val, "Stop", &c.Stop)
			delete(rawMsg, key)
		case "temperature":
			err = unpopulate(val, "Temperature", &c.Temperature)
			delete(rawMsg, key)
		case "top_p":
			err = unpopulate(val, "TopP", &c.TopP)
			delete(rawMsg, key)
		case "user":
			err = unpopulate(val, "User", &c.User)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type CompletionsUsage.
func (c CompletionsUsage) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "completion_tokens", c.CompletionTokens)
	populate(objectMap, "prompt_tokens", c.PromptTokens)
	populate(objectMap, "total_tokens", c.TotalTokens)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type CompletionsUsage.
func (c *CompletionsUsage) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "completion_tokens":
			err = unpopulate(val, "CompletionTokens", &c.CompletionTokens)
			delete(rawMsg, key)
		case "prompt_tokens":
			err = unpopulate(val, "PromptTokens", &c.PromptTokens)
			delete(rawMsg, key)
		case "total_tokens":
			err = unpopulate(val, "TotalTokens", &c.TotalTokens)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Deployment.
func (d Deployment) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "deploymentId", d.DeploymentID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Deployment.
func (d *Deployment) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "deploymentId":
			err = unpopulate(val, "DeploymentID", &d.DeploymentID)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type EmbeddingItem.
func (e EmbeddingItem) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "embedding", e.Embedding)
	populate(objectMap, "index", e.Index)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type EmbeddingItem.
func (e *EmbeddingItem) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "embedding":
			err = unpopulate(val, "Embedding", &e.Embedding)
			delete(rawMsg, key)
		case "index":
			err = unpopulate(val, "Index", &e.Index)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Embeddings.
func (e Embeddings) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "data", e.Data)
	populate(objectMap, "usage", e.Usage)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Embeddings.
func (e *Embeddings) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "data":
			err = unpopulate(val, "Data", &e.Data)
			delete(rawMsg, key)
		case "usage":
			err = unpopulate(val, "Usage", &e.Usage)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type EmbeddingsOptions.
func (e EmbeddingsOptions) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "input", e.Input)
	populate(objectMap, "model", e.Model)
	populate(objectMap, "user", e.User)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type EmbeddingsOptions.
func (e *EmbeddingsOptions) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "input":
			err = unpopulate(val, "Input", &e.Input)
			delete(rawMsg, key)
		case "model":
			err = unpopulate(val, "Model", &e.Model)
			delete(rawMsg, key)
		case "user":
			err = unpopulate(val, "User", &e.User)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type EmbeddingsUsage.
func (e EmbeddingsUsage) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "prompt_tokens", e.PromptTokens)
	populate(objectMap, "total_tokens", e.TotalTokens)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type EmbeddingsUsage.
func (e *EmbeddingsUsage) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "prompt_tokens":
			err = unpopulate(val, "PromptTokens", &e.PromptTokens)
			delete(rawMsg, key)
		case "total_tokens":
			err = unpopulate(val, "TotalTokens", &e.TotalTokens)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type EmbeddingsUsageAutoGenerated.
func (e EmbeddingsUsageAutoGenerated) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "prompt_tokens", e.PromptTokens)
	populate(objectMap, "total_tokens", e.TotalTokens)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type EmbeddingsUsageAutoGenerated.
func (e *EmbeddingsUsageAutoGenerated) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "prompt_tokens":
			err = unpopulate(val, "PromptTokens", &e.PromptTokens)
			delete(rawMsg, key)
		case "total_tokens":
			err = unpopulate(val, "TotalTokens", &e.TotalTokens)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ImageGenerationOptions.
func (i ImageGenerationOptions) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "n", i.N)
	populate(objectMap, "prompt", i.Prompt)
	populate(objectMap, "response_format", i.ResponseFormat)
	populate(objectMap, "size", i.Size)
	populate(objectMap, "user", i.User)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ImageGenerationOptions.
func (i *ImageGenerationOptions) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", i, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "n":
			err = unpopulate(val, "N", &i.N)
			delete(rawMsg, key)
		case "prompt":
			err = unpopulate(val, "Prompt", &i.Prompt)
			delete(rawMsg, key)
		case "response_format":
			err = unpopulate(val, "ResponseFormat", &i.ResponseFormat)
			delete(rawMsg, key)
		case "size":
			err = unpopulate(val, "Size", &i.Size)
			delete(rawMsg, key)
		case "user":
			err = unpopulate(val, "User", &i.User)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", i, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ImageGenerations.
func (i ImageGenerations) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "created", i.Created)
	populate(objectMap, "data", i.Data)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ImageGenerations.
func (i *ImageGenerations) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", i, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "created":
			err = unpopulate(val, "Created", &i.Created)
			delete(rawMsg, key)
		case "data":
			err = unpopulate(val, "Data", &i.Data)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", i, err)
		}
	}
	return nil
}

func populate(m map[string]any, k string, v any) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, fn string, v any) error {
	if data == nil {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	return nil
}
