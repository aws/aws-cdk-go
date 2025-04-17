package awswafv2


// Action setting to use in the place of a rule action that is configured inside the rule group.
//
// You specify one override for each rule whose action you want to change.
//
// You can use overrides for testing, for example you can override all of rule actions to `Count` and then monitor the resulting count metrics to understand how the rule group would handle your web traffic. You can also permanently override some or all actions, to modify how the rule group manages your web traffic.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ruleActionOverrideProperty := &RuleActionOverrideProperty{
//   	ActionToUse: &RuleActionProperty{
//   		Allow: &AllowActionProperty{
//   			CustomRequestHandling: &CustomRequestHandlingProperty{
//   				InsertHeaders: []interface{}{
//   					&CustomHTTPHeaderProperty{
//   						Name: jsii.String("name"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   		Block: &BlockActionProperty{
//   			CustomResponse: &CustomResponseProperty{
//   				ResponseCode: jsii.Number(123),
//
//   				// the properties below are optional
//   				CustomResponseBodyKey: jsii.String("customResponseBodyKey"),
//   				ResponseHeaders: []interface{}{
//   					&CustomHTTPHeaderProperty{
//   						Name: jsii.String("name"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   		Captcha: &CaptchaActionProperty{
//   			CustomRequestHandling: &CustomRequestHandlingProperty{
//   				InsertHeaders: []interface{}{
//   					&CustomHTTPHeaderProperty{
//   						Name: jsii.String("name"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   		Challenge: &ChallengeActionProperty{
//   			CustomRequestHandling: &CustomRequestHandlingProperty{
//   				InsertHeaders: []interface{}{
//   					&CustomHTTPHeaderProperty{
//   						Name: jsii.String("name"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   		Count: &CountActionProperty{
//   			CustomRequestHandling: &CustomRequestHandlingProperty{
//   				InsertHeaders: []interface{}{
//   					&CustomHTTPHeaderProperty{
//   						Name: jsii.String("name"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-ruleactionoverride.html
//
type CfnWebACL_RuleActionOverrideProperty struct {
	// The override action to use, in place of the configured action of the rule in the rule group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-ruleactionoverride.html#cfn-wafv2-webacl-ruleactionoverride-actiontouse
	//
	ActionToUse interface{} `field:"required" json:"actionToUse" yaml:"actionToUse"`
	// The name of the rule to override.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-ruleactionoverride.html#cfn-wafv2-webacl-ruleactionoverride-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
}

