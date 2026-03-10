package previewawsdeadlineevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.deadline@FleetSizeRecommendationChange event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   fleetSizeRecommendationChangeProps := &FleetSizeRecommendationChangeProps{
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
//   	FarmId: []*string{
//   		jsii.String("farmId"),
//   	},
//   	FleetId: []*string{
//   		jsii.String("fleetId"),
//   	},
//   	NewFleetSize: []*string{
//   		jsii.String("newFleetSize"),
//   	},
//   	OldFleetSize: []*string{
//   		jsii.String("oldFleetSize"),
//   	},
//   }
//
// Experimental.
type FleetSizeRecommendationChange_FleetSizeRecommendationChangeProps struct {
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// farmId property.
	//
	// Specify an array of string values to match this event if the actual value of farmId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FarmId *[]*string `field:"optional" json:"farmId" yaml:"farmId"`
	// fleetId property.
	//
	// Specify an array of string values to match this event if the actual value of fleetId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FleetId *[]*string `field:"optional" json:"fleetId" yaml:"fleetId"`
	// newFleetSize property.
	//
	// Specify an array of string values to match this event if the actual value of newFleetSize is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	NewFleetSize *[]*string `field:"optional" json:"newFleetSize" yaml:"newFleetSize"`
	// oldFleetSize property.
	//
	// Specify an array of string values to match this event if the actual value of oldFleetSize is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	OldFleetSize *[]*string `field:"optional" json:"oldFleetSize" yaml:"oldFleetSize"`
}

