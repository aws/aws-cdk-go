package awsevents


// Interface for Logging Configuration of the Event Bus.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   bus := awscdk.NewEventBus(this, jsii.String("Bus"), &EventBusProps{
//   	LogConfig: &LogConfig{
//   		IncludeDetail: awscdk.IncludeDetail_FULL,
//   		Level: awscdk.Level_TRACE,
//   	},
//   })
//
type LogConfig struct {
	// Whether EventBridge include detailed event information in the records it generates.
	// Default: no details.
	//
	IncludeDetail IncludeDetail `field:"optional" json:"includeDetail" yaml:"includeDetail"`
	// Logging level.
	// Default: OFF.
	//
	Level Level `field:"optional" json:"level" yaml:"level"`
}

