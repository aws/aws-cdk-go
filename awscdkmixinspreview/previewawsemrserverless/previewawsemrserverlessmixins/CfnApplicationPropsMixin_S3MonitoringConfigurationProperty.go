package previewawsemrserverlessmixins


// The Amazon S3 configuration for monitoring log publishing.
//
// You can configure your jobs to send log information to Amazon S3.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   s3MonitoringConfigurationProperty := &S3MonitoringConfigurationProperty{
//   	EncryptionKeyArn: jsii.String("encryptionKeyArn"),
//   	LogUri: jsii.String("logUri"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-s3monitoringconfiguration.html
//
type CfnApplicationPropsMixin_S3MonitoringConfigurationProperty struct {
	// The KMS key ARN to encrypt the logs published to the given Amazon S3 destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-s3monitoringconfiguration.html#cfn-emrserverless-application-s3monitoringconfiguration-encryptionkeyarn
	//
	EncryptionKeyArn *string `field:"optional" json:"encryptionKeyArn" yaml:"encryptionKeyArn"`
	// The Amazon S3 destination URI for log publishing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-s3monitoringconfiguration.html#cfn-emrserverless-application-s3monitoringconfiguration-loguri
	//
	LogUri *string `field:"optional" json:"logUri" yaml:"logUri"`
}

