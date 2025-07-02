package awsbedrockalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for SessionSummaryConfiguration.
//
// Example:
//   agent := bedrock.NewAgent(this, jsii.String("MyAgent"), &AgentProps{
//   	AgentName: jsii.String("MyAgent"),
//   	Instruction: jsii.String("Your instruction here"),
//   	FoundationModel: bedrock.BedrockFoundationModel_AMAZON_NOVA_LITE_V1(),
//   	Memory: awsbedrockalpha.Memory_SessionSummary(&SessionSummaryMemoryProps{
//   		MaxRecentSessions: jsii.Number(10),
//   		 // Keep the last 10 session summaries
//   		MemoryDuration: awscdk.Duration_Days(jsii.Number(20)),
//   	}),
//   })
//
// Experimental.
type SessionSummaryMemoryProps struct {
	// Maximum number of recent session summaries to include (min 1).
	// Default: 20.
	//
	// Experimental.
	MaxRecentSessions *float64 `field:"optional" json:"maxRecentSessions" yaml:"maxRecentSessions"`
	// Duration for which session summaries are retained (between 1 and 365 days).
	// Default: Duration.days(30)
	//
	// Experimental.
	MemoryDuration awscdk.Duration `field:"optional" json:"memoryDuration" yaml:"memoryDuration"`
}

