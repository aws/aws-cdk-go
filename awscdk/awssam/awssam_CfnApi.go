package awssam

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssam/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::Serverless::Api`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var authorizers interface{}
//   var definitionBody interface{}
//   var gatewayResponses interface{}
//   var methodSettings interface{}
//   var models interface{}
//
//   cfnApi := awscdk.Aws_sam.NewCfnApi(this, jsii.String("MyCfnApi"), &cfnApiProps{
//   	stageName: jsii.String("stageName"),
//
//   	// the properties below are optional
//   	accessLogSetting: &accessLogSettingProperty{
//   		destinationArn: jsii.String("destinationArn"),
//   		format: jsii.String("format"),
//   	},
//   	auth: &authProperty{
//   		addDefaultAuthorizerToCorsPreflight: jsii.Boolean(false),
//   		authorizers: authorizers,
//   		defaultAuthorizer: jsii.String("defaultAuthorizer"),
//   	},
//   	binaryMediaTypes: []*string{
//   		jsii.String("binaryMediaTypes"),
//   	},
//   	cacheClusterEnabled: jsii.Boolean(false),
//   	cacheClusterSize: jsii.String("cacheClusterSize"),
//   	canarySetting: &canarySettingProperty{
//   		deploymentId: jsii.String("deploymentId"),
//   		percentTraffic: jsii.Number(123),
//   		stageVariableOverrides: map[string]*string{
//   			"stageVariableOverridesKey": jsii.String("stageVariableOverrides"),
//   		},
//   		useStageCache: jsii.Boolean(false),
//   	},
//   	cors: jsii.String("cors"),
//   	definitionBody: definitionBody,
//   	definitionUri: jsii.String("definitionUri"),
//   	description: jsii.String("description"),
//   	domain: &domainConfigurationProperty{
//   		certificateArn: jsii.String("certificateArn"),
//   		domainName: jsii.String("domainName"),
//
//   		// the properties below are optional
//   		basePath: []*string{
//   			jsii.String("basePath"),
//   		},
//   		endpointConfiguration: jsii.String("endpointConfiguration"),
//   		mutualTlsAuthentication: &mutualTlsAuthenticationProperty{
//   			truststoreUri: jsii.String("truststoreUri"),
//   			truststoreVersion: jsii.String("truststoreVersion"),
//   		},
//   		ownershipVerificationCertificateArn: jsii.String("ownershipVerificationCertificateArn"),
//   		route53: &route53ConfigurationProperty{
//   			distributedDomainName: jsii.String("distributedDomainName"),
//   			evaluateTargetHealth: jsii.Boolean(false),
//   			hostedZoneId: jsii.String("hostedZoneId"),
//   			hostedZoneName: jsii.String("hostedZoneName"),
//   			ipV6: jsii.Boolean(false),
//   		},
//   		securityPolicy: jsii.String("securityPolicy"),
//   	},
//   	endpointConfiguration: jsii.String("endpointConfiguration"),
//   	gatewayResponses: gatewayResponses,
//   	methodSettings: []interface{}{
//   		methodSettings,
//   	},
//   	minimumCompressionSize: jsii.Number(123),
//   	models: models,
//   	name: jsii.String("name"),
//   	openApiVersion: jsii.String("openApiVersion"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	tracingEnabled: jsii.Boolean(false),
//   	variables: map[string]*string{
//   		"variablesKey": jsii.String("variables"),
//   	},
//   })
//
type CfnApi interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// `AWS::Serverless::Api.AccessLogSetting`.
	AccessLogSetting() interface{}
	SetAccessLogSetting(val interface{})
	// `AWS::Serverless::Api.Auth`.
	Auth() interface{}
	SetAuth(val interface{})
	// `AWS::Serverless::Api.BinaryMediaTypes`.
	BinaryMediaTypes() *[]*string
	SetBinaryMediaTypes(val *[]*string)
	// `AWS::Serverless::Api.CacheClusterEnabled`.
	CacheClusterEnabled() interface{}
	SetCacheClusterEnabled(val interface{})
	// `AWS::Serverless::Api.CacheClusterSize`.
	CacheClusterSize() *string
	SetCacheClusterSize(val *string)
	// `AWS::Serverless::Api.CanarySetting`.
	CanarySetting() interface{}
	SetCanarySetting(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// `AWS::Serverless::Api.Cors`.
	Cors() interface{}
	SetCors(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// `AWS::Serverless::Api.DefinitionBody`.
	DefinitionBody() interface{}
	SetDefinitionBody(val interface{})
	// `AWS::Serverless::Api.DefinitionUri`.
	DefinitionUri() interface{}
	SetDefinitionUri(val interface{})
	// `AWS::Serverless::Api.Description`.
	Description() *string
	SetDescription(val *string)
	// `AWS::Serverless::Api.Domain`.
	Domain() interface{}
	SetDomain(val interface{})
	// `AWS::Serverless::Api.EndpointConfiguration`.
	EndpointConfiguration() interface{}
	SetEndpointConfiguration(val interface{})
	// `AWS::Serverless::Api.GatewayResponses`.
	GatewayResponses() interface{}
	SetGatewayResponses(val interface{})
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
	// `AWS::Serverless::Api.MethodSettings`.
	MethodSettings() interface{}
	SetMethodSettings(val interface{})
	// `AWS::Serverless::Api.MinimumCompressionSize`.
	MinimumCompressionSize() *float64
	SetMinimumCompressionSize(val *float64)
	// `AWS::Serverless::Api.Models`.
	Models() interface{}
	SetModels(val interface{})
	// `AWS::Serverless::Api.Name`.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// `AWS::Serverless::Api.OpenApiVersion`.
	OpenApiVersion() *string
	SetOpenApiVersion(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// `AWS::Serverless::Api.StageName`.
	StageName() *string
	SetStageName(val *string)
	// `AWS::Serverless::Api.Tags`.
	Tags() awscdk.TagManager
	// `AWS::Serverless::Api.TracingEnabled`.
	TracingEnabled() interface{}
	SetTracingEnabled(val interface{})
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
	// `AWS::Serverless::Api.Variables`.
	Variables() interface{}
	SetVariables(val interface{})
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
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
	GetAtt(attributeName *string) awscdk.Reference
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
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
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

// The jsii proxy struct for CfnApi
type jsiiProxy_CfnApi struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnApi) AccessLogSetting() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessLogSetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) Auth() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"auth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) BinaryMediaTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"binaryMediaTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) CacheClusterEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cacheClusterEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) CacheClusterSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheClusterSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) CanarySetting() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"canarySetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) Cors() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) DefinitionBody() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"definitionBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) DefinitionUri() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"definitionUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) Domain() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) EndpointConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"endpointConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) GatewayResponses() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gatewayResponses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) MethodSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"methodSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) MinimumCompressionSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumCompressionSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) Models() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"models",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) OpenApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openApiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) StageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) TracingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tracingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) Variables() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"variables",
		&returns,
	)
	return returns
}


// Create a new `AWS::Serverless::Api`.
func NewCfnApi(scope constructs.Construct, id *string, props *CfnApiProps) CfnApi {
	_init_.Initialize()

	if err := validateNewCfnApiParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnApi{}

	_jsii_.Create(
		"aws-cdk-lib.aws_sam.CfnApi",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Serverless::Api`.
func NewCfnApi_Override(c CfnApi, scope constructs.Construct, id *string, props *CfnApiProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_sam.CfnApi",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnApi)SetAccessLogSetting(val interface{}) {
	if err := j.validateSetAccessLogSettingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessLogSetting",
		val,
	)
}

func (j *jsiiProxy_CfnApi)SetAuth(val interface{}) {
	if err := j.validateSetAuthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"auth",
		val,
	)
}

func (j *jsiiProxy_CfnApi)SetBinaryMediaTypes(val *[]*string) {
	_jsii_.Set(
		j,
		"binaryMediaTypes",
		val,
	)
}

func (j *jsiiProxy_CfnApi)SetCacheClusterEnabled(val interface{}) {
	if err := j.validateSetCacheClusterEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheClusterEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnApi)SetCacheClusterSize(val *string) {
	_jsii_.Set(
		j,
		"cacheClusterSize",
		val,
	)
}

func (j *jsiiProxy_CfnApi)SetCanarySetting(val interface{}) {
	if err := j.validateSetCanarySettingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"canarySetting",
		val,
	)
}

func (j *jsiiProxy_CfnApi)SetCors(val interface{}) {
	if err := j.validateSetCorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cors",
		val,
	)
}

func (j *jsiiProxy_CfnApi)SetDefinitionBody(val interface{}) {
	if err := j.validateSetDefinitionBodyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"definitionBody",
		val,
	)
}

func (j *jsiiProxy_CfnApi)SetDefinitionUri(val interface{}) {
	if err := j.validateSetDefinitionUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"definitionUri",
		val,
	)
}

func (j *jsiiProxy_CfnApi)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnApi)SetDomain(val interface{}) {
	if err := j.validateSetDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domain",
		val,
	)
}

func (j *jsiiProxy_CfnApi)SetEndpointConfiguration(val interface{}) {
	if err := j.validateSetEndpointConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnApi)SetGatewayResponses(val interface{}) {
	if err := j.validateSetGatewayResponsesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gatewayResponses",
		val,
	)
}

func (j *jsiiProxy_CfnApi)SetMethodSettings(val interface{}) {
	if err := j.validateSetMethodSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"methodSettings",
		val,
	)
}

func (j *jsiiProxy_CfnApi)SetMinimumCompressionSize(val *float64) {
	_jsii_.Set(
		j,
		"minimumCompressionSize",
		val,
	)
}

func (j *jsiiProxy_CfnApi)SetModels(val interface{}) {
	if err := j.validateSetModelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"models",
		val,
	)
}

func (j *jsiiProxy_CfnApi)SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnApi)SetOpenApiVersion(val *string) {
	_jsii_.Set(
		j,
		"openApiVersion",
		val,
	)
}

func (j *jsiiProxy_CfnApi)SetStageName(val *string) {
	if err := j.validateSetStageNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stageName",
		val,
	)
}

func (j *jsiiProxy_CfnApi)SetTracingEnabled(val interface{}) {
	if err := j.validateSetTracingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tracingEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnApi)SetVariables(val interface{}) {
	if err := j.validateSetVariablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"variables",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnApi_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnApi_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sam.CfnApi",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnApi_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnApi_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sam.CfnApi",
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
func CfnApi_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnApi_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sam.CfnApi",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnApi_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_sam.CfnApi",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func CfnApi_REQUIRED_TRANSFORM() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_sam.CfnApi",
		"REQUIRED_TRANSFORM",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnApi) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnApi) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnApi) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnApi) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnApi) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnApi) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnApi) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnApi) GetAtt(attributeName *string) awscdk.Reference {
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

func (c *jsiiProxy_CfnApi) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnApi) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnApi) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnApi) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnApi) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApi) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApi) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

