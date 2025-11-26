package previewawswafv2mixins


// In a `WebACL` , this is the action that you want AWS WAF to perform when a web request doesn't match any of the rules in the `WebACL` .
//
// The default action must be a terminating action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
//   			CustomResponseBodyKey: jsii.String("customResponseBodyKey"),
//   			ResponseCode: jsii.Number(123),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-defaultaction.html
//
type CfnWebACLPropsMixin_DefaultActionProperty struct {
	// Specifies that AWS WAF should allow requests by default.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-defaultaction.html#cfn-wafv2-webacl-defaultaction-allow
	//
	Allow interface{} `field:"optional" json:"allow" yaml:"allow"`
	// Specifies that AWS WAF should block requests by default.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-defaultaction.html#cfn-wafv2-webacl-defaultaction-block
	//
	Block interface{} `field:"optional" json:"block" yaml:"block"`
}

