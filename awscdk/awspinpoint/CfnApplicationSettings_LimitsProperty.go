package awspinpoint


// Specifies the default sending limits for campaigns in the application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   limitsProperty := &LimitsProperty{
//   	Daily: jsii.Number(123),
//   	MaximumDuration: jsii.Number(123),
//   	MessagesPerSecond: jsii.Number(123),
//   	Total: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-applicationsettings-limits.html
//
type CfnApplicationSettings_LimitsProperty struct {
	// The maximum number of messages that a campaign can send to a single endpoint during a 24-hour period.
	//
	// The maximum value is 100.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-applicationsettings-limits.html#cfn-pinpoint-applicationsettings-limits-daily
	//
	Daily *float64 `field:"optional" json:"daily" yaml:"daily"`
	// The maximum amount of time, in seconds, that a campaign can attempt to deliver a message after the scheduled start time for the campaign.
	//
	// The minimum value is 60 seconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-applicationsettings-limits.html#cfn-pinpoint-applicationsettings-limits-maximumduration
	//
	MaximumDuration *float64 `field:"optional" json:"maximumDuration" yaml:"maximumDuration"`
	// The maximum number of messages that a campaign can send each second.
	//
	// The minimum value is 1. The maximum value is 20,000.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-applicationsettings-limits.html#cfn-pinpoint-applicationsettings-limits-messagespersecond
	//
	MessagesPerSecond *float64 `field:"optional" json:"messagesPerSecond" yaml:"messagesPerSecond"`
	// The maximum number of messages that a campaign can send to a single endpoint during the course of the campaign.
	//
	// The maximum value is 100.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-applicationsettings-limits.html#cfn-pinpoint-applicationsettings-limits-total
	//
	Total *float64 `field:"optional" json:"total" yaml:"total"`
}

