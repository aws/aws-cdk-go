package awswafv2


// Specifies that AWS WAF should run a `CAPTCHA` check against the request:.
//
// - If the request includes a valid, unexpired `CAPTCHA` token, AWS WAF allows the web request inspection to proceed to the next rule, similar to a `CountAction` .
// - If the request doesn't include a valid, unexpired `CAPTCHA` token, AWS WAF discontinues the web ACL evaluation of the request and blocks it from going to its intended destination.
//
// AWS WAF generates a response that it sends back to the client, which includes the following:
//
// - The header `x-amzn-waf-action` with a value of `captcha` .
// - The HTTP status code `405 Method Not Allowed` .
// - If the request contains an `Accept` header with a value of `text/html` , the response includes a `CAPTCHA` challenge.
//
// You can configure the expiration time in the `CaptchaConfig` `ImmunityTimeProperty` setting at the rule and web ACL level. The rule setting overrides the web ACL setting.
//
// This action option is available for rules. It isn't available for web ACL default actions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   captchaActionProperty := &CaptchaActionProperty{
//   	CustomRequestHandling: &CustomRequestHandlingProperty{
//   		InsertHeaders: []interface{}{
//   			&CustomHTTPHeaderProperty{
//   				Name: jsii.String("name"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-captchaaction.html
//
type CfnWebACL_CaptchaActionProperty struct {
	// Defines custom handling for the web request, used when the `CAPTCHA` inspection determines that the request's token is valid and unexpired.
	//
	// For information about customizing web requests and responses, see [Customizing web requests and responses in AWS WAF](https://docs.aws.amazon.com/waf/latest/developerguide/waf-custom-request-response.html) in the *AWS WAF Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-captchaaction.html#cfn-wafv2-webacl-captchaaction-customrequesthandling
	//
	CustomRequestHandling interface{} `field:"optional" json:"customRequestHandling" yaml:"customRequestHandling"`
}

