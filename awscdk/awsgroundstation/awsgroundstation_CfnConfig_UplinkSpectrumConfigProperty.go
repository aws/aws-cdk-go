package awsgroundstation


// Defines a uplink spectrum.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   uplinkSpectrumConfigProperty := &uplinkSpectrumConfigProperty{
//   	centerFrequency: &frequencyProperty{
//   		units: jsii.String("units"),
//   		value: jsii.Number(123),
//   	},
//   	polarization: jsii.String("polarization"),
//   }
//
type CfnConfig_UplinkSpectrumConfigProperty struct {
	// The center frequency of the spectrum.
	//
	// Valid values are between 2200 to 2300 MHz and 7750 to 8400 MHz for downlink and 2025 to 2120 MHz for uplink.
	CenterFrequency interface{} `field:"optional" json:"centerFrequency" yaml:"centerFrequency"`
	// The polarization of the spectrum.
	//
	// Valid values are `"RIGHT_HAND"` and `"LEFT_HAND"` .
	Polarization *string `field:"optional" json:"polarization" yaml:"polarization"`
}

