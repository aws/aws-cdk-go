package awssecretsmanager

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Secret replica region.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var key key
//
//   replicaRegion := &ReplicaRegion{
//   	Region: jsii.String("region"),
//
//   	// the properties below are optional
//   	EncryptionKey: key,
//   }
//
type ReplicaRegion struct {
	// The name of the region.
	Region *string `field:"required" json:"region" yaml:"region"`
	// The customer-managed encryption key to use for encrypting the secret value.
	// Default: - A default KMS key for the account and region is used.
	//
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
}

