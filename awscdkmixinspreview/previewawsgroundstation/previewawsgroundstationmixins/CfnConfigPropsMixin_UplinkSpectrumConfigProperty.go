package previewawsgroundstationmixins


// Defines a uplink spectrum.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   uplinkSpectrumConfigProperty := &UplinkSpectrumConfigProperty{
//   	CenterFrequency: &FrequencyProperty{
//   		Units: jsii.String("units"),
//   		Value: jsii.Number(123),
//   	},
//   	Polarization: jsii.String("polarization"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-uplinkspectrumconfig.html
//
type CfnConfigPropsMixin_UplinkSpectrumConfigProperty struct {
	// The center frequency of the spectrum.
	//
	// Valid values are between 2200 to 2300 MHz and 7750 to 8400 MHz for downlink and 2025 to 2120 MHz for uplink.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-uplinkspectrumconfig.html#cfn-groundstation-config-uplinkspectrumconfig-centerfrequency
	//
	CenterFrequency interface{} `field:"optional" json:"centerFrequency" yaml:"centerFrequency"`
	// The polarization of the spectrum.
	//
	// Valid values are `"RIGHT_HAND"` and `"LEFT_HAND"` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-uplinkspectrumconfig.html#cfn-groundstation-config-uplinkspectrumconfig-polarization
	//
	Polarization *string `field:"optional" json:"polarization" yaml:"polarization"`
}

