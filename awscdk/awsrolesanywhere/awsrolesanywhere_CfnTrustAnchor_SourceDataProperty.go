package awsrolesanywhere


// The data field of the trust anchor depending on its type.
//
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
	// The root certificate of the AWS Private Certificate Authority specified by this ARN is used in trust validation for temporary credential requests.
	//
	// Included for trust anchors of type `AWS_ACM_PCA` .
	//
	// > This field is not supported in your region.
	AcmPcaArn *string `field:"optional" json:"acmPcaArn" yaml:"acmPcaArn"`
	// The PEM-encoded data for the certificate anchor.
	//
	// Included for trust anchors of type `CERTIFICATE_BUNDLE` .
	X509CertificateData *string `field:"optional" json:"x509CertificateData" yaml:"x509CertificateData"`
}

