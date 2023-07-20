package awswafv2


// A custom header for custom request and response handling.
//
// This is used in `CustomResponse` and `CustomRequestHandling`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customHTTPHeaderProperty := &CustomHTTPHeaderProperty{
//   	Name: jsii.String("name"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-customhttpheader.html
//
type CfnRuleGroup_CustomHTTPHeaderProperty struct {
	// The name of the custom header.
	//
	// For custom request header insertion, when AWS WAF inserts the header into the request, it prefixes this name `x-amzn-waf-` , to avoid confusion with the headers that are already in the request. For example, for the header name `sample` , AWS WAF inserts the header `x-amzn-waf-sample` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-customhttpheader.html#cfn-wafv2-rulegroup-customhttpheader-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The value of the custom header.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-customhttpheader.html#cfn-wafv2-rulegroup-customhttpheader-value
	//
	Value *string `field:"required" json:"value" yaml:"value"`
}

