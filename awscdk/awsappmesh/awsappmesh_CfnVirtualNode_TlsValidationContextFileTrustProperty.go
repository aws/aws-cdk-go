package awsappmesh


// An object that represents a Transport Layer Security (TLS) validation context trust for a local file.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tlsValidationContextFileTrustProperty := &tlsValidationContextFileTrustProperty{
//   	certificateChain: jsii.String("certificateChain"),
//   }
//
type CfnVirtualNode_TlsValidationContextFileTrustProperty struct {
	// The certificate trust chain for a certificate stored on the file system of the virtual node that the proxy is running on.
	CertificateChain *string `field:"required" json:"certificateChain" yaml:"certificateChain"`
}

