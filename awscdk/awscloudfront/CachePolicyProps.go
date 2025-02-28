package awscloudfront

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for creating a Cache Policy.
//
// Example:
//   // Creating a custom cache policy for a Distribution -- all parameters optional
//   var bucketOrigin s3Origin
//
//   myCachePolicy := cloudfront.NewCachePolicy(this, jsii.String("myCachePolicy"), &CachePolicyProps{
//   	CachePolicyName: jsii.String("MyPolicy"),
//   	Comment: jsii.String("A default policy"),
//   	DefaultTtl: awscdk.Duration_Days(jsii.Number(2)),
//   	MinTtl: awscdk.Duration_Minutes(jsii.Number(1)),
//   	MaxTtl: awscdk.Duration_*Days(jsii.Number(10)),
//   	CookieBehavior: cloudfront.CacheCookieBehavior_All(),
//   	HeaderBehavior: cloudfront.CacheHeaderBehavior_AllowList(jsii.String("X-CustomHeader")),
//   	QueryStringBehavior: cloudfront.CacheQueryStringBehavior_DenyList(jsii.String("username")),
//   	EnableAcceptEncodingGzip: jsii.Boolean(true),
//   	EnableAcceptEncodingBrotli: jsii.Boolean(true),
//   })
//   cloudfront.NewDistribution(this, jsii.String("myDistCustomPolicy"), &DistributionProps{
//   	DefaultBehavior: &BehaviorOptions{
//   		Origin: bucketOrigin,
//   		CachePolicy: myCachePolicy,
//   	},
//   })
//
type CachePolicyProps struct {
	// A unique name to identify the cache policy.
	//
	// The name must only include '-', '_', or alphanumeric characters.
	// Default: - generated from the `id`.
	//
	CachePolicyName *string `field:"optional" json:"cachePolicyName" yaml:"cachePolicyName"`
	// A comment to describe the cache policy.
	//
	// The comment cannot be longer than 128 characters.
	// Default: - no comment.
	//
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Determines whether any cookies in viewer requests are included in the cache key and automatically included in requests that CloudFront sends to the origin.
	// Default: CacheCookieBehavior.none()
	//
	CookieBehavior CacheCookieBehavior `field:"optional" json:"cookieBehavior" yaml:"cookieBehavior"`
	// The default amount of time for objects to stay in the CloudFront cache.
	//
	// Only used when the origin does not send Cache-Control or Expires headers with the object.
	// Default: - The greater of 1 day and ``minTtl``.
	//
	DefaultTtl awscdk.Duration `field:"optional" json:"defaultTtl" yaml:"defaultTtl"`
	// Whether to normalize and include the `Accept-Encoding` header in the cache key when the `Accept-Encoding` header is 'br'.
	// Default: false.
	//
	EnableAcceptEncodingBrotli *bool `field:"optional" json:"enableAcceptEncodingBrotli" yaml:"enableAcceptEncodingBrotli"`
	// Whether to normalize and include the `Accept-Encoding` header in the cache key when the `Accept-Encoding` header is 'gzip'.
	// Default: false.
	//
	EnableAcceptEncodingGzip *bool `field:"optional" json:"enableAcceptEncodingGzip" yaml:"enableAcceptEncodingGzip"`
	// Determines whether any HTTP headers are included in the cache key and automatically included in requests that CloudFront sends to the origin.
	// Default: CacheHeaderBehavior.none()
	//
	HeaderBehavior CacheHeaderBehavior `field:"optional" json:"headerBehavior" yaml:"headerBehavior"`
	// The maximum amount of time for objects to stay in the CloudFront cache.
	//
	// CloudFront uses this value only when the origin sends Cache-Control or Expires headers with the object.
	// Default: - The greater of 1 year and ``defaultTtl``.
	//
	MaxTtl awscdk.Duration `field:"optional" json:"maxTtl" yaml:"maxTtl"`
	// The minimum amount of time for objects to stay in the CloudFront cache.
	// Default: Duration.seconds(0)
	//
	MinTtl awscdk.Duration `field:"optional" json:"minTtl" yaml:"minTtl"`
	// Determines whether any query strings are included in the cache key and automatically included in requests that CloudFront sends to the origin.
	// Default: CacheQueryStringBehavior.none()
	//
	QueryStringBehavior CacheQueryStringBehavior `field:"optional" json:"queryStringBehavior" yaml:"queryStringBehavior"`
}

