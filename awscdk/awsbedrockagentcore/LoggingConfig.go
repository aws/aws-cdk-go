package awsbedrockagentcore


// Configuration for logging with log type and destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var loggingDestination LoggingDestination
//   var logType LogType
//
//   loggingConfig := &LoggingConfig{
//   	Destination: loggingDestination,
//   	LogType: logType,
//   }
//
type LoggingConfig struct {
	// The destination for logs.
	Destination LoggingDestination `field:"required" json:"destination" yaml:"destination"`
	// The type of logs to deliver.
	LogType LogType `field:"required" json:"logType" yaml:"logType"`
}

