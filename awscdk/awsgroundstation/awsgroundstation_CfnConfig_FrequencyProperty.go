package awsgroundstation


// Defines a frequency.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   frequencyProperty := &frequencyProperty{
//   	units: jsii.String("units"),
//   	value: jsii.Number(123),
//   }
//
type CfnConfig_FrequencyProperty struct {
	// The units of the frequency.
	Units *string `field:"optional" json:"units" yaml:"units"`
	// The value of the frequency.
	//
	// Valid values are between 2200 to 2300 MHz and 7750 to 8400 MHz for downlink and 2025 to 2120 MHz for uplink.
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

