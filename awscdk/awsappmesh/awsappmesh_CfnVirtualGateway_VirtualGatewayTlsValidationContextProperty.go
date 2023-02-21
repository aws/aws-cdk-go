package awsappmesh


// An object that represents a Transport Layer Security (TLS) validation context.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   virtualGatewayTlsValidationContextProperty := &virtualGatewayTlsValidationContextProperty{
//   	trust: &virtualGatewayTlsValidationContextTrustProperty{
//   		acm: &virtualGatewayTlsValidationContextAcmTrustProperty{
//   			certificateAuthorityArns: []*string{
//   				jsii.String("certificateAuthorityArns"),
//   			},
//   		},
//   		file: &virtualGatewayTlsValidationContextFileTrustProperty{
//   			certificateChain: jsii.String("certificateChain"),
//   		},
//   		sds: &virtualGatewayTlsValidationContextSdsTrustProperty{
//   			secretName: jsii.String("secretName"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	subjectAlternativeNames: &subjectAlternativeNamesProperty{
//   		match: &subjectAlternativeNameMatchersProperty{
//   			exact: []*string{
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

