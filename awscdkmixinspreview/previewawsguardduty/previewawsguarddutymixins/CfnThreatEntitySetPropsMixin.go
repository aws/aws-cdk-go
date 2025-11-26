package previewawsguarddutymixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsguardduty/previewawsguarddutymixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::GuardDuty::ThreatEntitySet` resource helps you create a list of known malicious IP addresses and domain names in your AWS environment.
//
// Once you activate this list, GuardDuty will use the entries in this list as an additional source of threat detection and generate findings when there is an activity associated with these known malicious IP addresses and domain names. GuardDuty continues to monitor independently of this custom threat entity set.
//
// Only the users of the GuardDuty administrator account can manage this list. These settings automatically apply to the member accounts.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnThreatEntitySetPropsMixin := awscdkmixinspreview.Mixins.NewCfnThreatEntitySetPropsMixin(&CfnThreatEntitySetMixinProps{
//   	Activate: jsii.Boolean(false),
//   	DetectorId: jsii.String("detectorId"),
//   	ExpectedBucketOwner: jsii.String("expectedBucketOwner"),
//   	Format: jsii.String("format"),
//   	Location: jsii.String("location"),
//   	Name: jsii.String("name"),
//   	Tags: []TagItemProperty{
//   		&TagItemProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-threatentityset.html
//
type CfnThreatEntitySetPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnThreatEntitySetMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnThreatEntitySetPropsMixin
type jsiiProxy_CfnThreatEntitySetPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnThreatEntitySetPropsMixin) Props() *CfnThreatEntitySetMixinProps {
	var returns *CfnThreatEntitySetMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnThreatEntitySetPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::GuardDuty::ThreatEntitySet`.
func NewCfnThreatEntitySetPropsMixin(props *CfnThreatEntitySetMixinProps, options *mixins.CfnPropertyMixinOptions) CfnThreatEntitySetPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnThreatEntitySetPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnThreatEntitySetPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_guardduty.mixins.CfnThreatEntitySetPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::GuardDuty::ThreatEntitySet`.
func NewCfnThreatEntitySetPropsMixin_Override(c CfnThreatEntitySetPropsMixin, props *CfnThreatEntitySetMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_guardduty.mixins.CfnThreatEntitySetPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnThreatEntitySetPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnThreatEntitySetPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_guardduty.mixins.CfnThreatEntitySetPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnThreatEntitySetPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_guardduty.mixins.CfnThreatEntitySetPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnThreatEntitySetPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnThreatEntitySetPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

