package awsappmesh


// An object that represents a Transport Layer Security (TLS) client policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   virtualGatewayClientPolicyTlsProperty := &virtualGatewayClientPolicyTlsProperty{
//   	validation: &virtualGatewayTlsValidationContextProperty{
//   		trust: &virtualGatewayTlsValidationContextTrustProperty{
//   			acm: &virtualGatewayTlsValidationContextAcmTrustProperty{
//   				certificateAuthorityArns: []*string{
//   					jsii.String("certificateAuthorityArns"),
//   				},
//   			},
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
//
//   	// the properties below are optional
//   	certificate: &virtualGatewayClientTlsCertificateProperty{
//   		file: &virtualGatewayListenerTlsFileCertificateProperty{
//   			certificateChain: jsii.String("certificateChain"),
//   			privateKey: jsii.String("privateKey"),
//   		},
//   		sds: &virtualGatewayListenerTlsSdsCertificateProperty{
//   			secretName: jsii.String("secretName"),
//   		},
//   	},
//   	enforce: jsii.Boolean(false),
//   	ports: []interface{}{
//   		jsii.Number(123),
//   	},
//   }
//
type CfnVirtualGateway_VirtualGatewayClientPolicyTlsProperty struct {
	// A reference to an object that represents a Transport Layer Security (TLS) validation context.
	Validation interface{} `field:"required" json:"validation" yaml:"validation"`
	// A reference to an object that represents a virtual gateway's client's Transport Layer Security (TLS) certificate.
	Certificate interface{} `field:"optional" json:"certificate" yaml:"certificate"`
	// Whether the policy is enforced.
	//
	// The default is `True` , if a value isn't specified.
	Enforce interface{} `field:"optional" json:"enforce" yaml:"enforce"`
	// One or more ports that the policy is enforced for.
	Ports interface{} `field:"optional" json:"ports" yaml:"ports"`
}

