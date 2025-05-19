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
//   transformEncryptionProperty := &TransformEncryptionProperty{
//   	MlUserDataEncryption: &MLUserDataEncryptionProperty{
//   		MlUserDataEncryptionMode: jsii.String("mlUserDataEncryptionMode"),
//
//   		// the properties below are optional
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	TaskRunSecurityConfigurationName: jsii.String("taskRunSecurityConfigurationName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-mltransform-transformencryption.html
//
type CfnMLTransform_TransformEncryptionProperty struct {
	// The encryption-at-rest settings of the transform that apply to accessing user data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-mltransform-transformencryption.html#cfn-glue-mltransform-transformencryption-mluserdataencryption
	//
	MlUserDataEncryption interface{} `field:"optional" json:"mlUserDataEncryption" yaml:"mlUserDataEncryption"`
	// The name of the security configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-mltransform-transformencryption.html#cfn-glue-mltransform-transformencryption-taskrunsecurityconfigurationname
	//
	TaskRunSecurityConfigurationName *string `field:"optional" json:"taskRunSecurityConfigurationName" yaml:"taskRunSecurityConfigurationName"`
}

