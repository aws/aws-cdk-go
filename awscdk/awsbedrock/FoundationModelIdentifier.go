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
//   model := bedrock.FoundationModel_FromFoundationModelId(this, jsii.String("Model"), bedrock.FoundationModelIdentifier_AMAZON_TITAN_TEXT_G1_EXPRESS_V1())
//
//   task := tasks.NewBedrockInvokeModel(this, jsii.String("Prompt Model"), &BedrockInvokeModelProps{
//   	Model: Model,
//   	Body: sfn.TaskInput_FromObject(map[string]interface{}{
//   		"inputText": jsii.String("Generate a list of five first names."),
//   		"textGenerationConfig": map[string]*f64{
//   			"maxTokenCount": jsii.Number(100),
//   			"temperature": jsii.Number(1),
//   		},
//   	}),
//   	ResultSelector: map[string]interface{}{
//   		"names": sfn.JsonPath_stringAt(jsii.String("$.Body.results[0].outputText")),
//   	},
//   })
//
// See: https://docs.aws.amazon.com/bedrock/latest/userguide/model-ids.html
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

func FoundationModelIdentifier_AI21_J2_GRANDE_INSTRUCT() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"AI21_J2_GRANDE_INSTRUCT",
		&returns,
	)
	return returns
}

func FoundationModelIdentifier_AI21_J2_JAMBA_INSTRUCT_V1_0() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"AI21_J2_JAMBA_INSTRUCT_V1_0",
		&returns,
	)
	return returns
}

func FoundationModelIdentifier_AI21_J2_JUMBO_INSTRUCT() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"AI21_J2_JUMBO_INSTRUCT",
		&returns,
	)
	return returns
}

func FoundationModelIdentifier_AI21_J2_MID() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"AI21_J2_MID",
		&returns,
	)
	return returns
}

func FoundationModelIdentifier_AI21_J2_ULTRA() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"AI21_J2_ULTRA",
		&returns,
	)
	return returns
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

func FoundationModelIdentifier_AI21_LABS_JURASSIC_2_ULTRA_V1_0_8K() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"AI21_LABS_JURASSIC_2_ULTRA_V1_0_8K",
		&returns,
	)
	return returns
}

func FoundationModelIdentifier_AMAZON_TITAN_EMBED_G1_TEXT_02() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"AMAZON_TITAN_EMBED_G1_TEXT_02",
		&returns,
	)
	return returns
}

func FoundationModelIdentifier_AMAZON_TITAN_EMBED_IMAGE_V1_0() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"AMAZON_TITAN_EMBED_IMAGE_V1_0",
		&returns,
	)
	return returns
}

func FoundationModelIdentifier_AMAZON_TITAN_EMBED_TEXT_V1_2_8K() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"AMAZON_TITAN_EMBED_TEXT_V1_2_8K",
		&returns,
	)
	return returns
}

func FoundationModelIdentifier_AMAZON_TITAN_EMBED_TEXT_V2_0() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"AMAZON_TITAN_EMBED_TEXT_V2_0",
		&returns,
	)
	return returns
}

func FoundationModelIdentifier_AMAZON_TITAN_EMBED_TEXT_V2_0_8K() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"AMAZON_TITAN_EMBED_TEXT_V2_0_8K",
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

func FoundationModelIdentifier_AMAZON_TITAN_IMAGE_GENERATOR_V1_0() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"AMAZON_TITAN_IMAGE_GENERATOR_V1_0",
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

func FoundationModelIdentifier_AMAZON_TITAN_TEXT_EXPRESS_V1_0_8K() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"AMAZON_TITAN_TEXT_EXPRESS_V1_0_8K",
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

func FoundationModelIdentifier_AMAZON_TITAN_TEXT_G1_LITE_V1() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"AMAZON_TITAN_TEXT_G1_LITE_V1",
		&returns,
	)
	return returns
}

func FoundationModelIdentifier_AMAZON_TITAN_TEXT_LITE_V1() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"AMAZON_TITAN_TEXT_LITE_V1",
		&returns,
	)
	return returns
}

func FoundationModelIdentifier_AMAZON_TITAN_TEXT_LITE_V1_0_4K() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"AMAZON_TITAN_TEXT_LITE_V1_0_4K",
		&returns,
	)
	return returns
}

func FoundationModelIdentifier_AMAZON_TITAN_TEXT_PREMIER_V1() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"AMAZON_TITAN_TEXT_PREMIER_V1",
		&returns,
	)
	return returns
}

func FoundationModelIdentifier_AMAZON_TITAN_TG1_LARGE() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"AMAZON_TITAN_TG1_LARGE",
		&returns,
	)
	return returns
}

func FoundationModelIdentifier_ANTHROPIC_CLAUDE_3_5_SONNET_20240620_V1_0() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"ANTHROPIC_CLAUDE_3_5_SONNET_20240620_V1_0",
		&returns,
	)
	return returns
}

func FoundationModelIdentifier_ANTHROPIC_CLAUDE_3_HAIKU_20240307_V1_0() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"ANTHROPIC_CLAUDE_3_HAIKU_20240307_V1_0",
		&returns,
	)
	return returns
}

func FoundationModelIdentifier_ANTHROPIC_CLAUDE_3_HAIKU_20240307_V1_0_200K() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"ANTHROPIC_CLAUDE_3_HAIKU_20240307_V1_0_200K",
		&returns,
	)
	return returns
}

func FoundationModelIdentifier_ANTHROPIC_CLAUDE_3_HAIKU_20240307_V1_0_48K() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"ANTHROPIC_CLAUDE_3_HAIKU_20240307_V1_0_48K",
		&returns,
	)
	return returns
}

func FoundationModelIdentifier_ANTHROPIC_CLAUDE_3_OPUS_20240229_V1_0() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"ANTHROPIC_CLAUDE_3_OPUS_20240229_V1_0",
		&returns,
	)
	return returns
}

func FoundationModelIdentifier_ANTHROPIC_CLAUDE_3_SONNET_20240229_V1_0() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"ANTHROPIC_CLAUDE_3_SONNET_20240229_V1_0",
		&returns,
	)
	return returns
}

func FoundationModelIdentifier_ANTHROPIC_CLAUDE_3_SONNET_20240229_V1_0_200K() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"ANTHROPIC_CLAUDE_3_SONNET_20240229_V1_0_200K",
		&returns,
	)
	return returns
}

func FoundationModelIdentifier_ANTHROPIC_CLAUDE_3_SONNET_20240229_V1_0_28K() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"ANTHROPIC_CLAUDE_3_SONNET_20240229_V1_0_28K",
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

func FoundationModelIdentifier_ANTHROPIC_CLAUDE_INSTANT_V1_2_100K() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"ANTHROPIC_CLAUDE_INSTANT_V1_2_100K",
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

func FoundationModelIdentifier_ANTHROPIC_CLAUDE_V2_0_100K() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"ANTHROPIC_CLAUDE_V2_0_100K",
		&returns,
	)
	return returns
}

func FoundationModelIdentifier_ANTHROPIC_CLAUDE_V2_0_18K() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"ANTHROPIC_CLAUDE_V2_0_18K",
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

func FoundationModelIdentifier_ANTHROPIC_CLAUDE_V2_1_18K() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"ANTHROPIC_CLAUDE_V2_1_18K",
		&returns,
	)
	return returns
}

func FoundationModelIdentifier_ANTHROPIC_CLAUDE_V2_1_200K() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"ANTHROPIC_CLAUDE_V2_1_200K",
		&returns,
	)
	return returns
}

func FoundationModelIdentifier_COHERE_COMMAND_LIGHT_TEXT_V14_7_4K() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"COHERE_COMMAND_LIGHT_TEXT_V14_7_4K",
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

func FoundationModelIdentifier_COHERE_COMMAND_R_PLUS_V1() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"COHERE_COMMAND_R_PLUS_V1",
		&returns,
	)
	return returns
}

func FoundationModelIdentifier_COHERE_COMMAND_R_V1() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"COHERE_COMMAND_R_V1",
		&returns,
	)
	return returns
}

func FoundationModelIdentifier_COHERE_COMMAND_TEXT_V14_7_4K() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"COHERE_COMMAND_TEXT_V14_7_4K",
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

func FoundationModelIdentifier_COHERE_EMBED_ENGLISH_V3_0_512() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"COHERE_EMBED_ENGLISH_V3_0_512",
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

func FoundationModelIdentifier_COHERE_EMBED_MULTILINGUAL_V3_0_512() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"COHERE_EMBED_MULTILINGUAL_V3_0_512",
		&returns,
	)
	return returns
}

func FoundationModelIdentifier_META_LLAMA_2_13B_CHAT_V1_0_4K() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"META_LLAMA_2_13B_CHAT_V1_0_4K",
		&returns,
	)
	return returns
}

func FoundationModelIdentifier_META_LLAMA_2_13B_V1() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"META_LLAMA_2_13B_V1",
		&returns,
	)
	return returns
}

func FoundationModelIdentifier_META_LLAMA_2_13B_V1_0_4K() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"META_LLAMA_2_13B_V1_0_4K",
		&returns,
	)
	return returns
}

func FoundationModelIdentifier_META_LLAMA_2_70B_CHAT_V1_0_4K() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"META_LLAMA_2_70B_CHAT_V1_0_4K",
		&returns,
	)
	return returns
}

func FoundationModelIdentifier_META_LLAMA_2_70B_V1() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"META_LLAMA_2_70B_V1",
		&returns,
	)
	return returns
}

func FoundationModelIdentifier_META_LLAMA_2_70B_V1_0_4K() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"META_LLAMA_2_70B_V1_0_4K",
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

func FoundationModelIdentifier_META_LLAMA_3_70_INSTRUCT_V1() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"META_LLAMA_3_70_INSTRUCT_V1",
		&returns,
	)
	return returns
}

func FoundationModelIdentifier_META_LLAMA_3_8B_INSTRUCT_V1() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"META_LLAMA_3_8B_INSTRUCT_V1",
		&returns,
	)
	return returns
}

func FoundationModelIdentifier_MISTRAL_LARGE_V0_1() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"MISTRAL_LARGE_V0_1",
		&returns,
	)
	return returns
}

func FoundationModelIdentifier_MISTRAL_MISTRAL_7B_INSTRUCT_V0_2() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"MISTRAL_MISTRAL_7B_INSTRUCT_V0_2",
		&returns,
	)
	return returns
}

func FoundationModelIdentifier_MISTRAL_MIXTRAL_8X7B_INSTRUCT_V0_1() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"MISTRAL_MIXTRAL_8X7B_INSTRUCT_V0_1",
		&returns,
	)
	return returns
}

func FoundationModelIdentifier_MISTRAL_SMALL_V0_1() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"MISTRAL_SMALL_V0_1",
		&returns,
	)
	return returns
}

func FoundationModelIdentifier_STABILITY_STABLE_DIFFUSION_XL() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"STABILITY_STABLE_DIFFUSION_XL",
		&returns,
	)
	return returns
}

func FoundationModelIdentifier_STABILITY_STABLE_DIFFUSION_XL_V0() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"STABILITY_STABLE_DIFFUSION_XL_V0",
		&returns,
	)
	return returns
}

func FoundationModelIdentifier_STABILITY_STABLE_DIFFUSION_XL_V1() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"STABILITY_STABLE_DIFFUSION_XL_V1",
		&returns,
	)
	return returns
}

func FoundationModelIdentifier_STABILITY_STABLE_DIFFUSION_XL_V1_0() FoundationModelIdentifier {
	_init_.Initialize()
	var returns FoundationModelIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		"STABILITY_STABLE_DIFFUSION_XL_V1_0",
		&returns,
	)
	return returns
}

