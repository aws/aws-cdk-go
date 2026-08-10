package awssagemaker


// Properties for defining a `CfnTrialComponent`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTrialComponentProps := &CfnTrialComponentProps{
//   	TrialComponentName: jsii.String("trialComponentName"),
//
//   	// the properties below are optional
//   	DisplayName: jsii.String("displayName"),
//   	InputArtifacts: map[string]interface{}{
//   		"inputArtifactsKey": &TrialComponentArtifactProperty{
//   			"value": jsii.String("value"),
//
//   			// the properties below are optional
//   			"mediaType": jsii.String("mediaType"),
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
//   			"value": jsii.String("value"),
//
//   			// the properties below are optional
//   			"mediaType": jsii.String("mediaType"),
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-trialcomponent.html
//
type CfnTrialComponentProps struct {
	// The name of the trial component.
	//
	// Must be unique in your AWS account and is not case-sensitive.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-trialcomponent.html#cfn-sagemaker-trialcomponent-trialcomponentname
	//
	TrialComponentName *string `field:"required" json:"trialComponentName" yaml:"trialComponentName"`
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
	Tags *[]*CfnTrialComponent_TagsItemsProperty `field:"optional" json:"tags" yaml:"tags"`
}

