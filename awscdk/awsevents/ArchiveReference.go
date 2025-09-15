package awsevents


// A reference to a Archive resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   archiveReference := &ArchiveReference{
//   	ArchiveArn: jsii.String("archiveArn"),
//   	ArchiveName: jsii.String("archiveName"),
//   }
//
type ArchiveReference struct {
	// The ARN of the Archive resource.
	ArchiveArn *string `field:"required" json:"archiveArn" yaml:"archiveArn"`
	// The ArchiveName of the Archive resource.
	ArchiveName *string `field:"required" json:"archiveName" yaml:"archiveName"`
}

