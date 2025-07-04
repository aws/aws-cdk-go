package awscloudfront


// An object that determines whether any URL query strings in viewer requests (and if so, which query strings) are included in the cache key and in requests that CloudFront sends to the origin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   queryStringsConfigProperty := &QueryStringsConfigProperty{
//   	QueryStringBehavior: jsii.String("queryStringBehavior"),
//
//   	// the properties below are optional
//   	QueryStrings: []*string{
//   		jsii.String("queryStrings"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-cachepolicy-querystringsconfig.html
//
type CfnCachePolicy_QueryStringsConfigProperty struct {
	// Determines whether any URL query strings in viewer requests are included in the cache key and in requests that CloudFront sends to the origin.
	//
	// Valid values are:
	//
	// - `none` – No query strings in viewer requests are included in the cache key or in requests that CloudFront sends to the origin. Even when this field is set to `none` , any query strings that are listed in an `OriginRequestPolicy` *are* included in origin requests.
	// - `whitelist` – Only the query strings in viewer requests that are listed in the `QueryStringNames` type are included in the cache key and in requests that CloudFront sends to the origin.
	// - `allExcept` – All query strings in viewer requests are included in the cache key and in requests that CloudFront sends to the origin, **except** those that are listed in the `QueryStringNames` type, which are not included.
	// - `all` – All query strings in viewer requests are included in the cache key and in requests that CloudFront sends to the origin.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-cachepolicy-querystringsconfig.html#cfn-cloudfront-cachepolicy-querystringsconfig-querystringbehavior
	//
	QueryStringBehavior *string `field:"required" json:"queryStringBehavior" yaml:"queryStringBehavior"`
	// Contains a list of query string names.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-cachepolicy-querystringsconfig.html#cfn-cloudfront-cachepolicy-querystringsconfig-querystrings
	//
	QueryStrings *[]*string `field:"optional" json:"queryStrings" yaml:"queryStrings"`
}

