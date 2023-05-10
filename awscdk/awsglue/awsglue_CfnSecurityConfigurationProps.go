package awsglue


// Properties for defining a `CfnSecurityConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSecurityConfigurationProps := &CfnSecurityConfigurationProps{
//   	EncryptionConfiguration: &EncryptionConfigurationProperty{
//   		CloudWatchEncryption: &CloudWatchEncryptionProperty{
//   			CloudWatchEncryptionMode: jsii.String("cloudWatchEncryptionMode"),
//   			KmsKeyArn: jsii.String("kmsKeyArn"),
//   		},
//   		JobBookmarksEncryption: &JobBookmarksEncryptionProperty{
//   			JobBookmarksEncryptionMode: jsii.String("jobBookmarksEncryptionMode"),
//   			KmsKeyArn: jsii.String("kmsKeyArn"),
//   		},
//   		S3Encryptions: []interface{}{
//   			&S3EncryptionProperty{
//   				KmsKeyArn: jsii.String("kmsKeyArn"),
//   				S3EncryptionMode: jsii.String("s3EncryptionMode"),
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
//   }
//
type CfnSecurityConfigurationProps struct {
	// The encryption configuration associated with this security configuration.
	EncryptionConfiguration interface{} `field:"required" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// The name of the security configuration.
	Name *string `field:"required" json:"name" yaml:"name"`
}

