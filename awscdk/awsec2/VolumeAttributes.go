package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Attributes required to import an existing EBS Volume into the Stack.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var key key
//
//   volumeAttributes := &VolumeAttributes{
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   	VolumeId: jsii.String("volumeId"),
//
//   	// the properties below are optional
//   	EncryptionKey: key,
//   }
//
type VolumeAttributes struct {
	// The availability zone that the EBS Volume is contained within (ex: us-west-2a).
	AvailabilityZone *string `field:"required" json:"availabilityZone" yaml:"availabilityZone"`
	// The EBS Volume's ID.
	VolumeId *string `field:"required" json:"volumeId" yaml:"volumeId"`
	// The customer-managed encryption key that is used to encrypt the Volume.
	// Default: None -- The EBS Volume is not using a customer-managed KMS key for encryption.
	//
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
}

