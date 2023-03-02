package awsglue


// The encryption-at-rest settings of the transform that apply to accessing user data.
//
// Machine learning
// transforms can access user data encrypted in Amazon S3 using KMS.
//
// Additionally, imported labels and trained transforms can now be encrypted using a customer provided
// KMS key.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   transformEncryptionProperty := &transformEncryptionProperty{
//   	mlUserDataEncryption: &mLUserDataEncryptionProperty{
//   		mlUserDataEncryptionMode: jsii.String("mlUserDataEncryptionMode"),
//
//   		// the properties below are optional
//   		kmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	taskRunSecurityConfigurationName: jsii.String("taskRunSecurityConfigurationName"),
//   }
//
type CfnMLTransform_TransformEncryptionProperty struct {
	// The encryption-at-rest settings of the transform that apply to accessing user data.
	MlUserDataEncryption interface{} `field:"optional" json:"mlUserDataEncryption" yaml:"mlUserDataEncryption"`
	// The name of the security configuration.
	TaskRunSecurityConfigurationName *string `field:"optional" json:"taskRunSecurityConfigurationName" yaml:"taskRunSecurityConfigurationName"`
}

