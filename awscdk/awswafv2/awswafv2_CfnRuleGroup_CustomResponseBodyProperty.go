package awswafv2


// The response body to use in a custom response to a web request.
//
// This is referenced by key from `CustomResponse` `CustomResponseBodyKey` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customResponseBodyProperty := &customResponseBodyProperty{
//   	content: jsii.String("content"),
//   	contentType: jsii.String("contentType"),
//   }
//
type CfnRuleGroup_CustomResponseBodyProperty struct {
	// The payload of the custom response.
	//
	// You can use JSON escape strings in JSON content. To do this, you must specify JSON content in the `ContentType` setting.
	//
	// For information about the limits on count and size for custom request and response settings, see [AWS WAF quotas](https://docs.aws.amazon.com/waf/latest/developerguide/limits.html) in the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) .
	Content *string `field:"required" json:"content" yaml:"content"`
	// The type of content in the payload that you are defining in the `Content` string.
	ContentType *string `field:"required" json:"contentType" yaml:"contentType"`
}

