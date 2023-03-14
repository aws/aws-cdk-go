package awsappmesh


// An object that represents the Transport Layer Security (TLS) properties for a listener.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   listenerTlsProperty := &listenerTlsProperty{
//   	certificate: &listenerTlsCertificateProperty{
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
//   	mode: jsii.String("mode"),
//
//   	// the properties below are optional
//   	validation: &listenerTlsValidationContextProperty{
//   		trust: &listenerTlsValidationContextTrustProperty{
//   			file: &tlsValidationContextFileTrustProperty{
//   				certificateChain: jsii.String("certificateChain"),
//   			},
//   			sds: &tlsValidationContextSdsTrustProperty{
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
type CfnVirtualNode_ListenerTlsProperty struct {
	// A reference to an object that represents a listener's Transport Layer Security (TLS) certificate.
	Certificate interface{} `field:"required" json:"certificate" yaml:"certificate"`
	// Specify one of the following modes.
	//
	// - ** STRICT – Listener only accepts connections with TLS enabled.
	// - ** PERMISSIVE – Listener accepts connections with or without TLS enabled.
	// - ** DISABLED – Listener only accepts connections without TLS.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// A reference to an object that represents a listener's Transport Layer Security (TLS) validation context.
	Validation interface{} `field:"optional" json:"validation" yaml:"validation"`
}

