package awsiam


// A reference to a ServerCertificate resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serverCertificateReference := &ServerCertificateReference{
//   	ServerCertificateArn: jsii.String("serverCertificateArn"),
//   	ServerCertificateName: jsii.String("serverCertificateName"),
//   }
//
type ServerCertificateReference struct {
	// The ARN of the ServerCertificate resource.
	ServerCertificateArn *string `field:"required" json:"serverCertificateArn" yaml:"serverCertificateArn"`
	// The ServerCertificateName of the ServerCertificate resource.
	ServerCertificateName *string `field:"required" json:"serverCertificateName" yaml:"serverCertificateName"`
}

