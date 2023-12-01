package awspipes

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awspipes/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Create a pipe.
//
// Amazon EventBridge Pipes connect event sources to targets and reduces the need for specialized knowledge and integration code.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPipe := awscdk.Aws_pipes.NewCfnPipe(this, jsii.String("MyCfnPipe"), &CfnPipeProps{
//   	RoleArn: jsii.String("roleArn"),
//   	Source: jsii.String("source"),
//   	Target: jsii.String("target"),
//
//   	// the properties below are optional
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
//   	SourceParameters: &PipeSourceParametersProperty{
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
//   		FilterCriteria: &FilterCriteriaProperty{
//   			Filters: []interface{}{
//   				&FilterProperty{
//   					Pattern: jsii.String("pattern"),
//   				},
//   			},
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
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	TargetParameters: &PipeTargetParametersProperty{
//   		BatchJobParameters: &PipeTargetBatchJobParametersProperty{
//   			JobDefinition: jsii.String("jobDefinition"),
//   			JobName: jsii.String("jobName"),
//
//   			// the properties below are optional
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
//   			TaskDefinitionArn: jsii.String("taskDefinitionArn"),
//
//   			// the properties below are optional
//   			CapacityProviderStrategy: []interface{}{
//   				&CapacityProviderStrategyItemProperty{
//   					CapacityProvider: jsii.String("capacityProvider"),
//
//   					// the properties below are optional
//   					Base: jsii.Number(123),
//   					Weight: jsii.Number(123),
//   				},
//   			},
//   			EnableEcsManagedTags: jsii.Boolean(false),
//   			EnableExecuteCommand: jsii.Boolean(false),
//   			Group: jsii.String("group"),
//   			LaunchType: jsii.String("launchType"),
//   			NetworkConfiguration: &NetworkConfigurationProperty{
//   				AwsvpcConfiguration: &AwsVpcConfigurationProperty{
//   					Subnets: []*string{
//   						jsii.String("subnets"),
//   					},
//
//   					// the properties below are optional
//   					AssignPublicIp: jsii.String("assignPublicIp"),
//   					SecurityGroups: []*string{
//   						jsii.String("securityGroups"),
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
//   			Tags: []cfnTag{
//   				&cfnTag{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			TaskCount: jsii.Number(123),
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
//   			Sqls: []*string{
//   				jsii.String("sqls"),
//   			},
//
//   			// the properties below are optional
//   			DbUser: jsii.String("dbUser"),
//   			SecretManagerArn: jsii.String("secretManagerArn"),
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
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pipes-pipe.html
//
type CfnPipe interface {
	awscdk.CfnResource
	awscdk.IInspectable
	awscdk.ITaggable
	// The ARN of the pipe.
	AttrArn() *string
	// The time the pipe was created.
	AttrCreationTime() *string
	// The state the pipe is in.
	AttrCurrentState() *string
	// When the pipe was last updated, in [ISO-8601 format](https://docs.aws.amazon.com/https://www.w3.org/TR/NOTE-datetime) (YYYY-MM-DDThh:mm:ss.sTZD).
	AttrLastModifiedTime() *string
	// The reason the pipe is in its current state.
	AttrStateReason() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// A description of the pipe.
	Description() *string
	SetDescription(val *string)
	// The state the pipe should be in.
	DesiredState() *string
	SetDesiredState(val *string)
	// The ARN of the enrichment resource.
	Enrichment() *string
	SetEnrichment(val *string)
	// The parameters required to set up enrichment on your pipe.
	EnrichmentParameters() interface{}
	SetEnrichmentParameters(val interface{})
	// The logging configuration settings for the pipe.
	LogConfiguration() interface{}
	SetLogConfiguration(val interface{})
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The name of the pipe.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The ARN of the role that allows the pipe to send data to the target.
	RoleArn() *string
	SetRoleArn(val *string)
	// The ARN of the source resource.
	Source() *string
	SetSource(val *string)
	// The parameters required to set up a source for your pipe.
	SourceParameters() interface{}
	SetSourceParameters(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Tag Manager which manages the tags for this resource.
	Tags() awscdk.TagManager
	// The list of key-value pairs to associate with the pipe.
	TagsRaw() *map[string]*string
	SetTagsRaw(val *map[string]*string)
	// The ARN of the target resource.
	Target() *string
	SetTarget(val *string)
	// The parameters required to set up a target for your pipe.
	TargetParameters() interface{}
	SetTargetParameters(val interface{})
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependency(target awscdk.CfnResource)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	// Deprecated: use addDependency.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//   "GlobalSecondaryIndexes": [
	//     {
	//       "Projection": {
	//         "NonKeyAttributes": [ "myattribute" ]
	//         ...
	//       }
	//       ...
	//     },
	//     {
	//       "ProjectionType": "INCLUDE"
	//       ...
	//     },
	//   ]
	//   ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Retrieves an array of resources this resource depends on.
	//
	// This assembles dependencies on resources across stacks (including nested stacks)
	// automatically.
	ObtainDependencies() *[]interface{}
	// Get a shallow copy of dependencies between this resource and other resources in the same stack.
	ObtainResourceDependencies() *[]awscdk.CfnResource
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	// Indicates that this resource no longer depends on another resource.
	//
	// This can be used for resources across stacks (including nested stacks)
	// and the dependency will automatically be removed from the relevant scope.
	RemoveDependency(target awscdk.CfnResource)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Replaces one dependency with another.
	ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource)
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnPipe
type jsiiProxy_CfnPipe struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__awscdkITaggable
}

func (j *jsiiProxy_CfnPipe) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipe) AttrCreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipe) AttrCurrentState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCurrentState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipe) AttrLastModifiedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLastModifiedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipe) AttrStateReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStateReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipe) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipe) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipe) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipe) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipe) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipe) DesiredState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desiredState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipe) Enrichment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enrichment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipe) EnrichmentParameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enrichmentParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipe) LogConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipe) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipe) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipe) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipe) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipe) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipe) Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipe) SourceParameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipe) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipe) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipe) TagsRaw() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipe) Target() *string {
	var returns *string
	_jsii_.Get(
		j,
		"target",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipe) TargetParameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipe) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipe) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


func NewCfnPipe(scope constructs.Construct, id *string, props *CfnPipeProps) CfnPipe {
	_init_.Initialize()

	if err := validateNewCfnPipeParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnPipe{}

	_jsii_.Create(
		"aws-cdk-lib.aws_pipes.CfnPipe",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnPipe_Override(c CfnPipe, scope constructs.Construct, id *string, props *CfnPipeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_pipes.CfnPipe",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnPipe)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnPipe)SetDesiredState(val *string) {
	_jsii_.Set(
		j,
		"desiredState",
		val,
	)
}

func (j *jsiiProxy_CfnPipe)SetEnrichment(val *string) {
	_jsii_.Set(
		j,
		"enrichment",
		val,
	)
}

func (j *jsiiProxy_CfnPipe)SetEnrichmentParameters(val interface{}) {
	if err := j.validateSetEnrichmentParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enrichmentParameters",
		val,
	)
}

func (j *jsiiProxy_CfnPipe)SetLogConfiguration(val interface{}) {
	if err := j.validateSetLogConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnPipe)SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnPipe)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_CfnPipe)SetSource(val *string) {
	if err := j.validateSetSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"source",
		val,
	)
}

func (j *jsiiProxy_CfnPipe)SetSourceParameters(val interface{}) {
	if err := j.validateSetSourceParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceParameters",
		val,
	)
}

func (j *jsiiProxy_CfnPipe)SetTagsRaw(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tagsRaw",
		val,
	)
}

func (j *jsiiProxy_CfnPipe)SetTarget(val *string) {
	if err := j.validateSetTargetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"target",
		val,
	)
}

func (j *jsiiProxy_CfnPipe)SetTargetParameters(val interface{}) {
	if err := j.validateSetTargetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetParameters",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnPipe_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPipe_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pipes.CfnPipe",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnPipe_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPipe_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pipes.CfnPipe",
		"isCfnResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnPipe_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPipe_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pipes.CfnPipe",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPipe_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_pipes.CfnPipe",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPipe) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnPipe) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnPipe) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnPipe) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnPipe) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnPipe) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnPipe) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnPipe) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnPipe) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName, typeHint},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPipe) GetMetadata(key *string) interface{} {
	if err := c.validateGetMetadataParameters(key); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPipe) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnPipe) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPipe) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPipe) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnPipe) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnPipe) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	if err := c.validateRenderPropertiesParameters(props); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPipe) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnPipe) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPipe) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPipe) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

