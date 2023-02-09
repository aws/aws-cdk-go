package awsrolesanywhere


// The trust anchor type and its related certificate data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceProperty := &sourceProperty{
//   	sourceData: &sourceDataProperty{
//   		acmPcaArn: jsii.String("acmPcaArn"),
//   		x509CertificateData: jsii.String("x509CertificateData"),
//   	},
//   	sourceType: jsii.String("sourceType"),
//   }
//
type CfnTrustAnchor_SourceProperty struct {
	// The data field of the trust anchor depending on its type.
	SourceData interface{} `field:"optional" json:"sourceData" yaml:"sourceData"`
	// The type of the trust anchor.
	SourceType *string `field:"optional" json:"sourceType" yaml:"sourceType"`
}

