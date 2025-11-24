package mixinsawsses


// An object that defines an Amazon EventBridge destination for email events.
//
// You can use Amazon EventBridge to send notifications when certain email events occur.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eventBridgeDestinationProperty := &EventBridgeDestinationProperty{
//   	EventBusArn: jsii.String("eventBusArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationseteventdestination-eventbridgedestination.html
//
type CfnConfigurationSetEventDestinationPropsMixin_EventBridgeDestinationProperty struct {
	// The Amazon Resource Name (ARN) of the Amazon EventBridge bus to publish email events to.
	//
	// Only the default bus is supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationseteventdestination-eventbridgedestination.html#cfn-ses-configurationseteventdestination-eventbridgedestination-eventbusarn
	//
	EventBusArn *string `field:"optional" json:"eventBusArn" yaml:"eventBusArn"`
}

