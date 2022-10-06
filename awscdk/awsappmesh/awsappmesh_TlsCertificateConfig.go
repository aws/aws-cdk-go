package awsappmesh


// A wrapper for the tls config returned by {@link TlsCertificate.bind}.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tlsCertificateConfig := &tlsCertificateConfig{
//   	tlsCertificate: &listenerTlsCertificateProperty{
//   		acm: &listenerTlsAcmCertificateProperty{
//   			certificateArn: jsii.String("certificateArn"),
//   		},
//   		file: &listenerTlsFileCertificateProperty{
//   			certificateChain: jsii.String("certificateChain"),
//   			privateKey: jsii.String("privateKey"),
//   		},
//   		sds: &listenerTlsSdsCertificateProperty{
//   			secretName: jsii.String("secretName"),
//   		},
//   	},
//   }
//
type TlsCertificateConfig struct {
	// The CFN shape for a TLS certificate.
	TlsCertificate *CfnVirtualNode_ListenerTlsCertificateProperty `field:"required" json:"tlsCertificate" yaml:"tlsCertificate"`
}

