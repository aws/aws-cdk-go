package awsappmesh


// An object that represents a listener's Transport Layer Security (TLS) validation context.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   listenerTlsValidationContextProperty := &ListenerTlsValidationContextProperty{
//   	Trust: &ListenerTlsValidationContextTrustProperty{
//   		File: &TlsValidationContextFileTrustProperty{
//   			CertificateChain: jsii.String("certificateChain"),
//   		},
//   		Sds: &TlsValidationContextSdsTrustProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-listenertlsvalidationcontext.html
//
type CfnVirtualNode_ListenerTlsValidationContextProperty struct {
	// A reference to where to retrieve the trust chain when validating a peer’s Transport Layer Security (TLS) certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-listenertlsvalidationcontext.html#cfn-appmesh-virtualnode-listenertlsvalidationcontext-trust
	//
	Trust interface{} `field:"required" json:"trust" yaml:"trust"`
	// A reference to an object that represents the SANs for a listener's Transport Layer Security (TLS) validation context.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-listenertlsvalidationcontext.html#cfn-appmesh-virtualnode-listenertlsvalidationcontext-subjectalternativenames
	//
	SubjectAlternativeNames interface{} `field:"optional" json:"subjectAlternativeNames" yaml:"subjectAlternativeNames"`
}

