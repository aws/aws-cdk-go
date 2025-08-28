package awsglue


// The encryption-at-rest settings of the transform that apply to accessing user data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mLUserDataEncryptionProperty := &MLUserDataEncryptionProperty{
//   	MlUserDataEncryptionMode: jsii.String("mlUserDataEncryptionMode"),
//
//   	// the properties below are optional
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-mltransform-mluserdataencryption.html
//
type CfnMLTransform_MLUserDataEncryptionProperty struct {
	// The encryption mode applied to user data. Valid values are:.
	//
	// - DISABLED: encryption is disabled.
	// - SSEKMS: use of server-side encryption with AWS Key Management Service (SSE-KMS) for user data
	// stored in Amazon S3.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-mltransform-mluserdataencryption.html#cfn-glue-mltransform-mluserdataencryption-mluserdataencryptionmode
	//
	MlUserDataEncryptionMode *string `field:"required" json:"mlUserDataEncryptionMode" yaml:"mlUserDataEncryptionMode"`
	// The ID for the customer-provided KMS key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-mltransform-mluserdataencryption.html#cfn-glue-mltransform-mluserdataencryption-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

