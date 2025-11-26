package previewawswafmixins


// Properties for CfnWebACLPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnWebACLMixinProps := &CfnWebACLMixinProps{
//   	DefaultAction: &WafActionProperty{
//   		Type: jsii.String("type"),
//   	},
//   	MetricName: jsii.String("metricName"),
//   	Name: jsii.String("name"),
//   	Rules: []interface{}{
//   		&ActivatedRuleProperty{
//   			Action: &WafActionProperty{
//   				Type: jsii.String("type"),
//   			},
//   			Priority: jsii.Number(123),
//   			RuleId: jsii.String("ruleId"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-webacl.html
//
type CfnWebACLMixinProps struct {
	// The action to perform if none of the `Rules` contained in the `WebACL` match.
	//
	// The action is specified by the `WafAction` object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-webacl.html#cfn-waf-webacl-defaultaction
	//
	DefaultAction interface{} `field:"optional" json:"defaultAction" yaml:"defaultAction"`
	// The name of the metrics for this `WebACL` .
	//
	// The name can contain only alphanumeric characters (A-Z, a-z, 0-9), with maximum length 128 and minimum length one. It can't contain whitespace or metric names reserved for AWS WAF , including "All" and "Default_Action." You can't change `MetricName` after you create the `WebACL` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-webacl.html#cfn-waf-webacl-metricname
	//
	MetricName *string `field:"optional" json:"metricName" yaml:"metricName"`
	// A friendly name or description of the `WebACL` .
	//
	// You can't change the name of a `WebACL` after you create it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-webacl.html#cfn-waf-webacl-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An array that contains the action for each `Rule` in a `WebACL` , the priority of the `Rule` , and the ID of the `Rule` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-webacl.html#cfn-waf-webacl-rules
	//
	Rules interface{} `field:"optional" json:"rules" yaml:"rules"`
}

