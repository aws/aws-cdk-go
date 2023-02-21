package awsgroundstation


// Provides information about how AWS Ground Station should configure an antenna for uplink during a contact.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   antennaUplinkConfigProperty := &antennaUplinkConfigProperty{
//   	spectrumConfig: &uplinkSpectrumConfigProperty{
//   		centerFrequency: &frequencyProperty{
//   			units: jsii.String("units"),
//   			value: jsii.Number(123),
//   		},
//   		polarization: jsii.String("polarization"),
//   	},
//   	targetEirp: &eirpProperty{
//   		units: jsii.String("units"),
//   		value: jsii.Number(123),
//   	},
//   	transmitDisabled: jsii.Boolean(false),
//   }
//
type CfnConfig_AntennaUplinkConfigProperty struct {
	// Defines the spectrum configuration.
	SpectrumConfig interface{} `field:"optional" json:"spectrumConfig" yaml:"spectrumConfig"`
	// The equivalent isotropically radiated power (EIRP) to use for uplink transmissions.
	//
	// Valid values are between 20.0 to 50.0 dBW.
	TargetEirp interface{} `field:"optional" json:"targetEirp" yaml:"targetEirp"`
	// Whether or not uplink transmit is disabled.
	TransmitDisabled interface{} `field:"optional" json:"transmitDisabled" yaml:"transmitDisabled"`
}

