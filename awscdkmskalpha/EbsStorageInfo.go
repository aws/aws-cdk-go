package awscdkmskalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// EBS volume information.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import msk_alpha "github.com/aws/aws-cdk-go/awscdkmskalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var keyRef iKeyRef
//
//   ebsStorageInfo := &EbsStorageInfo{
//   	EncryptionKey: keyRef,
//   	VolumeSize: jsii.Number(123),
//   }
//
// Experimental.
type EbsStorageInfo struct {
	// The AWS KMS key for encrypting data at rest.
	// Default: Uses AWS managed CMK (aws/kafka).
	//
	// Experimental.
	EncryptionKey awskms.IKeyRef `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// The size in GiB of the EBS volume for the data drive on each broker node.
	// Default: 1000.
	//
	// Experimental.
	VolumeSize *float64 `field:"optional" json:"volumeSize" yaml:"volumeSize"`
}

