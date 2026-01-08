package awsefs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsefs"
)

// Properties for configuring ReplicationConfiguration to replicate to an existing file system.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var fileSystemRef IFileSystemRef
//
//   existingFileSystemProps := &ExistingFileSystemProps{
//   	DestinationFileSystem: fileSystemRef,
//   }
//
type ExistingFileSystemProps struct {
	// The existing destination file system for the replication.
	DestinationFileSystem interfacesawsefs.IFileSystemRef `field:"required" json:"destinationFileSystem" yaml:"destinationFileSystem"`
}

