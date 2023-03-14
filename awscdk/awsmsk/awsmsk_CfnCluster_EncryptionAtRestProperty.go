package awsmsk


// The data volume encryption details.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   encryptionAtRestProperty := &encryptionAtRestProperty{
//   	dataVolumeKmsKeyId: jsii.String("dataVolumeKmsKeyId"),
//   }
//
type CfnCluster_EncryptionAtRestProperty struct {
	// The ARN of the Amazon KMS key for encrypting data at rest.
	//
	// If you don't specify a KMS key, MSK creates one for you and uses it on your behalf.
	DataVolumeKmsKeyId *string `field:"required" json:"dataVolumeKmsKeyId" yaml:"dataVolumeKmsKeyId"`
}

