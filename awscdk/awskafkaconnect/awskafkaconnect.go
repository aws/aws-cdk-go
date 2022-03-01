package awskafkaconnect

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskafkaconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::KafkaConnect::Connector`.
//
// TODO: EXAMPLE
//
type CfnConnector interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrConnectorArn() *string
	Capacity() interface{}
	SetCapacity(val interface{})
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	ConnectorConfiguration() interface{}
	SetConnectorConfiguration(val interface{})
	ConnectorDescription() *string
	SetConnectorDescription(val *string)
	ConnectorName() *string
	SetConnectorName(val *string)
	CreationStack() *[]*string
	KafkaCluster() interface{}
	SetKafkaCluster(val interface{})
	KafkaClusterClientAuthentication() interface{}
	SetKafkaClusterClientAuthentication(val interface{})
	KafkaClusterEncryptionInTransit() interface{}
	SetKafkaClusterEncryptionInTransit(val interface{})
	KafkaConnectVersion() *string
	SetKafkaConnectVersion(val *string)
	LogDelivery() interface{}
	SetLogDelivery(val interface{})
	LogicalId() *string
	Node() constructs.Node
	Plugins() interface{}
	SetPlugins(val interface{})
	Ref() *string
	ServiceExecutionRoleArn() *string
	SetServiceExecutionRoleArn(val *string)
	Stack() awscdk.Stack
	UpdatedProperites() *map[string]interface{}
	WorkerConfiguration() interface{}
	SetWorkerConfiguration(val interface{})
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnConnector
type jsiiProxy_CfnConnector struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnConnector) AttrConnectorArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrConnectorArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnector) Capacity() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"capacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnector) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnector) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnector) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnector) ConnectorConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectorConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnector) ConnectorDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectorDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnector) ConnectorName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectorName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnector) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnector) KafkaCluster() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kafkaCluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnector) KafkaClusterClientAuthentication() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kafkaClusterClientAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnector) KafkaClusterEncryptionInTransit() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kafkaClusterEncryptionInTransit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnector) KafkaConnectVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kafkaConnectVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnector) LogDelivery() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logDelivery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnector) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnector) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnector) Plugins() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"plugins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnector) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnector) ServiceExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceExecutionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnector) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnector) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnector) WorkerConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"workerConfiguration",
		&returns,
	)
	return returns
}


// Create a new `AWS::KafkaConnect::Connector`.
func NewCfnConnector(scope constructs.Construct, id *string, props *CfnConnectorProps) CfnConnector {
	_init_.Initialize()

	j := jsiiProxy_CfnConnector{}

	_jsii_.Create(
		"aws-cdk-lib.aws_kafkaconnect.CfnConnector",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::KafkaConnect::Connector`.
func NewCfnConnector_Override(c CfnConnector, scope constructs.Construct, id *string, props *CfnConnectorProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_kafkaconnect.CfnConnector",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnConnector) SetCapacity(val interface{}) {
	_jsii_.Set(
		j,
		"capacity",
		val,
	)
}

func (j *jsiiProxy_CfnConnector) SetConnectorConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"connectorConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnConnector) SetConnectorDescription(val *string) {
	_jsii_.Set(
		j,
		"connectorDescription",
		val,
	)
}

func (j *jsiiProxy_CfnConnector) SetConnectorName(val *string) {
	_jsii_.Set(
		j,
		"connectorName",
		val,
	)
}

func (j *jsiiProxy_CfnConnector) SetKafkaCluster(val interface{}) {
	_jsii_.Set(
		j,
		"kafkaCluster",
		val,
	)
}

func (j *jsiiProxy_CfnConnector) SetKafkaClusterClientAuthentication(val interface{}) {
	_jsii_.Set(
		j,
		"kafkaClusterClientAuthentication",
		val,
	)
}

func (j *jsiiProxy_CfnConnector) SetKafkaClusterEncryptionInTransit(val interface{}) {
	_jsii_.Set(
		j,
		"kafkaClusterEncryptionInTransit",
		val,
	)
}

func (j *jsiiProxy_CfnConnector) SetKafkaConnectVersion(val *string) {
	_jsii_.Set(
		j,
		"kafkaConnectVersion",
		val,
	)
}

func (j *jsiiProxy_CfnConnector) SetLogDelivery(val interface{}) {
	_jsii_.Set(
		j,
		"logDelivery",
		val,
	)
}

func (j *jsiiProxy_CfnConnector) SetPlugins(val interface{}) {
	_jsii_.Set(
		j,
		"plugins",
		val,
	)
}

func (j *jsiiProxy_CfnConnector) SetServiceExecutionRoleArn(val *string) {
	_jsii_.Set(
		j,
		"serviceExecutionRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnConnector) SetWorkerConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"workerConfiguration",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnConnector_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kafkaconnect.CfnConnector",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnConnector_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kafkaconnect.CfnConnector",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnConnector_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kafkaconnect.CfnConnector",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnConnector_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_kafkaconnect.CfnConnector",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnConnector) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnConnector) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnConnector) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

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
func (c *jsiiProxy_CfnConnector) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnConnector) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnConnector) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnConnector) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnConnector) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnConnector) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnConnector) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnConnector) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnConnector) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnConnector) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnConnector) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConnector) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// TODO: EXAMPLE
//
type CfnConnector_ApacheKafkaClusterProperty struct {
	// `CfnConnector.ApacheKafkaClusterProperty.BootstrapServers`.
	BootstrapServers *string `json:"bootstrapServers" yaml:"bootstrapServers"`
	// `CfnConnector.ApacheKafkaClusterProperty.Vpc`.
	Vpc interface{} `json:"vpc" yaml:"vpc"`
}

// TODO: EXAMPLE
//
type CfnConnector_AutoScalingProperty struct {
	// `CfnConnector.AutoScalingProperty.MaxWorkerCount`.
	MaxWorkerCount *float64 `json:"maxWorkerCount" yaml:"maxWorkerCount"`
	// `CfnConnector.AutoScalingProperty.McuCount`.
	McuCount *float64 `json:"mcuCount" yaml:"mcuCount"`
	// `CfnConnector.AutoScalingProperty.MinWorkerCount`.
	MinWorkerCount *float64 `json:"minWorkerCount" yaml:"minWorkerCount"`
	// `CfnConnector.AutoScalingProperty.ScaleInPolicy`.
	ScaleInPolicy interface{} `json:"scaleInPolicy" yaml:"scaleInPolicy"`
	// `CfnConnector.AutoScalingProperty.ScaleOutPolicy`.
	ScaleOutPolicy interface{} `json:"scaleOutPolicy" yaml:"scaleOutPolicy"`
}

// TODO: EXAMPLE
//
type CfnConnector_CapacityProperty struct {
	// `CfnConnector.CapacityProperty.AutoScaling`.
	AutoScaling interface{} `json:"autoScaling" yaml:"autoScaling"`
	// `CfnConnector.CapacityProperty.ProvisionedCapacity`.
	ProvisionedCapacity interface{} `json:"provisionedCapacity" yaml:"provisionedCapacity"`
}

// TODO: EXAMPLE
//
type CfnConnector_CloudWatchLogsLogDeliveryProperty struct {
	// `CfnConnector.CloudWatchLogsLogDeliveryProperty.Enabled`.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
	// `CfnConnector.CloudWatchLogsLogDeliveryProperty.LogGroup`.
	LogGroup *string `json:"logGroup" yaml:"logGroup"`
}

// TODO: EXAMPLE
//
type CfnConnector_CustomPluginProperty struct {
	// `CfnConnector.CustomPluginProperty.CustomPluginArn`.
	CustomPluginArn *string `json:"customPluginArn" yaml:"customPluginArn"`
	// `CfnConnector.CustomPluginProperty.Revision`.
	Revision *float64 `json:"revision" yaml:"revision"`
}

// TODO: EXAMPLE
//
type CfnConnector_FirehoseLogDeliveryProperty struct {
	// `CfnConnector.FirehoseLogDeliveryProperty.Enabled`.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
	// `CfnConnector.FirehoseLogDeliveryProperty.DeliveryStream`.
	DeliveryStream *string `json:"deliveryStream" yaml:"deliveryStream"`
}

// TODO: EXAMPLE
//
type CfnConnector_KafkaClusterClientAuthenticationProperty struct {
	// `CfnConnector.KafkaClusterClientAuthenticationProperty.AuthenticationType`.
	AuthenticationType *string `json:"authenticationType" yaml:"authenticationType"`
}

// TODO: EXAMPLE
//
type CfnConnector_KafkaClusterEncryptionInTransitProperty struct {
	// `CfnConnector.KafkaClusterEncryptionInTransitProperty.EncryptionType`.
	EncryptionType *string `json:"encryptionType" yaml:"encryptionType"`
}

// TODO: EXAMPLE
//
type CfnConnector_KafkaClusterProperty struct {
	// `CfnConnector.KafkaClusterProperty.ApacheKafkaCluster`.
	ApacheKafkaCluster interface{} `json:"apacheKafkaCluster" yaml:"apacheKafkaCluster"`
}

// TODO: EXAMPLE
//
type CfnConnector_LogDeliveryProperty struct {
	// `CfnConnector.LogDeliveryProperty.WorkerLogDelivery`.
	WorkerLogDelivery interface{} `json:"workerLogDelivery" yaml:"workerLogDelivery"`
}

// TODO: EXAMPLE
//
type CfnConnector_PluginProperty struct {
	// `CfnConnector.PluginProperty.CustomPlugin`.
	CustomPlugin interface{} `json:"customPlugin" yaml:"customPlugin"`
}

// TODO: EXAMPLE
//
type CfnConnector_ProvisionedCapacityProperty struct {
	// `CfnConnector.ProvisionedCapacityProperty.WorkerCount`.
	WorkerCount *float64 `json:"workerCount" yaml:"workerCount"`
	// `CfnConnector.ProvisionedCapacityProperty.McuCount`.
	McuCount *float64 `json:"mcuCount" yaml:"mcuCount"`
}

// TODO: EXAMPLE
//
type CfnConnector_S3LogDeliveryProperty struct {
	// `CfnConnector.S3LogDeliveryProperty.Enabled`.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
	// `CfnConnector.S3LogDeliveryProperty.Bucket`.
	Bucket *string `json:"bucket" yaml:"bucket"`
	// `CfnConnector.S3LogDeliveryProperty.Prefix`.
	Prefix *string `json:"prefix" yaml:"prefix"`
}

// TODO: EXAMPLE
//
type CfnConnector_ScaleInPolicyProperty struct {
	// `CfnConnector.ScaleInPolicyProperty.CpuUtilizationPercentage`.
	CpuUtilizationPercentage *float64 `json:"cpuUtilizationPercentage" yaml:"cpuUtilizationPercentage"`
}

// TODO: EXAMPLE
//
type CfnConnector_ScaleOutPolicyProperty struct {
	// `CfnConnector.ScaleOutPolicyProperty.CpuUtilizationPercentage`.
	CpuUtilizationPercentage *float64 `json:"cpuUtilizationPercentage" yaml:"cpuUtilizationPercentage"`
}

// TODO: EXAMPLE
//
type CfnConnector_VpcProperty struct {
	// `CfnConnector.VpcProperty.SecurityGroups`.
	SecurityGroups *[]*string `json:"securityGroups" yaml:"securityGroups"`
	// `CfnConnector.VpcProperty.Subnets`.
	Subnets *[]*string `json:"subnets" yaml:"subnets"`
}

// TODO: EXAMPLE
//
type CfnConnector_WorkerConfigurationProperty struct {
	// `CfnConnector.WorkerConfigurationProperty.Revision`.
	Revision *float64 `json:"revision" yaml:"revision"`
	// `CfnConnector.WorkerConfigurationProperty.WorkerConfigurationArn`.
	WorkerConfigurationArn *string `json:"workerConfigurationArn" yaml:"workerConfigurationArn"`
}

// TODO: EXAMPLE
//
type CfnConnector_WorkerLogDeliveryProperty struct {
	// `CfnConnector.WorkerLogDeliveryProperty.CloudWatchLogs`.
	CloudWatchLogs interface{} `json:"cloudWatchLogs" yaml:"cloudWatchLogs"`
	// `CfnConnector.WorkerLogDeliveryProperty.Firehose`.
	Firehose interface{} `json:"firehose" yaml:"firehose"`
	// `CfnConnector.WorkerLogDeliveryProperty.S3`.
	S3 interface{} `json:"s3" yaml:"s3"`
}

// Properties for defining a `CfnConnector`.
//
// TODO: EXAMPLE
//
type CfnConnectorProps struct {
	// `AWS::KafkaConnect::Connector.Capacity`.
	Capacity interface{} `json:"capacity" yaml:"capacity"`
	// `AWS::KafkaConnect::Connector.ConnectorConfiguration`.
	ConnectorConfiguration interface{} `json:"connectorConfiguration" yaml:"connectorConfiguration"`
	// `AWS::KafkaConnect::Connector.ConnectorName`.
	ConnectorName *string `json:"connectorName" yaml:"connectorName"`
	// `AWS::KafkaConnect::Connector.KafkaCluster`.
	KafkaCluster interface{} `json:"kafkaCluster" yaml:"kafkaCluster"`
	// `AWS::KafkaConnect::Connector.KafkaClusterClientAuthentication`.
	KafkaClusterClientAuthentication interface{} `json:"kafkaClusterClientAuthentication" yaml:"kafkaClusterClientAuthentication"`
	// `AWS::KafkaConnect::Connector.KafkaClusterEncryptionInTransit`.
	KafkaClusterEncryptionInTransit interface{} `json:"kafkaClusterEncryptionInTransit" yaml:"kafkaClusterEncryptionInTransit"`
	// `AWS::KafkaConnect::Connector.KafkaConnectVersion`.
	KafkaConnectVersion *string `json:"kafkaConnectVersion" yaml:"kafkaConnectVersion"`
	// `AWS::KafkaConnect::Connector.Plugins`.
	Plugins interface{} `json:"plugins" yaml:"plugins"`
	// `AWS::KafkaConnect::Connector.ServiceExecutionRoleArn`.
	ServiceExecutionRoleArn *string `json:"serviceExecutionRoleArn" yaml:"serviceExecutionRoleArn"`
	// `AWS::KafkaConnect::Connector.ConnectorDescription`.
	ConnectorDescription *string `json:"connectorDescription" yaml:"connectorDescription"`
	// `AWS::KafkaConnect::Connector.LogDelivery`.
	LogDelivery interface{} `json:"logDelivery" yaml:"logDelivery"`
	// `AWS::KafkaConnect::Connector.WorkerConfiguration`.
	WorkerConfiguration interface{} `json:"workerConfiguration" yaml:"workerConfiguration"`
}

