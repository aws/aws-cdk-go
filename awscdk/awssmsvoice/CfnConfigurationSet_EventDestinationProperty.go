package awssmsvoice


// Contains information about an event destination.
//
// Event destinations are associated with configuration sets, which enable you to publish message sending events to CloudWatch, Firehose, or Amazon SNS.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventDestinationProperty := &EventDestinationProperty{
//   	Enabled: jsii.Boolean(false),
//   	EventDestinationName: jsii.String("eventDestinationName"),
//   	MatchingEventTypes: []*string{
//   		jsii.String("matchingEventTypes"),
//   	},
//
//   	// the properties below are optional
//   	CloudWatchLogsDestination: &CloudWatchLogsDestinationProperty{
//   		IamRoleArn: jsii.String("iamRoleArn"),
//   		LogGroupArn: jsii.String("logGroupArn"),
//   	},
//   	KinesisFirehoseDestination: &KinesisFirehoseDestinationProperty{
//   		DeliveryStreamArn: jsii.String("deliveryStreamArn"),
//   		IamRoleArn: jsii.String("iamRoleArn"),
//   	},
//   	SnsDestination: &SnsDestinationProperty{
//   		TopicArn: jsii.String("topicArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-smsvoice-configurationset-eventdestination.html
//
type CfnConfigurationSet_EventDestinationProperty struct {
	// When set to true events will be logged.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-smsvoice-configurationset-eventdestination.html#cfn-smsvoice-configurationset-eventdestination-enabled
	//
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The name of the EventDestination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-smsvoice-configurationset-eventdestination.html#cfn-smsvoice-configurationset-eventdestination-eventdestinationname
	//
	EventDestinationName *string `field:"required" json:"eventDestinationName" yaml:"eventDestinationName"`
	// An array of event types that determine which events to log.
	//
	// > The `TEXT_SENT` event type is not supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-smsvoice-configurationset-eventdestination.html#cfn-smsvoice-configurationset-eventdestination-matchingeventtypes
	//
	MatchingEventTypes *[]*string `field:"required" json:"matchingEventTypes" yaml:"matchingEventTypes"`
	// An object that contains information about an event destination that sends logging events to Amazon CloudWatch logs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-smsvoice-configurationset-eventdestination.html#cfn-smsvoice-configurationset-eventdestination-cloudwatchlogsdestination
	//
	CloudWatchLogsDestination interface{} `field:"optional" json:"cloudWatchLogsDestination" yaml:"cloudWatchLogsDestination"`
	// An object that contains information about an event destination for logging to Amazon Data Firehose.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-smsvoice-configurationset-eventdestination.html#cfn-smsvoice-configurationset-eventdestination-kinesisfirehosedestination
	//
	KinesisFirehoseDestination interface{} `field:"optional" json:"kinesisFirehoseDestination" yaml:"kinesisFirehoseDestination"`
	// An object that contains information about an event destination that sends logging events to Amazon SNS.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-smsvoice-configurationset-eventdestination.html#cfn-smsvoice-configurationset-eventdestination-snsdestination
	//
	SnsDestination interface{} `field:"optional" json:"snsDestination" yaml:"snsDestination"`
}

