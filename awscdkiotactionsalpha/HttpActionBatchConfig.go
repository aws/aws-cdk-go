package awscdkiotactionsalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Configuration for batching HTTP action messages.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var topicRule TopicRule
//
//
//   topicRule.AddAction(
//   actions.NewHttpsAction(jsii.String("https://example.com/endpoint"), &HttpsActionProps{
//   	BatchConfig: &HttpActionBatchConfig{
//   		MaxBatchOpenDuration: awscdk.Duration_Millis(jsii.Number(100)),
//   		MaxBatchSize: jsii.Number(5),
//   		MaxBatchSizeBytes: awscdk.Size_Kibibytes(jsii.Number(1)),
//   	},
//   }))
//
// See: https://docs.aws.amazon.com/iot/latest/developerguide/http_batching.html
//
// Experimental.
type HttpActionBatchConfig struct {
	// The maximum amount of time an outgoing message waits for other messages to create the batch.
	//
	// Must be between 5 ms and 200 ms.
	// Default: Duration.millis(20)
	//
	// Experimental.
	MaxBatchOpenDuration awscdk.Duration `field:"optional" json:"maxBatchOpenDuration" yaml:"maxBatchOpenDuration"`
	// The maximum number of messages that are batched together in a single IoT rule action execution.
	//
	// Must be between 2 and 10.
	// Default: 10.
	//
	// Experimental.
	MaxBatchSize *float64 `field:"optional" json:"maxBatchSize" yaml:"maxBatchSize"`
	// Maximum size of a message batch.
	//
	// Must be between 100 bytes and 128 KiB.
	// Default: Size.kibibytes(5)
	//
	// Experimental.
	MaxBatchSizeBytes awscdk.Size `field:"optional" json:"maxBatchSizeBytes" yaml:"maxBatchSizeBytes"`
}

