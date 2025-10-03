package awsappmesh


// An object that represents a listener's Transport Layer Security (TLS) certificate.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   listenerTlsCertificateProperty := &ListenerTlsCertificateProperty{
//   	Acm: &ListenerTlsAcmCertificateProperty{
//   		CertificateArn: jsii.String("certificateArn"),
//   	},
//   	File: &ListenerTlsFileCertificateProperty{
//   		CertificateChain: jsii.String("certificateChain"),
//   		PrivateKey: jsii.String("privateKey"),
//   	},
//   	Sds: &ListenerTlsSdsCertificateProperty{
//   		SecretName: jsii.String("secretName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-listenertlscertificate.html
//
type CfnVirtualNode_ListenerTlsCertificateProperty struct {
	// A reference to an object that represents an AWS Certificate Manager certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-listenertlscertificate.html#cfn-appmesh-virtualnode-listenertlscertificate-acm
	//
	Acm interface{} `field:"optional" json:"acm" yaml:"acm"`
	// A reference to an object that represents a local file certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-listenertlscertificate.html#cfn-appmesh-virtualnode-listenertlscertificate-file
	//
	File interface{} `field:"optional" json:"file" yaml:"file"`
	// A reference to an object that represents a listener's Secret Discovery Service certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-listenertlscertificate.html#cfn-appmesh-virtualnode-listenertlscertificate-sds
	//
	Sds interface{} `field:"optional" json:"sds" yaml:"sds"`
}

