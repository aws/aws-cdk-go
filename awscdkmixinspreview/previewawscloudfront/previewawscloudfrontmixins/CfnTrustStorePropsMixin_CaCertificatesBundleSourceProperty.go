package previewawscloudfrontmixins


// A CA certificates bundle source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   caCertificatesBundleSourceProperty := &CaCertificatesBundleSourceProperty{
//   	CaCertificatesBundleS3Location: &CaCertificatesBundleS3LocationProperty{
//   		Bucket: jsii.String("bucket"),
//   		Key: jsii.String("key"),
//   		Region: jsii.String("region"),
//   		Version: jsii.String("version"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-truststore-cacertificatesbundlesource.html
//
type CfnTrustStorePropsMixin_CaCertificatesBundleSourceProperty struct {
	// The CA certificates bundle location in Amazon S3.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-truststore-cacertificatesbundlesource.html#cfn-cloudfront-truststore-cacertificatesbundlesource-cacertificatesbundles3location
	//
	CaCertificatesBundleS3Location interface{} `field:"optional" json:"caCertificatesBundleS3Location" yaml:"caCertificatesBundleS3Location"`
}

