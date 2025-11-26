package previewawsrbinmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsrbin/previewawsrbinmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a Recycle Bin retention rule. You can create two types of retention rules:.
//
// - *Tag-level retention rules* - These retention rules use resource tags to identify the resources to protect. For each retention rule, you specify one or more tag key and value pairs. Resources (of the specified type) that have at least one of these tag key and value pairs are automatically retained in the Recycle Bin upon deletion. Use this type of retention rule to protect specific resources in your account based on their tags.
// - *Region-level retention rules* - These retention rules, by default, apply to all of the resources (of the specified type) in the Region, even if the resources are not tagged. However, you can specify exclusion tags to exclude resources that have specific tags. Use this type of retention rule to protect all resources of a specific type in a Region.
//
// For more information, see [Create Recycle Bin retention rules](https://docs.aws.amazon.com/ebs/latest/userguide/recycle-bin.html) in the *Amazon EBS User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRulePropsMixin := awscdkmixinspreview.Mixins.NewCfnRulePropsMixin(&CfnRuleMixinProps{
//   	Description: jsii.String("description"),
//   	ExcludeResourceTags: []interface{}{
//   		&ResourceTagProperty{
//   			ResourceTagKey: jsii.String("resourceTagKey"),
//   			ResourceTagValue: jsii.String("resourceTagValue"),
//   		},
//   	},
//   	LockConfiguration: &UnlockDelayProperty{
//   		UnlockDelayUnit: jsii.String("unlockDelayUnit"),
//   		UnlockDelayValue: jsii.Number(123),
//   	},
//   	ResourceTags: []interface{}{
//   		&ResourceTagProperty{
//   			ResourceTagKey: jsii.String("resourceTagKey"),
//   			ResourceTagValue: jsii.String("resourceTagValue"),
//   		},
//   	},
//   	ResourceType: jsii.String("resourceType"),
//   	RetentionPeriod: &RetentionPeriodProperty{
//   		RetentionPeriodUnit: jsii.String("retentionPeriodUnit"),
//   		RetentionPeriodValue: jsii.Number(123),
//   	},
//   	Status: jsii.String("status"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rbin-rule.html
//
type CfnRulePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnRuleMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnRulePropsMixin
type jsiiProxy_CfnRulePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnRulePropsMixin) Props() *CfnRuleMixinProps {
	var returns *CfnRuleMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRulePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Rbin::Rule`.
func NewCfnRulePropsMixin(props *CfnRuleMixinProps, options *mixins.CfnPropertyMixinOptions) CfnRulePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnRulePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnRulePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_rbin.mixins.CfnRulePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Rbin::Rule`.
func NewCfnRulePropsMixin_Override(c CfnRulePropsMixin, props *CfnRuleMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_rbin.mixins.CfnRulePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnRulePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnRulePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_rbin.mixins.CfnRulePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRulePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_rbin.mixins.CfnRulePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnRulePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnRulePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

