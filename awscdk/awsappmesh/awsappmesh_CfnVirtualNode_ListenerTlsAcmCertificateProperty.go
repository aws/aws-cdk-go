package awsappmesh


// An object that represents an AWS Certificate Manager certificate.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   listenerTlsAcmCertificateProperty := &listenerTlsAcmCertificateProperty{
//   	certificateArn: jsii.String("certificateArn"),
//   }
//
type CfnVirtualNode_ListenerTlsAcmCertificateProperty struct {
	// The Amazon Resource Name (ARN) for the certificate.
	//
	// The certificate must meet specific requirements and you must have proxy authorization enabled. For more information, see [Transport Layer Security (TLS)](https://docs.aws.amazon.com/app-mesh/latest/userguide/tls.html#virtual-node-tls-prerequisites) .
	CertificateArn *string `field:"required" json:"certificateArn" yaml:"certificateArn"`
}

