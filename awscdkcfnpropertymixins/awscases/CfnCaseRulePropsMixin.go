package awscases

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awscases/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a new case rule.
//
// In the Amazon Connect admin website, case rules are known as *case field conditions* . For more information about case field conditions, see [Add case field conditions to a case template](https://docs.aws.amazon.com/connect/latest/adminguide/case-field-conditions.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var emptyValue interface{}
//   var mergeStrategy IMergeStrategy
//
//   cfnCaseRulePropsMixin := awscdkcfnpropertymixins.Aws_cases.NewCfnCaseRulePropsMixin(&CfnCaseRuleMixinProps{
//   	Description: jsii.String("description"),
//   	DomainId: jsii.String("domainId"),
//   	Name: jsii.String("name"),
//   	Rule: &CaseRuleDetailsProperty{
//   		Hidden: &HiddenCaseRuleProperty{
//   			Conditions: []interface{}{
//   				&BooleanConditionProperty{
//   					EqualTo: &BooleanOperandsProperty{
//   						OperandOne: &OperandOneProperty{
//   							FieldId: jsii.String("fieldId"),
//   						},
//   						OperandTwo: &OperandTwoProperty{
//   							BooleanValue: jsii.Boolean(false),
//   							DoubleValue: jsii.Number(123),
//   							EmptyValue: emptyValue,
//   							StringValue: jsii.String("stringValue"),
//   						},
//   						Result: jsii.Boolean(false),
//   					},
//   					NotEqualTo: &BooleanOperandsProperty{
//   						OperandOne: &OperandOneProperty{
//   							FieldId: jsii.String("fieldId"),
//   						},
//   						OperandTwo: &OperandTwoProperty{
//   							BooleanValue: jsii.Boolean(false),
//   							DoubleValue: jsii.Number(123),
//   							EmptyValue: emptyValue,
//   							StringValue: jsii.String("stringValue"),
//   						},
//   						Result: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   			DefaultValue: jsii.Boolean(false),
//   		},
//   		Required: &RequiredCaseRuleProperty{
//   			Conditions: []interface{}{
//   				&BooleanConditionProperty{
//   					EqualTo: &BooleanOperandsProperty{
//   						OperandOne: &OperandOneProperty{
//   							FieldId: jsii.String("fieldId"),
//   						},
//   						OperandTwo: &OperandTwoProperty{
//   							BooleanValue: jsii.Boolean(false),
//   							DoubleValue: jsii.Number(123),
//   							EmptyValue: emptyValue,
//   							StringValue: jsii.String("stringValue"),
//   						},
//   						Result: jsii.Boolean(false),
//   					},
//   					NotEqualTo: &BooleanOperandsProperty{
//   						OperandOne: &OperandOneProperty{
//   							FieldId: jsii.String("fieldId"),
//   						},
//   						OperandTwo: &OperandTwoProperty{
//   							BooleanValue: jsii.Boolean(false),
//   							DoubleValue: jsii.Number(123),
//   							EmptyValue: emptyValue,
//   							StringValue: jsii.String("stringValue"),
//   						},
//   						Result: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   			DefaultValue: jsii.Boolean(false),
//   		},
//   	},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cases-caserule.html
//
type CfnCaseRulePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnCaseRuleMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnCaseRulePropsMixin
type jsiiProxy_CfnCaseRulePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnCaseRulePropsMixin) Props() *CfnCaseRuleMixinProps {
	var returns *CfnCaseRuleMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCaseRulePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Cases::CaseRule`.
func NewCfnCaseRulePropsMixin(props *CfnCaseRuleMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnCaseRulePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnCaseRulePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCaseRulePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_cases.CfnCaseRulePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Cases::CaseRule`.
func NewCfnCaseRulePropsMixin_Override(c CfnCaseRulePropsMixin, props *CfnCaseRuleMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_cases.CfnCaseRulePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnCaseRulePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCaseRulePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_cases.CfnCaseRulePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCaseRulePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_cases.CfnCaseRulePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCaseRulePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnCaseRulePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

