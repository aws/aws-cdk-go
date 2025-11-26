package previewawswafv2mixins


// Custom request handling behavior that inserts custom headers into a web request.
//
// You can add custom request handling for AWS WAF to use when the rule action doesn't block the request. For example, `CaptchaAction` for requests with valid t okens, and `AllowAction` .
//
// For information about customizing web requests and responses, see [Customizing web requests and responses in AWS WAF](https://docs.aws.amazon.com/waf/latest/developerguide/waf-custom-request-response.html) in the *AWS WAF Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   customRequestHandlingProperty := &CustomRequestHandlingProperty{
//   	InsertHeaders: []interface{}{
//   		&CustomHTTPHeaderProperty{
//   			Name: jsii.String("name"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-customrequesthandling.html
//
type CfnWebACLPropsMixin_CustomRequestHandlingProperty struct {
	// The HTTP headers to insert into the request. Duplicate header names are not allowed.
	//
	// For information about the limits on count and size for custom request and response settings, see [AWS WAF quotas](https://docs.aws.amazon.com/waf/latest/developerguide/limits.html) in the *AWS WAF Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-customrequesthandling.html#cfn-wafv2-webacl-customrequesthandling-insertheaders
	//
	InsertHeaders interface{} `field:"optional" json:"insertHeaders" yaml:"insertHeaders"`
}

