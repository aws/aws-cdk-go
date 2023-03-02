package awsappmesh


// An object that represents a client policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clientPolicyProperty := &clientPolicyProperty{
//   	tls: &clientPolicyTlsProperty{
//   		validation: &tlsValidationContextProperty{
//   			trust: &tlsValidationContextTrustProperty{
//   				acm: &tlsValidationContextAcmTrustProperty{
//   					certificateAuthorityArns: []*string{
//   						jsii.String("certificateAuthorityArns"),
//   					},
//   				},
//   				file: &tlsValidationContextFileTrustProperty{
//   					certificateChain: jsii.String("certificateChain"),
//   				},
//   				sds: &tlsValidationContextSdsTrustProperty{
//   					secretName: jsii.String("secretName"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			subjectAlternativeNames: &subjectAlternativeNamesProperty{
//   				match: &subjectAlternativeNameMatchersProperty{
//   					exact: []*string{
//   						jsii.String("exact"),
//   					},
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		certificate: &clientTlsCertificateProperty{
//   			file: &listenerTlsFileCertificateProperty{
//   				certificateChain: jsii.String("certificateChain"),
//   				privateKey: jsii.String("privateKey"),
//   			},
//   			sds: &listenerTlsSdsCertificateProperty{
//   				secretName: jsii.String("secretName"),
//   			},
//   		},
//   		enforce: jsii.Boolean(false),
//   		ports: []interface{}{
//   			jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnVirtualNode_ClientPolicyProperty struct {
	// A reference to an object that represents a Transport Layer Security (TLS) client policy.
	Tls interface{} `field:"optional" json:"tls" yaml:"tls"`
}

