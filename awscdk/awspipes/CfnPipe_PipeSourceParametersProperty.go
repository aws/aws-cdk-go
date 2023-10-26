package awspipes


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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourceparameters.html
//
type CfnPipe_PipeSourceParametersProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourceparameters.html#cfn-pipes-pipe-pipesourceparameters-activemqbrokerparameters
	//
	ActiveMqBrokerParameters interface{} `field:"optional" json:"activeMqBrokerParameters" yaml:"activeMqBrokerParameters"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourceparameters.html#cfn-pipes-pipe-pipesourceparameters-dynamodbstreamparameters
	//
	DynamoDbStreamParameters interface{} `field:"optional" json:"dynamoDbStreamParameters" yaml:"dynamoDbStreamParameters"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourceparameters.html#cfn-pipes-pipe-pipesourceparameters-filtercriteria
	//
	FilterCriteria interface{} `field:"optional" json:"filterCriteria" yaml:"filterCriteria"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourceparameters.html#cfn-pipes-pipe-pipesourceparameters-kinesisstreamparameters
	//
	KinesisStreamParameters interface{} `field:"optional" json:"kinesisStreamParameters" yaml:"kinesisStreamParameters"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourceparameters.html#cfn-pipes-pipe-pipesourceparameters-managedstreamingkafkaparameters
	//
	ManagedStreamingKafkaParameters interface{} `field:"optional" json:"managedStreamingKafkaParameters" yaml:"managedStreamingKafkaParameters"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourceparameters.html#cfn-pipes-pipe-pipesourceparameters-rabbitmqbrokerparameters
	//
	RabbitMqBrokerParameters interface{} `field:"optional" json:"rabbitMqBrokerParameters" yaml:"rabbitMqBrokerParameters"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourceparameters.html#cfn-pipes-pipe-pipesourceparameters-selfmanagedkafkaparameters
	//
	SelfManagedKafkaParameters interface{} `field:"optional" json:"selfManagedKafkaParameters" yaml:"selfManagedKafkaParameters"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourceparameters.html#cfn-pipes-pipe-pipesourceparameters-sqsqueueparameters
	//
	SqsQueueParameters interface{} `field:"optional" json:"sqsQueueParameters" yaml:"sqsQueueParameters"`
}

