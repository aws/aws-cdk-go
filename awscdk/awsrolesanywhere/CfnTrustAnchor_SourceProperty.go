package awsrolesanywhere


// The trust anchor type and its related certificate data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceProperty := &SourceProperty{
//   	SourceData: &SourceDataProperty{
//   		AcmPcaArn: jsii.String("acmPcaArn"),
//   		X509CertificateData: jsii.String("x509CertificateData"),
//   	},
//   	SourceType: jsii.String("sourceType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rolesanywhere-trustanchor-source.html
//
type CfnTrustAnchor_SourceProperty struct {
	// The data field of the trust anchor depending on its type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rolesanywhere-trustanchor-source.html#cfn-rolesanywhere-trustanchor-source-sourcedata
	//
	SourceData interface{} `field:"optional" json:"sourceData" yaml:"sourceData"`
	// The type of the TrustAnchor.
	//
	// > `AWS_ACM_PCA` is not an allowed value in your region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rolesanywhere-trustanchor-source.html#cfn-rolesanywhere-trustanchor-source-sourcetype
	//
	SourceType *string `field:"optional" json:"sourceType" yaml:"sourceType"`
}

