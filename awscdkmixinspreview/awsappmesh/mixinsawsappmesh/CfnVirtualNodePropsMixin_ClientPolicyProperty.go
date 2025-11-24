package mixinsawsappmesh


// An object that represents a client policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   clientPolicyProperty := &ClientPolicyProperty{
//   	Tls: &ClientPolicyTlsProperty{
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
//   		Validation: &TlsValidationContextProperty{
//   			SubjectAlternativeNames: &SubjectAlternativeNamesProperty{
//   				Match: &SubjectAlternativeNameMatchersProperty{
//   					Exact: []*string{
//   						jsii.String("exact"),
//   					},
//   				},
//   			},
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
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-clientpolicy.html
//
type CfnVirtualNodePropsMixin_ClientPolicyProperty struct {
	// A reference to an object that represents a Transport Layer Security (TLS) client policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-clientpolicy.html#cfn-appmesh-virtualnode-clientpolicy-tls
	//
	Tls interface{} `field:"optional" json:"tls" yaml:"tls"`
}

