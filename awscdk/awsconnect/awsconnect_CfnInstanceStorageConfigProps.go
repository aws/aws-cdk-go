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
	// `AWS::Connect::InstanceStorageConfig.InstanceArn`.
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// `AWS::Connect::InstanceStorageConfig.ResourceType`.
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
	// `AWS::Connect::InstanceStorageConfig.StorageType`.
	StorageType *string `field:"required" json:"storageType" yaml:"storageType"`
	// `AWS::Connect::InstanceStorageConfig.KinesisFirehoseConfig`.
	KinesisFirehoseConfig interface{} `field:"optional" json:"kinesisFirehoseConfig" yaml:"kinesisFirehoseConfig"`
	// `AWS::Connect::InstanceStorageConfig.KinesisStreamConfig`.
	KinesisStreamConfig interface{} `field:"optional" json:"kinesisStreamConfig" yaml:"kinesisStreamConfig"`
	// `AWS::Connect::InstanceStorageConfig.KinesisVideoStreamConfig`.
	KinesisVideoStreamConfig interface{} `field:"optional" json:"kinesisVideoStreamConfig" yaml:"kinesisVideoStreamConfig"`
	// `AWS::Connect::InstanceStorageConfig.S3Config`.
	S3Config interface{} `field:"optional" json:"s3Config" yaml:"s3Config"`
}

