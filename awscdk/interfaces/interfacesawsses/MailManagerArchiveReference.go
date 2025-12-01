package interfacesawsses


// A reference to a MailManagerArchive resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mailManagerArchiveReference := &MailManagerArchiveReference{
//   	ArchiveArn: jsii.String("archiveArn"),
//   	ArchiveId: jsii.String("archiveId"),
//   }
//
type MailManagerArchiveReference struct {
	// The ARN of the MailManagerArchive resource.
	ArchiveArn *string `field:"required" json:"archiveArn" yaml:"archiveArn"`
	// The ArchiveId of the MailManagerArchive resource.
	ArchiveId *string `field:"required" json:"archiveId" yaml:"archiveId"`
}

