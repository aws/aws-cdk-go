package previewawskinesismixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawskinesis/previewawskinesismixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a Kinesis stream that captures and transports data records that are emitted from data sources.
//
// For information about creating streams, see [CreateStream](https://docs.aws.amazon.com/kinesis/latest/APIReference/API_CreateStream.html) in the Amazon Kinesis API Reference.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnStreamPropsMixin := awscdkmixinspreview.Mixins.NewCfnStreamPropsMixin(&CfnStreamMixinProps{
//   	DesiredShardLevelMetrics: []*string{
//   		jsii.String("desiredShardLevelMetrics"),
//   	},
//   	MaxRecordSizeInKiB: jsii.Number(123),
//   	Name: jsii.String("name"),
//   	RetentionPeriodHours: jsii.Number(123),
//   	ShardCount: jsii.Number(123),
//   	StreamEncryption: &StreamEncryptionProperty{
//   		EncryptionType: jsii.String("encryptionType"),
//   		KeyId: jsii.String("keyId"),
//   	},
//   	StreamModeDetails: &StreamModeDetailsProperty{
//   		StreamMode: jsii.String("streamMode"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	WarmThroughputMiBps: jsii.Number(123),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesis-stream.html
//
type CfnStreamPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnStreamMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnStreamPropsMixin
type jsiiProxy_CfnStreamPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnStreamPropsMixin) Props() *CfnStreamMixinProps {
	var returns *CfnStreamMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStreamPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Kinesis::Stream`.
func NewCfnStreamPropsMixin(props *CfnStreamMixinProps, options *mixins.CfnPropertyMixinOptions) CfnStreamPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnStreamPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnStreamPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_kinesis.mixins.CfnStreamPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Kinesis::Stream`.
func NewCfnStreamPropsMixin_Override(c CfnStreamPropsMixin, props *CfnStreamMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_kinesis.mixins.CfnStreamPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnStreamPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnStreamPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_kinesis.mixins.CfnStreamPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnStreamPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_kinesis.mixins.CfnStreamPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnStreamPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnStreamPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

