package awsappmesh


// An object that represents a client policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clientPolicyProperty := &ClientPolicyProperty{
//   	Tls: &ClientPolicyTlsProperty{
//   		Validation: &TlsValidationContextProperty{
//   			Trust: &TlsValidationContextTrustProperty{
//   				Acm: &TlsValidationContextAcmTrustProperty{
//   					CertificateAuthorityArns: []*string{
//   						jsii.String("certificateAuthorityArns"),
//   					},
//   				},
//   				File: &TlsValidationContextFileTrustProperty{
//   					CertificateChain: jsii.String("certificateChain"),
//   				},
//   				Sds: &TlsValidationContextSdsTrustProperty{
//   					SecretName: jsii.String("secretName"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			SubjectAlternativeNames: &SubjectAlternativeNamesProperty{
//   				Match: &SubjectAlternativeNameMatchersProperty{
//   					Exact: []*string{
//   						jsii.String("exact"),
//   					},
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		Certificate: &ClientTlsCertificateProperty{
//   			File: &ListenerTlsFileCertificateProperty{
//   				CertificateChain: jsii.String("certificateChain"),
//   				PrivateKey: jsii.String("privateKey"),
//   			},
//   			Sds: &ListenerTlsSdsCertificateProperty{
//   				SecretName: jsii.String("secretName"),
//   			},
//   		},
//   		Enforce: jsii.Boolean(false),
//   		Ports: []interface{}{
//   			jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnVirtualNode_ClientPolicyProperty struct {
	// A reference to an object that represents a Transport Layer Security (TLS) client policy.
	Tls interface{} `field:"optional" json:"tls" yaml:"tls"`
}

