package mixinsawsappmesh


// An object that represents the backends that a virtual node is expected to send outbound traffic to.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   backendProperty := &BackendProperty{
//   	VirtualService: &VirtualServiceBackendProperty{
//   		ClientPolicy: &ClientPolicyProperty{
//   			Tls: &ClientPolicyTlsProperty{
//   				Certificate: &ClientTlsCertificateProperty{
//   					File: &ListenerTlsFileCertificateProperty{
//   						CertificateChain: jsii.String("certificateChain"),
//   						PrivateKey: jsii.String("privateKey"),
//   					},
//   					Sds: &ListenerTlsSdsCertificateProperty{
//   						SecretName: jsii.String("secretName"),
//   					},
//   				},
//   				Enforce: jsii.Boolean(false),
//   				Ports: []interface{}{
//   					jsii.Number(123),
//   				},
//   				Validation: &TlsValidationContextProperty{
//   					SubjectAlternativeNames: &SubjectAlternativeNamesProperty{
//   						Match: &SubjectAlternativeNameMatchersProperty{
//   							Exact: []*string{
//   								jsii.String("exact"),
//   							},
//   						},
//   					},
//   					Trust: &TlsValidationContextTrustProperty{
//   						Acm: &TlsValidationContextAcmTrustProperty{
//   							CertificateAuthorityArns: []*string{
//   								jsii.String("certificateAuthorityArns"),
//   							},
//   						},
//   						File: &TlsValidationContextFileTrustProperty{
//   							CertificateChain: jsii.String("certificateChain"),
//   						},
//   						Sds: &TlsValidationContextSdsTrustProperty{
//   							SecretName: jsii.String("secretName"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		VirtualServiceName: jsii.String("virtualServiceName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-backend.html
//
type CfnVirtualNodePropsMixin_BackendProperty struct {
	// Specifies a virtual service to use as a backend.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-backend.html#cfn-appmesh-virtualnode-backend-virtualservice
	//
	VirtualService interface{} `field:"optional" json:"virtualService" yaml:"virtualService"`
}

