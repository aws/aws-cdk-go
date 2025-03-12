package awsevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnEventBus`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var policy interface{}
//
//   cfnEventBusProps := &CfnEventBusProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	DeadLetterConfig: &DeadLetterConfigProperty{
//   		Arn: jsii.String("arn"),
//   	},
//   	Description: jsii.String("description"),
//   	EventSourceName: jsii.String("eventSourceName"),
//   	KmsKeyIdentifier: jsii.String("kmsKeyIdentifier"),
//   	Policy: policy,
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-eventbus.html
//
type CfnEventBusProps struct {
	// The name of the new event bus.
	//
	// Custom event bus names can't contain the `/` character, but you can use the `/` character in partner event bus names. In addition, for partner event buses, the name must exactly match the name of the partner event source that this event bus is matched to.
	//
	// You can't use the name `default` for a custom event bus, as this name is already used for your account's default event bus.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-eventbus.html#cfn-events-eventbus-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Configuration details of the Amazon SQS queue for EventBridge to use as a dead-letter queue (DLQ).
	//
	// For more information, see [Using dead-letter queues to process undelivered events](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-rule-event-delivery.html#eb-rule-dlq) in the *EventBridge User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-eventbus.html#cfn-events-eventbus-deadletterconfig
	//
	DeadLetterConfig interface{} `field:"optional" json:"deadLetterConfig" yaml:"deadLetterConfig"`
	// The event bus description.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-eventbus.html#cfn-events-eventbus-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// If you are creating a partner event bus, this specifies the partner event source that the new event bus will be matched with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-eventbus.html#cfn-events-eventbus-eventsourcename
	//
	EventSourceName *string `field:"optional" json:"eventSourceName" yaml:"eventSourceName"`
	// The identifier of the AWS KMS customer managed key for EventBridge to use, if you choose to use a customer managed key to encrypt events on this event bus.
	//
	// The identifier can be the key Amazon Resource Name (ARN), KeyId, key alias, or key alias ARN.
	//
	// If you do not specify a customer managed key identifier, EventBridge uses an AWS owned key to encrypt events on the event bus.
	//
	// For more information, see [Managing keys](https://docs.aws.amazon.com/kms/latest/developerguide/getting-started.html) in the *AWS Key Management Service Developer Guide* .
	//
	// > Archives and schema discovery are not supported for event buses encrypted using a customer managed key. EventBridge returns an error if:
	// >
	// > - You call `[CreateArchive](https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_CreateArchive.html)` on an event bus set to use a customer managed key for encryption.
	// > - You call `[CreateDiscoverer](https://docs.aws.amazon.com/eventbridge/latest/schema-reference/v1-discoverers.html#CreateDiscoverer)` on an event bus set to use a customer managed key for encryption.
	// > - You call `[UpdatedEventBus](https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_UpdatedEventBus.html)` to set a customer managed key on an event bus with an archives or schema discovery enabled.
	// >
	// > To enable archives or schema discovery on an event bus, choose to use an AWS owned key . For more information, see [Data encryption in EventBridge](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-encryption.html) in the *Amazon EventBridge User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-eventbus.html#cfn-events-eventbus-kmskeyidentifier
	//
	KmsKeyIdentifier *string `field:"optional" json:"kmsKeyIdentifier" yaml:"kmsKeyIdentifier"`
	// The permissions policy of the event bus, describing which other AWS accounts can write events to this event bus.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-eventbus.html#cfn-events-eventbus-policy
	//
	Policy interface{} `field:"optional" json:"policy" yaml:"policy"`
	// Tags to associate with the event bus.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-eventbus.html#cfn-events-eventbus-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

