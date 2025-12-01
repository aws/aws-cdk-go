package awsathena


// Configuration settings for delivering logs to Amazon S3 buckets.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   managedLoggingConfigurationProperty := &ManagedLoggingConfigurationProperty{
//   	Enabled: jsii.Boolean(false),
//   	KmsKey: jsii.String("kmsKey"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-managedloggingconfiguration.html
//
type CfnWorkGroup_ManagedLoggingConfigurationProperty struct {
	// Enables mamanged log persistence.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-managedloggingconfiguration.html#cfn-athena-workgroup-managedloggingconfiguration-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The KMS key ARN to encrypt the logs stored in managed log persistence.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-managedloggingconfiguration.html#cfn-athena-workgroup-managedloggingconfiguration-kmskey
	//
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
}

