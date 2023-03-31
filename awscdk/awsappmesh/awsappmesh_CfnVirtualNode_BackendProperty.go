package awsappmesh


// An object that represents the backends that a virtual node is expected to send outbound traffic to.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   backendProperty := &backendProperty{
//   	virtualService: &virtualServiceBackendProperty{
//   		virtualServiceName: jsii.String("virtualServiceName"),
//
//   		// the properties below are optional
//   		clientPolicy: &clientPolicyProperty{
//   			tls: &clientPolicyTlsProperty{
//   				validation: &tlsValidationContextProperty{
//   					trust: &tlsValidationContextTrustProperty{
//   						acm: &tlsValidationContextAcmTrustProperty{
//   							certificateAuthorityArns: []*string{
//   								jsii.String("certificateAuthorityArns"),
//   							},
//   						},
//   						file: &tlsValidationContextFileTrustProperty{
//   							certificateChain: jsii.String("certificateChain"),
//   						},
//   						sds: &tlsValidationContextSdsTrustProperty{
//   							secretName: jsii.String("secretName"),
//   						},
//   					},
//
//   					// the properties below are optional
//   					subjectAlternativeNames: &subjectAlternativeNamesProperty{
//   						match: &subjectAlternativeNameMatchersProperty{
//   							exact: []*string{
//   								jsii.String("exact"),
//   							},
//   						},
//   					},
//   				},
//
//   				// the properties below are optional
//   				certificate: &clientTlsCertificateProperty{
//   					file: &listenerTlsFileCertificateProperty{
//   						certificateChain: jsii.String("certificateChain"),
//   						privateKey: jsii.String("privateKey"),
//   					},
//   					sds: &listenerTlsSdsCertificateProperty{
//   						secretName: jsii.String("secretName"),
//   					},
//   				},
//   				enforce: jsii.Boolean(false),
//   				ports: []interface{}{
//   					jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnVirtualNode_BackendProperty struct {
	// Specifies a virtual service to use as a backend.
	VirtualService interface{} `field:"optional" json:"virtualService" yaml:"virtualService"`
}

