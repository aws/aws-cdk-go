package awsentityresolution


// The properties of a rule condition that provides the ability to use more complex syntax.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   ruleConditionPropertiesProperty := &RuleConditionPropertiesProperty{
//   	MatchingConfig: &MatchingConfigProperty{
//   		EnableTransitiveMatching: jsii.Boolean(false),
//   	},
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
type CfnMatchingWorkflowPropsMixin_RuleConditionPropertiesProperty struct {
	// Configuration for matching behavior within rule condition properties.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-ruleconditionproperties.html#cfn-entityresolution-matchingworkflow-ruleconditionproperties-matchingconfig
	//
	MatchingConfig interface{} `field:"optional" json:"matchingConfig" yaml:"matchingConfig"`
	// A list of rule objects, each of which have fields `ruleName` and `condition` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-ruleconditionproperties.html#cfn-entityresolution-matchingworkflow-ruleconditionproperties-rules
	//
	Rules interface{} `field:"optional" json:"rules" yaml:"rules"`
}

