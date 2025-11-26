package previewawsbedrockagentcoremixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsbedrockagentcore/previewawsbedrockagentcoremixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Contains information about an agent runtime. An agent runtime is the execution environment for a Amazon Bedrock Agent.
//
// AgentCore Runtime is a secure, serverless runtime purpose-built for deploying and scaling dynamic AI agents and tools using any open-source framework including LangGraph, CrewAI, and Strands Agents, any protocol, and any model.
//
// For more information about using agent runtime in Amazon Bedrock AgentCore, see [Host agent or tools with Amazon Bedrock AgentCore Runtime](https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/agents-tools-runtime.html) .
//
// See the *Properties* section below for descriptions of both the required and optional properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRuntimePropsMixin := awscdkmixinspreview.Mixins.NewCfnRuntimePropsMixin(&CfnRuntimeMixinProps{
//   	AgentRuntimeArtifact: &AgentRuntimeArtifactProperty{
//   		CodeConfiguration: &CodeConfigurationProperty{
//   			Code: &CodeProperty{
//   				S3: &S3LocationProperty{
//   					Bucket: jsii.String("bucket"),
//   					Prefix: jsii.String("prefix"),
//   					VersionId: jsii.String("versionId"),
//   				},
//   			},
//   			EntryPoint: []*string{
//   				jsii.String("entryPoint"),
//   			},
//   			Runtime: jsii.String("runtime"),
//   		},
//   		ContainerConfiguration: &ContainerConfigurationProperty{
//   			ContainerUri: jsii.String("containerUri"),
//   		},
//   	},
//   	AgentRuntimeName: jsii.String("agentRuntimeName"),
//   	AuthorizerConfiguration: &AuthorizerConfigurationProperty{
//   		CustomJwtAuthorizer: &CustomJWTAuthorizerConfigurationProperty{
//   			AllowedAudience: []*string{
//   				jsii.String("allowedAudience"),
//   			},
//   			AllowedClients: []*string{
//   				jsii.String("allowedClients"),
//   			},
//   			DiscoveryUrl: jsii.String("discoveryUrl"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	EnvironmentVariables: map[string]*string{
//   		"environmentVariablesKey": jsii.String("environmentVariables"),
//   	},
//   	LifecycleConfiguration: &LifecycleConfigurationProperty{
//   		IdleRuntimeSessionTimeout: jsii.Number(123),
//   		MaxLifetime: jsii.Number(123),
//   	},
//   	NetworkConfiguration: &NetworkConfigurationProperty{
//   		NetworkMode: jsii.String("networkMode"),
//   		NetworkModeConfig: &VpcConfigProperty{
//   			SecurityGroups: []*string{
//   				jsii.String("securityGroups"),
//   			},
//   			Subnets: []*string{
//   				jsii.String("subnets"),
//   			},
//   		},
//   	},
//   	ProtocolConfiguration: jsii.String("protocolConfiguration"),
//   	RequestHeaderConfiguration: &RequestHeaderConfigurationProperty{
//   		RequestHeaderAllowlist: []*string{
//   			jsii.String("requestHeaderAllowlist"),
//   		},
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-runtime.html
//
type CfnRuntimePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnRuntimeMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnRuntimePropsMixin
type jsiiProxy_CfnRuntimePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnRuntimePropsMixin) Props() *CfnRuntimeMixinProps {
	var returns *CfnRuntimeMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuntimePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::BedrockAgentCore::Runtime`.
func NewCfnRuntimePropsMixin(props *CfnRuntimeMixinProps, options *mixins.CfnPropertyMixinOptions) CfnRuntimePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnRuntimePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnRuntimePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::BedrockAgentCore::Runtime`.
func NewCfnRuntimePropsMixin_Override(c CfnRuntimePropsMixin, props *CfnRuntimeMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnRuntimePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnRuntimePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRuntimePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnRuntimePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnRuntimePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

