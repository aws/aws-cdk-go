package awspipes


// The parameters for using a self-managed Apache Kafka stream as a source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pipeSourceSelfManagedKafkaParametersProperty := &PipeSourceSelfManagedKafkaParametersProperty{
//   	TopicName: jsii.String("topicName"),
//
//   	// the properties below are optional
//   	AdditionalBootstrapServers: []*string{
//   		jsii.String("additionalBootstrapServers"),
//   	},
//   	BatchSize: jsii.Number(123),
//   	ConsumerGroupId: jsii.String("consumerGroupId"),
//   	Credentials: &SelfManagedKafkaAccessConfigurationCredentialsProperty{
//   		BasicAuth: jsii.String("basicAuth"),
//   		ClientCertificateTlsAuth: jsii.String("clientCertificateTlsAuth"),
//   		SaslScram256Auth: jsii.String("saslScram256Auth"),
//   		SaslScram512Auth: jsii.String("saslScram512Auth"),
//   	},
//   	MaximumBatchingWindowInSeconds: jsii.Number(123),
//   	ServerRootCaCertificate: jsii.String("serverRootCaCertificate"),
//   	StartingPosition: jsii.String("startingPosition"),
//   	Vpc: &SelfManagedKafkaAccessConfigurationVpcProperty{
//   		SecurityGroup: []*string{
//   			jsii.String("securityGroup"),
//   		},
//   		Subnets: []*string{
//   			jsii.String("subnets"),
//   		},
//   	},
//   }
//
type CfnPipe_PipeSourceSelfManagedKafkaParametersProperty struct {
	// The name of the topic that the pipe will read from.
	TopicName *string `field:"required" json:"topicName" yaml:"topicName"`
	// An array of server URLs.
	AdditionalBootstrapServers *[]*string `field:"optional" json:"additionalBootstrapServers" yaml:"additionalBootstrapServers"`
	// The maximum number of records to include in each batch.
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// The name of the destination queue to consume.
	ConsumerGroupId *string `field:"optional" json:"consumerGroupId" yaml:"consumerGroupId"`
	// The credentials needed to access the resource.
	Credentials interface{} `field:"optional" json:"credentials" yaml:"credentials"`
	// The maximum length of a time to wait for events.
	MaximumBatchingWindowInSeconds *float64 `field:"optional" json:"maximumBatchingWindowInSeconds" yaml:"maximumBatchingWindowInSeconds"`
	// The ARN of the Secrets Manager secret used for certification.
	ServerRootCaCertificate *string `field:"optional" json:"serverRootCaCertificate" yaml:"serverRootCaCertificate"`
	// (Streams only) The position in a stream from which to start reading.
	StartingPosition *string `field:"optional" json:"startingPosition" yaml:"startingPosition"`
	// This structure specifies the VPC subnets and security groups for the stream, and whether a public IP address is to be used.
	Vpc interface{} `field:"optional" json:"vpc" yaml:"vpc"`
}

