package awscloudfront

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnTrustStore`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTrustStoreProps := &CfnTrustStoreProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	CaCertificatesBundleSource: &CaCertificatesBundleSourceProperty{
//   		CaCertificatesBundleS3Location: &CaCertificatesBundleS3LocationProperty{
//   			Bucket: jsii.String("bucket"),
//   			Key: jsii.String("key"),
//   			Region: jsii.String("region"),
//
//   			// the properties below are optional
//   			Version: jsii.String("version"),
//   		},
//   	},
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
type CfnTrustStoreProps struct {
	// The trust store's name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-truststore.html#cfn-cloudfront-truststore-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// A CA certificates bundle source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-truststore.html#cfn-cloudfront-truststore-cacertificatesbundlesource
	//
	CaCertificatesBundleSource interface{} `field:"optional" json:"caCertificatesBundleSource" yaml:"caCertificatesBundleSource"`
	// A complex type that contains zero or more `Tag` elements.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-truststore.html#cfn-cloudfront-truststore-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

