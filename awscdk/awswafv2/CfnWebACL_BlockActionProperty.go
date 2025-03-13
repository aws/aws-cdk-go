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
//   blockActionProperty := &BlockActionProperty{
//   	CustomResponse: &CustomResponseProperty{
//   		ResponseCode: jsii.Number(123),
//
//   		// the properties below are optional
//   		CustomResponseBodyKey: jsii.String("customResponseBodyKey"),
//   		ResponseHeaders: []interface{}{
//   			&CustomHTTPHeaderProperty{
//   				Name: jsii.String("name"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-blockaction.html
//
type CfnWebACL_BlockActionProperty struct {
	// Defines a custom response for the web request.
	//
	// For information about customizing web requests and responses, see [Customizing web requests and responses in AWS WAF](https://docs.aws.amazon.com/waf/latest/developerguide/waf-custom-request-response.html) in the *AWS WAF Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-blockaction.html#cfn-wafv2-webacl-blockaction-customresponse
	//
	CustomResponse interface{} `field:"optional" json:"customResponse" yaml:"customResponse"`
}

