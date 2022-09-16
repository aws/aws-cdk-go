package awsconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kinesisVideoStreamConfigProperty := &kinesisVideoStreamConfigProperty{
//   	prefix: jsii.String("prefix"),
//   	retentionPeriodHours: jsii.Number(123),
//
//   	// the properties below are optional
//   	encryptionConfig: &encryptionConfigProperty{
//   		encryptionType: jsii.String("encryptionType"),
//   		keyId: jsii.String("keyId"),
//   	},
//   }
//
type CfnInstanceStorageConfig_KinesisVideoStreamConfigProperty struct {
	// `CfnInstanceStorageConfig.KinesisVideoStreamConfigProperty.Prefix`.
	Prefix *string `field:"required" json:"prefix" yaml:"prefix"`
	// `CfnInstanceStorageConfig.KinesisVideoStreamConfigProperty.RetentionPeriodHours`.
	RetentionPeriodHours *float64 `field:"required" json:"retentionPeriodHours" yaml:"retentionPeriodHours"`
	// `CfnInstanceStorageConfig.KinesisVideoStreamConfigProperty.EncryptionConfig`.
	EncryptionConfig interface{} `field:"optional" json:"encryptionConfig" yaml:"encryptionConfig"`
}

