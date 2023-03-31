package awsec2


// Properties for defining a `CfnCapacityReservationFleet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCapacityReservationFleetProps := &cfnCapacityReservationFleetProps{
//   	allocationStrategy: jsii.String("allocationStrategy"),
//   	endDate: jsii.String("endDate"),
//   	instanceMatchCriteria: jsii.String("instanceMatchCriteria"),
//   	instanceTypeSpecifications: []interface{}{
//   		&instanceTypeSpecificationProperty{
//   			availabilityZone: jsii.String("availabilityZone"),
//   			availabilityZoneId: jsii.String("availabilityZoneId"),
//   			ebsOptimized: jsii.Boolean(false),
//   			instancePlatform: jsii.String("instancePlatform"),
//   			instanceType: jsii.String("instanceType"),
//   			priority: jsii.Number(123),
//   			weight: jsii.Number(123),
//   		},
//   	},
//   	noRemoveEndDate: jsii.Boolean(false),
//   	removeEndDate: jsii.Boolean(false),
//   	tagSpecifications: []interface{}{
//   		&tagSpecificationProperty{
//   			resourceType: jsii.String("resourceType"),
//   			tags: []cfnTag{
//   				&cfnTag{
//   					key: jsii.String("key"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	tenancy: jsii.String("tenancy"),
//   	totalTargetCapacity: jsii.Number(123),
//   }
//
type CfnCapacityReservationFleetProps struct {
	// The strategy used by the Capacity Reservation Fleet to determine which of the specified instance types to use.
	//
	// Currently, only the `prioritized` allocation strategy is supported. For more information, see [Allocation strategy](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/crfleet-concepts.html#allocation-strategy) in the Amazon EC2 User Guide.
	//
	// Valid values: `prioritized`.
	AllocationStrategy *string `field:"optional" json:"allocationStrategy" yaml:"allocationStrategy"`
	// The date and time at which the Capacity Reservation Fleet expires.
	//
	// When the Capacity Reservation Fleet expires, its state changes to `expired` and all of the Capacity Reservations in the Fleet expire.
	//
	// The Capacity Reservation Fleet expires within an hour after the specified time. For example, if you specify `5/31/2019` , `13:30:55` , the Capacity Reservation Fleet is guaranteed to expire between `13:30:55` and `14:30:55` on `5/31/2019` .
	EndDate *string `field:"optional" json:"endDate" yaml:"endDate"`
	// Indicates the type of instance launches that the Capacity Reservation Fleet accepts.
	//
	// All Capacity Reservations in the Fleet inherit this instance matching criteria.
	//
	// Currently, Capacity Reservation Fleets support `open` instance matching criteria only. This means that instances that have matching attributes (instance type, platform, and Availability Zone) run in the Capacity Reservations automatically. Instances do not need to explicitly target a Capacity Reservation Fleet to use its reserved capacity.
	InstanceMatchCriteria *string `field:"optional" json:"instanceMatchCriteria" yaml:"instanceMatchCriteria"`
	// Information about the instance types for which to reserve the capacity.
	InstanceTypeSpecifications interface{} `field:"optional" json:"instanceTypeSpecifications" yaml:"instanceTypeSpecifications"`
	// `AWS::EC2::CapacityReservationFleet.NoRemoveEndDate`.
	NoRemoveEndDate interface{} `field:"optional" json:"noRemoveEndDate" yaml:"noRemoveEndDate"`
	// `AWS::EC2::CapacityReservationFleet.RemoveEndDate`.
	RemoveEndDate interface{} `field:"optional" json:"removeEndDate" yaml:"removeEndDate"`
	// The tags to assign to the Capacity Reservation Fleet.
	//
	// The tags are automatically assigned to the Capacity Reservations in the Fleet.
	TagSpecifications interface{} `field:"optional" json:"tagSpecifications" yaml:"tagSpecifications"`
	// Indicates the tenancy of the Capacity Reservation Fleet.
	//
	// All Capacity Reservations in the Fleet inherit this tenancy. The Capacity Reservation Fleet can have one of the following tenancy settings:
	//
	// - `default` - The Capacity Reservation Fleet is created on hardware that is shared with other AWS accounts .
	// - `dedicated` - The Capacity Reservations are created on single-tenant hardware that is dedicated to a single AWS account .
	Tenancy *string `field:"optional" json:"tenancy" yaml:"tenancy"`
	// The total number of capacity units to be reserved by the Capacity Reservation Fleet.
	//
	// This value, together with the instance type weights that you assign to each instance type used by the Fleet determine the number of instances for which the Fleet reserves capacity. Both values are based on units that make sense for your workload. For more information, see [Total target capacity](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/crfleet-concepts.html#target-capacity) in the Amazon EC2 User Guide.
	TotalTargetCapacity *float64 `field:"optional" json:"totalTargetCapacity" yaml:"totalTargetCapacity"`
}

