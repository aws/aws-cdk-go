package interfacesawsacmpca


// A reference to a Permission resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   permissionReference := &PermissionReference{
//   	CertificateAuthorityArn: jsii.String("certificateAuthorityArn"),
//   	Principal: jsii.String("principal"),
//   }
//
type PermissionReference struct {
	// The CertificateAuthorityArn of the Permission resource.
	CertificateAuthorityArn *string `field:"required" json:"certificateAuthorityArn" yaml:"certificateAuthorityArn"`
	// The Principal of the Permission resource.
	Principal *string `field:"required" json:"principal" yaml:"principal"`
}

