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
//   spotOptionsRequestProperty := &SpotOptionsRequestProperty{
//   	AllocationStrategy: jsii.String("allocationStrategy"),
//   	InstanceInterruptionBehavior: jsii.String("instanceInterruptionBehavior"),
//   	InstancePoolsToUseCount: jsii.Number(123),
//   	MaintenanceStrategies: &MaintenanceStrategiesProperty{
//   		CapacityRebalance: &CapacityRebalanceProperty{
//   			ReplacementStrategy: jsii.String("replacementStrategy"),
//   			TerminationDelay: jsii.Number(123),
//   		},
//   	},
//   	MaxTotalPrice: jsii.String("maxTotalPrice"),
//   	MinTargetCapacity: jsii.Number(123),
//   	SingleAvailabilityZone: jsii.Boolean(false),
//   	SingleInstanceType: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-spotoptionsrequest.html
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-spotoptionsrequest.html#cfn-ec2-ec2fleet-spotoptionsrequest-allocationstrategy
	//
	AllocationStrategy *string `field:"optional" json:"allocationStrategy" yaml:"allocationStrategy"`
	// The behavior when a Spot Instance is interrupted.
	//
	// Default: `terminate`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-spotoptionsrequest.html#cfn-ec2-ec2fleet-spotoptionsrequest-instanceinterruptionbehavior
	//
	InstanceInterruptionBehavior *string `field:"optional" json:"instanceInterruptionBehavior" yaml:"instanceInterruptionBehavior"`
	// The number of Spot pools across which to allocate your target Spot capacity.
	//
	// Supported only when Spot `AllocationStrategy` is set to `lowest-price` . EC2 Fleet selects the cheapest Spot pools and evenly allocates your target Spot capacity across the number of Spot pools that you specify.
	//
	// Note that EC2 Fleet attempts to draw Spot Instances from the number of pools that you specify on a best effort basis. If a pool runs out of Spot capacity before fulfilling your target capacity, EC2 Fleet will continue to fulfill your request by drawing from the next cheapest pool. To ensure that your target capacity is met, you might receive Spot Instances from more than the number of pools that you specified. Similarly, if most of the pools have no Spot capacity, you might receive your full target capacity from fewer than the number of pools that you specified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-spotoptionsrequest.html#cfn-ec2-ec2fleet-spotoptionsrequest-instancepoolstousecount
	//
	InstancePoolsToUseCount *float64 `field:"optional" json:"instancePoolsToUseCount" yaml:"instancePoolsToUseCount"`
	// The strategies for managing your Spot Instances that are at an elevated risk of being interrupted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-spotoptionsrequest.html#cfn-ec2-ec2fleet-spotoptionsrequest-maintenancestrategies
	//
	MaintenanceStrategies interface{} `field:"optional" json:"maintenanceStrategies" yaml:"maintenanceStrategies"`
	// The maximum amount per hour for Spot Instances that you're willing to pay.
	//
	// We do not recommend using this parameter because it can lead to increased interruptions. If you do not specify this parameter, you will pay the current Spot price.
	//
	// > If you specify a maximum price, your Spot Instances will be interrupted more frequently than if you do not specify this parameter. > If your fleet includes T instances that are configured as `unlimited` , and if their average CPU usage exceeds the baseline utilization, you will incur a charge for surplus credits. The `MaxTotalPrice` does not account for surplus credits, and, if you use surplus credits, your final cost might be higher than what you specified for `MaxTotalPrice` . For more information, see [Surplus credits can incur charges](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances-unlimited-mode-concepts.html#unlimited-mode-surplus-credits) in the *Amazon EC2 User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-spotoptionsrequest.html#cfn-ec2-ec2fleet-spotoptionsrequest-maxtotalprice
	//
	MaxTotalPrice *string `field:"optional" json:"maxTotalPrice" yaml:"maxTotalPrice"`
	// The minimum target capacity for Spot Instances in the fleet.
	//
	// If this minimum capacity isn't reached, no instances are launched.
	//
	// Constraints: Maximum value of `1000` . Supported only for fleets of type `instant` .
	//
	// At least one of the following must be specified: `SingleAvailabilityZone` | `SingleInstanceType`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-spotoptionsrequest.html#cfn-ec2-ec2fleet-spotoptionsrequest-mintargetcapacity
	//
	MinTargetCapacity *float64 `field:"optional" json:"minTargetCapacity" yaml:"minTargetCapacity"`
	// Indicates that the fleet launches all Spot Instances into a single Availability Zone.
	//
	// Supported only for fleets of type `instant` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-spotoptionsrequest.html#cfn-ec2-ec2fleet-spotoptionsrequest-singleavailabilityzone
	//
	SingleAvailabilityZone interface{} `field:"optional" json:"singleAvailabilityZone" yaml:"singleAvailabilityZone"`
	// Indicates that the fleet uses a single instance type to launch all Spot Instances in the fleet.
	//
	// Supported only for fleets of type `instant` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-spotoptionsrequest.html#cfn-ec2-ec2fleet-spotoptionsrequest-singleinstancetype
	//
	SingleInstanceType interface{} `field:"optional" json:"singleInstanceType" yaml:"singleInstanceType"`
}

