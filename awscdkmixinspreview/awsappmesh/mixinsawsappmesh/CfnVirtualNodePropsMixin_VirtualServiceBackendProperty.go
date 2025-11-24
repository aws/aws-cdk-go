package mixinsawsappmesh


// An object that represents a virtual service backend for a virtual node.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   virtualServiceBackendProperty := &VirtualServiceBackendProperty{
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
//   	VirtualServiceName: jsii.String("virtualServiceName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-virtualservicebackend.html
//
type CfnVirtualNodePropsMixin_VirtualServiceBackendProperty struct {
	// A reference to an object that represents the client policy for a backend.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-virtualservicebackend.html#cfn-appmesh-virtualnode-virtualservicebackend-clientpolicy
	//
	ClientPolicy interface{} `field:"optional" json:"clientPolicy" yaml:"clientPolicy"`
	// The name of the virtual service that is acting as a virtual node backend.
	//
	// > App Mesh doesn't validate the existence of those virtual services specified in backends. This is to prevent a cyclic dependency between virtual nodes and virtual services creation. Make sure the virtual service name is correct. The virtual service can be created afterwards if it doesn't already exist.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-virtualservicebackend.html#cfn-appmesh-virtualnode-virtualservicebackend-virtualservicename
	//
	VirtualServiceName *string `field:"optional" json:"virtualServiceName" yaml:"virtualServiceName"`
}

