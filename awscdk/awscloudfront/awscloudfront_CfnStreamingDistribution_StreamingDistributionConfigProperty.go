package awscloudfront


// The RTMP distribution's configuration information.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   streamingDistributionConfigProperty := &StreamingDistributionConfigProperty{
//   	Comment: jsii.String("comment"),
//   	Enabled: jsii.Boolean(false),
//   	S3Origin: &S3OriginProperty{
//   		DomainName: jsii.String("domainName"),
//   		OriginAccessIdentity: jsii.String("originAccessIdentity"),
//   	},
//   	TrustedSigners: &TrustedSignersProperty{
//   		Enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		AwsAccountNumbers: []*string{
//   			jsii.String("awsAccountNumbers"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	Aliases: []*string{
//   		jsii.String("aliases"),
//   	},
//   	Logging: &LoggingProperty{
//   		Bucket: jsii.String("bucket"),
//   		Enabled: jsii.Boolean(false),
//   		Prefix: jsii.String("prefix"),
//   	},
//   	PriceClass: jsii.String("priceClass"),
//   }
//
type CfnStreamingDistribution_StreamingDistributionConfigProperty struct {
	// Any comments you want to include about the streaming distribution.
	Comment *string `field:"required" json:"comment" yaml:"comment"`
	// Whether the streaming distribution is enabled to accept user requests for content.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// A complex type that contains information about the Amazon S3 bucket from which you want CloudFront to get your media files for distribution.
	S3Origin interface{} `field:"required" json:"s3Origin" yaml:"s3Origin"`
	// A complex type that specifies any AWS accounts that you want to permit to create signed URLs for private content.
	//
	// If you want the distribution to use signed URLs, include this element; if you want the distribution to use public URLs, remove this element. For more information, see [Serving Private Content through CloudFront](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html) in the *Amazon CloudFront Developer Guide* .
	TrustedSigners interface{} `field:"required" json:"trustedSigners" yaml:"trustedSigners"`
	// A complex type that contains information about CNAMEs (alternate domain names), if any, for this streaming distribution.
	Aliases *[]*string `field:"optional" json:"aliases" yaml:"aliases"`
	// A complex type that controls whether access logs are written for the streaming distribution.
	Logging interface{} `field:"optional" json:"logging" yaml:"logging"`
	// A complex type that contains information about price class for this streaming distribution.
	PriceClass *string `field:"optional" json:"priceClass" yaml:"priceClass"`
}

