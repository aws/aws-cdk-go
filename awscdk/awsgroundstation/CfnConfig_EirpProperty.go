package awsgroundstation


// Defines an equivalent isotropically radiated power (EIRP).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eirpProperty := &EirpProperty{
//   	Units: jsii.String("units"),
//   	Value: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-eirp.html
//
type CfnConfig_EirpProperty struct {
	// The units of the EIRP.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-eirp.html#cfn-groundstation-config-eirp-units
	//
	Units *string `field:"optional" json:"units" yaml:"units"`
	// The value of the EIRP.
	//
	// Valid values are between 20.0 to 50.0 dBW.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-eirp.html#cfn-groundstation-config-eirp-value
	//
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

