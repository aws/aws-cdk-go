package awswafregional


// Properties for defining a `CfnWebACL`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnWebACLProps := &CfnWebACLProps{
//   	DefaultAction: &ActionProperty{
//   		Type: jsii.String("type"),
//   	},
//   	MetricName: jsii.String("metricName"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Rules: []interface{}{
//   		&RuleProperty{
//   			Action: &ActionProperty{
//   				Type: jsii.String("type"),
//   			},
//   			Priority: jsii.Number(123),
//   			RuleId: jsii.String("ruleId"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-webacl.html
//
type CfnWebACLProps struct {
	// The action to perform if none of the `Rules` contained in the `WebACL` match.
	//
	// The action is specified by the `WafAction` object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-webacl.html#cfn-wafregional-webacl-defaultaction
	//
	DefaultAction interface{} `field:"required" json:"defaultAction" yaml:"defaultAction"`
	// A name for the metrics for this `WebACL` .
	//
	// The name can contain only alphanumeric characters (A-Z, a-z, 0-9), with maximum length 128 and minimum length one. It can't contain whitespace or metric names reserved for AWS WAF, including "All" and "Default_Action." You can't change `MetricName` after you create the `WebACL` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-webacl.html#cfn-wafregional-webacl-metricname
	//
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// A friendly name or description of the `WebACL` .
	//
	// You can't change the name of a `WebACL` after you create it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-webacl.html#cfn-wafregional-webacl-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// An array that contains the action for each `Rule` in a `WebACL` , the priority of the `Rule` , and the ID of the `Rule` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-webacl.html#cfn-wafregional-webacl-rules
	//
	Rules interface{} `field:"optional" json:"rules" yaml:"rules"`
}

