package awsentityresolution


// An object containing `RuleName` , and `MatchingKeys` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ruleProperty := &RuleProperty{
//   	MatchingKeys: []*string{
//   		jsii.String("matchingKeys"),
//   	},
//   	RuleName: jsii.String("ruleName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-rule.html
//
type CfnMatchingWorkflow_RuleProperty struct {
	// A list of `MatchingKeys` .
	//
	// The `MatchingKeys` must have been defined in the `SchemaMapping` . Two records are considered to match according to this rule if all of the `MatchingKeys` match.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-rule.html#cfn-entityresolution-matchingworkflow-rule-matchingkeys
	//
	MatchingKeys *[]*string `field:"required" json:"matchingKeys" yaml:"matchingKeys"`
	// A name for the matching rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-rule.html#cfn-entityresolution-matchingworkflow-rule-rulename
	//
	RuleName *string `field:"required" json:"ruleName" yaml:"ruleName"`
}

