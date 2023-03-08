package awsglue


// The encryption-at-rest settings of the transform that apply to accessing user data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mLUserDataEncryptionProperty := &mLUserDataEncryptionProperty{
//   	mlUserDataEncryptionMode: jsii.String("mlUserDataEncryptionMode"),
//
//   	// the properties below are optional
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   }
//
type CfnMLTransform_MLUserDataEncryptionProperty struct {
	// The encryption mode applied to user data. Valid values are:.
	//
	// - DISABLED: encryption is disabled.
	// - SSEKMS: use of server-side encryption with AWS Key Management Service (SSE-KMS) for user data
	// stored in Amazon S3.
	MlUserDataEncryptionMode *string `field:"required" json:"mlUserDataEncryptionMode" yaml:"mlUserDataEncryptionMode"`
	// The ID for the customer-provided KMS key.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

