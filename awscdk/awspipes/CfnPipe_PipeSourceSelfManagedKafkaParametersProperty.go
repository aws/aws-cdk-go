package awspipes


// The parameters for using a self-managed Apache Kafka stream as a source.
//
// A *self managed* cluster refers to any Apache Kafka cluster not hosted by AWS . This includes both clusters you manage yourself, as well as those hosted by a third-party provider, such as [Confluent Cloud](https://docs.aws.amazon.com/https://www.confluent.io/) , [CloudKarafka](https://docs.aws.amazon.com/https://www.cloudkarafka.com/) , or [Redpanda](https://docs.aws.amazon.com/https://redpanda.com/) . For more information, see [Apache Kafka streams as a source](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes-kafka.html) in the *Amazon EventBridge User Guide* .
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourceselfmanagedkafkaparameters.html
//
type CfnPipe_PipeSourceSelfManagedKafkaParametersProperty struct {
	// The name of the topic that the pipe will read from.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourceselfmanagedkafkaparameters.html#cfn-pipes-pipe-pipesourceselfmanagedkafkaparameters-topicname
	//
	TopicName *string `field:"required" json:"topicName" yaml:"topicName"`
	// An array of server URLs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourceselfmanagedkafkaparameters.html#cfn-pipes-pipe-pipesourceselfmanagedkafkaparameters-additionalbootstrapservers
	//
	AdditionalBootstrapServers *[]*string `field:"optional" json:"additionalBootstrapServers" yaml:"additionalBootstrapServers"`
	// The maximum number of records to include in each batch.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourceselfmanagedkafkaparameters.html#cfn-pipes-pipe-pipesourceselfmanagedkafkaparameters-batchsize
	//
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// The name of the destination queue to consume.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourceselfmanagedkafkaparameters.html#cfn-pipes-pipe-pipesourceselfmanagedkafkaparameters-consumergroupid
	//
	ConsumerGroupId *string `field:"optional" json:"consumerGroupId" yaml:"consumerGroupId"`
	// The credentials needed to access the resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourceselfmanagedkafkaparameters.html#cfn-pipes-pipe-pipesourceselfmanagedkafkaparameters-credentials
	//
	Credentials interface{} `field:"optional" json:"credentials" yaml:"credentials"`
	// The maximum length of a time to wait for events.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourceselfmanagedkafkaparameters.html#cfn-pipes-pipe-pipesourceselfmanagedkafkaparameters-maximumbatchingwindowinseconds
	//
	MaximumBatchingWindowInSeconds *float64 `field:"optional" json:"maximumBatchingWindowInSeconds" yaml:"maximumBatchingWindowInSeconds"`
	// The ARN of the Secrets Manager secret used for certification.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourceselfmanagedkafkaparameters.html#cfn-pipes-pipe-pipesourceselfmanagedkafkaparameters-serverrootcacertificate
	//
	ServerRootCaCertificate *string `field:"optional" json:"serverRootCaCertificate" yaml:"serverRootCaCertificate"`
	// The position in a stream from which to start reading.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourceselfmanagedkafkaparameters.html#cfn-pipes-pipe-pipesourceselfmanagedkafkaparameters-startingposition
	//
	StartingPosition *string `field:"optional" json:"startingPosition" yaml:"startingPosition"`
	// This structure specifies the VPC subnets and security groups for the stream, and whether a public IP address is to be used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourceselfmanagedkafkaparameters.html#cfn-pipes-pipe-pipesourceselfmanagedkafkaparameters-vpc
	//
	Vpc interface{} `field:"optional" json:"vpc" yaml:"vpc"`
}

