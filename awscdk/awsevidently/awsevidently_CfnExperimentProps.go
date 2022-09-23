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
//   cfnExperimentProps := &cfnExperimentProps{
//   	metricGoals: []interface{}{
//   		&metricGoalObjectProperty{
//   			desiredChange: jsii.String("desiredChange"),
//   			entityIdKey: jsii.String("entityIdKey"),
//   			eventPattern: jsii.String("eventPattern"),
//   			metricName: jsii.String("metricName"),
//   			valueKey: jsii.String("valueKey"),
//
//   			// the properties below are optional
//   			unitLabel: jsii.String("unitLabel"),
//   		},
//   	},
//   	name: jsii.String("name"),
//   	onlineAbConfig: &onlineAbConfigObjectProperty{
//   		controlTreatmentName: jsii.String("controlTreatmentName"),
//   		treatmentWeights: []interface{}{
//   			&treatmentToWeightProperty{
//   				splitWeight: jsii.Number(123),
//   				treatment: jsii.String("treatment"),
//   			},
//   		},
//   	},
//   	project: jsii.String("project"),
//   	treatments: []interface{}{
//   		&treatmentObjectProperty{
//   			feature: jsii.String("feature"),
//   			treatmentName: jsii.String("treatmentName"),
//   			variation: jsii.String("variation"),
//
//   			// the properties below are optional
//   			description: jsii.String("description"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	randomizationSalt: jsii.String("randomizationSalt"),
//   	removeSegment: jsii.Boolean(false),
//   	runningStatus: &runningStatusObjectProperty{
//   		analysisCompleteTime: jsii.String("analysisCompleteTime"),
//   		desiredState: jsii.String("desiredState"),
//   		reason: jsii.String("reason"),
//   		status: jsii.String("status"),
//   	},
//   	samplingRate: jsii.Number(123),
//   	segment: jsii.String("segment"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
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
	// `AWS::Evidently::Experiment.RemoveSegment`.
	RemoveSegment interface{} `field:"optional" json:"removeSegment" yaml:"removeSegment"`
	// A structure that you can use to start and stop the experiment.
	RunningStatus interface{} `field:"optional" json:"runningStatus" yaml:"runningStatus"`
	// The portion of the available audience that you want to allocate to this experiment, in thousandths of a percent.
	//
	// The available audience is the total audience minus the audience that you have allocated to overrides or current launches of this feature.
	//
	// This is represented in thousandths of a percent. For example, specify 10,000 to allocate 10% of the available audience.
	SamplingRate *float64 `field:"optional" json:"samplingRate" yaml:"samplingRate"`
	// `AWS::Evidently::Experiment.Segment`.
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

