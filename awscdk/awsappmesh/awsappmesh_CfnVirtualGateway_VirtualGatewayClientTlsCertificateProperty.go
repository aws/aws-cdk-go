package awsappmesh


// An object that represents the virtual gateway's client's Transport Layer Security (TLS) certificate.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   virtualGatewayClientTlsCertificateProperty := &virtualGatewayClientTlsCertificateProperty{
//   	file: &virtualGatewayListenerTlsFileCertificateProperty{
//   		certificateChain: jsii.String("certificateChain"),
//   		privateKey: jsii.String("privateKey"),
//   	},
//   	sds: &virtualGatewayListenerTlsSdsCertificateProperty{
//   		secretName: jsii.String("secretName"),
//   	},
//   }
//
type CfnVirtualGateway_VirtualGatewayClientTlsCertificateProperty struct {
	// An object that represents a local file certificate.
	//
	// The certificate must meet specific requirements and you must have proxy authorization enabled. For more information, see [Transport Layer Security (TLS)](https://docs.aws.amazon.com/app-mesh/latest/userguide/tls.html) .
	File interface{} `field:"optional" json:"file" yaml:"file"`
	// A reference to an object that represents a virtual gateway's client's Secret Discovery Service certificate.
	Sds interface{} `field:"optional" json:"sds" yaml:"sds"`
}

