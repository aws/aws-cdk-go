package awsec2


// Information about the client certificate to be used for authentication.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   certificateAuthenticationRequestProperty := &CertificateAuthenticationRequestProperty{
//   	ClientRootCertificateChainArn: jsii.String("clientRootCertificateChainArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-certificateauthenticationrequest.html
//
type CfnClientVpnEndpoint_CertificateAuthenticationRequestProperty struct {
	// The ARN of the client certificate.
	//
	// The certificate must be signed by a certificate authority (CA) and it must be provisioned in AWS Certificate Manager (ACM).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-certificateauthenticationrequest.html#cfn-ec2-clientvpnendpoint-certificateauthenticationrequest-clientrootcertificatechainarn
	//
	ClientRootCertificateChainArn *string `field:"required" json:"clientRootCertificateChainArn" yaml:"clientRootCertificateChainArn"`
}

