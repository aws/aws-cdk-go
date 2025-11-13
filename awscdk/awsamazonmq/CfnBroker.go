package awsamazonmq

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsamazonmq/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsamazonmq"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a broker. Note: This API is asynchronous.
//
// To create a broker, you must either use the `AmazonMQFullAccess` IAM policy or include the following EC2 permissions in your IAM policy.
//
// - `ec2:CreateNetworkInterface`
//
// This permission is required to allow Amazon MQ to create an elastic network interface (ENI) on behalf of your account.
// - `ec2:CreateNetworkInterfacePermission`
//
// This permission is required to attach the ENI to the broker instance.
// - `ec2:DeleteNetworkInterface`
// - `ec2:DeleteNetworkInterfacePermission`
// - `ec2:DetachNetworkInterface`
// - `ec2:DescribeInternetGateways`
// - `ec2:DescribeNetworkInterfaces`
// - `ec2:DescribeNetworkInterfacePermissions`
// - `ec2:DescribeRouteTables`
// - `ec2:DescribeSecurityGroups`
// - `ec2:DescribeSubnets`
// - `ec2:DescribeVpcs`
//
// For more information, see [Create an IAM User and Get Your AWS Credentials](https://docs.aws.amazon.com//amazon-mq/latest/developer-guide/amazon-mq-setting-up.html#create-iam-user) and [Never Modify or Delete the Amazon MQ Elastic Network Interface](https://docs.aws.amazon.com//amazon-mq/latest/developer-guide/connecting-to-amazon-mq.html#never-modify-delete-elastic-network-interface) in the *Amazon MQ Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnBroker := awscdk.Aws_amazonmq.NewCfnBroker(this, jsii.String("MyCfnBroker"), &CfnBrokerProps{
//   	BrokerName: jsii.String("brokerName"),
//   	DeploymentMode: jsii.String("deploymentMode"),
//   	EngineType: jsii.String("engineType"),
//   	HostInstanceType: jsii.String("hostInstanceType"),
//   	PubliclyAccessible: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	AuthenticationStrategy: jsii.String("authenticationStrategy"),
//   	AutoMinorVersionUpgrade: jsii.Boolean(false),
//   	Configuration: &ConfigurationIdProperty{
//   		Id: jsii.String("id"),
//   		Revision: jsii.Number(123),
//   	},
//   	DataReplicationMode: jsii.String("dataReplicationMode"),
//   	DataReplicationPrimaryBrokerArn: jsii.String("dataReplicationPrimaryBrokerArn"),
//   	EncryptionOptions: &EncryptionOptionsProperty{
//   		UseAwsOwnedKey: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	EngineVersion: jsii.String("engineVersion"),
//   	LdapServerMetadata: &LdapServerMetadataProperty{
//   		Hosts: []*string{
//   			jsii.String("hosts"),
//   		},
//   		RoleBase: jsii.String("roleBase"),
//   		RoleSearchMatching: jsii.String("roleSearchMatching"),
//   		ServiceAccountUsername: jsii.String("serviceAccountUsername"),
//   		UserBase: jsii.String("userBase"),
//   		UserSearchMatching: jsii.String("userSearchMatching"),
//
//   		// the properties below are optional
//   		RoleName: jsii.String("roleName"),
//   		RoleSearchSubtree: jsii.Boolean(false),
//   		ServiceAccountPassword: jsii.String("serviceAccountPassword"),
//   		UserRoleName: jsii.String("userRoleName"),
//   		UserSearchSubtree: jsii.Boolean(false),
//   	},
//   	Logs: &LogListProperty{
//   		Audit: jsii.Boolean(false),
//   		General: jsii.Boolean(false),
//   	},
//   	MaintenanceWindowStartTime: &MaintenanceWindowProperty{
//   		DayOfWeek: jsii.String("dayOfWeek"),
//   		TimeOfDay: jsii.String("timeOfDay"),
//   		TimeZone: jsii.String("timeZone"),
//   	},
//   	SecurityGroups: []*string{
//   		jsii.String("securityGroups"),
//   	},
//   	StorageType: jsii.String("storageType"),
//   	SubnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   	Tags: []TagsEntryProperty{
//   		&TagsEntryProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Users: []interface{}{
//   		&UserProperty{
//   			Password: jsii.String("password"),
//   			Username: jsii.String("username"),
//
//   			// the properties below are optional
//   			ConsoleAccess: jsii.Boolean(false),
//   			Groups: []*string{
//   				jsii.String("groups"),
//   			},
//   			ReplicationUser: jsii.Boolean(false),
//   		},
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html
//
type CfnBroker interface {
	awscdk.CfnResource
	awscdk.IInspectable
	interfacesawsamazonmq.IBrokerRef
	awscdk.ITaggable
	// The AMQP endpoints of each broker instance as a list of strings.
	//
	// `amqp+ssl://b-4aada85d-a80c-4be0-9d30-e344a01b921e-1.mq.eu-central-amazonaws.com:5671`
	AttrAmqpEndpoints() *[]*string
	// The Amazon Resource Name (ARN) of the Amazon MQ broker.
	//
	// `arn:aws:mq:us-east-2:123456789012:broker:MyBroker:b-1234a5b6-78cd-901e-2fgh-3i45j6k178l9`.
	AttrArn() *string
	// The unique ID that Amazon MQ generates for the configuration.
	//
	// `c-1234a5b6-78cd-901e-2fgh-3i45j6k178l9`.
	AttrConfigurationId() *string
	// The revision number of the configuration.
	//
	// `1`.
	AttrConfigurationRevision() *float64
	AttrConsoleUrLs() *[]*string
	// The version in use.
	//
	// This may have more precision than the specified EngineVersion.
	AttrEngineVersionCurrent() *string
	AttrId() *string
	// The IP addresses of each broker instance as a list of strings. Does not apply to RabbitMQ brokers.
	//
	// `['198.51.100.2', '203.0.113.9']`
	AttrIpAddresses() *[]*string
	// The MQTT endpoints of each broker instance as a list of strings.
	//
	// `mqtt+ssl://b-4aada85d-a80c-4be0-9d30-e344a01b921e-1.mq.eu-central-amazonaws.com:8883`
	AttrMqttEndpoints() *[]*string
	// The OpenWire endpoints of each broker instance as a list of strings.
	//
	// `ssl://b-4aada85d-a80c-4be0-9d30-e344a01b921e-1.mq.eu-central-amazonaws.com:61617`
	AttrOpenWireEndpoints() *[]*string
	// The STOMP endpoints of each broker instance as a list of strings.
	//
	// `stomp+ssl://b-4aada85d-a80c-4be0-9d30-e344a01b921e-1.mq.eu-central-amazonaws.com:61614`
	AttrStompEndpoints() *[]*string
	// The WSS endpoints of each broker instance as a list of strings.
	//
	// `wss://b-4aada85d-a80c-4be0-9d30-e344a01b921e-1.mq.eu-central-amazonaws.com:61619`
	AttrWssEndpoints() *[]*string
	// Optional.
	AuthenticationStrategy() *string
	SetAuthenticationStrategy(val *string)
	// Enables automatic upgrades to new patch versions for brokers as new versions are released and supported by Amazon MQ.
	AutoMinorVersionUpgrade() interface{}
	SetAutoMinorVersionUpgrade(val interface{})
	// Required.
	BrokerName() *string
	SetBrokerName(val *string)
	// A reference to a Broker resource.
	BrokerRef() *interfacesawsamazonmq.BrokerReference
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// A list of information about the configuration.
	Configuration() interface{}
	SetConfiguration(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Defines whether this broker is a part of a data replication pair.
	DataReplicationMode() *string
	SetDataReplicationMode(val *string)
	// The Amazon Resource Name (ARN) of the primary broker that is used to replicate data from in a data replication pair, and is applied to the replica broker.
	DataReplicationPrimaryBrokerArn() *string
	SetDataReplicationPrimaryBrokerArn(val *string)
	// Required.
	DeploymentMode() *string
	SetDeploymentMode(val *string)
	// Encryption options for the broker.
	EncryptionOptions() interface{}
	SetEncryptionOptions(val interface{})
	// Required.
	EngineType() *string
	SetEngineType(val *string)
	// The broker engine version.
	EngineVersion() *string
	SetEngineVersion(val *string)
	Env() *interfaces.ResourceEnvironment
	// Required.
	HostInstanceType() *string
	SetHostInstanceType(val *string)
	// Optional.
	LdapServerMetadata() interface{}
	SetLdapServerMetadata(val interface{})
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
	// Enables Amazon CloudWatch logging for brokers.
	Logs() interface{}
	SetLogs(val interface{})
	// The parameters that determine the WeeklyStartTime.
	MaintenanceWindowStartTime() interface{}
	SetMaintenanceWindowStartTime(val interface{})
	// The tree node.
	Node() constructs.Node
	// Enables connections from applications outside of the VPC that hosts the broker's subnets.
	PubliclyAccessible() interface{}
	SetPubliclyAccessible(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The list of rules (1 minimum, 125 maximum) that authorize connections to brokers.
	SecurityGroups() *[]*string
	SetSecurityGroups(val *[]*string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The broker's storage type.
	StorageType() *string
	SetStorageType(val *string)
	// The list of groups that define which subnets and IP ranges the broker can use from different Availability Zones.
	SubnetIds() *[]*string
	SetSubnetIds(val *[]*string)
	// Tag Manager which manages the tags for this resource.
	Tags() awscdk.TagManager
	// Create tags when creating the broker.
	TagsRaw() *[]*CfnBroker_TagsEntryProperty
	SetTagsRaw(val *[]*CfnBroker_TagsEntryProperty)
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
	// The list of broker users (persons or applications) who can access queues and topics.
	Users() interface{}
	SetUsers(val interface{})
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

// The jsii proxy struct for CfnBroker
type jsiiProxy_CfnBroker struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__interfacesawsamazonmqIBrokerRef
	internal.Type__awscdkITaggable
}

func (j *jsiiProxy_CfnBroker) AttrAmqpEndpoints() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrAmqpEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBroker) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBroker) AttrConfigurationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrConfigurationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBroker) AttrConfigurationRevision() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrConfigurationRevision",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBroker) AttrConsoleUrLs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrConsoleUrLs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBroker) AttrEngineVersionCurrent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrEngineVersionCurrent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBroker) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBroker) AttrIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBroker) AttrMqttEndpoints() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrMqttEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBroker) AttrOpenWireEndpoints() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrOpenWireEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBroker) AttrStompEndpoints() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrStompEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBroker) AttrWssEndpoints() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrWssEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBroker) AuthenticationStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBroker) AutoMinorVersionUpgrade() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoMinorVersionUpgrade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBroker) BrokerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"brokerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBroker) BrokerRef() *interfacesawsamazonmq.BrokerReference {
	var returns *interfacesawsamazonmq.BrokerReference
	_jsii_.Get(
		j,
		"brokerRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBroker) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBroker) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBroker) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBroker) Configuration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBroker) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBroker) DataReplicationMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataReplicationMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBroker) DataReplicationPrimaryBrokerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataReplicationPrimaryBrokerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBroker) DeploymentMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBroker) EncryptionOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBroker) EngineType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBroker) EngineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBroker) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBroker) HostInstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostInstanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBroker) LdapServerMetadata() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ldapServerMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBroker) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBroker) Logs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBroker) MaintenanceWindowStartTime() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"maintenanceWindowStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBroker) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBroker) PubliclyAccessible() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publiclyAccessible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBroker) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBroker) SecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBroker) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBroker) StorageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBroker) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBroker) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBroker) TagsRaw() *[]*CfnBroker_TagsEntryProperty {
	var returns *[]*CfnBroker_TagsEntryProperty
	_jsii_.Get(
		j,
		"tagsRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBroker) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBroker) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBroker) Users() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"users",
		&returns,
	)
	return returns
}


// Create a new `AWS::AmazonMQ::Broker`.
func NewCfnBroker(scope constructs.Construct, id *string, props *CfnBrokerProps) CfnBroker {
	_init_.Initialize()

	if err := validateNewCfnBrokerParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnBroker{}

	_jsii_.Create(
		"aws-cdk-lib.aws_amazonmq.CfnBroker",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AmazonMQ::Broker`.
func NewCfnBroker_Override(c CfnBroker, scope constructs.Construct, id *string, props *CfnBrokerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_amazonmq.CfnBroker",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnBroker)SetAuthenticationStrategy(val *string) {
	_jsii_.Set(
		j,
		"authenticationStrategy",
		val,
	)
}

func (j *jsiiProxy_CfnBroker)SetAutoMinorVersionUpgrade(val interface{}) {
	if err := j.validateSetAutoMinorVersionUpgradeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoMinorVersionUpgrade",
		val,
	)
}

func (j *jsiiProxy_CfnBroker)SetBrokerName(val *string) {
	if err := j.validateSetBrokerNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"brokerName",
		val,
	)
}

func (j *jsiiProxy_CfnBroker)SetConfiguration(val interface{}) {
	if err := j.validateSetConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configuration",
		val,
	)
}

func (j *jsiiProxy_CfnBroker)SetDataReplicationMode(val *string) {
	_jsii_.Set(
		j,
		"dataReplicationMode",
		val,
	)
}

func (j *jsiiProxy_CfnBroker)SetDataReplicationPrimaryBrokerArn(val *string) {
	_jsii_.Set(
		j,
		"dataReplicationPrimaryBrokerArn",
		val,
	)
}

func (j *jsiiProxy_CfnBroker)SetDeploymentMode(val *string) {
	if err := j.validateSetDeploymentModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deploymentMode",
		val,
	)
}

func (j *jsiiProxy_CfnBroker)SetEncryptionOptions(val interface{}) {
	if err := j.validateSetEncryptionOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptionOptions",
		val,
	)
}

func (j *jsiiProxy_CfnBroker)SetEngineType(val *string) {
	if err := j.validateSetEngineTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engineType",
		val,
	)
}

func (j *jsiiProxy_CfnBroker)SetEngineVersion(val *string) {
	_jsii_.Set(
		j,
		"engineVersion",
		val,
	)
}

func (j *jsiiProxy_CfnBroker)SetHostInstanceType(val *string) {
	if err := j.validateSetHostInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostInstanceType",
		val,
	)
}

func (j *jsiiProxy_CfnBroker)SetLdapServerMetadata(val interface{}) {
	if err := j.validateSetLdapServerMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ldapServerMetadata",
		val,
	)
}

func (j *jsiiProxy_CfnBroker)SetLogs(val interface{}) {
	if err := j.validateSetLogsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logs",
		val,
	)
}

func (j *jsiiProxy_CfnBroker)SetMaintenanceWindowStartTime(val interface{}) {
	if err := j.validateSetMaintenanceWindowStartTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maintenanceWindowStartTime",
		val,
	)
}

func (j *jsiiProxy_CfnBroker)SetPubliclyAccessible(val interface{}) {
	if err := j.validateSetPubliclyAccessibleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publiclyAccessible",
		val,
	)
}

func (j *jsiiProxy_CfnBroker)SetSecurityGroups(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroups",
		val,
	)
}

func (j *jsiiProxy_CfnBroker)SetStorageType(val *string) {
	_jsii_.Set(
		j,
		"storageType",
		val,
	)
}

func (j *jsiiProxy_CfnBroker)SetSubnetIds(val *[]*string) {
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_CfnBroker)SetTagsRaw(val *[]*CfnBroker_TagsEntryProperty) {
	if err := j.validateSetTagsRawParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsRaw",
		val,
	)
}

func (j *jsiiProxy_CfnBroker)SetUsers(val interface{}) {
	if err := j.validateSetUsersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"users",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnBroker_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnBroker_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_amazonmq.CfnBroker",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnBroker_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnBroker_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_amazonmq.CfnBroker",
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
func CfnBroker_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnBroker_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_amazonmq.CfnBroker",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnBroker_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_amazonmq.CfnBroker",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnBroker) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnBroker) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnBroker) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnBroker) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnBroker) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnBroker) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnBroker) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnBroker) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnBroker) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnBroker) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnBroker) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnBroker) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBroker) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBroker) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnBroker) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnBroker) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnBroker) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnBroker) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBroker) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBroker) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

