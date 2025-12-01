package previewawscloudfrontmixins


// The CA certificates bundle location in Amazon S3.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   caCertificatesBundleS3LocationProperty := &CaCertificatesBundleS3LocationProperty{
//   	Bucket: jsii.String("bucket"),
//   	Key: jsii.String("key"),
//   	Region: jsii.String("region"),
//   	Version: jsii.String("version"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-truststore-cacertificatesbundles3location.html
//
type CfnTrustStorePropsMixin_CaCertificatesBundleS3LocationProperty struct {
	// The S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-truststore-cacertificatesbundles3location.html#cfn-cloudfront-truststore-cacertificatesbundles3location-bucket
	//
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// The location's key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-truststore-cacertificatesbundles3location.html#cfn-cloudfront-truststore-cacertificatesbundles3location-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The location's Region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-truststore-cacertificatesbundles3location.html#cfn-cloudfront-truststore-cacertificatesbundles3location-region
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The location's version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-truststore-cacertificatesbundles3location.html#cfn-cloudfront-truststore-cacertificatesbundles3location-version
	//
	Version *string `field:"optional" json:"version" yaml:"version"`
}

