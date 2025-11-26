package previewawscemixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsce/previewawscemixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::CE::CostCategory` resource creates groupings of cost that you can use across products in the Billing and Cost Management console, such as Cost Explorer and AWS Budgets.
//
// For more information, see [Managing Your Costs with Cost Categories](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/manage-cost-categories.html) in the *Billing and Cost Management User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCostCategoryPropsMixin := awscdkmixinspreview.Mixins.NewCfnCostCategoryPropsMixin(&CfnCostCategoryMixinProps{
//   	DefaultValue: jsii.String("defaultValue"),
//   	Name: jsii.String("name"),
//   	Rules: jsii.String("rules"),
//   	RuleVersion: jsii.String("ruleVersion"),
//   	SplitChargeRules: jsii.String("splitChargeRules"),
//   	Tags: []ResourceTagProperty{
//   		&ResourceTagProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ce-costcategory.html
//
type CfnCostCategoryPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnCostCategoryMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnCostCategoryPropsMixin
type jsiiProxy_CfnCostCategoryPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnCostCategoryPropsMixin) Props() *CfnCostCategoryMixinProps {
	var returns *CfnCostCategoryMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCostCategoryPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CE::CostCategory`.
func NewCfnCostCategoryPropsMixin(props *CfnCostCategoryMixinProps, options *mixins.CfnPropertyMixinOptions) CfnCostCategoryPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnCostCategoryPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCostCategoryPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ce.mixins.CfnCostCategoryPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CE::CostCategory`.
func NewCfnCostCategoryPropsMixin_Override(c CfnCostCategoryPropsMixin, props *CfnCostCategoryMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ce.mixins.CfnCostCategoryPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnCostCategoryPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCostCategoryPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ce.mixins.CfnCostCategoryPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCostCategoryPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_ce.mixins.CfnCostCategoryPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCostCategoryPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnCostCategoryPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

