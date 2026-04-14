package awsbedrockagentcore

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsbedrockagentcore/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::BedrockAgentCore::ApiKeyCredentialProvider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnApiKeyCredentialProviderPropsMixin := awscdkcfnpropertymixins.Aws_bedrockagentcore.NewCfnApiKeyCredentialProviderPropsMixin(&CfnApiKeyCredentialProviderMixinProps{
//   	ApiKey: jsii.String("apiKey"),
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-apikeycredentialprovider.html
//
type CfnApiKeyCredentialProviderPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnApiKeyCredentialProviderMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnApiKeyCredentialProviderPropsMixin
type jsiiProxy_CfnApiKeyCredentialProviderPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnApiKeyCredentialProviderPropsMixin) Props() *CfnApiKeyCredentialProviderMixinProps {
	var returns *CfnApiKeyCredentialProviderMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiKeyCredentialProviderPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::BedrockAgentCore::ApiKeyCredentialProvider`.
func NewCfnApiKeyCredentialProviderPropsMixin(props *CfnApiKeyCredentialProviderMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnApiKeyCredentialProviderPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnApiKeyCredentialProviderPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnApiKeyCredentialProviderPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnApiKeyCredentialProviderPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::BedrockAgentCore::ApiKeyCredentialProvider`.
func NewCfnApiKeyCredentialProviderPropsMixin_Override(c CfnApiKeyCredentialProviderPropsMixin, props *CfnApiKeyCredentialProviderMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnApiKeyCredentialProviderPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnApiKeyCredentialProviderPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnApiKeyCredentialProviderPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnApiKeyCredentialProviderPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnApiKeyCredentialProviderPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnApiKeyCredentialProviderPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnApiKeyCredentialProviderPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnApiKeyCredentialProviderPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

