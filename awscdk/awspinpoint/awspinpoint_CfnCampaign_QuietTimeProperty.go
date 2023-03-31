package awspinpoint


// Specifies the start and end times that define a time range when messages aren't sent to endpoints.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   quietTimeProperty := &quietTimeProperty{
//   	end: jsii.String("end"),
//   	start: jsii.String("start"),
//   }
//
type CfnCampaign_QuietTimeProperty struct {
	// The specific time when quiet time ends.
	//
	// This value has to use 24-hour notation and be in HH:MM format, where HH is the hour (with a leading zero, if applicable) and MM is the minutes. For example, use `02:30` to represent 2:30 AM, or `14:30` to represent 2:30 PM.
	End *string `field:"required" json:"end" yaml:"end"`
	// The specific time when quiet time begins.
	//
	// This value has to use 24-hour notation and be in HH:MM format, where HH is the hour (with a leading zero, if applicable) and MM is the minutes. For example, use `02:30` to represent 2:30 AM, or `14:30` to represent 2:30 PM.
	Start *string `field:"required" json:"start" yaml:"start"`
}

