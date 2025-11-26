package previewawsappmeshmixins


// An object that represents the default properties for a backend.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   virtualGatewayBackendDefaultsProperty := &VirtualGatewayBackendDefaultsProperty{
//   	ClientPolicy: &VirtualGatewayClientPolicyProperty{
//   		Tls: &VirtualGatewayClientPolicyTlsProperty{
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
//   			Validation: &VirtualGatewayTlsValidationContextProperty{
//   				SubjectAlternativeNames: &SubjectAlternativeNamesProperty{
//   					Match: &SubjectAlternativeNameMatchersProperty{
//   						Exact: []*string{
//   							jsii.String("exact"),
//   						},
//   					},
//   				},
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
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewaybackenddefaults.html
//
type CfnVirtualGatewayPropsMixin_VirtualGatewayBackendDefaultsProperty struct {
	// A reference to an object that represents a client policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewaybackenddefaults.html#cfn-appmesh-virtualgateway-virtualgatewaybackenddefaults-clientpolicy
	//
	ClientPolicy interface{} `field:"optional" json:"clientPolicy" yaml:"clientPolicy"`
}

