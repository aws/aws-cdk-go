package previewawsobservabilityadminmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsobservabilityadmin/previewawsobservabilityadminmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::ObservabilityAdmin::TelemetryPipelines.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTelemetryPipelinesPropsMixin := awscdkmixinspreview.Mixins.NewCfnTelemetryPipelinesPropsMixin(&CfnTelemetryPipelinesMixinProps{
//   	Configuration: &TelemetryPipelineConfigurationProperty{
//   		Body: jsii.String("body"),
//   	},
//   	Name: jsii.String("name"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-observabilityadmin-telemetrypipelines.html
//
type CfnTelemetryPipelinesPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnTelemetryPipelinesMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnTelemetryPipelinesPropsMixin
type jsiiProxy_CfnTelemetryPipelinesPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnTelemetryPipelinesPropsMixin) Props() *CfnTelemetryPipelinesMixinProps {
	var returns *CfnTelemetryPipelinesMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTelemetryPipelinesPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ObservabilityAdmin::TelemetryPipelines`.
func NewCfnTelemetryPipelinesPropsMixin(props *CfnTelemetryPipelinesMixinProps, options *mixins.CfnPropertyMixinOptions) CfnTelemetryPipelinesPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnTelemetryPipelinesPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTelemetryPipelinesPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_observabilityadmin.mixins.CfnTelemetryPipelinesPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ObservabilityAdmin::TelemetryPipelines`.
func NewCfnTelemetryPipelinesPropsMixin_Override(c CfnTelemetryPipelinesPropsMixin, props *CfnTelemetryPipelinesMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_observabilityadmin.mixins.CfnTelemetryPipelinesPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnTelemetryPipelinesPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTelemetryPipelinesPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_observabilityadmin.mixins.CfnTelemetryPipelinesPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTelemetryPipelinesPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_observabilityadmin.mixins.CfnTelemetryPipelinesPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTelemetryPipelinesPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnTelemetryPipelinesPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

