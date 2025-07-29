package awsconnect


// Configuration information of a Kinesis video stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kinesisVideoStreamConfigProperty := &KinesisVideoStreamConfigProperty{
//   	EncryptionConfig: &EncryptionConfigProperty{
//   		EncryptionType: jsii.String("encryptionType"),
//   		KeyId: jsii.String("keyId"),
//   	},
//   	Prefix: jsii.String("prefix"),
//   	RetentionPeriodHours: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-instancestorageconfig-kinesisvideostreamconfig.html
//
type CfnInstanceStorageConfig_KinesisVideoStreamConfigProperty struct {
	// The encryption configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-instancestorageconfig-kinesisvideostreamconfig.html#cfn-connect-instancestorageconfig-kinesisvideostreamconfig-encryptionconfig
	//
	EncryptionConfig interface{} `field:"required" json:"encryptionConfig" yaml:"encryptionConfig"`
	// The prefix of the video stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-instancestorageconfig-kinesisvideostreamconfig.html#cfn-connect-instancestorageconfig-kinesisvideostreamconfig-prefix
	//
	Prefix *string `field:"required" json:"prefix" yaml:"prefix"`
	// The number of hours data is retained in the stream.
	//
	// Kinesis Video Streams retains the data in a data store that is associated with the stream.
	//
	// The default value is 0, indicating that the stream does not persist data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-instancestorageconfig-kinesisvideostreamconfig.html#cfn-connect-instancestorageconfig-kinesisvideostreamconfig-retentionperiodhours
	//
	RetentionPeriodHours *float64 `field:"required" json:"retentionPeriodHours" yaml:"retentionPeriodHours"`
}

