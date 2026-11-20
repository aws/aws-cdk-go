package awssso

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awssso/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::SSO::TrustedTokenIssuer, a trusted token issuer configuration for IAM Identity Center that enables trusted identity propagation with external identity providers using OIDC.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnTrustedTokenIssuerPropsMixin := awscdkcfnpropertymixins.Aws_sso.NewCfnTrustedTokenIssuerPropsMixin(&CfnTrustedTokenIssuerMixinProps{
//   	InstanceArn: jsii.String("instanceArn"),
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TrustedTokenIssuerConfiguration: &TrustedTokenIssuerConfigurationProperty{
//   		OidcJwtConfiguration: &OidcJwtConfigurationProperty{
//   			ClaimAttributePath: jsii.String("claimAttributePath"),
//   			IdentityStoreAttributePath: jsii.String("identityStoreAttributePath"),
//   			IssuerUrl: jsii.String("issuerUrl"),
//   			JwksRetrievalOption: jsii.String("jwksRetrievalOption"),
//   		},
//   	},
//   	TrustedTokenIssuerType: jsii.String("trustedTokenIssuerType"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-trustedtokenissuer.html
//
type CfnTrustedTokenIssuerPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnTrustedTokenIssuerMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnTrustedTokenIssuerPropsMixin
type jsiiProxy_CfnTrustedTokenIssuerPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnTrustedTokenIssuerPropsMixin) Props() *CfnTrustedTokenIssuerMixinProps {
	var returns *CfnTrustedTokenIssuerMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrustedTokenIssuerPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SSO::TrustedTokenIssuer`.
func NewCfnTrustedTokenIssuerPropsMixin(props *CfnTrustedTokenIssuerMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnTrustedTokenIssuerPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnTrustedTokenIssuerPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTrustedTokenIssuerPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_sso.CfnTrustedTokenIssuerPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SSO::TrustedTokenIssuer`.
func NewCfnTrustedTokenIssuerPropsMixin_Override(c CfnTrustedTokenIssuerPropsMixin, props *CfnTrustedTokenIssuerMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_sso.CfnTrustedTokenIssuerPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnTrustedTokenIssuerPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTrustedTokenIssuerPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_sso.CfnTrustedTokenIssuerPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTrustedTokenIssuerPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_sso.CfnTrustedTokenIssuerPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTrustedTokenIssuerPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnTrustedTokenIssuerPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

