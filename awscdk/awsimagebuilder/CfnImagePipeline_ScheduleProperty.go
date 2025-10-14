package awsimagebuilder


// A schedule configures when and how often a pipeline will automatically create a new image.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scheduleProperty := &ScheduleProperty{
//   	AutoDisablePolicy: &AutoDisablePolicyProperty{
//   		FailureCount: jsii.Number(123),
//   	},
//   	PipelineExecutionStartCondition: jsii.String("pipelineExecutionStartCondition"),
//   	ScheduleExpression: jsii.String("scheduleExpression"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagepipeline-schedule.html
//
type CfnImagePipeline_ScheduleProperty struct {
	// The policy that configures when Image Builder should automatically disable a pipeline that is failing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagepipeline-schedule.html#cfn-imagebuilder-imagepipeline-schedule-autodisablepolicy
	//
	AutoDisablePolicy interface{} `field:"optional" json:"autoDisablePolicy" yaml:"autoDisablePolicy"`
	// The condition configures when the pipeline should trigger a new image build.
	//
	// When the `pipelineExecutionStartCondition` is set to `EXPRESSION_MATCH_AND_DEPENDENCY_UPDATES_AVAILABLE` , and you use semantic version filters on the base image or components in your image recipe, Image Builder will build a new image only when there are new versions of the image or components in your recipe that match the semantic version filter. When it is set to `EXPRESSION_MATCH_ONLY` , it will build a new image every time the CRON expression matches the current time. For semantic version syntax, see [CreateComponent](https://docs.aws.amazon.com/imagebuilder/latest/APIReference/API_CreateComponent.html) in the *Image Builder API Reference* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagepipeline-schedule.html#cfn-imagebuilder-imagepipeline-schedule-pipelineexecutionstartcondition
	//
	PipelineExecutionStartCondition *string `field:"optional" json:"pipelineExecutionStartCondition" yaml:"pipelineExecutionStartCondition"`
	// The cron expression determines how often EC2 Image Builder evaluates your `pipelineExecutionStartCondition` .
	//
	// For information on how to format a cron expression in Image Builder, see [Use cron expressions in EC2 Image Builder](https://docs.aws.amazon.com/imagebuilder/latest/userguide/image-builder-cron.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagepipeline-schedule.html#cfn-imagebuilder-imagepipeline-schedule-scheduleexpression
	//
	ScheduleExpression *string `field:"optional" json:"scheduleExpression" yaml:"scheduleExpression"`
}

