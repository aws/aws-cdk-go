package awsentityresolution


// An object that defines the list of matching rules to run in an ID mapping workflow.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   idMappingRuleBasedPropertiesProperty := &IdMappingRuleBasedPropertiesProperty{
//   	AttributeMatchingModel: jsii.String("attributeMatchingModel"),
//   	RecordMatchingModel: jsii.String("recordMatchingModel"),
//
//   	// the properties below are optional
//   	RuleDefinitionType: jsii.String("ruleDefinitionType"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idmappingworkflow-idmappingrulebasedproperties.html
//
type CfnIdMappingWorkflow_IdMappingRuleBasedPropertiesProperty struct {
	// The comparison type. You can either choose `ONE_TO_ONE` or `MANY_TO_MANY` as the `attributeMatchingModel` .
	//
	// If you choose `MANY_TO_MANY` , the system can match attributes across the sub-types of an attribute type. For example, if the value of the `Email` field of Profile A matches the value of the `BusinessEmail` field of Profile B, the two profiles are matched on the `Email` attribute type.
	//
	// If you choose `ONE_TO_ONE` , the system can only match attributes if the sub-types are an exact match. For example, for the `Email` attribute type, the system will only consider it a match if the value of the `Email` field of Profile A matches the value of the `Email` field of Profile B.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idmappingworkflow-idmappingrulebasedproperties.html#cfn-entityresolution-idmappingworkflow-idmappingrulebasedproperties-attributematchingmodel
	//
	AttributeMatchingModel *string `field:"required" json:"attributeMatchingModel" yaml:"attributeMatchingModel"`
	// The type of matching record that is allowed to be used in an ID mapping workflow.
	//
	// If the value is set to `ONE_SOURCE_TO_ONE_TARGET` , only one record in the source can be matched to the same record in the target.
	//
	// If the value is set to `MANY_SOURCE_TO_ONE_TARGET` , multiple records in the source can be matched to one record in the target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idmappingworkflow-idmappingrulebasedproperties.html#cfn-entityresolution-idmappingworkflow-idmappingrulebasedproperties-recordmatchingmodel
	//
	RecordMatchingModel *string `field:"required" json:"recordMatchingModel" yaml:"recordMatchingModel"`
	// The set of rules you can use in an ID mapping workflow.
	//
	// The limitations specified for the source or target to define the match rules must be compatible.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idmappingworkflow-idmappingrulebasedproperties.html#cfn-entityresolution-idmappingworkflow-idmappingrulebasedproperties-ruledefinitiontype
	//
	RuleDefinitionType *string `field:"optional" json:"ruleDefinitionType" yaml:"ruleDefinitionType"`
	// The rules that can be used for ID mapping.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idmappingworkflow-idmappingrulebasedproperties.html#cfn-entityresolution-idmappingworkflow-idmappingrulebasedproperties-rules
	//
	Rules interface{} `field:"optional" json:"rules" yaml:"rules"`
}

