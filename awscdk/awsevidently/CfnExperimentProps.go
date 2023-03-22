package awsevidently

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnExperiment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnExperimentProps := &CfnExperimentProps{
//   	MetricGoals: []interface{}{
//   		&MetricGoalObjectProperty{
//   			DesiredChange: jsii.String("desiredChange"),
//   			EntityIdKey: jsii.String("entityIdKey"),
//   			MetricName: jsii.String("metricName"),
//   			ValueKey: jsii.String("valueKey"),
//
//   			// the properties below are optional
//   			EventPattern: jsii.String("eventPattern"),
//   			UnitLabel: jsii.String("unitLabel"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	OnlineAbConfig: &OnlineAbConfigObjectProperty{
//   		ControlTreatmentName: jsii.String("controlTreatmentName"),
//   		TreatmentWeights: []interface{}{
//   			&TreatmentToWeightProperty{
//   				SplitWeight: jsii.Number(123),
//   				Treatment: jsii.String("treatment"),
//   			},
//   		},
//   	},
//   	Project: jsii.String("project"),
//   	Treatments: []interface{}{
//   		&TreatmentObjectProperty{
//   			Feature: jsii.String("feature"),
//   			TreatmentName: jsii.String("treatmentName"),
//   			Variation: jsii.String("variation"),
//
//   			// the properties below are optional
//   			Description: jsii.String("description"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	RandomizationSalt: jsii.String("randomizationSalt"),
//   	RemoveSegment: jsii.Boolean(false),
//   	RunningStatus: &RunningStatusObjectProperty{
//   		Status: jsii.String("status"),
//
//   		// the properties below are optional
//   		AnalysisCompleteTime: jsii.String("analysisCompleteTime"),
//   		DesiredState: jsii.String("desiredState"),
//   		Reason: jsii.String("reason"),
//   	},
//   	SamplingRate: jsii.Number(123),
//   	Segment: jsii.String("segment"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnExperimentProps struct {
	// An array of structures that defines the metrics used for the experiment, and whether a higher or lower value for each metric is the goal.
	//
	// You can use up to three metrics in an experiment.
	MetricGoals interface{} `field:"required" json:"metricGoals" yaml:"metricGoals"`
	// A name for the new experiment.
	Name *string `field:"required" json:"name" yaml:"name"`
	// A structure that contains the configuration of which variation to use as the "control" version.
	//
	// The "control" version is used for comparison with other variations. This structure also specifies how much experiment traffic is allocated to each variation.
	OnlineAbConfig interface{} `field:"required" json:"onlineAbConfig" yaml:"onlineAbConfig"`
	// The name or the ARN of the project where this experiment is to be created.
	Project *string `field:"required" json:"project" yaml:"project"`
	// An array of structures that describe the configuration of each feature variation used in the experiment.
	Treatments interface{} `field:"required" json:"treatments" yaml:"treatments"`
	// An optional description of the experiment.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// When Evidently assigns a particular user session to an experiment, it must use a randomization ID to determine which variation the user session is served.
	//
	// This randomization ID is a combination of the entity ID and `randomizationSalt` . If you omit `randomizationSalt` , Evidently uses the experiment name as the `randomizationSalt` .
	RandomizationSalt *string `field:"optional" json:"randomizationSalt" yaml:"randomizationSalt"`
	// Set this to `true` to remove the segment that is associated with this experiment.
	//
	// You can't use this parameter if the experiment is currently running.
	RemoveSegment interface{} `field:"optional" json:"removeSegment" yaml:"removeSegment"`
	// A structure that you can use to start and stop the experiment.
	RunningStatus interface{} `field:"optional" json:"runningStatus" yaml:"runningStatus"`
	// The portion of the available audience that you want to allocate to this experiment, in thousandths of a percent.
	//
	// The available audience is the total audience minus the audience that you have allocated to overrides or current launches of this feature.
	//
	// This is represented in thousandths of a percent. For example, specify 10,000 to allocate 10% of the available audience.
	SamplingRate *float64 `field:"optional" json:"samplingRate" yaml:"samplingRate"`
	// Specifies an audience *segment* to use in the experiment.
	//
	// When a segment is used in an experiment, only user sessions that match the segment pattern are used in the experiment.
	//
	// For more information, see [Segment rule pattern syntax](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Evidently-segments-syntax.html) .
	Segment *string `field:"optional" json:"segment" yaml:"segment"`
	// Assigns one or more tags (key-value pairs) to the experiment.
	//
	// Tags can help you organize and categorize your resources. You can also use them to scope user permissions by granting a user permission to access or change only resources with certain tag values.
	//
	// Tags don't have any semantic meaning to AWS and are interpreted strictly as strings of characters.
	//
	// You can associate as many as 50 tags with an experiment.
	//
	// For more information, see [Tagging AWS resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

