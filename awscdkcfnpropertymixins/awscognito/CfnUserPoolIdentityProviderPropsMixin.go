package awscognito

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awscognito/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::Cognito::UserPoolIdentityProvider` resource creates an identity provider for a user pool.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var attributeMapping interface{}
//   var mergeStrategy IMergeStrategy
//   var providerDetails interface{}
//
//   cfnUserPoolIdentityProviderPropsMixin := awscdkcfnpropertymixins.Aws_cognito.NewCfnUserPoolIdentityProviderPropsMixin(&CfnUserPoolIdentityProviderMixinProps{
//   	AttributeMapping: attributeMapping,
//   	IdpIdentifiers: []*string{
//   		jsii.String("idpIdentifiers"),
//   	},
//   	ProviderDetails: providerDetails,
//   	ProviderName: jsii.String("providerName"),
//   	ProviderType: jsii.String("providerType"),
//   	UserPoolId: jsii.String("userPoolId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolidentityprovider.html
//
type CfnUserPoolIdentityProviderPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnUserPoolIdentityProviderMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnUserPoolIdentityProviderPropsMixin
type jsiiProxy_CfnUserPoolIdentityProviderPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnUserPoolIdentityProviderPropsMixin) Props() *CfnUserPoolIdentityProviderMixinProps {
	var returns *CfnUserPoolIdentityProviderMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolIdentityProviderPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Cognito::UserPoolIdentityProvider`.
func NewCfnUserPoolIdentityProviderPropsMixin(props *CfnUserPoolIdentityProviderMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnUserPoolIdentityProviderPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnUserPoolIdentityProviderPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnUserPoolIdentityProviderPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolIdentityProviderPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Cognito::UserPoolIdentityProvider`.
func NewCfnUserPoolIdentityProviderPropsMixin_Override(c CfnUserPoolIdentityProviderPropsMixin, props *CfnUserPoolIdentityProviderMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolIdentityProviderPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnUserPoolIdentityProviderPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnUserPoolIdentityProviderPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolIdentityProviderPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnUserPoolIdentityProviderPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolIdentityProviderPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnUserPoolIdentityProviderPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnUserPoolIdentityProviderPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

