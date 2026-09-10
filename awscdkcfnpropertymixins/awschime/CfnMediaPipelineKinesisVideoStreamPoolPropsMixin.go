package awschime

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awschime/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for an Amazon Chime SDK Media Pipeline Kinesis Video Stream Pool.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnMediaPipelineKinesisVideoStreamPoolPropsMixin := awscdkcfnpropertymixins.Aws_chime.NewCfnMediaPipelineKinesisVideoStreamPoolPropsMixin(&CfnMediaPipelineKinesisVideoStreamPoolMixinProps{
//   	PoolName: jsii.String("poolName"),
//   	StreamConfiguration: &StreamConfigurationProperty{
//   		DataRetentionInHours: jsii.Number(123),
//   		Region: jsii.String("region"),
//   	},
//   	Tags: []TagsItemsProperty{
//   		&TagsItemsProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-mediapipelinekinesisvideostreampool.html
//
type CfnMediaPipelineKinesisVideoStreamPoolPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnMediaPipelineKinesisVideoStreamPoolMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnMediaPipelineKinesisVideoStreamPoolPropsMixin
type jsiiProxy_CfnMediaPipelineKinesisVideoStreamPoolPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnMediaPipelineKinesisVideoStreamPoolPropsMixin) Props() *CfnMediaPipelineKinesisVideoStreamPoolMixinProps {
	var returns *CfnMediaPipelineKinesisVideoStreamPoolMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMediaPipelineKinesisVideoStreamPoolPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Chime::MediaPipelineKinesisVideoStreamPool`.
func NewCfnMediaPipelineKinesisVideoStreamPoolPropsMixin(props *CfnMediaPipelineKinesisVideoStreamPoolMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnMediaPipelineKinesisVideoStreamPoolPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnMediaPipelineKinesisVideoStreamPoolPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnMediaPipelineKinesisVideoStreamPoolPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_chime.CfnMediaPipelineKinesisVideoStreamPoolPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Chime::MediaPipelineKinesisVideoStreamPool`.
func NewCfnMediaPipelineKinesisVideoStreamPoolPropsMixin_Override(c CfnMediaPipelineKinesisVideoStreamPoolPropsMixin, props *CfnMediaPipelineKinesisVideoStreamPoolMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_chime.CfnMediaPipelineKinesisVideoStreamPoolPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnMediaPipelineKinesisVideoStreamPoolPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnMediaPipelineKinesisVideoStreamPoolPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_chime.CfnMediaPipelineKinesisVideoStreamPoolPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnMediaPipelineKinesisVideoStreamPoolPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_chime.CfnMediaPipelineKinesisVideoStreamPoolPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnMediaPipelineKinesisVideoStreamPoolPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnMediaPipelineKinesisVideoStreamPoolPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

