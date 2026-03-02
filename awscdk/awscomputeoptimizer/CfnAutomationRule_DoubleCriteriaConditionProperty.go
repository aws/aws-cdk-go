package awscomputeoptimizer


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   doubleCriteriaConditionProperty := &DoubleCriteriaConditionProperty{
//   	Comparison: jsii.String("comparison"),
//   	Values: []interface{}{
//   		jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-computeoptimizer-automationrule-doublecriteriacondition.html
//
type CfnAutomationRule_DoubleCriteriaConditionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-computeoptimizer-automationrule-doublecriteriacondition.html#cfn-computeoptimizer-automationrule-doublecriteriacondition-comparison
	//
	Comparison *string `field:"optional" json:"comparison" yaml:"comparison"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-computeoptimizer-automationrule-doublecriteriacondition.html#cfn-computeoptimizer-automationrule-doublecriteriacondition-values
	//
	Values interface{} `field:"optional" json:"values" yaml:"values"`
}

