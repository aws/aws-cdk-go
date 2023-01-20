package awspipes


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
	// `CfnPipe.PipeSourceParametersProperty.ActiveMQBrokerParameters`.
	ActiveMqBrokerParameters interface{} `field:"optional" json:"activeMqBrokerParameters" yaml:"activeMqBrokerParameters"`
	// `CfnPipe.PipeSourceParametersProperty.DynamoDBStreamParameters`.
	DynamoDbStreamParameters interface{} `field:"optional" json:"dynamoDbStreamParameters" yaml:"dynamoDbStreamParameters"`
	// `CfnPipe.PipeSourceParametersProperty.FilterCriteria`.
	FilterCriteria interface{} `field:"optional" json:"filterCriteria" yaml:"filterCriteria"`
	// `CfnPipe.PipeSourceParametersProperty.KinesisStreamParameters`.
	KinesisStreamParameters interface{} `field:"optional" json:"kinesisStreamParameters" yaml:"kinesisStreamParameters"`
	// `CfnPipe.PipeSourceParametersProperty.ManagedStreamingKafkaParameters`.
	ManagedStreamingKafkaParameters interface{} `field:"optional" json:"managedStreamingKafkaParameters" yaml:"managedStreamingKafkaParameters"`
	// `CfnPipe.PipeSourceParametersProperty.RabbitMQBrokerParameters`.
	RabbitMqBrokerParameters interface{} `field:"optional" json:"rabbitMqBrokerParameters" yaml:"rabbitMqBrokerParameters"`
	// `CfnPipe.PipeSourceParametersProperty.SelfManagedKafkaParameters`.
	SelfManagedKafkaParameters interface{} `field:"optional" json:"selfManagedKafkaParameters" yaml:"selfManagedKafkaParameters"`
	// `CfnPipe.PipeSourceParametersProperty.SqsQueueParameters`.
	SqsQueueParameters interface{} `field:"optional" json:"sqsQueueParameters" yaml:"sqsQueueParameters"`
}

