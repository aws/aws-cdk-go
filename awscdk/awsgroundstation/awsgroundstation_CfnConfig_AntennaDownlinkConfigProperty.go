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
//   antennaDownlinkConfigProperty := &antennaDownlinkConfigProperty{
//   	spectrumConfig: &spectrumConfigProperty{
//   		bandwidth: &frequencyBandwidthProperty{
//   			units: jsii.String("units"),
//   			value: jsii.Number(123),
//   		},
//   		centerFrequency: &frequencyProperty{
//   			units: jsii.String("units"),
//   			value: jsii.Number(123),
//   		},
//   		polarization: jsii.String("polarization"),
//   	},
//   }
//
type CfnConfig_AntennaDownlinkConfigProperty struct {
	// Defines the spectrum configuration.
	SpectrumConfig interface{} `field:"optional" json:"spectrumConfig" yaml:"spectrumConfig"`
}

