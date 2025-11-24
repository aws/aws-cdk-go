package mixinsawsentityresolution


// An object which defines the ID mapping technique and any additional configurations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   idMappingTechniquesProperty := &IdMappingTechniquesProperty{
//   	IdMappingType: jsii.String("idMappingType"),
//   	NormalizationVersion: jsii.String("normalizationVersion"),
//   	ProviderProperties: &ProviderPropertiesProperty{
//   		IntermediateSourceConfiguration: &IntermediateSourceConfigurationProperty{
//   			IntermediateS3Path: jsii.String("intermediateS3Path"),
//   		},
//   		ProviderConfiguration: map[string]*string{
//   			"providerConfigurationKey": jsii.String("providerConfiguration"),
//   		},
//   		ProviderServiceArn: jsii.String("providerServiceArn"),
//   	},
//   	RuleBasedProperties: &IdMappingRuleBasedPropertiesProperty{
//   		AttributeMatchingModel: jsii.String("attributeMatchingModel"),
//   		RecordMatchingModel: jsii.String("recordMatchingModel"),
//   		RuleDefinitionType: jsii.String("ruleDefinitionType"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idmappingworkflow-idmappingtechniques.html
//
type CfnIdMappingWorkflowPropsMixin_IdMappingTechniquesProperty struct {
	// The type of ID mapping.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idmappingworkflow-idmappingtechniques.html#cfn-entityresolution-idmappingworkflow-idmappingtechniques-idmappingtype
	//
	IdMappingType *string `field:"optional" json:"idMappingType" yaml:"idMappingType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idmappingworkflow-idmappingtechniques.html#cfn-entityresolution-idmappingworkflow-idmappingtechniques-normalizationversion
	//
	NormalizationVersion *string `field:"optional" json:"normalizationVersion" yaml:"normalizationVersion"`
	// An object which defines any additional configurations required by the provider service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idmappingworkflow-idmappingtechniques.html#cfn-entityresolution-idmappingworkflow-idmappingtechniques-providerproperties
	//
	ProviderProperties interface{} `field:"optional" json:"providerProperties" yaml:"providerProperties"`
	// An object which defines any additional configurations required by rule-based matching.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idmappingworkflow-idmappingtechniques.html#cfn-entityresolution-idmappingworkflow-idmappingtechniques-rulebasedproperties
	//
	RuleBasedProperties interface{} `field:"optional" json:"ruleBasedProperties" yaml:"ruleBasedProperties"`
}

