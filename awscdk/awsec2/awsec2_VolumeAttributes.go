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
//   volumeAttributes := &volumeAttributes{
//   	availabilityZone: jsii.String("availabilityZone"),
//   	volumeId: jsii.String("volumeId"),
//
//   	// the properties below are optional
//   	encryptionKey: key,
//   }
//
type VolumeAttributes struct {
	// The availability zone that the EBS Volume is contained within (ex: us-west-2a).
	AvailabilityZone *string `field:"required" json:"availabilityZone" yaml:"availabilityZone"`
	// The EBS Volume's ID.
	VolumeId *string `field:"required" json:"volumeId" yaml:"volumeId"`
	// The customer-managed encryption key that is used to encrypt the Volume.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
}

