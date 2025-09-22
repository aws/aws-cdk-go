package awsiot


// A reference to a CertificateProvider resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   certificateProviderReference := &CertificateProviderReference{
//   	CertificateProviderArn: jsii.String("certificateProviderArn"),
//   	CertificateProviderName: jsii.String("certificateProviderName"),
//   }
//
type CertificateProviderReference struct {
	// The ARN of the CertificateProvider resource.
	CertificateProviderArn *string `field:"required" json:"certificateProviderArn" yaml:"certificateProviderArn"`
	// The CertificateProviderName of the CertificateProvider resource.
	CertificateProviderName *string `field:"required" json:"certificateProviderName" yaml:"certificateProviderName"`
}

