package previewawsec2events

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.ec2@EBSMultiVolumeSnapshotsCompletionStatus event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eBSMultiVolumeSnapshotsCompletionStatusProps := &EBSMultiVolumeSnapshotsCompletionStatusProps{
//   	Cause: []*string{
//   		jsii.String("cause"),
//   	},
//   	EndTime: []*string{
//   		jsii.String("endTime"),
//   	},
//   	Event: []*string{
//   		jsii.String("event"),
//   	},
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
//   	RequestId: []*string{
//   		jsii.String("requestId"),
//   	},
//   	Result: []*string{
//   		jsii.String("result"),
//   	},
//   	Snapshots: []Snapshot{
//   		&Snapshot{
//   			SnapshotId: []*string{
//   				jsii.String("snapshotId"),
//   			},
//   			Source: []*string{
//   				jsii.String("source"),
//   			},
//   			Status: []*string{
//   				jsii.String("status"),
//   			},
//   		},
//   	},
//   	StartTime: []*string{
//   		jsii.String("startTime"),
//   	},
//   }
//
// Experimental.
type EBSMultiVolumeSnapshotsCompletionStatus_EBSMultiVolumeSnapshotsCompletionStatusProps struct {
	// cause property.
	//
	// Specify an array of string values to match this event if the actual value of cause is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Cause *[]*string `field:"optional" json:"cause" yaml:"cause"`
	// endTime property.
	//
	// Specify an array of string values to match this event if the actual value of endTime is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EndTime *[]*string `field:"optional" json:"endTime" yaml:"endTime"`
	// event property.
	//
	// Specify an array of string values to match this event if the actual value of event is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Event *[]*string `field:"optional" json:"event" yaml:"event"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// request-id property.
	//
	// Specify an array of string values to match this event if the actual value of request-id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RequestId *[]*string `field:"optional" json:"requestId" yaml:"requestId"`
	// result property.
	//
	// Specify an array of string values to match this event if the actual value of result is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Result *[]*string `field:"optional" json:"result" yaml:"result"`
	// snapshots property.
	//
	// Specify an array of string values to match this event if the actual value of snapshots is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Snapshots *[]*EBSMultiVolumeSnapshotsCompletionStatus_Snapshot `field:"optional" json:"snapshots" yaml:"snapshots"`
	// startTime property.
	//
	// Specify an array of string values to match this event if the actual value of startTime is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StartTime *[]*string `field:"optional" json:"startTime" yaml:"startTime"`
}

