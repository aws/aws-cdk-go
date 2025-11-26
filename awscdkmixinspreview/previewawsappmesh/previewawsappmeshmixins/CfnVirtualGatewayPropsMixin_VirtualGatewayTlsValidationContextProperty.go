package previewawsappmeshmixins


// An object that represents a Transport Layer Security (TLS) validation context.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   virtualGatewayTlsValidationContextProperty := &VirtualGatewayTlsValidationContextProperty{
//   	SubjectAlternativeNames: &SubjectAlternativeNamesProperty{
//   		Match: &SubjectAlternativeNameMatchersProperty{
//   			Exact: []*string{
//   				jsii.String("exact"),
//   			},
//   		},
//   	},
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewaytlsvalidationcontext.html
//
type CfnVirtualGatewayPropsMixin_VirtualGatewayTlsValidationContextProperty struct {
	// A reference to an object that represents the SANs for a virtual gateway's listener's Transport Layer Security (TLS) validation context.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewaytlsvalidationcontext.html#cfn-appmesh-virtualgateway-virtualgatewaytlsvalidationcontext-subjectalternativenames
	//
	SubjectAlternativeNames interface{} `field:"optional" json:"subjectAlternativeNames" yaml:"subjectAlternativeNames"`
	// A reference to where to retrieve the trust chain when validating a peer’s Transport Layer Security (TLS) certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewaytlsvalidationcontext.html#cfn-appmesh-virtualgateway-virtualgatewaytlsvalidationcontext-trust
	//
	Trust interface{} `field:"optional" json:"trust" yaml:"trust"`
}

