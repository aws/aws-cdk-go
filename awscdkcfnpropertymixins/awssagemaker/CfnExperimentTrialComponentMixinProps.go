package awssagemaker


// Properties for CfnExperimentTrialComponentPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnExperimentTrialComponentMixinProps := &CfnExperimentTrialComponentMixinProps{
//   	DisplayName: jsii.String("displayName"),
//   	EndTime: jsii.String("endTime"),
//   	MetadataProperties: &MetadataPropertiesProperty{
//   		CommitId: jsii.String("commitId"),
//   		GeneratedBy: jsii.String("generatedBy"),
//   		ProjectId: jsii.String("projectId"),
//   		Repository: jsii.String("repository"),
//   	},
//   	StartTime: jsii.String("startTime"),
//   	Status: &StatusProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-experimenttrialcomponent.html
//
type CfnExperimentTrialComponentMixinProps struct {
	// The name of the component as displayed.
	//
	// The name doesn't need to be unique. If DisplayName isn't specified, TrialComponentName is displayed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-experimenttrialcomponent.html#cfn-sagemaker-experimenttrialcomponent-displayname
	//
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// When the component ended.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-experimenttrialcomponent.html#cfn-sagemaker-experimenttrialcomponent-endtime
	//
	EndTime *string `field:"optional" json:"endTime" yaml:"endTime"`
	// Metadata properties of the tracking entity, trial, or trial component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-experimenttrialcomponent.html#cfn-sagemaker-experimenttrialcomponent-metadataproperties
	//
	MetadataProperties interface{} `field:"optional" json:"metadataProperties" yaml:"metadataProperties"`
	// When the component started.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-experimenttrialcomponent.html#cfn-sagemaker-experimenttrialcomponent-starttime
	//
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
	// The status of the trial component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-experimenttrialcomponent.html#cfn-sagemaker-experimenttrialcomponent-status
	//
	Status interface{} `field:"optional" json:"status" yaml:"status"`
	// A list of tags to associate with the component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-experimenttrialcomponent.html#cfn-sagemaker-experimenttrialcomponent-tags
	//
	Tags *[]*CfnExperimentTrialComponentPropsMixin_TagsItemsProperty `field:"optional" json:"tags" yaml:"tags"`
	// The name of the trial component.
	//
	// The name must be unique in your AWS account and is not case-sensitive.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-experimenttrialcomponent.html#cfn-sagemaker-experimenttrialcomponent-trialcomponentname
	//
	TrialComponentName *string `field:"optional" json:"trialComponentName" yaml:"trialComponentName"`
}

