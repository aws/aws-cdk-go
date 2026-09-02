package awsmedialivealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Settings for inserting a DVB Service Description Table (SDT).
//
// Example:
//   var video EncodeConfiguration
//   var audio EncodeConfiguration
//
//
//   medialive.OutputGroupConfiguration_Udp(&UdpOutputGroupProps{
//   	Name: jsii.String("udp_out"),
//   	Destinations: []UdpOutputDestination{
//   		medialive.UdpOutputDestination_Udp(&TransportOutputDestinationProps{
//   			Address: jsii.String("203.0.113.5"),
//   			Port: jsii.Number(5000),
//   		}),
//   	},
//   	Outputs: []UdpOutputDefinition{
//   		&UdpOutputDefinition{
//   			Encodes: []EncodeConfiguration{
//   				video,
//   				audio,
//   			},
//   			OutputName: jsii.String("ts"),
//   			M2tsSettings: medialive.M2tsSettings_Of(&M2tsSettingsProps{
//   				Bitrate: awscdk.Bitrate_Mbps(jsii.Number(8)),
//   				RateMode: medialive.M2tsRateMode_VBR(),
//   				ProgramNum: jsii.Number(1),
//   				PatInterval: awscdk.Duration_Millis(jsii.Number(100)),
//   				PmtInterval: awscdk.Duration_*Millis(jsii.Number(100)),
//   				Scte35Control: medialive.M2tsScte35Control_PASSTHROUGH(),
//   				DvbSdtSettings: &DvbSdtSettings{
//   					OutputSdt: medialive.DvbSdtOutputMode_SDT_MANUAL(),
//   					ServiceName: jsii.String("My Service"),
//   					RepInterval: awscdk.Duration_*Millis(jsii.Number(2000)),
//   				},
//   			}),
//   		},
//   	},
//   })
//
// Experimental.
type DvbSdtSettings struct {
	// The method of inserting SDT information into the output stream.
	// Default: - service default.
	//
	// Experimental.
	OutputSdt DvbSdtOutputMode `field:"optional" json:"outputSdt" yaml:"outputSdt"`
	// The interval between instances of this table in the output transport stream.
	// Default: - service default.
	//
	// Experimental.
	RepInterval awscdk.Duration `field:"optional" json:"repInterval" yaml:"repInterval"`
	// The service name placed in the serviceDescriptor in the SDT (max 256 characters).
	// Default: - no service name.
	//
	// Experimental.
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
	// The service provider name placed in the serviceDescriptor in the SDT (max 256 characters).
	// Default: - no service provider name.
	//
	// Experimental.
	ServiceProviderName *string `field:"optional" json:"serviceProviderName" yaml:"serviceProviderName"`
}

