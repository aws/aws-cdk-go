package interfacesawseks


// A reference to a CertificateAuthority resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   certificateAuthorityReference := &CertificateAuthorityReference{
//   	CertificateAuthorityId: jsii.String("certificateAuthorityId"),
//   	ClusterName: jsii.String("clusterName"),
//   }
//
type CertificateAuthorityReference struct {
	// The Id of the CertificateAuthority resource.
	CertificateAuthorityId *string `field:"required" json:"certificateAuthorityId" yaml:"certificateAuthorityId"`
	// The ClusterName of the CertificateAuthority resource.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
}

