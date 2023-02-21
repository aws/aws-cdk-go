package awspinpointemail


// In Amazon Pinpoint, *events* include message sends, deliveries, opens, clicks, bounces, and complaints.
//
// *Event destinations* are places that you can send information about these events to. For example, you can send event data to Amazon SNS to receive notifications when you receive bounces or complaints, or you can use Amazon Kinesis Data Firehose to stream data to Amazon S3 for long-term storage.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventDestinationProperty := &eventDestinationProperty{
//   	matchingEventTypes: []*string{
//   		jsii.String("matchingEventTypes"),
//   	},
//
//   	// the properties below are optional
//   	cloudWatchDestination: &cloudWatchDestinationProperty{
//   		dimensionConfigurations: []interface{}{
//   			&dimensionConfigurationProperty{
//   				defaultDimensionValue: jsii.String("defaultDimensionValue"),
//   				dimensionName: jsii.String("dimensionName"),
//   				dimensionValueSource: jsii.String("dimensionValueSource"),
//   			},
//   		},
//   	},
//   	enabled: jsii.Boolean(false),
//   	kinesisFirehoseDestination: &kinesisFirehoseDestinationProperty{
//   		deliveryStreamArn: jsii.String("deliveryStreamArn"),
//   		iamRoleArn: jsii.String("iamRoleArn"),
//   	},
//   	pinpointDestination: &pinpointDestinationProperty{
//   		applicationArn: jsii.String("applicationArn"),
//   	},
//   	snsDestination: &snsDestinationProperty{
//   		topicArn: jsii.String("topicArn"),
//   	},
//   }
//
type CfnConfigurationSetEventDestination_EventDestinationProperty struct {
	// The types of events that Amazon Pinpoint sends to the specified event destinations.
	//
	// Acceptable values: `SEND` , `REJECT` , `BOUNCE` , `COMPLAINT` , `DELIVERY` , `OPEN` , `CLICK` , and `RENDERING_FAILURE` .
	MatchingEventTypes *[]*string `field:"required" json:"matchingEventTypes" yaml:"matchingEventTypes"`
	// An object that defines an Amazon CloudWatch destination for email events.
	//
	// You can use Amazon CloudWatch to monitor and gain insights on your email sending metrics.
	CloudWatchDestination interface{} `field:"optional" json:"cloudWatchDestination" yaml:"cloudWatchDestination"`
	// If `true` , the event destination is enabled.
	//
	// When the event destination is enabled, the specified event types are sent to the destinations in this `EventDestinationDefinition` .
	//
	// If `false` , the event destination is disabled. When the event destination is disabled, events aren't sent to the specified destinations.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// An object that defines an Amazon Kinesis Data Firehose destination for email events.
	//
	// You can use Amazon Kinesis Data Firehose to stream data to other services, such as Amazon S3 and Amazon Redshift.
	KinesisFirehoseDestination interface{} `field:"optional" json:"kinesisFirehoseDestination" yaml:"kinesisFirehoseDestination"`
	// An object that defines a Amazon Pinpoint destination for email events.
	//
	// You can use Amazon Pinpoint events to create attributes in Amazon Pinpoint projects. You can use these attributes to create segments for your campaigns.
	PinpointDestination interface{} `field:"optional" json:"pinpointDestination" yaml:"pinpointDestination"`
	// An object that defines an Amazon SNS destination for email events.
	//
	// You can use Amazon SNS to send notification when certain email events occur.
	SnsDestination interface{} `field:"optional" json:"snsDestination" yaml:"snsDestination"`
}

