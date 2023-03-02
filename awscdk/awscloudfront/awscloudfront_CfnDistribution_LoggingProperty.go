package awscloudfront


// A complex type that controls whether access logs are written for the distribution.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loggingProperty := &loggingProperty{
//   	bucket: jsii.String("bucket"),
//
//   	// the properties below are optional
//   	includeCookies: jsii.Boolean(false),
//   	prefix: jsii.String("prefix"),
//   }
//
type CfnDistribution_LoggingProperty struct {
	// The Amazon S3 bucket to store the access logs in, for example, `myawslogbucket.s3.amazonaws.com` .
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Specifies whether you want CloudFront to include cookies in access logs, specify `true` for `IncludeCookies` .
	//
	// If you choose to include cookies in logs, CloudFront logs all cookies regardless of how you configure the cache behaviors for this distribution. If you don't want to include cookies when you create a distribution or if you want to disable include cookies for an existing distribution, specify `false` for `IncludeCookies` .
	IncludeCookies interface{} `field:"optional" json:"includeCookies" yaml:"includeCookies"`
	// An optional string that you want CloudFront to prefix to the access log `filenames` for this distribution, for example, `myprefix/` .
	//
	// If you want to enable logging, but you don't want to specify a prefix, you still must include an empty `Prefix` element in the `Logging` element.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

