package previewawsobservabilityadminmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsobservabilityadmin/previewawsobservabilityadminmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a telemetry rule that defines how telemetry should be configured for AWS resources in your account.
//
// The rule specifies which resources should have telemetry enabled and how that telemetry data should be collected based on resource type, telemetry type, and selection criteria.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTelemetryRulePropsMixin := awscdkmixinspreview.Mixins.NewCfnTelemetryRulePropsMixin(&CfnTelemetryRuleMixinProps{
//   	Rule: &TelemetryRuleProperty{
//   		DestinationConfiguration: &TelemetryDestinationConfigurationProperty{
//   			CloudtrailParameters: &CloudtrailParametersProperty{
//   				AdvancedEventSelectors: []interface{}{
//   					&AdvancedEventSelectorProperty{
//   						FieldSelectors: []interface{}{
//   							&AdvancedFieldSelectorProperty{
//   								EndsWith: []*string{
//   									jsii.String("endsWith"),
//   								},
//   								EqualTo: []*string{
//   									jsii.String("equalTo"),
//   								},
//   								Field: jsii.String("field"),
//   								NotEndsWith: []*string{
//   									jsii.String("notEndsWith"),
//   								},
//   								NotEquals: []*string{
//   									jsii.String("notEquals"),
//   								},
//   								NotStartsWith: []*string{
//   									jsii.String("notStartsWith"),
//   								},
//   								StartsWith: []*string{
//   									jsii.String("startsWith"),
//   								},
//   							},
//   						},
//   						Name: jsii.String("name"),
//   					},
//   				},
//   			},
//   			DestinationPattern: jsii.String("destinationPattern"),
//   			DestinationType: jsii.String("destinationType"),
//   			ElbLoadBalancerLoggingParameters: &ELBLoadBalancerLoggingParametersProperty{
//   				FieldDelimiter: jsii.String("fieldDelimiter"),
//   				OutputFormat: jsii.String("outputFormat"),
//   			},
//   			LogDeliveryParameters: &LogDeliveryParametersProperty{
//   				LogTypes: []*string{
//   					jsii.String("logTypes"),
//   				},
//   			},
//   			RetentionInDays: jsii.Number(123),
//   			VpcFlowLogParameters: &VPCFlowLogParametersProperty{
//   				LogFormat: jsii.String("logFormat"),
//   				MaxAggregationInterval: jsii.Number(123),
//   				TrafficType: jsii.String("trafficType"),
//   			},
//   			WafLoggingParameters: &WAFLoggingParametersProperty{
//   				LoggingFilter: &LoggingFilterProperty{
//   					DefaultBehavior: jsii.String("defaultBehavior"),
//   					Filters: []interface{}{
//   						&FilterProperty{
//   							Behavior: jsii.String("behavior"),
//   							Conditions: []interface{}{
//   								&ConditionProperty{
//   									ActionCondition: &ActionConditionProperty{
//   										Action: jsii.String("action"),
//   									},
//   									LabelNameCondition: &LabelNameConditionProperty{
//   										LabelName: jsii.String("labelName"),
//   									},
//   								},
//   							},
//   							Requirement: jsii.String("requirement"),
//   						},
//   					},
//   				},
//   				LogType: jsii.String("logType"),
//   				RedactedFields: []interface{}{
//   					&FieldToMatchProperty{
//   						Method: jsii.String("method"),
//   						QueryString: jsii.String("queryString"),
//   						SingleHeader: &SingleHeaderProperty{
//   							Name: jsii.String("name"),
//   						},
//   						UriPath: jsii.String("uriPath"),
//   					},
//   				},
//   			},
//   		},
//   		ResourceType: jsii.String("resourceType"),
//   		SelectionCriteria: jsii.String("selectionCriteria"),
//   		TelemetrySourceTypes: []*string{
//   			jsii.String("telemetrySourceTypes"),
//   		},
//   		TelemetryType: jsii.String("telemetryType"),
//   	},
//   	RuleName: jsii.String("ruleName"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-observabilityadmin-telemetryrule.html
//
type CfnTelemetryRulePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnTelemetryRuleMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnTelemetryRulePropsMixin
type jsiiProxy_CfnTelemetryRulePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnTelemetryRulePropsMixin) Props() *CfnTelemetryRuleMixinProps {
	var returns *CfnTelemetryRuleMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTelemetryRulePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ObservabilityAdmin::TelemetryRule`.
func NewCfnTelemetryRulePropsMixin(props *CfnTelemetryRuleMixinProps, options *mixins.CfnPropertyMixinOptions) CfnTelemetryRulePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnTelemetryRulePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTelemetryRulePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_observabilityadmin.mixins.CfnTelemetryRulePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ObservabilityAdmin::TelemetryRule`.
func NewCfnTelemetryRulePropsMixin_Override(c CfnTelemetryRulePropsMixin, props *CfnTelemetryRuleMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_observabilityadmin.mixins.CfnTelemetryRulePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnTelemetryRulePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTelemetryRulePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_observabilityadmin.mixins.CfnTelemetryRulePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTelemetryRulePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_observabilityadmin.mixins.CfnTelemetryRulePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTelemetryRulePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnTelemetryRulePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

