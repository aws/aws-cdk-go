package awsdms

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsdms/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::DMS::Certificate`.
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
	Node() awscdk.ConstructNode
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
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
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

func (j *jsiiProxy_CfnCertificate) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
func NewCfnCertificate(scope awscdk.Construct, id *string, props *CfnCertificateProps) CfnCertificate {
	_init_.Initialize()

	j := jsiiProxy_CfnCertificate{}

	_jsii_.Create(
		"monocdk.aws_dms.CfnCertificate",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::DMS::Certificate`.
func NewCfnCertificate_Override(c CfnCertificate, scope awscdk.Construct, id *string, props *CfnCertificateProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_dms.CfnCertificate",
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
// Experimental.
func CfnCertificate_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_dms.CfnCertificate",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnCertificate_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_dms.CfnCertificate",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnCertificate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_dms.CfnCertificate",
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
		"monocdk.aws_dms.CfnCertificate",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnCertificate) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnCertificate) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnCertificate) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnCertificate) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnCertificate) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnCertificate) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
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
// Experimental.
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

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnCertificate) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
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

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnCertificate) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnCertificate) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `AWS::DMS::Certificate`.
type CfnCertificateProps struct {
	// `AWS::DMS::Certificate.CertificateIdentifier`.
	CertificateIdentifier *string `json:"certificateIdentifier"`
	// `AWS::DMS::Certificate.CertificatePem`.
	CertificatePem *string `json:"certificatePem"`
	// `AWS::DMS::Certificate.CertificateWallet`.
	CertificateWallet *string `json:"certificateWallet"`
}

// A CloudFormation `AWS::DMS::Endpoint`.
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
	Node() awscdk.ConstructNode
	OracleSettings() interface{}
	SetOracleSettings(val interface{})
	Password() *string
	SetPassword(val *string)
	Port() *float64
	SetPort(val *float64)
	PostgreSqlSettings() interface{}
	SetPostgreSqlSettings(val interface{})
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
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
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

func (j *jsiiProxy_CfnEndpoint) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
func NewCfnEndpoint(scope awscdk.Construct, id *string, props *CfnEndpointProps) CfnEndpoint {
	_init_.Initialize()

	j := jsiiProxy_CfnEndpoint{}

	_jsii_.Create(
		"monocdk.aws_dms.CfnEndpoint",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::DMS::Endpoint`.
func NewCfnEndpoint_Override(c CfnEndpoint, scope awscdk.Construct, id *string, props *CfnEndpointProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_dms.CfnEndpoint",
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
// Experimental.
func CfnEndpoint_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_dms.CfnEndpoint",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnEndpoint_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_dms.CfnEndpoint",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnEndpoint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_dms.CfnEndpoint",
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
		"monocdk.aws_dms.CfnEndpoint",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnEndpoint) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnEndpoint) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnEndpoint) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnEndpoint) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnEndpoint) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnEndpoint) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
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
// Experimental.
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

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnEndpoint) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
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

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnEndpoint) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnEndpoint) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnEndpoint_DocDbSettingsProperty struct {
	// `CfnEndpoint.DocDbSettingsProperty.SecretsManagerAccessRoleArn`.
	SecretsManagerAccessRoleArn *string `json:"secretsManagerAccessRoleArn"`
	// `CfnEndpoint.DocDbSettingsProperty.SecretsManagerSecretId`.
	SecretsManagerSecretId *string `json:"secretsManagerSecretId"`
}

type CfnEndpoint_DynamoDbSettingsProperty struct {
	// `CfnEndpoint.DynamoDbSettingsProperty.ServiceAccessRoleArn`.
	ServiceAccessRoleArn *string `json:"serviceAccessRoleArn"`
}

type CfnEndpoint_ElasticsearchSettingsProperty struct {
	// `CfnEndpoint.ElasticsearchSettingsProperty.EndpointUri`.
	EndpointUri *string `json:"endpointUri"`
	// `CfnEndpoint.ElasticsearchSettingsProperty.ErrorRetryDuration`.
	ErrorRetryDuration *float64 `json:"errorRetryDuration"`
	// `CfnEndpoint.ElasticsearchSettingsProperty.FullLoadErrorPercentage`.
	FullLoadErrorPercentage *float64 `json:"fullLoadErrorPercentage"`
	// `CfnEndpoint.ElasticsearchSettingsProperty.ServiceAccessRoleArn`.
	ServiceAccessRoleArn *string `json:"serviceAccessRoleArn"`
}

type CfnEndpoint_IbmDb2SettingsProperty struct {
	// `CfnEndpoint.IbmDb2SettingsProperty.SecretsManagerAccessRoleArn`.
	SecretsManagerAccessRoleArn *string `json:"secretsManagerAccessRoleArn"`
	// `CfnEndpoint.IbmDb2SettingsProperty.SecretsManagerSecretId`.
	SecretsManagerSecretId *string `json:"secretsManagerSecretId"`
}

type CfnEndpoint_KafkaSettingsProperty struct {
	// `CfnEndpoint.KafkaSettingsProperty.Broker`.
	Broker *string `json:"broker"`
	// `CfnEndpoint.KafkaSettingsProperty.Topic`.
	Topic *string `json:"topic"`
}

type CfnEndpoint_KinesisSettingsProperty struct {
	// `CfnEndpoint.KinesisSettingsProperty.MessageFormat`.
	MessageFormat *string `json:"messageFormat"`
	// `CfnEndpoint.KinesisSettingsProperty.ServiceAccessRoleArn`.
	ServiceAccessRoleArn *string `json:"serviceAccessRoleArn"`
	// `CfnEndpoint.KinesisSettingsProperty.StreamArn`.
	StreamArn *string `json:"streamArn"`
}

type CfnEndpoint_MicrosoftSqlServerSettingsProperty struct {
	// `CfnEndpoint.MicrosoftSqlServerSettingsProperty.SecretsManagerAccessRoleArn`.
	SecretsManagerAccessRoleArn *string `json:"secretsManagerAccessRoleArn"`
	// `CfnEndpoint.MicrosoftSqlServerSettingsProperty.SecretsManagerSecretId`.
	SecretsManagerSecretId *string `json:"secretsManagerSecretId"`
}

type CfnEndpoint_MongoDbSettingsProperty struct {
	// `CfnEndpoint.MongoDbSettingsProperty.AuthMechanism`.
	AuthMechanism *string `json:"authMechanism"`
	// `CfnEndpoint.MongoDbSettingsProperty.AuthSource`.
	AuthSource *string `json:"authSource"`
	// `CfnEndpoint.MongoDbSettingsProperty.AuthType`.
	AuthType *string `json:"authType"`
	// `CfnEndpoint.MongoDbSettingsProperty.DatabaseName`.
	DatabaseName *string `json:"databaseName"`
	// `CfnEndpoint.MongoDbSettingsProperty.DocsToInvestigate`.
	DocsToInvestigate *string `json:"docsToInvestigate"`
	// `CfnEndpoint.MongoDbSettingsProperty.ExtractDocId`.
	ExtractDocId *string `json:"extractDocId"`
	// `CfnEndpoint.MongoDbSettingsProperty.NestingLevel`.
	NestingLevel *string `json:"nestingLevel"`
	// `CfnEndpoint.MongoDbSettingsProperty.Password`.
	Password *string `json:"password"`
	// `CfnEndpoint.MongoDbSettingsProperty.Port`.
	Port *float64 `json:"port"`
	// `CfnEndpoint.MongoDbSettingsProperty.SecretsManagerAccessRoleArn`.
	SecretsManagerAccessRoleArn *string `json:"secretsManagerAccessRoleArn"`
	// `CfnEndpoint.MongoDbSettingsProperty.SecretsManagerSecretId`.
	SecretsManagerSecretId *string `json:"secretsManagerSecretId"`
	// `CfnEndpoint.MongoDbSettingsProperty.ServerName`.
	ServerName *string `json:"serverName"`
	// `CfnEndpoint.MongoDbSettingsProperty.Username`.
	Username *string `json:"username"`
}

type CfnEndpoint_MySqlSettingsProperty struct {
	// `CfnEndpoint.MySqlSettingsProperty.SecretsManagerAccessRoleArn`.
	SecretsManagerAccessRoleArn *string `json:"secretsManagerAccessRoleArn"`
	// `CfnEndpoint.MySqlSettingsProperty.SecretsManagerSecretId`.
	SecretsManagerSecretId *string `json:"secretsManagerSecretId"`
}

type CfnEndpoint_NeptuneSettingsProperty struct {
	// `CfnEndpoint.NeptuneSettingsProperty.ErrorRetryDuration`.
	ErrorRetryDuration *float64 `json:"errorRetryDuration"`
	// `CfnEndpoint.NeptuneSettingsProperty.IamAuthEnabled`.
	IamAuthEnabled interface{} `json:"iamAuthEnabled"`
	// `CfnEndpoint.NeptuneSettingsProperty.MaxFileSize`.
	MaxFileSize *float64 `json:"maxFileSize"`
	// `CfnEndpoint.NeptuneSettingsProperty.MaxRetryCount`.
	MaxRetryCount *float64 `json:"maxRetryCount"`
	// `CfnEndpoint.NeptuneSettingsProperty.S3BucketFolder`.
	S3BucketFolder *string `json:"s3BucketFolder"`
	// `CfnEndpoint.NeptuneSettingsProperty.S3BucketName`.
	S3BucketName *string `json:"s3BucketName"`
	// `CfnEndpoint.NeptuneSettingsProperty.ServiceAccessRoleArn`.
	ServiceAccessRoleArn *string `json:"serviceAccessRoleArn"`
}

type CfnEndpoint_OracleSettingsProperty struct {
	// `CfnEndpoint.OracleSettingsProperty.SecretsManagerAccessRoleArn`.
	SecretsManagerAccessRoleArn *string `json:"secretsManagerAccessRoleArn"`
	// `CfnEndpoint.OracleSettingsProperty.SecretsManagerOracleAsmAccessRoleArn`.
	SecretsManagerOracleAsmAccessRoleArn *string `json:"secretsManagerOracleAsmAccessRoleArn"`
	// `CfnEndpoint.OracleSettingsProperty.SecretsManagerOracleAsmSecretId`.
	SecretsManagerOracleAsmSecretId *string `json:"secretsManagerOracleAsmSecretId"`
	// `CfnEndpoint.OracleSettingsProperty.SecretsManagerSecretId`.
	SecretsManagerSecretId *string `json:"secretsManagerSecretId"`
}

type CfnEndpoint_PostgreSqlSettingsProperty struct {
	// `CfnEndpoint.PostgreSqlSettingsProperty.SecretsManagerAccessRoleArn`.
	SecretsManagerAccessRoleArn *string `json:"secretsManagerAccessRoleArn"`
	// `CfnEndpoint.PostgreSqlSettingsProperty.SecretsManagerSecretId`.
	SecretsManagerSecretId *string `json:"secretsManagerSecretId"`
}

type CfnEndpoint_RedshiftSettingsProperty struct {
	// `CfnEndpoint.RedshiftSettingsProperty.SecretsManagerAccessRoleArn`.
	SecretsManagerAccessRoleArn *string `json:"secretsManagerAccessRoleArn"`
	// `CfnEndpoint.RedshiftSettingsProperty.SecretsManagerSecretId`.
	SecretsManagerSecretId *string `json:"secretsManagerSecretId"`
}

type CfnEndpoint_S3SettingsProperty struct {
	// `CfnEndpoint.S3SettingsProperty.BucketFolder`.
	BucketFolder *string `json:"bucketFolder"`
	// `CfnEndpoint.S3SettingsProperty.BucketName`.
	BucketName *string `json:"bucketName"`
	// `CfnEndpoint.S3SettingsProperty.CompressionType`.
	CompressionType *string `json:"compressionType"`
	// `CfnEndpoint.S3SettingsProperty.CsvDelimiter`.
	CsvDelimiter *string `json:"csvDelimiter"`
	// `CfnEndpoint.S3SettingsProperty.CsvRowDelimiter`.
	CsvRowDelimiter *string `json:"csvRowDelimiter"`
	// `CfnEndpoint.S3SettingsProperty.ExternalTableDefinition`.
	ExternalTableDefinition *string `json:"externalTableDefinition"`
	// `CfnEndpoint.S3SettingsProperty.ServiceAccessRoleArn`.
	ServiceAccessRoleArn *string `json:"serviceAccessRoleArn"`
}

type CfnEndpoint_SybaseSettingsProperty struct {
	// `CfnEndpoint.SybaseSettingsProperty.SecretsManagerAccessRoleArn`.
	SecretsManagerAccessRoleArn *string `json:"secretsManagerAccessRoleArn"`
	// `CfnEndpoint.SybaseSettingsProperty.SecretsManagerSecretId`.
	SecretsManagerSecretId *string `json:"secretsManagerSecretId"`
}

// Properties for defining a `AWS::DMS::Endpoint`.
type CfnEndpointProps struct {
	// `AWS::DMS::Endpoint.EndpointType`.
	EndpointType *string `json:"endpointType"`
	// `AWS::DMS::Endpoint.EngineName`.
	EngineName *string `json:"engineName"`
	// `AWS::DMS::Endpoint.CertificateArn`.
	CertificateArn *string `json:"certificateArn"`
	// `AWS::DMS::Endpoint.DatabaseName`.
	DatabaseName *string `json:"databaseName"`
	// `AWS::DMS::Endpoint.DocDbSettings`.
	DocDbSettings interface{} `json:"docDbSettings"`
	// `AWS::DMS::Endpoint.DynamoDbSettings`.
	DynamoDbSettings interface{} `json:"dynamoDbSettings"`
	// `AWS::DMS::Endpoint.ElasticsearchSettings`.
	ElasticsearchSettings interface{} `json:"elasticsearchSettings"`
	// `AWS::DMS::Endpoint.EndpointIdentifier`.
	EndpointIdentifier *string `json:"endpointIdentifier"`
	// `AWS::DMS::Endpoint.ExtraConnectionAttributes`.
	ExtraConnectionAttributes *string `json:"extraConnectionAttributes"`
	// `AWS::DMS::Endpoint.IbmDb2Settings`.
	IbmDb2Settings interface{} `json:"ibmDb2Settings"`
	// `AWS::DMS::Endpoint.KafkaSettings`.
	KafkaSettings interface{} `json:"kafkaSettings"`
	// `AWS::DMS::Endpoint.KinesisSettings`.
	KinesisSettings interface{} `json:"kinesisSettings"`
	// `AWS::DMS::Endpoint.KmsKeyId`.
	KmsKeyId *string `json:"kmsKeyId"`
	// `AWS::DMS::Endpoint.MicrosoftSqlServerSettings`.
	MicrosoftSqlServerSettings interface{} `json:"microsoftSqlServerSettings"`
	// `AWS::DMS::Endpoint.MongoDbSettings`.
	MongoDbSettings interface{} `json:"mongoDbSettings"`
	// `AWS::DMS::Endpoint.MySqlSettings`.
	MySqlSettings interface{} `json:"mySqlSettings"`
	// `AWS::DMS::Endpoint.NeptuneSettings`.
	NeptuneSettings interface{} `json:"neptuneSettings"`
	// `AWS::DMS::Endpoint.OracleSettings`.
	OracleSettings interface{} `json:"oracleSettings"`
	// `AWS::DMS::Endpoint.Password`.
	Password *string `json:"password"`
	// `AWS::DMS::Endpoint.Port`.
	Port *float64 `json:"port"`
	// `AWS::DMS::Endpoint.PostgreSqlSettings`.
	PostgreSqlSettings interface{} `json:"postgreSqlSettings"`
	// `AWS::DMS::Endpoint.RedshiftSettings`.
	RedshiftSettings interface{} `json:"redshiftSettings"`
	// `AWS::DMS::Endpoint.ResourceIdentifier`.
	ResourceIdentifier *string `json:"resourceIdentifier"`
	// `AWS::DMS::Endpoint.S3Settings`.
	S3Settings interface{} `json:"s3Settings"`
	// `AWS::DMS::Endpoint.ServerName`.
	ServerName *string `json:"serverName"`
	// `AWS::DMS::Endpoint.SslMode`.
	SslMode *string `json:"sslMode"`
	// `AWS::DMS::Endpoint.SybaseSettings`.
	SybaseSettings interface{} `json:"sybaseSettings"`
	// `AWS::DMS::Endpoint.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
	// `AWS::DMS::Endpoint.Username`.
	Username *string `json:"username"`
}

// A CloudFormation `AWS::DMS::EventSubscription`.
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
	Node() awscdk.ConstructNode
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
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
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

func (j *jsiiProxy_CfnEventSubscription) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
func NewCfnEventSubscription(scope awscdk.Construct, id *string, props *CfnEventSubscriptionProps) CfnEventSubscription {
	_init_.Initialize()

	j := jsiiProxy_CfnEventSubscription{}

	_jsii_.Create(
		"monocdk.aws_dms.CfnEventSubscription",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::DMS::EventSubscription`.
func NewCfnEventSubscription_Override(c CfnEventSubscription, scope awscdk.Construct, id *string, props *CfnEventSubscriptionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_dms.CfnEventSubscription",
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
// Experimental.
func CfnEventSubscription_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_dms.CfnEventSubscription",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnEventSubscription_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_dms.CfnEventSubscription",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnEventSubscription_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_dms.CfnEventSubscription",
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
		"monocdk.aws_dms.CfnEventSubscription",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnEventSubscription) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnEventSubscription) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnEventSubscription) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnEventSubscription) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnEventSubscription) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnEventSubscription) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
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
// Experimental.
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

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnEventSubscription) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
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

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnEventSubscription) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnEventSubscription) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `AWS::DMS::EventSubscription`.
type CfnEventSubscriptionProps struct {
	// `AWS::DMS::EventSubscription.SnsTopicArn`.
	SnsTopicArn *string `json:"snsTopicArn"`
	// `AWS::DMS::EventSubscription.Enabled`.
	Enabled interface{} `json:"enabled"`
	// `AWS::DMS::EventSubscription.EventCategories`.
	EventCategories *[]*string `json:"eventCategories"`
	// `AWS::DMS::EventSubscription.SourceIds`.
	SourceIds *[]*string `json:"sourceIds"`
	// `AWS::DMS::EventSubscription.SourceType`.
	SourceType *string `json:"sourceType"`
	// `AWS::DMS::EventSubscription.SubscriptionName`.
	SubscriptionName *string `json:"subscriptionName"`
	// `AWS::DMS::EventSubscription.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
}

// A CloudFormation `AWS::DMS::ReplicationInstance`.
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
	Node() awscdk.ConstructNode
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
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
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

func (j *jsiiProxy_CfnReplicationInstance) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
func NewCfnReplicationInstance(scope awscdk.Construct, id *string, props *CfnReplicationInstanceProps) CfnReplicationInstance {
	_init_.Initialize()

	j := jsiiProxy_CfnReplicationInstance{}

	_jsii_.Create(
		"monocdk.aws_dms.CfnReplicationInstance",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::DMS::ReplicationInstance`.
func NewCfnReplicationInstance_Override(c CfnReplicationInstance, scope awscdk.Construct, id *string, props *CfnReplicationInstanceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_dms.CfnReplicationInstance",
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
// Experimental.
func CfnReplicationInstance_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_dms.CfnReplicationInstance",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnReplicationInstance_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_dms.CfnReplicationInstance",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnReplicationInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_dms.CfnReplicationInstance",
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
		"monocdk.aws_dms.CfnReplicationInstance",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnReplicationInstance) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnReplicationInstance) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnReplicationInstance) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnReplicationInstance) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnReplicationInstance) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnReplicationInstance) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
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
// Experimental.
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

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnReplicationInstance) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
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

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnReplicationInstance) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnReplicationInstance) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `AWS::DMS::ReplicationInstance`.
type CfnReplicationInstanceProps struct {
	// `AWS::DMS::ReplicationInstance.ReplicationInstanceClass`.
	ReplicationInstanceClass *string `json:"replicationInstanceClass"`
	// `AWS::DMS::ReplicationInstance.AllocatedStorage`.
	AllocatedStorage *float64 `json:"allocatedStorage"`
	// `AWS::DMS::ReplicationInstance.AllowMajorVersionUpgrade`.
	AllowMajorVersionUpgrade interface{} `json:"allowMajorVersionUpgrade"`
	// `AWS::DMS::ReplicationInstance.AutoMinorVersionUpgrade`.
	AutoMinorVersionUpgrade interface{} `json:"autoMinorVersionUpgrade"`
	// `AWS::DMS::ReplicationInstance.AvailabilityZone`.
	AvailabilityZone *string `json:"availabilityZone"`
	// `AWS::DMS::ReplicationInstance.EngineVersion`.
	EngineVersion *string `json:"engineVersion"`
	// `AWS::DMS::ReplicationInstance.KmsKeyId`.
	KmsKeyId *string `json:"kmsKeyId"`
	// `AWS::DMS::ReplicationInstance.MultiAZ`.
	MultiAz interface{} `json:"multiAz"`
	// `AWS::DMS::ReplicationInstance.PreferredMaintenanceWindow`.
	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow"`
	// `AWS::DMS::ReplicationInstance.PubliclyAccessible`.
	PubliclyAccessible interface{} `json:"publiclyAccessible"`
	// `AWS::DMS::ReplicationInstance.ReplicationInstanceIdentifier`.
	ReplicationInstanceIdentifier *string `json:"replicationInstanceIdentifier"`
	// `AWS::DMS::ReplicationInstance.ReplicationSubnetGroupIdentifier`.
	ReplicationSubnetGroupIdentifier *string `json:"replicationSubnetGroupIdentifier"`
	// `AWS::DMS::ReplicationInstance.ResourceIdentifier`.
	ResourceIdentifier *string `json:"resourceIdentifier"`
	// `AWS::DMS::ReplicationInstance.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
	// `AWS::DMS::ReplicationInstance.VpcSecurityGroupIds`.
	VpcSecurityGroupIds *[]*string `json:"vpcSecurityGroupIds"`
}

// A CloudFormation `AWS::DMS::ReplicationSubnetGroup`.
type CfnReplicationSubnetGroup interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	Node() awscdk.ConstructNode
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
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
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

func (j *jsiiProxy_CfnReplicationSubnetGroup) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
func NewCfnReplicationSubnetGroup(scope awscdk.Construct, id *string, props *CfnReplicationSubnetGroupProps) CfnReplicationSubnetGroup {
	_init_.Initialize()

	j := jsiiProxy_CfnReplicationSubnetGroup{}

	_jsii_.Create(
		"monocdk.aws_dms.CfnReplicationSubnetGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::DMS::ReplicationSubnetGroup`.
func NewCfnReplicationSubnetGroup_Override(c CfnReplicationSubnetGroup, scope awscdk.Construct, id *string, props *CfnReplicationSubnetGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_dms.CfnReplicationSubnetGroup",
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
// Experimental.
func CfnReplicationSubnetGroup_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_dms.CfnReplicationSubnetGroup",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnReplicationSubnetGroup_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_dms.CfnReplicationSubnetGroup",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnReplicationSubnetGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_dms.CfnReplicationSubnetGroup",
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
		"monocdk.aws_dms.CfnReplicationSubnetGroup",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnReplicationSubnetGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnReplicationSubnetGroup) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnReplicationSubnetGroup) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnReplicationSubnetGroup) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnReplicationSubnetGroup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnReplicationSubnetGroup) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
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
// Experimental.
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

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnReplicationSubnetGroup) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
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

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnReplicationSubnetGroup) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnReplicationSubnetGroup) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `AWS::DMS::ReplicationSubnetGroup`.
type CfnReplicationSubnetGroupProps struct {
	// `AWS::DMS::ReplicationSubnetGroup.ReplicationSubnetGroupDescription`.
	ReplicationSubnetGroupDescription *string `json:"replicationSubnetGroupDescription"`
	// `AWS::DMS::ReplicationSubnetGroup.SubnetIds`.
	SubnetIds *[]*string `json:"subnetIds"`
	// `AWS::DMS::ReplicationSubnetGroup.ReplicationSubnetGroupIdentifier`.
	ReplicationSubnetGroupIdentifier *string `json:"replicationSubnetGroupIdentifier"`
	// `AWS::DMS::ReplicationSubnetGroup.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
}

// A CloudFormation `AWS::DMS::ReplicationTask`.
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
	Node() awscdk.ConstructNode
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
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
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

func (j *jsiiProxy_CfnReplicationTask) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
func NewCfnReplicationTask(scope awscdk.Construct, id *string, props *CfnReplicationTaskProps) CfnReplicationTask {
	_init_.Initialize()

	j := jsiiProxy_CfnReplicationTask{}

	_jsii_.Create(
		"monocdk.aws_dms.CfnReplicationTask",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::DMS::ReplicationTask`.
func NewCfnReplicationTask_Override(c CfnReplicationTask, scope awscdk.Construct, id *string, props *CfnReplicationTaskProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_dms.CfnReplicationTask",
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
// Experimental.
func CfnReplicationTask_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_dms.CfnReplicationTask",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnReplicationTask_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_dms.CfnReplicationTask",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnReplicationTask_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_dms.CfnReplicationTask",
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
		"monocdk.aws_dms.CfnReplicationTask",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnReplicationTask) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnReplicationTask) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnReplicationTask) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnReplicationTask) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnReplicationTask) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnReplicationTask) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
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
// Experimental.
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

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnReplicationTask) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
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

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnReplicationTask) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnReplicationTask) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `AWS::DMS::ReplicationTask`.
type CfnReplicationTaskProps struct {
	// `AWS::DMS::ReplicationTask.MigrationType`.
	MigrationType *string `json:"migrationType"`
	// `AWS::DMS::ReplicationTask.ReplicationInstanceArn`.
	ReplicationInstanceArn *string `json:"replicationInstanceArn"`
	// `AWS::DMS::ReplicationTask.SourceEndpointArn`.
	SourceEndpointArn *string `json:"sourceEndpointArn"`
	// `AWS::DMS::ReplicationTask.TableMappings`.
	TableMappings *string `json:"tableMappings"`
	// `AWS::DMS::ReplicationTask.TargetEndpointArn`.
	TargetEndpointArn *string `json:"targetEndpointArn"`
	// `AWS::DMS::ReplicationTask.CdcStartPosition`.
	CdcStartPosition *string `json:"cdcStartPosition"`
	// `AWS::DMS::ReplicationTask.CdcStartTime`.
	CdcStartTime *float64 `json:"cdcStartTime"`
	// `AWS::DMS::ReplicationTask.CdcStopPosition`.
	CdcStopPosition *string `json:"cdcStopPosition"`
	// `AWS::DMS::ReplicationTask.ReplicationTaskIdentifier`.
	ReplicationTaskIdentifier *string `json:"replicationTaskIdentifier"`
	// `AWS::DMS::ReplicationTask.ReplicationTaskSettings`.
	ReplicationTaskSettings *string `json:"replicationTaskSettings"`
	// `AWS::DMS::ReplicationTask.ResourceIdentifier`.
	ResourceIdentifier *string `json:"resourceIdentifier"`
	// `AWS::DMS::ReplicationTask.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
	// `AWS::DMS::ReplicationTask.TaskData`.
	TaskData *string `json:"taskData"`
}

