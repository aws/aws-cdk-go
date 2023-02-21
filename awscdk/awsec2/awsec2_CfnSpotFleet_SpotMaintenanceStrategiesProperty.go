package awsec2


// The strategies for managing your Spot Instances that are at an elevated risk of being interrupted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   spotMaintenanceStrategiesProperty := &spotMaintenanceStrategiesProperty{
//   	capacityRebalance: &spotCapacityRebalanceProperty{
//   		replacementStrategy: jsii.String("replacementStrategy"),
//   		terminationDelay: jsii.Number(123),
//   	},
//   }
//
type CfnSpotFleet_SpotMaintenanceStrategiesProperty struct {
	// The Spot Instance replacement strategy to use when Amazon EC2 emits a signal that your Spot Instance is at an elevated risk of being interrupted.
	//
	// For more information, see [Capacity rebalancing](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-fleet-capacity-rebalance.html) in the *Amazon EC2 User Guide for Linux Instances* .
	CapacityRebalance interface{} `field:"optional" json:"capacityRebalance" yaml:"capacityRebalance"`
}

