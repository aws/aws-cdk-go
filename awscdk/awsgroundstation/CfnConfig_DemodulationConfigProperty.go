package awsgroundstation


// Defines demodulation settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   demodulationConfigProperty := &DemodulationConfigProperty{
//   	UnvalidatedJson: jsii.String("unvalidatedJson"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-demodulationconfig.html
//
type CfnConfig_DemodulationConfigProperty struct {
	// The demodulation settings are in JSON format and define parameters for demodulation, for example which modulation scheme (e.g. PSK, QPSK, etc.) and matched filter to use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-demodulationconfig.html#cfn-groundstation-config-demodulationconfig-unvalidatedjson
	//
	UnvalidatedJson *string `field:"optional" json:"unvalidatedJson" yaml:"unvalidatedJson"`
}

