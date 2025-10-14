package awsappmesh


// An object that represents a Transport Layer Security (TLS) validation context trust for an Certificate Manager certificate.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   virtualGatewayTlsValidationContextAcmTrustProperty := &VirtualGatewayTlsValidationContextAcmTrustProperty{
//   	CertificateAuthorityArns: []*string{
//   		jsii.String("certificateAuthorityArns"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewaytlsvalidationcontextacmtrust.html
//
type CfnVirtualGateway_VirtualGatewayTlsValidationContextAcmTrustProperty struct {
	// One or more ACM Amazon Resource Name (ARN)s.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewaytlsvalidationcontextacmtrust.html#cfn-appmesh-virtualgateway-virtualgatewaytlsvalidationcontextacmtrust-certificateauthorityarns
	//
	CertificateAuthorityArns *[]*string `field:"required" json:"certificateAuthorityArns" yaml:"certificateAuthorityArns"`
}

