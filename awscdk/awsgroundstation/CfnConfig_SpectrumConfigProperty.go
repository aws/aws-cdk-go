package awsgroundstation


// Defines a spectrum.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   spectrumConfigProperty := &SpectrumConfigProperty{
//   	Bandwidth: &FrequencyBandwidthProperty{
//   		Units: jsii.String("units"),
//   		Value: jsii.Number(123),
//   	},
//   	CenterFrequency: &FrequencyProperty{
//   		Units: jsii.String("units"),
//   		Value: jsii.Number(123),
//   	},
//   	Polarization: jsii.String("polarization"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-spectrumconfig.html
//
type CfnConfig_SpectrumConfigProperty struct {
	// The bandwidth of the spectrum. AWS Ground Station currently has the following bandwidth limitations:.
	//
	// - For `AntennaDownlinkDemodDecodeconfig` , valid values are between 125 kHz to 650 MHz.
	// - For `AntennaDownlinkconfig` , valid values are between 10 kHz to 54 MHz.
	// - For `AntennaUplinkConfig` , valid values are between 10 kHz to 54 MHz.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-spectrumconfig.html#cfn-groundstation-config-spectrumconfig-bandwidth
	//
	Bandwidth interface{} `field:"optional" json:"bandwidth" yaml:"bandwidth"`
	// The center frequency of the spectrum.
	//
	// Valid values are between 2200 to 2300 MHz and 7750 to 8400 MHz for downlink and 2025 to 2120 MHz for uplink.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-spectrumconfig.html#cfn-groundstation-config-spectrumconfig-centerfrequency
	//
	CenterFrequency interface{} `field:"optional" json:"centerFrequency" yaml:"centerFrequency"`
	// The polarization of the spectrum.
	//
	// Valid values are `"RIGHT_HAND"` and `"LEFT_HAND"` . Capturing both `"RIGHT_HAND"` and `"LEFT_HAND"` polarization requires two separate configs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-spectrumconfig.html#cfn-groundstation-config-spectrumconfig-polarization
	//
	Polarization *string `field:"optional" json:"polarization" yaml:"polarization"`
}

