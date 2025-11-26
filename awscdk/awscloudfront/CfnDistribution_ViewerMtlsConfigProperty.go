package awscloudfront


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   viewerMtlsConfigProperty := &ViewerMtlsConfigProperty{
//   	Mode: jsii.String("mode"),
//   	TrustStoreConfig: &TrustStoreConfigProperty{
//   		TrustStoreId: jsii.String("trustStoreId"),
//
//   		// the properties below are optional
//   		AdvertiseTrustStoreCaNames: jsii.Boolean(false),
//   		IgnoreCertificateExpiry: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-viewermtlsconfig.html
//
type CfnDistribution_ViewerMtlsConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-viewermtlsconfig.html#cfn-cloudfront-distribution-viewermtlsconfig-mode
	//
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-viewermtlsconfig.html#cfn-cloudfront-distribution-viewermtlsconfig-truststoreconfig
	//
	TrustStoreConfig interface{} `field:"optional" json:"trustStoreConfig" yaml:"trustStoreConfig"`
}

