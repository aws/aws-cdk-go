package previewawsiotsitewisemixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsiotsitewise/previewawsiotsitewisemixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a gateway, which is a virtual or edge device that delivers industrial data streams from local servers to AWS IoT SiteWise .
//
// For more information, see [Ingesting data using a gateway](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/gateway-connector.html) in the *AWS IoT SiteWise User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnGatewayPropsMixin := awscdkmixinspreview.Mixins.NewCfnGatewayPropsMixin(&CfnGatewayMixinProps{
//   	GatewayCapabilitySummaries: []interface{}{
//   		&GatewayCapabilitySummaryProperty{
//   			CapabilityConfiguration: jsii.String("capabilityConfiguration"),
//   			CapabilityNamespace: jsii.String("capabilityNamespace"),
//   		},
//   	},
//   	GatewayName: jsii.String("gatewayName"),
//   	GatewayPlatform: &GatewayPlatformProperty{
//   		Greengrass: &GreengrassProperty{
//   			GroupArn: jsii.String("groupArn"),
//   		},
//   		GreengrassV2: &GreengrassV2Property{
//   			CoreDeviceOperatingSystem: jsii.String("coreDeviceOperatingSystem"),
//   			CoreDeviceThingName: jsii.String("coreDeviceThingName"),
//   		},
//   		SiemensIe: &SiemensIEProperty{
//   			IotCoreThingName: jsii.String("iotCoreThingName"),
//   		},
//   	},
//   	GatewayVersion: jsii.String("gatewayVersion"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-gateway.html
//
type CfnGatewayPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnGatewayMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnGatewayPropsMixin
type jsiiProxy_CfnGatewayPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnGatewayPropsMixin) Props() *CfnGatewayMixinProps {
	var returns *CfnGatewayMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGatewayPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::IoTSiteWise::Gateway`.
func NewCfnGatewayPropsMixin(props *CfnGatewayMixinProps, options *mixins.CfnPropertyMixinOptions) CfnGatewayPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnGatewayPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnGatewayPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iotsitewise.mixins.CfnGatewayPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::IoTSiteWise::Gateway`.
func NewCfnGatewayPropsMixin_Override(c CfnGatewayPropsMixin, props *CfnGatewayMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iotsitewise.mixins.CfnGatewayPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnGatewayPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnGatewayPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_iotsitewise.mixins.CfnGatewayPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnGatewayPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_iotsitewise.mixins.CfnGatewayPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnGatewayPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnGatewayPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

