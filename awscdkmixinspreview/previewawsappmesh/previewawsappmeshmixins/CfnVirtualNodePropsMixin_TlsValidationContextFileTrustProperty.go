package previewawsappmeshmixins


// An object that represents a Transport Layer Security (TLS) validation context trust for a local file.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   tlsValidationContextFileTrustProperty := &TlsValidationContextFileTrustProperty{
//   	CertificateChain: jsii.String("certificateChain"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-tlsvalidationcontextfiletrust.html
//
type CfnVirtualNodePropsMixin_TlsValidationContextFileTrustProperty struct {
	// The certificate trust chain for a certificate stored on the file system of the virtual node that the proxy is running on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-tlsvalidationcontextfiletrust.html#cfn-appmesh-virtualnode-tlsvalidationcontextfiletrust-certificatechain
	//
	CertificateChain *string `field:"optional" json:"certificateChain" yaml:"certificateChain"`
}

