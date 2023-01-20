package awsappmesh


// An object that represents a virtual service backend for a virtual node.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   virtualServiceBackendProperty := &virtualServiceBackendProperty{
//   	virtualServiceName: jsii.String("virtualServiceName"),
//
//   	// the properties below are optional
//   	clientPolicy: &clientPolicyProperty{
//   		tls: &clientPolicyTlsProperty{
//   			validation: &tlsValidationContextProperty{
//   				trust: &tlsValidationContextTrustProperty{
//   					acm: &tlsValidationContextAcmTrustProperty{
//   						certificateAuthorityArns: []*string{
//   							jsii.String("certificateAuthorityArns"),
//   						},
//   					},
//   					file: &tlsValidationContextFileTrustProperty{
//   						certificateChain: jsii.String("certificateChain"),
//   					},
//   					sds: &tlsValidationContextSdsTrustProperty{
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
//   			certificate: &clientTlsCertificateProperty{
//   				file: &listenerTlsFileCertificateProperty{
//   					certificateChain: jsii.String("certificateChain"),
//   					privateKey: jsii.String("privateKey"),
//   				},
//   				sds: &listenerTlsSdsCertificateProperty{
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
type CfnVirtualNode_VirtualServiceBackendProperty struct {
	// The name of the virtual service that is acting as a virtual node backend.
	VirtualServiceName *string `field:"required" json:"virtualServiceName" yaml:"virtualServiceName"`
	// A reference to an object that represents the client policy for a backend.
	ClientPolicy interface{} `field:"optional" json:"clientPolicy" yaml:"clientPolicy"`
}

