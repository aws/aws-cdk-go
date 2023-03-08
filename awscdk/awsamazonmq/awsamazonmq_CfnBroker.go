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
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnBroker := awscdk.Aws_amazonmq.NewCfnBroker(this, jsii.String("MyCfnBroker"), &cfnBrokerProps{
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

	if err := validateNewCfnBrokerParameters(scope, id, props); err != nil {
		panic(err)
	}
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
	if err := j.validateSetEngineVersionParameters(val); err != nil {
		panic(err)
	}
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
// Experimental.
func CfnBroker_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnBroker_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
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

	if err := validateCfnBroker_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
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

	if err := validateCfnBroker_IsConstructParameters(x); err != nil {
		panic(err)
	}
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
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
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

func (c *jsiiProxy_CfnBroker) GetAtt(attributeName *string) awscdk.Reference {
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

func (c *jsiiProxy_CfnBroker) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnBroker) OnSynthesize(session constructs.ISynthesisSession) {
	if err := c.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
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
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
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
	if err := c.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
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
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

