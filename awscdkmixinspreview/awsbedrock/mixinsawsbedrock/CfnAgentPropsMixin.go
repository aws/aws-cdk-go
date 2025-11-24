package mixinsawsbedrock

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsbedrock/mixinsawsbedrock/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies an agent as a resource in a top-level template. Minimally, you must specify the following properties:.
//
// - AgentName – Specify a name for the agent.
// - AgentResourceRoleArn – Specify the Amazon Resource Name (ARN) of the service role with permissions to invoke API operations on the agent. For more information, see [Create a service role for Agents for Amazon Bedrock](https://docs.aws.amazon.com/bedrock/latest/userguide/agents-permissions.html) .
// - FoundationModel – Specify the model ID of a foundation model to use when invoking the agent. For more information, see [Supported regions and models for Agents for Amazon Bedrock](https://docs.aws.amazon.com//bedrock/latest/userguide/agents-supported.html) .
//
// For more information about using agents in Amazon Bedrock , see [Agents for Amazon Bedrock](https://docs.aws.amazon.com/bedrock/latest/userguide/agents.html) .
//
// See the *Properties* section below for descriptions of both the required and optional properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var additionalModelRequestFields interface{}
//
//   cfnAgentPropsMixin := awscdkmixinspreview.Mixins.NewCfnAgentPropsMixin(&CfnAgentMixinProps{
//   	ActionGroups: []interface{}{
//   		&AgentActionGroupProperty{
//   			ActionGroupExecutor: &ActionGroupExecutorProperty{
//   				CustomControl: jsii.String("customControl"),
//   				Lambda: jsii.String("lambda"),
//   			},
//   			ActionGroupName: jsii.String("actionGroupName"),
//   			ActionGroupState: jsii.String("actionGroupState"),
//   			ApiSchema: &APISchemaProperty{
//   				Payload: jsii.String("payload"),
//   				S3: &S3IdentifierProperty{
//   					S3BucketName: jsii.String("s3BucketName"),
//   					S3ObjectKey: jsii.String("s3ObjectKey"),
//   				},
//   			},
//   			Description: jsii.String("description"),
//   			FunctionSchema: &FunctionSchemaProperty{
//   				Functions: []interface{}{
//   					&FunctionProperty{
//   						Description: jsii.String("description"),
//   						Name: jsii.String("name"),
//   						Parameters: map[string]interface{}{
//   							"parametersKey": &ParameterDetailProperty{
//   								"description": jsii.String("description"),
//   								"required": jsii.Boolean(false),
//   								"type": jsii.String("type"),
//   							},
//   						},
//   						RequireConfirmation: jsii.String("requireConfirmation"),
//   					},
//   				},
//   			},
//   			ParentActionGroupSignature: jsii.String("parentActionGroupSignature"),
//   			SkipResourceInUseCheckOnDelete: jsii.Boolean(false),
//   		},
//   	},
//   	AgentCollaboration: jsii.String("agentCollaboration"),
//   	AgentCollaborators: []interface{}{
//   		&AgentCollaboratorProperty{
//   			AgentDescriptor: &AgentDescriptorProperty{
//   				AliasArn: jsii.String("aliasArn"),
//   			},
//   			CollaborationInstruction: jsii.String("collaborationInstruction"),
//   			CollaboratorName: jsii.String("collaboratorName"),
//   			RelayConversationHistory: jsii.String("relayConversationHistory"),
//   		},
//   	},
//   	AgentName: jsii.String("agentName"),
//   	AgentResourceRoleArn: jsii.String("agentResourceRoleArn"),
//   	AutoPrepare: jsii.Boolean(false),
//   	CustomerEncryptionKeyArn: jsii.String("customerEncryptionKeyArn"),
//   	CustomOrchestration: &CustomOrchestrationProperty{
//   		Executor: &OrchestrationExecutorProperty{
//   			Lambda: jsii.String("lambda"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	FoundationModel: jsii.String("foundationModel"),
//   	GuardrailConfiguration: &GuardrailConfigurationProperty{
//   		GuardrailIdentifier: jsii.String("guardrailIdentifier"),
//   		GuardrailVersion: jsii.String("guardrailVersion"),
//   	},
//   	IdleSessionTtlInSeconds: jsii.Number(123),
//   	Instruction: jsii.String("instruction"),
//   	KnowledgeBases: []interface{}{
//   		&AgentKnowledgeBaseProperty{
//   			Description: jsii.String("description"),
//   			KnowledgeBaseId: jsii.String("knowledgeBaseId"),
//   			KnowledgeBaseState: jsii.String("knowledgeBaseState"),
//   		},
//   	},
//   	MemoryConfiguration: &MemoryConfigurationProperty{
//   		EnabledMemoryTypes: []*string{
//   			jsii.String("enabledMemoryTypes"),
//   		},
//   		SessionSummaryConfiguration: &SessionSummaryConfigurationProperty{
//   			MaxRecentSessions: jsii.Number(123),
//   		},
//   		StorageDays: jsii.Number(123),
//   	},
//   	OrchestrationType: jsii.String("orchestrationType"),
//   	PromptOverrideConfiguration: &PromptOverrideConfigurationProperty{
//   		OverrideLambda: jsii.String("overrideLambda"),
//   		PromptConfigurations: []interface{}{
//   			&PromptConfigurationProperty{
//   				AdditionalModelRequestFields: additionalModelRequestFields,
//   				BasePromptTemplate: jsii.String("basePromptTemplate"),
//   				FoundationModel: jsii.String("foundationModel"),
//   				InferenceConfiguration: &InferenceConfigurationProperty{
//   					MaximumLength: jsii.Number(123),
//   					StopSequences: []*string{
//   						jsii.String("stopSequences"),
//   					},
//   					Temperature: jsii.Number(123),
//   					TopK: jsii.Number(123),
//   					TopP: jsii.Number(123),
//   				},
//   				ParserMode: jsii.String("parserMode"),
//   				PromptCreationMode: jsii.String("promptCreationMode"),
//   				PromptState: jsii.String("promptState"),
//   				PromptType: jsii.String("promptType"),
//   			},
//   		},
//   	},
//   	SkipResourceInUseCheckOnDelete: jsii.Boolean(false),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	TestAliasTags: map[string]*string{
//   		"testAliasTagsKey": jsii.String("testAliasTags"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-agent.html
//
type CfnAgentPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnAgentMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAgentPropsMixin
type jsiiProxy_CfnAgentPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnAgentPropsMixin) Props() *CfnAgentMixinProps {
	var returns *CfnAgentMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAgentPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Bedrock::Agent`.
func NewCfnAgentPropsMixin(props *CfnAgentMixinProps, options *mixins.CfnPropertyMixinOptions) CfnAgentPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAgentPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAgentPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnAgentPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Bedrock::Agent`.
func NewCfnAgentPropsMixin_Override(c CfnAgentPropsMixin, props *CfnAgentMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnAgentPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnAgentPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAgentPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnAgentPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAgentPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnAgentPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAgentPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"applyTo",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAgentPropsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

