package awsbedrockalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface representing a prompt variant configuration.
//
// Variants are specific sets of inputs that guide FMs on Amazon Bedrock to
// generate an appropriate response or output for a given task or instruction.
// Experimental.
type IPromptVariant interface {
	// The generative AI resource configuration.
	// Experimental.
	GenAiResource() PromptGenAiResource
	// The inference configuration.
	// Experimental.
	InferenceConfiguration() PromptInferenceConfiguration
	// The unique identifier of the model with which to run inference on the prompt.
	// Experimental.
	ModelId() *string
	// The name of the prompt variant.
	// Experimental.
	Name() *string
	// The template configuration.
	// Experimental.
	TemplateConfiguration() PromptTemplateConfiguration
	// The type of prompt template.
	// Default: - Text.
	//
	// Experimental.
	TemplateType() PromptTemplateType
}

// The jsii proxy for IPromptVariant
type jsiiProxy_IPromptVariant struct {
	_ byte // padding
}

func (j *jsiiProxy_IPromptVariant) GenAiResource() PromptGenAiResource {
	var returns PromptGenAiResource
	_jsii_.Get(
		j,
		"genAiResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPromptVariant) InferenceConfiguration() PromptInferenceConfiguration {
	var returns PromptInferenceConfiguration
	_jsii_.Get(
		j,
		"inferenceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPromptVariant) ModelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPromptVariant) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPromptVariant) TemplateConfiguration() PromptTemplateConfiguration {
	var returns PromptTemplateConfiguration
	_jsii_.Get(
		j,
		"templateConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPromptVariant) TemplateType() PromptTemplateType {
	var returns PromptTemplateType
	_jsii_.Get(
		j,
		"templateType",
		&returns,
	)
	return returns
}

