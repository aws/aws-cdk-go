package awsappmesh


// An object that represents the Transport Layer Security (TLS) properties for a listener.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   virtualGatewayListenerTlsProperty := &VirtualGatewayListenerTlsProperty{
//   	Certificate: &VirtualGatewayListenerTlsCertificateProperty{
//   		Acm: &VirtualGatewayListenerTlsAcmCertificateProperty{
//   			CertificateArn: jsii.String("certificateArn"),
//   		},
//   		File: &VirtualGatewayListenerTlsFileCertificateProperty{
//   			CertificateChain: jsii.String("certificateChain"),
//   			PrivateKey: jsii.String("privateKey"),
//   		},
//   		Sds: &VirtualGatewayListenerTlsSdsCertificateProperty{
//   			SecretName: jsii.String("secretName"),
//   		},
//   	},
//   	Mode: jsii.String("mode"),
//
//   	// the properties below are optional
//   	Validation: &VirtualGatewayListenerTlsValidationContextProperty{
//   		Trust: &VirtualGatewayListenerTlsValidationContextTrustProperty{
//   			File: &VirtualGatewayTlsValidationContextFileTrustProperty{
//   				CertificateChain: jsii.String("certificateChain"),
//   			},
//   			Sds: &VirtualGatewayTlsValidationContextSdsTrustProperty{
//   				SecretName: jsii.String("secretName"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		SubjectAlternativeNames: &SubjectAlternativeNamesProperty{
//   			Match: &SubjectAlternativeNameMatchersProperty{
//   				Exact: []*string{
//   					jsii.String("exact"),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnVirtualGateway_VirtualGatewayListenerTlsProperty struct {
	// An object that represents a Transport Layer Security (TLS) certificate.
	Certificate interface{} `field:"required" json:"certificate" yaml:"certificate"`
	// Specify one of the following modes.
	//
	// - ** STRICT – Listener only accepts connections with TLS enabled.
	// - ** PERMISSIVE – Listener accepts connections with or without TLS enabled.
	// - ** DISABLED – Listener only accepts connections without TLS.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// A reference to an object that represents a virtual gateway's listener's Transport Layer Security (TLS) validation context.
	Validation interface{} `field:"optional" json:"validation" yaml:"validation"`
}

