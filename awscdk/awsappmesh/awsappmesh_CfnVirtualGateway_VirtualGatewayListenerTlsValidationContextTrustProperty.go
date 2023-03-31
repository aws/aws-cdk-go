package awsappmesh


// An object that represents a virtual gateway's listener's Transport Layer Security (TLS) validation context trust.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   virtualGatewayListenerTlsValidationContextTrustProperty := &virtualGatewayListenerTlsValidationContextTrustProperty{
//   	file: &virtualGatewayTlsValidationContextFileTrustProperty{
//   		certificateChain: jsii.String("certificateChain"),
//   	},
//   	sds: &virtualGatewayTlsValidationContextSdsTrustProperty{
//   		secretName: jsii.String("secretName"),
//   	},
//   }
//
type CfnVirtualGateway_VirtualGatewayListenerTlsValidationContextTrustProperty struct {
	// An object that represents a Transport Layer Security (TLS) validation context trust for a local file.
	File interface{} `field:"optional" json:"file" yaml:"file"`
	// A reference to an object that represents a virtual gateway's listener's Transport Layer Security (TLS) Secret Discovery Service validation context trust.
	Sds interface{} `field:"optional" json:"sds" yaml:"sds"`
}

