package awsses


// An object that contains Event bus ARN associated with the event bridge destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventBridgeDestinationProperty := &EventBridgeDestinationProperty{
//   	EventBusArn: jsii.String("eventBusArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationseteventdestination-eventbridgedestination.html
//
type CfnConfigurationSetEventDestination_EventBridgeDestinationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationseteventdestination-eventbridgedestination.html#cfn-ses-configurationseteventdestination-eventbridgedestination-eventbusarn
	//
	EventBusArn *string `field:"required" json:"eventBusArn" yaml:"eventBusArn"`
}

