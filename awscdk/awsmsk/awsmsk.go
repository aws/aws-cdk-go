package awsmsk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmsk/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::MSK::Cluster`.
type CfnCluster interface {
	awscdk.CfnResource
	awscdk.IInspectable
	BrokerNodeGroupInfo() interface{}
	SetBrokerNodeGroupInfo(val interface{})
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	ClientAuthentication() interface{}
	SetClientAuthentication(val interface{})
	ClusterName() *string
	SetClusterName(val *string)
	ConfigurationInfo() interface{}
	SetConfigurationInfo(val interface{})
	CreationStack() *[]*string
	EncryptionInfo() interface{}
	SetEncryptionInfo(val interface{})
	EnhancedMonitoring() *string
	SetEnhancedMonitoring(val *string)
	KafkaVersion() *string
	SetKafkaVersion(val *string)
	LoggingInfo() interface{}
	SetLoggingInfo(val interface{})
	LogicalId() *string
	Node() constructs.Node
	NumberOfBrokerNodes() *float64
	SetNumberOfBrokerNodes(val *float64)
	OpenMonitoring() interface{}
	SetOpenMonitoring(val interface{})
	Ref() *string
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
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

// The jsii proxy struct for CfnCluster
type jsiiProxy_CfnCluster struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnCluster) BrokerNodeGroupInfo() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"brokerNodeGroupInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) ClientAuthentication() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) ConfigurationInfo() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"configurationInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) EncryptionInfo() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) EnhancedMonitoring() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enhancedMonitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) KafkaVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kafkaVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) LoggingInfo() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loggingInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) NumberOfBrokerNodes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfBrokerNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) OpenMonitoring() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"openMonitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::MSK::Cluster`.
func NewCfnCluster(scope constructs.Construct, id *string, props *CfnClusterProps) CfnCluster {
	_init_.Initialize()

	j := jsiiProxy_CfnCluster{}

	_jsii_.Create(
		"aws-cdk-lib.aws_msk.CfnCluster",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::MSK::Cluster`.
func NewCfnCluster_Override(c CfnCluster, scope constructs.Construct, id *string, props *CfnClusterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_msk.CfnCluster",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnCluster) SetBrokerNodeGroupInfo(val interface{}) {
	_jsii_.Set(
		j,
		"brokerNodeGroupInfo",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetClientAuthentication(val interface{}) {
	_jsii_.Set(
		j,
		"clientAuthentication",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetClusterName(val *string) {
	_jsii_.Set(
		j,
		"clusterName",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetConfigurationInfo(val interface{}) {
	_jsii_.Set(
		j,
		"configurationInfo",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetEncryptionInfo(val interface{}) {
	_jsii_.Set(
		j,
		"encryptionInfo",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetEnhancedMonitoring(val *string) {
	_jsii_.Set(
		j,
		"enhancedMonitoring",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetKafkaVersion(val *string) {
	_jsii_.Set(
		j,
		"kafkaVersion",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetLoggingInfo(val interface{}) {
	_jsii_.Set(
		j,
		"loggingInfo",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetNumberOfBrokerNodes(val *float64) {
	_jsii_.Set(
		j,
		"numberOfBrokerNodes",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetOpenMonitoring(val interface{}) {
	_jsii_.Set(
		j,
		"openMonitoring",
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
func CfnCluster_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_msk.CfnCluster",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnCluster_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_msk.CfnCluster",
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
func CfnCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_msk.CfnCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCluster_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_msk.CfnCluster",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnCluster) AddDeletionOverride(path *string) {
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
// Experimental.
func (c *jsiiProxy_CfnCluster) AddDependsOn(target awscdk.CfnResource) {
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
// Experimental.
func (c *jsiiProxy_CfnCluster) AddMetadata(key *string, value interface{}) {
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
// Experimental.
func (c *jsiiProxy_CfnCluster) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnCluster) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
// Experimental.
func (c *jsiiProxy_CfnCluster) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnCluster) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
// Experimental.
func (c *jsiiProxy_CfnCluster) GetAtt(attributeName *string) awscdk.Reference {
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
// Experimental.
func (c *jsiiProxy_CfnCluster) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnCluster) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnCluster) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnCluster) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
// Experimental.
func (c *jsiiProxy_CfnCluster) ShouldSynthesize() *bool {
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
// Experimental.
func (c *jsiiProxy_CfnCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnCluster) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnCluster_BrokerLogsProperty struct {
	// `CfnCluster.BrokerLogsProperty.CloudWatchLogs`.
	CloudWatchLogs interface{} `json:"cloudWatchLogs"`
	// `CfnCluster.BrokerLogsProperty.Firehose`.
	Firehose interface{} `json:"firehose"`
	// `CfnCluster.BrokerLogsProperty.S3`.
	S3 interface{} `json:"s3"`
}

type CfnCluster_BrokerNodeGroupInfoProperty struct {
	// `CfnCluster.BrokerNodeGroupInfoProperty.ClientSubnets`.
	ClientSubnets *[]*string `json:"clientSubnets"`
	// `CfnCluster.BrokerNodeGroupInfoProperty.InstanceType`.
	InstanceType *string `json:"instanceType"`
	// `CfnCluster.BrokerNodeGroupInfoProperty.BrokerAZDistribution`.
	BrokerAzDistribution *string `json:"brokerAzDistribution"`
	// `CfnCluster.BrokerNodeGroupInfoProperty.SecurityGroups`.
	SecurityGroups *[]*string `json:"securityGroups"`
	// `CfnCluster.BrokerNodeGroupInfoProperty.StorageInfo`.
	StorageInfo interface{} `json:"storageInfo"`
}

type CfnCluster_ClientAuthenticationProperty struct {
	// `CfnCluster.ClientAuthenticationProperty.Sasl`.
	Sasl interface{} `json:"sasl"`
	// `CfnCluster.ClientAuthenticationProperty.Tls`.
	Tls interface{} `json:"tls"`
}

type CfnCluster_CloudWatchLogsProperty struct {
	// `CfnCluster.CloudWatchLogsProperty.Enabled`.
	Enabled interface{} `json:"enabled"`
	// `CfnCluster.CloudWatchLogsProperty.LogGroup`.
	LogGroup *string `json:"logGroup"`
}

type CfnCluster_ConfigurationInfoProperty struct {
	// `CfnCluster.ConfigurationInfoProperty.Arn`.
	Arn *string `json:"arn"`
	// `CfnCluster.ConfigurationInfoProperty.Revision`.
	Revision *float64 `json:"revision"`
}

type CfnCluster_EBSStorageInfoProperty struct {
	// `CfnCluster.EBSStorageInfoProperty.VolumeSize`.
	VolumeSize *float64 `json:"volumeSize"`
}

type CfnCluster_EncryptionAtRestProperty struct {
	// `CfnCluster.EncryptionAtRestProperty.DataVolumeKMSKeyId`.
	DataVolumeKmsKeyId *string `json:"dataVolumeKmsKeyId"`
}

type CfnCluster_EncryptionInTransitProperty struct {
	// `CfnCluster.EncryptionInTransitProperty.ClientBroker`.
	ClientBroker *string `json:"clientBroker"`
	// `CfnCluster.EncryptionInTransitProperty.InCluster`.
	InCluster interface{} `json:"inCluster"`
}

type CfnCluster_EncryptionInfoProperty struct {
	// `CfnCluster.EncryptionInfoProperty.EncryptionAtRest`.
	EncryptionAtRest interface{} `json:"encryptionAtRest"`
	// `CfnCluster.EncryptionInfoProperty.EncryptionInTransit`.
	EncryptionInTransit interface{} `json:"encryptionInTransit"`
}

type CfnCluster_FirehoseProperty struct {
	// `CfnCluster.FirehoseProperty.Enabled`.
	Enabled interface{} `json:"enabled"`
	// `CfnCluster.FirehoseProperty.DeliveryStream`.
	DeliveryStream *string `json:"deliveryStream"`
}

type CfnCluster_IamProperty struct {
	// `CfnCluster.IamProperty.Enabled`.
	Enabled interface{} `json:"enabled"`
}

type CfnCluster_JmxExporterProperty struct {
	// `CfnCluster.JmxExporterProperty.EnabledInBroker`.
	EnabledInBroker interface{} `json:"enabledInBroker"`
}

type CfnCluster_LoggingInfoProperty struct {
	// `CfnCluster.LoggingInfoProperty.BrokerLogs`.
	BrokerLogs interface{} `json:"brokerLogs"`
}

type CfnCluster_NodeExporterProperty struct {
	// `CfnCluster.NodeExporterProperty.EnabledInBroker`.
	EnabledInBroker interface{} `json:"enabledInBroker"`
}

type CfnCluster_OpenMonitoringProperty struct {
	// `CfnCluster.OpenMonitoringProperty.Prometheus`.
	Prometheus interface{} `json:"prometheus"`
}

type CfnCluster_PrometheusProperty struct {
	// `CfnCluster.PrometheusProperty.JmxExporter`.
	JmxExporter interface{} `json:"jmxExporter"`
	// `CfnCluster.PrometheusProperty.NodeExporter`.
	NodeExporter interface{} `json:"nodeExporter"`
}

type CfnCluster_S3Property struct {
	// `CfnCluster.S3Property.Enabled`.
	Enabled interface{} `json:"enabled"`
	// `CfnCluster.S3Property.Bucket`.
	Bucket *string `json:"bucket"`
	// `CfnCluster.S3Property.Prefix`.
	Prefix *string `json:"prefix"`
}

type CfnCluster_SaslProperty struct {
	// `CfnCluster.SaslProperty.Iam`.
	Iam interface{} `json:"iam"`
	// `CfnCluster.SaslProperty.Scram`.
	Scram interface{} `json:"scram"`
}

type CfnCluster_ScramProperty struct {
	// `CfnCluster.ScramProperty.Enabled`.
	Enabled interface{} `json:"enabled"`
}

type CfnCluster_StorageInfoProperty struct {
	// `CfnCluster.StorageInfoProperty.EBSStorageInfo`.
	EbsStorageInfo interface{} `json:"ebsStorageInfo"`
}

type CfnCluster_TlsProperty struct {
	// `CfnCluster.TlsProperty.CertificateAuthorityArnList`.
	CertificateAuthorityArnList *[]*string `json:"certificateAuthorityArnList"`
}

// Properties for defining a `AWS::MSK::Cluster`.
type CfnClusterProps struct {
	// `AWS::MSK::Cluster.BrokerNodeGroupInfo`.
	BrokerNodeGroupInfo interface{} `json:"brokerNodeGroupInfo"`
	// `AWS::MSK::Cluster.ClusterName`.
	ClusterName *string `json:"clusterName"`
	// `AWS::MSK::Cluster.KafkaVersion`.
	KafkaVersion *string `json:"kafkaVersion"`
	// `AWS::MSK::Cluster.NumberOfBrokerNodes`.
	NumberOfBrokerNodes *float64 `json:"numberOfBrokerNodes"`
	// `AWS::MSK::Cluster.ClientAuthentication`.
	ClientAuthentication interface{} `json:"clientAuthentication"`
	// `AWS::MSK::Cluster.ConfigurationInfo`.
	ConfigurationInfo interface{} `json:"configurationInfo"`
	// `AWS::MSK::Cluster.EncryptionInfo`.
	EncryptionInfo interface{} `json:"encryptionInfo"`
	// `AWS::MSK::Cluster.EnhancedMonitoring`.
	EnhancedMonitoring *string `json:"enhancedMonitoring"`
	// `AWS::MSK::Cluster.LoggingInfo`.
	LoggingInfo interface{} `json:"loggingInfo"`
	// `AWS::MSK::Cluster.OpenMonitoring`.
	OpenMonitoring interface{} `json:"openMonitoring"`
	// `AWS::MSK::Cluster.Tags`.
	Tags interface{} `json:"tags"`
}

