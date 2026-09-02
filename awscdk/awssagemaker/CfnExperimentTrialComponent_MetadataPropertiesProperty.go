package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metadataPropertiesProperty := &MetadataPropertiesProperty{
//   	CommitId: jsii.String("commitId"),
//   	GeneratedBy: jsii.String("generatedBy"),
//   	ProjectId: jsii.String("projectId"),
//   	Repository: jsii.String("repository"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-experimenttrialcomponent-metadataproperties.html
//
type CfnExperimentTrialComponent_MetadataPropertiesProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-experimenttrialcomponent-metadataproperties.html#cfn-sagemaker-experimenttrialcomponent-metadataproperties-commitid
	//
	CommitId *string `field:"optional" json:"commitId" yaml:"commitId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-experimenttrialcomponent-metadataproperties.html#cfn-sagemaker-experimenttrialcomponent-metadataproperties-generatedby
	//
	GeneratedBy *string `field:"optional" json:"generatedBy" yaml:"generatedBy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-experimenttrialcomponent-metadataproperties.html#cfn-sagemaker-experimenttrialcomponent-metadataproperties-projectid
	//
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-experimenttrialcomponent-metadataproperties.html#cfn-sagemaker-experimenttrialcomponent-metadataproperties-repository
	//
	Repository *string `field:"optional" json:"repository" yaml:"repository"`
}

