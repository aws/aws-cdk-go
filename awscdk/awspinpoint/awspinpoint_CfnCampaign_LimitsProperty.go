package awspinpoint


// Specifies the limits on the messages that a campaign can send.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   limitsProperty := &limitsProperty{
//   	daily: jsii.Number(123),
//   	maximumDuration: jsii.Number(123),
//   	messagesPerSecond: jsii.Number(123),
//   	session: jsii.Number(123),
//   	total: jsii.Number(123),
//   }
//
type CfnCampaign_LimitsProperty struct {
	// The maximum number of messages that a campaign can send to a single endpoint during a 24-hour period.
	//
	// The maximum value is 100.
	Daily *float64 `field:"optional" json:"daily" yaml:"daily"`
	// The maximum amount of time, in seconds, that a campaign can attempt to deliver a message after the scheduled start time for the campaign.
	//
	// The minimum value is 60 seconds.
	MaximumDuration *float64 `field:"optional" json:"maximumDuration" yaml:"maximumDuration"`
	// The maximum number of messages that a campaign can send each second.
	//
	// The minimum value is 50. The maximum value is 20,000.
	MessagesPerSecond *float64 `field:"optional" json:"messagesPerSecond" yaml:"messagesPerSecond"`
	// `CfnCampaign.LimitsProperty.Session`.
	Session *float64 `field:"optional" json:"session" yaml:"session"`
	// The maximum number of messages that a campaign can send to a single endpoint during the course of the campaign.
	//
	// The maximum value is 100.
	Total *float64 `field:"optional" json:"total" yaml:"total"`
}

