package mixinsawsappmesh


// An object that represents the virtual gateway's client's Transport Layer Security (TLS) certificate.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   virtualGatewayClientTlsCertificateProperty := &VirtualGatewayClientTlsCertificateProperty{
//   	File: &VirtualGatewayListenerTlsFileCertificateProperty{
//   		CertificateChain: jsii.String("certificateChain"),
//   		PrivateKey: jsii.String("privateKey"),
//   	},
//   	Sds: &VirtualGatewayListenerTlsSdsCertificateProperty{
//   		SecretName: jsii.String("secretName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewayclienttlscertificate.html
//
type CfnVirtualGatewayPropsMixin_VirtualGatewayClientTlsCertificateProperty struct {
	// An object that represents a local file certificate.
	//
	// The certificate must meet specific requirements and you must have proxy authorization enabled. For more information, see [Transport Layer Security (TLS)](https://docs.aws.amazon.com/app-mesh/latest/userguide/tls.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewayclienttlscertificate.html#cfn-appmesh-virtualgateway-virtualgatewayclienttlscertificate-file
	//
	File interface{} `field:"optional" json:"file" yaml:"file"`
	// A reference to an object that represents a virtual gateway's client's Secret Discovery Service certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewayclienttlscertificate.html#cfn-appmesh-virtualgateway-virtualgatewayclienttlscertificate-sds
	//
	Sds interface{} `field:"optional" json:"sds" yaml:"sds"`
}

