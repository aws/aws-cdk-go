package previewawsrtbfabricmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsrtbfabric/previewawsrtbfabricmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::RTBFabric::OutboundExternalLink Resource Type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnOutboundExternalLinkPropsMixin := awscdkmixinspreview.Mixins.NewCfnOutboundExternalLinkPropsMixin(&CfnOutboundExternalLinkMixinProps{
//   	GatewayId: jsii.String("gatewayId"),
//   	LinkAttributes: &LinkAttributesProperty{
//   		CustomerProvidedId: jsii.String("customerProvidedId"),
//   		ResponderErrorMasking: []interface{}{
//   			&ResponderErrorMaskingForHttpCodeProperty{
//   				Action: jsii.String("action"),
//   				HttpCode: jsii.String("httpCode"),
//   				LoggingTypes: []*string{
//   					jsii.String("loggingTypes"),
//   				},
//   				ResponseLoggingPercentage: jsii.Number(123),
//   			},
//   		},
//   	},
//   	LinkLogSettings: &LinkLogSettingsProperty{
//   		ApplicationLogs: &ApplicationLogsProperty{
//   			LinkApplicationLogSampling: &LinkApplicationLogSamplingProperty{
//   				ErrorLog: jsii.Number(123),
//   				FilterLog: jsii.Number(123),
//   			},
//   		},
//   	},
//   	PublicEndpoint: jsii.String("publicEndpoint"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rtbfabric-outboundexternallink.html
//
type CfnOutboundExternalLinkPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnOutboundExternalLinkMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnOutboundExternalLinkPropsMixin
type jsiiProxy_CfnOutboundExternalLinkPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnOutboundExternalLinkPropsMixin) Props() *CfnOutboundExternalLinkMixinProps {
	var returns *CfnOutboundExternalLinkMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOutboundExternalLinkPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::RTBFabric::OutboundExternalLink`.
func NewCfnOutboundExternalLinkPropsMixin(props *CfnOutboundExternalLinkMixinProps, options *mixins.CfnPropertyMixinOptions) CfnOutboundExternalLinkPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnOutboundExternalLinkPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnOutboundExternalLinkPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_rtbfabric.mixins.CfnOutboundExternalLinkPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::RTBFabric::OutboundExternalLink`.
func NewCfnOutboundExternalLinkPropsMixin_Override(c CfnOutboundExternalLinkPropsMixin, props *CfnOutboundExternalLinkMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_rtbfabric.mixins.CfnOutboundExternalLinkPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnOutboundExternalLinkPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnOutboundExternalLinkPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_rtbfabric.mixins.CfnOutboundExternalLinkPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnOutboundExternalLinkPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_rtbfabric.mixins.CfnOutboundExternalLinkPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnOutboundExternalLinkPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnOutboundExternalLinkPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

