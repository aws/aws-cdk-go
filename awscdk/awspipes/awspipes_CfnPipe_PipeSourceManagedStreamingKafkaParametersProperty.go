package awspipes


// The parameters for using an MSK stream as a source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pipeSourceManagedStreamingKafkaParametersProperty := &pipeSourceManagedStreamingKafkaParametersProperty{
//   	topicName: jsii.String("topicName"),
//
//   	// the properties below are optional
//   	batchSize: jsii.Number(123),
//   	consumerGroupId: jsii.String("consumerGroupId"),
//   	credentials: &mSKAccessCredentialsProperty{
//   		clientCertificateTlsAuth: jsii.String("clientCertificateTlsAuth"),
//   		saslScram512Auth: jsii.String("saslScram512Auth"),
//   	},
//   	maximumBatchingWindowInSeconds: jsii.Number(123),
//   	startingPosition: jsii.String("startingPosition"),
//   }
//
type CfnPipe_PipeSourceManagedStreamingKafkaParametersProperty struct {
	// The name of the topic that the pipe will read from.
	TopicName *string `field:"required" json:"topicName" yaml:"topicName"`
	// The maximum number of records to include in each batch.
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// The name of the destination queue to consume.
	ConsumerGroupId *string `field:"optional" json:"consumerGroupId" yaml:"consumerGroupId"`
	// The credentials needed to access the resource.
	Credentials interface{} `field:"optional" json:"credentials" yaml:"credentials"`
	// The maximum length of a time to wait for events.
	MaximumBatchingWindowInSeconds *float64 `field:"optional" json:"maximumBatchingWindowInSeconds" yaml:"maximumBatchingWindowInSeconds"`
	// (Streams only) The position in a stream from which to start reading.
	StartingPosition *string `field:"optional" json:"startingPosition" yaml:"startingPosition"`
}

