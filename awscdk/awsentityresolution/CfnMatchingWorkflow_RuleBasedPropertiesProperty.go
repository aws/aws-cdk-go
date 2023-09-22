package awsentityresolution


// An object which defines the list of matching rules to run and has a field `Rules` , which is a list of rule objects.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ruleBasedPropertiesProperty := &RuleBasedPropertiesProperty{
//   	AttributeMatchingModel: jsii.String("attributeMatchingModel"),
//   	Rules: []interface{}{
//   		&RuleProperty{
//   			MatchingKeys: []*string{
//   				jsii.String("matchingKeys"),
//   			},
//   			RuleName: jsii.String("ruleName"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-rulebasedproperties.html
//
type CfnMatchingWorkflow_RuleBasedPropertiesProperty struct {
	// The comparison type.
	//
	// You can either choose `ONE_TO_ONE` or `MANY_TO_MANY` as the AttributeMatchingModel. When choosing `MANY_TO_MANY` , the system can match attributes across the sub-types of an attribute type. For example, if the value of the `Email` field of Profile A and the value of `BusinessEmail` field of Profile B matches, the two profiles are matched on the `Email` type. When choosing `ONE_TO_ONE` ,the system can only match if the sub-types are exact matches. For example, only when the value of the `Email` field of Profile A and the value of the `Email` field of Profile B matches, the two profiles are matched on the `Email` type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-rulebasedproperties.html#cfn-entityresolution-matchingworkflow-rulebasedproperties-attributematchingmodel
	//
	AttributeMatchingModel *string `field:"required" json:"attributeMatchingModel" yaml:"attributeMatchingModel"`
	// A list of `Rule` objects, each of which have fields `RuleName` and `MatchingKeys` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-rulebasedproperties.html#cfn-entityresolution-matchingworkflow-rulebasedproperties-rules
	//
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
}

