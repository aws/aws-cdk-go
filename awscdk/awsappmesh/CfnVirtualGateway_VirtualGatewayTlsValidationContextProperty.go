package awsappmesh


// An object that represents a Transport Layer Security (TLS) validation context.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   virtualGatewayTlsValidationContextProperty := &VirtualGatewayTlsValidationContextProperty{
//   	Trust: &VirtualGatewayTlsValidationContextTrustProperty{
//   		Acm: &VirtualGatewayTlsValidationContextAcmTrustProperty{
//   			CertificateAuthorityArns: []*string{
//   				jsii.String("certificateAuthorityArns"),
//   			},
//   		},
//   		File: &VirtualGatewayTlsValidationContextFileTrustProperty{
//   			CertificateChain: jsii.String("certificateChain"),
//   		},
//   		Sds: &VirtualGatewayTlsValidationContextSdsTrustProperty{
//   			SecretName: jsii.String("secretName"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	SubjectAlternativeNames: &SubjectAlternativeNamesProperty{
//   		Match: &SubjectAlternativeNameMatchersProperty{
//   			Exact: []*string{
//   				jsii.String("exact"),
//   			},
//   		},
//   	},
//   }
//
type CfnVirtualGateway_VirtualGatewayTlsValidationContextProperty struct {
	// A reference to where to retrieve the trust chain when validating a peer’s Transport Layer Security (TLS) certificate.
	Trust interface{} `field:"required" json:"trust" yaml:"trust"`
	// A reference to an object that represents the SANs for a virtual gateway's listener's Transport Layer Security (TLS) validation context.
	SubjectAlternativeNames interface{} `field:"optional" json:"subjectAlternativeNames" yaml:"subjectAlternativeNames"`
}

