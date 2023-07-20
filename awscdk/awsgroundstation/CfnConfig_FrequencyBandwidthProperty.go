package awsgroundstation


// Defines a bandwidth.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   frequencyBandwidthProperty := &FrequencyBandwidthProperty{
//   	Units: jsii.String("units"),
//   	Value: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-frequencybandwidth.html
//
type CfnConfig_FrequencyBandwidthProperty struct {
	// The units of the bandwidth.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-frequencybandwidth.html#cfn-groundstation-config-frequencybandwidth-units
	//
	Units *string `field:"optional" json:"units" yaml:"units"`
	// The value of the bandwidth. AWS Ground Station currently has the following bandwidth limitations:.
	//
	// - For `AntennaDownlinkDemodDecodeconfig` , valid values are between 125 kHz to 650 MHz.
	// - For `AntennaDownlinkconfig` , valid values are between 10 kHz to 54 MHz.
	// - For `AntennaUplinkConfig` , valid values are between 10 kHz to 54 MHz.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-frequencybandwidth.html#cfn-groundstation-config-frequencybandwidth-value
	//
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

