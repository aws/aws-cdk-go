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
// Experimental.
type LoggingConfig struct {
	// The destination for logs.
	// Experimental.
	Destination LoggingDestination `field:"required" json:"destination" yaml:"destination"`
	// The type of logs to deliver.
	// Experimental.
	LogType LogType `field:"required" json:"logType" yaml:"logType"`
}

