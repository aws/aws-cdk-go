package awspipes


// The parameters required to set up a source for your pipe.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pipeSourceParametersProperty := &pipeSourceParametersProperty{
//   	activeMqBrokerParameters: &pipeSourceActiveMQBrokerParametersProperty{
//   		credentials: &mQBrokerAccessCredentialsProperty{
//   			basicAuth: jsii.String("basicAuth"),
//   		},
//   		queueName: jsii.String("queueName"),
//
//   		// the properties below are optional
//   		batchSize: jsii.Number(123),
//   		maximumBatchingWindowInSeconds: jsii.Number(123),
//   	},
//   	dynamoDbStreamParameters: &pipeSourceDynamoDBStreamParametersProperty{
//   		startingPosition: jsii.String("startingPosition"),
//
//   		// the properties below are optional
//   		batchSize: jsii.Number(123),
//   		deadLetterConfig: &deadLetterConfigProperty{
//   			arn: jsii.String("arn"),
//   		},
//   		maximumBatchingWindowInSeconds: jsii.Number(123),
//   		maximumRecordAgeInSeconds: jsii.Number(123),
//   		maximumRetryAttempts: jsii.Number(123),
//   		onPartialBatchItemFailure: jsii.String("onPartialBatchItemFailure"),
//   		parallelizationFactor: jsii.Number(123),
//   	},
//   	filterCriteria: &filterCriteriaProperty{
//   		filters: []interface{}{
//   			&filterProperty{
//   				pattern: jsii.String("pattern"),
//   			},
//   		},
//   	},
//   	kinesisStreamParameters: &pipeSourceKinesisStreamParametersProperty{
//   		startingPosition: jsii.String("startingPosition"),
//
//   		// the properties below are optional
//   		batchSize: jsii.Number(123),
//   		deadLetterConfig: &deadLetterConfigProperty{
//   			arn: jsii.String("arn"),
//   		},
//   		maximumBatchingWindowInSeconds: jsii.Number(123),
//   		maximumRecordAgeInSeconds: jsii.Number(123),
//   		maximumRetryAttempts: jsii.Number(123),
//   		onPartialBatchItemFailure: jsii.String("onPartialBatchItemFailure"),
//   		parallelizationFactor: jsii.Number(123),
//   		startingPositionTimestamp: jsii.String("startingPositionTimestamp"),
//   	},
//   	managedStreamingKafkaParameters: &pipeSourceManagedStreamingKafkaParametersProperty{
//   		topicName: jsii.String("topicName"),
//
//   		// the properties below are optional
//   		batchSize: jsii.Number(123),
//   		consumerGroupId: jsii.String("consumerGroupId"),
//   		credentials: &mSKAccessCredentialsProperty{
//   			clientCertificateTlsAuth: jsii.String("clientCertificateTlsAuth"),
//   			saslScram512Auth: jsii.String("saslScram512Auth"),
//   		},
//   		maximumBatchingWindowInSeconds: jsii.Number(123),
//   		startingPosition: jsii.String("startingPosition"),
//   	},
//   	rabbitMqBrokerParameters: &pipeSourceRabbitMQBrokerParametersProperty{
//   		credentials: &mQBrokerAccessCredentialsProperty{
//   			basicAuth: jsii.String("basicAuth"),
//   		},
//   		queueName: jsii.String("queueName"),
//
//   		// the properties below are optional
//   		batchSize: jsii.Number(123),
//   		maximumBatchingWindowInSeconds: jsii.Number(123),
//   		virtualHost: jsii.String("virtualHost"),
//   	},
//   	selfManagedKafkaParameters: &pipeSourceSelfManagedKafkaParametersProperty{
//   		topicName: jsii.String("topicName"),
//
//   		// the properties below are optional
//   		additionalBootstrapServers: []*string{
//   			jsii.String("additionalBootstrapServers"),
//   		},
//   		batchSize: jsii.Number(123),
//   		consumerGroupId: jsii.String("consumerGroupId"),
//   		credentials: &selfManagedKafkaAccessConfigurationCredentialsProperty{
//   			basicAuth: jsii.String("basicAuth"),
//   			clientCertificateTlsAuth: jsii.String("clientCertificateTlsAuth"),
//   			saslScram256Auth: jsii.String("saslScram256Auth"),
//   			saslScram512Auth: jsii.String("saslScram512Auth"),
//   		},
//   		maximumBatchingWindowInSeconds: jsii.Number(123),
//   		serverRootCaCertificate: jsii.String("serverRootCaCertificate"),
//   		startingPosition: jsii.String("startingPosition"),
//   		vpc: &selfManagedKafkaAccessConfigurationVpcProperty{
//   			securityGroup: []*string{
//   				jsii.String("securityGroup"),
//   			},
//   			subnets: []*string{
//   				jsii.String("subnets"),
//   			},
//   		},
//   	},
//   	sqsQueueParameters: &pipeSourceSqsQueueParametersProperty{
//   		batchSize: jsii.Number(123),
//   		maximumBatchingWindowInSeconds: jsii.Number(123),
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

