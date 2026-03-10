package previewawsimagebuilderevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.imagebuilder@EC2ImageBuilderCVEDetected event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eC2ImageBuilderCVEDetectedProps := &EC2ImageBuilderCVEDetectedProps{
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
//   	FindingSeverityCounts: &FindingSeverityCounts{
//   		All: []*string{
//   			jsii.String("all"),
//   		},
//   		Critical: []*string{
//   			jsii.String("critical"),
//   		},
//   		High: []*string{
//   			jsii.String("high"),
//   		},
//   		Medium: []*string{
//   			jsii.String("medium"),
//   		},
//   	},
//   	ResourceId: []*string{
//   		jsii.String("resourceId"),
//   	},
//   }
//
// Experimental.
type EC2ImageBuilderCVEDetected_EC2ImageBuilderCVEDetectedProps struct {
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// finding-severity-counts property.
	//
	// Specify an array of string values to match this event if the actual value of finding-severity-counts is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FindingSeverityCounts *EC2ImageBuilderCVEDetected_FindingSeverityCounts `field:"optional" json:"findingSeverityCounts" yaml:"findingSeverityCounts"`
	// resource-id property.
	//
	// Specify an array of string values to match this event if the actual value of resource-id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ResourceId *[]*string `field:"optional" json:"resourceId" yaml:"resourceId"`
}

