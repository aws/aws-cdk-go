package awsgroundstation


// Defines an equivalent isotropically radiated power (EIRP).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eirpProperty := &eirpProperty{
//   	units: jsii.String("units"),
//   	value: jsii.Number(123),
//   }
//
type CfnConfig_EirpProperty struct {
	// The units of the EIRP.
	Units *string `field:"optional" json:"units" yaml:"units"`
	// The value of the EIRP.
	//
	// Valid values are between 20.0 to 50.0 dBW.
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

