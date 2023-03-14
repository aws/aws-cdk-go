package awsec2


// Specifies information about an instance type to use in a Capacity Reservation Fleet.
//
// `InstanceTypeSpecification` is a property of the [AWS::EC2::CapacityReservationFleet](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-capacityreservationfleet.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   instanceTypeSpecificationProperty := &instanceTypeSpecificationProperty{
//   	availabilityZone: jsii.String("availabilityZone"),
//   	availabilityZoneId: jsii.String("availabilityZoneId"),
//   	ebsOptimized: jsii.Boolean(false),
//   	instancePlatform: jsii.String("instancePlatform"),
//   	instanceType: jsii.String("instanceType"),
//   	priority: jsii.Number(123),
//   	weight: jsii.Number(123),
//   }
//
type CfnCapacityReservationFleet_InstanceTypeSpecificationProperty struct {
	// The Availability Zone in which the Capacity Reservation Fleet reserves the capacity.
	//
	// A Capacity Reservation Fleet can't span Availability Zones. All instance type specifications that you specify for the Fleet must use the same Availability Zone.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The ID of the Availability Zone in which the Capacity Reservation Fleet reserves the capacity.
	//
	// A Capacity Reservation Fleet can't span Availability Zones. All instance type specifications that you specify for the Fleet must use the same Availability Zone.
	AvailabilityZoneId *string `field:"optional" json:"availabilityZoneId" yaml:"availabilityZoneId"`
	// Indicates whether the Capacity Reservation Fleet supports EBS-optimized instances types.
	//
	// This optimization provides dedicated throughput to Amazon EBS and an optimized configuration stack to provide optimal I/O performance. This optimization isn't available with all instance types. Additional usage charges apply when using EBS-optimized instance types.
	EbsOptimized interface{} `field:"optional" json:"ebsOptimized" yaml:"ebsOptimized"`
	// The type of operating system for which the Capacity Reservation Fleet reserves capacity.
	InstancePlatform *string `field:"optional" json:"instancePlatform" yaml:"instancePlatform"`
	// The instance type for which the Capacity Reservation Fleet reserves capacity.
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// The priority to assign to the instance type.
	//
	// This value is used to determine which of the instance types specified for the Fleet should be prioritized for use. A lower value indicates a high priority. For more information, see [Instance type priority](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/crfleet-concepts.html#instance-priority) in the Amazon EC2 User Guide.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// The number of capacity units provided by the specified instance type.
	//
	// This value, together with the total target capacity that you specify for the Fleet determine the number of instances for which the Fleet reserves capacity. Both values are based on units that make sense for your workload. For more information, see [Total target capacity](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/crfleet-concepts.html#target-capacity) in the Amazon EC2 User Guide.
	//
	// Valid Range: Minimum value of `0.001` . Maximum value of `99.999` .
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

