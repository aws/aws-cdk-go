package previewawscloudfrontmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   viewerMtlsConfigProperty := &ViewerMtlsConfigProperty{
//   	Mode: jsii.String("mode"),
//   	TrustStoreConfig: &TrustStoreConfigProperty{
//   		AdvertiseTrustStoreCaNames: jsii.Boolean(false),
//   		IgnoreCertificateExpiry: jsii.Boolean(false),
//   		TrustStoreId: jsii.String("trustStoreId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-viewermtlsconfig.html
//
type CfnDistributionPropsMixin_ViewerMtlsConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-viewermtlsconfig.html#cfn-cloudfront-distribution-viewermtlsconfig-mode
	//
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-viewermtlsconfig.html#cfn-cloudfront-distribution-viewermtlsconfig-truststoreconfig
	//
	TrustStoreConfig interface{} `field:"optional" json:"trustStoreConfig" yaml:"trustStoreConfig"`
}

