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
type CfnCapacityReservation_CapacityAssignmentProperty struct {
	// The list of workgroup names for the capacity assignment.
	WorkgroupNames *[]*string `field:"required" json:"workgroupNames" yaml:"workgroupNames"`
}

