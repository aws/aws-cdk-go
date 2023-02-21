package awsec2


// Information about the client certificate to be used for authentication.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   certificateAuthenticationRequestProperty := &certificateAuthenticationRequestProperty{
//   	clientRootCertificateChainArn: jsii.String("clientRootCertificateChainArn"),
//   }
//
type CfnClientVpnEndpoint_CertificateAuthenticationRequestProperty struct {
	// The ARN of the client certificate.
	//
	// The certificate must be signed by a certificate authority (CA) and it must be provisioned in AWS Certificate Manager (ACM).
	ClientRootCertificateChainArn *string `field:"required" json:"clientRootCertificateChainArn" yaml:"clientRootCertificateChainArn"`
}

