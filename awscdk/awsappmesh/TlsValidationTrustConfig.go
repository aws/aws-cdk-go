package awsappmesh


// All Properties for TLS Validation Trusts for both Client Policy and Listener.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tlsValidationTrustConfig := &TlsValidationTrustConfig{
//   	TlsValidationTrust: &TlsValidationContextTrustProperty{
//   		Acm: &TlsValidationContextAcmTrustProperty{
//   			CertificateAuthorityArns: []*string{
//   				jsii.String("certificateAuthorityArns"),
//   			},
//   		},
//   		File: &TlsValidationContextFileTrustProperty{
//   			CertificateChain: jsii.String("certificateChain"),
//   		},
//   		Sds: &TlsValidationContextSdsTrustProperty{
//   			SecretName: jsii.String("secretName"),
//   		},
//   	},
//   }
//
type TlsValidationTrustConfig struct {
	// VirtualNode CFN configuration for client policy's TLS Validation Trust.
	TlsValidationTrust *CfnVirtualNode_TlsValidationContextTrustProperty `field:"required" json:"tlsValidationTrust" yaml:"tlsValidationTrust"`
}

