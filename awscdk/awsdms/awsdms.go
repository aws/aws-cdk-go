package awsdms

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsdms/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::DMS::Certificate`.
//
// The `AWS::DMS::Certificate` resource creates an Secure Sockets Layer (SSL) certificate that encrypts connections between AWS DMS endpoints and the replication instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCertificate := awscdk.Aws_dms.NewCfnCertificate(this, jsii.String("MyCfnCertificate"), &cfnCertificateProps{
//   	certificateIdentifier: jsii.String("certificateIdentifier"),
//   	certificatePem: jsii.String("certificatePem"),
//   	certificateWallet: jsii.String("certificateWallet"),
//   })
//
type CfnCertificate interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// A customer-assigned name for the certificate.
	//
	// Identifiers must begin with a letter and must contain only ASCII letters, digits, and hyphens. They can't end with a hyphen or contain two consecutive hyphens.
	CertificateIdentifier() *string
	SetCertificateIdentifier(val *string)
	// The contents of a `.pem` file, which contains an X.509 certificate.
	CertificatePem() *string
	SetCertificatePem(val *string)
	// The location of an imported Oracle Wallet certificate for use with SSL.
	//
	// An example is: `filebase64("${path.root}/rds-ca-2019-root.sso")`
	CertificateWallet() *string
	SetCertificateWallet(val *string)
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

func (c *jsiiProxy_CfnCertificate) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnCertificate) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCertificate) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnCertificate) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnCertificate) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnCertificate) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnCertificate) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

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

func (c *jsiiProxy_CfnCertificate) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnCertificate) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnCertificate) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_CfnCertificate) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

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

func (c *jsiiProxy_CfnCertificate) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_CfnCertificate) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnCertificate`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCertificateProps := &cfnCertificateProps{
//   	certificateIdentifier: jsii.String("certificateIdentifier"),
//   	certificatePem: jsii.String("certificatePem"),
//   	certificateWallet: jsii.String("certificateWallet"),
//   }
//
type CfnCertificateProps struct {
	// A customer-assigned name for the certificate.
	//
	// Identifiers must begin with a letter and must contain only ASCII letters, digits, and hyphens. They can't end with a hyphen or contain two consecutive hyphens.
	CertificateIdentifier *string `field:"optional" json:"certificateIdentifier" yaml:"certificateIdentifier"`
	// The contents of a `.pem` file, which contains an X.509 certificate.
	CertificatePem *string `field:"optional" json:"certificatePem" yaml:"certificatePem"`
	// The location of an imported Oracle Wallet certificate for use with SSL.
	//
	// An example is: `filebase64("${path.root}/rds-ca-2019-root.sso")`
	CertificateWallet *string `field:"optional" json:"certificateWallet" yaml:"certificateWallet"`
}

// A CloudFormation `AWS::DMS::Endpoint`.
//
// The `AWS::DMS::Endpoint` resource specifies an AWS DMS endpoint.
//
// Currently, AWS CloudFormation supports all AWS DMS endpoint types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEndpoint := awscdk.Aws_dms.NewCfnEndpoint(this, jsii.String("MyCfnEndpoint"), &cfnEndpointProps{
//   	endpointType: jsii.String("endpointType"),
//   	engineName: jsii.String("engineName"),
//
//   	// the properties below are optional
//   	certificateArn: jsii.String("certificateArn"),
//   	databaseName: jsii.String("databaseName"),
//   	docDbSettings: &docDbSettingsProperty{
//   		docsToInvestigate: jsii.Number(123),
//   		extractDocId: jsii.Boolean(false),
//   		nestingLevel: jsii.String("nestingLevel"),
//   		secretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   		secretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   	},
//   	dynamoDbSettings: &dynamoDbSettingsProperty{
//   		serviceAccessRoleArn: jsii.String("serviceAccessRoleArn"),
//   	},
//   	elasticsearchSettings: &elasticsearchSettingsProperty{
//   		endpointUri: jsii.String("endpointUri"),
//   		errorRetryDuration: jsii.Number(123),
//   		fullLoadErrorPercentage: jsii.Number(123),
//   		serviceAccessRoleArn: jsii.String("serviceAccessRoleArn"),
//   	},
//   	endpointIdentifier: jsii.String("endpointIdentifier"),
//   	extraConnectionAttributes: jsii.String("extraConnectionAttributes"),
//   	gcpMySqlSettings: &gcpMySQLSettingsProperty{
//   		afterConnectScript: jsii.String("afterConnectScript"),
//   		cleanSourceMetadataOnMismatch: jsii.Boolean(false),
//   		databaseName: jsii.String("databaseName"),
//   		eventsPollInterval: jsii.Number(123),
//   		maxFileSize: jsii.Number(123),
//   		parallelLoadThreads: jsii.Number(123),
//   		password: jsii.String("password"),
//   		port: jsii.Number(123),
//   		secretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   		secretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   		serverName: jsii.String("serverName"),
//   		serverTimezone: jsii.String("serverTimezone"),
//   		username: jsii.String("username"),
//   	},
//   	ibmDb2Settings: &ibmDb2SettingsProperty{
//   		currentLsn: jsii.String("currentLsn"),
//   		maxKBytesPerRead: jsii.Number(123),
//   		secretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   		secretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   		setDataCaptureChanges: jsii.Boolean(false),
//   	},
//   	kafkaSettings: &kafkaSettingsProperty{
//   		broker: jsii.String("broker"),
//   		includeControlDetails: jsii.Boolean(false),
//   		includeNullAndEmpty: jsii.Boolean(false),
//   		includePartitionValue: jsii.Boolean(false),
//   		includeTableAlterOperations: jsii.Boolean(false),
//   		includeTransactionDetails: jsii.Boolean(false),
//   		messageFormat: jsii.String("messageFormat"),
//   		messageMaxBytes: jsii.Number(123),
//   		noHexPrefix: jsii.Boolean(false),
//   		partitionIncludeSchemaTable: jsii.Boolean(false),
//   		saslPassword: jsii.String("saslPassword"),
//   		saslUserName: jsii.String("saslUserName"),
//   		securityProtocol: jsii.String("securityProtocol"),
//   		sslCaCertificateArn: jsii.String("sslCaCertificateArn"),
//   		sslClientCertificateArn: jsii.String("sslClientCertificateArn"),
//   		sslClientKeyArn: jsii.String("sslClientKeyArn"),
//   		sslClientKeyPassword: jsii.String("sslClientKeyPassword"),
//   		topic: jsii.String("topic"),
//   	},
//   	kinesisSettings: &kinesisSettingsProperty{
//   		includeControlDetails: jsii.Boolean(false),
//   		includeNullAndEmpty: jsii.Boolean(false),
//   		includePartitionValue: jsii.Boolean(false),
//   		includeTableAlterOperations: jsii.Boolean(false),
//   		includeTransactionDetails: jsii.Boolean(false),
//   		messageFormat: jsii.String("messageFormat"),
//   		noHexPrefix: jsii.Boolean(false),
//   		partitionIncludeSchemaTable: jsii.Boolean(false),
//   		serviceAccessRoleArn: jsii.String("serviceAccessRoleArn"),
//   		streamArn: jsii.String("streamArn"),
//   	},
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	microsoftSqlServerSettings: &microsoftSqlServerSettingsProperty{
//   		bcpPacketSize: jsii.Number(123),
//   		controlTablesFileGroup: jsii.String("controlTablesFileGroup"),
//   		querySingleAlwaysOnNode: jsii.Boolean(false),
//   		readBackupOnly: jsii.Boolean(false),
//   		safeguardPolicy: jsii.String("safeguardPolicy"),
//   		secretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   		secretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   		useBcpFullLoad: jsii.Boolean(false),
//   		useThirdPartyBackupDevice: jsii.Boolean(false),
//   	},
//   	mongoDbSettings: &mongoDbSettingsProperty{
//   		authMechanism: jsii.String("authMechanism"),
//   		authSource: jsii.String("authSource"),
//   		authType: jsii.String("authType"),
//   		databaseName: jsii.String("databaseName"),
//   		docsToInvestigate: jsii.String("docsToInvestigate"),
//   		extractDocId: jsii.String("extractDocId"),
//   		nestingLevel: jsii.String("nestingLevel"),
//   		password: jsii.String("password"),
//   		port: jsii.Number(123),
//   		secretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   		secretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   		serverName: jsii.String("serverName"),
//   		username: jsii.String("username"),
//   	},
//   	mySqlSettings: &mySqlSettingsProperty{
//   		afterConnectScript: jsii.String("afterConnectScript"),
//   		cleanSourceMetadataOnMismatch: jsii.Boolean(false),
//   		eventsPollInterval: jsii.Number(123),
//   		maxFileSize: jsii.Number(123),
//   		parallelLoadThreads: jsii.Number(123),
//   		secretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   		secretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   		serverTimezone: jsii.String("serverTimezone"),
//   		targetDbType: jsii.String("targetDbType"),
//   	},
//   	neptuneSettings: &neptuneSettingsProperty{
//   		errorRetryDuration: jsii.Number(123),
//   		iamAuthEnabled: jsii.Boolean(false),
//   		maxFileSize: jsii.Number(123),
//   		maxRetryCount: jsii.Number(123),
//   		s3BucketFolder: jsii.String("s3BucketFolder"),
//   		s3BucketName: jsii.String("s3BucketName"),
//   		serviceAccessRoleArn: jsii.String("serviceAccessRoleArn"),
//   	},
//   	oracleSettings: &oracleSettingsProperty{
//   		accessAlternateDirectly: jsii.Boolean(false),
//   		additionalArchivedLogDestId: jsii.Number(123),
//   		addSupplementalLogging: jsii.Boolean(false),
//   		allowSelectNestedTables: jsii.Boolean(false),
//   		archivedLogDestId: jsii.Number(123),
//   		archivedLogsOnly: jsii.Boolean(false),
//   		asmPassword: jsii.String("asmPassword"),
//   		asmServer: jsii.String("asmServer"),
//   		asmUser: jsii.String("asmUser"),
//   		charLengthSemantics: jsii.String("charLengthSemantics"),
//   		directPathNoLog: jsii.Boolean(false),
//   		directPathParallelLoad: jsii.Boolean(false),
//   		enableHomogenousTablespace: jsii.Boolean(false),
//   		extraArchivedLogDestIds: []interface{}{
//   			jsii.Number(123),
//   		},
//   		failTasksOnLobTruncation: jsii.Boolean(false),
//   		numberDatatypeScale: jsii.Number(123),
//   		oraclePathPrefix: jsii.String("oraclePathPrefix"),
//   		parallelAsmReadThreads: jsii.Number(123),
//   		readAheadBlocks: jsii.Number(123),
//   		readTableSpaceName: jsii.Boolean(false),
//   		replacePathPrefix: jsii.Boolean(false),
//   		retryInterval: jsii.Number(123),
//   		secretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   		secretsManagerOracleAsmAccessRoleArn: jsii.String("secretsManagerOracleAsmAccessRoleArn"),
//   		secretsManagerOracleAsmSecretId: jsii.String("secretsManagerOracleAsmSecretId"),
//   		secretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   		securityDbEncryption: jsii.String("securityDbEncryption"),
//   		securityDbEncryptionName: jsii.String("securityDbEncryptionName"),
//   		spatialDataOptionToGeoJsonFunctionName: jsii.String("spatialDataOptionToGeoJsonFunctionName"),
//   		standbyDelayTime: jsii.Number(123),
//   		useAlternateFolderForOnline: jsii.Boolean(false),
//   		useBFile: jsii.Boolean(false),
//   		useDirectPathFullLoad: jsii.Boolean(false),
//   		useLogminerReader: jsii.Boolean(false),
//   		usePathPrefix: jsii.String("usePathPrefix"),
//   	},
//   	password: jsii.String("password"),
//   	port: jsii.Number(123),
//   	postgreSqlSettings: &postgreSqlSettingsProperty{
//   		afterConnectScript: jsii.String("afterConnectScript"),
//   		captureDdls: jsii.Boolean(false),
//   		ddlArtifactsSchema: jsii.String("ddlArtifactsSchema"),
//   		executeTimeout: jsii.Number(123),
//   		failTasksOnLobTruncation: jsii.Boolean(false),
//   		heartbeatEnable: jsii.Boolean(false),
//   		heartbeatFrequency: jsii.Number(123),
//   		heartbeatSchema: jsii.String("heartbeatSchema"),
//   		maxFileSize: jsii.Number(123),
//   		pluginName: jsii.String("pluginName"),
//   		secretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   		secretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   		slotName: jsii.String("slotName"),
//   	},
//   	redisSettings: &redisSettingsProperty{
//   		authPassword: jsii.String("authPassword"),
//   		authType: jsii.String("authType"),
//   		authUserName: jsii.String("authUserName"),
//   		port: jsii.Number(123),
//   		serverName: jsii.String("serverName"),
//   		sslCaCertificateArn: jsii.String("sslCaCertificateArn"),
//   		sslSecurityProtocol: jsii.String("sslSecurityProtocol"),
//   	},
//   	redshiftSettings: &redshiftSettingsProperty{
//   		acceptAnyDate: jsii.Boolean(false),
//   		afterConnectScript: jsii.String("afterConnectScript"),
//   		bucketFolder: jsii.String("bucketFolder"),
//   		bucketName: jsii.String("bucketName"),
//   		caseSensitiveNames: jsii.Boolean(false),
//   		compUpdate: jsii.Boolean(false),
//   		connectionTimeout: jsii.Number(123),
//   		dateFormat: jsii.String("dateFormat"),
//   		emptyAsNull: jsii.Boolean(false),
//   		encryptionMode: jsii.String("encryptionMode"),
//   		explicitIds: jsii.Boolean(false),
//   		fileTransferUploadStreams: jsii.Number(123),
//   		loadTimeout: jsii.Number(123),
//   		maxFileSize: jsii.Number(123),
//   		removeQuotes: jsii.Boolean(false),
//   		replaceChars: jsii.String("replaceChars"),
//   		replaceInvalidChars: jsii.String("replaceInvalidChars"),
//   		secretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   		secretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   		serverSideEncryptionKmsKeyId: jsii.String("serverSideEncryptionKmsKeyId"),
//   		serviceAccessRoleArn: jsii.String("serviceAccessRoleArn"),
//   		timeFormat: jsii.String("timeFormat"),
//   		trimBlanks: jsii.Boolean(false),
//   		truncateColumns: jsii.Boolean(false),
//   		writeBufferSize: jsii.Number(123),
//   	},
//   	resourceIdentifier: jsii.String("resourceIdentifier"),
//   	s3Settings: &s3SettingsProperty{
//   		addColumnName: jsii.Boolean(false),
//   		bucketFolder: jsii.String("bucketFolder"),
//   		bucketName: jsii.String("bucketName"),
//   		cannedAclForObjects: jsii.String("cannedAclForObjects"),
//   		cdcInsertsAndUpdates: jsii.Boolean(false),
//   		cdcInsertsOnly: jsii.Boolean(false),
//   		cdcMaxBatchInterval: jsii.Number(123),
//   		cdcMinFileSize: jsii.Number(123),
//   		cdcPath: jsii.String("cdcPath"),
//   		compressionType: jsii.String("compressionType"),
//   		csvDelimiter: jsii.String("csvDelimiter"),
//   		csvNoSupValue: jsii.String("csvNoSupValue"),
//   		csvNullValue: jsii.String("csvNullValue"),
//   		csvRowDelimiter: jsii.String("csvRowDelimiter"),
//   		dataFormat: jsii.String("dataFormat"),
//   		dataPageSize: jsii.Number(123),
//   		datePartitionDelimiter: jsii.String("datePartitionDelimiter"),
//   		datePartitionEnabled: jsii.Boolean(false),
//   		datePartitionSequence: jsii.String("datePartitionSequence"),
//   		datePartitionTimezone: jsii.String("datePartitionTimezone"),
//   		dictPageSizeLimit: jsii.Number(123),
//   		enableStatistics: jsii.Boolean(false),
//   		encodingType: jsii.String("encodingType"),
//   		encryptionMode: jsii.String("encryptionMode"),
//   		externalTableDefinition: jsii.String("externalTableDefinition"),
//   		ignoreHeaderRows: jsii.Number(123),
//   		includeOpForFullLoad: jsii.Boolean(false),
//   		maxFileSize: jsii.Number(123),
//   		parquetTimestampInMillisecond: jsii.Boolean(false),
//   		parquetVersion: jsii.String("parquetVersion"),
//   		preserveTransactions: jsii.Boolean(false),
//   		rfc4180: jsii.Boolean(false),
//   		rowGroupLength: jsii.Number(123),
//   		serverSideEncryptionKmsKeyId: jsii.String("serverSideEncryptionKmsKeyId"),
//   		serviceAccessRoleArn: jsii.String("serviceAccessRoleArn"),
//   		timestampColumnName: jsii.String("timestampColumnName"),
//   		useCsvNoSupValue: jsii.Boolean(false),
//   		useTaskStartTimeForFullLoadTimestamp: jsii.Boolean(false),
//   	},
//   	serverName: jsii.String("serverName"),
//   	sslMode: jsii.String("sslMode"),
//   	sybaseSettings: &sybaseSettingsProperty{
//   		secretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   		secretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	username: jsii.String("username"),
//   })
//
type CfnEndpoint interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// A value that can be used for cross-account validation.
	AttrExternalId() *string
	// The Amazon Resource Name (ARN) for the certificate.
	CertificateArn() *string
	SetCertificateArn(val *string)
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
	// The name of the endpoint database.
	//
	// For a MySQL source or target endpoint, don't specify `DatabaseName` . To migrate to a specific database, use this setting and `targetDbType` .
	DatabaseName() *string
	SetDatabaseName(val *string)
	// Settings in JSON format for the source and target DocumentDB endpoint.
	//
	// For more information about other available settings, see [Using extra connections attributes with Amazon DocumentDB as a source](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.DocumentDB.html#CHAP_Source.DocumentDB.ECAs) and [Using Amazon DocumentDB as a target for AWS Database Migration Service](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.DocumentDB.html) in the *AWS Database Migration Service User Guide* .
	DocDbSettings() interface{}
	SetDocDbSettings(val interface{})
	// Settings in JSON format for the target Amazon DynamoDB endpoint.
	//
	// For information about other available settings, see [Using object mapping to migrate data to DynamoDB](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.DynamoDB.html#CHAP_Target.DynamoDB.ObjectMapping) in the *AWS Database Migration Service User Guide* .
	DynamoDbSettings() interface{}
	SetDynamoDbSettings(val interface{})
	// Settings in JSON format for the target OpenSearch endpoint.
	//
	// For more information about the available settings, see [Extra connection attributes when using OpenSearch as a target for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Elasticsearch.html#CHAP_Target.Elasticsearch.Configuration) in the *AWS Database Migration Service User Guide* .
	ElasticsearchSettings() interface{}
	SetElasticsearchSettings(val interface{})
	// The database endpoint identifier.
	//
	// Identifiers must begin with a letter and must contain only ASCII letters, digits, and hyphens. They can't end with a hyphen, or contain two consecutive hyphens.
	EndpointIdentifier() *string
	SetEndpointIdentifier(val *string)
	// The type of endpoint.
	//
	// Valid values are `source` and `target` .
	EndpointType() *string
	SetEndpointType(val *string)
	// The type of engine for the endpoint, depending on the `EndpointType` value.
	//
	// *Valid values* : `mysql` | `oracle` | `postgres` | `mariadb` | `aurora` | `aurora-postgresql` | `opensearch` | `redshift` | `s3` | `db2` | `azuredb` | `sybase` | `dynamodb` | `mongodb` | `kinesis` | `kafka` | `elasticsearch` | `docdb` | `sqlserver` | `neptune`.
	EngineName() *string
	SetEngineName(val *string)
	// Additional attributes associated with the connection.
	//
	// Each attribute is specified as a name-value pair associated by an equal sign (=). Multiple attributes are separated by a semicolon (;) with no additional white space. For information on the attributes available for connecting your source or target endpoint, see [Working with AWS DMS Endpoints](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Endpoints.html) in the *AWS Database Migration Service User Guide* .
	ExtraConnectionAttributes() *string
	SetExtraConnectionAttributes(val *string)
	// Settings in JSON format for the source GCP MySQL endpoint.
	//
	// These settings are much the same as the settings for any MySQL-compatible endpoint. For more information, see [Extra connection attributes when using MySQL as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.MySQL.html#CHAP_Source.MySQL.ConnectionAttrib) in the *AWS Database Migration Service User Guide* .
	GcpMySqlSettings() interface{}
	SetGcpMySqlSettings(val interface{})
	// Settings in JSON format for the source IBM Db2 LUW endpoint.
	//
	// For information about other available settings, see [Extra connection attributes when using Db2 LUW as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.DB2.html#CHAP_Source.DB2.ConnectionAttrib) in the *AWS Database Migration Service User Guide* .
	IbmDb2Settings() interface{}
	SetIbmDb2Settings(val interface{})
	// Settings in JSON format for the target Apache Kafka endpoint.
	//
	// For more information about other available settings, see [Using object mapping to migrate data to a Kafka topic](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Kafka.html#CHAP_Target.Kafka.ObjectMapping) in the *AWS Database Migration Service User Guide* .
	KafkaSettings() interface{}
	SetKafkaSettings(val interface{})
	// Settings in JSON format for the target endpoint for Amazon Kinesis Data Streams.
	//
	// For more information about other available settings, see [Using object mapping to migrate data to a Kinesis data stream](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Kinesis.html#CHAP_Target.Kinesis.ObjectMapping) in the *AWS Database Migration Service User Guide* .
	KinesisSettings() interface{}
	SetKinesisSettings(val interface{})
	// An AWS KMS key identifier that is used to encrypt the connection parameters for the endpoint.
	//
	// If you don't specify a value for the `KmsKeyId` parameter, AWS DMS uses your default encryption key.
	//
	// AWS KMS creates the default encryption key for your AWS account . Your AWS account has a different default encryption key for each AWS Region .
	KmsKeyId() *string
	SetKmsKeyId(val *string)
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
	// Settings in JSON format for the source and target Microsoft SQL Server endpoint.
	//
	// For information about other available settings, see [Extra connection attributes when using SQL Server as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.SQLServer.html#CHAP_Source.SQLServer.ConnectionAttrib) and [Extra connection attributes when using SQL Server as a target for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.SQLServer.html#CHAP_Target.SQLServer.ConnectionAttrib) in the *AWS Database Migration Service User Guide* .
	MicrosoftSqlServerSettings() interface{}
	SetMicrosoftSqlServerSettings(val interface{})
	// Settings in JSON format for the source MongoDB endpoint.
	//
	// For more information about the available settings, see [Using MongoDB as a target for AWS Database Migration Service](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.MongoDB.html#CHAP_Source.MongoDB.Configuration) in the *AWS Database Migration Service User Guide* .
	MongoDbSettings() interface{}
	SetMongoDbSettings(val interface{})
	// Settings in JSON format for the source and target MySQL endpoint.
	//
	// For information about other available settings, see [Extra connection attributes when using MySQL as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.MySQL.html#CHAP_Source.MySQL.ConnectionAttrib) and [Extra connection attributes when using a MySQL-compatible database as a target for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.MySQL.html#CHAP_Target.MySQL.ConnectionAttrib) in the *AWS Database Migration Service User Guide* .
	MySqlSettings() interface{}
	SetMySqlSettings(val interface{})
	// Settings in JSON format for the target Amazon Neptune endpoint.
	//
	// For more information about the available settings, see [Specifying endpoint settings for Amazon Neptune as a target](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Neptune.html#CHAP_Target.Neptune.EndpointSettings) in the *AWS Database Migration Service User Guide* .
	NeptuneSettings() interface{}
	SetNeptuneSettings(val interface{})
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Settings in JSON format for the source and target Oracle endpoint.
	//
	// For information about other available settings, see [Extra connection attributes when using Oracle as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.Oracle.html#CHAP_Source.Oracle.ConnectionAttrib) and [Extra connection attributes when using Oracle as a target for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Oracle.html#CHAP_Target.Oracle.ConnectionAttrib) in the *AWS Database Migration Service User Guide* .
	OracleSettings() interface{}
	SetOracleSettings(val interface{})
	// The password to be used to log in to the endpoint database.
	Password() *string
	SetPassword(val *string)
	// The port used by the endpoint database.
	Port() *float64
	SetPort(val *float64)
	// Settings in JSON format for the source and target PostgreSQL endpoint.
	//
	// For information about other available settings, see [Extra connection attributes when using PostgreSQL as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.PostgreSQL.html#CHAP_Source.PostgreSQL.ConnectionAttrib) and [Extra connection attributes when using PostgreSQL as a target for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.PostgreSQL.html#CHAP_Target.PostgreSQL.ConnectionAttrib) in the *AWS Database Migration Service User Guide* .
	PostgreSqlSettings() interface{}
	SetPostgreSqlSettings(val interface{})
	// Settings in JSON format for the target Redis endpoint.
	//
	// For information about other available settings, see [Specifying endpoint settings for Redis as a target](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Redis.html#CHAP_Target.Redis.EndpointSettings) in the *AWS Database Migration Service User Guide* .
	RedisSettings() interface{}
	SetRedisSettings(val interface{})
	// Settings in JSON format for the Amazon Redshift endpoint.
	//
	// For more information about other available settings, see [Extra connection attributes when using Amazon Redshift as a target for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Redshift.html#CHAP_Target.Redshift.ConnectionAttrib) in the *AWS Database Migration Service User Guide* .
	RedshiftSettings() interface{}
	SetRedshiftSettings(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// A display name for the resource identifier at the end of the `EndpointArn` response parameter that is returned in the created `Endpoint` object.
	//
	// The value for this parameter can have up to 31 characters. It can contain only ASCII letters, digits, and hyphen ('-'). Also, it can't end with a hyphen or contain two consecutive hyphens, and can only begin with a letter, such as `Example-App-ARN1` .
	//
	// For example, this value might result in the `EndpointArn` value `arn:aws:dms:eu-west-1:012345678901:rep:Example-App-ARN1` . If you don't specify a `ResourceIdentifier` value, AWS DMS generates a default identifier value for the end of `EndpointArn` .
	ResourceIdentifier() *string
	SetResourceIdentifier(val *string)
	// Settings in JSON format for the source and target Amazon S3 endpoint.
	//
	// For more information about other available settings, see [Extra connection attributes when using Amazon S3 as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.S3.html#CHAP_Source.S3.Configuring) and [Extra connection attributes when using Amazon S3 as a target for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.S3.html#CHAP_Target.S3.Configuring) in the *AWS Database Migration Service User Guide* .
	S3Settings() interface{}
	SetS3Settings(val interface{})
	// The name of the server where the endpoint database resides.
	ServerName() *string
	SetServerName(val *string)
	// The Secure Sockets Layer (SSL) mode to use for the SSL connection. The default is `none` .
	//
	// > When `engine_name` is set to S3, the only allowed value is `none` .
	SslMode() *string
	SetSslMode(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// Settings in JSON format for the source and target SAP ASE endpoint.
	//
	// For information about other available settings, see [Extra connection attributes when using SAP ASE as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.SAP.html#CHAP_Source.SAP.ConnectionAttrib) and [Extra connection attributes when using SAP ASE as a target for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.SAP.html#CHAP_Target.SAP.ConnectionAttrib) in the *AWS Database Migration Service User Guide* .
	SybaseSettings() interface{}
	SetSybaseSettings(val interface{})
	// One or more tags to be assigned to the endpoint.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// The user name to be used to log in to the endpoint database.
	Username() *string
	SetUsername(val *string)
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

func (c *jsiiProxy_CfnEndpoint) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnEndpoint) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnEndpoint) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnEndpoint) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnEndpoint) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnEndpoint) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnEndpoint) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

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

func (c *jsiiProxy_CfnEndpoint) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnEndpoint) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnEndpoint) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_CfnEndpoint) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

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

func (c *jsiiProxy_CfnEndpoint) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_CfnEndpoint) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Provides information that defines a DocumentDB endpoint.
//
// This information includes the output format of records applied to the endpoint and details of transaction and control table data information. For more information about other available settings, see [Using extra connections attributes with Amazon DocumentDB as a source](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.DocumentDB.html#CHAP_Source.DocumentDB.ECAs) and [Using Amazon DocumentDB as a target for AWS Database Migration Service](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.DocumentDB.html) in the *AWS Database Migration Service User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   docDbSettingsProperty := &docDbSettingsProperty{
//   	docsToInvestigate: jsii.Number(123),
//   	extractDocId: jsii.Boolean(false),
//   	nestingLevel: jsii.String("nestingLevel"),
//   	secretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   	secretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   }
//
type CfnEndpoint_DocDbSettingsProperty struct {
	// Indicates the number of documents to preview to determine the document organization.
	//
	// Use this setting when `NestingLevel` is set to `"one"` .
	//
	// Must be a positive value greater than `0` . Default value is `1000` .
	DocsToInvestigate *float64 `field:"optional" json:"docsToInvestigate" yaml:"docsToInvestigate"`
	// Specifies the document ID. Use this setting when `NestingLevel` is set to `"none"` .
	//
	// Default value is `"false"` .
	ExtractDocId interface{} `field:"optional" json:"extractDocId" yaml:"extractDocId"`
	// Specifies either document or table mode.
	//
	// Default value is `"none"` . Specify `"none"` to use document mode. Specify `"one"` to use table mode.
	NestingLevel *string `field:"optional" json:"nestingLevel" yaml:"nestingLevel"`
	// The full Amazon Resource Name (ARN) of the IAM role that specifies AWS DMS as the trusted entity and grants the required permissions to access the value in `SecretsManagerSecret` .
	//
	// The role must allow the `iam:PassRole` action. `SecretsManagerSecret` has the value of the AWS Secrets Manager secret that allows access to the DocumentDB endpoint.
	//
	// > You can specify one of two sets of values for these permissions. You can specify the values for this setting and `SecretsManagerSecretId` . Or you can specify clear-text values for `UserName` , `Password` , `ServerName` , and `Port` . You can't specify both.
	// >
	// > For more information on creating this `SecretsManagerSecret` , the corresponding `SecretsManagerAccessRoleArn` , and the `SecretsManagerSecretId` that is required to access it, see [Using secrets to access AWS Database Migration Service resources](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Security.html#security-iam-secretsmanager) in the *AWS Database Migration Service User Guide* .
	SecretsManagerAccessRoleArn *string `field:"optional" json:"secretsManagerAccessRoleArn" yaml:"secretsManagerAccessRoleArn"`
	// The full ARN, partial ARN, or display name of the `SecretsManagerSecret` that contains the DocumentDB endpoint connection details.
	SecretsManagerSecretId *string `field:"optional" json:"secretsManagerSecretId" yaml:"secretsManagerSecretId"`
}

// Provides information, including the Amazon Resource Name (ARN) of the IAM role used to define an Amazon DynamoDB target endpoint.
//
// This information also includes the output format of records applied to the endpoint and details of transaction and control table data information. For information about other available settings, see [Using object mapping to migrate data to DynamoDB](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.DynamoDB.html#CHAP_Target.DynamoDB.ObjectMapping) in the *AWS Database Migration Service User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dynamoDbSettingsProperty := &dynamoDbSettingsProperty{
//   	serviceAccessRoleArn: jsii.String("serviceAccessRoleArn"),
//   }
//
type CfnEndpoint_DynamoDbSettingsProperty struct {
	// The Amazon Resource Name (ARN) used by the service to access the IAM role.
	//
	// The role must allow the `iam:PassRole` action.
	ServiceAccessRoleArn *string `field:"optional" json:"serviceAccessRoleArn" yaml:"serviceAccessRoleArn"`
}

// Provides information that defines an OpenSearch endpoint.
//
// This information includes the output format of records applied to the endpoint and details of transaction and control table data information. For more information about the available settings, see [Extra connection attributes when using OpenSearch as a target for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Elasticsearch.html#CHAP_Target.Elasticsearch.Configuration) in the *AWS Database Migration Service User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   elasticsearchSettingsProperty := &elasticsearchSettingsProperty{
//   	endpointUri: jsii.String("endpointUri"),
//   	errorRetryDuration: jsii.Number(123),
//   	fullLoadErrorPercentage: jsii.Number(123),
//   	serviceAccessRoleArn: jsii.String("serviceAccessRoleArn"),
//   }
//
type CfnEndpoint_ElasticsearchSettingsProperty struct {
	// The endpoint for the OpenSearch cluster.
	//
	// AWS DMS uses HTTPS if a transport protocol (either HTTP or HTTPS) isn't specified.
	EndpointUri *string `field:"optional" json:"endpointUri" yaml:"endpointUri"`
	// The maximum number of seconds for which DMS retries failed API requests to the OpenSearch cluster.
	ErrorRetryDuration *float64 `field:"optional" json:"errorRetryDuration" yaml:"errorRetryDuration"`
	// The maximum percentage of records that can fail to be written before a full load operation stops.
	//
	// To avoid early failure, this counter is only effective after 1,000 records are transferred. OpenSearch also has the concept of error monitoring during the last 10 minutes of an Observation Window. If transfer of all records fail in the last 10 minutes, the full load operation stops.
	FullLoadErrorPercentage *float64 `field:"optional" json:"fullLoadErrorPercentage" yaml:"fullLoadErrorPercentage"`
	// The Amazon Resource Name (ARN) used by the service to access the IAM role.
	//
	// The role must allow the `iam:PassRole` action.
	ServiceAccessRoleArn *string `field:"optional" json:"serviceAccessRoleArn" yaml:"serviceAccessRoleArn"`
}

// Provides information that defines a GCP MySQL endpoint.
//
// This information includes the output format of records applied to the endpoint and details of transaction and control table data information. These settings are much the same as the settings for any MySQL-compatible endpoint. For more information, see [Extra connection attributes when using MySQL as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.MySQL.html#CHAP_Source.MySQL.ConnectionAttrib) in the *AWS Database Migration Service User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gcpMySQLSettingsProperty := &gcpMySQLSettingsProperty{
//   	afterConnectScript: jsii.String("afterConnectScript"),
//   	cleanSourceMetadataOnMismatch: jsii.Boolean(false),
//   	databaseName: jsii.String("databaseName"),
//   	eventsPollInterval: jsii.Number(123),
//   	maxFileSize: jsii.Number(123),
//   	parallelLoadThreads: jsii.Number(123),
//   	password: jsii.String("password"),
//   	port: jsii.Number(123),
//   	secretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   	secretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   	serverName: jsii.String("serverName"),
//   	serverTimezone: jsii.String("serverTimezone"),
//   	username: jsii.String("username"),
//   }
//
type CfnEndpoint_GcpMySQLSettingsProperty struct {
	// Specifies a script to run immediately after AWS DMS connects to the endpoint.
	//
	// The migration task continues running regardless if the SQL statement succeeds or fails.
	//
	// For this parameter, provide the code of the script itself, not the name of a file containing the script.
	AfterConnectScript *string `field:"optional" json:"afterConnectScript" yaml:"afterConnectScript"`
	// Adjusts the behavior of AWS DMS when migrating from an SQL Server source database that is hosted as part of an Always On availability group cluster.
	//
	// If you need AWS DMS to poll all the nodes in the Always On cluster for transaction backups, set this attribute to `false` .
	CleanSourceMetadataOnMismatch interface{} `field:"optional" json:"cleanSourceMetadataOnMismatch" yaml:"cleanSourceMetadataOnMismatch"`
	// Database name for the endpoint.
	//
	// For a MySQL source or target endpoint, don't explicitly specify the database using the `DatabaseName` request parameter on either the `CreateEndpoint` or `ModifyEndpoint` API call. Specifying `DatabaseName` when you create or modify a MySQL endpoint replicates all the task tables to this single database. For MySQL endpoints, you specify the database only when you specify the schema in the table-mapping rules of the AWS DMS task.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// Specifies how often to check the binary log for new changes/events when the database is idle.
	//
	// The default is five seconds.
	//
	// Example: `eventsPollInterval=5;`
	//
	// In the example, AWS DMS checks for changes in the binary logs every five seconds.
	EventsPollInterval *float64 `field:"optional" json:"eventsPollInterval" yaml:"eventsPollInterval"`
	// Specifies the maximum size (in KB) of any .csv file used to transfer data to a MySQL-compatible database.
	//
	// Example: `maxFileSize=512`.
	MaxFileSize *float64 `field:"optional" json:"maxFileSize" yaml:"maxFileSize"`
	// Improves performance when loading data into the MySQL-compatible target database.
	//
	// Specifies how many threads to use to load the data into the MySQL-compatible target database. Setting a large number of threads can have an adverse effect on database performance, because a separate connection is required for each thread. The default is one.
	//
	// Example: `parallelLoadThreads=1`.
	ParallelLoadThreads *float64 `field:"optional" json:"parallelLoadThreads" yaml:"parallelLoadThreads"`
	// Endpoint connection password.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The port used by the endpoint database.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The full Amazon Resource Name (ARN) of the IAM role that specifies AWS DMS as the trusted entity and grants the required permissions to access the value in `SecretsManagerSecret.` The role must allow the `iam:PassRole` action. `SecretsManagerSecret` has the value of the AWS Secrets Manager secret that allows access to the MySQL endpoint.
	//
	// > You can specify one of two sets of values for these permissions. You can specify the values for this setting and `SecretsManagerSecretId` . Or you can specify clear-text values for `UserName` , `Password` , `ServerName` , and `Port` . You can't specify both.
	// >
	// > For more information on creating this `SecretsManagerSecret` , the corresponding `SecretsManagerAccessRoleArn` , and the `SecretsManagerSecretId` required to access it, see [Using secrets to access AWS Database Migration Service resources](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Security.html#security-iam-secretsmanager) in the *AWS Database Migration Service User Guide* .
	SecretsManagerAccessRoleArn *string `field:"optional" json:"secretsManagerAccessRoleArn" yaml:"secretsManagerAccessRoleArn"`
	// The full ARN, partial ARN, or display name of the `SecretsManagerSecret` that contains the MySQL endpoint connection details.
	SecretsManagerSecretId *string `field:"optional" json:"secretsManagerSecretId" yaml:"secretsManagerSecretId"`
	// Endpoint TCP port.
	ServerName *string `field:"optional" json:"serverName" yaml:"serverName"`
	// Specifies the time zone for the source MySQL database. Don't enclose time zones in single quotation marks.
	//
	// Example: `serverTimezone=US/Pacific;`.
	ServerTimezone *string `field:"optional" json:"serverTimezone" yaml:"serverTimezone"`
	// Endpoint connection user name.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

// Provides information that defines an IBMDB2 endpoint.
//
// This information includes the output format of records applied to the endpoint and details of transaction and control table data information. For more information about other available settings, see [Extra connection attributes when using Db2 LUW as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.DB2.html#CHAP_Source.DB2.ConnectionAttrib) in the *AWS Database Migration Service User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ibmDb2SettingsProperty := &ibmDb2SettingsProperty{
//   	currentLsn: jsii.String("currentLsn"),
//   	maxKBytesPerRead: jsii.Number(123),
//   	secretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   	secretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   	setDataCaptureChanges: jsii.Boolean(false),
//   }
//
type CfnEndpoint_IbmDb2SettingsProperty struct {
	// For ongoing replication (CDC), use CurrentLSN to specify a log sequence number (LSN) where you want the replication to start.
	CurrentLsn *string `field:"optional" json:"currentLsn" yaml:"currentLsn"`
	// Maximum number of bytes per read, as a NUMBER value.
	//
	// The default is 64 KB.
	MaxKBytesPerRead *float64 `field:"optional" json:"maxKBytesPerRead" yaml:"maxKBytesPerRead"`
	// The full Amazon Resource Name (ARN) of the IAM role that specifies AWS DMS as the trusted entity and grants the required permissions to access the value in `SecretsManagerSecret` .
	//
	// The role must allow the `iam:PassRole` action. `SecretsManagerSecret` has the value ofthe AWS Secrets Manager secret that allows access to the Db2 LUW endpoint.
	//
	// > You can specify one of two sets of values for these permissions. You can specify the values for this setting and `SecretsManagerSecretId` . Or you can specify clear-text values for `UserName` , `Password` , `ServerName` , and `Port` . You can't specify both.
	// >
	// > For more information on creating this `SecretsManagerSecret` , the corresponding `SecretsManagerAccessRoleArn` , and the `SecretsManagerSecretId` that is required to access it, see [Using secrets to access AWS Database Migration Service resources](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Security.html#security-iam-secretsmanager) in the *AWS Database Migration Service User Guide* .
	SecretsManagerAccessRoleArn *string `field:"optional" json:"secretsManagerAccessRoleArn" yaml:"secretsManagerAccessRoleArn"`
	// The full ARN, partial ARN, or display name of the `SecretsManagerSecret` that contains the IBMDB2 endpoint connection details.
	SecretsManagerSecretId *string `field:"optional" json:"secretsManagerSecretId" yaml:"secretsManagerSecretId"`
	// Enables ongoing replication (CDC) as a BOOLEAN value.
	//
	// The default is true.
	SetDataCaptureChanges interface{} `field:"optional" json:"setDataCaptureChanges" yaml:"setDataCaptureChanges"`
}

// Provides information that describes an Apache Kafka endpoint.
//
// This information includes the output format of records applied to the endpoint and details of transaction and control table data information. For more information about other available settings, see [Using object mapping to migrate data to a Kafka topic](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Kafka.html#CHAP_Target.Kafka.ObjectMapping) in the *AWS Database Migration Service User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kafkaSettingsProperty := &kafkaSettingsProperty{
//   	broker: jsii.String("broker"),
//   	includeControlDetails: jsii.Boolean(false),
//   	includeNullAndEmpty: jsii.Boolean(false),
//   	includePartitionValue: jsii.Boolean(false),
//   	includeTableAlterOperations: jsii.Boolean(false),
//   	includeTransactionDetails: jsii.Boolean(false),
//   	messageFormat: jsii.String("messageFormat"),
//   	messageMaxBytes: jsii.Number(123),
//   	noHexPrefix: jsii.Boolean(false),
//   	partitionIncludeSchemaTable: jsii.Boolean(false),
//   	saslPassword: jsii.String("saslPassword"),
//   	saslUserName: jsii.String("saslUserName"),
//   	securityProtocol: jsii.String("securityProtocol"),
//   	sslCaCertificateArn: jsii.String("sslCaCertificateArn"),
//   	sslClientCertificateArn: jsii.String("sslClientCertificateArn"),
//   	sslClientKeyArn: jsii.String("sslClientKeyArn"),
//   	sslClientKeyPassword: jsii.String("sslClientKeyPassword"),
//   	topic: jsii.String("topic"),
//   }
//
type CfnEndpoint_KafkaSettingsProperty struct {
	// A comma-separated list of one or more broker locations in your Kafka cluster that host your Kafka instance.
	//
	// Specify each broker location in the form `*broker-hostname-or-ip* : *port*` . For example, `"ec2-12-345-678-901.compute-1.amazonaws.com:2345"` . For more information and examples of specifying a list of broker locations, see [Using Apache Kafka as a target for AWS Database Migration Service](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Kafka.html) in the *AWS Database Migration Service User Guide* .
	Broker *string `field:"optional" json:"broker" yaml:"broker"`
	// Shows detailed control information for table definition, column definition, and table and column changes in the Kafka message output.
	//
	// The default is `false` .
	IncludeControlDetails interface{} `field:"optional" json:"includeControlDetails" yaml:"includeControlDetails"`
	// Include NULL and empty columns for records migrated to the endpoint.
	//
	// The default is `false` .
	IncludeNullAndEmpty interface{} `field:"optional" json:"includeNullAndEmpty" yaml:"includeNullAndEmpty"`
	// Shows the partition value within the Kafka message output unless the partition type is `schema-table-type` .
	//
	// The default is `false` .
	IncludePartitionValue interface{} `field:"optional" json:"includePartitionValue" yaml:"includePartitionValue"`
	// Includes any data definition language (DDL) operations that change the table in the control data, such as `rename-table` , `drop-table` , `add-column` , `drop-column` , and `rename-column` .
	//
	// The default is `false` .
	IncludeTableAlterOperations interface{} `field:"optional" json:"includeTableAlterOperations" yaml:"includeTableAlterOperations"`
	// Provides detailed transaction information from the source database.
	//
	// This information includes a commit timestamp, a log position, and values for `transaction_id` , previous `transaction_id` , and `transaction_record_id` (the record offset within a transaction). The default is `false` .
	IncludeTransactionDetails interface{} `field:"optional" json:"includeTransactionDetails" yaml:"includeTransactionDetails"`
	// The output format for the records created on the endpoint.
	//
	// The message format is `JSON` (default) or `JSON_UNFORMATTED` (a single line with no tab).
	MessageFormat *string `field:"optional" json:"messageFormat" yaml:"messageFormat"`
	// The maximum size in bytes for records created on the endpoint The default is 1,000,000.
	MessageMaxBytes *float64 `field:"optional" json:"messageMaxBytes" yaml:"messageMaxBytes"`
	// Set this optional parameter to `true` to avoid adding a '0x' prefix to raw data in hexadecimal format.
	//
	// For example, by default, AWS DMS adds a '0x' prefix to the LOB column type in hexadecimal format moving from an Oracle source to a Kafka target. Use the `NoHexPrefix` endpoint setting to enable migration of RAW data type columns without adding the '0x' prefix.
	NoHexPrefix interface{} `field:"optional" json:"noHexPrefix" yaml:"noHexPrefix"`
	// Prefixes schema and table names to partition values, when the partition type is `primary-key-type` .
	//
	// Doing this increases data distribution among Kafka partitions. For example, suppose that a SysBench schema has thousands of tables and each table has only limited range for a primary key. In this case, the same primary key is sent from thousands of tables to the same partition, which causes throttling. The default is `false` .
	PartitionIncludeSchemaTable interface{} `field:"optional" json:"partitionIncludeSchemaTable" yaml:"partitionIncludeSchemaTable"`
	// The secure password that you created when you first set up your Amazon MSK cluster to validate a client identity and make an encrypted connection between server and client using SASL-SSL authentication.
	SaslPassword *string `field:"optional" json:"saslPassword" yaml:"saslPassword"`
	// The secure user name you created when you first set up your Amazon MSK cluster to validate a client identity and make an encrypted connection between server and client using SASL-SSL authentication.
	SaslUserName *string `field:"optional" json:"saslUserName" yaml:"saslUserName"`
	// Set secure connection to a Kafka target endpoint using Transport Layer Security (TLS).
	//
	// Options include `ssl-encryption` , `ssl-authentication` , and `sasl-ssl` . `sasl-ssl` requires `SaslUsername` and `SaslPassword` .
	SecurityProtocol *string `field:"optional" json:"securityProtocol" yaml:"securityProtocol"`
	// The Amazon Resource Name (ARN) for the private certificate authority (CA) cert that AWS DMS uses to securely connect to your Kafka target endpoint.
	SslCaCertificateArn *string `field:"optional" json:"sslCaCertificateArn" yaml:"sslCaCertificateArn"`
	// The Amazon Resource Name (ARN) of the client certificate used to securely connect to a Kafka target endpoint.
	SslClientCertificateArn *string `field:"optional" json:"sslClientCertificateArn" yaml:"sslClientCertificateArn"`
	// The Amazon Resource Name (ARN) for the client private key used to securely connect to a Kafka target endpoint.
	SslClientKeyArn *string `field:"optional" json:"sslClientKeyArn" yaml:"sslClientKeyArn"`
	// The password for the client private key used to securely connect to a Kafka target endpoint.
	SslClientKeyPassword *string `field:"optional" json:"sslClientKeyPassword" yaml:"sslClientKeyPassword"`
	// The topic to which you migrate the data.
	//
	// If you don't specify a topic, AWS DMS specifies `"kafka-default-topic"` as the migration topic.
	Topic *string `field:"optional" json:"topic" yaml:"topic"`
}

// Provides information that describes an Amazon Kinesis Data Stream endpoint.
//
// This information includes the output format of records applied to the endpoint and details of transaction and control table data information. For more information about other available settings, see [Using object mapping to migrate data to a Kinesis data stream](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Kinesis.html#CHAP_Target.Kinesis.ObjectMapping) in the *AWS Database Migration Service User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kinesisSettingsProperty := &kinesisSettingsProperty{
//   	includeControlDetails: jsii.Boolean(false),
//   	includeNullAndEmpty: jsii.Boolean(false),
//   	includePartitionValue: jsii.Boolean(false),
//   	includeTableAlterOperations: jsii.Boolean(false),
//   	includeTransactionDetails: jsii.Boolean(false),
//   	messageFormat: jsii.String("messageFormat"),
//   	noHexPrefix: jsii.Boolean(false),
//   	partitionIncludeSchemaTable: jsii.Boolean(false),
//   	serviceAccessRoleArn: jsii.String("serviceAccessRoleArn"),
//   	streamArn: jsii.String("streamArn"),
//   }
//
type CfnEndpoint_KinesisSettingsProperty struct {
	// Shows detailed control information for table definition, column definition, and table and column changes in the Kinesis message output.
	//
	// The default is `false` .
	IncludeControlDetails interface{} `field:"optional" json:"includeControlDetails" yaml:"includeControlDetails"`
	// Include NULL and empty columns for records migrated to the endpoint.
	//
	// The default is `false` .
	IncludeNullAndEmpty interface{} `field:"optional" json:"includeNullAndEmpty" yaml:"includeNullAndEmpty"`
	// Shows the partition value within the Kinesis message output, unless the partition type is `schema-table-type` .
	//
	// The default is `false` .
	IncludePartitionValue interface{} `field:"optional" json:"includePartitionValue" yaml:"includePartitionValue"`
	// Includes any data definition language (DDL) operations that change the table in the control data, such as `rename-table` , `drop-table` , `add-column` , `drop-column` , and `rename-column` .
	//
	// The default is `false` .
	IncludeTableAlterOperations interface{} `field:"optional" json:"includeTableAlterOperations" yaml:"includeTableAlterOperations"`
	// Provides detailed transaction information from the source database.
	//
	// This information includes a commit timestamp, a log position, and values for `transaction_id` , previous `transaction_id` , and `transaction_record_id` (the record offset within a transaction). The default is `false` .
	IncludeTransactionDetails interface{} `field:"optional" json:"includeTransactionDetails" yaml:"includeTransactionDetails"`
	// The output format for the records created on the endpoint.
	//
	// The message format is `JSON` (default) or `JSON_UNFORMATTED` (a single line with no tab).
	MessageFormat *string `field:"optional" json:"messageFormat" yaml:"messageFormat"`
	// Set this optional parameter to `true` to avoid adding a '0x' prefix to raw data in hexadecimal format.
	//
	// For example, by default, AWS DMS adds a '0x' prefix to the LOB column type in hexadecimal format moving from an Oracle source to an Amazon Kinesis target. Use the `NoHexPrefix` endpoint setting to enable migration of RAW data type columns without adding the '0x' prefix.
	NoHexPrefix interface{} `field:"optional" json:"noHexPrefix" yaml:"noHexPrefix"`
	// Prefixes schema and table names to partition values, when the partition type is `primary-key-type` .
	//
	// Doing this increases data distribution among Kinesis shards. For example, suppose that a SysBench schema has thousands of tables and each table has only limited range for a primary key. In this case, the same primary key is sent from thousands of tables to the same shard, which causes throttling. The default is `false` .
	PartitionIncludeSchemaTable interface{} `field:"optional" json:"partitionIncludeSchemaTable" yaml:"partitionIncludeSchemaTable"`
	// The Amazon Resource Name (ARN) for the IAM role that AWS DMS uses to write to the Kinesis data stream.
	//
	// The role must allow the `iam:PassRole` action.
	ServiceAccessRoleArn *string `field:"optional" json:"serviceAccessRoleArn" yaml:"serviceAccessRoleArn"`
	// The Amazon Resource Name (ARN) for the Amazon Kinesis Data Streams endpoint.
	StreamArn *string `field:"optional" json:"streamArn" yaml:"streamArn"`
}

// Provides information that defines a Microsoft SQL Server endpoint.
//
// This information includes the output format of records applied to the endpoint and details of transaction and control table data information. For information about other available settings, see [Extra connection attributes when using SQL Server as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.SQLServer.html#CHAP_Source.SQLServer.ConnectionAttrib) and [Extra connection attributes when using SQL Server as a target for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.SQLServer.html#CHAP_Target.SQLServer.ConnectionAttrib) in the *AWS Database Migration Service User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   microsoftSqlServerSettingsProperty := &microsoftSqlServerSettingsProperty{
//   	bcpPacketSize: jsii.Number(123),
//   	controlTablesFileGroup: jsii.String("controlTablesFileGroup"),
//   	querySingleAlwaysOnNode: jsii.Boolean(false),
//   	readBackupOnly: jsii.Boolean(false),
//   	safeguardPolicy: jsii.String("safeguardPolicy"),
//   	secretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   	secretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   	useBcpFullLoad: jsii.Boolean(false),
//   	useThirdPartyBackupDevice: jsii.Boolean(false),
//   }
//
type CfnEndpoint_MicrosoftSqlServerSettingsProperty struct {
	// The maximum size of the packets (in bytes) used to transfer data using BCP.
	BcpPacketSize *float64 `field:"optional" json:"bcpPacketSize" yaml:"bcpPacketSize"`
	// Specifies a file group for the AWS DMS internal tables.
	//
	// When the replication task starts, all the internal AWS DMS control tables (awsdms_ apply_exception, awsdms_apply, awsdms_changes) are created for the specified file group.
	ControlTablesFileGroup *string `field:"optional" json:"controlTablesFileGroup" yaml:"controlTablesFileGroup"`
	// Cleans and recreates table metadata information on the replication instance when a mismatch occurs.
	//
	// An example is a situation where running an alter DDL statement on a table might result in different information about the table cached in the replication instance.
	QuerySingleAlwaysOnNode interface{} `field:"optional" json:"querySingleAlwaysOnNode" yaml:"querySingleAlwaysOnNode"`
	// When this attribute is set to `Y` , AWS DMS only reads changes from transaction log backups and doesn't read from the active transaction log file during ongoing replication.
	//
	// Setting this parameter to `Y` enables you to control active transaction log file growth during full load and ongoing replication tasks. However, it can add some source latency to ongoing replication.
	ReadBackupOnly interface{} `field:"optional" json:"readBackupOnly" yaml:"readBackupOnly"`
	// Use this attribute to minimize the need to access the backup log and enable AWS DMS to prevent truncation using one of the following two methods.
	//
	// *Start transactions in the database:* This is the default method. When this method is used, AWS DMS prevents TLOG truncation by mimicking a transaction in the database. As long as such a transaction is open, changes that appear after the transaction started aren't truncated. If you need Microsoft Replication to be enabled in your database, then you must choose this method.
	//
	// *Exclusively use sp_repldone within a single task* : When this method is used, AWS DMS reads the changes and then uses sp_repldone to mark the TLOG transactions as ready for truncation. Although this method doesn't involve any transactional activities, it can only be used when Microsoft Replication isn't running. Also, when using this method, only one AWS DMS task can access the database at any given time. Therefore, if you need to run parallel AWS DMS tasks against the same database, use the default method.
	SafeguardPolicy *string `field:"optional" json:"safeguardPolicy" yaml:"safeguardPolicy"`
	// The full Amazon Resource Name (ARN) of the IAM role that specifies AWS DMS as the trusted entity and grants the required permissions to access the value in `SecretsManagerSecret` .
	//
	// The role must allow the `iam:PassRole` action. `SecretsManagerSecret` has the value of the AWS Secrets Manager secret that allows access to the SQL Server endpoint.
	//
	// > You can specify one of two sets of values for these permissions. You can specify the values for this setting and `SecretsManagerSecretId` . Or you can specify clear-text values for `UserName` , `Password` , `ServerName` , and `Port` . You can't specify both.
	// >
	// > For more information on creating this `SecretsManagerSecret` , the corresponding `SecretsManagerAccessRoleArn` , and the `SecretsManagerSecretId` that is required to access it, see [Using secrets to access AWS Database Migration Service resources](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Security.html#security-iam-secretsmanager) in the *AWS Database Migration Service User Guide* .
	SecretsManagerAccessRoleArn *string `field:"optional" json:"secretsManagerAccessRoleArn" yaml:"secretsManagerAccessRoleArn"`
	// The full ARN, partial ARN, or display name of the `SecretsManagerSecret` that contains the MicrosoftSQLServer endpoint connection details.
	SecretsManagerSecretId *string `field:"optional" json:"secretsManagerSecretId" yaml:"secretsManagerSecretId"`
	// Use this to attribute to transfer data for full-load operations using BCP.
	//
	// When the target table contains an identity column that does not exist in the source table, you must disable the use BCP for loading table option.
	UseBcpFullLoad interface{} `field:"optional" json:"useBcpFullLoad" yaml:"useBcpFullLoad"`
	// When this attribute is set to `Y` , DMS processes third-party transaction log backups if they are created in native format.
	UseThirdPartyBackupDevice interface{} `field:"optional" json:"useThirdPartyBackupDevice" yaml:"useThirdPartyBackupDevice"`
}

// Provides information that defines a MongoDB endpoint.
//
// This information includes the output format of records applied to the endpoint and details of transaction and control table data information. For more information about other available settings, see [Endpoint configuration settings when using MongoDB as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.MongoDB.html#CHAP_Source.MongoDB.Configuration) in the *AWS Database Migration Service User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mongoDbSettingsProperty := &mongoDbSettingsProperty{
//   	authMechanism: jsii.String("authMechanism"),
//   	authSource: jsii.String("authSource"),
//   	authType: jsii.String("authType"),
//   	databaseName: jsii.String("databaseName"),
//   	docsToInvestigate: jsii.String("docsToInvestigate"),
//   	extractDocId: jsii.String("extractDocId"),
//   	nestingLevel: jsii.String("nestingLevel"),
//   	password: jsii.String("password"),
//   	port: jsii.Number(123),
//   	secretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   	secretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   	serverName: jsii.String("serverName"),
//   	username: jsii.String("username"),
//   }
//
type CfnEndpoint_MongoDbSettingsProperty struct {
	// The authentication mechanism you use to access the MongoDB source endpoint.
	//
	// For the default value, in MongoDB version 2.x, `"default"` is `"mongodb_cr"` . For MongoDB version 3.x or later, `"default"` is `"scram_sha_1"` . This setting isn't used when `AuthType` is set to `"no"` .
	AuthMechanism *string `field:"optional" json:"authMechanism" yaml:"authMechanism"`
	// The MongoDB database name. This setting isn't used when `AuthType` is set to `"no"` .
	//
	// The default is `"admin"` .
	AuthSource *string `field:"optional" json:"authSource" yaml:"authSource"`
	// The authentication type you use to access the MongoDB source endpoint.
	//
	// When set to `"no"` , user name and password parameters are not used and can be empty.
	AuthType *string `field:"optional" json:"authType" yaml:"authType"`
	// The database name on the MongoDB source endpoint.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// Indicates the number of documents to preview to determine the document organization.
	//
	// Use this setting when `NestingLevel` is set to `"one"` .
	//
	// Must be a positive value greater than `0` . Default value is `1000` .
	DocsToInvestigate *string `field:"optional" json:"docsToInvestigate" yaml:"docsToInvestigate"`
	// Specifies the document ID. Use this setting when `NestingLevel` is set to `"none"` .
	//
	// Default value is `"false"` .
	ExtractDocId *string `field:"optional" json:"extractDocId" yaml:"extractDocId"`
	// Specifies either document or table mode.
	//
	// Default value is `"none"` . Specify `"none"` to use document mode. Specify `"one"` to use table mode.
	NestingLevel *string `field:"optional" json:"nestingLevel" yaml:"nestingLevel"`
	// The password for the user account you use to access the MongoDB source endpoint.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The port value for the MongoDB source endpoint.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The full Amazon Resource Name (ARN) of the IAM role that specifies AWS DMS as the trusted entity and grants the required permissions to access the value in `SecretsManagerSecret` .
	//
	// The role must allow the `iam:PassRole` action. `SecretsManagerSecret` has the value of the AWS Secrets Manager secret that allows access to the MongoDB endpoint.
	//
	// > You can specify one of two sets of values for these permissions. You can specify the values for this setting and `SecretsManagerSecretId` . Or you can specify clear-text values for `UserName` , `Password` , `ServerName` , and `Port` . You can't specify both.
	// >
	// > For more information on creating this `SecretsManagerSecret` , the corresponding `SecretsManagerAccessRoleArn` , and the `SecretsManagerSecretId` that is required to access it, see [Using secrets to access AWS Database Migration Service resources](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Security.html#security-iam-secretsmanager) in the *AWS Database Migration Service User Guide* .
	SecretsManagerAccessRoleArn *string `field:"optional" json:"secretsManagerAccessRoleArn" yaml:"secretsManagerAccessRoleArn"`
	// The full ARN, partial ARN, or display name of the `SecretsManagerSecret` that contains the MongoDB endpoint connection details.
	SecretsManagerSecretId *string `field:"optional" json:"secretsManagerSecretId" yaml:"secretsManagerSecretId"`
	// The name of the server on the MongoDB source endpoint.
	ServerName *string `field:"optional" json:"serverName" yaml:"serverName"`
	// The user name you use to access the MongoDB source endpoint.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

// Provides information that defines a MySQL endpoint.
//
// This information includes the output format of records applied to the endpoint and details of transaction and control table data information. For information about other available settings, see [Extra connection attributes when using MySQL as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.MySQL.html#CHAP_Source.MySQL.ConnectionAttrib) and [Extra connection attributes when using a MySQL-compatible database as a target for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.MySQL.html#CHAP_Target.MySQL.ConnectionAttrib) in the *AWS Database Migration Service User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mySqlSettingsProperty := &mySqlSettingsProperty{
//   	afterConnectScript: jsii.String("afterConnectScript"),
//   	cleanSourceMetadataOnMismatch: jsii.Boolean(false),
//   	eventsPollInterval: jsii.Number(123),
//   	maxFileSize: jsii.Number(123),
//   	parallelLoadThreads: jsii.Number(123),
//   	secretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   	secretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   	serverTimezone: jsii.String("serverTimezone"),
//   	targetDbType: jsii.String("targetDbType"),
//   }
//
type CfnEndpoint_MySqlSettingsProperty struct {
	// Specifies a script to run immediately after AWS DMS connects to the endpoint.
	//
	// The migration task continues running regardless if the SQL statement succeeds or fails.
	//
	// For this parameter, provide the code of the script itself, not the name of a file containing the script.
	AfterConnectScript *string `field:"optional" json:"afterConnectScript" yaml:"afterConnectScript"`
	// Adjusts the behavior of DMS when migrating from an SQL Server source database that is hosted as part of an Always On availability group cluster.
	//
	// If you need DMS to poll all the nodes in the Always On cluster for transaction backups, set this attribute to `false` .
	CleanSourceMetadataOnMismatch interface{} `field:"optional" json:"cleanSourceMetadataOnMismatch" yaml:"cleanSourceMetadataOnMismatch"`
	// Specifies how often to check the binary log for new changes/events when the database is idle.
	//
	// The default is five seconds.
	//
	// Example: `eventsPollInterval=5;`
	//
	// In the example, AWS DMS checks for changes in the binary logs every five seconds.
	EventsPollInterval *float64 `field:"optional" json:"eventsPollInterval" yaml:"eventsPollInterval"`
	// Specifies the maximum size (in KB) of any .csv file used to transfer data to a MySQL-compatible database.
	//
	// Example: `maxFileSize=512`.
	MaxFileSize *float64 `field:"optional" json:"maxFileSize" yaml:"maxFileSize"`
	// Improves performance when loading data into the MySQL-compatible target database.
	//
	// Specifies how many threads to use to load the data into the MySQL-compatible target database. Setting a large number of threads can have an adverse effect on database performance, because a separate connection is required for each thread. The default is one.
	//
	// Example: `parallelLoadThreads=1`.
	ParallelLoadThreads *float64 `field:"optional" json:"parallelLoadThreads" yaml:"parallelLoadThreads"`
	// The full Amazon Resource Name (ARN) of the IAM role that specifies AWS DMS as the trusted entity and grants the required permissions to access the value in `SecretsManagerSecret` .
	//
	// The role must allow the `iam:PassRole` action. `SecretsManagerSecret` has the value of the AWS Secrets Manager secret that allows access to the MySQL endpoint.
	//
	// > You can specify one of two sets of values for these permissions. You can specify the values for this setting and `SecretsManagerSecretId` . Or you can specify clear-text values for `UserName` , `Password` , `ServerName` , and `Port` . You can't specify both.
	// >
	// > For more information on creating this `SecretsManagerSecret` , the corresponding `SecretsManagerAccessRoleArn` , and the `SecretsManagerSecretId` that is required to access it, see [Using secrets to access AWS Database Migration Service resources](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Security.html#security-iam-secretsmanager) in the *AWS Database Migration Service User Guide* .
	SecretsManagerAccessRoleArn *string `field:"optional" json:"secretsManagerAccessRoleArn" yaml:"secretsManagerAccessRoleArn"`
	// The full ARN, partial ARN, or display name of the `SecretsManagerSecret` that contains the MySQL endpoint connection details.
	SecretsManagerSecretId *string `field:"optional" json:"secretsManagerSecretId" yaml:"secretsManagerSecretId"`
	// Specifies the time zone for the source MySQL database.
	//
	// Example: `serverTimezone=US/Pacific;`
	//
	// Note: Do not enclose time zones in single quotes.
	ServerTimezone *string `field:"optional" json:"serverTimezone" yaml:"serverTimezone"`
	// Specifies where to migrate source tables on the target, either to a single database or multiple databases.
	//
	// If you specify `SPECIFIC_DATABASE` , specify the database name using the `DatabaseName` parameter of the `Endpoint` object.
	//
	// Example: `targetDbType=MULTIPLE_DATABASES`.
	TargetDbType *string `field:"optional" json:"targetDbType" yaml:"targetDbType"`
}

// Provides information that defines an Amazon Neptune endpoint.
//
// This information includes the output format of records applied to the endpoint and details of transaction and control table data information. For more information about the available settings, see [Specifying endpoint settings for Amazon Neptune as a target](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Neptune.html#CHAP_Target.Neptune.EndpointSettings) in the *AWS Database Migration Service User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   neptuneSettingsProperty := &neptuneSettingsProperty{
//   	errorRetryDuration: jsii.Number(123),
//   	iamAuthEnabled: jsii.Boolean(false),
//   	maxFileSize: jsii.Number(123),
//   	maxRetryCount: jsii.Number(123),
//   	s3BucketFolder: jsii.String("s3BucketFolder"),
//   	s3BucketName: jsii.String("s3BucketName"),
//   	serviceAccessRoleArn: jsii.String("serviceAccessRoleArn"),
//   }
//
type CfnEndpoint_NeptuneSettingsProperty struct {
	// The number of milliseconds for AWS DMS to wait to retry a bulk-load of migrated graph data to the Neptune target database before raising an error.
	//
	// The default is 250.
	ErrorRetryDuration *float64 `field:"optional" json:"errorRetryDuration" yaml:"errorRetryDuration"`
	// If you want IAM authorization enabled for this endpoint, set this parameter to `true` .
	//
	// Then attach the appropriate IAM policy document to your service role specified by `ServiceAccessRoleArn` . The default is `false` .
	IamAuthEnabled interface{} `field:"optional" json:"iamAuthEnabled" yaml:"iamAuthEnabled"`
	// The maximum size in kilobytes of migrated graph data stored in a .csv file before AWS DMS bulk-loads the data to the Neptune target database. The default is 1,048,576 KB. If the bulk load is successful, AWS DMS clears the bucket, ready to store the next batch of migrated graph data.
	MaxFileSize *float64 `field:"optional" json:"maxFileSize" yaml:"maxFileSize"`
	// The number of times for AWS DMS to retry a bulk load of migrated graph data to the Neptune target database before raising an error.
	//
	// The default is 5.
	MaxRetryCount *float64 `field:"optional" json:"maxRetryCount" yaml:"maxRetryCount"`
	// A folder path where you want AWS DMS to store migrated graph data in the S3 bucket specified by `S3BucketName`.
	S3BucketFolder *string `field:"optional" json:"s3BucketFolder" yaml:"s3BucketFolder"`
	// The name of the Amazon S3 bucket where AWS DMS can temporarily store migrated graph data in .csv files before bulk-loading it to the Neptune target database. AWS DMS maps the SQL source data to graph data before storing it in these .csv files.
	S3BucketName *string `field:"optional" json:"s3BucketName" yaml:"s3BucketName"`
	// The Amazon Resource Name (ARN) of the service role that you created for the Neptune target endpoint.
	//
	// The role must allow the `iam:PassRole` action.
	//
	// For more information, see [Creating an IAM Service Role for Accessing Amazon Neptune as a Target](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Neptune.html#CHAP_Target.Neptune.ServiceRole) in the *AWS Database Migration Service User Guide* .
	ServiceAccessRoleArn *string `field:"optional" json:"serviceAccessRoleArn" yaml:"serviceAccessRoleArn"`
}

// Provides information that defines an Oracle endpoint.
//
// This information includes the output format of records applied to the endpoint and details of transaction and control table data information. For information about other available settings, see [Extra connection attributes when using Oracle as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.Oracle.html#CHAP_Source.Oracle.ConnectionAttrib) and [Extra connection attributes when using Oracle as a target for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Oracle.html#CHAP_Target.Oracle.ConnectionAttrib) in the *AWS Database Migration Service User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   oracleSettingsProperty := &oracleSettingsProperty{
//   	accessAlternateDirectly: jsii.Boolean(false),
//   	additionalArchivedLogDestId: jsii.Number(123),
//   	addSupplementalLogging: jsii.Boolean(false),
//   	allowSelectNestedTables: jsii.Boolean(false),
//   	archivedLogDestId: jsii.Number(123),
//   	archivedLogsOnly: jsii.Boolean(false),
//   	asmPassword: jsii.String("asmPassword"),
//   	asmServer: jsii.String("asmServer"),
//   	asmUser: jsii.String("asmUser"),
//   	charLengthSemantics: jsii.String("charLengthSemantics"),
//   	directPathNoLog: jsii.Boolean(false),
//   	directPathParallelLoad: jsii.Boolean(false),
//   	enableHomogenousTablespace: jsii.Boolean(false),
//   	extraArchivedLogDestIds: []interface{}{
//   		jsii.Number(123),
//   	},
//   	failTasksOnLobTruncation: jsii.Boolean(false),
//   	numberDatatypeScale: jsii.Number(123),
//   	oraclePathPrefix: jsii.String("oraclePathPrefix"),
//   	parallelAsmReadThreads: jsii.Number(123),
//   	readAheadBlocks: jsii.Number(123),
//   	readTableSpaceName: jsii.Boolean(false),
//   	replacePathPrefix: jsii.Boolean(false),
//   	retryInterval: jsii.Number(123),
//   	secretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   	secretsManagerOracleAsmAccessRoleArn: jsii.String("secretsManagerOracleAsmAccessRoleArn"),
//   	secretsManagerOracleAsmSecretId: jsii.String("secretsManagerOracleAsmSecretId"),
//   	secretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   	securityDbEncryption: jsii.String("securityDbEncryption"),
//   	securityDbEncryptionName: jsii.String("securityDbEncryptionName"),
//   	spatialDataOptionToGeoJsonFunctionName: jsii.String("spatialDataOptionToGeoJsonFunctionName"),
//   	standbyDelayTime: jsii.Number(123),
//   	useAlternateFolderForOnline: jsii.Boolean(false),
//   	useBFile: jsii.Boolean(false),
//   	useDirectPathFullLoad: jsii.Boolean(false),
//   	useLogminerReader: jsii.Boolean(false),
//   	usePathPrefix: jsii.String("usePathPrefix"),
//   }
//
type CfnEndpoint_OracleSettingsProperty struct {
	// Set this attribute to `false` in order to use the Binary Reader to capture change data for an Amazon RDS for Oracle as the source.
	//
	// This tells the DMS instance to not access redo logs through any specified path prefix replacement using direct file access.
	AccessAlternateDirectly interface{} `field:"optional" json:"accessAlternateDirectly" yaml:"accessAlternateDirectly"`
	// Set this attribute with `ArchivedLogDestId` in a primary/ standby setup.
	//
	// This attribute is useful in the case of a switchover. In this case, AWS DMS needs to know which destination to get archive redo logs from to read changes. This need arises because the previous primary instance is now a standby instance after switchover.
	//
	// Although AWS DMS supports the use of the Oracle `RESETLOGS` option to open the database, never use `RESETLOGS` unless necessary. For additional information about `RESETLOGS` , see [RMAN Data Repair Concepts](https://docs.aws.amazon.com/https://docs.oracle.com/en/database/oracle/oracle-database/19/bradv/rman-data-repair-concepts.html#GUID-1805CCF7-4AF2-482D-B65A-998192F89C2B) in the *Oracle Database Backup and Recovery User's Guide* .
	AdditionalArchivedLogDestId *float64 `field:"optional" json:"additionalArchivedLogDestId" yaml:"additionalArchivedLogDestId"`
	// Set this attribute to set up table-level supplemental logging for the Oracle database.
	//
	// This attribute enables PRIMARY KEY supplemental logging on all tables selected for a migration task.
	//
	// If you use this option, you still need to enable database-level supplemental logging.
	AddSupplementalLogging interface{} `field:"optional" json:"addSupplementalLogging" yaml:"addSupplementalLogging"`
	// Set this attribute to `true` to enable replication of Oracle tables containing columns that are nested tables or defined types.
	AllowSelectNestedTables interface{} `field:"optional" json:"allowSelectNestedTables" yaml:"allowSelectNestedTables"`
	// Specifies the ID of the destination for the archived redo logs.
	//
	// This value should be the same as a number in the dest_id column of the v$archived_log view. If you work with an additional redo log destination, use the `AdditionalArchivedLogDestId` option to specify the additional destination ID. Doing this improves performance by ensuring that the correct logs are accessed from the outset.
	ArchivedLogDestId *float64 `field:"optional" json:"archivedLogDestId" yaml:"archivedLogDestId"`
	// When this field is set to `Y` , AWS DMS only accesses the archived redo logs.
	//
	// If the archived redo logs are stored on Oracle ASM only, the AWS DMS user account needs to be granted ASM privileges.
	ArchivedLogsOnly interface{} `field:"optional" json:"archivedLogsOnly" yaml:"archivedLogsOnly"`
	// For an Oracle source endpoint, your Oracle Automatic Storage Management (ASM) password.
	//
	// You can set this value from the `*asm_user_password*` value. You set this value as part of the comma-separated value that you set to the `Password` request parameter when you create the endpoint to access transaction logs using Binary Reader. For more information, see [Configuration for change data capture (CDC) on an Oracle source database](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.Oracle.html#dms/latest/userguide/CHAP_Source.Oracle.html#CHAP_Source.Oracle.CDC.Configuration) .
	AsmPassword *string `field:"optional" json:"asmPassword" yaml:"asmPassword"`
	// For an Oracle source endpoint, your ASM server address.
	//
	// You can set this value from the `asm_server` value. You set `asm_server` as part of the extra connection attribute string to access an Oracle server with Binary Reader that uses ASM. For more information, see [Configuration for change data capture (CDC) on an Oracle source database](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.Oracle.html#dms/latest/userguide/CHAP_Source.Oracle.html#CHAP_Source.Oracle.CDC.Configuration) .
	AsmServer *string `field:"optional" json:"asmServer" yaml:"asmServer"`
	// For an Oracle source endpoint, your ASM user name.
	//
	// You can set this value from the `asm_user` value. You set `asm_user` as part of the extra connection attribute string to access an Oracle server with Binary Reader that uses ASM. For more information, see [Configuration for change data capture (CDC) on an Oracle source database](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.Oracle.html#dms/latest/userguide/CHAP_Source.Oracle.html#CHAP_Source.Oracle.CDC.Configuration) .
	AsmUser *string `field:"optional" json:"asmUser" yaml:"asmUser"`
	// Specifies whether the length of a character column is in bytes or in characters.
	//
	// To indicate that the character column length is in characters, set this attribute to `CHAR` . Otherwise, the character column length is in bytes.
	//
	// Example: `charLengthSemantics=CHAR;`.
	CharLengthSemantics *string `field:"optional" json:"charLengthSemantics" yaml:"charLengthSemantics"`
	// When set to `true` , this attribute helps to increase the commit rate on the Oracle target database by writing directly to tables and not writing a trail to database logs.
	DirectPathNoLog interface{} `field:"optional" json:"directPathNoLog" yaml:"directPathNoLog"`
	// When set to `true` , this attribute specifies a parallel load when `useDirectPathFullLoad` is set to `Y` .
	//
	// This attribute also only applies when you use the AWS DMS parallel load feature. Note that the target table cannot have any constraints or indexes.
	DirectPathParallelLoad interface{} `field:"optional" json:"directPathParallelLoad" yaml:"directPathParallelLoad"`
	// Set this attribute to enable homogenous tablespace replication and create existing tables or indexes under the same tablespace on the target.
	EnableHomogenousTablespace interface{} `field:"optional" json:"enableHomogenousTablespace" yaml:"enableHomogenousTablespace"`
	// Specifies the IDs of one more destinations for one or more archived redo logs.
	//
	// These IDs are the values of the `dest_id` column in the `v$archived_log` view. Use this setting with the `archivedLogDestId` extra connection attribute in a primary-to-single setup or a primary-to-multiple-standby setup.
	//
	// This setting is useful in a switchover when you use an Oracle Data Guard database as a source. In this case, AWS DMS needs information about what destination to get archive redo logs from to read changes. AWS DMS needs this because after the switchover the previous primary is a standby instance. For example, in a primary-to-single standby setup you might apply the following settings.
	//
	// `archivedLogDestId=1; ExtraArchivedLogDestIds=[2]`
	//
	// In a primary-to-multiple-standby setup, you might apply the following settings.
	//
	// `archivedLogDestId=1; ExtraArchivedLogDestIds=[2,3,4]`
	//
	// Although AWS DMS supports the use of the Oracle `RESETLOGS` option to open the database, never use `RESETLOGS` unless it's necessary. For more information about `RESETLOGS` , see [RMAN Data Repair Concepts](https://docs.aws.amazon.com/https://docs.oracle.com/en/database/oracle/oracle-database/19/bradv/rman-data-repair-concepts.html#GUID-1805CCF7-4AF2-482D-B65A-998192F89C2B) in the *Oracle Database Backup and Recovery User's Guide* .
	ExtraArchivedLogDestIds interface{} `field:"optional" json:"extraArchivedLogDestIds" yaml:"extraArchivedLogDestIds"`
	// When set to `true` , this attribute causes a task to fail if the actual size of an LOB column is greater than the specified `LobMaxSize` .
	//
	// If a task is set to limited LOB mode and this option is set to `true` , the task fails instead of truncating the LOB data.
	FailTasksOnLobTruncation interface{} `field:"optional" json:"failTasksOnLobTruncation" yaml:"failTasksOnLobTruncation"`
	// Specifies the number scale.
	//
	// You can select a scale up to 38, or you can select FLOAT. By default, the NUMBER data type is converted to precision 38, scale 10.
	//
	// Example: `numberDataTypeScale=12`.
	NumberDatatypeScale *float64 `field:"optional" json:"numberDatatypeScale" yaml:"numberDatatypeScale"`
	// Set this string attribute to the required value in order to use the Binary Reader to capture change data for an Amazon RDS for Oracle as the source.
	//
	// This value specifies the default Oracle root used to access the redo logs.
	OraclePathPrefix *string `field:"optional" json:"oraclePathPrefix" yaml:"oraclePathPrefix"`
	// Set this attribute to change the number of threads that DMS configures to perform a change data capture (CDC) load using Oracle Automatic Storage Management (ASM).
	//
	// You can specify an integer value between 2 (the default) and 8 (the maximum). Use this attribute together with the `readAheadBlocks` attribute.
	ParallelAsmReadThreads *float64 `field:"optional" json:"parallelAsmReadThreads" yaml:"parallelAsmReadThreads"`
	// Set this attribute to change the number of read-ahead blocks that DMS configures to perform a change data capture (CDC) load using Oracle Automatic Storage Management (ASM).
	//
	// You can specify an integer value between 1000 (the default) and 200,000 (the maximum).
	ReadAheadBlocks *float64 `field:"optional" json:"readAheadBlocks" yaml:"readAheadBlocks"`
	// When set to `true` , this attribute supports tablespace replication.
	ReadTableSpaceName interface{} `field:"optional" json:"readTableSpaceName" yaml:"readTableSpaceName"`
	// Set this attribute to true in order to use the Binary Reader to capture change data for an Amazon RDS for Oracle as the source.
	//
	// This setting tells DMS instance to replace the default Oracle root with the specified `usePathPrefix` setting to access the redo logs.
	ReplacePathPrefix interface{} `field:"optional" json:"replacePathPrefix" yaml:"replacePathPrefix"`
	// Specifies the number of seconds that the system waits before resending a query.
	//
	// Example: `retryInterval=6;`.
	RetryInterval *float64 `field:"optional" json:"retryInterval" yaml:"retryInterval"`
	// The full Amazon Resource Name (ARN) of the IAM role that specifies AWS DMS as the trusted entity and grants the required permissions to access the value in `SecretsManagerSecret` .
	//
	// The role must allow the `iam:PassRole` action. `SecretsManagerSecret` has the value of the AWS Secrets Manager secret that allows access to the Oracle endpoint.
	//
	// > You can specify one of two sets of values for these permissions. You can specify the values for this setting and `SecretsManagerSecretId` . Or you can specify clear-text values for `UserName` , `Password` , `ServerName` , and `Port` . You can't specify both.
	// >
	// > For more information on creating this `SecretsManagerSecret` , the corresponding `SecretsManagerAccessRoleArn` , and the `SecretsManagerSecretId` that is required to access it, see [Using secrets to access AWS Database Migration Service resources](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Security.html#security-iam-secretsmanager) in the *AWS Database Migration Service User Guide* .
	SecretsManagerAccessRoleArn *string `field:"optional" json:"secretsManagerAccessRoleArn" yaml:"secretsManagerAccessRoleArn"`
	// Required only if your Oracle endpoint uses Advanced Storage Manager (ASM).
	//
	// The full ARN of the IAM role that specifies AWS DMS as the trusted entity and grants the required permissions to access the `SecretsManagerOracleAsmSecret` . This `SecretsManagerOracleAsmSecret` has the secret value that allows access to the Oracle ASM of the endpoint.
	//
	// > You can specify one of two sets of values for these permissions. You can specify the values for this setting and `SecretsManagerOracleAsmSecretId` . Or you can specify clear-text values for `AsmUserName` , `AsmPassword` , and `AsmServerName` . You can't specify both.
	// >
	// > For more information on creating this `SecretsManagerOracleAsmSecret` , the corresponding `SecretsManagerOracleAsmAccessRoleArn` , and the `SecretsManagerOracleAsmSecretId` that is required to access it, see [Using secrets to access AWS Database Migration Service resources](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Security.html#security-iam-secretsmanager) in the *AWS Database Migration Service User Guide* .
	SecretsManagerOracleAsmAccessRoleArn *string `field:"optional" json:"secretsManagerOracleAsmAccessRoleArn" yaml:"secretsManagerOracleAsmAccessRoleArn"`
	// Required only if your Oracle endpoint uses Advanced Storage Manager (ASM).
	//
	// The full ARN, partial ARN, or display name of the `SecretsManagerOracleAsmSecret` that contains the Oracle ASM connection details for the Oracle endpoint.
	SecretsManagerOracleAsmSecretId *string `field:"optional" json:"secretsManagerOracleAsmSecretId" yaml:"secretsManagerOracleAsmSecretId"`
	// The full ARN, partial ARN, or display name of the `SecretsManagerSecret` that contains the Oracle endpoint connection details.
	SecretsManagerSecretId *string `field:"optional" json:"secretsManagerSecretId" yaml:"secretsManagerSecretId"`
	// For an Oracle source endpoint, the transparent data encryption (TDE) password required by AWM DMS to access Oracle redo logs encrypted by TDE using Binary Reader.
	//
	// It is also the `*TDE_Password*` part of the comma-separated value you set to the `Password` request parameter when you create the endpoint. The `SecurityDbEncryptian` setting is related to this `SecurityDbEncryptionName` setting. For more information, see [Supported encryption methods for using Oracle as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.Oracle.html#CHAP_Source.Oracle.Encryption) in the *AWS Database Migration Service User Guide* .
	SecurityDbEncryption *string `field:"optional" json:"securityDbEncryption" yaml:"securityDbEncryption"`
	// For an Oracle source endpoint, the name of a key used for the transparent data encryption (TDE) of the columns and tablespaces in an Oracle source database that is encrypted using TDE.
	//
	// The key value is the value of the `SecurityDbEncryption` setting. For more information on setting the key name value of `SecurityDbEncryptionName` , see the information and example for setting the `securityDbEncryptionName` extra connection attribute in [Supported encryption methods for using Oracle as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.Oracle.html#CHAP_Source.Oracle.Encryption) in the *AWS Database Migration Service User Guide* .
	SecurityDbEncryptionName *string `field:"optional" json:"securityDbEncryptionName" yaml:"securityDbEncryptionName"`
	// Use this attribute to convert `SDO_GEOMETRY` to `GEOJSON` format.
	//
	// By default, DMS calls the `SDO2GEOJSON` custom function if present and accessible. Or you can create your own custom function that mimics the operation of `SDOGEOJSON` and set `SpatialDataOptionToGeoJsonFunctionName` to call it instead.
	SpatialDataOptionToGeoJsonFunctionName *string `field:"optional" json:"spatialDataOptionToGeoJsonFunctionName" yaml:"spatialDataOptionToGeoJsonFunctionName"`
	// Use this attribute to specify a time in minutes for the delay in standby sync.
	//
	// If the source is an Oracle Active Data Guard standby database, use this attribute to specify the time lag between primary and standby databases.
	//
	// In AWS DMS , you can create an Oracle CDC task that uses an Active Data Guard standby instance as a source for replicating ongoing changes. Doing this eliminates the need to connect to an active database that might be in production.
	StandbyDelayTime *float64 `field:"optional" json:"standbyDelayTime" yaml:"standbyDelayTime"`
	// Set this attribute to `true` in order to use the Binary Reader to capture change data for an Amazon RDS for Oracle as the source.
	//
	// This tells the DMS instance to use any specified prefix replacement to access all online redo logs.
	UseAlternateFolderForOnline interface{} `field:"optional" json:"useAlternateFolderForOnline" yaml:"useAlternateFolderForOnline"`
	// Set this attribute to Y to capture change data using the Binary Reader utility.
	//
	// Set `UseLogminerReader` to N to set this attribute to Y. To use Binary Reader with Amazon RDS for Oracle as the source, you set additional attributes. For more information about using this setting with Oracle Automatic Storage Management (ASM), see [Using Oracle LogMiner or AWS DMS Binary Reader for CDC](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.Oracle.html#CHAP_Source.Oracle.CDC) .
	UseBFile interface{} `field:"optional" json:"useBFile" yaml:"useBFile"`
	// Set this attribute to Y to have AWS DMS use a direct path full load.
	//
	// Specify this value to use the direct path protocol in the Oracle Call Interface (OCI). By using this OCI protocol, you can bulk-load Oracle target tables during a full load.
	UseDirectPathFullLoad interface{} `field:"optional" json:"useDirectPathFullLoad" yaml:"useDirectPathFullLoad"`
	// Set this attribute to Y to capture change data using the Oracle LogMiner utility (the default).
	//
	// Set this attribute to N if you want to access the redo logs as a binary file. When you set `UseLogminerReader` to N, also set `UseBfile` to Y. For more information on this setting and using Oracle ASM, see [Using Oracle LogMiner or AWS DMS Binary Reader for CDC](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.Oracle.html#CHAP_Source.Oracle.CDC) in the *AWS DMS User Guide* .
	UseLogminerReader interface{} `field:"optional" json:"useLogminerReader" yaml:"useLogminerReader"`
	// Set this string attribute to the required value in order to use the Binary Reader to capture change data for an Amazon RDS for Oracle as the source.
	//
	// This value specifies the path prefix used to replace the default Oracle root to access the redo logs.
	UsePathPrefix *string `field:"optional" json:"usePathPrefix" yaml:"usePathPrefix"`
}

// Provides information that defines a PostgreSQL endpoint.
//
// This information includes the output format of records applied to the endpoint and details of transaction and control table data information. For information about other available settings, see [Extra connection attributes when using PostgreSQL as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.PostgreSQL.html#CHAP_Source.PostgreSQL.ConnectionAttrib) and [Extra connection attributes when using PostgreSQL as a target for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.PostgreSQL.html#CHAP_Target.PostgreSQL.ConnectionAttrib) in the *AWS Database Migration Service User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   postgreSqlSettingsProperty := &postgreSqlSettingsProperty{
//   	afterConnectScript: jsii.String("afterConnectScript"),
//   	captureDdls: jsii.Boolean(false),
//   	ddlArtifactsSchema: jsii.String("ddlArtifactsSchema"),
//   	executeTimeout: jsii.Number(123),
//   	failTasksOnLobTruncation: jsii.Boolean(false),
//   	heartbeatEnable: jsii.Boolean(false),
//   	heartbeatFrequency: jsii.Number(123),
//   	heartbeatSchema: jsii.String("heartbeatSchema"),
//   	maxFileSize: jsii.Number(123),
//   	pluginName: jsii.String("pluginName"),
//   	secretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   	secretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   	slotName: jsii.String("slotName"),
//   }
//
type CfnEndpoint_PostgreSqlSettingsProperty struct {
	// For use with change data capture (CDC) only, this attribute has AWS DMS bypass foreign keys and user triggers to reduce the time it takes to bulk load data.
	//
	// Example: `afterConnectScript=SET session_replication_role='replica'`.
	AfterConnectScript *string `field:"optional" json:"afterConnectScript" yaml:"afterConnectScript"`
	// To capture DDL events, AWS DMS creates various artifacts in the PostgreSQL database when the task starts.
	//
	// You can later remove these artifacts.
	//
	// If this value is set to `N` , you don't have to create tables or triggers on the source database.
	CaptureDdls interface{} `field:"optional" json:"captureDdls" yaml:"captureDdls"`
	// The schema in which the operational DDL database artifacts are created.
	//
	// Example: `ddlArtifactsSchema=xyzddlschema;`.
	DdlArtifactsSchema *string `field:"optional" json:"ddlArtifactsSchema" yaml:"ddlArtifactsSchema"`
	// Sets the client statement timeout for the PostgreSQL instance, in seconds. The default value is 60 seconds.
	//
	// Example: `executeTimeout=100;`.
	ExecuteTimeout *float64 `field:"optional" json:"executeTimeout" yaml:"executeTimeout"`
	// When set to `true` , this value causes a task to fail if the actual size of a LOB column is greater than the specified `LobMaxSize` .
	//
	// If task is set to Limited LOB mode and this option is set to true, the task fails instead of truncating the LOB data.
	FailTasksOnLobTruncation interface{} `field:"optional" json:"failTasksOnLobTruncation" yaml:"failTasksOnLobTruncation"`
	// The write-ahead log (WAL) heartbeat feature mimics a dummy transaction.
	//
	// By doing this, it prevents idle logical replication slots from holding onto old WAL logs, which can result in storage full situations on the source. This heartbeat keeps `restart_lsn` moving and prevents storage full scenarios.
	HeartbeatEnable interface{} `field:"optional" json:"heartbeatEnable" yaml:"heartbeatEnable"`
	// Sets the WAL heartbeat frequency (in minutes).
	HeartbeatFrequency *float64 `field:"optional" json:"heartbeatFrequency" yaml:"heartbeatFrequency"`
	// Sets the schema in which the heartbeat artifacts are created.
	HeartbeatSchema *string `field:"optional" json:"heartbeatSchema" yaml:"heartbeatSchema"`
	// Specifies the maximum size (in KB) of any .csv file used to transfer data to PostgreSQL.
	//
	// Example: `maxFileSize=512`.
	MaxFileSize *float64 `field:"optional" json:"maxFileSize" yaml:"maxFileSize"`
	// Specifies the plugin to use to create a replication slot.
	PluginName *string `field:"optional" json:"pluginName" yaml:"pluginName"`
	// The full Amazon Resource Name (ARN) of the IAM role that specifies AWS DMS as the trusted entity and grants the required permissions to access the value in `SecretsManagerSecret` .
	//
	// The role must allow the `iam:PassRole` action. `SecretsManagerSecret` has the value of the AWS Secrets Manager secret that allows access to the PostgreSQL endpoint.
	//
	// > You can specify one of two sets of values for these permissions. You can specify the values for this setting and `SecretsManagerSecretId` . Or you can specify clear-text values for `UserName` , `Password` , `ServerName` , and `Port` . You can't specify both.
	// >
	// > For more information on creating this `SecretsManagerSecret` , the corresponding `SecretsManagerAccessRoleArn` , and the `SecretsManagerSecretId` that is required to access it, see [Using secrets to access AWS Database Migration Service resources](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Security.html#security-iam-secretsmanager) in the *AWS Database Migration Service User Guide* .
	SecretsManagerAccessRoleArn *string `field:"optional" json:"secretsManagerAccessRoleArn" yaml:"secretsManagerAccessRoleArn"`
	// The full ARN, partial ARN, or display name of the `SecretsManagerSecret` that contains the PostgreSQL endpoint connection details.
	SecretsManagerSecretId *string `field:"optional" json:"secretsManagerSecretId" yaml:"secretsManagerSecretId"`
	// Sets the name of a previously created logical replication slot for a change data capture (CDC) load of the PostgreSQL source instance.
	//
	// When used with the `CdcStartPosition` request parameter for the AWS DMS API , this attribute also makes it possible to use native CDC start points. DMS verifies that the specified logical replication slot exists before starting the CDC load task. It also verifies that the task was created with a valid setting of `CdcStartPosition` . If the specified slot doesn't exist or the task doesn't have a valid `CdcStartPosition` setting, DMS raises an error.
	//
	// For more information about setting the `CdcStartPosition` request parameter, see [Determining a CDC native start point](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Task.CDC.html#CHAP_Task.CDC.StartPoint.Native) in the *AWS Database Migration Service User Guide* . For more information about using `CdcStartPosition` , see [CreateReplicationTask](https://docs.aws.amazon.com/dms/latest/APIReference/API_CreateReplicationTask.html) , [StartReplicationTask](https://docs.aws.amazon.com/dms/latest/APIReference/API_StartReplicationTask.html) , and [ModifyReplicationTask](https://docs.aws.amazon.com/dms/latest/APIReference/API_ModifyReplicationTask.html) .
	SlotName *string `field:"optional" json:"slotName" yaml:"slotName"`
}

// Provides information that defines a Redis target endpoint.
//
// This information includes the output format of records applied to the endpoint and details of transaction and control table data information. For information about other available settings, see [Specifying endpoint settings for Redis as a target](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Redis.html#CHAP_Target.Redis.EndpointSettings) in the *AWS Database Migration Service User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   redisSettingsProperty := &redisSettingsProperty{
//   	authPassword: jsii.String("authPassword"),
//   	authType: jsii.String("authType"),
//   	authUserName: jsii.String("authUserName"),
//   	port: jsii.Number(123),
//   	serverName: jsii.String("serverName"),
//   	sslCaCertificateArn: jsii.String("sslCaCertificateArn"),
//   	sslSecurityProtocol: jsii.String("sslSecurityProtocol"),
//   }
//
type CfnEndpoint_RedisSettingsProperty struct {
	// The password provided with the `auth-role` and `auth-token` options of the `AuthType` setting for a Redis target endpoint.
	AuthPassword *string `field:"optional" json:"authPassword" yaml:"authPassword"`
	// The type of authentication to perform when connecting to a Redis target.
	//
	// Options include `none` , `auth-token` , and `auth-role` . The `auth-token` option requires an `AuthPassword` value to be provided. The `auth-role` option requires `AuthUserName` and `AuthPassword` values to be provided.
	AuthType *string `field:"optional" json:"authType" yaml:"authType"`
	// The user name provided with the `auth-role` option of the `AuthType` setting for a Redis target endpoint.
	AuthUserName *string `field:"optional" json:"authUserName" yaml:"authUserName"`
	// Transmission Control Protocol (TCP) port for the endpoint.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Fully qualified domain name of the endpoint.
	ServerName *string `field:"optional" json:"serverName" yaml:"serverName"`
	// The Amazon Resource Name (ARN) for the certificate authority (CA) that DMS uses to connect to your Redis target endpoint.
	SslCaCertificateArn *string `field:"optional" json:"sslCaCertificateArn" yaml:"sslCaCertificateArn"`
	// The connection to a Redis target endpoint using Transport Layer Security (TLS).
	//
	// Valid values include `plaintext` and `ssl-encryption` . The default is `ssl-encryption` . The `ssl-encryption` option makes an encrypted connection. Optionally, you can identify an Amazon Resource Name (ARN) for an SSL certificate authority (CA) using the `SslCaCertificateArn` setting. If an ARN isn't given for a CA, DMS uses the Amazon root CA.
	//
	// The `plaintext` option doesn't provide Transport Layer Security (TLS) encryption for traffic between endpoint and database.
	SslSecurityProtocol *string `field:"optional" json:"sslSecurityProtocol" yaml:"sslSecurityProtocol"`
}

// Provides information that defines an Amazon Redshift endpoint.
//
// This information includes the output format of records applied to the endpoint and details of transaction and control table data information. For more information about other available settings, see [Extra connection attributes when using Amazon Redshift as a target for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Redshift.html#CHAP_Target.Redshift.ConnectionAttrib) in the *AWS Database Migration Service User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   redshiftSettingsProperty := &redshiftSettingsProperty{
//   	acceptAnyDate: jsii.Boolean(false),
//   	afterConnectScript: jsii.String("afterConnectScript"),
//   	bucketFolder: jsii.String("bucketFolder"),
//   	bucketName: jsii.String("bucketName"),
//   	caseSensitiveNames: jsii.Boolean(false),
//   	compUpdate: jsii.Boolean(false),
//   	connectionTimeout: jsii.Number(123),
//   	dateFormat: jsii.String("dateFormat"),
//   	emptyAsNull: jsii.Boolean(false),
//   	encryptionMode: jsii.String("encryptionMode"),
//   	explicitIds: jsii.Boolean(false),
//   	fileTransferUploadStreams: jsii.Number(123),
//   	loadTimeout: jsii.Number(123),
//   	maxFileSize: jsii.Number(123),
//   	removeQuotes: jsii.Boolean(false),
//   	replaceChars: jsii.String("replaceChars"),
//   	replaceInvalidChars: jsii.String("replaceInvalidChars"),
//   	secretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   	secretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   	serverSideEncryptionKmsKeyId: jsii.String("serverSideEncryptionKmsKeyId"),
//   	serviceAccessRoleArn: jsii.String("serviceAccessRoleArn"),
//   	timeFormat: jsii.String("timeFormat"),
//   	trimBlanks: jsii.Boolean(false),
//   	truncateColumns: jsii.Boolean(false),
//   	writeBufferSize: jsii.Number(123),
//   }
//
type CfnEndpoint_RedshiftSettingsProperty struct {
	// A value that indicates to allow any date format, including invalid formats such as 00/00/00 00:00:00, to be loaded without generating an error.
	//
	// You can choose `true` or `false` (the default).
	//
	// This parameter applies only to TIMESTAMP and DATE columns. Always use ACCEPTANYDATE with the DATEFORMAT parameter. If the date format for the data doesn't match the DATEFORMAT specification, Amazon Redshift inserts a NULL value into that field.
	AcceptAnyDate interface{} `field:"optional" json:"acceptAnyDate" yaml:"acceptAnyDate"`
	// Code to run after connecting.
	//
	// This parameter should contain the code itself, not the name of a file containing the code.
	AfterConnectScript *string `field:"optional" json:"afterConnectScript" yaml:"afterConnectScript"`
	// An S3 folder where the comma-separated-value (.csv) files are stored before being uploaded to the target Redshift cluster.
	//
	// For full load mode, AWS DMS converts source records into .csv files and loads them to the *BucketFolder/TableID* path. AWS DMS uses the Redshift `COPY` command to upload the .csv files to the target table. The files are deleted once the `COPY` operation has finished. For more information, see [COPY](https://docs.aws.amazon.com/redshift/latest/dg/r_COPY.html) in the *Amazon Redshift Database Developer Guide* .
	//
	// For change-data-capture (CDC) mode, AWS DMS creates a *NetChanges* table, and loads the .csv files to this *BucketFolder/NetChangesTableID* path.
	BucketFolder *string `field:"optional" json:"bucketFolder" yaml:"bucketFolder"`
	// The name of the intermediate S3 bucket used to store .csv files before uploading data to Redshift.
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// If Amazon Redshift is configured to support case sensitive schema names, set `CaseSensitiveNames` to `true` .
	//
	// The default is `false` .
	CaseSensitiveNames interface{} `field:"optional" json:"caseSensitiveNames" yaml:"caseSensitiveNames"`
	// If you set `CompUpdate` to `true` Amazon Redshift applies automatic compression if the table is empty.
	//
	// This applies even if the table columns already have encodings other than `RAW` . If you set `CompUpdate` to `false` , automatic compression is disabled and existing column encodings aren't changed. The default is `true` .
	CompUpdate interface{} `field:"optional" json:"compUpdate" yaml:"compUpdate"`
	// A value that sets the amount of time to wait (in milliseconds) before timing out, beginning from when you initially establish a connection.
	ConnectionTimeout *float64 `field:"optional" json:"connectionTimeout" yaml:"connectionTimeout"`
	// The date format that you are using.
	//
	// Valid values are `auto` (case-sensitive), your date format string enclosed in quotes, or NULL. If this parameter is left unset (NULL), it defaults to a format of 'YYYY-MM-DD'. Using `auto` recognizes most strings, even some that aren't supported when you use a date format string.
	//
	// If your date and time values use formats different from each other, set this to `auto` .
	DateFormat *string `field:"optional" json:"dateFormat" yaml:"dateFormat"`
	// A value that specifies whether AWS DMS should migrate empty CHAR and VARCHAR fields as NULL.
	//
	// A value of `true` sets empty CHAR and VARCHAR fields to null. The default is `false` .
	EmptyAsNull interface{} `field:"optional" json:"emptyAsNull" yaml:"emptyAsNull"`
	// The type of server-side encryption that you want to use for your data.
	//
	// This encryption type is part of the endpoint settings or the extra connections attributes for Amazon S3. You can choose either `SSE_S3` (the default) or `SSE_KMS` .
	//
	// > For the `ModifyEndpoint` operation, you can change the existing value of the `EncryptionMode` parameter from `SSE_KMS` to `SSE_S3` . But you can’t change the existing value from `SSE_S3` to `SSE_KMS` .
	//
	// To use `SSE_S3` , create an AWS Identity and Access Management (IAM) role with a policy that allows `"arn:aws:s3:::*"` to use the following actions: `"s3:PutObject", "s3:ListBucket"`.
	EncryptionMode *string `field:"optional" json:"encryptionMode" yaml:"encryptionMode"`
	// This setting is only valid for a full-load migration task.
	//
	// Set `ExplicitIds` to `true` to have tables with `IDENTITY` columns override their auto-generated values with explicit values loaded from the source data files used to populate the tables. The default is `false` .
	ExplicitIds interface{} `field:"optional" json:"explicitIds" yaml:"explicitIds"`
	// The number of threads used to upload a single file.
	//
	// This parameter accepts a value from 1 through 64. It defaults to 10.
	//
	// The number of parallel streams used to upload a single .csv file to an S3 bucket using S3 Multipart Upload. For more information, see [Multipart upload overview](https://docs.aws.amazon.com/AmazonS3/latest/dev/mpuoverview.html) .
	//
	// `FileTransferUploadStreams` accepts a value from 1 through 64. It defaults to 10.
	FileTransferUploadStreams *float64 `field:"optional" json:"fileTransferUploadStreams" yaml:"fileTransferUploadStreams"`
	// The amount of time to wait (in milliseconds) before timing out of operations performed by AWS DMS on a Redshift cluster, such as Redshift COPY, INSERT, DELETE, and UPDATE.
	LoadTimeout *float64 `field:"optional" json:"loadTimeout" yaml:"loadTimeout"`
	// The maximum size (in KB) of any .csv file used to load data on an S3 bucket and transfer data to Amazon Redshift. It defaults to 1048576KB (1 GB).
	MaxFileSize *float64 `field:"optional" json:"maxFileSize" yaml:"maxFileSize"`
	// A value that specifies to remove surrounding quotation marks from strings in the incoming data.
	//
	// All characters within the quotation marks, including delimiters, are retained. Choose `true` to remove quotation marks. The default is `false` .
	RemoveQuotes interface{} `field:"optional" json:"removeQuotes" yaml:"removeQuotes"`
	// A value that specifies to replaces the invalid characters specified in `ReplaceInvalidChars` , substituting the specified characters instead.
	//
	// The default is `"?"` .
	ReplaceChars *string `field:"optional" json:"replaceChars" yaml:"replaceChars"`
	// A list of characters that you want to replace.
	//
	// Use with `ReplaceChars` .
	ReplaceInvalidChars *string `field:"optional" json:"replaceInvalidChars" yaml:"replaceInvalidChars"`
	// The full Amazon Resource Name (ARN) of the IAM role that specifies AWS DMS as the trusted entity and grants the required permissions to access the value in `SecretsManagerSecret` .
	//
	// The role must allow the `iam:PassRole` action. `SecretsManagerSecret` has the value of the AWS Secrets Manager secret that allows access to the Amazon Redshift endpoint.
	//
	// > You can specify one of two sets of values for these permissions. You can specify the values for this setting and `SecretsManagerSecretId` . Or you can specify clear-text values for `UserName` , `Password` , `ServerName` , and `Port` . You can't specify both.
	// >
	// > For more information on creating this `SecretsManagerSecret` , the corresponding `SecretsManagerAccessRoleArn` , and the `SecretsManagerSecretId` that is required to access it, see [Using secrets to access AWS Database Migration Service resources](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Security.html#security-iam-secretsmanager) in the *AWS Database Migration Service User Guide* .
	SecretsManagerAccessRoleArn *string `field:"optional" json:"secretsManagerAccessRoleArn" yaml:"secretsManagerAccessRoleArn"`
	// The full ARN, partial ARN, or display name of the `SecretsManagerSecret` that contains the Amazon Redshift endpoint connection details.
	SecretsManagerSecretId *string `field:"optional" json:"secretsManagerSecretId" yaml:"secretsManagerSecretId"`
	// The AWS KMS key ID.
	//
	// If you are using `SSE_KMS` for the `EncryptionMode` , provide this key ID. The key that you use needs an attached policy that enables IAM user permissions and allows use of the key.
	ServerSideEncryptionKmsKeyId *string `field:"optional" json:"serverSideEncryptionKmsKeyId" yaml:"serverSideEncryptionKmsKeyId"`
	// The Amazon Resource Name (ARN) of the IAM role that has access to the Amazon Redshift service.
	//
	// The role must allow the `iam:PassRole` action.
	ServiceAccessRoleArn *string `field:"optional" json:"serviceAccessRoleArn" yaml:"serviceAccessRoleArn"`
	// The time format that you want to use.
	//
	// Valid values are `auto` (case-sensitive), `'timeformat_string'` , `'epochsecs'` , or `'epochmillisecs'` . It defaults to 10. Using `auto` recognizes most strings, even some that aren't supported when you use a time format string.
	//
	// If your date and time values use formats different from each other, set this parameter to `auto` .
	TimeFormat *string `field:"optional" json:"timeFormat" yaml:"timeFormat"`
	// A value that specifies to remove the trailing white space characters from a VARCHAR string.
	//
	// This parameter applies only to columns with a VARCHAR data type. Choose `true` to remove unneeded white space. The default is `false` .
	TrimBlanks interface{} `field:"optional" json:"trimBlanks" yaml:"trimBlanks"`
	// A value that specifies to truncate data in columns to the appropriate number of characters, so that the data fits in the column.
	//
	// This parameter applies only to columns with a VARCHAR or CHAR data type, and rows with a size of 4 MB or less. Choose `true` to truncate data. The default is `false` .
	TruncateColumns interface{} `field:"optional" json:"truncateColumns" yaml:"truncateColumns"`
	// The size (in KB) of the in-memory file write buffer used when generating .csv files on the local disk at the DMS replication instance. The default value is 1000 (buffer size is 1000KB).
	WriteBufferSize *float64 `field:"optional" json:"writeBufferSize" yaml:"writeBufferSize"`
}

// Provides information that defines an Amazon S3 endpoint.
//
// This information includes the output format of records applied to the endpoint and details of transaction and control table data information. For more information about the available settings, see [Extra connection attributes when using Amazon S3 as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.S3.html#CHAP_Source.S3.Configuring) and [Extra connection attributes when using Amazon S3 as a target for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.S3.html#CHAP_Target.S3.Configuring) in the *AWS Database Migration Service User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3SettingsProperty := &s3SettingsProperty{
//   	addColumnName: jsii.Boolean(false),
//   	bucketFolder: jsii.String("bucketFolder"),
//   	bucketName: jsii.String("bucketName"),
//   	cannedAclForObjects: jsii.String("cannedAclForObjects"),
//   	cdcInsertsAndUpdates: jsii.Boolean(false),
//   	cdcInsertsOnly: jsii.Boolean(false),
//   	cdcMaxBatchInterval: jsii.Number(123),
//   	cdcMinFileSize: jsii.Number(123),
//   	cdcPath: jsii.String("cdcPath"),
//   	compressionType: jsii.String("compressionType"),
//   	csvDelimiter: jsii.String("csvDelimiter"),
//   	csvNoSupValue: jsii.String("csvNoSupValue"),
//   	csvNullValue: jsii.String("csvNullValue"),
//   	csvRowDelimiter: jsii.String("csvRowDelimiter"),
//   	dataFormat: jsii.String("dataFormat"),
//   	dataPageSize: jsii.Number(123),
//   	datePartitionDelimiter: jsii.String("datePartitionDelimiter"),
//   	datePartitionEnabled: jsii.Boolean(false),
//   	datePartitionSequence: jsii.String("datePartitionSequence"),
//   	datePartitionTimezone: jsii.String("datePartitionTimezone"),
//   	dictPageSizeLimit: jsii.Number(123),
//   	enableStatistics: jsii.Boolean(false),
//   	encodingType: jsii.String("encodingType"),
//   	encryptionMode: jsii.String("encryptionMode"),
//   	externalTableDefinition: jsii.String("externalTableDefinition"),
//   	ignoreHeaderRows: jsii.Number(123),
//   	includeOpForFullLoad: jsii.Boolean(false),
//   	maxFileSize: jsii.Number(123),
//   	parquetTimestampInMillisecond: jsii.Boolean(false),
//   	parquetVersion: jsii.String("parquetVersion"),
//   	preserveTransactions: jsii.Boolean(false),
//   	rfc4180: jsii.Boolean(false),
//   	rowGroupLength: jsii.Number(123),
//   	serverSideEncryptionKmsKeyId: jsii.String("serverSideEncryptionKmsKeyId"),
//   	serviceAccessRoleArn: jsii.String("serviceAccessRoleArn"),
//   	timestampColumnName: jsii.String("timestampColumnName"),
//   	useCsvNoSupValue: jsii.Boolean(false),
//   	useTaskStartTimeForFullLoadTimestamp: jsii.Boolean(false),
//   }
//
type CfnEndpoint_S3SettingsProperty struct {
	// An optional parameter that, when set to `true` or `y` , you can use to add column name information to the .csv output file.
	//
	// The default value is `false` . Valid values are `true` , `false` , `y` , and `n` .
	AddColumnName interface{} `field:"optional" json:"addColumnName" yaml:"addColumnName"`
	// An optional parameter to set a folder name in the S3 bucket.
	//
	// If provided, tables are created in the path `*bucketFolder* / *schema_name* / *table_name* /` . If this parameter isn't specified, the path used is `*schema_name* / *table_name* /` .
	BucketFolder *string `field:"optional" json:"bucketFolder" yaml:"bucketFolder"`
	// The name of the S3 bucket.
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// A value that enables AWS DMS to specify a predefined (canned) access control list (ACL) for objects created in an Amazon S3 bucket as .csv or .parquet files. For more information about Amazon S3 canned ACLs, see [Canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) in the *Amazon S3 Developer Guide* .
	//
	// The default value is NONE. Valid values include NONE, PRIVATE, PUBLIC_READ, PUBLIC_READ_WRITE, AUTHENTICATED_READ, AWS_EXEC_READ, BUCKET_OWNER_READ, and BUCKET_OWNER_FULL_CONTROL.
	CannedAclForObjects *string `field:"optional" json:"cannedAclForObjects" yaml:"cannedAclForObjects"`
	// A value that enables a change data capture (CDC) load to write INSERT and UPDATE operations to .csv or .parquet (columnar storage) output files. The default setting is `false` , but when `CdcInsertsAndUpdates` is set to `true` or `y` , only INSERTs and UPDATEs from the source database are migrated to the .csv or .parquet file.
	//
	// For .csv file format only, how these INSERTs and UPDATEs are recorded depends on the value of the `IncludeOpForFullLoad` parameter. If `IncludeOpForFullLoad` is set to `true` , the first field of every CDC record is set to either `I` or `U` to indicate INSERT and UPDATE operations at the source. But if `IncludeOpForFullLoad` is set to `false` , CDC records are written without an indication of INSERT or UPDATE operations at the source. For more information about how these settings work together, see [Indicating Source DB Operations in Migrated S3 Data](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.S3.html#CHAP_Target.S3.Configuring.InsertOps) in the *AWS Database Migration Service User Guide* .
	//
	// > AWS DMS supports the use of the `CdcInsertsAndUpdates` parameter in versions 3.3.1 and later.
	// >
	// > `CdcInsertsOnly` and `CdcInsertsAndUpdates` can't both be set to `true` for the same endpoint. Set either `CdcInsertsOnly` or `CdcInsertsAndUpdates` to `true` for the same endpoint, but not both.
	CdcInsertsAndUpdates interface{} `field:"optional" json:"cdcInsertsAndUpdates" yaml:"cdcInsertsAndUpdates"`
	// A value that enables a change data capture (CDC) load to write only INSERT operations to .csv or columnar storage (.parquet) output files. By default (the `false` setting), the first field in a .csv or .parquet record contains the letter I (INSERT), U (UPDATE), or D (DELETE). These values indicate whether the row was inserted, updated, or deleted at the source database for a CDC load to the target.
	//
	// If `CdcInsertsOnly` is set to `true` or `y` , only INSERTs from the source database are migrated to the .csv or .parquet file. For .csv format only, how these INSERTs are recorded depends on the value of `IncludeOpForFullLoad` . If `IncludeOpForFullLoad` is set to `true` , the first field of every CDC record is set to I to indicate the INSERT operation at the source. If `IncludeOpForFullLoad` is set to `false` , every CDC record is written without a first field to indicate the INSERT operation at the source. For more information about how these settings work together, see [Indicating Source DB Operations in Migrated S3 Data](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.S3.html#CHAP_Target.S3.Configuring.InsertOps) in the *AWS Database Migration Service User Guide* .
	//
	// > AWS DMS supports the interaction described preceding between the `CdcInsertsOnly` and `IncludeOpForFullLoad` parameters in versions 3.1.4 and later.
	// >
	// > `CdcInsertsOnly` and `CdcInsertsAndUpdates` can't both be set to `true` for the same endpoint. Set either `CdcInsertsOnly` or `CdcInsertsAndUpdates` to `true` for the same endpoint, but not both.
	CdcInsertsOnly interface{} `field:"optional" json:"cdcInsertsOnly" yaml:"cdcInsertsOnly"`
	// Maximum length of the interval, defined in seconds, after which to output a file to Amazon S3.
	//
	// When `CdcMaxBatchInterval` and `CdcMinFileSize` are both specified, the file write is triggered by whichever parameter condition is met first within an AWS DMS CloudFormation template.
	//
	// The default value is 60 seconds.
	CdcMaxBatchInterval *float64 `field:"optional" json:"cdcMaxBatchInterval" yaml:"cdcMaxBatchInterval"`
	// Minimum file size, defined in megabytes, to reach for a file output to Amazon S3.
	//
	// When `CdcMinFileSize` and `CdcMaxBatchInterval` are both specified, the file write is triggered by whichever parameter condition is met first within an AWS DMS CloudFormation template.
	//
	// The default value is 32 MB.
	CdcMinFileSize *float64 `field:"optional" json:"cdcMinFileSize" yaml:"cdcMinFileSize"`
	// Specifies the folder path of CDC files.
	//
	// For an S3 source, this setting is required if a task captures change data; otherwise, it's optional. If `CdcPath` is set, AWS DMS reads CDC files from this path and replicates the data changes to the target endpoint. For an S3 target if you set [`PreserveTransactions`](https://docs.aws.amazon.com/dms/latest/APIReference/API_S3Settings.html#DMS-Type-S3Settings-PreserveTransactions) to `true` , AWS DMS verifies that you have set this parameter to a folder path on your S3 target where AWS DMS can save the transaction order for the CDC load. AWS DMS creates this CDC folder path in either your S3 target working directory or the S3 target location specified by [`BucketFolder`](https://docs.aws.amazon.com/dms/latest/APIReference/API_S3Settings.html#DMS-Type-S3Settings-BucketFolder) and [`BucketName`](https://docs.aws.amazon.com/dms/latest/APIReference/API_S3Settings.html#DMS-Type-S3Settings-BucketName) .
	//
	// For example, if you specify `CdcPath` as `MyChangedData` , and you specify `BucketName` as `MyTargetBucket` but do not specify `BucketFolder` , AWS DMS creates the CDC folder path following: `MyTargetBucket/MyChangedData` .
	//
	// If you specify the same `CdcPath` , and you specify `BucketName` as `MyTargetBucket` and `BucketFolder` as `MyTargetData` , AWS DMS creates the CDC folder path following: `MyTargetBucket/MyTargetData/MyChangedData` .
	//
	// For more information on CDC including transaction order on an S3 target, see [Capturing data changes (CDC) including transaction order on the S3 target](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.S3.html#CHAP_Target.S3.EndpointSettings.CdcPath) .
	//
	// > This setting is supported in AWS DMS versions 3.4.2 and later.
	CdcPath *string `field:"optional" json:"cdcPath" yaml:"cdcPath"`
	// An optional parameter.
	//
	// When set to GZIP it enables the service to compress the target files. To allow the service to write the target files uncompressed, either set this parameter to NONE (the default) or don't specify the parameter at all. This parameter applies to both .csv and .parquet file formats.
	CompressionType *string `field:"optional" json:"compressionType" yaml:"compressionType"`
	// The delimiter used to separate columns in the .csv file for both source and target. The default is a comma.
	CsvDelimiter *string `field:"optional" json:"csvDelimiter" yaml:"csvDelimiter"`
	// This setting only applies if your Amazon S3 output files during a change data capture (CDC) load are written in .csv format. If [`UseCsvNoSupValue`](https://docs.aws.amazon.com/dms/latest/APIReference/API_S3Settings.html#DMS-Type-S3Settings-UseCsvNoSupValue) is set to true, specify a string value that you want AWS DMS to use for all columns not included in the supplemental log. If you do not specify a string value, AWS DMS uses the null value for these columns regardless of the `UseCsvNoSupValue` setting.
	//
	// > This setting is supported in AWS DMS versions 3.4.1 and later.
	CsvNoSupValue *string `field:"optional" json:"csvNoSupValue" yaml:"csvNoSupValue"`
	// An optional parameter that specifies how AWS DMS treats null values.
	//
	// While handling the null value, you can use this parameter to pass a user-defined string as null when writing to the target. For example, when target columns are not nullable, you can use this option to differentiate between the empty string value and the null value. So, if you set this parameter value to the empty string ("" or ''), AWS DMS treats the empty string as the null value instead of `NULL` .
	//
	// The default value is `NULL` . Valid values include any valid string.
	CsvNullValue *string `field:"optional" json:"csvNullValue" yaml:"csvNullValue"`
	// The delimiter used to separate rows in the .csv file for both source and target.
	//
	// The default is a carriage return ( `\n` ).
	CsvRowDelimiter *string `field:"optional" json:"csvRowDelimiter" yaml:"csvRowDelimiter"`
	// The format of the data that you want to use for output. You can choose one of the following:.
	//
	// - `csv` : This is a row-based file format with comma-separated values (.csv).
	// - `parquet` : Apache Parquet (.parquet) is a columnar storage file format that features efficient compression and provides faster query response.
	DataFormat *string `field:"optional" json:"dataFormat" yaml:"dataFormat"`
	// The size of one data page in bytes.
	//
	// This parameter defaults to 1024 * 1024 bytes (1 MiB). This number is used for .parquet file format only.
	DataPageSize *float64 `field:"optional" json:"dataPageSize" yaml:"dataPageSize"`
	// Specifies a date separating delimiter to use during folder partitioning.
	//
	// The default value is `SLASH` . Use this parameter when `DatePartitionedEnabled` is set to `true` .
	DatePartitionDelimiter *string `field:"optional" json:"datePartitionDelimiter" yaml:"datePartitionDelimiter"`
	// When set to `true` , this parameter partitions S3 bucket folders based on transaction commit dates.
	//
	// The default value is `false` . For more information about date-based folder partitioning, see [Using date-based folder partitioning](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.S3.html#CHAP_Target.S3.DatePartitioning) .
	DatePartitionEnabled interface{} `field:"optional" json:"datePartitionEnabled" yaml:"datePartitionEnabled"`
	// Identifies the sequence of the date format to use during folder partitioning.
	//
	// The default value is `YYYYMMDD` . Use this parameter when `DatePartitionedEnabled` is set to `true` .
	DatePartitionSequence *string `field:"optional" json:"datePartitionSequence" yaml:"datePartitionSequence"`
	// When creating an S3 target endpoint, set `DatePartitionTimezone` to convert the current UTC time into a specified time zone.
	//
	// The conversion occurs when a date partition folder is created and a change data capture (CDC) file name is generated. The time zone format is Area/Location. Use this parameter when `DatePartitionedEnabled` is set to `true` , as shown in the following example.
	//
	// `s3-settings='{"DatePartitionEnabled": true, "DatePartitionSequence": "YYYYMMDDHH", "DatePartitionDelimiter": "SLASH", "DatePartitionTimezone":" *Asia/Seoul* ", "BucketName": "dms-nattarat-test"}'`.
	DatePartitionTimezone *string `field:"optional" json:"datePartitionTimezone" yaml:"datePartitionTimezone"`
	// The maximum size of an encoded dictionary page of a column.
	//
	// If the dictionary page exceeds this, this column is stored using an encoding type of `PLAIN` . This parameter defaults to 1024 * 1024 bytes (1 MiB), the maximum size of a dictionary page before it reverts to `PLAIN` encoding. This size is used for .parquet file format only.
	DictPageSizeLimit *float64 `field:"optional" json:"dictPageSizeLimit" yaml:"dictPageSizeLimit"`
	// A value that enables statistics for Parquet pages and row groups.
	//
	// Choose `true` to enable statistics, `false` to disable. Statistics include `NULL` , `DISTINCT` , `MAX` , and `MIN` values. This parameter defaults to `true` . This value is used for .parquet file format only.
	EnableStatistics interface{} `field:"optional" json:"enableStatistics" yaml:"enableStatistics"`
	// The type of encoding that you're using:.
	//
	// - `RLE_DICTIONARY` uses a combination of bit-packing and run-length encoding to store repeated values more efficiently. This is the default.
	// - `PLAIN` doesn't use encoding at all. Values are stored as they are.
	// - `PLAIN_DICTIONARY` builds a dictionary of the values encountered in a given column. The dictionary is stored in a dictionary page for each column chunk.
	EncodingType *string `field:"optional" json:"encodingType" yaml:"encodingType"`
	// The type of server-side encryption that you want to use for your data.
	//
	// This encryption type is part of the endpoint settings or the extra connections attributes for Amazon S3. You can choose either `SSE_S3` (the default) or `SSE_KMS` .
	//
	// > For the `ModifyEndpoint` operation, you can change the existing value of the `EncryptionMode` parameter from `SSE_KMS` to `SSE_S3` . But you can’t change the existing value from `SSE_S3` to `SSE_KMS` .
	//
	// To use `SSE_S3` , you need an IAM role with permission to allow `"arn:aws:s3:::dms-*"` to use the following actions:
	//
	// - `s3:CreateBucket`
	// - `s3:ListBucket`
	// - `s3:DeleteBucket`
	// - `s3:GetBucketLocation`
	// - `s3:GetObject`
	// - `s3:PutObject`
	// - `s3:DeleteObject`
	// - `s3:GetObjectVersion`
	// - `s3:GetBucketPolicy`
	// - `s3:PutBucketPolicy`
	// - `s3:DeleteBucketPolicy`.
	EncryptionMode *string `field:"optional" json:"encryptionMode" yaml:"encryptionMode"`
	// The external table definition.
	//
	// Conditional: If `S3` is used as a source then `ExternalTableDefinition` is required.
	ExternalTableDefinition *string `field:"optional" json:"externalTableDefinition" yaml:"externalTableDefinition"`
	// When this value is set to 1, AWS DMS ignores the first row header in a .csv file. A value of 1 turns on the feature; a value of 0 turns off the feature.
	//
	// The default is 0.
	IgnoreHeaderRows *float64 `field:"optional" json:"ignoreHeaderRows" yaml:"ignoreHeaderRows"`
	// A value that enables a full load to write INSERT operations to the comma-separated value (.csv) output files only to indicate how the rows were added to the source database.
	//
	// > AWS DMS supports the `IncludeOpForFullLoad` parameter in versions 3.1.4 and later.
	//
	// For full load, records can only be inserted. By default (the `false` setting), no information is recorded in these output files for a full load to indicate that the rows were inserted at the source database. If `IncludeOpForFullLoad` is set to `true` or `y` , the INSERT is recorded as an I annotation in the first field of the .csv file. This allows the format of your target records from a full load to be consistent with the target records from a CDC load.
	//
	// > This setting works together with the `CdcInsertsOnly` and the `CdcInsertsAndUpdates` parameters for output to .csv files only. For more information about how these settings work together, see [Indicating Source DB Operations in Migrated S3 Data](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.S3.html#CHAP_Target.S3.Configuring.InsertOps) in the *AWS Database Migration Service User Guide* .
	IncludeOpForFullLoad interface{} `field:"optional" json:"includeOpForFullLoad" yaml:"includeOpForFullLoad"`
	// A value that specifies the maximum size (in KB) of any .csv file to be created while migrating to an S3 target during full load.
	//
	// The default value is 1,048,576 KB (1 GB). Valid values include 1 to 1,048,576.
	MaxFileSize *float64 `field:"optional" json:"maxFileSize" yaml:"maxFileSize"`
	// A value that specifies the precision of any `TIMESTAMP` column values that are written to an Amazon S3 object file in .parquet format.
	//
	// > AWS DMS supports the `ParquetTimestampInMillisecond` parameter in versions 3.1.4 and later.
	//
	// When `ParquetTimestampInMillisecond` is set to `true` or `y` , AWS DMS writes all `TIMESTAMP` columns in a .parquet formatted file with millisecond precision. Otherwise, DMS writes them with microsecond precision.
	//
	// Currently, Amazon Athena and AWS Glue can handle only millisecond precision for `TIMESTAMP` values. Set this parameter to `true` for S3 endpoint object files that are .parquet formatted only if you plan to query or process the data with Athena or AWS Glue .
	//
	// > AWS DMS writes any `TIMESTAMP` column values written to an S3 file in .csv format with microsecond precision.
	// >
	// > Setting `ParquetTimestampInMillisecond` has no effect on the string format of the timestamp column value that is inserted by setting the `TimestampColumnName` parameter.
	ParquetTimestampInMillisecond interface{} `field:"optional" json:"parquetTimestampInMillisecond" yaml:"parquetTimestampInMillisecond"`
	// The version of the Apache Parquet format that you want to use: `parquet_1_0` (the default) or `parquet_2_0` .
	ParquetVersion *string `field:"optional" json:"parquetVersion" yaml:"parquetVersion"`
	// If this setting is set to `true` , AWS DMS saves the transaction order for a change data capture (CDC) load on the Amazon S3 target specified by [`CdcPath`](https://docs.aws.amazon.com/dms/latest/APIReference/API_S3Settings.html#DMS-Type-S3Settings-CdcPath) . For more information, see [Capturing data changes (CDC) including transaction order on the S3 target](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.S3.html#CHAP_Target.S3.EndpointSettings.CdcPath) .
	//
	// > This setting is supported in AWS DMS versions 3.4.2 and later.
	PreserveTransactions interface{} `field:"optional" json:"preserveTransactions" yaml:"preserveTransactions"`
	// For an S3 source, when this value is set to `true` or `y` , each leading double quotation mark has to be followed by an ending double quotation mark.
	//
	// This formatting complies with RFC 4180. When this value is set to `false` or `n` , string literals are copied to the target as is. In this case, a delimiter (row or column) signals the end of the field. Thus, you can't use a delimiter as part of the string, because it signals the end of the value.
	//
	// For an S3 target, an optional parameter used to set behavior to comply with RFC 4180 for data migrated to Amazon S3 using .csv file format only. When this value is set to `true` or `y` using Amazon S3 as a target, if the data has quotation marks or newline characters in it, AWS DMS encloses the entire column with an additional pair of double quotation marks ("). Every quotation mark within the data is repeated twice.
	//
	// The default value is `true` . Valid values include `true` , `false` , `y` , and `n` .
	Rfc4180 interface{} `field:"optional" json:"rfc4180" yaml:"rfc4180"`
	// The number of rows in a row group.
	//
	// A smaller row group size provides faster reads. But as the number of row groups grows, the slower writes become. This parameter defaults to 10,000 rows. This number is used for .parquet file format only.
	//
	// If you choose a value larger than the maximum, `RowGroupLength` is set to the max row group length in bytes (64 * 1024 * 1024).
	RowGroupLength *float64 `field:"optional" json:"rowGroupLength" yaml:"rowGroupLength"`
	// If you are using `SSE_KMS` for the `EncryptionMode` , provide the AWS KMS key ID.
	//
	// The key that you use needs an attached policy that enables IAM user permissions and allows use of the key.
	//
	// Here is a CLI example: `aws dms create-endpoint --endpoint-identifier *value* --endpoint-type target --engine-name s3 --s3-settings ServiceAccessRoleArn= *value* ,BucketFolder= *value* ,BucketName= *value* ,EncryptionMode=SSE_KMS,ServerSideEncryptionKmsKeyId= *value*`.
	ServerSideEncryptionKmsKeyId *string `field:"optional" json:"serverSideEncryptionKmsKeyId" yaml:"serverSideEncryptionKmsKeyId"`
	// A required parameter that specifies the Amazon Resource Name (ARN) used by the service to access the IAM role.
	//
	// The role must allow the `iam:PassRole` action. It enables AWS DMS to read and write objects from an S3 bucket.
	ServiceAccessRoleArn *string `field:"optional" json:"serviceAccessRoleArn" yaml:"serviceAccessRoleArn"`
	// A value that when nonblank causes AWS DMS to add a column with timestamp information to the endpoint data for an Amazon S3 target.
	//
	// > AWS DMS supports the `TimestampColumnName` parameter in versions 3.1.4 and later.
	//
	// AWS DMS includes an additional `STRING` column in the .csv or .parquet object files of your migrated data when you set `TimestampColumnName` to a nonblank value.
	//
	// For a full load, each row of this timestamp column contains a timestamp for when the data was transferred from the source to the target by DMS.
	//
	// For a change data capture (CDC) load, each row of the timestamp column contains the timestamp for the commit of that row in the source database.
	//
	// The string format for this timestamp column value is `yyyy-MM-dd HH:mm:ss.SSSSSS` . By default, the precision of this value is in microseconds. For a CDC load, the rounding of the precision depends on the commit timestamp supported by DMS for the source database.
	//
	// When the `AddColumnName` parameter is set to `true` , DMS also includes a name for the timestamp column that you set with `TimestampColumnName` .
	TimestampColumnName *string `field:"optional" json:"timestampColumnName" yaml:"timestampColumnName"`
	// This setting applies if the S3 output files during a change data capture (CDC) load are written in .csv format. If this setting is set to `true` for columns not included in the supplemental log, AWS DMS uses the value specified by [`CsvNoSupValue`](https://docs.aws.amazon.com/dms/latest/APIReference/API_S3Settings.html#DMS-Type-S3Settings-CsvNoSupValue) . If this setting isn't set or is set to `false` , AWS DMS uses the null value for these columns.
	//
	// > This setting is supported in AWS DMS versions 3.4.1 and later.
	UseCsvNoSupValue interface{} `field:"optional" json:"useCsvNoSupValue" yaml:"useCsvNoSupValue"`
	// When set to true, this parameter uses the task start time as the timestamp column value instead of the time data is written to target.
	//
	// For full load, when `useTaskStartTimeForFullLoadTimestamp` is set to `true` , each row of the timestamp column contains the task start time. For CDC loads, each row of the timestamp column contains the transaction commit time.
	//
	// When `useTaskStartTimeForFullLoadTimestamp` is set to `false` , the full load timestamp in the timestamp column increments with the time data arrives at the target.
	UseTaskStartTimeForFullLoadTimestamp interface{} `field:"optional" json:"useTaskStartTimeForFullLoadTimestamp" yaml:"useTaskStartTimeForFullLoadTimestamp"`
}

// Provides information that defines a SAP ASE endpoint.
//
// This information includes the output format of records applied to the endpoint and details of transaction and control table data information. For information about other available settings, see [Extra connection attributes when using SAP ASE as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.SAP.html#CHAP_Source.SAP.ConnectionAttrib) and [Extra connection attributes when using SAP ASE as a target for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.SAP.html#CHAP_Target.SAP.ConnectionAttrib) in the *AWS Database Migration Service User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sybaseSettingsProperty := &sybaseSettingsProperty{
//   	secretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   	secretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   }
//
type CfnEndpoint_SybaseSettingsProperty struct {
	// The full Amazon Resource Name (ARN) of the IAM role that specifies AWS DMS as the trusted entity and grants the required permissions to access the value in `SecretsManagerSecret` .
	//
	// The role must allow the `iam:PassRole` action. `SecretsManagerSecret` has the value of the AWS Secrets Manager secret that allows access to the SAP ASE endpoint.
	//
	// > You can specify one of two sets of values for these permissions. You can specify the values for this setting and `SecretsManagerSecretId` . Or you can specify clear-text values for `UserName` , `Password` , `ServerName` , and `Port` . You can't specify both.
	// >
	// > For more information on creating this `SecretsManagerSecret` , the corresponding `SecretsManagerAccessRoleArn` , and the `SecretsManagerSecretId` that is required to access it, see [Using secrets to access AWS Database Migration Service resources](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Security.html#security-iam-secretsmanager) in the *AWS Database Migration Service User Guide* .
	SecretsManagerAccessRoleArn *string `field:"optional" json:"secretsManagerAccessRoleArn" yaml:"secretsManagerAccessRoleArn"`
	// The full ARN, partial ARN, or display name of the `SecretsManagerSecret` that contains the SAP SAE endpoint connection details.
	SecretsManagerSecretId *string `field:"optional" json:"secretsManagerSecretId" yaml:"secretsManagerSecretId"`
}

// Properties for defining a `CfnEndpoint`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEndpointProps := &cfnEndpointProps{
//   	endpointType: jsii.String("endpointType"),
//   	engineName: jsii.String("engineName"),
//
//   	// the properties below are optional
//   	certificateArn: jsii.String("certificateArn"),
//   	databaseName: jsii.String("databaseName"),
//   	docDbSettings: &docDbSettingsProperty{
//   		docsToInvestigate: jsii.Number(123),
//   		extractDocId: jsii.Boolean(false),
//   		nestingLevel: jsii.String("nestingLevel"),
//   		secretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   		secretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   	},
//   	dynamoDbSettings: &dynamoDbSettingsProperty{
//   		serviceAccessRoleArn: jsii.String("serviceAccessRoleArn"),
//   	},
//   	elasticsearchSettings: &elasticsearchSettingsProperty{
//   		endpointUri: jsii.String("endpointUri"),
//   		errorRetryDuration: jsii.Number(123),
//   		fullLoadErrorPercentage: jsii.Number(123),
//   		serviceAccessRoleArn: jsii.String("serviceAccessRoleArn"),
//   	},
//   	endpointIdentifier: jsii.String("endpointIdentifier"),
//   	extraConnectionAttributes: jsii.String("extraConnectionAttributes"),
//   	gcpMySqlSettings: &gcpMySQLSettingsProperty{
//   		afterConnectScript: jsii.String("afterConnectScript"),
//   		cleanSourceMetadataOnMismatch: jsii.Boolean(false),
//   		databaseName: jsii.String("databaseName"),
//   		eventsPollInterval: jsii.Number(123),
//   		maxFileSize: jsii.Number(123),
//   		parallelLoadThreads: jsii.Number(123),
//   		password: jsii.String("password"),
//   		port: jsii.Number(123),
//   		secretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   		secretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   		serverName: jsii.String("serverName"),
//   		serverTimezone: jsii.String("serverTimezone"),
//   		username: jsii.String("username"),
//   	},
//   	ibmDb2Settings: &ibmDb2SettingsProperty{
//   		currentLsn: jsii.String("currentLsn"),
//   		maxKBytesPerRead: jsii.Number(123),
//   		secretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   		secretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   		setDataCaptureChanges: jsii.Boolean(false),
//   	},
//   	kafkaSettings: &kafkaSettingsProperty{
//   		broker: jsii.String("broker"),
//   		includeControlDetails: jsii.Boolean(false),
//   		includeNullAndEmpty: jsii.Boolean(false),
//   		includePartitionValue: jsii.Boolean(false),
//   		includeTableAlterOperations: jsii.Boolean(false),
//   		includeTransactionDetails: jsii.Boolean(false),
//   		messageFormat: jsii.String("messageFormat"),
//   		messageMaxBytes: jsii.Number(123),
//   		noHexPrefix: jsii.Boolean(false),
//   		partitionIncludeSchemaTable: jsii.Boolean(false),
//   		saslPassword: jsii.String("saslPassword"),
//   		saslUserName: jsii.String("saslUserName"),
//   		securityProtocol: jsii.String("securityProtocol"),
//   		sslCaCertificateArn: jsii.String("sslCaCertificateArn"),
//   		sslClientCertificateArn: jsii.String("sslClientCertificateArn"),
//   		sslClientKeyArn: jsii.String("sslClientKeyArn"),
//   		sslClientKeyPassword: jsii.String("sslClientKeyPassword"),
//   		topic: jsii.String("topic"),
//   	},
//   	kinesisSettings: &kinesisSettingsProperty{
//   		includeControlDetails: jsii.Boolean(false),
//   		includeNullAndEmpty: jsii.Boolean(false),
//   		includePartitionValue: jsii.Boolean(false),
//   		includeTableAlterOperations: jsii.Boolean(false),
//   		includeTransactionDetails: jsii.Boolean(false),
//   		messageFormat: jsii.String("messageFormat"),
//   		noHexPrefix: jsii.Boolean(false),
//   		partitionIncludeSchemaTable: jsii.Boolean(false),
//   		serviceAccessRoleArn: jsii.String("serviceAccessRoleArn"),
//   		streamArn: jsii.String("streamArn"),
//   	},
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	microsoftSqlServerSettings: &microsoftSqlServerSettingsProperty{
//   		bcpPacketSize: jsii.Number(123),
//   		controlTablesFileGroup: jsii.String("controlTablesFileGroup"),
//   		querySingleAlwaysOnNode: jsii.Boolean(false),
//   		readBackupOnly: jsii.Boolean(false),
//   		safeguardPolicy: jsii.String("safeguardPolicy"),
//   		secretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   		secretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   		useBcpFullLoad: jsii.Boolean(false),
//   		useThirdPartyBackupDevice: jsii.Boolean(false),
//   	},
//   	mongoDbSettings: &mongoDbSettingsProperty{
//   		authMechanism: jsii.String("authMechanism"),
//   		authSource: jsii.String("authSource"),
//   		authType: jsii.String("authType"),
//   		databaseName: jsii.String("databaseName"),
//   		docsToInvestigate: jsii.String("docsToInvestigate"),
//   		extractDocId: jsii.String("extractDocId"),
//   		nestingLevel: jsii.String("nestingLevel"),
//   		password: jsii.String("password"),
//   		port: jsii.Number(123),
//   		secretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   		secretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   		serverName: jsii.String("serverName"),
//   		username: jsii.String("username"),
//   	},
//   	mySqlSettings: &mySqlSettingsProperty{
//   		afterConnectScript: jsii.String("afterConnectScript"),
//   		cleanSourceMetadataOnMismatch: jsii.Boolean(false),
//   		eventsPollInterval: jsii.Number(123),
//   		maxFileSize: jsii.Number(123),
//   		parallelLoadThreads: jsii.Number(123),
//   		secretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   		secretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   		serverTimezone: jsii.String("serverTimezone"),
//   		targetDbType: jsii.String("targetDbType"),
//   	},
//   	neptuneSettings: &neptuneSettingsProperty{
//   		errorRetryDuration: jsii.Number(123),
//   		iamAuthEnabled: jsii.Boolean(false),
//   		maxFileSize: jsii.Number(123),
//   		maxRetryCount: jsii.Number(123),
//   		s3BucketFolder: jsii.String("s3BucketFolder"),
//   		s3BucketName: jsii.String("s3BucketName"),
//   		serviceAccessRoleArn: jsii.String("serviceAccessRoleArn"),
//   	},
//   	oracleSettings: &oracleSettingsProperty{
//   		accessAlternateDirectly: jsii.Boolean(false),
//   		additionalArchivedLogDestId: jsii.Number(123),
//   		addSupplementalLogging: jsii.Boolean(false),
//   		allowSelectNestedTables: jsii.Boolean(false),
//   		archivedLogDestId: jsii.Number(123),
//   		archivedLogsOnly: jsii.Boolean(false),
//   		asmPassword: jsii.String("asmPassword"),
//   		asmServer: jsii.String("asmServer"),
//   		asmUser: jsii.String("asmUser"),
//   		charLengthSemantics: jsii.String("charLengthSemantics"),
//   		directPathNoLog: jsii.Boolean(false),
//   		directPathParallelLoad: jsii.Boolean(false),
//   		enableHomogenousTablespace: jsii.Boolean(false),
//   		extraArchivedLogDestIds: []interface{}{
//   			jsii.Number(123),
//   		},
//   		failTasksOnLobTruncation: jsii.Boolean(false),
//   		numberDatatypeScale: jsii.Number(123),
//   		oraclePathPrefix: jsii.String("oraclePathPrefix"),
//   		parallelAsmReadThreads: jsii.Number(123),
//   		readAheadBlocks: jsii.Number(123),
//   		readTableSpaceName: jsii.Boolean(false),
//   		replacePathPrefix: jsii.Boolean(false),
//   		retryInterval: jsii.Number(123),
//   		secretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   		secretsManagerOracleAsmAccessRoleArn: jsii.String("secretsManagerOracleAsmAccessRoleArn"),
//   		secretsManagerOracleAsmSecretId: jsii.String("secretsManagerOracleAsmSecretId"),
//   		secretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   		securityDbEncryption: jsii.String("securityDbEncryption"),
//   		securityDbEncryptionName: jsii.String("securityDbEncryptionName"),
//   		spatialDataOptionToGeoJsonFunctionName: jsii.String("spatialDataOptionToGeoJsonFunctionName"),
//   		standbyDelayTime: jsii.Number(123),
//   		useAlternateFolderForOnline: jsii.Boolean(false),
//   		useBFile: jsii.Boolean(false),
//   		useDirectPathFullLoad: jsii.Boolean(false),
//   		useLogminerReader: jsii.Boolean(false),
//   		usePathPrefix: jsii.String("usePathPrefix"),
//   	},
//   	password: jsii.String("password"),
//   	port: jsii.Number(123),
//   	postgreSqlSettings: &postgreSqlSettingsProperty{
//   		afterConnectScript: jsii.String("afterConnectScript"),
//   		captureDdls: jsii.Boolean(false),
//   		ddlArtifactsSchema: jsii.String("ddlArtifactsSchema"),
//   		executeTimeout: jsii.Number(123),
//   		failTasksOnLobTruncation: jsii.Boolean(false),
//   		heartbeatEnable: jsii.Boolean(false),
//   		heartbeatFrequency: jsii.Number(123),
//   		heartbeatSchema: jsii.String("heartbeatSchema"),
//   		maxFileSize: jsii.Number(123),
//   		pluginName: jsii.String("pluginName"),
//   		secretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   		secretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   		slotName: jsii.String("slotName"),
//   	},
//   	redisSettings: &redisSettingsProperty{
//   		authPassword: jsii.String("authPassword"),
//   		authType: jsii.String("authType"),
//   		authUserName: jsii.String("authUserName"),
//   		port: jsii.Number(123),
//   		serverName: jsii.String("serverName"),
//   		sslCaCertificateArn: jsii.String("sslCaCertificateArn"),
//   		sslSecurityProtocol: jsii.String("sslSecurityProtocol"),
//   	},
//   	redshiftSettings: &redshiftSettingsProperty{
//   		acceptAnyDate: jsii.Boolean(false),
//   		afterConnectScript: jsii.String("afterConnectScript"),
//   		bucketFolder: jsii.String("bucketFolder"),
//   		bucketName: jsii.String("bucketName"),
//   		caseSensitiveNames: jsii.Boolean(false),
//   		compUpdate: jsii.Boolean(false),
//   		connectionTimeout: jsii.Number(123),
//   		dateFormat: jsii.String("dateFormat"),
//   		emptyAsNull: jsii.Boolean(false),
//   		encryptionMode: jsii.String("encryptionMode"),
//   		explicitIds: jsii.Boolean(false),
//   		fileTransferUploadStreams: jsii.Number(123),
//   		loadTimeout: jsii.Number(123),
//   		maxFileSize: jsii.Number(123),
//   		removeQuotes: jsii.Boolean(false),
//   		replaceChars: jsii.String("replaceChars"),
//   		replaceInvalidChars: jsii.String("replaceInvalidChars"),
//   		secretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   		secretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   		serverSideEncryptionKmsKeyId: jsii.String("serverSideEncryptionKmsKeyId"),
//   		serviceAccessRoleArn: jsii.String("serviceAccessRoleArn"),
//   		timeFormat: jsii.String("timeFormat"),
//   		trimBlanks: jsii.Boolean(false),
//   		truncateColumns: jsii.Boolean(false),
//   		writeBufferSize: jsii.Number(123),
//   	},
//   	resourceIdentifier: jsii.String("resourceIdentifier"),
//   	s3Settings: &s3SettingsProperty{
//   		addColumnName: jsii.Boolean(false),
//   		bucketFolder: jsii.String("bucketFolder"),
//   		bucketName: jsii.String("bucketName"),
//   		cannedAclForObjects: jsii.String("cannedAclForObjects"),
//   		cdcInsertsAndUpdates: jsii.Boolean(false),
//   		cdcInsertsOnly: jsii.Boolean(false),
//   		cdcMaxBatchInterval: jsii.Number(123),
//   		cdcMinFileSize: jsii.Number(123),
//   		cdcPath: jsii.String("cdcPath"),
//   		compressionType: jsii.String("compressionType"),
//   		csvDelimiter: jsii.String("csvDelimiter"),
//   		csvNoSupValue: jsii.String("csvNoSupValue"),
//   		csvNullValue: jsii.String("csvNullValue"),
//   		csvRowDelimiter: jsii.String("csvRowDelimiter"),
//   		dataFormat: jsii.String("dataFormat"),
//   		dataPageSize: jsii.Number(123),
//   		datePartitionDelimiter: jsii.String("datePartitionDelimiter"),
//   		datePartitionEnabled: jsii.Boolean(false),
//   		datePartitionSequence: jsii.String("datePartitionSequence"),
//   		datePartitionTimezone: jsii.String("datePartitionTimezone"),
//   		dictPageSizeLimit: jsii.Number(123),
//   		enableStatistics: jsii.Boolean(false),
//   		encodingType: jsii.String("encodingType"),
//   		encryptionMode: jsii.String("encryptionMode"),
//   		externalTableDefinition: jsii.String("externalTableDefinition"),
//   		ignoreHeaderRows: jsii.Number(123),
//   		includeOpForFullLoad: jsii.Boolean(false),
//   		maxFileSize: jsii.Number(123),
//   		parquetTimestampInMillisecond: jsii.Boolean(false),
//   		parquetVersion: jsii.String("parquetVersion"),
//   		preserveTransactions: jsii.Boolean(false),
//   		rfc4180: jsii.Boolean(false),
//   		rowGroupLength: jsii.Number(123),
//   		serverSideEncryptionKmsKeyId: jsii.String("serverSideEncryptionKmsKeyId"),
//   		serviceAccessRoleArn: jsii.String("serviceAccessRoleArn"),
//   		timestampColumnName: jsii.String("timestampColumnName"),
//   		useCsvNoSupValue: jsii.Boolean(false),
//   		useTaskStartTimeForFullLoadTimestamp: jsii.Boolean(false),
//   	},
//   	serverName: jsii.String("serverName"),
//   	sslMode: jsii.String("sslMode"),
//   	sybaseSettings: &sybaseSettingsProperty{
//   		secretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   		secretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	username: jsii.String("username"),
//   }
//
type CfnEndpointProps struct {
	// The type of endpoint.
	//
	// Valid values are `source` and `target` .
	EndpointType *string `field:"required" json:"endpointType" yaml:"endpointType"`
	// The type of engine for the endpoint, depending on the `EndpointType` value.
	//
	// *Valid values* : `mysql` | `oracle` | `postgres` | `mariadb` | `aurora` | `aurora-postgresql` | `opensearch` | `redshift` | `s3` | `db2` | `azuredb` | `sybase` | `dynamodb` | `mongodb` | `kinesis` | `kafka` | `elasticsearch` | `docdb` | `sqlserver` | `neptune`.
	EngineName *string `field:"required" json:"engineName" yaml:"engineName"`
	// The Amazon Resource Name (ARN) for the certificate.
	CertificateArn *string `field:"optional" json:"certificateArn" yaml:"certificateArn"`
	// The name of the endpoint database.
	//
	// For a MySQL source or target endpoint, don't specify `DatabaseName` . To migrate to a specific database, use this setting and `targetDbType` .
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// Settings in JSON format for the source and target DocumentDB endpoint.
	//
	// For more information about other available settings, see [Using extra connections attributes with Amazon DocumentDB as a source](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.DocumentDB.html#CHAP_Source.DocumentDB.ECAs) and [Using Amazon DocumentDB as a target for AWS Database Migration Service](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.DocumentDB.html) in the *AWS Database Migration Service User Guide* .
	DocDbSettings interface{} `field:"optional" json:"docDbSettings" yaml:"docDbSettings"`
	// Settings in JSON format for the target Amazon DynamoDB endpoint.
	//
	// For information about other available settings, see [Using object mapping to migrate data to DynamoDB](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.DynamoDB.html#CHAP_Target.DynamoDB.ObjectMapping) in the *AWS Database Migration Service User Guide* .
	DynamoDbSettings interface{} `field:"optional" json:"dynamoDbSettings" yaml:"dynamoDbSettings"`
	// Settings in JSON format for the target OpenSearch endpoint.
	//
	// For more information about the available settings, see [Extra connection attributes when using OpenSearch as a target for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Elasticsearch.html#CHAP_Target.Elasticsearch.Configuration) in the *AWS Database Migration Service User Guide* .
	ElasticsearchSettings interface{} `field:"optional" json:"elasticsearchSettings" yaml:"elasticsearchSettings"`
	// The database endpoint identifier.
	//
	// Identifiers must begin with a letter and must contain only ASCII letters, digits, and hyphens. They can't end with a hyphen, or contain two consecutive hyphens.
	EndpointIdentifier *string `field:"optional" json:"endpointIdentifier" yaml:"endpointIdentifier"`
	// Additional attributes associated with the connection.
	//
	// Each attribute is specified as a name-value pair associated by an equal sign (=). Multiple attributes are separated by a semicolon (;) with no additional white space. For information on the attributes available for connecting your source or target endpoint, see [Working with AWS DMS Endpoints](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Endpoints.html) in the *AWS Database Migration Service User Guide* .
	ExtraConnectionAttributes *string `field:"optional" json:"extraConnectionAttributes" yaml:"extraConnectionAttributes"`
	// Settings in JSON format for the source GCP MySQL endpoint.
	//
	// These settings are much the same as the settings for any MySQL-compatible endpoint. For more information, see [Extra connection attributes when using MySQL as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.MySQL.html#CHAP_Source.MySQL.ConnectionAttrib) in the *AWS Database Migration Service User Guide* .
	GcpMySqlSettings interface{} `field:"optional" json:"gcpMySqlSettings" yaml:"gcpMySqlSettings"`
	// Settings in JSON format for the source IBM Db2 LUW endpoint.
	//
	// For information about other available settings, see [Extra connection attributes when using Db2 LUW as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.DB2.html#CHAP_Source.DB2.ConnectionAttrib) in the *AWS Database Migration Service User Guide* .
	IbmDb2Settings interface{} `field:"optional" json:"ibmDb2Settings" yaml:"ibmDb2Settings"`
	// Settings in JSON format for the target Apache Kafka endpoint.
	//
	// For more information about other available settings, see [Using object mapping to migrate data to a Kafka topic](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Kafka.html#CHAP_Target.Kafka.ObjectMapping) in the *AWS Database Migration Service User Guide* .
	KafkaSettings interface{} `field:"optional" json:"kafkaSettings" yaml:"kafkaSettings"`
	// Settings in JSON format for the target endpoint for Amazon Kinesis Data Streams.
	//
	// For more information about other available settings, see [Using object mapping to migrate data to a Kinesis data stream](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Kinesis.html#CHAP_Target.Kinesis.ObjectMapping) in the *AWS Database Migration Service User Guide* .
	KinesisSettings interface{} `field:"optional" json:"kinesisSettings" yaml:"kinesisSettings"`
	// An AWS KMS key identifier that is used to encrypt the connection parameters for the endpoint.
	//
	// If you don't specify a value for the `KmsKeyId` parameter, AWS DMS uses your default encryption key.
	//
	// AWS KMS creates the default encryption key for your AWS account . Your AWS account has a different default encryption key for each AWS Region .
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Settings in JSON format for the source and target Microsoft SQL Server endpoint.
	//
	// For information about other available settings, see [Extra connection attributes when using SQL Server as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.SQLServer.html#CHAP_Source.SQLServer.ConnectionAttrib) and [Extra connection attributes when using SQL Server as a target for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.SQLServer.html#CHAP_Target.SQLServer.ConnectionAttrib) in the *AWS Database Migration Service User Guide* .
	MicrosoftSqlServerSettings interface{} `field:"optional" json:"microsoftSqlServerSettings" yaml:"microsoftSqlServerSettings"`
	// Settings in JSON format for the source MongoDB endpoint.
	//
	// For more information about the available settings, see [Using MongoDB as a target for AWS Database Migration Service](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.MongoDB.html#CHAP_Source.MongoDB.Configuration) in the *AWS Database Migration Service User Guide* .
	MongoDbSettings interface{} `field:"optional" json:"mongoDbSettings" yaml:"mongoDbSettings"`
	// Settings in JSON format for the source and target MySQL endpoint.
	//
	// For information about other available settings, see [Extra connection attributes when using MySQL as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.MySQL.html#CHAP_Source.MySQL.ConnectionAttrib) and [Extra connection attributes when using a MySQL-compatible database as a target for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.MySQL.html#CHAP_Target.MySQL.ConnectionAttrib) in the *AWS Database Migration Service User Guide* .
	MySqlSettings interface{} `field:"optional" json:"mySqlSettings" yaml:"mySqlSettings"`
	// Settings in JSON format for the target Amazon Neptune endpoint.
	//
	// For more information about the available settings, see [Specifying endpoint settings for Amazon Neptune as a target](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Neptune.html#CHAP_Target.Neptune.EndpointSettings) in the *AWS Database Migration Service User Guide* .
	NeptuneSettings interface{} `field:"optional" json:"neptuneSettings" yaml:"neptuneSettings"`
	// Settings in JSON format for the source and target Oracle endpoint.
	//
	// For information about other available settings, see [Extra connection attributes when using Oracle as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.Oracle.html#CHAP_Source.Oracle.ConnectionAttrib) and [Extra connection attributes when using Oracle as a target for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Oracle.html#CHAP_Target.Oracle.ConnectionAttrib) in the *AWS Database Migration Service User Guide* .
	OracleSettings interface{} `field:"optional" json:"oracleSettings" yaml:"oracleSettings"`
	// The password to be used to log in to the endpoint database.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The port used by the endpoint database.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Settings in JSON format for the source and target PostgreSQL endpoint.
	//
	// For information about other available settings, see [Extra connection attributes when using PostgreSQL as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.PostgreSQL.html#CHAP_Source.PostgreSQL.ConnectionAttrib) and [Extra connection attributes when using PostgreSQL as a target for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.PostgreSQL.html#CHAP_Target.PostgreSQL.ConnectionAttrib) in the *AWS Database Migration Service User Guide* .
	PostgreSqlSettings interface{} `field:"optional" json:"postgreSqlSettings" yaml:"postgreSqlSettings"`
	// Settings in JSON format for the target Redis endpoint.
	//
	// For information about other available settings, see [Specifying endpoint settings for Redis as a target](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Redis.html#CHAP_Target.Redis.EndpointSettings) in the *AWS Database Migration Service User Guide* .
	RedisSettings interface{} `field:"optional" json:"redisSettings" yaml:"redisSettings"`
	// Settings in JSON format for the Amazon Redshift endpoint.
	//
	// For more information about other available settings, see [Extra connection attributes when using Amazon Redshift as a target for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Redshift.html#CHAP_Target.Redshift.ConnectionAttrib) in the *AWS Database Migration Service User Guide* .
	RedshiftSettings interface{} `field:"optional" json:"redshiftSettings" yaml:"redshiftSettings"`
	// A display name for the resource identifier at the end of the `EndpointArn` response parameter that is returned in the created `Endpoint` object.
	//
	// The value for this parameter can have up to 31 characters. It can contain only ASCII letters, digits, and hyphen ('-'). Also, it can't end with a hyphen or contain two consecutive hyphens, and can only begin with a letter, such as `Example-App-ARN1` .
	//
	// For example, this value might result in the `EndpointArn` value `arn:aws:dms:eu-west-1:012345678901:rep:Example-App-ARN1` . If you don't specify a `ResourceIdentifier` value, AWS DMS generates a default identifier value for the end of `EndpointArn` .
	ResourceIdentifier *string `field:"optional" json:"resourceIdentifier" yaml:"resourceIdentifier"`
	// Settings in JSON format for the source and target Amazon S3 endpoint.
	//
	// For more information about other available settings, see [Extra connection attributes when using Amazon S3 as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.S3.html#CHAP_Source.S3.Configuring) and [Extra connection attributes when using Amazon S3 as a target for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.S3.html#CHAP_Target.S3.Configuring) in the *AWS Database Migration Service User Guide* .
	S3Settings interface{} `field:"optional" json:"s3Settings" yaml:"s3Settings"`
	// The name of the server where the endpoint database resides.
	ServerName *string `field:"optional" json:"serverName" yaml:"serverName"`
	// The Secure Sockets Layer (SSL) mode to use for the SSL connection. The default is `none` .
	//
	// > When `engine_name` is set to S3, the only allowed value is `none` .
	SslMode *string `field:"optional" json:"sslMode" yaml:"sslMode"`
	// Settings in JSON format for the source and target SAP ASE endpoint.
	//
	// For information about other available settings, see [Extra connection attributes when using SAP ASE as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.SAP.html#CHAP_Source.SAP.ConnectionAttrib) and [Extra connection attributes when using SAP ASE as a target for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.SAP.html#CHAP_Target.SAP.ConnectionAttrib) in the *AWS Database Migration Service User Guide* .
	SybaseSettings interface{} `field:"optional" json:"sybaseSettings" yaml:"sybaseSettings"`
	// One or more tags to be assigned to the endpoint.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The user name to be used to log in to the endpoint database.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

// A CloudFormation `AWS::DMS::EventSubscription`.
//
// Use the `AWS::DMS::EventSubscription` resource to get notifications for AWS Database Migration Service events through the Amazon Simple Notification Service . For more information, see [Working with events and notifications in AWS Database Migration Service](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Events.html) in the *AWS Database Migration Service User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEventSubscription := awscdk.Aws_dms.NewCfnEventSubscription(this, jsii.String("MyCfnEventSubscription"), &cfnEventSubscriptionProps{
//   	snsTopicArn: jsii.String("snsTopicArn"),
//
//   	// the properties below are optional
//   	enabled: jsii.Boolean(false),
//   	eventCategories: []*string{
//   		jsii.String("eventCategories"),
//   	},
//   	sourceIds: []*string{
//   		jsii.String("sourceIds"),
//   	},
//   	sourceType: jsii.String("sourceType"),
//   	subscriptionName: jsii.String("subscriptionName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnEventSubscription interface {
	awscdk.CfnResource
	awscdk.IInspectable
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
	// Indicates whether to activate the subscription.
	//
	// If you don't specify this property, AWS CloudFormation activates the subscription.
	Enabled() interface{}
	SetEnabled(val interface{})
	// A list of event categories for a source type that you want to subscribe to.
	//
	// If you don't specify this property, you are notified about all event categories. For more information, see [Working with Events and Notifications](https://docs.aws.amazon.com//dms/latest/userguide/CHAP_Events.html) in the *AWS DMS User Guide* .
	EventCategories() *[]*string
	SetEventCategories(val *[]*string)
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
	// The Amazon Resource Name (ARN) of the Amazon SNS topic created for event notification.
	//
	// The ARN is created by Amazon SNS when you create a topic and subscribe to it.
	SnsTopicArn() *string
	SetSnsTopicArn(val *string)
	// A list of identifiers for which AWS DMS provides notification events.
	//
	// If you don't specify a value, notifications are provided for all sources.
	//
	// If you specify multiple values, they must be of the same type. For example, if you specify a database instance ID, then all of the other values must be database instance IDs.
	SourceIds() *[]*string
	SetSourceIds(val *[]*string)
	// The type of AWS DMS resource that generates the events.
	//
	// For example, if you want to be notified of events generated by a replication instance, you set this parameter to `replication-instance` . If this value isn't specified, all events are returned.
	//
	// *Valid values* : `replication-instance` | `replication-task`.
	SourceType() *string
	SetSourceType(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// The name of the AWS DMS event notification subscription.
	//
	// This name must be less than 255 characters.
	SubscriptionName() *string
	SetSubscriptionName(val *string)
	// One or more tags to be assigned to the event subscription.
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

func (c *jsiiProxy_CfnEventSubscription) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnEventSubscription) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnEventSubscription) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnEventSubscription) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnEventSubscription) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnEventSubscription) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnEventSubscription) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

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

func (c *jsiiProxy_CfnEventSubscription) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnEventSubscription) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnEventSubscription) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_CfnEventSubscription) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

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

func (c *jsiiProxy_CfnEventSubscription) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_CfnEventSubscription) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnEventSubscription`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEventSubscriptionProps := &cfnEventSubscriptionProps{
//   	snsTopicArn: jsii.String("snsTopicArn"),
//
//   	// the properties below are optional
//   	enabled: jsii.Boolean(false),
//   	eventCategories: []*string{
//   		jsii.String("eventCategories"),
//   	},
//   	sourceIds: []*string{
//   		jsii.String("sourceIds"),
//   	},
//   	sourceType: jsii.String("sourceType"),
//   	subscriptionName: jsii.String("subscriptionName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnEventSubscriptionProps struct {
	// The Amazon Resource Name (ARN) of the Amazon SNS topic created for event notification.
	//
	// The ARN is created by Amazon SNS when you create a topic and subscribe to it.
	SnsTopicArn *string `field:"required" json:"snsTopicArn" yaml:"snsTopicArn"`
	// Indicates whether to activate the subscription.
	//
	// If you don't specify this property, AWS CloudFormation activates the subscription.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// A list of event categories for a source type that you want to subscribe to.
	//
	// If you don't specify this property, you are notified about all event categories. For more information, see [Working with Events and Notifications](https://docs.aws.amazon.com//dms/latest/userguide/CHAP_Events.html) in the *AWS DMS User Guide* .
	EventCategories *[]*string `field:"optional" json:"eventCategories" yaml:"eventCategories"`
	// A list of identifiers for which AWS DMS provides notification events.
	//
	// If you don't specify a value, notifications are provided for all sources.
	//
	// If you specify multiple values, they must be of the same type. For example, if you specify a database instance ID, then all of the other values must be database instance IDs.
	SourceIds *[]*string `field:"optional" json:"sourceIds" yaml:"sourceIds"`
	// The type of AWS DMS resource that generates the events.
	//
	// For example, if you want to be notified of events generated by a replication instance, you set this parameter to `replication-instance` . If this value isn't specified, all events are returned.
	//
	// *Valid values* : `replication-instance` | `replication-task`.
	SourceType *string `field:"optional" json:"sourceType" yaml:"sourceType"`
	// The name of the AWS DMS event notification subscription.
	//
	// This name must be less than 255 characters.
	SubscriptionName *string `field:"optional" json:"subscriptionName" yaml:"subscriptionName"`
	// One or more tags to be assigned to the event subscription.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::DMS::ReplicationInstance`.
//
// The `AWS::DMS::ReplicationInstance` resource creates an AWS DMS replication instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnReplicationInstance := awscdk.Aws_dms.NewCfnReplicationInstance(this, jsii.String("MyCfnReplicationInstance"), &cfnReplicationInstanceProps{
//   	replicationInstanceClass: jsii.String("replicationInstanceClass"),
//
//   	// the properties below are optional
//   	allocatedStorage: jsii.Number(123),
//   	allowMajorVersionUpgrade: jsii.Boolean(false),
//   	autoMinorVersionUpgrade: jsii.Boolean(false),
//   	availabilityZone: jsii.String("availabilityZone"),
//   	engineVersion: jsii.String("engineVersion"),
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	multiAz: jsii.Boolean(false),
//   	preferredMaintenanceWindow: jsii.String("preferredMaintenanceWindow"),
//   	publiclyAccessible: jsii.Boolean(false),
//   	replicationInstanceIdentifier: jsii.String("replicationInstanceIdentifier"),
//   	replicationSubnetGroupIdentifier: jsii.String("replicationSubnetGroupIdentifier"),
//   	resourceIdentifier: jsii.String("resourceIdentifier"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	vpcSecurityGroupIds: []*string{
//   		jsii.String("vpcSecurityGroupIds"),
//   	},
//   })
//
type CfnReplicationInstance interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The amount of storage (in gigabytes) to be initially allocated for the replication instance.
	AllocatedStorage() *float64
	SetAllocatedStorage(val *float64)
	// Indicates that major version upgrades are allowed.
	//
	// Changing this parameter does not result in an outage, and the change is asynchronously applied as soon as possible.
	//
	// This parameter must be set to `true` when specifying a value for the `EngineVersion` parameter that is a different major version than the replication instance's current version.
	AllowMajorVersionUpgrade() interface{}
	SetAllowMajorVersionUpgrade(val interface{})
	// One or more private IP addresses for the replication instance.
	AttrReplicationInstancePrivateIpAddresses() *string
	// One or more public IP addresses for the replication instance.
	AttrReplicationInstancePublicIpAddresses() *string
	// A value that indicates whether minor engine upgrades are applied automatically to the replication instance during the maintenance window.
	//
	// This parameter defaults to `true` .
	//
	// Default: `true`.
	AutoMinorVersionUpgrade() interface{}
	SetAutoMinorVersionUpgrade(val interface{})
	// The Availability Zone that the replication instance will be created in.
	//
	// The default value is a random, system-chosen Availability Zone in the endpoint's AWS Region , for example `us-east-1d` .
	AvailabilityZone() *string
	SetAvailabilityZone(val *string)
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
	// The engine version number of the replication instance.
	//
	// If an engine version number is not specified when a replication instance is created, the default is the latest engine version available.
	EngineVersion() *string
	SetEngineVersion(val *string)
	// An AWS KMS key identifier that is used to encrypt the data on the replication instance.
	//
	// If you don't specify a value for the `KmsKeyId` parameter, AWS DMS uses your default encryption key.
	//
	// AWS KMS creates the default encryption key for your AWS account . Your AWS account has a different default encryption key for each AWS Region .
	KmsKeyId() *string
	SetKmsKeyId(val *string)
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
	// Specifies whether the replication instance is a Multi-AZ deployment.
	//
	// You can't set the `AvailabilityZone` parameter if the Multi-AZ parameter is set to `true` .
	MultiAz() interface{}
	SetMultiAz(val interface{})
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// The weekly time range during which system maintenance can occur, in UTC.
	//
	// *Format* : `ddd:hh24:mi-ddd:hh24:mi`
	//
	// *Default* : A 30-minute window selected at random from an 8-hour block of time per AWS Region , occurring on a random day of the week.
	//
	// *Valid days* ( `ddd` ): `Mon` | `Tue` | `Wed` | `Thu` | `Fri` | `Sat` | `Sun`
	//
	// *Constraints* : Minimum 30-minute window.
	PreferredMaintenanceWindow() *string
	SetPreferredMaintenanceWindow(val *string)
	// Specifies the accessibility options for the replication instance.
	//
	// A value of `true` represents an instance with a public IP address. A value of `false` represents an instance with a private IP address. The default value is `true` .
	PubliclyAccessible() interface{}
	SetPubliclyAccessible(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The compute and memory capacity of the replication instance as defined for the specified replication instance class.
	//
	// For example, to specify the instance class dms.c4.large, set this parameter to `"dms.c4.large"` . For more information on the settings and capacities for the available replication instance classes, see [Selecting the right AWS DMS replication instance for your migration](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_ReplicationInstance.html#CHAP_ReplicationInstance.InDepth) in the *AWS Database Migration Service User Guide* .
	ReplicationInstanceClass() *string
	SetReplicationInstanceClass(val *string)
	// The replication instance identifier. This parameter is stored as a lowercase string.
	//
	// Constraints:
	//
	// - Must contain 1-63 alphanumeric characters or hyphens.
	// - First character must be a letter.
	// - Can't end with a hyphen or contain two consecutive hyphens.
	//
	// Example: `myrepinstance`.
	ReplicationInstanceIdentifier() *string
	SetReplicationInstanceIdentifier(val *string)
	// A subnet group to associate with the replication instance.
	ReplicationSubnetGroupIdentifier() *string
	SetReplicationSubnetGroupIdentifier(val *string)
	// A display name for the resource identifier at the end of the `EndpointArn` response parameter that is returned in the created `Endpoint` object.
	//
	// The value for this parameter can have up to 31 characters. It can contain only ASCII letters, digits, and hyphen ('-'). Also, it can't end with a hyphen or contain two consecutive hyphens, and can only begin with a letter, such as `Example-App-ARN1` . For example, this value might result in the `EndpointArn` value `arn:aws:dms:eu-west-1:012345678901:rep:Example-App-ARN1` . If you don't specify a `ResourceIdentifier` value, AWS DMS generates a default identifier value for the end of `EndpointArn` .
	ResourceIdentifier() *string
	SetResourceIdentifier(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// One or more tags to be assigned to the replication instance.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// Specifies the virtual private cloud (VPC) security group to be used with the replication instance.
	//
	// The VPC security group must work with the VPC containing the replication instance.
	VpcSecurityGroupIds() *[]*string
	SetVpcSecurityGroupIds(val *[]*string)
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

func (j *jsiiProxy_CfnReplicationInstance) AttrReplicationInstancePrivateIpAddresses() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrReplicationInstancePrivateIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationInstance) AttrReplicationInstancePublicIpAddresses() *string {
	var returns *string
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

func (c *jsiiProxy_CfnReplicationInstance) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnReplicationInstance) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnReplicationInstance) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnReplicationInstance) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnReplicationInstance) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnReplicationInstance) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnReplicationInstance) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

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

func (c *jsiiProxy_CfnReplicationInstance) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnReplicationInstance) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnReplicationInstance) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_CfnReplicationInstance) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

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

func (c *jsiiProxy_CfnReplicationInstance) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_CfnReplicationInstance) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnReplicationInstance`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnReplicationInstanceProps := &cfnReplicationInstanceProps{
//   	replicationInstanceClass: jsii.String("replicationInstanceClass"),
//
//   	// the properties below are optional
//   	allocatedStorage: jsii.Number(123),
//   	allowMajorVersionUpgrade: jsii.Boolean(false),
//   	autoMinorVersionUpgrade: jsii.Boolean(false),
//   	availabilityZone: jsii.String("availabilityZone"),
//   	engineVersion: jsii.String("engineVersion"),
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	multiAz: jsii.Boolean(false),
//   	preferredMaintenanceWindow: jsii.String("preferredMaintenanceWindow"),
//   	publiclyAccessible: jsii.Boolean(false),
//   	replicationInstanceIdentifier: jsii.String("replicationInstanceIdentifier"),
//   	replicationSubnetGroupIdentifier: jsii.String("replicationSubnetGroupIdentifier"),
//   	resourceIdentifier: jsii.String("resourceIdentifier"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	vpcSecurityGroupIds: []*string{
//   		jsii.String("vpcSecurityGroupIds"),
//   	},
//   }
//
type CfnReplicationInstanceProps struct {
	// The compute and memory capacity of the replication instance as defined for the specified replication instance class.
	//
	// For example, to specify the instance class dms.c4.large, set this parameter to `"dms.c4.large"` . For more information on the settings and capacities for the available replication instance classes, see [Selecting the right AWS DMS replication instance for your migration](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_ReplicationInstance.html#CHAP_ReplicationInstance.InDepth) in the *AWS Database Migration Service User Guide* .
	ReplicationInstanceClass *string `field:"required" json:"replicationInstanceClass" yaml:"replicationInstanceClass"`
	// The amount of storage (in gigabytes) to be initially allocated for the replication instance.
	AllocatedStorage *float64 `field:"optional" json:"allocatedStorage" yaml:"allocatedStorage"`
	// Indicates that major version upgrades are allowed.
	//
	// Changing this parameter does not result in an outage, and the change is asynchronously applied as soon as possible.
	//
	// This parameter must be set to `true` when specifying a value for the `EngineVersion` parameter that is a different major version than the replication instance's current version.
	AllowMajorVersionUpgrade interface{} `field:"optional" json:"allowMajorVersionUpgrade" yaml:"allowMajorVersionUpgrade"`
	// A value that indicates whether minor engine upgrades are applied automatically to the replication instance during the maintenance window.
	//
	// This parameter defaults to `true` .
	//
	// Default: `true`.
	AutoMinorVersionUpgrade interface{} `field:"optional" json:"autoMinorVersionUpgrade" yaml:"autoMinorVersionUpgrade"`
	// The Availability Zone that the replication instance will be created in.
	//
	// The default value is a random, system-chosen Availability Zone in the endpoint's AWS Region , for example `us-east-1d` .
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The engine version number of the replication instance.
	//
	// If an engine version number is not specified when a replication instance is created, the default is the latest engine version available.
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// An AWS KMS key identifier that is used to encrypt the data on the replication instance.
	//
	// If you don't specify a value for the `KmsKeyId` parameter, AWS DMS uses your default encryption key.
	//
	// AWS KMS creates the default encryption key for your AWS account . Your AWS account has a different default encryption key for each AWS Region .
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Specifies whether the replication instance is a Multi-AZ deployment.
	//
	// You can't set the `AvailabilityZone` parameter if the Multi-AZ parameter is set to `true` .
	MultiAz interface{} `field:"optional" json:"multiAz" yaml:"multiAz"`
	// The weekly time range during which system maintenance can occur, in UTC.
	//
	// *Format* : `ddd:hh24:mi-ddd:hh24:mi`
	//
	// *Default* : A 30-minute window selected at random from an 8-hour block of time per AWS Region , occurring on a random day of the week.
	//
	// *Valid days* ( `ddd` ): `Mon` | `Tue` | `Wed` | `Thu` | `Fri` | `Sat` | `Sun`
	//
	// *Constraints* : Minimum 30-minute window.
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// Specifies the accessibility options for the replication instance.
	//
	// A value of `true` represents an instance with a public IP address. A value of `false` represents an instance with a private IP address. The default value is `true` .
	PubliclyAccessible interface{} `field:"optional" json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// The replication instance identifier. This parameter is stored as a lowercase string.
	//
	// Constraints:
	//
	// - Must contain 1-63 alphanumeric characters or hyphens.
	// - First character must be a letter.
	// - Can't end with a hyphen or contain two consecutive hyphens.
	//
	// Example: `myrepinstance`.
	ReplicationInstanceIdentifier *string `field:"optional" json:"replicationInstanceIdentifier" yaml:"replicationInstanceIdentifier"`
	// A subnet group to associate with the replication instance.
	ReplicationSubnetGroupIdentifier *string `field:"optional" json:"replicationSubnetGroupIdentifier" yaml:"replicationSubnetGroupIdentifier"`
	// A display name for the resource identifier at the end of the `EndpointArn` response parameter that is returned in the created `Endpoint` object.
	//
	// The value for this parameter can have up to 31 characters. It can contain only ASCII letters, digits, and hyphen ('-'). Also, it can't end with a hyphen or contain two consecutive hyphens, and can only begin with a letter, such as `Example-App-ARN1` . For example, this value might result in the `EndpointArn` value `arn:aws:dms:eu-west-1:012345678901:rep:Example-App-ARN1` . If you don't specify a `ResourceIdentifier` value, AWS DMS generates a default identifier value for the end of `EndpointArn` .
	ResourceIdentifier *string `field:"optional" json:"resourceIdentifier" yaml:"resourceIdentifier"`
	// One or more tags to be assigned to the replication instance.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Specifies the virtual private cloud (VPC) security group to be used with the replication instance.
	//
	// The VPC security group must work with the VPC containing the replication instance.
	VpcSecurityGroupIds *[]*string `field:"optional" json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
}

// A CloudFormation `AWS::DMS::ReplicationSubnetGroup`.
//
// The `AWS::DMS::ReplicationSubnetGroup` resource creates an AWS DMS replication subnet group. Subnet groups must contain at least two subnets in two different Availability Zones in the same AWS Region .
//
// > Resource creation fails if the `dms-vpc-role` AWS Identity and Access Management ( IAM ) role doesn't already exist. For more information, see [Creating the IAM Roles to Use With the AWS CLI and AWS DMS API](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Security.APIRole.html) in the *AWS Database Migration Service User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnReplicationSubnetGroup := awscdk.Aws_dms.NewCfnReplicationSubnetGroup(this, jsii.String("MyCfnReplicationSubnetGroup"), &cfnReplicationSubnetGroupProps{
//   	replicationSubnetGroupDescription: jsii.String("replicationSubnetGroupDescription"),
//   	subnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//
//   	// the properties below are optional
//   	replicationSubnetGroupIdentifier: jsii.String("replicationSubnetGroupIdentifier"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnReplicationSubnetGroup interface {
	awscdk.CfnResource
	awscdk.IInspectable
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
	// The description for the subnet group.
	ReplicationSubnetGroupDescription() *string
	SetReplicationSubnetGroupDescription(val *string)
	// The identifier for the replication subnet group.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique ID and uses that ID for the identifier.
	ReplicationSubnetGroupIdentifier() *string
	SetReplicationSubnetGroupIdentifier(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// One or more subnet IDs to be assigned to the subnet group.
	SubnetIds() *[]*string
	SetSubnetIds(val *[]*string)
	// One or more tags to be assigned to the subnet group.
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

func (c *jsiiProxy_CfnReplicationSubnetGroup) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnReplicationSubnetGroup) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnReplicationSubnetGroup) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnReplicationSubnetGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnReplicationSubnetGroup) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnReplicationSubnetGroup) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnReplicationSubnetGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

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

func (c *jsiiProxy_CfnReplicationSubnetGroup) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnReplicationSubnetGroup) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnReplicationSubnetGroup) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_CfnReplicationSubnetGroup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

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

func (c *jsiiProxy_CfnReplicationSubnetGroup) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_CfnReplicationSubnetGroup) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnReplicationSubnetGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnReplicationSubnetGroupProps := &cfnReplicationSubnetGroupProps{
//   	replicationSubnetGroupDescription: jsii.String("replicationSubnetGroupDescription"),
//   	subnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//
//   	// the properties below are optional
//   	replicationSubnetGroupIdentifier: jsii.String("replicationSubnetGroupIdentifier"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnReplicationSubnetGroupProps struct {
	// The description for the subnet group.
	ReplicationSubnetGroupDescription *string `field:"required" json:"replicationSubnetGroupDescription" yaml:"replicationSubnetGroupDescription"`
	// One or more subnet IDs to be assigned to the subnet group.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// The identifier for the replication subnet group.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique ID and uses that ID for the identifier.
	ReplicationSubnetGroupIdentifier *string `field:"optional" json:"replicationSubnetGroupIdentifier" yaml:"replicationSubnetGroupIdentifier"`
	// One or more tags to be assigned to the subnet group.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::DMS::ReplicationTask`.
//
// The `AWS::DMS::ReplicationTask` resource creates an AWS DMS replication task.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnReplicationTask := awscdk.Aws_dms.NewCfnReplicationTask(this, jsii.String("MyCfnReplicationTask"), &cfnReplicationTaskProps{
//   	migrationType: jsii.String("migrationType"),
//   	replicationInstanceArn: jsii.String("replicationInstanceArn"),
//   	sourceEndpointArn: jsii.String("sourceEndpointArn"),
//   	tableMappings: jsii.String("tableMappings"),
//   	targetEndpointArn: jsii.String("targetEndpointArn"),
//
//   	// the properties below are optional
//   	cdcStartPosition: jsii.String("cdcStartPosition"),
//   	cdcStartTime: jsii.Number(123),
//   	cdcStopPosition: jsii.String("cdcStopPosition"),
//   	replicationTaskIdentifier: jsii.String("replicationTaskIdentifier"),
//   	replicationTaskSettings: jsii.String("replicationTaskSettings"),
//   	resourceIdentifier: jsii.String("resourceIdentifier"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	taskData: jsii.String("taskData"),
//   })
//
type CfnReplicationTask interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Indicates when you want a change data capture (CDC) operation to start.
	//
	// Use either `CdcStartPosition` or `CdcStartTime` to specify when you want a CDC operation to start. Specifying both values results in an error.
	//
	// The value can be in date, checkpoint, log sequence number (LSN), or system change number (SCN) format.
	//
	// Here is a date example: `--cdc-start-position "2018-03-08T12:12:12"`
	//
	// Here is a checkpoint example: `--cdc-start-position "checkpoint:V1#27#mysql-bin-changelog.157832:1975:-1:2002:677883278264080:mysql-bin-changelog.157832:1876#0#0#*#0#93"`
	//
	// Here is an LSN example: `--cdc-start-position “mysql-bin-changelog.000024:373”`
	//
	// > When you use this task setting with a source PostgreSQL database, a logical replication slot should already be created and associated with the source endpoint. You can verify this by setting the `slotName` extra connection attribute to the name of this logical replication slot. For more information, see [Extra Connection Attributes When Using PostgreSQL as a Source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.PostgreSQL.html#CHAP_Source.PostgreSQL.ConnectionAttrib) in the *AWS Database Migration Service User Guide* .
	CdcStartPosition() *string
	SetCdcStartPosition(val *string)
	// Indicates the start time for a change data capture (CDC) operation.
	CdcStartTime() *float64
	SetCdcStartTime(val *float64)
	// Indicates when you want a change data capture (CDC) operation to stop.
	//
	// The value can be either server time or commit time.
	//
	// Here is a server time example: `--cdc-stop-position "server_time:2018-02-09T12:12:12"`
	//
	// Here is a commit time example: `--cdc-stop-position "commit_time: 2018-02-09T12:12:12"`.
	CdcStopPosition() *string
	SetCdcStopPosition(val *string)
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
	// The migration type.
	//
	// Valid values: `full-load` | `cdc` | `full-load-and-cdc`.
	MigrationType() *string
	SetMigrationType(val *string)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The Amazon Resource Name (ARN) of a replication instance.
	ReplicationInstanceArn() *string
	SetReplicationInstanceArn(val *string)
	// An identifier for the replication task.
	//
	// Constraints:
	//
	// - Must contain 1-255 alphanumeric characters or hyphens.
	// - First character must be a letter.
	// - Cannot end with a hyphen or contain two consecutive hyphens.
	ReplicationTaskIdentifier() *string
	SetReplicationTaskIdentifier(val *string)
	// Overall settings for the task, in JSON format.
	//
	// For more information, see [Specifying Task Settings for AWS Database Migration Service Tasks](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TaskSettings.html) in the *AWS Database Migration Service User Guide* .
	ReplicationTaskSettings() *string
	SetReplicationTaskSettings(val *string)
	// A display name for the resource identifier at the end of the `EndpointArn` response parameter that is returned in the created `Endpoint` object.
	//
	// The value for this parameter can have up to 31 characters. It can contain only ASCII letters, digits, and hyphen ('-'). Also, it can't end with a hyphen or contain two consecutive hyphens, and can only begin with a letter, such as `Example-App-ARN1` .
	//
	// For example, this value might result in the `EndpointArn` value `arn:aws:dms:eu-west-1:012345678901:rep:Example-App-ARN1` . If you don't specify a `ResourceIdentifier` value, AWS DMS generates a default identifier value for the end of `EndpointArn` .
	ResourceIdentifier() *string
	SetResourceIdentifier(val *string)
	// An Amazon Resource Name (ARN) that uniquely identifies the source endpoint.
	SourceEndpointArn() *string
	SetSourceEndpointArn(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// The table mappings for the task, in JSON format.
	//
	// For more information, see [Using Table Mapping to Specify Task Settings](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TableMapping.html) in the *AWS Database Migration Service User Guide* .
	TableMappings() *string
	SetTableMappings(val *string)
	// One or more tags to be assigned to the replication task.
	Tags() awscdk.TagManager
	// An Amazon Resource Name (ARN) that uniquely identifies the target endpoint.
	TargetEndpointArn() *string
	SetTargetEndpointArn(val *string)
	// `AWS::DMS::ReplicationTask.TaskData`.
	TaskData() *string
	SetTaskData(val *string)
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

func (c *jsiiProxy_CfnReplicationTask) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnReplicationTask) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnReplicationTask) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnReplicationTask) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnReplicationTask) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnReplicationTask) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnReplicationTask) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

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

func (c *jsiiProxy_CfnReplicationTask) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnReplicationTask) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnReplicationTask) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_CfnReplicationTask) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

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

func (c *jsiiProxy_CfnReplicationTask) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_CfnReplicationTask) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnReplicationTask`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnReplicationTaskProps := &cfnReplicationTaskProps{
//   	migrationType: jsii.String("migrationType"),
//   	replicationInstanceArn: jsii.String("replicationInstanceArn"),
//   	sourceEndpointArn: jsii.String("sourceEndpointArn"),
//   	tableMappings: jsii.String("tableMappings"),
//   	targetEndpointArn: jsii.String("targetEndpointArn"),
//
//   	// the properties below are optional
//   	cdcStartPosition: jsii.String("cdcStartPosition"),
//   	cdcStartTime: jsii.Number(123),
//   	cdcStopPosition: jsii.String("cdcStopPosition"),
//   	replicationTaskIdentifier: jsii.String("replicationTaskIdentifier"),
//   	replicationTaskSettings: jsii.String("replicationTaskSettings"),
//   	resourceIdentifier: jsii.String("resourceIdentifier"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	taskData: jsii.String("taskData"),
//   }
//
type CfnReplicationTaskProps struct {
	// The migration type.
	//
	// Valid values: `full-load` | `cdc` | `full-load-and-cdc`.
	MigrationType *string `field:"required" json:"migrationType" yaml:"migrationType"`
	// The Amazon Resource Name (ARN) of a replication instance.
	ReplicationInstanceArn *string `field:"required" json:"replicationInstanceArn" yaml:"replicationInstanceArn"`
	// An Amazon Resource Name (ARN) that uniquely identifies the source endpoint.
	SourceEndpointArn *string `field:"required" json:"sourceEndpointArn" yaml:"sourceEndpointArn"`
	// The table mappings for the task, in JSON format.
	//
	// For more information, see [Using Table Mapping to Specify Task Settings](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TableMapping.html) in the *AWS Database Migration Service User Guide* .
	TableMappings *string `field:"required" json:"tableMappings" yaml:"tableMappings"`
	// An Amazon Resource Name (ARN) that uniquely identifies the target endpoint.
	TargetEndpointArn *string `field:"required" json:"targetEndpointArn" yaml:"targetEndpointArn"`
	// Indicates when you want a change data capture (CDC) operation to start.
	//
	// Use either `CdcStartPosition` or `CdcStartTime` to specify when you want a CDC operation to start. Specifying both values results in an error.
	//
	// The value can be in date, checkpoint, log sequence number (LSN), or system change number (SCN) format.
	//
	// Here is a date example: `--cdc-start-position "2018-03-08T12:12:12"`
	//
	// Here is a checkpoint example: `--cdc-start-position "checkpoint:V1#27#mysql-bin-changelog.157832:1975:-1:2002:677883278264080:mysql-bin-changelog.157832:1876#0#0#*#0#93"`
	//
	// Here is an LSN example: `--cdc-start-position “mysql-bin-changelog.000024:373”`
	//
	// > When you use this task setting with a source PostgreSQL database, a logical replication slot should already be created and associated with the source endpoint. You can verify this by setting the `slotName` extra connection attribute to the name of this logical replication slot. For more information, see [Extra Connection Attributes When Using PostgreSQL as a Source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.PostgreSQL.html#CHAP_Source.PostgreSQL.ConnectionAttrib) in the *AWS Database Migration Service User Guide* .
	CdcStartPosition *string `field:"optional" json:"cdcStartPosition" yaml:"cdcStartPosition"`
	// Indicates the start time for a change data capture (CDC) operation.
	CdcStartTime *float64 `field:"optional" json:"cdcStartTime" yaml:"cdcStartTime"`
	// Indicates when you want a change data capture (CDC) operation to stop.
	//
	// The value can be either server time or commit time.
	//
	// Here is a server time example: `--cdc-stop-position "server_time:2018-02-09T12:12:12"`
	//
	// Here is a commit time example: `--cdc-stop-position "commit_time: 2018-02-09T12:12:12"`.
	CdcStopPosition *string `field:"optional" json:"cdcStopPosition" yaml:"cdcStopPosition"`
	// An identifier for the replication task.
	//
	// Constraints:
	//
	// - Must contain 1-255 alphanumeric characters or hyphens.
	// - First character must be a letter.
	// - Cannot end with a hyphen or contain two consecutive hyphens.
	ReplicationTaskIdentifier *string `field:"optional" json:"replicationTaskIdentifier" yaml:"replicationTaskIdentifier"`
	// Overall settings for the task, in JSON format.
	//
	// For more information, see [Specifying Task Settings for AWS Database Migration Service Tasks](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TaskSettings.html) in the *AWS Database Migration Service User Guide* .
	ReplicationTaskSettings *string `field:"optional" json:"replicationTaskSettings" yaml:"replicationTaskSettings"`
	// A display name for the resource identifier at the end of the `EndpointArn` response parameter that is returned in the created `Endpoint` object.
	//
	// The value for this parameter can have up to 31 characters. It can contain only ASCII letters, digits, and hyphen ('-'). Also, it can't end with a hyphen or contain two consecutive hyphens, and can only begin with a letter, such as `Example-App-ARN1` .
	//
	// For example, this value might result in the `EndpointArn` value `arn:aws:dms:eu-west-1:012345678901:rep:Example-App-ARN1` . If you don't specify a `ResourceIdentifier` value, AWS DMS generates a default identifier value for the end of `EndpointArn` .
	ResourceIdentifier *string `field:"optional" json:"resourceIdentifier" yaml:"resourceIdentifier"`
	// One or more tags to be assigned to the replication task.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// `AWS::DMS::ReplicationTask.TaskData`.
	TaskData *string `field:"optional" json:"taskData" yaml:"taskData"`
}

