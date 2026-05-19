package awsbedrockagentcorealpha


// Configuration parameters for a self managed memory strategy existing built-in default prompts/models.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var topic Topic
//
//   selfManagedStrategyProps := &SelfManagedStrategyProps{
//   	InvocationConfiguration: &InvocationConfiguration{
//   		S3Location: &Location{
//   			BucketName: jsii.String("bucketName"),
//   			ObjectKey: jsii.String("objectKey"),
//
//   			// the properties below are optional
//   			ObjectVersion: jsii.String("objectVersion"),
//   		},
//   		Topic: topic,
//   	},
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	HistoricalContextWindowSize: jsii.Number(123),
//   	TriggerConditions: &TriggerConditions{
//   		MessageBasedTrigger: jsii.Number(123),
//   		TimeBasedTrigger: cdk.Duration_Minutes(jsii.Number(30)),
//   		TokenBasedTrigger: jsii.Number(123),
//   	},
//   }
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

