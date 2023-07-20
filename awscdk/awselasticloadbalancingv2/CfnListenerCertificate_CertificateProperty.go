package awselasticloadbalancingv2


// Specifies an SSL server certificate for the certificate list of a secure listener.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   certificateProperty := &CertificateProperty{
//   	CertificateArn: jsii.String("certificateArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenercertificate-certificate.html
//
type CfnListenerCertificate_CertificateProperty struct {
	// The Amazon Resource Name (ARN) of the certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenercertificate-certificate.html#cfn-elasticloadbalancingv2-listenercertificate-certificate-certificatearn
	//
	CertificateArn *string `field:"optional" json:"certificateArn" yaml:"certificateArn"`
}

