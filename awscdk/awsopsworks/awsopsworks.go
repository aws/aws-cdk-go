package awsopsworks

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsopsworks/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::OpsWorks::App`.
//
// TODO: EXAMPLE
//
type CfnApp interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AppSource() interface{}
	SetAppSource(val interface{})
	Attributes() interface{}
	SetAttributes(val interface{})
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	DataSources() interface{}
	SetDataSources(val interface{})
	Description() *string
	SetDescription(val *string)
	Domains() *[]*string
	SetDomains(val *[]*string)
	EnableSsl() interface{}
	SetEnableSsl(val interface{})
	Environment() interface{}
	SetEnvironment(val interface{})
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Ref() *string
	Shortname() *string
	SetShortname(val *string)
	SslConfiguration() interface{}
	SetSslConfiguration(val interface{})
	Stack() awscdk.Stack
	StackId() *string
	SetStackId(val *string)
	Type() *string
	SetType(val *string)
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

// The jsii proxy struct for CfnApp
type jsiiProxy_CfnApp struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnApp) AppSource() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"appSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) Attributes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"attributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) DataSources() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataSources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) Domains() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"domains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) EnableSsl() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSsl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) Environment() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) Shortname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shortname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) SslConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sslConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) StackId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::OpsWorks::App`.
func NewCfnApp(scope constructs.Construct, id *string, props *CfnAppProps) CfnApp {
	_init_.Initialize()

	j := jsiiProxy_CfnApp{}

	_jsii_.Create(
		"aws-cdk-lib.aws_opsworks.CfnApp",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::OpsWorks::App`.
func NewCfnApp_Override(c CfnApp, scope constructs.Construct, id *string, props *CfnAppProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_opsworks.CfnApp",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnApp) SetAppSource(val interface{}) {
	_jsii_.Set(
		j,
		"appSource",
		val,
	)
}

func (j *jsiiProxy_CfnApp) SetAttributes(val interface{}) {
	_jsii_.Set(
		j,
		"attributes",
		val,
	)
}

func (j *jsiiProxy_CfnApp) SetDataSources(val interface{}) {
	_jsii_.Set(
		j,
		"dataSources",
		val,
	)
}

func (j *jsiiProxy_CfnApp) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnApp) SetDomains(val *[]*string) {
	_jsii_.Set(
		j,
		"domains",
		val,
	)
}

func (j *jsiiProxy_CfnApp) SetEnableSsl(val interface{}) {
	_jsii_.Set(
		j,
		"enableSsl",
		val,
	)
}

func (j *jsiiProxy_CfnApp) SetEnvironment(val interface{}) {
	_jsii_.Set(
		j,
		"environment",
		val,
	)
}

func (j *jsiiProxy_CfnApp) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnApp) SetShortname(val *string) {
	_jsii_.Set(
		j,
		"shortname",
		val,
	)
}

func (j *jsiiProxy_CfnApp) SetSslConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"sslConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnApp) SetStackId(val *string) {
	_jsii_.Set(
		j,
		"stackId",
		val,
	)
}

func (j *jsiiProxy_CfnApp) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
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
func CfnApp_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_opsworks.CfnApp",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnApp_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_opsworks.CfnApp",
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
func CfnApp_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_opsworks.CfnApp",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnApp_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_opsworks.CfnApp",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnApp) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnApp) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnApp) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnApp) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnApp) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnApp) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnApp) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnApp) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnApp) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnApp) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnApp) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnApp) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnApp) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnApp) ToString() *string {
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
func (c *jsiiProxy_CfnApp) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// TODO: EXAMPLE
//
type CfnApp_DataSourceProperty struct {
	// `CfnApp.DataSourceProperty.Arn`.
	Arn *string `json:"arn"`
	// `CfnApp.DataSourceProperty.DatabaseName`.
	DatabaseName *string `json:"databaseName"`
	// `CfnApp.DataSourceProperty.Type`.
	Type *string `json:"type"`
}

// TODO: EXAMPLE
//
type CfnApp_EnvironmentVariableProperty struct {
	// `CfnApp.EnvironmentVariableProperty.Key`.
	Key *string `json:"key"`
	// `CfnApp.EnvironmentVariableProperty.Secure`.
	Secure interface{} `json:"secure"`
	// `CfnApp.EnvironmentVariableProperty.Value`.
	Value *string `json:"value"`
}

// TODO: EXAMPLE
//
type CfnApp_SourceProperty struct {
	// `CfnApp.SourceProperty.Password`.
	Password *string `json:"password"`
	// `CfnApp.SourceProperty.Revision`.
	Revision *string `json:"revision"`
	// `CfnApp.SourceProperty.SshKey`.
	SshKey *string `json:"sshKey"`
	// `CfnApp.SourceProperty.Type`.
	Type *string `json:"type"`
	// `CfnApp.SourceProperty.Url`.
	Url *string `json:"url"`
	// `CfnApp.SourceProperty.Username`.
	Username *string `json:"username"`
}

// TODO: EXAMPLE
//
type CfnApp_SslConfigurationProperty struct {
	// `CfnApp.SslConfigurationProperty.Certificate`.
	Certificate *string `json:"certificate"`
	// `CfnApp.SslConfigurationProperty.Chain`.
	Chain *string `json:"chain"`
	// `CfnApp.SslConfigurationProperty.PrivateKey`.
	PrivateKey *string `json:"privateKey"`
}

// Properties for defining a `AWS::OpsWorks::App`.
//
// TODO: EXAMPLE
//
type CfnAppProps struct {
	// `AWS::OpsWorks::App.AppSource`.
	AppSource interface{} `json:"appSource"`
	// `AWS::OpsWorks::App.Attributes`.
	Attributes interface{} `json:"attributes"`
	// `AWS::OpsWorks::App.DataSources`.
	DataSources interface{} `json:"dataSources"`
	// `AWS::OpsWorks::App.Description`.
	Description *string `json:"description"`
	// `AWS::OpsWorks::App.Domains`.
	Domains *[]*string `json:"domains"`
	// `AWS::OpsWorks::App.EnableSsl`.
	EnableSsl interface{} `json:"enableSsl"`
	// `AWS::OpsWorks::App.Environment`.
	Environment interface{} `json:"environment"`
	// `AWS::OpsWorks::App.Name`.
	Name *string `json:"name"`
	// `AWS::OpsWorks::App.Shortname`.
	Shortname *string `json:"shortname"`
	// `AWS::OpsWorks::App.SslConfiguration`.
	SslConfiguration interface{} `json:"sslConfiguration"`
	// `AWS::OpsWorks::App.StackId`.
	StackId *string `json:"stackId"`
	// `AWS::OpsWorks::App.Type`.
	Type *string `json:"type"`
}

// A CloudFormation `AWS::OpsWorks::ElasticLoadBalancerAttachment`.
//
// TODO: EXAMPLE
//
type CfnElasticLoadBalancerAttachment interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	ElasticLoadBalancerName() *string
	SetElasticLoadBalancerName(val *string)
	LayerId() *string
	SetLayerId(val *string)
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

// The jsii proxy struct for CfnElasticLoadBalancerAttachment
type jsiiProxy_CfnElasticLoadBalancerAttachment struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnElasticLoadBalancerAttachment) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnElasticLoadBalancerAttachment) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnElasticLoadBalancerAttachment) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnElasticLoadBalancerAttachment) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnElasticLoadBalancerAttachment) ElasticLoadBalancerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticLoadBalancerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnElasticLoadBalancerAttachment) LayerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"layerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnElasticLoadBalancerAttachment) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnElasticLoadBalancerAttachment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnElasticLoadBalancerAttachment) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnElasticLoadBalancerAttachment) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnElasticLoadBalancerAttachment) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::OpsWorks::ElasticLoadBalancerAttachment`.
func NewCfnElasticLoadBalancerAttachment(scope constructs.Construct, id *string, props *CfnElasticLoadBalancerAttachmentProps) CfnElasticLoadBalancerAttachment {
	_init_.Initialize()

	j := jsiiProxy_CfnElasticLoadBalancerAttachment{}

	_jsii_.Create(
		"aws-cdk-lib.aws_opsworks.CfnElasticLoadBalancerAttachment",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::OpsWorks::ElasticLoadBalancerAttachment`.
func NewCfnElasticLoadBalancerAttachment_Override(c CfnElasticLoadBalancerAttachment, scope constructs.Construct, id *string, props *CfnElasticLoadBalancerAttachmentProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_opsworks.CfnElasticLoadBalancerAttachment",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnElasticLoadBalancerAttachment) SetElasticLoadBalancerName(val *string) {
	_jsii_.Set(
		j,
		"elasticLoadBalancerName",
		val,
	)
}

func (j *jsiiProxy_CfnElasticLoadBalancerAttachment) SetLayerId(val *string) {
	_jsii_.Set(
		j,
		"layerId",
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
func CfnElasticLoadBalancerAttachment_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_opsworks.CfnElasticLoadBalancerAttachment",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnElasticLoadBalancerAttachment_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_opsworks.CfnElasticLoadBalancerAttachment",
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
func CfnElasticLoadBalancerAttachment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_opsworks.CfnElasticLoadBalancerAttachment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnElasticLoadBalancerAttachment_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_opsworks.CfnElasticLoadBalancerAttachment",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnElasticLoadBalancerAttachment) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnElasticLoadBalancerAttachment) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnElasticLoadBalancerAttachment) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnElasticLoadBalancerAttachment) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnElasticLoadBalancerAttachment) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnElasticLoadBalancerAttachment) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnElasticLoadBalancerAttachment) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnElasticLoadBalancerAttachment) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnElasticLoadBalancerAttachment) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnElasticLoadBalancerAttachment) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnElasticLoadBalancerAttachment) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnElasticLoadBalancerAttachment) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnElasticLoadBalancerAttachment) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnElasticLoadBalancerAttachment) ToString() *string {
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
func (c *jsiiProxy_CfnElasticLoadBalancerAttachment) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `AWS::OpsWorks::ElasticLoadBalancerAttachment`.
//
// TODO: EXAMPLE
//
type CfnElasticLoadBalancerAttachmentProps struct {
	// `AWS::OpsWorks::ElasticLoadBalancerAttachment.ElasticLoadBalancerName`.
	ElasticLoadBalancerName *string `json:"elasticLoadBalancerName"`
	// `AWS::OpsWorks::ElasticLoadBalancerAttachment.LayerId`.
	LayerId *string `json:"layerId"`
}

// A CloudFormation `AWS::OpsWorks::Instance`.
//
// TODO: EXAMPLE
//
type CfnInstance interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AgentVersion() *string
	SetAgentVersion(val *string)
	AmiId() *string
	SetAmiId(val *string)
	Architecture() *string
	SetArchitecture(val *string)
	AttrAvailabilityZone() *string
	AttrPrivateDnsName() *string
	AttrPrivateIp() *string
	AttrPublicDnsName() *string
	AttrPublicIp() *string
	AutoScalingType() *string
	SetAutoScalingType(val *string)
	AvailabilityZone() *string
	SetAvailabilityZone(val *string)
	BlockDeviceMappings() interface{}
	SetBlockDeviceMappings(val interface{})
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	EbsOptimized() interface{}
	SetEbsOptimized(val interface{})
	ElasticIps() *[]*string
	SetElasticIps(val *[]*string)
	Hostname() *string
	SetHostname(val *string)
	InstallUpdatesOnBoot() interface{}
	SetInstallUpdatesOnBoot(val interface{})
	InstanceType() *string
	SetInstanceType(val *string)
	LayerIds() *[]*string
	SetLayerIds(val *[]*string)
	LogicalId() *string
	Node() constructs.Node
	Os() *string
	SetOs(val *string)
	Ref() *string
	RootDeviceType() *string
	SetRootDeviceType(val *string)
	SshKeyName() *string
	SetSshKeyName(val *string)
	Stack() awscdk.Stack
	StackId() *string
	SetStackId(val *string)
	SubnetId() *string
	SetSubnetId(val *string)
	Tenancy() *string
	SetTenancy(val *string)
	TimeBasedAutoScaling() interface{}
	SetTimeBasedAutoScaling(val interface{})
	UpdatedProperites() *map[string]interface{}
	VirtualizationType() *string
	SetVirtualizationType(val *string)
	Volumes() *[]*string
	SetVolumes(val *[]*string)
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

// The jsii proxy struct for CfnInstance
type jsiiProxy_CfnInstance struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnInstance) AgentVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AmiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"amiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) Architecture() *string {
	var returns *string
	_jsii_.Get(
		j,
		"architecture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AttrAvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAvailabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AttrPrivateDnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrPrivateDnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AttrPrivateIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrPrivateIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AttrPublicDnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrPublicDnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AttrPublicIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrPublicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AutoScalingType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoScalingType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) BlockDeviceMappings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blockDeviceMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) EbsOptimized() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsOptimized",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) ElasticIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"elasticIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) InstallUpdatesOnBoot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"installUpdatesOnBoot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) LayerIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"layerIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) Os() *string {
	var returns *string
	_jsii_.Get(
		j,
		"os",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) RootDeviceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootDeviceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) SshKeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sshKeyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) StackId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) Tenancy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenancy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) TimeBasedAutoScaling() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeBasedAutoScaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) VirtualizationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualizationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) Volumes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"volumes",
		&returns,
	)
	return returns
}


// Create a new `AWS::OpsWorks::Instance`.
func NewCfnInstance(scope constructs.Construct, id *string, props *CfnInstanceProps) CfnInstance {
	_init_.Initialize()

	j := jsiiProxy_CfnInstance{}

	_jsii_.Create(
		"aws-cdk-lib.aws_opsworks.CfnInstance",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::OpsWorks::Instance`.
func NewCfnInstance_Override(c CfnInstance, scope constructs.Construct, id *string, props *CfnInstanceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_opsworks.CfnInstance",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnInstance) SetAgentVersion(val *string) {
	_jsii_.Set(
		j,
		"agentVersion",
		val,
	)
}

func (j *jsiiProxy_CfnInstance) SetAmiId(val *string) {
	_jsii_.Set(
		j,
		"amiId",
		val,
	)
}

func (j *jsiiProxy_CfnInstance) SetArchitecture(val *string) {
	_jsii_.Set(
		j,
		"architecture",
		val,
	)
}

func (j *jsiiProxy_CfnInstance) SetAutoScalingType(val *string) {
	_jsii_.Set(
		j,
		"autoScalingType",
		val,
	)
}

func (j *jsiiProxy_CfnInstance) SetAvailabilityZone(val *string) {
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_CfnInstance) SetBlockDeviceMappings(val interface{}) {
	_jsii_.Set(
		j,
		"blockDeviceMappings",
		val,
	)
}

func (j *jsiiProxy_CfnInstance) SetEbsOptimized(val interface{}) {
	_jsii_.Set(
		j,
		"ebsOptimized",
		val,
	)
}

func (j *jsiiProxy_CfnInstance) SetElasticIps(val *[]*string) {
	_jsii_.Set(
		j,
		"elasticIps",
		val,
	)
}

func (j *jsiiProxy_CfnInstance) SetHostname(val *string) {
	_jsii_.Set(
		j,
		"hostname",
		val,
	)
}

func (j *jsiiProxy_CfnInstance) SetInstallUpdatesOnBoot(val interface{}) {
	_jsii_.Set(
		j,
		"installUpdatesOnBoot",
		val,
	)
}

func (j *jsiiProxy_CfnInstance) SetInstanceType(val *string) {
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_CfnInstance) SetLayerIds(val *[]*string) {
	_jsii_.Set(
		j,
		"layerIds",
		val,
	)
}

func (j *jsiiProxy_CfnInstance) SetOs(val *string) {
	_jsii_.Set(
		j,
		"os",
		val,
	)
}

func (j *jsiiProxy_CfnInstance) SetRootDeviceType(val *string) {
	_jsii_.Set(
		j,
		"rootDeviceType",
		val,
	)
}

func (j *jsiiProxy_CfnInstance) SetSshKeyName(val *string) {
	_jsii_.Set(
		j,
		"sshKeyName",
		val,
	)
}

func (j *jsiiProxy_CfnInstance) SetStackId(val *string) {
	_jsii_.Set(
		j,
		"stackId",
		val,
	)
}

func (j *jsiiProxy_CfnInstance) SetSubnetId(val *string) {
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_CfnInstance) SetTenancy(val *string) {
	_jsii_.Set(
		j,
		"tenancy",
		val,
	)
}

func (j *jsiiProxy_CfnInstance) SetTimeBasedAutoScaling(val interface{}) {
	_jsii_.Set(
		j,
		"timeBasedAutoScaling",
		val,
	)
}

func (j *jsiiProxy_CfnInstance) SetVirtualizationType(val *string) {
	_jsii_.Set(
		j,
		"virtualizationType",
		val,
	)
}

func (j *jsiiProxy_CfnInstance) SetVolumes(val *[]*string) {
	_jsii_.Set(
		j,
		"volumes",
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
func CfnInstance_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_opsworks.CfnInstance",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnInstance_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_opsworks.CfnInstance",
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
func CfnInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_opsworks.CfnInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnInstance_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_opsworks.CfnInstance",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnInstance) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnInstance) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnInstance) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnInstance) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnInstance) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnInstance) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnInstance) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnInstance) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnInstance) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnInstance) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnInstance) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnInstance) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnInstance) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnInstance) ToString() *string {
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
func (c *jsiiProxy_CfnInstance) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// TODO: EXAMPLE
//
type CfnInstance_BlockDeviceMappingProperty struct {
	// `CfnInstance.BlockDeviceMappingProperty.DeviceName`.
	DeviceName *string `json:"deviceName"`
	// `CfnInstance.BlockDeviceMappingProperty.Ebs`.
	Ebs interface{} `json:"ebs"`
	// `CfnInstance.BlockDeviceMappingProperty.NoDevice`.
	NoDevice *string `json:"noDevice"`
	// `CfnInstance.BlockDeviceMappingProperty.VirtualName`.
	VirtualName *string `json:"virtualName"`
}

// TODO: EXAMPLE
//
type CfnInstance_EbsBlockDeviceProperty struct {
	// `CfnInstance.EbsBlockDeviceProperty.DeleteOnTermination`.
	DeleteOnTermination interface{} `json:"deleteOnTermination"`
	// `CfnInstance.EbsBlockDeviceProperty.Iops`.
	Iops *float64 `json:"iops"`
	// `CfnInstance.EbsBlockDeviceProperty.SnapshotId`.
	SnapshotId *string `json:"snapshotId"`
	// `CfnInstance.EbsBlockDeviceProperty.VolumeSize`.
	VolumeSize *float64 `json:"volumeSize"`
	// `CfnInstance.EbsBlockDeviceProperty.VolumeType`.
	VolumeType *string `json:"volumeType"`
}

// TODO: EXAMPLE
//
type CfnInstance_TimeBasedAutoScalingProperty struct {
	// `CfnInstance.TimeBasedAutoScalingProperty.Friday`.
	Friday interface{} `json:"friday"`
	// `CfnInstance.TimeBasedAutoScalingProperty.Monday`.
	Monday interface{} `json:"monday"`
	// `CfnInstance.TimeBasedAutoScalingProperty.Saturday`.
	Saturday interface{} `json:"saturday"`
	// `CfnInstance.TimeBasedAutoScalingProperty.Sunday`.
	Sunday interface{} `json:"sunday"`
	// `CfnInstance.TimeBasedAutoScalingProperty.Thursday`.
	Thursday interface{} `json:"thursday"`
	// `CfnInstance.TimeBasedAutoScalingProperty.Tuesday`.
	Tuesday interface{} `json:"tuesday"`
	// `CfnInstance.TimeBasedAutoScalingProperty.Wednesday`.
	Wednesday interface{} `json:"wednesday"`
}

// Properties for defining a `AWS::OpsWorks::Instance`.
//
// TODO: EXAMPLE
//
type CfnInstanceProps struct {
	// `AWS::OpsWorks::Instance.AgentVersion`.
	AgentVersion *string `json:"agentVersion"`
	// `AWS::OpsWorks::Instance.AmiId`.
	AmiId *string `json:"amiId"`
	// `AWS::OpsWorks::Instance.Architecture`.
	Architecture *string `json:"architecture"`
	// `AWS::OpsWorks::Instance.AutoScalingType`.
	AutoScalingType *string `json:"autoScalingType"`
	// `AWS::OpsWorks::Instance.AvailabilityZone`.
	AvailabilityZone *string `json:"availabilityZone"`
	// `AWS::OpsWorks::Instance.BlockDeviceMappings`.
	BlockDeviceMappings interface{} `json:"blockDeviceMappings"`
	// `AWS::OpsWorks::Instance.EbsOptimized`.
	EbsOptimized interface{} `json:"ebsOptimized"`
	// `AWS::OpsWorks::Instance.ElasticIps`.
	ElasticIps *[]*string `json:"elasticIps"`
	// `AWS::OpsWorks::Instance.Hostname`.
	Hostname *string `json:"hostname"`
	// `AWS::OpsWorks::Instance.InstallUpdatesOnBoot`.
	InstallUpdatesOnBoot interface{} `json:"installUpdatesOnBoot"`
	// `AWS::OpsWorks::Instance.InstanceType`.
	InstanceType *string `json:"instanceType"`
	// `AWS::OpsWorks::Instance.LayerIds`.
	LayerIds *[]*string `json:"layerIds"`
	// `AWS::OpsWorks::Instance.Os`.
	Os *string `json:"os"`
	// `AWS::OpsWorks::Instance.RootDeviceType`.
	RootDeviceType *string `json:"rootDeviceType"`
	// `AWS::OpsWorks::Instance.SshKeyName`.
	SshKeyName *string `json:"sshKeyName"`
	// `AWS::OpsWorks::Instance.StackId`.
	StackId *string `json:"stackId"`
	// `AWS::OpsWorks::Instance.SubnetId`.
	SubnetId *string `json:"subnetId"`
	// `AWS::OpsWorks::Instance.Tenancy`.
	Tenancy *string `json:"tenancy"`
	// `AWS::OpsWorks::Instance.TimeBasedAutoScaling`.
	TimeBasedAutoScaling interface{} `json:"timeBasedAutoScaling"`
	// `AWS::OpsWorks::Instance.VirtualizationType`.
	VirtualizationType *string `json:"virtualizationType"`
	// `AWS::OpsWorks::Instance.Volumes`.
	Volumes *[]*string `json:"volumes"`
}

// A CloudFormation `AWS::OpsWorks::Layer`.
//
// TODO: EXAMPLE
//
type CfnLayer interface {
	awscdk.CfnResource
	awscdk.IInspectable
	Attributes() interface{}
	SetAttributes(val interface{})
	AutoAssignElasticIps() interface{}
	SetAutoAssignElasticIps(val interface{})
	AutoAssignPublicIps() interface{}
	SetAutoAssignPublicIps(val interface{})
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	CustomInstanceProfileArn() *string
	SetCustomInstanceProfileArn(val *string)
	CustomJson() interface{}
	SetCustomJson(val interface{})
	CustomRecipes() interface{}
	SetCustomRecipes(val interface{})
	CustomSecurityGroupIds() *[]*string
	SetCustomSecurityGroupIds(val *[]*string)
	EnableAutoHealing() interface{}
	SetEnableAutoHealing(val interface{})
	InstallUpdatesOnBoot() interface{}
	SetInstallUpdatesOnBoot(val interface{})
	LifecycleEventConfiguration() interface{}
	SetLifecycleEventConfiguration(val interface{})
	LoadBasedAutoScaling() interface{}
	SetLoadBasedAutoScaling(val interface{})
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Packages() *[]*string
	SetPackages(val *[]*string)
	Ref() *string
	Shortname() *string
	SetShortname(val *string)
	Stack() awscdk.Stack
	StackId() *string
	SetStackId(val *string)
	Tags() awscdk.TagManager
	Type() *string
	SetType(val *string)
	UpdatedProperites() *map[string]interface{}
	UseEbsOptimizedInstances() interface{}
	SetUseEbsOptimizedInstances(val interface{})
	VolumeConfigurations() interface{}
	SetVolumeConfigurations(val interface{})
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

// The jsii proxy struct for CfnLayer
type jsiiProxy_CfnLayer struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnLayer) Attributes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"attributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayer) AutoAssignElasticIps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAssignElasticIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayer) AutoAssignPublicIps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAssignPublicIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayer) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayer) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayer) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayer) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayer) CustomInstanceProfileArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customInstanceProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayer) CustomJson() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayer) CustomRecipes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customRecipes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayer) CustomSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customSecurityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayer) EnableAutoHealing() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAutoHealing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayer) InstallUpdatesOnBoot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"installUpdatesOnBoot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayer) LifecycleEventConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lifecycleEventConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayer) LoadBasedAutoScaling() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loadBasedAutoScaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayer) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayer) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayer) Packages() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"packages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayer) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayer) Shortname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shortname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayer) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayer) StackId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayer) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayer) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayer) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayer) UseEbsOptimizedInstances() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useEbsOptimizedInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayer) VolumeConfigurations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"volumeConfigurations",
		&returns,
	)
	return returns
}


// Create a new `AWS::OpsWorks::Layer`.
func NewCfnLayer(scope constructs.Construct, id *string, props *CfnLayerProps) CfnLayer {
	_init_.Initialize()

	j := jsiiProxy_CfnLayer{}

	_jsii_.Create(
		"aws-cdk-lib.aws_opsworks.CfnLayer",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::OpsWorks::Layer`.
func NewCfnLayer_Override(c CfnLayer, scope constructs.Construct, id *string, props *CfnLayerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_opsworks.CfnLayer",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnLayer) SetAttributes(val interface{}) {
	_jsii_.Set(
		j,
		"attributes",
		val,
	)
}

func (j *jsiiProxy_CfnLayer) SetAutoAssignElasticIps(val interface{}) {
	_jsii_.Set(
		j,
		"autoAssignElasticIps",
		val,
	)
}

func (j *jsiiProxy_CfnLayer) SetAutoAssignPublicIps(val interface{}) {
	_jsii_.Set(
		j,
		"autoAssignPublicIps",
		val,
	)
}

func (j *jsiiProxy_CfnLayer) SetCustomInstanceProfileArn(val *string) {
	_jsii_.Set(
		j,
		"customInstanceProfileArn",
		val,
	)
}

func (j *jsiiProxy_CfnLayer) SetCustomJson(val interface{}) {
	_jsii_.Set(
		j,
		"customJson",
		val,
	)
}

func (j *jsiiProxy_CfnLayer) SetCustomRecipes(val interface{}) {
	_jsii_.Set(
		j,
		"customRecipes",
		val,
	)
}

func (j *jsiiProxy_CfnLayer) SetCustomSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"customSecurityGroupIds",
		val,
	)
}

func (j *jsiiProxy_CfnLayer) SetEnableAutoHealing(val interface{}) {
	_jsii_.Set(
		j,
		"enableAutoHealing",
		val,
	)
}

func (j *jsiiProxy_CfnLayer) SetInstallUpdatesOnBoot(val interface{}) {
	_jsii_.Set(
		j,
		"installUpdatesOnBoot",
		val,
	)
}

func (j *jsiiProxy_CfnLayer) SetLifecycleEventConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"lifecycleEventConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnLayer) SetLoadBasedAutoScaling(val interface{}) {
	_jsii_.Set(
		j,
		"loadBasedAutoScaling",
		val,
	)
}

func (j *jsiiProxy_CfnLayer) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnLayer) SetPackages(val *[]*string) {
	_jsii_.Set(
		j,
		"packages",
		val,
	)
}

func (j *jsiiProxy_CfnLayer) SetShortname(val *string) {
	_jsii_.Set(
		j,
		"shortname",
		val,
	)
}

func (j *jsiiProxy_CfnLayer) SetStackId(val *string) {
	_jsii_.Set(
		j,
		"stackId",
		val,
	)
}

func (j *jsiiProxy_CfnLayer) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_CfnLayer) SetUseEbsOptimizedInstances(val interface{}) {
	_jsii_.Set(
		j,
		"useEbsOptimizedInstances",
		val,
	)
}

func (j *jsiiProxy_CfnLayer) SetVolumeConfigurations(val interface{}) {
	_jsii_.Set(
		j,
		"volumeConfigurations",
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
func CfnLayer_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_opsworks.CfnLayer",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnLayer_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_opsworks.CfnLayer",
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
func CfnLayer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_opsworks.CfnLayer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLayer_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_opsworks.CfnLayer",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnLayer) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnLayer) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnLayer) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnLayer) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnLayer) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnLayer) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnLayer) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnLayer) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnLayer) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnLayer) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnLayer) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnLayer) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnLayer) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnLayer) ToString() *string {
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
func (c *jsiiProxy_CfnLayer) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// TODO: EXAMPLE
//
type CfnLayer_AutoScalingThresholdsProperty struct {
	// `CfnLayer.AutoScalingThresholdsProperty.CpuThreshold`.
	CpuThreshold *float64 `json:"cpuThreshold"`
	// `CfnLayer.AutoScalingThresholdsProperty.IgnoreMetricsTime`.
	IgnoreMetricsTime *float64 `json:"ignoreMetricsTime"`
	// `CfnLayer.AutoScalingThresholdsProperty.InstanceCount`.
	InstanceCount *float64 `json:"instanceCount"`
	// `CfnLayer.AutoScalingThresholdsProperty.LoadThreshold`.
	LoadThreshold *float64 `json:"loadThreshold"`
	// `CfnLayer.AutoScalingThresholdsProperty.MemoryThreshold`.
	MemoryThreshold *float64 `json:"memoryThreshold"`
	// `CfnLayer.AutoScalingThresholdsProperty.ThresholdsWaitTime`.
	ThresholdsWaitTime *float64 `json:"thresholdsWaitTime"`
}

// TODO: EXAMPLE
//
type CfnLayer_LifecycleEventConfigurationProperty struct {
	// `CfnLayer.LifecycleEventConfigurationProperty.ShutdownEventConfiguration`.
	ShutdownEventConfiguration interface{} `json:"shutdownEventConfiguration"`
}

// TODO: EXAMPLE
//
type CfnLayer_LoadBasedAutoScalingProperty struct {
	// `CfnLayer.LoadBasedAutoScalingProperty.DownScaling`.
	DownScaling interface{} `json:"downScaling"`
	// `CfnLayer.LoadBasedAutoScalingProperty.Enable`.
	Enable interface{} `json:"enable"`
	// `CfnLayer.LoadBasedAutoScalingProperty.UpScaling`.
	UpScaling interface{} `json:"upScaling"`
}

// TODO: EXAMPLE
//
type CfnLayer_RecipesProperty struct {
	// `CfnLayer.RecipesProperty.Configure`.
	Configure *[]*string `json:"configure"`
	// `CfnLayer.RecipesProperty.Deploy`.
	Deploy *[]*string `json:"deploy"`
	// `CfnLayer.RecipesProperty.Setup`.
	Setup *[]*string `json:"setup"`
	// `CfnLayer.RecipesProperty.Shutdown`.
	Shutdown *[]*string `json:"shutdown"`
	// `CfnLayer.RecipesProperty.Undeploy`.
	Undeploy *[]*string `json:"undeploy"`
}

// TODO: EXAMPLE
//
type CfnLayer_ShutdownEventConfigurationProperty struct {
	// `CfnLayer.ShutdownEventConfigurationProperty.DelayUntilElbConnectionsDrained`.
	DelayUntilElbConnectionsDrained interface{} `json:"delayUntilElbConnectionsDrained"`
	// `CfnLayer.ShutdownEventConfigurationProperty.ExecutionTimeout`.
	ExecutionTimeout *float64 `json:"executionTimeout"`
}

// TODO: EXAMPLE
//
type CfnLayer_VolumeConfigurationProperty struct {
	// `CfnLayer.VolumeConfigurationProperty.Encrypted`.
	Encrypted interface{} `json:"encrypted"`
	// `CfnLayer.VolumeConfigurationProperty.Iops`.
	Iops *float64 `json:"iops"`
	// `CfnLayer.VolumeConfigurationProperty.MountPoint`.
	MountPoint *string `json:"mountPoint"`
	// `CfnLayer.VolumeConfigurationProperty.NumberOfDisks`.
	NumberOfDisks *float64 `json:"numberOfDisks"`
	// `CfnLayer.VolumeConfigurationProperty.RaidLevel`.
	RaidLevel *float64 `json:"raidLevel"`
	// `CfnLayer.VolumeConfigurationProperty.Size`.
	Size *float64 `json:"size"`
	// `CfnLayer.VolumeConfigurationProperty.VolumeType`.
	VolumeType *string `json:"volumeType"`
}

// Properties for defining a `AWS::OpsWorks::Layer`.
//
// TODO: EXAMPLE
//
type CfnLayerProps struct {
	// `AWS::OpsWorks::Layer.Attributes`.
	Attributes interface{} `json:"attributes"`
	// `AWS::OpsWorks::Layer.AutoAssignElasticIps`.
	AutoAssignElasticIps interface{} `json:"autoAssignElasticIps"`
	// `AWS::OpsWorks::Layer.AutoAssignPublicIps`.
	AutoAssignPublicIps interface{} `json:"autoAssignPublicIps"`
	// `AWS::OpsWorks::Layer.CustomInstanceProfileArn`.
	CustomInstanceProfileArn *string `json:"customInstanceProfileArn"`
	// `AWS::OpsWorks::Layer.CustomJson`.
	CustomJson interface{} `json:"customJson"`
	// `AWS::OpsWorks::Layer.CustomRecipes`.
	CustomRecipes interface{} `json:"customRecipes"`
	// `AWS::OpsWorks::Layer.CustomSecurityGroupIds`.
	CustomSecurityGroupIds *[]*string `json:"customSecurityGroupIds"`
	// `AWS::OpsWorks::Layer.EnableAutoHealing`.
	EnableAutoHealing interface{} `json:"enableAutoHealing"`
	// `AWS::OpsWorks::Layer.InstallUpdatesOnBoot`.
	InstallUpdatesOnBoot interface{} `json:"installUpdatesOnBoot"`
	// `AWS::OpsWorks::Layer.LifecycleEventConfiguration`.
	LifecycleEventConfiguration interface{} `json:"lifecycleEventConfiguration"`
	// `AWS::OpsWorks::Layer.LoadBasedAutoScaling`.
	LoadBasedAutoScaling interface{} `json:"loadBasedAutoScaling"`
	// `AWS::OpsWorks::Layer.Name`.
	Name *string `json:"name"`
	// `AWS::OpsWorks::Layer.Packages`.
	Packages *[]*string `json:"packages"`
	// `AWS::OpsWorks::Layer.Shortname`.
	Shortname *string `json:"shortname"`
	// `AWS::OpsWorks::Layer.StackId`.
	StackId *string `json:"stackId"`
	// `AWS::OpsWorks::Layer.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
	// `AWS::OpsWorks::Layer.Type`.
	Type *string `json:"type"`
	// `AWS::OpsWorks::Layer.UseEbsOptimizedInstances`.
	UseEbsOptimizedInstances interface{} `json:"useEbsOptimizedInstances"`
	// `AWS::OpsWorks::Layer.VolumeConfigurations`.
	VolumeConfigurations interface{} `json:"volumeConfigurations"`
}

// A CloudFormation `AWS::OpsWorks::Stack`.
//
// TODO: EXAMPLE
//
type CfnStack interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AgentVersion() *string
	SetAgentVersion(val *string)
	Attributes() interface{}
	SetAttributes(val interface{})
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	ChefConfiguration() interface{}
	SetChefConfiguration(val interface{})
	CloneAppIds() *[]*string
	SetCloneAppIds(val *[]*string)
	ClonePermissions() interface{}
	SetClonePermissions(val interface{})
	ConfigurationManager() interface{}
	SetConfigurationManager(val interface{})
	CreationStack() *[]*string
	CustomCookbooksSource() interface{}
	SetCustomCookbooksSource(val interface{})
	CustomJson() interface{}
	SetCustomJson(val interface{})
	DefaultAvailabilityZone() *string
	SetDefaultAvailabilityZone(val *string)
	DefaultInstanceProfileArn() *string
	SetDefaultInstanceProfileArn(val *string)
	DefaultOs() *string
	SetDefaultOs(val *string)
	DefaultRootDeviceType() *string
	SetDefaultRootDeviceType(val *string)
	DefaultSshKeyName() *string
	SetDefaultSshKeyName(val *string)
	DefaultSubnetId() *string
	SetDefaultSubnetId(val *string)
	EcsClusterArn() *string
	SetEcsClusterArn(val *string)
	ElasticIps() interface{}
	SetElasticIps(val interface{})
	HostnameTheme() *string
	SetHostnameTheme(val *string)
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	RdsDbInstances() interface{}
	SetRdsDbInstances(val interface{})
	Ref() *string
	ServiceRoleArn() *string
	SetServiceRoleArn(val *string)
	SourceStackId() *string
	SetSourceStackId(val *string)
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	UseCustomCookbooks() interface{}
	SetUseCustomCookbooks(val interface{})
	UseOpsworksSecurityGroups() interface{}
	SetUseOpsworksSecurityGroups(val interface{})
	VpcId() *string
	SetVpcId(val *string)
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

// The jsii proxy struct for CfnStack
type jsiiProxy_CfnStack struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnStack) AgentVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) Attributes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"attributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) ChefConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"chefConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) CloneAppIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cloneAppIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) ClonePermissions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clonePermissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) ConfigurationManager() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"configurationManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) CustomCookbooksSource() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customCookbooksSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) CustomJson() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) DefaultAvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultAvailabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) DefaultInstanceProfileArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultInstanceProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) DefaultOs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultOs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) DefaultRootDeviceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRootDeviceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) DefaultSshKeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSshKeyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) DefaultSubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) EcsClusterArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ecsClusterArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) ElasticIps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"elasticIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) HostnameTheme() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnameTheme",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) RdsDbInstances() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rdsDbInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) ServiceRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) SourceStackId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceStackId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) UseCustomCookbooks() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useCustomCookbooks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) UseOpsworksSecurityGroups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useOpsworksSecurityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}


// Create a new `AWS::OpsWorks::Stack`.
func NewCfnStack(scope constructs.Construct, id *string, props *CfnStackProps) CfnStack {
	_init_.Initialize()

	j := jsiiProxy_CfnStack{}

	_jsii_.Create(
		"aws-cdk-lib.aws_opsworks.CfnStack",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::OpsWorks::Stack`.
func NewCfnStack_Override(c CfnStack, scope constructs.Construct, id *string, props *CfnStackProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_opsworks.CfnStack",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnStack) SetAgentVersion(val *string) {
	_jsii_.Set(
		j,
		"agentVersion",
		val,
	)
}

func (j *jsiiProxy_CfnStack) SetAttributes(val interface{}) {
	_jsii_.Set(
		j,
		"attributes",
		val,
	)
}

func (j *jsiiProxy_CfnStack) SetChefConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"chefConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnStack) SetCloneAppIds(val *[]*string) {
	_jsii_.Set(
		j,
		"cloneAppIds",
		val,
	)
}

func (j *jsiiProxy_CfnStack) SetClonePermissions(val interface{}) {
	_jsii_.Set(
		j,
		"clonePermissions",
		val,
	)
}

func (j *jsiiProxy_CfnStack) SetConfigurationManager(val interface{}) {
	_jsii_.Set(
		j,
		"configurationManager",
		val,
	)
}

func (j *jsiiProxy_CfnStack) SetCustomCookbooksSource(val interface{}) {
	_jsii_.Set(
		j,
		"customCookbooksSource",
		val,
	)
}

func (j *jsiiProxy_CfnStack) SetCustomJson(val interface{}) {
	_jsii_.Set(
		j,
		"customJson",
		val,
	)
}

func (j *jsiiProxy_CfnStack) SetDefaultAvailabilityZone(val *string) {
	_jsii_.Set(
		j,
		"defaultAvailabilityZone",
		val,
	)
}

func (j *jsiiProxy_CfnStack) SetDefaultInstanceProfileArn(val *string) {
	_jsii_.Set(
		j,
		"defaultInstanceProfileArn",
		val,
	)
}

func (j *jsiiProxy_CfnStack) SetDefaultOs(val *string) {
	_jsii_.Set(
		j,
		"defaultOs",
		val,
	)
}

func (j *jsiiProxy_CfnStack) SetDefaultRootDeviceType(val *string) {
	_jsii_.Set(
		j,
		"defaultRootDeviceType",
		val,
	)
}

func (j *jsiiProxy_CfnStack) SetDefaultSshKeyName(val *string) {
	_jsii_.Set(
		j,
		"defaultSshKeyName",
		val,
	)
}

func (j *jsiiProxy_CfnStack) SetDefaultSubnetId(val *string) {
	_jsii_.Set(
		j,
		"defaultSubnetId",
		val,
	)
}

func (j *jsiiProxy_CfnStack) SetEcsClusterArn(val *string) {
	_jsii_.Set(
		j,
		"ecsClusterArn",
		val,
	)
}

func (j *jsiiProxy_CfnStack) SetElasticIps(val interface{}) {
	_jsii_.Set(
		j,
		"elasticIps",
		val,
	)
}

func (j *jsiiProxy_CfnStack) SetHostnameTheme(val *string) {
	_jsii_.Set(
		j,
		"hostnameTheme",
		val,
	)
}

func (j *jsiiProxy_CfnStack) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnStack) SetRdsDbInstances(val interface{}) {
	_jsii_.Set(
		j,
		"rdsDbInstances",
		val,
	)
}

func (j *jsiiProxy_CfnStack) SetServiceRoleArn(val *string) {
	_jsii_.Set(
		j,
		"serviceRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnStack) SetSourceStackId(val *string) {
	_jsii_.Set(
		j,
		"sourceStackId",
		val,
	)
}

func (j *jsiiProxy_CfnStack) SetUseCustomCookbooks(val interface{}) {
	_jsii_.Set(
		j,
		"useCustomCookbooks",
		val,
	)
}

func (j *jsiiProxy_CfnStack) SetUseOpsworksSecurityGroups(val interface{}) {
	_jsii_.Set(
		j,
		"useOpsworksSecurityGroups",
		val,
	)
}

func (j *jsiiProxy_CfnStack) SetVpcId(val *string) {
	_jsii_.Set(
		j,
		"vpcId",
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
func CfnStack_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_opsworks.CfnStack",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnStack_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_opsworks.CfnStack",
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
func CfnStack_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_opsworks.CfnStack",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnStack_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_opsworks.CfnStack",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnStack) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnStack) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnStack) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnStack) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnStack) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnStack) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnStack) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnStack) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnStack) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnStack) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnStack) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnStack) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnStack) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnStack) ToString() *string {
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
func (c *jsiiProxy_CfnStack) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// TODO: EXAMPLE
//
type CfnStack_ChefConfigurationProperty struct {
	// `CfnStack.ChefConfigurationProperty.BerkshelfVersion`.
	BerkshelfVersion *string `json:"berkshelfVersion"`
	// `CfnStack.ChefConfigurationProperty.ManageBerkshelf`.
	ManageBerkshelf interface{} `json:"manageBerkshelf"`
}

// TODO: EXAMPLE
//
type CfnStack_ElasticIpProperty struct {
	// `CfnStack.ElasticIpProperty.Ip`.
	Ip *string `json:"ip"`
	// `CfnStack.ElasticIpProperty.Name`.
	Name *string `json:"name"`
}

// TODO: EXAMPLE
//
type CfnStack_RdsDbInstanceProperty struct {
	// `CfnStack.RdsDbInstanceProperty.DbPassword`.
	DbPassword *string `json:"dbPassword"`
	// `CfnStack.RdsDbInstanceProperty.DbUser`.
	DbUser *string `json:"dbUser"`
	// `CfnStack.RdsDbInstanceProperty.RdsDbInstanceArn`.
	RdsDbInstanceArn *string `json:"rdsDbInstanceArn"`
}

// TODO: EXAMPLE
//
type CfnStack_SourceProperty struct {
	// `CfnStack.SourceProperty.Password`.
	Password *string `json:"password"`
	// `CfnStack.SourceProperty.Revision`.
	Revision *string `json:"revision"`
	// `CfnStack.SourceProperty.SshKey`.
	SshKey *string `json:"sshKey"`
	// `CfnStack.SourceProperty.Type`.
	Type *string `json:"type"`
	// `CfnStack.SourceProperty.Url`.
	Url *string `json:"url"`
	// `CfnStack.SourceProperty.Username`.
	Username *string `json:"username"`
}

// TODO: EXAMPLE
//
type CfnStack_StackConfigurationManagerProperty struct {
	// `CfnStack.StackConfigurationManagerProperty.Name`.
	Name *string `json:"name"`
	// `CfnStack.StackConfigurationManagerProperty.Version`.
	Version *string `json:"version"`
}

// Properties for defining a `AWS::OpsWorks::Stack`.
//
// TODO: EXAMPLE
//
type CfnStackProps struct {
	// `AWS::OpsWorks::Stack.AgentVersion`.
	AgentVersion *string `json:"agentVersion"`
	// `AWS::OpsWorks::Stack.Attributes`.
	Attributes interface{} `json:"attributes"`
	// `AWS::OpsWorks::Stack.ChefConfiguration`.
	ChefConfiguration interface{} `json:"chefConfiguration"`
	// `AWS::OpsWorks::Stack.CloneAppIds`.
	CloneAppIds *[]*string `json:"cloneAppIds"`
	// `AWS::OpsWorks::Stack.ClonePermissions`.
	ClonePermissions interface{} `json:"clonePermissions"`
	// `AWS::OpsWorks::Stack.ConfigurationManager`.
	ConfigurationManager interface{} `json:"configurationManager"`
	// `AWS::OpsWorks::Stack.CustomCookbooksSource`.
	CustomCookbooksSource interface{} `json:"customCookbooksSource"`
	// `AWS::OpsWorks::Stack.CustomJson`.
	CustomJson interface{} `json:"customJson"`
	// `AWS::OpsWorks::Stack.DefaultAvailabilityZone`.
	DefaultAvailabilityZone *string `json:"defaultAvailabilityZone"`
	// `AWS::OpsWorks::Stack.DefaultInstanceProfileArn`.
	DefaultInstanceProfileArn *string `json:"defaultInstanceProfileArn"`
	// `AWS::OpsWorks::Stack.DefaultOs`.
	DefaultOs *string `json:"defaultOs"`
	// `AWS::OpsWorks::Stack.DefaultRootDeviceType`.
	DefaultRootDeviceType *string `json:"defaultRootDeviceType"`
	// `AWS::OpsWorks::Stack.DefaultSshKeyName`.
	DefaultSshKeyName *string `json:"defaultSshKeyName"`
	// `AWS::OpsWorks::Stack.DefaultSubnetId`.
	DefaultSubnetId *string `json:"defaultSubnetId"`
	// `AWS::OpsWorks::Stack.EcsClusterArn`.
	EcsClusterArn *string `json:"ecsClusterArn"`
	// `AWS::OpsWorks::Stack.ElasticIps`.
	ElasticIps interface{} `json:"elasticIps"`
	// `AWS::OpsWorks::Stack.HostnameTheme`.
	HostnameTheme *string `json:"hostnameTheme"`
	// `AWS::OpsWorks::Stack.Name`.
	Name *string `json:"name"`
	// `AWS::OpsWorks::Stack.RdsDbInstances`.
	RdsDbInstances interface{} `json:"rdsDbInstances"`
	// `AWS::OpsWorks::Stack.ServiceRoleArn`.
	ServiceRoleArn *string `json:"serviceRoleArn"`
	// `AWS::OpsWorks::Stack.SourceStackId`.
	SourceStackId *string `json:"sourceStackId"`
	// `AWS::OpsWorks::Stack.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
	// `AWS::OpsWorks::Stack.UseCustomCookbooks`.
	UseCustomCookbooks interface{} `json:"useCustomCookbooks"`
	// `AWS::OpsWorks::Stack.UseOpsworksSecurityGroups`.
	UseOpsworksSecurityGroups interface{} `json:"useOpsworksSecurityGroups"`
	// `AWS::OpsWorks::Stack.VpcId`.
	VpcId *string `json:"vpcId"`
}

// A CloudFormation `AWS::OpsWorks::UserProfile`.
//
// TODO: EXAMPLE
//
type CfnUserProfile interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AllowSelfManagement() interface{}
	SetAllowSelfManagement(val interface{})
	AttrSshUsername() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	IamUserArn() *string
	SetIamUserArn(val *string)
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	SshPublicKey() *string
	SetSshPublicKey(val *string)
	SshUsername() *string
	SetSshUsername(val *string)
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

// The jsii proxy struct for CfnUserProfile
type jsiiProxy_CfnUserProfile struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnUserProfile) AllowSelfManagement() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowSelfManagement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserProfile) AttrSshUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSshUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserProfile) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserProfile) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserProfile) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserProfile) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserProfile) IamUserArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamUserArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserProfile) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserProfile) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserProfile) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserProfile) SshPublicKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sshPublicKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserProfile) SshUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sshUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserProfile) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserProfile) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::OpsWorks::UserProfile`.
func NewCfnUserProfile(scope constructs.Construct, id *string, props *CfnUserProfileProps) CfnUserProfile {
	_init_.Initialize()

	j := jsiiProxy_CfnUserProfile{}

	_jsii_.Create(
		"aws-cdk-lib.aws_opsworks.CfnUserProfile",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::OpsWorks::UserProfile`.
func NewCfnUserProfile_Override(c CfnUserProfile, scope constructs.Construct, id *string, props *CfnUserProfileProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_opsworks.CfnUserProfile",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnUserProfile) SetAllowSelfManagement(val interface{}) {
	_jsii_.Set(
		j,
		"allowSelfManagement",
		val,
	)
}

func (j *jsiiProxy_CfnUserProfile) SetIamUserArn(val *string) {
	_jsii_.Set(
		j,
		"iamUserArn",
		val,
	)
}

func (j *jsiiProxy_CfnUserProfile) SetSshPublicKey(val *string) {
	_jsii_.Set(
		j,
		"sshPublicKey",
		val,
	)
}

func (j *jsiiProxy_CfnUserProfile) SetSshUsername(val *string) {
	_jsii_.Set(
		j,
		"sshUsername",
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
func CfnUserProfile_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_opsworks.CfnUserProfile",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnUserProfile_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_opsworks.CfnUserProfile",
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
func CfnUserProfile_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_opsworks.CfnUserProfile",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnUserProfile_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_opsworks.CfnUserProfile",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnUserProfile) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnUserProfile) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnUserProfile) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnUserProfile) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnUserProfile) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnUserProfile) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnUserProfile) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnUserProfile) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnUserProfile) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnUserProfile) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnUserProfile) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnUserProfile) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnUserProfile) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnUserProfile) ToString() *string {
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
func (c *jsiiProxy_CfnUserProfile) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `AWS::OpsWorks::UserProfile`.
//
// TODO: EXAMPLE
//
type CfnUserProfileProps struct {
	// `AWS::OpsWorks::UserProfile.AllowSelfManagement`.
	AllowSelfManagement interface{} `json:"allowSelfManagement"`
	// `AWS::OpsWorks::UserProfile.IamUserArn`.
	IamUserArn *string `json:"iamUserArn"`
	// `AWS::OpsWorks::UserProfile.SshPublicKey`.
	SshPublicKey *string `json:"sshPublicKey"`
	// `AWS::OpsWorks::UserProfile.SshUsername`.
	SshUsername *string `json:"sshUsername"`
}

// A CloudFormation `AWS::OpsWorks::Volume`.
//
// TODO: EXAMPLE
//
type CfnVolume interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	Ec2VolumeId() *string
	SetEc2VolumeId(val *string)
	LogicalId() *string
	MountPoint() *string
	SetMountPoint(val *string)
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Ref() *string
	Stack() awscdk.Stack
	StackId() *string
	SetStackId(val *string)
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

// The jsii proxy struct for CfnVolume
type jsiiProxy_CfnVolume struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnVolume) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVolume) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVolume) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVolume) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVolume) Ec2VolumeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ec2VolumeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVolume) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVolume) MountPoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mountPoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVolume) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVolume) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVolume) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVolume) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVolume) StackId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVolume) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::OpsWorks::Volume`.
func NewCfnVolume(scope constructs.Construct, id *string, props *CfnVolumeProps) CfnVolume {
	_init_.Initialize()

	j := jsiiProxy_CfnVolume{}

	_jsii_.Create(
		"aws-cdk-lib.aws_opsworks.CfnVolume",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::OpsWorks::Volume`.
func NewCfnVolume_Override(c CfnVolume, scope constructs.Construct, id *string, props *CfnVolumeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_opsworks.CfnVolume",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnVolume) SetEc2VolumeId(val *string) {
	_jsii_.Set(
		j,
		"ec2VolumeId",
		val,
	)
}

func (j *jsiiProxy_CfnVolume) SetMountPoint(val *string) {
	_jsii_.Set(
		j,
		"mountPoint",
		val,
	)
}

func (j *jsiiProxy_CfnVolume) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnVolume) SetStackId(val *string) {
	_jsii_.Set(
		j,
		"stackId",
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
func CfnVolume_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_opsworks.CfnVolume",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnVolume_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_opsworks.CfnVolume",
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
func CfnVolume_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_opsworks.CfnVolume",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnVolume_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_opsworks.CfnVolume",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnVolume) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnVolume) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnVolume) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnVolume) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnVolume) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnVolume) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnVolume) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnVolume) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnVolume) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnVolume) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnVolume) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnVolume) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnVolume) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnVolume) ToString() *string {
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
func (c *jsiiProxy_CfnVolume) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `AWS::OpsWorks::Volume`.
//
// TODO: EXAMPLE
//
type CfnVolumeProps struct {
	// `AWS::OpsWorks::Volume.Ec2VolumeId`.
	Ec2VolumeId *string `json:"ec2VolumeId"`
	// `AWS::OpsWorks::Volume.MountPoint`.
	MountPoint *string `json:"mountPoint"`
	// `AWS::OpsWorks::Volume.Name`.
	Name *string `json:"name"`
	// `AWS::OpsWorks::Volume.StackId`.
	StackId *string `json:"stackId"`
}

