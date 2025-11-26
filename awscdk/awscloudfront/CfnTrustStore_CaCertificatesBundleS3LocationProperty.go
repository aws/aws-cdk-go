package awscloudfront


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   caCertificatesBundleS3LocationProperty := &CaCertificatesBundleS3LocationProperty{
//   	Bucket: jsii.String("bucket"),
//   	Key: jsii.String("key"),
//   	Region: jsii.String("region"),
//
//   	// the properties below are optional
//   	Version: jsii.String("version"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-truststore-cacertificatesbundles3location.html
//
type CfnTrustStore_CaCertificatesBundleS3LocationProperty struct {
	// The S3 bucket containing the CA certificates bundle PEM file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-truststore-cacertificatesbundles3location.html#cfn-cloudfront-truststore-cacertificatesbundles3location-bucket
	//
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The S3 object key of the CA certificates bundle PEM file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-truststore-cacertificatesbundles3location.html#cfn-cloudfront-truststore-cacertificatesbundles3location-key
	//
	Key *string `field:"required" json:"key" yaml:"key"`
	// The S3 bucket region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-truststore-cacertificatesbundles3location.html#cfn-cloudfront-truststore-cacertificatesbundles3location-region
	//
	Region *string `field:"required" json:"region" yaml:"region"`
	// The S3 object version of the CA certificates bundle PEM file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-truststore-cacertificatesbundles3location.html#cfn-cloudfront-truststore-cacertificatesbundles3location-version
	//
	Version *string `field:"optional" json:"version" yaml:"version"`
}

