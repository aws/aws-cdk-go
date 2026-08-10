package awssagemaker


// Associates a SageMaker job as a trial component with an experiment and trial.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   experimentConfigProperty := &ExperimentConfigProperty{
//   	ExperimentName: jsii.String("experimentName"),
//   	TrialComponentDisplayName: jsii.String("trialComponentDisplayName"),
//   	TrialName: jsii.String("trialName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-transformjob-experimentconfig.html
//
type CfnTransformJob_ExperimentConfigProperty struct {
	// The name of an existing experiment to associate with the trial component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-transformjob-experimentconfig.html#cfn-sagemaker-transformjob-experimentconfig-experimentname
	//
	ExperimentName *string `field:"optional" json:"experimentName" yaml:"experimentName"`
	// The display name for the trial component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-transformjob-experimentconfig.html#cfn-sagemaker-transformjob-experimentconfig-trialcomponentdisplayname
	//
	TrialComponentDisplayName *string `field:"optional" json:"trialComponentDisplayName" yaml:"trialComponentDisplayName"`
	// The name of an existing trial to associate the trial component with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-transformjob-experimentconfig.html#cfn-sagemaker-transformjob-experimentconfig-trialname
	//
	TrialName *string `field:"optional" json:"trialName" yaml:"trialName"`
}

