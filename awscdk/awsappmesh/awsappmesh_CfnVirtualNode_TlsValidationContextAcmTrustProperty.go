package awsappmesh


// An object that represents a Transport Layer Security (TLS) validation context trust for an AWS Certificate Manager certificate.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tlsValidationContextAcmTrustProperty := &tlsValidationContextAcmTrustProperty{
//   	certificateAuthorityArns: []*string{
//   		jsii.String("certificateAuthorityArns"),
//   	},
//   }
//
type CfnVirtualNode_TlsValidationContextAcmTrustProperty struct {
	// One or more ACM Amazon Resource Name (ARN)s.
	CertificateAuthorityArns *[]*string `field:"required" json:"certificateAuthorityArns" yaml:"certificateAuthorityArns"`
}

