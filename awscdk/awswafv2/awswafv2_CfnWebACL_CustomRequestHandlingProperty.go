package awswafv2


// Custom request handling behavior that inserts custom headers into a web request.
//
// You can add custom request handling for the rule actions allow and count.
//
// For information about customizing web requests and responses, see [Customizing web requests and responses in AWS WAF](https://docs.aws.amazon.com/waf/latest/developerguide/waf-custom-request-response.html) in the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customRequestHandlingProperty := &customRequestHandlingProperty{
//   	insertHeaders: []interface{}{
//   		&customHTTPHeaderProperty{
//   			name: jsii.String("name"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnWebACL_CustomRequestHandlingProperty struct {
	// The HTTP headers to insert into the request. Duplicate header names are not allowed.
	//
	// For information about the limits on count and size for custom request and response settings, see [AWS WAF quotas](https://docs.aws.amazon.com/waf/latest/developerguide/limits.html) in the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) .
	InsertHeaders interface{} `field:"required" json:"insertHeaders" yaml:"insertHeaders"`
}

