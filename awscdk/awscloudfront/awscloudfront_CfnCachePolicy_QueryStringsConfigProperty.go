package awscloudfront


// An object that determines whether any URL query strings in viewer requests (and if so, which query strings) are included in the cache key and automatically included in requests that CloudFront sends to the origin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   queryStringsConfigProperty := &queryStringsConfigProperty{
//   	queryStringBehavior: jsii.String("queryStringBehavior"),
//
//   	// the properties below are optional
//   	queryStrings: []*string{
//   		jsii.String("queryStrings"),
//   	},
//   }
//
type CfnCachePolicy_QueryStringsConfigProperty struct {
	// Determines whether any URL query strings in viewer requests are included in the cache key and automatically included in requests that CloudFront sends to the origin.
	//
	// Valid values are:
	//
	// - `none` – Query strings in viewer requests are not included in the cache key and are not automatically included in requests that CloudFront sends to the origin. Even when this field is set to `none` , any query strings that are listed in an `OriginRequestPolicy` *are* included in origin requests.
	// - `whitelist` – The query strings in viewer requests that are listed in the `QueryStringNames` type are included in the cache key and automatically included in requests that CloudFront sends to the origin.
	// - `allExcept` – All query strings in viewer requests that are **not** listed in the `QueryStringNames` type are included in the cache key and automatically included in requests that CloudFront sends to the origin.
	// - `all` – All query strings in viewer requests are included in the cache key and are automatically included in requests that CloudFront sends to the origin.
	QueryStringBehavior *string `field:"required" json:"queryStringBehavior" yaml:"queryStringBehavior"`
	// Contains a list of query string names.
	QueryStrings *[]*string `field:"optional" json:"queryStrings" yaml:"queryStrings"`
}

