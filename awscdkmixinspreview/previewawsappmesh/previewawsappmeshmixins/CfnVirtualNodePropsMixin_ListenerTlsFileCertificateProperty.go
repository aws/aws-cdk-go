package previewawsappmeshmixins


// An object that represents a local file certificate.
//
// The certificate must meet specific requirements and you must have proxy authorization enabled. For more information, see [Transport Layer Security (TLS)](https://docs.aws.amazon.com/app-mesh/latest/userguide/tls.html#virtual-node-tls-prerequisites) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   listenerTlsFileCertificateProperty := &ListenerTlsFileCertificateProperty{
//   	CertificateChain: jsii.String("certificateChain"),
//   	PrivateKey: jsii.String("privateKey"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-listenertlsfilecertificate.html
//
type CfnVirtualNodePropsMixin_ListenerTlsFileCertificateProperty struct {
	// The certificate chain for the certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-listenertlsfilecertificate.html#cfn-appmesh-virtualnode-listenertlsfilecertificate-certificatechain
	//
	CertificateChain *string `field:"optional" json:"certificateChain" yaml:"certificateChain"`
	// The private key for a certificate stored on the file system of the virtual node that the proxy is running on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-listenertlsfilecertificate.html#cfn-appmesh-virtualnode-listenertlsfilecertificate-privatekey
	//
	PrivateKey *string `field:"optional" json:"privateKey" yaml:"privateKey"`
}

