package mixinsawsentityresolution


// An object which defines the list of matching rules to run in a matching workflow.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ruleBasedPropertiesProperty := &RuleBasedPropertiesProperty{
//   	AttributeMatchingModel: jsii.String("attributeMatchingModel"),
//   	MatchPurpose: jsii.String("matchPurpose"),
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
type CfnMatchingWorkflowPropsMixin_RuleBasedPropertiesProperty struct {
	// The comparison type. You can choose `ONE_TO_ONE` or `MANY_TO_MANY` as the `attributeMatchingModel` .
	//
	// If you choose `ONE_TO_ONE` , the system can only match attributes if the sub-types are an exact match. For example, for the `Email` attribute type, the system will only consider it a match if the value of the `Email` field of Profile A matches the value of the `Email` field of Profile B.
	//
	// If you choose `MANY_TO_MANY` , the system can match attributes across the sub-types of an attribute type. For example, if the value of the `Email` field of Profile A and the value of `BusinessEmail` field of Profile B matches, the two profiles are matched on the `Email` attribute type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-rulebasedproperties.html#cfn-entityresolution-matchingworkflow-rulebasedproperties-attributematchingmodel
	//
	AttributeMatchingModel *string `field:"optional" json:"attributeMatchingModel" yaml:"attributeMatchingModel"`
	// An indicator of whether to generate IDs and index the data or not.
	//
	// If you choose `IDENTIFIER_GENERATION` , the process generates IDs and indexes the data.
	//
	// If you choose `INDEXING` , the process indexes the data without generating IDs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-rulebasedproperties.html#cfn-entityresolution-matchingworkflow-rulebasedproperties-matchpurpose
	//
	MatchPurpose *string `field:"optional" json:"matchPurpose" yaml:"matchPurpose"`
	// A list of `Rule` objects, each of which have fields `RuleName` and `MatchingKeys` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-rulebasedproperties.html#cfn-entityresolution-matchingworkflow-rulebasedproperties-rules
	//
	Rules interface{} `field:"optional" json:"rules" yaml:"rules"`
}

