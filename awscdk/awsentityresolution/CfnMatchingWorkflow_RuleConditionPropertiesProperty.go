package awsentityresolution


// The properties of a rule condition that provides the ability to use more complex syntax.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ruleConditionPropertiesProperty := &RuleConditionPropertiesProperty{
//   	Rules: []interface{}{
//   		&RuleConditionProperty{
//   			Condition: jsii.String("condition"),
//   			RuleName: jsii.String("ruleName"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-ruleconditionproperties.html
//
type CfnMatchingWorkflow_RuleConditionPropertiesProperty struct {
	// A list of rule objects, each of which have fields `ruleName` and `condition` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-ruleconditionproperties.html#cfn-entityresolution-matchingworkflow-ruleconditionproperties-rules
	//
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
}

