package awscloudfront


// An object that determines whether any HTTP headers (and if so, which headers) are included in the cache key and automatically included in requests that CloudFront sends to the origin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   headersConfigProperty := &headersConfigProperty{
//   	headerBehavior: jsii.String("headerBehavior"),
//
//   	// the properties below are optional
//   	headers: []*string{
//   		jsii.String("headers"),
//   	},
//   }
//
type CfnCachePolicy_HeadersConfigProperty struct {
	// Determines whether any HTTP headers are included in the cache key and automatically included in requests that CloudFront sends to the origin.
	//
	// Valid values are:
	//
	// - `none` – HTTP headers are not included in the cache key and are not automatically included in requests that CloudFront sends to the origin. Even when this field is set to `none` , any headers that are listed in an `OriginRequestPolicy` *are* included in origin requests.
	// - `whitelist` – The HTTP headers that are listed in the `Headers` type are included in the cache key and are automatically included in requests that CloudFront sends to the origin.
	HeaderBehavior *string `field:"required" json:"headerBehavior" yaml:"headerBehavior"`
	// Contains a list of HTTP header names.
	Headers *[]*string `field:"optional" json:"headers" yaml:"headers"`
}

