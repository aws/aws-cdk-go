package awsimagebuilderalpha


// Indicates whether the pipeline is enabled to be triggered by the provided schedule.
//
// Example:
//   advancedSchedulePipeline := imagebuilder.NewImagePipeline(this, jsii.String("AdvancedSchedulePipeline"), &ImagePipelineProps{
//   	Recipe: exampleImageRecipe,
//   	Schedule: &ImagePipelineSchedule{
//   		Expression: events.Schedule_Rate(awscdk.Duration_Days(jsii.Number(7))),
//   		// Only trigger when dependencies are updated (new base images, components, etc.)
//   		StartCondition: imagebuilder.ScheduleStartCondition_EXPRESSION_MATCH_AND_DEPENDENCY_UPDATES_AVAILABLE,
//   		// Automatically disable after 3 consecutive failures
//   		AutoDisableFailureCount: jsii.Number(3),
//   	},
//   	// Start enabled
//   	Status: imagebuilder.ImagePipelineStatus_ENABLED,
//   })
//
// Experimental.
type ImagePipelineStatus string

const (
	// Indicates that the pipeline is enabled for scheduling.
	// Experimental.
	ImagePipelineStatus_ENABLED ImagePipelineStatus = "ENABLED"
	// Indicates that the pipeline is disabled and will not be triggered on the schedule.
	// Experimental.
	ImagePipelineStatus_DISABLED ImagePipelineStatus = "DISABLED"
)

