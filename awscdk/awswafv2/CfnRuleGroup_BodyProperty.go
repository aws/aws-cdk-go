package awswafv2


// Inspect the body of the web request. The body immediately follows the request headers.
//
// This is used to indicate the web request component to inspect, in the `FieldToMatch` specification.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   bodyProperty := &BodyProperty{
//   	OversizeHandling: jsii.String("oversizeHandling"),
//   }
//
type CfnRuleGroup_BodyProperty struct {
	// What AWS WAF should do if the body is larger than AWS WAF can inspect.
	//
	// AWS WAF does not support inspecting the entire contents of the web request body if the body exceeds the limit for the resource type. If the body is larger than the limit, the underlying host service only forwards the contents that are below the limit to AWS WAF for inspection.
	//
	// The default limit is 8 KB (8,192 kilobytes) for regional resources and 16 KB (16,384 kilobytes) for CloudFront distributions. For CloudFront distributions, you can increase the limit in the web ACL `AssociationConfig` , for additional processing fees.
	//
	// The options for oversize handling are the following:
	//
	// - `CONTINUE` - Inspect the available body contents normally, according to the rule inspection criteria.
	// - `MATCH` - Treat the web request as matching the rule statement. AWS WAF applies the rule action to the request.
	// - `NO_MATCH` - Treat the web request as not matching the rule statement.
	//
	// You can combine the `MATCH` or `NO_MATCH` settings for oversize handling with your rule and web ACL action settings, so that you block any request whose body is over the limit.
	//
	// Default: `CONTINUE`.
	OversizeHandling *string `field:"optional" json:"oversizeHandling" yaml:"oversizeHandling"`
}

