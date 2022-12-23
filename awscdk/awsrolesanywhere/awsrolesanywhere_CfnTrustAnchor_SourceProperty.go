package awsrolesanywhere


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
	// `CfnTrustAnchor.SourceProperty.SourceData`.
	SourceData interface{} `field:"optional" json:"sourceData" yaml:"sourceData"`
	// `CfnTrustAnchor.SourceProperty.SourceType`.
	SourceType *string `field:"optional" json:"sourceType" yaml:"sourceType"`
}

