package awsentityresolution


// The rule-based properties of an ID namespace.
//
// These properties define how the ID namespace can be used in an ID mapping workflow.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   namespaceRuleBasedPropertiesProperty := &NamespaceRuleBasedPropertiesProperty{
//   	AttributeMatchingModel: jsii.String("attributeMatchingModel"),
//   	RecordMatchingModels: []*string{
//   		jsii.String("recordMatchingModels"),
//   	},
//   	RuleDefinitionTypes: []*string{
//   		jsii.String("ruleDefinitionTypes"),
//   	},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idnamespace-namespacerulebasedproperties.html
//
type CfnIdNamespace_NamespaceRuleBasedPropertiesProperty struct {
	// The comparison type. You can either choose `ONE_TO_ONE` or `MANY_TO_MANY` as the `attributeMatchingModel` .
	//
	// If you choose `MANY_TO_MANY` , the system can match attributes across the sub-types of an attribute type. For example, if the value of the `Email` field of Profile A matches the value of `BusinessEmail` field of Profile B, the two profiles are matched on the `Email` attribute type.
	//
	// If you choose `ONE_TO_ONE` , the system can only match attributes if the sub-types are an exact match. For example, for the `Email` attribute type, the system will only consider it a match if the value of the `Email` field of Profile A matches the value of the `Email` field of Profile B.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idnamespace-namespacerulebasedproperties.html#cfn-entityresolution-idnamespace-namespacerulebasedproperties-attributematchingmodel
	//
	AttributeMatchingModel *string `field:"optional" json:"attributeMatchingModel" yaml:"attributeMatchingModel"`
	// The type of matching record that is allowed to be used in an ID mapping workflow.
	//
	// If the value is set to `ONE_SOURCE_TO_ONE_TARGET` , only one record in the source is matched to one record in the target.
	//
	// If the value is set to `MANY_SOURCE_TO_ONE_TARGET` , all matching records in the source are matched to one record in the target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idnamespace-namespacerulebasedproperties.html#cfn-entityresolution-idnamespace-namespacerulebasedproperties-recordmatchingmodels
	//
	RecordMatchingModels *[]*string `field:"optional" json:"recordMatchingModels" yaml:"recordMatchingModels"`
	// The sets of rules you can use in an ID mapping workflow.
	//
	// The limitations specified for the source and target must be compatible.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idnamespace-namespacerulebasedproperties.html#cfn-entityresolution-idnamespace-namespacerulebasedproperties-ruledefinitiontypes
	//
	RuleDefinitionTypes *[]*string `field:"optional" json:"ruleDefinitionTypes" yaml:"ruleDefinitionTypes"`
	// The rules for the ID namespace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idnamespace-namespacerulebasedproperties.html#cfn-entityresolution-idnamespace-namespacerulebasedproperties-rules
	//
	Rules interface{} `field:"optional" json:"rules" yaml:"rules"`
}

