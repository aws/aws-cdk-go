package awsdms

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdms/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::DMS::Certificate`.
//
// The `AWS::DMS::Certificate` resource creates an SSL certificate that encrypts connections between AWS DMS endpoints and the replication instance.
//
// TODO: EXAMPLE
//
type CfnCertificate interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CertificateIdentifier() *string
	SetCertificateIdentifier(val *string)
	CertificatePem() *string
	SetCertificatePem(val *string)
	CertificateWallet() *string
	SetCertificateWallet(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
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

// The jsii proxy struct for CfnCertificate
type jsiiProxy_CfnCertificate struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnCertificate) CertificateIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificate) CertificatePem() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificatePem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificate) CertificateWallet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateWallet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificate) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificate) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificate) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificate) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificate) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificate) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificate) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificate) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::DMS::Certificate`.
func NewCfnCertificate(scope constructs.Construct, id *string, props *CfnCertificateProps) CfnCertificate {
	_init_.Initialize()

	j := jsiiProxy_CfnCertificate{}

	_jsii_.Create(
		"aws-cdk-lib.aws_dms.CfnCertificate",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::DMS::Certificate`.
func NewCfnCertificate_Override(c CfnCertificate, scope constructs.Construct, id *string, props *CfnCertificateProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_dms.CfnCertificate",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnCertificate) SetCertificateIdentifier(val *string) {
	_jsii_.Set(
		j,
		"certificateIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnCertificate) SetCertificatePem(val *string) {
	_jsii_.Set(
		j,
		"certificatePem",
		val,
	)
}

func (j *jsiiProxy_CfnCertificate) SetCertificateWallet(val *string) {
	_jsii_.Set(
		j,
		"certificateWallet",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnCertificate_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_dms.CfnCertificate",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnCertificate_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_dms.CfnCertificate",
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
func CfnCertificate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_dms.CfnCertificate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCertificate_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_dms.CfnCertificate",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnCertificate) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnCertificate) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnCertificate) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnCertificate) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnCertificate) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnCertificate) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnCertificate) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnCertificate) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnCertificate) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnCertificate) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnCertificate) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnCertificate) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnCertificate) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnCertificate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCertificate) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnCertificate`.
//
// TODO: EXAMPLE
//
type CfnCertificateProps struct {
	// A customer-assigned name for the certificate.
	//
	// Identifiers must begin with a letter and must contain only ASCII letters, digits, and hyphens. They can't end with a hyphen or contain two consecutive hyphens.
	CertificateIdentifier *string `json:"certificateIdentifier" yaml:"certificateIdentifier"`
	// The contents of a `.pem` file, which contains an X.509 certificate.
	CertificatePem *string `json:"certificatePem" yaml:"certificatePem"`
	// The location of an imported Oracle Wallet certificate for use with SSL.
	//
	// Example: `filebase64("${path.root}/rds-ca-2019-root.sso")`
	CertificateWallet *string `json:"certificateWallet" yaml:"certificateWallet"`
}

// A CloudFormation `AWS::DMS::Endpoint`.
//
// The `AWS::DMS::Endpoint` resource creates an AWS DMS endpoint.
//
// Currently, the only endpoint setting types that AWS CloudFormation supports are *DynamoDBSettings* , *ElasticSearchSettings* , and *NeptuneSettings* .
//
// TODO: EXAMPLE
//
type CfnEndpoint interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrExternalId() *string
	CertificateArn() *string
	SetCertificateArn(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	DatabaseName() *string
	SetDatabaseName(val *string)
	DocDbSettings() interface{}
	SetDocDbSettings(val interface{})
	DynamoDbSettings() interface{}
	SetDynamoDbSettings(val interface{})
	ElasticsearchSettings() interface{}
	SetElasticsearchSettings(val interface{})
	EndpointIdentifier() *string
	SetEndpointIdentifier(val *string)
	EndpointType() *string
	SetEndpointType(val *string)
	EngineName() *string
	SetEngineName(val *string)
	ExtraConnectionAttributes() *string
	SetExtraConnectionAttributes(val *string)
	GcpMySqlSettings() interface{}
	SetGcpMySqlSettings(val interface{})
	IbmDb2Settings() interface{}
	SetIbmDb2Settings(val interface{})
	KafkaSettings() interface{}
	SetKafkaSettings(val interface{})
	KinesisSettings() interface{}
	SetKinesisSettings(val interface{})
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	LogicalId() *string
	MicrosoftSqlServerSettings() interface{}
	SetMicrosoftSqlServerSettings(val interface{})
	MongoDbSettings() interface{}
	SetMongoDbSettings(val interface{})
	MySqlSettings() interface{}
	SetMySqlSettings(val interface{})
	NeptuneSettings() interface{}
	SetNeptuneSettings(val interface{})
	Node() constructs.Node
	OracleSettings() interface{}
	SetOracleSettings(val interface{})
	Password() *string
	SetPassword(val *string)
	Port() *float64
	SetPort(val *float64)
	PostgreSqlSettings() interface{}
	SetPostgreSqlSettings(val interface{})
	RedisSettings() interface{}
	SetRedisSettings(val interface{})
	RedshiftSettings() interface{}
	SetRedshiftSettings(val interface{})
	Ref() *string
	ResourceIdentifier() *string
	SetResourceIdentifier(val *string)
	S3Settings() interface{}
	SetS3Settings(val interface{})
	ServerName() *string
	SetServerName(val *string)
	SslMode() *string
	SetSslMode(val *string)
	Stack() awscdk.Stack
	SybaseSettings() interface{}
	SetSybaseSettings(val interface{})
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	Username() *string
	SetUsername(val *string)
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

// The jsii proxy struct for CfnEndpoint
type jsiiProxy_CfnEndpoint struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnEndpoint) AttrExternalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrExternalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) CertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) DocDbSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"docDbSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) DynamoDbSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamoDbSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) ElasticsearchSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"elasticsearchSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) EndpointIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) EndpointType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) EngineName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) ExtraConnectionAttributes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extraConnectionAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) GcpMySqlSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gcpMySqlSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) IbmDb2Settings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ibmDb2Settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) KafkaSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kafkaSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) KinesisSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kinesisSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) MicrosoftSqlServerSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"microsoftSqlServerSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) MongoDbSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mongoDbSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) MySqlSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mySqlSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) NeptuneSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"neptuneSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) OracleSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oracleSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) PostgreSqlSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"postgreSqlSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) RedisSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"redisSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) RedshiftSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"redshiftSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) ResourceIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) S3Settings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3Settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) ServerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) SslMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) SybaseSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sybaseSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}


// Create a new `AWS::DMS::Endpoint`.
func NewCfnEndpoint(scope constructs.Construct, id *string, props *CfnEndpointProps) CfnEndpoint {
	_init_.Initialize()

	j := jsiiProxy_CfnEndpoint{}

	_jsii_.Create(
		"aws-cdk-lib.aws_dms.CfnEndpoint",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::DMS::Endpoint`.
func NewCfnEndpoint_Override(c CfnEndpoint, scope constructs.Construct, id *string, props *CfnEndpointProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_dms.CfnEndpoint",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnEndpoint) SetCertificateArn(val *string) {
	_jsii_.Set(
		j,
		"certificateArn",
		val,
	)
}

func (j *jsiiProxy_CfnEndpoint) SetDatabaseName(val *string) {
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_CfnEndpoint) SetDocDbSettings(val interface{}) {
	_jsii_.Set(
		j,
		"docDbSettings",
		val,
	)
}

func (j *jsiiProxy_CfnEndpoint) SetDynamoDbSettings(val interface{}) {
	_jsii_.Set(
		j,
		"dynamoDbSettings",
		val,
	)
}

func (j *jsiiProxy_CfnEndpoint) SetElasticsearchSettings(val interface{}) {
	_jsii_.Set(
		j,
		"elasticsearchSettings",
		val,
	)
}

func (j *jsiiProxy_CfnEndpoint) SetEndpointIdentifier(val *string) {
	_jsii_.Set(
		j,
		"endpointIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnEndpoint) SetEndpointType(val *string) {
	_jsii_.Set(
		j,
		"endpointType",
		val,
	)
}

func (j *jsiiProxy_CfnEndpoint) SetEngineName(val *string) {
	_jsii_.Set(
		j,
		"engineName",
		val,
	)
}

func (j *jsiiProxy_CfnEndpoint) SetExtraConnectionAttributes(val *string) {
	_jsii_.Set(
		j,
		"extraConnectionAttributes",
		val,
	)
}

func (j *jsiiProxy_CfnEndpoint) SetGcpMySqlSettings(val interface{}) {
	_jsii_.Set(
		j,
		"gcpMySqlSettings",
		val,
	)
}

func (j *jsiiProxy_CfnEndpoint) SetIbmDb2Settings(val interface{}) {
	_jsii_.Set(
		j,
		"ibmDb2Settings",
		val,
	)
}

func (j *jsiiProxy_CfnEndpoint) SetKafkaSettings(val interface{}) {
	_jsii_.Set(
		j,
		"kafkaSettings",
		val,
	)
}

func (j *jsiiProxy_CfnEndpoint) SetKinesisSettings(val interface{}) {
	_jsii_.Set(
		j,
		"kinesisSettings",
		val,
	)
}

func (j *jsiiProxy_CfnEndpoint) SetKmsKeyId(val *string) {
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_CfnEndpoint) SetMicrosoftSqlServerSettings(val interface{}) {
	_jsii_.Set(
		j,
		"microsoftSqlServerSettings",
		val,
	)
}

func (j *jsiiProxy_CfnEndpoint) SetMongoDbSettings(val interface{}) {
	_jsii_.Set(
		j,
		"mongoDbSettings",
		val,
	)
}

func (j *jsiiProxy_CfnEndpoint) SetMySqlSettings(val interface{}) {
	_jsii_.Set(
		j,
		"mySqlSettings",
		val,
	)
}

func (j *jsiiProxy_CfnEndpoint) SetNeptuneSettings(val interface{}) {
	_jsii_.Set(
		j,
		"neptuneSettings",
		val,
	)
}

func (j *jsiiProxy_CfnEndpoint) SetOracleSettings(val interface{}) {
	_jsii_.Set(
		j,
		"oracleSettings",
		val,
	)
}

func (j *jsiiProxy_CfnEndpoint) SetPassword(val *string) {
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_CfnEndpoint) SetPort(val *float64) {
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_CfnEndpoint) SetPostgreSqlSettings(val interface{}) {
	_jsii_.Set(
		j,
		"postgreSqlSettings",
		val,
	)
}

func (j *jsiiProxy_CfnEndpoint) SetRedisSettings(val interface{}) {
	_jsii_.Set(
		j,
		"redisSettings",
		val,
	)
}

func (j *jsiiProxy_CfnEndpoint) SetRedshiftSettings(val interface{}) {
	_jsii_.Set(
		j,
		"redshiftSettings",
		val,
	)
}

func (j *jsiiProxy_CfnEndpoint) SetResourceIdentifier(val *string) {
	_jsii_.Set(
		j,
		"resourceIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnEndpoint) SetS3Settings(val interface{}) {
	_jsii_.Set(
		j,
		"s3Settings",
		val,
	)
}

func (j *jsiiProxy_CfnEndpoint) SetServerName(val *string) {
	_jsii_.Set(
		j,
		"serverName",
		val,
	)
}

func (j *jsiiProxy_CfnEndpoint) SetSslMode(val *string) {
	_jsii_.Set(
		j,
		"sslMode",
		val,
	)
}

func (j *jsiiProxy_CfnEndpoint) SetSybaseSettings(val interface{}) {
	_jsii_.Set(
		j,
		"sybaseSettings",
		val,
	)
}

func (j *jsiiProxy_CfnEndpoint) SetUsername(val *string) {
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnEndpoint_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_dms.CfnEndpoint",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnEndpoint_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_dms.CfnEndpoint",
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
func CfnEndpoint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_dms.CfnEndpoint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEndpoint_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_dms.CfnEndpoint",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnEndpoint) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnEndpoint) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnEndpoint) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnEndpoint) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnEndpoint) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnEndpoint) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnEndpoint) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnEndpoint) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnEndpoint) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnEndpoint) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnEndpoint) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnEndpoint) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnEndpoint) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnEndpoint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEndpoint) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Not currently supported by AWS CloudFormation .
//
// TODO: EXAMPLE
//
type CfnEndpoint_DocDbSettingsProperty struct {
	// Not currently supported by AWS CloudFormation .
	SecretsManagerAccessRoleArn *string `json:"secretsManagerAccessRoleArn" yaml:"secretsManagerAccessRoleArn"`
	// Not currently supported by AWS CloudFormation .
	SecretsManagerSecretId *string `json:"secretsManagerSecretId" yaml:"secretsManagerSecretId"`
}

// Provides the Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role used to define an Amazon DynamoDB target endpoint.
//
// TODO: EXAMPLE
//
type CfnEndpoint_DynamoDbSettingsProperty struct {
	// The Amazon Resource Name (ARN) used by the service to access the IAM role.
	//
	// The role must allow the `iam:PassRole` action.
	ServiceAccessRoleArn *string `json:"serviceAccessRoleArn" yaml:"serviceAccessRoleArn"`
}

// Provides information that defines an OpenSearch endpoint.
//
// TODO: EXAMPLE
//
type CfnEndpoint_ElasticsearchSettingsProperty struct {
	// The endpoint for the OpenSearch cluster.
	//
	// AWS DMS uses HTTPS if a transport protocol (http/https) is not specified.
	EndpointUri *string `json:"endpointUri" yaml:"endpointUri"`
	// The maximum number of seconds for which DMS retries failed API requests to the OpenSearch cluster.
	ErrorRetryDuration *float64 `json:"errorRetryDuration" yaml:"errorRetryDuration"`
	// The maximum percentage of records that can fail to be written before a full load operation stops.
	//
	// To avoid early failure, this counter is only effective after 1000 records are transferred. OpenSearch also has the concept of error monitoring during the last 10 minutes of an Observation Window. If transfer of all records fail in the last 10 minutes, the full load operation stops.
	FullLoadErrorPercentage *float64 `json:"fullLoadErrorPercentage" yaml:"fullLoadErrorPercentage"`
	// The Amazon Resource Name (ARN) used by the service to access the IAM role.
	//
	// The role must allow the `iam:PassRole` action.
	ServiceAccessRoleArn *string `json:"serviceAccessRoleArn" yaml:"serviceAccessRoleArn"`
}

// Settings in JSON format for the source GCP MySQL endpoint.
//
// TODO: EXAMPLE
//
type CfnEndpoint_GcpMySQLSettingsProperty struct {
	// Specifies a script to run immediately after AWS DMS connects to the endpoint.
	//
	// The migration task continues running regardless if the SQL statement succeeds or fails.
	//
	// For this parameter, provide the code of the script itself, not the name of a file containing the script.
	AfterConnectScript *string `json:"afterConnectScript" yaml:"afterConnectScript"`
	// Adjusts the behavior of AWS DMS when migrating from an SQL Server source database that is hosted as part of an Always On availability group cluster.
	//
	// If you need AWS DMS to poll all the nodes in the Always On cluster for transaction backups, set this attribute to `false` .
	CleanSourceMetadataOnMismatch interface{} `json:"cleanSourceMetadataOnMismatch" yaml:"cleanSourceMetadataOnMismatch"`
	// Database name for the endpoint.
	//
	// For a MySQL source or target endpoint, don't explicitly specify the database using the `DatabaseName` request parameter on either the `CreateEndpoint` or `ModifyEndpoint` API call. Specifying `DatabaseName` when you create or modify a MySQL endpoint replicates all the task tables to this single database. For MySQL endpoints, you specify the database only when you specify the schema in the table-mapping rules of the AWS DMS task.
	DatabaseName *string `json:"databaseName" yaml:"databaseName"`
	// Specifies how often to check the binary log for new changes/events when the database is idle.
	//
	// The default is five seconds.
	//
	// Example: `eventsPollInterval=5;`
	//
	// In the example, AWS DMS checks for changes in the binary logs every five seconds.
	EventsPollInterval *float64 `json:"eventsPollInterval" yaml:"eventsPollInterval"`
	// Specifies the maximum size (in KB) of any .csv file used to transfer data to a MySQL-compatible database.
	//
	// Example: `maxFileSize=512`
	MaxFileSize *float64 `json:"maxFileSize" yaml:"maxFileSize"`
	// Improves performance when loading data into the MySQL-compatible target database.
	//
	// Specifies how many threads to use to load the data into the MySQL-compatible target database. Setting a large number of threads can have an adverse effect on database performance, because a separate connection is required for each thread. The default is one.
	//
	// Example: `parallelLoadThreads=1`
	ParallelLoadThreads *float64 `json:"parallelLoadThreads" yaml:"parallelLoadThreads"`
	// Endpoint connection password.
	Password *string `json:"password" yaml:"password"`
	// `CfnEndpoint.GcpMySQLSettingsProperty.Port`.
	Port *float64 `json:"port" yaml:"port"`
	// The full Amazon Resource Name (ARN) of the IAM role that specifies AWS DMS as the trusted entity and grants the required permissions to access the value in `SecretsManagerSecret.` The role must allow the `iam:PassRole` action. `SecretsManagerSecret` has the value of the AWS Secrets Manager secret that allows access to the MySQL endpoint.
	//
	// > You can specify one of two sets of values for these permissions. You can specify the values for this setting and `SecretsManagerSecretId` . Or you can specify clear-text values for `UserName` , `Password` , `ServerName` , and `Port` . You can't specify both. For more information on creating this `SecretsManagerSecret` and the `SecretsManagerAccessRoleArn` and `SecretsManagerSecretId` required to access it, see [Using secrets to access AWS Database Migration Service resources](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Security.html#security-iam-secretsmanager) in the AWS Database Migration Service User Guide.
	SecretsManagerAccessRoleArn *string `json:"secretsManagerAccessRoleArn" yaml:"secretsManagerAccessRoleArn"`
	// The full ARN, partial ARN, or friendly name of the `SecretsManagerSecret` that contains the MySQL endpoint connection details.
	SecretsManagerSecretId *string `json:"secretsManagerSecretId" yaml:"secretsManagerSecretId"`
	// Endpoint TCP port.
	ServerName *string `json:"serverName" yaml:"serverName"`
	// Specifies the time zone for the source MySQL database.
	//
	// Example: `serverTimezone=US/Pacific;`
	//
	// Note: Do not enclose time zones in single quotes.
	ServerTimezone *string `json:"serverTimezone" yaml:"serverTimezone"`
	// Endpoint connection user name.
	Username *string `json:"username" yaml:"username"`
}

// Not currently supported by AWS CloudFormation .
//
// TODO: EXAMPLE
//
type CfnEndpoint_IbmDb2SettingsProperty struct {
	// Not currently supported by AWS CloudFormation .
	SecretsManagerAccessRoleArn *string `json:"secretsManagerAccessRoleArn" yaml:"secretsManagerAccessRoleArn"`
	// Not currently supported by AWS CloudFormation .
	SecretsManagerSecretId *string `json:"secretsManagerSecretId" yaml:"secretsManagerSecretId"`
}

// Not currently supported by AWS CloudFormation .
//
// TODO: EXAMPLE
//
type CfnEndpoint_KafkaSettingsProperty struct {
	// Not currently supported by AWS CloudFormation .
	Broker *string `json:"broker" yaml:"broker"`
	// Shows detailed control information for table definition, column definition, and table and column changes in the Kafka message output.
	//
	// The default is `false` .
	IncludeControlDetails interface{} `json:"includeControlDetails" yaml:"includeControlDetails"`
	// Include NULL and empty columns for records migrated to the endpoint.
	//
	// The default is `false` .
	IncludeNullAndEmpty interface{} `json:"includeNullAndEmpty" yaml:"includeNullAndEmpty"`
	// Includes any data definition language (DDL) operations that change the table in the control data, such as `rename-table` , `drop-table` , `add-column` , `drop-column` , and `rename-column` .
	//
	// The default is `false` .
	IncludeTableAlterOperations interface{} `json:"includeTableAlterOperations" yaml:"includeTableAlterOperations"`
	// Provides detailed transaction information from the source database.
	//
	// This information includes a commit timestamp, a log position, and values for `transaction_id` , previous `transaction_id` , and `transaction_record_id` (the record offset within a transaction). The default is `false` .
	IncludeTransactionDetails interface{} `json:"includeTransactionDetails" yaml:"includeTransactionDetails"`
	// Set this optional parameter to `true` to avoid adding a '0x' prefix to raw data in hexadecimal format.
	//
	// For example, by default, AWS DMS adds a '0x' prefix to the LOB column type in hexadecimal format moving from an Oracle source to a Kafka target. Use the `NoHexPrefix` endpoint setting to enable migration of RAW data type columns without adding the '0x' prefix.
	NoHexPrefix interface{} `json:"noHexPrefix" yaml:"noHexPrefix"`
	// Prefixes schema and table names to partition values, when the partition type is `primary-key-type` .
	//
	// Doing this increases data distribution among Kafka partitions. For example, suppose that a SysBench schema has thousands of tables and each table has only limited range for a primary key. In this case, the same primary key is sent from thousands of tables to the same partition, which causes throttling. The default is `false` .
	PartitionIncludeSchemaTable interface{} `json:"partitionIncludeSchemaTable" yaml:"partitionIncludeSchemaTable"`
	// The secure password you created when you first set up your MSK cluster to validate a client identity and make an encrypted connection between server and client using SASL-SSL authentication.
	SaslPassword *string `json:"saslPassword" yaml:"saslPassword"`
	// The secure user name you created when you first set up your MSK cluster to validate a client identity and make an encrypted connection between server and client using SASL-SSL authentication.
	SaslUserName *string `json:"saslUserName" yaml:"saslUserName"`
	// Set secure connection to a Kafka target endpoint using Transport Layer Security (TLS).
	//
	// Options include `ssl-encryption` , `ssl-authentication` , and `sasl-ssl` . `sasl-ssl` requires `SaslUsername` and `SaslPassword` .
	SecurityProtocol *string `json:"securityProtocol" yaml:"securityProtocol"`
	// The Amazon Resource Name (ARN) for the private certificate authority (CA) cert that AWS DMS uses to securely connect to your Kafka target endpoint.
	SslCaCertificateArn *string `json:"sslCaCertificateArn" yaml:"sslCaCertificateArn"`
	// The Amazon Resource Name (ARN) of the client certificate used to securely connect to a Kafka target endpoint.
	SslClientCertificateArn *string `json:"sslClientCertificateArn" yaml:"sslClientCertificateArn"`
	// The Amazon Resource Name (ARN) for the client private key used to securely connect to a Kafka target endpoint.
	SslClientKeyArn *string `json:"sslClientKeyArn" yaml:"sslClientKeyArn"`
	// The password for the client private key used to securely connect to a Kafka target endpoint.
	SslClientKeyPassword *string `json:"sslClientKeyPassword" yaml:"sslClientKeyPassword"`
	// Not currently supported by AWS CloudFormation .
	Topic *string `json:"topic" yaml:"topic"`
}

// Not currently supported by AWS CloudFormation ..
//
// TODO: EXAMPLE
//
type CfnEndpoint_KinesisSettingsProperty struct {
	// Shows detailed control information for table definition, column definition, and table and column changes in the Kinesis message output.
	//
	// The default is `false` .
	IncludeControlDetails interface{} `json:"includeControlDetails" yaml:"includeControlDetails"`
	// Include NULL and empty columns for records migrated to the endpoint.
	//
	// The default is `false` .
	IncludeNullAndEmpty interface{} `json:"includeNullAndEmpty" yaml:"includeNullAndEmpty"`
	// Includes any data definition language (DDL) operations that change the table in the control data, such as `rename-table` , `drop-table` , `add-column` , `drop-column` , and `rename-column` .
	//
	// The default is `false` .
	IncludeTableAlterOperations interface{} `json:"includeTableAlterOperations" yaml:"includeTableAlterOperations"`
	// Provides detailed transaction information from the source database.
	//
	// This information includes a commit timestamp, a log position, and values for `transaction_id` , previous `transaction_id` , and `transaction_record_id` (the record offset within a transaction). The default is `false` .
	IncludeTransactionDetails interface{} `json:"includeTransactionDetails" yaml:"includeTransactionDetails"`
	// Not currently supported by AWS CloudFormation .
	MessageFormat *string `json:"messageFormat" yaml:"messageFormat"`
	// Set this optional parameter to `true` to avoid adding a '0x' prefix to raw data in hexadecimal format.
	//
	// For example, by default, AWS DMS adds a '0x' prefix to the LOB column type in hexadecimal format moving from an Oracle source to an Amazon Kinesis target. Use the `NoHexPrefix` endpoint setting to enable migration of RAW data type columns without adding the '0x' prefix.
	NoHexPrefix interface{} `json:"noHexPrefix" yaml:"noHexPrefix"`
	// Prefixes schema and table names to partition values, when the partition type is `primary-key-type` .
	//
	// Doing this increases data distribution among Kinesis shards. For example, suppose that a SysBench schema has thousands of tables and each table has only limited range for a primary key. In this case, the same primary key is sent from thousands of tables to the same shard, which causes throttling. The default is `false` .
	PartitionIncludeSchemaTable interface{} `json:"partitionIncludeSchemaTable" yaml:"partitionIncludeSchemaTable"`
	// Not currently supported by AWS CloudFormation .
	ServiceAccessRoleArn *string `json:"serviceAccessRoleArn" yaml:"serviceAccessRoleArn"`
	// Not currently supported by AWS CloudFormation .
	StreamArn *string `json:"streamArn" yaml:"streamArn"`
}

// Not currently supported by AWS CloudFormation .
//
// TODO: EXAMPLE
//
type CfnEndpoint_MicrosoftSqlServerSettingsProperty struct {
	// Not currently supported by AWS CloudFormation .
	SecretsManagerAccessRoleArn *string `json:"secretsManagerAccessRoleArn" yaml:"secretsManagerAccessRoleArn"`
	// Not currently supported by AWS CloudFormation .
	SecretsManagerSecretId *string `json:"secretsManagerSecretId" yaml:"secretsManagerSecretId"`
}

// Not currently supported by AWS CloudFormation .
//
// TODO: EXAMPLE
//
type CfnEndpoint_MongoDbSettingsProperty struct {
	// Not currently supported by AWS CloudFormation .
	AuthMechanism *string `json:"authMechanism" yaml:"authMechanism"`
	// Not currently supported by AWS CloudFormation .
	AuthSource *string `json:"authSource" yaml:"authSource"`
	// Not currently supported by AWS CloudFormation .
	AuthType *string `json:"authType" yaml:"authType"`
	// Not currently supported by AWS CloudFormation .
	DatabaseName *string `json:"databaseName" yaml:"databaseName"`
	// Not currently supported by AWS CloudFormation .
	DocsToInvestigate *string `json:"docsToInvestigate" yaml:"docsToInvestigate"`
	// Not currently supported by AWS CloudFormation .
	ExtractDocId *string `json:"extractDocId" yaml:"extractDocId"`
	// Not currently supported by AWS CloudFormation .
	NestingLevel *string `json:"nestingLevel" yaml:"nestingLevel"`
	// Not currently supported by AWS CloudFormation .
	Password *string `json:"password" yaml:"password"`
	// Not currently supported by AWS CloudFormation .
	Port *float64 `json:"port" yaml:"port"`
	// Not currently supported by AWS CloudFormation .
	SecretsManagerAccessRoleArn *string `json:"secretsManagerAccessRoleArn" yaml:"secretsManagerAccessRoleArn"`
	// Not currently supported by AWS CloudFormation .
	SecretsManagerSecretId *string `json:"secretsManagerSecretId" yaml:"secretsManagerSecretId"`
	// Not currently supported by AWS CloudFormation .
	ServerName *string `json:"serverName" yaml:"serverName"`
	// Not currently supported by AWS CloudFormation .
	Username *string `json:"username" yaml:"username"`
}

// Not currently supported by AWS CloudFormation .
//
// TODO: EXAMPLE
//
type CfnEndpoint_MySqlSettingsProperty struct {
	// Not currently supported by AWS CloudFormation .
	SecretsManagerAccessRoleArn *string `json:"secretsManagerAccessRoleArn" yaml:"secretsManagerAccessRoleArn"`
	// Not currently supported by AWS CloudFormation .
	SecretsManagerSecretId *string `json:"secretsManagerSecretId" yaml:"secretsManagerSecretId"`
}

// Provides information that defines an Amazon Neptune endpoint.
//
// TODO: EXAMPLE
//
type CfnEndpoint_NeptuneSettingsProperty struct {
	// The number of milliseconds for AWS DMS to wait to retry a bulk-load of migrated graph data to the Neptune target database before raising an error.
	//
	// The default is 250.
	ErrorRetryDuration *float64 `json:"errorRetryDuration" yaml:"errorRetryDuration"`
	// If you want AWS Identity and Access Management (IAM) authorization enabled for this endpoint, set this parameter to `true` .
	//
	// Then attach the appropriate IAM policy document to your service role specified by `ServiceAccessRoleArn` . The default is `false` .
	IamAuthEnabled interface{} `json:"iamAuthEnabled" yaml:"iamAuthEnabled"`
	// The maximum size in kilobytes of migrated graph data stored in a .csv file before AWS DMS bulk-loads the data to the Neptune target database. The default is 1,048,576 KB. If the bulk load is successful, AWS DMS clears the bucket, ready to store the next batch of migrated graph data.
	MaxFileSize *float64 `json:"maxFileSize" yaml:"maxFileSize"`
	// The number of times for AWS DMS to retry a bulk load of migrated graph data to the Neptune target database before raising an error.
	//
	// The default is 5.
	MaxRetryCount *float64 `json:"maxRetryCount" yaml:"maxRetryCount"`
	// A folder path where you want AWS DMS to store migrated graph data in the S3 bucket specified by `S3BucketName`.
	S3BucketFolder *string `json:"s3BucketFolder" yaml:"s3BucketFolder"`
	// The name of the Amazon S3 bucket where AWS DMS can temporarily store migrated graph data in .csv files before bulk-loading it to the Neptune target database. AWS DMS maps the SQL source data to graph data before storing it in these .csv files.
	S3BucketName *string `json:"s3BucketName" yaml:"s3BucketName"`
	// The Amazon Resource Name (ARN) of the service role that you created for the Neptune target endpoint.
	//
	// The role must allow the `iam:PassRole` action. For more information, see [Creating an IAM Service Role for Accessing Amazon Neptune as a Target](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Neptune.html#CHAP_Target.Neptune.ServiceRole) in the *AWS Database Migration Service User Guide.*
	ServiceAccessRoleArn *string `json:"serviceAccessRoleArn" yaml:"serviceAccessRoleArn"`
}

// Not currently supported by AWS CloudFormation .
//
// TODO: EXAMPLE
//
type CfnEndpoint_OracleSettingsProperty struct {
	// Not currently supported by AWS CloudFormation .
	SecretsManagerAccessRoleArn *string `json:"secretsManagerAccessRoleArn" yaml:"secretsManagerAccessRoleArn"`
	// Not currently supported by AWS CloudFormation .
	SecretsManagerOracleAsmAccessRoleArn *string `json:"secretsManagerOracleAsmAccessRoleArn" yaml:"secretsManagerOracleAsmAccessRoleArn"`
	// Not currently supported by AWS CloudFormation .
	SecretsManagerOracleAsmSecretId *string `json:"secretsManagerOracleAsmSecretId" yaml:"secretsManagerOracleAsmSecretId"`
	// Not currently supported by AWS CloudFormation .
	SecretsManagerSecretId *string `json:"secretsManagerSecretId" yaml:"secretsManagerSecretId"`
}

// Not currently supported by AWS CloudFormation .
//
// TODO: EXAMPLE
//
type CfnEndpoint_PostgreSqlSettingsProperty struct {
	// Not currently supported by AWS CloudFormation .
	SecretsManagerAccessRoleArn *string `json:"secretsManagerAccessRoleArn" yaml:"secretsManagerAccessRoleArn"`
	// Not currently supported by AWS CloudFormation .
	SecretsManagerSecretId *string `json:"secretsManagerSecretId" yaml:"secretsManagerSecretId"`
}

// Provides information that defines a Redis target endpoint.
//
// TODO: EXAMPLE
//
type CfnEndpoint_RedisSettingsProperty struct {
	// The password provided with the `auth-role` and `auth-token` options of the `AuthType` setting for a Redis target endpoint.
	AuthPassword *string `json:"authPassword" yaml:"authPassword"`
	// The type of authentication to perform when connecting to a Redis target.
	//
	// Options include `none` , `auth-token` , and `auth-role` . The `auth-token` option requires an `AuthPassword` value to be provided. The `auth-role` option requires `AuthUserName` and `AuthPassword` values to be provided.
	AuthType *string `json:"authType" yaml:"authType"`
	// The user name provided with the `auth-role` option of the `AuthType` setting for a Redis target endpoint.
	AuthUserName *string `json:"authUserName" yaml:"authUserName"`
	// Transmission Control Protocol (TCP) port for the endpoint.
	Port *float64 `json:"port" yaml:"port"`
	// Fully qualified domain name of the endpoint.
	ServerName *string `json:"serverName" yaml:"serverName"`
	// The Amazon Resource Name (ARN) for the certificate authority (CA) that DMS uses to connect to your Redis target endpoint.
	SslCaCertificateArn *string `json:"sslCaCertificateArn" yaml:"sslCaCertificateArn"`
	// The connection to a Redis target endpoint using Transport Layer Security (TLS).
	//
	// Valid values include `plaintext` and `ssl-encryption` . The default is `ssl-encryption` . The `ssl-encryption` option makes an encrypted connection. Optionally, you can identify an Amazon Resource Name (ARN) for an SSL certificate authority (CA) using the `SslCaCertificateArn` setting. If an ARN isn't given for a CA, DMS uses the Amazon root CA.
	//
	// The `plaintext` option doesn't provide Transport Layer Security (TLS) encryption for traffic between endpoint and database.
	SslSecurityProtocol *string `json:"sslSecurityProtocol" yaml:"sslSecurityProtocol"`
}

// Not currently supported by AWS CloudFormation .
//
// TODO: EXAMPLE
//
type CfnEndpoint_RedshiftSettingsProperty struct {
	// Not currently supported by AWS CloudFormation .
	SecretsManagerAccessRoleArn *string `json:"secretsManagerAccessRoleArn" yaml:"secretsManagerAccessRoleArn"`
	// Not currently supported by AWS CloudFormation .
	SecretsManagerSecretId *string `json:"secretsManagerSecretId" yaml:"secretsManagerSecretId"`
}

// Not currently supported by AWS CloudFormation .
//
// TODO: EXAMPLE
//
type CfnEndpoint_S3SettingsProperty struct {
	// Not currently supported by AWS CloudFormation .
	BucketFolder *string `json:"bucketFolder" yaml:"bucketFolder"`
	// Not currently supported by AWS CloudFormation .
	BucketName *string `json:"bucketName" yaml:"bucketName"`
	// Not currently supported by AWS CloudFormation .
	CompressionType *string `json:"compressionType" yaml:"compressionType"`
	// Not currently supported by AWS CloudFormation .
	CsvDelimiter *string `json:"csvDelimiter" yaml:"csvDelimiter"`
	// Not currently supported by AWS CloudFormation .
	CsvRowDelimiter *string `json:"csvRowDelimiter" yaml:"csvRowDelimiter"`
	// Not currently supported by AWS CloudFormation .
	ExternalTableDefinition *string `json:"externalTableDefinition" yaml:"externalTableDefinition"`
	// Not currently supported by AWS CloudFormation .
	ServiceAccessRoleArn *string `json:"serviceAccessRoleArn" yaml:"serviceAccessRoleArn"`
}

// Not currently supported by AWS CloudFormation .
//
// TODO: EXAMPLE
//
type CfnEndpoint_SybaseSettingsProperty struct {
	// Not currently supported by AWS CloudFormation .
	SecretsManagerAccessRoleArn *string `json:"secretsManagerAccessRoleArn" yaml:"secretsManagerAccessRoleArn"`
	// Not currently supported by AWS CloudFormation .
	SecretsManagerSecretId *string `json:"secretsManagerSecretId" yaml:"secretsManagerSecretId"`
}

// Properties for defining a `CfnEndpoint`.
//
// TODO: EXAMPLE
//
type CfnEndpointProps struct {
	// The type of endpoint.
	//
	// Valid values are `source` and `target` .
	EndpointType *string `json:"endpointType" yaml:"endpointType"`
	// The type of engine for the endpoint.
	//
	// Valid values, depending on the `EndpointType` value, include `"mysql"` , `"oracle"` , `"postgres"` , `"mariadb"` , `"aurora"` , `"aurora-postgresql"` , `"opensearch"` , `"redshift"` , `"s3"` , `"db2"` , `"azuredb"` , `"sybase"` , `"dynamodb"` , `"mongodb"` , `"kinesis"` , `"kafka"` , `"elasticsearch"` , `"docdb"` , `"sqlserver"` , and `"neptune"` .
	EngineName *string `json:"engineName" yaml:"engineName"`
	// The Amazon Resource Name (ARN) for the certificate.
	CertificateArn *string `json:"certificateArn" yaml:"certificateArn"`
	// The name of the endpoint database.
	//
	// For a MySQL source or target endpoint, do not specify DatabaseName. To migrate to a specific database, use this setting and `targetDbType` .
	DatabaseName *string `json:"databaseName" yaml:"databaseName"`
	// Settings in JSON format for the source DocumentDB endpoint.
	//
	// For more information about the available settings, see the configuration properties section in [Using DocumentDB as a Target for AWS Database Migration Service](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.DocumentDB.html) in the *AWS Database Migration Service User Guide.*
	DocDbSettings interface{} `json:"docDbSettings" yaml:"docDbSettings"`
	// Settings in JSON format for the target Amazon DynamoDB endpoint.
	//
	// For information about other available settings, see [Using Object Mapping to Migrate Data to DynamoDB](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.DynamoDB.html#CHAP_Target.DynamoDB.ObjectMapping) in the *AWS Database Migration Service User Guide.*
	DynamoDbSettings interface{} `json:"dynamoDbSettings" yaml:"dynamoDbSettings"`
	// Settings in JSON format for the target OpenSearch endpoint.
	//
	// For more information about the available settings, see [Extra Connection Attributes When Using OpenSearch as a Target for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Elasticsearch.html#CHAP_Target.Elasticsearch.Configuration) in the *AWS Database Migration Service User Guide* .
	ElasticsearchSettings interface{} `json:"elasticsearchSettings" yaml:"elasticsearchSettings"`
	// The database endpoint identifier.
	//
	// Identifiers must begin with a letter and must contain only ASCII letters, digits, and hyphens. They can't end with a hyphen, or contain two consecutive hyphens.
	EndpointIdentifier *string `json:"endpointIdentifier" yaml:"endpointIdentifier"`
	// Additional attributes associated with the connection.
	//
	// Each attribute is specified as a name-value pair associated by an equal sign (=). Multiple attributes are separated by a semicolon (;) with no additional white space. For information on the attributes available for connecting your source or target endpoint, see [Working with AWS DMS Endpoints](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Endpoints.html) in the *AWS Database Migration Service User Guide.*
	ExtraConnectionAttributes *string `json:"extraConnectionAttributes" yaml:"extraConnectionAttributes"`
	// Settings in JSON format for the source GCP MySQL endpoint.
	GcpMySqlSettings interface{} `json:"gcpMySqlSettings" yaml:"gcpMySqlSettings"`
	// Not currently supported by AWS CloudFormation .
	IbmDb2Settings interface{} `json:"ibmDb2Settings" yaml:"ibmDb2Settings"`
	// Settings in JSON format for the target Apache Kafka endpoint.
	//
	// For more information about the available settings, see [Using object mapping to migrate data to a Kafka topic](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Kafka.html#CHAP_Target.Kafka.ObjectMapping) in the *AWS Database Migration Service User Guide.*
	KafkaSettings interface{} `json:"kafkaSettings" yaml:"kafkaSettings"`
	// Settings in JSON format for the target endpoint for Amazon Kinesis Data Streams.
	//
	// For more information about the available settings, see [Using Amazon Kinesis Data Streams as a Target for AWS Database Migration Service](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Kinesis.html) in the *AWS Database Migration Service User Guide.*
	KinesisSettings interface{} `json:"kinesisSettings" yaml:"kinesisSettings"`
	// An AWS KMS key identifier that is used to encrypt the connection parameters for the endpoint.
	//
	// If you don't specify a value for the `KmsKeyId` parameter, then AWS DMS uses your default encryption key.
	//
	// AWS KMS creates the default encryption key for your AWS account . Your AWS account has a different default encryption key for each AWS Region .
	KmsKeyId *string `json:"kmsKeyId" yaml:"kmsKeyId"`
	// Not currently supported by AWS CloudFormation .
	MicrosoftSqlServerSettings interface{} `json:"microsoftSqlServerSettings" yaml:"microsoftSqlServerSettings"`
	// Not currently supported by AWS CloudFormation .
	MongoDbSettings interface{} `json:"mongoDbSettings" yaml:"mongoDbSettings"`
	// Settings in JSON format for the source and target MySQL endpoint.
	//
	// For information about other available settings, see [Extra connection attributes when using MySQL as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.MySQL.html#CHAP_Source.MySQL.ConnectionAttrib) and [Extra connection attributes when using a MySQL-compatible database as a target for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.MySQL.html#CHAP_Target.MySQL.ConnectionAttrib) in the *AWS Database Migration Service User Guide.*
	MySqlSettings interface{} `json:"mySqlSettings" yaml:"mySqlSettings"`
	// `AWS::DMS::Endpoint.NeptuneSettings`.
	NeptuneSettings interface{} `json:"neptuneSettings" yaml:"neptuneSettings"`
	// Settings in JSON format for the source and target Oracle endpoint.
	//
	// For information about other available settings, see [Extra connection attributes when using Oracle as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.Oracle.html#CHAP_Source.Oracle.ConnectionAttrib) and [Extra connection attributes when using Oracle as a target for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Oracle.html#CHAP_Target.Oracle.ConnectionAttrib) in the *AWS Database Migration Service User Guide.*
	OracleSettings interface{} `json:"oracleSettings" yaml:"oracleSettings"`
	// The password to be used to log in to the endpoint database.
	Password *string `json:"password" yaml:"password"`
	// The port used by the endpoint database.
	Port *float64 `json:"port" yaml:"port"`
	// Not currently supported by AWS CloudFormation .
	PostgreSqlSettings interface{} `json:"postgreSqlSettings" yaml:"postgreSqlSettings"`
	// Settings in JSON format for the target Redis endpoint.
	RedisSettings interface{} `json:"redisSettings" yaml:"redisSettings"`
	// Not currently supported by AWS CloudFormation .
	RedshiftSettings interface{} `json:"redshiftSettings" yaml:"redshiftSettings"`
	// A friendly name for the resource identifier at the end of the `EndpointArn` response parameter that is returned in the created `Endpoint` object.
	//
	// The value for this parameter can have up to 31 characters. It can contain only ASCII letters, digits, and hyphen ('-'). Also, it can't end with a hyphen or contain two consecutive hyphens, and can only begin with a letter, such as `Example-App-ARN1` . For example, this value might result in the `EndpointArn` value `arn:aws:dms:eu-west-1:012345678901:rep:Example-App-ARN1` . If you don't specify a `ResourceIdentifier` value, AWS DMS generates a default identifier value for the end of `EndpointArn` .
	ResourceIdentifier *string `json:"resourceIdentifier" yaml:"resourceIdentifier"`
	// Settings in JSON format for the target Amazon S3 endpoint.
	//
	// For more information about the available settings, see [Extra Connection Attributes When Using Amazon S3 as a Target for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.S3.html#CHAP_Target.S3.Configuring) in the *AWS Database Migration Service User Guide.*
	S3Settings interface{} `json:"s3Settings" yaml:"s3Settings"`
	// The name of the server where the endpoint database resides.
	ServerName *string `json:"serverName" yaml:"serverName"`
	// The Secure Sockets Layer (SSL) mode to use for the SSL connection. The default is `none` .
	//
	// > When `engine_name` is set to S3, then the only allowed value is `none` .
	SslMode *string `json:"sslMode" yaml:"sslMode"`
	// Settings in JSON format for the source and target SAP ASE endpoint.
	//
	// For information about other available settings, see [Extra connection attributes when using SAP ASE as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.SAP.html#CHAP_Source.SAP.ConnectionAttrib) and [Extra connection attributes when using SAP ASE as a target for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.SAP.html#CHAP_Target.SAP.ConnectionAttrib) in the *AWS Database Migration Service User Guide.*
	SybaseSettings interface{} `json:"sybaseSettings" yaml:"sybaseSettings"`
	// One or more tags to be assigned to the endpoint.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
	// The user name to be used to log in to the endpoint database.
	Username *string `json:"username" yaml:"username"`
}

// A CloudFormation `AWS::DMS::EventSubscription`.
//
// Use the `AWS::DMS::EventSubscription` resource to get notifications for AWS Database Migration Service events through the Amazon Simple Notification Service. For more information, see [Using AWS DMS Event Notification](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Events.html) in the *AWS Database Migration Service User Guide.*
//
// TODO: EXAMPLE
//
type CfnEventSubscription interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	Enabled() interface{}
	SetEnabled(val interface{})
	EventCategories() *[]*string
	SetEventCategories(val *[]*string)
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	SnsTopicArn() *string
	SetSnsTopicArn(val *string)
	SourceIds() *[]*string
	SetSourceIds(val *[]*string)
	SourceType() *string
	SetSourceType(val *string)
	Stack() awscdk.Stack
	SubscriptionName() *string
	SetSubscriptionName(val *string)
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

// The jsii proxy struct for CfnEventSubscription
type jsiiProxy_CfnEventSubscription struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnEventSubscription) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSubscription) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSubscription) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSubscription) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSubscription) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSubscription) EventCategories() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"eventCategories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSubscription) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSubscription) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSubscription) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSubscription) SnsTopicArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snsTopicArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSubscription) SourceIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSubscription) SourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSubscription) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSubscription) SubscriptionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscriptionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSubscription) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSubscription) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::DMS::EventSubscription`.
func NewCfnEventSubscription(scope constructs.Construct, id *string, props *CfnEventSubscriptionProps) CfnEventSubscription {
	_init_.Initialize()

	j := jsiiProxy_CfnEventSubscription{}

	_jsii_.Create(
		"aws-cdk-lib.aws_dms.CfnEventSubscription",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::DMS::EventSubscription`.
func NewCfnEventSubscription_Override(c CfnEventSubscription, scope constructs.Construct, id *string, props *CfnEventSubscriptionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_dms.CfnEventSubscription",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnEventSubscription) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_CfnEventSubscription) SetEventCategories(val *[]*string) {
	_jsii_.Set(
		j,
		"eventCategories",
		val,
	)
}

func (j *jsiiProxy_CfnEventSubscription) SetSnsTopicArn(val *string) {
	_jsii_.Set(
		j,
		"snsTopicArn",
		val,
	)
}

func (j *jsiiProxy_CfnEventSubscription) SetSourceIds(val *[]*string) {
	_jsii_.Set(
		j,
		"sourceIds",
		val,
	)
}

func (j *jsiiProxy_CfnEventSubscription) SetSourceType(val *string) {
	_jsii_.Set(
		j,
		"sourceType",
		val,
	)
}

func (j *jsiiProxy_CfnEventSubscription) SetSubscriptionName(val *string) {
	_jsii_.Set(
		j,
		"subscriptionName",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnEventSubscription_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_dms.CfnEventSubscription",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnEventSubscription_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_dms.CfnEventSubscription",
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
func CfnEventSubscription_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_dms.CfnEventSubscription",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEventSubscription_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_dms.CfnEventSubscription",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnEventSubscription) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnEventSubscription) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnEventSubscription) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnEventSubscription) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnEventSubscription) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnEventSubscription) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnEventSubscription) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnEventSubscription) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnEventSubscription) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnEventSubscription) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnEventSubscription) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnEventSubscription) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnEventSubscription) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnEventSubscription) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEventSubscription) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnEventSubscription`.
//
// TODO: EXAMPLE
//
type CfnEventSubscriptionProps struct {
	// The Amazon Resource Name (ARN) of the Amazon SNS topic created for event notification.
	//
	// The ARN is created by Amazon SNS when you create a topic and subscribe to it.
	SnsTopicArn *string `json:"snsTopicArn" yaml:"snsTopicArn"`
	// Indicates whether to activate the subscription.
	//
	// If you don't specify this property, AWS CloudFormation activates the subscription.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
	// A list of event categories for a source type that you want to subscribe to.
	//
	// If you don't specify this property, you are notified about all event categories. For more information, see [Working with Events and Notifications](https://docs.aws.amazon.com//dms/latest/userguide/CHAP_Events.html) in the *AWS DMS User Guide* .
	EventCategories *[]*string `json:"eventCategories" yaml:"eventCategories"`
	// A list of identifiers for which AWS DMS provides notification events.
	//
	// If you don't specify a value, notifications are provided for all sources.
	//
	// If you specify multiple values, they must be of the same type. For example, if you specify a database instance ID, then all of the other values must be database instance IDs.
	SourceIds *[]*string `json:"sourceIds" yaml:"sourceIds"`
	// The type of AWS DMS resource that generates the events.
	//
	// For example, if you want to be notified of events generated by a replication instance, you set this parameter to `replication-instance` . If this value isn't specified, all events are returned.
	//
	// Valid values: `replication-instance` | `replication-task`
	SourceType *string `json:"sourceType" yaml:"sourceType"`
	// The name of the AWS DMS event notification subscription.
	//
	// This name must be less than 255 characters.
	SubscriptionName *string `json:"subscriptionName" yaml:"subscriptionName"`
	// One or more tags to be assigned to the event subscription.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::DMS::ReplicationInstance`.
//
// The `AWS::DMS::ReplicationInstance` resource creates an AWS DMS replication instance.
//
// TODO: EXAMPLE
//
type CfnReplicationInstance interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AllocatedStorage() *float64
	SetAllocatedStorage(val *float64)
	AllowMajorVersionUpgrade() interface{}
	SetAllowMajorVersionUpgrade(val interface{})
	AttrReplicationInstancePrivateIpAddresses() *[]*string
	AttrReplicationInstancePublicIpAddresses() *[]*string
	AutoMinorVersionUpgrade() interface{}
	SetAutoMinorVersionUpgrade(val interface{})
	AvailabilityZone() *string
	SetAvailabilityZone(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	EngineVersion() *string
	SetEngineVersion(val *string)
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	LogicalId() *string
	MultiAz() interface{}
	SetMultiAz(val interface{})
	Node() constructs.Node
	PreferredMaintenanceWindow() *string
	SetPreferredMaintenanceWindow(val *string)
	PubliclyAccessible() interface{}
	SetPubliclyAccessible(val interface{})
	Ref() *string
	ReplicationInstanceClass() *string
	SetReplicationInstanceClass(val *string)
	ReplicationInstanceIdentifier() *string
	SetReplicationInstanceIdentifier(val *string)
	ReplicationSubnetGroupIdentifier() *string
	SetReplicationSubnetGroupIdentifier(val *string)
	ResourceIdentifier() *string
	SetResourceIdentifier(val *string)
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	VpcSecurityGroupIds() *[]*string
	SetVpcSecurityGroupIds(val *[]*string)
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

// The jsii proxy struct for CfnReplicationInstance
type jsiiProxy_CfnReplicationInstance struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnReplicationInstance) AllocatedStorage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocatedStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationInstance) AllowMajorVersionUpgrade() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowMajorVersionUpgrade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationInstance) AttrReplicationInstancePrivateIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrReplicationInstancePrivateIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationInstance) AttrReplicationInstancePublicIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrReplicationInstancePublicIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationInstance) AutoMinorVersionUpgrade() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoMinorVersionUpgrade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationInstance) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationInstance) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationInstance) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationInstance) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationInstance) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationInstance) EngineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationInstance) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationInstance) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationInstance) MultiAz() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiAz",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationInstance) PreferredMaintenanceWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredMaintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationInstance) PubliclyAccessible() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publiclyAccessible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationInstance) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationInstance) ReplicationInstanceClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationInstanceClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationInstance) ReplicationInstanceIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationInstanceIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationInstance) ReplicationSubnetGroupIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationSubnetGroupIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationInstance) ResourceIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationInstance) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationInstance) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationInstance) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationInstance) VpcSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroupIds",
		&returns,
	)
	return returns
}


// Create a new `AWS::DMS::ReplicationInstance`.
func NewCfnReplicationInstance(scope constructs.Construct, id *string, props *CfnReplicationInstanceProps) CfnReplicationInstance {
	_init_.Initialize()

	j := jsiiProxy_CfnReplicationInstance{}

	_jsii_.Create(
		"aws-cdk-lib.aws_dms.CfnReplicationInstance",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::DMS::ReplicationInstance`.
func NewCfnReplicationInstance_Override(c CfnReplicationInstance, scope constructs.Construct, id *string, props *CfnReplicationInstanceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_dms.CfnReplicationInstance",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnReplicationInstance) SetAllocatedStorage(val *float64) {
	_jsii_.Set(
		j,
		"allocatedStorage",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationInstance) SetAllowMajorVersionUpgrade(val interface{}) {
	_jsii_.Set(
		j,
		"allowMajorVersionUpgrade",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationInstance) SetAutoMinorVersionUpgrade(val interface{}) {
	_jsii_.Set(
		j,
		"autoMinorVersionUpgrade",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationInstance) SetAvailabilityZone(val *string) {
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationInstance) SetEngineVersion(val *string) {
	_jsii_.Set(
		j,
		"engineVersion",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationInstance) SetKmsKeyId(val *string) {
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationInstance) SetMultiAz(val interface{}) {
	_jsii_.Set(
		j,
		"multiAz",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationInstance) SetPreferredMaintenanceWindow(val *string) {
	_jsii_.Set(
		j,
		"preferredMaintenanceWindow",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationInstance) SetPubliclyAccessible(val interface{}) {
	_jsii_.Set(
		j,
		"publiclyAccessible",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationInstance) SetReplicationInstanceClass(val *string) {
	_jsii_.Set(
		j,
		"replicationInstanceClass",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationInstance) SetReplicationInstanceIdentifier(val *string) {
	_jsii_.Set(
		j,
		"replicationInstanceIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationInstance) SetReplicationSubnetGroupIdentifier(val *string) {
	_jsii_.Set(
		j,
		"replicationSubnetGroupIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationInstance) SetResourceIdentifier(val *string) {
	_jsii_.Set(
		j,
		"resourceIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationInstance) SetVpcSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"vpcSecurityGroupIds",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnReplicationInstance_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_dms.CfnReplicationInstance",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnReplicationInstance_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_dms.CfnReplicationInstance",
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
func CfnReplicationInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_dms.CfnReplicationInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnReplicationInstance_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_dms.CfnReplicationInstance",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnReplicationInstance) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnReplicationInstance) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnReplicationInstance) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnReplicationInstance) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnReplicationInstance) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnReplicationInstance) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnReplicationInstance) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnReplicationInstance) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnReplicationInstance) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnReplicationInstance) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnReplicationInstance) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnReplicationInstance) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnReplicationInstance) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnReplicationInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnReplicationInstance) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnReplicationInstance`.
//
// TODO: EXAMPLE
//
type CfnReplicationInstanceProps struct {
	// The compute and memory capacity of the replication instance as defined for the specified replication instance class.
	//
	// For example to specify the instance class dms.c4.large, set this parameter to `"dms.c4.large"` .
	//
	// For more information on the settings and capacities for the available replication instance classes, see [Selecting the right AWS DMS replication instance for your migration](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_ReplicationInstance.html#CHAP_ReplicationInstance.InDepth) .
	ReplicationInstanceClass *string `json:"replicationInstanceClass" yaml:"replicationInstanceClass"`
	// The amount of storage (in gigabytes) to be initially allocated for the replication instance.
	AllocatedStorage *float64 `json:"allocatedStorage" yaml:"allocatedStorage"`
	// Indicates that major version upgrades are allowed.
	//
	// Changing this parameter does not result in an outage, and the change is asynchronously applied as soon as possible.
	//
	// This parameter must be set to `true` when specifying a value for the `EngineVersion` parameter that is a different major version than the replication instance's current version.
	AllowMajorVersionUpgrade interface{} `json:"allowMajorVersionUpgrade" yaml:"allowMajorVersionUpgrade"`
	// A value that indicates whether minor engine upgrades are applied automatically to the replication instance during the maintenance window.
	//
	// This parameter defaults to `true` .
	//
	// Default: `true`
	AutoMinorVersionUpgrade interface{} `json:"autoMinorVersionUpgrade" yaml:"autoMinorVersionUpgrade"`
	// The Availability Zone that the replication instance will be created in.
	//
	// The default value is a random, system-chosen Availability Zone in the endpoint's AWS Region , for example: `us-east-1d`
	AvailabilityZone *string `json:"availabilityZone" yaml:"availabilityZone"`
	// The engine version number of the replication instance.
	//
	// If an engine version number is not specified when a replication instance is created, the default is the latest engine version available.
	EngineVersion *string `json:"engineVersion" yaml:"engineVersion"`
	// An AWS KMS key identifier that is used to encrypt the data on the replication instance.
	//
	// If you don't specify a value for the `KmsKeyId` parameter, then AWS DMS uses your default encryption key.
	//
	// AWS KMS creates the default encryption key for your AWS account . Your AWS account has a different default encryption key for each AWS Region .
	KmsKeyId *string `json:"kmsKeyId" yaml:"kmsKeyId"`
	// Specifies whether the replication instance is a Multi-AZ deployment.
	//
	// You can't set the `AvailabilityZone` parameter if the Multi-AZ parameter is set to `true` .
	MultiAz interface{} `json:"multiAz" yaml:"multiAz"`
	// The weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).
	//
	// Format: `ddd:hh24:mi-ddd:hh24:mi`
	//
	// Default: A 30-minute window selected at random from an 8-hour block of time per AWS Region , occurring on a random day of the week.
	//
	// Valid Days: Mon, Tue, Wed, Thu, Fri, Sat, Sun
	//
	// Constraints: Minimum 30-minute window.
	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// Specifies the accessibility options for the replication instance.
	//
	// A value of `true` represents an instance with a public IP address. A value of `false` represents an instance with a private IP address. The default value is `true` .
	PubliclyAccessible interface{} `json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// The replication instance identifier. This parameter is stored as a lowercase string.
	//
	// Constraints:
	//
	// - Must contain 1-63 alphanumeric characters or hyphens.
	// - First character must be a letter.
	// - Can't end with a hyphen or contain two consecutive hyphens.
	//
	// Example: `myrepinstance`
	ReplicationInstanceIdentifier *string `json:"replicationInstanceIdentifier" yaml:"replicationInstanceIdentifier"`
	// A subnet group to associate with the replication instance.
	ReplicationSubnetGroupIdentifier *string `json:"replicationSubnetGroupIdentifier" yaml:"replicationSubnetGroupIdentifier"`
	// A friendly name for the resource identifier at the end of the `EndpointArn` response parameter that is returned in the created `Endpoint` object.
	//
	// The value for this parameter can have up to 31 characters. It can contain only ASCII letters, digits, and hyphen ('-'). Also, it can't end with a hyphen or contain two consecutive hyphens, and can only begin with a letter, such as `Example-App-ARN1` . For example, this value might result in the `EndpointArn` value `arn:aws:dms:eu-west-1:012345678901:rep:Example-App-ARN1` . If you don't specify a `ResourceIdentifier` value, AWS DMS generates a default identifier value for the end of `EndpointArn` .
	ResourceIdentifier *string `json:"resourceIdentifier" yaml:"resourceIdentifier"`
	// One or more tags to be assigned to the replication instance.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
	// Specifies the VPC security group to be used with the replication instance.
	//
	// The VPC security group must work with the VPC containing the replication instance.
	VpcSecurityGroupIds *[]*string `json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
}

// A CloudFormation `AWS::DMS::ReplicationSubnetGroup`.
//
// The `AWS::DMS::ReplicationSubnetGroup` resource creates an AWS DMS replication subnet group. Subnet groups must contain at least two subnets in two different Availability Zones in the same region.
//
// > Resource creation will fail if the `dms-vpc-role` IAM role doesn't already exist. For more information, see [Creating the IAM Roles to Use With the AWS CLI and AWS DMS API](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Security.APIRole.html) in the *AWS Database Migration Service User Guide.*
//
// TODO: EXAMPLE
//
type CfnReplicationSubnetGroup interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	ReplicationSubnetGroupDescription() *string
	SetReplicationSubnetGroupDescription(val *string)
	ReplicationSubnetGroupIdentifier() *string
	SetReplicationSubnetGroupIdentifier(val *string)
	Stack() awscdk.Stack
	SubnetIds() *[]*string
	SetSubnetIds(val *[]*string)
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

// The jsii proxy struct for CfnReplicationSubnetGroup
type jsiiProxy_CfnReplicationSubnetGroup struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnReplicationSubnetGroup) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationSubnetGroup) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationSubnetGroup) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationSubnetGroup) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationSubnetGroup) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationSubnetGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationSubnetGroup) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationSubnetGroup) ReplicationSubnetGroupDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationSubnetGroupDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationSubnetGroup) ReplicationSubnetGroupIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationSubnetGroupIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationSubnetGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationSubnetGroup) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationSubnetGroup) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationSubnetGroup) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::DMS::ReplicationSubnetGroup`.
func NewCfnReplicationSubnetGroup(scope constructs.Construct, id *string, props *CfnReplicationSubnetGroupProps) CfnReplicationSubnetGroup {
	_init_.Initialize()

	j := jsiiProxy_CfnReplicationSubnetGroup{}

	_jsii_.Create(
		"aws-cdk-lib.aws_dms.CfnReplicationSubnetGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::DMS::ReplicationSubnetGroup`.
func NewCfnReplicationSubnetGroup_Override(c CfnReplicationSubnetGroup, scope constructs.Construct, id *string, props *CfnReplicationSubnetGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_dms.CfnReplicationSubnetGroup",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnReplicationSubnetGroup) SetReplicationSubnetGroupDescription(val *string) {
	_jsii_.Set(
		j,
		"replicationSubnetGroupDescription",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationSubnetGroup) SetReplicationSubnetGroupIdentifier(val *string) {
	_jsii_.Set(
		j,
		"replicationSubnetGroupIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationSubnetGroup) SetSubnetIds(val *[]*string) {
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnReplicationSubnetGroup_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_dms.CfnReplicationSubnetGroup",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnReplicationSubnetGroup_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_dms.CfnReplicationSubnetGroup",
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
func CfnReplicationSubnetGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_dms.CfnReplicationSubnetGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnReplicationSubnetGroup_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_dms.CfnReplicationSubnetGroup",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnReplicationSubnetGroup) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnReplicationSubnetGroup) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnReplicationSubnetGroup) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnReplicationSubnetGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnReplicationSubnetGroup) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnReplicationSubnetGroup) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnReplicationSubnetGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnReplicationSubnetGroup) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnReplicationSubnetGroup) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnReplicationSubnetGroup) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnReplicationSubnetGroup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnReplicationSubnetGroup) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnReplicationSubnetGroup) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnReplicationSubnetGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnReplicationSubnetGroup) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnReplicationSubnetGroup`.
//
// TODO: EXAMPLE
//
type CfnReplicationSubnetGroupProps struct {
	// The description for the subnet group.
	ReplicationSubnetGroupDescription *string `json:"replicationSubnetGroupDescription" yaml:"replicationSubnetGroupDescription"`
	// One or more subnet IDs to be assigned to the subnet group.
	SubnetIds *[]*string `json:"subnetIds" yaml:"subnetIds"`
	// The identifier for the replication subnet group.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique ID and uses that ID for the identifier.
	ReplicationSubnetGroupIdentifier *string `json:"replicationSubnetGroupIdentifier" yaml:"replicationSubnetGroupIdentifier"`
	// One or more tags to be assigned to the subnet group.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::DMS::ReplicationTask`.
//
// The `AWS::DMS::ReplicationTask` resource creates an AWS DMS replication task.
//
// TODO: EXAMPLE
//
type CfnReplicationTask interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CdcStartPosition() *string
	SetCdcStartPosition(val *string)
	CdcStartTime() *float64
	SetCdcStartTime(val *float64)
	CdcStopPosition() *string
	SetCdcStopPosition(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	MigrationType() *string
	SetMigrationType(val *string)
	Node() constructs.Node
	Ref() *string
	ReplicationInstanceArn() *string
	SetReplicationInstanceArn(val *string)
	ReplicationTaskIdentifier() *string
	SetReplicationTaskIdentifier(val *string)
	ReplicationTaskSettings() *string
	SetReplicationTaskSettings(val *string)
	ResourceIdentifier() *string
	SetResourceIdentifier(val *string)
	SourceEndpointArn() *string
	SetSourceEndpointArn(val *string)
	Stack() awscdk.Stack
	TableMappings() *string
	SetTableMappings(val *string)
	Tags() awscdk.TagManager
	TargetEndpointArn() *string
	SetTargetEndpointArn(val *string)
	TaskData() *string
	SetTaskData(val *string)
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

// The jsii proxy struct for CfnReplicationTask
type jsiiProxy_CfnReplicationTask struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnReplicationTask) CdcStartPosition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdcStartPosition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationTask) CdcStartTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cdcStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationTask) CdcStopPosition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdcStopPosition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationTask) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationTask) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationTask) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationTask) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationTask) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationTask) MigrationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"migrationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationTask) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationTask) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationTask) ReplicationInstanceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationInstanceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationTask) ReplicationTaskIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationTaskIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationTask) ReplicationTaskSettings() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationTaskSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationTask) ResourceIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationTask) SourceEndpointArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceEndpointArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationTask) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationTask) TableMappings() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationTask) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationTask) TargetEndpointArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetEndpointArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationTask) TaskData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationTask) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::DMS::ReplicationTask`.
func NewCfnReplicationTask(scope constructs.Construct, id *string, props *CfnReplicationTaskProps) CfnReplicationTask {
	_init_.Initialize()

	j := jsiiProxy_CfnReplicationTask{}

	_jsii_.Create(
		"aws-cdk-lib.aws_dms.CfnReplicationTask",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::DMS::ReplicationTask`.
func NewCfnReplicationTask_Override(c CfnReplicationTask, scope constructs.Construct, id *string, props *CfnReplicationTaskProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_dms.CfnReplicationTask",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnReplicationTask) SetCdcStartPosition(val *string) {
	_jsii_.Set(
		j,
		"cdcStartPosition",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationTask) SetCdcStartTime(val *float64) {
	_jsii_.Set(
		j,
		"cdcStartTime",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationTask) SetCdcStopPosition(val *string) {
	_jsii_.Set(
		j,
		"cdcStopPosition",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationTask) SetMigrationType(val *string) {
	_jsii_.Set(
		j,
		"migrationType",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationTask) SetReplicationInstanceArn(val *string) {
	_jsii_.Set(
		j,
		"replicationInstanceArn",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationTask) SetReplicationTaskIdentifier(val *string) {
	_jsii_.Set(
		j,
		"replicationTaskIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationTask) SetReplicationTaskSettings(val *string) {
	_jsii_.Set(
		j,
		"replicationTaskSettings",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationTask) SetResourceIdentifier(val *string) {
	_jsii_.Set(
		j,
		"resourceIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationTask) SetSourceEndpointArn(val *string) {
	_jsii_.Set(
		j,
		"sourceEndpointArn",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationTask) SetTableMappings(val *string) {
	_jsii_.Set(
		j,
		"tableMappings",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationTask) SetTargetEndpointArn(val *string) {
	_jsii_.Set(
		j,
		"targetEndpointArn",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationTask) SetTaskData(val *string) {
	_jsii_.Set(
		j,
		"taskData",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnReplicationTask_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_dms.CfnReplicationTask",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnReplicationTask_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_dms.CfnReplicationTask",
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
func CfnReplicationTask_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_dms.CfnReplicationTask",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnReplicationTask_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_dms.CfnReplicationTask",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnReplicationTask) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnReplicationTask) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnReplicationTask) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnReplicationTask) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnReplicationTask) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnReplicationTask) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnReplicationTask) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnReplicationTask) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnReplicationTask) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnReplicationTask) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnReplicationTask) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnReplicationTask) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnReplicationTask) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnReplicationTask) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnReplicationTask) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnReplicationTask`.
//
// TODO: EXAMPLE
//
type CfnReplicationTaskProps struct {
	// The migration type.
	//
	// Valid values: `full-load` | `cdc` | `full-load-and-cdc`
	MigrationType *string `json:"migrationType" yaml:"migrationType"`
	// The Amazon Resource Name (ARN) of a replication instance.
	ReplicationInstanceArn *string `json:"replicationInstanceArn" yaml:"replicationInstanceArn"`
	// An Amazon Resource Name (ARN) that uniquely identifies the source endpoint.
	SourceEndpointArn *string `json:"sourceEndpointArn" yaml:"sourceEndpointArn"`
	// The table mappings for the task, in JSON format.
	//
	// For more information, see [Using Table Mapping to Specify Task Settings](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TableMapping.html) in the *AWS Database Migration Service User Guide.*
	TableMappings *string `json:"tableMappings" yaml:"tableMappings"`
	// An Amazon Resource Name (ARN) that uniquely identifies the target endpoint.
	TargetEndpointArn *string `json:"targetEndpointArn" yaml:"targetEndpointArn"`
	// Indicates when you want a change data capture (CDC) operation to start.
	//
	// Use either CdcStartPosition or CdcStartTime to specify when you want a CDC operation to start. Specifying both values results in an error.
	//
	// The value can be in date, checkpoint, or LSN/SCN format.
	//
	// Date Example: --cdc-start-position “2018-03-08T12:12:12”
	//
	// Checkpoint Example: --cdc-start-position "checkpoint:V1#27#mysql-bin-changelog.157832:1975:-1:2002:677883278264080:mysql-bin-changelog.157832:1876#0#0#*#0#93"
	//
	// LSN Example: --cdc-start-position “mysql-bin-changelog.000024:373”
	//
	// > When you use this task setting with a source PostgreSQL database, a logical replication slot should already be created and associated with the source endpoint. You can verify this by setting the `slotName` extra connection attribute to the name of this logical replication slot. For more information, see [Extra Connection Attributes When Using PostgreSQL as a Source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.PostgreSQL.html#CHAP_Source.PostgreSQL.ConnectionAttrib) .
	CdcStartPosition *string `json:"cdcStartPosition" yaml:"cdcStartPosition"`
	// Indicates the start time for a change data capture (CDC) operation.
	CdcStartTime *float64 `json:"cdcStartTime" yaml:"cdcStartTime"`
	// Indicates when you want a change data capture (CDC) operation to stop.
	//
	// The value can be either server time or commit time.
	//
	// Server time example: --cdc-stop-position “server_time:2018-02-09T12:12:12”
	//
	// Commit time example: --cdc-stop-position “commit_time: 2018-02-09T12:12:12 “
	CdcStopPosition *string `json:"cdcStopPosition" yaml:"cdcStopPosition"`
	// An identifier for the replication task.
	//
	// Constraints:
	//
	// - Must contain 1-255 alphanumeric characters or hyphens.
	// - First character must be a letter.
	// - Cannot end with a hyphen or contain two consecutive hyphens.
	ReplicationTaskIdentifier *string `json:"replicationTaskIdentifier" yaml:"replicationTaskIdentifier"`
	// Overall settings for the task, in JSON format.
	//
	// For more information, see [Specifying Task Settings for AWS Database Migration Service Tasks](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TaskSettings.html) in the *AWS Database Migration Service User Guide.*
	ReplicationTaskSettings *string `json:"replicationTaskSettings" yaml:"replicationTaskSettings"`
	// A friendly name for the resource identifier at the end of the `EndpointArn` response parameter that is returned in the created `Endpoint` object.
	//
	// The value for this parameter can have up to 31 characters. It can contain only ASCII letters, digits, and hyphen ('-'). Also, it can't end with a hyphen or contain two consecutive hyphens, and can only begin with a letter, such as `Example-App-ARN1` . For example, this value might result in the `EndpointArn` value `arn:aws:dms:eu-west-1:012345678901:rep:Example-App-ARN1` . If you don't specify a `ResourceIdentifier` value, AWS DMS generates a default identifier value for the end of `EndpointArn` .
	ResourceIdentifier *string `json:"resourceIdentifier" yaml:"resourceIdentifier"`
	// One or more tags to be assigned to the replication task.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
	// `AWS::DMS::ReplicationTask.TaskData`.
	TaskData *string `json:"taskData" yaml:"taskData"`
}

