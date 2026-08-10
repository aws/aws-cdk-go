package awssagemaker


// Properties for CfnExperimentPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnExperimentMixinProps := &CfnExperimentMixinProps{
//   	Description: jsii.String("description"),
//   	DisplayName: jsii.String("displayName"),
//   	ExperimentName: jsii.String("experimentName"),
//   	Tags: []TagsItemsProperty{
//   		&TagsItemsProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-experiment.html
//
type CfnExperimentMixinProps struct {
	// The description of the experiment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-experiment.html#cfn-sagemaker-experiment-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the experiment as displayed.
	//
	// The name does not need to be unique.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-experiment.html#cfn-sagemaker-experiment-displayname
	//
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The name of the experiment.
	//
	// Must be unique in your AWS account and is not case-sensitive.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-experiment.html#cfn-sagemaker-experiment-experimentname
	//
	ExperimentName *string `field:"optional" json:"experimentName" yaml:"experimentName"`
	// A list of tags to associate with the experiment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-experiment.html#cfn-sagemaker-experiment-tags
	//
	Tags *[]*CfnExperimentPropsMixin_TagsItemsProperty `field:"optional" json:"tags" yaml:"tags"`
}

