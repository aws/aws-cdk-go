package awsappmesh


// An object that represents the Transport Layer Security (TLS) properties for a listener.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   virtualGatewayListenerTlsProperty := &virtualGatewayListenerTlsProperty{
//   	certificate: &virtualGatewayListenerTlsCertificateProperty{
//   		acm: &virtualGatewayListenerTlsAcmCertificateProperty{
//   			certificateArn: jsii.String("certificateArn"),
//   		},
//   		file: &virtualGatewayListenerTlsFileCertificateProperty{
//   			certificateChain: jsii.String("certificateChain"),
//   			privateKey: jsii.String("privateKey"),
//   		},
//   		sds: &virtualGatewayListenerTlsSdsCertificateProperty{
//   			secretName: jsii.String("secretName"),
//   		},
//   	},
//   	mode: jsii.String("mode"),
//
//   	// the properties below are optional
//   	validation: &virtualGatewayListenerTlsValidationContextProperty{
//   		trust: &virtualGatewayListenerTlsValidationContextTrustProperty{
//   			file: &virtualGatewayTlsValidationContextFileTrustProperty{
//   				certificateChain: jsii.String("certificateChain"),
//   			},
//   			sds: &virtualGatewayTlsValidationContextSdsTrustProperty{
//   				secretName: jsii.String("secretName"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		subjectAlternativeNames: &subjectAlternativeNamesProperty{
//   			match: &subjectAlternativeNameMatchersProperty{
//   				exact: []*string{
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

