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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourceactivemqbrokerparameters.html
//
type CfnPipe_PipeSourceActiveMQBrokerParametersProperty struct {
	// The credentials needed to access the resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourceactivemqbrokerparameters.html#cfn-pipes-pipe-pipesourceactivemqbrokerparameters-credentials
	//
	Credentials interface{} `field:"required" json:"credentials" yaml:"credentials"`
	// The name of the destination queue to consume.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourceactivemqbrokerparameters.html#cfn-pipes-pipe-pipesourceactivemqbrokerparameters-queuename
	//
	QueueName *string `field:"required" json:"queueName" yaml:"queueName"`
	// The maximum number of records to include in each batch.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourceactivemqbrokerparameters.html#cfn-pipes-pipe-pipesourceactivemqbrokerparameters-batchsize
	//
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// The maximum length of a time to wait for events.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourceactivemqbrokerparameters.html#cfn-pipes-pipe-pipesourceactivemqbrokerparameters-maximumbatchingwindowinseconds
	//
	MaximumBatchingWindowInSeconds *float64 `field:"optional" json:"maximumBatchingWindowInSeconds" yaml:"maximumBatchingWindowInSeconds"`
}

