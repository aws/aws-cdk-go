package awssagemaker


// Properties for CfnTrialComponentPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnTrialComponentMixinProps := &CfnTrialComponentMixinProps{
//   	DisplayName: jsii.String("displayName"),
//   	InputArtifacts: map[string]interface{}{
//   		"inputArtifactsKey": &TrialComponentArtifactProperty{
//   			"mediaType": jsii.String("mediaType"),
//   			"value": jsii.String("value"),
//   		},
//   	},
//   	MetadataProperties: &MetadataPropertiesProperty{
//   		CommitId: jsii.String("commitId"),
//   		GeneratedBy: jsii.String("generatedBy"),
//   		ProjectId: jsii.String("projectId"),
//   		Repository: jsii.String("repository"),
//   	},
//   	OutputArtifacts: map[string]interface{}{
//   		"outputArtifactsKey": &TrialComponentArtifactProperty{
//   			"mediaType": jsii.String("mediaType"),
//   			"value": jsii.String("value"),
//   		},
//   	},
//   	Parameters: map[string]interface{}{
//   		"parametersKey": &TrialComponentParameterValueProperty{
//   			"numberValue": jsii.Number(123),
//   			"stringValue": jsii.String("stringValue"),
//   		},
//   	},
//   	Status: &TrialComponentStatusProperty{
//   		Message: jsii.String("message"),
//   		PrimaryStatus: jsii.String("primaryStatus"),
//   	},
//   	Tags: []TagsItemsProperty{
//   		&TagsItemsProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TrialComponentName: jsii.String("trialComponentName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-trialcomponent.html
//
type CfnTrialComponentMixinProps struct {
	// The name of the component as displayed.
	//
	// If DisplayName isn't specified, TrialComponentName is displayed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-trialcomponent.html#cfn-sagemaker-trialcomponent-displayname
	//
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The input artifacts for the component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-trialcomponent.html#cfn-sagemaker-trialcomponent-inputartifacts
	//
	InputArtifacts interface{} `field:"optional" json:"inputArtifacts" yaml:"inputArtifacts"`
	// Metadata properties of the tracking entity, trial, or trial component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-trialcomponent.html#cfn-sagemaker-trialcomponent-metadataproperties
	//
	MetadataProperties interface{} `field:"optional" json:"metadataProperties" yaml:"metadataProperties"`
	// The output artifacts for the component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-trialcomponent.html#cfn-sagemaker-trialcomponent-outputartifacts
	//
	OutputArtifacts interface{} `field:"optional" json:"outputArtifacts" yaml:"outputArtifacts"`
	// The hyperparameters for the component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-trialcomponent.html#cfn-sagemaker-trialcomponent-parameters
	//
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// The status of the trial component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-trialcomponent.html#cfn-sagemaker-trialcomponent-status
	//
	Status interface{} `field:"optional" json:"status" yaml:"status"`
	// A list of tags to associate with the trial component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-trialcomponent.html#cfn-sagemaker-trialcomponent-tags
	//
	Tags *[]*CfnTrialComponentPropsMixin_TagsItemsProperty `field:"optional" json:"tags" yaml:"tags"`
	// The name of the trial component.
	//
	// Must be unique in your AWS account and is not case-sensitive.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-trialcomponent.html#cfn-sagemaker-trialcomponent-trialcomponentname
	//
	TrialComponentName *string `field:"optional" json:"trialComponentName" yaml:"trialComponentName"`
}

