package awsmsk


// Includes encryption-related information, such as the Amazon KMS key used for encrypting data at rest and whether you want MSK to encrypt your data in transit.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   encryptionInfoProperty := &EncryptionInfoProperty{
//   	EncryptionAtRest: &EncryptionAtRestProperty{
//   		DataVolumeKmsKeyId: jsii.String("dataVolumeKmsKeyId"),
//   	},
//   	EncryptionInTransit: &EncryptionInTransitProperty{
//   		ClientBroker: jsii.String("clientBroker"),
//   		InCluster: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-encryptioninfo.html
//
type CfnCluster_EncryptionInfoProperty struct {
	// The data-volume encryption details.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-encryptioninfo.html#cfn-msk-cluster-encryptioninfo-encryptionatrest
	//
	EncryptionAtRest interface{} `field:"optional" json:"encryptionAtRest" yaml:"encryptionAtRest"`
	// The details for encryption in transit.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-encryptioninfo.html#cfn-msk-cluster-encryptioninfo-encryptionintransit
	//
	EncryptionInTransit interface{} `field:"optional" json:"encryptionInTransit" yaml:"encryptionInTransit"`
}

