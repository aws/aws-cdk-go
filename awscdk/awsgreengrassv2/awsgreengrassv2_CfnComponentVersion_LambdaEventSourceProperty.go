package awsgreengrassv2


// Contains information about an event source for an AWS Lambda function.
//
// The event source defines the topics on which this Lambda function subscribes to receive messages that run the function.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaEventSourceProperty := &lambdaEventSourceProperty{
//   	topic: jsii.String("topic"),
//   	type: jsii.String("type"),
//   }
//
type CfnComponentVersion_LambdaEventSourceProperty struct {
	// The topic to which to subscribe to receive event messages.
	Topic *string `field:"optional" json:"topic" yaml:"topic"`
	// The type of event source. Choose from the following options:.
	//
	// - `PUB_SUB` – Subscribe to local publish/subscribe messages. This event source type doesn't support MQTT wildcards ( `+` and `#` ) in the event source topic.
	// - `IOT_CORE` – Subscribe to AWS IoT Core MQTT messages. This event source type supports MQTT wildcards ( `+` and `#` ) in the event source topic.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

