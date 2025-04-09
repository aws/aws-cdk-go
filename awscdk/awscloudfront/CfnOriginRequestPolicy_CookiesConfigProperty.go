package awscloudfront


// An object that determines whether any cookies in viewer requests (and if so, which cookies) are included in requests that CloudFront sends to the origin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cookiesConfigProperty := &CookiesConfigProperty{
//   	CookieBehavior: jsii.String("cookieBehavior"),
//
//   	// the properties below are optional
//   	Cookies: []*string{
//   		jsii.String("cookies"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-originrequestpolicy-cookiesconfig.html
//
type CfnOriginRequestPolicy_CookiesConfigProperty struct {
	// Determines whether cookies in viewer requests are included in requests that CloudFront sends to the origin. Valid values are:.
	//
	// - `none` – No cookies in viewer requests are included in requests that CloudFront sends to the origin. Even when this field is set to `none` , any cookies that are listed in a `CachePolicy` *are* included in origin requests.
	// - `whitelist` – Only the cookies in viewer requests that are listed in the `CookieNames` type are included in requests that CloudFront sends to the origin.
	// - `all` – All cookies in viewer requests are included in requests that CloudFront sends to the origin.
	// - `allExcept` – All cookies in viewer requests are included in requests that CloudFront sends to the origin, **except** for those listed in the `CookieNames` type, which are not included.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-originrequestpolicy-cookiesconfig.html#cfn-cloudfront-originrequestpolicy-cookiesconfig-cookiebehavior
	//
	CookieBehavior *string `field:"required" json:"cookieBehavior" yaml:"cookieBehavior"`
	// Contains a list of cookie names.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-originrequestpolicy-cookiesconfig.html#cfn-cloudfront-originrequestpolicy-cookiesconfig-cookies
	//
	Cookies *[]*string `field:"optional" json:"cookies" yaml:"cookies"`
}

