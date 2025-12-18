package previewawsbedrockagentcoremixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsbedrockagentcore/previewawsbedrockagentcoremixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Amazon Bedrock AgentCore Gateway provides a unified connectivity layer between agents and the tools and resources they need to interact with.
//
// For more information about creating a gateway, see [Set up an Amazon Bedrock AgentCore gateway](https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/gateway-building.html) .
//
// See the *Properties* section below for descriptions of both the required and optional properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnGatewayPropsMixin := awscdkmixinspreview.Mixins.NewCfnGatewayPropsMixin(&CfnGatewayMixinProps{
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
//   	AuthorizerType: jsii.String("authorizerType"),
//   	Description: jsii.String("description"),
//   	ExceptionLevel: jsii.String("exceptionLevel"),
//   	InterceptorConfigurations: []interface{}{
//   		&GatewayInterceptorConfigurationProperty{
//   			InputConfiguration: &InterceptorInputConfigurationProperty{
//   				PassRequestHeaders: jsii.Boolean(false),
//   			},
//   			InterceptionPoints: []*string{
//   				jsii.String("interceptionPoints"),
//   			},
//   			Interceptor: &InterceptorConfigurationProperty{
//   				Lambda: &LambdaInterceptorConfigurationProperty{
//   					Arn: jsii.String("arn"),
//   				},
//   			},
//   		},
//   	},
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	Name: jsii.String("name"),
//   	ProtocolConfiguration: &GatewayProtocolConfigurationProperty{
//   		Mcp: &MCPGatewayConfigurationProperty{
//   			Instructions: jsii.String("instructions"),
//   			SearchType: jsii.String("searchType"),
//   			SupportedVersions: []*string{
//   				jsii.String("supportedVersions"),
//   			},
//   		},
//   	},
//   	ProtocolType: jsii.String("protocolType"),
//   	RoleArn: jsii.String("roleArn"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-gateway.html
//
type CfnGatewayPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnGatewayMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnGatewayPropsMixin
type jsiiProxy_CfnGatewayPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnGatewayPropsMixin) Props() *CfnGatewayMixinProps {
	var returns *CfnGatewayMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGatewayPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::BedrockAgentCore::Gateway`.
func NewCfnGatewayPropsMixin(props *CfnGatewayMixinProps, options *mixins.CfnPropertyMixinOptions) CfnGatewayPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnGatewayPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnGatewayPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::BedrockAgentCore::Gateway`.
func NewCfnGatewayPropsMixin_Override(c CfnGatewayPropsMixin, props *CfnGatewayMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnGatewayPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnGatewayPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnGatewayPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnGatewayPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnGatewayPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

