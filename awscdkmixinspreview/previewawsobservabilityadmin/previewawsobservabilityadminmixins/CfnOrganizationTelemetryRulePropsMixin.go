package previewawsobservabilityadminmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsobservabilityadmin/previewawsobservabilityadminmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Retrieves the details of a specific organization centralization rule.
//
// This operation can only be called by the organization's management account or a delegated administrator account.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnOrganizationTelemetryRulePropsMixin := awscdkmixinspreview.Mixins.NewCfnOrganizationTelemetryRulePropsMixin(&CfnOrganizationTelemetryRuleMixinProps{
//   	Rule: &TelemetryRuleProperty{
//   		DestinationConfiguration: &TelemetryDestinationConfigurationProperty{
//   			DestinationPattern: jsii.String("destinationPattern"),
//   			DestinationType: jsii.String("destinationType"),
//   			RetentionInDays: jsii.Number(123),
//   			VpcFlowLogParameters: &VPCFlowLogParametersProperty{
//   				LogFormat: jsii.String("logFormat"),
//   				MaxAggregationInterval: jsii.Number(123),
//   				TrafficType: jsii.String("trafficType"),
//   			},
//   		},
//   		ResourceType: jsii.String("resourceType"),
//   		Scope: jsii.String("scope"),
//   		SelectionCriteria: jsii.String("selectionCriteria"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-observabilityadmin-organizationtelemetryrule.html
//
type CfnOrganizationTelemetryRulePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnOrganizationTelemetryRuleMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnOrganizationTelemetryRulePropsMixin
type jsiiProxy_CfnOrganizationTelemetryRulePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnOrganizationTelemetryRulePropsMixin) Props() *CfnOrganizationTelemetryRuleMixinProps {
	var returns *CfnOrganizationTelemetryRuleMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOrganizationTelemetryRulePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ObservabilityAdmin::OrganizationTelemetryRule`.
func NewCfnOrganizationTelemetryRulePropsMixin(props *CfnOrganizationTelemetryRuleMixinProps, options *mixins.CfnPropertyMixinOptions) CfnOrganizationTelemetryRulePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnOrganizationTelemetryRulePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnOrganizationTelemetryRulePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_observabilityadmin.mixins.CfnOrganizationTelemetryRulePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ObservabilityAdmin::OrganizationTelemetryRule`.
func NewCfnOrganizationTelemetryRulePropsMixin_Override(c CfnOrganizationTelemetryRulePropsMixin, props *CfnOrganizationTelemetryRuleMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_observabilityadmin.mixins.CfnOrganizationTelemetryRulePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnOrganizationTelemetryRulePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnOrganizationTelemetryRulePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_observabilityadmin.mixins.CfnOrganizationTelemetryRulePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnOrganizationTelemetryRulePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_observabilityadmin.mixins.CfnOrganizationTelemetryRulePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnOrganizationTelemetryRulePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnOrganizationTelemetryRulePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

