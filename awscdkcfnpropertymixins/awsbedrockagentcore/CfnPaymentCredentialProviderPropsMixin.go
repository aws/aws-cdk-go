package awsbedrockagentcore

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsbedrockagentcore/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::BedrockAgentCore::PaymentCredentialProvider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnPaymentCredentialProviderPropsMixin := awscdkcfnpropertymixins.Aws_bedrockagentcore.NewCfnPaymentCredentialProviderPropsMixin(&CfnPaymentCredentialProviderMixinProps{
//   	CredentialProviderVendor: jsii.String("credentialProviderVendor"),
//   	Name: jsii.String("name"),
//   	ProviderConfigurationInput: &PaymentProviderConfigurationInputProperty{
//   		CoinbaseCdpConfiguration: &CoinbaseCdpConfigurationInputProperty{
//   			ApiKeyId: jsii.String("apiKeyId"),
//   			ApiKeySecret: jsii.String("apiKeySecret"),
//   			WalletSecret: jsii.String("walletSecret"),
//   		},
//   		StripePrivyConfiguration: &StripePrivyConfigurationInputProperty{
//   			AppId: jsii.String("appId"),
//   			AppSecret: jsii.String("appSecret"),
//   			AuthorizationId: jsii.String("authorizationId"),
//   			AuthorizationPrivateKey: jsii.String("authorizationPrivateKey"),
//   		},
//   	},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-paymentcredentialprovider.html
//
type CfnPaymentCredentialProviderPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnPaymentCredentialProviderMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnPaymentCredentialProviderPropsMixin
type jsiiProxy_CfnPaymentCredentialProviderPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnPaymentCredentialProviderPropsMixin) Props() *CfnPaymentCredentialProviderMixinProps {
	var returns *CfnPaymentCredentialProviderMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPaymentCredentialProviderPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::BedrockAgentCore::PaymentCredentialProvider`.
func NewCfnPaymentCredentialProviderPropsMixin(props *CfnPaymentCredentialProviderMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnPaymentCredentialProviderPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnPaymentCredentialProviderPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnPaymentCredentialProviderPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnPaymentCredentialProviderPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::BedrockAgentCore::PaymentCredentialProvider`.
func NewCfnPaymentCredentialProviderPropsMixin_Override(c CfnPaymentCredentialProviderPropsMixin, props *CfnPaymentCredentialProviderMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnPaymentCredentialProviderPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnPaymentCredentialProviderPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPaymentCredentialProviderPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnPaymentCredentialProviderPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPaymentCredentialProviderPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnPaymentCredentialProviderPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPaymentCredentialProviderPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnPaymentCredentialProviderPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

