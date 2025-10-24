package awsbedrockalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents identifiers for default prompt routers in Bedrock.
//
// These are pre-configured routers provided by AWS that route between
// different models in the same family for optimal performance and cost.
//
// Example:
//   // Create a prompt router for intelligent model selection
//   promptRouter := bedrock.PromptRouter_FromDefaultId(bedrock.DefaultPromptRouterIdentifier_ANTHROPIC_CLAUDE_V1(), jsii.String("us-east-1"))
//
//   // Use the prompt router with a prompt variant
//   variant := bedrock.PromptVariant_Text(&TextPromptVariantProps{
//   	VariantName: jsii.String("variant1"),
//   	PromptText: jsii.String("What is the capital of France?"),
//   	Model: promptRouter,
//   })
//
//   bedrock.NewPrompt(this, jsii.String("Prompt"), &PromptProps{
//   	PromptName: jsii.String("prompt-router-test"),
//   	Variants: []IPromptVariant{
//   		variant,
//   	},
//   })
//
// Experimental.
type DefaultPromptRouterIdentifier interface {
	// The unique identifier for this prompt router.
	// Experimental.
	PromptRouterId() *string
	// The foundation models that this router can route between.
	// Experimental.
	RoutingModels() *[]BedrockFoundationModel
}

// The jsii proxy struct for DefaultPromptRouterIdentifier
type jsiiProxy_DefaultPromptRouterIdentifier struct {
	_ byte // padding
}

func (j *jsiiProxy_DefaultPromptRouterIdentifier) PromptRouterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"promptRouterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultPromptRouterIdentifier) RoutingModels() *[]BedrockFoundationModel {
	var returns *[]BedrockFoundationModel
	_jsii_.Get(
		j,
		"routingModels",
		&returns,
	)
	return returns
}


func DefaultPromptRouterIdentifier_ANTHROPIC_CLAUDE_V1() DefaultPromptRouterIdentifier {
	_init_.Initialize()
	var returns DefaultPromptRouterIdentifier
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.DefaultPromptRouterIdentifier",
		"ANTHROPIC_CLAUDE_V1",
		&returns,
	)
	return returns
}

func DefaultPromptRouterIdentifier_META_LLAMA_3_1() DefaultPromptRouterIdentifier {
	_init_.Initialize()
	var returns DefaultPromptRouterIdentifier
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.DefaultPromptRouterIdentifier",
		"META_LLAMA_3_1",
		&returns,
	)
	return returns
}

