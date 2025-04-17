package awsdeadline


// The fleet amount and attribute capabilities.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fleetAmountCapabilityProperty := &FleetAmountCapabilityProperty{
//   	Min: jsii.Number(123),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Max: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-fleetamountcapability.html
//
type CfnFleet_FleetAmountCapabilityProperty struct {
	// The minimum amount of fleet worker capability.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-fleetamountcapability.html#cfn-deadline-fleet-fleetamountcapability-min
	//
	Min *float64 `field:"required" json:"min" yaml:"min"`
	// The name of the fleet capability.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-fleetamountcapability.html#cfn-deadline-fleet-fleetamountcapability-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The maximum amount of the fleet worker capability.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-fleetamountcapability.html#cfn-deadline-fleet-fleetamountcapability-max
	//
	Max *float64 `field:"optional" json:"max" yaml:"max"`
}

