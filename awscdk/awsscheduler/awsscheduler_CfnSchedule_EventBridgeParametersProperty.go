package awsscheduler


// The templated target type for the EventBridge [`PutEvents`](https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_PutEvents.html) API operation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventBridgeParametersProperty := &eventBridgeParametersProperty{
//   	detailType: jsii.String("detailType"),
//   	source: jsii.String("source"),
//   }
//
type CfnSchedule_EventBridgeParametersProperty struct {
	// A free-form string, with a maximum of 128 characters, used to decide what fields to expect in the event detail.
	DetailType *string `field:"required" json:"detailType" yaml:"detailType"`
	// The source of the event.
	Source *string `field:"required" json:"source" yaml:"source"`
}

