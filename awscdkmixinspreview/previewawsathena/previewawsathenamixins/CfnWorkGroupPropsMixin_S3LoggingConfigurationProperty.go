package previewawsathenamixins


// Configuration settings for delivering logs to Amazon S3 buckets.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   s3LoggingConfigurationProperty := &S3LoggingConfigurationProperty{
//   	Enabled: jsii.Boolean(false),
//   	KmsKey: jsii.String("kmsKey"),
//   	LogLocation: jsii.String("logLocation"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-s3loggingconfiguration.html
//
type CfnWorkGroupPropsMixin_S3LoggingConfigurationProperty struct {
	// Enables S3 log delivery.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-s3loggingconfiguration.html#cfn-athena-workgroup-s3loggingconfiguration-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The KMS key ARN to encrypt the logs published to the given Amazon S3 destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-s3loggingconfiguration.html#cfn-athena-workgroup-s3loggingconfiguration-kmskey
	//
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// The Amazon S3 destination URI for log publishing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-s3loggingconfiguration.html#cfn-athena-workgroup-s3loggingconfiguration-loglocation
	//
	LogLocation *string `field:"optional" json:"logLocation" yaml:"logLocation"`
}

