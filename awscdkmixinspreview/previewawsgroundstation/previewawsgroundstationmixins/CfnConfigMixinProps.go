package previewawsgroundstationmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnConfigPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnConfigMixinProps := &CfnConfigMixinProps{
//   	ConfigData: &ConfigDataProperty{
//   		AntennaDownlinkConfig: &AntennaDownlinkConfigProperty{
//   			SpectrumConfig: &SpectrumConfigProperty{
//   				Bandwidth: &FrequencyBandwidthProperty{
//   					Units: jsii.String("units"),
//   					Value: jsii.Number(123),
//   				},
//   				CenterFrequency: &FrequencyProperty{
//   					Units: jsii.String("units"),
//   					Value: jsii.Number(123),
//   				},
//   				Polarization: jsii.String("polarization"),
//   			},
//   		},
//   		AntennaDownlinkDemodDecodeConfig: &AntennaDownlinkDemodDecodeConfigProperty{
//   			DecodeConfig: &DecodeConfigProperty{
//   				UnvalidatedJson: jsii.String("unvalidatedJson"),
//   			},
//   			DemodulationConfig: &DemodulationConfigProperty{
//   				UnvalidatedJson: jsii.String("unvalidatedJson"),
//   			},
//   			SpectrumConfig: &SpectrumConfigProperty{
//   				Bandwidth: &FrequencyBandwidthProperty{
//   					Units: jsii.String("units"),
//   					Value: jsii.Number(123),
//   				},
//   				CenterFrequency: &FrequencyProperty{
//   					Units: jsii.String("units"),
//   					Value: jsii.Number(123),
//   				},
//   				Polarization: jsii.String("polarization"),
//   			},
//   		},
//   		AntennaUplinkConfig: &AntennaUplinkConfigProperty{
//   			SpectrumConfig: &UplinkSpectrumConfigProperty{
//   				CenterFrequency: &FrequencyProperty{
//   					Units: jsii.String("units"),
//   					Value: jsii.Number(123),
//   				},
//   				Polarization: jsii.String("polarization"),
//   			},
//   			TargetEirp: &EirpProperty{
//   				Units: jsii.String("units"),
//   				Value: jsii.Number(123),
//   			},
//   			TransmitDisabled: jsii.Boolean(false),
//   		},
//   		DataflowEndpointConfig: &DataflowEndpointConfigProperty{
//   			DataflowEndpointName: jsii.String("dataflowEndpointName"),
//   			DataflowEndpointRegion: jsii.String("dataflowEndpointRegion"),
//   		},
//   		S3RecordingConfig: &S3RecordingConfigProperty{
//   			BucketArn: jsii.String("bucketArn"),
//   			Prefix: jsii.String("prefix"),
//   			RoleArn: jsii.String("roleArn"),
//   		},
//   		TelemetrySinkConfig: &TelemetrySinkConfigProperty{
//   			TelemetrySinkData: &TelemetrySinkDataProperty{
//   				KinesisDataStreamData: &KinesisDataStreamDataProperty{
//   					KinesisDataStreamArn: jsii.String("kinesisDataStreamArn"),
//   					KinesisRoleArn: jsii.String("kinesisRoleArn"),
//   				},
//   			},
//   			TelemetrySinkType: jsii.String("telemetrySinkType"),
//   		},
//   		TrackingConfig: &TrackingConfigProperty{
//   			Autotrack: jsii.String("autotrack"),
//   		},
//   		UplinkEchoConfig: &UplinkEchoConfigProperty{
//   			AntennaUplinkConfigArn: jsii.String("antennaUplinkConfigArn"),
//   			Enabled: jsii.Boolean(false),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-groundstation-config.html
//
type CfnConfigMixinProps struct {
	// Object containing the parameters of a config.
	//
	// Only one subtype may be specified per config. See the subtype definitions for a description of each config subtype.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-groundstation-config.html#cfn-groundstation-config-configdata
	//
	ConfigData interface{} `field:"optional" json:"configData" yaml:"configData"`
	// The name of the config object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-groundstation-config.html#cfn-groundstation-config-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Tags assigned to a resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-groundstation-config.html#cfn-groundstation-config-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

