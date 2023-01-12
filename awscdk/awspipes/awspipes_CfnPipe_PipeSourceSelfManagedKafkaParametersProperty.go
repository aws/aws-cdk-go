package awspipes


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pipeSourceSelfManagedKafkaParametersProperty := &pipeSourceSelfManagedKafkaParametersProperty{
//   	topicName: jsii.String("topicName"),
//
//   	// the properties below are optional
//   	additionalBootstrapServers: []*string{
//   		jsii.String("additionalBootstrapServers"),
//   	},
//   	batchSize: jsii.Number(123),
//   	consumerGroupId: jsii.String("consumerGroupId"),
//   	credentials: &selfManagedKafkaAccessConfigurationCredentialsProperty{
//   		basicAuth: jsii.String("basicAuth"),
//   		clientCertificateTlsAuth: jsii.String("clientCertificateTlsAuth"),
//   		saslScram256Auth: jsii.String("saslScram256Auth"),
//   		saslScram512Auth: jsii.String("saslScram512Auth"),
//   	},
//   	maximumBatchingWindowInSeconds: jsii.Number(123),
//   	serverRootCaCertificate: jsii.String("serverRootCaCertificate"),
//   	startingPosition: jsii.String("startingPosition"),
//   	vpc: &selfManagedKafkaAccessConfigurationVpcProperty{
//   		securityGroup: []*string{
//   			jsii.String("securityGroup"),
//   		},
//   		subnets: []*string{
//   			jsii.String("subnets"),
//   		},
//   	},
//   }
//
type CfnPipe_PipeSourceSelfManagedKafkaParametersProperty struct {
	// `CfnPipe.PipeSourceSelfManagedKafkaParametersProperty.TopicName`.
	TopicName *string `field:"required" json:"topicName" yaml:"topicName"`
	// `CfnPipe.PipeSourceSelfManagedKafkaParametersProperty.AdditionalBootstrapServers`.
	AdditionalBootstrapServers *[]*string `field:"optional" json:"additionalBootstrapServers" yaml:"additionalBootstrapServers"`
	// `CfnPipe.PipeSourceSelfManagedKafkaParametersProperty.BatchSize`.
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// `CfnPipe.PipeSourceSelfManagedKafkaParametersProperty.ConsumerGroupID`.
	ConsumerGroupId *string `field:"optional" json:"consumerGroupId" yaml:"consumerGroupId"`
	// `CfnPipe.PipeSourceSelfManagedKafkaParametersProperty.Credentials`.
	Credentials interface{} `field:"optional" json:"credentials" yaml:"credentials"`
	// `CfnPipe.PipeSourceSelfManagedKafkaParametersProperty.MaximumBatchingWindowInSeconds`.
	MaximumBatchingWindowInSeconds *float64 `field:"optional" json:"maximumBatchingWindowInSeconds" yaml:"maximumBatchingWindowInSeconds"`
	// `CfnPipe.PipeSourceSelfManagedKafkaParametersProperty.ServerRootCaCertificate`.
	ServerRootCaCertificate *string `field:"optional" json:"serverRootCaCertificate" yaml:"serverRootCaCertificate"`
	// `CfnPipe.PipeSourceSelfManagedKafkaParametersProperty.StartingPosition`.
	StartingPosition *string `field:"optional" json:"startingPosition" yaml:"startingPosition"`
	// `CfnPipe.PipeSourceSelfManagedKafkaParametersProperty.Vpc`.
	Vpc interface{} `field:"optional" json:"vpc" yaml:"vpc"`
}

