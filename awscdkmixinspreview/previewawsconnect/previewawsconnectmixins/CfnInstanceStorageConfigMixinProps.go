package previewawsconnectmixins


// Properties for CfnInstanceStorageConfigPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnInstanceStorageConfigMixinProps := &CfnInstanceStorageConfigMixinProps{
//   	InstanceArn: jsii.String("instanceArn"),
//   	KinesisFirehoseConfig: &KinesisFirehoseConfigProperty{
//   		FirehoseArn: jsii.String("firehoseArn"),
//   	},
//   	KinesisStreamConfig: &KinesisStreamConfigProperty{
//   		StreamArn: jsii.String("streamArn"),
//   	},
//   	KinesisVideoStreamConfig: &KinesisVideoStreamConfigProperty{
//   		EncryptionConfig: &EncryptionConfigProperty{
//   			EncryptionType: jsii.String("encryptionType"),
//   			KeyId: jsii.String("keyId"),
//   		},
//   		Prefix: jsii.String("prefix"),
//   		RetentionPeriodHours: jsii.Number(123),
//   	},
//   	ResourceType: jsii.String("resourceType"),
//   	S3Config: &S3ConfigProperty{
//   		BucketName: jsii.String("bucketName"),
//   		BucketPrefix: jsii.String("bucketPrefix"),
//   		EncryptionConfig: &EncryptionConfigProperty{
//   			EncryptionType: jsii.String("encryptionType"),
//   			KeyId: jsii.String("keyId"),
//   		},
//   	},
//   	StorageType: jsii.String("storageType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-instancestorageconfig.html
//
type CfnInstanceStorageConfigMixinProps struct {
	// The Amazon Resource Name (ARN) of the instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-instancestorageconfig.html#cfn-connect-instancestorageconfig-instancearn
	//
	InstanceArn *string `field:"optional" json:"instanceArn" yaml:"instanceArn"`
	// The configuration of the Kinesis Firehose delivery stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-instancestorageconfig.html#cfn-connect-instancestorageconfig-kinesisfirehoseconfig
	//
	KinesisFirehoseConfig interface{} `field:"optional" json:"kinesisFirehoseConfig" yaml:"kinesisFirehoseConfig"`
	// The configuration of the Kinesis data stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-instancestorageconfig.html#cfn-connect-instancestorageconfig-kinesisstreamconfig
	//
	KinesisStreamConfig interface{} `field:"optional" json:"kinesisStreamConfig" yaml:"kinesisStreamConfig"`
	// The configuration of the Kinesis video stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-instancestorageconfig.html#cfn-connect-instancestorageconfig-kinesisvideostreamconfig
	//
	KinesisVideoStreamConfig interface{} `field:"optional" json:"kinesisVideoStreamConfig" yaml:"kinesisVideoStreamConfig"`
	// A valid resource type.
	//
	// Following are the valid resource types: `CHAT_TRANSCRIPTS` | `CALL_RECORDINGS` | `SCHEDULED_REPORTS` | `MEDIA_STREAMS` | `CONTACT_TRACE_RECORDS` | `AGENT_EVENTS`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-instancestorageconfig.html#cfn-connect-instancestorageconfig-resourcetype
	//
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
	// The S3 bucket configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-instancestorageconfig.html#cfn-connect-instancestorageconfig-s3config
	//
	S3Config interface{} `field:"optional" json:"s3Config" yaml:"s3Config"`
	// A valid storage type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-instancestorageconfig.html#cfn-connect-instancestorageconfig-storagetype
	//
	StorageType *string `field:"optional" json:"storageType" yaml:"storageType"`
}

