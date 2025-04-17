package awsmsk


// The data-volume encryption details.
//
// You can't update encryption at rest settings for existing clusters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   encryptionAtRestProperty := &EncryptionAtRestProperty{
//   	DataVolumeKmsKeyId: jsii.String("dataVolumeKmsKeyId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-encryptionatrest.html
//
type CfnCluster_EncryptionAtRestProperty struct {
	// The ARN of the Amazon KMS key for encrypting data at rest.
	//
	// If you don't specify a KMS key, MSK creates one for you and uses it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-encryptionatrest.html#cfn-msk-cluster-encryptionatrest-datavolumekmskeyid
	//
	DataVolumeKmsKeyId *string `field:"required" json:"dataVolumeKmsKeyId" yaml:"dataVolumeKmsKeyId"`
}

