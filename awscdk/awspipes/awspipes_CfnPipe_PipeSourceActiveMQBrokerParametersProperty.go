package awspipes


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pipeSourceActiveMQBrokerParametersProperty := &pipeSourceActiveMQBrokerParametersProperty{
//   	credentials: &mQBrokerAccessCredentialsProperty{
//   		basicAuth: jsii.String("basicAuth"),
//   	},
//   	queueName: jsii.String("queueName"),
//
//   	// the properties below are optional
//   	batchSize: jsii.Number(123),
//   	maximumBatchingWindowInSeconds: jsii.Number(123),
//   }
//
type CfnPipe_PipeSourceActiveMQBrokerParametersProperty struct {
	// `CfnPipe.PipeSourceActiveMQBrokerParametersProperty.Credentials`.
	Credentials interface{} `field:"required" json:"credentials" yaml:"credentials"`
	// `CfnPipe.PipeSourceActiveMQBrokerParametersProperty.QueueName`.
	QueueName *string `field:"required" json:"queueName" yaml:"queueName"`
	// `CfnPipe.PipeSourceActiveMQBrokerParametersProperty.BatchSize`.
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// `CfnPipe.PipeSourceActiveMQBrokerParametersProperty.MaximumBatchingWindowInSeconds`.
	MaximumBatchingWindowInSeconds *float64 `field:"optional" json:"maximumBatchingWindowInSeconds" yaml:"maximumBatchingWindowInSeconds"`
}

