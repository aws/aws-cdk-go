package awswafv2


// A custom response to send to the client.
//
// You can define a custom response for rule actions and default web ACL actions that are set to the block action.
//
// For information about customizing web requests and responses, see [Customizing web requests and responses in AWS WAF](https://docs.aws.amazon.com/waf/latest/developerguide/waf-custom-request-response.html) in the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customResponseProperty := &customResponseProperty{
//   	responseCode: jsii.Number(123),
//
//   	// the properties below are optional
//   	customResponseBodyKey: jsii.String("customResponseBodyKey"),
//   	responseHeaders: []interface{}{
//   		&customHTTPHeaderProperty{
//   			name: jsii.String("name"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnWebACL_CustomResponseProperty struct {
	// The HTTP status code to return to the client.
	//
	// For a list of status codes that you can use in your custom reqponses, see [Supported status codes for custom response](https://docs.aws.amazon.com/waf/latest/developerguide/customizing-the-response-status-codes.html) in the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) .
	ResponseCode *float64 `field:"required" json:"responseCode" yaml:"responseCode"`
	// References the response body that you want AWS WAF to return to the web request client.
	//
	// You can define a custom response for a rule action or a default web ACL action that is set to block. To do this, you first define the response body key and value in the `CustomResponseBodies` setting for the `WebACL` or `RuleGroup` where you want to use it. Then, in the rule action or web ACL default action `BlockAction` setting, you reference the response body using this key.
	CustomResponseBodyKey *string `field:"optional" json:"customResponseBodyKey" yaml:"customResponseBodyKey"`
	// The HTTP headers to use in the response. Duplicate header names are not allowed.
	//
	// For information about the limits on count and size for custom request and response settings, see [AWS WAF quotas](https://docs.aws.amazon.com/waf/latest/developerguide/limits.html) in the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) .
	ResponseHeaders interface{} `field:"optional" json:"responseHeaders" yaml:"responseHeaders"`
}

