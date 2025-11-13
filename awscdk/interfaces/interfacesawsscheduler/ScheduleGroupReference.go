package interfacesawsscheduler


// A reference to a ScheduleGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scheduleGroupReference := &ScheduleGroupReference{
//   	ScheduleGroupArn: jsii.String("scheduleGroupArn"),
//   	ScheduleGroupName: jsii.String("scheduleGroupName"),
//   }
//
type ScheduleGroupReference struct {
	// The ARN of the ScheduleGroup resource.
	ScheduleGroupArn *string `field:"required" json:"scheduleGroupArn" yaml:"scheduleGroupArn"`
	// The Name of the ScheduleGroup resource.
	ScheduleGroupName *string `field:"required" json:"scheduleGroupName" yaml:"scheduleGroupName"`
}

