package awsappflow

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsappflow/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::AppFlow::ConnectorProfile`.
type CfnConnectorProfile interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrConnectorProfileArn() *string
	AttrCredentialsArn() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	ConnectionMode() *string
	SetConnectionMode(val *string)
	ConnectorProfileConfig() interface{}
	SetConnectorProfileConfig(val interface{})
	ConnectorProfileName() *string
	SetConnectorProfileName(val *string)
	ConnectorType() *string
	SetConnectorType(val *string)
	CreationStack() *[]*string
	KmsArn() *string
	SetKmsArn(val *string)
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

// The jsii proxy struct for CfnConnectorProfile
type jsiiProxy_CfnConnectorProfile struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnConnectorProfile) AttrConnectorProfileArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrConnectorProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectorProfile) AttrCredentialsArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCredentialsArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectorProfile) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectorProfile) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectorProfile) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectorProfile) ConnectionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectorProfile) ConnectorProfileConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectorProfileConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectorProfile) ConnectorProfileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectorProfileName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectorProfile) ConnectorType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectorType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectorProfile) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectorProfile) KmsArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectorProfile) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectorProfile) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectorProfile) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectorProfile) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectorProfile) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::AppFlow::ConnectorProfile`.
func NewCfnConnectorProfile(scope awscdk.Construct, id *string, props *CfnConnectorProfileProps) CfnConnectorProfile {
	_init_.Initialize()

	j := jsiiProxy_CfnConnectorProfile{}

	_jsii_.Create(
		"monocdk.aws_appflow.CfnConnectorProfile",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AppFlow::ConnectorProfile`.
func NewCfnConnectorProfile_Override(c CfnConnectorProfile, scope awscdk.Construct, id *string, props *CfnConnectorProfileProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appflow.CfnConnectorProfile",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnConnectorProfile) SetConnectionMode(val *string) {
	_jsii_.Set(
		j,
		"connectionMode",
		val,
	)
}

func (j *jsiiProxy_CfnConnectorProfile) SetConnectorProfileConfig(val interface{}) {
	_jsii_.Set(
		j,
		"connectorProfileConfig",
		val,
	)
}

func (j *jsiiProxy_CfnConnectorProfile) SetConnectorProfileName(val *string) {
	_jsii_.Set(
		j,
		"connectorProfileName",
		val,
	)
}

func (j *jsiiProxy_CfnConnectorProfile) SetConnectorType(val *string) {
	_jsii_.Set(
		j,
		"connectorType",
		val,
	)
}

func (j *jsiiProxy_CfnConnectorProfile) SetKmsArn(val *string) {
	_jsii_.Set(
		j,
		"kmsArn",
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
func CfnConnectorProfile_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appflow.CfnConnectorProfile",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnConnectorProfile_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appflow.CfnConnectorProfile",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnConnectorProfile_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appflow.CfnConnectorProfile",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnConnectorProfile_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_appflow.CfnConnectorProfile",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnConnectorProfile) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnConnectorProfile) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnConnectorProfile) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnConnectorProfile) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnConnectorProfile) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnConnectorProfile) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnConnectorProfile) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnConnectorProfile) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnConnectorProfile) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnConnectorProfile) Inspect(inspector awscdk.TreeInspector) {
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
func (c *jsiiProxy_CfnConnectorProfile) OnPrepare() {
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
func (c *jsiiProxy_CfnConnectorProfile) OnSynthesize(session constructs.ISynthesisSession) {
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
func (c *jsiiProxy_CfnConnectorProfile) OnValidate() *[]*string {
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
func (c *jsiiProxy_CfnConnectorProfile) OverrideLogicalId(newLogicalId *string) {
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
func (c *jsiiProxy_CfnConnectorProfile) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnConnectorProfile) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnConnectorProfile) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnConnectorProfile) Synthesize(session awscdk.ISynthesisSession) {
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
func (c *jsiiProxy_CfnConnectorProfile) ToString() *string {
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
func (c *jsiiProxy_CfnConnectorProfile) Validate() *[]*string {
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
func (c *jsiiProxy_CfnConnectorProfile) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnConnectorProfile_AmplitudeConnectorProfileCredentialsProperty struct {
	// `CfnConnectorProfile.AmplitudeConnectorProfileCredentialsProperty.ApiKey`.
	ApiKey *string `json:"apiKey"`
	// `CfnConnectorProfile.AmplitudeConnectorProfileCredentialsProperty.SecretKey`.
	SecretKey *string `json:"secretKey"`
}

type CfnConnectorProfile_ConnectorOAuthRequestProperty struct {
	// `CfnConnectorProfile.ConnectorOAuthRequestProperty.AuthCode`.
	AuthCode *string `json:"authCode"`
	// `CfnConnectorProfile.ConnectorOAuthRequestProperty.RedirectUri`.
	RedirectUri *string `json:"redirectUri"`
}

type CfnConnectorProfile_ConnectorProfileConfigProperty struct {
	// `CfnConnectorProfile.ConnectorProfileConfigProperty.ConnectorProfileCredentials`.
	ConnectorProfileCredentials interface{} `json:"connectorProfileCredentials"`
	// `CfnConnectorProfile.ConnectorProfileConfigProperty.ConnectorProfileProperties`.
	ConnectorProfileProperties interface{} `json:"connectorProfileProperties"`
}

type CfnConnectorProfile_ConnectorProfileCredentialsProperty struct {
	// `CfnConnectorProfile.ConnectorProfileCredentialsProperty.Amplitude`.
	Amplitude interface{} `json:"amplitude"`
	// `CfnConnectorProfile.ConnectorProfileCredentialsProperty.Datadog`.
	Datadog interface{} `json:"datadog"`
	// `CfnConnectorProfile.ConnectorProfileCredentialsProperty.Dynatrace`.
	Dynatrace interface{} `json:"dynatrace"`
	// `CfnConnectorProfile.ConnectorProfileCredentialsProperty.GoogleAnalytics`.
	GoogleAnalytics interface{} `json:"googleAnalytics"`
	// `CfnConnectorProfile.ConnectorProfileCredentialsProperty.InforNexus`.
	InforNexus interface{} `json:"inforNexus"`
	// `CfnConnectorProfile.ConnectorProfileCredentialsProperty.Marketo`.
	Marketo interface{} `json:"marketo"`
	// `CfnConnectorProfile.ConnectorProfileCredentialsProperty.Redshift`.
	Redshift interface{} `json:"redshift"`
	// `CfnConnectorProfile.ConnectorProfileCredentialsProperty.Salesforce`.
	Salesforce interface{} `json:"salesforce"`
	// `CfnConnectorProfile.ConnectorProfileCredentialsProperty.SAPOData`.
	SapoData interface{} `json:"sapoData"`
	// `CfnConnectorProfile.ConnectorProfileCredentialsProperty.ServiceNow`.
	ServiceNow interface{} `json:"serviceNow"`
	// `CfnConnectorProfile.ConnectorProfileCredentialsProperty.Singular`.
	Singular interface{} `json:"singular"`
	// `CfnConnectorProfile.ConnectorProfileCredentialsProperty.Slack`.
	Slack interface{} `json:"slack"`
	// `CfnConnectorProfile.ConnectorProfileCredentialsProperty.Snowflake`.
	Snowflake interface{} `json:"snowflake"`
	// `CfnConnectorProfile.ConnectorProfileCredentialsProperty.Trendmicro`.
	Trendmicro interface{} `json:"trendmicro"`
	// `CfnConnectorProfile.ConnectorProfileCredentialsProperty.Veeva`.
	Veeva interface{} `json:"veeva"`
	// `CfnConnectorProfile.ConnectorProfileCredentialsProperty.Zendesk`.
	Zendesk interface{} `json:"zendesk"`
}

type CfnConnectorProfile_ConnectorProfilePropertiesProperty struct {
	// `CfnConnectorProfile.ConnectorProfilePropertiesProperty.Datadog`.
	Datadog interface{} `json:"datadog"`
	// `CfnConnectorProfile.ConnectorProfilePropertiesProperty.Dynatrace`.
	Dynatrace interface{} `json:"dynatrace"`
	// `CfnConnectorProfile.ConnectorProfilePropertiesProperty.InforNexus`.
	InforNexus interface{} `json:"inforNexus"`
	// `CfnConnectorProfile.ConnectorProfilePropertiesProperty.Marketo`.
	Marketo interface{} `json:"marketo"`
	// `CfnConnectorProfile.ConnectorProfilePropertiesProperty.Redshift`.
	Redshift interface{} `json:"redshift"`
	// `CfnConnectorProfile.ConnectorProfilePropertiesProperty.Salesforce`.
	Salesforce interface{} `json:"salesforce"`
	// `CfnConnectorProfile.ConnectorProfilePropertiesProperty.SAPOData`.
	SapoData interface{} `json:"sapoData"`
	// `CfnConnectorProfile.ConnectorProfilePropertiesProperty.ServiceNow`.
	ServiceNow interface{} `json:"serviceNow"`
	// `CfnConnectorProfile.ConnectorProfilePropertiesProperty.Slack`.
	Slack interface{} `json:"slack"`
	// `CfnConnectorProfile.ConnectorProfilePropertiesProperty.Snowflake`.
	Snowflake interface{} `json:"snowflake"`
	// `CfnConnectorProfile.ConnectorProfilePropertiesProperty.Veeva`.
	Veeva interface{} `json:"veeva"`
	// `CfnConnectorProfile.ConnectorProfilePropertiesProperty.Zendesk`.
	Zendesk interface{} `json:"zendesk"`
}

type CfnConnectorProfile_DatadogConnectorProfileCredentialsProperty struct {
	// `CfnConnectorProfile.DatadogConnectorProfileCredentialsProperty.ApiKey`.
	ApiKey *string `json:"apiKey"`
	// `CfnConnectorProfile.DatadogConnectorProfileCredentialsProperty.ApplicationKey`.
	ApplicationKey *string `json:"applicationKey"`
}

type CfnConnectorProfile_DatadogConnectorProfilePropertiesProperty struct {
	// `CfnConnectorProfile.DatadogConnectorProfilePropertiesProperty.InstanceUrl`.
	InstanceUrl *string `json:"instanceUrl"`
}

type CfnConnectorProfile_DynatraceConnectorProfileCredentialsProperty struct {
	// `CfnConnectorProfile.DynatraceConnectorProfileCredentialsProperty.ApiToken`.
	ApiToken *string `json:"apiToken"`
}

type CfnConnectorProfile_DynatraceConnectorProfilePropertiesProperty struct {
	// `CfnConnectorProfile.DynatraceConnectorProfilePropertiesProperty.InstanceUrl`.
	InstanceUrl *string `json:"instanceUrl"`
}

type CfnConnectorProfile_GoogleAnalyticsConnectorProfileCredentialsProperty struct {
	// `CfnConnectorProfile.GoogleAnalyticsConnectorProfileCredentialsProperty.ClientId`.
	ClientId *string `json:"clientId"`
	// `CfnConnectorProfile.GoogleAnalyticsConnectorProfileCredentialsProperty.ClientSecret`.
	ClientSecret *string `json:"clientSecret"`
	// `CfnConnectorProfile.GoogleAnalyticsConnectorProfileCredentialsProperty.AccessToken`.
	AccessToken *string `json:"accessToken"`
	// `CfnConnectorProfile.GoogleAnalyticsConnectorProfileCredentialsProperty.ConnectorOAuthRequest`.
	ConnectorOAuthRequest interface{} `json:"connectorOAuthRequest"`
	// `CfnConnectorProfile.GoogleAnalyticsConnectorProfileCredentialsProperty.RefreshToken`.
	RefreshToken *string `json:"refreshToken"`
}

type CfnConnectorProfile_InforNexusConnectorProfileCredentialsProperty struct {
	// `CfnConnectorProfile.InforNexusConnectorProfileCredentialsProperty.AccessKeyId`.
	AccessKeyId *string `json:"accessKeyId"`
	// `CfnConnectorProfile.InforNexusConnectorProfileCredentialsProperty.Datakey`.
	Datakey *string `json:"datakey"`
	// `CfnConnectorProfile.InforNexusConnectorProfileCredentialsProperty.SecretAccessKey`.
	SecretAccessKey *string `json:"secretAccessKey"`
	// `CfnConnectorProfile.InforNexusConnectorProfileCredentialsProperty.UserId`.
	UserId *string `json:"userId"`
}

type CfnConnectorProfile_InforNexusConnectorProfilePropertiesProperty struct {
	// `CfnConnectorProfile.InforNexusConnectorProfilePropertiesProperty.InstanceUrl`.
	InstanceUrl *string `json:"instanceUrl"`
}

type CfnConnectorProfile_MarketoConnectorProfileCredentialsProperty struct {
	// `CfnConnectorProfile.MarketoConnectorProfileCredentialsProperty.ClientId`.
	ClientId *string `json:"clientId"`
	// `CfnConnectorProfile.MarketoConnectorProfileCredentialsProperty.ClientSecret`.
	ClientSecret *string `json:"clientSecret"`
	// `CfnConnectorProfile.MarketoConnectorProfileCredentialsProperty.AccessToken`.
	AccessToken *string `json:"accessToken"`
	// `CfnConnectorProfile.MarketoConnectorProfileCredentialsProperty.ConnectorOAuthRequest`.
	ConnectorOAuthRequest interface{} `json:"connectorOAuthRequest"`
}

type CfnConnectorProfile_MarketoConnectorProfilePropertiesProperty struct {
	// `CfnConnectorProfile.MarketoConnectorProfilePropertiesProperty.InstanceUrl`.
	InstanceUrl *string `json:"instanceUrl"`
}

type CfnConnectorProfile_OAuthPropertiesProperty struct {
	// `CfnConnectorProfile.OAuthPropertiesProperty.AuthCodeUrl`.
	AuthCodeUrl *string `json:"authCodeUrl"`
	// `CfnConnectorProfile.OAuthPropertiesProperty.OAuthScopes`.
	OAuthScopes *[]*string `json:"oAuthScopes"`
	// `CfnConnectorProfile.OAuthPropertiesProperty.TokenUrl`.
	TokenUrl *string `json:"tokenUrl"`
}

type CfnConnectorProfile_RedshiftConnectorProfileCredentialsProperty struct {
	// `CfnConnectorProfile.RedshiftConnectorProfileCredentialsProperty.Password`.
	Password *string `json:"password"`
	// `CfnConnectorProfile.RedshiftConnectorProfileCredentialsProperty.Username`.
	Username *string `json:"username"`
}

type CfnConnectorProfile_RedshiftConnectorProfilePropertiesProperty struct {
	// `CfnConnectorProfile.RedshiftConnectorProfilePropertiesProperty.BucketName`.
	BucketName *string `json:"bucketName"`
	// `CfnConnectorProfile.RedshiftConnectorProfilePropertiesProperty.DatabaseUrl`.
	DatabaseUrl *string `json:"databaseUrl"`
	// `CfnConnectorProfile.RedshiftConnectorProfilePropertiesProperty.RoleArn`.
	RoleArn *string `json:"roleArn"`
	// `CfnConnectorProfile.RedshiftConnectorProfilePropertiesProperty.BucketPrefix`.
	BucketPrefix *string `json:"bucketPrefix"`
}

type CfnConnectorProfile_SAPODataConnectorProfileCredentialsProperty struct {
	// `CfnConnectorProfile.SAPODataConnectorProfileCredentialsProperty.BasicAuthCredentials`.
	BasicAuthCredentials interface{} `json:"basicAuthCredentials"`
	// `CfnConnectorProfile.SAPODataConnectorProfileCredentialsProperty.OAuthCredentials`.
	OAuthCredentials interface{} `json:"oAuthCredentials"`
}

type CfnConnectorProfile_SAPODataConnectorProfilePropertiesProperty struct {
	// `CfnConnectorProfile.SAPODataConnectorProfilePropertiesProperty.ApplicationHostUrl`.
	ApplicationHostUrl *string `json:"applicationHostUrl"`
	// `CfnConnectorProfile.SAPODataConnectorProfilePropertiesProperty.ApplicationServicePath`.
	ApplicationServicePath *string `json:"applicationServicePath"`
	// `CfnConnectorProfile.SAPODataConnectorProfilePropertiesProperty.ClientNumber`.
	ClientNumber *string `json:"clientNumber"`
	// `CfnConnectorProfile.SAPODataConnectorProfilePropertiesProperty.LogonLanguage`.
	LogonLanguage *string `json:"logonLanguage"`
	// `CfnConnectorProfile.SAPODataConnectorProfilePropertiesProperty.OAuthProperties`.
	OAuthProperties interface{} `json:"oAuthProperties"`
	// `CfnConnectorProfile.SAPODataConnectorProfilePropertiesProperty.PortNumber`.
	PortNumber *float64 `json:"portNumber"`
	// `CfnConnectorProfile.SAPODataConnectorProfilePropertiesProperty.PrivateLinkServiceName`.
	PrivateLinkServiceName *string `json:"privateLinkServiceName"`
}

type CfnConnectorProfile_SalesforceConnectorProfileCredentialsProperty struct {
	// `CfnConnectorProfile.SalesforceConnectorProfileCredentialsProperty.AccessToken`.
	AccessToken *string `json:"accessToken"`
	// `CfnConnectorProfile.SalesforceConnectorProfileCredentialsProperty.ClientCredentialsArn`.
	ClientCredentialsArn *string `json:"clientCredentialsArn"`
	// `CfnConnectorProfile.SalesforceConnectorProfileCredentialsProperty.ConnectorOAuthRequest`.
	ConnectorOAuthRequest interface{} `json:"connectorOAuthRequest"`
	// `CfnConnectorProfile.SalesforceConnectorProfileCredentialsProperty.RefreshToken`.
	RefreshToken *string `json:"refreshToken"`
}

type CfnConnectorProfile_SalesforceConnectorProfilePropertiesProperty struct {
	// `CfnConnectorProfile.SalesforceConnectorProfilePropertiesProperty.InstanceUrl`.
	InstanceUrl *string `json:"instanceUrl"`
	// `CfnConnectorProfile.SalesforceConnectorProfilePropertiesProperty.isSandboxEnvironment`.
	IsSandboxEnvironment interface{} `json:"isSandboxEnvironment"`
}

type CfnConnectorProfile_ServiceNowConnectorProfileCredentialsProperty struct {
	// `CfnConnectorProfile.ServiceNowConnectorProfileCredentialsProperty.Password`.
	Password *string `json:"password"`
	// `CfnConnectorProfile.ServiceNowConnectorProfileCredentialsProperty.Username`.
	Username *string `json:"username"`
}

type CfnConnectorProfile_ServiceNowConnectorProfilePropertiesProperty struct {
	// `CfnConnectorProfile.ServiceNowConnectorProfilePropertiesProperty.InstanceUrl`.
	InstanceUrl *string `json:"instanceUrl"`
}

type CfnConnectorProfile_SingularConnectorProfileCredentialsProperty struct {
	// `CfnConnectorProfile.SingularConnectorProfileCredentialsProperty.ApiKey`.
	ApiKey *string `json:"apiKey"`
}

type CfnConnectorProfile_SlackConnectorProfileCredentialsProperty struct {
	// `CfnConnectorProfile.SlackConnectorProfileCredentialsProperty.ClientId`.
	ClientId *string `json:"clientId"`
	// `CfnConnectorProfile.SlackConnectorProfileCredentialsProperty.ClientSecret`.
	ClientSecret *string `json:"clientSecret"`
	// `CfnConnectorProfile.SlackConnectorProfileCredentialsProperty.AccessToken`.
	AccessToken *string `json:"accessToken"`
	// `CfnConnectorProfile.SlackConnectorProfileCredentialsProperty.ConnectorOAuthRequest`.
	ConnectorOAuthRequest interface{} `json:"connectorOAuthRequest"`
}

type CfnConnectorProfile_SlackConnectorProfilePropertiesProperty struct {
	// `CfnConnectorProfile.SlackConnectorProfilePropertiesProperty.InstanceUrl`.
	InstanceUrl *string `json:"instanceUrl"`
}

type CfnConnectorProfile_SnowflakeConnectorProfileCredentialsProperty struct {
	// `CfnConnectorProfile.SnowflakeConnectorProfileCredentialsProperty.Password`.
	Password *string `json:"password"`
	// `CfnConnectorProfile.SnowflakeConnectorProfileCredentialsProperty.Username`.
	Username *string `json:"username"`
}

type CfnConnectorProfile_SnowflakeConnectorProfilePropertiesProperty struct {
	// `CfnConnectorProfile.SnowflakeConnectorProfilePropertiesProperty.BucketName`.
	BucketName *string `json:"bucketName"`
	// `CfnConnectorProfile.SnowflakeConnectorProfilePropertiesProperty.Stage`.
	Stage *string `json:"stage"`
	// `CfnConnectorProfile.SnowflakeConnectorProfilePropertiesProperty.Warehouse`.
	Warehouse *string `json:"warehouse"`
	// `CfnConnectorProfile.SnowflakeConnectorProfilePropertiesProperty.AccountName`.
	AccountName *string `json:"accountName"`
	// `CfnConnectorProfile.SnowflakeConnectorProfilePropertiesProperty.BucketPrefix`.
	BucketPrefix *string `json:"bucketPrefix"`
	// `CfnConnectorProfile.SnowflakeConnectorProfilePropertiesProperty.PrivateLinkServiceName`.
	PrivateLinkServiceName *string `json:"privateLinkServiceName"`
	// `CfnConnectorProfile.SnowflakeConnectorProfilePropertiesProperty.Region`.
	Region *string `json:"region"`
}

type CfnConnectorProfile_TrendmicroConnectorProfileCredentialsProperty struct {
	// `CfnConnectorProfile.TrendmicroConnectorProfileCredentialsProperty.ApiSecretKey`.
	ApiSecretKey *string `json:"apiSecretKey"`
}

type CfnConnectorProfile_VeevaConnectorProfileCredentialsProperty struct {
	// `CfnConnectorProfile.VeevaConnectorProfileCredentialsProperty.Password`.
	Password *string `json:"password"`
	// `CfnConnectorProfile.VeevaConnectorProfileCredentialsProperty.Username`.
	Username *string `json:"username"`
}

type CfnConnectorProfile_VeevaConnectorProfilePropertiesProperty struct {
	// `CfnConnectorProfile.VeevaConnectorProfilePropertiesProperty.InstanceUrl`.
	InstanceUrl *string `json:"instanceUrl"`
}

type CfnConnectorProfile_ZendeskConnectorProfileCredentialsProperty struct {
	// `CfnConnectorProfile.ZendeskConnectorProfileCredentialsProperty.ClientId`.
	ClientId *string `json:"clientId"`
	// `CfnConnectorProfile.ZendeskConnectorProfileCredentialsProperty.ClientSecret`.
	ClientSecret *string `json:"clientSecret"`
	// `CfnConnectorProfile.ZendeskConnectorProfileCredentialsProperty.AccessToken`.
	AccessToken *string `json:"accessToken"`
	// `CfnConnectorProfile.ZendeskConnectorProfileCredentialsProperty.ConnectorOAuthRequest`.
	ConnectorOAuthRequest interface{} `json:"connectorOAuthRequest"`
}

type CfnConnectorProfile_ZendeskConnectorProfilePropertiesProperty struct {
	// `CfnConnectorProfile.ZendeskConnectorProfilePropertiesProperty.InstanceUrl`.
	InstanceUrl *string `json:"instanceUrl"`
}

// Properties for defining a `AWS::AppFlow::ConnectorProfile`.
type CfnConnectorProfileProps struct {
	// `AWS::AppFlow::ConnectorProfile.ConnectionMode`.
	ConnectionMode *string `json:"connectionMode"`
	// `AWS::AppFlow::ConnectorProfile.ConnectorProfileName`.
	ConnectorProfileName *string `json:"connectorProfileName"`
	// `AWS::AppFlow::ConnectorProfile.ConnectorType`.
	ConnectorType *string `json:"connectorType"`
	// `AWS::AppFlow::ConnectorProfile.ConnectorProfileConfig`.
	ConnectorProfileConfig interface{} `json:"connectorProfileConfig"`
	// `AWS::AppFlow::ConnectorProfile.KMSArn`.
	KmsArn *string `json:"kmsArn"`
}

// A CloudFormation `AWS::AppFlow::Flow`.
type CfnFlow interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrFlowArn() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	Description() *string
	SetDescription(val *string)
	DestinationFlowConfigList() interface{}
	SetDestinationFlowConfigList(val interface{})
	FlowName() *string
	SetFlowName(val *string)
	KmsArn() *string
	SetKmsArn(val *string)
	LogicalId() *string
	Node() awscdk.ConstructNode
	Ref() *string
	SourceFlowConfig() interface{}
	SetSourceFlowConfig(val interface{})
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	Tasks() interface{}
	SetTasks(val interface{})
	TriggerConfig() interface{}
	SetTriggerConfig(val interface{})
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

// The jsii proxy struct for CfnFlow
type jsiiProxy_CfnFlow struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnFlow) AttrFlowArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrFlowArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) DestinationFlowConfigList() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"destinationFlowConfigList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) FlowName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flowName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) KmsArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) SourceFlowConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceFlowConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) Tasks() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tasks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) TriggerConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"triggerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::AppFlow::Flow`.
func NewCfnFlow(scope awscdk.Construct, id *string, props *CfnFlowProps) CfnFlow {
	_init_.Initialize()

	j := jsiiProxy_CfnFlow{}

	_jsii_.Create(
		"monocdk.aws_appflow.CfnFlow",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AppFlow::Flow`.
func NewCfnFlow_Override(c CfnFlow, scope awscdk.Construct, id *string, props *CfnFlowProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appflow.CfnFlow",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnFlow) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnFlow) SetDestinationFlowConfigList(val interface{}) {
	_jsii_.Set(
		j,
		"destinationFlowConfigList",
		val,
	)
}

func (j *jsiiProxy_CfnFlow) SetFlowName(val *string) {
	_jsii_.Set(
		j,
		"flowName",
		val,
	)
}

func (j *jsiiProxy_CfnFlow) SetKmsArn(val *string) {
	_jsii_.Set(
		j,
		"kmsArn",
		val,
	)
}

func (j *jsiiProxy_CfnFlow) SetSourceFlowConfig(val interface{}) {
	_jsii_.Set(
		j,
		"sourceFlowConfig",
		val,
	)
}

func (j *jsiiProxy_CfnFlow) SetTasks(val interface{}) {
	_jsii_.Set(
		j,
		"tasks",
		val,
	)
}

func (j *jsiiProxy_CfnFlow) SetTriggerConfig(val interface{}) {
	_jsii_.Set(
		j,
		"triggerConfig",
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
func CfnFlow_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appflow.CfnFlow",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnFlow_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appflow.CfnFlow",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnFlow_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appflow.CfnFlow",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFlow_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_appflow.CfnFlow",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnFlow) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnFlow) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnFlow) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnFlow) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnFlow) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnFlow) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnFlow) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnFlow) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnFlow) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnFlow) Inspect(inspector awscdk.TreeInspector) {
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
func (c *jsiiProxy_CfnFlow) OnPrepare() {
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
func (c *jsiiProxy_CfnFlow) OnSynthesize(session constructs.ISynthesisSession) {
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
func (c *jsiiProxy_CfnFlow) OnValidate() *[]*string {
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
func (c *jsiiProxy_CfnFlow) OverrideLogicalId(newLogicalId *string) {
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
func (c *jsiiProxy_CfnFlow) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnFlow) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnFlow) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnFlow) Synthesize(session awscdk.ISynthesisSession) {
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
func (c *jsiiProxy_CfnFlow) ToString() *string {
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
func (c *jsiiProxy_CfnFlow) Validate() *[]*string {
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
func (c *jsiiProxy_CfnFlow) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnFlow_AggregationConfigProperty struct {
	// `CfnFlow.AggregationConfigProperty.AggregationType`.
	AggregationType *string `json:"aggregationType"`
}

type CfnFlow_AmplitudeSourcePropertiesProperty struct {
	// `CfnFlow.AmplitudeSourcePropertiesProperty.Object`.
	Object *string `json:"object"`
}

type CfnFlow_ConnectorOperatorProperty struct {
	// `CfnFlow.ConnectorOperatorProperty.Amplitude`.
	Amplitude *string `json:"amplitude"`
	// `CfnFlow.ConnectorOperatorProperty.Datadog`.
	Datadog *string `json:"datadog"`
	// `CfnFlow.ConnectorOperatorProperty.Dynatrace`.
	Dynatrace *string `json:"dynatrace"`
	// `CfnFlow.ConnectorOperatorProperty.GoogleAnalytics`.
	GoogleAnalytics *string `json:"googleAnalytics"`
	// `CfnFlow.ConnectorOperatorProperty.InforNexus`.
	InforNexus *string `json:"inforNexus"`
	// `CfnFlow.ConnectorOperatorProperty.Marketo`.
	Marketo *string `json:"marketo"`
	// `CfnFlow.ConnectorOperatorProperty.S3`.
	S3 *string `json:"s3"`
	// `CfnFlow.ConnectorOperatorProperty.Salesforce`.
	Salesforce *string `json:"salesforce"`
	// `CfnFlow.ConnectorOperatorProperty.SAPOData`.
	SapoData *string `json:"sapoData"`
	// `CfnFlow.ConnectorOperatorProperty.ServiceNow`.
	ServiceNow *string `json:"serviceNow"`
	// `CfnFlow.ConnectorOperatorProperty.Singular`.
	Singular *string `json:"singular"`
	// `CfnFlow.ConnectorOperatorProperty.Slack`.
	Slack *string `json:"slack"`
	// `CfnFlow.ConnectorOperatorProperty.Trendmicro`.
	Trendmicro *string `json:"trendmicro"`
	// `CfnFlow.ConnectorOperatorProperty.Veeva`.
	Veeva *string `json:"veeva"`
	// `CfnFlow.ConnectorOperatorProperty.Zendesk`.
	Zendesk *string `json:"zendesk"`
}

type CfnFlow_DatadogSourcePropertiesProperty struct {
	// `CfnFlow.DatadogSourcePropertiesProperty.Object`.
	Object *string `json:"object"`
}

type CfnFlow_DestinationConnectorPropertiesProperty struct {
	// `CfnFlow.DestinationConnectorPropertiesProperty.EventBridge`.
	EventBridge interface{} `json:"eventBridge"`
	// `CfnFlow.DestinationConnectorPropertiesProperty.LookoutMetrics`.
	LookoutMetrics interface{} `json:"lookoutMetrics"`
	// `CfnFlow.DestinationConnectorPropertiesProperty.Redshift`.
	Redshift interface{} `json:"redshift"`
	// `CfnFlow.DestinationConnectorPropertiesProperty.S3`.
	S3 interface{} `json:"s3"`
	// `CfnFlow.DestinationConnectorPropertiesProperty.Salesforce`.
	Salesforce interface{} `json:"salesforce"`
	// `CfnFlow.DestinationConnectorPropertiesProperty.Snowflake`.
	Snowflake interface{} `json:"snowflake"`
	// `CfnFlow.DestinationConnectorPropertiesProperty.Upsolver`.
	Upsolver interface{} `json:"upsolver"`
	// `CfnFlow.DestinationConnectorPropertiesProperty.Zendesk`.
	Zendesk interface{} `json:"zendesk"`
}

type CfnFlow_DestinationFlowConfigProperty struct {
	// `CfnFlow.DestinationFlowConfigProperty.ConnectorType`.
	ConnectorType *string `json:"connectorType"`
	// `CfnFlow.DestinationFlowConfigProperty.DestinationConnectorProperties`.
	DestinationConnectorProperties interface{} `json:"destinationConnectorProperties"`
	// `CfnFlow.DestinationFlowConfigProperty.ConnectorProfileName`.
	ConnectorProfileName *string `json:"connectorProfileName"`
}

type CfnFlow_DynatraceSourcePropertiesProperty struct {
	// `CfnFlow.DynatraceSourcePropertiesProperty.Object`.
	Object *string `json:"object"`
}

type CfnFlow_ErrorHandlingConfigProperty struct {
	// `CfnFlow.ErrorHandlingConfigProperty.BucketName`.
	BucketName *string `json:"bucketName"`
	// `CfnFlow.ErrorHandlingConfigProperty.BucketPrefix`.
	BucketPrefix *string `json:"bucketPrefix"`
	// `CfnFlow.ErrorHandlingConfigProperty.FailOnFirstError`.
	FailOnFirstError interface{} `json:"failOnFirstError"`
}

type CfnFlow_EventBridgeDestinationPropertiesProperty struct {
	// `CfnFlow.EventBridgeDestinationPropertiesProperty.Object`.
	Object *string `json:"object"`
	// `CfnFlow.EventBridgeDestinationPropertiesProperty.ErrorHandlingConfig`.
	ErrorHandlingConfig interface{} `json:"errorHandlingConfig"`
}

type CfnFlow_GoogleAnalyticsSourcePropertiesProperty struct {
	// `CfnFlow.GoogleAnalyticsSourcePropertiesProperty.Object`.
	Object *string `json:"object"`
}

type CfnFlow_IncrementalPullConfigProperty struct {
	// `CfnFlow.IncrementalPullConfigProperty.DatetimeTypeFieldName`.
	DatetimeTypeFieldName *string `json:"datetimeTypeFieldName"`
}

type CfnFlow_InforNexusSourcePropertiesProperty struct {
	// `CfnFlow.InforNexusSourcePropertiesProperty.Object`.
	Object *string `json:"object"`
}

type CfnFlow_LookoutMetricsDestinationPropertiesProperty struct {
	// `CfnFlow.LookoutMetricsDestinationPropertiesProperty.Object`.
	Object *string `json:"object"`
}

type CfnFlow_MarketoSourcePropertiesProperty struct {
	// `CfnFlow.MarketoSourcePropertiesProperty.Object`.
	Object *string `json:"object"`
}

type CfnFlow_PrefixConfigProperty struct {
	// `CfnFlow.PrefixConfigProperty.PrefixFormat`.
	PrefixFormat *string `json:"prefixFormat"`
	// `CfnFlow.PrefixConfigProperty.PrefixType`.
	PrefixType *string `json:"prefixType"`
}

type CfnFlow_RedshiftDestinationPropertiesProperty struct {
	// `CfnFlow.RedshiftDestinationPropertiesProperty.IntermediateBucketName`.
	IntermediateBucketName *string `json:"intermediateBucketName"`
	// `CfnFlow.RedshiftDestinationPropertiesProperty.Object`.
	Object *string `json:"object"`
	// `CfnFlow.RedshiftDestinationPropertiesProperty.BucketPrefix`.
	BucketPrefix *string `json:"bucketPrefix"`
	// `CfnFlow.RedshiftDestinationPropertiesProperty.ErrorHandlingConfig`.
	ErrorHandlingConfig interface{} `json:"errorHandlingConfig"`
}

type CfnFlow_S3DestinationPropertiesProperty struct {
	// `CfnFlow.S3DestinationPropertiesProperty.BucketName`.
	BucketName *string `json:"bucketName"`
	// `CfnFlow.S3DestinationPropertiesProperty.BucketPrefix`.
	BucketPrefix *string `json:"bucketPrefix"`
	// `CfnFlow.S3DestinationPropertiesProperty.S3OutputFormatConfig`.
	S3OutputFormatConfig interface{} `json:"s3OutputFormatConfig"`
}

type CfnFlow_S3InputFormatConfigProperty struct {
	// `CfnFlow.S3InputFormatConfigProperty.S3InputFileType`.
	S3InputFileType *string `json:"s3InputFileType"`
}

type CfnFlow_S3OutputFormatConfigProperty struct {
	// `CfnFlow.S3OutputFormatConfigProperty.AggregationConfig`.
	AggregationConfig interface{} `json:"aggregationConfig"`
	// `CfnFlow.S3OutputFormatConfigProperty.FileType`.
	FileType *string `json:"fileType"`
	// `CfnFlow.S3OutputFormatConfigProperty.PrefixConfig`.
	PrefixConfig interface{} `json:"prefixConfig"`
}

type CfnFlow_S3SourcePropertiesProperty struct {
	// `CfnFlow.S3SourcePropertiesProperty.BucketName`.
	BucketName *string `json:"bucketName"`
	// `CfnFlow.S3SourcePropertiesProperty.BucketPrefix`.
	BucketPrefix *string `json:"bucketPrefix"`
	// `CfnFlow.S3SourcePropertiesProperty.S3InputFormatConfig`.
	S3InputFormatConfig interface{} `json:"s3InputFormatConfig"`
}

type CfnFlow_SAPODataSourcePropertiesProperty struct {
	// `CfnFlow.SAPODataSourcePropertiesProperty.ObjectPath`.
	ObjectPath *string `json:"objectPath"`
}

type CfnFlow_SalesforceDestinationPropertiesProperty struct {
	// `CfnFlow.SalesforceDestinationPropertiesProperty.Object`.
	Object *string `json:"object"`
	// `CfnFlow.SalesforceDestinationPropertiesProperty.ErrorHandlingConfig`.
	ErrorHandlingConfig interface{} `json:"errorHandlingConfig"`
	// `CfnFlow.SalesforceDestinationPropertiesProperty.IdFieldNames`.
	IdFieldNames *[]*string `json:"idFieldNames"`
	// `CfnFlow.SalesforceDestinationPropertiesProperty.WriteOperationType`.
	WriteOperationType *string `json:"writeOperationType"`
}

type CfnFlow_SalesforceSourcePropertiesProperty struct {
	// `CfnFlow.SalesforceSourcePropertiesProperty.Object`.
	Object *string `json:"object"`
	// `CfnFlow.SalesforceSourcePropertiesProperty.EnableDynamicFieldUpdate`.
	EnableDynamicFieldUpdate interface{} `json:"enableDynamicFieldUpdate"`
	// `CfnFlow.SalesforceSourcePropertiesProperty.IncludeDeletedRecords`.
	IncludeDeletedRecords interface{} `json:"includeDeletedRecords"`
}

type CfnFlow_ScheduledTriggerPropertiesProperty struct {
	// `CfnFlow.ScheduledTriggerPropertiesProperty.ScheduleExpression`.
	ScheduleExpression *string `json:"scheduleExpression"`
	// `CfnFlow.ScheduledTriggerPropertiesProperty.DataPullMode`.
	DataPullMode *string `json:"dataPullMode"`
	// `CfnFlow.ScheduledTriggerPropertiesProperty.ScheduleEndTime`.
	ScheduleEndTime *float64 `json:"scheduleEndTime"`
	// `CfnFlow.ScheduledTriggerPropertiesProperty.ScheduleOffset`.
	ScheduleOffset *float64 `json:"scheduleOffset"`
	// `CfnFlow.ScheduledTriggerPropertiesProperty.ScheduleStartTime`.
	ScheduleStartTime *float64 `json:"scheduleStartTime"`
	// `CfnFlow.ScheduledTriggerPropertiesProperty.TimeZone`.
	TimeZone *string `json:"timeZone"`
}

type CfnFlow_ServiceNowSourcePropertiesProperty struct {
	// `CfnFlow.ServiceNowSourcePropertiesProperty.Object`.
	Object *string `json:"object"`
}

type CfnFlow_SingularSourcePropertiesProperty struct {
	// `CfnFlow.SingularSourcePropertiesProperty.Object`.
	Object *string `json:"object"`
}

type CfnFlow_SlackSourcePropertiesProperty struct {
	// `CfnFlow.SlackSourcePropertiesProperty.Object`.
	Object *string `json:"object"`
}

type CfnFlow_SnowflakeDestinationPropertiesProperty struct {
	// `CfnFlow.SnowflakeDestinationPropertiesProperty.IntermediateBucketName`.
	IntermediateBucketName *string `json:"intermediateBucketName"`
	// `CfnFlow.SnowflakeDestinationPropertiesProperty.Object`.
	Object *string `json:"object"`
	// `CfnFlow.SnowflakeDestinationPropertiesProperty.BucketPrefix`.
	BucketPrefix *string `json:"bucketPrefix"`
	// `CfnFlow.SnowflakeDestinationPropertiesProperty.ErrorHandlingConfig`.
	ErrorHandlingConfig interface{} `json:"errorHandlingConfig"`
}

type CfnFlow_SourceConnectorPropertiesProperty struct {
	// `CfnFlow.SourceConnectorPropertiesProperty.Amplitude`.
	Amplitude interface{} `json:"amplitude"`
	// `CfnFlow.SourceConnectorPropertiesProperty.Datadog`.
	Datadog interface{} `json:"datadog"`
	// `CfnFlow.SourceConnectorPropertiesProperty.Dynatrace`.
	Dynatrace interface{} `json:"dynatrace"`
	// `CfnFlow.SourceConnectorPropertiesProperty.GoogleAnalytics`.
	GoogleAnalytics interface{} `json:"googleAnalytics"`
	// `CfnFlow.SourceConnectorPropertiesProperty.InforNexus`.
	InforNexus interface{} `json:"inforNexus"`
	// `CfnFlow.SourceConnectorPropertiesProperty.Marketo`.
	Marketo interface{} `json:"marketo"`
	// `CfnFlow.SourceConnectorPropertiesProperty.S3`.
	S3 interface{} `json:"s3"`
	// `CfnFlow.SourceConnectorPropertiesProperty.Salesforce`.
	Salesforce interface{} `json:"salesforce"`
	// `CfnFlow.SourceConnectorPropertiesProperty.SAPOData`.
	SapoData interface{} `json:"sapoData"`
	// `CfnFlow.SourceConnectorPropertiesProperty.ServiceNow`.
	ServiceNow interface{} `json:"serviceNow"`
	// `CfnFlow.SourceConnectorPropertiesProperty.Singular`.
	Singular interface{} `json:"singular"`
	// `CfnFlow.SourceConnectorPropertiesProperty.Slack`.
	Slack interface{} `json:"slack"`
	// `CfnFlow.SourceConnectorPropertiesProperty.Trendmicro`.
	Trendmicro interface{} `json:"trendmicro"`
	// `CfnFlow.SourceConnectorPropertiesProperty.Veeva`.
	Veeva interface{} `json:"veeva"`
	// `CfnFlow.SourceConnectorPropertiesProperty.Zendesk`.
	Zendesk interface{} `json:"zendesk"`
}

type CfnFlow_SourceFlowConfigProperty struct {
	// `CfnFlow.SourceFlowConfigProperty.ConnectorType`.
	ConnectorType *string `json:"connectorType"`
	// `CfnFlow.SourceFlowConfigProperty.SourceConnectorProperties`.
	SourceConnectorProperties interface{} `json:"sourceConnectorProperties"`
	// `CfnFlow.SourceFlowConfigProperty.ConnectorProfileName`.
	ConnectorProfileName *string `json:"connectorProfileName"`
	// `CfnFlow.SourceFlowConfigProperty.IncrementalPullConfig`.
	IncrementalPullConfig interface{} `json:"incrementalPullConfig"`
}

type CfnFlow_TaskPropertiesObjectProperty struct {
	// `CfnFlow.TaskPropertiesObjectProperty.Key`.
	Key *string `json:"key"`
	// `CfnFlow.TaskPropertiesObjectProperty.Value`.
	Value *string `json:"value"`
}

type CfnFlow_TaskProperty struct {
	// `CfnFlow.TaskProperty.SourceFields`.
	SourceFields *[]*string `json:"sourceFields"`
	// `CfnFlow.TaskProperty.TaskType`.
	TaskType *string `json:"taskType"`
	// `CfnFlow.TaskProperty.ConnectorOperator`.
	ConnectorOperator interface{} `json:"connectorOperator"`
	// `CfnFlow.TaskProperty.DestinationField`.
	DestinationField *string `json:"destinationField"`
	// `CfnFlow.TaskProperty.TaskProperties`.
	TaskProperties interface{} `json:"taskProperties"`
}

type CfnFlow_TrendmicroSourcePropertiesProperty struct {
	// `CfnFlow.TrendmicroSourcePropertiesProperty.Object`.
	Object *string `json:"object"`
}

type CfnFlow_TriggerConfigProperty struct {
	// `CfnFlow.TriggerConfigProperty.TriggerType`.
	TriggerType *string `json:"triggerType"`
	// `CfnFlow.TriggerConfigProperty.TriggerProperties`.
	TriggerProperties interface{} `json:"triggerProperties"`
}

type CfnFlow_UpsolverDestinationPropertiesProperty struct {
	// `CfnFlow.UpsolverDestinationPropertiesProperty.BucketName`.
	BucketName *string `json:"bucketName"`
	// `CfnFlow.UpsolverDestinationPropertiesProperty.S3OutputFormatConfig`.
	S3OutputFormatConfig interface{} `json:"s3OutputFormatConfig"`
	// `CfnFlow.UpsolverDestinationPropertiesProperty.BucketPrefix`.
	BucketPrefix *string `json:"bucketPrefix"`
}

type CfnFlow_UpsolverS3OutputFormatConfigProperty struct {
	// `CfnFlow.UpsolverS3OutputFormatConfigProperty.PrefixConfig`.
	PrefixConfig interface{} `json:"prefixConfig"`
	// `CfnFlow.UpsolverS3OutputFormatConfigProperty.AggregationConfig`.
	AggregationConfig interface{} `json:"aggregationConfig"`
	// `CfnFlow.UpsolverS3OutputFormatConfigProperty.FileType`.
	FileType *string `json:"fileType"`
}

type CfnFlow_VeevaSourcePropertiesProperty struct {
	// `CfnFlow.VeevaSourcePropertiesProperty.Object`.
	Object *string `json:"object"`
	// `CfnFlow.VeevaSourcePropertiesProperty.DocumentType`.
	DocumentType *string `json:"documentType"`
	// `CfnFlow.VeevaSourcePropertiesProperty.IncludeAllVersions`.
	IncludeAllVersions interface{} `json:"includeAllVersions"`
	// `CfnFlow.VeevaSourcePropertiesProperty.IncludeRenditions`.
	IncludeRenditions interface{} `json:"includeRenditions"`
	// `CfnFlow.VeevaSourcePropertiesProperty.IncludeSourceFiles`.
	IncludeSourceFiles interface{} `json:"includeSourceFiles"`
}

type CfnFlow_ZendeskDestinationPropertiesProperty struct {
	// `CfnFlow.ZendeskDestinationPropertiesProperty.Object`.
	Object *string `json:"object"`
	// `CfnFlow.ZendeskDestinationPropertiesProperty.ErrorHandlingConfig`.
	ErrorHandlingConfig interface{} `json:"errorHandlingConfig"`
	// `CfnFlow.ZendeskDestinationPropertiesProperty.IdFieldNames`.
	IdFieldNames *[]*string `json:"idFieldNames"`
	// `CfnFlow.ZendeskDestinationPropertiesProperty.WriteOperationType`.
	WriteOperationType *string `json:"writeOperationType"`
}

type CfnFlow_ZendeskSourcePropertiesProperty struct {
	// `CfnFlow.ZendeskSourcePropertiesProperty.Object`.
	Object *string `json:"object"`
}

// Properties for defining a `AWS::AppFlow::Flow`.
type CfnFlowProps struct {
	// `AWS::AppFlow::Flow.DestinationFlowConfigList`.
	DestinationFlowConfigList interface{} `json:"destinationFlowConfigList"`
	// `AWS::AppFlow::Flow.FlowName`.
	FlowName *string `json:"flowName"`
	// `AWS::AppFlow::Flow.SourceFlowConfig`.
	SourceFlowConfig interface{} `json:"sourceFlowConfig"`
	// `AWS::AppFlow::Flow.Tasks`.
	Tasks interface{} `json:"tasks"`
	// `AWS::AppFlow::Flow.TriggerConfig`.
	TriggerConfig interface{} `json:"triggerConfig"`
	// `AWS::AppFlow::Flow.Description`.
	Description *string `json:"description"`
	// `AWS::AppFlow::Flow.KMSArn`.
	KmsArn *string `json:"kmsArn"`
	// `AWS::AppFlow::Flow.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
}

