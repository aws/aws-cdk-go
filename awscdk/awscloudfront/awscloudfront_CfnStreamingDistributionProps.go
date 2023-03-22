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
//   cfnStreamingDistributionProps := &cfnStreamingDistributionProps{
//   	streamingDistributionConfig: &streamingDistributionConfigProperty{
//   		comment: jsii.String("comment"),
//   		enabled: jsii.Boolean(false),
//   		s3Origin: &s3OriginProperty{
//   			domainName: jsii.String("domainName"),
//   			originAccessIdentity: jsii.String("originAccessIdentity"),
//   		},
//   		trustedSigners: &trustedSignersProperty{
//   			enabled: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			awsAccountNumbers: []*string{
//   				jsii.String("awsAccountNumbers"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		aliases: []*string{
//   			jsii.String("aliases"),
//   		},
//   		logging: &loggingProperty{
//   			bucket: jsii.String("bucket"),
//   			enabled: jsii.Boolean(false),
//   			prefix: jsii.String("prefix"),
//   		},
//   		priceClass: jsii.String("priceClass"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
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

