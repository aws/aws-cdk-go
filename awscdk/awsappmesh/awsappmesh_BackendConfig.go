package awsappmesh


// Properties for a backend.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   backendConfig := &BackendConfig{
//   	VirtualServiceBackend: &BackendProperty{
//   		VirtualService: &VirtualServiceBackendProperty{
//   			VirtualServiceName: jsii.String("virtualServiceName"),
//
//   			// the properties below are optional
//   			ClientPolicy: &ClientPolicyProperty{
//   				Tls: &ClientPolicyTlsProperty{
//   					Validation: &TlsValidationContextProperty{
//   						Trust: &TlsValidationContextTrustProperty{
//   							Acm: &TlsValidationContextAcmTrustProperty{
//   								CertificateAuthorityArns: []*string{
//   									jsii.String("certificateAuthorityArns"),
//   								},
//   							},
//   							File: &TlsValidationContextFileTrustProperty{
//   								CertificateChain: jsii.String("certificateChain"),
//   							},
//   							Sds: &TlsValidationContextSdsTrustProperty{
//   								SecretName: jsii.String("secretName"),
//   							},
//   						},
//
//   						// the properties below are optional
//   						SubjectAlternativeNames: &SubjectAlternativeNamesProperty{
//   							Match: &SubjectAlternativeNameMatchersProperty{
//   								Exact: []*string{
//   									jsii.String("exact"),
//   								},
//   							},
//   						},
//   					},
//
//   					// the properties below are optional
//   					Certificate: &ClientTlsCertificateProperty{
//   						File: &ListenerTlsFileCertificateProperty{
//   							CertificateChain: jsii.String("certificateChain"),
//   							PrivateKey: jsii.String("privateKey"),
//   						},
//   						Sds: &ListenerTlsSdsCertificateProperty{
//   							SecretName: jsii.String("secretName"),
//   						},
//   					},
//   					Enforce: jsii.Boolean(false),
//   					Ports: []interface{}{
//   						jsii.Number(123),
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
// Experimental.
type BackendConfig struct {
	// Config for a Virtual Service backend.
	// Experimental.
	VirtualServiceBackend *CfnVirtualNode_BackendProperty `field:"required" json:"virtualServiceBackend" yaml:"virtualServiceBackend"`
}

