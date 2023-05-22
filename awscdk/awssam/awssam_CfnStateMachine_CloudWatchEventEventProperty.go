package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var pattern interface{}
//
//   cloudWatchEventEventProperty := &CloudWatchEventEventProperty{
//   	Pattern: pattern,
//
//   	// the properties below are optional
//   	EventBusName: jsii.String("eventBusName"),
//   	Input: jsii.String("input"),
//   	InputPath: jsii.String("inputPath"),
//   }
//
type CfnStateMachine_CloudWatchEventEventProperty struct {
	// `CfnStateMachine.CloudWatchEventEventProperty.Pattern`.
	Pattern interface{} `field:"required" json:"pattern" yaml:"pattern"`
	// `CfnStateMachine.CloudWatchEventEventProperty.EventBusName`.
	EventBusName *string `field:"optional" json:"eventBusName" yaml:"eventBusName"`
	// `CfnStateMachine.CloudWatchEventEventProperty.Input`.
	Input *string `field:"optional" json:"input" yaml:"input"`
	// `CfnStateMachine.CloudWatchEventEventProperty.InputPath`.
	InputPath *string `field:"optional" json:"inputPath" yaml:"inputPath"`
}

