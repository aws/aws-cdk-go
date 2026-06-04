package awsbedrockagentcorealpha


// Configuration for logging with log type and destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   var loggingDestination LoggingDestination
//   var logType LogType
//
//   loggingConfig := &LoggingConfig{
//   	Destination: loggingDestination,
//   	LogType: logType,
//   }
//
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
type LoggingConfig struct {
	// The destination for logs.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	Destination LoggingDestination `field:"required" json:"destination" yaml:"destination"`
	// The type of logs to deliver.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	LogType LogType `field:"required" json:"logType" yaml:"logType"`
}

