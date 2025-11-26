package previewawsgroundstationmixins


// Config objects provide information to Ground Station about how to configure the antenna and how data flows during a contact.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   configDataProperty := &ConfigDataProperty{
//   	AntennaDownlinkConfig: &AntennaDownlinkConfigProperty{
//   		SpectrumConfig: &SpectrumConfigProperty{
//   			Bandwidth: &FrequencyBandwidthProperty{
//   				Units: jsii.String("units"),
//   				Value: jsii.Number(123),
//   			},
//   			CenterFrequency: &FrequencyProperty{
//   				Units: jsii.String("units"),
//   				Value: jsii.Number(123),
//   			},
//   			Polarization: jsii.String("polarization"),
//   		},
//   	},
//   	AntennaDownlinkDemodDecodeConfig: &AntennaDownlinkDemodDecodeConfigProperty{
//   		DecodeConfig: &DecodeConfigProperty{
//   			UnvalidatedJson: jsii.String("unvalidatedJson"),
//   		},
//   		DemodulationConfig: &DemodulationConfigProperty{
//   			UnvalidatedJson: jsii.String("unvalidatedJson"),
//   		},
//   		SpectrumConfig: &SpectrumConfigProperty{
//   			Bandwidth: &FrequencyBandwidthProperty{
//   				Units: jsii.String("units"),
//   				Value: jsii.Number(123),
//   			},
//   			CenterFrequency: &FrequencyProperty{
//   				Units: jsii.String("units"),
//   				Value: jsii.Number(123),
//   			},
//   			Polarization: jsii.String("polarization"),
//   		},
//   	},
//   	AntennaUplinkConfig: &AntennaUplinkConfigProperty{
//   		SpectrumConfig: &UplinkSpectrumConfigProperty{
//   			CenterFrequency: &FrequencyProperty{
//   				Units: jsii.String("units"),
//   				Value: jsii.Number(123),
//   			},
//   			Polarization: jsii.String("polarization"),
//   		},
//   		TargetEirp: &EirpProperty{
//   			Units: jsii.String("units"),
//   			Value: jsii.Number(123),
//   		},
//   		TransmitDisabled: jsii.Boolean(false),
//   	},
//   	DataflowEndpointConfig: &DataflowEndpointConfigProperty{
//   		DataflowEndpointName: jsii.String("dataflowEndpointName"),
//   		DataflowEndpointRegion: jsii.String("dataflowEndpointRegion"),
//   	},
//   	S3RecordingConfig: &S3RecordingConfigProperty{
//   		BucketArn: jsii.String("bucketArn"),
//   		Prefix: jsii.String("prefix"),
//   		RoleArn: jsii.String("roleArn"),
//   	},
//   	TrackingConfig: &TrackingConfigProperty{
//   		Autotrack: jsii.String("autotrack"),
//   	},
//   	UplinkEchoConfig: &UplinkEchoConfigProperty{
//   		AntennaUplinkConfigArn: jsii.String("antennaUplinkConfigArn"),
//   		Enabled: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-configdata.html
//
type CfnConfigPropsMixin_ConfigDataProperty struct {
	// Provides information for an antenna downlink config object.
	//
	// Antenna downlink config objects are used to provide parameters for downlinks where no demodulation or decoding is performed by Ground Station (RF over IP downlinks).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-configdata.html#cfn-groundstation-config-configdata-antennadownlinkconfig
	//
	AntennaDownlinkConfig interface{} `field:"optional" json:"antennaDownlinkConfig" yaml:"antennaDownlinkConfig"`
	// Provides information for a downlink demod decode config object.
	//
	// Downlink demod decode config objects are used to provide parameters for downlinks where the Ground Station service will demodulate and decode the downlinked data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-configdata.html#cfn-groundstation-config-configdata-antennadownlinkdemoddecodeconfig
	//
	AntennaDownlinkDemodDecodeConfig interface{} `field:"optional" json:"antennaDownlinkDemodDecodeConfig" yaml:"antennaDownlinkDemodDecodeConfig"`
	// Provides information for an uplink config object.
	//
	// Uplink config objects are used to provide parameters for uplink contacts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-configdata.html#cfn-groundstation-config-configdata-antennauplinkconfig
	//
	AntennaUplinkConfig interface{} `field:"optional" json:"antennaUplinkConfig" yaml:"antennaUplinkConfig"`
	// Provides information for a dataflow endpoint config object.
	//
	// Dataflow endpoint config objects are used to provide parameters about which IP endpoint(s) to use during a contact. Dataflow endpoints are where Ground Station sends data during a downlink contact and where Ground Station receives data to send to the satellite during an uplink contact.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-configdata.html#cfn-groundstation-config-configdata-dataflowendpointconfig
	//
	DataflowEndpointConfig interface{} `field:"optional" json:"dataflowEndpointConfig" yaml:"dataflowEndpointConfig"`
	// Provides information for an S3 recording config object.
	//
	// S3 recording config objects are used to provide parameters for S3 recording during downlink contacts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-configdata.html#cfn-groundstation-config-configdata-s3recordingconfig
	//
	S3RecordingConfig interface{} `field:"optional" json:"s3RecordingConfig" yaml:"s3RecordingConfig"`
	// Provides information for a tracking config object.
	//
	// Tracking config objects are used to provide parameters about how to track the satellite through the sky during a contact.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-configdata.html#cfn-groundstation-config-configdata-trackingconfig
	//
	TrackingConfig interface{} `field:"optional" json:"trackingConfig" yaml:"trackingConfig"`
	// Provides information for an uplink echo config object.
	//
	// Uplink echo config objects are used to provide parameters for uplink echo during uplink contacts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-configdata.html#cfn-groundstation-config-configdata-uplinkechoconfig
	//
	UplinkEchoConfig interface{} `field:"optional" json:"uplinkEchoConfig" yaml:"uplinkEchoConfig"`
}

