package previewawsconnectmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsconnect/previewawsconnectmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates an evaluation form for the specified Amazon Connect instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var evaluationFormSectionProperty_ EvaluationFormSectionProperty
//
//   cfnEvaluationFormPropsMixin := awscdkmixinspreview.Mixins.NewCfnEvaluationFormPropsMixin(&CfnEvaluationFormMixinProps{
//   	AutoEvaluationConfiguration: &AutoEvaluationConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	Description: jsii.String("description"),
//   	InstanceArn: jsii.String("instanceArn"),
//   	Items: []interface{}{
//   		&EvaluationFormBaseItemProperty{
//   			Section: &EvaluationFormSectionProperty{
//   				Instructions: jsii.String("instructions"),
//   				Items: []interface{}{
//   					&EvaluationFormItemProperty{
//   						Question: &EvaluationFormQuestionProperty{
//   							Enablement: &EvaluationFormItemEnablementConfigurationProperty{
//   								Action: jsii.String("action"),
//   								Condition: &EvaluationFormItemEnablementConditionProperty{
//   									Operands: []interface{}{
//   										&EvaluationFormItemEnablementConditionOperandProperty{
//   											Expression: &EvaluationFormItemEnablementExpressionProperty{
//   												Comparator: jsii.String("comparator"),
//   												Source: &EvaluationFormItemEnablementSourceProperty{
//   													RefId: jsii.String("refId"),
//   													Type: jsii.String("type"),
//   												},
//   												Values: []interface{}{
//   													&EvaluationFormItemEnablementSourceValueProperty{
//   														RefId: jsii.String("refId"),
//   														Type: jsii.String("type"),
//   													},
//   												},
//   											},
//   										},
//   									},
//   									Operator: jsii.String("operator"),
//   								},
//   								DefaultAction: jsii.String("defaultAction"),
//   							},
//   							Instructions: jsii.String("instructions"),
//   							NotApplicableEnabled: jsii.Boolean(false),
//   							QuestionType: jsii.String("questionType"),
//   							QuestionTypeProperties: &EvaluationFormQuestionTypePropertiesProperty{
//   								Numeric: &EvaluationFormNumericQuestionPropertiesProperty{
//   									Automation: &EvaluationFormNumericQuestionAutomationProperty{
//   										AnswerSource: &EvaluationFormQuestionAutomationAnswerSourceProperty{
//   											SourceType: jsii.String("sourceType"),
//   										},
//   										PropertyValue: &NumericQuestionPropertyValueAutomationProperty{
//   											Label: jsii.String("label"),
//   										},
//   									},
//   									MaxValue: jsii.Number(123),
//   									MinValue: jsii.Number(123),
//   									Options: []interface{}{
//   										&EvaluationFormNumericQuestionOptionProperty{
//   											AutomaticFail: jsii.Boolean(false),
//   											AutomaticFailConfiguration: &AutomaticFailConfigurationProperty{
//   												TargetSection: jsii.String("targetSection"),
//   											},
//   											MaxValue: jsii.Number(123),
//   											MinValue: jsii.Number(123),
//   											Score: jsii.Number(123),
//   										},
//   									},
//   								},
//   								SingleSelect: &EvaluationFormSingleSelectQuestionPropertiesProperty{
//   									Automation: &EvaluationFormSingleSelectQuestionAutomationProperty{
//   										AnswerSource: &EvaluationFormQuestionAutomationAnswerSourceProperty{
//   											SourceType: jsii.String("sourceType"),
//   										},
//   										DefaultOptionRefId: jsii.String("defaultOptionRefId"),
//   										Options: []interface{}{
//   											&EvaluationFormSingleSelectQuestionAutomationOptionProperty{
//   												RuleCategory: &SingleSelectQuestionRuleCategoryAutomationProperty{
//   													Category: jsii.String("category"),
//   													Condition: jsii.String("condition"),
//   													OptionRefId: jsii.String("optionRefId"),
//   												},
//   											},
//   										},
//   									},
//   									DisplayAs: jsii.String("displayAs"),
//   									Options: []interface{}{
//   										&EvaluationFormSingleSelectQuestionOptionProperty{
//   											AutomaticFail: jsii.Boolean(false),
//   											AutomaticFailConfiguration: &AutomaticFailConfigurationProperty{
//   												TargetSection: jsii.String("targetSection"),
//   											},
//   											RefId: jsii.String("refId"),
//   											Score: jsii.Number(123),
//   											Text: jsii.String("text"),
//   										},
//   									},
//   								},
//   								Text: &EvaluationFormTextQuestionPropertiesProperty{
//   									Automation: &EvaluationFormTextQuestionAutomationProperty{
//   										AnswerSource: &EvaluationFormQuestionAutomationAnswerSourceProperty{
//   											SourceType: jsii.String("sourceType"),
//   										},
//   									},
//   								},
//   							},
//   							RefId: jsii.String("refId"),
//   							Title: jsii.String("title"),
//   							Weight: jsii.Number(123),
//   						},
//   						Section: evaluationFormSectionProperty_,
//   					},
//   				},
//   				RefId: jsii.String("refId"),
//   				Title: jsii.String("title"),
//   				Weight: jsii.Number(123),
//   			},
//   		},
//   	},
//   	ScoringStrategy: &ScoringStrategyProperty{
//   		Mode: jsii.String("mode"),
//   		Status: jsii.String("status"),
//   	},
//   	Status: jsii.String("status"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Title: jsii.String("title"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-evaluationform.html
//
type CfnEvaluationFormPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnEvaluationFormMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnEvaluationFormPropsMixin
type jsiiProxy_CfnEvaluationFormPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnEvaluationFormPropsMixin) Props() *CfnEvaluationFormMixinProps {
	var returns *CfnEvaluationFormMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEvaluationFormPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Connect::EvaluationForm`.
func NewCfnEvaluationFormPropsMixin(props *CfnEvaluationFormMixinProps, options *mixins.CfnPropertyMixinOptions) CfnEvaluationFormPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnEvaluationFormPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnEvaluationFormPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_connect.mixins.CfnEvaluationFormPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Connect::EvaluationForm`.
func NewCfnEvaluationFormPropsMixin_Override(c CfnEvaluationFormPropsMixin, props *CfnEvaluationFormMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_connect.mixins.CfnEvaluationFormPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnEvaluationFormPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEvaluationFormPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_connect.mixins.CfnEvaluationFormPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEvaluationFormPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_connect.mixins.CfnEvaluationFormPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnEvaluationFormPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnEvaluationFormPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

