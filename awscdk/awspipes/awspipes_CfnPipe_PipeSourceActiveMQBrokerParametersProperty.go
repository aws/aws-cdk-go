package awspipes


// The parameters for using an Active MQ broker as a source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pipeSourceActiveMQBrokerParametersProperty := &PipeSourceActiveMQBrokerParametersProperty{
//   	Credentials: &MQBrokerAccessCredentialsProperty{
//   		BasicAuth: jsii.String("basicAuth"),
//   	},
//   	QueueName: jsii.String("queueName"),
//
//   	// the properties below are optional
//   	BatchSize: jsii.Number(123),
//   	MaximumBatchingWindowInSeconds: jsii.Number(123),
//   }
//
type CfnPipe_PipeSourceActiveMQBrokerParametersProperty struct {
	// The credentials needed to access the resource.
	Credentials interface{} `field:"required" json:"credentials" yaml:"credentials"`
	// The name of the destination queue to consume.
	QueueName *string `field:"required" json:"queueName" yaml:"queueName"`
	// The maximum number of records to include in each batch.
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// The maximum length of a time to wait for events.
	MaximumBatchingWindowInSeconds *float64 `field:"optional" json:"maximumBatchingWindowInSeconds" yaml:"maximumBatchingWindowInSeconds"`
}

