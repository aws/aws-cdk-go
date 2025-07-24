package awscloudfront


// A cache policy configuration.
//
// This configuration determines the following:
//
// - The values that CloudFront includes in the cache key. These values can include HTTP headers, cookies, and URL query strings. CloudFront uses the cache key to find an object in its cache that it can return to the viewer.
// - The default, minimum, and maximum time to live (TTL) values that you want objects to stay in the CloudFront cache.
//
// The headers, cookies, and query strings that are included in the cache key are also included in requests that CloudFront sends to the origin. CloudFront sends a request when it can't find a valid object in its cache that matches the request's cache key. If you want to send values to the origin but *not* include them in the cache key, use `OriginRequestPolicy` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cachePolicyConfigProperty := &CachePolicyConfigProperty{
//   	DefaultTtl: jsii.Number(123),
//   	MaxTtl: jsii.Number(123),
//   	MinTtl: jsii.Number(123),
//   	Name: jsii.String("name"),
//   	ParametersInCacheKeyAndForwardedToOrigin: &ParametersInCacheKeyAndForwardedToOriginProperty{
//   		CookiesConfig: &CookiesConfigProperty{
//   			CookieBehavior: jsii.String("cookieBehavior"),
//
//   			// the properties below are optional
//   			Cookies: []*string{
//   				jsii.String("cookies"),
//   			},
//   		},
//   		EnableAcceptEncodingGzip: jsii.Boolean(false),
//   		HeadersConfig: &HeadersConfigProperty{
//   			HeaderBehavior: jsii.String("headerBehavior"),
//
//   			// the properties below are optional
//   			Headers: []*string{
//   				jsii.String("headers"),
//   			},
//   		},
//   		QueryStringsConfig: &QueryStringsConfigProperty{
//   			QueryStringBehavior: jsii.String("queryStringBehavior"),
//
//   			// the properties below are optional
//   			QueryStrings: []*string{
//   				jsii.String("queryStrings"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		EnableAcceptEncodingBrotli: jsii.Boolean(false),
//   	},
//
//   	// the properties below are optional
//   	Comment: jsii.String("comment"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-cachepolicy-cachepolicyconfig.html
//
type CfnCachePolicy_CachePolicyConfigProperty struct {
	// The default amount of time, in seconds, that you want objects to stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
	//
	// CloudFront uses this value as the object's time to live (TTL) only when the origin does *not* send `Cache-Control` or `Expires` headers with the object. For more information, see [Managing How Long Content Stays in an Edge Cache (Expiration)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Expiration.html) in the *Amazon CloudFront Developer Guide* .
	//
	// The default value for this field is 86400 seconds (one day). If the value of `MinTTL` is more than 86400 seconds, then the default value for this field is the same as the value of `MinTTL` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-cachepolicy-cachepolicyconfig.html#cfn-cloudfront-cachepolicy-cachepolicyconfig-defaultttl
	//
	DefaultTtl *float64 `field:"required" json:"defaultTtl" yaml:"defaultTtl"`
	// The maximum amount of time, in seconds, that objects stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
	//
	// CloudFront uses this value only when the origin sends `Cache-Control` or `Expires` headers with the object. For more information, see [Managing How Long Content Stays in an Edge Cache (Expiration)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Expiration.html) in the *Amazon CloudFront Developer Guide* .
	//
	// The default value for this field is 31536000 seconds (one year). If the value of `MinTTL` or `DefaultTTL` is more than 31536000 seconds, then the default value for this field is the same as the value of `DefaultTTL` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-cachepolicy-cachepolicyconfig.html#cfn-cloudfront-cachepolicy-cachepolicyconfig-maxttl
	//
	MaxTtl *float64 `field:"required" json:"maxTtl" yaml:"maxTtl"`
	// The minimum amount of time, in seconds, that you want objects to stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
	//
	// For more information, see [Managing How Long Content Stays in an Edge Cache (Expiration)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Expiration.html) in the *Amazon CloudFront Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-cachepolicy-cachepolicyconfig.html#cfn-cloudfront-cachepolicy-cachepolicyconfig-minttl
	//
	MinTtl *float64 `field:"required" json:"minTtl" yaml:"minTtl"`
	// A unique name to identify the cache policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-cachepolicy-cachepolicyconfig.html#cfn-cloudfront-cachepolicy-cachepolicyconfig-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The HTTP headers, cookies, and URL query strings to include in the cache key.
	//
	// The values included in the cache key are also included in requests that CloudFront sends to the origin.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-cachepolicy-cachepolicyconfig.html#cfn-cloudfront-cachepolicy-cachepolicyconfig-parametersincachekeyandforwardedtoorigin
	//
	ParametersInCacheKeyAndForwardedToOrigin interface{} `field:"required" json:"parametersInCacheKeyAndForwardedToOrigin" yaml:"parametersInCacheKeyAndForwardedToOrigin"`
	// A comment to describe the cache policy.
	//
	// The comment cannot be longer than 128 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-cachepolicy-cachepolicyconfig.html#cfn-cloudfront-cachepolicy-cachepolicyconfig-comment
	//
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
}

