package awsmsk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmsk/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::MSK::Cluster`.
//
// The `AWS::MSK::Cluster` resource creates an Amazon MSK cluster . For more information, see [What Is Amazon MSK?](https://docs.aws.amazon.com/msk/latest/developerguide/what-is-msk.html) in the *Amazon MSK Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCluster := awscdk.Aws_msk.NewCfnCluster(this, jsii.String("MyCfnCluster"), &CfnClusterProps{
//   	BrokerNodeGroupInfo: &BrokerNodeGroupInfoProperty{
//   		ClientSubnets: []*string{
//   			jsii.String("clientSubnets"),
//   		},
//   		InstanceType: jsii.String("instanceType"),
//
//   		// the properties below are optional
//   		BrokerAzDistribution: jsii.String("brokerAzDistribution"),
//   		ConnectivityInfo: &ConnectivityInfoProperty{
//   			PublicAccess: &PublicAccessProperty{
//   				Type: jsii.String("type"),
//   			},
//   			VpcConnectivity: &VpcConnectivityProperty{
//   				ClientAuthentication: &VpcConnectivityClientAuthenticationProperty{
//   					Sasl: &VpcConnectivitySaslProperty{
//   						Iam: &VpcConnectivityIamProperty{
//   							Enabled: jsii.Boolean(false),
//   						},
//   						Scram: &VpcConnectivityScramProperty{
//   							Enabled: jsii.Boolean(false),
//   						},
//   					},
//   					Tls: &VpcConnectivityTlsProperty{
//   						Enabled: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   		},
//   		SecurityGroups: []*string{
//   			jsii.String("securityGroups"),
//   		},
//   		StorageInfo: &StorageInfoProperty{
//   			EbsStorageInfo: &EBSStorageInfoProperty{
//   				ProvisionedThroughput: &ProvisionedThroughputProperty{
//   					Enabled: jsii.Boolean(false),
//   					VolumeThroughput: jsii.Number(123),
//   				},
//   				VolumeSize: jsii.Number(123),
//   			},
//   		},
//   	},
//   	ClusterName: jsii.String("clusterName"),
//   	KafkaVersion: jsii.String("kafkaVersion"),
//   	NumberOfBrokerNodes: jsii.Number(123),
//
//   	// the properties below are optional
//   	ClientAuthentication: &ClientAuthenticationProperty{
//   		Sasl: &SaslProperty{
//   			Iam: &IamProperty{
//   				Enabled: jsii.Boolean(false),
//   			},
//   			Scram: &ScramProperty{
//   				Enabled: jsii.Boolean(false),
//   			},
//   		},
//   		Tls: &TlsProperty{
//   			CertificateAuthorityArnList: []*string{
//   				jsii.String("certificateAuthorityArnList"),
//   			},
//   			Enabled: jsii.Boolean(false),
//   		},
//   		Unauthenticated: &UnauthenticatedProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   	},
//   	ConfigurationInfo: &ConfigurationInfoProperty{
//   		Arn: jsii.String("arn"),
//   		Revision: jsii.Number(123),
//   	},
//   	CurrentVersion: jsii.String("currentVersion"),
//   	EncryptionInfo: &EncryptionInfoProperty{
//   		EncryptionAtRest: &EncryptionAtRestProperty{
//   			DataVolumeKmsKeyId: jsii.String("dataVolumeKmsKeyId"),
//   		},
//   		EncryptionInTransit: &EncryptionInTransitProperty{
//   			ClientBroker: jsii.String("clientBroker"),
//   			InCluster: jsii.Boolean(false),
//   		},
//   	},
//   	EnhancedMonitoring: jsii.String("enhancedMonitoring"),
//   	LoggingInfo: &LoggingInfoProperty{
//   		BrokerLogs: &BrokerLogsProperty{
//   			CloudWatchLogs: &CloudWatchLogsProperty{
//   				Enabled: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				LogGroup: jsii.String("logGroup"),
//   			},
//   			Firehose: &FirehoseProperty{
//   				Enabled: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				DeliveryStream: jsii.String("deliveryStream"),
//   			},
//   			S3: &S3Property{
//   				Enabled: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				Bucket: jsii.String("bucket"),
//   				Prefix: jsii.String("prefix"),
//   			},
//   		},
//   	},
//   	OpenMonitoring: &OpenMonitoringProperty{
//   		Prometheus: &PrometheusProperty{
//   			JmxExporter: &JmxExporterProperty{
//   				EnabledInBroker: jsii.Boolean(false),
//   			},
//   			NodeExporter: &NodeExporterProperty{
//   				EnabledInBroker: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	StorageMode: jsii.String("storageMode"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   })
//
type CfnCluster interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	// The setup to be used for brokers in the cluster.
	//
	// AWS CloudFormation may replace the cluster when you update certain `BrokerNodeGroupInfo` properties. To understand the update behavior for your use case, you should review the child properties for [`BrokerNodeGroupInfo`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-brokernodegroupinfo.html#aws-properties-msk-cluster-brokernodegroupinfo-properties) .
	BrokerNodeGroupInfo() interface{}
	SetBrokerNodeGroupInfo(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Includes information related to client authentication.
	ClientAuthentication() interface{}
	SetClientAuthentication(val interface{})
	// The name of the cluster.
	ClusterName() *string
	SetClusterName(val *string)
	// The Amazon MSK configuration to use for the cluster.
	ConfigurationInfo() interface{}
	SetConfigurationInfo(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The version of the cluster that you want to update.
	CurrentVersion() *string
	SetCurrentVersion(val *string)
	// Includes all encryption-related information.
	EncryptionInfo() interface{}
	SetEncryptionInfo(val interface{})
	// Specifies the level of monitoring for the MSK cluster.
	//
	// The possible values are `DEFAULT` , `PER_BROKER` , and `PER_TOPIC_PER_BROKER` .
	EnhancedMonitoring() *string
	SetEnhancedMonitoring(val *string)
	// The version of Apache Kafka.
	//
	// For more information, see [Supported Apache Kafka versions](https://docs.aws.amazon.com/msk/latest/developerguide/supported-kafka-versions.html) in the Amazon MSK Developer Guide.
	KafkaVersion() *string
	SetKafkaVersion(val *string)
	// You can configure your Amazon MSK cluster to send broker logs to different destination types.
	//
	// This is a container for the configuration details related to broker logs.
	LoggingInfo() interface{}
	SetLoggingInfo(val interface{})
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
	// The number of broker nodes you want in the Amazon MSK cluster.
	//
	// You can submit an update to increase the number of broker nodes in a cluster.
	NumberOfBrokerNodes() *float64
	SetNumberOfBrokerNodes(val *float64)
	// The settings for open monitoring.
	OpenMonitoring() interface{}
	SetOpenMonitoring(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// This controls storage mode for supported storage tiers.
	StorageMode() *string
	SetStorageMode(val *string)
	// A map of key:value pairs to apply to this resource.
	//
	// Both key and value are of type String.
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

// The jsii proxy struct for CfnCluster
type jsiiProxy_CfnCluster struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnCluster) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
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

func (j *jsiiProxy_CfnCluster) CurrentVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"currentVersion",
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

func (j *jsiiProxy_CfnCluster) StorageMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageMode",
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

func (j *jsiiProxy_CfnCluster) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::MSK::Cluster`.
func NewCfnCluster(scope constructs.Construct, id *string, props *CfnClusterProps) CfnCluster {
	_init_.Initialize()

	if err := validateNewCfnClusterParameters(scope, id, props); err != nil {
		panic(err)
	}
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

func (j *jsiiProxy_CfnCluster)SetBrokerNodeGroupInfo(val interface{}) {
	if err := j.validateSetBrokerNodeGroupInfoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"brokerNodeGroupInfo",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetClientAuthentication(val interface{}) {
	if err := j.validateSetClientAuthenticationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientAuthentication",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetClusterName(val *string) {
	if err := j.validateSetClusterNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterName",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetConfigurationInfo(val interface{}) {
	if err := j.validateSetConfigurationInfoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configurationInfo",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetCurrentVersion(val *string) {
	_jsii_.Set(
		j,
		"currentVersion",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetEncryptionInfo(val interface{}) {
	if err := j.validateSetEncryptionInfoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptionInfo",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetEnhancedMonitoring(val *string) {
	_jsii_.Set(
		j,
		"enhancedMonitoring",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetKafkaVersion(val *string) {
	if err := j.validateSetKafkaVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kafkaVersion",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetLoggingInfo(val interface{}) {
	if err := j.validateSetLoggingInfoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loggingInfo",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetNumberOfBrokerNodes(val *float64) {
	if err := j.validateSetNumberOfBrokerNodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numberOfBrokerNodes",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetOpenMonitoring(val interface{}) {
	if err := j.validateSetOpenMonitoringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"openMonitoring",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetStorageMode(val *string) {
	_jsii_.Set(
		j,
		"storageMode",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnCluster_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCluster_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
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
func CfnCluster_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnCluster_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
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
func CfnCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
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

func (c *jsiiProxy_CfnCluster) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnCluster) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCluster) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCluster) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnCluster) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnCluster) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnCluster) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnCluster) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnCluster) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnCluster) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnCluster) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnCluster) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCluster) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCluster) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnCluster) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCluster) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnCluster) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

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

func (c *jsiiProxy_CfnCluster) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

