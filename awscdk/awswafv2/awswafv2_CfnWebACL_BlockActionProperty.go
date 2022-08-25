package awswafv2


// Specifies that AWS WAF should block the request and optionally defines additional custom handling for the response to the web request.
//
// This is used in the context of other settings, for example to specify values for a rule action or a web ACL default action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   blockActionProperty := &blockActionProperty{
//   	customResponse: &customResponseProperty{
//   		responseCode: jsii.Number(123),
//
//   		// the properties below are optional
//   		customResponseBodyKey: jsii.String("customResponseBodyKey"),
//   		responseHeaders: []interface{}{
//   			&customHTTPHeaderProperty{
//   				name: jsii.String("name"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   	},
//   }
//
type CfnWebACL_BlockActionProperty struct {
	// Defines a custom response for the web request.
	//
	// For information about customizing web requests and responses, see [Customizing web requests and responses in AWS WAF](https://docs.aws.amazon.com/waf/latest/developerguide/waf-custom-request-response.html) in the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) .
	CustomResponse interface{} `field:"optional" json:"customResponse" yaml:"customResponse"`
}

