package awsevents


// The logging configuration settings for the event bus.
//
// For more information, see [Configuring logs for event buses](https://docs.aws.amazon.com/eb-event-bus-logs.html) in the *EventBridge User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   logConfigProperty := &LogConfigProperty{
//   	IncludeDetail: jsii.String("includeDetail"),
//   	Level: jsii.String("level"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-eventbus-logconfig.html
//
type CfnEventBus_LogConfigProperty struct {
	// Whether EventBridge include detailed event information in the records it generates.
	//
	// Detailed data can be useful for troubleshooting and debugging. This information includes details of the event itself, as well as target details.
	//
	// For more information, see [Including detail data in event bus logs](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-event-bus-logs.html#eb-event-logs-data) in the *EventBridge User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-eventbus-logconfig.html#cfn-events-eventbus-logconfig-includedetail
	//
	IncludeDetail *string `field:"optional" json:"includeDetail" yaml:"includeDetail"`
	// The level of logging detail to include. This applies to all log destinations for the event bus.
	//
	// For more information, see [Specifying event bus log level](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-event-bus-logs.html#eb-event-bus-logs-level) in the *EventBridge User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-eventbus-logconfig.html#cfn-events-eventbus-logconfig-level
	//
	Level *string `field:"optional" json:"level" yaml:"level"`
}

