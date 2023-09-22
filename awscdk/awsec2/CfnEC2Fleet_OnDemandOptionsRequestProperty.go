package awsec2


// Specifies the allocation strategy of On-Demand Instances in an EC2 Fleet.
//
// `OnDemandOptionsRequest` is a property of the [AWS::EC2::EC2Fleet](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   onDemandOptionsRequestProperty := &OnDemandOptionsRequestProperty{
//   	AllocationStrategy: jsii.String("allocationStrategy"),
//   	CapacityReservationOptions: &CapacityReservationOptionsRequestProperty{
//   		UsageStrategy: jsii.String("usageStrategy"),
//   	},
//   	MaxTotalPrice: jsii.String("maxTotalPrice"),
//   	MinTargetCapacity: jsii.Number(123),
//   	SingleAvailabilityZone: jsii.Boolean(false),
//   	SingleInstanceType: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-ondemandoptionsrequest.html
//
type CfnEC2Fleet_OnDemandOptionsRequestProperty struct {
	// The strategy that determines the order of the launch template overrides to use in fulfilling On-Demand capacity.
	//
	// `lowest-price` - EC2 Fleet uses price to determine the order, launching the lowest price first.
	//
	// `prioritized` - EC2 Fleet uses the priority that you assigned to each launch template override, launching the highest priority first.
	//
	// Default: `lowest-price`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-ondemandoptionsrequest.html#cfn-ec2-ec2fleet-ondemandoptionsrequest-allocationstrategy
	//
	AllocationStrategy *string `field:"optional" json:"allocationStrategy" yaml:"allocationStrategy"`
	// The strategy for using unused Capacity Reservations for fulfilling On-Demand capacity.
	//
	// Supported only for fleets of type `instant` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-ondemandoptionsrequest.html#cfn-ec2-ec2fleet-ondemandoptionsrequest-capacityreservationoptions
	//
	CapacityReservationOptions interface{} `field:"optional" json:"capacityReservationOptions" yaml:"capacityReservationOptions"`
	// The maximum amount per hour for On-Demand Instances that you're willing to pay.
	//
	// > If your fleet includes T instances that are configured as `unlimited` , and if their average CPU usage exceeds the baseline utilization, you will incur a charge for surplus credits. The `MaxTotalPrice` does not account for surplus credits, and, if you use surplus credits, your final cost might be higher than what you specified for `MaxTotalPrice` . For more information, see [Surplus credits can incur charges](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances-unlimited-mode-concepts.html#unlimited-mode-surplus-credits) in the *EC2 User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-ondemandoptionsrequest.html#cfn-ec2-ec2fleet-ondemandoptionsrequest-maxtotalprice
	//
	MaxTotalPrice *string `field:"optional" json:"maxTotalPrice" yaml:"maxTotalPrice"`
	// The minimum target capacity for On-Demand Instances in the fleet.
	//
	// If the minimum target capacity is not reached, the fleet launches no instances.
	//
	// Supported only for fleets of type `instant` .
	//
	// At least one of the following must be specified: `SingleAvailabilityZone` | `SingleInstanceType`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-ondemandoptionsrequest.html#cfn-ec2-ec2fleet-ondemandoptionsrequest-mintargetcapacity
	//
	MinTargetCapacity *float64 `field:"optional" json:"minTargetCapacity" yaml:"minTargetCapacity"`
	// Indicates that the fleet launches all On-Demand Instances into a single Availability Zone.
	//
	// Supported only for fleets of type `instant` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-ondemandoptionsrequest.html#cfn-ec2-ec2fleet-ondemandoptionsrequest-singleavailabilityzone
	//
	SingleAvailabilityZone interface{} `field:"optional" json:"singleAvailabilityZone" yaml:"singleAvailabilityZone"`
	// Indicates that the fleet uses a single instance type to launch all On-Demand Instances in the fleet.
	//
	// Supported only for fleets of type `instant` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-ondemandoptionsrequest.html#cfn-ec2-ec2fleet-ondemandoptionsrequest-singleinstancetype
	//
	SingleInstanceType interface{} `field:"optional" json:"singleInstanceType" yaml:"singleInstanceType"`
}

