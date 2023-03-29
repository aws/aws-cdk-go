package awswafv2


// In a `WebACL` , this is the action that you want AWS WAF to perform when a web request doesn't match any of the rules in the `WebACL` .
//
// The default action must be a terminating action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   defaultActionProperty := &DefaultActionProperty{
//   	Allow: &AllowActionProperty{
//   		CustomRequestHandling: &CustomRequestHandlingProperty{
//   			InsertHeaders: []interface{}{
//   				&CustomHTTPHeaderProperty{
//   					Name: jsii.String("name"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	Block: &BlockActionProperty{
//   		CustomResponse: &CustomResponseProperty{
//   			ResponseCode: jsii.Number(123),
//
//   			// the properties below are optional
//   			CustomResponseBodyKey: jsii.String("customResponseBodyKey"),
//   			ResponseHeaders: []interface{}{
//   				&CustomHTTPHeaderProperty{
//   					Name: jsii.String("name"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnWebACL_DefaultActionProperty struct {
	// Specifies that AWS WAF should allow requests by default.
	Allow interface{} `field:"optional" json:"allow" yaml:"allow"`
	// Specifies that AWS WAF should block requests by default.
	Block interface{} `field:"optional" json:"block" yaml:"block"`
}

