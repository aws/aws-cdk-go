package awscomputeoptimizer


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stringCriteriaConditionProperty := &StringCriteriaConditionProperty{
//   	Comparison: jsii.String("comparison"),
//   	Values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-computeoptimizer-automationrule-stringcriteriacondition.html
//
type CfnAutomationRule_StringCriteriaConditionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-computeoptimizer-automationrule-stringcriteriacondition.html#cfn-computeoptimizer-automationrule-stringcriteriacondition-comparison
	//
	Comparison *string `field:"optional" json:"comparison" yaml:"comparison"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-computeoptimizer-automationrule-stringcriteriacondition.html#cfn-computeoptimizer-automationrule-stringcriteriacondition-values
	//
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

