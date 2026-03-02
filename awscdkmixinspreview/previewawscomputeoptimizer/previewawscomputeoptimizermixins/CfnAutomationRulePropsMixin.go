package previewawscomputeoptimizermixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawscomputeoptimizer/previewawscomputeoptimizermixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates an AWS Compute Optimizer automation rule that automatically implements recommended actions based on your defined criteria and schedule.
//
// Automation rules are global resources that manage automated actions across all AWS Regions where Compute Optimizer Automation is available. Organization-level rules can only be created by the management account or delegated administrator.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAutomationRulePropsMixin := awscdkmixinspreview.Mixins.NewCfnAutomationRulePropsMixin(&CfnAutomationRuleMixinProps{
//   	Criteria: &CriteriaProperty{
//   		EbsVolumeSizeInGib: []interface{}{
//   			&IntegerCriteriaConditionProperty{
//   				Comparison: jsii.String("comparison"),
//   				Values: []interface{}{
//   					jsii.Number(123),
//   				},
//   			},
//   		},
//   		EbsVolumeType: []interface{}{
//   			&StringCriteriaConditionProperty{
//   				Comparison: jsii.String("comparison"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   		EstimatedMonthlySavings: []interface{}{
//   			&DoubleCriteriaConditionProperty{
//   				Comparison: jsii.String("comparison"),
//   				Values: []interface{}{
//   					jsii.Number(123),
//   				},
//   			},
//   		},
//   		LookBackPeriodInDays: []interface{}{
//   			&IntegerCriteriaConditionProperty{
//   				Comparison: jsii.String("comparison"),
//   				Values: []interface{}{
//   					jsii.Number(123),
//   				},
//   			},
//   		},
//   		Region: []interface{}{
//   			&StringCriteriaConditionProperty{
//   				Comparison: jsii.String("comparison"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   		ResourceArn: []interface{}{
//   			&StringCriteriaConditionProperty{
//   				Comparison: jsii.String("comparison"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   		ResourceTag: []interface{}{
//   			&ResourceTagsCriteriaConditionProperty{
//   				Comparison: jsii.String("comparison"),
//   				Key: jsii.String("key"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   		RestartNeeded: []interface{}{
//   			&StringCriteriaConditionProperty{
//   				Comparison: jsii.String("comparison"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	OrganizationConfiguration: &OrganizationConfigurationProperty{
//   		AccountIds: []*string{
//   			jsii.String("accountIds"),
//   		},
//   		RuleApplyOrder: jsii.String("ruleApplyOrder"),
//   	},
//   	Priority: jsii.String("priority"),
//   	RecommendedActionTypes: []*string{
//   		jsii.String("recommendedActionTypes"),
//   	},
//   	RuleType: jsii.String("ruleType"),
//   	Schedule: &ScheduleProperty{
//   		ExecutionWindowInMinutes: jsii.Number(123),
//   		ScheduleExpression: jsii.String("scheduleExpression"),
//   		ScheduleExpressionTimezone: jsii.String("scheduleExpressionTimezone"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-computeoptimizer-automationrule.html
//
type CfnAutomationRulePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnAutomationRuleMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAutomationRulePropsMixin
type jsiiProxy_CfnAutomationRulePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnAutomationRulePropsMixin) Props() *CfnAutomationRuleMixinProps {
	var returns *CfnAutomationRuleMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutomationRulePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ComputeOptimizer::AutomationRule`.
func NewCfnAutomationRulePropsMixin(props *CfnAutomationRuleMixinProps, options *mixins.CfnPropertyMixinOptions) CfnAutomationRulePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAutomationRulePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAutomationRulePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_computeoptimizer.mixins.CfnAutomationRulePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ComputeOptimizer::AutomationRule`.
func NewCfnAutomationRulePropsMixin_Override(c CfnAutomationRulePropsMixin, props *CfnAutomationRuleMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_computeoptimizer.mixins.CfnAutomationRulePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnAutomationRulePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAutomationRulePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_computeoptimizer.mixins.CfnAutomationRulePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAutomationRulePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_computeoptimizer.mixins.CfnAutomationRulePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAutomationRulePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnAutomationRulePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

