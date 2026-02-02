package previewawsrtbfabricmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsrtbfabric/previewawsrtbfabricmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a new link between gateways.
//
// Establishes a connection that allows gateways to communicate and exchange bid requests and responses.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLinkPropsMixin := awscdkmixinspreview.Mixins.NewCfnLinkPropsMixin(&CfnLinkMixinProps{
//   	GatewayId: jsii.String("gatewayId"),
//   	HttpResponderAllowed: jsii.Boolean(false),
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
//   	ModuleConfigurationList: []interface{}{
//   		&ModuleConfigurationProperty{
//   			DependsOn: []*string{
//   				jsii.String("dependsOn"),
//   			},
//   			ModuleParameters: &ModuleParametersProperty{
//   				NoBid: &NoBidModuleParametersProperty{
//   					PassThroughPercentage: jsii.Number(123),
//   					Reason: jsii.String("reason"),
//   					ReasonCode: jsii.Number(123),
//   				},
//   				OpenRtbAttribute: &OpenRtbAttributeModuleParametersProperty{
//   					Action: &ActionProperty{
//   						HeaderTag: &HeaderTagActionProperty{
//   							Name: jsii.String("name"),
//   							Value: jsii.String("value"),
//   						},
//   						NoBid: &NoBidActionProperty{
//   							NoBidReasonCode: jsii.Number(123),
//   						},
//   					},
//   					FilterConfiguration: []interface{}{
//   						&FilterProperty{
//   							Criteria: []interface{}{
//   								&FilterCriterionProperty{
//   									Path: jsii.String("path"),
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   							},
//   						},
//   					},
//   					FilterType: jsii.String("filterType"),
//   					HoldbackPercentage: jsii.Number(123),
//   				},
//   			},
//   			Name: jsii.String("name"),
//   			Version: jsii.String("version"),
//   		},
//   	},
//   	PeerGatewayId: jsii.String("peerGatewayId"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rtbfabric-link.html
//
type CfnLinkPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnLinkMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnLinkPropsMixin
type jsiiProxy_CfnLinkPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnLinkPropsMixin) Props() *CfnLinkMixinProps {
	var returns *CfnLinkMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLinkPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::RTBFabric::Link`.
func NewCfnLinkPropsMixin(props *CfnLinkMixinProps, options *mixins.CfnPropertyMixinOptions) CfnLinkPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnLinkPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnLinkPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_rtbfabric.mixins.CfnLinkPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::RTBFabric::Link`.
func NewCfnLinkPropsMixin_Override(c CfnLinkPropsMixin, props *CfnLinkMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_rtbfabric.mixins.CfnLinkPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnLinkPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnLinkPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_rtbfabric.mixins.CfnLinkPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLinkPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_rtbfabric.mixins.CfnLinkPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnLinkPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnLinkPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

