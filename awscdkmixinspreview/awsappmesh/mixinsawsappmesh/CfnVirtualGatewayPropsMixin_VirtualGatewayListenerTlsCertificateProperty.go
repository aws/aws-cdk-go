package mixinsawsappmesh


// An object that represents a listener's Transport Layer Security (TLS) certificate.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   virtualGatewayListenerTlsCertificateProperty := &VirtualGatewayListenerTlsCertificateProperty{
//   	Acm: &VirtualGatewayListenerTlsAcmCertificateProperty{
//   		CertificateArn: jsii.String("certificateArn"),
//   	},
//   	File: &VirtualGatewayListenerTlsFileCertificateProperty{
//   		CertificateChain: jsii.String("certificateChain"),
//   		PrivateKey: jsii.String("privateKey"),
//   	},
//   	Sds: &VirtualGatewayListenerTlsSdsCertificateProperty{
//   		SecretName: jsii.String("secretName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewaylistenertlscertificate.html
//
type CfnVirtualGatewayPropsMixin_VirtualGatewayListenerTlsCertificateProperty struct {
	// A reference to an object that represents an Certificate Manager certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewaylistenertlscertificate.html#cfn-appmesh-virtualgateway-virtualgatewaylistenertlscertificate-acm
	//
	Acm interface{} `field:"optional" json:"acm" yaml:"acm"`
	// A reference to an object that represents a local file certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewaylistenertlscertificate.html#cfn-appmesh-virtualgateway-virtualgatewaylistenertlscertificate-file
	//
	File interface{} `field:"optional" json:"file" yaml:"file"`
	// A reference to an object that represents a virtual gateway's listener's Secret Discovery Service certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewaylistenertlscertificate.html#cfn-appmesh-virtualgateway-virtualgatewaylistenertlscertificate-sds
	//
	Sds interface{} `field:"optional" json:"sds" yaml:"sds"`
}

