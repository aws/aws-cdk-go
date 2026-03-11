package awslicensemanager

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awslicensemanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies a grant.
//
// A grant shares the use of license entitlements with specific AWS accounts . For more information, see [Granted licenses](https://docs.aws.amazon.com/license-manager/latest/userguide/granted-licenses.html) in the *License Manager User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnGrantPropsMixin := awscdkcfnpropertymixins.Aws_licensemanager.NewCfnGrantPropsMixin(&CfnGrantMixinProps{
//   	AllowedOperations: []*string{
//   		jsii.String("allowedOperations"),
//   	},
//   	GrantName: jsii.String("grantName"),
//   	HomeRegion: jsii.String("homeRegion"),
//   	LicenseArn: jsii.String("licenseArn"),
//   	Principals: []*string{
//   		jsii.String("principals"),
//   	},
//   	Status: jsii.String("status"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html
//
type CfnGrantPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnGrantMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnGrantPropsMixin
type jsiiProxy_CfnGrantPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnGrantPropsMixin) Props() *CfnGrantMixinProps {
	var returns *CfnGrantMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGrantPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::LicenseManager::Grant`.
func NewCfnGrantPropsMixin(props *CfnGrantMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnGrantPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnGrantPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnGrantPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_licensemanager.CfnGrantPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::LicenseManager::Grant`.
func NewCfnGrantPropsMixin_Override(c CfnGrantPropsMixin, props *CfnGrantMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_licensemanager.CfnGrantPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnGrantPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnGrantPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_licensemanager.CfnGrantPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnGrantPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_licensemanager.CfnGrantPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnGrantPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnGrantPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

