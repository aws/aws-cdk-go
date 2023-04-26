package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/awskms"
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
// Experimental.
type VolumeAttributes struct {
	// The availability zone that the EBS Volume is contained within (ex: us-west-2a).
	// Experimental.
	AvailabilityZone *string `field:"required" json:"availabilityZone" yaml:"availabilityZone"`
	// The EBS Volume's ID.
	// Experimental.
	VolumeId *string `field:"required" json:"volumeId" yaml:"volumeId"`
	// The customer-managed encryption key that is used to encrypt the Volume.
	// Experimental.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
}

