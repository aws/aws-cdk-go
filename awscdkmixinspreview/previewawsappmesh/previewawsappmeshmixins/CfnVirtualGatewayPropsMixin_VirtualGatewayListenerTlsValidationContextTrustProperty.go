package previewawsappmeshmixins


// An object that represents a virtual gateway's listener's Transport Layer Security (TLS) validation context trust.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   virtualGatewayListenerTlsValidationContextTrustProperty := &VirtualGatewayListenerTlsValidationContextTrustProperty{
//   	File: &VirtualGatewayTlsValidationContextFileTrustProperty{
//   		CertificateChain: jsii.String("certificateChain"),
//   	},
//   	Sds: &VirtualGatewayTlsValidationContextSdsTrustProperty{
//   		SecretName: jsii.String("secretName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewaylistenertlsvalidationcontexttrust.html
//
type CfnVirtualGatewayPropsMixin_VirtualGatewayListenerTlsValidationContextTrustProperty struct {
	// An object that represents a Transport Layer Security (TLS) validation context trust for a local file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewaylistenertlsvalidationcontexttrust.html#cfn-appmesh-virtualgateway-virtualgatewaylistenertlsvalidationcontexttrust-file
	//
	File interface{} `field:"optional" json:"file" yaml:"file"`
	// A reference to an object that represents a virtual gateway's listener's Transport Layer Security (TLS) Secret Discovery Service validation context trust.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewaylistenertlsvalidationcontexttrust.html#cfn-appmesh-virtualgateway-virtualgatewaylistenertlsvalidationcontexttrust-sds
	//
	Sds interface{} `field:"optional" json:"sds" yaml:"sds"`
}

