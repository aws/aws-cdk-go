package previewawsnetworkfirewallevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.networkfirewall@FirewallAttachmentStatusChanged event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   firewallAttachmentStatusChangedProps := &FirewallAttachmentStatusChangedProps{
//   	Data: []DataItem{
//   		&DataItem{
//   			AvailabilityZone: []*string{
//   				jsii.String("availabilityZone"),
//   			},
//   			CurrentAttachmentStatus: []*string{
//   				jsii.String("currentAttachmentStatus"),
//   			},
//   			EndpointId: []*string{
//   				jsii.String("endpointId"),
//   			},
//   			PreviousAttachmentStatus: []*string{
//   				jsii.String("previousAttachmentStatus"),
//   			},
//   		},
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
//   	Metadata: &Metadata{
//   		StateChangeId: []*string{
//   			jsii.String("stateChangeId"),
//   		},
//   	},
//   	Version: []*string{
//   		jsii.String("version"),
//   	},
//   }
//
// Experimental.
type FirewallAttachmentStatusChanged_FirewallAttachmentStatusChangedProps struct {
	// data property.
	//
	// Specify an array of string values to match this event if the actual value of data is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Data *[]*FirewallAttachmentStatusChanged_DataItem `field:"optional" json:"data" yaml:"data"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// metadata property.
	//
	// Specify an array of string values to match this event if the actual value of metadata is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Metadata *FirewallAttachmentStatusChanged_Metadata `field:"optional" json:"metadata" yaml:"metadata"`
	// version property.
	//
	// Specify an array of string values to match this event if the actual value of version is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Version *[]*string `field:"optional" json:"version" yaml:"version"`
}

