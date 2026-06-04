package awsbedrockagentcorealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Trigger conditions for self managed memory strategy When first condition is met, batched payloads are sent to specified S3 bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   triggerConditions := &TriggerConditions{
//   	MessageBasedTrigger: jsii.Number(123),
//   	TimeBasedTrigger: cdk.Duration_Minutes(jsii.Number(30)),
//   	TokenBasedTrigger: jsii.Number(123),
//   }
//
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
type TriggerConditions struct {
	// Triggers memory processing when specified number of new messages is reached.
	// Default: 1.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	MessageBasedTrigger *float64 `field:"optional" json:"messageBasedTrigger" yaml:"messageBasedTrigger"`
	// Triggers memory processing when the session has been idle for the specified duration.
	//
	// Value in seconds.
	// Default: - 10 seconds.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	TimeBasedTrigger awscdk.Duration `field:"optional" json:"timeBasedTrigger" yaml:"timeBasedTrigger"`
	// Triggers memory processing when the token size reaches the specified threshold.
	// Default: 100.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	TokenBasedTrigger *float64 `field:"optional" json:"tokenBasedTrigger" yaml:"tokenBasedTrigger"`
}

