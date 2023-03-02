package awsmsk


// Includes encryption-related information, such as the Amazon KMS key used for encrypting data at rest and whether you want MSK to encrypt your data in transit.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   encryptionInfoProperty := &encryptionInfoProperty{
//   	encryptionAtRest: &encryptionAtRestProperty{
//   		dataVolumeKmsKeyId: jsii.String("dataVolumeKmsKeyId"),
//   	},
//   	encryptionInTransit: &encryptionInTransitProperty{
//   		clientBroker: jsii.String("clientBroker"),
//   		inCluster: jsii.Boolean(false),
//   	},
//   }
//
type CfnCluster_EncryptionInfoProperty struct {
	// The data-volume encryption details.
	EncryptionAtRest interface{} `field:"optional" json:"encryptionAtRest" yaml:"encryptionAtRest"`
	// The details for encryption in transit.
	EncryptionInTransit interface{} `field:"optional" json:"encryptionInTransit" yaml:"encryptionInTransit"`
}

