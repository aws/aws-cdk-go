package awsconnect


// Properties for defining a `CfnInstanceStorageConfig`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnInstanceStorageConfigProps := &cfnInstanceStorageConfigProps{
//   	instanceArn: jsii.String("instanceArn"),
//   	resourceType: jsii.String("resourceType"),
//   	storageType: jsii.String("storageType"),
//
//   	// the properties below are optional
//   	kinesisFirehoseConfig: &kinesisFirehoseConfigProperty{
//   		firehoseArn: jsii.String("firehoseArn"),
//   	},
//   	kinesisStreamConfig: &kinesisStreamConfigProperty{
//   		streamArn: jsii.String("streamArn"),
//   	},
//   	kinesisVideoStreamConfig: &kinesisVideoStreamConfigProperty{
//   		prefix: jsii.String("prefix"),
//   		retentionPeriodHours: jsii.Number(123),
//
//   		// the properties below are optional
//   		encryptionConfig: &encryptionConfigProperty{
//   			encryptionType: jsii.String("encryptionType"),
//   			keyId: jsii.String("keyId"),
//   		},
//   	},
//   	s3Config: &s3ConfigProperty{
//   		bucketName: jsii.String("bucketName"),
//   		bucketPrefix: jsii.String("bucketPrefix"),
//
//   		// the properties below are optional
//   		encryptionConfig: &encryptionConfigProperty{
//   			encryptionType: jsii.String("encryptionType"),
//   			keyId: jsii.String("keyId"),
//   		},
//   	},
//   }
//
type CfnInstanceStorageConfigProps struct {
	// The Amazon Resource Name (ARN) of the instance.
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// A valid resource type.
	//
	// Following are the valid resource types: `CHAT_TRANSCRIPTS` | `CALL_RECORDINGS` | `SCHEDULED_REPORTS` | `MEDIA_STREAMS` | `CONTACT_TRACE_RECORDS` | `AGENT_EVENTS`.
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
	// A valid storage type.
	StorageType *string `field:"required" json:"storageType" yaml:"storageType"`
	// The configuration of the Kinesis Firehose delivery stream.
	KinesisFirehoseConfig interface{} `field:"optional" json:"kinesisFirehoseConfig" yaml:"kinesisFirehoseConfig"`
	// The configuration of the Kinesis data stream.
	KinesisStreamConfig interface{} `field:"optional" json:"kinesisStreamConfig" yaml:"kinesisStreamConfig"`
	// The configuration of the Kinesis video stream.
	KinesisVideoStreamConfig interface{} `field:"optional" json:"kinesisVideoStreamConfig" yaml:"kinesisVideoStreamConfig"`
	// The S3 bucket configuration.
	S3Config interface{} `field:"optional" json:"s3Config" yaml:"s3Config"`
}

