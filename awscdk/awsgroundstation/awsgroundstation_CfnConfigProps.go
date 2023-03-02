package awsgroundstation

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnConfig`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConfigProps := &cfnConfigProps{
//   	configData: &configDataProperty{
//   		antennaDownlinkConfig: &antennaDownlinkConfigProperty{
//   			spectrumConfig: &spectrumConfigProperty{
//   				bandwidth: &frequencyBandwidthProperty{
//   					units: jsii.String("units"),
//   					value: jsii.Number(123),
//   				},
//   				centerFrequency: &frequencyProperty{
//   					units: jsii.String("units"),
//   					value: jsii.Number(123),
//   				},
//   				polarization: jsii.String("polarization"),
//   			},
//   		},
//   		antennaDownlinkDemodDecodeConfig: &antennaDownlinkDemodDecodeConfigProperty{
//   			decodeConfig: &decodeConfigProperty{
//   				unvalidatedJson: jsii.String("unvalidatedJson"),
//   			},
//   			demodulationConfig: &demodulationConfigProperty{
//   				unvalidatedJson: jsii.String("unvalidatedJson"),
//   			},
//   			spectrumConfig: &spectrumConfigProperty{
//   				bandwidth: &frequencyBandwidthProperty{
//   					units: jsii.String("units"),
//   					value: jsii.Number(123),
//   				},
//   				centerFrequency: &frequencyProperty{
//   					units: jsii.String("units"),
//   					value: jsii.Number(123),
//   				},
//   				polarization: jsii.String("polarization"),
//   			},
//   		},
//   		antennaUplinkConfig: &antennaUplinkConfigProperty{
//   			spectrumConfig: &uplinkSpectrumConfigProperty{
//   				centerFrequency: &frequencyProperty{
//   					units: jsii.String("units"),
//   					value: jsii.Number(123),
//   				},
//   				polarization: jsii.String("polarization"),
//   			},
//   			targetEirp: &eirpProperty{
//   				units: jsii.String("units"),
//   				value: jsii.Number(123),
//   			},
//   			transmitDisabled: jsii.Boolean(false),
//   		},
//   		dataflowEndpointConfig: &dataflowEndpointConfigProperty{
//   			dataflowEndpointName: jsii.String("dataflowEndpointName"),
//   			dataflowEndpointRegion: jsii.String("dataflowEndpointRegion"),
//   		},
//   		s3RecordingConfig: &s3RecordingConfigProperty{
//   			bucketArn: jsii.String("bucketArn"),
//   			prefix: jsii.String("prefix"),
//   			roleArn: jsii.String("roleArn"),
//   		},
//   		trackingConfig: &trackingConfigProperty{
//   			autotrack: jsii.String("autotrack"),
//   		},
//   		uplinkEchoConfig: &uplinkEchoConfigProperty{
//   			antennaUplinkConfigArn: jsii.String("antennaUplinkConfigArn"),
//   			enabled: jsii.Boolean(false),
//   		},
//   	},
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnConfigProps struct {
	// Object containing the parameters of a config.
	//
	// Only one subtype may be specified per config. See the subtype definitions for a description of each config subtype.
	ConfigData interface{} `field:"required" json:"configData" yaml:"configData"`
	// The name of the config object.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Tags assigned to a resource.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

