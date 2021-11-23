package awsamazonmq

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsamazonmq/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::AmazonMQ::Broker`.
//
// TODO: EXAMPLE
//
type CfnBroker interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrAmqpEndpoints() *[]*string
	AttrArn() *string
	AttrConfigurationId() *string
	AttrConfigurationRevision() *float64
	AttrIpAddresses() *[]*string
	AttrMqttEndpoints() *[]*string
	AttrOpenWireEndpoints() *[]*string
	AttrStompEndpoints() *[]*string
	AttrWssEndpoints() *[]*string
	AuthenticationStrategy() *string
	SetAuthenticationStrategy(val *string)
	AutoMinorVersionUpgrade() interface{}
	SetAutoMinorVersionUpgrade(val interface{})
	BrokerName() *string
	SetBrokerName(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	Configuration() interface{}
	SetConfiguration(val interface{})
	CreationStack() *[]*string
	DeploymentMode() *string
	SetDeploymentMode(val *string)
	EncryptionOptions() interface{}
	SetEncryptionOptions(val interface{})
	EngineType() *string
	SetEngineType(val *string)
	EngineVersion() *string
	SetEngineVersion(val *string)
	HostInstanceType() *string
	SetHostInstanceType(val *string)
	LdapServerMetadata() interface{}
	SetLdapServerMetadata(val interface{})
	LogicalId() *string
	Logs() interface{}
	SetLogs(val interface{})
	MaintenanceWindowStartTime() interface{}
	SetMaintenanceWindowStartTime(val interface{})
	Node() constructs.Node
	PubliclyAccessible() interface{}
	SetPubliclyAccessible(val interface{})
	Ref() *string
	SecurityGroups() *[]*string
	SetSecurityGroups(val *[]*string)
	Stack() awscdk.Stack
	StorageType() *string
	SetStorageType(val *string)
	SubnetIds() *[]*string
	SetSubnetIds(val *[]*string)
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	Users() interface{}
	SetUsers(val interface{})
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
func NewCfnBroker(scope constructs.Construct, id *string, props *CfnBrokerProps) CfnBroker {
	_init_.Initialize()

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
		"aws-cdk-lib.aws_amazonmq.CfnBroker",
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
		"aws-cdk-lib.aws_amazonmq.CfnBroker",
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
func CfnBroker_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

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

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnBroker) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnBroker) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnBroker) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnBroker) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnBroker) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnBroker) AddPropertyOverride(propertyPath *string, value interface{}) {
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
// Experimental.
func (c *jsiiProxy_CfnBroker) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
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

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnBroker) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnBroker) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
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

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
// Experimental.
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

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
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

// Experimental.
func (c *jsiiProxy_CfnBroker) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// TODO: EXAMPLE
//
type CfnBroker_ConfigurationIdProperty struct {
	// `CfnBroker.ConfigurationIdProperty.Id`.
	Id *string `json:"id"`
	// `CfnBroker.ConfigurationIdProperty.Revision`.
	Revision *float64 `json:"revision"`
}

// TODO: EXAMPLE
//
type CfnBroker_EncryptionOptionsProperty struct {
	// `CfnBroker.EncryptionOptionsProperty.KmsKeyId`.
	KmsKeyId *string `json:"kmsKeyId"`
	// `CfnBroker.EncryptionOptionsProperty.UseAwsOwnedKey`.
	UseAwsOwnedKey interface{} `json:"useAwsOwnedKey"`
}

// TODO: EXAMPLE
//
type CfnBroker_LdapServerMetadataProperty struct {
	// `CfnBroker.LdapServerMetadataProperty.Hosts`.
	Hosts *[]*string `json:"hosts"`
	// `CfnBroker.LdapServerMetadataProperty.RoleBase`.
	RoleBase *string `json:"roleBase"`
	// `CfnBroker.LdapServerMetadataProperty.RoleName`.
	RoleName *string `json:"roleName"`
	// `CfnBroker.LdapServerMetadataProperty.RoleSearchMatching`.
	RoleSearchMatching *string `json:"roleSearchMatching"`
	// `CfnBroker.LdapServerMetadataProperty.RoleSearchSubtree`.
	RoleSearchSubtree interface{} `json:"roleSearchSubtree"`
	// `CfnBroker.LdapServerMetadataProperty.ServiceAccountPassword`.
	ServiceAccountPassword *string `json:"serviceAccountPassword"`
	// `CfnBroker.LdapServerMetadataProperty.ServiceAccountUsername`.
	ServiceAccountUsername *string `json:"serviceAccountUsername"`
	// `CfnBroker.LdapServerMetadataProperty.UserBase`.
	UserBase *string `json:"userBase"`
	// `CfnBroker.LdapServerMetadataProperty.UserRoleName`.
	UserRoleName *string `json:"userRoleName"`
	// `CfnBroker.LdapServerMetadataProperty.UserSearchMatching`.
	UserSearchMatching *string `json:"userSearchMatching"`
	// `CfnBroker.LdapServerMetadataProperty.UserSearchSubtree`.
	UserSearchSubtree interface{} `json:"userSearchSubtree"`
}

// TODO: EXAMPLE
//
type CfnBroker_LogListProperty struct {
	// `CfnBroker.LogListProperty.Audit`.
	Audit interface{} `json:"audit"`
	// `CfnBroker.LogListProperty.General`.
	General interface{} `json:"general"`
}

// TODO: EXAMPLE
//
type CfnBroker_MaintenanceWindowProperty struct {
	// `CfnBroker.MaintenanceWindowProperty.DayOfWeek`.
	DayOfWeek *string `json:"dayOfWeek"`
	// `CfnBroker.MaintenanceWindowProperty.TimeOfDay`.
	TimeOfDay *string `json:"timeOfDay"`
	// `CfnBroker.MaintenanceWindowProperty.TimeZone`.
	TimeZone *string `json:"timeZone"`
}

// TODO: EXAMPLE
//
type CfnBroker_TagsEntryProperty struct {
	// `CfnBroker.TagsEntryProperty.Key`.
	Key *string `json:"key"`
	// `CfnBroker.TagsEntryProperty.Value`.
	Value *string `json:"value"`
}

// TODO: EXAMPLE
//
type CfnBroker_UserProperty struct {
	// `CfnBroker.UserProperty.ConsoleAccess`.
	ConsoleAccess interface{} `json:"consoleAccess"`
	// `CfnBroker.UserProperty.Groups`.
	Groups *[]*string `json:"groups"`
	// `CfnBroker.UserProperty.Password`.
	Password *string `json:"password"`
	// `CfnBroker.UserProperty.Username`.
	Username *string `json:"username"`
}

// Properties for defining a `AWS::AmazonMQ::Broker`.
//
// TODO: EXAMPLE
//
type CfnBrokerProps struct {
	// `AWS::AmazonMQ::Broker.AuthenticationStrategy`.
	AuthenticationStrategy *string `json:"authenticationStrategy"`
	// `AWS::AmazonMQ::Broker.AutoMinorVersionUpgrade`.
	AutoMinorVersionUpgrade interface{} `json:"autoMinorVersionUpgrade"`
	// `AWS::AmazonMQ::Broker.BrokerName`.
	BrokerName *string `json:"brokerName"`
	// `AWS::AmazonMQ::Broker.Configuration`.
	Configuration interface{} `json:"configuration"`
	// `AWS::AmazonMQ::Broker.DeploymentMode`.
	DeploymentMode *string `json:"deploymentMode"`
	// `AWS::AmazonMQ::Broker.EncryptionOptions`.
	EncryptionOptions interface{} `json:"encryptionOptions"`
	// `AWS::AmazonMQ::Broker.EngineType`.
	EngineType *string `json:"engineType"`
	// `AWS::AmazonMQ::Broker.EngineVersion`.
	EngineVersion *string `json:"engineVersion"`
	// `AWS::AmazonMQ::Broker.HostInstanceType`.
	HostInstanceType *string `json:"hostInstanceType"`
	// `AWS::AmazonMQ::Broker.LdapServerMetadata`.
	LdapServerMetadata interface{} `json:"ldapServerMetadata"`
	// `AWS::AmazonMQ::Broker.Logs`.
	Logs interface{} `json:"logs"`
	// `AWS::AmazonMQ::Broker.MaintenanceWindowStartTime`.
	MaintenanceWindowStartTime interface{} `json:"maintenanceWindowStartTime"`
	// `AWS::AmazonMQ::Broker.PubliclyAccessible`.
	PubliclyAccessible interface{} `json:"publiclyAccessible"`
	// `AWS::AmazonMQ::Broker.SecurityGroups`.
	SecurityGroups *[]*string `json:"securityGroups"`
	// `AWS::AmazonMQ::Broker.StorageType`.
	StorageType *string `json:"storageType"`
	// `AWS::AmazonMQ::Broker.SubnetIds`.
	SubnetIds *[]*string `json:"subnetIds"`
	// `AWS::AmazonMQ::Broker.Tags`.
	Tags *[]*CfnBroker_TagsEntryProperty `json:"tags"`
	// `AWS::AmazonMQ::Broker.Users`.
	Users interface{} `json:"users"`
}

// A CloudFormation `AWS::AmazonMQ::Configuration`.
//
// TODO: EXAMPLE
//
type CfnConfiguration interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	AttrId() *string
	AttrRevision() *float64
	AuthenticationStrategy() *string
	SetAuthenticationStrategy(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	Data() *string
	SetData(val *string)
	Description() *string
	SetDescription(val *string)
	EngineType() *string
	SetEngineType(val *string)
	EngineVersion() *string
	SetEngineVersion(val *string)
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
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

func (j *jsiiProxy_CfnConfiguration) Node() constructs.Node {
	var returns constructs.Node
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
func NewCfnConfiguration(scope constructs.Construct, id *string, props *CfnConfigurationProps) CfnConfiguration {
	_init_.Initialize()

	j := jsiiProxy_CfnConfiguration{}

	_jsii_.Create(
		"aws-cdk-lib.aws_amazonmq.CfnConfiguration",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AmazonMQ::Configuration`.
func NewCfnConfiguration_Override(c CfnConfiguration, scope constructs.Construct, id *string, props *CfnConfigurationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_amazonmq.CfnConfiguration",
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
		"aws-cdk-lib.aws_amazonmq.CfnConfiguration",
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
		"aws-cdk-lib.aws_amazonmq.CfnConfiguration",
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
func CfnConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_amazonmq.CfnConfiguration",
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
		"aws-cdk-lib.aws_amazonmq.CfnConfiguration",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnConfiguration) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnConfiguration) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnConfiguration) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnConfiguration) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnConfiguration) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnConfiguration) AddPropertyOverride(propertyPath *string, value interface{}) {
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
// Experimental.
func (c *jsiiProxy_CfnConfiguration) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
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

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnConfiguration) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnConfiguration) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
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

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
// Experimental.
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

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
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

// Experimental.
func (c *jsiiProxy_CfnConfiguration) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// TODO: EXAMPLE
//
type CfnConfiguration_TagsEntryProperty struct {
	// `CfnConfiguration.TagsEntryProperty.Key`.
	Key *string `json:"key"`
	// `CfnConfiguration.TagsEntryProperty.Value`.
	Value *string `json:"value"`
}

// A CloudFormation `AWS::AmazonMQ::ConfigurationAssociation`.
//
// TODO: EXAMPLE
//
type CfnConfigurationAssociation interface {
	awscdk.CfnResource
	awscdk.IInspectable
	Broker() *string
	SetBroker(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	Configuration() interface{}
	SetConfiguration(val interface{})
	CreationStack() *[]*string
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	Stack() awscdk.Stack
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

func (j *jsiiProxy_CfnConfigurationAssociation) Node() constructs.Node {
	var returns constructs.Node
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
func NewCfnConfigurationAssociation(scope constructs.Construct, id *string, props *CfnConfigurationAssociationProps) CfnConfigurationAssociation {
	_init_.Initialize()

	j := jsiiProxy_CfnConfigurationAssociation{}

	_jsii_.Create(
		"aws-cdk-lib.aws_amazonmq.CfnConfigurationAssociation",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AmazonMQ::ConfigurationAssociation`.
func NewCfnConfigurationAssociation_Override(c CfnConfigurationAssociation, scope constructs.Construct, id *string, props *CfnConfigurationAssociationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_amazonmq.CfnConfigurationAssociation",
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
		"aws-cdk-lib.aws_amazonmq.CfnConfigurationAssociation",
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
		"aws-cdk-lib.aws_amazonmq.CfnConfigurationAssociation",
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
func CfnConfigurationAssociation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_amazonmq.CfnConfigurationAssociation",
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
		"aws-cdk-lib.aws_amazonmq.CfnConfigurationAssociation",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnConfigurationAssociation) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnConfigurationAssociation) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnConfigurationAssociation) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnConfigurationAssociation) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnConfigurationAssociation) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnConfigurationAssociation) AddPropertyOverride(propertyPath *string, value interface{}) {
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
// Experimental.
func (c *jsiiProxy_CfnConfigurationAssociation) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
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

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnConfigurationAssociation) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnConfigurationAssociation) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
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

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
// Experimental.
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

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
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

// Experimental.
func (c *jsiiProxy_CfnConfigurationAssociation) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// TODO: EXAMPLE
//
type CfnConfigurationAssociation_ConfigurationIdProperty struct {
	// `CfnConfigurationAssociation.ConfigurationIdProperty.Id`.
	Id *string `json:"id"`
	// `CfnConfigurationAssociation.ConfigurationIdProperty.Revision`.
	Revision *float64 `json:"revision"`
}

// Properties for defining a `AWS::AmazonMQ::ConfigurationAssociation`.
//
// TODO: EXAMPLE
//
type CfnConfigurationAssociationProps struct {
	// `AWS::AmazonMQ::ConfigurationAssociation.Broker`.
	Broker *string `json:"broker"`
	// `AWS::AmazonMQ::ConfigurationAssociation.Configuration`.
	Configuration interface{} `json:"configuration"`
}

// Properties for defining a `AWS::AmazonMQ::Configuration`.
//
// TODO: EXAMPLE
//
type CfnConfigurationProps struct {
	// `AWS::AmazonMQ::Configuration.AuthenticationStrategy`.
	AuthenticationStrategy *string `json:"authenticationStrategy"`
	// `AWS::AmazonMQ::Configuration.Data`.
	Data *string `json:"data"`
	// `AWS::AmazonMQ::Configuration.Description`.
	Description *string `json:"description"`
	// `AWS::AmazonMQ::Configuration.EngineType`.
	EngineType *string `json:"engineType"`
	// `AWS::AmazonMQ::Configuration.EngineVersion`.
	EngineVersion *string `json:"engineVersion"`
	// `AWS::AmazonMQ::Configuration.Name`.
	Name *string `json:"name"`
	// `AWS::AmazonMQ::Configuration.Tags`.
	Tags *[]*CfnConfiguration_TagsEntryProperty `json:"tags"`
}

