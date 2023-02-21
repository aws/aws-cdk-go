package awsglue


// Properties for defining a `CfnSecurityConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSecurityConfigurationProps := &cfnSecurityConfigurationProps{
//   	encryptionConfiguration: &encryptionConfigurationProperty{
//   		cloudWatchEncryption: &cloudWatchEncryptionProperty{
//   			cloudWatchEncryptionMode: jsii.String("cloudWatchEncryptionMode"),
//   			kmsKeyArn: jsii.String("kmsKeyArn"),
//   		},
//   		jobBookmarksEncryption: &jobBookmarksEncryptionProperty{
//   			jobBookmarksEncryptionMode: jsii.String("jobBookmarksEncryptionMode"),
//   			kmsKeyArn: jsii.String("kmsKeyArn"),
//   		},
//   		s3Encryptions: []interface{}{
//   			&s3EncryptionProperty{
//   				kmsKeyArn: jsii.String("kmsKeyArn"),
//   				s3EncryptionMode: jsii.String("s3EncryptionMode"),
//   			},
//   		},
//   	},
//   	name: jsii.String("name"),
//   }
//
type CfnSecurityConfigurationProps struct {
	// The encryption configuration associated with this security configuration.
	EncryptionConfiguration interface{} `field:"required" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// The name of the security configuration.
	Name *string `field:"required" json:"name" yaml:"name"`
}

