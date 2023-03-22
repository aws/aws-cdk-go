package awsrekognition

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnStreamProcessor`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var polygonRegionsOfInterest interface{}
//
//   cfnStreamProcessorProps := &cfnStreamProcessorProps{
//   	kinesisVideoStream: &kinesisVideoStreamProperty{
//   		arn: jsii.String("arn"),
//   	},
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	boundingBoxRegionsOfInterest: []interface{}{
//   		&boundingBoxProperty{
//   			height: jsii.Number(123),
//   			left: jsii.Number(123),
//   			top: jsii.Number(123),
//   			width: jsii.Number(123),
//   		},
//   	},
//   	connectedHomeSettings: &connectedHomeSettingsProperty{
//   		labels: []*string{
//   			jsii.String("labels"),
//   		},
//
//   		// the properties below are optional
//   		minConfidence: jsii.Number(123),
//   	},
//   	dataSharingPreference: &dataSharingPreferenceProperty{
//   		optIn: jsii.Boolean(false),
//   	},
//   	faceSearchSettings: &faceSearchSettingsProperty{
//   		collectionId: jsii.String("collectionId"),
//
//   		// the properties below are optional
//   		faceMatchThreshold: jsii.Number(123),
//   	},
//   	kinesisDataStream: &kinesisDataStreamProperty{
//   		arn: jsii.String("arn"),
//   	},
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	name: jsii.String("name"),
//   	notificationChannel: &notificationChannelProperty{
//   		arn: jsii.String("arn"),
//   	},
//   	polygonRegionsOfInterest: polygonRegionsOfInterest,
//   	s3Destination: &s3DestinationProperty{
//   		bucketName: jsii.String("bucketName"),
//
//   		// the properties below are optional
//   		objectKeyPrefix: jsii.String("objectKeyPrefix"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnStreamProcessorProps struct {
	// The Kinesis video stream that provides the source of the streaming video for an Amazon Rekognition Video stream processor.
	//
	// For more information, see [KinesisVideoStream](https://docs.aws.amazon.com/rekognition/latest/APIReference/API_KinesisVideoStream) .
	KinesisVideoStream interface{} `field:"required" json:"kinesisVideoStream" yaml:"kinesisVideoStream"`
	// The ARN of the IAM role that allows access to the stream processor.
	//
	// The IAM role provides Rekognition read permissions to the Kinesis stream. It also provides write permissions to an Amazon S3 bucket and Amazon Simple Notification Service topic for a connected home stream processor. This is required for both face search and connected home stream processors. For information about constraints, see the RoleArn section of [CreateStreamProcessor](https://docs.aws.amazon.com/rekognition/latest/APIReference/API_CreateStreamProcessor) .
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// List of BoundingBox objects, each of which denotes a region of interest on screen.
	//
	// For more information, see the BoundingBox field of [RegionOfInterest](https://docs.aws.amazon.com/rekognition/latest/APIReference/API_RegionOfInterest) .
	BoundingBoxRegionsOfInterest interface{} `field:"optional" json:"boundingBoxRegionsOfInterest" yaml:"boundingBoxRegionsOfInterest"`
	// Connected home settings to use on a streaming video.
	//
	// You can use a stream processor for connected home features and select what you want the stream processor to detect, such as people or pets. When the stream processor has started, one notification is sent for each object class specified. For more information, see the ConnectedHome section of [StreamProcessorSettings](https://docs.aws.amazon.com/rekognition/latest/APIReference/API_StreamProcessorSettings) .
	ConnectedHomeSettings interface{} `field:"optional" json:"connectedHomeSettings" yaml:"connectedHomeSettings"`
	// Allows you to opt in or opt out to share data with Rekognition to improve model performance.
	//
	// You can choose this option at the account level or on a per-stream basis. Note that if you opt out at the account level this setting is ignored on individual streams. For more information, see [StreamProcessorDataSharingPreference](https://docs.aws.amazon.com/rekognition/latest/APIReference/API_StreamProcessorDataSharingPreference) .
	DataSharingPreference interface{} `field:"optional" json:"dataSharingPreference" yaml:"dataSharingPreference"`
	// The input parameters used to recognize faces in a streaming video analyzed by an Amazon Rekognition stream processor.
	//
	// For more information regarding the contents of the parameters, see [FaceSearchSettings](https://docs.aws.amazon.com/rekognition/latest/APIReference/API_FaceSearchSettings) .
	FaceSearchSettings interface{} `field:"optional" json:"faceSearchSettings" yaml:"faceSearchSettings"`
	// Amazon Rekognition's Video Stream Processor takes a Kinesis video stream as input.
	//
	// This is the Amazon Kinesis Data Streams instance to which the Amazon Rekognition stream processor streams the analysis results. This must be created within the constraints specified at [KinesisDataStream](https://docs.aws.amazon.com/rekognition/latest/APIReference/API_KinesisDataStream) .
	KinesisDataStream interface{} `field:"optional" json:"kinesisDataStream" yaml:"kinesisDataStream"`
	// The identifier for your Amazon Key Management Service key (Amazon KMS key).
	//
	// Optional parameter for connected home stream processors used to encrypt results and data published to your Amazon S3 bucket. For more information, see the KMSKeyId section of [CreateStreamProcessor](https://docs.aws.amazon.com/rekognition/latest/APIReference/API_CreateStreamProcessor) .
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The Name attribute specifies the name of the stream processor and it must be within the constraints described in the Name section of [StreamProcessor](https://docs.aws.amazon.com/rekognition/latest/APIReference/API_StreamProcessor) . If you don't specify a name, Amazon CloudFormation generates a unique ID and uses that ID for the stream processor name.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The Amazon Simple Notification Service topic to which Amazon Rekognition publishes the object detection results and completion status of a video analysis operation.
	//
	// Amazon Rekognition publishes a notification the first time an object of interest or a person is detected in the video stream. Amazon Rekognition also publishes an end-of-session notification with a summary when the stream processing session is complete. For more information, see [StreamProcessorNotificationChannel](https://docs.aws.amazon.com/rekognition/latest/APIReference/API_StreamProcessorNotificationChannel) .
	NotificationChannel interface{} `field:"optional" json:"notificationChannel" yaml:"notificationChannel"`
	// A set of ordered lists of [Point](https://docs.aws.amazon.com/rekognition/latest/APIReference/API_Point) objects. Each entry of the set contains a polygon denoting a region of interest on the screen. Each polygon is an ordered list of [Point](https://docs.aws.amazon.com/rekognition/latest/APIReference/API_Point) objects. For more information, see the Polygon field of [RegionOfInterest](https://docs.aws.amazon.com/rekognition/latest/APIReference/API_RegionOfInterest) .
	PolygonRegionsOfInterest interface{} `field:"optional" json:"polygonRegionsOfInterest" yaml:"polygonRegionsOfInterest"`
	// The Amazon S3 bucket location to which Amazon Rekognition publishes the detailed inference results of a video analysis operation.
	//
	// For more information, see the S3Destination section of [StreamProcessorOutput](https://docs.aws.amazon.com/rekognition/latest/APIReference/API_StreamProcessorOutput) .
	S3Destination interface{} `field:"optional" json:"s3Destination" yaml:"s3Destination"`
	// A set of tags (key-value pairs) that you want to attach to the stream processor.
	//
	// For more information, see the Tags section of [CreateStreamProcessor](https://docs.aws.amazon.com/rekognition/latest/APIReference/API_CreateStreamProcessor) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

