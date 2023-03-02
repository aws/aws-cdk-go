package awsappmesh


// An object that represents the client's certificate.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clientTlsCertificateProperty := &clientTlsCertificateProperty{
//   	file: &listenerTlsFileCertificateProperty{
//   		certificateChain: jsii.String("certificateChain"),
//   		privateKey: jsii.String("privateKey"),
//   	},
//   	sds: &listenerTlsSdsCertificateProperty{
//   		secretName: jsii.String("secretName"),
//   	},
//   }
//
type CfnVirtualNode_ClientTlsCertificateProperty struct {
	// An object that represents a local file certificate.
	//
	// The certificate must meet specific requirements and you must have proxy authorization enabled. For more information, see [Transport Layer Security (TLS)](https://docs.aws.amazon.com/app-mesh/latest/userguide/tls.html) .
	File interface{} `field:"optional" json:"file" yaml:"file"`
	// A reference to an object that represents a client's TLS Secret Discovery Service certificate.
	Sds interface{} `field:"optional" json:"sds" yaml:"sds"`
}

