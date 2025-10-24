package awss3

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// The properties for the destination bucket for granting replication permission.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket Bucket
//   var key Key
//
//   grantReplicationPermissionDestinationProps := &GrantReplicationPermissionDestinationProps{
//   	Bucket: bucket,
//
//   	// the properties below are optional
//   	EncryptionKey: key,
//   }
//
type GrantReplicationPermissionDestinationProps struct {
	// The destination bucket.
	Bucket IBucket `field:"required" json:"bucket" yaml:"bucket"`
	// The KMS key to use for encryption if a destination bucket needs to be encrypted with a customer-managed KMS key.
	// Default: - no KMS key is used for replication.
	//
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
}

