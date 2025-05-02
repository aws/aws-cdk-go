package awsec2


// Information about instance capacity usage for a Capacity Reservation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   capacityAllocationProperty := &CapacityAllocationProperty{
//   	AllocationType: jsii.String("allocationType"),
//   	Count: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-capacityreservation-capacityallocation.html
//
type CfnCapacityReservation_CapacityAllocationProperty struct {
	// The usage type.
	//
	// `used` indicates that the instance capacity is in use by instances that are running in the Capacity Reservation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-capacityreservation-capacityallocation.html#cfn-ec2-capacityreservation-capacityallocation-allocationtype
	//
	AllocationType *string `field:"optional" json:"allocationType" yaml:"allocationType"`
	// The amount of instance capacity associated with the usage.
	//
	// For example a value of `4` indicates that instance capacity for 4 instances is currently in use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-capacityreservation-capacityallocation.html#cfn-ec2-capacityreservation-capacityallocation-count
	//
	Count *float64 `field:"optional" json:"count" yaml:"count"`
}

