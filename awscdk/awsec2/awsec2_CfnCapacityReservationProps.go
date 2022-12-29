package awsec2


// Properties for defining a `CfnCapacityReservation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCapacityReservationProps := &cfnCapacityReservationProps{
//   	availabilityZone: jsii.String("availabilityZone"),
//   	instanceCount: jsii.Number(123),
//   	instancePlatform: jsii.String("instancePlatform"),
//   	instanceType: jsii.String("instanceType"),
//
//   	// the properties below are optional
//   	ebsOptimized: jsii.Boolean(false),
//   	endDate: jsii.String("endDate"),
//   	endDateType: jsii.String("endDateType"),
//   	ephemeralStorage: jsii.Boolean(false),
//   	instanceMatchCriteria: jsii.String("instanceMatchCriteria"),
//   	outPostArn: jsii.String("outPostArn"),
//   	placementGroupArn: jsii.String("placementGroupArn"),
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
//   }
//
type CfnCapacityReservationProps struct {
	// The Availability Zone in which to create the Capacity Reservation.
	AvailabilityZone *string `field:"required" json:"availabilityZone" yaml:"availabilityZone"`
	// The number of instances for which to reserve capacity.
	//
	// Valid range: 1 - 1000.
	InstanceCount *float64 `field:"required" json:"instanceCount" yaml:"instanceCount"`
	// The type of operating system for which to reserve capacity.
	InstancePlatform *string `field:"required" json:"instancePlatform" yaml:"instancePlatform"`
	// The instance type for which to reserve capacity.
	//
	// For more information, see [Instance types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html) in the *Amazon EC2 User Guide* .
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// Indicates whether the Capacity Reservation supports EBS-optimized instances.
	//
	// This optimization provides dedicated throughput to Amazon EBS and an optimized configuration stack to provide optimal I/O performance. This optimization isn't available with all instance types. Additional usage charges apply when using an EBS- optimized instance.
	EbsOptimized interface{} `field:"optional" json:"ebsOptimized" yaml:"ebsOptimized"`
	// The date and time at which the Capacity Reservation expires.
	//
	// When a Capacity Reservation expires, the reserved capacity is released and you can no longer launch instances into it. The Capacity Reservation's state changes to `expired` when it reaches its end date and time.
	//
	// You must provide an `EndDate` value if `EndDateType` is `limited` . Omit `EndDate` if `EndDateType` is `unlimited` .
	//
	// If the `EndDateType` is `limited` , the Capacity Reservation is cancelled within an hour from the specified time. For example, if you specify 5/31/2019, 13:30:55, the Capacity Reservation is guaranteed to end between 13:30:55 and 14:30:55 on 5/31/2019.
	EndDate *string `field:"optional" json:"endDate" yaml:"endDate"`
	// Indicates the way in which the Capacity Reservation ends.
	//
	// A Capacity Reservation can have one of the following end types:
	//
	// - `unlimited` - The Capacity Reservation remains active until you explicitly cancel it. Do not provide an `EndDate` if the `EndDateType` is `unlimited` .
	// - `limited` - The Capacity Reservation expires automatically at a specified date and time. You must provide an `EndDate` value if the `EndDateType` value is `limited` .
	EndDateType *string `field:"optional" json:"endDateType" yaml:"endDateType"`
	// Indicates whether the Capacity Reservation supports instances with temporary, block-level storage.
	EphemeralStorage interface{} `field:"optional" json:"ephemeralStorage" yaml:"ephemeralStorage"`
	// Indicates the type of instance launches that the Capacity Reservation accepts. The options include:.
	//
	// - `open` - The Capacity Reservation automatically matches all instances that have matching attributes (instance type, platform, and Availability Zone). Instances that have matching attributes run in the Capacity Reservation automatically without specifying any additional parameters.
	// - `targeted` - The Capacity Reservation only accepts instances that have matching attributes (instance type, platform, and Availability Zone), and explicitly target the Capacity Reservation. This ensures that only permitted instances can use the reserved capacity.
	//
	// Default: `open`.
	InstanceMatchCriteria *string `field:"optional" json:"instanceMatchCriteria" yaml:"instanceMatchCriteria"`
	// The Amazon Resource Name (ARN) of the Outpost on which to create the Capacity Reservation.
	OutPostArn *string `field:"optional" json:"outPostArn" yaml:"outPostArn"`
	// The Amazon Resource Name (ARN) of the cluster placement group in which to create the Capacity Reservation.
	//
	// For more information, see [Capacity Reservations for cluster placement groups](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/cr-cpg.html) in the *Amazon EC2 User Guide* .
	PlacementGroupArn *string `field:"optional" json:"placementGroupArn" yaml:"placementGroupArn"`
	// The tags to apply to the Capacity Reservation during launch.
	TagSpecifications interface{} `field:"optional" json:"tagSpecifications" yaml:"tagSpecifications"`
	// Indicates the tenancy of the Capacity Reservation. A Capacity Reservation can have one of the following tenancy settings:.
	//
	// - `default` - The Capacity Reservation is created on hardware that is shared with other AWS accounts .
	// - `dedicated` - The Capacity Reservation is created on single-tenant hardware that is dedicated to a single AWS account .
	Tenancy *string `field:"optional" json:"tenancy" yaml:"tenancy"`
}

