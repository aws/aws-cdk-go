package awsathena


// A mapping between one or more workgroups and a capacity reservation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   capacityAssignmentProperty := &CapacityAssignmentProperty{
//   	WorkgroupNames: []*string{
//   		jsii.String("workgroupNames"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-capacityreservation-capacityassignment.html
//
type CfnCapacityReservation_CapacityAssignmentProperty struct {
	// The list of workgroup names for the capacity assignment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-capacityreservation-capacityassignment.html#cfn-athena-capacityreservation-capacityassignment-workgroupnames
	//
	WorkgroupNames *[]*string `field:"required" json:"workgroupNames" yaml:"workgroupNames"`
}

