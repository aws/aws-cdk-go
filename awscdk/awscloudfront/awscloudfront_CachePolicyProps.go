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
//   myCachePolicy := cloudfront.NewCachePolicy(this, jsii.String("myCachePolicy"), &cachePolicyProps{
//   	cachePolicyName: jsii.String("MyPolicy"),
//   	comment: jsii.String("A default policy"),
//   	defaultTtl: awscdk.Duration.days(jsii.Number(2)),
//   	minTtl: awscdk.Duration.minutes(jsii.Number(1)),
//   	maxTtl: awscdk.Duration.days(jsii.Number(10)),
//   	cookieBehavior: cloudfront.cacheCookieBehavior.all(),
//   	headerBehavior: cloudfront.cacheHeaderBehavior.allowList(jsii.String("X-CustomHeader")),
//   	queryStringBehavior: cloudfront.cacheQueryStringBehavior.denyList(jsii.String("username")),
//   	enableAcceptEncodingGzip: jsii.Boolean(true),
//   	enableAcceptEncodingBrotli: jsii.Boolean(true),
//   })
//   cloudfront.NewDistribution(this, jsii.String("myDistCustomPolicy"), &distributionProps{
//   	defaultBehavior: &behaviorOptions{
//   		origin: bucketOrigin,
//   		cachePolicy: myCachePolicy,
//   	},
//   })
//
type CachePolicyProps struct {
	// A unique name to identify the cache policy.
	//
	// The name must only include '-', '_', or alphanumeric characters.
	CachePolicyName *string `field:"optional" json:"cachePolicyName" yaml:"cachePolicyName"`
	// A comment to describe the cache policy.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Determines whether any cookies in viewer requests are included in the cache key and automatically included in requests that CloudFront sends to the origin.
	CookieBehavior CacheCookieBehavior `field:"optional" json:"cookieBehavior" yaml:"cookieBehavior"`
	// The default amount of time for objects to stay in the CloudFront cache.
	//
	// Only used when the origin does not send Cache-Control or Expires headers with the object.
	DefaultTtl awscdk.Duration `field:"optional" json:"defaultTtl" yaml:"defaultTtl"`
	// Whether to normalize and include the `Accept-Encoding` header in the cache key when the `Accept-Encoding` header is 'br'.
	EnableAcceptEncodingBrotli *bool `field:"optional" json:"enableAcceptEncodingBrotli" yaml:"enableAcceptEncodingBrotli"`
	// Whether to normalize and include the `Accept-Encoding` header in the cache key when the `Accept-Encoding` header is 'gzip'.
	EnableAcceptEncodingGzip *bool `field:"optional" json:"enableAcceptEncodingGzip" yaml:"enableAcceptEncodingGzip"`
	// Determines whether any HTTP headers are included in the cache key and automatically included in requests that CloudFront sends to the origin.
	HeaderBehavior CacheHeaderBehavior `field:"optional" json:"headerBehavior" yaml:"headerBehavior"`
	// The maximum amount of time for objects to stay in the CloudFront cache.
	//
	// CloudFront uses this value only when the origin sends Cache-Control or Expires headers with the object.
	MaxTtl awscdk.Duration `field:"optional" json:"maxTtl" yaml:"maxTtl"`
	// The minimum amount of time for objects to stay in the CloudFront cache.
	MinTtl awscdk.Duration `field:"optional" json:"minTtl" yaml:"minTtl"`
	// Determines whether any query strings are included in the cache key and automatically included in requests that CloudFront sends to the origin.
	QueryStringBehavior CacheQueryStringBehavior `field:"optional" json:"queryStringBehavior" yaml:"queryStringBehavior"`
}

