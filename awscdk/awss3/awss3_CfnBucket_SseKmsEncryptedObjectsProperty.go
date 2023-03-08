package awss3


// A container for filter information for the selection of S3 objects encrypted with AWS KMS.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sseKmsEncryptedObjectsProperty := &sseKmsEncryptedObjectsProperty{
//   	status: jsii.String("status"),
//   }
//
type CfnBucket_SseKmsEncryptedObjectsProperty struct {
	// Specifies whether Amazon S3 replicates objects created with server-side encryption using an AWS KMS key stored in AWS Key Management Service.
	Status *string `field:"required" json:"status" yaml:"status"`
}

