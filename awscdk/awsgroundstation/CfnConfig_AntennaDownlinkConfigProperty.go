package awsgroundstation


// Provides information about how AWS Ground Station should configure an antenna for downlink during a contact.
//
// Use an antenna downlink config in a mission profile to receive the downlink data in raw DigIF format.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   antennaDownlinkConfigProperty := &AntennaDownlinkConfigProperty{
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
type CfnConfig_AntennaDownlinkConfigProperty struct {
	// Defines the spectrum configuration.
	SpectrumConfig interface{} `field:"optional" json:"spectrumConfig" yaml:"spectrumConfig"`
}

