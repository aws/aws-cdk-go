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
//   listenerTlsFileCertificateProperty := &listenerTlsFileCertificateProperty{
//   	certificateChain: jsii.String("certificateChain"),
//   	privateKey: jsii.String("privateKey"),
//   }
//
type CfnVirtualNode_ListenerTlsFileCertificateProperty struct {
	// The certificate chain for the certificate.
	CertificateChain *string `field:"required" json:"certificateChain" yaml:"certificateChain"`
	// The private key for a certificate stored on the file system of the virtual node that the proxy is running on.
	PrivateKey *string `field:"required" json:"privateKey" yaml:"privateKey"`
}

