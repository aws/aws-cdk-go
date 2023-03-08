package awsappmesh


// An object that represents a virtual service backend for a virtual node.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   virtualServiceBackendProperty := &VirtualServiceBackendProperty{
//   	VirtualServiceName: jsii.String("virtualServiceName"),
//
//   	// the properties below are optional
//   	ClientPolicy: &ClientPolicyProperty{
//   		Tls: &ClientPolicyTlsProperty{
//   			Validation: &TlsValidationContextProperty{
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
//   		},
//   	},
//   }
//
type CfnVirtualNode_VirtualServiceBackendProperty struct {
	// The name of the virtual service that is acting as a virtual node backend.
	//
	// > App Mesh doesn't validate the existence of those virtual services specified in backends. This is to prevent a cyclic dependency between virtual nodes and virtual services creation. Make sure the virtual service name is correct. The virtual service can be created afterwards if it doesn't already exist.
	VirtualServiceName *string `field:"required" json:"virtualServiceName" yaml:"virtualServiceName"`
	// A reference to an object that represents the client policy for a backend.
	ClientPolicy interface{} `field:"optional" json:"clientPolicy" yaml:"clientPolicy"`
}

