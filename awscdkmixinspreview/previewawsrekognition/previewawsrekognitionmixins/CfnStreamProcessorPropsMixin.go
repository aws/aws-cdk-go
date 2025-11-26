package previewawsrekognitionmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsrekognition/previewawsrekognitionmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::Rekognition::StreamProcessor` type creates a stream processor used to detect and recognize faces or to detect connected home labels in a streaming video.
//
// Amazon Rekognition Video is a consumer of live video from Amazon Kinesis Video Streams. There are two different settings for stream processors in Amazon Rekognition, one for detecting faces and one for connected home features.
//
// If you are creating a stream processor for detecting faces, you provide a Kinesis video stream (input) and a Kinesis data stream (output). You also specify the face recognition criteria in FaceSearchSettings. For example, the collection containing faces that you want to recognize.
//
// If you are creating a stream processor for detection of connected home labels, you provide a Kinesis video stream for input, and for output an Amazon S3 bucket and an Amazon SNS topic. You can also provide a KMS key ID to encrypt the data sent to your Amazon S3 bucket. You specify what you want to detect in ConnectedHomeSettings, such as people, packages, and pets.
//
// You can also specify where in the frame you want Amazon Rekognition to monitor with BoundingBoxRegionsOfInterest and PolygonRegionsOfInterest. The Name is used to manage the stream processor and it is the identifier for the stream processor. The `AWS::Rekognition::StreamProcessor` resource creates a stream processor in the same Region where you create the Amazon CloudFormation stack.
//
// For more information, see [CreateStreamProcessor](https://docs.aws.amazon.com/rekognition/latest/APIReference/API_CreateStreamProcessor) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var polygonRegionsOfInterest interface{}
//
//   cfnStreamProcessorPropsMixin := awscdkmixinspreview.Mixins.NewCfnStreamProcessorPropsMixin(&CfnStreamProcessorMixinProps{
//   	BoundingBoxRegionsOfInterest: []interface{}{
//   		&BoundingBoxProperty{
//   			Height: jsii.Number(123),
//   			Left: jsii.Number(123),
//   			Top: jsii.Number(123),
//   			Width: jsii.Number(123),
//   		},
//   	},
//   	ConnectedHomeSettings: &ConnectedHomeSettingsProperty{
//   		Labels: []*string{
//   			jsii.String("labels"),
//   		},
//   		MinConfidence: jsii.Number(123),
//   	},
//   	DataSharingPreference: &DataSharingPreferenceProperty{
//   		OptIn: jsii.Boolean(false),
//   	},
//   	FaceSearchSettings: &FaceSearchSettingsProperty{
//   		CollectionId: jsii.String("collectionId"),
//   		FaceMatchThreshold: jsii.Number(123),
//   	},
//   	KinesisDataStream: &KinesisDataStreamProperty{
//   		Arn: jsii.String("arn"),
//   	},
//   	KinesisVideoStream: &KinesisVideoStreamProperty{
//   		Arn: jsii.String("arn"),
//   	},
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	Name: jsii.String("name"),
//   	NotificationChannel: &NotificationChannelProperty{
//   		Arn: jsii.String("arn"),
//   	},
//   	PolygonRegionsOfInterest: polygonRegionsOfInterest,
//   	RoleArn: jsii.String("roleArn"),
//   	S3Destination: &S3DestinationProperty{
//   		BucketName: jsii.String("bucketName"),
//   		ObjectKeyPrefix: jsii.String("objectKeyPrefix"),
//   	},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rekognition-streamprocessor.html
//
type CfnStreamProcessorPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnStreamProcessorMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnStreamProcessorPropsMixin
type jsiiProxy_CfnStreamProcessorPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnStreamProcessorPropsMixin) Props() *CfnStreamProcessorMixinProps {
	var returns *CfnStreamProcessorMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStreamProcessorPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Rekognition::StreamProcessor`.
func NewCfnStreamProcessorPropsMixin(props *CfnStreamProcessorMixinProps, options *mixins.CfnPropertyMixinOptions) CfnStreamProcessorPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnStreamProcessorPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnStreamProcessorPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_rekognition.mixins.CfnStreamProcessorPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Rekognition::StreamProcessor`.
func NewCfnStreamProcessorPropsMixin_Override(c CfnStreamProcessorPropsMixin, props *CfnStreamProcessorMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_rekognition.mixins.CfnStreamProcessorPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnStreamProcessorPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnStreamProcessorPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_rekognition.mixins.CfnStreamProcessorPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnStreamProcessorPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_rekognition.mixins.CfnStreamProcessorPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnStreamProcessorPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnStreamProcessorPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

