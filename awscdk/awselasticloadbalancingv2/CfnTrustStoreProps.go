package awselasticloadbalancingv2

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
//   	CaCertificatesBundleS3Bucket: jsii.String("caCertificatesBundleS3Bucket"),
//   	CaCertificatesBundleS3Key: jsii.String("caCertificatesBundleS3Key"),
//   	CaCertificatesBundleS3ObjectVersion: jsii.String("caCertificatesBundleS3ObjectVersion"),
//   	Name: jsii.String("name"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-truststore.html
//
type CfnTrustStoreProps struct {
	// The Amazon S3 bucket for the ca certificates bundle.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-truststore.html#cfn-elasticloadbalancingv2-truststore-cacertificatesbundles3bucket
	//
	CaCertificatesBundleS3Bucket *string `field:"optional" json:"caCertificatesBundleS3Bucket" yaml:"caCertificatesBundleS3Bucket"`
	// The Amazon S3 path for the ca certificates bundle.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-truststore.html#cfn-elasticloadbalancingv2-truststore-cacertificatesbundles3key
	//
	CaCertificatesBundleS3Key *string `field:"optional" json:"caCertificatesBundleS3Key" yaml:"caCertificatesBundleS3Key"`
	// The Amazon S3 object version for the ca certificates bundle.
	//
	// If undefined the current version is used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-truststore.html#cfn-elasticloadbalancingv2-truststore-cacertificatesbundles3objectversion
	//
	CaCertificatesBundleS3ObjectVersion *string `field:"optional" json:"caCertificatesBundleS3ObjectVersion" yaml:"caCertificatesBundleS3ObjectVersion"`
	// The name of the trust store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-truststore.html#cfn-elasticloadbalancingv2-truststore-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The tags to assign to the trust store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-truststore.html#cfn-elasticloadbalancingv2-truststore-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

