package awsappmesh


// An object that represents a listener's Transport Layer Security (TLS) validation context.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   listenerTlsValidationContextProperty := &listenerTlsValidationContextProperty{
//   	trust: &listenerTlsValidationContextTrustProperty{
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
type CfnVirtualNode_ListenerTlsValidationContextProperty struct {
	// A reference to where to retrieve the trust chain when validating a peer’s Transport Layer Security (TLS) certificate.
	Trust interface{} `field:"required" json:"trust" yaml:"trust"`
	// A reference to an object that represents the SANs for a listener's Transport Layer Security (TLS) validation context.
	SubjectAlternativeNames interface{} `field:"optional" json:"subjectAlternativeNames" yaml:"subjectAlternativeNames"`
}

