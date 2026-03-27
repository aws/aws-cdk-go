package awsobservabilityadmin

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsobservabilityadmin/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// AWS::ObservabilityAdmin::TelemetryEnrichment cloudformation resource enables the resource tags for telemetry feature in CloudWatch to enrich infrastructure metrics with AWS resource tags.
//
// For more details: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/resource-tags-for-telemetry.html
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnTelemetryEnrichmentPropsMixin := awscdkcfnpropertymixins.Aws_observabilityadmin.NewCfnTelemetryEnrichmentPropsMixin(&CfnTelemetryEnrichmentMixinProps{
//   	Scope: jsii.String("scope"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-observabilityadmin-telemetryenrichment.html
//
type CfnTelemetryEnrichmentPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnTelemetryEnrichmentMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnTelemetryEnrichmentPropsMixin
type jsiiProxy_CfnTelemetryEnrichmentPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnTelemetryEnrichmentPropsMixin) Props() *CfnTelemetryEnrichmentMixinProps {
	var returns *CfnTelemetryEnrichmentMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTelemetryEnrichmentPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ObservabilityAdmin::TelemetryEnrichment`.
func NewCfnTelemetryEnrichmentPropsMixin(props *CfnTelemetryEnrichmentMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnTelemetryEnrichmentPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnTelemetryEnrichmentPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTelemetryEnrichmentPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_observabilityadmin.CfnTelemetryEnrichmentPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ObservabilityAdmin::TelemetryEnrichment`.
func NewCfnTelemetryEnrichmentPropsMixin_Override(c CfnTelemetryEnrichmentPropsMixin, props *CfnTelemetryEnrichmentMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_observabilityadmin.CfnTelemetryEnrichmentPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnTelemetryEnrichmentPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTelemetryEnrichmentPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_observabilityadmin.CfnTelemetryEnrichmentPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTelemetryEnrichmentPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_observabilityadmin.CfnTelemetryEnrichmentPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTelemetryEnrichmentPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnTelemetryEnrichmentPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

