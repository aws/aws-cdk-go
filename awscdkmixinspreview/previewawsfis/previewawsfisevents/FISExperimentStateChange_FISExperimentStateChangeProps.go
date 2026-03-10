package previewawsfisevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.fis@FISExperimentStateChange event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   fISExperimentStateChangeProps := &FISExperimentStateChangeProps{
//   	EventMetadata: &AWSEventMetadataProps{
//   		Region: []*string{
//   			jsii.String("region"),
//   		},
//   		Resources: []*string{
//   			jsii.String("resources"),
//   		},
//   		Version: []*string{
//   			jsii.String("version"),
//   		},
//   	},
//   	ExperimentId: []*string{
//   		jsii.String("experimentId"),
//   	},
//   	NewState: &NewState{
//   		Reason: []*string{
//   			jsii.String("reason"),
//   		},
//   		Status: []*string{
//   			jsii.String("status"),
//   		},
//   	},
//   	OldState: &OldState{
//   		Reason: []*string{
//   			jsii.String("reason"),
//   		},
//   		Status: []*string{
//   			jsii.String("status"),
//   		},
//   	},
//   }
//
// Experimental.
type FISExperimentStateChange_FISExperimentStateChangeProps struct {
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// experiment-id property.
	//
	// Specify an array of string values to match this event if the actual value of experiment-id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ExperimentId *[]*string `field:"optional" json:"experimentId" yaml:"experimentId"`
	// new-state property.
	//
	// Specify an array of string values to match this event if the actual value of new-state is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	NewState *FISExperimentStateChange_NewState `field:"optional" json:"newState" yaml:"newState"`
	// old-state property.
	//
	// Specify an array of string values to match this event if the actual value of old-state is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	OldState *FISExperimentStateChange_OldState `field:"optional" json:"oldState" yaml:"oldState"`
}

