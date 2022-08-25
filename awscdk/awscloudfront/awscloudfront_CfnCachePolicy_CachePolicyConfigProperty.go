package awscloudfront


// A cache policy configuration.
//
// This configuration determines the following:
//
// - The values that CloudFront includes in the cache key. These values can include HTTP headers, cookies, and URL query strings. CloudFront uses the cache key to find an object in its cache that it can return to the viewer.
// - The default, minimum, and maximum time to live (TTL) values that you want objects to stay in the CloudFront cache.
//
// The headers, cookies, and query strings that are included in the cache key are automatically included in requests that CloudFront sends to the origin. CloudFront sends a request when it can’t find a valid object in its cache that matches the request’s cache key. If you want to send values to the origin but *not* include them in the cache key, use `OriginRequestPolicy` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cachePolicyConfigProperty := &cachePolicyConfigProperty{
//   	defaultTtl: jsii.Number(123),
//   	maxTtl: jsii.Number(123),
//   	minTtl: jsii.Number(123),
//   	name: jsii.String("name"),
//   	parametersInCacheKeyAndForwardedToOrigin: &parametersInCacheKeyAndForwardedToOriginProperty{
//   		cookiesConfig: &cookiesConfigProperty{
//   			cookieBehavior: jsii.String("cookieBehavior"),
//
//   			// the properties below are optional
//   			cookies: []*string{
//   				jsii.String("cookies"),
//   			},
//   		},
//   		enableAcceptEncodingGzip: jsii.Boolean(false),
//   		headersConfig: &headersConfigProperty{
//   			headerBehavior: jsii.String("headerBehavior"),
//
//   			// the properties below are optional
//   			headers: []*string{
//   				jsii.String("headers"),
//   			},
//   		},
//   		queryStringsConfig: &queryStringsConfigProperty{
//   			queryStringBehavior: jsii.String("queryStringBehavior"),
//
//   			// the properties below are optional
//   			queryStrings: []*string{
//   				jsii.String("queryStrings"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		enableAcceptEncodingBrotli: jsii.Boolean(false),
//   	},
//
//   	// the properties below are optional
//   	comment: jsii.String("comment"),
//   }
//
type CfnCachePolicy_CachePolicyConfigProperty struct {
	// The default amount of time, in seconds, that you want objects to stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
	//
	// CloudFront uses this value as the object’s time to live (TTL) only when the origin does *not* send `Cache-Control` or `Expires` headers with the object. For more information, see [Managing How Long Content Stays in an Edge Cache (Expiration)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Expiration.html) in the *Amazon CloudFront Developer Guide* .
	//
	// The default value for this field is 86400 seconds (one day). If the value of `MinTTL` is more than 86400 seconds, then the default value for this field is the same as the value of `MinTTL` .
	DefaultTtl *float64 `field:"required" json:"defaultTtl" yaml:"defaultTtl"`
	// The maximum amount of time, in seconds, that objects stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
	//
	// CloudFront uses this value only when the origin sends `Cache-Control` or `Expires` headers with the object. For more information, see [Managing How Long Content Stays in an Edge Cache (Expiration)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Expiration.html) in the *Amazon CloudFront Developer Guide* .
	//
	// The default value for this field is 31536000 seconds (one year). If the value of `MinTTL` or `DefaultTTL` is more than 31536000 seconds, then the default value for this field is the same as the value of `DefaultTTL` .
	MaxTtl *float64 `field:"required" json:"maxTtl" yaml:"maxTtl"`
	// The minimum amount of time, in seconds, that you want objects to stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
	//
	// For more information, see [Managing How Long Content Stays in an Edge Cache (Expiration)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Expiration.html) in the *Amazon CloudFront Developer Guide* .
	MinTtl *float64 `field:"required" json:"minTtl" yaml:"minTtl"`
	// A unique name to identify the cache policy.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The HTTP headers, cookies, and URL query strings to include in the cache key.
	//
	// The values included in the cache key are automatically included in requests that CloudFront sends to the origin.
	ParametersInCacheKeyAndForwardedToOrigin interface{} `field:"required" json:"parametersInCacheKeyAndForwardedToOrigin" yaml:"parametersInCacheKeyAndForwardedToOrigin"`
	// A comment to describe the cache policy.
	//
	// The comment cannot be longer than 128 characters.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
}

