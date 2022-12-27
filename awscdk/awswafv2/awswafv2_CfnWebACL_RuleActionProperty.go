package awswafv2


// The action that AWS WAF should take on a web request when it matches a rule's statement.
//
// Settings at the web ACL level can override the rule action setting.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ruleActionProperty := &ruleActionProperty{
//   	allow: &allowActionProperty{
//   		customRequestHandling: &customRequestHandlingProperty{
//   			insertHeaders: []interface{}{
//   				&customHTTPHeaderProperty{
//   					name: jsii.String("name"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	block: &blockActionProperty{
//   		customResponse: &customResponseProperty{
//   			responseCode: jsii.Number(123),
//
//   			// the properties below are optional
//   			customResponseBodyKey: jsii.String("customResponseBodyKey"),
//   			responseHeaders: []interface{}{
//   				&customHTTPHeaderProperty{
//   					name: jsii.String("name"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	captcha: &captchaActionProperty{
//   		customRequestHandling: &customRequestHandlingProperty{
//   			insertHeaders: []interface{}{
//   				&customHTTPHeaderProperty{
//   					name: jsii.String("name"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	challenge: &challengeActionProperty{
//   		customRequestHandling: &customRequestHandlingProperty{
//   			insertHeaders: []interface{}{
//   				&customHTTPHeaderProperty{
//   					name: jsii.String("name"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	count: &countActionProperty{
//   		customRequestHandling: &customRequestHandlingProperty{
//   			insertHeaders: []interface{}{
//   				&customHTTPHeaderProperty{
//   					name: jsii.String("name"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnWebACL_RuleActionProperty struct {
	// Instructs AWS WAF to allow the web request.
	Allow interface{} `field:"optional" json:"allow" yaml:"allow"`
	// Instructs AWS WAF to block the web request.
	Block interface{} `field:"optional" json:"block" yaml:"block"`
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
	Captcha interface{} `field:"optional" json:"captcha" yaml:"captcha"`
	// `CfnWebACL.RuleActionProperty.Challenge`.
	Challenge interface{} `field:"optional" json:"challenge" yaml:"challenge"`
	// Instructs AWS WAF to count the web request and allow it.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
}

