package awscloudfront


// This object determines the values that CloudFront includes in the cache key.
//
// These values can include HTTP headers, cookies, and URL query strings. CloudFront uses the cache key to find an object in its cache that it can return to the viewer.
//
// The headers, cookies, and query strings that are included in the cache key are automatically included in requests that CloudFront sends to the origin. CloudFront sends a request when it can’t find an object in its cache that matches the request’s cache key. If you want to send values to the origin but *not* include them in the cache key, use `OriginRequestPolicy` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parametersInCacheKeyAndForwardedToOriginProperty := &parametersInCacheKeyAndForwardedToOriginProperty{
//   	cookiesConfig: &cookiesConfigProperty{
//   		cookieBehavior: jsii.String("cookieBehavior"),
//
//   		// the properties below are optional
//   		cookies: []*string{
//   			jsii.String("cookies"),
//   		},
//   	},
//   	enableAcceptEncodingGzip: jsii.Boolean(false),
//   	headersConfig: &headersConfigProperty{
//   		headerBehavior: jsii.String("headerBehavior"),
//
//   		// the properties below are optional
//   		headers: []*string{
//   			jsii.String("headers"),
//   		},
//   	},
//   	queryStringsConfig: &queryStringsConfigProperty{
//   		queryStringBehavior: jsii.String("queryStringBehavior"),
//
//   		// the properties below are optional
//   		queryStrings: []*string{
//   			jsii.String("queryStrings"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	enableAcceptEncodingBrotli: jsii.Boolean(false),
//   }
//
type CfnCachePolicy_ParametersInCacheKeyAndForwardedToOriginProperty struct {
	// An object that determines whether any cookies in viewer requests (and if so, which cookies) are included in the cache key and automatically included in requests that CloudFront sends to the origin.
	CookiesConfig interface{} `field:"required" json:"cookiesConfig" yaml:"cookiesConfig"`
	// A flag that can affect whether the `Accept-Encoding` HTTP header is included in the cache key and included in requests that CloudFront sends to the origin.
	//
	// This field is related to the `EnableAcceptEncodingBrotli` field. If one or both of these fields is `true` *and* the viewer request includes the `Accept-Encoding` header, then CloudFront does the following:
	//
	// - Normalizes the value of the viewer’s `Accept-Encoding` header
	// - Includes the normalized header in the cache key
	// - Includes the normalized header in the request to the origin, if a request is necessary
	//
	// For more information, see [Compression support](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html#cache-policy-compressed-objects) in the *Amazon CloudFront Developer Guide* .
	//
	// If you set this value to `true` , and this cache behavior also has an origin request policy attached, do not include the `Accept-Encoding` header in the origin request policy. CloudFront always includes the `Accept-Encoding` header in origin requests when the value of this field is `true` , so including this header in an origin request policy has no effect.
	//
	// If both of these fields are `false` , then CloudFront treats the `Accept-Encoding` header the same as any other HTTP header in the viewer request. By default, it’s not included in the cache key and it’s not included in origin requests. In this case, you can manually add `Accept-Encoding` to the headers whitelist like any other HTTP header.
	EnableAcceptEncodingGzip interface{} `field:"required" json:"enableAcceptEncodingGzip" yaml:"enableAcceptEncodingGzip"`
	// An object that determines whether any HTTP headers (and if so, which headers) are included in the cache key and automatically included in requests that CloudFront sends to the origin.
	HeadersConfig interface{} `field:"required" json:"headersConfig" yaml:"headersConfig"`
	// An object that determines whether any URL query strings in viewer requests (and if so, which query strings) are included in the cache key and automatically included in requests that CloudFront sends to the origin.
	QueryStringsConfig interface{} `field:"required" json:"queryStringsConfig" yaml:"queryStringsConfig"`
	// A flag that can affect whether the `Accept-Encoding` HTTP header is included in the cache key and included in requests that CloudFront sends to the origin.
	//
	// This field is related to the `EnableAcceptEncodingGzip` field. If one or both of these fields is `true` *and* the viewer request includes the `Accept-Encoding` header, then CloudFront does the following:
	//
	// - Normalizes the value of the viewer’s `Accept-Encoding` header
	// - Includes the normalized header in the cache key
	// - Includes the normalized header in the request to the origin, if a request is necessary
	//
	// For more information, see [Compression support](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html#cache-policy-compressed-objects) in the *Amazon CloudFront Developer Guide* .
	//
	// If you set this value to `true` , and this cache behavior also has an origin request policy attached, do not include the `Accept-Encoding` header in the origin request policy. CloudFront always includes the `Accept-Encoding` header in origin requests when the value of this field is `true` , so including this header in an origin request policy has no effect.
	//
	// If both of these fields are `false` , then CloudFront treats the `Accept-Encoding` header the same as any other HTTP header in the viewer request. By default, it’s not included in the cache key and it’s not included in origin requests. In this case, you can manually add `Accept-Encoding` to the headers whitelist like any other HTTP header.
	EnableAcceptEncodingBrotli interface{} `field:"optional" json:"enableAcceptEncodingBrotli" yaml:"enableAcceptEncodingBrotli"`
}

