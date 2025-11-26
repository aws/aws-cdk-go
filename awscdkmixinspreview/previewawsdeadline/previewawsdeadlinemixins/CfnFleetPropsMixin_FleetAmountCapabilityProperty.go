package previewawsdeadlinemixins


// The fleet amount and attribute capabilities.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   fleetAmountCapabilityProperty := &FleetAmountCapabilityProperty{
//   	Max: jsii.Number(123),
//   	Min: jsii.Number(123),
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-fleetamountcapability.html
//
type CfnFleetPropsMixin_FleetAmountCapabilityProperty struct {
	// The maximum amount of the fleet worker capability.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-fleetamountcapability.html#cfn-deadline-fleet-fleetamountcapability-max
	//
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// The minimum amount of fleet worker capability.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-fleetamountcapability.html#cfn-deadline-fleet-fleetamountcapability-min
	//
	Min *float64 `field:"optional" json:"min" yaml:"min"`
	// The name of the fleet capability.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-fleetamountcapability.html#cfn-deadline-fleet-fleetamountcapability-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

