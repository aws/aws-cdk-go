package awswafv2


// Specifies that AWS WAF should allow the request and optionally defines additional custom handling for the request.
//
// This is used in the context of other settings, for example to specify values for a rule action or a web ACL default action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   allowActionProperty := &AllowActionProperty{
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
type CfnWebACL_AllowActionProperty struct {
	// Defines custom handling for the web request.
	//
	// For information about customizing web requests and responses, see [Customizing web requests and responses in AWS WAF](https://docs.aws.amazon.com/waf/latest/developerguide/waf-custom-request-response.html) in the *AWS WAF Developer Guide* .
	CustomRequestHandling interface{} `field:"optional" json:"customRequestHandling" yaml:"customRequestHandling"`
}

