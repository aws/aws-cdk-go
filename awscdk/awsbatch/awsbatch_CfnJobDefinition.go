package awsbatch

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbatch/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::Batch::JobDefinition`.
//
// The `AWS::Batch::JobDefinition` resource specifies the parameters for an AWS Batch job definition. For more information, see [Job Definitions](https://docs.aws.amazon.com/batch/latest/userguide/job_definitions.html) in the ** .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var limits interface{}
//   var options interface{}
//   var parameters interface{}
//   var requests interface{}
//   var tags interface{}
//
//   cfnJobDefinition := awscdk.Aws_batch.NewCfnJobDefinition(this, jsii.String("MyCfnJobDefinition"), &cfnJobDefinitionProps{
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	containerProperties: &containerPropertiesProperty{
//   		image: jsii.String("image"),
//
//   		// the properties below are optional
//   		command: []*string{
//   			jsii.String("command"),
//   		},
//   		environment: []interface{}{
//   			&environmentProperty{
//   				name: jsii.String("name"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		executionRoleArn: jsii.String("executionRoleArn"),
//   		fargatePlatformConfiguration: &fargatePlatformConfigurationProperty{
//   			platformVersion: jsii.String("platformVersion"),
//   		},
//   		instanceType: jsii.String("instanceType"),
//   		jobRoleArn: jsii.String("jobRoleArn"),
//   		linuxParameters: &linuxParametersProperty{
//   			devices: []interface{}{
//   				&deviceProperty{
//   					containerPath: jsii.String("containerPath"),
//   					hostPath: jsii.String("hostPath"),
//   					permissions: []*string{
//   						jsii.String("permissions"),
//   					},
//   				},
//   			},
//   			initProcessEnabled: jsii.Boolean(false),
//   			maxSwap: jsii.Number(123),
//   			sharedMemorySize: jsii.Number(123),
//   			swappiness: jsii.Number(123),
//   			tmpfs: []interface{}{
//   				&tmpfsProperty{
//   					containerPath: jsii.String("containerPath"),
//   					size: jsii.Number(123),
//
//   					// the properties below are optional
//   					mountOptions: []*string{
//   						jsii.String("mountOptions"),
//   					},
//   				},
//   			},
//   		},
//   		logConfiguration: &logConfigurationProperty{
//   			logDriver: jsii.String("logDriver"),
//
//   			// the properties below are optional
//   			options: options,
//   			secretOptions: []interface{}{
//   				&secretProperty{
//   					name: jsii.String("name"),
//   					valueFrom: jsii.String("valueFrom"),
//   				},
//   			},
//   		},
//   		memory: jsii.Number(123),
//   		mountPoints: []interface{}{
//   			&mountPointsProperty{
//   				containerPath: jsii.String("containerPath"),
//   				readOnly: jsii.Boolean(false),
//   				sourceVolume: jsii.String("sourceVolume"),
//   			},
//   		},
//   		networkConfiguration: &networkConfigurationProperty{
//   			assignPublicIp: jsii.String("assignPublicIp"),
//   		},
//   		privileged: jsii.Boolean(false),
//   		readonlyRootFilesystem: jsii.Boolean(false),
//   		resourceRequirements: []interface{}{
//   			&resourceRequirementProperty{
//   				type: jsii.String("type"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		secrets: []interface{}{
//   			&secretProperty{
//   				name: jsii.String("name"),
//   				valueFrom: jsii.String("valueFrom"),
//   			},
//   		},
//   		ulimits: []interface{}{
//   			&ulimitProperty{
//   				hardLimit: jsii.Number(123),
//   				name: jsii.String("name"),
//   				softLimit: jsii.Number(123),
//   			},
//   		},
//   		user: jsii.String("user"),
//   		vcpus: jsii.Number(123),
//   		volumes: []interface{}{
//   			&volumesProperty{
//   				efsVolumeConfiguration: &efsVolumeConfigurationProperty{
//   					fileSystemId: jsii.String("fileSystemId"),
//
//   					// the properties below are optional
//   					authorizationConfig: &authorizationConfigProperty{
//   						accessPointId: jsii.String("accessPointId"),
//   						iam: jsii.String("iam"),
//   					},
//   					rootDirectory: jsii.String("rootDirectory"),
//   					transitEncryption: jsii.String("transitEncryption"),
//   					transitEncryptionPort: jsii.Number(123),
//   				},
//   				host: &volumesHostProperty{
//   					sourcePath: jsii.String("sourcePath"),
//   				},
//   				name: jsii.String("name"),
//   			},
//   		},
//   	},
//   	eksProperties: &eksPropertiesProperty{
//   		podProperties: &podPropertiesProperty{
//   			containers: []interface{}{
//   				&eksContainerProperty{
//   					image: jsii.String("image"),
//
//   					// the properties below are optional
//   					args: []*string{
//   						jsii.String("args"),
//   					},
//   					command: []*string{
//   						jsii.String("command"),
//   					},
//   					env: []interface{}{
//   						&eksContainerEnvironmentVariableProperty{
//   							name: jsii.String("name"),
//
//   							// the properties below are optional
//   							value: jsii.String("value"),
//   						},
//   					},
//   					imagePullPolicy: jsii.String("imagePullPolicy"),
//   					name: jsii.String("name"),
//   					resources: &resourcesProperty{
//   						limits: limits,
//   						requests: requests,
//   					},
//   					securityContext: &securityContextProperty{
//   						privileged: jsii.Boolean(false),
//   						readOnlyRootFilesystem: jsii.Boolean(false),
//   						runAsGroup: jsii.Number(123),
//   						runAsNonRoot: jsii.Boolean(false),
//   						runAsUser: jsii.Number(123),
//   					},
//   					volumeMounts: []interface{}{
//   						&eksContainerVolumeMountProperty{
//   							mountPath: jsii.String("mountPath"),
//   							name: jsii.String("name"),
//   							readOnly: jsii.Boolean(false),
//   						},
//   					},
//   				},
//   			},
//   			dnsPolicy: jsii.String("dnsPolicy"),
//   			hostNetwork: jsii.Boolean(false),
//   			serviceAccountName: jsii.String("serviceAccountName"),
//   			volumes: []interface{}{
//   				&eksVolumeProperty{
//   					name: jsii.String("name"),
//
//   					// the properties below are optional
//   					emptyDir: &emptyDirProperty{
//   						medium: jsii.String("medium"),
//   						sizeLimit: jsii.String("sizeLimit"),
//   					},
//   					hostPath: &hostPathProperty{
//   						path: jsii.String("path"),
//   					},
//   					secret: &secretProperty{
//   						name: jsii.String("name"),
//   						valueFrom: jsii.String("valueFrom"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	jobDefinitionName: jsii.String("jobDefinitionName"),
//   	nodeProperties: &nodePropertiesProperty{
//   		mainNode: jsii.Number(123),
//   		nodeRangeProperties: []interface{}{
//   			&nodeRangePropertyProperty{
//   				targetNodes: jsii.String("targetNodes"),
//
//   				// the properties below are optional
//   				container: &containerPropertiesProperty{
//   					image: jsii.String("image"),
//
//   					// the properties below are optional
//   					command: []*string{
//   						jsii.String("command"),
//   					},
//   					environment: []interface{}{
//   						&environmentProperty{
//   							name: jsii.String("name"),
//   							value: jsii.String("value"),
//   						},
//   					},
//   					executionRoleArn: jsii.String("executionRoleArn"),
//   					fargatePlatformConfiguration: &fargatePlatformConfigurationProperty{
//   						platformVersion: jsii.String("platformVersion"),
//   					},
//   					instanceType: jsii.String("instanceType"),
//   					jobRoleArn: jsii.String("jobRoleArn"),
//   					linuxParameters: &linuxParametersProperty{
//   						devices: []interface{}{
//   							&deviceProperty{
//   								containerPath: jsii.String("containerPath"),
//   								hostPath: jsii.String("hostPath"),
//   								permissions: []*string{
//   									jsii.String("permissions"),
//   								},
//   							},
//   						},
//   						initProcessEnabled: jsii.Boolean(false),
//   						maxSwap: jsii.Number(123),
//   						sharedMemorySize: jsii.Number(123),
//   						swappiness: jsii.Number(123),
//   						tmpfs: []interface{}{
//   							&tmpfsProperty{
//   								containerPath: jsii.String("containerPath"),
//   								size: jsii.Number(123),
//
//   								// the properties below are optional
//   								mountOptions: []*string{
//   									jsii.String("mountOptions"),
//   								},
//   							},
//   						},
//   					},
//   					logConfiguration: &logConfigurationProperty{
//   						logDriver: jsii.String("logDriver"),
//
//   						// the properties below are optional
//   						options: options,
//   						secretOptions: []interface{}{
//   							&secretProperty{
//   								name: jsii.String("name"),
//   								valueFrom: jsii.String("valueFrom"),
//   							},
//   						},
//   					},
//   					memory: jsii.Number(123),
//   					mountPoints: []interface{}{
//   						&mountPointsProperty{
//   							containerPath: jsii.String("containerPath"),
//   							readOnly: jsii.Boolean(false),
//   							sourceVolume: jsii.String("sourceVolume"),
//   						},
//   					},
//   					networkConfiguration: &networkConfigurationProperty{
//   						assignPublicIp: jsii.String("assignPublicIp"),
//   					},
//   					privileged: jsii.Boolean(false),
//   					readonlyRootFilesystem: jsii.Boolean(false),
//   					resourceRequirements: []interface{}{
//   						&resourceRequirementProperty{
//   							type: jsii.String("type"),
//   							value: jsii.String("value"),
//   						},
//   					},
//   					secrets: []interface{}{
//   						&secretProperty{
//   							name: jsii.String("name"),
//   							valueFrom: jsii.String("valueFrom"),
//   						},
//   					},
//   					ulimits: []interface{}{
//   						&ulimitProperty{
//   							hardLimit: jsii.Number(123),
//   							name: jsii.String("name"),
//   							softLimit: jsii.Number(123),
//   						},
//   					},
//   					user: jsii.String("user"),
//   					vcpus: jsii.Number(123),
//   					volumes: []interface{}{
//   						&volumesProperty{
//   							efsVolumeConfiguration: &efsVolumeConfigurationProperty{
//   								fileSystemId: jsii.String("fileSystemId"),
//
//   								// the properties below are optional
//   								authorizationConfig: &authorizationConfigProperty{
//   									accessPointId: jsii.String("accessPointId"),
//   									iam: jsii.String("iam"),
//   								},
//   								rootDirectory: jsii.String("rootDirectory"),
//   								transitEncryption: jsii.String("transitEncryption"),
//   								transitEncryptionPort: jsii.Number(123),
//   							},
//   							host: &volumesHostProperty{
//   								sourcePath: jsii.String("sourcePath"),
//   							},
//   							name: jsii.String("name"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		numNodes: jsii.Number(123),
//   	},
//   	parameters: parameters,
//   	platformCapabilities: []*string{
//   		jsii.String("platformCapabilities"),
//   	},
//   	propagateTags: jsii.Boolean(false),
//   	retryStrategy: &retryStrategyProperty{
//   		attempts: jsii.Number(123),
//   		evaluateOnExit: []interface{}{
//   			&evaluateOnExitProperty{
//   				action: jsii.String("action"),
//
//   				// the properties below are optional
//   				onExitCode: jsii.String("onExitCode"),
//   				onReason: jsii.String("onReason"),
//   				onStatusReason: jsii.String("onStatusReason"),
//   			},
//   		},
//   	},
//   	schedulingPriority: jsii.Number(123),
//   	tags: tags,
//   	timeout: &timeoutProperty{
//   		attemptDurationSeconds: jsii.Number(123),
//   	},
//   })
//
type CfnJobDefinition interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// An object with various properties specific to container-based jobs.
	ContainerProperties() interface{}
	SetContainerProperties(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// `AWS::Batch::JobDefinition.EksProperties`.
	EksProperties() interface{}
	SetEksProperties(val interface{})
	// The name of the job definition.
	JobDefinitionName() *string
	SetJobDefinitionName(val *string)
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
	// The tree node.
	Node() constructs.Node
	// An object with various properties specific to multi-node parallel jobs.
	//
	// > If the job runs on Fargate resources, then you must not specify `nodeProperties` ; use `containerProperties` instead.
	NodeProperties() interface{}
	SetNodeProperties(val interface{})
	// Default parameters or parameter substitution placeholders that are set in the job definition.
	//
	// Parameters are specified as a key-value pair mapping. Parameters in a `SubmitJob` request override any corresponding parameter defaults from the job definition. For more information about specifying parameters, see [Job definition parameters](https://docs.aws.amazon.com/batch/latest/userguide/job_definition_parameters.html) in the *AWS Batch User Guide* .
	Parameters() interface{}
	SetParameters(val interface{})
	// The platform capabilities required by the job definition.
	//
	// If no value is specified, it defaults to `EC2` . Jobs run on Fargate resources specify `FARGATE` .
	PlatformCapabilities() *[]*string
	SetPlatformCapabilities(val *[]*string)
	// Specifies whether to propagate the tags from the job or job definition to the corresponding Amazon ECS task.
	//
	// If no value is specified, the tags aren't propagated. Tags can only be propagated to the tasks during task creation. For tags with the same name, job tags are given priority over job definitions tags. If the total number of combined tags from the job and job definition is over 50, the job is moved to the `FAILED` state.
	PropagateTags() interface{}
	SetPropagateTags(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The retry strategy to use for failed jobs that are submitted with this job definition.
	RetryStrategy() interface{}
	SetRetryStrategy(val interface{})
	// The scheduling priority of the job definition.
	//
	// This only affects jobs in job queues with a fair share policy. Jobs with a higher scheduling priority are scheduled before jobs with a lower scheduling priority.
	SchedulingPriority() *float64
	SetSchedulingPriority(val *float64)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The tags applied to the job definition.
	Tags() awscdk.TagManager
	// The timeout configuration for jobs that are submitted with this job definition.
	//
	// You can specify a timeout duration after which AWS Batch terminates your jobs if they haven't finished.
	Timeout() interface{}
	SetTimeout(val interface{})
	// The type of job definition.
	//
	// For more information about multi-node parallel jobs, see [Creating a multi-node parallel job definition](https://docs.aws.amazon.com/batch/latest/userguide/multi-node-job-def.html) in the *AWS Batch User Guide* .
	//
	// > If the job is run on Fargate resources, then `multinode` isn't supported.
	Type() *string
	SetType(val *string)
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
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
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

// The jsii proxy struct for CfnJobDefinition
type jsiiProxy_CfnJobDefinition struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnJobDefinition) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobDefinition) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobDefinition) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobDefinition) ContainerProperties() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"containerProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobDefinition) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobDefinition) EksProperties() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"eksProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobDefinition) JobDefinitionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobDefinitionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobDefinition) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobDefinition) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobDefinition) NodeProperties() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nodeProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobDefinition) Parameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobDefinition) PlatformCapabilities() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"platformCapabilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobDefinition) PropagateTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"propagateTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobDefinition) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobDefinition) RetryStrategy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retryStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobDefinition) SchedulingPriority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"schedulingPriority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobDefinition) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobDefinition) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobDefinition) Timeout() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobDefinition) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobDefinition) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobDefinition) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::Batch::JobDefinition`.
func NewCfnJobDefinition(scope constructs.Construct, id *string, props *CfnJobDefinitionProps) CfnJobDefinition {
	_init_.Initialize()

	if err := validateNewCfnJobDefinitionParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnJobDefinition{}

	_jsii_.Create(
		"aws-cdk-lib.aws_batch.CfnJobDefinition",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Batch::JobDefinition`.
func NewCfnJobDefinition_Override(c CfnJobDefinition, scope constructs.Construct, id *string, props *CfnJobDefinitionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_batch.CfnJobDefinition",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnJobDefinition)SetContainerProperties(val interface{}) {
	if err := j.validateSetContainerPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerProperties",
		val,
	)
}

func (j *jsiiProxy_CfnJobDefinition)SetEksProperties(val interface{}) {
	if err := j.validateSetEksPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eksProperties",
		val,
	)
}

func (j *jsiiProxy_CfnJobDefinition)SetJobDefinitionName(val *string) {
	_jsii_.Set(
		j,
		"jobDefinitionName",
		val,
	)
}

func (j *jsiiProxy_CfnJobDefinition)SetNodeProperties(val interface{}) {
	if err := j.validateSetNodePropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeProperties",
		val,
	)
}

func (j *jsiiProxy_CfnJobDefinition)SetParameters(val interface{}) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_CfnJobDefinition)SetPlatformCapabilities(val *[]*string) {
	_jsii_.Set(
		j,
		"platformCapabilities",
		val,
	)
}

func (j *jsiiProxy_CfnJobDefinition)SetPropagateTags(val interface{}) {
	if err := j.validateSetPropagateTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"propagateTags",
		val,
	)
}

func (j *jsiiProxy_CfnJobDefinition)SetRetryStrategy(val interface{}) {
	if err := j.validateSetRetryStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retryStrategy",
		val,
	)
}

func (j *jsiiProxy_CfnJobDefinition)SetSchedulingPriority(val *float64) {
	_jsii_.Set(
		j,
		"schedulingPriority",
		val,
	)
}

func (j *jsiiProxy_CfnJobDefinition)SetTimeout(val interface{}) {
	if err := j.validateSetTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (j *jsiiProxy_CfnJobDefinition)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnJobDefinition_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnJobDefinition_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_batch.CfnJobDefinition",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnJobDefinition_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnJobDefinition_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_batch.CfnJobDefinition",
		"isCfnResource",
		[]interface{}{construct},
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
func CfnJobDefinition_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnJobDefinition_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_batch.CfnJobDefinition",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnJobDefinition_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_batch.CfnJobDefinition",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnJobDefinition) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnJobDefinition) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnJobDefinition) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnJobDefinition) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnJobDefinition) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnJobDefinition) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnJobDefinition) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnJobDefinition) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnJobDefinition) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnJobDefinition) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnJobDefinition) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnJobDefinition) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnJobDefinition) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnJobDefinition) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnJobDefinition) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

