package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scheduleEventProperty := &ScheduleEventProperty{
//   	Schedule: jsii.String("schedule"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Enabled: jsii.Boolean(false),
//   	Input: jsii.String("input"),
//   	Name: jsii.String("name"),
//   }
//
type CfnFunction_ScheduleEventProperty struct {
	// `CfnFunction.ScheduleEventProperty.Schedule`.
	Schedule *string `field:"required" json:"schedule" yaml:"schedule"`
	// `CfnFunction.ScheduleEventProperty.Description`.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `CfnFunction.ScheduleEventProperty.Enabled`.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// `CfnFunction.ScheduleEventProperty.Input`.
	Input *string `field:"optional" json:"input" yaml:"input"`
	// `CfnFunction.ScheduleEventProperty.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

