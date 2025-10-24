package awsefs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Properties for configuring ReplicationConfiguration to replicate to a new Regional file system.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var key Key
//
//   regionalFileSystemProps := &RegionalFileSystemProps{
//   	KmsKey: key,
//   	Region: jsii.String("region"),
//   }
//
type RegionalFileSystemProps struct {
	// AWS KMS key used to protect the encrypted file system.
	// Default: - use service-managed KMS key for Amazon EFS.
	//
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// The AWS Region in which the destination file system is located.
	// Default: - the region of the stack.
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
}

