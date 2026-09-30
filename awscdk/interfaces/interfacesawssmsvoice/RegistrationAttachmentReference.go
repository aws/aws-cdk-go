package interfacesawssmsvoice


// A reference to a RegistrationAttachment resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   registrationAttachmentReference := &RegistrationAttachmentReference{
//   	RegistrationAttachmentArn: jsii.String("registrationAttachmentArn"),
//   }
//
type RegistrationAttachmentReference struct {
	// The RegistrationAttachmentArn of the RegistrationAttachment resource.
	RegistrationAttachmentArn *string `field:"required" json:"registrationAttachmentArn" yaml:"registrationAttachmentArn"`
}

