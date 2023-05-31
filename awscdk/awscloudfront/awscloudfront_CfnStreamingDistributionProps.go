package awscloudfront

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnStreamingDistribution`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnStreamingDistributionProps := &CfnStreamingDistributionProps{
//   	StreamingDistributionConfig: &StreamingDistributionConfigProperty{
//   		Comment: jsii.String("comment"),
//   		Enabled: jsii.Boolean(false),
//   		S3Origin: &S3OriginProperty{
//   			DomainName: jsii.String("domainName"),
//   			OriginAccessIdentity: jsii.String("originAccessIdentity"),
//   		},
//   		TrustedSigners: &TrustedSignersProperty{
//   			Enabled: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			AwsAccountNumbers: []*string{
//   				jsii.String("awsAccountNumbers"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		Aliases: []*string{
//   			jsii.String("aliases"),
//   		},
//   		Logging: &LoggingProperty{
//   			Bucket: jsii.String("bucket"),
//   			Enabled: jsii.Boolean(false),
//   			Prefix: jsii.String("prefix"),
//   		},
//   		PriceClass: jsii.String("priceClass"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnStreamingDistributionProps struct {
	// The current configuration information for the RTMP distribution.
	StreamingDistributionConfig interface{} `field:"required" json:"streamingDistributionConfig" yaml:"streamingDistributionConfig"`
	// A complex type that contains zero or more `Tag` elements.
	Tags *[]*awscdk.CfnTag `field:"required" json:"tags" yaml:"tags"`
}

