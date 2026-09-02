package awsmedialivealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm"
)

// Connection details for an ESAM POIS (Placement Opportunity Information System) endpoint.
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
type PoisEndpoint struct {
	// The POIS endpoint URL that MediaLive sends signal conditioning information to.
	// Experimental.
	Url *string `field:"required" json:"url" yaml:"url"`
	// An SSM parameter holding the password for the POIS endpoint.
	//
	// The channel role is granted
	// read access to the parameter automatically.
	// Default: - no credentials.
	//
	// Experimental.
	Password awsssm.IStringParameter `field:"optional" json:"password" yaml:"password"`
	// The username used to connect to the POIS endpoint.
	// Default: - no credentials.
	//
	// Experimental.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

