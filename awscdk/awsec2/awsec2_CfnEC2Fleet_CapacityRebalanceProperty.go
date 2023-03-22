package awsec2


// The Spot Instance replacement strategy to use when Amazon EC2 emits a rebalance notification signal that your Spot Instance is at an elevated risk of being interrupted.
//
// For more information, see [Capacity rebalancing](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-fleet-capacity-rebalance.html) in the *Amazon EC2 User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   capacityRebalanceProperty := &capacityRebalanceProperty{
//   	replacementStrategy: jsii.String("replacementStrategy"),
//   	terminationDelay: jsii.Number(123),
//   }
//
type CfnEC2Fleet_CapacityRebalanceProperty struct {
	// The replacement strategy to use. Only available for fleets of type `maintain` .
	//
	// `launch` - EC2 Fleet launches a replacement Spot Instance when a rebalance notification is emitted for an existing Spot Instance in the fleet. EC2 Fleet does not terminate the instances that receive a rebalance notification. You can terminate the old instances, or you can leave them running. You are charged for all instances while they are running.
	//
	// `launch-before-terminate` - EC2 Fleet launches a replacement Spot Instance when a rebalance notification is emitted for an existing Spot Instance in the fleet, and then, after a delay that you specify (in `TerminationDelay` ), terminates the instances that received a rebalance notification.
	ReplacementStrategy *string `field:"optional" json:"replacementStrategy" yaml:"replacementStrategy"`
	// The amount of time (in seconds) that Amazon EC2 waits before terminating the old Spot Instance after launching a new replacement Spot Instance.
	//
	// Required when `ReplacementStrategy` is set to `launch-before-terminate` .
	//
	// Not valid when `ReplacementStrategy` is set to `launch` .
	//
	// Valid values: Minimum value of `120` seconds. Maximum value of `7200` seconds.
	TerminationDelay *float64 `field:"optional" json:"terminationDelay" yaml:"terminationDelay"`
}

