package awsmsk

import (
	"github.com/aws/aws-cdk-go/awscdk/awskms"
)

// EBS volume information.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var key key
//
//   ebsStorageInfo := &ebsStorageInfo{
//   	encryptionKey: key,
//   	volumeSize: jsii.Number(123),
//   }
//
// Experimental.
type EbsStorageInfo struct {
	// The AWS KMS key for encrypting data at rest.
	// Experimental.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// The size in GiB of the EBS volume for the data drive on each broker node.
	// Experimental.
	VolumeSize *float64 `field:"optional" json:"volumeSize" yaml:"volumeSize"`
}

