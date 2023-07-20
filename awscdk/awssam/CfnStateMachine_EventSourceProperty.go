package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventSourceProperty := &EventSourceProperty{
//   	Properties: &ApiEventProperty{
//   		Method: jsii.String("method"),
//   		Path: jsii.String("path"),
//
//   		// the properties below are optional
//   		RestApiId: jsii.String("restApiId"),
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-statemachine-eventsource.html
//
type CfnStateMachine_EventSourceProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-statemachine-eventsource.html#cfn-serverless-statemachine-eventsource-properties
	//
	Properties interface{} `field:"required" json:"properties" yaml:"properties"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-statemachine-eventsource.html#cfn-serverless-statemachine-eventsource-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
}

