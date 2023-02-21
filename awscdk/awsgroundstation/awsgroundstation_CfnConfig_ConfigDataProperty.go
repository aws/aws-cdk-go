package awsgroundstation


// Config objects provide information to Ground Station about how to configure the antenna and how data flows during a contact.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   configDataProperty := &configDataProperty{
//   	antennaDownlinkConfig: &antennaDownlinkConfigProperty{
//   		spectrumConfig: &spectrumConfigProperty{
//   			bandwidth: &frequencyBandwidthProperty{
//   				units: jsii.String("units"),
//   				value: jsii.Number(123),
//   			},
//   			centerFrequency: &frequencyProperty{
//   				units: jsii.String("units"),
//   				value: jsii.Number(123),
//   			},
//   			polarization: jsii.String("polarization"),
//   		},
//   	},
//   	antennaDownlinkDemodDecodeConfig: &antennaDownlinkDemodDecodeConfigProperty{
//   		decodeConfig: &decodeConfigProperty{
//   			unvalidatedJson: jsii.String("unvalidatedJson"),
//   		},
//   		demodulationConfig: &demodulationConfigProperty{
//   			unvalidatedJson: jsii.String("unvalidatedJson"),
//   		},
//   		spectrumConfig: &spectrumConfigProperty{
//   			bandwidth: &frequencyBandwidthProperty{
//   				units: jsii.String("units"),
//   				value: jsii.Number(123),
//   			},
//   			centerFrequency: &frequencyProperty{
//   				units: jsii.String("units"),
//   				value: jsii.Number(123),
//   			},
//   			polarization: jsii.String("polarization"),
//   		},
//   	},
//   	antennaUplinkConfig: &antennaUplinkConfigProperty{
//   		spectrumConfig: &uplinkSpectrumConfigProperty{
//   			centerFrequency: &frequencyProperty{
//   				units: jsii.String("units"),
//   				value: jsii.Number(123),
//   			},
//   			polarization: jsii.String("polarization"),
//   		},
//   		targetEirp: &eirpProperty{
//   			units: jsii.String("units"),
//   			value: jsii.Number(123),
//   		},
//   		transmitDisabled: jsii.Boolean(false),
//   	},
//   	dataflowEndpointConfig: &dataflowEndpointConfigProperty{
//   		dataflowEndpointName: jsii.String("dataflowEndpointName"),
//   		dataflowEndpointRegion: jsii.String("dataflowEndpointRegion"),
//   	},
//   	s3RecordingConfig: &s3RecordingConfigProperty{
//   		bucketArn: jsii.String("bucketArn"),
//   		prefix: jsii.String("prefix"),
//   		roleArn: jsii.String("roleArn"),
//   	},
//   	trackingConfig: &trackingConfigProperty{
//   		autotrack: jsii.String("autotrack"),
//   	},
//   	uplinkEchoConfig: &uplinkEchoConfigProperty{
//   		antennaUplinkConfigArn: jsii.String("antennaUplinkConfigArn"),
//   		enabled: jsii.Boolean(false),
//   	},
//   }
//
type CfnConfig_ConfigDataProperty struct {
	// Provides information for an antenna downlink config object.
	//
	// Antenna downlink config objects are used to provide parameters for downlinks where no demodulation or decoding is performed by Ground Station (RF over IP downlinks).
	AntennaDownlinkConfig interface{} `field:"optional" json:"antennaDownlinkConfig" yaml:"antennaDownlinkConfig"`
	// Provides information for a downlink demod decode config object.
	//
	// Downlink demod decode config objects are used to provide parameters for downlinks where the Ground Station service will demodulate and decode the downlinked data.
	AntennaDownlinkDemodDecodeConfig interface{} `field:"optional" json:"antennaDownlinkDemodDecodeConfig" yaml:"antennaDownlinkDemodDecodeConfig"`
	// Provides information for an uplink config object.
	//
	// Uplink config objects are used to provide parameters for uplink contacts.
	AntennaUplinkConfig interface{} `field:"optional" json:"antennaUplinkConfig" yaml:"antennaUplinkConfig"`
	// Provides information for a dataflow endpoint config object.
	//
	// Dataflow endpoint config objects are used to provide parameters about which IP endpoint(s) to use during a contact. Dataflow endpoints are where Ground Station sends data during a downlink contact and where Ground Station receives data to send to the satellite during an uplink contact.
	DataflowEndpointConfig interface{} `field:"optional" json:"dataflowEndpointConfig" yaml:"dataflowEndpointConfig"`
	// Provides information for an S3 recording config object.
	//
	// S3 recording config objects are used to provide parameters for S3 recording during downlink contacts.
	S3RecordingConfig interface{} `field:"optional" json:"s3RecordingConfig" yaml:"s3RecordingConfig"`
	// Provides information for a tracking config object.
	//
	// Tracking config objects are used to provide parameters about how to track the satellite through the sky during a contact.
	TrackingConfig interface{} `field:"optional" json:"trackingConfig" yaml:"trackingConfig"`
	// Provides information for an uplink echo config object.
	//
	// Uplink echo config objects are used to provide parameters for uplink echo during uplink contacts.
	UplinkEchoConfig interface{} `field:"optional" json:"uplinkEchoConfig" yaml:"uplinkEchoConfig"`
}

