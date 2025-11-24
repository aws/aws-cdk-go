package mixinsawsgroundstation


// Provides information about how AWS Ground Station should configure an antenna for uplink during a contact.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   antennaUplinkConfigProperty := &AntennaUplinkConfigProperty{
//   	SpectrumConfig: &UplinkSpectrumConfigProperty{
//   		CenterFrequency: &FrequencyProperty{
//   			Units: jsii.String("units"),
//   			Value: jsii.Number(123),
//   		},
//   		Polarization: jsii.String("polarization"),
//   	},
//   	TargetEirp: &EirpProperty{
//   		Units: jsii.String("units"),
//   		Value: jsii.Number(123),
//   	},
//   	TransmitDisabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-antennauplinkconfig.html
//
type CfnConfigPropsMixin_AntennaUplinkConfigProperty struct {
	// Defines the spectrum configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-antennauplinkconfig.html#cfn-groundstation-config-antennauplinkconfig-spectrumconfig
	//
	SpectrumConfig interface{} `field:"optional" json:"spectrumConfig" yaml:"spectrumConfig"`
	// The equivalent isotropically radiated power (EIRP) to use for uplink transmissions.
	//
	// Valid values are between 20.0 to 50.0 dBW.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-antennauplinkconfig.html#cfn-groundstation-config-antennauplinkconfig-targeteirp
	//
	TargetEirp interface{} `field:"optional" json:"targetEirp" yaml:"targetEirp"`
	// Whether or not uplink transmit is disabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-antennauplinkconfig.html#cfn-groundstation-config-antennauplinkconfig-transmitdisabled
	//
	TransmitDisabled interface{} `field:"optional" json:"transmitDisabled" yaml:"transmitDisabled"`
}

