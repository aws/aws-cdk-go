package interfacesawsec2


// A reference to a EnclaveCertificateIamRoleAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   enclaveCertificateIamRoleAssociationReference := &EnclaveCertificateIamRoleAssociationReference{
//   	CertificateArn: jsii.String("certificateArn"),
//   	RoleArn: jsii.String("roleArn"),
//   }
//
type EnclaveCertificateIamRoleAssociationReference struct {
	// The CertificateArn of the EnclaveCertificateIamRoleAssociation resource.
	CertificateArn *string `field:"required" json:"certificateArn" yaml:"certificateArn"`
	// The RoleArn of the EnclaveCertificateIamRoleAssociation resource.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}

