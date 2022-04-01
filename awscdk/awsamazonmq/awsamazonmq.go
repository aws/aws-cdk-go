package awsamazonmq

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsamazonmq/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::AmazonMQ::Broker`.
//
// A *broker* is a message broker environment running on Amazon MQ . It is the basic building block of Amazon MQ .
//
// The `AWS::AmazonMQ::Broker` resource lets you create Amazon MQ for ActiveMQ and Amazon MQ for RabbitMQ brokers, add configuration changes or modify users for a speified ActiveMQ broker, return information about the specified broker, and delete the broker. For more information, see [How Amazon MQ works](https://docs.aws.amazon.com//amazon-mq/latest/developer-guide/amazon-mq-how-it-works.html) in the *Amazon MQ Developer Guide* .
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
// - `ec2:DescribeVpcs`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import amazonmq "github.com/aws/aws-cdk-go/awscdk/aws_amazonmq"
//   cfnBroker := amazonmq.NewCfnBroker(this, jsii.String("MyCfnBroker"), &cfnBrokerProps{
//   	autoMinorVersionUpgrade: jsii.Boolean(false),
//   	brokerName: jsii.String("brokerName"),
//   	deploymentMode: jsii.String("deploymentMode"),
//   	engineType: jsii.String("engineType"),
//   	engineVersion: jsii.String("engineVersion"),
//   	hostInstanceType: jsii.String("hostInstanceType"),
//   	publiclyAccessible: jsii.Boolean(false),
//   	users: []interface{}{
//   		&userProperty{
//   			password: jsii.String("password"),
//   			username: jsii.String("username"),
//
//   			// the properties below are optional
//   			consoleAccess: jsii.Boolean(false),
//   			groups: []*string{
//   				jsii.String("groups"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	authenticationStrategy: jsii.String("authenticationStrategy"),
//   	configuration: &configurationIdProperty{
//   		id: jsii.String("id"),
//   		revision: jsii.Number(123),
//   	},
//   	encryptionOptions: &encryptionOptionsProperty{
//   		useAwsOwnedKey: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		kmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	ldapServerMetadata: &ldapServerMetadataProperty{
//   		hosts: []*string{
//   			jsii.String("hosts"),
//   		},
//   		roleBase: jsii.String("roleBase"),
//   		roleSearchMatching: jsii.String("roleSearchMatching"),
//   		serviceAccountPassword: jsii.String("serviceAccountPassword"),
//   		serviceAccountUsername: jsii.String("serviceAccountUsername"),
//   		userBase: jsii.String("userBase"),
//   		userSearchMatching: jsii.String("userSearchMatching"),
//
//   		// the properties below are optional
//   		roleName: jsii.String("roleName"),
//   		roleSearchSubtree: jsii.Boolean(false),
//   		userRoleName: jsii.String("userRoleName"),
//   		userSearchSubtree: jsii.Boolean(false),
//   	},
//   	logs: &logListProperty{
//   		audit: jsii.Boolean(false),
//   		general: jsii.Boolean(false),
//   	},
//   	maintenanceWindowStartTime: &maintenanceWindowProperty{
//   		dayOfWeek: jsii.String("dayOfWeek"),
//   		timeOfDay: jsii.String("timeOfDay"),
//   		timeZone: jsii.String("timeZone"),
//   	},
//   	securityGroups: []*string{
//   		jsii.String("securityGroups"),
//   	},
//   	storageType: jsii.String("storageType"),
//   	subnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   	tags: []tagsEntryProperty{
//   		&tagsEntryProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnBroker interface {
	awscdk.CfnResource
	awscdk.IInspectable
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
	//
	// The authentication strategy used to secure the broker. The default is `SIMPLE` .
	AuthenticationStrategy() *string
	SetAuthenticationStrategy(val *string)
	// Enables automatic upgrades to new minor versions for brokers, as new broker engine versions are released and supported by Amazon MQ.
	//
	// Automatic upgrades occur during the scheduled maintenance window of the broker or after a manual broker reboot.
	AutoMinorVersionUpgrade() interface{}
	SetAutoMinorVersionUpgrade(val interface{})
	// The name of the broker.
	//
	// This value must be unique in your AWS account , 1-50 characters long, must contain only letters, numbers, dashes, and underscores, and must not contain white spaces, brackets, wildcard characters, or special characters.
	//
	// > Do not add personally identifiable information (PII) or other confidential or sensitive information in broker names. Broker names are accessible to other AWS services, including C CloudWatch Logs . Broker names are not intended to be used for private or sensitive data.
	BrokerName() *string
	SetBrokerName(val *string)
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// A list of information about the configuration.
	//
	// Does not apply to RabbitMQ brokers.
	Configuration() interface{}
	SetConfiguration(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// The deployment mode of the broker. Available values:.
	//
	// - `SINGLE_INSTANCE`
	// - `ACTIVE_STANDBY_MULTI_AZ`
	// - `CLUSTER_MULTI_AZ`.
	DeploymentMode() *string
	SetDeploymentMode(val *string)
	// Encryption options for the broker.
	//
	// Does not apply to RabbitMQ brokers.
	EncryptionOptions() interface{}
	SetEncryptionOptions(val interface{})
	// The type of broker engine.
	//
	// Currently, Amazon MQ supports `ACTIVEMQ` and `RABBITMQ` .
	EngineType() *string
	SetEngineType(val *string)
	// The version of the broker engine.
	//
	// For a list of supported engine versions, see [Engine](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/broker-engine.html) in the *Amazon MQ Developer Guide* .
	EngineVersion() *string
	SetEngineVersion(val *string)
	// The broker's instance type.
	HostInstanceType() *string
	SetHostInstanceType(val *string)
	// Optional.
	//
	// The metadata of the LDAP server used to authenticate and authorize connections to the broker. Does not apply to RabbitMQ brokers.
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
	// Experimental.
	LogicalId() *string
	// Enables Amazon CloudWatch logging for brokers.
	Logs() interface{}
	SetLogs(val interface{})
	// The scheduled time period relative to UTC during which Amazon MQ begins to apply pending updates or patches to the broker.
	MaintenanceWindowStartTime() interface{}
	SetMaintenanceWindowStartTime(val interface{})
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Enables connections from applications outside of the VPC that hosts the broker's subnets.
	PubliclyAccessible() interface{}
	SetPubliclyAccessible(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The list of rules (1 minimum, 125 maximum) that authorize connections to brokers.
	SecurityGroups() *[]*string
	SetSecurityGroups(val *[]*string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// The broker's storage type.
	StorageType() *string
	SetStorageType(val *string)
	// The list of groups that define which subnets and IP ranges the broker can use from different Availability Zones.
	//
	// If you specify more than one subnet, the subnets must be in different Availability Zones. Amazon MQ will not be able to create VPC endpoints for your broker with multiple subnets in the same Availability Zone. A SINGLE_INSTANCE deployment requires one subnet (for example, the default subnet). An ACTIVE_STANDBY_MULTI_AZ deployment (ACTIVEMQ) requires two subnets. A CLUSTER_MULTI_AZ deployment (RABBITMQ) has no subnet requirements when deployed with public accessibility, deployment without public accessibility requires at least one subnet.
	//
	// > If you specify subnets in a shared VPC for a RabbitMQ broker, the associated VPC to which the specified subnets belong must be owned by your AWS account . Amazon MQ will not be able to create VPC enpoints in VPCs that are not owned by your AWS account .
	SubnetIds() *[]*string
	SetSubnetIds(val *[]*string)
	// An array of key-value pairs.
	//
	// For more information, see [Using Cost Allocation Tags](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html) in the *Billing and Cost Management User Guide* .
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// The list of broker users (persons or applications) who can access queues and topics.
	//
	// For Amazon MQ for RabbitMQ brokers, one and only one administrative user is accepted and created when a broker is first provisioned. All subsequent RabbitMQ users are created by via the RabbitMQ web console or by using the RabbitMQ management API.
	Users() interface{}
	SetUsers(val interface{})
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

// The jsii proxy struct for CfnBroker
type jsiiProxy_CfnBroker struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
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

func (j *jsiiProxy_CfnBroker) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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

func (j *jsiiProxy_CfnBroker) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
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
func NewCfnBroker(scope awscdk.Construct, id *string, props *CfnBrokerProps) CfnBroker {
	_init_.Initialize()

	j := jsiiProxy_CfnBroker{}

	_jsii_.Create(
		"monocdk.aws_amazonmq.CfnBroker",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AmazonMQ::Broker`.
func NewCfnBroker_Override(c CfnBroker, scope awscdk.Construct, id *string, props *CfnBrokerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_amazonmq.CfnBroker",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnBroker) SetAuthenticationStrategy(val *string) {
	_jsii_.Set(
		j,
		"authenticationStrategy",
		val,
	)
}

func (j *jsiiProxy_CfnBroker) SetAutoMinorVersionUpgrade(val interface{}) {
	_jsii_.Set(
		j,
		"autoMinorVersionUpgrade",
		val,
	)
}

func (j *jsiiProxy_CfnBroker) SetBrokerName(val *string) {
	_jsii_.Set(
		j,
		"brokerName",
		val,
	)
}

func (j *jsiiProxy_CfnBroker) SetConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"configuration",
		val,
	)
}

func (j *jsiiProxy_CfnBroker) SetDeploymentMode(val *string) {
	_jsii_.Set(
		j,
		"deploymentMode",
		val,
	)
}

func (j *jsiiProxy_CfnBroker) SetEncryptionOptions(val interface{}) {
	_jsii_.Set(
		j,
		"encryptionOptions",
		val,
	)
}

func (j *jsiiProxy_CfnBroker) SetEngineType(val *string) {
	_jsii_.Set(
		j,
		"engineType",
		val,
	)
}

func (j *jsiiProxy_CfnBroker) SetEngineVersion(val *string) {
	_jsii_.Set(
		j,
		"engineVersion",
		val,
	)
}

func (j *jsiiProxy_CfnBroker) SetHostInstanceType(val *string) {
	_jsii_.Set(
		j,
		"hostInstanceType",
		val,
	)
}

func (j *jsiiProxy_CfnBroker) SetLdapServerMetadata(val interface{}) {
	_jsii_.Set(
		j,
		"ldapServerMetadata",
		val,
	)
}

func (j *jsiiProxy_CfnBroker) SetLogs(val interface{}) {
	_jsii_.Set(
		j,
		"logs",
		val,
	)
}

func (j *jsiiProxy_CfnBroker) SetMaintenanceWindowStartTime(val interface{}) {
	_jsii_.Set(
		j,
		"maintenanceWindowStartTime",
		val,
	)
}

func (j *jsiiProxy_CfnBroker) SetPubliclyAccessible(val interface{}) {
	_jsii_.Set(
		j,
		"publiclyAccessible",
		val,
	)
}

func (j *jsiiProxy_CfnBroker) SetSecurityGroups(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroups",
		val,
	)
}

func (j *jsiiProxy_CfnBroker) SetStorageType(val *string) {
	_jsii_.Set(
		j,
		"storageType",
		val,
	)
}

func (j *jsiiProxy_CfnBroker) SetSubnetIds(val *[]*string) {
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_CfnBroker) SetUsers(val interface{}) {
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
// Experimental.
func CfnBroker_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_amazonmq.CfnBroker",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnBroker_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_amazonmq.CfnBroker",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnBroker_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_amazonmq.CfnBroker",
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
		"monocdk.aws_amazonmq.CfnBroker",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnBroker) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnBroker) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnBroker) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnBroker) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnBroker) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnBroker) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnBroker) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnBroker) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBroker) GetMetadata(key *string) interface{} {
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
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnBroker) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnBroker) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnBroker) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBroker) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnBroker) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnBroker) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
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

func (c *jsiiProxy_CfnBroker) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
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

func (c *jsiiProxy_CfnBroker) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBroker) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A list of information about the configuration.
//
// > Does not apply to RabbitMQ brokers.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import amazonmq "github.com/aws/aws-cdk-go/awscdk/aws_amazonmq"
//   configurationIdProperty := &configurationIdProperty{
//   	id: jsii.String("id"),
//   	revision: jsii.Number(123),
//   }
//
type CfnBroker_ConfigurationIdProperty struct {
	// The unique ID that Amazon MQ generates for the configuration.
	Id *string `json:"id" yaml:"id"`
	// The revision number of the configuration.
	Revision *float64 `json:"revision" yaml:"revision"`
}

// Encryption options for the broker.
//
// > Does not apply to RabbitMQ brokers.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import amazonmq "github.com/aws/aws-cdk-go/awscdk/aws_amazonmq"
//   encryptionOptionsProperty := &encryptionOptionsProperty{
//   	useAwsOwnedKey: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   }
//
type CfnBroker_EncryptionOptionsProperty struct {
	// Enables the use of an AWS owned CMK using AWS KMS (KMS).
	//
	// Set to `true` by default, if no value is provided, for example, for RabbitMQ brokers.
	UseAwsOwnedKey interface{} `json:"useAwsOwnedKey" yaml:"useAwsOwnedKey"`
	// The customer master key (CMK) to use for the A AWS KMS (KMS).
	//
	// This key is used to encrypt your data at rest. If not provided, Amazon MQ will use a default CMK to encrypt your data.
	KmsKeyId *string `json:"kmsKeyId" yaml:"kmsKeyId"`
}

// Optional. The metadata of the LDAP server used to authenticate and authorize connections to the broker.
//
// > Does not apply to RabbitMQ brokers.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import amazonmq "github.com/aws/aws-cdk-go/awscdk/aws_amazonmq"
//   ldapServerMetadataProperty := &ldapServerMetadataProperty{
//   	hosts: []*string{
//   		jsii.String("hosts"),
//   	},
//   	roleBase: jsii.String("roleBase"),
//   	roleSearchMatching: jsii.String("roleSearchMatching"),
//   	serviceAccountPassword: jsii.String("serviceAccountPassword"),
//   	serviceAccountUsername: jsii.String("serviceAccountUsername"),
//   	userBase: jsii.String("userBase"),
//   	userSearchMatching: jsii.String("userSearchMatching"),
//
//   	// the properties below are optional
//   	roleName: jsii.String("roleName"),
//   	roleSearchSubtree: jsii.Boolean(false),
//   	userRoleName: jsii.String("userRoleName"),
//   	userSearchSubtree: jsii.Boolean(false),
//   }
//
type CfnBroker_LdapServerMetadataProperty struct {
	// Specifies the location of the LDAP server such as AWS Directory Service for Microsoft Active Directory .
	//
	// Optional failover server.
	Hosts *[]*string `json:"hosts" yaml:"hosts"`
	// The distinguished name of the node in the directory information tree (DIT) to search for roles or groups.
	//
	// For example, `ou=group` , `ou=corp` , `dc=corp` , `dc=example` , `dc=com` .
	RoleBase *string `json:"roleBase" yaml:"roleBase"`
	// The LDAP search filter used to find roles within the roleBase.
	//
	// The distinguished name of the user matched by userSearchMatching is substituted into the `{0}` placeholder in the search filter. The client's username is substituted into the `{1}` placeholder. For example, if you set this option to `(member=uid={1})` for the user janedoe, the search filter becomes `(member=uid=janedoe)` after string substitution. It matches all role entries that have a member attribute equal to `uid=janedoe` under the subtree selected by the `RoleBases` .
	RoleSearchMatching *string `json:"roleSearchMatching" yaml:"roleSearchMatching"`
	// Service account password.
	//
	// A service account is an account in your LDAP server that has access to initiate a connection. For example, `cn=admin` , `dc=corp` , `dc=example` , `dc=com` .
	ServiceAccountPassword *string `json:"serviceAccountPassword" yaml:"serviceAccountPassword"`
	// Service account username.
	//
	// A service account is an account in your LDAP server that has access to initiate a connection. For example, `cn=admin` , `ou=corp` , `dc=corp` , `dc=example` , `dc=com` .
	ServiceAccountUsername *string `json:"serviceAccountUsername" yaml:"serviceAccountUsername"`
	// Select a particular subtree of the directory information tree (DIT) to search for user entries.
	//
	// The subtree is specified by a DN, which specifies the base node of the subtree. For example, by setting this option to `ou=Users` , `ou=corp` , `dc=corp` , `dc=example` , `dc=com` , the search for user entries is restricted to the subtree beneath `ou=Users` , `ou=corp` , `dc=corp` , `dc=example` , `dc=com` .
	UserBase *string `json:"userBase" yaml:"userBase"`
	// The LDAP search filter used to find users within the `userBase` .
	//
	// The client's username is substituted into the `{0}` placeholder in the search filter. For example, if this option is set to `(uid={0})` and the received username is `janedoe` , the search filter becomes `(uid=janedoe)` after string substitution. It will result in matching an entry like `uid=janedoe` , `ou=Users` , `ou=corp` , `dc=corp` , `dc=example` , `dc=com` .
	UserSearchMatching *string `json:"userSearchMatching" yaml:"userSearchMatching"`
	// The group name attribute in a role entry whose value is the name of that role.
	//
	// For example, you can specify `cn` for a group entry's common name. If authentication succeeds, then the user is assigned the the value of the `cn` attribute for each role entry that they are a member of.
	RoleName *string `json:"roleName" yaml:"roleName"`
	// The directory search scope for the role.
	//
	// If set to true, scope is to search the entire subtree.
	RoleSearchSubtree interface{} `json:"roleSearchSubtree" yaml:"roleSearchSubtree"`
	// The name of the LDAP attribute in the user's directory entry for the user's group membership.
	//
	// In some cases, user roles may be identified by the value of an attribute in the user's directory entry. The `UserRoleName` option allows you to provide the name of this attribute.
	UserRoleName *string `json:"userRoleName" yaml:"userRoleName"`
	// The directory search scope for the user.
	//
	// If set to true, scope is to search the entire subtree.
	UserSearchSubtree interface{} `json:"userSearchSubtree" yaml:"userSearchSubtree"`
}

// The list of information about logs to be enabled for the specified broker.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import amazonmq "github.com/aws/aws-cdk-go/awscdk/aws_amazonmq"
//   logListProperty := &logListProperty{
//   	audit: jsii.Boolean(false),
//   	general: jsii.Boolean(false),
//   }
//
type CfnBroker_LogListProperty struct {
	// Enables audit logging.
	//
	// Every user management action made using JMX or the ActiveMQ Web Console is logged. Does not apply to RabbitMQ brokers.
	Audit interface{} `json:"audit" yaml:"audit"`
	// Enables general logging.
	General interface{} `json:"general" yaml:"general"`
}

// The parameters that determine the `WeeklyStartTime` to apply pending updates or patches to the broker.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import amazonmq "github.com/aws/aws-cdk-go/awscdk/aws_amazonmq"
//   maintenanceWindowProperty := &maintenanceWindowProperty{
//   	dayOfWeek: jsii.String("dayOfWeek"),
//   	timeOfDay: jsii.String("timeOfDay"),
//   	timeZone: jsii.String("timeZone"),
//   }
//
type CfnBroker_MaintenanceWindowProperty struct {
	// The day of the week.
	DayOfWeek *string `json:"dayOfWeek" yaml:"dayOfWeek"`
	// The time, in 24-hour format.
	TimeOfDay *string `json:"timeOfDay" yaml:"timeOfDay"`
	// The time zone, UTC by default, in either the Country/City format, or the UTC offset format.
	TimeZone *string `json:"timeZone" yaml:"timeZone"`
}

// A key-value pair to associate with the broker.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import amazonmq "github.com/aws/aws-cdk-go/awscdk/aws_amazonmq"
//   tagsEntryProperty := &tagsEntryProperty{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
//
type CfnBroker_TagsEntryProperty struct {
	// The key in a key-value pair.
	Key *string `json:"key" yaml:"key"`
	// The value in a key-value pair.
	Value *string `json:"value" yaml:"value"`
}

// The list of broker users (persons or applications) who can access queues and topics.
//
// For Amazon MQ for RabbitMQ brokers, one and only one administrative user is accepted and created when a broker is first provisioned. All subsequent broker users are created via the RabbitMQ web console or by using the RabbitMQ management API.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import amazonmq "github.com/aws/aws-cdk-go/awscdk/aws_amazonmq"
//   userProperty := &userProperty{
//   	password: jsii.String("password"),
//   	username: jsii.String("username"),
//
//   	// the properties below are optional
//   	consoleAccess: jsii.Boolean(false),
//   	groups: []*string{
//   		jsii.String("groups"),
//   	},
//   }
//
type CfnBroker_UserProperty struct {
	// The password of the user.
	//
	// This value must be at least 12 characters long, must contain at least 4 unique characters, and must not contain commas, colons, or equal signs (,:=).
	Password *string `json:"password" yaml:"password"`
	// The username of the broker user.
	//
	// For Amazon MQ for ActiveMQ brokers, this value can contain only alphanumeric characters, dashes, periods, underscores, and tildes (- . _ ~). For Amazon MQ for RabbitMQ brokers, this value can contain only alphanumeric characters, dashes, periods, underscores (- . _). This value must not contain a tilde (~) character. Amazon MQ prohibts using guest as a valid usename. This value must be 2-100 characters long.
	//
	// > Do not add personally identifiable information (PII) or other confidential or sensitive information in broker usernames. Broker usernames are accessible to other AWS services, including CloudWatch Logs . Broker usernames are not intended to be used for private or sensitive data.
	Username *string `json:"username" yaml:"username"`
	// Enables access to the ActiveMQ web console for the ActiveMQ user.
	//
	// Does not apply to RabbitMQ brokers.
	ConsoleAccess interface{} `json:"consoleAccess" yaml:"consoleAccess"`
	// The list of groups (20 maximum) to which the ActiveMQ user belongs.
	//
	// This value can contain only alphanumeric characters, dashes, periods, underscores, and tildes (- . _ ~). This value must be 2-100 characters long. Does not apply to RabbitMQ brokers.
	Groups *[]*string `json:"groups" yaml:"groups"`
}

// Properties for defining a `CfnBroker`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import amazonmq "github.com/aws/aws-cdk-go/awscdk/aws_amazonmq"
//   cfnBrokerProps := &cfnBrokerProps{
//   	autoMinorVersionUpgrade: jsii.Boolean(false),
//   	brokerName: jsii.String("brokerName"),
//   	deploymentMode: jsii.String("deploymentMode"),
//   	engineType: jsii.String("engineType"),
//   	engineVersion: jsii.String("engineVersion"),
//   	hostInstanceType: jsii.String("hostInstanceType"),
//   	publiclyAccessible: jsii.Boolean(false),
//   	users: []interface{}{
//   		&userProperty{
//   			password: jsii.String("password"),
//   			username: jsii.String("username"),
//
//   			// the properties below are optional
//   			consoleAccess: jsii.Boolean(false),
//   			groups: []*string{
//   				jsii.String("groups"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	authenticationStrategy: jsii.String("authenticationStrategy"),
//   	configuration: &configurationIdProperty{
//   		id: jsii.String("id"),
//   		revision: jsii.Number(123),
//   	},
//   	encryptionOptions: &encryptionOptionsProperty{
//   		useAwsOwnedKey: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		kmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	ldapServerMetadata: &ldapServerMetadataProperty{
//   		hosts: []*string{
//   			jsii.String("hosts"),
//   		},
//   		roleBase: jsii.String("roleBase"),
//   		roleSearchMatching: jsii.String("roleSearchMatching"),
//   		serviceAccountPassword: jsii.String("serviceAccountPassword"),
//   		serviceAccountUsername: jsii.String("serviceAccountUsername"),
//   		userBase: jsii.String("userBase"),
//   		userSearchMatching: jsii.String("userSearchMatching"),
//
//   		// the properties below are optional
//   		roleName: jsii.String("roleName"),
//   		roleSearchSubtree: jsii.Boolean(false),
//   		userRoleName: jsii.String("userRoleName"),
//   		userSearchSubtree: jsii.Boolean(false),
//   	},
//   	logs: &logListProperty{
//   		audit: jsii.Boolean(false),
//   		general: jsii.Boolean(false),
//   	},
//   	maintenanceWindowStartTime: &maintenanceWindowProperty{
//   		dayOfWeek: jsii.String("dayOfWeek"),
//   		timeOfDay: jsii.String("timeOfDay"),
//   		timeZone: jsii.String("timeZone"),
//   	},
//   	securityGroups: []*string{
//   		jsii.String("securityGroups"),
//   	},
//   	storageType: jsii.String("storageType"),
//   	subnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   	tags: []tagsEntryProperty{
//   		&tagsEntryProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnBrokerProps struct {
	// Enables automatic upgrades to new minor versions for brokers, as new broker engine versions are released and supported by Amazon MQ.
	//
	// Automatic upgrades occur during the scheduled maintenance window of the broker or after a manual broker reboot.
	AutoMinorVersionUpgrade interface{} `json:"autoMinorVersionUpgrade" yaml:"autoMinorVersionUpgrade"`
	// The name of the broker.
	//
	// This value must be unique in your AWS account , 1-50 characters long, must contain only letters, numbers, dashes, and underscores, and must not contain white spaces, brackets, wildcard characters, or special characters.
	//
	// > Do not add personally identifiable information (PII) or other confidential or sensitive information in broker names. Broker names are accessible to other AWS services, including C CloudWatch Logs . Broker names are not intended to be used for private or sensitive data.
	BrokerName *string `json:"brokerName" yaml:"brokerName"`
	// The deployment mode of the broker. Available values:.
	//
	// - `SINGLE_INSTANCE`
	// - `ACTIVE_STANDBY_MULTI_AZ`
	// - `CLUSTER_MULTI_AZ`.
	DeploymentMode *string `json:"deploymentMode" yaml:"deploymentMode"`
	// The type of broker engine.
	//
	// Currently, Amazon MQ supports `ACTIVEMQ` and `RABBITMQ` .
	EngineType *string `json:"engineType" yaml:"engineType"`
	// The version of the broker engine.
	//
	// For a list of supported engine versions, see [Engine](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/broker-engine.html) in the *Amazon MQ Developer Guide* .
	EngineVersion *string `json:"engineVersion" yaml:"engineVersion"`
	// The broker's instance type.
	HostInstanceType *string `json:"hostInstanceType" yaml:"hostInstanceType"`
	// Enables connections from applications outside of the VPC that hosts the broker's subnets.
	PubliclyAccessible interface{} `json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// The list of broker users (persons or applications) who can access queues and topics.
	//
	// For Amazon MQ for RabbitMQ brokers, one and only one administrative user is accepted and created when a broker is first provisioned. All subsequent RabbitMQ users are created by via the RabbitMQ web console or by using the RabbitMQ management API.
	Users interface{} `json:"users" yaml:"users"`
	// Optional.
	//
	// The authentication strategy used to secure the broker. The default is `SIMPLE` .
	AuthenticationStrategy *string `json:"authenticationStrategy" yaml:"authenticationStrategy"`
	// A list of information about the configuration.
	//
	// Does not apply to RabbitMQ brokers.
	Configuration interface{} `json:"configuration" yaml:"configuration"`
	// Encryption options for the broker.
	//
	// Does not apply to RabbitMQ brokers.
	EncryptionOptions interface{} `json:"encryptionOptions" yaml:"encryptionOptions"`
	// Optional.
	//
	// The metadata of the LDAP server used to authenticate and authorize connections to the broker. Does not apply to RabbitMQ brokers.
	LdapServerMetadata interface{} `json:"ldapServerMetadata" yaml:"ldapServerMetadata"`
	// Enables Amazon CloudWatch logging for brokers.
	Logs interface{} `json:"logs" yaml:"logs"`
	// The scheduled time period relative to UTC during which Amazon MQ begins to apply pending updates or patches to the broker.
	MaintenanceWindowStartTime interface{} `json:"maintenanceWindowStartTime" yaml:"maintenanceWindowStartTime"`
	// The list of rules (1 minimum, 125 maximum) that authorize connections to brokers.
	SecurityGroups *[]*string `json:"securityGroups" yaml:"securityGroups"`
	// The broker's storage type.
	StorageType *string `json:"storageType" yaml:"storageType"`
	// The list of groups that define which subnets and IP ranges the broker can use from different Availability Zones.
	//
	// If you specify more than one subnet, the subnets must be in different Availability Zones. Amazon MQ will not be able to create VPC endpoints for your broker with multiple subnets in the same Availability Zone. A SINGLE_INSTANCE deployment requires one subnet (for example, the default subnet). An ACTIVE_STANDBY_MULTI_AZ deployment (ACTIVEMQ) requires two subnets. A CLUSTER_MULTI_AZ deployment (RABBITMQ) has no subnet requirements when deployed with public accessibility, deployment without public accessibility requires at least one subnet.
	//
	// > If you specify subnets in a shared VPC for a RabbitMQ broker, the associated VPC to which the specified subnets belong must be owned by your AWS account . Amazon MQ will not be able to create VPC enpoints in VPCs that are not owned by your AWS account .
	SubnetIds *[]*string `json:"subnetIds" yaml:"subnetIds"`
	// An array of key-value pairs.
	//
	// For more information, see [Using Cost Allocation Tags](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html) in the *Billing and Cost Management User Guide* .
	Tags *[]*CfnBroker_TagsEntryProperty `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::AmazonMQ::Configuration`.
//
// Creates a new configuration for the specified configuration name. Amazon MQ uses the default configuration (the engine type and version).
//
// > Does not apply to RabbitMQ brokers.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import amazonmq "github.com/aws/aws-cdk-go/awscdk/aws_amazonmq"
//   cfnConfiguration := amazonmq.NewCfnConfiguration(this, jsii.String("MyCfnConfiguration"), &cfnConfigurationProps{
//   	data: jsii.String("data"),
//   	engineType: jsii.String("engineType"),
//   	engineVersion: jsii.String("engineVersion"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	authenticationStrategy: jsii.String("authenticationStrategy"),
//   	description: jsii.String("description"),
//   	tags: []tagsEntryProperty{
//   		&tagsEntryProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnConfiguration interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of the Amazon MQ configuration.
	//
	// `arn:aws:mq:us-east-2:123456789012:configuration:MyConfigurationDevelopment:c-1234a5b6-78cd-901e-2fgh-3i45j6k178l9`.
	AttrArn() *string
	// The ID of the Amazon MQ configuration.
	//
	// `c-1234a5b6-78cd-901e-2fgh-3i45j6k178l9`.
	AttrId() *string
	// The revision number of the configuration.
	//
	// `1`.
	AttrRevision() *float64
	// Optional.
	//
	// The authentication strategy associated with the configuration. The default is `SIMPLE` .
	AuthenticationStrategy() *string
	SetAuthenticationStrategy(val *string)
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
	// The base64-encoded XML configuration.
	Data() *string
	SetData(val *string)
	// The description of the configuration.
	Description() *string
	SetDescription(val *string)
	// The type of broker engine.
	//
	// Note: Currently, Amazon MQ only supports ACTIVEMQ for creating and editing broker configurations.
	EngineType() *string
	SetEngineType(val *string)
	// The version of the broker engine.
	//
	// For a list of supported engine versions, see [](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/broker-engine.html)
	EngineVersion() *string
	SetEngineVersion(val *string)
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
	// The name of the configuration.
	//
	// This value can contain only alphanumeric characters, dashes, periods, underscores, and tildes (- . _ ~). This value must be 1-150 characters long.
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
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// Create tags when creating the configuration.
	Tags() awscdk.TagManager
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

// The jsii proxy struct for CfnConfiguration
type jsiiProxy_CfnConfiguration struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnConfiguration) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfiguration) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfiguration) AttrRevision() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrRevision",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfiguration) AuthenticationStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfiguration) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfiguration) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfiguration) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfiguration) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfiguration) Data() *string {
	var returns *string
	_jsii_.Get(
		j,
		"data",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfiguration) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfiguration) EngineType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfiguration) EngineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfiguration) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfiguration) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfiguration) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfiguration) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfiguration) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfiguration) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfiguration) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::AmazonMQ::Configuration`.
func NewCfnConfiguration(scope awscdk.Construct, id *string, props *CfnConfigurationProps) CfnConfiguration {
	_init_.Initialize()

	j := jsiiProxy_CfnConfiguration{}

	_jsii_.Create(
		"monocdk.aws_amazonmq.CfnConfiguration",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AmazonMQ::Configuration`.
func NewCfnConfiguration_Override(c CfnConfiguration, scope awscdk.Construct, id *string, props *CfnConfigurationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_amazonmq.CfnConfiguration",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnConfiguration) SetAuthenticationStrategy(val *string) {
	_jsii_.Set(
		j,
		"authenticationStrategy",
		val,
	)
}

func (j *jsiiProxy_CfnConfiguration) SetData(val *string) {
	_jsii_.Set(
		j,
		"data",
		val,
	)
}

func (j *jsiiProxy_CfnConfiguration) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnConfiguration) SetEngineType(val *string) {
	_jsii_.Set(
		j,
		"engineType",
		val,
	)
}

func (j *jsiiProxy_CfnConfiguration) SetEngineVersion(val *string) {
	_jsii_.Set(
		j,
		"engineVersion",
		val,
	)
}

func (j *jsiiProxy_CfnConfiguration) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
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
func CfnConfiguration_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_amazonmq.CfnConfiguration",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnConfiguration_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_amazonmq.CfnConfiguration",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_amazonmq.CfnConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnConfiguration_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_amazonmq.CfnConfiguration",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnConfiguration) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnConfiguration) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnConfiguration) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnConfiguration) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnConfiguration) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnConfiguration) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnConfiguration) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnConfiguration) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConfiguration) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConfiguration) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnConfiguration) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnConfiguration) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnConfiguration) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConfiguration) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnConfiguration) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnConfiguration) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConfiguration) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConfiguration) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConfiguration) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConfiguration) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A key-value pair to associate with the configuration.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import amazonmq "github.com/aws/aws-cdk-go/awscdk/aws_amazonmq"
//   tagsEntryProperty := &tagsEntryProperty{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
//
type CfnConfiguration_TagsEntryProperty struct {
	// The key in a key-value pair.
	Key *string `json:"key" yaml:"key"`
	// The value in a key-value pair.
	Value *string `json:"value" yaml:"value"`
}

// A CloudFormation `AWS::AmazonMQ::ConfigurationAssociation`.
//
// Use the AWS CloudFormation `AWS::AmazonMQ::ConfigurationAssociation` resource to associate a configuration with a broker, or return information about the specified ConfigurationAssociation. Only use one per broker, and don't use a configuration on the broker resource if you have associated a configuration with that broker.
//
// > Does not apply to RabbitMQ brokers.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import amazonmq "github.com/aws/aws-cdk-go/awscdk/aws_amazonmq"
//   cfnConfigurationAssociation := amazonmq.NewCfnConfigurationAssociation(this, jsii.String("MyCfnConfigurationAssociation"), &cfnConfigurationAssociationProps{
//   	broker: jsii.String("broker"),
//   	configuration: &configurationIdProperty{
//   		id: jsii.String("id"),
//   		revision: jsii.Number(123),
//   	},
//   })
//
type CfnConfigurationAssociation interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The broker to associate with a configuration.
	Broker() *string
	SetBroker(val *string)
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// The configuration to associate with a broker.
	Configuration() interface{}
	SetConfiguration(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
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
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
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

// The jsii proxy struct for CfnConfigurationAssociation
type jsiiProxy_CfnConfigurationAssociation struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnConfigurationAssociation) Broker() *string {
	var returns *string
	_jsii_.Get(
		j,
		"broker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationAssociation) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationAssociation) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationAssociation) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationAssociation) Configuration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationAssociation) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationAssociation) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationAssociation) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationAssociation) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationAssociation) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationAssociation) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::AmazonMQ::ConfigurationAssociation`.
func NewCfnConfigurationAssociation(scope awscdk.Construct, id *string, props *CfnConfigurationAssociationProps) CfnConfigurationAssociation {
	_init_.Initialize()

	j := jsiiProxy_CfnConfigurationAssociation{}

	_jsii_.Create(
		"monocdk.aws_amazonmq.CfnConfigurationAssociation",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AmazonMQ::ConfigurationAssociation`.
func NewCfnConfigurationAssociation_Override(c CfnConfigurationAssociation, scope awscdk.Construct, id *string, props *CfnConfigurationAssociationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_amazonmq.CfnConfigurationAssociation",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnConfigurationAssociation) SetBroker(val *string) {
	_jsii_.Set(
		j,
		"broker",
		val,
	)
}

func (j *jsiiProxy_CfnConfigurationAssociation) SetConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"configuration",
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
func CfnConfigurationAssociation_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_amazonmq.CfnConfigurationAssociation",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnConfigurationAssociation_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_amazonmq.CfnConfigurationAssociation",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnConfigurationAssociation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_amazonmq.CfnConfigurationAssociation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnConfigurationAssociation_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_amazonmq.CfnConfigurationAssociation",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnConfigurationAssociation) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnConfigurationAssociation) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnConfigurationAssociation) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnConfigurationAssociation) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnConfigurationAssociation) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnConfigurationAssociation) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnConfigurationAssociation) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnConfigurationAssociation) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConfigurationAssociation) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConfigurationAssociation) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnConfigurationAssociation) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnConfigurationAssociation) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnConfigurationAssociation) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConfigurationAssociation) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnConfigurationAssociation) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnConfigurationAssociation) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConfigurationAssociation) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConfigurationAssociation) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnConfigurationAssociation) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConfigurationAssociation) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConfigurationAssociation) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The `ConfigurationId` property type specifies a configuration Id and the revision of a configuration.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import amazonmq "github.com/aws/aws-cdk-go/awscdk/aws_amazonmq"
//   configurationIdProperty := &configurationIdProperty{
//   	id: jsii.String("id"),
//   	revision: jsii.Number(123),
//   }
//
type CfnConfigurationAssociation_ConfigurationIdProperty struct {
	// The unique ID that Amazon MQ generates for the configuration.
	Id *string `json:"id" yaml:"id"`
	// The revision number of the configuration.
	Revision *float64 `json:"revision" yaml:"revision"`
}

// Properties for defining a `CfnConfigurationAssociation`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import amazonmq "github.com/aws/aws-cdk-go/awscdk/aws_amazonmq"
//   cfnConfigurationAssociationProps := &cfnConfigurationAssociationProps{
//   	broker: jsii.String("broker"),
//   	configuration: &configurationIdProperty{
//   		id: jsii.String("id"),
//   		revision: jsii.Number(123),
//   	},
//   }
//
type CfnConfigurationAssociationProps struct {
	// The broker to associate with a configuration.
	Broker *string `json:"broker" yaml:"broker"`
	// The configuration to associate with a broker.
	Configuration interface{} `json:"configuration" yaml:"configuration"`
}

// Properties for defining a `CfnConfiguration`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import amazonmq "github.com/aws/aws-cdk-go/awscdk/aws_amazonmq"
//   cfnConfigurationProps := &cfnConfigurationProps{
//   	data: jsii.String("data"),
//   	engineType: jsii.String("engineType"),
//   	engineVersion: jsii.String("engineVersion"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	authenticationStrategy: jsii.String("authenticationStrategy"),
//   	description: jsii.String("description"),
//   	tags: []tagsEntryProperty{
//   		&tagsEntryProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnConfigurationProps struct {
	// The base64-encoded XML configuration.
	Data *string `json:"data" yaml:"data"`
	// The type of broker engine.
	//
	// Note: Currently, Amazon MQ only supports ACTIVEMQ for creating and editing broker configurations.
	EngineType *string `json:"engineType" yaml:"engineType"`
	// The version of the broker engine.
	//
	// For a list of supported engine versions, see [](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/broker-engine.html)
	EngineVersion *string `json:"engineVersion" yaml:"engineVersion"`
	// The name of the configuration.
	//
	// This value can contain only alphanumeric characters, dashes, periods, underscores, and tildes (- . _ ~). This value must be 1-150 characters long.
	Name *string `json:"name" yaml:"name"`
	// Optional.
	//
	// The authentication strategy associated with the configuration. The default is `SIMPLE` .
	AuthenticationStrategy *string `json:"authenticationStrategy" yaml:"authenticationStrategy"`
	// The description of the configuration.
	Description *string `json:"description" yaml:"description"`
	// Create tags when creating the configuration.
	Tags *[]*CfnConfiguration_TagsEntryProperty `json:"tags" yaml:"tags"`
}

