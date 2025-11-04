package awsappsync

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::AppSync::GraphQLApi` resource creates a new AWS AppSync GraphQL API.
//
// This is the top-level construct for your application. For more information, see [Quick Start](https://docs.aws.amazon.com/appsync/latest/devguide/quickstart.html) in the *AWS AppSync Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGraphQLApi := awscdk.Aws_appsync.NewCfnGraphQLApi(this, jsii.String("MyCfnGraphQLApi"), &CfnGraphQLApiProps{
//   	AuthenticationType: jsii.String("authenticationType"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	AdditionalAuthenticationProviders: []interface{}{
//   		&AdditionalAuthenticationProviderProperty{
//   			AuthenticationType: jsii.String("authenticationType"),
//
//   			// the properties below are optional
//   			LambdaAuthorizerConfig: &LambdaAuthorizerConfigProperty{
//   				AuthorizerResultTtlInSeconds: jsii.Number(123),
//   				AuthorizerUri: jsii.String("authorizerUri"),
//   				IdentityValidationExpression: jsii.String("identityValidationExpression"),
//   			},
//   			OpenIdConnectConfig: &OpenIDConnectConfigProperty{
//   				AuthTtl: jsii.Number(123),
//   				ClientId: jsii.String("clientId"),
//   				IatTtl: jsii.Number(123),
//   				Issuer: jsii.String("issuer"),
//   			},
//   			UserPoolConfig: &CognitoUserPoolConfigProperty{
//   				AppIdClientRegex: jsii.String("appIdClientRegex"),
//   				AwsRegion: jsii.String("awsRegion"),
//   				UserPoolId: jsii.String("userPoolId"),
//   			},
//   		},
//   	},
//   	ApiType: jsii.String("apiType"),
//   	EnhancedMetricsConfig: &EnhancedMetricsConfigProperty{
//   		DataSourceLevelMetricsBehavior: jsii.String("dataSourceLevelMetricsBehavior"),
//   		OperationLevelMetricsConfig: jsii.String("operationLevelMetricsConfig"),
//   		ResolverLevelMetricsBehavior: jsii.String("resolverLevelMetricsBehavior"),
//   	},
//   	EnvironmentVariables: map[string]*string{
//   		"environmentVariablesKey": jsii.String("environmentVariables"),
//   	},
//   	IntrospectionConfig: jsii.String("introspectionConfig"),
//   	LambdaAuthorizerConfig: &LambdaAuthorizerConfigProperty{
//   		AuthorizerResultTtlInSeconds: jsii.Number(123),
//   		AuthorizerUri: jsii.String("authorizerUri"),
//   		IdentityValidationExpression: jsii.String("identityValidationExpression"),
//   	},
//   	LogConfig: &LogConfigProperty{
//   		CloudWatchLogsRoleArn: jsii.String("cloudWatchLogsRoleArn"),
//   		ExcludeVerboseContent: jsii.Boolean(false),
//   		FieldLogLevel: jsii.String("fieldLogLevel"),
//   	},
//   	MergedApiExecutionRoleArn: jsii.String("mergedApiExecutionRoleArn"),
//   	OpenIdConnectConfig: &OpenIDConnectConfigProperty{
//   		AuthTtl: jsii.Number(123),
//   		ClientId: jsii.String("clientId"),
//   		IatTtl: jsii.Number(123),
//   		Issuer: jsii.String("issuer"),
//   	},
//   	OwnerContact: jsii.String("ownerContact"),
//   	QueryDepthLimit: jsii.Number(123),
//   	ResolverCountLimit: jsii.Number(123),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	UserPoolConfig: &UserPoolConfigProperty{
//   		AppIdClientRegex: jsii.String("appIdClientRegex"),
//   		AwsRegion: jsii.String("awsRegion"),
//   		DefaultAction: jsii.String("defaultAction"),
//   		UserPoolId: jsii.String("userPoolId"),
//   	},
//   	Visibility: jsii.String("visibility"),
//   	XrayEnabled: jsii.Boolean(false),
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-graphqlapi.html
//
type CfnGraphQLApi interface {
	awscdk.CfnResource
	IGraphQLApiRef
	awscdk.IInspectable
	awscdk.ITaggable
	// A list of additional authentication providers for the `GraphqlApi` API.
	AdditionalAuthenticationProviders() interface{}
	SetAdditionalAuthenticationProviders(val interface{})
	// The value that indicates whether the GraphQL API is a standard API ( `GRAPHQL` ) or merged API ( `MERGED` ).
	ApiType() *string
	SetApiType(val *string)
	// Unique AWS AppSync GraphQL API identifier.
	AttrApiId() *string
	// The Amazon Resource Name (ARN) of the API key, such as `arn:aws:appsync:us-east-1:123456789012:apis/graphqlapiid` .
	AttrArn() *string
	// The fully qualified domain name (FQDN) of the endpoint URL of your GraphQL API.
	AttrGraphQlDns() *string
	// The GraphQL endpoint ARN.
	AttrGraphQlEndpointArn() *string
	// The Endpoint URL of your GraphQL API.
	AttrGraphQlUrl() *string
	// The fully qualified domain name (FQDN) of the real-time endpoint URL of your GraphQL API.
	AttrRealtimeDns() *string
	// The GraphQL API real-time endpoint URL.
	//
	// For more information, see [Discovering the real-time endpoint from the GraphQL endpoint](https://docs.aws.amazon.com/appsync/latest/devguide/real-time-websocket-client.html#handshake-details-to-establish-the-websocket-connection) .
	AttrRealtimeUrl() *string
	// Security configuration for your GraphQL API.
	AuthenticationType() *string
	SetAuthenticationType(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Enables and controls the enhanced metrics feature.
	EnhancedMetricsConfig() interface{}
	SetEnhancedMetricsConfig(val interface{})
	Env() *awscdk.ResourceEnvironment
	// A map containing the list of resources with their properties and environment variables.
	EnvironmentVariables() interface{}
	SetEnvironmentVariables(val interface{})
	// A reference to a GraphQLApi resource.
	GraphQlApiRef() *GraphQLApiReference
	// Sets the value of the GraphQL API to enable ( `ENABLED` ) or disable ( `DISABLED` ) introspection.
	IntrospectionConfig() *string
	SetIntrospectionConfig(val *string)
	// A `LambdaAuthorizerConfig` holds configuration on how to authorize AWS AppSync API access when using the `AWS_LAMBDA` authorizer mode.
	LambdaAuthorizerConfig() interface{}
	SetLambdaAuthorizerConfig(val interface{})
	// The Amazon CloudWatch Logs configuration.
	LogConfig() interface{}
	SetLogConfig(val interface{})
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
	// The AWS Identity and Access Management service role ARN for a merged API.
	MergedApiExecutionRoleArn() *string
	SetMergedApiExecutionRoleArn(val *string)
	// The API name.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// The OpenID Connect configuration.
	OpenIdConnectConfig() interface{}
	SetOpenIdConnectConfig(val interface{})
	// The owner contact information for an API resource.
	OwnerContact() *string
	SetOwnerContact(val *string)
	// The maximum depth a query can have in a single request.
	QueryDepthLimit() *float64
	SetQueryDepthLimit(val *float64)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The maximum number of resolvers that can be invoked in a single request.
	ResolverCountLimit() *float64
	SetResolverCountLimit(val *float64)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Tag Manager which manages the tags for this resource.
	Tags() awscdk.TagManager
	// An arbitrary set of tags (key-value pairs) for this GraphQL API.
	TagsRaw() *[]*awscdk.CfnTag
	SetTagsRaw(val *[]*awscdk.CfnTag)
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
	// Optional authorization configuration for using Amazon Cognito user pools with your GraphQL endpoint.
	UserPoolConfig() interface{}
	SetUserPoolConfig(val interface{})
	// Sets the scope of the GraphQL API to public ( `GLOBAL` ) or private ( `PRIVATE` ).
	Visibility() *string
	SetVisibility(val *string)
	// A flag indicating whether to use AWS X-Ray tracing for this `GraphqlApi` .
	XrayEnabled() interface{}
	SetXrayEnabled(val interface{})
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

// The jsii proxy struct for CfnGraphQLApi
type jsiiProxy_CfnGraphQLApi struct {
	internal.Type__awscdkCfnResource
	jsiiProxy_IGraphQLApiRef
	internal.Type__awscdkIInspectable
	internal.Type__awscdkITaggable
}

func (j *jsiiProxy_CfnGraphQLApi) AdditionalAuthenticationProviders() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"additionalAuthenticationProviders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) ApiType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) AttrApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrApiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) AttrGraphQlDns() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrGraphQlDns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) AttrGraphQlEndpointArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrGraphQlEndpointArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) AttrGraphQlUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrGraphQlUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) AttrRealtimeDns() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrRealtimeDns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) AttrRealtimeUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrRealtimeUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) AuthenticationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) EnhancedMetricsConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enhancedMetricsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) EnvironmentVariables() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"environmentVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) GraphQlApiRef() *GraphQLApiReference {
	var returns *GraphQLApiReference
	_jsii_.Get(
		j,
		"graphQlApiRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) IntrospectionConfig() *string {
	var returns *string
	_jsii_.Get(
		j,
		"introspectionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) LambdaAuthorizerConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaAuthorizerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) LogConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) MergedApiExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mergedApiExecutionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) OpenIdConnectConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"openIdConnectConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) OwnerContact() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerContact",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) QueryDepthLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queryDepthLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) ResolverCountLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"resolverCountLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) TagsRaw() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tagsRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) UserPoolConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userPoolConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) Visibility() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visibility",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) XrayEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"xrayEnabled",
		&returns,
	)
	return returns
}


func NewCfnGraphQLApi(scope constructs.Construct, id *string, props *CfnGraphQLApiProps) CfnGraphQLApi {
	_init_.Initialize()

	if err := validateNewCfnGraphQLApiParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnGraphQLApi{}

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.CfnGraphQLApi",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnGraphQLApi_Override(c CfnGraphQLApi, scope constructs.Construct, id *string, props *CfnGraphQLApiProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.CfnGraphQLApi",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnGraphQLApi)SetAdditionalAuthenticationProviders(val interface{}) {
	if err := j.validateSetAdditionalAuthenticationProvidersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalAuthenticationProviders",
		val,
	)
}

func (j *jsiiProxy_CfnGraphQLApi)SetApiType(val *string) {
	_jsii_.Set(
		j,
		"apiType",
		val,
	)
}

func (j *jsiiProxy_CfnGraphQLApi)SetAuthenticationType(val *string) {
	if err := j.validateSetAuthenticationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationType",
		val,
	)
}

func (j *jsiiProxy_CfnGraphQLApi)SetEnhancedMetricsConfig(val interface{}) {
	if err := j.validateSetEnhancedMetricsConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enhancedMetricsConfig",
		val,
	)
}

func (j *jsiiProxy_CfnGraphQLApi)SetEnvironmentVariables(val interface{}) {
	if err := j.validateSetEnvironmentVariablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environmentVariables",
		val,
	)
}

func (j *jsiiProxy_CfnGraphQLApi)SetIntrospectionConfig(val *string) {
	_jsii_.Set(
		j,
		"introspectionConfig",
		val,
	)
}

func (j *jsiiProxy_CfnGraphQLApi)SetLambdaAuthorizerConfig(val interface{}) {
	if err := j.validateSetLambdaAuthorizerConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lambdaAuthorizerConfig",
		val,
	)
}

func (j *jsiiProxy_CfnGraphQLApi)SetLogConfig(val interface{}) {
	if err := j.validateSetLogConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logConfig",
		val,
	)
}

func (j *jsiiProxy_CfnGraphQLApi)SetMergedApiExecutionRoleArn(val *string) {
	_jsii_.Set(
		j,
		"mergedApiExecutionRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnGraphQLApi)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnGraphQLApi)SetOpenIdConnectConfig(val interface{}) {
	if err := j.validateSetOpenIdConnectConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"openIdConnectConfig",
		val,
	)
}

func (j *jsiiProxy_CfnGraphQLApi)SetOwnerContact(val *string) {
	_jsii_.Set(
		j,
		"ownerContact",
		val,
	)
}

func (j *jsiiProxy_CfnGraphQLApi)SetQueryDepthLimit(val *float64) {
	_jsii_.Set(
		j,
		"queryDepthLimit",
		val,
	)
}

func (j *jsiiProxy_CfnGraphQLApi)SetResolverCountLimit(val *float64) {
	_jsii_.Set(
		j,
		"resolverCountLimit",
		val,
	)
}

func (j *jsiiProxy_CfnGraphQLApi)SetTagsRaw(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsRawParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsRaw",
		val,
	)
}

func (j *jsiiProxy_CfnGraphQLApi)SetUserPoolConfig(val interface{}) {
	if err := j.validateSetUserPoolConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userPoolConfig",
		val,
	)
}

func (j *jsiiProxy_CfnGraphQLApi)SetVisibility(val *string) {
	_jsii_.Set(
		j,
		"visibility",
		val,
	)
}

func (j *jsiiProxy_CfnGraphQLApi)SetXrayEnabled(val interface{}) {
	if err := j.validateSetXrayEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"xrayEnabled",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnGraphQLApi_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnGraphQLApi_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.CfnGraphQLApi",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnGraphQLApi_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnGraphQLApi_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.CfnGraphQLApi",
		"isCfnResource",
		[]interface{}{x},
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
func CfnGraphQLApi_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnGraphQLApi_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.CfnGraphQLApi",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnGraphQLApi_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_appsync.CfnGraphQLApi",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnGraphQLApi) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnGraphQLApi) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnGraphQLApi) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnGraphQLApi) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnGraphQLApi) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnGraphQLApi) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnGraphQLApi) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnGraphQLApi) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnGraphQLApi) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnGraphQLApi) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnGraphQLApi) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnGraphQLApi) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGraphQLApi) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGraphQLApi) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnGraphQLApi) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnGraphQLApi) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnGraphQLApi) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnGraphQLApi) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGraphQLApi) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGraphQLApi) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

