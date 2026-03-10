package previewawsimagebuilderevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.imagebuilder@EC2ImageBuilderImageStateChange event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eC2ImageBuilderImageStateChangeProps := &EC2ImageBuilderImageStateChangeProps{
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
//   	PreviousState: &ImageState{
//   		Reason: []*string{
//   			jsii.String("reason"),
//   		},
//   		Status: []*string{
//   			jsii.String("status"),
//   		},
//   	},
//   	State: &ImageState{
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
type EC2ImageBuilderImageStateChange_EC2ImageBuilderImageStateChangeProps struct {
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// previous-state property.
	//
	// Specify an array of string values to match this event if the actual value of previous-state is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PreviousState *EC2ImageBuilderImageStateChange_ImageState `field:"optional" json:"previousState" yaml:"previousState"`
	// state property.
	//
	// Specify an array of string values to match this event if the actual value of state is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	State *EC2ImageBuilderImageStateChange_ImageState `field:"optional" json:"state" yaml:"state"`
}

