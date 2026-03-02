package awscomputeoptimizer


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceTagsCriteriaConditionProperty := &ResourceTagsCriteriaConditionProperty{
//   	Comparison: jsii.String("comparison"),
//   	Key: jsii.String("key"),
//   	Values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-computeoptimizer-automationrule-resourcetagscriteriacondition.html
//
type CfnAutomationRule_ResourceTagsCriteriaConditionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-computeoptimizer-automationrule-resourcetagscriteriacondition.html#cfn-computeoptimizer-automationrule-resourcetagscriteriacondition-comparison
	//
	Comparison *string `field:"optional" json:"comparison" yaml:"comparison"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-computeoptimizer-automationrule-resourcetagscriteriacondition.html#cfn-computeoptimizer-automationrule-resourcetagscriteriacondition-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-computeoptimizer-automationrule-resourcetagscriteriacondition.html#cfn-computeoptimizer-automationrule-resourcetagscriteriacondition-values
	//
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

