package previewawspipesmixins


// The parameters required to set up a source for your pipe.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   pipeSourceParametersProperty := &PipeSourceParametersProperty{
//   	ActiveMqBrokerParameters: &PipeSourceActiveMQBrokerParametersProperty{
//   		BatchSize: jsii.Number(123),
//   		Credentials: &MQBrokerAccessCredentialsProperty{
//   			BasicAuth: jsii.String("basicAuth"),
//   		},
//   		MaximumBatchingWindowInSeconds: jsii.Number(123),
//   		QueueName: jsii.String("queueName"),
//   	},
//   	DynamoDbStreamParameters: &PipeSourceDynamoDBStreamParametersProperty{
//   		BatchSize: jsii.Number(123),
//   		DeadLetterConfig: &DeadLetterConfigProperty{
//   			Arn: jsii.String("arn"),
//   		},
//   		MaximumBatchingWindowInSeconds: jsii.Number(123),
//   		MaximumRecordAgeInSeconds: jsii.Number(123),
//   		MaximumRetryAttempts: jsii.Number(123),
//   		OnPartialBatchItemFailure: jsii.String("onPartialBatchItemFailure"),
//   		ParallelizationFactor: jsii.Number(123),
//   		StartingPosition: jsii.String("startingPosition"),
//   	},
//   	FilterCriteria: &FilterCriteriaProperty{
//   		Filters: []interface{}{
//   			&FilterProperty{
//   				Pattern: jsii.String("pattern"),
//   			},
//   		},
//   	},
//   	KinesisStreamParameters: &PipeSourceKinesisStreamParametersProperty{
//   		BatchSize: jsii.Number(123),
//   		DeadLetterConfig: &DeadLetterConfigProperty{
//   			Arn: jsii.String("arn"),
//   		},
//   		MaximumBatchingWindowInSeconds: jsii.Number(123),
//   		MaximumRecordAgeInSeconds: jsii.Number(123),
//   		MaximumRetryAttempts: jsii.Number(123),
//   		OnPartialBatchItemFailure: jsii.String("onPartialBatchItemFailure"),
//   		ParallelizationFactor: jsii.Number(123),
//   		StartingPosition: jsii.String("startingPosition"),
//   		StartingPositionTimestamp: jsii.String("startingPositionTimestamp"),
//   	},
//   	ManagedStreamingKafkaParameters: &PipeSourceManagedStreamingKafkaParametersProperty{
//   		BatchSize: jsii.Number(123),
//   		ConsumerGroupId: jsii.String("consumerGroupId"),
//   		Credentials: &MSKAccessCredentialsProperty{
//   			ClientCertificateTlsAuth: jsii.String("clientCertificateTlsAuth"),
//   			SaslScram512Auth: jsii.String("saslScram512Auth"),
//   		},
//   		MaximumBatchingWindowInSeconds: jsii.Number(123),
//   		StartingPosition: jsii.String("startingPosition"),
//   		TopicName: jsii.String("topicName"),
//   	},
//   	RabbitMqBrokerParameters: &PipeSourceRabbitMQBrokerParametersProperty{
//   		BatchSize: jsii.Number(123),
//   		Credentials: &MQBrokerAccessCredentialsProperty{
//   			BasicAuth: jsii.String("basicAuth"),
//   		},
//   		MaximumBatchingWindowInSeconds: jsii.Number(123),
//   		QueueName: jsii.String("queueName"),
//   		VirtualHost: jsii.String("virtualHost"),
//   	},
//   	SelfManagedKafkaParameters: &PipeSourceSelfManagedKafkaParametersProperty{
//   		AdditionalBootstrapServers: []*string{
//   			jsii.String("additionalBootstrapServers"),
//   		},
//   		BatchSize: jsii.Number(123),
//   		ConsumerGroupId: jsii.String("consumerGroupId"),
//   		Credentials: &SelfManagedKafkaAccessConfigurationCredentialsProperty{
//   			BasicAuth: jsii.String("basicAuth"),
//   			ClientCertificateTlsAuth: jsii.String("clientCertificateTlsAuth"),
//   			SaslScram256Auth: jsii.String("saslScram256Auth"),
//   			SaslScram512Auth: jsii.String("saslScram512Auth"),
//   		},
//   		MaximumBatchingWindowInSeconds: jsii.Number(123),
//   		ServerRootCaCertificate: jsii.String("serverRootCaCertificate"),
//   		StartingPosition: jsii.String("startingPosition"),
//   		TopicName: jsii.String("topicName"),
//   		Vpc: &SelfManagedKafkaAccessConfigurationVpcProperty{
//   			SecurityGroup: []*string{
//   				jsii.String("securityGroup"),
//   			},
//   			Subnets: []*string{
//   				jsii.String("subnets"),
//   			},
//   		},
//   	},
//   	SqsQueueParameters: &PipeSourceSqsQueueParametersProperty{
//   		BatchSize: jsii.Number(123),
//   		MaximumBatchingWindowInSeconds: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourceparameters.html
//
type CfnPipePropsMixin_PipeSourceParametersProperty struct {
	// The parameters for using an Active MQ broker as a source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourceparameters.html#cfn-pipes-pipe-pipesourceparameters-activemqbrokerparameters
	//
	ActiveMqBrokerParameters interface{} `field:"optional" json:"activeMqBrokerParameters" yaml:"activeMqBrokerParameters"`
	// The parameters for using a DynamoDB stream as a source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourceparameters.html#cfn-pipes-pipe-pipesourceparameters-dynamodbstreamparameters
	//
	DynamoDbStreamParameters interface{} `field:"optional" json:"dynamoDbStreamParameters" yaml:"dynamoDbStreamParameters"`
	// The collection of event patterns used to filter events.
	//
	// To remove a filter, specify a `FilterCriteria` object with an empty array of `Filter` objects.
	//
	// For more information, see [Events and Event Patterns](https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-and-event-patterns.html) in the *Amazon EventBridge User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourceparameters.html#cfn-pipes-pipe-pipesourceparameters-filtercriteria
	//
	FilterCriteria interface{} `field:"optional" json:"filterCriteria" yaml:"filterCriteria"`
	// The parameters for using a Kinesis stream as a source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourceparameters.html#cfn-pipes-pipe-pipesourceparameters-kinesisstreamparameters
	//
	KinesisStreamParameters interface{} `field:"optional" json:"kinesisStreamParameters" yaml:"kinesisStreamParameters"`
	// The parameters for using an MSK stream as a source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourceparameters.html#cfn-pipes-pipe-pipesourceparameters-managedstreamingkafkaparameters
	//
	ManagedStreamingKafkaParameters interface{} `field:"optional" json:"managedStreamingKafkaParameters" yaml:"managedStreamingKafkaParameters"`
	// The parameters for using a Rabbit MQ broker as a source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourceparameters.html#cfn-pipes-pipe-pipesourceparameters-rabbitmqbrokerparameters
	//
	RabbitMqBrokerParameters interface{} `field:"optional" json:"rabbitMqBrokerParameters" yaml:"rabbitMqBrokerParameters"`
	// The parameters for using a self-managed Apache Kafka stream as a source.
	//
	// A *self managed* cluster refers to any Apache Kafka cluster not hosted by AWS . This includes both clusters you manage yourself, as well as those hosted by a third-party provider, such as [Confluent Cloud](https://docs.aws.amazon.com/https://www.confluent.io/) , [CloudKarafka](https://docs.aws.amazon.com/https://www.cloudkarafka.com/) , or [Redpanda](https://docs.aws.amazon.com/https://redpanda.com/) . For more information, see [Apache Kafka streams as a source](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes-kafka.html) in the *Amazon EventBridge User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourceparameters.html#cfn-pipes-pipe-pipesourceparameters-selfmanagedkafkaparameters
	//
	SelfManagedKafkaParameters interface{} `field:"optional" json:"selfManagedKafkaParameters" yaml:"selfManagedKafkaParameters"`
	// The parameters for using a Amazon SQS stream as a source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourceparameters.html#cfn-pipes-pipe-pipesourceparameters-sqsqueueparameters
	//
	SqsQueueParameters interface{} `field:"optional" json:"sqsQueueParameters" yaml:"sqsQueueParameters"`
}

