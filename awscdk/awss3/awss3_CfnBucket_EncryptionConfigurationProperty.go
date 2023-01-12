package awss3


// Specifies encryption-related information for an Amazon S3 bucket that is a destination for replicated objects.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   encryptionConfigurationProperty := &encryptionConfigurationProperty{
//   	replicaKmsKeyId: jsii.String("replicaKmsKeyId"),
//   }
//
type CfnBucket_EncryptionConfigurationProperty struct {
	// Specifies the ID (Key ARN or Alias ARN) of the customer managed AWS KMS key stored in AWS Key Management Service (KMS) for the destination bucket.
	//
	// Amazon S3 uses this key to encrypt replica objects. Amazon S3 only supports symmetric, customer managed KMS keys. For more information, see [Using symmetric and asymmetric keys](https://docs.aws.amazon.com//kms/latest/developerguide/symmetric-asymmetric.html) in the *AWS Key Management Service Developer Guide* .
	ReplicaKmsKeyId *string `field:"required" json:"replicaKmsKeyId" yaml:"replicaKmsKeyId"`
}

