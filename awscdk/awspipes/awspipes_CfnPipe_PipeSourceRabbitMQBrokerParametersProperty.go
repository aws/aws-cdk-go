package awspipes


// The parameters for using a Rabbit MQ broker as a source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pipeSourceRabbitMQBrokerParametersProperty := &PipeSourceRabbitMQBrokerParametersProperty{
//   	Credentials: &MQBrokerAccessCredentialsProperty{
//   		BasicAuth: jsii.String("basicAuth"),
//   	},
//   	QueueName: jsii.String("queueName"),
//
//   	// the properties below are optional
//   	BatchSize: jsii.Number(123),
//   	MaximumBatchingWindowInSeconds: jsii.Number(123),
//   	VirtualHost: jsii.String("virtualHost"),
//   }
//
type CfnPipe_PipeSourceRabbitMQBrokerParametersProperty struct {
	// The credentials needed to access the resource.
	Credentials interface{} `field:"required" json:"credentials" yaml:"credentials"`
	// The name of the destination queue to consume.
	QueueName *string `field:"required" json:"queueName" yaml:"queueName"`
	// The maximum number of records to include in each batch.
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// The maximum length of a time to wait for events.
	MaximumBatchingWindowInSeconds *float64 `field:"optional" json:"maximumBatchingWindowInSeconds" yaml:"maximumBatchingWindowInSeconds"`
	// The name of the virtual host associated with the source broker.
	VirtualHost *string `field:"optional" json:"virtualHost" yaml:"virtualHost"`
}

