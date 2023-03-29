package awsconnect


// Properties for defining a `CfnInstanceStorageConfig`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnInstanceStorageConfigProps := &CfnInstanceStorageConfigProps{
//   	InstanceArn: jsii.String("instanceArn"),
//   	ResourceType: jsii.String("resourceType"),
//   	StorageType: jsii.String("storageType"),
//
//   	// the properties below are optional
//   	KinesisFirehoseConfig: &KinesisFirehoseConfigProperty{
//   		FirehoseArn: jsii.String("firehoseArn"),
//   	},
//   	KinesisStreamConfig: &KinesisStreamConfigProperty{
//   		StreamArn: jsii.String("streamArn"),
//   	},
//   	KinesisVideoStreamConfig: &KinesisVideoStreamConfigProperty{
//   		Prefix: jsii.String("prefix"),
//   		RetentionPeriodHours: jsii.Number(123),
//
//   		// the properties below are optional
//   		EncryptionConfig: &EncryptionConfigProperty{
//   			EncryptionType: jsii.String("encryptionType"),
//   			KeyId: jsii.String("keyId"),
//   		},
//   	},
//   	S3Config: &S3ConfigProperty{
//   		BucketName: jsii.String("bucketName"),
//   		BucketPrefix: jsii.String("bucketPrefix"),
//
//   		// the properties below are optional
//   		EncryptionConfig: &EncryptionConfigProperty{
//   			EncryptionType: jsii.String("encryptionType"),
//   			KeyId: jsii.String("keyId"),
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

