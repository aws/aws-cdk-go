package awsimagebuilderalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// The schedule settings for the image pipeline, which defines when a pipeline should be triggered.
//
// Example:
//   dailyPipeline := imagebuilder.NewImagePipeline(this, jsii.String("DailyPipeline"), &ImagePipelineProps{
//   	Recipe: exampleContainerRecipe,
//   	Schedule: &ImagePipelineSchedule{
//   		Expression: events.Schedule_Rate(awscdk.Duration_Days(jsii.Number(1))),
//   	},
//   })
//
// Experimental.
type ImagePipelineSchedule struct {
	// The schedule expression to use.
	//
	// This can either be a cron expression or a rate expression.
	// Experimental.
	Expression awsevents.Schedule `field:"required" json:"expression" yaml:"expression"`
	// The number of consecutive failures allowed before the pipeline is automatically disabled.
	//
	// This value must be
	// between 1 and 10.
	// Default: - no auto-disable policy is configured and the pipeline is not automatically disabled on consecutive
	// failures.
	//
	// Experimental.
	AutoDisableFailureCount *float64 `field:"optional" json:"autoDisableFailureCount" yaml:"autoDisableFailureCount"`
	// The start condition for the pipeline, indicating the condition under which a pipeline should be triggered.
	// Default: ScheduleStartCondition.EXPRESSION_MATCH_AND_DEPENDENCY_UPDATES_AVAILABLE
	//
	// Experimental.
	StartCondition ScheduleStartCondition `field:"optional" json:"startCondition" yaml:"startCondition"`
}

