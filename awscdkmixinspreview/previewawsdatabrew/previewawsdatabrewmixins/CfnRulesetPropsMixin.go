package previewawsdatabrewmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsdatabrew/previewawsdatabrewmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies a new ruleset that can be used in a profile job to validate the data quality of a dataset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRulesetPropsMixin := awscdkmixinspreview.Mixins.NewCfnRulesetPropsMixin(&CfnRulesetMixinProps{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Rules: []interface{}{
//   		&RuleProperty{
//   			CheckExpression: jsii.String("checkExpression"),
//   			ColumnSelectors: []interface{}{
//   				&ColumnSelectorProperty{
//   					Name: jsii.String("name"),
//   					Regex: jsii.String("regex"),
//   				},
//   			},
//   			Disabled: jsii.Boolean(false),
//   			Name: jsii.String("name"),
//   			SubstitutionMap: []interface{}{
//   				&SubstitutionValueProperty{
//   					Value: jsii.String("value"),
//   					ValueReference: jsii.String("valueReference"),
//   				},
//   			},
//   			Threshold: &ThresholdProperty{
//   				Type: jsii.String("type"),
//   				Unit: jsii.String("unit"),
//   				Value: jsii.Number(123),
//   			},
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TargetArn: jsii.String("targetArn"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-ruleset.html
//
type CfnRulesetPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnRulesetMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnRulesetPropsMixin
type jsiiProxy_CfnRulesetPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnRulesetPropsMixin) Props() *CfnRulesetMixinProps {
	var returns *CfnRulesetMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRulesetPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::DataBrew::Ruleset`.
func NewCfnRulesetPropsMixin(props *CfnRulesetMixinProps, options *mixins.CfnPropertyMixinOptions) CfnRulesetPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnRulesetPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnRulesetPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnRulesetPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::DataBrew::Ruleset`.
func NewCfnRulesetPropsMixin_Override(c CfnRulesetPropsMixin, props *CfnRulesetMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnRulesetPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnRulesetPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnRulesetPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnRulesetPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRulesetPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnRulesetPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnRulesetPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnRulesetPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

