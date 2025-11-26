package previewawslicensemanagermixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslicensemanager/previewawslicensemanagermixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies a grant.
//
// A grant shares the use of license entitlements with specific AWS accounts . For more information, see [Granted licenses](https://docs.aws.amazon.com/license-manager/latest/userguide/granted-licenses.html) in the *License Manager User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnGrantPropsMixin := awscdkmixinspreview.Mixins.NewCfnGrantPropsMixin(&CfnGrantMixinProps{
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
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html
//
type CfnGrantPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnGrantMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnGrantPropsMixin
type jsiiProxy_CfnGrantPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
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

func (j *jsiiProxy_CfnGrantPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::LicenseManager::Grant`.
func NewCfnGrantPropsMixin(props *CfnGrantMixinProps, options *mixins.CfnPropertyMixinOptions) CfnGrantPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnGrantPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnGrantPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_licensemanager.mixins.CfnGrantPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::LicenseManager::Grant`.
func NewCfnGrantPropsMixin_Override(c CfnGrantPropsMixin, props *CfnGrantMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_licensemanager.mixins.CfnGrantPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnGrantPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnGrantPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_licensemanager.mixins.CfnGrantPropsMixin",
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
		"@aws-cdk/mixins-preview.aws_licensemanager.mixins.CfnGrantPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnGrantPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

