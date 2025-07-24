package awscdkpipesalpha


// Source properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import pipes_alpha "github.com/aws/aws-cdk-go/awscdkpipesalpha"
//
//   sourceConfig := &SourceConfig{
//   	SourceParameters: &SourceParameters{
//   		ActiveMqBrokerParameters: &PipeSourceActiveMQBrokerParametersProperty{
//   			Credentials: &MQBrokerAccessCredentialsProperty{
//   				BasicAuth: jsii.String("basicAuth"),
//   			},
//   			QueueName: jsii.String("queueName"),
//
//   			// the properties below are optional
//   			BatchSize: jsii.Number(123),
//   			MaximumBatchingWindowInSeconds: jsii.Number(123),
//   		},
//   		DynamoDbStreamParameters: &PipeSourceDynamoDBStreamParametersProperty{
//   			StartingPosition: jsii.String("startingPosition"),
//
//   			// the properties below are optional
//   			BatchSize: jsii.Number(123),
//   			DeadLetterConfig: &DeadLetterConfigProperty{
//   				Arn: jsii.String("arn"),
//   			},
//   			MaximumBatchingWindowInSeconds: jsii.Number(123),
//   			MaximumRecordAgeInSeconds: jsii.Number(123),
//   			MaximumRetryAttempts: jsii.Number(123),
//   			OnPartialBatchItemFailure: jsii.String("onPartialBatchItemFailure"),
//   			ParallelizationFactor: jsii.Number(123),
//   		},
//   		KinesisStreamParameters: &PipeSourceKinesisStreamParametersProperty{
//   			StartingPosition: jsii.String("startingPosition"),
//
//   			// the properties below are optional
//   			BatchSize: jsii.Number(123),
//   			DeadLetterConfig: &DeadLetterConfigProperty{
//   				Arn: jsii.String("arn"),
//   			},
//   			MaximumBatchingWindowInSeconds: jsii.Number(123),
//   			MaximumRecordAgeInSeconds: jsii.Number(123),
//   			MaximumRetryAttempts: jsii.Number(123),
//   			OnPartialBatchItemFailure: jsii.String("onPartialBatchItemFailure"),
//   			ParallelizationFactor: jsii.Number(123),
//   			StartingPositionTimestamp: jsii.String("startingPositionTimestamp"),
//   		},
//   		ManagedStreamingKafkaParameters: &PipeSourceManagedStreamingKafkaParametersProperty{
//   			TopicName: jsii.String("topicName"),
//
//   			// the properties below are optional
//   			BatchSize: jsii.Number(123),
//   			ConsumerGroupId: jsii.String("consumerGroupId"),
//   			Credentials: &MSKAccessCredentialsProperty{
//   				ClientCertificateTlsAuth: jsii.String("clientCertificateTlsAuth"),
//   				SaslScram512Auth: jsii.String("saslScram512Auth"),
//   			},
//   			MaximumBatchingWindowInSeconds: jsii.Number(123),
//   			StartingPosition: jsii.String("startingPosition"),
//   		},
//   		RabbitMqBrokerParameters: &PipeSourceRabbitMQBrokerParametersProperty{
//   			Credentials: &MQBrokerAccessCredentialsProperty{
//   				BasicAuth: jsii.String("basicAuth"),
//   			},
//   			QueueName: jsii.String("queueName"),
//
//   			// the properties below are optional
//   			BatchSize: jsii.Number(123),
//   			MaximumBatchingWindowInSeconds: jsii.Number(123),
//   			VirtualHost: jsii.String("virtualHost"),
//   		},
//   		SelfManagedKafkaParameters: &PipeSourceSelfManagedKafkaParametersProperty{
//   			TopicName: jsii.String("topicName"),
//
//   			// the properties below are optional
//   			AdditionalBootstrapServers: []*string{
//   				jsii.String("additionalBootstrapServers"),
//   			},
//   			BatchSize: jsii.Number(123),
//   			ConsumerGroupId: jsii.String("consumerGroupId"),
//   			Credentials: &SelfManagedKafkaAccessConfigurationCredentialsProperty{
//   				BasicAuth: jsii.String("basicAuth"),
//   				ClientCertificateTlsAuth: jsii.String("clientCertificateTlsAuth"),
//   				SaslScram256Auth: jsii.String("saslScram256Auth"),
//   				SaslScram512Auth: jsii.String("saslScram512Auth"),
//   			},
//   			MaximumBatchingWindowInSeconds: jsii.Number(123),
//   			ServerRootCaCertificate: jsii.String("serverRootCaCertificate"),
//   			StartingPosition: jsii.String("startingPosition"),
//   			Vpc: &SelfManagedKafkaAccessConfigurationVpcProperty{
//   				SecurityGroup: []*string{
//   					jsii.String("securityGroup"),
//   				},
//   				Subnets: []*string{
//   					jsii.String("subnets"),
//   				},
//   			},
//   		},
//   		SqsQueueParameters: &PipeSourceSqsQueueParametersProperty{
//   			BatchSize: jsii.Number(123),
//   			MaximumBatchingWindowInSeconds: jsii.Number(123),
//   		},
//   	},
//   }
//
// Experimental.
type SourceConfig struct {
	// The parameters required to set up a source for your pipe.
	// Default: - none.
	//
	// Experimental.
	SourceParameters *SourceParameters `field:"optional" json:"sourceParameters" yaml:"sourceParameters"`
}

