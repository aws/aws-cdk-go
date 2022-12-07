package awsscheduler


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
	// `CfnSchedule.EventBridgeParametersProperty.DetailType`.
	DetailType *string `field:"required" json:"detailType" yaml:"detailType"`
	// `CfnSchedule.EventBridgeParametersProperty.Source`.
	Source *string `field:"required" json:"source" yaml:"source"`
}

