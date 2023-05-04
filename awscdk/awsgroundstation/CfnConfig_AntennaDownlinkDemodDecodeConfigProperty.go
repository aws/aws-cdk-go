package awsgroundstation


// Provides information about how AWS Ground Station should configure an antenna for downlink during a contact.
//
// Use an antenna downlink demod decode config in a mission profile to receive the downlink data that has been demodulated and decoded.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
type CfnConfig_AntennaDownlinkDemodDecodeConfigProperty struct {
	// Defines how the RF signal will be decoded.
	DecodeConfig interface{} `field:"optional" json:"decodeConfig" yaml:"decodeConfig"`
	// Defines how the RF signal will be demodulated.
	DemodulationConfig interface{} `field:"optional" json:"demodulationConfig" yaml:"demodulationConfig"`
	// Defines the spectrum configuration.
	SpectrumConfig interface{} `field:"optional" json:"spectrumConfig" yaml:"spectrumConfig"`
}

