package awspipes


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pipeSourceRabbitMQBrokerParametersProperty := &pipeSourceRabbitMQBrokerParametersProperty{
//   	credentials: &mQBrokerAccessCredentialsProperty{
//   		basicAuth: jsii.String("basicAuth"),
//   	},
//   	queueName: jsii.String("queueName"),
//
//   	// the properties below are optional
//   	batchSize: jsii.Number(123),
//   	maximumBatchingWindowInSeconds: jsii.Number(123),
//   	virtualHost: jsii.String("virtualHost"),
//   }
//
type CfnPipe_PipeSourceRabbitMQBrokerParametersProperty struct {
	// `CfnPipe.PipeSourceRabbitMQBrokerParametersProperty.Credentials`.
	Credentials interface{} `field:"required" json:"credentials" yaml:"credentials"`
	// `CfnPipe.PipeSourceRabbitMQBrokerParametersProperty.QueueName`.
	QueueName *string `field:"required" json:"queueName" yaml:"queueName"`
	// `CfnPipe.PipeSourceRabbitMQBrokerParametersProperty.BatchSize`.
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// `CfnPipe.PipeSourceRabbitMQBrokerParametersProperty.MaximumBatchingWindowInSeconds`.
	MaximumBatchingWindowInSeconds *float64 `field:"optional" json:"maximumBatchingWindowInSeconds" yaml:"maximumBatchingWindowInSeconds"`
	// `CfnPipe.PipeSourceRabbitMQBrokerParametersProperty.VirtualHost`.
	VirtualHost *string `field:"optional" json:"virtualHost" yaml:"virtualHost"`
}

