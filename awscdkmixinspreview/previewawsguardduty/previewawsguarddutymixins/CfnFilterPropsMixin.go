package previewawsguarddutymixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsguardduty/previewawsguarddutymixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::GuardDuty::Filter` resource specifies a new filter defined by the provided `findingCriteria` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var criterion interface{}
//
//   cfnFilterPropsMixin := awscdkmixinspreview.Mixins.NewCfnFilterPropsMixin(&CfnFilterMixinProps{
//   	Action: jsii.String("action"),
//   	Description: jsii.String("description"),
//   	DetectorId: jsii.String("detectorId"),
//   	FindingCriteria: &FindingCriteriaProperty{
//   		Criterion: criterion,
//   		ItemType: &ConditionProperty{
//   			Eq: []*string{
//   				jsii.String("eq"),
//   			},
//   			EqualTo: []*string{
//   				jsii.String("equalTo"),
//   			},
//   			GreaterThan: jsii.Number(123),
//   			GreaterThanOrEqual: jsii.Number(123),
//   			Gt: jsii.Number(123),
//   			Gte: jsii.Number(123),
//   			LessThan: jsii.Number(123),
//   			LessThanOrEqual: jsii.Number(123),
//   			Lt: jsii.Number(123),
//   			Lte: jsii.Number(123),
//   			Neq: []*string{
//   				jsii.String("neq"),
//   			},
//   			NotEquals: []*string{
//   				jsii.String("notEquals"),
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Rank: jsii.Number(123),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-filter.html
//
type CfnFilterPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnFilterMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnFilterPropsMixin
type jsiiProxy_CfnFilterPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnFilterPropsMixin) Props() *CfnFilterMixinProps {
	var returns *CfnFilterMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFilterPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::GuardDuty::Filter`.
func NewCfnFilterPropsMixin(props *CfnFilterMixinProps, options *mixins.CfnPropertyMixinOptions) CfnFilterPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnFilterPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnFilterPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_guardduty.mixins.CfnFilterPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::GuardDuty::Filter`.
func NewCfnFilterPropsMixin_Override(c CfnFilterPropsMixin, props *CfnFilterMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_guardduty.mixins.CfnFilterPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnFilterPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnFilterPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_guardduty.mixins.CfnFilterPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFilterPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_guardduty.mixins.CfnFilterPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnFilterPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnFilterPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

