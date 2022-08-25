package awsgroundstation


// Defines decoding settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   decodeConfigProperty := &decodeConfigProperty{
//   	unvalidatedJson: jsii.String("unvalidatedJson"),
//   }
//
type CfnConfig_DecodeConfigProperty struct {
	// The decoding settings are in JSON format and define a set of steps to perform to decode the data.
	UnvalidatedJson *string `field:"optional" json:"unvalidatedJson" yaml:"unvalidatedJson"`
}

