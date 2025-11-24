package mixinsawsiotwireless

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsiotwireless/mixinsawsiotwireless/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a new service profile.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnServiceProfilePropsMixin := awscdkmixinspreview.Mixins.NewCfnServiceProfilePropsMixin(&CfnServiceProfileMixinProps{
//   	LoRaWan: &LoRaWANServiceProfileProperty{
//   		AddGwMetadata: jsii.Boolean(false),
//   		ChannelMask: jsii.String("channelMask"),
//   		DevStatusReqFreq: jsii.Number(123),
//   		DlBucketSize: jsii.Number(123),
//   		DlRate: jsii.Number(123),
//   		DlRatePolicy: jsii.String("dlRatePolicy"),
//   		DrMax: jsii.Number(123),
//   		DrMin: jsii.Number(123),
//   		HrAllowed: jsii.Boolean(false),
//   		MinGwDiversity: jsii.Number(123),
//   		NwkGeoLoc: jsii.Boolean(false),
//   		PrAllowed: jsii.Boolean(false),
//   		RaAllowed: jsii.Boolean(false),
//   		ReportDevStatusBattery: jsii.Boolean(false),
//   		ReportDevStatusMargin: jsii.Boolean(false),
//   		TargetPer: jsii.Number(123),
//   		UlBucketSize: jsii.Number(123),
//   		UlRate: jsii.Number(123),
//   		UlRatePolicy: jsii.String("ulRatePolicy"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-serviceprofile.html
//
type CfnServiceProfilePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnServiceProfileMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnServiceProfilePropsMixin
type jsiiProxy_CfnServiceProfilePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnServiceProfilePropsMixin) Props() *CfnServiceProfileMixinProps {
	var returns *CfnServiceProfileMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceProfilePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::IoTWireless::ServiceProfile`.
func NewCfnServiceProfilePropsMixin(props *CfnServiceProfileMixinProps, options *mixins.CfnPropertyMixinOptions) CfnServiceProfilePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnServiceProfilePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnServiceProfilePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iotwireless.mixins.CfnServiceProfilePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::IoTWireless::ServiceProfile`.
func NewCfnServiceProfilePropsMixin_Override(c CfnServiceProfilePropsMixin, props *CfnServiceProfileMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iotwireless.mixins.CfnServiceProfilePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnServiceProfilePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnServiceProfilePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_iotwireless.mixins.CfnServiceProfilePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnServiceProfilePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_iotwireless.mixins.CfnServiceProfilePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnServiceProfilePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnServiceProfilePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

