package awskendra


// Provides the identifier of the AWS KMS customer master key (CMK) used to encrypt data indexed by Amazon Kendra.
//
// We suggest that you use a CMK from your account to help secure your index. Amazon Kendra doesn't support asymmetric CMKs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serverSideEncryptionConfigurationProperty := &serverSideEncryptionConfigurationProperty{
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   }
//
type CfnIndex_ServerSideEncryptionConfigurationProperty struct {
	// The identifier of the AWS KMS key .
	//
	// Amazon Kendra doesn't support asymmetric keys.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

