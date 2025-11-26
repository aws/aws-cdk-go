package previewawsentityresolutionmixins


// An object containing `idMappingType` , `providerProperties` , and `ruleBasedProperties` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   idNamespaceIdMappingWorkflowPropertiesProperty := &IdNamespaceIdMappingWorkflowPropertiesProperty{
//   	IdMappingType: jsii.String("idMappingType"),
//   	ProviderProperties: &NamespaceProviderPropertiesProperty{
//   		ProviderConfiguration: map[string]*string{
//   			"providerConfigurationKey": jsii.String("providerConfiguration"),
//   		},
//   		ProviderServiceArn: jsii.String("providerServiceArn"),
//   	},
//   	RuleBasedProperties: &NamespaceRuleBasedPropertiesProperty{
//   		AttributeMatchingModel: jsii.String("attributeMatchingModel"),
//   		RecordMatchingModels: []*string{
//   			jsii.String("recordMatchingModels"),
//   		},
//   		RuleDefinitionTypes: []*string{
//   			jsii.String("ruleDefinitionTypes"),
//   		},
//   		Rules: []interface{}{
//   			&RuleProperty{
//   				MatchingKeys: []*string{
//   					jsii.String("matchingKeys"),
//   				},
//   				RuleName: jsii.String("ruleName"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idnamespace-idnamespaceidmappingworkflowproperties.html
//
type CfnIdNamespacePropsMixin_IdNamespaceIdMappingWorkflowPropertiesProperty struct {
	// The type of ID mapping.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idnamespace-idnamespaceidmappingworkflowproperties.html#cfn-entityresolution-idnamespace-idnamespaceidmappingworkflowproperties-idmappingtype
	//
	IdMappingType *string `field:"optional" json:"idMappingType" yaml:"idMappingType"`
	// An object which defines any additional configurations required by the provider service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idnamespace-idnamespaceidmappingworkflowproperties.html#cfn-entityresolution-idnamespace-idnamespaceidmappingworkflowproperties-providerproperties
	//
	ProviderProperties interface{} `field:"optional" json:"providerProperties" yaml:"providerProperties"`
	// An object which defines any additional configurations required by rule-based matching.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idnamespace-idnamespaceidmappingworkflowproperties.html#cfn-entityresolution-idnamespace-idnamespaceidmappingworkflowproperties-rulebasedproperties
	//
	RuleBasedProperties interface{} `field:"optional" json:"ruleBasedProperties" yaml:"ruleBasedProperties"`
}

