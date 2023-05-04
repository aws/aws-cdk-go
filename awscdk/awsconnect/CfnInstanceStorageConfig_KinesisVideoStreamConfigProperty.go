package awsconnect


// Configuration information of a Kinesis video stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kinesisVideoStreamConfigProperty := &KinesisVideoStreamConfigProperty{
//   	Prefix: jsii.String("prefix"),
//   	RetentionPeriodHours: jsii.Number(123),
//
//   	// the properties below are optional
//   	EncryptionConfig: &EncryptionConfigProperty{
//   		EncryptionType: jsii.String("encryptionType"),
//   		KeyId: jsii.String("keyId"),
//   	},
//   }
//
type CfnInstanceStorageConfig_KinesisVideoStreamConfigProperty struct {
	// The prefix of the video stream.
	Prefix *string `field:"required" json:"prefix" yaml:"prefix"`
	// The number of hours data is retained in the stream.
	//
	// Kinesis Video Streams retains the data in a data store that is associated with the stream.
	//
	// The default value is 0, indicating that the stream does not persist data.
	RetentionPeriodHours *float64 `field:"required" json:"retentionPeriodHours" yaml:"retentionPeriodHours"`
	// The encryption configuration.
	EncryptionConfig interface{} `field:"optional" json:"encryptionConfig" yaml:"encryptionConfig"`
}

