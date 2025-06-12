package awsappmesh


// An object that represents a local file certificate.
//
// The certificate must meet specific requirements and you must have proxy authorization enabled. For more information, see [Transport Layer Security (TLS)](https://docs.aws.amazon.com/app-mesh/latest/userguide/tls.html#virtual-node-tls-prerequisites) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   virtualGatewayListenerTlsFileCertificateProperty := &VirtualGatewayListenerTlsFileCertificateProperty{
//   	CertificateChain: jsii.String("certificateChain"),
//   	PrivateKey: jsii.String("privateKey"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewaylistenertlsfilecertificate.html
//
type CfnVirtualGateway_VirtualGatewayListenerTlsFileCertificateProperty struct {
	// The certificate chain for the certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewaylistenertlsfilecertificate.html#cfn-appmesh-virtualgateway-virtualgatewaylistenertlsfilecertificate-certificatechain
	//
	CertificateChain *string `field:"required" json:"certificateChain" yaml:"certificateChain"`
	// The private key for a certificate stored on the file system of the mesh endpoint that the proxy is running on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewaylistenertlsfilecertificate.html#cfn-appmesh-virtualgateway-virtualgatewaylistenertlsfilecertificate-privatekey
	//
	PrivateKey *string `field:"required" json:"privateKey" yaml:"privateKey"`
}

