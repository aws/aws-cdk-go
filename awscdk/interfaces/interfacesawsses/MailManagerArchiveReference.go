package interfacesawsses


// A reference to a MailManagerArchive resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mailManagerArchiveReference := &MailManagerArchiveReference{
//   	ArchiveId: jsii.String("archiveId"),
//   }
//
type MailManagerArchiveReference struct {
	// The ArchiveId of the MailManagerArchive resource.
	ArchiveId *string `field:"required" json:"archiveId" yaml:"archiveId"`
}

