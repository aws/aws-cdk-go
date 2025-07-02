package awsefs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Properties for the ReplicationConfiguration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var fileSystem fileSystem
//   var key key
//
//   replicationConfigurationProps := &ReplicationConfigurationProps{
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   	DestinationFileSystem: fileSystem,
//   	KmsKey: key,
//   	Region: jsii.String("region"),
//   }
//
type ReplicationConfigurationProps struct {
	// The availability zone name of the destination file system.
	//
	// One zone file system is used as the destination file system when this property is set.
	// Default: - no availability zone is set.
	//
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The existing destination file system for the replication.
	// Default: - None.
	//
	DestinationFileSystem IFileSystem `field:"optional" json:"destinationFileSystem" yaml:"destinationFileSystem"`
	// AWS KMS key used to protect the encrypted file system.
	// Default: - use service-managed KMS key for Amazon EFS.
	//
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// The AWS Region in which the destination file system is located.
	// Default: - the region of the stack.
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
}

