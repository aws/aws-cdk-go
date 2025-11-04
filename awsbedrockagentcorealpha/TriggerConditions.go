package awsbedrockagentcorealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Trigger conditions for self managed memory strategy When first condition is met, batched payloads are sent to specified S3 bucket.
//
// Example:
//   bucket := s3.NewBucket(this, jsii.String("memoryBucket"), &BucketProps{
//   	BucketName: jsii.String("test-memory"),
//   	RemovalPolicy: cdk.RemovalPolicy_DESTROY,
//   	AutoDeleteObjects: jsii.Boolean(true),
//   })
//
//   topic := sns.NewTopic(this, jsii.String("topic"))
//
//   // Create a custom semantic memory strategy
//   selfManagedStrategy := agentcore.MemoryStrategy_UsingSelfManaged(&SelfManagedStrategyProps{
//   	Name: jsii.String("selfManagedStrategy"),
//   	Description: jsii.String("self managed memory strategy"),
//   	HistoricalContextWindowSize: jsii.Number(5),
//   	InvocationConfiguration: &InvocationConfiguration{
//   		Topic: topic,
//   		S3Location: &Location{
//   			BucketName: bucket.BucketName,
//   			ObjectKey: jsii.String("memory/"),
//   		},
//   	},
//   	TriggerConditions: &TriggerConditions{
//   		MessageBasedTrigger: jsii.Number(1),
//   		TimeBasedTrigger: cdk.Duration_Seconds(jsii.Number(10)),
//   		TokenBasedTrigger: jsii.Number(100),
//   	},
//   })
//
//   // Create memory with custom strategy
//   memory := agentcore.NewMemory(this, jsii.String("MyMemory"), &MemoryProps{
//   	MemoryName: jsii.String("my-custom-memory"),
//   	Description: jsii.String("Memory with custom strategy"),
//   	ExpirationDuration: cdk.Duration_Days(jsii.Number(90)),
//   	MemoryStrategies: []IMemoryStrategy{
//   		selfManagedStrategy,
//   	},
//   })
//
// Experimental.
type TriggerConditions struct {
	// Triggers memory processing when specified number of new messages is reached.
	// Default: - 1.
	//
	// Experimental.
	MessageBasedTrigger *float64 `field:"optional" json:"messageBasedTrigger" yaml:"messageBasedTrigger"`
	// Triggers memory processing when the session has been idle for the specified duration.
	//
	// Value in seconds.
	// Default: - 10.
	//
	// Experimental.
	TimeBasedTrigger awscdk.Duration `field:"optional" json:"timeBasedTrigger" yaml:"timeBasedTrigger"`
	// Triggers memory processing when the token size reaches the specified threshold.
	// Default: - 100.
	//
	// Experimental.
	TokenBasedTrigger *float64 `field:"optional" json:"tokenBasedTrigger" yaml:"tokenBasedTrigger"`
}

