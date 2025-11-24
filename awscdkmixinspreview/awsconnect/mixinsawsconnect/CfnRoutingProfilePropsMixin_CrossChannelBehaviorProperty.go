package mixinsawsconnect


// Defines the cross-channel routing behavior that allows an agent working on a contact in one channel to be offered a contact from a different channel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   crossChannelBehaviorProperty := &CrossChannelBehaviorProperty{
//   	BehaviorType: jsii.String("behaviorType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-routingprofile-crosschannelbehavior.html
//
type CfnRoutingProfilePropsMixin_CrossChannelBehaviorProperty struct {
	// Specifies the other channels that can be routed to an agent handling their current channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-routingprofile-crosschannelbehavior.html#cfn-connect-routingprofile-crosschannelbehavior-behaviortype
	//
	BehaviorType *string `field:"optional" json:"behaviorType" yaml:"behaviorType"`
}

