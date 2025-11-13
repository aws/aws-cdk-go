package interfacesawsdatabrew


// A reference to a Schedule resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scheduleReference := &ScheduleReference{
//   	ScheduleName: jsii.String("scheduleName"),
//   }
//
type ScheduleReference struct {
	// The Name of the Schedule resource.
	ScheduleName *string `field:"required" json:"scheduleName" yaml:"scheduleName"`
}

