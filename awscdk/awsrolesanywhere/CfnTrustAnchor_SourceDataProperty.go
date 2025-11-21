package awsrolesanywhere


// A union object representing the data field of the TrustAnchor depending on its type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceDataProperty := &SourceDataProperty{
//   	AcmPcaArn: jsii.String("acmPcaArn"),
//   	X509CertificateData: jsii.String("x509CertificateData"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rolesanywhere-trustanchor-sourcedata.html
//
type CfnTrustAnchor_SourceDataProperty struct {
	// The root certificate of the Private Certificate Authority specified by this ARN is used in trust validation for temporary credential requests.
	//
	// Included for trust anchors of type `AWS_ACM_PCA` .
	//
	// > This field is not supported in your region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rolesanywhere-trustanchor-sourcedata.html#cfn-rolesanywhere-trustanchor-sourcedata-acmpcaarn
	//
	AcmPcaArn *string `field:"optional" json:"acmPcaArn" yaml:"acmPcaArn"`
	// The PEM-encoded data for the certificate anchor.
	//
	// Included for trust anchors of type `CERTIFICATE_BUNDLE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rolesanywhere-trustanchor-sourcedata.html#cfn-rolesanywhere-trustanchor-sourcedata-x509certificatedata
	//
	X509CertificateData *string `field:"optional" json:"x509CertificateData" yaml:"x509CertificateData"`
}

