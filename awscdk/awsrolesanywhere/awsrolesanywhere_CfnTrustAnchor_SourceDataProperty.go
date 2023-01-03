package awsrolesanywhere


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceDataProperty := &sourceDataProperty{
//   	acmPcaArn: jsii.String("acmPcaArn"),
//   	x509CertificateData: jsii.String("x509CertificateData"),
//   }
//
type CfnTrustAnchor_SourceDataProperty struct {
	// `CfnTrustAnchor.SourceDataProperty.AcmPcaArn`.
	AcmPcaArn *string `field:"optional" json:"acmPcaArn" yaml:"acmPcaArn"`
	// `CfnTrustAnchor.SourceDataProperty.X509CertificateData`.
	X509CertificateData *string `field:"optional" json:"x509CertificateData" yaml:"x509CertificateData"`
}

