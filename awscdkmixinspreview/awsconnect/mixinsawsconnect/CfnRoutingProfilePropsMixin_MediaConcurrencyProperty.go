package mixinsawsconnect


// Contains information about which channels are supported, and how many contacts an agent can have on a channel simultaneously.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   mediaConcurrencyProperty := &MediaConcurrencyProperty{
//   	Channel: jsii.String("channel"),
//   	Concurrency: jsii.Number(123),
//   	CrossChannelBehavior: &CrossChannelBehaviorProperty{
//   		BehaviorType: jsii.String("behaviorType"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-routingprofile-mediaconcurrency.html
//
type CfnRoutingProfilePropsMixin_MediaConcurrencyProperty struct {
	// The channels that agents can handle in the Contact Control Panel (CCP).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-routingprofile-mediaconcurrency.html#cfn-connect-routingprofile-mediaconcurrency-channel
	//
	Channel *string `field:"optional" json:"channel" yaml:"channel"`
	// The number of contacts an agent can have on a channel simultaneously.
	//
	// Valid Range for `VOICE` : Minimum value of 1. Maximum value of 1.
	//
	// Valid Range for `CHAT` : Minimum value of 1. Maximum value of 10.
	//
	// Valid Range for `TASK` : Minimum value of 1. Maximum value of 10.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-routingprofile-mediaconcurrency.html#cfn-connect-routingprofile-mediaconcurrency-concurrency
	//
	Concurrency *float64 `field:"optional" json:"concurrency" yaml:"concurrency"`
	// Defines the cross-channel routing behavior for each channel that is enabled for this Routing Profile.
	//
	// For example, this allows you to offer an agent a different contact from another channel when they are currently working with a contact from a Voice channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-routingprofile-mediaconcurrency.html#cfn-connect-routingprofile-mediaconcurrency-crosschannelbehavior
	//
	CrossChannelBehavior interface{} `field:"optional" json:"crossChannelBehavior" yaml:"crossChannelBehavior"`
}

