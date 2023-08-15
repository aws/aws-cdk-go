package awseventstargets

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// Customize the Kinesis Stream Event Target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var ruleTargetInput ruleTargetInput
//
//   kinesisStreamProps := &KinesisStreamProps{
//   	Message: ruleTargetInput,
//   	PartitionKeyPath: jsii.String("partitionKeyPath"),
//   }
//
type KinesisStreamProps struct {
	// The message to send to the stream.
	//
	// Must be a valid JSON text passed to the target stream.
	// Default: - the entire CloudWatch event.
	//
	Message awsevents.RuleTargetInput `field:"optional" json:"message" yaml:"message"`
	// Partition Key Path for records sent to this stream.
	// Default: - eventId as the partition key.
	//
	PartitionKeyPath *string `field:"optional" json:"partitionKeyPath" yaml:"partitionKeyPath"`
}

