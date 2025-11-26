package previewawsnetworkmanagerevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for CoreNetwork aws.networkmanager@NetworkManagerSegmentUpdate event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   networkManagerSegmentUpdateProps := &NetworkManagerSegmentUpdateProps{
//   	AttachmentArn: []*string{
//   		jsii.String("attachmentArn"),
//   	},
//   	ChangeDescription: []*string{
//   		jsii.String("changeDescription"),
//   	},
//   	ChangeType: []*string{
//   		jsii.String("changeType"),
//   	},
//   	CoreNetworkArn: []*string{
//   		jsii.String("coreNetworkArn"),
//   	},
//   	EdgeLocation: []*string{
//   		jsii.String("edgeLocation"),
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
//   	NetworkFunctionGroupName: []*string{
//   		jsii.String("networkFunctionGroupName"),
//   	},
//   	NewNetworkFunctionGroupName: []*string{
//   		jsii.String("newNetworkFunctionGroupName"),
//   	},
//   	NewSegmentName: []*string{
//   		jsii.String("newSegmentName"),
//   	},
//   	PreviousNetworkFunctionGroupName: []*string{
//   		jsii.String("previousNetworkFunctionGroupName"),
//   	},
//   	PreviousSegmentName: []*string{
//   		jsii.String("previousSegmentName"),
//   	},
//   	SegmentName: []*string{
//   		jsii.String("segmentName"),
//   	},
//   }
//
// Experimental.
type CoreNetworkEvents_NetworkManagerSegmentUpdate_NetworkManagerSegmentUpdateProps struct {
	// attachmentArn property.
	//
	// Specify an array of string values to match this event if the actual value of attachmentArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AttachmentArn *[]*string `field:"optional" json:"attachmentArn" yaml:"attachmentArn"`
	// changeDescription property.
	//
	// Specify an array of string values to match this event if the actual value of changeDescription is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ChangeDescription *[]*string `field:"optional" json:"changeDescription" yaml:"changeDescription"`
	// changeType property.
	//
	// Specify an array of string values to match this event if the actual value of changeType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ChangeType *[]*string `field:"optional" json:"changeType" yaml:"changeType"`
	// coreNetworkArn property.
	//
	// Specify an array of string values to match this event if the actual value of coreNetworkArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Filter with the CoreNetwork reference.
	//
	// Experimental.
	CoreNetworkArn *[]*string `field:"optional" json:"coreNetworkArn" yaml:"coreNetworkArn"`
	// edgeLocation property.
	//
	// Specify an array of string values to match this event if the actual value of edgeLocation is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EdgeLocation *[]*string `field:"optional" json:"edgeLocation" yaml:"edgeLocation"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// networkFunctionGroupName property.
	//
	// Specify an array of string values to match this event if the actual value of networkFunctionGroupName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	NetworkFunctionGroupName *[]*string `field:"optional" json:"networkFunctionGroupName" yaml:"networkFunctionGroupName"`
	// newNetworkFunctionGroupName property.
	//
	// Specify an array of string values to match this event if the actual value of newNetworkFunctionGroupName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	NewNetworkFunctionGroupName *[]*string `field:"optional" json:"newNetworkFunctionGroupName" yaml:"newNetworkFunctionGroupName"`
	// newSegmentName property.
	//
	// Specify an array of string values to match this event if the actual value of newSegmentName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	NewSegmentName *[]*string `field:"optional" json:"newSegmentName" yaml:"newSegmentName"`
	// previousNetworkFunctionGroupName property.
	//
	// Specify an array of string values to match this event if the actual value of previousNetworkFunctionGroupName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PreviousNetworkFunctionGroupName *[]*string `field:"optional" json:"previousNetworkFunctionGroupName" yaml:"previousNetworkFunctionGroupName"`
	// previousSegmentName property.
	//
	// Specify an array of string values to match this event if the actual value of previousSegmentName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PreviousSegmentName *[]*string `field:"optional" json:"previousSegmentName" yaml:"previousSegmentName"`
	// segmentName property.
	//
	// Specify an array of string values to match this event if the actual value of segmentName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SegmentName *[]*string `field:"optional" json:"segmentName" yaml:"segmentName"`
}

