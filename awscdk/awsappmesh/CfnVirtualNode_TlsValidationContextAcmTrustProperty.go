package awsappmesh


// An object that represents a Transport Layer Security (TLS) validation context trust for an Certificate Manager certificate.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tlsValidationContextAcmTrustProperty := &TlsValidationContextAcmTrustProperty{
//   	CertificateAuthorityArns: []*string{
//   		jsii.String("certificateAuthorityArns"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-tlsvalidationcontextacmtrust.html
//
type CfnVirtualNode_TlsValidationContextAcmTrustProperty struct {
	// One or more ACM Amazon Resource Name (ARN)s.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-tlsvalidationcontextacmtrust.html#cfn-appmesh-virtualnode-tlsvalidationcontextacmtrust-certificateauthorityarns
	//
	CertificateAuthorityArns *[]*string `field:"required" json:"certificateAuthorityArns" yaml:"certificateAuthorityArns"`
}

