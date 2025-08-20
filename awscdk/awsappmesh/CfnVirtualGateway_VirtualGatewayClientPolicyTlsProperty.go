package awsappmesh


// An object that represents a Transport Layer Security (TLS) client policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   virtualGatewayClientPolicyTlsProperty := &VirtualGatewayClientPolicyTlsProperty{
//   	Validation: &VirtualGatewayTlsValidationContextProperty{
//   		Trust: &VirtualGatewayTlsValidationContextTrustProperty{
//   			Acm: &VirtualGatewayTlsValidationContextAcmTrustProperty{
//   				CertificateAuthorityArns: []*string{
//   					jsii.String("certificateAuthorityArns"),
//   				},
//   			},
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
//
//   	// the properties below are optional
//   	Certificate: &VirtualGatewayClientTlsCertificateProperty{
//   		File: &VirtualGatewayListenerTlsFileCertificateProperty{
//   			CertificateChain: jsii.String("certificateChain"),
//   			PrivateKey: jsii.String("privateKey"),
//   		},
//   		Sds: &VirtualGatewayListenerTlsSdsCertificateProperty{
//   			SecretName: jsii.String("secretName"),
//   		},
//   	},
//   	Enforce: jsii.Boolean(false),
//   	Ports: []interface{}{
//   		jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewayclientpolicytls.html
//
type CfnVirtualGateway_VirtualGatewayClientPolicyTlsProperty struct {
	// A reference to an object that represents a Transport Layer Security (TLS) validation context.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewayclientpolicytls.html#cfn-appmesh-virtualgateway-virtualgatewayclientpolicytls-validation
	//
	Validation interface{} `field:"required" json:"validation" yaml:"validation"`
	// A reference to an object that represents a virtual gateway's client's Transport Layer Security (TLS) certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewayclientpolicytls.html#cfn-appmesh-virtualgateway-virtualgatewayclientpolicytls-certificate
	//
	Certificate interface{} `field:"optional" json:"certificate" yaml:"certificate"`
	// Whether the policy is enforced.
	//
	// The default is `True` , if a value isn't specified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewayclientpolicytls.html#cfn-appmesh-virtualgateway-virtualgatewayclientpolicytls-enforce
	//
	Enforce interface{} `field:"optional" json:"enforce" yaml:"enforce"`
	// One or more ports that the policy is enforced for.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewayclientpolicytls.html#cfn-appmesh-virtualgateway-virtualgatewayclientpolicytls-ports
	//
	Ports interface{} `field:"optional" json:"ports" yaml:"ports"`
}

