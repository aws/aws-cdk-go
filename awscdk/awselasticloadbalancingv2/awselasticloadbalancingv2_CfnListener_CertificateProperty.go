package awselasticloadbalancingv2


// Specifies an SSL server certificate to use as the default certificate for a secure listener.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   certificateProperty := &certificateProperty{
//   	certificateArn: jsii.String("certificateArn"),
//   }
//
type CfnListener_CertificateProperty struct {
	// The Amazon Resource Name (ARN) of the certificate.
	CertificateArn *string `field:"optional" json:"certificateArn" yaml:"certificateArn"`
}

