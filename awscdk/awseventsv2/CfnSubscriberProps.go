package awseventsv2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnSubscriber`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSubscriberProps := &CfnSubscriberProps{
//   	EventBusArn: jsii.String("eventBusArn"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	BatchConfiguration: &BatchConfigurationProperty{
//   		MaxBatchSize: jsii.Number(123),
//   		MaxBatchWindowInSeconds: jsii.Number(123),
//   	},
//   	Description: jsii.String("description"),
//   	FilterConfiguration: &FilterConfigurationProperty{
//   		Filters: []interface{}{
//   			&FilterProperty{
//   				Pattern: jsii.String("pattern"),
//   				Scope: jsii.String("scope"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		Language: jsii.String("language"),
//   	},
//   	LogConfiguration: &LogConfigurationProperty{
//   		IncludePayload: jsii.String("includePayload"),
//   		Level: jsii.String("level"),
//   	},
//   	OnFailureConfiguration: &OnFailureConfigurationProperty{
//   		Arn: jsii.String("arn"),
//   	},
//   	PointInTimeConfiguration: &PointInTimeConfigurationProperty{
//   		PointType: jsii.String("pointType"),
//
//   		// the properties below are optional
//   		EndPoint: jsii.Number(123),
//   		StartingPoint: jsii.Number(123),
//   	},
//   	ResumePosition: jsii.String("resumePosition"),
//   	RetryPolicy: &RetryPolicyProperty{
//   		MaxEventAgeInSeconds: jsii.Number(123),
//   		MaxRetryAttempts: jsii.Number(123),
//   		RetryStrategy: jsii.String("retryStrategy"),
//   	},
//   	StartingPosition: jsii.String("startingPosition"),
//   	State: jsii.String("state"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Transformer: &TransformerProperty{
//   		JsonataConfiguration: &JsonataConfigurationProperty{
//   			Expression: jsii.String("expression"),
//   		},
//   		Type: jsii.String("type"),
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventsv2-subscriber.html
//
type CfnSubscriberProps struct {
	// The ARN of the event bus this subscriber belongs to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventsv2-subscriber.html#cfn-eventsv2-subscriber-eventbusarn
	//
	EventBusArn *string `field:"required" json:"eventBusArn" yaml:"eventBusArn"`
	// The name of the subscriber.
	//
	// The first character must be alphanumeric; the remaining characters may also include '.', '-', and '_'.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventsv2-subscriber.html#cfn-eventsv2-subscriber-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Configuration for batching events into a single delivery.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventsv2-subscriber.html#cfn-eventsv2-subscriber-batchconfiguration
	//
	BatchConfiguration interface{} `field:"optional" json:"batchConfiguration" yaml:"batchConfiguration"`
	// A description of the subscriber.
	//
	// Control characters and Unicode line separators are not allowed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventsv2-subscriber.html#cfn-eventsv2-subscriber-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Configuration for filtering which events are delivered to the target.
	//
	// An event must match every filter to be delivered.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventsv2-subscriber.html#cfn-eventsv2-subscriber-filterconfiguration
	//
	FilterConfiguration interface{} `field:"optional" json:"filterConfiguration" yaml:"filterConfiguration"`
	// Delivery logging configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventsv2-subscriber.html#cfn-eventsv2-subscriber-logconfiguration
	//
	LogConfiguration interface{} `field:"optional" json:"logConfiguration" yaml:"logConfiguration"`
	// The destination for events that could not be delivered.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventsv2-subscriber.html#cfn-eventsv2-subscriber-onfailureconfiguration
	//
	OnFailureConfiguration interface{} `field:"optional" json:"onFailureConfiguration" yaml:"onFailureConfiguration"`
	// The point in time to start delivering events from, used when StartingPosition is POINT_IN_TIME.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventsv2-subscriber.html#cfn-eventsv2-subscriber-pointintimeconfiguration
	//
	PointInTimeConfiguration interface{} `field:"optional" json:"pointInTimeConfiguration" yaml:"pointInTimeConfiguration"`
	// Resume-time control, never returned by the service.
	//
	// Applied only when an update transitions State from STOPPED to RUNNING: LAST_PROCESSED (default) resumes from the last processed event, LATEST skips to the newest. Ignored on create and on any update that does not perform that transition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventsv2-subscriber.html#cfn-eventsv2-subscriber-resumeposition
	//
	ResumePosition *string `field:"optional" json:"resumePosition" yaml:"resumePosition"`
	// The retry policy for failed deliveries.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventsv2-subscriber.html#cfn-eventsv2-subscriber-retrypolicy
	//
	RetryPolicy interface{} `field:"optional" json:"retryPolicy" yaml:"retryPolicy"`
	// Where the subscriber starts reading events: LATEST starts from the newest events;
	//
	// POINT_IN_TIME starts from the point specified in PointInTimeConfiguration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventsv2-subscriber.html#cfn-eventsv2-subscriber-startingposition
	//
	StartingPosition *string `field:"optional" json:"startingPosition" yaml:"startingPosition"`
	// The run state of the subscriber.
	//
	// Events are delivered only while the state is RUNNING. Setting the state to STOPPED pauses delivery. When an update sets a stopped subscriber back to RUNNING, ResumePosition controls where delivery resumes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventsv2-subscriber.html#cfn-eventsv2-subscriber-state
	//
	State *string `field:"optional" json:"state" yaml:"state"`
	// The tags assigned to the subscriber.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventsv2-subscriber.html#cfn-eventsv2-subscriber-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Configuration for transforming events before delivery: the raw payload, the payload with its metadata envelope, or the output of a JSONata expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventsv2-subscriber.html#cfn-eventsv2-subscriber-transformer
	//
	Transformer interface{} `field:"optional" json:"transformer" yaml:"transformer"`
	// The delivery ordering mode of the subscriber.
	//
	// FIFO delivers events in order within an event group; UNORDERED delivers without an ordering guarantee.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventsv2-subscriber.html#cfn-eventsv2-subscriber-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

