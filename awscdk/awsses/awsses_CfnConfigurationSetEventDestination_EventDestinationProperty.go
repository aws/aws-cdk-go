package awsses


// Contains information about an event destination.
//
// > When you create or update an event destination, you must provide one, and only one, destination. The destination can be Amazon CloudWatch, Amazon Kinesis Firehose or Amazon Simple Notification Service (Amazon SNS).
//
// Event destinations are associated with configuration sets, which enable you to publish email sending events to Amazon CloudWatch, Amazon Kinesis Firehose, or Amazon Simple Notification Service (Amazon SNS). For information about using configuration sets, see the [Amazon SES Developer Guide](https://docs.aws.amazon.com/ses/latest/dg/monitor-sending-activity.html) .
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
//   	name: jsii.String("name"),
//   	snsDestination: &snsDestinationProperty{
//   		topicArn: jsii.String("topicArn"),
//   	},
//   }
//
type CfnConfigurationSetEventDestination_EventDestinationProperty struct {
	// The type of email sending events to publish to the event destination.
	//
	// - `send` - The call was successful and Amazon SES is attempting to deliver the email.
	// - `reject` - Amazon SES determined that the email contained a virus and rejected it.
	// - `bounce` - The recipient's mail server permanently rejected the email. This corresponds to a hard bounce.
	// - `complaint` - The recipient marked the email as spam.
	// - `delivery` - Amazon SES successfully delivered the email to the recipient's mail server.
	// - `open` - The recipient received the email and opened it in their email client.
	// - `click` - The recipient clicked one or more links in the email.
	// - `renderingFailure` - Amazon SES did not send the email because of a template rendering issue.
	MatchingEventTypes *[]*string `field:"required" json:"matchingEventTypes" yaml:"matchingEventTypes"`
	// An object that contains the names, default values, and sources of the dimensions associated with an Amazon CloudWatch event destination.
	CloudWatchDestination interface{} `field:"optional" json:"cloudWatchDestination" yaml:"cloudWatchDestination"`
	// Sets whether Amazon SES publishes events to this destination when you send an email with the associated configuration set.
	//
	// Set to `true` to enable publishing to this destination; set to `false` to prevent publishing to this destination. The default value is `false` .
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// An object that contains the delivery stream ARN and the IAM role ARN associated with an Amazon Kinesis Firehose event destination.
	KinesisFirehoseDestination interface{} `field:"optional" json:"kinesisFirehoseDestination" yaml:"kinesisFirehoseDestination"`
	// The name of the event destination. The name must meet the following requirements:.
	//
	// - Contain only ASCII letters (a-z, A-Z), numbers (0-9), underscores (_), or dashes (-).
	// - Contain 64 characters or fewer.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// `CfnConfigurationSetEventDestination.EventDestinationProperty.SnsDestination`.
	SnsDestination interface{} `field:"optional" json:"snsDestination" yaml:"snsDestination"`
}

