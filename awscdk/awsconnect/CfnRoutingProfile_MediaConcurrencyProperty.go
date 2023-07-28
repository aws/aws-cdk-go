package awsconnect


// Contains information about which channels are supported, and how many contacts an agent can have on a channel simultaneously.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mediaConcurrencyProperty := &MediaConcurrencyProperty{
//   	Channel: jsii.String("channel"),
//   	Concurrency: jsii.Number(123),
//
//   	// the properties below are optional
//   	CrossChannelBehavior: &CrossChannelBehaviorProperty{
//   		BehaviorType: jsii.String("behaviorType"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-routingprofile-mediaconcurrency.html
//
type CfnRoutingProfile_MediaConcurrencyProperty struct {
	// The channels that agents can handle in the Contact Control Panel (CCP).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-routingprofile-mediaconcurrency.html#cfn-connect-routingprofile-mediaconcurrency-channel
	//
	Channel *string `field:"required" json:"channel" yaml:"channel"`
	// The number of contacts an agent can have on a channel simultaneously.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-routingprofile-mediaconcurrency.html#cfn-connect-routingprofile-mediaconcurrency-concurrency
	//
	Concurrency *float64 `field:"required" json:"concurrency" yaml:"concurrency"`
	// Defines the cross-channel routing behavior that allows an agent working on a contact in one channel to be offered a contact from a different channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-routingprofile-mediaconcurrency.html#cfn-connect-routingprofile-mediaconcurrency-crosschannelbehavior
	//
	CrossChannelBehavior interface{} `field:"optional" json:"crossChannelBehavior" yaml:"crossChannelBehavior"`
}

