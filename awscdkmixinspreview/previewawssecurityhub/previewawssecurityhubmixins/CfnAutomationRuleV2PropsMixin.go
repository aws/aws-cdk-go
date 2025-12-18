package previewawssecurityhubmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawssecurityhub/previewawssecurityhubmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a V2 automation rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAutomationRuleV2PropsMixin := awscdkmixinspreview.Mixins.NewCfnAutomationRuleV2PropsMixin(&CfnAutomationRuleV2MixinProps{
//   	Actions: []interface{}{
//   		&AutomationRulesActionV2Property{
//   			ExternalIntegrationConfiguration: &ExternalIntegrationConfigurationProperty{
//   				ConnectorArn: jsii.String("connectorArn"),
//   			},
//   			FindingFieldsUpdate: &AutomationRulesFindingFieldsUpdateV2Property{
//   				Comment: jsii.String("comment"),
//   				SeverityId: jsii.Number(123),
//   				StatusId: jsii.Number(123),
//   			},
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	Criteria: &CriteriaProperty{
//   		OcsfFindingCriteria: &OcsfFindingFiltersProperty{
//   			CompositeFilters: []interface{}{
//   				&CompositeFilterProperty{
//   					BooleanFilters: []interface{}{
//   						&OcsfBooleanFilterProperty{
//   							FieldName: jsii.String("fieldName"),
//   							Filter: &BooleanFilterProperty{
//   								Value: jsii.Boolean(false),
//   							},
//   						},
//   					},
//   					DateFilters: []interface{}{
//   						&OcsfDateFilterProperty{
//   							FieldName: jsii.String("fieldName"),
//   							Filter: &DateFilterProperty{
//   								DateRange: &DateRangeProperty{
//   									Unit: jsii.String("unit"),
//   									Value: jsii.Number(123),
//   								},
//   								End: jsii.String("end"),
//   								Start: jsii.String("start"),
//   							},
//   						},
//   					},
//   					MapFilters: []interface{}{
//   						&OcsfMapFilterProperty{
//   							FieldName: jsii.String("fieldName"),
//   							Filter: &MapFilterProperty{
//   								Comparison: jsii.String("comparison"),
//   								Key: jsii.String("key"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   					},
//   					NumberFilters: []interface{}{
//   						&OcsfNumberFilterProperty{
//   							FieldName: jsii.String("fieldName"),
//   							Filter: &NumberFilterProperty{
//   								Eq: jsii.Number(123),
//   								Gte: jsii.Number(123),
//   								Lte: jsii.Number(123),
//   							},
//   						},
//   					},
//   					Operator: jsii.String("operator"),
//   					StringFilters: []interface{}{
//   						&OcsfStringFilterProperty{
//   							FieldName: jsii.String("fieldName"),
//   							Filter: &StringFilterProperty{
//   								Comparison: jsii.String("comparison"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   			CompositeOperator: jsii.String("compositeOperator"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	RuleName: jsii.String("ruleName"),
//   	RuleOrder: jsii.Number(123),
//   	RuleStatus: jsii.String("ruleStatus"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-automationrulev2.html
//
type CfnAutomationRuleV2PropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnAutomationRuleV2MixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAutomationRuleV2PropsMixin
type jsiiProxy_CfnAutomationRuleV2PropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnAutomationRuleV2PropsMixin) Props() *CfnAutomationRuleV2MixinProps {
	var returns *CfnAutomationRuleV2MixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutomationRuleV2PropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SecurityHub::AutomationRuleV2`.
func NewCfnAutomationRuleV2PropsMixin(props *CfnAutomationRuleV2MixinProps, options *mixins.CfnPropertyMixinOptions) CfnAutomationRuleV2PropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAutomationRuleV2PropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAutomationRuleV2PropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_securityhub.mixins.CfnAutomationRuleV2PropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SecurityHub::AutomationRuleV2`.
func NewCfnAutomationRuleV2PropsMixin_Override(c CfnAutomationRuleV2PropsMixin, props *CfnAutomationRuleV2MixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_securityhub.mixins.CfnAutomationRuleV2PropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnAutomationRuleV2PropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAutomationRuleV2PropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_securityhub.mixins.CfnAutomationRuleV2PropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAutomationRuleV2PropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_securityhub.mixins.CfnAutomationRuleV2PropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAutomationRuleV2PropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnAutomationRuleV2PropsMixin) Supports(construct constructs.IConstruct) *bool {
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

