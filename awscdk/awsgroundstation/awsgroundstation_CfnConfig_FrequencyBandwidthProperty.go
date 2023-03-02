package awsgroundstation


// Defines a bandwidth.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   frequencyBandwidthProperty := &frequencyBandwidthProperty{
//   	units: jsii.String("units"),
//   	value: jsii.Number(123),
//   }
//
type CfnConfig_FrequencyBandwidthProperty struct {
	// The units of the bandwidth.
	Units *string `field:"optional" json:"units" yaml:"units"`
	// The value of the bandwidth. AWS Ground Station currently has the following bandwidth limitations:.
	//
	// - For `AntennaDownlinkDemodDecodeconfig` , valid values are between 125 kHz to 650 MHz.
	// - For `AntennaDownlinkconfig` , valid values are between 10 kHz to 54 MHz.
	// - For `AntennaUplinkConfig` , valid values are between 10 kHz to 54 MHz.
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

