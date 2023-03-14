package awsappmesh


// An object that represents a listener's Transport Layer Security (TLS) certificate.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   virtualGatewayListenerTlsCertificateProperty := &virtualGatewayListenerTlsCertificateProperty{
//   	acm: &virtualGatewayListenerTlsAcmCertificateProperty{
//   		certificateArn: jsii.String("certificateArn"),
//   	},
//   	file: &virtualGatewayListenerTlsFileCertificateProperty{
//   		certificateChain: jsii.String("certificateChain"),
//   		privateKey: jsii.String("privateKey"),
//   	},
//   	sds: &virtualGatewayListenerTlsSdsCertificateProperty{
//   		secretName: jsii.String("secretName"),
//   	},
//   }
//
type CfnVirtualGateway_VirtualGatewayListenerTlsCertificateProperty struct {
	// A reference to an object that represents an AWS Certificate Manager certificate.
	Acm interface{} `field:"optional" json:"acm" yaml:"acm"`
	// A reference to an object that represents a local file certificate.
	File interface{} `field:"optional" json:"file" yaml:"file"`
	// A reference to an object that represents a virtual gateway's listener's Secret Discovery Service certificate.
	Sds interface{} `field:"optional" json:"sds" yaml:"sds"`
}

