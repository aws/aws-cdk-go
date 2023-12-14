package awsbedrock

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The model identifiers for the Bedrock base foundation models.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   bedrock.FoundationModel_FromFoundationModelId(this, jsii.String("Model"), bedrock.FoundationModelIdentifier_ANTHROPIC_CLAUDE_V2())
//
// See: https://docs.aws.amazon.com/bedrock/latest/userguide/model-ids-arns.html
//
type FoundationModelIdentifier interface {
	// the model identifier.
	ModelId() *string
}

// The jsii proxy struct for FoundationModelIdentifier
type jsiiProxy_FoundationModelIdentifier struct {
	_ byte // padding
}

func (j *jsiiProxy_FoundationModelIdentifier) ModelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelId",
		&returns,
	)
	return returns
}


// Constructor for foundation model identifier.
func NewFoundationModelIdentifier(modelId *string) FoundationModelIdentifier {
	_init_.Initialize()

	if err := validateNewFoundationModelIdentifierParameters(modelId); err != nil {
		panic(err)
	}
	j := jsiiProxy_FoundationModelIdentifier{}

	_jsii_.Create(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		[]interface{}{modelId},
		&j,
	)

	return &j
}

// Constructor for foundation model identifier.
func NewFoundationModelIdentifier_Override(f FoundationModelIdentifier, modelId *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		[]interface{}{modelId},
		f,
	)
}

func FoundationModelIdentifier_AI21_LABS_JURASSIC_2_MID_V1() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"AI21_LABS_JURASSIC_2_MID_V1",
		&returns,
	)
	return returns
}

func FoundationModelIdentifier_AI21_LABS_JURASSIC_2_ULTRA_V1() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"AI21_LABS_JURASSIC_2_ULTRA_V1",
		&returns,
	)
	return returns
}

func FoundationModelIdentifier_AMAZON_TITAN_EMBEDDINGS_G1_TEXT_V1() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"AMAZON_TITAN_EMBEDDINGS_G1_TEXT_V1",
		&returns,
	)
	return returns
}

func FoundationModelIdentifier_AMAZON_TITAN_IMAGE_GENERATOR_G1_V1() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"AMAZON_TITAN_IMAGE_GENERATOR_G1_V1",
		&returns,
	)
	return returns
}

func FoundationModelIdentifier_AMAZON_TITAN_MULTIMODAL_EMBEDDINGS_G1_V1() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"AMAZON_TITAN_MULTIMODAL_EMBEDDINGS_G1_V1",
		&returns,
	)
	return returns
}

func FoundationModelIdentifier_AMAZON_TITAN_TEXT_G1_EXPRESS_V1() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"AMAZON_TITAN_TEXT_G1_EXPRESS_V1",
		&returns,
	)
	return returns
}

func FoundationModelIdentifier_ANTHROPIC_CLAUDE_INSTANT_V1() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"ANTHROPIC_CLAUDE_INSTANT_V1",
		&returns,
	)
	return returns
}

func FoundationModelIdentifier_ANTHROPIC_CLAUDE_V1() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"ANTHROPIC_CLAUDE_V1",
		&returns,
	)
	return returns
}

func FoundationModelIdentifier_ANTHROPIC_CLAUDE_V2() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"ANTHROPIC_CLAUDE_V2",
		&returns,
	)
	return returns
}

func FoundationModelIdentifier_ANTHROPIC_CLAUDE_V2_1() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"ANTHROPIC_CLAUDE_V2_1",
		&returns,
	)
	return returns
}

func FoundationModelIdentifier_COHERE_COMMAND_LIGHT_V14() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"COHERE_COMMAND_LIGHT_V14",
		&returns,
	)
	return returns
}

func FoundationModelIdentifier_COHERE_COMMAND_V14() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"COHERE_COMMAND_V14",
		&returns,
	)
	return returns
}

func FoundationModelIdentifier_COHERE_EMBED_ENGLISH_V3() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"COHERE_EMBED_ENGLISH_V3",
		&returns,
	)
	return returns
}

func FoundationModelIdentifier_COHERE_EMBED_MULTILINGUAL_V3() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"COHERE_EMBED_MULTILINGUAL_V3",
		&returns,
	)
	return returns
}

func FoundationModelIdentifier_META_LLAMA_2_CHAT_13B_V1() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"META_LLAMA_2_CHAT_13B_V1",
		&returns,
	)
	return returns
}

func FoundationModelIdentifier_META_LLAMA_2_CHAT_70B_V1() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"META_LLAMA_2_CHAT_70B_V1",
		&returns,
	)
	return returns
}

