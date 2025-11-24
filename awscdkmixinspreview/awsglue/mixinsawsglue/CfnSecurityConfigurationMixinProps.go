package mixinsawsglue


// Properties for CfnSecurityConfigurationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnSecurityConfigurationMixinProps := &CfnSecurityConfigurationMixinProps{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-securityconfiguration.html
//
type CfnSecurityConfigurationMixinProps struct {
	// The encryption configuration associated with this security configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-securityconfiguration.html#cfn-glue-securityconfiguration-encryptionconfiguration
	//
	EncryptionConfiguration interface{} `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// The name of the security configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-securityconfiguration.html#cfn-glue-securityconfiguration-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

