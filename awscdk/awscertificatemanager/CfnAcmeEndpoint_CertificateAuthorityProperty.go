package awscertificatemanager


// The certificate authority configuration for the ACME endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   certificateAuthorityProperty := &CertificateAuthorityProperty{
//   	PublicCertificateAuthority: &PublicCertificateAuthorityProperty{
//   		AllowedKeyAlgorithms: []*string{
//   			jsii.String("allowedKeyAlgorithms"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-certificatemanager-acmeendpoint-certificateauthority.html
//
type CfnAcmeEndpoint_CertificateAuthorityProperty struct {
	// Configuration for the public certificate authority.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-certificatemanager-acmeendpoint-certificateauthority.html#cfn-certificatemanager-acmeendpoint-certificateauthority-publiccertificateauthority
	//
	PublicCertificateAuthority interface{} `field:"required" json:"publicCertificateAuthority" yaml:"publicCertificateAuthority"`
}

