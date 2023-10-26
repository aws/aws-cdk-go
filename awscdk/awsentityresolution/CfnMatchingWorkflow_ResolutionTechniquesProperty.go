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
	// The type of matching.
	//
	// There are two types of matching: `RULE_MATCHING` and `ML_MATCHING` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-resolutiontechniques.html#cfn-entityresolution-matchingworkflow-resolutiontechniques-resolutiontype
	//
	ResolutionType *string `field:"optional" json:"resolutionType" yaml:"resolutionType"`
	// An object which defines the list of matching rules to run and has a field `Rules` , which is a list of rule objects.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-resolutiontechniques.html#cfn-entityresolution-matchingworkflow-resolutiontechniques-rulebasedproperties
	//
	RuleBasedProperties interface{} `field:"optional" json:"ruleBasedProperties" yaml:"ruleBasedProperties"`
}

