package awsiot


// An object that contains information about a server certificate.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serverCertificateSummaryProperty := &serverCertificateSummaryProperty{
//   	serverCertificateArn: jsii.String("serverCertificateArn"),
//   	serverCertificateStatus: jsii.String("serverCertificateStatus"),
//   	serverCertificateStatusDetail: jsii.String("serverCertificateStatusDetail"),
//   }
//
type CfnDomainConfiguration_ServerCertificateSummaryProperty struct {
	// The ARN of the server certificate.
	ServerCertificateArn *string `field:"optional" json:"serverCertificateArn" yaml:"serverCertificateArn"`
	// The status of the server certificate.
	ServerCertificateStatus *string `field:"optional" json:"serverCertificateStatus" yaml:"serverCertificateStatus"`
	// Details that explain the status of the server certificate.
	ServerCertificateStatusDetail *string `field:"optional" json:"serverCertificateStatusDetail" yaml:"serverCertificateStatusDetail"`
}

