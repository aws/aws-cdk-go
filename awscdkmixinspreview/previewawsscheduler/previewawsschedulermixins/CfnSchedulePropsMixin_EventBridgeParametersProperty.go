package previewawsschedulermixins


// The templated target type for the EventBridge [`PutEvents`](https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_PutEvents.html) API operation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eventBridgeParametersProperty := &EventBridgeParametersProperty{
//   	DetailType: jsii.String("detailType"),
//   	Source: jsii.String("source"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scheduler-schedule-eventbridgeparameters.html
//
type CfnSchedulePropsMixin_EventBridgeParametersProperty struct {
	// A free-form string, with a maximum of 128 characters, used to decide what fields to expect in the event detail.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scheduler-schedule-eventbridgeparameters.html#cfn-scheduler-schedule-eventbridgeparameters-detailtype
	//
	DetailType *string `field:"optional" json:"detailType" yaml:"detailType"`
	// The source of the event.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scheduler-schedule-eventbridgeparameters.html#cfn-scheduler-schedule-eventbridgeparameters-source
	//
	Source *string `field:"optional" json:"source" yaml:"source"`
}

