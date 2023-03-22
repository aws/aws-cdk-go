package awsappmesh


// A wrapper for the tls config returned by `TlsCertificate.bind`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tlsCertificateConfig := &TlsCertificateConfig{
//   	TlsCertificate: &ListenerTlsCertificateProperty{
//   		Acm: &ListenerTlsAcmCertificateProperty{
//   			CertificateArn: jsii.String("certificateArn"),
//   		},
//   		File: &ListenerTlsFileCertificateProperty{
//   			CertificateChain: jsii.String("certificateChain"),
//   			PrivateKey: jsii.String("privateKey"),
//   		},
//   		Sds: &ListenerTlsSdsCertificateProperty{
//   			SecretName: jsii.String("secretName"),
//   		},
//   	},
//   }
//
type TlsCertificateConfig struct {
	// The CFN shape for a TLS certificate.
	TlsCertificate *CfnVirtualNode_ListenerTlsCertificateProperty `field:"required" json:"tlsCertificate" yaml:"tlsCertificate"`
}

