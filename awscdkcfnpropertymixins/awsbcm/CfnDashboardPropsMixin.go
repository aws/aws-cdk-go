package awsbcm

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsbcm/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Definition of AWS::BCM::Dashboard Resource Type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var costAndUsageExpressionProperty_ CostAndUsageExpressionProperty
//   var expressionProperty_ ExpressionProperty
//   var mergeStrategy IMergeStrategy
//   var table interface{}
//
//   cfnDashboardPropsMixin := awscdkcfnpropertymixins.Aws_bcm.NewCfnDashboardPropsMixin(&CfnDashboardMixinProps{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Widgets: []interface{}{
//   		&WidgetProperty{
//   			Configs: []interface{}{
//   				&WidgetConfigProperty{
//   					DisplayConfig: &DisplayConfigProperty{
//   						Graph: map[string]interface{}{
//   							"graphKey": &GraphDisplayConfigProperty{
//   								"visualType": jsii.String("visualType"),
//   							},
//   						},
//   						Table: table,
//   					},
//   					QueryParameters: &QueryParametersProperty{
//   						CostAndUsage: &CostAndUsageQueryProperty{
//   							Filter: &CostAndUsageExpressionProperty{
//   								And: []interface{}{
//   									costAndUsageExpressionProperty_,
//   								},
//   								CostCategories: &CostCategoryValuesProperty{
//   									Key: jsii.String("key"),
//   									MatchOptions: []*string{
//   										jsii.String("matchOptions"),
//   									},
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   								Dimensions: &DimensionValuesProperty{
//   									Key: jsii.String("key"),
//   									MatchOptions: []*string{
//   										jsii.String("matchOptions"),
//   									},
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   								Not: costAndUsageExpressionProperty_,
//   								Or: []interface{}{
//   									costAndUsageExpressionProperty_,
//   								},
//   								Tags: &TagValuesProperty{
//   									Key: jsii.String("key"),
//   									MatchOptions: []*string{
//   										jsii.String("matchOptions"),
//   									},
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   							},
//   							Granularity: jsii.String("granularity"),
//   							GroupBy: []interface{}{
//   								&GroupDefinitionProperty{
//   									Key: jsii.String("key"),
//   									Type: jsii.String("type"),
//   								},
//   							},
//   							Metrics: []*string{
//   								jsii.String("metrics"),
//   							},
//   							TimeRange: &DateTimeRangeProperty{
//   								EndTime: &DateTimeValueProperty{
//   									Type: jsii.String("type"),
//   									Value: jsii.String("value"),
//   								},
//   								StartTime: &DateTimeValueProperty{
//   									Type: jsii.String("type"),
//   									Value: jsii.String("value"),
//   								},
//   							},
//   						},
//   						ReservationCoverage: &ReservationCoverageQueryProperty{
//   							Filter: &ExpressionProperty{
//   								And: []interface{}{
//   									expressionProperty_,
//   								},
//   								CostCategories: &CostCategoryValuesProperty{
//   									Key: jsii.String("key"),
//   									MatchOptions: []*string{
//   										jsii.String("matchOptions"),
//   									},
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   								Dimensions: &DimensionValuesProperty{
//   									Key: jsii.String("key"),
//   									MatchOptions: []*string{
//   										jsii.String("matchOptions"),
//   									},
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   								Not: expressionProperty_,
//   								Tags: &TagValuesProperty{
//   									Key: jsii.String("key"),
//   									MatchOptions: []*string{
//   										jsii.String("matchOptions"),
//   									},
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   							},
//   							Granularity: jsii.String("granularity"),
//   							GroupBy: []interface{}{
//   								&GroupDefinitionProperty{
//   									Key: jsii.String("key"),
//   									Type: jsii.String("type"),
//   								},
//   							},
//   							Metrics: []*string{
//   								jsii.String("metrics"),
//   							},
//   							TimeRange: &DateTimeRangeProperty{
//   								EndTime: &DateTimeValueProperty{
//   									Type: jsii.String("type"),
//   									Value: jsii.String("value"),
//   								},
//   								StartTime: &DateTimeValueProperty{
//   									Type: jsii.String("type"),
//   									Value: jsii.String("value"),
//   								},
//   							},
//   						},
//   						ReservationUtilization: &ReservationUtilizationQueryProperty{
//   							Filter: &ExpressionProperty{
//   								And: []interface{}{
//   									expressionProperty_,
//   								},
//   								CostCategories: &CostCategoryValuesProperty{
//   									Key: jsii.String("key"),
//   									MatchOptions: []*string{
//   										jsii.String("matchOptions"),
//   									},
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   								Dimensions: &DimensionValuesProperty{
//   									Key: jsii.String("key"),
//   									MatchOptions: []*string{
//   										jsii.String("matchOptions"),
//   									},
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   								Not: expressionProperty_,
//   								Tags: &TagValuesProperty{
//   									Key: jsii.String("key"),
//   									MatchOptions: []*string{
//   										jsii.String("matchOptions"),
//   									},
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   							},
//   							Granularity: jsii.String("granularity"),
//   							GroupBy: []interface{}{
//   								&GroupDefinitionProperty{
//   									Key: jsii.String("key"),
//   									Type: jsii.String("type"),
//   								},
//   							},
//   							TimeRange: &DateTimeRangeProperty{
//   								EndTime: &DateTimeValueProperty{
//   									Type: jsii.String("type"),
//   									Value: jsii.String("value"),
//   								},
//   								StartTime: &DateTimeValueProperty{
//   									Type: jsii.String("type"),
//   									Value: jsii.String("value"),
//   								},
//   							},
//   						},
//   						SavingsPlansCoverage: &SavingsPlansCoverageQueryProperty{
//   							Filter: &ExpressionProperty{
//   								And: []interface{}{
//   									expressionProperty_,
//   								},
//   								CostCategories: &CostCategoryValuesProperty{
//   									Key: jsii.String("key"),
//   									MatchOptions: []*string{
//   										jsii.String("matchOptions"),
//   									},
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   								Dimensions: &DimensionValuesProperty{
//   									Key: jsii.String("key"),
//   									MatchOptions: []*string{
//   										jsii.String("matchOptions"),
//   									},
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   								Not: expressionProperty_,
//   								Tags: &TagValuesProperty{
//   									Key: jsii.String("key"),
//   									MatchOptions: []*string{
//   										jsii.String("matchOptions"),
//   									},
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   							},
//   							Granularity: jsii.String("granularity"),
//   							GroupBy: []interface{}{
//   								&GroupDefinitionProperty{
//   									Key: jsii.String("key"),
//   									Type: jsii.String("type"),
//   								},
//   							},
//   							Metrics: []*string{
//   								jsii.String("metrics"),
//   							},
//   							TimeRange: &DateTimeRangeProperty{
//   								EndTime: &DateTimeValueProperty{
//   									Type: jsii.String("type"),
//   									Value: jsii.String("value"),
//   								},
//   								StartTime: &DateTimeValueProperty{
//   									Type: jsii.String("type"),
//   									Value: jsii.String("value"),
//   								},
//   							},
//   						},
//   						SavingsPlansUtilization: &SavingsPlansUtilizationQueryProperty{
//   							Filter: &ExpressionProperty{
//   								And: []interface{}{
//   									expressionProperty_,
//   								},
//   								CostCategories: &CostCategoryValuesProperty{
//   									Key: jsii.String("key"),
//   									MatchOptions: []*string{
//   										jsii.String("matchOptions"),
//   									},
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   								Dimensions: &DimensionValuesProperty{
//   									Key: jsii.String("key"),
//   									MatchOptions: []*string{
//   										jsii.String("matchOptions"),
//   									},
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   								Not: expressionProperty_,
//   								Tags: &TagValuesProperty{
//   									Key: jsii.String("key"),
//   									MatchOptions: []*string{
//   										jsii.String("matchOptions"),
//   									},
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   							},
//   							Granularity: jsii.String("granularity"),
//   							TimeRange: &DateTimeRangeProperty{
//   								EndTime: &DateTimeValueProperty{
//   									Type: jsii.String("type"),
//   									Value: jsii.String("value"),
//   								},
//   								StartTime: &DateTimeValueProperty{
//   									Type: jsii.String("type"),
//   									Value: jsii.String("value"),
//   								},
//   							},
//   						},
//   					},
//   				},
//   			},
//   			Description: jsii.String("description"),
//   			Height: jsii.Number(123),
//   			HorizontalOffset: jsii.Number(123),
//   			Title: jsii.String("title"),
//   			Width: jsii.Number(123),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bcm-dashboard.html
//
type CfnDashboardPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnDashboardMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDashboardPropsMixin
type jsiiProxy_CfnDashboardPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnDashboardPropsMixin) Props() *CfnDashboardMixinProps {
	var returns *CfnDashboardMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDashboardPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::BCM::Dashboard`.
func NewCfnDashboardPropsMixin(props *CfnDashboardMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnDashboardPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDashboardPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDashboardPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_bcm.CfnDashboardPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::BCM::Dashboard`.
func NewCfnDashboardPropsMixin_Override(c CfnDashboardPropsMixin, props *CfnDashboardMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_bcm.CfnDashboardPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnDashboardPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDashboardPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_bcm.CfnDashboardPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDashboardPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_bcm.CfnDashboardPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDashboardPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnDashboardPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

