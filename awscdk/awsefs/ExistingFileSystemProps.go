package awsefs


// Properties for configuring ReplicationConfiguration to replicate to an existing file system.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var fileSystem fileSystem
//
//   existingFileSystemProps := &ExistingFileSystemProps{
//   	DestinationFileSystem: fileSystem,
//   }
//
type ExistingFileSystemProps struct {
	// The existing destination file system for the replication.
	DestinationFileSystem IFileSystem `field:"required" json:"destinationFileSystem" yaml:"destinationFileSystem"`
}

