package awsentityresolution


// An object which defines the `resolutionType` and the `ruleBasedProperties` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resolutionTechniquesProperty := &ResolutionTechniquesProperty{
//   	ProviderProperties: &ProviderPropertiesProperty{
//   		ProviderServiceArn: jsii.String("providerServiceArn"),
//
//   		// the properties below are optional
//   		IntermediateSourceConfiguration: &IntermediateSourceConfigurationProperty{
//   			IntermediateS3Path: jsii.String("intermediateS3Path"),
//   		},
//   		ProviderConfiguration: map[string]*string{
//   			"providerConfigurationKey": jsii.String("providerConfiguration"),
//   		},
//   	},
//   	ResolutionType: jsii.String("resolutionType"),
//   	RuleBasedProperties: &RuleBasedPropertiesProperty{
//   		AttributeMatchingModel: jsii.String("attributeMatchingModel"),
//   		Rules: []interface{}{
//   			&RuleProperty{
//   				MatchingKeys: []*string{
//   					jsii.String("matchingKeys"),
//   				},
//   				RuleName: jsii.String("ruleName"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		MatchPurpose: jsii.String("matchPurpose"),
//   	},
//   	RuleConditionProperties: &RuleConditionPropertiesProperty{
//   		Rules: []interface{}{
//   			&RuleConditionProperty{
//   				Condition: jsii.String("condition"),
//   				RuleName: jsii.String("ruleName"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-resolutiontechniques.html
//
type CfnMatchingWorkflow_ResolutionTechniquesProperty struct {
	// The properties of the provider service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-resolutiontechniques.html#cfn-entityresolution-matchingworkflow-resolutiontechniques-providerproperties
	//
	ProviderProperties interface{} `field:"optional" json:"providerProperties" yaml:"providerProperties"`
	// The type of matching workflow to create. Specify one of the following types:.
	//
	// - `RULE_MATCHING` : Match records using configurable rule-based criteria
	// - `ML_MATCHING` : Match records using machine learning models
	// - `PROVIDER` : Match records using a third-party matching provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-resolutiontechniques.html#cfn-entityresolution-matchingworkflow-resolutiontechniques-resolutiontype
	//
	ResolutionType *string `field:"optional" json:"resolutionType" yaml:"resolutionType"`
	// An object which defines the list of matching rules to run and has a field `rules` , which is a list of rule objects.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-resolutiontechniques.html#cfn-entityresolution-matchingworkflow-resolutiontechniques-rulebasedproperties
	//
	RuleBasedProperties interface{} `field:"optional" json:"ruleBasedProperties" yaml:"ruleBasedProperties"`
	// An object containing the `rules` for a matching workflow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-resolutiontechniques.html#cfn-entityresolution-matchingworkflow-resolutiontechniques-ruleconditionproperties
	//
	RuleConditionProperties interface{} `field:"optional" json:"ruleConditionProperties" yaml:"ruleConditionProperties"`
}

