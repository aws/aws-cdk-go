package awsappmesh


// An object that represents the default properties for a backend.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   virtualGatewayBackendDefaultsProperty := &virtualGatewayBackendDefaultsProperty{
//   	clientPolicy: &virtualGatewayClientPolicyProperty{
//   		tls: &virtualGatewayClientPolicyTlsProperty{
//   			validation: &virtualGatewayTlsValidationContextProperty{
//   				trust: &virtualGatewayTlsValidationContextTrustProperty{
//   					acm: &virtualGatewayTlsValidationContextAcmTrustProperty{
//   						certificateAuthorityArns: []*string{
//   							jsii.String("certificateAuthorityArns"),
//   						},
//   					},
//   					file: &virtualGatewayTlsValidationContextFileTrustProperty{
//   						certificateChain: jsii.String("certificateChain"),
//   					},
//   					sds: &virtualGatewayTlsValidationContextSdsTrustProperty{
//   						secretName: jsii.String("secretName"),
//   					},
//   				},
//
//   				// the properties below are optional
//   				subjectAlternativeNames: &subjectAlternativeNamesProperty{
//   					match: &subjectAlternativeNameMatchersProperty{
//   						exact: []*string{
//   							jsii.String("exact"),
//   						},
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			certificate: &virtualGatewayClientTlsCertificateProperty{
//   				file: &virtualGatewayListenerTlsFileCertificateProperty{
//   					certificateChain: jsii.String("certificateChain"),
//   					privateKey: jsii.String("privateKey"),
//   				},
//   				sds: &virtualGatewayListenerTlsSdsCertificateProperty{
//   					secretName: jsii.String("secretName"),
//   				},
//   			},
//   			enforce: jsii.Boolean(false),
//   			ports: []interface{}{
//   				jsii.Number(123),
//   			},
//   		},
//   	},
//   }
//
type CfnVirtualGateway_VirtualGatewayBackendDefaultsProperty struct {
	// A reference to an object that represents a client policy.
	ClientPolicy interface{} `field:"optional" json:"clientPolicy" yaml:"clientPolicy"`
}

