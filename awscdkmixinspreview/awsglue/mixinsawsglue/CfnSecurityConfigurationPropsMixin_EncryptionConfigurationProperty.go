package mixinsawsglue


// Specifies an encryption configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   encryptionConfigurationProperty := &EncryptionConfigurationProperty{
//   	CloudWatchEncryption: &CloudWatchEncryptionProperty{
//   		CloudWatchEncryptionMode: jsii.String("cloudWatchEncryptionMode"),
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   	},
//   	JobBookmarksEncryption: &JobBookmarksEncryptionProperty{
//   		JobBookmarksEncryptionMode: jsii.String("jobBookmarksEncryptionMode"),
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   	},
//   	S3Encryptions: []interface{}{
//   		&S3EncryptionProperty{
//   			KmsKeyArn: jsii.String("kmsKeyArn"),
//   			S3EncryptionMode: jsii.String("s3EncryptionMode"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-securityconfiguration-encryptionconfiguration.html
//
type CfnSecurityConfigurationPropsMixin_EncryptionConfigurationProperty struct {
	// The encryption configuration for Amazon CloudWatch.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-securityconfiguration-encryptionconfiguration.html#cfn-glue-securityconfiguration-encryptionconfiguration-cloudwatchencryption
	//
	CloudWatchEncryption interface{} `field:"optional" json:"cloudWatchEncryption" yaml:"cloudWatchEncryption"`
	// The encryption configuration for job bookmarks.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-securityconfiguration-encryptionconfiguration.html#cfn-glue-securityconfiguration-encryptionconfiguration-jobbookmarksencryption
	//
	JobBookmarksEncryption interface{} `field:"optional" json:"jobBookmarksEncryption" yaml:"jobBookmarksEncryption"`
	// The encyption configuration for Amazon Simple Storage Service (Amazon S3) data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-securityconfiguration-encryptionconfiguration.html#cfn-glue-securityconfiguration-encryptionconfiguration-s3encryptions
	//
	S3Encryptions interface{} `field:"optional" json:"s3Encryptions" yaml:"s3Encryptions"`
}

