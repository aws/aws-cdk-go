package awsec2


// Specifies the configuration of Spot Instances for an EC2 Fleet.
//
// `SpotOptionsRequest` is a property of the [AWS::EC2::EC2Fleet](https://docs.aws.amazon.com//AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   spotOptionsRequestProperty := &spotOptionsRequestProperty{
//   	allocationStrategy: jsii.String("allocationStrategy"),
//   	instanceInterruptionBehavior: jsii.String("instanceInterruptionBehavior"),
//   	instancePoolsToUseCount: jsii.Number(123),
//   	maintenanceStrategies: &maintenanceStrategiesProperty{
//   		capacityRebalance: &capacityRebalanceProperty{
//   			replacementStrategy: jsii.String("replacementStrategy"),
//   			terminationDelay: jsii.Number(123),
//   		},
//   	},
//   	maxTotalPrice: jsii.String("maxTotalPrice"),
//   	minTargetCapacity: jsii.Number(123),
//   	singleAvailabilityZone: jsii.Boolean(false),
//   	singleInstanceType: jsii.Boolean(false),
//   }
//
type CfnEC2Fleet_SpotOptionsRequestProperty struct {
	// Indicates how to allocate the target Spot Instance capacity across the Spot Instance pools specified by the EC2 Fleet.
	//
	// If the allocation strategy is `lowestPrice` , EC2 Fleet launches instances from the Spot Instance pools with the lowest price. This is the default allocation strategy.
	//
	// If the allocation strategy is `diversified` , EC2 Fleet launches instances from all the Spot Instance pools that you specify.
	//
	// If the allocation strategy is `capacityOptimized` , EC2 Fleet launches instances from Spot Instance pools that are optimally chosen based on the available Spot Instance capacity.
	//
	// *Allowed Values* : `lowestPrice` | `diversified` | `capacityOptimized` | `capacityOptimizedPrioritized`.
	AllocationStrategy *string `field:"optional" json:"allocationStrategy" yaml:"allocationStrategy"`
	// The behavior when a Spot Instance is interrupted.
	//
	// Default: `terminate`.
	InstanceInterruptionBehavior *string `field:"optional" json:"instanceInterruptionBehavior" yaml:"instanceInterruptionBehavior"`
	// The number of Spot pools across which to allocate your target Spot capacity.
	//
	// Supported only when Spot `AllocationStrategy` is set to `lowest-price` . EC2 Fleet selects the cheapest Spot pools and evenly allocates your target Spot capacity across the number of Spot pools that you specify.
	//
	// Note that EC2 Fleet attempts to draw Spot Instances from the number of pools that you specify on a best effort basis. If a pool runs out of Spot capacity before fulfilling your target capacity, EC2 Fleet will continue to fulfill your request by drawing from the next cheapest pool. To ensure that your target capacity is met, you might receive Spot Instances from more than the number of pools that you specified. Similarly, if most of the pools have no Spot capacity, you might receive your full target capacity from fewer than the number of pools that you specified.
	InstancePoolsToUseCount *float64 `field:"optional" json:"instancePoolsToUseCount" yaml:"instancePoolsToUseCount"`
	// The strategies for managing your Spot Instances that are at an elevated risk of being interrupted.
	MaintenanceStrategies interface{} `field:"optional" json:"maintenanceStrategies" yaml:"maintenanceStrategies"`
	// The maximum amount per hour for Spot Instances that you're willing to pay.
	MaxTotalPrice *string `field:"optional" json:"maxTotalPrice" yaml:"maxTotalPrice"`
	// The minimum target capacity for Spot Instances in the fleet.
	//
	// If the minimum target capacity is not reached, the fleet launches no instances.
	//
	// Supported only for fleets of type `instant` .
	//
	// At least one of the following must be specified: `SingleAvailabilityZone` | `SingleInstanceType`.
	MinTargetCapacity *float64 `field:"optional" json:"minTargetCapacity" yaml:"minTargetCapacity"`
	// Indicates that the fleet launches all Spot Instances into a single Availability Zone.
	//
	// Supported only for fleets of type `instant` .
	SingleAvailabilityZone interface{} `field:"optional" json:"singleAvailabilityZone" yaml:"singleAvailabilityZone"`
	// Indicates that the fleet uses a single instance type to launch all Spot Instances in the fleet.
	//
	// Supported only for fleets of type `instant` .
	SingleInstanceType interface{} `field:"optional" json:"singleInstanceType" yaml:"singleInstanceType"`
}

