package awsmedialivealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Settings for ESAM (Event Signaling and Management) ad avail handling.
//
// MediaLive signals ad avail
// events to an external POIS (Placement Opportunity Information System) endpoint.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var stack Stack
//   var input IInput
//   var bucket IBucket
//   var video EncodeConfiguration
//   var poisPassword StringParameter
//
//
//   medialive.NewChannel(stack, jsii.String("Channel"), &ChannelProps{
//   	Inputs: []InputAttachment{
//   		&InputAttachment{
//   			Input: *Input,
//   		},
//   	},
//   	AvailSettings: medialive.AvailSettings_Esam(&EsamSettings{
//   		Pois: &PoisEndpoint{
//   			Url: jsii.String("https://pois.example.com/esam"),
//   			Username: jsii.String("pois-user"),
//   			Password: poisPassword,
//   		},
//   		AcquisitionPointId: jsii.String("acquisition-point-1"),
//   		AdAvailOffset: awscdk.Duration_Millis(jsii.Number(200)),
//   	}),
//   	Scte35SegmentationScope: medialive.Scte35SegmentationScope_SCTE35_ENABLED_OUTPUT_GROUPS(),
//   	OutputGroups: []OutputGroupConfiguration{
//   		medialive.OutputGroupConfiguration_Hls(&HlsOutputGroupProps{
//   			Name: jsii.String("hls"),
//   			Destinations: []OutputDestination{
//   				medialive.OutputDestination_ToBucket(bucket, jsii.String("live/stream")),
//   			},
//   			Outputs: []HlsOutputDefinition{
//   				&HlsOutputDefinition{
//   					Encodes: []EncodeConfiguration{
//   						video,
//   					},
//   					OutputName: jsii.String("hls_out"),
//   				},
//   			},
//   		}),
//   	},
//   })
//
// Experimental.
type EsamSettings struct {
	// The acquisition point identity sent to the POIS in MCC requests.
	// Experimental.
	AcquisitionPointId *string `field:"required" json:"acquisitionPointId" yaml:"acquisitionPointId"`
	// The POIS endpoint connection details — URL and optional credentials.
	// Experimental.
	Pois *PoisEndpoint `field:"required" json:"pois" yaml:"pois"`
	// Offset added to the input ad avail PTS time.
	// Default: - service default.
	//
	// Experimental.
	AdAvailOffset awscdk.Duration `field:"optional" json:"adAvailOffset" yaml:"adAvailOffset"`
	// The ID of a zone the POIS uses to control the placement of ad avails.
	// Default: - service default.
	//
	// Experimental.
	ZoneIdentity *string `field:"optional" json:"zoneIdentity" yaml:"zoneIdentity"`
}

