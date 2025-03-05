package awskinesis


// Properties for a Kinesis Stream Consumer.
//
// Example:
//   lambdaRole := iam.NewRole(this, jsii.String("Role"), &RoleProps{
//   	AssumedBy: iam.NewServicePrincipal(jsii.String("lambda.amazonaws.com")),
//   	Description: jsii.String("Example role..."),
//   })
//
//   stream := kinesis.NewStream(this, jsii.String("MyEncryptedStream"), &StreamProps{
//   	Encryption: kinesis.StreamEncryption_KMS,
//   })
//   streamConsumer := kinesis.NewStreamConsumer(this, jsii.String("MyStreamConsumer"), &StreamConsumerProps{
//   	StreamConsumerName: jsii.String("MyStreamConsumer"),
//   	Stream: Stream,
//   })
//
//   // give lambda permissions to read stream via the stream consumer
//   streamConsumer.grantRead(lambdaRole)
//
type StreamConsumerProps struct {
	// The Kinesis data stream to associate this consumer with.
	Stream IStream `field:"required" json:"stream" yaml:"stream"`
	// The name of the stream consumer.
	StreamConsumerName *string `field:"required" json:"streamConsumerName" yaml:"streamConsumerName"`
}

