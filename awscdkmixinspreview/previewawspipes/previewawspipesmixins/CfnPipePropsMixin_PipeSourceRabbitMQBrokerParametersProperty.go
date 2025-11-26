package previewawspipesmixins


// The parameters for using a Rabbit MQ broker as a source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   pipeSourceRabbitMQBrokerParametersProperty := &PipeSourceRabbitMQBrokerParametersProperty{
//   	BatchSize: jsii.Number(123),
//   	Credentials: &MQBrokerAccessCredentialsProperty{
//   		BasicAuth: jsii.String("basicAuth"),
//   	},
//   	MaximumBatchingWindowInSeconds: jsii.Number(123),
//   	QueueName: jsii.String("queueName"),
//   	VirtualHost: jsii.String("virtualHost"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourcerabbitmqbrokerparameters.html
//
type CfnPipePropsMixin_PipeSourceRabbitMQBrokerParametersProperty struct {
	// The maximum number of records to include in each batch.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourcerabbitmqbrokerparameters.html#cfn-pipes-pipe-pipesourcerabbitmqbrokerparameters-batchsize
	//
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// The credentials needed to access the resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourcerabbitmqbrokerparameters.html#cfn-pipes-pipe-pipesourcerabbitmqbrokerparameters-credentials
	//
	Credentials interface{} `field:"optional" json:"credentials" yaml:"credentials"`
	// The maximum length of a time to wait for events.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourcerabbitmqbrokerparameters.html#cfn-pipes-pipe-pipesourcerabbitmqbrokerparameters-maximumbatchingwindowinseconds
	//
	MaximumBatchingWindowInSeconds *float64 `field:"optional" json:"maximumBatchingWindowInSeconds" yaml:"maximumBatchingWindowInSeconds"`
	// The name of the destination queue to consume.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourcerabbitmqbrokerparameters.html#cfn-pipes-pipe-pipesourcerabbitmqbrokerparameters-queuename
	//
	QueueName *string `field:"optional" json:"queueName" yaml:"queueName"`
	// The name of the virtual host associated with the source broker.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourcerabbitmqbrokerparameters.html#cfn-pipes-pipe-pipesourcerabbitmqbrokerparameters-virtualhost
	//
	VirtualHost *string `field:"optional" json:"virtualHost" yaml:"virtualHost"`
}

