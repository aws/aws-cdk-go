package mixinsawsgroundstation


// Provides information about how AWS Ground Station should configure an antenna for downlink during a contact.
//
// Use an antenna downlink demod decode config in a mission profile to receive the downlink data that has been demodulated and decoded.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   antennaDownlinkDemodDecodeConfigProperty := &AntennaDownlinkDemodDecodeConfigProperty{
//   	DecodeConfig: &DecodeConfigProperty{
//   		UnvalidatedJson: jsii.String("unvalidatedJson"),
//   	},
//   	DemodulationConfig: &DemodulationConfigProperty{
//   		UnvalidatedJson: jsii.String("unvalidatedJson"),
//   	},
//   	SpectrumConfig: &SpectrumConfigProperty{
//   		Bandwidth: &FrequencyBandwidthProperty{
//   			Units: jsii.String("units"),
//   			Value: jsii.Number(123),
//   		},
//   		CenterFrequency: &FrequencyProperty{
//   			Units: jsii.String("units"),
//   			Value: jsii.Number(123),
//   		},
//   		Polarization: jsii.String("polarization"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-antennadownlinkdemoddecodeconfig.html
//
type CfnConfigPropsMixin_AntennaDownlinkDemodDecodeConfigProperty struct {
	// Defines how the RF signal will be decoded.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-antennadownlinkdemoddecodeconfig.html#cfn-groundstation-config-antennadownlinkdemoddecodeconfig-decodeconfig
	//
	DecodeConfig interface{} `field:"optional" json:"decodeConfig" yaml:"decodeConfig"`
	// Defines how the RF signal will be demodulated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-antennadownlinkdemoddecodeconfig.html#cfn-groundstation-config-antennadownlinkdemoddecodeconfig-demodulationconfig
	//
	DemodulationConfig interface{} `field:"optional" json:"demodulationConfig" yaml:"demodulationConfig"`
	// Defines the spectrum configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-antennadownlinkdemoddecodeconfig.html#cfn-groundstation-config-antennadownlinkdemoddecodeconfig-spectrumconfig
	//
	SpectrumConfig interface{} `field:"optional" json:"spectrumConfig" yaml:"spectrumConfig"`
}

