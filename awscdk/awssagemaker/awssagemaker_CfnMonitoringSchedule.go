package awssagemaker

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::SageMaker::MonitoringSchedule`.
//
// The `AWS::SageMaker::MonitoringSchedule` resource is an Amazon SageMaker resource type that regularly starts SageMaker processing Jobs to monitor the data captured for a SageMaker endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var json interface{}
//
//   cfnMonitoringSchedule := awscdk.Aws_sagemaker.NewCfnMonitoringSchedule(this, jsii.String("MyCfnMonitoringSchedule"), &cfnMonitoringScheduleProps{
//   	monitoringScheduleConfig: &monitoringScheduleConfigProperty{
//   		monitoringJobDefinition: &monitoringJobDefinitionProperty{
//   			monitoringAppSpecification: &monitoringAppSpecificationProperty{
//   				imageUri: jsii.String("imageUri"),
//
//   				// the properties below are optional
//   				containerArguments: []*string{
//   					jsii.String("containerArguments"),
//   				},
//   				containerEntrypoint: []*string{
//   					jsii.String("containerEntrypoint"),
//   				},
//   				postAnalyticsProcessorSourceUri: jsii.String("postAnalyticsProcessorSourceUri"),
//   				recordPreprocessorSourceUri: jsii.String("recordPreprocessorSourceUri"),
//   			},
//   			monitoringInputs: []interface{}{
//   				&monitoringInputProperty{
//   					batchTransformInput: &batchTransformInputProperty{
//   						dataCapturedDestinationS3Uri: jsii.String("dataCapturedDestinationS3Uri"),
//   						datasetFormat: &datasetFormatProperty{
//   							csv: &csvProperty{
//   								header: jsii.Boolean(false),
//   							},
//   							json: json,
//   							parquet: jsii.Boolean(false),
//   						},
//   						localPath: jsii.String("localPath"),
//
//   						// the properties below are optional
//   						s3DataDistributionType: jsii.String("s3DataDistributionType"),
//   						s3InputMode: jsii.String("s3InputMode"),
//   					},
//   					endpointInput: &endpointInputProperty{
//   						endpointName: jsii.String("endpointName"),
//   						localPath: jsii.String("localPath"),
//
//   						// the properties below are optional
//   						s3DataDistributionType: jsii.String("s3DataDistributionType"),
//   						s3InputMode: jsii.String("s3InputMode"),
//   					},
//   				},
//   			},
//   			monitoringOutputConfig: &monitoringOutputConfigProperty{
//   				monitoringOutputs: []interface{}{
//   					&monitoringOutputProperty{
//   						s3Output: &s3OutputProperty{
//   							localPath: jsii.String("localPath"),
//   							s3Uri: jsii.String("s3Uri"),
//
//   							// the properties below are optional
//   							s3UploadMode: jsii.String("s3UploadMode"),
//   						},
//   					},
//   				},
//
//   				// the properties below are optional
//   				kmsKeyId: jsii.String("kmsKeyId"),
//   			},
//   			monitoringResources: &monitoringResourcesProperty{
//   				clusterConfig: &clusterConfigProperty{
//   					instanceCount: jsii.Number(123),
//   					instanceType: jsii.String("instanceType"),
//   					volumeSizeInGb: jsii.Number(123),
//
//   					// the properties below are optional
//   					volumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   				},
//   			},
//   			roleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			baselineConfig: &baselineConfigProperty{
//   				constraintsResource: &constraintsResourceProperty{
//   					s3Uri: jsii.String("s3Uri"),
//   				},
//   				statisticsResource: &statisticsResourceProperty{
//   					s3Uri: jsii.String("s3Uri"),
//   				},
//   			},
//   			environment: map[string]*string{
//   				"environmentKey": jsii.String("environment"),
//   			},
//   			networkConfig: &networkConfigProperty{
//   				enableInterContainerTrafficEncryption: jsii.Boolean(false),
//   				enableNetworkIsolation: jsii.Boolean(false),
//   				vpcConfig: &vpcConfigProperty{
//   					securityGroupIds: []*string{
//   						jsii.String("securityGroupIds"),
//   					},
//   					subnets: []*string{
//   						jsii.String("subnets"),
//   					},
//   				},
//   			},
//   			stoppingCondition: &stoppingConditionProperty{
//   				maxRuntimeInSeconds: jsii.Number(123),
//   			},
//   		},
//   		monitoringJobDefinitionName: jsii.String("monitoringJobDefinitionName"),
//   		monitoringType: jsii.String("monitoringType"),
//   		scheduleConfig: &scheduleConfigProperty{
//   			scheduleExpression: jsii.String("scheduleExpression"),
//   		},
//   	},
//   	monitoringScheduleName: jsii.String("monitoringScheduleName"),
//
//   	// the properties below are optional
//   	endpointName: jsii.String("endpointName"),
//   	failureReason: jsii.String("failureReason"),
//   	lastMonitoringExecutionSummary: &monitoringExecutionSummaryProperty{
//   		creationTime: jsii.String("creationTime"),
//   		lastModifiedTime: jsii.String("lastModifiedTime"),
//   		monitoringExecutionStatus: jsii.String("monitoringExecutionStatus"),
//   		monitoringScheduleName: jsii.String("monitoringScheduleName"),
//   		scheduledTime: jsii.String("scheduledTime"),
//
//   		// the properties below are optional
//   		endpointName: jsii.String("endpointName"),
//   		failureReason: jsii.String("failureReason"),
//   		processingJobArn: jsii.String("processingJobArn"),
//   	},
//   	monitoringScheduleStatus: jsii.String("monitoringScheduleStatus"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnMonitoringSchedule interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The time when the monitoring schedule was created.
	AttrCreationTime() *string
	// The last time that the monitoring schedule was modified.
	AttrLastModifiedTime() *string
	// The Amazon Resource Name (ARN) of the monitoring schedule.
	AttrMonitoringScheduleArn() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The name of the endpoint using the monitoring schedule.
	EndpointName() *string
	SetEndpointName(val *string)
	// Contains the reason a monitoring job failed, if it failed.
	FailureReason() *string
	SetFailureReason(val *string)
	// Describes metadata on the last execution to run, if there was one.
	LastMonitoringExecutionSummary() interface{}
	SetLastMonitoringExecutionSummary(val interface{})
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
	// The configuration object that specifies the monitoring schedule and defines the monitoring job.
	MonitoringScheduleConfig() interface{}
	SetMonitoringScheduleConfig(val interface{})
	// The name of the monitoring schedule.
	MonitoringScheduleName() *string
	SetMonitoringScheduleName(val *string)
	// The status of the monitoring schedule.
	MonitoringScheduleStatus() *string
	SetMonitoringScheduleStatus(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags() awscdk.TagManager
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
	GetAtt(attributeName *string) awscdk.Reference
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

// The jsii proxy struct for CfnMonitoringSchedule
type jsiiProxy_CfnMonitoringSchedule struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnMonitoringSchedule) AttrCreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitoringSchedule) AttrLastModifiedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLastModifiedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitoringSchedule) AttrMonitoringScheduleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrMonitoringScheduleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitoringSchedule) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitoringSchedule) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitoringSchedule) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitoringSchedule) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitoringSchedule) EndpointName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitoringSchedule) FailureReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failureReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitoringSchedule) LastMonitoringExecutionSummary() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lastMonitoringExecutionSummary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitoringSchedule) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitoringSchedule) MonitoringScheduleConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"monitoringScheduleConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitoringSchedule) MonitoringScheduleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringScheduleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitoringSchedule) MonitoringScheduleStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringScheduleStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitoringSchedule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitoringSchedule) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitoringSchedule) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitoringSchedule) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitoringSchedule) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitoringSchedule) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::SageMaker::MonitoringSchedule`.
func NewCfnMonitoringSchedule(scope constructs.Construct, id *string, props *CfnMonitoringScheduleProps) CfnMonitoringSchedule {
	_init_.Initialize()

	if err := validateNewCfnMonitoringScheduleParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnMonitoringSchedule{}

	_jsii_.Create(
		"aws-cdk-lib.aws_sagemaker.CfnMonitoringSchedule",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::SageMaker::MonitoringSchedule`.
func NewCfnMonitoringSchedule_Override(c CfnMonitoringSchedule, scope constructs.Construct, id *string, props *CfnMonitoringScheduleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_sagemaker.CfnMonitoringSchedule",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnMonitoringSchedule)SetEndpointName(val *string) {
	_jsii_.Set(
		j,
		"endpointName",
		val,
	)
}

func (j *jsiiProxy_CfnMonitoringSchedule)SetFailureReason(val *string) {
	_jsii_.Set(
		j,
		"failureReason",
		val,
	)
}

func (j *jsiiProxy_CfnMonitoringSchedule)SetLastMonitoringExecutionSummary(val interface{}) {
	if err := j.validateSetLastMonitoringExecutionSummaryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lastMonitoringExecutionSummary",
		val,
	)
}

func (j *jsiiProxy_CfnMonitoringSchedule)SetMonitoringScheduleConfig(val interface{}) {
	if err := j.validateSetMonitoringScheduleConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitoringScheduleConfig",
		val,
	)
}

func (j *jsiiProxy_CfnMonitoringSchedule)SetMonitoringScheduleName(val *string) {
	if err := j.validateSetMonitoringScheduleNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitoringScheduleName",
		val,
	)
}

func (j *jsiiProxy_CfnMonitoringSchedule)SetMonitoringScheduleStatus(val *string) {
	_jsii_.Set(
		j,
		"monitoringScheduleStatus",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnMonitoringSchedule_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnMonitoringSchedule_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sagemaker.CfnMonitoringSchedule",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnMonitoringSchedule_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnMonitoringSchedule_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sagemaker.CfnMonitoringSchedule",
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
func CfnMonitoringSchedule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnMonitoringSchedule_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sagemaker.CfnMonitoringSchedule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnMonitoringSchedule_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_sagemaker.CfnMonitoringSchedule",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnMonitoringSchedule) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnMonitoringSchedule) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnMonitoringSchedule) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnMonitoringSchedule) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnMonitoringSchedule) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnMonitoringSchedule) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnMonitoringSchedule) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnMonitoringSchedule) GetAtt(attributeName *string) awscdk.Reference {
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

func (c *jsiiProxy_CfnMonitoringSchedule) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnMonitoringSchedule) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnMonitoringSchedule) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnMonitoringSchedule) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnMonitoringSchedule) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMonitoringSchedule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMonitoringSchedule) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

