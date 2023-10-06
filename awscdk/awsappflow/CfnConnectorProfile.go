package awsappflow

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappflow/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::AppFlow::ConnectorProfile` resource is an Amazon AppFlow resource type that specifies the configuration profile for an instance of a connector.
//
// This includes the provided name, credentials ARN, connection-mode, and so on. The fields that are common to all types of connector profiles are explicitly specified under the `Properties` field. The rest of the connector-specific properties are specified under `Properties/ConnectorProfileConfig` .
//
// > If you want to use AWS CloudFormation to create a connector profile for connectors that implement OAuth (such as Salesforce, Slack, Zendesk, and Google Analytics), you must fetch the access and refresh tokens. You can do this by implementing your own UI for OAuth, or by retrieving the tokens from elsewhere. Alternatively, you can use the Amazon AppFlow console to create the connector profile, and then use that connector profile in the flow creation CloudFormation template.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConnectorProfile := awscdk.Aws_appflow.NewCfnConnectorProfile(this, jsii.String("MyCfnConnectorProfile"), &CfnConnectorProfileProps{
//   	ConnectionMode: jsii.String("connectionMode"),
//   	ConnectorProfileName: jsii.String("connectorProfileName"),
//   	ConnectorType: jsii.String("connectorType"),
//
//   	// the properties below are optional
//   	ConnectorLabel: jsii.String("connectorLabel"),
//   	ConnectorProfileConfig: &ConnectorProfileConfigProperty{
//   		ConnectorProfileCredentials: &ConnectorProfileCredentialsProperty{
//   			Amplitude: &AmplitudeConnectorProfileCredentialsProperty{
//   				ApiKey: jsii.String("apiKey"),
//   				SecretKey: jsii.String("secretKey"),
//   			},
//   			CustomConnector: &CustomConnectorProfileCredentialsProperty{
//   				AuthenticationType: jsii.String("authenticationType"),
//
//   				// the properties below are optional
//   				ApiKey: &ApiKeyCredentialsProperty{
//   					ApiKey: jsii.String("apiKey"),
//
//   					// the properties below are optional
//   					ApiSecretKey: jsii.String("apiSecretKey"),
//   				},
//   				Basic: &BasicAuthCredentialsProperty{
//   					Password: jsii.String("password"),
//   					Username: jsii.String("username"),
//   				},
//   				Custom: &CustomAuthCredentialsProperty{
//   					CustomAuthenticationType: jsii.String("customAuthenticationType"),
//
//   					// the properties below are optional
//   					CredentialsMap: map[string]*string{
//   						"credentialsMapKey": jsii.String("credentialsMap"),
//   					},
//   				},
//   				Oauth2: &OAuth2CredentialsProperty{
//   					AccessToken: jsii.String("accessToken"),
//   					ClientId: jsii.String("clientId"),
//   					ClientSecret: jsii.String("clientSecret"),
//   					OAuthRequest: &ConnectorOAuthRequestProperty{
//   						AuthCode: jsii.String("authCode"),
//   						RedirectUri: jsii.String("redirectUri"),
//   					},
//   					RefreshToken: jsii.String("refreshToken"),
//   				},
//   			},
//   			Datadog: &DatadogConnectorProfileCredentialsProperty{
//   				ApiKey: jsii.String("apiKey"),
//   				ApplicationKey: jsii.String("applicationKey"),
//   			},
//   			Dynatrace: &DynatraceConnectorProfileCredentialsProperty{
//   				ApiToken: jsii.String("apiToken"),
//   			},
//   			GoogleAnalytics: &GoogleAnalyticsConnectorProfileCredentialsProperty{
//   				ClientId: jsii.String("clientId"),
//   				ClientSecret: jsii.String("clientSecret"),
//
//   				// the properties below are optional
//   				AccessToken: jsii.String("accessToken"),
//   				ConnectorOAuthRequest: &ConnectorOAuthRequestProperty{
//   					AuthCode: jsii.String("authCode"),
//   					RedirectUri: jsii.String("redirectUri"),
//   				},
//   				RefreshToken: jsii.String("refreshToken"),
//   			},
//   			InforNexus: &InforNexusConnectorProfileCredentialsProperty{
//   				AccessKeyId: jsii.String("accessKeyId"),
//   				Datakey: jsii.String("datakey"),
//   				SecretAccessKey: jsii.String("secretAccessKey"),
//   				UserId: jsii.String("userId"),
//   			},
//   			Marketo: &MarketoConnectorProfileCredentialsProperty{
//   				ClientId: jsii.String("clientId"),
//   				ClientSecret: jsii.String("clientSecret"),
//
//   				// the properties below are optional
//   				AccessToken: jsii.String("accessToken"),
//   				ConnectorOAuthRequest: &ConnectorOAuthRequestProperty{
//   					AuthCode: jsii.String("authCode"),
//   					RedirectUri: jsii.String("redirectUri"),
//   				},
//   			},
//   			Pardot: &PardotConnectorProfileCredentialsProperty{
//   				AccessToken: jsii.String("accessToken"),
//   				ClientCredentialsArn: jsii.String("clientCredentialsArn"),
//   				ConnectorOAuthRequest: &ConnectorOAuthRequestProperty{
//   					AuthCode: jsii.String("authCode"),
//   					RedirectUri: jsii.String("redirectUri"),
//   				},
//   				RefreshToken: jsii.String("refreshToken"),
//   			},
//   			Redshift: &RedshiftConnectorProfileCredentialsProperty{
//   				Password: jsii.String("password"),
//   				Username: jsii.String("username"),
//   			},
//   			Salesforce: &SalesforceConnectorProfileCredentialsProperty{
//   				AccessToken: jsii.String("accessToken"),
//   				ClientCredentialsArn: jsii.String("clientCredentialsArn"),
//   				ConnectorOAuthRequest: &ConnectorOAuthRequestProperty{
//   					AuthCode: jsii.String("authCode"),
//   					RedirectUri: jsii.String("redirectUri"),
//   				},
//   				JwtToken: jsii.String("jwtToken"),
//   				OAuth2GrantType: jsii.String("oAuth2GrantType"),
//   				RefreshToken: jsii.String("refreshToken"),
//   			},
//   			SapoData: &SAPODataConnectorProfileCredentialsProperty{
//   				BasicAuthCredentials: &BasicAuthCredentialsProperty{
//   					Password: jsii.String("password"),
//   					Username: jsii.String("username"),
//   				},
//   				OAuthCredentials: &OAuthCredentialsProperty{
//   					AccessToken: jsii.String("accessToken"),
//   					ClientId: jsii.String("clientId"),
//   					ClientSecret: jsii.String("clientSecret"),
//   					ConnectorOAuthRequest: &ConnectorOAuthRequestProperty{
//   						AuthCode: jsii.String("authCode"),
//   						RedirectUri: jsii.String("redirectUri"),
//   					},
//   					RefreshToken: jsii.String("refreshToken"),
//   				},
//   			},
//   			ServiceNow: &ServiceNowConnectorProfileCredentialsProperty{
//   				OAuth2Credentials: &OAuth2CredentialsProperty{
//   					AccessToken: jsii.String("accessToken"),
//   					ClientId: jsii.String("clientId"),
//   					ClientSecret: jsii.String("clientSecret"),
//   					OAuthRequest: &ConnectorOAuthRequestProperty{
//   						AuthCode: jsii.String("authCode"),
//   						RedirectUri: jsii.String("redirectUri"),
//   					},
//   					RefreshToken: jsii.String("refreshToken"),
//   				},
//   				Password: jsii.String("password"),
//   				Username: jsii.String("username"),
//   			},
//   			Singular: &SingularConnectorProfileCredentialsProperty{
//   				ApiKey: jsii.String("apiKey"),
//   			},
//   			Slack: &SlackConnectorProfileCredentialsProperty{
//   				ClientId: jsii.String("clientId"),
//   				ClientSecret: jsii.String("clientSecret"),
//
//   				// the properties below are optional
//   				AccessToken: jsii.String("accessToken"),
//   				ConnectorOAuthRequest: &ConnectorOAuthRequestProperty{
//   					AuthCode: jsii.String("authCode"),
//   					RedirectUri: jsii.String("redirectUri"),
//   				},
//   			},
//   			Snowflake: &SnowflakeConnectorProfileCredentialsProperty{
//   				Password: jsii.String("password"),
//   				Username: jsii.String("username"),
//   			},
//   			Trendmicro: &TrendmicroConnectorProfileCredentialsProperty{
//   				ApiSecretKey: jsii.String("apiSecretKey"),
//   			},
//   			Veeva: &VeevaConnectorProfileCredentialsProperty{
//   				Password: jsii.String("password"),
//   				Username: jsii.String("username"),
//   			},
//   			Zendesk: &ZendeskConnectorProfileCredentialsProperty{
//   				ClientId: jsii.String("clientId"),
//   				ClientSecret: jsii.String("clientSecret"),
//
//   				// the properties below are optional
//   				AccessToken: jsii.String("accessToken"),
//   				ConnectorOAuthRequest: &ConnectorOAuthRequestProperty{
//   					AuthCode: jsii.String("authCode"),
//   					RedirectUri: jsii.String("redirectUri"),
//   				},
//   			},
//   		},
//   		ConnectorProfileProperties: &ConnectorProfilePropertiesProperty{
//   			CustomConnector: &CustomConnectorProfilePropertiesProperty{
//   				OAuth2Properties: &OAuth2PropertiesProperty{
//   					OAuth2GrantType: jsii.String("oAuth2GrantType"),
//   					TokenUrl: jsii.String("tokenUrl"),
//   					TokenUrlCustomProperties: map[string]*string{
//   						"tokenUrlCustomPropertiesKey": jsii.String("tokenUrlCustomProperties"),
//   					},
//   				},
//   				ProfileProperties: map[string]*string{
//   					"profilePropertiesKey": jsii.String("profileProperties"),
//   				},
//   			},
//   			Datadog: &DatadogConnectorProfilePropertiesProperty{
//   				InstanceUrl: jsii.String("instanceUrl"),
//   			},
//   			Dynatrace: &DynatraceConnectorProfilePropertiesProperty{
//   				InstanceUrl: jsii.String("instanceUrl"),
//   			},
//   			InforNexus: &InforNexusConnectorProfilePropertiesProperty{
//   				InstanceUrl: jsii.String("instanceUrl"),
//   			},
//   			Marketo: &MarketoConnectorProfilePropertiesProperty{
//   				InstanceUrl: jsii.String("instanceUrl"),
//   			},
//   			Pardot: &PardotConnectorProfilePropertiesProperty{
//   				BusinessUnitId: jsii.String("businessUnitId"),
//
//   				// the properties below are optional
//   				InstanceUrl: jsii.String("instanceUrl"),
//   				IsSandboxEnvironment: jsii.Boolean(false),
//   			},
//   			Redshift: &RedshiftConnectorProfilePropertiesProperty{
//   				BucketName: jsii.String("bucketName"),
//   				RoleArn: jsii.String("roleArn"),
//
//   				// the properties below are optional
//   				BucketPrefix: jsii.String("bucketPrefix"),
//   				ClusterIdentifier: jsii.String("clusterIdentifier"),
//   				DataApiRoleArn: jsii.String("dataApiRoleArn"),
//   				DatabaseName: jsii.String("databaseName"),
//   				DatabaseUrl: jsii.String("databaseUrl"),
//   				IsRedshiftServerless: jsii.Boolean(false),
//   				WorkgroupName: jsii.String("workgroupName"),
//   			},
//   			Salesforce: &SalesforceConnectorProfilePropertiesProperty{
//   				InstanceUrl: jsii.String("instanceUrl"),
//   				IsSandboxEnvironment: jsii.Boolean(false),
//   				UsePrivateLinkForMetadataAndAuthorization: jsii.Boolean(false),
//   			},
//   			SapoData: &SAPODataConnectorProfilePropertiesProperty{
//   				ApplicationHostUrl: jsii.String("applicationHostUrl"),
//   				ApplicationServicePath: jsii.String("applicationServicePath"),
//   				ClientNumber: jsii.String("clientNumber"),
//   				DisableSso: jsii.Boolean(false),
//   				LogonLanguage: jsii.String("logonLanguage"),
//   				OAuthProperties: &OAuthPropertiesProperty{
//   					AuthCodeUrl: jsii.String("authCodeUrl"),
//   					OAuthScopes: []*string{
//   						jsii.String("oAuthScopes"),
//   					},
//   					TokenUrl: jsii.String("tokenUrl"),
//   				},
//   				PortNumber: jsii.Number(123),
//   				PrivateLinkServiceName: jsii.String("privateLinkServiceName"),
//   			},
//   			ServiceNow: &ServiceNowConnectorProfilePropertiesProperty{
//   				InstanceUrl: jsii.String("instanceUrl"),
//   			},
//   			Slack: &SlackConnectorProfilePropertiesProperty{
//   				InstanceUrl: jsii.String("instanceUrl"),
//   			},
//   			Snowflake: &SnowflakeConnectorProfilePropertiesProperty{
//   				BucketName: jsii.String("bucketName"),
//   				Stage: jsii.String("stage"),
//   				Warehouse: jsii.String("warehouse"),
//
//   				// the properties below are optional
//   				AccountName: jsii.String("accountName"),
//   				BucketPrefix: jsii.String("bucketPrefix"),
//   				PrivateLinkServiceName: jsii.String("privateLinkServiceName"),
//   				Region: jsii.String("region"),
//   			},
//   			Veeva: &VeevaConnectorProfilePropertiesProperty{
//   				InstanceUrl: jsii.String("instanceUrl"),
//   			},
//   			Zendesk: &ZendeskConnectorProfilePropertiesProperty{
//   				InstanceUrl: jsii.String("instanceUrl"),
//   			},
//   		},
//   	},
//   	KmsArn: jsii.String("kmsArn"),
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appflow-connectorprofile.html
//
type CfnConnectorProfile interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of the connector profile.
	AttrConnectorProfileArn() *string
	// The Amazon Resource Name (ARN) of the connector profile credentials.
	AttrCredentialsArn() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Indicates the connection mode and if it is public or private.
	ConnectionMode() *string
	SetConnectionMode(val *string)
	// The label for the connector profile being created.
	ConnectorLabel() *string
	SetConnectorLabel(val *string)
	// Defines the connector-specific configuration and credentials.
	ConnectorProfileConfig() interface{}
	SetConnectorProfileConfig(val interface{})
	// The name of the connector profile.
	ConnectorProfileName() *string
	SetConnectorProfileName(val *string)
	// The type of connector, such as Salesforce, Amplitude, and so on.
	ConnectorType() *string
	SetConnectorType(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The ARN (Amazon Resource Name) of the Key Management Service (KMS) key you provide for encryption.
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
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
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

func (j *jsiiProxy_CfnConnectorProfile) ConnectorLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectorLabel",
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

func (j *jsiiProxy_CfnConnectorProfile) Node() constructs.Node {
	var returns constructs.Node
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

func (j *jsiiProxy_CfnConnectorProfile) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


func NewCfnConnectorProfile(scope constructs.Construct, id *string, props *CfnConnectorProfileProps) CfnConnectorProfile {
	_init_.Initialize()

	if err := validateNewCfnConnectorProfileParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnConnectorProfile{}

	_jsii_.Create(
		"aws-cdk-lib.aws_appflow.CfnConnectorProfile",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnConnectorProfile_Override(c CfnConnectorProfile, scope constructs.Construct, id *string, props *CfnConnectorProfileProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appflow.CfnConnectorProfile",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnConnectorProfile)SetConnectionMode(val *string) {
	if err := j.validateSetConnectionModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionMode",
		val,
	)
}

func (j *jsiiProxy_CfnConnectorProfile)SetConnectorLabel(val *string) {
	_jsii_.Set(
		j,
		"connectorLabel",
		val,
	)
}

func (j *jsiiProxy_CfnConnectorProfile)SetConnectorProfileConfig(val interface{}) {
	if err := j.validateSetConnectorProfileConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectorProfileConfig",
		val,
	)
}

func (j *jsiiProxy_CfnConnectorProfile)SetConnectorProfileName(val *string) {
	if err := j.validateSetConnectorProfileNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectorProfileName",
		val,
	)
}

func (j *jsiiProxy_CfnConnectorProfile)SetConnectorType(val *string) {
	if err := j.validateSetConnectorTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectorType",
		val,
	)
}

func (j *jsiiProxy_CfnConnectorProfile)SetKmsArn(val *string) {
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
func CfnConnectorProfile_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnConnectorProfile_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appflow.CfnConnectorProfile",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnConnectorProfile_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnConnectorProfile_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appflow.CfnConnectorProfile",
		"isCfnResource",
		[]interface{}{construct},
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
func CfnConnectorProfile_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnConnectorProfile_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appflow.CfnConnectorProfile",
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
		"aws-cdk-lib.aws_appflow.CfnConnectorProfile",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnConnectorProfile) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnConnectorProfile) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnConnectorProfile) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnConnectorProfile) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnConnectorProfile) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnConnectorProfile) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnConnectorProfile) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnConnectorProfile) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnConnectorProfile) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnConnectorProfile) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnConnectorProfile) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnConnectorProfile) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConnectorProfile) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConnectorProfile) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnConnectorProfile) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnConnectorProfile) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnConnectorProfile) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
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

func (c *jsiiProxy_CfnConnectorProfile) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

