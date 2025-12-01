package previewawscloudfrontmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnTrustStorePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTrustStoreMixinProps := &CfnTrustStoreMixinProps{
//   	CaCertificatesBundleSource: &CaCertificatesBundleSourceProperty{
//   		CaCertificatesBundleS3Location: &CaCertificatesBundleS3LocationProperty{
//   			Bucket: jsii.String("bucket"),
//   			Key: jsii.String("key"),
//   			Region: jsii.String("region"),
//   			Version: jsii.String("version"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-truststore.html
//
type CfnTrustStoreMixinProps struct {
	// A CA certificates bundle source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-truststore.html#cfn-cloudfront-truststore-cacertificatesbundlesource
	//
	CaCertificatesBundleSource interface{} `field:"optional" json:"caCertificatesBundleSource" yaml:"caCertificatesBundleSource"`
	// The trust store's name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-truststore.html#cfn-cloudfront-truststore-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A complex type that contains zero or more `Tag` elements.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-truststore.html#cfn-cloudfront-truststore-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

