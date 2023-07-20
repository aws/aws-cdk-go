package awsappmesh


// An object that represents a Transport Layer Security (TLS) validation context trust.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   virtualGatewayTlsValidationContextTrustProperty := &VirtualGatewayTlsValidationContextTrustProperty{
//   	Acm: &VirtualGatewayTlsValidationContextAcmTrustProperty{
//   		CertificateAuthorityArns: []*string{
//   			jsii.String("certificateAuthorityArns"),
//   		},
//   	},
//   	File: &VirtualGatewayTlsValidationContextFileTrustProperty{
//   		CertificateChain: jsii.String("certificateChain"),
//   	},
//   	Sds: &VirtualGatewayTlsValidationContextSdsTrustProperty{
//   		SecretName: jsii.String("secretName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewaytlsvalidationcontexttrust.html
//
type CfnVirtualGateway_VirtualGatewayTlsValidationContextTrustProperty struct {
	// A reference to an object that represents a Transport Layer Security (TLS) validation context trust for an AWS Certificate Manager certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewaytlsvalidationcontexttrust.html#cfn-appmesh-virtualgateway-virtualgatewaytlsvalidationcontexttrust-acm
	//
	Acm interface{} `field:"optional" json:"acm" yaml:"acm"`
	// An object that represents a Transport Layer Security (TLS) validation context trust for a local file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewaytlsvalidationcontexttrust.html#cfn-appmesh-virtualgateway-virtualgatewaytlsvalidationcontexttrust-file
	//
	File interface{} `field:"optional" json:"file" yaml:"file"`
	// A reference to an object that represents a virtual gateway's Transport Layer Security (TLS) Secret Discovery Service validation context trust.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewaytlsvalidationcontexttrust.html#cfn-appmesh-virtualgateway-virtualgatewaytlsvalidationcontexttrust-sds
	//
	Sds interface{} `field:"optional" json:"sds" yaml:"sds"`
}

