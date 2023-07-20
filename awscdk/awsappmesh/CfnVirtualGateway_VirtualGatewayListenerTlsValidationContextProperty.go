package awsappmesh


// An object that represents a virtual gateway's listener's Transport Layer Security (TLS) validation context.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   virtualGatewayListenerTlsValidationContextProperty := &VirtualGatewayListenerTlsValidationContextProperty{
//   	Trust: &VirtualGatewayListenerTlsValidationContextTrustProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewaylistenertlsvalidationcontext.html
//
type CfnVirtualGateway_VirtualGatewayListenerTlsValidationContextProperty struct {
	// A reference to where to retrieve the trust chain when validating a peer’s Transport Layer Security (TLS) certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewaylistenertlsvalidationcontext.html#cfn-appmesh-virtualgateway-virtualgatewaylistenertlsvalidationcontext-trust
	//
	Trust interface{} `field:"required" json:"trust" yaml:"trust"`
	// A reference to an object that represents the SANs for a virtual gateway listener's Transport Layer Security (TLS) validation context.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewaylistenertlsvalidationcontext.html#cfn-appmesh-virtualgateway-virtualgatewaylistenertlsvalidationcontext-subjectalternativenames
	//
	SubjectAlternativeNames interface{} `field:"optional" json:"subjectAlternativeNames" yaml:"subjectAlternativeNames"`
}

