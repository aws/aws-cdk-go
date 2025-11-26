package previewawsappmeshmixins


// An object that represents the default properties for a backend.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   backendDefaultsProperty := &BackendDefaultsProperty{
//   	ClientPolicy: &ClientPolicyProperty{
//   		Tls: &ClientPolicyTlsProperty{
//   			Certificate: &ClientTlsCertificateProperty{
//   				File: &ListenerTlsFileCertificateProperty{
//   					CertificateChain: jsii.String("certificateChain"),
//   					PrivateKey: jsii.String("privateKey"),
//   				},
//   				Sds: &ListenerTlsSdsCertificateProperty{
//   					SecretName: jsii.String("secretName"),
//   				},
//   			},
//   			Enforce: jsii.Boolean(false),
//   			Ports: []interface{}{
//   				jsii.Number(123),
//   			},
//   			Validation: &TlsValidationContextProperty{
//   				SubjectAlternativeNames: &SubjectAlternativeNamesProperty{
//   					Match: &SubjectAlternativeNameMatchersProperty{
//   						Exact: []*string{
//   							jsii.String("exact"),
//   						},
//   					},
//   				},
//   				Trust: &TlsValidationContextTrustProperty{
//   					Acm: &TlsValidationContextAcmTrustProperty{
//   						CertificateAuthorityArns: []*string{
//   							jsii.String("certificateAuthorityArns"),
//   						},
//   					},
//   					File: &TlsValidationContextFileTrustProperty{
//   						CertificateChain: jsii.String("certificateChain"),
//   					},
//   					Sds: &TlsValidationContextSdsTrustProperty{
//   						SecretName: jsii.String("secretName"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-backenddefaults.html
//
type CfnVirtualNodePropsMixin_BackendDefaultsProperty struct {
	// A reference to an object that represents a client policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-backenddefaults.html#cfn-appmesh-virtualnode-backenddefaults-clientpolicy
	//
	ClientPolicy interface{} `field:"optional" json:"clientPolicy" yaml:"clientPolicy"`
}

