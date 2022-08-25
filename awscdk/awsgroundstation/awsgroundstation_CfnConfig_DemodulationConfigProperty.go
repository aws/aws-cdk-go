package awsgroundstation


// Defines demodulation settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   demodulationConfigProperty := &demodulationConfigProperty{
//   	unvalidatedJson: jsii.String("unvalidatedJson"),
//   }
//
type CfnConfig_DemodulationConfigProperty struct {
	// The demodulation settings are in JSON format and define parameters for demodulation, for example which modulation scheme (e.g. PSK, QPSK, etc.) and matched filter to use.
	UnvalidatedJson *string `field:"optional" json:"unvalidatedJson" yaml:"unvalidatedJson"`
}

