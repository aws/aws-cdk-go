package awsmediaconnect

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsmediaconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource schema for AWS::MediaConnect::FlowMediaStream.
//
// A media stream represents a single track or stream of media containing video, audio, or ancillary data that is transported using the SMPTE 2110 JPEG XS or CDI protocol.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnFlowMediaStreamPropsMixin := awscdkcfnpropertymixins.Aws_mediaconnect.NewCfnFlowMediaStreamPropsMixin(&CfnFlowMediaStreamMixinProps{
//   	Attributes: &MediaStreamAttributesProperty{
//   		Fmtp: &FmtpProperty{
//   			ChannelOrder: jsii.String("channelOrder"),
//   			Colorimetry: jsii.String("colorimetry"),
//   			ExactFramerate: jsii.String("exactFramerate"),
//   			Par: jsii.String("par"),
//   			Range: jsii.String("range"),
//   			ScanMode: jsii.String("scanMode"),
//   			Tcs: jsii.String("tcs"),
//   		},
//   		Lang: jsii.String("lang"),
//   	},
//   	ClockRate: jsii.Number(123),
//   	Description: jsii.String("description"),
//   	FlowArn: jsii.String("flowArn"),
//   	MediaStreamId: jsii.Number(123),
//   	MediaStreamName: jsii.String("mediaStreamName"),
//   	MediaStreamType: jsii.String("mediaStreamType"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VideoFormat: jsii.String("videoFormat"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowmediastream.html
//
type CfnFlowMediaStreamPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnFlowMediaStreamMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnFlowMediaStreamPropsMixin
type jsiiProxy_CfnFlowMediaStreamPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnFlowMediaStreamPropsMixin) Props() *CfnFlowMediaStreamMixinProps {
	var returns *CfnFlowMediaStreamMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowMediaStreamPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::MediaConnect::FlowMediaStream`.
func NewCfnFlowMediaStreamPropsMixin(props *CfnFlowMediaStreamMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnFlowMediaStreamPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnFlowMediaStreamPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnFlowMediaStreamPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowMediaStreamPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::MediaConnect::FlowMediaStream`.
func NewCfnFlowMediaStreamPropsMixin_Override(c CfnFlowMediaStreamPropsMixin, props *CfnFlowMediaStreamMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowMediaStreamPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnFlowMediaStreamPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnFlowMediaStreamPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowMediaStreamPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFlowMediaStreamPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowMediaStreamPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnFlowMediaStreamPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnFlowMediaStreamPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

