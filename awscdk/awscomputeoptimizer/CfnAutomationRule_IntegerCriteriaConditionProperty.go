package awscomputeoptimizer


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   integerCriteriaConditionProperty := &IntegerCriteriaConditionProperty{
//   	Comparison: jsii.String("comparison"),
//   	Values: []interface{}{
//   		jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-computeoptimizer-automationrule-integercriteriacondition.html
//
type CfnAutomationRule_IntegerCriteriaConditionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-computeoptimizer-automationrule-integercriteriacondition.html#cfn-computeoptimizer-automationrule-integercriteriacondition-comparison
	//
	Comparison *string `field:"optional" json:"comparison" yaml:"comparison"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-computeoptimizer-automationrule-integercriteriacondition.html#cfn-computeoptimizer-automationrule-integercriteriacondition-values
	//
	Values interface{} `field:"optional" json:"values" yaml:"values"`
}

