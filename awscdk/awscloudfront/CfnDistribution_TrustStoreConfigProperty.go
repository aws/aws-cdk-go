package awscloudfront


// A trust store configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   trustStoreConfigProperty := &TrustStoreConfigProperty{
//   	TrustStoreId: jsii.String("trustStoreId"),
//
//   	// the properties below are optional
//   	AdvertiseTrustStoreCaNames: jsii.Boolean(false),
//   	IgnoreCertificateExpiry: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-truststoreconfig.html
//
type CfnDistribution_TrustStoreConfigProperty struct {
	// The trust store ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-truststoreconfig.html#cfn-cloudfront-distribution-truststoreconfig-truststoreid
	//
	TrustStoreId *string `field:"required" json:"trustStoreId" yaml:"trustStoreId"`
	// The configuration to use to advertise trust store CA names.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-truststoreconfig.html#cfn-cloudfront-distribution-truststoreconfig-advertisetruststorecanames
	//
	AdvertiseTrustStoreCaNames interface{} `field:"optional" json:"advertiseTrustStoreCaNames" yaml:"advertiseTrustStoreCaNames"`
	// The configuration to use to ignore certificate expiration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-truststoreconfig.html#cfn-cloudfront-distribution-truststoreconfig-ignorecertificateexpiry
	//
	IgnoreCertificateExpiry interface{} `field:"optional" json:"ignoreCertificateExpiry" yaml:"ignoreCertificateExpiry"`
}

