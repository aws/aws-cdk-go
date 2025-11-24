package mixinsawscloudfront


// An origin request policy configuration.
//
// This configuration determines the values that CloudFront includes in requests that it sends to the origin. Each request that CloudFront sends to the origin includes the following:
//
// - The request body and the URL path (without the domain name) from the viewer request.
// - The headers that CloudFront automatically includes in every origin request, including `Host` , `User-Agent` , and `X-Amz-Cf-Id` .
// - All HTTP headers, cookies, and URL query strings that are specified in the cache policy or the origin request policy. These can include items from the viewer request and, in the case of headers, additional ones that are added by CloudFront.
//
// CloudFront sends a request when it can't find an object in its cache that matches the request. If you want to send values to the origin and also include them in the cache key, use `CachePolicy` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   originRequestPolicyConfigProperty := &OriginRequestPolicyConfigProperty{
//   	Comment: jsii.String("comment"),
//   	CookiesConfig: &CookiesConfigProperty{
//   		CookieBehavior: jsii.String("cookieBehavior"),
//   		Cookies: []*string{
//   			jsii.String("cookies"),
//   		},
//   	},
//   	HeadersConfig: &HeadersConfigProperty{
//   		HeaderBehavior: jsii.String("headerBehavior"),
//   		Headers: []*string{
//   			jsii.String("headers"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	QueryStringsConfig: &QueryStringsConfigProperty{
//   		QueryStringBehavior: jsii.String("queryStringBehavior"),
//   		QueryStrings: []*string{
//   			jsii.String("queryStrings"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-originrequestpolicy-originrequestpolicyconfig.html
//
type CfnOriginRequestPolicyPropsMixin_OriginRequestPolicyConfigProperty struct {
	// A comment to describe the origin request policy.
	//
	// The comment cannot be longer than 128 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-originrequestpolicy-originrequestpolicyconfig.html#cfn-cloudfront-originrequestpolicy-originrequestpolicyconfig-comment
	//
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// The cookies from viewer requests to include in origin requests.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-originrequestpolicy-originrequestpolicyconfig.html#cfn-cloudfront-originrequestpolicy-originrequestpolicyconfig-cookiesconfig
	//
	CookiesConfig interface{} `field:"optional" json:"cookiesConfig" yaml:"cookiesConfig"`
	// The HTTP headers to include in origin requests.
	//
	// These can include headers from viewer requests and additional headers added by CloudFront.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-originrequestpolicy-originrequestpolicyconfig.html#cfn-cloudfront-originrequestpolicy-originrequestpolicyconfig-headersconfig
	//
	HeadersConfig interface{} `field:"optional" json:"headersConfig" yaml:"headersConfig"`
	// A unique name to identify the origin request policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-originrequestpolicy-originrequestpolicyconfig.html#cfn-cloudfront-originrequestpolicy-originrequestpolicyconfig-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The URL query strings from viewer requests to include in origin requests.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-originrequestpolicy-originrequestpolicyconfig.html#cfn-cloudfront-originrequestpolicy-originrequestpolicyconfig-querystringsconfig
	//
	QueryStringsConfig interface{} `field:"optional" json:"queryStringsConfig" yaml:"queryStringsConfig"`
}

