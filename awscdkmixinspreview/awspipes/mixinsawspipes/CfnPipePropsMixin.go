package mixinsawspipes

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awspipes/mixinsawspipes/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies a pipe.
//
// Amazon EventBridge Pipes connect event sources to targets and reduces the need for specialized knowledge and integration code.
//
// > As an aid to help you jumpstart developing CloudFormation templates, the EventBridge console enables you to create templates from the existing pipes in your account. For more information, see [Generate an CloudFormation template from EventBridge Pipes](https://docs.aws.amazon.com/eventbridge/latest/userguide/pipes-generate-template.html) in the *Amazon EventBridge User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnPipePropsMixin := awscdkmixinspreview.Mixins.NewCfnPipePropsMixin(&CfnPipeMixinProps{
//   	Description: jsii.String("description"),
//   	DesiredState: jsii.String("desiredState"),
//   	Enrichment: jsii.String("enrichment"),
//   	EnrichmentParameters: &PipeEnrichmentParametersProperty{
//   		HttpParameters: &PipeEnrichmentHttpParametersProperty{
//   			HeaderParameters: map[string]*string{
//   				"headerParametersKey": jsii.String("headerParameters"),
//   			},
//   			PathParameterValues: []*string{
//   				jsii.String("pathParameterValues"),
//   			},
//   			QueryStringParameters: map[string]*string{
//   				"queryStringParametersKey": jsii.String("queryStringParameters"),
//   			},
//   		},
//   		InputTemplate: jsii.String("inputTemplate"),
//   	},
//   	KmsKeyIdentifier: jsii.String("kmsKeyIdentifier"),
//   	LogConfiguration: &PipeLogConfigurationProperty{
//   		CloudwatchLogsLogDestination: &CloudwatchLogsLogDestinationProperty{
//   			LogGroupArn: jsii.String("logGroupArn"),
//   		},
//   		FirehoseLogDestination: &FirehoseLogDestinationProperty{
//   			DeliveryStreamArn: jsii.String("deliveryStreamArn"),
//   		},
//   		IncludeExecutionData: []*string{
//   			jsii.String("includeExecutionData"),
//   		},
//   		Level: jsii.String("level"),
//   		S3LogDestination: &S3LogDestinationProperty{
//   			BucketName: jsii.String("bucketName"),
//   			BucketOwner: jsii.String("bucketOwner"),
//   			OutputFormat: jsii.String("outputFormat"),
//   			Prefix: jsii.String("prefix"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	RoleArn: jsii.String("roleArn"),
//   	Source: jsii.String("source"),
//   	SourceParameters: &PipeSourceParametersProperty{
//   		ActiveMqBrokerParameters: &PipeSourceActiveMQBrokerParametersProperty{
//   			BatchSize: jsii.Number(123),
//   			Credentials: &MQBrokerAccessCredentialsProperty{
//   				BasicAuth: jsii.String("basicAuth"),
//   			},
//   			MaximumBatchingWindowInSeconds: jsii.Number(123),
//   			QueueName: jsii.String("queueName"),
//   		},
//   		DynamoDbStreamParameters: &PipeSourceDynamoDBStreamParametersProperty{
//   			BatchSize: jsii.Number(123),
//   			DeadLetterConfig: &DeadLetterConfigProperty{
//   				Arn: jsii.String("arn"),
//   			},
//   			MaximumBatchingWindowInSeconds: jsii.Number(123),
//   			MaximumRecordAgeInSeconds: jsii.Number(123),
//   			MaximumRetryAttempts: jsii.Number(123),
//   			OnPartialBatchItemFailure: jsii.String("onPartialBatchItemFailure"),
//   			ParallelizationFactor: jsii.Number(123),
//   			StartingPosition: jsii.String("startingPosition"),
//   		},
//   		FilterCriteria: &FilterCriteriaProperty{
//   			Filters: []interface{}{
//   				&FilterProperty{
//   					Pattern: jsii.String("pattern"),
//   				},
//   			},
//   		},
//   		KinesisStreamParameters: &PipeSourceKinesisStreamParametersProperty{
//   			BatchSize: jsii.Number(123),
//   			DeadLetterConfig: &DeadLetterConfigProperty{
//   				Arn: jsii.String("arn"),
//   			},
//   			MaximumBatchingWindowInSeconds: jsii.Number(123),
//   			MaximumRecordAgeInSeconds: jsii.Number(123),
//   			MaximumRetryAttempts: jsii.Number(123),
//   			OnPartialBatchItemFailure: jsii.String("onPartialBatchItemFailure"),
//   			ParallelizationFactor: jsii.Number(123),
//   			StartingPosition: jsii.String("startingPosition"),
//   			StartingPositionTimestamp: jsii.String("startingPositionTimestamp"),
//   		},
//   		ManagedStreamingKafkaParameters: &PipeSourceManagedStreamingKafkaParametersProperty{
//   			BatchSize: jsii.Number(123),
//   			ConsumerGroupId: jsii.String("consumerGroupId"),
//   			Credentials: &MSKAccessCredentialsProperty{
//   				ClientCertificateTlsAuth: jsii.String("clientCertificateTlsAuth"),
//   				SaslScram512Auth: jsii.String("saslScram512Auth"),
//   			},
//   			MaximumBatchingWindowInSeconds: jsii.Number(123),
//   			StartingPosition: jsii.String("startingPosition"),
//   			TopicName: jsii.String("topicName"),
//   		},
//   		RabbitMqBrokerParameters: &PipeSourceRabbitMQBrokerParametersProperty{
//   			BatchSize: jsii.Number(123),
//   			Credentials: &MQBrokerAccessCredentialsProperty{
//   				BasicAuth: jsii.String("basicAuth"),
//   			},
//   			MaximumBatchingWindowInSeconds: jsii.Number(123),
//   			QueueName: jsii.String("queueName"),
//   			VirtualHost: jsii.String("virtualHost"),
//   		},
//   		SelfManagedKafkaParameters: &PipeSourceSelfManagedKafkaParametersProperty{
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
//   			TopicName: jsii.String("topicName"),
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
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	Target: jsii.String("target"),
//   	TargetParameters: &PipeTargetParametersProperty{
//   		BatchJobParameters: &PipeTargetBatchJobParametersProperty{
//   			ArrayProperties: &BatchArrayPropertiesProperty{
//   				Size: jsii.Number(123),
//   			},
//   			ContainerOverrides: &BatchContainerOverridesProperty{
//   				Command: []*string{
//   					jsii.String("command"),
//   				},
//   				Environment: []interface{}{
//   					&BatchEnvironmentVariableProperty{
//   						Name: jsii.String("name"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   				InstanceType: jsii.String("instanceType"),
//   				ResourceRequirements: []interface{}{
//   					&BatchResourceRequirementProperty{
//   						Type: jsii.String("type"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   			},
//   			DependsOn: []interface{}{
//   				&BatchJobDependencyProperty{
//   					JobId: jsii.String("jobId"),
//   					Type: jsii.String("type"),
//   				},
//   			},
//   			JobDefinition: jsii.String("jobDefinition"),
//   			JobName: jsii.String("jobName"),
//   			Parameters: map[string]*string{
//   				"parametersKey": jsii.String("parameters"),
//   			},
//   			RetryStrategy: &BatchRetryStrategyProperty{
//   				Attempts: jsii.Number(123),
//   			},
//   		},
//   		CloudWatchLogsParameters: &PipeTargetCloudWatchLogsParametersProperty{
//   			LogStreamName: jsii.String("logStreamName"),
//   			Timestamp: jsii.String("timestamp"),
//   		},
//   		EcsTaskParameters: &PipeTargetEcsTaskParametersProperty{
//   			CapacityProviderStrategy: []interface{}{
//   				&CapacityProviderStrategyItemProperty{
//   					Base: jsii.Number(123),
//   					CapacityProvider: jsii.String("capacityProvider"),
//   					Weight: jsii.Number(123),
//   				},
//   			},
//   			EnableEcsManagedTags: jsii.Boolean(false),
//   			EnableExecuteCommand: jsii.Boolean(false),
//   			Group: jsii.String("group"),
//   			LaunchType: jsii.String("launchType"),
//   			NetworkConfiguration: &NetworkConfigurationProperty{
//   				AwsvpcConfiguration: &AwsVpcConfigurationProperty{
//   					AssignPublicIp: jsii.String("assignPublicIp"),
//   					SecurityGroups: []*string{
//   						jsii.String("securityGroups"),
//   					},
//   					Subnets: []*string{
//   						jsii.String("subnets"),
//   					},
//   				},
//   			},
//   			Overrides: &EcsTaskOverrideProperty{
//   				ContainerOverrides: []interface{}{
//   					&EcsContainerOverrideProperty{
//   						Command: []*string{
//   							jsii.String("command"),
//   						},
//   						Cpu: jsii.Number(123),
//   						Environment: []interface{}{
//   							&EcsEnvironmentVariableProperty{
//   								Name: jsii.String("name"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   						EnvironmentFiles: []interface{}{
//   							&EcsEnvironmentFileProperty{
//   								Type: jsii.String("type"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   						Memory: jsii.Number(123),
//   						MemoryReservation: jsii.Number(123),
//   						Name: jsii.String("name"),
//   						ResourceRequirements: []interface{}{
//   							&EcsResourceRequirementProperty{
//   								Type: jsii.String("type"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   				Cpu: jsii.String("cpu"),
//   				EphemeralStorage: &EcsEphemeralStorageProperty{
//   					SizeInGiB: jsii.Number(123),
//   				},
//   				ExecutionRoleArn: jsii.String("executionRoleArn"),
//   				InferenceAcceleratorOverrides: []interface{}{
//   					&EcsInferenceAcceleratorOverrideProperty{
//   						DeviceName: jsii.String("deviceName"),
//   						DeviceType: jsii.String("deviceType"),
//   					},
//   				},
//   				Memory: jsii.String("memory"),
//   				TaskRoleArn: jsii.String("taskRoleArn"),
//   			},
//   			PlacementConstraints: []interface{}{
//   				&PlacementConstraintProperty{
//   					Expression: jsii.String("expression"),
//   					Type: jsii.String("type"),
//   				},
//   			},
//   			PlacementStrategy: []interface{}{
//   				&PlacementStrategyProperty{
//   					Field: jsii.String("field"),
//   					Type: jsii.String("type"),
//   				},
//   			},
//   			PlatformVersion: jsii.String("platformVersion"),
//   			PropagateTags: jsii.String("propagateTags"),
//   			ReferenceId: jsii.String("referenceId"),
//   			Tags: []CfnTag{
//   				&CfnTag{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			TaskCount: jsii.Number(123),
//   			TaskDefinitionArn: jsii.String("taskDefinitionArn"),
//   		},
//   		EventBridgeEventBusParameters: &PipeTargetEventBridgeEventBusParametersProperty{
//   			DetailType: jsii.String("detailType"),
//   			EndpointId: jsii.String("endpointId"),
//   			Resources: []*string{
//   				jsii.String("resources"),
//   			},
//   			Source: jsii.String("source"),
//   			Time: jsii.String("time"),
//   		},
//   		HttpParameters: &PipeTargetHttpParametersProperty{
//   			HeaderParameters: map[string]*string{
//   				"headerParametersKey": jsii.String("headerParameters"),
//   			},
//   			PathParameterValues: []*string{
//   				jsii.String("pathParameterValues"),
//   			},
//   			QueryStringParameters: map[string]*string{
//   				"queryStringParametersKey": jsii.String("queryStringParameters"),
//   			},
//   		},
//   		InputTemplate: jsii.String("inputTemplate"),
//   		KinesisStreamParameters: &PipeTargetKinesisStreamParametersProperty{
//   			PartitionKey: jsii.String("partitionKey"),
//   		},
//   		LambdaFunctionParameters: &PipeTargetLambdaFunctionParametersProperty{
//   			InvocationType: jsii.String("invocationType"),
//   		},
//   		RedshiftDataParameters: &PipeTargetRedshiftDataParametersProperty{
//   			Database: jsii.String("database"),
//   			DbUser: jsii.String("dbUser"),
//   			SecretManagerArn: jsii.String("secretManagerArn"),
//   			Sqls: []*string{
//   				jsii.String("sqls"),
//   			},
//   			StatementName: jsii.String("statementName"),
//   			WithEvent: jsii.Boolean(false),
//   		},
//   		SageMakerPipelineParameters: &PipeTargetSageMakerPipelineParametersProperty{
//   			PipelineParameterList: []interface{}{
//   				&SageMakerPipelineParameterProperty{
//   					Name: jsii.String("name"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   		SqsQueueParameters: &PipeTargetSqsQueueParametersProperty{
//   			MessageDeduplicationId: jsii.String("messageDeduplicationId"),
//   			MessageGroupId: jsii.String("messageGroupId"),
//   		},
//   		StepFunctionStateMachineParameters: &PipeTargetStateMachineParametersProperty{
//   			InvocationType: jsii.String("invocationType"),
//   		},
//   		TimestreamParameters: &PipeTargetTimestreamParametersProperty{
//   			DimensionMappings: []interface{}{
//   				&DimensionMappingProperty{
//   					DimensionName: jsii.String("dimensionName"),
//   					DimensionValue: jsii.String("dimensionValue"),
//   					DimensionValueType: jsii.String("dimensionValueType"),
//   				},
//   			},
//   			EpochTimeUnit: jsii.String("epochTimeUnit"),
//   			MultiMeasureMappings: []interface{}{
//   				&MultiMeasureMappingProperty{
//   					MultiMeasureAttributeMappings: []interface{}{
//   						&MultiMeasureAttributeMappingProperty{
//   							MeasureValue: jsii.String("measureValue"),
//   							MeasureValueType: jsii.String("measureValueType"),
//   							MultiMeasureAttributeName: jsii.String("multiMeasureAttributeName"),
//   						},
//   					},
//   					MultiMeasureName: jsii.String("multiMeasureName"),
//   				},
//   			},
//   			SingleMeasureMappings: []interface{}{
//   				&SingleMeasureMappingProperty{
//   					MeasureName: jsii.String("measureName"),
//   					MeasureValue: jsii.String("measureValue"),
//   					MeasureValueType: jsii.String("measureValueType"),
//   				},
//   			},
//   			TimeFieldType: jsii.String("timeFieldType"),
//   			TimestampFormat: jsii.String("timestampFormat"),
//   			TimeValue: jsii.String("timeValue"),
//   			VersionValue: jsii.String("versionValue"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pipes-pipe.html
//
type CfnPipePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnPipeMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnPipePropsMixin
type jsiiProxy_CfnPipePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnPipePropsMixin) Props() *CfnPipeMixinProps {
	var returns *CfnPipeMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Pipes::Pipe`.
func NewCfnPipePropsMixin(props *CfnPipeMixinProps, options *mixins.CfnPropertyMixinOptions) CfnPipePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnPipePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnPipePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_pipes.mixins.CfnPipePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Pipes::Pipe`.
func NewCfnPipePropsMixin_Override(c CfnPipePropsMixin, props *CfnPipeMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_pipes.mixins.CfnPipePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnPipePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPipePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_pipes.mixins.CfnPipePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPipePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_pipes.mixins.CfnPipePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPipePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"applyTo",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPipePropsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

