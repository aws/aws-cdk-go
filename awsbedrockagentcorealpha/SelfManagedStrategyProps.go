package awsbedrockagentcorealpha


// Configuration parameters for a self managed memory strategy existing built-in default prompts/models.
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
type SelfManagedStrategyProps struct {
	// The name for the strategy.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description of the strategy.
	// Default: - No description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Invocation configuration for self managed memory strategy.
	// Experimental.
	InvocationConfiguration *InvocationConfiguration `field:"required" json:"invocationConfiguration" yaml:"invocationConfiguration"`
	// Define the number of previous events to be included when processing memory.
	//
	// A larger history window provides more context from past conversations.
	// Default: 4.
	//
	// Experimental.
	HistoricalContextWindowSize *float64 `field:"optional" json:"historicalContextWindowSize" yaml:"historicalContextWindowSize"`
	// Trigger conditions for self managed memory strategy.
	// Default: - undefined.
	//
	// Experimental.
	TriggerConditions *TriggerConditions `field:"optional" json:"triggerConditions" yaml:"triggerConditions"`
}

