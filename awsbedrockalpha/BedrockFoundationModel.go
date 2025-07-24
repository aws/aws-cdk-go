package awsbedrockalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrock"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Bedrock models.
//
// If you need to use a model name that doesn't exist as a static member, you
// can instantiate a `BedrockFoundationModel` object, e.g: `new BedrockFoundationModel('my-model')`.
//
// Example:
//   // Create a specialized agent
//   customerSupportAgent := bedrock.NewAgent(this, jsii.String("CustomerSupportAgent"), &AgentProps{
//   	Instruction: jsii.String("You specialize in answering customer support questions."),
//   	FoundationModel: bedrock.BedrockFoundationModel_AMAZON_NOVA_LITE_V1(),
//   })
//
//   // Create an agent alias
//   customerSupportAlias := bedrock.NewAgentAlias(this, jsii.String("CustomerSupportAlias"), &AgentAliasProps{
//   	Agent: customerSupportAgent,
//   	AgentAliasName: jsii.String("production"),
//   })
//
//   // Create a main agent that collaborates with the specialized agent
//   mainAgent := bedrock.NewAgent(this, jsii.String("MainAgent"), &AgentProps{
//   	Instruction: jsii.String("You route specialized questions to other agents."),
//   	FoundationModel: bedrock.BedrockFoundationModel_AMAZON_NOVA_LITE_V1(),
//   	AgentCollaboration: map[string]interface{}{
//   		"type": bedrock.AgentCollaboratorType_SUPERVISOR,
//   		"collaborators": []AgentCollaborator{
//   			bedrock.NewAgentCollaborator(&AgentCollaboratorProps{
//   				"agentAlias": customerSupportAlias,
//   				"collaborationInstruction": jsii.String("Route customer support questions to this agent."),
//   				"collaboratorName": jsii.String("CustomerSupport"),
//   				"relayConversationHistory": jsii.Boolean(true),
//   			}),
//   		},
//   	},
//   })
//
// Experimental.
type BedrockFoundationModel interface {
	IBedrockInvokable
	// The ARN used for invoking the model.
	//
	// This is the same as modelArn for foundation models.
	// Experimental.
	InvokableArn() *string
	// The ARN of the foundation model.
	//
	// Format: arn:${Partition}:bedrock:${Region}::foundation-model/${ResourceId}.
	// Experimental.
	ModelArn() *string
	// The unique identifier of the foundation model.
	// Experimental.
	ModelId() *string
	// The vector types supported by this model for embeddings.
	//
	// Defines whether the model supports floating-point or binary vectors.
	// Experimental.
	SupportedVectorType() *[]VectorType
	// Whether this model supports integration with Bedrock Agents.
	//
	// When true, the model can be used with Bedrock Agents for automated task execution.
	// Experimental.
	SupportsAgents() *bool
	// Whether this model supports cross-region inference.
	//
	// When true, the model can be used with Cross-Region Inference Profiles.
	// Experimental.
	SupportsCrossRegion() *bool
	// Whether this model supports integration with Bedrock Knowledge Base.
	//
	// When true, the model can be used for knowledge base operations.
	// Experimental.
	SupportsKnowledgeBase() *bool
	// The dimensionality of the vector embeddings produced by this model.
	//
	// Only applicable for embedding models.
	// Experimental.
	VectorDimensions() *float64
	// Returns the ARN of the foundation model in the following format: `arn:${Partition}:bedrock:${Region}::foundation-model/${ResourceId}`.
	// Experimental.
	AsArn() *string
	// Returns the IModel.
	// Experimental.
	AsIModel() awsbedrock.IModel
	// Gives the appropriate policies to invoke and use the Foundation Model in the stack region.
	// Experimental.
	GrantInvoke(grantee awsiam.IGrantable) awsiam.Grant
	// Gives the appropriate policies to invoke and use the Foundation Model in all regions.
	// Experimental.
	GrantInvokeAllRegions(grantee awsiam.IGrantable) awsiam.Grant
	// Returns a string representation of an object.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BedrockFoundationModel
type jsiiProxy_BedrockFoundationModel struct {
	jsiiProxy_IBedrockInvokable
}

func (j *jsiiProxy_BedrockFoundationModel) InvokableArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"invokableArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockFoundationModel) ModelArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockFoundationModel) ModelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockFoundationModel) SupportedVectorType() *[]VectorType {
	var returns *[]VectorType
	_jsii_.Get(
		j,
		"supportedVectorType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockFoundationModel) SupportsAgents() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"supportsAgents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockFoundationModel) SupportsCrossRegion() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"supportsCrossRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockFoundationModel) SupportsKnowledgeBase() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"supportsKnowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockFoundationModel) VectorDimensions() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vectorDimensions",
		&returns,
	)
	return returns
}


// Experimental.
func NewBedrockFoundationModel(value *string, props *BedrockFoundationModelProps) BedrockFoundationModel {
	_init_.Initialize()

	if err := validateNewBedrockFoundationModelParameters(value, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockFoundationModel{}

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-alpha.BedrockFoundationModel",
		[]interface{}{value, props},
		&j,
	)

	return &j
}

// Experimental.
func NewBedrockFoundationModel_Override(b BedrockFoundationModel, value *string, props *BedrockFoundationModelProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-alpha.BedrockFoundationModel",
		[]interface{}{value, props},
		b,
	)
}

// Creates a BedrockFoundationModel from a FoundationModel.
//
// Use this method when you have a foundation model from the CDK.
//
// Returns: A new BedrockFoundationModel instance.
// Experimental.
func BedrockFoundationModel_FromCdkFoundationModel(modelId awsbedrock.FoundationModel, props *BedrockFoundationModelProps) BedrockFoundationModel {
	_init_.Initialize()

	if err := validateBedrockFoundationModel_FromCdkFoundationModelParameters(modelId, props); err != nil {
		panic(err)
	}
	var returns BedrockFoundationModel

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-alpha.BedrockFoundationModel",
		"fromCdkFoundationModel",
		[]interface{}{modelId, props},
		&returns,
	)

	return returns
}

// Creates a BedrockFoundationModel from a FoundationModelIdentifier.
//
// Use this method when you have a model identifier from the CDK.
//
// Returns: A new BedrockFoundationModel instance.
// Experimental.
func BedrockFoundationModel_FromCdkFoundationModelId(modelId awsbedrock.FoundationModelIdentifier, props *BedrockFoundationModelProps) BedrockFoundationModel {
	_init_.Initialize()

	if err := validateBedrockFoundationModel_FromCdkFoundationModelIdParameters(modelId, props); err != nil {
		panic(err)
	}
	var returns BedrockFoundationModel

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-alpha.BedrockFoundationModel",
		"fromCdkFoundationModelId",
		[]interface{}{modelId, props},
		&returns,
	)

	return returns
}

func BedrockFoundationModel_AI21_JAMBA_1_5_LARGE_V1() BedrockFoundationModel {
	_init_.Initialize()
	var returns BedrockFoundationModel
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.BedrockFoundationModel",
		"AI21_JAMBA_1_5_LARGE_V1",
		&returns,
	)
	return returns
}

func BedrockFoundationModel_AI21_JAMBA_1_5_MINI_V1() BedrockFoundationModel {
	_init_.Initialize()
	var returns BedrockFoundationModel
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.BedrockFoundationModel",
		"AI21_JAMBA_1_5_MINI_V1",
		&returns,
	)
	return returns
}

func BedrockFoundationModel_AI21_JAMBA_INSTRUCT_V1() BedrockFoundationModel {
	_init_.Initialize()
	var returns BedrockFoundationModel
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.BedrockFoundationModel",
		"AI21_JAMBA_INSTRUCT_V1",
		&returns,
	)
	return returns
}

func BedrockFoundationModel_AMAZON_NOVA_LITE_V1() BedrockFoundationModel {
	_init_.Initialize()
	var returns BedrockFoundationModel
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.BedrockFoundationModel",
		"AMAZON_NOVA_LITE_V1",
		&returns,
	)
	return returns
}

func BedrockFoundationModel_AMAZON_NOVA_MICRO_V1() BedrockFoundationModel {
	_init_.Initialize()
	var returns BedrockFoundationModel
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.BedrockFoundationModel",
		"AMAZON_NOVA_MICRO_V1",
		&returns,
	)
	return returns
}

func BedrockFoundationModel_AMAZON_NOVA_PREMIER_V1() BedrockFoundationModel {
	_init_.Initialize()
	var returns BedrockFoundationModel
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.BedrockFoundationModel",
		"AMAZON_NOVA_PREMIER_V1",
		&returns,
	)
	return returns
}

func BedrockFoundationModel_AMAZON_NOVA_PRO_V1() BedrockFoundationModel {
	_init_.Initialize()
	var returns BedrockFoundationModel
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.BedrockFoundationModel",
		"AMAZON_NOVA_PRO_V1",
		&returns,
	)
	return returns
}

func BedrockFoundationModel_AMAZON_TITAN_PREMIER_V1_0() BedrockFoundationModel {
	_init_.Initialize()
	var returns BedrockFoundationModel
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.BedrockFoundationModel",
		"AMAZON_TITAN_PREMIER_V1_0",
		&returns,
	)
	return returns
}

func BedrockFoundationModel_AMAZON_TITAN_TEXT_EXPRESS_V1() BedrockFoundationModel {
	_init_.Initialize()
	var returns BedrockFoundationModel
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.BedrockFoundationModel",
		"AMAZON_TITAN_TEXT_EXPRESS_V1",
		&returns,
	)
	return returns
}

func BedrockFoundationModel_ANTHROPIC_CLAUDE_3_5_HAIKU_V1_0() BedrockFoundationModel {
	_init_.Initialize()
	var returns BedrockFoundationModel
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.BedrockFoundationModel",
		"ANTHROPIC_CLAUDE_3_5_HAIKU_V1_0",
		&returns,
	)
	return returns
}

func BedrockFoundationModel_ANTHROPIC_CLAUDE_3_5_SONNET_V1_0() BedrockFoundationModel {
	_init_.Initialize()
	var returns BedrockFoundationModel
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.BedrockFoundationModel",
		"ANTHROPIC_CLAUDE_3_5_SONNET_V1_0",
		&returns,
	)
	return returns
}

func BedrockFoundationModel_ANTHROPIC_CLAUDE_3_5_SONNET_V2_0() BedrockFoundationModel {
	_init_.Initialize()
	var returns BedrockFoundationModel
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.BedrockFoundationModel",
		"ANTHROPIC_CLAUDE_3_5_SONNET_V2_0",
		&returns,
	)
	return returns
}

func BedrockFoundationModel_ANTHROPIC_CLAUDE_3_7_SONNET_V1_0() BedrockFoundationModel {
	_init_.Initialize()
	var returns BedrockFoundationModel
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.BedrockFoundationModel",
		"ANTHROPIC_CLAUDE_3_7_SONNET_V1_0",
		&returns,
	)
	return returns
}

func BedrockFoundationModel_ANTHROPIC_CLAUDE_HAIKU_V1_0() BedrockFoundationModel {
	_init_.Initialize()
	var returns BedrockFoundationModel
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.BedrockFoundationModel",
		"ANTHROPIC_CLAUDE_HAIKU_V1_0",
		&returns,
	)
	return returns
}

func BedrockFoundationModel_ANTHROPIC_CLAUDE_INSTANT_V1_2() BedrockFoundationModel {
	_init_.Initialize()
	var returns BedrockFoundationModel
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.BedrockFoundationModel",
		"ANTHROPIC_CLAUDE_INSTANT_V1_2",
		&returns,
	)
	return returns
}

func BedrockFoundationModel_ANTHROPIC_CLAUDE_OPUS_V1_0() BedrockFoundationModel {
	_init_.Initialize()
	var returns BedrockFoundationModel
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.BedrockFoundationModel",
		"ANTHROPIC_CLAUDE_OPUS_V1_0",
		&returns,
	)
	return returns
}

func BedrockFoundationModel_ANTHROPIC_CLAUDE_SONNET_V1_0() BedrockFoundationModel {
	_init_.Initialize()
	var returns BedrockFoundationModel
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.BedrockFoundationModel",
		"ANTHROPIC_CLAUDE_SONNET_V1_0",
		&returns,
	)
	return returns
}

func BedrockFoundationModel_ANTHROPIC_CLAUDE_V2() BedrockFoundationModel {
	_init_.Initialize()
	var returns BedrockFoundationModel
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.BedrockFoundationModel",
		"ANTHROPIC_CLAUDE_V2",
		&returns,
	)
	return returns
}

func BedrockFoundationModel_ANTHROPIC_CLAUDE_V2_1() BedrockFoundationModel {
	_init_.Initialize()
	var returns BedrockFoundationModel
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.BedrockFoundationModel",
		"ANTHROPIC_CLAUDE_V2_1",
		&returns,
	)
	return returns
}

func BedrockFoundationModel_COHERE_EMBED_ENGLISH_V3() BedrockFoundationModel {
	_init_.Initialize()
	var returns BedrockFoundationModel
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.BedrockFoundationModel",
		"COHERE_EMBED_ENGLISH_V3",
		&returns,
	)
	return returns
}

func BedrockFoundationModel_COHERE_EMBED_MULTILINGUAL_V3() BedrockFoundationModel {
	_init_.Initialize()
	var returns BedrockFoundationModel
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.BedrockFoundationModel",
		"COHERE_EMBED_MULTILINGUAL_V3",
		&returns,
	)
	return returns
}

func BedrockFoundationModel_DEEPSEEK_R1_V1() BedrockFoundationModel {
	_init_.Initialize()
	var returns BedrockFoundationModel
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.BedrockFoundationModel",
		"DEEPSEEK_R1_V1",
		&returns,
	)
	return returns
}

func BedrockFoundationModel_META_LLAMA_3_1_70B_INSTRUCT_V1() BedrockFoundationModel {
	_init_.Initialize()
	var returns BedrockFoundationModel
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.BedrockFoundationModel",
		"META_LLAMA_3_1_70B_INSTRUCT_V1",
		&returns,
	)
	return returns
}

func BedrockFoundationModel_META_LLAMA_3_1_8B_INSTRUCT_V1() BedrockFoundationModel {
	_init_.Initialize()
	var returns BedrockFoundationModel
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.BedrockFoundationModel",
		"META_LLAMA_3_1_8B_INSTRUCT_V1",
		&returns,
	)
	return returns
}

func BedrockFoundationModel_META_LLAMA_3_2_11B_INSTRUCT_V1() BedrockFoundationModel {
	_init_.Initialize()
	var returns BedrockFoundationModel
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.BedrockFoundationModel",
		"META_LLAMA_3_2_11B_INSTRUCT_V1",
		&returns,
	)
	return returns
}

func BedrockFoundationModel_META_LLAMA_3_2_1B_INSTRUCT_V1() BedrockFoundationModel {
	_init_.Initialize()
	var returns BedrockFoundationModel
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.BedrockFoundationModel",
		"META_LLAMA_3_2_1B_INSTRUCT_V1",
		&returns,
	)
	return returns
}

func BedrockFoundationModel_META_LLAMA_3_2_3B_INSTRUCT_V1() BedrockFoundationModel {
	_init_.Initialize()
	var returns BedrockFoundationModel
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.BedrockFoundationModel",
		"META_LLAMA_3_2_3B_INSTRUCT_V1",
		&returns,
	)
	return returns
}

func BedrockFoundationModel_META_LLAMA_3_3_70B_INSTRUCT_V1() BedrockFoundationModel {
	_init_.Initialize()
	var returns BedrockFoundationModel
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.BedrockFoundationModel",
		"META_LLAMA_3_3_70B_INSTRUCT_V1",
		&returns,
	)
	return returns
}

func BedrockFoundationModel_META_LLAMA_4_MAVERICK_17B_INSTRUCT_V1() BedrockFoundationModel {
	_init_.Initialize()
	var returns BedrockFoundationModel
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.BedrockFoundationModel",
		"META_LLAMA_4_MAVERICK_17B_INSTRUCT_V1",
		&returns,
	)
	return returns
}

func BedrockFoundationModel_META_LLAMA_4_SCOUT_17B_INSTRUCT_V1() BedrockFoundationModel {
	_init_.Initialize()
	var returns BedrockFoundationModel
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.BedrockFoundationModel",
		"META_LLAMA_4_SCOUT_17B_INSTRUCT_V1",
		&returns,
	)
	return returns
}

func BedrockFoundationModel_MISTRAL_7B_INSTRUCT_V0() BedrockFoundationModel {
	_init_.Initialize()
	var returns BedrockFoundationModel
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.BedrockFoundationModel",
		"MISTRAL_7B_INSTRUCT_V0",
		&returns,
	)
	return returns
}

func BedrockFoundationModel_MISTRAL_LARGE_2402_V1() BedrockFoundationModel {
	_init_.Initialize()
	var returns BedrockFoundationModel
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.BedrockFoundationModel",
		"MISTRAL_LARGE_2402_V1",
		&returns,
	)
	return returns
}

func BedrockFoundationModel_MISTRAL_LARGE_2407_V1() BedrockFoundationModel {
	_init_.Initialize()
	var returns BedrockFoundationModel
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.BedrockFoundationModel",
		"MISTRAL_LARGE_2407_V1",
		&returns,
	)
	return returns
}

func BedrockFoundationModel_MISTRAL_MIXTRAL_8X7B_INSTRUCT_V0() BedrockFoundationModel {
	_init_.Initialize()
	var returns BedrockFoundationModel
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.BedrockFoundationModel",
		"MISTRAL_MIXTRAL_8X7B_INSTRUCT_V0",
		&returns,
	)
	return returns
}

func BedrockFoundationModel_MISTRAL_PIXTRAL_LARGE_2502_V1() BedrockFoundationModel {
	_init_.Initialize()
	var returns BedrockFoundationModel
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.BedrockFoundationModel",
		"MISTRAL_PIXTRAL_LARGE_2502_V1",
		&returns,
	)
	return returns
}

func BedrockFoundationModel_MISTRAL_SMALL_2402_V1() BedrockFoundationModel {
	_init_.Initialize()
	var returns BedrockFoundationModel
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.BedrockFoundationModel",
		"MISTRAL_SMALL_2402_V1",
		&returns,
	)
	return returns
}

func BedrockFoundationModel_TITAN_EMBED_TEXT_V1() BedrockFoundationModel {
	_init_.Initialize()
	var returns BedrockFoundationModel
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.BedrockFoundationModel",
		"TITAN_EMBED_TEXT_V1",
		&returns,
	)
	return returns
}

func BedrockFoundationModel_TITAN_EMBED_TEXT_V2_1024() BedrockFoundationModel {
	_init_.Initialize()
	var returns BedrockFoundationModel
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.BedrockFoundationModel",
		"TITAN_EMBED_TEXT_V2_1024",
		&returns,
	)
	return returns
}

func BedrockFoundationModel_TITAN_EMBED_TEXT_V2_256() BedrockFoundationModel {
	_init_.Initialize()
	var returns BedrockFoundationModel
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.BedrockFoundationModel",
		"TITAN_EMBED_TEXT_V2_256",
		&returns,
	)
	return returns
}

func BedrockFoundationModel_TITAN_EMBED_TEXT_V2_512() BedrockFoundationModel {
	_init_.Initialize()
	var returns BedrockFoundationModel
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.BedrockFoundationModel",
		"TITAN_EMBED_TEXT_V2_512",
		&returns,
	)
	return returns
}

func (b *jsiiProxy_BedrockFoundationModel) AsArn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"asArn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockFoundationModel) AsIModel() awsbedrock.IModel {
	var returns awsbedrock.IModel

	_jsii_.Invoke(
		b,
		"asIModel",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockFoundationModel) GrantInvoke(grantee awsiam.IGrantable) awsiam.Grant {
	if err := b.validateGrantInvokeParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		b,
		"grantInvoke",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockFoundationModel) GrantInvokeAllRegions(grantee awsiam.IGrantable) awsiam.Grant {
	if err := b.validateGrantInvokeAllRegionsParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		b,
		"grantInvokeAllRegions",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockFoundationModel) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

