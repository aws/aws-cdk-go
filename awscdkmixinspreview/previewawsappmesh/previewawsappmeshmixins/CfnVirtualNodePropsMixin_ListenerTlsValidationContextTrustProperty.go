package previewawsappmeshmixins


// An object that represents a listener's Transport Layer Security (TLS) validation context trust.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   listenerTlsValidationContextTrustProperty := &ListenerTlsValidationContextTrustProperty{
//   	File: &TlsValidationContextFileTrustProperty{
//   		CertificateChain: jsii.String("certificateChain"),
//   	},
//   	Sds: &TlsValidationContextSdsTrustProperty{
//   		SecretName: jsii.String("secretName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-listenertlsvalidationcontexttrust.html
//
type CfnVirtualNodePropsMixin_ListenerTlsValidationContextTrustProperty struct {
	// An object that represents a Transport Layer Security (TLS) validation context trust for a local file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-listenertlsvalidationcontexttrust.html#cfn-appmesh-virtualnode-listenertlsvalidationcontexttrust-file
	//
	File interface{} `field:"optional" json:"file" yaml:"file"`
	// A reference to an object that represents a listener's Transport Layer Security (TLS) Secret Discovery Service validation context trust.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-listenertlsvalidationcontexttrust.html#cfn-appmesh-virtualnode-listenertlsvalidationcontexttrust-sds
	//
	Sds interface{} `field:"optional" json:"sds" yaml:"sds"`
}

