package awsimagebuilderalpha


// The start condition for the pipeline, indicating the condition under which a pipeline should be triggered.
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
type ScheduleStartCondition string

const (
	// Indicates to trigger a pipeline whenever its schedule is met.
	// Experimental.
	ScheduleStartCondition_EXPRESSION_MATCH_ONLY ScheduleStartCondition = "EXPRESSION_MATCH_ONLY"
	// Indicates to trigger a pipeline whenever its schedule is met, and there are matching dependency updates available, such as new versions of components or images to use in the pipeline build.
	// Experimental.
	ScheduleStartCondition_EXPRESSION_MATCH_AND_DEPENDENCY_UPDATES_AVAILABLE ScheduleStartCondition = "EXPRESSION_MATCH_AND_DEPENDENCY_UPDATES_AVAILABLE"
)

