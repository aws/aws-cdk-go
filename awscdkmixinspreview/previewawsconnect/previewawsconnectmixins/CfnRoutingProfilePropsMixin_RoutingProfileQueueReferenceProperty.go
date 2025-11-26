package previewawsconnectmixins


// Contains the channel and queue identifier for a routing profile.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   routingProfileQueueReferenceProperty := &RoutingProfileQueueReferenceProperty{
//   	Channel: jsii.String("channel"),
//   	QueueArn: jsii.String("queueArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-routingprofile-routingprofilequeuereference.html
//
type CfnRoutingProfilePropsMixin_RoutingProfileQueueReferenceProperty struct {
	// The channels agents can handle in the Contact Control Panel (CCP) for this routing profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-routingprofile-routingprofilequeuereference.html#cfn-connect-routingprofile-routingprofilequeuereference-channel
	//
	Channel *string `field:"optional" json:"channel" yaml:"channel"`
	// The Amazon Resource Name (ARN) of the queue.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-routingprofile-routingprofilequeuereference.html#cfn-connect-routingprofile-routingprofilequeuereference-queuearn
	//
	QueueArn *string `field:"optional" json:"queueArn" yaml:"queueArn"`
}

