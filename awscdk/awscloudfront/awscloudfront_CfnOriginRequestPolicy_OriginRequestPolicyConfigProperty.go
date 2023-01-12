package awscloudfront


// An origin request policy configuration.
//
// This configuration determines the values that CloudFront includes in requests that it sends to the origin. Each request that CloudFront sends to the origin includes the following:
//
// - The request body and the URL path (without the domain name) from the viewer request.
// - The headers that CloudFront automatically includes in every origin request, including `Host` , `User-Agent` , and `X-Amz-Cf-Id` .
// - All HTTP headers, cookies, and URL query strings that are specified in the cache policy or the origin request policy. These can include items from the viewer request and, in the case of headers, additional ones that are added by CloudFront.
//
// CloudFront sends a request when it can’t find an object in its cache that matches the request. If you want to send values to the origin and also include them in the cache key, use `CachePolicy` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   originRequestPolicyConfigProperty := &originRequestPolicyConfigProperty{
//   	cookiesConfig: &cookiesConfigProperty{
//   		cookieBehavior: jsii.String("cookieBehavior"),
//
//   		// the properties below are optional
//   		cookies: []*string{
//   			jsii.String("cookies"),
//   		},
//   	},
//   	headersConfig: &headersConfigProperty{
//   		headerBehavior: jsii.String("headerBehavior"),
//
//   		// the properties below are optional
//   		headers: []*string{
//   			jsii.String("headers"),
//   		},
//   	},
//   	name: jsii.String("name"),
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
//   	comment: jsii.String("comment"),
//   }
//
type CfnOriginRequestPolicy_OriginRequestPolicyConfigProperty struct {
	// The cookies from viewer requests to include in origin requests.
	CookiesConfig interface{} `field:"required" json:"cookiesConfig" yaml:"cookiesConfig"`
	// The HTTP headers to include in origin requests.
	//
	// These can include headers from viewer requests and additional headers added by CloudFront.
	HeadersConfig interface{} `field:"required" json:"headersConfig" yaml:"headersConfig"`
	// A unique name to identify the origin request policy.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The URL query strings from viewer requests to include in origin requests.
	QueryStringsConfig interface{} `field:"required" json:"queryStringsConfig" yaml:"queryStringsConfig"`
	// A comment to describe the origin request policy.
	//
	// The comment cannot be longer than 128 characters.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
}

