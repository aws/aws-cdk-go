package awsappflow

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsappflow/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::AppFlow::ConnectorProfile`.
//
// The `AWS::AppFlow::ConnectorProfile` resource is an Amazon AppFlow resource type that specifies the configuration profile for an instance of a connector. This includes the provided name, credentials ARN, connection-mode, and so on. The fields that are common to all types of connector profiles are explicitly specified under the `Properties` field. The rest of the connector-specific properties are specified under `Properties/ConnectorProfileConfig` .
//
// > If you want to use AWS CloudFormation to create a connector profile for connectors that implement OAuth (such as Salesforce, Slack, Zendesk, and Google Analytics), you must fetch the access and refresh tokens. You can do this by implementing your own UI for OAuth, or by retrieving the tokens from elsewhere. Alternatively, you can use the Amazon AppFlow console to create the connector profile, and then use that connector profile in the flow creation CloudFormation template.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//
//   var basicAuthCredentials interface{}
//   var oAuthCredentials interface{}
//   cfnConnectorProfile := appflow.NewCfnConnectorProfile(this, jsii.String("MyCfnConnectorProfile"), &cfnConnectorProfileProps{
//   	connectionMode: jsii.String("connectionMode"),
//   	connectorProfileName: jsii.String("connectorProfileName"),
//   	connectorType: jsii.String("connectorType"),
//
//   	// the properties below are optional
//   	connectorProfileConfig: &connectorProfileConfigProperty{
//   		connectorProfileCredentials: &connectorProfileCredentialsProperty{
//   			amplitude: &amplitudeConnectorProfileCredentialsProperty{
//   				apiKey: jsii.String("apiKey"),
//   				secretKey: jsii.String("secretKey"),
//   			},
//   			datadog: &datadogConnectorProfileCredentialsProperty{
//   				apiKey: jsii.String("apiKey"),
//   				applicationKey: jsii.String("applicationKey"),
//   			},
//   			dynatrace: &dynatraceConnectorProfileCredentialsProperty{
//   				apiToken: jsii.String("apiToken"),
//   			},
//   			googleAnalytics: &googleAnalyticsConnectorProfileCredentialsProperty{
//   				clientId: jsii.String("clientId"),
//   				clientSecret: jsii.String("clientSecret"),
//
//   				// the properties below are optional
//   				accessToken: jsii.String("accessToken"),
//   				connectorOAuthRequest: &connectorOAuthRequestProperty{
//   					authCode: jsii.String("authCode"),
//   					redirectUri: jsii.String("redirectUri"),
//   				},
//   				refreshToken: jsii.String("refreshToken"),
//   			},
//   			inforNexus: &inforNexusConnectorProfileCredentialsProperty{
//   				accessKeyId: jsii.String("accessKeyId"),
//   				datakey: jsii.String("datakey"),
//   				secretAccessKey: jsii.String("secretAccessKey"),
//   				userId: jsii.String("userId"),
//   			},
//   			marketo: &marketoConnectorProfileCredentialsProperty{
//   				clientId: jsii.String("clientId"),
//   				clientSecret: jsii.String("clientSecret"),
//
//   				// the properties below are optional
//   				accessToken: jsii.String("accessToken"),
//   				connectorOAuthRequest: &connectorOAuthRequestProperty{
//   					authCode: jsii.String("authCode"),
//   					redirectUri: jsii.String("redirectUri"),
//   				},
//   			},
//   			redshift: &redshiftConnectorProfileCredentialsProperty{
//   				password: jsii.String("password"),
//   				username: jsii.String("username"),
//   			},
//   			salesforce: &salesforceConnectorProfileCredentialsProperty{
//   				accessToken: jsii.String("accessToken"),
//   				clientCredentialsArn: jsii.String("clientCredentialsArn"),
//   				connectorOAuthRequest: &connectorOAuthRequestProperty{
//   					authCode: jsii.String("authCode"),
//   					redirectUri: jsii.String("redirectUri"),
//   				},
//   				refreshToken: jsii.String("refreshToken"),
//   			},
//   			sapoData: &sAPODataConnectorProfileCredentialsProperty{
//   				basicAuthCredentials: basicAuthCredentials,
//   				oAuthCredentials: oAuthCredentials,
//   			},
//   			serviceNow: &serviceNowConnectorProfileCredentialsProperty{
//   				password: jsii.String("password"),
//   				username: jsii.String("username"),
//   			},
//   			singular: &singularConnectorProfileCredentialsProperty{
//   				apiKey: jsii.String("apiKey"),
//   			},
//   			slack: &slackConnectorProfileCredentialsProperty{
//   				clientId: jsii.String("clientId"),
//   				clientSecret: jsii.String("clientSecret"),
//
//   				// the properties below are optional
//   				accessToken: jsii.String("accessToken"),
//   				connectorOAuthRequest: &connectorOAuthRequestProperty{
//   					authCode: jsii.String("authCode"),
//   					redirectUri: jsii.String("redirectUri"),
//   				},
//   			},
//   			snowflake: &snowflakeConnectorProfileCredentialsProperty{
//   				password: jsii.String("password"),
//   				username: jsii.String("username"),
//   			},
//   			trendmicro: &trendmicroConnectorProfileCredentialsProperty{
//   				apiSecretKey: jsii.String("apiSecretKey"),
//   			},
//   			veeva: &veevaConnectorProfileCredentialsProperty{
//   				password: jsii.String("password"),
//   				username: jsii.String("username"),
//   			},
//   			zendesk: &zendeskConnectorProfileCredentialsProperty{
//   				clientId: jsii.String("clientId"),
//   				clientSecret: jsii.String("clientSecret"),
//
//   				// the properties below are optional
//   				accessToken: jsii.String("accessToken"),
//   				connectorOAuthRequest: &connectorOAuthRequestProperty{
//   					authCode: jsii.String("authCode"),
//   					redirectUri: jsii.String("redirectUri"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		connectorProfileProperties: &connectorProfilePropertiesProperty{
//   			datadog: &datadogConnectorProfilePropertiesProperty{
//   				instanceUrl: jsii.String("instanceUrl"),
//   			},
//   			dynatrace: &dynatraceConnectorProfilePropertiesProperty{
//   				instanceUrl: jsii.String("instanceUrl"),
//   			},
//   			inforNexus: &inforNexusConnectorProfilePropertiesProperty{
//   				instanceUrl: jsii.String("instanceUrl"),
//   			},
//   			marketo: &marketoConnectorProfilePropertiesProperty{
//   				instanceUrl: jsii.String("instanceUrl"),
//   			},
//   			redshift: &redshiftConnectorProfilePropertiesProperty{
//   				bucketName: jsii.String("bucketName"),
//   				databaseUrl: jsii.String("databaseUrl"),
//   				roleArn: jsii.String("roleArn"),
//
//   				// the properties below are optional
//   				bucketPrefix: jsii.String("bucketPrefix"),
//   			},
//   			salesforce: &salesforceConnectorProfilePropertiesProperty{
//   				instanceUrl: jsii.String("instanceUrl"),
//   				isSandboxEnvironment: jsii.Boolean(false),
//   			},
//   			sapoData: &sAPODataConnectorProfilePropertiesProperty{
//   				applicationHostUrl: jsii.String("applicationHostUrl"),
//   				applicationServicePath: jsii.String("applicationServicePath"),
//   				clientNumber: jsii.String("clientNumber"),
//   				logonLanguage: jsii.String("logonLanguage"),
//   				oAuthProperties: &oAuthPropertiesProperty{
//   					authCodeUrl: jsii.String("authCodeUrl"),
//   					oAuthScopes: []*string{
//   						jsii.String("oAuthScopes"),
//   					},
//   					tokenUrl: jsii.String("tokenUrl"),
//   				},
//   				portNumber: jsii.Number(123),
//   				privateLinkServiceName: jsii.String("privateLinkServiceName"),
//   			},
//   			serviceNow: &serviceNowConnectorProfilePropertiesProperty{
//   				instanceUrl: jsii.String("instanceUrl"),
//   			},
//   			slack: &slackConnectorProfilePropertiesProperty{
//   				instanceUrl: jsii.String("instanceUrl"),
//   			},
//   			snowflake: &snowflakeConnectorProfilePropertiesProperty{
//   				bucketName: jsii.String("bucketName"),
//   				stage: jsii.String("stage"),
//   				warehouse: jsii.String("warehouse"),
//
//   				// the properties below are optional
//   				accountName: jsii.String("accountName"),
//   				bucketPrefix: jsii.String("bucketPrefix"),
//   				privateLinkServiceName: jsii.String("privateLinkServiceName"),
//   				region: jsii.String("region"),
//   			},
//   			veeva: &veevaConnectorProfilePropertiesProperty{
//   				instanceUrl: jsii.String("instanceUrl"),
//   			},
//   			zendesk: &zendeskConnectorProfilePropertiesProperty{
//   				instanceUrl: jsii.String("instanceUrl"),
//   			},
//   		},
//   	},
//   	kmsArn: jsii.String("kmsArn"),
//   })
//
type CfnConnectorProfile interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of the connector profile.
	AttrConnectorProfileArn() *string
	// The Amazon Resource Name (ARN) of the connector profile credentials.
	AttrCredentialsArn() *string
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Indicates the connection mode and if it is public or private.
	ConnectionMode() *string
	SetConnectionMode(val *string)
	// Defines the connector-specific configuration and credentials.
	ConnectorProfileConfig() interface{}
	SetConnectorProfileConfig(val interface{})
	// The name of the connector profile.
	//
	// The name is unique for each `ConnectorProfile` in the AWS account .
	ConnectorProfileName() *string
	SetConnectorProfileName(val *string)
	// The type of connector, such as Salesforce, Amplitude, and so on.
	ConnectorType() *string
	SetConnectorType(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// The ARN (Amazon Resource Name) of the Key Management Service (KMS) key you provide for encryption.
	//
	// This is required if you do not want to use the Amazon AppFlow-managed KMS key. If you don't provide anything here, Amazon AppFlow uses the Amazon AppFlow-managed KMS key.
	KmsArn() *string
	SetKmsArn(val *string)
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

func (c *jsiiProxy_CfnConnectorProfile) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnConnectorProfile) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnConnectorProfile) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnConnectorProfile) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnConnectorProfile) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnConnectorProfile) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnConnectorProfile) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

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

func (c *jsiiProxy_CfnConnectorProfile) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnConnectorProfile) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnConnectorProfile) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_CfnConnectorProfile) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

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

func (c *jsiiProxy_CfnConnectorProfile) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_CfnConnectorProfile) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The `AmplitudeConnectorProfileCredentials` property type specifies the connector-specific credentials required when using Amplitude.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   amplitudeConnectorProfileCredentialsProperty := &amplitudeConnectorProfileCredentialsProperty{
//   	apiKey: jsii.String("apiKey"),
//   	secretKey: jsii.String("secretKey"),
//   }
//
type CfnConnectorProfile_AmplitudeConnectorProfileCredentialsProperty struct {
	// A unique alphanumeric identifier used to authenticate a user, developer, or calling program to your API.
	ApiKey *string `json:"apiKey" yaml:"apiKey"`
	// The Secret Access Key portion of the credentials.
	SecretKey *string `json:"secretKey" yaml:"secretKey"`
}

// The `ConnectorOAuthRequest` property type specifies the select connectors for which the OAuth workflow is supported, such as Salesforce, Google Analytics, Marketo, Zendesk, and Slack.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   connectorOAuthRequestProperty := &connectorOAuthRequestProperty{
//   	authCode: jsii.String("authCode"),
//   	redirectUri: jsii.String("redirectUri"),
//   }
//
type CfnConnectorProfile_ConnectorOAuthRequestProperty struct {
	// The code provided by the connector when it has been authenticated via the connected app.
	AuthCode *string `json:"authCode" yaml:"authCode"`
	// The URL to which the authentication server redirects the browser after authorization has been granted.
	RedirectUri *string `json:"redirectUri" yaml:"redirectUri"`
}

// Defines the connector-specific configuration and credentials for the connector profile.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//
//   var basicAuthCredentials interface{}
//   var oAuthCredentials interface{}
//   connectorProfileConfigProperty := &connectorProfileConfigProperty{
//   	connectorProfileCredentials: &connectorProfileCredentialsProperty{
//   		amplitude: &amplitudeConnectorProfileCredentialsProperty{
//   			apiKey: jsii.String("apiKey"),
//   			secretKey: jsii.String("secretKey"),
//   		},
//   		datadog: &datadogConnectorProfileCredentialsProperty{
//   			apiKey: jsii.String("apiKey"),
//   			applicationKey: jsii.String("applicationKey"),
//   		},
//   		dynatrace: &dynatraceConnectorProfileCredentialsProperty{
//   			apiToken: jsii.String("apiToken"),
//   		},
//   		googleAnalytics: &googleAnalyticsConnectorProfileCredentialsProperty{
//   			clientId: jsii.String("clientId"),
//   			clientSecret: jsii.String("clientSecret"),
//
//   			// the properties below are optional
//   			accessToken: jsii.String("accessToken"),
//   			connectorOAuthRequest: &connectorOAuthRequestProperty{
//   				authCode: jsii.String("authCode"),
//   				redirectUri: jsii.String("redirectUri"),
//   			},
//   			refreshToken: jsii.String("refreshToken"),
//   		},
//   		inforNexus: &inforNexusConnectorProfileCredentialsProperty{
//   			accessKeyId: jsii.String("accessKeyId"),
//   			datakey: jsii.String("datakey"),
//   			secretAccessKey: jsii.String("secretAccessKey"),
//   			userId: jsii.String("userId"),
//   		},
//   		marketo: &marketoConnectorProfileCredentialsProperty{
//   			clientId: jsii.String("clientId"),
//   			clientSecret: jsii.String("clientSecret"),
//
//   			// the properties below are optional
//   			accessToken: jsii.String("accessToken"),
//   			connectorOAuthRequest: &connectorOAuthRequestProperty{
//   				authCode: jsii.String("authCode"),
//   				redirectUri: jsii.String("redirectUri"),
//   			},
//   		},
//   		redshift: &redshiftConnectorProfileCredentialsProperty{
//   			password: jsii.String("password"),
//   			username: jsii.String("username"),
//   		},
//   		salesforce: &salesforceConnectorProfileCredentialsProperty{
//   			accessToken: jsii.String("accessToken"),
//   			clientCredentialsArn: jsii.String("clientCredentialsArn"),
//   			connectorOAuthRequest: &connectorOAuthRequestProperty{
//   				authCode: jsii.String("authCode"),
//   				redirectUri: jsii.String("redirectUri"),
//   			},
//   			refreshToken: jsii.String("refreshToken"),
//   		},
//   		sapoData: &sAPODataConnectorProfileCredentialsProperty{
//   			basicAuthCredentials: basicAuthCredentials,
//   			oAuthCredentials: oAuthCredentials,
//   		},
//   		serviceNow: &serviceNowConnectorProfileCredentialsProperty{
//   			password: jsii.String("password"),
//   			username: jsii.String("username"),
//   		},
//   		singular: &singularConnectorProfileCredentialsProperty{
//   			apiKey: jsii.String("apiKey"),
//   		},
//   		slack: &slackConnectorProfileCredentialsProperty{
//   			clientId: jsii.String("clientId"),
//   			clientSecret: jsii.String("clientSecret"),
//
//   			// the properties below are optional
//   			accessToken: jsii.String("accessToken"),
//   			connectorOAuthRequest: &connectorOAuthRequestProperty{
//   				authCode: jsii.String("authCode"),
//   				redirectUri: jsii.String("redirectUri"),
//   			},
//   		},
//   		snowflake: &snowflakeConnectorProfileCredentialsProperty{
//   			password: jsii.String("password"),
//   			username: jsii.String("username"),
//   		},
//   		trendmicro: &trendmicroConnectorProfileCredentialsProperty{
//   			apiSecretKey: jsii.String("apiSecretKey"),
//   		},
//   		veeva: &veevaConnectorProfileCredentialsProperty{
//   			password: jsii.String("password"),
//   			username: jsii.String("username"),
//   		},
//   		zendesk: &zendeskConnectorProfileCredentialsProperty{
//   			clientId: jsii.String("clientId"),
//   			clientSecret: jsii.String("clientSecret"),
//
//   			// the properties below are optional
//   			accessToken: jsii.String("accessToken"),
//   			connectorOAuthRequest: &connectorOAuthRequestProperty{
//   				authCode: jsii.String("authCode"),
//   				redirectUri: jsii.String("redirectUri"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	connectorProfileProperties: &connectorProfilePropertiesProperty{
//   		datadog: &datadogConnectorProfilePropertiesProperty{
//   			instanceUrl: jsii.String("instanceUrl"),
//   		},
//   		dynatrace: &dynatraceConnectorProfilePropertiesProperty{
//   			instanceUrl: jsii.String("instanceUrl"),
//   		},
//   		inforNexus: &inforNexusConnectorProfilePropertiesProperty{
//   			instanceUrl: jsii.String("instanceUrl"),
//   		},
//   		marketo: &marketoConnectorProfilePropertiesProperty{
//   			instanceUrl: jsii.String("instanceUrl"),
//   		},
//   		redshift: &redshiftConnectorProfilePropertiesProperty{
//   			bucketName: jsii.String("bucketName"),
//   			databaseUrl: jsii.String("databaseUrl"),
//   			roleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			bucketPrefix: jsii.String("bucketPrefix"),
//   		},
//   		salesforce: &salesforceConnectorProfilePropertiesProperty{
//   			instanceUrl: jsii.String("instanceUrl"),
//   			isSandboxEnvironment: jsii.Boolean(false),
//   		},
//   		sapoData: &sAPODataConnectorProfilePropertiesProperty{
//   			applicationHostUrl: jsii.String("applicationHostUrl"),
//   			applicationServicePath: jsii.String("applicationServicePath"),
//   			clientNumber: jsii.String("clientNumber"),
//   			logonLanguage: jsii.String("logonLanguage"),
//   			oAuthProperties: &oAuthPropertiesProperty{
//   				authCodeUrl: jsii.String("authCodeUrl"),
//   				oAuthScopes: []*string{
//   					jsii.String("oAuthScopes"),
//   				},
//   				tokenUrl: jsii.String("tokenUrl"),
//   			},
//   			portNumber: jsii.Number(123),
//   			privateLinkServiceName: jsii.String("privateLinkServiceName"),
//   		},
//   		serviceNow: &serviceNowConnectorProfilePropertiesProperty{
//   			instanceUrl: jsii.String("instanceUrl"),
//   		},
//   		slack: &slackConnectorProfilePropertiesProperty{
//   			instanceUrl: jsii.String("instanceUrl"),
//   		},
//   		snowflake: &snowflakeConnectorProfilePropertiesProperty{
//   			bucketName: jsii.String("bucketName"),
//   			stage: jsii.String("stage"),
//   			warehouse: jsii.String("warehouse"),
//
//   			// the properties below are optional
//   			accountName: jsii.String("accountName"),
//   			bucketPrefix: jsii.String("bucketPrefix"),
//   			privateLinkServiceName: jsii.String("privateLinkServiceName"),
//   			region: jsii.String("region"),
//   		},
//   		veeva: &veevaConnectorProfilePropertiesProperty{
//   			instanceUrl: jsii.String("instanceUrl"),
//   		},
//   		zendesk: &zendeskConnectorProfilePropertiesProperty{
//   			instanceUrl: jsii.String("instanceUrl"),
//   		},
//   	},
//   }
//
type CfnConnectorProfile_ConnectorProfileConfigProperty struct {
	// The connector-specific credentials required by each connector.
	ConnectorProfileCredentials interface{} `json:"connectorProfileCredentials" yaml:"connectorProfileCredentials"`
	// The connector-specific properties of the profile configuration.
	ConnectorProfileProperties interface{} `json:"connectorProfileProperties" yaml:"connectorProfileProperties"`
}

// The `ConnectorProfileCredentials` property type specifies the connector-specific credentials required by a given connector.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//
//   var basicAuthCredentials interface{}
//   var oAuthCredentials interface{}
//   connectorProfileCredentialsProperty := &connectorProfileCredentialsProperty{
//   	amplitude: &amplitudeConnectorProfileCredentialsProperty{
//   		apiKey: jsii.String("apiKey"),
//   		secretKey: jsii.String("secretKey"),
//   	},
//   	datadog: &datadogConnectorProfileCredentialsProperty{
//   		apiKey: jsii.String("apiKey"),
//   		applicationKey: jsii.String("applicationKey"),
//   	},
//   	dynatrace: &dynatraceConnectorProfileCredentialsProperty{
//   		apiToken: jsii.String("apiToken"),
//   	},
//   	googleAnalytics: &googleAnalyticsConnectorProfileCredentialsProperty{
//   		clientId: jsii.String("clientId"),
//   		clientSecret: jsii.String("clientSecret"),
//
//   		// the properties below are optional
//   		accessToken: jsii.String("accessToken"),
//   		connectorOAuthRequest: &connectorOAuthRequestProperty{
//   			authCode: jsii.String("authCode"),
//   			redirectUri: jsii.String("redirectUri"),
//   		},
//   		refreshToken: jsii.String("refreshToken"),
//   	},
//   	inforNexus: &inforNexusConnectorProfileCredentialsProperty{
//   		accessKeyId: jsii.String("accessKeyId"),
//   		datakey: jsii.String("datakey"),
//   		secretAccessKey: jsii.String("secretAccessKey"),
//   		userId: jsii.String("userId"),
//   	},
//   	marketo: &marketoConnectorProfileCredentialsProperty{
//   		clientId: jsii.String("clientId"),
//   		clientSecret: jsii.String("clientSecret"),
//
//   		// the properties below are optional
//   		accessToken: jsii.String("accessToken"),
//   		connectorOAuthRequest: &connectorOAuthRequestProperty{
//   			authCode: jsii.String("authCode"),
//   			redirectUri: jsii.String("redirectUri"),
//   		},
//   	},
//   	redshift: &redshiftConnectorProfileCredentialsProperty{
//   		password: jsii.String("password"),
//   		username: jsii.String("username"),
//   	},
//   	salesforce: &salesforceConnectorProfileCredentialsProperty{
//   		accessToken: jsii.String("accessToken"),
//   		clientCredentialsArn: jsii.String("clientCredentialsArn"),
//   		connectorOAuthRequest: &connectorOAuthRequestProperty{
//   			authCode: jsii.String("authCode"),
//   			redirectUri: jsii.String("redirectUri"),
//   		},
//   		refreshToken: jsii.String("refreshToken"),
//   	},
//   	sapoData: &sAPODataConnectorProfileCredentialsProperty{
//   		basicAuthCredentials: basicAuthCredentials,
//   		oAuthCredentials: oAuthCredentials,
//   	},
//   	serviceNow: &serviceNowConnectorProfileCredentialsProperty{
//   		password: jsii.String("password"),
//   		username: jsii.String("username"),
//   	},
//   	singular: &singularConnectorProfileCredentialsProperty{
//   		apiKey: jsii.String("apiKey"),
//   	},
//   	slack: &slackConnectorProfileCredentialsProperty{
//   		clientId: jsii.String("clientId"),
//   		clientSecret: jsii.String("clientSecret"),
//
//   		// the properties below are optional
//   		accessToken: jsii.String("accessToken"),
//   		connectorOAuthRequest: &connectorOAuthRequestProperty{
//   			authCode: jsii.String("authCode"),
//   			redirectUri: jsii.String("redirectUri"),
//   		},
//   	},
//   	snowflake: &snowflakeConnectorProfileCredentialsProperty{
//   		password: jsii.String("password"),
//   		username: jsii.String("username"),
//   	},
//   	trendmicro: &trendmicroConnectorProfileCredentialsProperty{
//   		apiSecretKey: jsii.String("apiSecretKey"),
//   	},
//   	veeva: &veevaConnectorProfileCredentialsProperty{
//   		password: jsii.String("password"),
//   		username: jsii.String("username"),
//   	},
//   	zendesk: &zendeskConnectorProfileCredentialsProperty{
//   		clientId: jsii.String("clientId"),
//   		clientSecret: jsii.String("clientSecret"),
//
//   		// the properties below are optional
//   		accessToken: jsii.String("accessToken"),
//   		connectorOAuthRequest: &connectorOAuthRequestProperty{
//   			authCode: jsii.String("authCode"),
//   			redirectUri: jsii.String("redirectUri"),
//   		},
//   	},
//   }
//
type CfnConnectorProfile_ConnectorProfileCredentialsProperty struct {
	// The connector-specific credentials required when using Amplitude.
	Amplitude interface{} `json:"amplitude" yaml:"amplitude"`
	// The connector-specific credentials required when using Datadog.
	Datadog interface{} `json:"datadog" yaml:"datadog"`
	// The connector-specific credentials required when using Dynatrace.
	Dynatrace interface{} `json:"dynatrace" yaml:"dynatrace"`
	// The connector-specific credentials required when using Google Analytics.
	GoogleAnalytics interface{} `json:"googleAnalytics" yaml:"googleAnalytics"`
	// The connector-specific credentials required when using Infor Nexus.
	InforNexus interface{} `json:"inforNexus" yaml:"inforNexus"`
	// The connector-specific credentials required when using Marketo.
	Marketo interface{} `json:"marketo" yaml:"marketo"`
	// The connector-specific credentials required when using Amazon Redshift.
	Redshift interface{} `json:"redshift" yaml:"redshift"`
	// The connector-specific credentials required when using Salesforce.
	Salesforce interface{} `json:"salesforce" yaml:"salesforce"`
	// `CfnConnectorProfile.ConnectorProfileCredentialsProperty.SAPOData`.
	SapoData interface{} `json:"sapoData" yaml:"sapoData"`
	// The connector-specific credentials required when using ServiceNow.
	ServiceNow interface{} `json:"serviceNow" yaml:"serviceNow"`
	// The connector-specific credentials required when using Singular.
	Singular interface{} `json:"singular" yaml:"singular"`
	// The connector-specific credentials required when using Slack.
	Slack interface{} `json:"slack" yaml:"slack"`
	// The connector-specific credentials required when using Snowflake.
	Snowflake interface{} `json:"snowflake" yaml:"snowflake"`
	// The connector-specific credentials required when using Trend Micro.
	Trendmicro interface{} `json:"trendmicro" yaml:"trendmicro"`
	// The connector-specific credentials required when using Veeva.
	Veeva interface{} `json:"veeva" yaml:"veeva"`
	// The connector-specific credentials required when using Zendesk.
	Zendesk interface{} `json:"zendesk" yaml:"zendesk"`
}

// The `ConnectorProfileProperties` property type specifies the connector-specific profile properties required by each connector.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   connectorProfilePropertiesProperty := &connectorProfilePropertiesProperty{
//   	datadog: &datadogConnectorProfilePropertiesProperty{
//   		instanceUrl: jsii.String("instanceUrl"),
//   	},
//   	dynatrace: &dynatraceConnectorProfilePropertiesProperty{
//   		instanceUrl: jsii.String("instanceUrl"),
//   	},
//   	inforNexus: &inforNexusConnectorProfilePropertiesProperty{
//   		instanceUrl: jsii.String("instanceUrl"),
//   	},
//   	marketo: &marketoConnectorProfilePropertiesProperty{
//   		instanceUrl: jsii.String("instanceUrl"),
//   	},
//   	redshift: &redshiftConnectorProfilePropertiesProperty{
//   		bucketName: jsii.String("bucketName"),
//   		databaseUrl: jsii.String("databaseUrl"),
//   		roleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		bucketPrefix: jsii.String("bucketPrefix"),
//   	},
//   	salesforce: &salesforceConnectorProfilePropertiesProperty{
//   		instanceUrl: jsii.String("instanceUrl"),
//   		isSandboxEnvironment: jsii.Boolean(false),
//   	},
//   	sapoData: &sAPODataConnectorProfilePropertiesProperty{
//   		applicationHostUrl: jsii.String("applicationHostUrl"),
//   		applicationServicePath: jsii.String("applicationServicePath"),
//   		clientNumber: jsii.String("clientNumber"),
//   		logonLanguage: jsii.String("logonLanguage"),
//   		oAuthProperties: &oAuthPropertiesProperty{
//   			authCodeUrl: jsii.String("authCodeUrl"),
//   			oAuthScopes: []*string{
//   				jsii.String("oAuthScopes"),
//   			},
//   			tokenUrl: jsii.String("tokenUrl"),
//   		},
//   		portNumber: jsii.Number(123),
//   		privateLinkServiceName: jsii.String("privateLinkServiceName"),
//   	},
//   	serviceNow: &serviceNowConnectorProfilePropertiesProperty{
//   		instanceUrl: jsii.String("instanceUrl"),
//   	},
//   	slack: &slackConnectorProfilePropertiesProperty{
//   		instanceUrl: jsii.String("instanceUrl"),
//   	},
//   	snowflake: &snowflakeConnectorProfilePropertiesProperty{
//   		bucketName: jsii.String("bucketName"),
//   		stage: jsii.String("stage"),
//   		warehouse: jsii.String("warehouse"),
//
//   		// the properties below are optional
//   		accountName: jsii.String("accountName"),
//   		bucketPrefix: jsii.String("bucketPrefix"),
//   		privateLinkServiceName: jsii.String("privateLinkServiceName"),
//   		region: jsii.String("region"),
//   	},
//   	veeva: &veevaConnectorProfilePropertiesProperty{
//   		instanceUrl: jsii.String("instanceUrl"),
//   	},
//   	zendesk: &zendeskConnectorProfilePropertiesProperty{
//   		instanceUrl: jsii.String("instanceUrl"),
//   	},
//   }
//
type CfnConnectorProfile_ConnectorProfilePropertiesProperty struct {
	// The connector-specific properties required by Datadog.
	Datadog interface{} `json:"datadog" yaml:"datadog"`
	// The connector-specific properties required by Dynatrace.
	Dynatrace interface{} `json:"dynatrace" yaml:"dynatrace"`
	// The connector-specific properties required by Infor Nexus.
	InforNexus interface{} `json:"inforNexus" yaml:"inforNexus"`
	// The connector-specific properties required by Marketo.
	Marketo interface{} `json:"marketo" yaml:"marketo"`
	// The connector-specific properties required by Amazon Redshift.
	Redshift interface{} `json:"redshift" yaml:"redshift"`
	// The connector-specific properties required by Salesforce.
	Salesforce interface{} `json:"salesforce" yaml:"salesforce"`
	// `CfnConnectorProfile.ConnectorProfilePropertiesProperty.SAPOData`.
	SapoData interface{} `json:"sapoData" yaml:"sapoData"`
	// The connector-specific properties required by serviceNow.
	ServiceNow interface{} `json:"serviceNow" yaml:"serviceNow"`
	// The connector-specific properties required by Slack.
	Slack interface{} `json:"slack" yaml:"slack"`
	// The connector-specific properties required by Snowflake.
	Snowflake interface{} `json:"snowflake" yaml:"snowflake"`
	// The connector-specific properties required by Veeva.
	Veeva interface{} `json:"veeva" yaml:"veeva"`
	// The connector-specific properties required by Zendesk.
	Zendesk interface{} `json:"zendesk" yaml:"zendesk"`
}

// The `DatadogConnectorProfileCredentials` property type specifies the connector-specific credentials required by Datadog.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   datadogConnectorProfileCredentialsProperty := &datadogConnectorProfileCredentialsProperty{
//   	apiKey: jsii.String("apiKey"),
//   	applicationKey: jsii.String("applicationKey"),
//   }
//
type CfnConnectorProfile_DatadogConnectorProfileCredentialsProperty struct {
	// A unique alphanumeric identifier used to authenticate a user, developer, or calling program to your API.
	ApiKey *string `json:"apiKey" yaml:"apiKey"`
	// Application keys, in conjunction with your API key, give you full access to Datadog’s programmatic API.
	//
	// Application keys are associated with the user account that created them. The application key is used to log all requests made to the API.
	ApplicationKey *string `json:"applicationKey" yaml:"applicationKey"`
}

// The `DatadogConnectorProfileProperties` property type specifies the connector-specific profile properties required by Datadog.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   datadogConnectorProfilePropertiesProperty := &datadogConnectorProfilePropertiesProperty{
//   	instanceUrl: jsii.String("instanceUrl"),
//   }
//
type CfnConnectorProfile_DatadogConnectorProfilePropertiesProperty struct {
	// The location of the Datadog resource.
	InstanceUrl *string `json:"instanceUrl" yaml:"instanceUrl"`
}

// The `DynatraceConnectorProfileCredentials` property type specifies the connector-specific profile credentials required by Dynatrace.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   dynatraceConnectorProfileCredentialsProperty := &dynatraceConnectorProfileCredentialsProperty{
//   	apiToken: jsii.String("apiToken"),
//   }
//
type CfnConnectorProfile_DynatraceConnectorProfileCredentialsProperty struct {
	// The API tokens used by Dynatrace API to authenticate various API calls.
	ApiToken *string `json:"apiToken" yaml:"apiToken"`
}

// The `DynatraceConnectorProfileProperties` property type specifies the connector-specific profile properties required by Dynatrace.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   dynatraceConnectorProfilePropertiesProperty := &dynatraceConnectorProfilePropertiesProperty{
//   	instanceUrl: jsii.String("instanceUrl"),
//   }
//
type CfnConnectorProfile_DynatraceConnectorProfilePropertiesProperty struct {
	// The location of the Dynatrace resource.
	InstanceUrl *string `json:"instanceUrl" yaml:"instanceUrl"`
}

// The `GoogleAnalyticsConnectorProfileCredentials` property type specifies the connector-specific profile credentials required by Google Analytics.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   googleAnalyticsConnectorProfileCredentialsProperty := &googleAnalyticsConnectorProfileCredentialsProperty{
//   	clientId: jsii.String("clientId"),
//   	clientSecret: jsii.String("clientSecret"),
//
//   	// the properties below are optional
//   	accessToken: jsii.String("accessToken"),
//   	connectorOAuthRequest: &connectorOAuthRequestProperty{
//   		authCode: jsii.String("authCode"),
//   		redirectUri: jsii.String("redirectUri"),
//   	},
//   	refreshToken: jsii.String("refreshToken"),
//   }
//
type CfnConnectorProfile_GoogleAnalyticsConnectorProfileCredentialsProperty struct {
	// The identifier for the desired client.
	ClientId *string `json:"clientId" yaml:"clientId"`
	// The client secret used by the OAuth client to authenticate to the authorization server.
	ClientSecret *string `json:"clientSecret" yaml:"clientSecret"`
	// The credentials used to access protected Google Analytics resources.
	AccessToken *string `json:"accessToken" yaml:"accessToken"`
	// Used by select connectors for which the OAuth workflow is supported, such as Salesforce, Google Analytics, Marketo, Zendesk, and Slack.
	ConnectorOAuthRequest interface{} `json:"connectorOAuthRequest" yaml:"connectorOAuthRequest"`
	// The credentials used to acquire new access tokens.
	//
	// This is required only for OAuth2 access tokens, and is not required for OAuth1 access tokens.
	RefreshToken *string `json:"refreshToken" yaml:"refreshToken"`
}

// The `InforNexusConnectorProfileCredentials` property type specifies the connector-specific profile credentials required by Infor Nexus.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   inforNexusConnectorProfileCredentialsProperty := &inforNexusConnectorProfileCredentialsProperty{
//   	accessKeyId: jsii.String("accessKeyId"),
//   	datakey: jsii.String("datakey"),
//   	secretAccessKey: jsii.String("secretAccessKey"),
//   	userId: jsii.String("userId"),
//   }
//
type CfnConnectorProfile_InforNexusConnectorProfileCredentialsProperty struct {
	// The Access Key portion of the credentials.
	AccessKeyId *string `json:"accessKeyId" yaml:"accessKeyId"`
	// The encryption keys used to encrypt data.
	Datakey *string `json:"datakey" yaml:"datakey"`
	// The secret key used to sign requests.
	SecretAccessKey *string `json:"secretAccessKey" yaml:"secretAccessKey"`
	// The identifier for the user.
	UserId *string `json:"userId" yaml:"userId"`
}

// The `InforNexusConnectorProfileProperties` property type specifies the connector-specific profile properties required by Infor Nexus.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   inforNexusConnectorProfilePropertiesProperty := &inforNexusConnectorProfilePropertiesProperty{
//   	instanceUrl: jsii.String("instanceUrl"),
//   }
//
type CfnConnectorProfile_InforNexusConnectorProfilePropertiesProperty struct {
	// The location of the Infor Nexus resource.
	InstanceUrl *string `json:"instanceUrl" yaml:"instanceUrl"`
}

// The `MarketoConnectorProfileCredentials` property type specifies the connector-specific profile credentials required by Marketo.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   marketoConnectorProfileCredentialsProperty := &marketoConnectorProfileCredentialsProperty{
//   	clientId: jsii.String("clientId"),
//   	clientSecret: jsii.String("clientSecret"),
//
//   	// the properties below are optional
//   	accessToken: jsii.String("accessToken"),
//   	connectorOAuthRequest: &connectorOAuthRequestProperty{
//   		authCode: jsii.String("authCode"),
//   		redirectUri: jsii.String("redirectUri"),
//   	},
//   }
//
type CfnConnectorProfile_MarketoConnectorProfileCredentialsProperty struct {
	// The identifier for the desired client.
	ClientId *string `json:"clientId" yaml:"clientId"`
	// The client secret used by the OAuth client to authenticate to the authorization server.
	ClientSecret *string `json:"clientSecret" yaml:"clientSecret"`
	// The credentials used to access protected Marketo resources.
	AccessToken *string `json:"accessToken" yaml:"accessToken"`
	// Used by select connectors for which the OAuth workflow is supported, such as Salesforce, Google Analytics, Marketo, Zendesk, and Slack.
	ConnectorOAuthRequest interface{} `json:"connectorOAuthRequest" yaml:"connectorOAuthRequest"`
}

// The `MarketoConnectorProfileProperties` property type specifies the connector-specific profile properties required when using Marketo.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   marketoConnectorProfilePropertiesProperty := &marketoConnectorProfilePropertiesProperty{
//   	instanceUrl: jsii.String("instanceUrl"),
//   }
//
type CfnConnectorProfile_MarketoConnectorProfilePropertiesProperty struct {
	// The location of the Marketo resource.
	InstanceUrl *string `json:"instanceUrl" yaml:"instanceUrl"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   oAuthPropertiesProperty := &oAuthPropertiesProperty{
//   	authCodeUrl: jsii.String("authCodeUrl"),
//   	oAuthScopes: []*string{
//   		jsii.String("oAuthScopes"),
//   	},
//   	tokenUrl: jsii.String("tokenUrl"),
//   }
//
type CfnConnectorProfile_OAuthPropertiesProperty struct {
	// `CfnConnectorProfile.OAuthPropertiesProperty.AuthCodeUrl`.
	AuthCodeUrl *string `json:"authCodeUrl" yaml:"authCodeUrl"`
	// `CfnConnectorProfile.OAuthPropertiesProperty.OAuthScopes`.
	OAuthScopes *[]*string `json:"oAuthScopes" yaml:"oAuthScopes"`
	// `CfnConnectorProfile.OAuthPropertiesProperty.TokenUrl`.
	TokenUrl *string `json:"tokenUrl" yaml:"tokenUrl"`
}

// The `RedshiftConnectorProfileCredentials` property type specifies the connector-specific profile credentials required when using Amazon Redshift.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   redshiftConnectorProfileCredentialsProperty := &redshiftConnectorProfileCredentialsProperty{
//   	password: jsii.String("password"),
//   	username: jsii.String("username"),
//   }
//
type CfnConnectorProfile_RedshiftConnectorProfileCredentialsProperty struct {
	// The password that corresponds to the user name.
	Password *string `json:"password" yaml:"password"`
	// The name of the user.
	Username *string `json:"username" yaml:"username"`
}

// The `RedshiftConnectorProfileProperties` property type specifies the connector-specific profile properties when using Amazon Redshift.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   redshiftConnectorProfilePropertiesProperty := &redshiftConnectorProfilePropertiesProperty{
//   	bucketName: jsii.String("bucketName"),
//   	databaseUrl: jsii.String("databaseUrl"),
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	bucketPrefix: jsii.String("bucketPrefix"),
//   }
//
type CfnConnectorProfile_RedshiftConnectorProfilePropertiesProperty struct {
	// A name for the associated Amazon S3 bucket.
	BucketName *string `json:"bucketName" yaml:"bucketName"`
	// The JDBC URL of the Amazon Redshift cluster.
	DatabaseUrl *string `json:"databaseUrl" yaml:"databaseUrl"`
	// The Amazon Resource Name (ARN) of the IAM role.
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
	// The object key for the destination bucket in which Amazon AppFlow places the files.
	BucketPrefix *string `json:"bucketPrefix" yaml:"bucketPrefix"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//
//   var basicAuthCredentials interface{}
//   var oAuthCredentials interface{}
//   sAPODataConnectorProfileCredentialsProperty := &sAPODataConnectorProfileCredentialsProperty{
//   	basicAuthCredentials: basicAuthCredentials,
//   	oAuthCredentials: oAuthCredentials,
//   }
//
type CfnConnectorProfile_SAPODataConnectorProfileCredentialsProperty struct {
	// `CfnConnectorProfile.SAPODataConnectorProfileCredentialsProperty.BasicAuthCredentials`.
	BasicAuthCredentials interface{} `json:"basicAuthCredentials" yaml:"basicAuthCredentials"`
	// `CfnConnectorProfile.SAPODataConnectorProfileCredentialsProperty.OAuthCredentials`.
	OAuthCredentials interface{} `json:"oAuthCredentials" yaml:"oAuthCredentials"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   sAPODataConnectorProfilePropertiesProperty := &sAPODataConnectorProfilePropertiesProperty{
//   	applicationHostUrl: jsii.String("applicationHostUrl"),
//   	applicationServicePath: jsii.String("applicationServicePath"),
//   	clientNumber: jsii.String("clientNumber"),
//   	logonLanguage: jsii.String("logonLanguage"),
//   	oAuthProperties: &oAuthPropertiesProperty{
//   		authCodeUrl: jsii.String("authCodeUrl"),
//   		oAuthScopes: []*string{
//   			jsii.String("oAuthScopes"),
//   		},
//   		tokenUrl: jsii.String("tokenUrl"),
//   	},
//   	portNumber: jsii.Number(123),
//   	privateLinkServiceName: jsii.String("privateLinkServiceName"),
//   }
//
type CfnConnectorProfile_SAPODataConnectorProfilePropertiesProperty struct {
	// `CfnConnectorProfile.SAPODataConnectorProfilePropertiesProperty.ApplicationHostUrl`.
	ApplicationHostUrl *string `json:"applicationHostUrl" yaml:"applicationHostUrl"`
	// `CfnConnectorProfile.SAPODataConnectorProfilePropertiesProperty.ApplicationServicePath`.
	ApplicationServicePath *string `json:"applicationServicePath" yaml:"applicationServicePath"`
	// `CfnConnectorProfile.SAPODataConnectorProfilePropertiesProperty.ClientNumber`.
	ClientNumber *string `json:"clientNumber" yaml:"clientNumber"`
	// `CfnConnectorProfile.SAPODataConnectorProfilePropertiesProperty.LogonLanguage`.
	LogonLanguage *string `json:"logonLanguage" yaml:"logonLanguage"`
	// `CfnConnectorProfile.SAPODataConnectorProfilePropertiesProperty.OAuthProperties`.
	OAuthProperties interface{} `json:"oAuthProperties" yaml:"oAuthProperties"`
	// `CfnConnectorProfile.SAPODataConnectorProfilePropertiesProperty.PortNumber`.
	PortNumber *float64 `json:"portNumber" yaml:"portNumber"`
	// `CfnConnectorProfile.SAPODataConnectorProfilePropertiesProperty.PrivateLinkServiceName`.
	PrivateLinkServiceName *string `json:"privateLinkServiceName" yaml:"privateLinkServiceName"`
}

// The `SalesforceConnectorProfileCredentials` property type specifies the connector-specific profile credentials required when using Salesforce.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   salesforceConnectorProfileCredentialsProperty := &salesforceConnectorProfileCredentialsProperty{
//   	accessToken: jsii.String("accessToken"),
//   	clientCredentialsArn: jsii.String("clientCredentialsArn"),
//   	connectorOAuthRequest: &connectorOAuthRequestProperty{
//   		authCode: jsii.String("authCode"),
//   		redirectUri: jsii.String("redirectUri"),
//   	},
//   	refreshToken: jsii.String("refreshToken"),
//   }
//
type CfnConnectorProfile_SalesforceConnectorProfileCredentialsProperty struct {
	// The credentials used to access protected Salesforce resources.
	AccessToken *string `json:"accessToken" yaml:"accessToken"`
	// The secret manager ARN, which contains the client ID and client secret of the connected app.
	ClientCredentialsArn *string `json:"clientCredentialsArn" yaml:"clientCredentialsArn"`
	// Used by select connectors for which the OAuth workflow is supported, such as Salesforce, Google Analytics, Marketo, Zendesk, and Slack.
	ConnectorOAuthRequest interface{} `json:"connectorOAuthRequest" yaml:"connectorOAuthRequest"`
	// The credentials used to acquire new access tokens.
	RefreshToken *string `json:"refreshToken" yaml:"refreshToken"`
}

// The `SalesforceConnectorProfileProperties` property type specifies the connector-specific profile properties required when using Salesforce.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   salesforceConnectorProfilePropertiesProperty := &salesforceConnectorProfilePropertiesProperty{
//   	instanceUrl: jsii.String("instanceUrl"),
//   	isSandboxEnvironment: jsii.Boolean(false),
//   }
//
type CfnConnectorProfile_SalesforceConnectorProfilePropertiesProperty struct {
	// The location of the Salesforce resource.
	InstanceUrl *string `json:"instanceUrl" yaml:"instanceUrl"`
	// Indicates whether the connector profile applies to a sandbox or production environment.
	IsSandboxEnvironment interface{} `json:"isSandboxEnvironment" yaml:"isSandboxEnvironment"`
}

// The `ServiceNowConnectorProfileCredentials` property type specifies the connector-specific profile credentials required when using ServiceNow.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   serviceNowConnectorProfileCredentialsProperty := &serviceNowConnectorProfileCredentialsProperty{
//   	password: jsii.String("password"),
//   	username: jsii.String("username"),
//   }
//
type CfnConnectorProfile_ServiceNowConnectorProfileCredentialsProperty struct {
	// The password that corresponds to the user name.
	Password *string `json:"password" yaml:"password"`
	// The name of the user.
	Username *string `json:"username" yaml:"username"`
}

// The `ServiceNowConnectorProfileProperties` property type specifies the connector-specific profile properties required when using ServiceNow.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   serviceNowConnectorProfilePropertiesProperty := &serviceNowConnectorProfilePropertiesProperty{
//   	instanceUrl: jsii.String("instanceUrl"),
//   }
//
type CfnConnectorProfile_ServiceNowConnectorProfilePropertiesProperty struct {
	// The location of the ServiceNow resource.
	InstanceUrl *string `json:"instanceUrl" yaml:"instanceUrl"`
}

// The `SingularConnectorProfileCredentials` property type specifies the connector-specific profile credentials required when using Singular.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   singularConnectorProfileCredentialsProperty := &singularConnectorProfileCredentialsProperty{
//   	apiKey: jsii.String("apiKey"),
//   }
//
type CfnConnectorProfile_SingularConnectorProfileCredentialsProperty struct {
	// A unique alphanumeric identifier used to authenticate a user, developer, or calling program to your API.
	ApiKey *string `json:"apiKey" yaml:"apiKey"`
}

// The `SlackConnectorProfileCredentials` property type specifies the connector-specific profile credentials required when using Slack.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   slackConnectorProfileCredentialsProperty := &slackConnectorProfileCredentialsProperty{
//   	clientId: jsii.String("clientId"),
//   	clientSecret: jsii.String("clientSecret"),
//
//   	// the properties below are optional
//   	accessToken: jsii.String("accessToken"),
//   	connectorOAuthRequest: &connectorOAuthRequestProperty{
//   		authCode: jsii.String("authCode"),
//   		redirectUri: jsii.String("redirectUri"),
//   	},
//   }
//
type CfnConnectorProfile_SlackConnectorProfileCredentialsProperty struct {
	// The identifier for the client.
	ClientId *string `json:"clientId" yaml:"clientId"`
	// The client secret used by the OAuth client to authenticate to the authorization server.
	ClientSecret *string `json:"clientSecret" yaml:"clientSecret"`
	// The credentials used to access protected Slack resources.
	AccessToken *string `json:"accessToken" yaml:"accessToken"`
	// Used by select connectors for which the OAuth workflow is supported, such as Salesforce, Google Analytics, Marketo, Zendesk, and Slack.
	ConnectorOAuthRequest interface{} `json:"connectorOAuthRequest" yaml:"connectorOAuthRequest"`
}

// The `SlackConnectorProfileProperties` property type specifies the connector-specific profile properties required when using Slack.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   slackConnectorProfilePropertiesProperty := &slackConnectorProfilePropertiesProperty{
//   	instanceUrl: jsii.String("instanceUrl"),
//   }
//
type CfnConnectorProfile_SlackConnectorProfilePropertiesProperty struct {
	// The location of the Slack resource.
	InstanceUrl *string `json:"instanceUrl" yaml:"instanceUrl"`
}

// The `SnowflakeConnectorProfileCredentials` property type specifies the connector-specific profile credentials required when using Snowflake.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   snowflakeConnectorProfileCredentialsProperty := &snowflakeConnectorProfileCredentialsProperty{
//   	password: jsii.String("password"),
//   	username: jsii.String("username"),
//   }
//
type CfnConnectorProfile_SnowflakeConnectorProfileCredentialsProperty struct {
	// The password that corresponds to the user name.
	Password *string `json:"password" yaml:"password"`
	// The name of the user.
	Username *string `json:"username" yaml:"username"`
}

// The `SnowflakeConnectorProfileProperties` property type specifies the connector-specific profile properties required when using Snowflake.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   snowflakeConnectorProfilePropertiesProperty := &snowflakeConnectorProfilePropertiesProperty{
//   	bucketName: jsii.String("bucketName"),
//   	stage: jsii.String("stage"),
//   	warehouse: jsii.String("warehouse"),
//
//   	// the properties below are optional
//   	accountName: jsii.String("accountName"),
//   	bucketPrefix: jsii.String("bucketPrefix"),
//   	privateLinkServiceName: jsii.String("privateLinkServiceName"),
//   	region: jsii.String("region"),
//   }
//
type CfnConnectorProfile_SnowflakeConnectorProfilePropertiesProperty struct {
	// The name of the Amazon S3 bucket associated with Snowflake.
	BucketName *string `json:"bucketName" yaml:"bucketName"`
	// The name of the Amazon S3 stage that was created while setting up an Amazon S3 stage in the Snowflake account.
	//
	// This is written in the following format: < Database>< Schema><Stage Name>.
	Stage *string `json:"stage" yaml:"stage"`
	// The name of the Snowflake warehouse.
	Warehouse *string `json:"warehouse" yaml:"warehouse"`
	// The name of the account.
	AccountName *string `json:"accountName" yaml:"accountName"`
	// The bucket path that refers to the Amazon S3 bucket associated with Snowflake.
	BucketPrefix *string `json:"bucketPrefix" yaml:"bucketPrefix"`
	// The Snowflake Private Link service name to be used for private data transfers.
	PrivateLinkServiceName *string `json:"privateLinkServiceName" yaml:"privateLinkServiceName"`
	// The AWS Region of the Snowflake account.
	Region *string `json:"region" yaml:"region"`
}

// The `TrendmicroConnectorProfileCredentials` property type specifies the connector-specific profile credentials required when using Trend Micro.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   trendmicroConnectorProfileCredentialsProperty := &trendmicroConnectorProfileCredentialsProperty{
//   	apiSecretKey: jsii.String("apiSecretKey"),
//   }
//
type CfnConnectorProfile_TrendmicroConnectorProfileCredentialsProperty struct {
	// The Secret Access Key portion of the credentials.
	ApiSecretKey *string `json:"apiSecretKey" yaml:"apiSecretKey"`
}

// The `VeevaConnectorProfileCredentials` property type specifies the connector-specific profile credentials required when using Veeva.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   veevaConnectorProfileCredentialsProperty := &veevaConnectorProfileCredentialsProperty{
//   	password: jsii.String("password"),
//   	username: jsii.String("username"),
//   }
//
type CfnConnectorProfile_VeevaConnectorProfileCredentialsProperty struct {
	// The password that corresponds to the user name.
	Password *string `json:"password" yaml:"password"`
	// The name of the user.
	Username *string `json:"username" yaml:"username"`
}

// The `VeevaConnectorProfileProperties` property type specifies the connector-specific profile properties required when using Veeva.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   veevaConnectorProfilePropertiesProperty := &veevaConnectorProfilePropertiesProperty{
//   	instanceUrl: jsii.String("instanceUrl"),
//   }
//
type CfnConnectorProfile_VeevaConnectorProfilePropertiesProperty struct {
	// The location of the Veeva resource.
	InstanceUrl *string `json:"instanceUrl" yaml:"instanceUrl"`
}

// The `ZendeskConnectorProfileCredentials` property type specifies the connector-specific profile credentials required when using Zendesk.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   zendeskConnectorProfileCredentialsProperty := &zendeskConnectorProfileCredentialsProperty{
//   	clientId: jsii.String("clientId"),
//   	clientSecret: jsii.String("clientSecret"),
//
//   	// the properties below are optional
//   	accessToken: jsii.String("accessToken"),
//   	connectorOAuthRequest: &connectorOAuthRequestProperty{
//   		authCode: jsii.String("authCode"),
//   		redirectUri: jsii.String("redirectUri"),
//   	},
//   }
//
type CfnConnectorProfile_ZendeskConnectorProfileCredentialsProperty struct {
	// The identifier for the desired client.
	ClientId *string `json:"clientId" yaml:"clientId"`
	// The client secret used by the OAuth client to authenticate to the authorization server.
	ClientSecret *string `json:"clientSecret" yaml:"clientSecret"`
	// The credentials used to access protected Zendesk resources.
	AccessToken *string `json:"accessToken" yaml:"accessToken"`
	// Used by select connectors for which the OAuth workflow is supported, such as Salesforce, Google Analytics, Marketo, Zendesk, and Slack.
	ConnectorOAuthRequest interface{} `json:"connectorOAuthRequest" yaml:"connectorOAuthRequest"`
}

// The `ZendeskConnectorProfileProperties` property type specifies the connector-specific profile properties required when using Zendesk.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   zendeskConnectorProfilePropertiesProperty := &zendeskConnectorProfilePropertiesProperty{
//   	instanceUrl: jsii.String("instanceUrl"),
//   }
//
type CfnConnectorProfile_ZendeskConnectorProfilePropertiesProperty struct {
	// The location of the Zendesk resource.
	InstanceUrl *string `json:"instanceUrl" yaml:"instanceUrl"`
}

// Properties for defining a `CfnConnectorProfile`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//
//   var basicAuthCredentials interface{}
//   var oAuthCredentials interface{}
//   cfnConnectorProfileProps := &cfnConnectorProfileProps{
//   	connectionMode: jsii.String("connectionMode"),
//   	connectorProfileName: jsii.String("connectorProfileName"),
//   	connectorType: jsii.String("connectorType"),
//
//   	// the properties below are optional
//   	connectorProfileConfig: &connectorProfileConfigProperty{
//   		connectorProfileCredentials: &connectorProfileCredentialsProperty{
//   			amplitude: &amplitudeConnectorProfileCredentialsProperty{
//   				apiKey: jsii.String("apiKey"),
//   				secretKey: jsii.String("secretKey"),
//   			},
//   			datadog: &datadogConnectorProfileCredentialsProperty{
//   				apiKey: jsii.String("apiKey"),
//   				applicationKey: jsii.String("applicationKey"),
//   			},
//   			dynatrace: &dynatraceConnectorProfileCredentialsProperty{
//   				apiToken: jsii.String("apiToken"),
//   			},
//   			googleAnalytics: &googleAnalyticsConnectorProfileCredentialsProperty{
//   				clientId: jsii.String("clientId"),
//   				clientSecret: jsii.String("clientSecret"),
//
//   				// the properties below are optional
//   				accessToken: jsii.String("accessToken"),
//   				connectorOAuthRequest: &connectorOAuthRequestProperty{
//   					authCode: jsii.String("authCode"),
//   					redirectUri: jsii.String("redirectUri"),
//   				},
//   				refreshToken: jsii.String("refreshToken"),
//   			},
//   			inforNexus: &inforNexusConnectorProfileCredentialsProperty{
//   				accessKeyId: jsii.String("accessKeyId"),
//   				datakey: jsii.String("datakey"),
//   				secretAccessKey: jsii.String("secretAccessKey"),
//   				userId: jsii.String("userId"),
//   			},
//   			marketo: &marketoConnectorProfileCredentialsProperty{
//   				clientId: jsii.String("clientId"),
//   				clientSecret: jsii.String("clientSecret"),
//
//   				// the properties below are optional
//   				accessToken: jsii.String("accessToken"),
//   				connectorOAuthRequest: &connectorOAuthRequestProperty{
//   					authCode: jsii.String("authCode"),
//   					redirectUri: jsii.String("redirectUri"),
//   				},
//   			},
//   			redshift: &redshiftConnectorProfileCredentialsProperty{
//   				password: jsii.String("password"),
//   				username: jsii.String("username"),
//   			},
//   			salesforce: &salesforceConnectorProfileCredentialsProperty{
//   				accessToken: jsii.String("accessToken"),
//   				clientCredentialsArn: jsii.String("clientCredentialsArn"),
//   				connectorOAuthRequest: &connectorOAuthRequestProperty{
//   					authCode: jsii.String("authCode"),
//   					redirectUri: jsii.String("redirectUri"),
//   				},
//   				refreshToken: jsii.String("refreshToken"),
//   			},
//   			sapoData: &sAPODataConnectorProfileCredentialsProperty{
//   				basicAuthCredentials: basicAuthCredentials,
//   				oAuthCredentials: oAuthCredentials,
//   			},
//   			serviceNow: &serviceNowConnectorProfileCredentialsProperty{
//   				password: jsii.String("password"),
//   				username: jsii.String("username"),
//   			},
//   			singular: &singularConnectorProfileCredentialsProperty{
//   				apiKey: jsii.String("apiKey"),
//   			},
//   			slack: &slackConnectorProfileCredentialsProperty{
//   				clientId: jsii.String("clientId"),
//   				clientSecret: jsii.String("clientSecret"),
//
//   				// the properties below are optional
//   				accessToken: jsii.String("accessToken"),
//   				connectorOAuthRequest: &connectorOAuthRequestProperty{
//   					authCode: jsii.String("authCode"),
//   					redirectUri: jsii.String("redirectUri"),
//   				},
//   			},
//   			snowflake: &snowflakeConnectorProfileCredentialsProperty{
//   				password: jsii.String("password"),
//   				username: jsii.String("username"),
//   			},
//   			trendmicro: &trendmicroConnectorProfileCredentialsProperty{
//   				apiSecretKey: jsii.String("apiSecretKey"),
//   			},
//   			veeva: &veevaConnectorProfileCredentialsProperty{
//   				password: jsii.String("password"),
//   				username: jsii.String("username"),
//   			},
//   			zendesk: &zendeskConnectorProfileCredentialsProperty{
//   				clientId: jsii.String("clientId"),
//   				clientSecret: jsii.String("clientSecret"),
//
//   				// the properties below are optional
//   				accessToken: jsii.String("accessToken"),
//   				connectorOAuthRequest: &connectorOAuthRequestProperty{
//   					authCode: jsii.String("authCode"),
//   					redirectUri: jsii.String("redirectUri"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		connectorProfileProperties: &connectorProfilePropertiesProperty{
//   			datadog: &datadogConnectorProfilePropertiesProperty{
//   				instanceUrl: jsii.String("instanceUrl"),
//   			},
//   			dynatrace: &dynatraceConnectorProfilePropertiesProperty{
//   				instanceUrl: jsii.String("instanceUrl"),
//   			},
//   			inforNexus: &inforNexusConnectorProfilePropertiesProperty{
//   				instanceUrl: jsii.String("instanceUrl"),
//   			},
//   			marketo: &marketoConnectorProfilePropertiesProperty{
//   				instanceUrl: jsii.String("instanceUrl"),
//   			},
//   			redshift: &redshiftConnectorProfilePropertiesProperty{
//   				bucketName: jsii.String("bucketName"),
//   				databaseUrl: jsii.String("databaseUrl"),
//   				roleArn: jsii.String("roleArn"),
//
//   				// the properties below are optional
//   				bucketPrefix: jsii.String("bucketPrefix"),
//   			},
//   			salesforce: &salesforceConnectorProfilePropertiesProperty{
//   				instanceUrl: jsii.String("instanceUrl"),
//   				isSandboxEnvironment: jsii.Boolean(false),
//   			},
//   			sapoData: &sAPODataConnectorProfilePropertiesProperty{
//   				applicationHostUrl: jsii.String("applicationHostUrl"),
//   				applicationServicePath: jsii.String("applicationServicePath"),
//   				clientNumber: jsii.String("clientNumber"),
//   				logonLanguage: jsii.String("logonLanguage"),
//   				oAuthProperties: &oAuthPropertiesProperty{
//   					authCodeUrl: jsii.String("authCodeUrl"),
//   					oAuthScopes: []*string{
//   						jsii.String("oAuthScopes"),
//   					},
//   					tokenUrl: jsii.String("tokenUrl"),
//   				},
//   				portNumber: jsii.Number(123),
//   				privateLinkServiceName: jsii.String("privateLinkServiceName"),
//   			},
//   			serviceNow: &serviceNowConnectorProfilePropertiesProperty{
//   				instanceUrl: jsii.String("instanceUrl"),
//   			},
//   			slack: &slackConnectorProfilePropertiesProperty{
//   				instanceUrl: jsii.String("instanceUrl"),
//   			},
//   			snowflake: &snowflakeConnectorProfilePropertiesProperty{
//   				bucketName: jsii.String("bucketName"),
//   				stage: jsii.String("stage"),
//   				warehouse: jsii.String("warehouse"),
//
//   				// the properties below are optional
//   				accountName: jsii.String("accountName"),
//   				bucketPrefix: jsii.String("bucketPrefix"),
//   				privateLinkServiceName: jsii.String("privateLinkServiceName"),
//   				region: jsii.String("region"),
//   			},
//   			veeva: &veevaConnectorProfilePropertiesProperty{
//   				instanceUrl: jsii.String("instanceUrl"),
//   			},
//   			zendesk: &zendeskConnectorProfilePropertiesProperty{
//   				instanceUrl: jsii.String("instanceUrl"),
//   			},
//   		},
//   	},
//   	kmsArn: jsii.String("kmsArn"),
//   }
//
type CfnConnectorProfileProps struct {
	// Indicates the connection mode and if it is public or private.
	ConnectionMode *string `json:"connectionMode" yaml:"connectionMode"`
	// The name of the connector profile.
	//
	// The name is unique for each `ConnectorProfile` in the AWS account .
	ConnectorProfileName *string `json:"connectorProfileName" yaml:"connectorProfileName"`
	// The type of connector, such as Salesforce, Amplitude, and so on.
	ConnectorType *string `json:"connectorType" yaml:"connectorType"`
	// Defines the connector-specific configuration and credentials.
	ConnectorProfileConfig interface{} `json:"connectorProfileConfig" yaml:"connectorProfileConfig"`
	// The ARN (Amazon Resource Name) of the Key Management Service (KMS) key you provide for encryption.
	//
	// This is required if you do not want to use the Amazon AppFlow-managed KMS key. If you don't provide anything here, Amazon AppFlow uses the Amazon AppFlow-managed KMS key.
	KmsArn *string `json:"kmsArn" yaml:"kmsArn"`
}

// A CloudFormation `AWS::AppFlow::Flow`.
//
// The `AWS::AppFlow::Flow` resource is an Amazon AppFlow resource type that specifies a new flow.
//
// > If you want to use AWS CloudFormation to create a connector profile for connectors that implement OAuth (such as Salesforce, Slack, Zendesk, and Google Analytics), you must fetch the access and refresh tokens. You can do this by implementing your own UI for OAuth, or by retrieving the tokens from elsewhere. Alternatively, you can use the Amazon AppFlow console to create the connector profile, and then use that connector profile in the flow creation CloudFormation template.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   cfnFlow := appflow.NewCfnFlow(this, jsii.String("MyCfnFlow"), &cfnFlowProps{
//   	destinationFlowConfigList: []interface{}{
//   		&destinationFlowConfigProperty{
//   			connectorType: jsii.String("connectorType"),
//   			destinationConnectorProperties: &destinationConnectorPropertiesProperty{
//   				eventBridge: &eventBridgeDestinationPropertiesProperty{
//   					object: jsii.String("object"),
//
//   					// the properties below are optional
//   					errorHandlingConfig: &errorHandlingConfigProperty{
//   						bucketName: jsii.String("bucketName"),
//   						bucketPrefix: jsii.String("bucketPrefix"),
//   						failOnFirstError: jsii.Boolean(false),
//   					},
//   				},
//   				lookoutMetrics: &lookoutMetricsDestinationPropertiesProperty{
//   					object: jsii.String("object"),
//   				},
//   				marketo: &marketoDestinationPropertiesProperty{
//   					object: jsii.String("object"),
//
//   					// the properties below are optional
//   					errorHandlingConfig: &errorHandlingConfigProperty{
//   						bucketName: jsii.String("bucketName"),
//   						bucketPrefix: jsii.String("bucketPrefix"),
//   						failOnFirstError: jsii.Boolean(false),
//   					},
//   				},
//   				redshift: &redshiftDestinationPropertiesProperty{
//   					intermediateBucketName: jsii.String("intermediateBucketName"),
//   					object: jsii.String("object"),
//
//   					// the properties below are optional
//   					bucketPrefix: jsii.String("bucketPrefix"),
//   					errorHandlingConfig: &errorHandlingConfigProperty{
//   						bucketName: jsii.String("bucketName"),
//   						bucketPrefix: jsii.String("bucketPrefix"),
//   						failOnFirstError: jsii.Boolean(false),
//   					},
//   				},
//   				s3: &s3DestinationPropertiesProperty{
//   					bucketName: jsii.String("bucketName"),
//
//   					// the properties below are optional
//   					bucketPrefix: jsii.String("bucketPrefix"),
//   					s3OutputFormatConfig: &s3OutputFormatConfigProperty{
//   						aggregationConfig: &aggregationConfigProperty{
//   							aggregationType: jsii.String("aggregationType"),
//   						},
//   						fileType: jsii.String("fileType"),
//   						prefixConfig: &prefixConfigProperty{
//   							prefixFormat: jsii.String("prefixFormat"),
//   							prefixType: jsii.String("prefixType"),
//   						},
//   					},
//   				},
//   				salesforce: &salesforceDestinationPropertiesProperty{
//   					object: jsii.String("object"),
//
//   					// the properties below are optional
//   					errorHandlingConfig: &errorHandlingConfigProperty{
//   						bucketName: jsii.String("bucketName"),
//   						bucketPrefix: jsii.String("bucketPrefix"),
//   						failOnFirstError: jsii.Boolean(false),
//   					},
//   					idFieldNames: []*string{
//   						jsii.String("idFieldNames"),
//   					},
//   					writeOperationType: jsii.String("writeOperationType"),
//   				},
//   				sapoData: &sAPODataDestinationPropertiesProperty{
//   					objectPath: jsii.String("objectPath"),
//
//   					// the properties below are optional
//   					errorHandlingConfig: &errorHandlingConfigProperty{
//   						bucketName: jsii.String("bucketName"),
//   						bucketPrefix: jsii.String("bucketPrefix"),
//   						failOnFirstError: jsii.Boolean(false),
//   					},
//   					idFieldNames: []*string{
//   						jsii.String("idFieldNames"),
//   					},
//   					successResponseHandlingConfig: &successResponseHandlingConfigProperty{
//   						bucketName: jsii.String("bucketName"),
//   						bucketPrefix: jsii.String("bucketPrefix"),
//   					},
//   					writeOperationType: jsii.String("writeOperationType"),
//   				},
//   				snowflake: &snowflakeDestinationPropertiesProperty{
//   					intermediateBucketName: jsii.String("intermediateBucketName"),
//   					object: jsii.String("object"),
//
//   					// the properties below are optional
//   					bucketPrefix: jsii.String("bucketPrefix"),
//   					errorHandlingConfig: &errorHandlingConfigProperty{
//   						bucketName: jsii.String("bucketName"),
//   						bucketPrefix: jsii.String("bucketPrefix"),
//   						failOnFirstError: jsii.Boolean(false),
//   					},
//   				},
//   				upsolver: &upsolverDestinationPropertiesProperty{
//   					bucketName: jsii.String("bucketName"),
//   					s3OutputFormatConfig: &upsolverS3OutputFormatConfigProperty{
//   						prefixConfig: &prefixConfigProperty{
//   							prefixFormat: jsii.String("prefixFormat"),
//   							prefixType: jsii.String("prefixType"),
//   						},
//
//   						// the properties below are optional
//   						aggregationConfig: &aggregationConfigProperty{
//   							aggregationType: jsii.String("aggregationType"),
//   						},
//   						fileType: jsii.String("fileType"),
//   					},
//
//   					// the properties below are optional
//   					bucketPrefix: jsii.String("bucketPrefix"),
//   				},
//   				zendesk: &zendeskDestinationPropertiesProperty{
//   					object: jsii.String("object"),
//
//   					// the properties below are optional
//   					errorHandlingConfig: &errorHandlingConfigProperty{
//   						bucketName: jsii.String("bucketName"),
//   						bucketPrefix: jsii.String("bucketPrefix"),
//   						failOnFirstError: jsii.Boolean(false),
//   					},
//   					idFieldNames: []*string{
//   						jsii.String("idFieldNames"),
//   					},
//   					writeOperationType: jsii.String("writeOperationType"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			connectorProfileName: jsii.String("connectorProfileName"),
//   		},
//   	},
//   	flowName: jsii.String("flowName"),
//   	sourceFlowConfig: &sourceFlowConfigProperty{
//   		connectorType: jsii.String("connectorType"),
//   		sourceConnectorProperties: &sourceConnectorPropertiesProperty{
//   			amplitude: &amplitudeSourcePropertiesProperty{
//   				object: jsii.String("object"),
//   			},
//   			datadog: &datadogSourcePropertiesProperty{
//   				object: jsii.String("object"),
//   			},
//   			dynatrace: &dynatraceSourcePropertiesProperty{
//   				object: jsii.String("object"),
//   			},
//   			googleAnalytics: &googleAnalyticsSourcePropertiesProperty{
//   				object: jsii.String("object"),
//   			},
//   			inforNexus: &inforNexusSourcePropertiesProperty{
//   				object: jsii.String("object"),
//   			},
//   			marketo: &marketoSourcePropertiesProperty{
//   				object: jsii.String("object"),
//   			},
//   			s3: &s3SourcePropertiesProperty{
//   				bucketName: jsii.String("bucketName"),
//   				bucketPrefix: jsii.String("bucketPrefix"),
//
//   				// the properties below are optional
//   				s3InputFormatConfig: &s3InputFormatConfigProperty{
//   					s3InputFileType: jsii.String("s3InputFileType"),
//   				},
//   			},
//   			salesforce: &salesforceSourcePropertiesProperty{
//   				object: jsii.String("object"),
//
//   				// the properties below are optional
//   				enableDynamicFieldUpdate: jsii.Boolean(false),
//   				includeDeletedRecords: jsii.Boolean(false),
//   			},
//   			sapoData: &sAPODataSourcePropertiesProperty{
//   				objectPath: jsii.String("objectPath"),
//   			},
//   			serviceNow: &serviceNowSourcePropertiesProperty{
//   				object: jsii.String("object"),
//   			},
//   			singular: &singularSourcePropertiesProperty{
//   				object: jsii.String("object"),
//   			},
//   			slack: &slackSourcePropertiesProperty{
//   				object: jsii.String("object"),
//   			},
//   			trendmicro: &trendmicroSourcePropertiesProperty{
//   				object: jsii.String("object"),
//   			},
//   			veeva: &veevaSourcePropertiesProperty{
//   				object: jsii.String("object"),
//
//   				// the properties below are optional
//   				documentType: jsii.String("documentType"),
//   				includeAllVersions: jsii.Boolean(false),
//   				includeRenditions: jsii.Boolean(false),
//   				includeSourceFiles: jsii.Boolean(false),
//   			},
//   			zendesk: &zendeskSourcePropertiesProperty{
//   				object: jsii.String("object"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		connectorProfileName: jsii.String("connectorProfileName"),
//   		incrementalPullConfig: &incrementalPullConfigProperty{
//   			datetimeTypeFieldName: jsii.String("datetimeTypeFieldName"),
//   		},
//   	},
//   	tasks: []interface{}{
//   		&taskProperty{
//   			sourceFields: []*string{
//   				jsii.String("sourceFields"),
//   			},
//   			taskType: jsii.String("taskType"),
//
//   			// the properties below are optional
//   			connectorOperator: &connectorOperatorProperty{
//   				amplitude: jsii.String("amplitude"),
//   				datadog: jsii.String("datadog"),
//   				dynatrace: jsii.String("dynatrace"),
//   				googleAnalytics: jsii.String("googleAnalytics"),
//   				inforNexus: jsii.String("inforNexus"),
//   				marketo: jsii.String("marketo"),
//   				s3: jsii.String("s3"),
//   				salesforce: jsii.String("salesforce"),
//   				sapoData: jsii.String("sapoData"),
//   				serviceNow: jsii.String("serviceNow"),
//   				singular: jsii.String("singular"),
//   				slack: jsii.String("slack"),
//   				trendmicro: jsii.String("trendmicro"),
//   				veeva: jsii.String("veeva"),
//   				zendesk: jsii.String("zendesk"),
//   			},
//   			destinationField: jsii.String("destinationField"),
//   			taskProperties: []interface{}{
//   				&taskPropertiesObjectProperty{
//   					key: jsii.String("key"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	triggerConfig: &triggerConfigProperty{
//   		triggerType: jsii.String("triggerType"),
//
//   		// the properties below are optional
//   		triggerProperties: &scheduledTriggerPropertiesProperty{
//   			scheduleExpression: jsii.String("scheduleExpression"),
//
//   			// the properties below are optional
//   			dataPullMode: jsii.String("dataPullMode"),
//   			scheduleEndTime: jsii.Number(123),
//   			scheduleOffset: jsii.Number(123),
//   			scheduleStartTime: jsii.Number(123),
//   			timeZone: jsii.String("timeZone"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	kmsArn: jsii.String("kmsArn"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnFlow interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The flow's Amazon Resource Name (ARN).
	AttrFlowArn() *string
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
	// A user-entered description of the flow.
	Description() *string
	SetDescription(val *string)
	// The configuration that controls how Amazon AppFlow places data in the destination connector.
	DestinationFlowConfigList() interface{}
	SetDestinationFlowConfigList(val interface{})
	// The specified name of the flow.
	//
	// Spaces are not allowed. Use underscores (_) or hyphens (-) only.
	FlowName() *string
	SetFlowName(val *string)
	// The ARN (Amazon Resource Name) of the Key Management Service (KMS) key you provide for encryption.
	//
	// This is required if you do not want to use the Amazon AppFlow-managed KMS key. If you don't provide anything here, Amazon AppFlow uses the Amazon AppFlow-managed KMS key.
	KmsArn() *string
	SetKmsArn(val *string)
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
	// Contains information about the configuration of the source connector used in the flow.
	SourceFlowConfig() interface{}
	SetSourceFlowConfig(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// The tags used to organize, track, or control access for your flow.
	Tags() awscdk.TagManager
	// A list of tasks that Amazon AppFlow performs while transferring the data in the flow run.
	Tasks() interface{}
	SetTasks(val interface{})
	// The trigger settings that determine how and when Amazon AppFlow runs the specified flow.
	TriggerConfig() interface{}
	SetTriggerConfig(val interface{})
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

func (c *jsiiProxy_CfnFlow) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnFlow) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnFlow) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnFlow) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnFlow) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnFlow) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnFlow) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

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

func (c *jsiiProxy_CfnFlow) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnFlow) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnFlow) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_CfnFlow) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

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

func (c *jsiiProxy_CfnFlow) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_CfnFlow) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The aggregation settings that you can use to customize the output format of your flow data.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   aggregationConfigProperty := &aggregationConfigProperty{
//   	aggregationType: jsii.String("aggregationType"),
//   }
//
type CfnFlow_AggregationConfigProperty struct {
	// Specifies whether Amazon AppFlow aggregates the flow records into a single file, or leave them unaggregated.
	AggregationType *string `json:"aggregationType" yaml:"aggregationType"`
}

// The properties that are applied when Amplitude is being used as a source.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   amplitudeSourcePropertiesProperty := &amplitudeSourcePropertiesProperty{
//   	object: jsii.String("object"),
//   }
//
type CfnFlow_AmplitudeSourcePropertiesProperty struct {
	// The object specified in the Amplitude flow source.
	Object *string `json:"object" yaml:"object"`
}

// The operation to be performed on the provided source fields.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   connectorOperatorProperty := &connectorOperatorProperty{
//   	amplitude: jsii.String("amplitude"),
//   	datadog: jsii.String("datadog"),
//   	dynatrace: jsii.String("dynatrace"),
//   	googleAnalytics: jsii.String("googleAnalytics"),
//   	inforNexus: jsii.String("inforNexus"),
//   	marketo: jsii.String("marketo"),
//   	s3: jsii.String("s3"),
//   	salesforce: jsii.String("salesforce"),
//   	sapoData: jsii.String("sapoData"),
//   	serviceNow: jsii.String("serviceNow"),
//   	singular: jsii.String("singular"),
//   	slack: jsii.String("slack"),
//   	trendmicro: jsii.String("trendmicro"),
//   	veeva: jsii.String("veeva"),
//   	zendesk: jsii.String("zendesk"),
//   }
//
type CfnFlow_ConnectorOperatorProperty struct {
	// The operation to be performed on the provided Amplitude source fields.
	Amplitude *string `json:"amplitude" yaml:"amplitude"`
	// The operation to be performed on the provided Datadog source fields.
	Datadog *string `json:"datadog" yaml:"datadog"`
	// The operation to be performed on the provided Dynatrace source fields.
	Dynatrace *string `json:"dynatrace" yaml:"dynatrace"`
	// The operation to be performed on the provided Google Analytics source fields.
	GoogleAnalytics *string `json:"googleAnalytics" yaml:"googleAnalytics"`
	// The operation to be performed on the provided Infor Nexus source fields.
	InforNexus *string `json:"inforNexus" yaml:"inforNexus"`
	// The operation to be performed on the provided Marketo source fields.
	Marketo *string `json:"marketo" yaml:"marketo"`
	// The operation to be performed on the provided Amazon S3 source fields.
	S3 *string `json:"s3" yaml:"s3"`
	// The operation to be performed on the provided Salesforce source fields.
	Salesforce *string `json:"salesforce" yaml:"salesforce"`
	// `CfnFlow.ConnectorOperatorProperty.SAPOData`.
	SapoData *string `json:"sapoData" yaml:"sapoData"`
	// The operation to be performed on the provided ServiceNow source fields.
	ServiceNow *string `json:"serviceNow" yaml:"serviceNow"`
	// The operation to be performed on the provided Singular source fields.
	Singular *string `json:"singular" yaml:"singular"`
	// The operation to be performed on the provided Slack source fields.
	Slack *string `json:"slack" yaml:"slack"`
	// The operation to be performed on the provided Trend Micro source fields.
	Trendmicro *string `json:"trendmicro" yaml:"trendmicro"`
	// The operation to be performed on the provided Veeva source fields.
	Veeva *string `json:"veeva" yaml:"veeva"`
	// The operation to be performed on the provided Zendesk source fields.
	Zendesk *string `json:"zendesk" yaml:"zendesk"`
}

// The properties that are applied when Datadog is being used as a source.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   datadogSourcePropertiesProperty := &datadogSourcePropertiesProperty{
//   	object: jsii.String("object"),
//   }
//
type CfnFlow_DatadogSourcePropertiesProperty struct {
	// The object specified in the Datadog flow source.
	Object *string `json:"object" yaml:"object"`
}

// This stores the information that is required to query a particular connector.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   destinationConnectorPropertiesProperty := &destinationConnectorPropertiesProperty{
//   	eventBridge: &eventBridgeDestinationPropertiesProperty{
//   		object: jsii.String("object"),
//
//   		// the properties below are optional
//   		errorHandlingConfig: &errorHandlingConfigProperty{
//   			bucketName: jsii.String("bucketName"),
//   			bucketPrefix: jsii.String("bucketPrefix"),
//   			failOnFirstError: jsii.Boolean(false),
//   		},
//   	},
//   	lookoutMetrics: &lookoutMetricsDestinationPropertiesProperty{
//   		object: jsii.String("object"),
//   	},
//   	marketo: &marketoDestinationPropertiesProperty{
//   		object: jsii.String("object"),
//
//   		// the properties below are optional
//   		errorHandlingConfig: &errorHandlingConfigProperty{
//   			bucketName: jsii.String("bucketName"),
//   			bucketPrefix: jsii.String("bucketPrefix"),
//   			failOnFirstError: jsii.Boolean(false),
//   		},
//   	},
//   	redshift: &redshiftDestinationPropertiesProperty{
//   		intermediateBucketName: jsii.String("intermediateBucketName"),
//   		object: jsii.String("object"),
//
//   		// the properties below are optional
//   		bucketPrefix: jsii.String("bucketPrefix"),
//   		errorHandlingConfig: &errorHandlingConfigProperty{
//   			bucketName: jsii.String("bucketName"),
//   			bucketPrefix: jsii.String("bucketPrefix"),
//   			failOnFirstError: jsii.Boolean(false),
//   		},
//   	},
//   	s3: &s3DestinationPropertiesProperty{
//   		bucketName: jsii.String("bucketName"),
//
//   		// the properties below are optional
//   		bucketPrefix: jsii.String("bucketPrefix"),
//   		s3OutputFormatConfig: &s3OutputFormatConfigProperty{
//   			aggregationConfig: &aggregationConfigProperty{
//   				aggregationType: jsii.String("aggregationType"),
//   			},
//   			fileType: jsii.String("fileType"),
//   			prefixConfig: &prefixConfigProperty{
//   				prefixFormat: jsii.String("prefixFormat"),
//   				prefixType: jsii.String("prefixType"),
//   			},
//   		},
//   	},
//   	salesforce: &salesforceDestinationPropertiesProperty{
//   		object: jsii.String("object"),
//
//   		// the properties below are optional
//   		errorHandlingConfig: &errorHandlingConfigProperty{
//   			bucketName: jsii.String("bucketName"),
//   			bucketPrefix: jsii.String("bucketPrefix"),
//   			failOnFirstError: jsii.Boolean(false),
//   		},
//   		idFieldNames: []*string{
//   			jsii.String("idFieldNames"),
//   		},
//   		writeOperationType: jsii.String("writeOperationType"),
//   	},
//   	sapoData: &sAPODataDestinationPropertiesProperty{
//   		objectPath: jsii.String("objectPath"),
//
//   		// the properties below are optional
//   		errorHandlingConfig: &errorHandlingConfigProperty{
//   			bucketName: jsii.String("bucketName"),
//   			bucketPrefix: jsii.String("bucketPrefix"),
//   			failOnFirstError: jsii.Boolean(false),
//   		},
//   		idFieldNames: []*string{
//   			jsii.String("idFieldNames"),
//   		},
//   		successResponseHandlingConfig: &successResponseHandlingConfigProperty{
//   			bucketName: jsii.String("bucketName"),
//   			bucketPrefix: jsii.String("bucketPrefix"),
//   		},
//   		writeOperationType: jsii.String("writeOperationType"),
//   	},
//   	snowflake: &snowflakeDestinationPropertiesProperty{
//   		intermediateBucketName: jsii.String("intermediateBucketName"),
//   		object: jsii.String("object"),
//
//   		// the properties below are optional
//   		bucketPrefix: jsii.String("bucketPrefix"),
//   		errorHandlingConfig: &errorHandlingConfigProperty{
//   			bucketName: jsii.String("bucketName"),
//   			bucketPrefix: jsii.String("bucketPrefix"),
//   			failOnFirstError: jsii.Boolean(false),
//   		},
//   	},
//   	upsolver: &upsolverDestinationPropertiesProperty{
//   		bucketName: jsii.String("bucketName"),
//   		s3OutputFormatConfig: &upsolverS3OutputFormatConfigProperty{
//   			prefixConfig: &prefixConfigProperty{
//   				prefixFormat: jsii.String("prefixFormat"),
//   				prefixType: jsii.String("prefixType"),
//   			},
//
//   			// the properties below are optional
//   			aggregationConfig: &aggregationConfigProperty{
//   				aggregationType: jsii.String("aggregationType"),
//   			},
//   			fileType: jsii.String("fileType"),
//   		},
//
//   		// the properties below are optional
//   		bucketPrefix: jsii.String("bucketPrefix"),
//   	},
//   	zendesk: &zendeskDestinationPropertiesProperty{
//   		object: jsii.String("object"),
//
//   		// the properties below are optional
//   		errorHandlingConfig: &errorHandlingConfigProperty{
//   			bucketName: jsii.String("bucketName"),
//   			bucketPrefix: jsii.String("bucketPrefix"),
//   			failOnFirstError: jsii.Boolean(false),
//   		},
//   		idFieldNames: []*string{
//   			jsii.String("idFieldNames"),
//   		},
//   		writeOperationType: jsii.String("writeOperationType"),
//   	},
//   }
//
type CfnFlow_DestinationConnectorPropertiesProperty struct {
	// The properties required to query Amazon EventBridge.
	EventBridge interface{} `json:"eventBridge" yaml:"eventBridge"`
	// The properties required to query Amazon Lookout for Metrics.
	LookoutMetrics interface{} `json:"lookoutMetrics" yaml:"lookoutMetrics"`
	// The properties required to query Marketo.
	Marketo interface{} `json:"marketo" yaml:"marketo"`
	// The properties required to query Amazon Redshift.
	Redshift interface{} `json:"redshift" yaml:"redshift"`
	// The properties required to query Amazon S3.
	S3 interface{} `json:"s3" yaml:"s3"`
	// The properties required to query Salesforce.
	Salesforce interface{} `json:"salesforce" yaml:"salesforce"`
	// The properties required to query SAPOData.
	SapoData interface{} `json:"sapoData" yaml:"sapoData"`
	// The properties required to query Snowflake.
	Snowflake interface{} `json:"snowflake" yaml:"snowflake"`
	// The properties required to query Upsolver.
	Upsolver interface{} `json:"upsolver" yaml:"upsolver"`
	// `CfnFlow.DestinationConnectorPropertiesProperty.Zendesk`.
	Zendesk interface{} `json:"zendesk" yaml:"zendesk"`
}

// The `DestinationFlowConfig` property type specifies information about the configuration of destination connectors present in the flow.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   destinationFlowConfigProperty := &destinationFlowConfigProperty{
//   	connectorType: jsii.String("connectorType"),
//   	destinationConnectorProperties: &destinationConnectorPropertiesProperty{
//   		eventBridge: &eventBridgeDestinationPropertiesProperty{
//   			object: jsii.String("object"),
//
//   			// the properties below are optional
//   			errorHandlingConfig: &errorHandlingConfigProperty{
//   				bucketName: jsii.String("bucketName"),
//   				bucketPrefix: jsii.String("bucketPrefix"),
//   				failOnFirstError: jsii.Boolean(false),
//   			},
//   		},
//   		lookoutMetrics: &lookoutMetricsDestinationPropertiesProperty{
//   			object: jsii.String("object"),
//   		},
//   		marketo: &marketoDestinationPropertiesProperty{
//   			object: jsii.String("object"),
//
//   			// the properties below are optional
//   			errorHandlingConfig: &errorHandlingConfigProperty{
//   				bucketName: jsii.String("bucketName"),
//   				bucketPrefix: jsii.String("bucketPrefix"),
//   				failOnFirstError: jsii.Boolean(false),
//   			},
//   		},
//   		redshift: &redshiftDestinationPropertiesProperty{
//   			intermediateBucketName: jsii.String("intermediateBucketName"),
//   			object: jsii.String("object"),
//
//   			// the properties below are optional
//   			bucketPrefix: jsii.String("bucketPrefix"),
//   			errorHandlingConfig: &errorHandlingConfigProperty{
//   				bucketName: jsii.String("bucketName"),
//   				bucketPrefix: jsii.String("bucketPrefix"),
//   				failOnFirstError: jsii.Boolean(false),
//   			},
//   		},
//   		s3: &s3DestinationPropertiesProperty{
//   			bucketName: jsii.String("bucketName"),
//
//   			// the properties below are optional
//   			bucketPrefix: jsii.String("bucketPrefix"),
//   			s3OutputFormatConfig: &s3OutputFormatConfigProperty{
//   				aggregationConfig: &aggregationConfigProperty{
//   					aggregationType: jsii.String("aggregationType"),
//   				},
//   				fileType: jsii.String("fileType"),
//   				prefixConfig: &prefixConfigProperty{
//   					prefixFormat: jsii.String("prefixFormat"),
//   					prefixType: jsii.String("prefixType"),
//   				},
//   			},
//   		},
//   		salesforce: &salesforceDestinationPropertiesProperty{
//   			object: jsii.String("object"),
//
//   			// the properties below are optional
//   			errorHandlingConfig: &errorHandlingConfigProperty{
//   				bucketName: jsii.String("bucketName"),
//   				bucketPrefix: jsii.String("bucketPrefix"),
//   				failOnFirstError: jsii.Boolean(false),
//   			},
//   			idFieldNames: []*string{
//   				jsii.String("idFieldNames"),
//   			},
//   			writeOperationType: jsii.String("writeOperationType"),
//   		},
//   		sapoData: &sAPODataDestinationPropertiesProperty{
//   			objectPath: jsii.String("objectPath"),
//
//   			// the properties below are optional
//   			errorHandlingConfig: &errorHandlingConfigProperty{
//   				bucketName: jsii.String("bucketName"),
//   				bucketPrefix: jsii.String("bucketPrefix"),
//   				failOnFirstError: jsii.Boolean(false),
//   			},
//   			idFieldNames: []*string{
//   				jsii.String("idFieldNames"),
//   			},
//   			successResponseHandlingConfig: &successResponseHandlingConfigProperty{
//   				bucketName: jsii.String("bucketName"),
//   				bucketPrefix: jsii.String("bucketPrefix"),
//   			},
//   			writeOperationType: jsii.String("writeOperationType"),
//   		},
//   		snowflake: &snowflakeDestinationPropertiesProperty{
//   			intermediateBucketName: jsii.String("intermediateBucketName"),
//   			object: jsii.String("object"),
//
//   			// the properties below are optional
//   			bucketPrefix: jsii.String("bucketPrefix"),
//   			errorHandlingConfig: &errorHandlingConfigProperty{
//   				bucketName: jsii.String("bucketName"),
//   				bucketPrefix: jsii.String("bucketPrefix"),
//   				failOnFirstError: jsii.Boolean(false),
//   			},
//   		},
//   		upsolver: &upsolverDestinationPropertiesProperty{
//   			bucketName: jsii.String("bucketName"),
//   			s3OutputFormatConfig: &upsolverS3OutputFormatConfigProperty{
//   				prefixConfig: &prefixConfigProperty{
//   					prefixFormat: jsii.String("prefixFormat"),
//   					prefixType: jsii.String("prefixType"),
//   				},
//
//   				// the properties below are optional
//   				aggregationConfig: &aggregationConfigProperty{
//   					aggregationType: jsii.String("aggregationType"),
//   				},
//   				fileType: jsii.String("fileType"),
//   			},
//
//   			// the properties below are optional
//   			bucketPrefix: jsii.String("bucketPrefix"),
//   		},
//   		zendesk: &zendeskDestinationPropertiesProperty{
//   			object: jsii.String("object"),
//
//   			// the properties below are optional
//   			errorHandlingConfig: &errorHandlingConfigProperty{
//   				bucketName: jsii.String("bucketName"),
//   				bucketPrefix: jsii.String("bucketPrefix"),
//   				failOnFirstError: jsii.Boolean(false),
//   			},
//   			idFieldNames: []*string{
//   				jsii.String("idFieldNames"),
//   			},
//   			writeOperationType: jsii.String("writeOperationType"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	connectorProfileName: jsii.String("connectorProfileName"),
//   }
//
type CfnFlow_DestinationFlowConfigProperty struct {
	// The type of destination connector, such as Sales force, Amazon S3, and so on.
	//
	// *Allowed Values* : `EventBridge | Redshift | S3 | Salesforce | Snowflake`.
	ConnectorType *string `json:"connectorType" yaml:"connectorType"`
	// This stores the information that is required to query a particular connector.
	DestinationConnectorProperties interface{} `json:"destinationConnectorProperties" yaml:"destinationConnectorProperties"`
	// The name of the connector profile.
	//
	// This name must be unique for each connector profile in the AWS account .
	ConnectorProfileName *string `json:"connectorProfileName" yaml:"connectorProfileName"`
}

// The properties that are applied when Dynatrace is being used as a source.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   dynatraceSourcePropertiesProperty := &dynatraceSourcePropertiesProperty{
//   	object: jsii.String("object"),
//   }
//
type CfnFlow_DynatraceSourcePropertiesProperty struct {
	// The object specified in the Dynatrace flow source.
	Object *string `json:"object" yaml:"object"`
}

// The settings that determine how Amazon AppFlow handles an error when placing data in the destination.
//
// For example, this setting would determine if the flow should fail after one insertion error, or continue and attempt to insert every record regardless of the initial failure. `ErrorHandlingConfig` is a part of the destination connector details.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   errorHandlingConfigProperty := &errorHandlingConfigProperty{
//   	bucketName: jsii.String("bucketName"),
//   	bucketPrefix: jsii.String("bucketPrefix"),
//   	failOnFirstError: jsii.Boolean(false),
//   }
//
type CfnFlow_ErrorHandlingConfigProperty struct {
	// Specifies the name of the Amazon S3 bucket.
	BucketName *string `json:"bucketName" yaml:"bucketName"`
	// Specifies the Amazon S3 bucket prefix.
	BucketPrefix *string `json:"bucketPrefix" yaml:"bucketPrefix"`
	// Specifies if the flow should fail after the first instance of a failure when attempting to place data in the destination.
	FailOnFirstError interface{} `json:"failOnFirstError" yaml:"failOnFirstError"`
}

// The properties that are applied when Amazon EventBridge is being used as a destination.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   eventBridgeDestinationPropertiesProperty := &eventBridgeDestinationPropertiesProperty{
//   	object: jsii.String("object"),
//
//   	// the properties below are optional
//   	errorHandlingConfig: &errorHandlingConfigProperty{
//   		bucketName: jsii.String("bucketName"),
//   		bucketPrefix: jsii.String("bucketPrefix"),
//   		failOnFirstError: jsii.Boolean(false),
//   	},
//   }
//
type CfnFlow_EventBridgeDestinationPropertiesProperty struct {
	// The object specified in the Amazon EventBridge flow destination.
	Object *string `json:"object" yaml:"object"`
	// The object specified in the Amplitude flow source.
	ErrorHandlingConfig interface{} `json:"errorHandlingConfig" yaml:"errorHandlingConfig"`
}

// The properties that are applied when Google Analytics is being used as a source.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   googleAnalyticsSourcePropertiesProperty := &googleAnalyticsSourcePropertiesProperty{
//   	object: jsii.String("object"),
//   }
//
type CfnFlow_GoogleAnalyticsSourcePropertiesProperty struct {
	// The object specified in the Google Analytics flow source.
	Object *string `json:"object" yaml:"object"`
}

// Specifies the configuration used when importing incremental records from the source.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   incrementalPullConfigProperty := &incrementalPullConfigProperty{
//   	datetimeTypeFieldName: jsii.String("datetimeTypeFieldName"),
//   }
//
type CfnFlow_IncrementalPullConfigProperty struct {
	// A field that specifies the date time or timestamp field as the criteria to use when importing incremental records from the source.
	DatetimeTypeFieldName *string `json:"datetimeTypeFieldName" yaml:"datetimeTypeFieldName"`
}

// The properties that are applied when Infor Nexus is being used as a source.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   inforNexusSourcePropertiesProperty := &inforNexusSourcePropertiesProperty{
//   	object: jsii.String("object"),
//   }
//
type CfnFlow_InforNexusSourcePropertiesProperty struct {
	// The object specified in the Infor Nexus flow source.
	Object *string `json:"object" yaml:"object"`
}

// The properties that are applied when Amazon Lookout for Metrics is used as a destination.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   lookoutMetricsDestinationPropertiesProperty := &lookoutMetricsDestinationPropertiesProperty{
//   	object: jsii.String("object"),
//   }
//
type CfnFlow_LookoutMetricsDestinationPropertiesProperty struct {
	// The object specified in the Amazon Lookout for Metrics flow destination.
	Object *string `json:"object" yaml:"object"`
}

// The properties that Amazon AppFlow applies when you use Marketo as a flow destination.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   marketoDestinationPropertiesProperty := &marketoDestinationPropertiesProperty{
//   	object: jsii.String("object"),
//
//   	// the properties below are optional
//   	errorHandlingConfig: &errorHandlingConfigProperty{
//   		bucketName: jsii.String("bucketName"),
//   		bucketPrefix: jsii.String("bucketPrefix"),
//   		failOnFirstError: jsii.Boolean(false),
//   	},
//   }
//
type CfnFlow_MarketoDestinationPropertiesProperty struct {
	// The object specified in the Marketo flow destination.
	Object *string `json:"object" yaml:"object"`
	// `CfnFlow.MarketoDestinationPropertiesProperty.ErrorHandlingConfig`.
	ErrorHandlingConfig interface{} `json:"errorHandlingConfig" yaml:"errorHandlingConfig"`
}

// The properties that are applied when Marketo is being used as a source.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   marketoSourcePropertiesProperty := &marketoSourcePropertiesProperty{
//   	object: jsii.String("object"),
//   }
//
type CfnFlow_MarketoSourcePropertiesProperty struct {
	// The object specified in the Marketo flow source.
	Object *string `json:"object" yaml:"object"`
}

// Determines the prefix that Amazon AppFlow applies to the destination folder name.
//
// You can name your destination folders according to the flow frequency and date.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   prefixConfigProperty := &prefixConfigProperty{
//   	prefixFormat: jsii.String("prefixFormat"),
//   	prefixType: jsii.String("prefixType"),
//   }
//
type CfnFlow_PrefixConfigProperty struct {
	// Determines the level of granularity that's included in the prefix.
	PrefixFormat *string `json:"prefixFormat" yaml:"prefixFormat"`
	// Determines the format of the prefix, and whether it applies to the file name, file path, or both.
	PrefixType *string `json:"prefixType" yaml:"prefixType"`
}

// The properties that are applied when Amazon Redshift is being used as a destination.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   redshiftDestinationPropertiesProperty := &redshiftDestinationPropertiesProperty{
//   	intermediateBucketName: jsii.String("intermediateBucketName"),
//   	object: jsii.String("object"),
//
//   	// the properties below are optional
//   	bucketPrefix: jsii.String("bucketPrefix"),
//   	errorHandlingConfig: &errorHandlingConfigProperty{
//   		bucketName: jsii.String("bucketName"),
//   		bucketPrefix: jsii.String("bucketPrefix"),
//   		failOnFirstError: jsii.Boolean(false),
//   	},
//   }
//
type CfnFlow_RedshiftDestinationPropertiesProperty struct {
	// The intermediate bucket that Amazon AppFlow uses when moving data into Amazon Redshift.
	IntermediateBucketName *string `json:"intermediateBucketName" yaml:"intermediateBucketName"`
	// The object specified in the Amazon Redshift flow destination.
	Object *string `json:"object" yaml:"object"`
	// The object key for the bucket in which Amazon AppFlow places the destination files.
	BucketPrefix *string `json:"bucketPrefix" yaml:"bucketPrefix"`
	// The settings that determine how Amazon AppFlow handles an error when placing data in the Amazon Redshift destination.
	//
	// For example, this setting would determine if the flow should fail after one insertion error, or continue and attempt to insert every record regardless of the initial failure. `ErrorHandlingConfig` is a part of the destination connector details.
	ErrorHandlingConfig interface{} `json:"errorHandlingConfig" yaml:"errorHandlingConfig"`
}

// The properties that are applied when Amazon S3 is used as a destination.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   s3DestinationPropertiesProperty := &s3DestinationPropertiesProperty{
//   	bucketName: jsii.String("bucketName"),
//
//   	// the properties below are optional
//   	bucketPrefix: jsii.String("bucketPrefix"),
//   	s3OutputFormatConfig: &s3OutputFormatConfigProperty{
//   		aggregationConfig: &aggregationConfigProperty{
//   			aggregationType: jsii.String("aggregationType"),
//   		},
//   		fileType: jsii.String("fileType"),
//   		prefixConfig: &prefixConfigProperty{
//   			prefixFormat: jsii.String("prefixFormat"),
//   			prefixType: jsii.String("prefixType"),
//   		},
//   	},
//   }
//
type CfnFlow_S3DestinationPropertiesProperty struct {
	// The Amazon S3 bucket name in which Amazon AppFlow places the transferred data.
	BucketName *string `json:"bucketName" yaml:"bucketName"`
	// The object key for the destination bucket in which Amazon AppFlow places the files.
	BucketPrefix *string `json:"bucketPrefix" yaml:"bucketPrefix"`
	// The configuration that determines how Amazon AppFlow should format the flow output data when Amazon S3 is used as the destination.
	S3OutputFormatConfig interface{} `json:"s3OutputFormatConfig" yaml:"s3OutputFormatConfig"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   s3InputFormatConfigProperty := &s3InputFormatConfigProperty{
//   	s3InputFileType: jsii.String("s3InputFileType"),
//   }
//
type CfnFlow_S3InputFormatConfigProperty struct {
	// `CfnFlow.S3InputFormatConfigProperty.S3InputFileType`.
	S3InputFileType *string `json:"s3InputFileType" yaml:"s3InputFileType"`
}

// The configuration that determines how Amazon AppFlow should format the flow output data when Amazon S3 is used as the destination.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   s3OutputFormatConfigProperty := &s3OutputFormatConfigProperty{
//   	aggregationConfig: &aggregationConfigProperty{
//   		aggregationType: jsii.String("aggregationType"),
//   	},
//   	fileType: jsii.String("fileType"),
//   	prefixConfig: &prefixConfigProperty{
//   		prefixFormat: jsii.String("prefixFormat"),
//   		prefixType: jsii.String("prefixType"),
//   	},
//   }
//
type CfnFlow_S3OutputFormatConfigProperty struct {
	// The aggregation settings that you can use to customize the output format of your flow data.
	AggregationConfig interface{} `json:"aggregationConfig" yaml:"aggregationConfig"`
	// Indicates the file type that Amazon AppFlow places in the Amazon S3 bucket.
	FileType *string `json:"fileType" yaml:"fileType"`
	// Determines the prefix that Amazon AppFlow applies to the folder name in the Amazon S3 bucket.
	//
	// You can name folders according to the flow frequency and date.
	PrefixConfig interface{} `json:"prefixConfig" yaml:"prefixConfig"`
}

// The properties that are applied when Amazon S3 is being used as the flow source.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   s3SourcePropertiesProperty := &s3SourcePropertiesProperty{
//   	bucketName: jsii.String("bucketName"),
//   	bucketPrefix: jsii.String("bucketPrefix"),
//
//   	// the properties below are optional
//   	s3InputFormatConfig: &s3InputFormatConfigProperty{
//   		s3InputFileType: jsii.String("s3InputFileType"),
//   	},
//   }
//
type CfnFlow_S3SourcePropertiesProperty struct {
	// The Amazon S3 bucket name where the source files are stored.
	BucketName *string `json:"bucketName" yaml:"bucketName"`
	// The object key for the Amazon S3 bucket in which the source files are stored.
	BucketPrefix *string `json:"bucketPrefix" yaml:"bucketPrefix"`
	// `CfnFlow.S3SourcePropertiesProperty.S3InputFormatConfig`.
	S3InputFormatConfig interface{} `json:"s3InputFormatConfig" yaml:"s3InputFormatConfig"`
}

// The properties that are applied when using SAPOData as a flow destination.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   sAPODataDestinationPropertiesProperty := &sAPODataDestinationPropertiesProperty{
//   	objectPath: jsii.String("objectPath"),
//
//   	// the properties below are optional
//   	errorHandlingConfig: &errorHandlingConfigProperty{
//   		bucketName: jsii.String("bucketName"),
//   		bucketPrefix: jsii.String("bucketPrefix"),
//   		failOnFirstError: jsii.Boolean(false),
//   	},
//   	idFieldNames: []*string{
//   		jsii.String("idFieldNames"),
//   	},
//   	successResponseHandlingConfig: &successResponseHandlingConfigProperty{
//   		bucketName: jsii.String("bucketName"),
//   		bucketPrefix: jsii.String("bucketPrefix"),
//   	},
//   	writeOperationType: jsii.String("writeOperationType"),
//   }
//
type CfnFlow_SAPODataDestinationPropertiesProperty struct {
	// The object path specified in the SAPOData flow destination.
	ObjectPath *string `json:"objectPath" yaml:"objectPath"`
	// `CfnFlow.SAPODataDestinationPropertiesProperty.ErrorHandlingConfig`.
	ErrorHandlingConfig interface{} `json:"errorHandlingConfig" yaml:"errorHandlingConfig"`
	// `CfnFlow.SAPODataDestinationPropertiesProperty.IdFieldNames`.
	IdFieldNames *[]*string `json:"idFieldNames" yaml:"idFieldNames"`
	// Determines how Amazon AppFlow handles the success response that it gets from the connector after placing data.
	//
	// For example, this setting would determine where to write the response from a destination connector upon a successful insert operation.
	SuccessResponseHandlingConfig interface{} `json:"successResponseHandlingConfig" yaml:"successResponseHandlingConfig"`
	// `CfnFlow.SAPODataDestinationPropertiesProperty.WriteOperationType`.
	WriteOperationType *string `json:"writeOperationType" yaml:"writeOperationType"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   sAPODataSourcePropertiesProperty := &sAPODataSourcePropertiesProperty{
//   	objectPath: jsii.String("objectPath"),
//   }
//
type CfnFlow_SAPODataSourcePropertiesProperty struct {
	// `CfnFlow.SAPODataSourcePropertiesProperty.ObjectPath`.
	ObjectPath *string `json:"objectPath" yaml:"objectPath"`
}

// The properties that are applied when Salesforce is being used as a destination.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   salesforceDestinationPropertiesProperty := &salesforceDestinationPropertiesProperty{
//   	object: jsii.String("object"),
//
//   	// the properties below are optional
//   	errorHandlingConfig: &errorHandlingConfigProperty{
//   		bucketName: jsii.String("bucketName"),
//   		bucketPrefix: jsii.String("bucketPrefix"),
//   		failOnFirstError: jsii.Boolean(false),
//   	},
//   	idFieldNames: []*string{
//   		jsii.String("idFieldNames"),
//   	},
//   	writeOperationType: jsii.String("writeOperationType"),
//   }
//
type CfnFlow_SalesforceDestinationPropertiesProperty struct {
	// The object specified in the Salesforce flow destination.
	Object *string `json:"object" yaml:"object"`
	// The settings that determine how Amazon AppFlow handles an error when placing data in the Salesforce destination.
	//
	// For example, this setting would determine if the flow should fail after one insertion error, or continue and attempt to insert every record regardless of the initial failure. `ErrorHandlingConfig` is a part of the destination connector details.
	ErrorHandlingConfig interface{} `json:"errorHandlingConfig" yaml:"errorHandlingConfig"`
	// The name of the field that Amazon AppFlow uses as an ID when performing a write operation such as update or delete.
	IdFieldNames *[]*string `json:"idFieldNames" yaml:"idFieldNames"`
	// This specifies the type of write operation to be performed in Salesforce.
	//
	// When the value is `UPSERT` , then `idFieldNames` is required.
	WriteOperationType *string `json:"writeOperationType" yaml:"writeOperationType"`
}

// The properties that are applied when Salesforce is being used as a source.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   salesforceSourcePropertiesProperty := &salesforceSourcePropertiesProperty{
//   	object: jsii.String("object"),
//
//   	// the properties below are optional
//   	enableDynamicFieldUpdate: jsii.Boolean(false),
//   	includeDeletedRecords: jsii.Boolean(false),
//   }
//
type CfnFlow_SalesforceSourcePropertiesProperty struct {
	// The object specified in the Salesforce flow source.
	Object *string `json:"object" yaml:"object"`
	// The flag that enables dynamic fetching of new (recently added) fields in the Salesforce objects while running a flow.
	EnableDynamicFieldUpdate interface{} `json:"enableDynamicFieldUpdate" yaml:"enableDynamicFieldUpdate"`
	// Indicates whether Amazon AppFlow includes deleted files in the flow run.
	IncludeDeletedRecords interface{} `json:"includeDeletedRecords" yaml:"includeDeletedRecords"`
}

// Specifies the configuration details of a schedule-triggered flow as defined by the user.
//
// Currently, these settings only apply to the `Scheduled` trigger type.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   scheduledTriggerPropertiesProperty := &scheduledTriggerPropertiesProperty{
//   	scheduleExpression: jsii.String("scheduleExpression"),
//
//   	// the properties below are optional
//   	dataPullMode: jsii.String("dataPullMode"),
//   	scheduleEndTime: jsii.Number(123),
//   	scheduleOffset: jsii.Number(123),
//   	scheduleStartTime: jsii.Number(123),
//   	timeZone: jsii.String("timeZone"),
//   }
//
type CfnFlow_ScheduledTriggerPropertiesProperty struct {
	// The scheduling expression that determines the rate at which the scheduled flow will run, for example: `rate(5minutes)` .
	ScheduleExpression *string `json:"scheduleExpression" yaml:"scheduleExpression"`
	// Specifies whether a scheduled flow has an incremental data transfer or a complete data transfer for each flow run.
	DataPullMode *string `json:"dataPullMode" yaml:"dataPullMode"`
	// Specifies the scheduled end time for a schedule-triggered flow.
	ScheduleEndTime *float64 `json:"scheduleEndTime" yaml:"scheduleEndTime"`
	// Specifies the optional offset that is added to the time interval for a schedule-triggered flow.
	ScheduleOffset *float64 `json:"scheduleOffset" yaml:"scheduleOffset"`
	// Specifies the scheduled start time for a schedule-triggered flow.
	ScheduleStartTime *float64 `json:"scheduleStartTime" yaml:"scheduleStartTime"`
	// Specifies the time zone used when referring to the date and time of a scheduled-triggered flow, such as `America/New_York` .
	TimeZone *string `json:"timeZone" yaml:"timeZone"`
}

// The properties that are applied when ServiceNow is being used as a source.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   serviceNowSourcePropertiesProperty := &serviceNowSourcePropertiesProperty{
//   	object: jsii.String("object"),
//   }
//
type CfnFlow_ServiceNowSourcePropertiesProperty struct {
	// The object specified in the ServiceNow flow source.
	Object *string `json:"object" yaml:"object"`
}

// The properties that are applied when Singular is being used as a source.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   singularSourcePropertiesProperty := &singularSourcePropertiesProperty{
//   	object: jsii.String("object"),
//   }
//
type CfnFlow_SingularSourcePropertiesProperty struct {
	// The object specified in the Singular flow source.
	Object *string `json:"object" yaml:"object"`
}

// The properties that are applied when Slack is being used as a source.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   slackSourcePropertiesProperty := &slackSourcePropertiesProperty{
//   	object: jsii.String("object"),
//   }
//
type CfnFlow_SlackSourcePropertiesProperty struct {
	// The object specified in the Slack flow source.
	Object *string `json:"object" yaml:"object"`
}

// The properties that are applied when Snowflake is being used as a destination.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   snowflakeDestinationPropertiesProperty := &snowflakeDestinationPropertiesProperty{
//   	intermediateBucketName: jsii.String("intermediateBucketName"),
//   	object: jsii.String("object"),
//
//   	// the properties below are optional
//   	bucketPrefix: jsii.String("bucketPrefix"),
//   	errorHandlingConfig: &errorHandlingConfigProperty{
//   		bucketName: jsii.String("bucketName"),
//   		bucketPrefix: jsii.String("bucketPrefix"),
//   		failOnFirstError: jsii.Boolean(false),
//   	},
//   }
//
type CfnFlow_SnowflakeDestinationPropertiesProperty struct {
	// The intermediate bucket that Amazon AppFlow uses when moving data into Snowflake.
	IntermediateBucketName *string `json:"intermediateBucketName" yaml:"intermediateBucketName"`
	// The object specified in the Snowflake flow destination.
	Object *string `json:"object" yaml:"object"`
	// The object key for the destination bucket in which Amazon AppFlow places the files.
	BucketPrefix *string `json:"bucketPrefix" yaml:"bucketPrefix"`
	// The settings that determine how Amazon AppFlow handles an error when placing data in the Snowflake destination.
	//
	// For example, this setting would determine if the flow should fail after one insertion error, or continue and attempt to insert every record regardless of the initial failure. `ErrorHandlingConfig` is a part of the destination connector details.
	ErrorHandlingConfig interface{} `json:"errorHandlingConfig" yaml:"errorHandlingConfig"`
}

// Specifies the information that is required to query a particular connector.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   sourceConnectorPropertiesProperty := &sourceConnectorPropertiesProperty{
//   	amplitude: &amplitudeSourcePropertiesProperty{
//   		object: jsii.String("object"),
//   	},
//   	datadog: &datadogSourcePropertiesProperty{
//   		object: jsii.String("object"),
//   	},
//   	dynatrace: &dynatraceSourcePropertiesProperty{
//   		object: jsii.String("object"),
//   	},
//   	googleAnalytics: &googleAnalyticsSourcePropertiesProperty{
//   		object: jsii.String("object"),
//   	},
//   	inforNexus: &inforNexusSourcePropertiesProperty{
//   		object: jsii.String("object"),
//   	},
//   	marketo: &marketoSourcePropertiesProperty{
//   		object: jsii.String("object"),
//   	},
//   	s3: &s3SourcePropertiesProperty{
//   		bucketName: jsii.String("bucketName"),
//   		bucketPrefix: jsii.String("bucketPrefix"),
//
//   		// the properties below are optional
//   		s3InputFormatConfig: &s3InputFormatConfigProperty{
//   			s3InputFileType: jsii.String("s3InputFileType"),
//   		},
//   	},
//   	salesforce: &salesforceSourcePropertiesProperty{
//   		object: jsii.String("object"),
//
//   		// the properties below are optional
//   		enableDynamicFieldUpdate: jsii.Boolean(false),
//   		includeDeletedRecords: jsii.Boolean(false),
//   	},
//   	sapoData: &sAPODataSourcePropertiesProperty{
//   		objectPath: jsii.String("objectPath"),
//   	},
//   	serviceNow: &serviceNowSourcePropertiesProperty{
//   		object: jsii.String("object"),
//   	},
//   	singular: &singularSourcePropertiesProperty{
//   		object: jsii.String("object"),
//   	},
//   	slack: &slackSourcePropertiesProperty{
//   		object: jsii.String("object"),
//   	},
//   	trendmicro: &trendmicroSourcePropertiesProperty{
//   		object: jsii.String("object"),
//   	},
//   	veeva: &veevaSourcePropertiesProperty{
//   		object: jsii.String("object"),
//
//   		// the properties below are optional
//   		documentType: jsii.String("documentType"),
//   		includeAllVersions: jsii.Boolean(false),
//   		includeRenditions: jsii.Boolean(false),
//   		includeSourceFiles: jsii.Boolean(false),
//   	},
//   	zendesk: &zendeskSourcePropertiesProperty{
//   		object: jsii.String("object"),
//   	},
//   }
//
type CfnFlow_SourceConnectorPropertiesProperty struct {
	// Specifies the information that is required for querying Amplitude.
	Amplitude interface{} `json:"amplitude" yaml:"amplitude"`
	// Specifies the information that is required for querying Datadog.
	Datadog interface{} `json:"datadog" yaml:"datadog"`
	// Specifies the information that is required for querying Dynatrace.
	Dynatrace interface{} `json:"dynatrace" yaml:"dynatrace"`
	// Specifies the information that is required for querying Google Analytics.
	GoogleAnalytics interface{} `json:"googleAnalytics" yaml:"googleAnalytics"`
	// Specifies the information that is required for querying Infor Nexus.
	InforNexus interface{} `json:"inforNexus" yaml:"inforNexus"`
	// Specifies the information that is required for querying Marketo.
	Marketo interface{} `json:"marketo" yaml:"marketo"`
	// Specifies the information that is required for querying Amazon S3.
	S3 interface{} `json:"s3" yaml:"s3"`
	// Specifies the information that is required for querying Salesforce.
	Salesforce interface{} `json:"salesforce" yaml:"salesforce"`
	// `CfnFlow.SourceConnectorPropertiesProperty.SAPOData`.
	SapoData interface{} `json:"sapoData" yaml:"sapoData"`
	// Specifies the information that is required for querying ServiceNow.
	ServiceNow interface{} `json:"serviceNow" yaml:"serviceNow"`
	// Specifies the information that is required for querying Singular.
	Singular interface{} `json:"singular" yaml:"singular"`
	// Specifies the information that is required for querying Slack.
	Slack interface{} `json:"slack" yaml:"slack"`
	// Specifies the information that is required for querying Trend Micro.
	Trendmicro interface{} `json:"trendmicro" yaml:"trendmicro"`
	// Specifies the information that is required for querying Veeva.
	Veeva interface{} `json:"veeva" yaml:"veeva"`
	// Specifies the information that is required for querying Zendesk.
	Zendesk interface{} `json:"zendesk" yaml:"zendesk"`
}

// Contains information about the configuration of the source connector used in the flow.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   sourceFlowConfigProperty := &sourceFlowConfigProperty{
//   	connectorType: jsii.String("connectorType"),
//   	sourceConnectorProperties: &sourceConnectorPropertiesProperty{
//   		amplitude: &amplitudeSourcePropertiesProperty{
//   			object: jsii.String("object"),
//   		},
//   		datadog: &datadogSourcePropertiesProperty{
//   			object: jsii.String("object"),
//   		},
//   		dynatrace: &dynatraceSourcePropertiesProperty{
//   			object: jsii.String("object"),
//   		},
//   		googleAnalytics: &googleAnalyticsSourcePropertiesProperty{
//   			object: jsii.String("object"),
//   		},
//   		inforNexus: &inforNexusSourcePropertiesProperty{
//   			object: jsii.String("object"),
//   		},
//   		marketo: &marketoSourcePropertiesProperty{
//   			object: jsii.String("object"),
//   		},
//   		s3: &s3SourcePropertiesProperty{
//   			bucketName: jsii.String("bucketName"),
//   			bucketPrefix: jsii.String("bucketPrefix"),
//
//   			// the properties below are optional
//   			s3InputFormatConfig: &s3InputFormatConfigProperty{
//   				s3InputFileType: jsii.String("s3InputFileType"),
//   			},
//   		},
//   		salesforce: &salesforceSourcePropertiesProperty{
//   			object: jsii.String("object"),
//
//   			// the properties below are optional
//   			enableDynamicFieldUpdate: jsii.Boolean(false),
//   			includeDeletedRecords: jsii.Boolean(false),
//   		},
//   		sapoData: &sAPODataSourcePropertiesProperty{
//   			objectPath: jsii.String("objectPath"),
//   		},
//   		serviceNow: &serviceNowSourcePropertiesProperty{
//   			object: jsii.String("object"),
//   		},
//   		singular: &singularSourcePropertiesProperty{
//   			object: jsii.String("object"),
//   		},
//   		slack: &slackSourcePropertiesProperty{
//   			object: jsii.String("object"),
//   		},
//   		trendmicro: &trendmicroSourcePropertiesProperty{
//   			object: jsii.String("object"),
//   		},
//   		veeva: &veevaSourcePropertiesProperty{
//   			object: jsii.String("object"),
//
//   			// the properties below are optional
//   			documentType: jsii.String("documentType"),
//   			includeAllVersions: jsii.Boolean(false),
//   			includeRenditions: jsii.Boolean(false),
//   			includeSourceFiles: jsii.Boolean(false),
//   		},
//   		zendesk: &zendeskSourcePropertiesProperty{
//   			object: jsii.String("object"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	connectorProfileName: jsii.String("connectorProfileName"),
//   	incrementalPullConfig: &incrementalPullConfigProperty{
//   		datetimeTypeFieldName: jsii.String("datetimeTypeFieldName"),
//   	},
//   }
//
type CfnFlow_SourceFlowConfigProperty struct {
	// The type of source connector, such as Salesforce, Amplitude, and so on.
	//
	// *Allowed Values* : S3 | Amplitude | Datadog | Dynatrace | Googleanalytics | Infornexus | Salesforce | Servicenow | Singular | Slack | Trendmicro | Veeva | Zendesk.
	ConnectorType *string `json:"connectorType" yaml:"connectorType"`
	// Specifies the information that is required to query a particular source connector.
	SourceConnectorProperties interface{} `json:"sourceConnectorProperties" yaml:"sourceConnectorProperties"`
	// The name of the connector profile.
	//
	// This name must be unique for each connector profile in the AWS account .
	ConnectorProfileName *string `json:"connectorProfileName" yaml:"connectorProfileName"`
	// Defines the configuration for a scheduled incremental data pull.
	//
	// If a valid configuration is provided, the fields specified in the configuration are used when querying for the incremental data pull.
	IncrementalPullConfig interface{} `json:"incrementalPullConfig" yaml:"incrementalPullConfig"`
}

// Determines how Amazon AppFlow handles the success response that it gets from the connector after placing data.
//
// For example, this setting would determine where to write the response from the destination connector upon a successful insert operation.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   successResponseHandlingConfigProperty := &successResponseHandlingConfigProperty{
//   	bucketName: jsii.String("bucketName"),
//   	bucketPrefix: jsii.String("bucketPrefix"),
//   }
//
type CfnFlow_SuccessResponseHandlingConfigProperty struct {
	// The name of the Amazon S3 bucket.
	BucketName *string `json:"bucketName" yaml:"bucketName"`
	// The Amazon S3 bucket prefix.
	BucketPrefix *string `json:"bucketPrefix" yaml:"bucketPrefix"`
}

// A map used to store task-related information.
//
// The execution service looks for particular information based on the `TaskType` .
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   taskPropertiesObjectProperty := &taskPropertiesObjectProperty{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
//
type CfnFlow_TaskPropertiesObjectProperty struct {
	// The task property key.
	//
	// *Allowed Values* : `VALUE | VALUES | DATA_TYPE | UPPER_BOUND | LOWER_BOUND | SOURCE_DATA_TYPE | DESTINATION_DATA_TYPE | VALIDATION_ACTION | MASK_VALUE | MASK_LENGTH | TRUNCATE_LENGTH | MATH_OPERATION_FIELDS_ORDER | CONCAT_FORMAT | SUBFIELD_CATEGORY_MAP` | `EXCLUDE_SOURCE_FIELDS_LIST`.
	Key *string `json:"key" yaml:"key"`
	// The task property value.
	Value *string `json:"value" yaml:"value"`
}

// The `Task` property type specifies the class for modeling different type of tasks.
//
// Task implementation varies based on the `TaskType` .
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   taskProperty := &taskProperty{
//   	sourceFields: []*string{
//   		jsii.String("sourceFields"),
//   	},
//   	taskType: jsii.String("taskType"),
//
//   	// the properties below are optional
//   	connectorOperator: &connectorOperatorProperty{
//   		amplitude: jsii.String("amplitude"),
//   		datadog: jsii.String("datadog"),
//   		dynatrace: jsii.String("dynatrace"),
//   		googleAnalytics: jsii.String("googleAnalytics"),
//   		inforNexus: jsii.String("inforNexus"),
//   		marketo: jsii.String("marketo"),
//   		s3: jsii.String("s3"),
//   		salesforce: jsii.String("salesforce"),
//   		sapoData: jsii.String("sapoData"),
//   		serviceNow: jsii.String("serviceNow"),
//   		singular: jsii.String("singular"),
//   		slack: jsii.String("slack"),
//   		trendmicro: jsii.String("trendmicro"),
//   		veeva: jsii.String("veeva"),
//   		zendesk: jsii.String("zendesk"),
//   	},
//   	destinationField: jsii.String("destinationField"),
//   	taskProperties: []interface{}{
//   		&taskPropertiesObjectProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnFlow_TaskProperty struct {
	// The source fields to which a particular task is applied.
	SourceFields *[]*string `json:"sourceFields" yaml:"sourceFields"`
	// Specifies the particular task implementation that Amazon AppFlow performs.
	//
	// *Allowed values* : `Arithmetic` | `Filter` | `Map` | `Map_all` | `Mask` | `Merge` | `Truncate` | `Validate`.
	TaskType *string `json:"taskType" yaml:"taskType"`
	// The operation to be performed on the provided source fields.
	ConnectorOperator interface{} `json:"connectorOperator" yaml:"connectorOperator"`
	// A field in a destination connector, or a field value against which Amazon AppFlow validates a source field.
	DestinationField *string `json:"destinationField" yaml:"destinationField"`
	// A map used to store task-related information.
	//
	// The execution service looks for particular information based on the `TaskType` .
	TaskProperties interface{} `json:"taskProperties" yaml:"taskProperties"`
}

// The properties that are applied when using Trend Micro as a flow source.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   trendmicroSourcePropertiesProperty := &trendmicroSourcePropertiesProperty{
//   	object: jsii.String("object"),
//   }
//
type CfnFlow_TrendmicroSourcePropertiesProperty struct {
	// The object specified in the Trend Micro flow source.
	Object *string `json:"object" yaml:"object"`
}

// The trigger settings that determine how and when Amazon AppFlow runs the specified flow.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   triggerConfigProperty := &triggerConfigProperty{
//   	triggerType: jsii.String("triggerType"),
//
//   	// the properties below are optional
//   	triggerProperties: &scheduledTriggerPropertiesProperty{
//   		scheduleExpression: jsii.String("scheduleExpression"),
//
//   		// the properties below are optional
//   		dataPullMode: jsii.String("dataPullMode"),
//   		scheduleEndTime: jsii.Number(123),
//   		scheduleOffset: jsii.Number(123),
//   		scheduleStartTime: jsii.Number(123),
//   		timeZone: jsii.String("timeZone"),
//   	},
//   }
//
type CfnFlow_TriggerConfigProperty struct {
	// Specifies the type of flow trigger.
	//
	// This can be `OnDemand` , `Scheduled` , or `Event` .
	TriggerType *string `json:"triggerType" yaml:"triggerType"`
	// Specifies the configuration details of a schedule-triggered flow as defined by the user.
	//
	// Currently, these settings only apply to the `Scheduled` trigger type.
	TriggerProperties interface{} `json:"triggerProperties" yaml:"triggerProperties"`
}

// The properties that are applied when Upsolver is used as a destination.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   upsolverDestinationPropertiesProperty := &upsolverDestinationPropertiesProperty{
//   	bucketName: jsii.String("bucketName"),
//   	s3OutputFormatConfig: &upsolverS3OutputFormatConfigProperty{
//   		prefixConfig: &prefixConfigProperty{
//   			prefixFormat: jsii.String("prefixFormat"),
//   			prefixType: jsii.String("prefixType"),
//   		},
//
//   		// the properties below are optional
//   		aggregationConfig: &aggregationConfigProperty{
//   			aggregationType: jsii.String("aggregationType"),
//   		},
//   		fileType: jsii.String("fileType"),
//   	},
//
//   	// the properties below are optional
//   	bucketPrefix: jsii.String("bucketPrefix"),
//   }
//
type CfnFlow_UpsolverDestinationPropertiesProperty struct {
	// The Upsolver Amazon S3 bucket name in which Amazon AppFlow places the transferred data.
	BucketName *string `json:"bucketName" yaml:"bucketName"`
	// The configuration that determines how data is formatted when Upsolver is used as the flow destination.
	S3OutputFormatConfig interface{} `json:"s3OutputFormatConfig" yaml:"s3OutputFormatConfig"`
	// The object key for the destination Upsolver Amazon S3 bucket in which Amazon AppFlow places the files.
	BucketPrefix *string `json:"bucketPrefix" yaml:"bucketPrefix"`
}

// The configuration that determines how Amazon AppFlow formats the flow output data when Upsolver is used as the destination.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   upsolverS3OutputFormatConfigProperty := &upsolverS3OutputFormatConfigProperty{
//   	prefixConfig: &prefixConfigProperty{
//   		prefixFormat: jsii.String("prefixFormat"),
//   		prefixType: jsii.String("prefixType"),
//   	},
//
//   	// the properties below are optional
//   	aggregationConfig: &aggregationConfigProperty{
//   		aggregationType: jsii.String("aggregationType"),
//   	},
//   	fileType: jsii.String("fileType"),
//   }
//
type CfnFlow_UpsolverS3OutputFormatConfigProperty struct {
	// Determines the prefix that Amazon AppFlow applies to the destination folder name.
	//
	// You can name your destination folders according to the flow frequency and date.
	PrefixConfig interface{} `json:"prefixConfig" yaml:"prefixConfig"`
	// The aggregation settings that you can use to customize the output format of your flow data.
	AggregationConfig interface{} `json:"aggregationConfig" yaml:"aggregationConfig"`
	// Indicates the file type that Amazon AppFlow places in the Upsolver Amazon S3 bucket.
	FileType *string `json:"fileType" yaml:"fileType"`
}

// The properties that are applied when using Veeva as a flow source.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   veevaSourcePropertiesProperty := &veevaSourcePropertiesProperty{
//   	object: jsii.String("object"),
//
//   	// the properties below are optional
//   	documentType: jsii.String("documentType"),
//   	includeAllVersions: jsii.Boolean(false),
//   	includeRenditions: jsii.Boolean(false),
//   	includeSourceFiles: jsii.Boolean(false),
//   }
//
type CfnFlow_VeevaSourcePropertiesProperty struct {
	// The object specified in the Veeva flow source.
	Object *string `json:"object" yaml:"object"`
	// `CfnFlow.VeevaSourcePropertiesProperty.DocumentType`.
	DocumentType *string `json:"documentType" yaml:"documentType"`
	// `CfnFlow.VeevaSourcePropertiesProperty.IncludeAllVersions`.
	IncludeAllVersions interface{} `json:"includeAllVersions" yaml:"includeAllVersions"`
	// `CfnFlow.VeevaSourcePropertiesProperty.IncludeRenditions`.
	IncludeRenditions interface{} `json:"includeRenditions" yaml:"includeRenditions"`
	// `CfnFlow.VeevaSourcePropertiesProperty.IncludeSourceFiles`.
	IncludeSourceFiles interface{} `json:"includeSourceFiles" yaml:"includeSourceFiles"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   zendeskDestinationPropertiesProperty := &zendeskDestinationPropertiesProperty{
//   	object: jsii.String("object"),
//
//   	// the properties below are optional
//   	errorHandlingConfig: &errorHandlingConfigProperty{
//   		bucketName: jsii.String("bucketName"),
//   		bucketPrefix: jsii.String("bucketPrefix"),
//   		failOnFirstError: jsii.Boolean(false),
//   	},
//   	idFieldNames: []*string{
//   		jsii.String("idFieldNames"),
//   	},
//   	writeOperationType: jsii.String("writeOperationType"),
//   }
//
type CfnFlow_ZendeskDestinationPropertiesProperty struct {
	// `CfnFlow.ZendeskDestinationPropertiesProperty.Object`.
	Object *string `json:"object" yaml:"object"`
	// `CfnFlow.ZendeskDestinationPropertiesProperty.ErrorHandlingConfig`.
	ErrorHandlingConfig interface{} `json:"errorHandlingConfig" yaml:"errorHandlingConfig"`
	// `CfnFlow.ZendeskDestinationPropertiesProperty.IdFieldNames`.
	IdFieldNames *[]*string `json:"idFieldNames" yaml:"idFieldNames"`
	// `CfnFlow.ZendeskDestinationPropertiesProperty.WriteOperationType`.
	WriteOperationType *string `json:"writeOperationType" yaml:"writeOperationType"`
}

// The properties that are applied when using Zendesk as a flow source.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   zendeskSourcePropertiesProperty := &zendeskSourcePropertiesProperty{
//   	object: jsii.String("object"),
//   }
//
type CfnFlow_ZendeskSourcePropertiesProperty struct {
	// The object specified in the Zendesk flow source.
	Object *string `json:"object" yaml:"object"`
}

// Properties for defining a `CfnFlow`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import appflow "github.com/aws/aws-cdk-go/awscdk/aws_appflow"
//   cfnFlowProps := &cfnFlowProps{
//   	destinationFlowConfigList: []interface{}{
//   		&destinationFlowConfigProperty{
//   			connectorType: jsii.String("connectorType"),
//   			destinationConnectorProperties: &destinationConnectorPropertiesProperty{
//   				eventBridge: &eventBridgeDestinationPropertiesProperty{
//   					object: jsii.String("object"),
//
//   					// the properties below are optional
//   					errorHandlingConfig: &errorHandlingConfigProperty{
//   						bucketName: jsii.String("bucketName"),
//   						bucketPrefix: jsii.String("bucketPrefix"),
//   						failOnFirstError: jsii.Boolean(false),
//   					},
//   				},
//   				lookoutMetrics: &lookoutMetricsDestinationPropertiesProperty{
//   					object: jsii.String("object"),
//   				},
//   				marketo: &marketoDestinationPropertiesProperty{
//   					object: jsii.String("object"),
//
//   					// the properties below are optional
//   					errorHandlingConfig: &errorHandlingConfigProperty{
//   						bucketName: jsii.String("bucketName"),
//   						bucketPrefix: jsii.String("bucketPrefix"),
//   						failOnFirstError: jsii.Boolean(false),
//   					},
//   				},
//   				redshift: &redshiftDestinationPropertiesProperty{
//   					intermediateBucketName: jsii.String("intermediateBucketName"),
//   					object: jsii.String("object"),
//
//   					// the properties below are optional
//   					bucketPrefix: jsii.String("bucketPrefix"),
//   					errorHandlingConfig: &errorHandlingConfigProperty{
//   						bucketName: jsii.String("bucketName"),
//   						bucketPrefix: jsii.String("bucketPrefix"),
//   						failOnFirstError: jsii.Boolean(false),
//   					},
//   				},
//   				s3: &s3DestinationPropertiesProperty{
//   					bucketName: jsii.String("bucketName"),
//
//   					// the properties below are optional
//   					bucketPrefix: jsii.String("bucketPrefix"),
//   					s3OutputFormatConfig: &s3OutputFormatConfigProperty{
//   						aggregationConfig: &aggregationConfigProperty{
//   							aggregationType: jsii.String("aggregationType"),
//   						},
//   						fileType: jsii.String("fileType"),
//   						prefixConfig: &prefixConfigProperty{
//   							prefixFormat: jsii.String("prefixFormat"),
//   							prefixType: jsii.String("prefixType"),
//   						},
//   					},
//   				},
//   				salesforce: &salesforceDestinationPropertiesProperty{
//   					object: jsii.String("object"),
//
//   					// the properties below are optional
//   					errorHandlingConfig: &errorHandlingConfigProperty{
//   						bucketName: jsii.String("bucketName"),
//   						bucketPrefix: jsii.String("bucketPrefix"),
//   						failOnFirstError: jsii.Boolean(false),
//   					},
//   					idFieldNames: []*string{
//   						jsii.String("idFieldNames"),
//   					},
//   					writeOperationType: jsii.String("writeOperationType"),
//   				},
//   				sapoData: &sAPODataDestinationPropertiesProperty{
//   					objectPath: jsii.String("objectPath"),
//
//   					// the properties below are optional
//   					errorHandlingConfig: &errorHandlingConfigProperty{
//   						bucketName: jsii.String("bucketName"),
//   						bucketPrefix: jsii.String("bucketPrefix"),
//   						failOnFirstError: jsii.Boolean(false),
//   					},
//   					idFieldNames: []*string{
//   						jsii.String("idFieldNames"),
//   					},
//   					successResponseHandlingConfig: &successResponseHandlingConfigProperty{
//   						bucketName: jsii.String("bucketName"),
//   						bucketPrefix: jsii.String("bucketPrefix"),
//   					},
//   					writeOperationType: jsii.String("writeOperationType"),
//   				},
//   				snowflake: &snowflakeDestinationPropertiesProperty{
//   					intermediateBucketName: jsii.String("intermediateBucketName"),
//   					object: jsii.String("object"),
//
//   					// the properties below are optional
//   					bucketPrefix: jsii.String("bucketPrefix"),
//   					errorHandlingConfig: &errorHandlingConfigProperty{
//   						bucketName: jsii.String("bucketName"),
//   						bucketPrefix: jsii.String("bucketPrefix"),
//   						failOnFirstError: jsii.Boolean(false),
//   					},
//   				},
//   				upsolver: &upsolverDestinationPropertiesProperty{
//   					bucketName: jsii.String("bucketName"),
//   					s3OutputFormatConfig: &upsolverS3OutputFormatConfigProperty{
//   						prefixConfig: &prefixConfigProperty{
//   							prefixFormat: jsii.String("prefixFormat"),
//   							prefixType: jsii.String("prefixType"),
//   						},
//
//   						// the properties below are optional
//   						aggregationConfig: &aggregationConfigProperty{
//   							aggregationType: jsii.String("aggregationType"),
//   						},
//   						fileType: jsii.String("fileType"),
//   					},
//
//   					// the properties below are optional
//   					bucketPrefix: jsii.String("bucketPrefix"),
//   				},
//   				zendesk: &zendeskDestinationPropertiesProperty{
//   					object: jsii.String("object"),
//
//   					// the properties below are optional
//   					errorHandlingConfig: &errorHandlingConfigProperty{
//   						bucketName: jsii.String("bucketName"),
//   						bucketPrefix: jsii.String("bucketPrefix"),
//   						failOnFirstError: jsii.Boolean(false),
//   					},
//   					idFieldNames: []*string{
//   						jsii.String("idFieldNames"),
//   					},
//   					writeOperationType: jsii.String("writeOperationType"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			connectorProfileName: jsii.String("connectorProfileName"),
//   		},
//   	},
//   	flowName: jsii.String("flowName"),
//   	sourceFlowConfig: &sourceFlowConfigProperty{
//   		connectorType: jsii.String("connectorType"),
//   		sourceConnectorProperties: &sourceConnectorPropertiesProperty{
//   			amplitude: &amplitudeSourcePropertiesProperty{
//   				object: jsii.String("object"),
//   			},
//   			datadog: &datadogSourcePropertiesProperty{
//   				object: jsii.String("object"),
//   			},
//   			dynatrace: &dynatraceSourcePropertiesProperty{
//   				object: jsii.String("object"),
//   			},
//   			googleAnalytics: &googleAnalyticsSourcePropertiesProperty{
//   				object: jsii.String("object"),
//   			},
//   			inforNexus: &inforNexusSourcePropertiesProperty{
//   				object: jsii.String("object"),
//   			},
//   			marketo: &marketoSourcePropertiesProperty{
//   				object: jsii.String("object"),
//   			},
//   			s3: &s3SourcePropertiesProperty{
//   				bucketName: jsii.String("bucketName"),
//   				bucketPrefix: jsii.String("bucketPrefix"),
//
//   				// the properties below are optional
//   				s3InputFormatConfig: &s3InputFormatConfigProperty{
//   					s3InputFileType: jsii.String("s3InputFileType"),
//   				},
//   			},
//   			salesforce: &salesforceSourcePropertiesProperty{
//   				object: jsii.String("object"),
//
//   				// the properties below are optional
//   				enableDynamicFieldUpdate: jsii.Boolean(false),
//   				includeDeletedRecords: jsii.Boolean(false),
//   			},
//   			sapoData: &sAPODataSourcePropertiesProperty{
//   				objectPath: jsii.String("objectPath"),
//   			},
//   			serviceNow: &serviceNowSourcePropertiesProperty{
//   				object: jsii.String("object"),
//   			},
//   			singular: &singularSourcePropertiesProperty{
//   				object: jsii.String("object"),
//   			},
//   			slack: &slackSourcePropertiesProperty{
//   				object: jsii.String("object"),
//   			},
//   			trendmicro: &trendmicroSourcePropertiesProperty{
//   				object: jsii.String("object"),
//   			},
//   			veeva: &veevaSourcePropertiesProperty{
//   				object: jsii.String("object"),
//
//   				// the properties below are optional
//   				documentType: jsii.String("documentType"),
//   				includeAllVersions: jsii.Boolean(false),
//   				includeRenditions: jsii.Boolean(false),
//   				includeSourceFiles: jsii.Boolean(false),
//   			},
//   			zendesk: &zendeskSourcePropertiesProperty{
//   				object: jsii.String("object"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		connectorProfileName: jsii.String("connectorProfileName"),
//   		incrementalPullConfig: &incrementalPullConfigProperty{
//   			datetimeTypeFieldName: jsii.String("datetimeTypeFieldName"),
//   		},
//   	},
//   	tasks: []interface{}{
//   		&taskProperty{
//   			sourceFields: []*string{
//   				jsii.String("sourceFields"),
//   			},
//   			taskType: jsii.String("taskType"),
//
//   			// the properties below are optional
//   			connectorOperator: &connectorOperatorProperty{
//   				amplitude: jsii.String("amplitude"),
//   				datadog: jsii.String("datadog"),
//   				dynatrace: jsii.String("dynatrace"),
//   				googleAnalytics: jsii.String("googleAnalytics"),
//   				inforNexus: jsii.String("inforNexus"),
//   				marketo: jsii.String("marketo"),
//   				s3: jsii.String("s3"),
//   				salesforce: jsii.String("salesforce"),
//   				sapoData: jsii.String("sapoData"),
//   				serviceNow: jsii.String("serviceNow"),
//   				singular: jsii.String("singular"),
//   				slack: jsii.String("slack"),
//   				trendmicro: jsii.String("trendmicro"),
//   				veeva: jsii.String("veeva"),
//   				zendesk: jsii.String("zendesk"),
//   			},
//   			destinationField: jsii.String("destinationField"),
//   			taskProperties: []interface{}{
//   				&taskPropertiesObjectProperty{
//   					key: jsii.String("key"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	triggerConfig: &triggerConfigProperty{
//   		triggerType: jsii.String("triggerType"),
//
//   		// the properties below are optional
//   		triggerProperties: &scheduledTriggerPropertiesProperty{
//   			scheduleExpression: jsii.String("scheduleExpression"),
//
//   			// the properties below are optional
//   			dataPullMode: jsii.String("dataPullMode"),
//   			scheduleEndTime: jsii.Number(123),
//   			scheduleOffset: jsii.Number(123),
//   			scheduleStartTime: jsii.Number(123),
//   			timeZone: jsii.String("timeZone"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	kmsArn: jsii.String("kmsArn"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnFlowProps struct {
	// The configuration that controls how Amazon AppFlow places data in the destination connector.
	DestinationFlowConfigList interface{} `json:"destinationFlowConfigList" yaml:"destinationFlowConfigList"`
	// The specified name of the flow.
	//
	// Spaces are not allowed. Use underscores (_) or hyphens (-) only.
	FlowName *string `json:"flowName" yaml:"flowName"`
	// Contains information about the configuration of the source connector used in the flow.
	SourceFlowConfig interface{} `json:"sourceFlowConfig" yaml:"sourceFlowConfig"`
	// A list of tasks that Amazon AppFlow performs while transferring the data in the flow run.
	Tasks interface{} `json:"tasks" yaml:"tasks"`
	// The trigger settings that determine how and when Amazon AppFlow runs the specified flow.
	TriggerConfig interface{} `json:"triggerConfig" yaml:"triggerConfig"`
	// A user-entered description of the flow.
	Description *string `json:"description" yaml:"description"`
	// The ARN (Amazon Resource Name) of the Key Management Service (KMS) key you provide for encryption.
	//
	// This is required if you do not want to use the Amazon AppFlow-managed KMS key. If you don't provide anything here, Amazon AppFlow uses the Amazon AppFlow-managed KMS key.
	KmsArn *string `json:"kmsArn" yaml:"kmsArn"`
	// The tags used to organize, track, or control access for your flow.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

