package awscloudfront


// A complex type that controls whether access logs are written for the streaming distribution.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loggingProperty := &loggingProperty{
//   	bucket: jsii.String("bucket"),
//   	enabled: jsii.Boolean(false),
//   	prefix: jsii.String("prefix"),
//   }
//
type CfnStreamingDistribution_LoggingProperty struct {
	// The Amazon S3 bucket to store the access logs in, for example, `myawslogbucket.s3.amazonaws.com` .
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Specifies whether you want CloudFront to save access logs to an Amazon S3 bucket.
	//
	// If you don't want to enable logging when you create a streaming distribution or if you want to disable logging for an existing streaming distribution, specify `false` for `Enabled` , and specify `empty Bucket` and `Prefix` elements. If you specify `false` for `Enabled` but you specify values for `Bucket` and `Prefix` , the values are automatically deleted.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// An optional string that you want CloudFront to prefix to the access log filenames for this streaming distribution, for example, `myprefix/` .
	//
	// If you want to enable logging, but you don't want to specify a prefix, you still must include an empty `Prefix` element in the `Logging` element.
	Prefix *string `field:"required" json:"prefix" yaml:"prefix"`
}

