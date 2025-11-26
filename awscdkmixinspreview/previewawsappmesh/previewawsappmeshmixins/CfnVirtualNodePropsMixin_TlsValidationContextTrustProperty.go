package previewawsappmeshmixins


// An object that represents a Transport Layer Security (TLS) validation context trust.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   tlsValidationContextTrustProperty := &TlsValidationContextTrustProperty{
//   	Acm: &TlsValidationContextAcmTrustProperty{
//   		CertificateAuthorityArns: []*string{
//   			jsii.String("certificateAuthorityArns"),
//   		},
//   	},
//   	File: &TlsValidationContextFileTrustProperty{
//   		CertificateChain: jsii.String("certificateChain"),
//   	},
//   	Sds: &TlsValidationContextSdsTrustProperty{
//   		SecretName: jsii.String("secretName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-tlsvalidationcontexttrust.html
//
type CfnVirtualNodePropsMixin_TlsValidationContextTrustProperty struct {
	// A reference to an object that represents a Transport Layer Security (TLS) validation context trust for an Certificate Manager certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-tlsvalidationcontexttrust.html#cfn-appmesh-virtualnode-tlsvalidationcontexttrust-acm
	//
	Acm interface{} `field:"optional" json:"acm" yaml:"acm"`
	// An object that represents a Transport Layer Security (TLS) validation context trust for a local file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-tlsvalidationcontexttrust.html#cfn-appmesh-virtualnode-tlsvalidationcontexttrust-file
	//
	File interface{} `field:"optional" json:"file" yaml:"file"`
	// A reference to an object that represents a Transport Layer Security (TLS) Secret Discovery Service validation context trust.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-tlsvalidationcontexttrust.html#cfn-appmesh-virtualnode-tlsvalidationcontexttrust-sds
	//
	Sds interface{} `field:"optional" json:"sds" yaml:"sds"`
}

