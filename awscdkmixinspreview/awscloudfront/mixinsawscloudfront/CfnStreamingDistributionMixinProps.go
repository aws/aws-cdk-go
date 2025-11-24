package mixinsawscloudfront

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnStreamingDistributionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnStreamingDistributionMixinProps := &CfnStreamingDistributionMixinProps{
//   	StreamingDistributionConfig: &StreamingDistributionConfigProperty{
//   		Aliases: []*string{
//   			jsii.String("aliases"),
//   		},
//   		Comment: jsii.String("comment"),
//   		Enabled: jsii.Boolean(false),
//   		Logging: &LoggingProperty{
//   			Bucket: jsii.String("bucket"),
//   			Enabled: jsii.Boolean(false),
//   			Prefix: jsii.String("prefix"),
//   		},
//   		PriceClass: jsii.String("priceClass"),
//   		S3Origin: &S3OriginProperty{
//   			DomainName: jsii.String("domainName"),
//   			OriginAccessIdentity: jsii.String("originAccessIdentity"),
//   		},
//   		TrustedSigners: &TrustedSignersProperty{
//   			AwsAccountNumbers: []*string{
//   				jsii.String("awsAccountNumbers"),
//   			},
//   			Enabled: jsii.Boolean(false),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-streamingdistribution.html
//
type CfnStreamingDistributionMixinProps struct {
	// The current configuration information for the RTMP distribution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-streamingdistribution.html#cfn-cloudfront-streamingdistribution-streamingdistributionconfig
	//
	StreamingDistributionConfig interface{} `field:"optional" json:"streamingDistributionConfig" yaml:"streamingDistributionConfig"`
	// A complex type that contains zero or more `Tag` elements.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-streamingdistribution.html#cfn-cloudfront-streamingdistribution-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

