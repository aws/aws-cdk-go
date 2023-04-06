package awspipes


// The parameters required to set up a source for your pipe.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pipeSourceParametersProperty := &PipeSourceParametersProperty{
//   	ActiveMqBrokerParameters: &PipeSourceActiveMQBrokerParametersProperty{
//   		Credentials: &MQBrokerAccessCredentialsProperty{
//   			BasicAuth: jsii.String("basicAuth"),
//   		},
//   		QueueName: jsii.String("queueName"),
//
//   		// the properties below are optional
//   		BatchSize: jsii.Number(123),
//   		MaximumBatchingWindowInSeconds: jsii.Number(123),
//   	},
//   	DynamoDbStreamParameters: &PipeSourceDynamoDBStreamParametersProperty{
//   		StartingPosition: jsii.String("startingPosition"),
//
//   		// the properties below are optional
//   		BatchSize: jsii.Number(123),
//   		DeadLetterConfig: &DeadLetterConfigProperty{
//   			Arn: jsii.String("arn"),
//   		},
//   		MaximumBatchingWindowInSeconds: jsii.Number(123),
//   		MaximumRecordAgeInSeconds: jsii.Number(123),
//   		MaximumRetryAttempts: jsii.Number(123),
//   		OnPartialBatchItemFailure: jsii.String("onPartialBatchItemFailure"),
//   		ParallelizationFactor: jsii.Number(123),
//   	},
//   	FilterCriteria: &FilterCriteriaProperty{
//   		Filters: []interface{}{
//   			&FilterProperty{
//   				Pattern: jsii.String("pattern"),
//   			},
//   		},
//   	},
//   	KinesisStreamParameters: &PipeSourceKinesisStreamParametersProperty{
//   		StartingPosition: jsii.String("startingPosition"),
//
//   		// the properties below are optional
//   		BatchSize: jsii.Number(123),
//   		DeadLetterConfig: &DeadLetterConfigProperty{
//   			Arn: jsii.String("arn"),
//   		},
//   		MaximumBatchingWindowInSeconds: jsii.Number(123),
//   		MaximumRecordAgeInSeconds: jsii.Number(123),
//   		MaximumRetryAttempts: jsii.Number(123),
//   		OnPartialBatchItemFailure: jsii.String("onPartialBatchItemFailure"),
//   		ParallelizationFactor: jsii.Number(123),
//   		StartingPositionTimestamp: jsii.String("startingPositionTimestamp"),
//   	},
//   	ManagedStreamingKafkaParameters: &PipeSourceManagedStreamingKafkaParametersProperty{
//   		TopicName: jsii.String("topicName"),
//
//   		// the properties below are optional
//   		BatchSize: jsii.Number(123),
//   		ConsumerGroupId: jsii.String("consumerGroupId"),
//   		Credentials: &MSKAccessCredentialsProperty{
//   			ClientCertificateTlsAuth: jsii.String("clientCertificateTlsAuth"),
//   			SaslScram512Auth: jsii.String("saslScram512Auth"),
//   		},
//   		MaximumBatchingWindowInSeconds: jsii.Number(123),
//   		StartingPosition: jsii.String("startingPosition"),
//   	},
//   	RabbitMqBrokerParameters: &PipeSourceRabbitMQBrokerParametersProperty{
//   		Credentials: &MQBrokerAccessCredentialsProperty{
//   			BasicAuth: jsii.String("basicAuth"),
//   		},
//   		QueueName: jsii.String("queueName"),
//
//   		// the properties below are optional
//   		BatchSize: jsii.Number(123),
//   		MaximumBatchingWindowInSeconds: jsii.Number(123),
//   		VirtualHost: jsii.String("virtualHost"),
//   	},
//   	SelfManagedKafkaParameters: &PipeSourceSelfManagedKafkaParametersProperty{
//   		TopicName: jsii.String("topicName"),
//
//   		// the properties below are optional
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
type CfnPipe_PipeSourceParametersProperty struct {
	// The parameters for using an Active MQ broker as a source.
	ActiveMqBrokerParameters interface{} `field:"optional" json:"activeMqBrokerParameters" yaml:"activeMqBrokerParameters"`
	// The parameters for using a DynamoDB stream as a source.
	DynamoDbStreamParameters interface{} `field:"optional" json:"dynamoDbStreamParameters" yaml:"dynamoDbStreamParameters"`
	// The collection of event patterns used to filter events.
	//
	// To remove a filter, specify a `FilterCriteria` object with an empty array of `Filter` objects.
	//
	// For more information, see [Events and Event Patterns](https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-and-event-patterns.html) in the *Amazon EventBridge User Guide* .
	FilterCriteria interface{} `field:"optional" json:"filterCriteria" yaml:"filterCriteria"`
	// The parameters for using a Kinesis stream as a source.
	KinesisStreamParameters interface{} `field:"optional" json:"kinesisStreamParameters" yaml:"kinesisStreamParameters"`
	// The parameters for using an MSK stream as a source.
	ManagedStreamingKafkaParameters interface{} `field:"optional" json:"managedStreamingKafkaParameters" yaml:"managedStreamingKafkaParameters"`
	// The parameters for using a Rabbit MQ broker as a source.
	RabbitMqBrokerParameters interface{} `field:"optional" json:"rabbitMqBrokerParameters" yaml:"rabbitMqBrokerParameters"`
	// The parameters for using a self-managed Apache Kafka stream as a source.
	SelfManagedKafkaParameters interface{} `field:"optional" json:"selfManagedKafkaParameters" yaml:"selfManagedKafkaParameters"`
	// The parameters for using a Amazon SQS stream as a source.
	SqsQueueParameters interface{} `field:"optional" json:"sqsQueueParameters" yaml:"sqsQueueParameters"`
}

