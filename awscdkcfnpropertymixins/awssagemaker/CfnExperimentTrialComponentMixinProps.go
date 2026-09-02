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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-experimenttrialcomponent.html#cfn-sagemaker-experimenttrialcomponent-displayname
	//
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-experimenttrialcomponent.html#cfn-sagemaker-experimenttrialcomponent-endtime
	//
	EndTime *string `field:"optional" json:"endTime" yaml:"endTime"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-experimenttrialcomponent.html#cfn-sagemaker-experimenttrialcomponent-metadataproperties
	//
	MetadataProperties interface{} `field:"optional" json:"metadataProperties" yaml:"metadataProperties"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-experimenttrialcomponent.html#cfn-sagemaker-experimenttrialcomponent-starttime
	//
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-experimenttrialcomponent.html#cfn-sagemaker-experimenttrialcomponent-status
	//
	Status interface{} `field:"optional" json:"status" yaml:"status"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-experimenttrialcomponent.html#cfn-sagemaker-experimenttrialcomponent-tags
	//
	Tags *[]*CfnExperimentTrialComponentPropsMixin_TagsItemsProperty `field:"optional" json:"tags" yaml:"tags"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-experimenttrialcomponent.html#cfn-sagemaker-experimenttrialcomponent-trialcomponentname
	//
	TrialComponentName *string `field:"optional" json:"trialComponentName" yaml:"trialComponentName"`
}

