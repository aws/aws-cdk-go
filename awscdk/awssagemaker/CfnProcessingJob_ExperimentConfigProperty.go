package awssagemaker


// Associates a SageMaker job as a trial component with an experiment and trial.
//
// Specified when you call the [CreateProcessingJob](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateProcessingJob.html) API.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   experimentConfigProperty := &ExperimentConfigProperty{
//   	ExperimentName: jsii.String("experimentName"),
//   	RunName: jsii.String("runName"),
//   	TrialComponentDisplayName: jsii.String("trialComponentDisplayName"),
//   	TrialName: jsii.String("trialName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-experimentconfig.html
//
type CfnProcessingJob_ExperimentConfigProperty struct {
	// The name of an existing experiment to associate with the trial component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-experimentconfig.html#cfn-sagemaker-processingjob-experimentconfig-experimentname
	//
	ExperimentName *string `field:"optional" json:"experimentName" yaml:"experimentName"`
	// The name of the experiment run to associate with the trial component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-experimentconfig.html#cfn-sagemaker-processingjob-experimentconfig-runname
	//
	RunName *string `field:"optional" json:"runName" yaml:"runName"`
	// The display name for the trial component.
	//
	// If this key isn't specified, the display name is the trial component name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-experimentconfig.html#cfn-sagemaker-processingjob-experimentconfig-trialcomponentdisplayname
	//
	TrialComponentDisplayName *string `field:"optional" json:"trialComponentDisplayName" yaml:"trialComponentDisplayName"`
	// The name of an existing trial to associate the trial component with.
	//
	// If not specified, a new trial is created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-experimentconfig.html#cfn-sagemaker-processingjob-experimentconfig-trialname
	//
	TrialName *string `field:"optional" json:"trialName" yaml:"trialName"`
}

