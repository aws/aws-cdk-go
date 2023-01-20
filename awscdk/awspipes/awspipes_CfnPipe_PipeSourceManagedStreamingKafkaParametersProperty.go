package awspipes


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
	// `CfnPipe.PipeSourceManagedStreamingKafkaParametersProperty.TopicName`.
	TopicName *string `field:"required" json:"topicName" yaml:"topicName"`
	// `CfnPipe.PipeSourceManagedStreamingKafkaParametersProperty.BatchSize`.
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// `CfnPipe.PipeSourceManagedStreamingKafkaParametersProperty.ConsumerGroupID`.
	ConsumerGroupId *string `field:"optional" json:"consumerGroupId" yaml:"consumerGroupId"`
	// `CfnPipe.PipeSourceManagedStreamingKafkaParametersProperty.Credentials`.
	Credentials interface{} `field:"optional" json:"credentials" yaml:"credentials"`
	// `CfnPipe.PipeSourceManagedStreamingKafkaParametersProperty.MaximumBatchingWindowInSeconds`.
	MaximumBatchingWindowInSeconds *float64 `field:"optional" json:"maximumBatchingWindowInSeconds" yaml:"maximumBatchingWindowInSeconds"`
	// `CfnPipe.PipeSourceManagedStreamingKafkaParametersProperty.StartingPosition`.
	StartingPosition *string `field:"optional" json:"startingPosition" yaml:"startingPosition"`
}

