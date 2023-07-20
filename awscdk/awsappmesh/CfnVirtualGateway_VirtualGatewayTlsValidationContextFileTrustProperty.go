package awsappmesh


// An object that represents a Transport Layer Security (TLS) validation context trust for a local file.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   virtualGatewayTlsValidationContextFileTrustProperty := &VirtualGatewayTlsValidationContextFileTrustProperty{
//   	CertificateChain: jsii.String("certificateChain"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewaytlsvalidationcontextfiletrust.html
//
type CfnVirtualGateway_VirtualGatewayTlsValidationContextFileTrustProperty struct {
	// The certificate trust chain for a certificate stored on the file system of the virtual node that the proxy is running on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewaytlsvalidationcontextfiletrust.html#cfn-appmesh-virtualgateway-virtualgatewaytlsvalidationcontextfiletrust-certificatechain
	//
	CertificateChain *string `field:"required" json:"certificateChain" yaml:"certificateChain"`
}

