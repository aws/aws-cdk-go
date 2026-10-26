package interfacesawsmediaconvert


// A reference to a CertificateAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   certificateAssociationReference := &CertificateAssociationReference{
//   	CertificateAssociationArn: jsii.String("certificateAssociationArn"),
//   }
//
type CertificateAssociationReference struct {
	// The Arn of the CertificateAssociation resource.
	CertificateAssociationArn *string `field:"required" json:"certificateAssociationArn" yaml:"certificateAssociationArn"`
}

