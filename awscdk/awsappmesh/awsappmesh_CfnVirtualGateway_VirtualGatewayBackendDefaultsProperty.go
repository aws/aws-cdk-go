package awsappmesh


// An object that represents the default properties for a backend.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   virtualGatewayBackendDefaultsProperty := &VirtualGatewayBackendDefaultsProperty{
//   	ClientPolicy: &VirtualGatewayClientPolicyProperty{
//   		Tls: &VirtualGatewayClientPolicyTlsProperty{
//   			Validation: &VirtualGatewayTlsValidationContextProperty{
//   				Trust: &VirtualGatewayTlsValidationContextTrustProperty{
//   					Acm: &VirtualGatewayTlsValidationContextAcmTrustProperty{
//   						CertificateAuthorityArns: []*string{
//   							jsii.String("certificateAuthorityArns"),
//   						},
//   					},
//   					File: &VirtualGatewayTlsValidationContextFileTrustProperty{
//   						CertificateChain: jsii.String("certificateChain"),
//   					},
//   					Sds: &VirtualGatewayTlsValidationContextSdsTrustProperty{
//   						SecretName: jsii.String("secretName"),
//   					},
//   				},
//
//   				// the properties below are optional
//   				SubjectAlternativeNames: &SubjectAlternativeNamesProperty{
//   					Match: &SubjectAlternativeNameMatchersProperty{
//   						Exact: []*string{
//   							jsii.String("exact"),
//   						},
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			Certificate: &VirtualGatewayClientTlsCertificateProperty{
//   				File: &VirtualGatewayListenerTlsFileCertificateProperty{
//   					CertificateChain: jsii.String("certificateChain"),
//   					PrivateKey: jsii.String("privateKey"),
//   				},
//   				Sds: &VirtualGatewayListenerTlsSdsCertificateProperty{
//   					SecretName: jsii.String("secretName"),
//   				},
//   			},
//   			Enforce: jsii.Boolean(false),
//   			Ports: []interface{}{
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

