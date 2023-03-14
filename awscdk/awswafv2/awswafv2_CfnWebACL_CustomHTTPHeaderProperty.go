package awswafv2


// A custom header for custom request and response handling.
//
// This is used in `CustomResponse` and `CustomRequestHandling` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customHTTPHeaderProperty := &customHTTPHeaderProperty{
//   	name: jsii.String("name"),
//   	value: jsii.String("value"),
//   }
//
type CfnWebACL_CustomHTTPHeaderProperty struct {
	// The name of the custom header.
	//
	// For custom request header insertion, when AWS WAF inserts the header into the request, it prefixes this name `x-amzn-waf-` , to avoid confusion with the headers that are already in the request. For example, for the header name `sample` , AWS WAF inserts the header `x-amzn-waf-sample` .
	Name *string `field:"required" json:"name" yaml:"name"`
	// The value of the custom header.
	Value *string `field:"required" json:"value" yaml:"value"`
}

