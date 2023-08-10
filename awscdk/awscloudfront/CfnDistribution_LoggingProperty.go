package awscloudfront


// A complex type that controls whether access logs are written for the distribution.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loggingProperty := &LoggingProperty{
//   	Bucket: jsii.String("bucket"),
//
//   	// the properties below are optional
//   	IncludeCookies: jsii.Boolean(false),
//   	Prefix: jsii.String("prefix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-logging.html
//
type CfnDistribution_LoggingProperty struct {
	// The Amazon S3 bucket to store the access logs in, for example, `myawslogbucket.s3.amazonaws.com` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-logging.html#cfn-cloudfront-distribution-logging-bucket
	//
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Specifies whether you want CloudFront to include cookies in access logs, specify `true` for `IncludeCookies` .
	//
	// If you choose to include cookies in logs, CloudFront logs all cookies regardless of how you configure the cache behaviors for this distribution. If you don't want to include cookies when you create a distribution or if you want to disable include cookies for an existing distribution, specify `false` for `IncludeCookies` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-logging.html#cfn-cloudfront-distribution-logging-includecookies
	//
	IncludeCookies interface{} `field:"optional" json:"includeCookies" yaml:"includeCookies"`
	// An optional string that you want CloudFront to prefix to the access log `filenames` for this distribution, for example, `myprefix/` .
	//
	// If you want to enable logging, but you don't want to specify a prefix, you still must include an empty `Prefix` element in the `Logging` element.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-logging.html#cfn-cloudfront-distribution-logging-prefix
	//
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

