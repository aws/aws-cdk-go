package awsappmesh


// An object that represents how the proxy will validate its peer during Transport Layer Security (TLS) negotiation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tlsValidationContextProperty := &tlsValidationContextProperty{
//   	trust: &tlsValidationContextTrustProperty{
//   		acm: &tlsValidationContextAcmTrustProperty{
//   			certificateAuthorityArns: []*string{
//   				jsii.String("certificateAuthorityArns"),
//   			},
//   		},
//   		file: &tlsValidationContextFileTrustProperty{
//   			certificateChain: jsii.String("certificateChain"),
//   		},
//   		sds: &tlsValidationContextSdsTrustProperty{
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
type CfnVirtualNode_TlsValidationContextProperty struct {
	// A reference to where to retrieve the trust chain when validating a peer’s Transport Layer Security (TLS) certificate.
	Trust interface{} `field:"required" json:"trust" yaml:"trust"`
	// A reference to an object that represents the SANs for a Transport Layer Security (TLS) validation context.
	SubjectAlternativeNames interface{} `field:"optional" json:"subjectAlternativeNames" yaml:"subjectAlternativeNames"`
}

