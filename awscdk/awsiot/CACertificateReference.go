package awsiot


// A reference to a CACertificate resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cACertificateReference := &CACertificateReference{
//   	CaCertificateArn: jsii.String("caCertificateArn"),
//   	CaCertificateId: jsii.String("caCertificateId"),
//   }
//
type CACertificateReference struct {
	// The ARN of the CACertificate resource.
	CaCertificateArn *string `field:"required" json:"caCertificateArn" yaml:"caCertificateArn"`
	// The Id of the CACertificate resource.
	CaCertificateId *string `field:"required" json:"caCertificateId" yaml:"caCertificateId"`
}

