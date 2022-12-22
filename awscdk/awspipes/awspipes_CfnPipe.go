package awspipes

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awspipes/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::Pipes::Pipe`.
//
// Create a pipe. Amazon EventBridge Pipes connect event sources to targets and reduces the need for specialized knowledge and integration code.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPipe := awscdk.Aws_pipes.NewCfnPipe(this, jsii.String("MyCfnPipe"), &cfnPipeProps{
//   	roleArn: jsii.String("roleArn"),
//   	source: jsii.String("source"),
//   	target: jsii.String("target"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	desiredState: jsii.String("desiredState"),
//   	enrichment: jsii.String("enrichment"),
//   	enrichmentParameters: &pipeEnrichmentParametersProperty{
//   		httpParameters: &pipeEnrichmentHttpParametersProperty{
//   			headerParameters: map[string]*string{
//   				"headerParametersKey": jsii.String("headerParameters"),
//   			},
//   			pathParameterValues: []*string{
//   				jsii.String("pathParameterValues"),
//   			},
//   			queryStringParameters: map[string]*string{
//   				"queryStringParametersKey": jsii.String("queryStringParameters"),
//   			},
//   		},
//   		inputTemplate: jsii.String("inputTemplate"),
//   	},
//   	name: jsii.String("name"),
//   	sourceParameters: &pipeSourceParametersProperty{
//   		activeMqBrokerParameters: &pipeSourceActiveMQBrokerParametersProperty{
//   			credentials: &mQBrokerAccessCredentialsProperty{
//   				basicAuth: jsii.String("basicAuth"),
//   			},
//   			queueName: jsii.String("queueName"),
//
//   			// the properties below are optional
//   			batchSize: jsii.Number(123),
//   			maximumBatchingWindowInSeconds: jsii.Number(123),
//   		},
//   		dynamoDbStreamParameters: &pipeSourceDynamoDBStreamParametersProperty{
//   			startingPosition: jsii.String("startingPosition"),
//
//   			// the properties below are optional
//   			batchSize: jsii.Number(123),
//   			deadLetterConfig: &deadLetterConfigProperty{
//   				arn: jsii.String("arn"),
//   			},
//   			maximumBatchingWindowInSeconds: jsii.Number(123),
//   			maximumRecordAgeInSeconds: jsii.Number(123),
//   			maximumRetryAttempts: jsii.Number(123),
//   			onPartialBatchItemFailure: jsii.String("onPartialBatchItemFailure"),
//   			parallelizationFactor: jsii.Number(123),
//   		},
//   		filterCriteria: &filterCriteriaProperty{
//   			filters: []interface{}{
//   				&filterProperty{
//   					pattern: jsii.String("pattern"),
//   				},
//   			},
//   		},
//   		kinesisStreamParameters: &pipeSourceKinesisStreamParametersProperty{
//   			startingPosition: jsii.String("startingPosition"),
//
//   			// the properties below are optional
//   			batchSize: jsii.Number(123),
//   			deadLetterConfig: &deadLetterConfigProperty{
//   				arn: jsii.String("arn"),
//   			},
//   			maximumBatchingWindowInSeconds: jsii.Number(123),
//   			maximumRecordAgeInSeconds: jsii.Number(123),
//   			maximumRetryAttempts: jsii.Number(123),
//   			onPartialBatchItemFailure: jsii.String("onPartialBatchItemFailure"),
//   			parallelizationFactor: jsii.Number(123),
//   			startingPositionTimestamp: jsii.String("startingPositionTimestamp"),
//   		},
//   		managedStreamingKafkaParameters: &pipeSourceManagedStreamingKafkaParametersProperty{
//   			topicName: jsii.String("topicName"),
//
//   			// the properties below are optional
//   			batchSize: jsii.Number(123),
//   			consumerGroupId: jsii.String("consumerGroupId"),
//   			credentials: &mSKAccessCredentialsProperty{
//   				clientCertificateTlsAuth: jsii.String("clientCertificateTlsAuth"),
//   				saslScram512Auth: jsii.String("saslScram512Auth"),
//   			},
//   			maximumBatchingWindowInSeconds: jsii.Number(123),
//   			startingPosition: jsii.String("startingPosition"),
//   		},
//   		rabbitMqBrokerParameters: &pipeSourceRabbitMQBrokerParametersProperty{
//   			credentials: &mQBrokerAccessCredentialsProperty{
//   				basicAuth: jsii.String("basicAuth"),
//   			},
//   			queueName: jsii.String("queueName"),
//
//   			// the properties below are optional
//   			batchSize: jsii.Number(123),
//   			maximumBatchingWindowInSeconds: jsii.Number(123),
//   			virtualHost: jsii.String("virtualHost"),
//   		},
//   		selfManagedKafkaParameters: &pipeSourceSelfManagedKafkaParametersProperty{
//   			topicName: jsii.String("topicName"),
//
//   			// the properties below are optional
//   			additionalBootstrapServers: []*string{
//   				jsii.String("additionalBootstrapServers"),
//   			},
//   			batchSize: jsii.Number(123),
//   			consumerGroupId: jsii.String("consumerGroupId"),
//   			credentials: &selfManagedKafkaAccessConfigurationCredentialsProperty{
//   				basicAuth: jsii.String("basicAuth"),
//   				clientCertificateTlsAuth: jsii.String("clientCertificateTlsAuth"),
//   				saslScram256Auth: jsii.String("saslScram256Auth"),
//   				saslScram512Auth: jsii.String("saslScram512Auth"),
//   			},
//   			maximumBatchingWindowInSeconds: jsii.Number(123),
//   			serverRootCaCertificate: jsii.String("serverRootCaCertificate"),
//   			startingPosition: jsii.String("startingPosition"),
//   			vpc: &selfManagedKafkaAccessConfigurationVpcProperty{
//   				securityGroup: []*string{
//   					jsii.String("securityGroup"),
//   				},
//   				subnets: []*string{
//   					jsii.String("subnets"),
//   				},
//   			},
//   		},
//   		sqsQueueParameters: &pipeSourceSqsQueueParametersProperty{
//   			batchSize: jsii.Number(123),
//   			maximumBatchingWindowInSeconds: jsii.Number(123),
//   		},
//   	},
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	targetParameters: &pipeTargetParametersProperty{
//   		batchJobParameters: &pipeTargetBatchJobParametersProperty{
//   			jobDefinition: jsii.String("jobDefinition"),
//   			jobName: jsii.String("jobName"),
//
//   			// the properties below are optional
//   			arrayProperties: &batchArrayPropertiesProperty{
//   				size: jsii.Number(123),
//   			},
//   			containerOverrides: &batchContainerOverridesProperty{
//   				command: []*string{
//   					jsii.String("command"),
//   				},
//   				environment: []interface{}{
//   					&batchEnvironmentVariableProperty{
//   						name: jsii.String("name"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   				instanceType: jsii.String("instanceType"),
//   				resourceRequirements: []interface{}{
//   					&batchResourceRequirementProperty{
//   						type: jsii.String("type"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   			},
//   			dependsOn: []interface{}{
//   				&batchJobDependencyProperty{
//   					jobId: jsii.String("jobId"),
//   					type: jsii.String("type"),
//   				},
//   			},
//   			parameters: map[string]*string{
//   				"parametersKey": jsii.String("parameters"),
//   			},
//   			retryStrategy: &batchRetryStrategyProperty{
//   				attempts: jsii.Number(123),
//   			},
//   		},
//   		cloudWatchLogsParameters: &pipeTargetCloudWatchLogsParametersProperty{
//   			logStreamName: jsii.String("logStreamName"),
//   			timestamp: jsii.String("timestamp"),
//   		},
//   		ecsTaskParameters: &pipeTargetEcsTaskParametersProperty{
//   			taskDefinitionArn: jsii.String("taskDefinitionArn"),
//
//   			// the properties below are optional
//   			capacityProviderStrategy: []interface{}{
//   				&capacityProviderStrategyItemProperty{
//   					capacityProvider: jsii.String("capacityProvider"),
//
//   					// the properties below are optional
//   					base: jsii.Number(123),
//   					weight: jsii.Number(123),
//   				},
//   			},
//   			enableEcsManagedTags: jsii.Boolean(false),
//   			enableExecuteCommand: jsii.Boolean(false),
//   			group: jsii.String("group"),
//   			launchType: jsii.String("launchType"),
//   			networkConfiguration: &networkConfigurationProperty{
//   				awsvpcConfiguration: &awsVpcConfigurationProperty{
//   					subnets: []*string{
//   						jsii.String("subnets"),
//   					},
//
//   					// the properties below are optional
//   					assignPublicIp: jsii.String("assignPublicIp"),
//   					securityGroups: []*string{
//   						jsii.String("securityGroups"),
//   					},
//   				},
//   			},
//   			overrides: &ecsTaskOverrideProperty{
//   				containerOverrides: []interface{}{
//   					&ecsContainerOverrideProperty{
//   						command: []*string{
//   							jsii.String("command"),
//   						},
//   						cpu: jsii.Number(123),
//   						environment: []interface{}{
//   							&ecsEnvironmentVariableProperty{
//   								name: jsii.String("name"),
//   								value: jsii.String("value"),
//   							},
//   						},
//   						environmentFiles: []interface{}{
//   							&ecsEnvironmentFileProperty{
//   								type: jsii.String("type"),
//   								value: jsii.String("value"),
//   							},
//   						},
//   						memory: jsii.Number(123),
//   						memoryReservation: jsii.Number(123),
//   						name: jsii.String("name"),
//   						resourceRequirements: []interface{}{
//   							&ecsResourceRequirementProperty{
//   								type: jsii.String("type"),
//   								value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   				cpu: jsii.String("cpu"),
//   				ephemeralStorage: &ecsEphemeralStorageProperty{
//   					sizeInGiB: jsii.Number(123),
//   				},
//   				executionRoleArn: jsii.String("executionRoleArn"),
//   				inferenceAcceleratorOverrides: []interface{}{
//   					&ecsInferenceAcceleratorOverrideProperty{
//   						deviceName: jsii.String("deviceName"),
//   						deviceType: jsii.String("deviceType"),
//   					},
//   				},
//   				memory: jsii.String("memory"),
//   				taskRoleArn: jsii.String("taskRoleArn"),
//   			},
//   			placementConstraints: []interface{}{
//   				&placementConstraintProperty{
//   					expression: jsii.String("expression"),
//   					type: jsii.String("type"),
//   				},
//   			},
//   			placementStrategy: []interface{}{
//   				&placementStrategyProperty{
//   					field: jsii.String("field"),
//   					type: jsii.String("type"),
//   				},
//   			},
//   			platformVersion: jsii.String("platformVersion"),
//   			propagateTags: jsii.String("propagateTags"),
//   			referenceId: jsii.String("referenceId"),
//   			tags: []cfnTag{
//   				&cfnTag{
//   					key: jsii.String("key"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   			taskCount: jsii.Number(123),
//   		},
//   		eventBridgeEventBusParameters: &pipeTargetEventBridgeEventBusParametersProperty{
//   			detailType: jsii.String("detailType"),
//   			endpointId: jsii.String("endpointId"),
//   			resources: []*string{
//   				jsii.String("resources"),
//   			},
//   			source: jsii.String("source"),
//   			time: jsii.String("time"),
//   		},
//   		httpParameters: &pipeTargetHttpParametersProperty{
//   			headerParameters: map[string]*string{
//   				"headerParametersKey": jsii.String("headerParameters"),
//   			},
//   			pathParameterValues: []*string{
//   				jsii.String("pathParameterValues"),
//   			},
//   			queryStringParameters: map[string]*string{
//   				"queryStringParametersKey": jsii.String("queryStringParameters"),
//   			},
//   		},
//   		inputTemplate: jsii.String("inputTemplate"),
//   		kinesisStreamParameters: &pipeTargetKinesisStreamParametersProperty{
//   			partitionKey: jsii.String("partitionKey"),
//   		},
//   		lambdaFunctionParameters: &pipeTargetLambdaFunctionParametersProperty{
//   			invocationType: jsii.String("invocationType"),
//   		},
//   		redshiftDataParameters: &pipeTargetRedshiftDataParametersProperty{
//   			database: jsii.String("database"),
//   			sqls: []*string{
//   				jsii.String("sqls"),
//   			},
//
//   			// the properties below are optional
//   			dbUser: jsii.String("dbUser"),
//   			secretManagerArn: jsii.String("secretManagerArn"),
//   			statementName: jsii.String("statementName"),
//   			withEvent: jsii.Boolean(false),
//   		},
//   		sageMakerPipelineParameters: &pipeTargetSageMakerPipelineParametersProperty{
//   			pipelineParameterList: []interface{}{
//   				&sageMakerPipelineParameterProperty{
//   					name: jsii.String("name"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   		sqsQueueParameters: &pipeTargetSqsQueueParametersProperty{
//   			messageDeduplicationId: jsii.String("messageDeduplicationId"),
//   			messageGroupId: jsii.String("messageGroupId"),
//   		},
//   		stepFunctionStateMachineParameters: &pipeTargetStateMachineParametersProperty{
//   			invocationType: jsii.String("invocationType"),
//   		},
//   	},
//   })
//
type CfnPipe interface {
	awscdk.CfnResource
	awscdk.IInspectable
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
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
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
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The name of the pipe.
	Name() *string
	SetName(val *string)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
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
	// Experimental.
	Stack() awscdk.Stack
	// The list of key-value pairs to associate with the pipe.
	Tags() awscdk.TagManager
	// The ARN of the target resource.
	Target() *string
	SetTarget(val *string)
	// The parameters required to set up a target for your pipe.
	TargetParameters() interface{}
	SetTargetParameters(val interface{})
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
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
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnPipe
type jsiiProxy_CfnPipe struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
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

func (j *jsiiProxy_CfnPipe) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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


// Create a new `AWS::Pipes::Pipe`.
func NewCfnPipe(scope awscdk.Construct, id *string, props *CfnPipeProps) CfnPipe {
	_init_.Initialize()

	if err := validateNewCfnPipeParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnPipe{}

	_jsii_.Create(
		"monocdk.aws_pipes.CfnPipe",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Pipes::Pipe`.
func NewCfnPipe_Override(c CfnPipe, scope awscdk.Construct, id *string, props *CfnPipeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_pipes.CfnPipe",
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
// Experimental.
func CfnPipe_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPipe_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_pipes.CfnPipe",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnPipe_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnPipe_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_pipes.CfnPipe",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnPipe_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPipe_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_pipes.CfnPipe",
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
		"monocdk.aws_pipes.CfnPipe",
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

func (c *jsiiProxy_CfnPipe) GetAtt(attributeName *string) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
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

func (c *jsiiProxy_CfnPipe) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnPipe) OnSynthesize(session constructs.ISynthesisSession) {
	if err := c.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnPipe) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
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

func (c *jsiiProxy_CfnPipe) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
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

func (c *jsiiProxy_CfnPipe) Synthesize(session awscdk.ISynthesisSession) {
	if err := c.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
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

func (c *jsiiProxy_CfnPipe) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
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

