package previewawsgameliftstreamsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsgameliftstreams/previewawsgameliftstreamsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::GameLiftStreams::StreamGroup` resource defines a group of compute resources that will be running and streaming your game.
//
// When you create a stream group, you specify the hardware configuration (CPU, GPU, RAM) that will run your game (known as the *stream class* ), the geographical locations where your game can run, and the number of streams that can run simultaneously in each location (known as *stream capacity* ). Stream groups manage how Amazon GameLift Streams allocates resources and handles concurrent streams, allowing you to effectively manage capacity and costs.
//
// There are two types of stream capacity: always-on and on-demand.
//
// - *Always-on* : The streaming capacity that is allocated and ready to handle stream requests without delay. You pay for this capacity whether it's in use or not. Best for quickest time from streaming request to streaming session. Default is 1 (2 for high stream classes) when creating a stream group or adding a location.
// - *On-demand* : The streaming capacity that Amazon GameLift Streams can allocate in response to stream requests, and then de-allocate when the session has terminated. This offers a cost control measure at the expense of a greater startup time (typically under 5 minutes). Default is 0 when creating a stream group or adding a location.
//
// Values for capacity must be whole number multiples of the tenancy value of the stream group's stream class.
//
// > Application association is not currently supported in CloudFormation . To link additional applications to a stream group, use the Amazon GameLift Streams console or the AWS CLI .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnStreamGroupPropsMixin := awscdkmixinspreview.Mixins.NewCfnStreamGroupPropsMixin(&CfnStreamGroupMixinProps{
//   	DefaultApplication: &DefaultApplicationProperty{
//   		Arn: jsii.String("arn"),
//   		Id: jsii.String("id"),
//   	},
//   	Description: jsii.String("description"),
//   	LocationConfigurations: []interface{}{
//   		&LocationConfigurationProperty{
//   			AlwaysOnCapacity: jsii.Number(123),
//   			LocationName: jsii.String("locationName"),
//   			MaximumCapacity: jsii.Number(123),
//   			OnDemandCapacity: jsii.Number(123),
//   			TargetIdleCapacity: jsii.Number(123),
//   		},
//   	},
//   	StreamClass: jsii.String("streamClass"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gameliftstreams-streamgroup.html
//
type CfnStreamGroupPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnStreamGroupMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnStreamGroupPropsMixin
type jsiiProxy_CfnStreamGroupPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnStreamGroupPropsMixin) Props() *CfnStreamGroupMixinProps {
	var returns *CfnStreamGroupMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStreamGroupPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::GameLiftStreams::StreamGroup`.
func NewCfnStreamGroupPropsMixin(props *CfnStreamGroupMixinProps, options *mixins.CfnPropertyMixinOptions) CfnStreamGroupPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnStreamGroupPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnStreamGroupPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_gameliftstreams.mixins.CfnStreamGroupPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::GameLiftStreams::StreamGroup`.
func NewCfnStreamGroupPropsMixin_Override(c CfnStreamGroupPropsMixin, props *CfnStreamGroupMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_gameliftstreams.mixins.CfnStreamGroupPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnStreamGroupPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnStreamGroupPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_gameliftstreams.mixins.CfnStreamGroupPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnStreamGroupPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_gameliftstreams.mixins.CfnStreamGroupPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnStreamGroupPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnStreamGroupPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

