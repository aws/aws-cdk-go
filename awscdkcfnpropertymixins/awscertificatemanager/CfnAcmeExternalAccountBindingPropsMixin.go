package awscertificatemanager

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awscertificatemanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::CertificateManager::AcmeExternalAccountBinding.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnAcmeExternalAccountBindingPropsMixin := awscdkcfnpropertymixins.Aws_certificatemanager.NewCfnAcmeExternalAccountBindingPropsMixin(&CfnAcmeExternalAccountBindingMixinProps{
//   	AcmeEndpointArn: jsii.String("acmeEndpointArn"),
//   	Expiration: &ExpirationProperty{
//   		Type: jsii.String("type"),
//   		Value: jsii.Number(123),
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   	Tags: []TagsItemsProperty{
//   		&TagsItemsProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-certificatemanager-acmeexternalaccountbinding.html
//
type CfnAcmeExternalAccountBindingPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnAcmeExternalAccountBindingMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAcmeExternalAccountBindingPropsMixin
type jsiiProxy_CfnAcmeExternalAccountBindingPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnAcmeExternalAccountBindingPropsMixin) Props() *CfnAcmeExternalAccountBindingMixinProps {
	var returns *CfnAcmeExternalAccountBindingMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAcmeExternalAccountBindingPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CertificateManager::AcmeExternalAccountBinding`.
func NewCfnAcmeExternalAccountBindingPropsMixin(props *CfnAcmeExternalAccountBindingMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnAcmeExternalAccountBindingPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAcmeExternalAccountBindingPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAcmeExternalAccountBindingPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_certificatemanager.CfnAcmeExternalAccountBindingPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CertificateManager::AcmeExternalAccountBinding`.
func NewCfnAcmeExternalAccountBindingPropsMixin_Override(c CfnAcmeExternalAccountBindingPropsMixin, props *CfnAcmeExternalAccountBindingMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_certificatemanager.CfnAcmeExternalAccountBindingPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnAcmeExternalAccountBindingPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAcmeExternalAccountBindingPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_certificatemanager.CfnAcmeExternalAccountBindingPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAcmeExternalAccountBindingPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_certificatemanager.CfnAcmeExternalAccountBindingPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAcmeExternalAccountBindingPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnAcmeExternalAccountBindingPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

