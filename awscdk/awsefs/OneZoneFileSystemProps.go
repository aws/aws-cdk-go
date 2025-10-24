package awsefs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Properties for configuring ReplicationConfiguration to replicate to a new One Zone file system.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var key Key
//
//   oneZoneFileSystemProps := &OneZoneFileSystemProps{
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   	Region: jsii.String("region"),
//
//   	// the properties below are optional
//   	KmsKey: key,
//   }
//
type OneZoneFileSystemProps struct {
	// The availability zone name of the destination file system.
	//
	// One zone file system is used as the destination file system when this property is set.
	AvailabilityZone *string `field:"required" json:"availabilityZone" yaml:"availabilityZone"`
	// The AWS Region in which the destination file system is located.
	Region *string `field:"required" json:"region" yaml:"region"`
	// AWS KMS key used to protect the encrypted file system.
	// Default: - use service-managed KMS key for Amazon EFS.
	//
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
}

