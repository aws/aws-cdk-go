package awsappsync


// The data source.
//
// This can be an API destination, resource, or AWS service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventBridgeConfigProperty := &eventBridgeConfigProperty{
//   	eventBusArn: jsii.String("eventBusArn"),
//   }
//
type CfnDataSource_EventBridgeConfigProperty struct {
	// The event bus pipeline's ARN.
	//
	// For more information about event buses, see [EventBridge event buses](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-event-bus.html) .
	EventBusArn *string `field:"required" json:"eventBusArn" yaml:"eventBusArn"`
}

