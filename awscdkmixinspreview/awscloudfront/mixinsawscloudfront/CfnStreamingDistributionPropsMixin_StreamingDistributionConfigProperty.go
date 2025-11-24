package mixinsawscloudfront


// The RTMP distribution's configuration information.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   streamingDistributionConfigProperty := &StreamingDistributionConfigProperty{
//   	Aliases: []*string{
//   		jsii.String("aliases"),
//   	},
//   	Comment: jsii.String("comment"),
//   	Enabled: jsii.Boolean(false),
//   	Logging: &LoggingProperty{
//   		Bucket: jsii.String("bucket"),
//   		Enabled: jsii.Boolean(false),
//   		Prefix: jsii.String("prefix"),
//   	},
//   	PriceClass: jsii.String("priceClass"),
//   	S3Origin: &S3OriginProperty{
//   		DomainName: jsii.String("domainName"),
//   		OriginAccessIdentity: jsii.String("originAccessIdentity"),
//   	},
//   	TrustedSigners: &TrustedSignersProperty{
//   		AwsAccountNumbers: []*string{
//   			jsii.String("awsAccountNumbers"),
//   		},
//   		Enabled: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-streamingdistribution-streamingdistributionconfig.html
//
type CfnStreamingDistributionPropsMixin_StreamingDistributionConfigProperty struct {
	// A complex type that contains information about CNAMEs (alternate domain names), if any, for this streaming distribution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-streamingdistribution-streamingdistributionconfig.html#cfn-cloudfront-streamingdistribution-streamingdistributionconfig-aliases
	//
	Aliases *[]*string `field:"optional" json:"aliases" yaml:"aliases"`
	// Any comments you want to include about the streaming distribution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-streamingdistribution-streamingdistributionconfig.html#cfn-cloudfront-streamingdistribution-streamingdistributionconfig-comment
	//
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Whether the streaming distribution is enabled to accept user requests for content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-streamingdistribution-streamingdistributionconfig.html#cfn-cloudfront-streamingdistribution-streamingdistributionconfig-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// A complex type that controls whether access logs are written for the streaming distribution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-streamingdistribution-streamingdistributionconfig.html#cfn-cloudfront-streamingdistribution-streamingdistributionconfig-logging
	//
	Logging interface{} `field:"optional" json:"logging" yaml:"logging"`
	// A complex type that contains information about price class for this streaming distribution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-streamingdistribution-streamingdistributionconfig.html#cfn-cloudfront-streamingdistribution-streamingdistributionconfig-priceclass
	//
	PriceClass *string `field:"optional" json:"priceClass" yaml:"priceClass"`
	// A complex type that contains information about the Amazon S3 bucket from which you want CloudFront to get your media files for distribution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-streamingdistribution-streamingdistributionconfig.html#cfn-cloudfront-streamingdistribution-streamingdistributionconfig-s3origin
	//
	S3Origin interface{} `field:"optional" json:"s3Origin" yaml:"s3Origin"`
	// A complex type that specifies any AWS accounts that you want to permit to create signed URLs for private content.
	//
	// If you want the distribution to use signed URLs, include this element; if you want the distribution to use public URLs, remove this element. For more information, see [Serving Private Content through CloudFront](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html) in the *Amazon CloudFront Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-streamingdistribution-streamingdistributionconfig.html#cfn-cloudfront-streamingdistribution-streamingdistributionconfig-trustedsigners
	//
	TrustedSigners interface{} `field:"optional" json:"trustedSigners" yaml:"trustedSigners"`
}

