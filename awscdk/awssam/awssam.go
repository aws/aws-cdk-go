package awssam

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awssam/internal"
	"github.com/aws/constructs-go/constructs/v3"
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
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// `AWS::Serverless::Api.Cors`.
	Cors() interface{}
	SetCors(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
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
	// Experimental.
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
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// `AWS::Serverless::Api.OpenApiVersion`.
	OpenApiVersion() *string
	SetOpenApiVersion(val *string)
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
	// `AWS::Serverless::Api.StageName`.
	StageName() *string
	SetStageName(val *string)
	// `AWS::Serverless::Api.Tags`.
	Tags() awscdk.TagManager
	// `AWS::Serverless::Api.TracingEnabled`.
	TracingEnabled() interface{}
	SetTracingEnabled(val interface{})
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// `AWS::Serverless::Api.Variables`.
	Variables() interface{}
	SetVariables(val interface{})
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

func (j *jsiiProxy_CfnApi) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
func NewCfnApi(scope awscdk.Construct, id *string, props *CfnApiProps) CfnApi {
	_init_.Initialize()

	j := jsiiProxy_CfnApi{}

	_jsii_.Create(
		"monocdk.aws_sam.CfnApi",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Serverless::Api`.
func NewCfnApi_Override(c CfnApi, scope awscdk.Construct, id *string, props *CfnApiProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_sam.CfnApi",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnApi) SetAccessLogSetting(val interface{}) {
	_jsii_.Set(
		j,
		"accessLogSetting",
		val,
	)
}

func (j *jsiiProxy_CfnApi) SetAuth(val interface{}) {
	_jsii_.Set(
		j,
		"auth",
		val,
	)
}

func (j *jsiiProxy_CfnApi) SetBinaryMediaTypes(val *[]*string) {
	_jsii_.Set(
		j,
		"binaryMediaTypes",
		val,
	)
}

func (j *jsiiProxy_CfnApi) SetCacheClusterEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"cacheClusterEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnApi) SetCacheClusterSize(val *string) {
	_jsii_.Set(
		j,
		"cacheClusterSize",
		val,
	)
}

func (j *jsiiProxy_CfnApi) SetCanarySetting(val interface{}) {
	_jsii_.Set(
		j,
		"canarySetting",
		val,
	)
}

func (j *jsiiProxy_CfnApi) SetCors(val interface{}) {
	_jsii_.Set(
		j,
		"cors",
		val,
	)
}

func (j *jsiiProxy_CfnApi) SetDefinitionBody(val interface{}) {
	_jsii_.Set(
		j,
		"definitionBody",
		val,
	)
}

func (j *jsiiProxy_CfnApi) SetDefinitionUri(val interface{}) {
	_jsii_.Set(
		j,
		"definitionUri",
		val,
	)
}

func (j *jsiiProxy_CfnApi) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnApi) SetDomain(val interface{}) {
	_jsii_.Set(
		j,
		"domain",
		val,
	)
}

func (j *jsiiProxy_CfnApi) SetEndpointConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"endpointConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnApi) SetGatewayResponses(val interface{}) {
	_jsii_.Set(
		j,
		"gatewayResponses",
		val,
	)
}

func (j *jsiiProxy_CfnApi) SetMethodSettings(val interface{}) {
	_jsii_.Set(
		j,
		"methodSettings",
		val,
	)
}

func (j *jsiiProxy_CfnApi) SetMinimumCompressionSize(val *float64) {
	_jsii_.Set(
		j,
		"minimumCompressionSize",
		val,
	)
}

func (j *jsiiProxy_CfnApi) SetModels(val interface{}) {
	_jsii_.Set(
		j,
		"models",
		val,
	)
}

func (j *jsiiProxy_CfnApi) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnApi) SetOpenApiVersion(val *string) {
	_jsii_.Set(
		j,
		"openApiVersion",
		val,
	)
}

func (j *jsiiProxy_CfnApi) SetStageName(val *string) {
	_jsii_.Set(
		j,
		"stageName",
		val,
	)
}

func (j *jsiiProxy_CfnApi) SetTracingEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"tracingEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnApi) SetVariables(val interface{}) {
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
// Experimental.
func CfnApi_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sam.CfnApi",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnApi_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sam.CfnApi",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnApi_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sam.CfnApi",
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
		"monocdk.aws_sam.CfnApi",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func CfnApi_REQUIRED_TRANSFORM() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_sam.CfnApi",
		"REQUIRED_TRANSFORM",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnApi) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnApi) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnApi) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnApi) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnApi) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnApi) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnApi) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnApi) GetAtt(attributeName *string) awscdk.Reference {
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
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnApi) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnApi) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnApi) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApi) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnApi) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnApi) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnApi) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
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

func (c *jsiiProxy_CfnApi) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApi) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessLogSettingProperty := &accessLogSettingProperty{
//   	destinationArn: jsii.String("destinationArn"),
//   	format: jsii.String("format"),
//   }
//
type CfnApi_AccessLogSettingProperty struct {
	// `CfnApi.AccessLogSettingProperty.DestinationArn`.
	DestinationArn *string `field:"optional" json:"destinationArn" yaml:"destinationArn"`
	// `CfnApi.AccessLogSettingProperty.Format`.
	Format *string `field:"optional" json:"format" yaml:"format"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var authorizers interface{}
//
//   authProperty := &authProperty{
//   	addDefaultAuthorizerToCorsPreflight: jsii.Boolean(false),
//   	authorizers: authorizers,
//   	defaultAuthorizer: jsii.String("defaultAuthorizer"),
//   }
//
type CfnApi_AuthProperty struct {
	// `CfnApi.AuthProperty.AddDefaultAuthorizerToCorsPreflight`.
	AddDefaultAuthorizerToCorsPreflight interface{} `field:"optional" json:"addDefaultAuthorizerToCorsPreflight" yaml:"addDefaultAuthorizerToCorsPreflight"`
	// `CfnApi.AuthProperty.Authorizers`.
	Authorizers interface{} `field:"optional" json:"authorizers" yaml:"authorizers"`
	// `CfnApi.AuthProperty.DefaultAuthorizer`.
	DefaultAuthorizer *string `field:"optional" json:"defaultAuthorizer" yaml:"defaultAuthorizer"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   canarySettingProperty := &canarySettingProperty{
//   	deploymentId: jsii.String("deploymentId"),
//   	percentTraffic: jsii.Number(123),
//   	stageVariableOverrides: map[string]*string{
//   		"stageVariableOverridesKey": jsii.String("stageVariableOverrides"),
//   	},
//   	useStageCache: jsii.Boolean(false),
//   }
//
type CfnApi_CanarySettingProperty struct {
	// `CfnApi.CanarySettingProperty.DeploymentId`.
	DeploymentId *string `field:"optional" json:"deploymentId" yaml:"deploymentId"`
	// `CfnApi.CanarySettingProperty.PercentTraffic`.
	PercentTraffic *float64 `field:"optional" json:"percentTraffic" yaml:"percentTraffic"`
	// `CfnApi.CanarySettingProperty.StageVariableOverrides`.
	StageVariableOverrides interface{} `field:"optional" json:"stageVariableOverrides" yaml:"stageVariableOverrides"`
	// `CfnApi.CanarySettingProperty.UseStageCache`.
	UseStageCache interface{} `field:"optional" json:"useStageCache" yaml:"useStageCache"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   corsConfigurationProperty := &corsConfigurationProperty{
//   	allowOrigin: jsii.String("allowOrigin"),
//
//   	// the properties below are optional
//   	allowCredentials: jsii.Boolean(false),
//   	allowHeaders: jsii.String("allowHeaders"),
//   	allowMethods: jsii.String("allowMethods"),
//   	maxAge: jsii.String("maxAge"),
//   }
//
type CfnApi_CorsConfigurationProperty struct {
	// `CfnApi.CorsConfigurationProperty.AllowOrigin`.
	AllowOrigin *string `field:"required" json:"allowOrigin" yaml:"allowOrigin"`
	// `CfnApi.CorsConfigurationProperty.AllowCredentials`.
	AllowCredentials interface{} `field:"optional" json:"allowCredentials" yaml:"allowCredentials"`
	// `CfnApi.CorsConfigurationProperty.AllowHeaders`.
	AllowHeaders *string `field:"optional" json:"allowHeaders" yaml:"allowHeaders"`
	// `CfnApi.CorsConfigurationProperty.AllowMethods`.
	AllowMethods *string `field:"optional" json:"allowMethods" yaml:"allowMethods"`
	// `CfnApi.CorsConfigurationProperty.MaxAge`.
	MaxAge *string `field:"optional" json:"maxAge" yaml:"maxAge"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   domainConfigurationProperty := &domainConfigurationProperty{
//   	certificateArn: jsii.String("certificateArn"),
//   	domainName: jsii.String("domainName"),
//
//   	// the properties below are optional
//   	basePath: []*string{
//   		jsii.String("basePath"),
//   	},
//   	endpointConfiguration: jsii.String("endpointConfiguration"),
//   	mutualTlsAuthentication: &mutualTlsAuthenticationProperty{
//   		truststoreUri: jsii.String("truststoreUri"),
//   		truststoreVersion: jsii.String("truststoreVersion"),
//   	},
//   	ownershipVerificationCertificateArn: jsii.String("ownershipVerificationCertificateArn"),
//   	route53: &route53ConfigurationProperty{
//   		distributedDomainName: jsii.String("distributedDomainName"),
//   		evaluateTargetHealth: jsii.Boolean(false),
//   		hostedZoneId: jsii.String("hostedZoneId"),
//   		hostedZoneName: jsii.String("hostedZoneName"),
//   		ipV6: jsii.Boolean(false),
//   	},
//   	securityPolicy: jsii.String("securityPolicy"),
//   }
//
type CfnApi_DomainConfigurationProperty struct {
	// `CfnApi.DomainConfigurationProperty.CertificateArn`.
	CertificateArn *string `field:"required" json:"certificateArn" yaml:"certificateArn"`
	// `CfnApi.DomainConfigurationProperty.DomainName`.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// `CfnApi.DomainConfigurationProperty.BasePath`.
	BasePath *[]*string `field:"optional" json:"basePath" yaml:"basePath"`
	// `CfnApi.DomainConfigurationProperty.EndpointConfiguration`.
	EndpointConfiguration *string `field:"optional" json:"endpointConfiguration" yaml:"endpointConfiguration"`
	// `CfnApi.DomainConfigurationProperty.MutualTlsAuthentication`.
	MutualTlsAuthentication interface{} `field:"optional" json:"mutualTlsAuthentication" yaml:"mutualTlsAuthentication"`
	// `CfnApi.DomainConfigurationProperty.OwnershipVerificationCertificateArn`.
	OwnershipVerificationCertificateArn *string `field:"optional" json:"ownershipVerificationCertificateArn" yaml:"ownershipVerificationCertificateArn"`
	// `CfnApi.DomainConfigurationProperty.Route53`.
	Route53 interface{} `field:"optional" json:"route53" yaml:"route53"`
	// `CfnApi.DomainConfigurationProperty.SecurityPolicy`.
	SecurityPolicy *string `field:"optional" json:"securityPolicy" yaml:"securityPolicy"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   endpointConfigurationProperty := &endpointConfigurationProperty{
//   	type: jsii.String("type"),
//   	vpcEndpointIds: []*string{
//   		jsii.String("vpcEndpointIds"),
//   	},
//   }
//
type CfnApi_EndpointConfigurationProperty struct {
	// `CfnApi.EndpointConfigurationProperty.Type`.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// `CfnApi.EndpointConfigurationProperty.VpcEndpointIds`.
	VpcEndpointIds *[]*string `field:"optional" json:"vpcEndpointIds" yaml:"vpcEndpointIds"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mutualTlsAuthenticationProperty := &mutualTlsAuthenticationProperty{
//   	truststoreUri: jsii.String("truststoreUri"),
//   	truststoreVersion: jsii.String("truststoreVersion"),
//   }
//
type CfnApi_MutualTlsAuthenticationProperty struct {
	// `CfnApi.MutualTlsAuthenticationProperty.TruststoreUri`.
	TruststoreUri *string `field:"optional" json:"truststoreUri" yaml:"truststoreUri"`
	// `CfnApi.MutualTlsAuthenticationProperty.TruststoreVersion`.
	TruststoreVersion *string `field:"optional" json:"truststoreVersion" yaml:"truststoreVersion"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   route53ConfigurationProperty := &route53ConfigurationProperty{
//   	distributedDomainName: jsii.String("distributedDomainName"),
//   	evaluateTargetHealth: jsii.Boolean(false),
//   	hostedZoneId: jsii.String("hostedZoneId"),
//   	hostedZoneName: jsii.String("hostedZoneName"),
//   	ipV6: jsii.Boolean(false),
//   }
//
type CfnApi_Route53ConfigurationProperty struct {
	// `CfnApi.Route53ConfigurationProperty.DistributedDomainName`.
	DistributedDomainName *string `field:"optional" json:"distributedDomainName" yaml:"distributedDomainName"`
	// `CfnApi.Route53ConfigurationProperty.EvaluateTargetHealth`.
	EvaluateTargetHealth interface{} `field:"optional" json:"evaluateTargetHealth" yaml:"evaluateTargetHealth"`
	// `CfnApi.Route53ConfigurationProperty.HostedZoneId`.
	HostedZoneId *string `field:"optional" json:"hostedZoneId" yaml:"hostedZoneId"`
	// `CfnApi.Route53ConfigurationProperty.HostedZoneName`.
	HostedZoneName *string `field:"optional" json:"hostedZoneName" yaml:"hostedZoneName"`
	// `CfnApi.Route53ConfigurationProperty.IpV6`.
	IpV6 interface{} `field:"optional" json:"ipV6" yaml:"ipV6"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3LocationProperty := &s3LocationProperty{
//   	bucket: jsii.String("bucket"),
//   	key: jsii.String("key"),
//   	version: jsii.Number(123),
//   }
//
type CfnApi_S3LocationProperty struct {
	// `CfnApi.S3LocationProperty.Bucket`.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// `CfnApi.S3LocationProperty.Key`.
	Key *string `field:"required" json:"key" yaml:"key"`
	// `CfnApi.S3LocationProperty.Version`.
	Version *float64 `field:"required" json:"version" yaml:"version"`
}

// Properties for defining a `CfnApi`.
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
//   cfnApiProps := &cfnApiProps{
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
//   }
//
type CfnApiProps struct {
	// `AWS::Serverless::Api.StageName`.
	StageName *string `field:"required" json:"stageName" yaml:"stageName"`
	// `AWS::Serverless::Api.AccessLogSetting`.
	AccessLogSetting interface{} `field:"optional" json:"accessLogSetting" yaml:"accessLogSetting"`
	// `AWS::Serverless::Api.Auth`.
	Auth interface{} `field:"optional" json:"auth" yaml:"auth"`
	// `AWS::Serverless::Api.BinaryMediaTypes`.
	BinaryMediaTypes *[]*string `field:"optional" json:"binaryMediaTypes" yaml:"binaryMediaTypes"`
	// `AWS::Serverless::Api.CacheClusterEnabled`.
	CacheClusterEnabled interface{} `field:"optional" json:"cacheClusterEnabled" yaml:"cacheClusterEnabled"`
	// `AWS::Serverless::Api.CacheClusterSize`.
	CacheClusterSize *string `field:"optional" json:"cacheClusterSize" yaml:"cacheClusterSize"`
	// `AWS::Serverless::Api.CanarySetting`.
	CanarySetting interface{} `field:"optional" json:"canarySetting" yaml:"canarySetting"`
	// `AWS::Serverless::Api.Cors`.
	Cors interface{} `field:"optional" json:"cors" yaml:"cors"`
	// `AWS::Serverless::Api.DefinitionBody`.
	DefinitionBody interface{} `field:"optional" json:"definitionBody" yaml:"definitionBody"`
	// `AWS::Serverless::Api.DefinitionUri`.
	DefinitionUri interface{} `field:"optional" json:"definitionUri" yaml:"definitionUri"`
	// `AWS::Serverless::Api.Description`.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `AWS::Serverless::Api.Domain`.
	Domain interface{} `field:"optional" json:"domain" yaml:"domain"`
	// `AWS::Serverless::Api.EndpointConfiguration`.
	EndpointConfiguration interface{} `field:"optional" json:"endpointConfiguration" yaml:"endpointConfiguration"`
	// `AWS::Serverless::Api.GatewayResponses`.
	GatewayResponses interface{} `field:"optional" json:"gatewayResponses" yaml:"gatewayResponses"`
	// `AWS::Serverless::Api.MethodSettings`.
	MethodSettings interface{} `field:"optional" json:"methodSettings" yaml:"methodSettings"`
	// `AWS::Serverless::Api.MinimumCompressionSize`.
	MinimumCompressionSize *float64 `field:"optional" json:"minimumCompressionSize" yaml:"minimumCompressionSize"`
	// `AWS::Serverless::Api.Models`.
	Models interface{} `field:"optional" json:"models" yaml:"models"`
	// `AWS::Serverless::Api.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// `AWS::Serverless::Api.OpenApiVersion`.
	OpenApiVersion *string `field:"optional" json:"openApiVersion" yaml:"openApiVersion"`
	// `AWS::Serverless::Api.Tags`.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// `AWS::Serverless::Api.TracingEnabled`.
	TracingEnabled interface{} `field:"optional" json:"tracingEnabled" yaml:"tracingEnabled"`
	// `AWS::Serverless::Api.Variables`.
	Variables interface{} `field:"optional" json:"variables" yaml:"variables"`
}

// A CloudFormation `AWS::Serverless::Application`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplication := awscdk.Aws_sam.NewCfnApplication(this, jsii.String("MyCfnApplication"), &cfnApplicationProps{
//   	location: jsii.String("location"),
//
//   	// the properties below are optional
//   	notificationArns: []*string{
//   		jsii.String("notificationArns"),
//   	},
//   	parameters: map[string]*string{
//   		"parametersKey": jsii.String("parameters"),
//   	},
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	timeoutInMinutes: jsii.Number(123),
//   })
//
type CfnApplication interface {
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
	// `AWS::Serverless::Application.Location`.
	Location() interface{}
	SetLocation(val interface{})
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
	// `AWS::Serverless::Application.NotificationArns`.
	NotificationArns() *[]*string
	SetNotificationArns(val *[]*string)
	// `AWS::Serverless::Application.Parameters`.
	Parameters() interface{}
	SetParameters(val interface{})
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
	// `AWS::Serverless::Application.Tags`.
	Tags() awscdk.TagManager
	// `AWS::Serverless::Application.TimeoutInMinutes`.
	TimeoutInMinutes() *float64
	SetTimeoutInMinutes(val *float64)
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

// The jsii proxy struct for CfnApplication
type jsiiProxy_CfnApplication struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnApplication) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) Location() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) NotificationArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notificationArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) Parameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) TimeoutInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Serverless::Application`.
func NewCfnApplication(scope awscdk.Construct, id *string, props *CfnApplicationProps) CfnApplication {
	_init_.Initialize()

	j := jsiiProxy_CfnApplication{}

	_jsii_.Create(
		"monocdk.aws_sam.CfnApplication",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Serverless::Application`.
func NewCfnApplication_Override(c CfnApplication, scope awscdk.Construct, id *string, props *CfnApplicationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_sam.CfnApplication",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnApplication) SetLocation(val interface{}) {
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_CfnApplication) SetNotificationArns(val *[]*string) {
	_jsii_.Set(
		j,
		"notificationArns",
		val,
	)
}

func (j *jsiiProxy_CfnApplication) SetParameters(val interface{}) {
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_CfnApplication) SetTimeoutInMinutes(val *float64) {
	_jsii_.Set(
		j,
		"timeoutInMinutes",
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
func CfnApplication_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sam.CfnApplication",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnApplication_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sam.CfnApplication",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnApplication_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sam.CfnApplication",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnApplication_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_sam.CfnApplication",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func CfnApplication_REQUIRED_TRANSFORM() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_sam.CfnApplication",
		"REQUIRED_TRANSFORM",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnApplication) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnApplication) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnApplication) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnApplication) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnApplication) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnApplication) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnApplication) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnApplication) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApplication) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApplication) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnApplication) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnApplication) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnApplication) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApplication) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnApplication) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnApplication) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApplication) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApplication) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnApplication) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApplication) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApplication) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   applicationLocationProperty := &applicationLocationProperty{
//   	applicationId: jsii.String("applicationId"),
//   	semanticVersion: jsii.String("semanticVersion"),
//   }
//
type CfnApplication_ApplicationLocationProperty struct {
	// `CfnApplication.ApplicationLocationProperty.ApplicationId`.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// `CfnApplication.ApplicationLocationProperty.SemanticVersion`.
	SemanticVersion *string `field:"required" json:"semanticVersion" yaml:"semanticVersion"`
}

// Properties for defining a `CfnApplication`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationProps := &cfnApplicationProps{
//   	location: jsii.String("location"),
//
//   	// the properties below are optional
//   	notificationArns: []*string{
//   		jsii.String("notificationArns"),
//   	},
//   	parameters: map[string]*string{
//   		"parametersKey": jsii.String("parameters"),
//   	},
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	timeoutInMinutes: jsii.Number(123),
//   }
//
type CfnApplicationProps struct {
	// `AWS::Serverless::Application.Location`.
	Location interface{} `field:"required" json:"location" yaml:"location"`
	// `AWS::Serverless::Application.NotificationArns`.
	NotificationArns *[]*string `field:"optional" json:"notificationArns" yaml:"notificationArns"`
	// `AWS::Serverless::Application.Parameters`.
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// `AWS::Serverless::Application.Tags`.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// `AWS::Serverless::Application.TimeoutInMinutes`.
	TimeoutInMinutes *float64 `field:"optional" json:"timeoutInMinutes" yaml:"timeoutInMinutes"`
}

// A CloudFormation `AWS::Serverless::Function`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var assumeRolePolicyDocument interface{}
//
//   cfnFunction := awscdk.Aws_sam.NewCfnFunction(this, jsii.String("MyCfnFunction"), &cfnFunctionProps{
//   	architectures: []*string{
//   		jsii.String("architectures"),
//   	},
//   	assumeRolePolicyDocument: assumeRolePolicyDocument,
//   	autoPublishAlias: jsii.String("autoPublishAlias"),
//   	autoPublishCodeSha256: jsii.String("autoPublishCodeSha256"),
//   	codeSigningConfigArn: jsii.String("codeSigningConfigArn"),
//   	codeUri: jsii.String("codeUri"),
//   	deadLetterQueue: &deadLetterQueueProperty{
//   		targetArn: jsii.String("targetArn"),
//   		type: jsii.String("type"),
//   	},
//   	deploymentPreference: &deploymentPreferenceProperty{
//   		enabled: jsii.Boolean(false),
//   		type: jsii.String("type"),
//
//   		// the properties below are optional
//   		alarms: []*string{
//   			jsii.String("alarms"),
//   		},
//   		hooks: &hooksProperty{
//   			postTraffic: jsii.String("postTraffic"),
//   			preTraffic: jsii.String("preTraffic"),
//   		},
//   	},
//   	description: jsii.String("description"),
//   	environment: &functionEnvironmentProperty{
//   		variables: map[string]*string{
//   			"variablesKey": jsii.String("variables"),
//   		},
//   	},
//   	eventInvokeConfig: &eventInvokeConfigProperty{
//   		destinationConfig: &eventInvokeDestinationConfigProperty{
//   			onFailure: &destinationProperty{
//   				destination: jsii.String("destination"),
//
//   				// the properties below are optional
//   				type: jsii.String("type"),
//   			},
//   			onSuccess: &destinationProperty{
//   				destination: jsii.String("destination"),
//
//   				// the properties below are optional
//   				type: jsii.String("type"),
//   			},
//   		},
//   		maximumEventAgeInSeconds: jsii.Number(123),
//   		maximumRetryAttempts: jsii.Number(123),
//   	},
//   	events: map[string]interface{}{
//   		"eventsKey": &EventSourceProperty{
//   			"properties": &S3EventProperty{
//   				"variables": map[string]*string{
//   					"variablesKey": jsii.String("variables"),
//   				},
//   			},
//   			"type": jsii.String("type"),
//   		},
//   	},
//   	fileSystemConfigs: []interface{}{
//   		&fileSystemConfigProperty{
//   			arn: jsii.String("arn"),
//   			localMountPath: jsii.String("localMountPath"),
//   		},
//   	},
//   	functionName: jsii.String("functionName"),
//   	handler: jsii.String("handler"),
//   	imageConfig: &imageConfigProperty{
//   		command: []*string{
//   			jsii.String("command"),
//   		},
//   		entryPoint: []*string{
//   			jsii.String("entryPoint"),
//   		},
//   		workingDirectory: jsii.String("workingDirectory"),
//   	},
//   	imageUri: jsii.String("imageUri"),
//   	inlineCode: jsii.String("inlineCode"),
//   	kmsKeyArn: jsii.String("kmsKeyArn"),
//   	layers: []*string{
//   		jsii.String("layers"),
//   	},
//   	memorySize: jsii.Number(123),
//   	packageType: jsii.String("packageType"),
//   	permissionsBoundary: jsii.String("permissionsBoundary"),
//   	policies: jsii.String("policies"),
//   	provisionedConcurrencyConfig: &provisionedConcurrencyConfigProperty{
//   		provisionedConcurrentExecutions: jsii.String("provisionedConcurrentExecutions"),
//   	},
//   	reservedConcurrentExecutions: jsii.Number(123),
//   	role: jsii.String("role"),
//   	runtime: jsii.String("runtime"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	timeout: jsii.Number(123),
//   	tracing: jsii.String("tracing"),
//   	versionDescription: jsii.String("versionDescription"),
//   	vpcConfig: &vpcConfigProperty{
//   		securityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		subnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   })
//
type CfnFunction interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// `AWS::Serverless::Function.Architectures`.
	Architectures() *[]*string
	SetArchitectures(val *[]*string)
	// `AWS::Serverless::Function.AssumeRolePolicyDocument`.
	AssumeRolePolicyDocument() interface{}
	SetAssumeRolePolicyDocument(val interface{})
	// `AWS::Serverless::Function.AutoPublishAlias`.
	AutoPublishAlias() *string
	SetAutoPublishAlias(val *string)
	// `AWS::Serverless::Function.AutoPublishCodeSha256`.
	AutoPublishCodeSha256() *string
	SetAutoPublishCodeSha256(val *string)
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// `AWS::Serverless::Function.CodeSigningConfigArn`.
	CodeSigningConfigArn() *string
	SetCodeSigningConfigArn(val *string)
	// `AWS::Serverless::Function.CodeUri`.
	CodeUri() interface{}
	SetCodeUri(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// `AWS::Serverless::Function.DeadLetterQueue`.
	DeadLetterQueue() interface{}
	SetDeadLetterQueue(val interface{})
	// `AWS::Serverless::Function.DeploymentPreference`.
	DeploymentPreference() interface{}
	SetDeploymentPreference(val interface{})
	// `AWS::Serverless::Function.Description`.
	Description() *string
	SetDescription(val *string)
	// `AWS::Serverless::Function.Environment`.
	Environment() interface{}
	SetEnvironment(val interface{})
	// `AWS::Serverless::Function.EventInvokeConfig`.
	EventInvokeConfig() interface{}
	SetEventInvokeConfig(val interface{})
	// `AWS::Serverless::Function.Events`.
	Events() interface{}
	SetEvents(val interface{})
	// `AWS::Serverless::Function.FileSystemConfigs`.
	FileSystemConfigs() interface{}
	SetFileSystemConfigs(val interface{})
	// `AWS::Serverless::Function.FunctionName`.
	FunctionName() *string
	SetFunctionName(val *string)
	// `AWS::Serverless::Function.Handler`.
	Handler() *string
	SetHandler(val *string)
	// `AWS::Serverless::Function.ImageConfig`.
	ImageConfig() interface{}
	SetImageConfig(val interface{})
	// `AWS::Serverless::Function.ImageUri`.
	ImageUri() *string
	SetImageUri(val *string)
	// `AWS::Serverless::Function.InlineCode`.
	InlineCode() *string
	SetInlineCode(val *string)
	// `AWS::Serverless::Function.KmsKeyArn`.
	KmsKeyArn() *string
	SetKmsKeyArn(val *string)
	// `AWS::Serverless::Function.Layers`.
	Layers() *[]*string
	SetLayers(val *[]*string)
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
	// `AWS::Serverless::Function.MemorySize`.
	MemorySize() *float64
	SetMemorySize(val *float64)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// `AWS::Serverless::Function.PackageType`.
	PackageType() *string
	SetPackageType(val *string)
	// `AWS::Serverless::Function.PermissionsBoundary`.
	PermissionsBoundary() *string
	SetPermissionsBoundary(val *string)
	// `AWS::Serverless::Function.Policies`.
	Policies() interface{}
	SetPolicies(val interface{})
	// `AWS::Serverless::Function.ProvisionedConcurrencyConfig`.
	ProvisionedConcurrencyConfig() interface{}
	SetProvisionedConcurrencyConfig(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// `AWS::Serverless::Function.ReservedConcurrentExecutions`.
	ReservedConcurrentExecutions() *float64
	SetReservedConcurrentExecutions(val *float64)
	// `AWS::Serverless::Function.Role`.
	Role() *string
	SetRole(val *string)
	// `AWS::Serverless::Function.Runtime`.
	Runtime() *string
	SetRuntime(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// `AWS::Serverless::Function.Tags`.
	Tags() awscdk.TagManager
	// `AWS::Serverless::Function.Timeout`.
	Timeout() *float64
	SetTimeout(val *float64)
	// `AWS::Serverless::Function.Tracing`.
	Tracing() *string
	SetTracing(val *string)
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// `AWS::Serverless::Function.VersionDescription`.
	VersionDescription() *string
	SetVersionDescription(val *string)
	// `AWS::Serverless::Function.VpcConfig`.
	VpcConfig() interface{}
	SetVpcConfig(val interface{})
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

// The jsii proxy struct for CfnFunction
type jsiiProxy_CfnFunction struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnFunction) Architectures() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"architectures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) AssumeRolePolicyDocument() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assumeRolePolicyDocument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) AutoPublishAlias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoPublishAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) AutoPublishCodeSha256() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoPublishCodeSha256",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) CodeSigningConfigArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeSigningConfigArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) CodeUri() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"codeUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) DeadLetterQueue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deadLetterQueue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) DeploymentPreference() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deploymentPreference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) Environment() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) EventInvokeConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"eventInvokeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) Events() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"events",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) FileSystemConfigs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fileSystemConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) FunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) Handler() *string {
	var returns *string
	_jsii_.Get(
		j,
		"handler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) ImageConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"imageConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) ImageUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) InlineCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inlineCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) KmsKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) Layers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"layers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) MemorySize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) PackageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) PermissionsBoundary() *string {
	var returns *string
	_jsii_.Get(
		j,
		"permissionsBoundary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) Policies() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"policies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) ProvisionedConcurrencyConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"provisionedConcurrencyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) ReservedConcurrentExecutions() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"reservedConcurrentExecutions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) Runtime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) Timeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) Tracing() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tracing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) VersionDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) VpcConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpcConfig",
		&returns,
	)
	return returns
}


// Create a new `AWS::Serverless::Function`.
func NewCfnFunction(scope awscdk.Construct, id *string, props *CfnFunctionProps) CfnFunction {
	_init_.Initialize()

	j := jsiiProxy_CfnFunction{}

	_jsii_.Create(
		"monocdk.aws_sam.CfnFunction",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Serverless::Function`.
func NewCfnFunction_Override(c CfnFunction, scope awscdk.Construct, id *string, props *CfnFunctionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_sam.CfnFunction",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnFunction) SetArchitectures(val *[]*string) {
	_jsii_.Set(
		j,
		"architectures",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetAssumeRolePolicyDocument(val interface{}) {
	_jsii_.Set(
		j,
		"assumeRolePolicyDocument",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetAutoPublishAlias(val *string) {
	_jsii_.Set(
		j,
		"autoPublishAlias",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetAutoPublishCodeSha256(val *string) {
	_jsii_.Set(
		j,
		"autoPublishCodeSha256",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetCodeSigningConfigArn(val *string) {
	_jsii_.Set(
		j,
		"codeSigningConfigArn",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetCodeUri(val interface{}) {
	_jsii_.Set(
		j,
		"codeUri",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetDeadLetterQueue(val interface{}) {
	_jsii_.Set(
		j,
		"deadLetterQueue",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetDeploymentPreference(val interface{}) {
	_jsii_.Set(
		j,
		"deploymentPreference",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetEnvironment(val interface{}) {
	_jsii_.Set(
		j,
		"environment",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetEventInvokeConfig(val interface{}) {
	_jsii_.Set(
		j,
		"eventInvokeConfig",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetEvents(val interface{}) {
	_jsii_.Set(
		j,
		"events",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetFileSystemConfigs(val interface{}) {
	_jsii_.Set(
		j,
		"fileSystemConfigs",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetFunctionName(val *string) {
	_jsii_.Set(
		j,
		"functionName",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetHandler(val *string) {
	_jsii_.Set(
		j,
		"handler",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetImageConfig(val interface{}) {
	_jsii_.Set(
		j,
		"imageConfig",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetImageUri(val *string) {
	_jsii_.Set(
		j,
		"imageUri",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetInlineCode(val *string) {
	_jsii_.Set(
		j,
		"inlineCode",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetKmsKeyArn(val *string) {
	_jsii_.Set(
		j,
		"kmsKeyArn",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetLayers(val *[]*string) {
	_jsii_.Set(
		j,
		"layers",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetMemorySize(val *float64) {
	_jsii_.Set(
		j,
		"memorySize",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetPackageType(val *string) {
	_jsii_.Set(
		j,
		"packageType",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetPermissionsBoundary(val *string) {
	_jsii_.Set(
		j,
		"permissionsBoundary",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetPolicies(val interface{}) {
	_jsii_.Set(
		j,
		"policies",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetProvisionedConcurrencyConfig(val interface{}) {
	_jsii_.Set(
		j,
		"provisionedConcurrencyConfig",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetReservedConcurrentExecutions(val *float64) {
	_jsii_.Set(
		j,
		"reservedConcurrentExecutions",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetRole(val *string) {
	_jsii_.Set(
		j,
		"role",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetRuntime(val *string) {
	_jsii_.Set(
		j,
		"runtime",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetTimeout(val *float64) {
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetTracing(val *string) {
	_jsii_.Set(
		j,
		"tracing",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetVersionDescription(val *string) {
	_jsii_.Set(
		j,
		"versionDescription",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetVpcConfig(val interface{}) {
	_jsii_.Set(
		j,
		"vpcConfig",
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
func CfnFunction_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sam.CfnFunction",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnFunction_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sam.CfnFunction",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnFunction_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sam.CfnFunction",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFunction_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_sam.CfnFunction",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func CfnFunction_REQUIRED_TRANSFORM() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_sam.CfnFunction",
		"REQUIRED_TRANSFORM",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnFunction) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnFunction) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnFunction) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnFunction) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnFunction) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnFunction) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnFunction) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnFunction) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFunction) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFunction) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnFunction) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnFunction) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnFunction) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFunction) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnFunction) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnFunction) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFunction) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFunction) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnFunction) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFunction) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFunction) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   alexaSkillEventProperty := &alexaSkillEventProperty{
//   	variables: map[string]*string{
//   		"variablesKey": jsii.String("variables"),
//   	},
//   }
//
type CfnFunction_AlexaSkillEventProperty struct {
	// `CfnFunction.AlexaSkillEventProperty.Variables`.
	Variables interface{} `field:"optional" json:"variables" yaml:"variables"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var customStatements interface{}
//
//   apiEventProperty := &apiEventProperty{
//   	method: jsii.String("method"),
//   	path: jsii.String("path"),
//
//   	// the properties below are optional
//   	auth: &authProperty{
//   		apiKeyRequired: jsii.Boolean(false),
//   		authorizationScopes: []*string{
//   			jsii.String("authorizationScopes"),
//   		},
//   		authorizer: jsii.String("authorizer"),
//   		resourcePolicy: &authResourcePolicyProperty{
//   			awsAccountBlacklist: []*string{
//   				jsii.String("awsAccountBlacklist"),
//   			},
//   			awsAccountWhitelist: []*string{
//   				jsii.String("awsAccountWhitelist"),
//   			},
//   			customStatements: []interface{}{
//   				customStatements,
//   			},
//   			intrinsicVpcBlacklist: []*string{
//   				jsii.String("intrinsicVpcBlacklist"),
//   			},
//   			intrinsicVpceBlacklist: []*string{
//   				jsii.String("intrinsicVpceBlacklist"),
//   			},
//   			intrinsicVpceWhitelist: []*string{
//   				jsii.String("intrinsicVpceWhitelist"),
//   			},
//   			intrinsicVpcWhitelist: []*string{
//   				jsii.String("intrinsicVpcWhitelist"),
//   			},
//   			ipRangeBlacklist: []*string{
//   				jsii.String("ipRangeBlacklist"),
//   			},
//   			ipRangeWhitelist: []*string{
//   				jsii.String("ipRangeWhitelist"),
//   			},
//   			sourceVpcBlacklist: []*string{
//   				jsii.String("sourceVpcBlacklist"),
//   			},
//   			sourceVpcWhitelist: []*string{
//   				jsii.String("sourceVpcWhitelist"),
//   			},
//   		},
//   	},
//   	requestModel: &requestModelProperty{
//   		model: jsii.String("model"),
//
//   		// the properties below are optional
//   		required: jsii.Boolean(false),
//   		validateBody: jsii.Boolean(false),
//   		validateParameters: jsii.Boolean(false),
//   	},
//   	requestParameters: []interface{}{
//   		jsii.String("requestParameters"),
//   	},
//   	restApiId: jsii.String("restApiId"),
//   }
//
type CfnFunction_ApiEventProperty struct {
	// `CfnFunction.ApiEventProperty.Method`.
	Method *string `field:"required" json:"method" yaml:"method"`
	// `CfnFunction.ApiEventProperty.Path`.
	Path *string `field:"required" json:"path" yaml:"path"`
	// `CfnFunction.ApiEventProperty.Auth`.
	Auth interface{} `field:"optional" json:"auth" yaml:"auth"`
	// `CfnFunction.ApiEventProperty.RequestModel`.
	RequestModel interface{} `field:"optional" json:"requestModel" yaml:"requestModel"`
	// `CfnFunction.ApiEventProperty.RequestParameters`.
	RequestParameters interface{} `field:"optional" json:"requestParameters" yaml:"requestParameters"`
	// `CfnFunction.ApiEventProperty.RestApiId`.
	RestApiId *string `field:"optional" json:"restApiId" yaml:"restApiId"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var customStatements interface{}
//
//   authProperty := &authProperty{
//   	apiKeyRequired: jsii.Boolean(false),
//   	authorizationScopes: []*string{
//   		jsii.String("authorizationScopes"),
//   	},
//   	authorizer: jsii.String("authorizer"),
//   	resourcePolicy: &authResourcePolicyProperty{
//   		awsAccountBlacklist: []*string{
//   			jsii.String("awsAccountBlacklist"),
//   		},
//   		awsAccountWhitelist: []*string{
//   			jsii.String("awsAccountWhitelist"),
//   		},
//   		customStatements: []interface{}{
//   			customStatements,
//   		},
//   		intrinsicVpcBlacklist: []*string{
//   			jsii.String("intrinsicVpcBlacklist"),
//   		},
//   		intrinsicVpceBlacklist: []*string{
//   			jsii.String("intrinsicVpceBlacklist"),
//   		},
//   		intrinsicVpceWhitelist: []*string{
//   			jsii.String("intrinsicVpceWhitelist"),
//   		},
//   		intrinsicVpcWhitelist: []*string{
//   			jsii.String("intrinsicVpcWhitelist"),
//   		},
//   		ipRangeBlacklist: []*string{
//   			jsii.String("ipRangeBlacklist"),
//   		},
//   		ipRangeWhitelist: []*string{
//   			jsii.String("ipRangeWhitelist"),
//   		},
//   		sourceVpcBlacklist: []*string{
//   			jsii.String("sourceVpcBlacklist"),
//   		},
//   		sourceVpcWhitelist: []*string{
//   			jsii.String("sourceVpcWhitelist"),
//   		},
//   	},
//   }
//
type CfnFunction_AuthProperty struct {
	// `CfnFunction.AuthProperty.ApiKeyRequired`.
	ApiKeyRequired interface{} `field:"optional" json:"apiKeyRequired" yaml:"apiKeyRequired"`
	// `CfnFunction.AuthProperty.AuthorizationScopes`.
	AuthorizationScopes *[]*string `field:"optional" json:"authorizationScopes" yaml:"authorizationScopes"`
	// `CfnFunction.AuthProperty.Authorizer`.
	Authorizer *string `field:"optional" json:"authorizer" yaml:"authorizer"`
	// `CfnFunction.AuthProperty.ResourcePolicy`.
	ResourcePolicy interface{} `field:"optional" json:"resourcePolicy" yaml:"resourcePolicy"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var customStatements interface{}
//
//   authResourcePolicyProperty := &authResourcePolicyProperty{
//   	awsAccountBlacklist: []*string{
//   		jsii.String("awsAccountBlacklist"),
//   	},
//   	awsAccountWhitelist: []*string{
//   		jsii.String("awsAccountWhitelist"),
//   	},
//   	customStatements: []interface{}{
//   		customStatements,
//   	},
//   	intrinsicVpcBlacklist: []*string{
//   		jsii.String("intrinsicVpcBlacklist"),
//   	},
//   	intrinsicVpceBlacklist: []*string{
//   		jsii.String("intrinsicVpceBlacklist"),
//   	},
//   	intrinsicVpceWhitelist: []*string{
//   		jsii.String("intrinsicVpceWhitelist"),
//   	},
//   	intrinsicVpcWhitelist: []*string{
//   		jsii.String("intrinsicVpcWhitelist"),
//   	},
//   	ipRangeBlacklist: []*string{
//   		jsii.String("ipRangeBlacklist"),
//   	},
//   	ipRangeWhitelist: []*string{
//   		jsii.String("ipRangeWhitelist"),
//   	},
//   	sourceVpcBlacklist: []*string{
//   		jsii.String("sourceVpcBlacklist"),
//   	},
//   	sourceVpcWhitelist: []*string{
//   		jsii.String("sourceVpcWhitelist"),
//   	},
//   }
//
type CfnFunction_AuthResourcePolicyProperty struct {
	// `CfnFunction.AuthResourcePolicyProperty.AwsAccountBlacklist`.
	AwsAccountBlacklist *[]*string `field:"optional" json:"awsAccountBlacklist" yaml:"awsAccountBlacklist"`
	// `CfnFunction.AuthResourcePolicyProperty.AwsAccountWhitelist`.
	AwsAccountWhitelist *[]*string `field:"optional" json:"awsAccountWhitelist" yaml:"awsAccountWhitelist"`
	// `CfnFunction.AuthResourcePolicyProperty.CustomStatements`.
	CustomStatements interface{} `field:"optional" json:"customStatements" yaml:"customStatements"`
	// `CfnFunction.AuthResourcePolicyProperty.IntrinsicVpcBlacklist`.
	IntrinsicVpcBlacklist *[]*string `field:"optional" json:"intrinsicVpcBlacklist" yaml:"intrinsicVpcBlacklist"`
	// `CfnFunction.AuthResourcePolicyProperty.IntrinsicVpceBlacklist`.
	IntrinsicVpceBlacklist *[]*string `field:"optional" json:"intrinsicVpceBlacklist" yaml:"intrinsicVpceBlacklist"`
	// `CfnFunction.AuthResourcePolicyProperty.IntrinsicVpceWhitelist`.
	IntrinsicVpceWhitelist *[]*string `field:"optional" json:"intrinsicVpceWhitelist" yaml:"intrinsicVpceWhitelist"`
	// `CfnFunction.AuthResourcePolicyProperty.IntrinsicVpcWhitelist`.
	IntrinsicVpcWhitelist *[]*string `field:"optional" json:"intrinsicVpcWhitelist" yaml:"intrinsicVpcWhitelist"`
	// `CfnFunction.AuthResourcePolicyProperty.IpRangeBlacklist`.
	IpRangeBlacklist *[]*string `field:"optional" json:"ipRangeBlacklist" yaml:"ipRangeBlacklist"`
	// `CfnFunction.AuthResourcePolicyProperty.IpRangeWhitelist`.
	IpRangeWhitelist *[]*string `field:"optional" json:"ipRangeWhitelist" yaml:"ipRangeWhitelist"`
	// `CfnFunction.AuthResourcePolicyProperty.SourceVpcBlacklist`.
	SourceVpcBlacklist *[]*string `field:"optional" json:"sourceVpcBlacklist" yaml:"sourceVpcBlacklist"`
	// `CfnFunction.AuthResourcePolicyProperty.SourceVpcWhitelist`.
	SourceVpcWhitelist *[]*string `field:"optional" json:"sourceVpcWhitelist" yaml:"sourceVpcWhitelist"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   bucketSAMPTProperty := &bucketSAMPTProperty{
//   	bucketName: jsii.String("bucketName"),
//   }
//
type CfnFunction_BucketSAMPTProperty struct {
	// `CfnFunction.BucketSAMPTProperty.BucketName`.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var pattern interface{}
//
//   cloudWatchEventEventProperty := &cloudWatchEventEventProperty{
//   	pattern: pattern,
//
//   	// the properties below are optional
//   	input: jsii.String("input"),
//   	inputPath: jsii.String("inputPath"),
//   }
//
type CfnFunction_CloudWatchEventEventProperty struct {
	// `CfnFunction.CloudWatchEventEventProperty.Pattern`.
	Pattern interface{} `field:"required" json:"pattern" yaml:"pattern"`
	// `CfnFunction.CloudWatchEventEventProperty.Input`.
	Input *string `field:"optional" json:"input" yaml:"input"`
	// `CfnFunction.CloudWatchEventEventProperty.InputPath`.
	InputPath *string `field:"optional" json:"inputPath" yaml:"inputPath"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudWatchLogsEventProperty := &cloudWatchLogsEventProperty{
//   	filterPattern: jsii.String("filterPattern"),
//   	logGroupName: jsii.String("logGroupName"),
//   }
//
type CfnFunction_CloudWatchLogsEventProperty struct {
	// `CfnFunction.CloudWatchLogsEventProperty.FilterPattern`.
	FilterPattern *string `field:"required" json:"filterPattern" yaml:"filterPattern"`
	// `CfnFunction.CloudWatchLogsEventProperty.LogGroupName`.
	LogGroupName *string `field:"required" json:"logGroupName" yaml:"logGroupName"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   collectionSAMPTProperty := &collectionSAMPTProperty{
//   	collectionId: jsii.String("collectionId"),
//   }
//
type CfnFunction_CollectionSAMPTProperty struct {
	// `CfnFunction.CollectionSAMPTProperty.CollectionId`.
	CollectionId *string `field:"required" json:"collectionId" yaml:"collectionId"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deadLetterQueueProperty := &deadLetterQueueProperty{
//   	targetArn: jsii.String("targetArn"),
//   	type: jsii.String("type"),
//   }
//
type CfnFunction_DeadLetterQueueProperty struct {
	// `CfnFunction.DeadLetterQueueProperty.TargetArn`.
	TargetArn *string `field:"required" json:"targetArn" yaml:"targetArn"`
	// `CfnFunction.DeadLetterQueueProperty.Type`.
	Type *string `field:"required" json:"type" yaml:"type"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deploymentPreferenceProperty := &deploymentPreferenceProperty{
//   	enabled: jsii.Boolean(false),
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	alarms: []*string{
//   		jsii.String("alarms"),
//   	},
//   	hooks: &hooksProperty{
//   		postTraffic: jsii.String("postTraffic"),
//   		preTraffic: jsii.String("preTraffic"),
//   	},
//   }
//
type CfnFunction_DeploymentPreferenceProperty struct {
	// `CfnFunction.DeploymentPreferenceProperty.Enabled`.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// `CfnFunction.DeploymentPreferenceProperty.Type`.
	Type *string `field:"required" json:"type" yaml:"type"`
	// `CfnFunction.DeploymentPreferenceProperty.Alarms`.
	Alarms *[]*string `field:"optional" json:"alarms" yaml:"alarms"`
	// `CfnFunction.DeploymentPreferenceProperty.Hooks`.
	Hooks interface{} `field:"optional" json:"hooks" yaml:"hooks"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   destinationConfigProperty := &destinationConfigProperty{
//   	onFailure: &destinationProperty{
//   		destination: jsii.String("destination"),
//
//   		// the properties below are optional
//   		type: jsii.String("type"),
//   	},
//   }
//
type CfnFunction_DestinationConfigProperty struct {
	// `CfnFunction.DestinationConfigProperty.OnFailure`.
	OnFailure interface{} `field:"required" json:"onFailure" yaml:"onFailure"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   destinationProperty := &destinationProperty{
//   	destination: jsii.String("destination"),
//
//   	// the properties below are optional
//   	type: jsii.String("type"),
//   }
//
type CfnFunction_DestinationProperty struct {
	// `CfnFunction.DestinationProperty.Destination`.
	Destination *string `field:"required" json:"destination" yaml:"destination"`
	// `CfnFunction.DestinationProperty.Type`.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   domainSAMPTProperty := &domainSAMPTProperty{
//   	domainName: jsii.String("domainName"),
//   }
//
type CfnFunction_DomainSAMPTProperty struct {
	// `CfnFunction.DomainSAMPTProperty.DomainName`.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dynamoDBEventProperty := &dynamoDBEventProperty{
//   	startingPosition: jsii.String("startingPosition"),
//   	stream: jsii.String("stream"),
//
//   	// the properties below are optional
//   	batchSize: jsii.Number(123),
//   	bisectBatchOnFunctionError: jsii.Boolean(false),
//   	destinationConfig: &destinationConfigProperty{
//   		onFailure: &destinationProperty{
//   			destination: jsii.String("destination"),
//
//   			// the properties below are optional
//   			type: jsii.String("type"),
//   		},
//   	},
//   	enabled: jsii.Boolean(false),
//   	maximumBatchingWindowInSeconds: jsii.Number(123),
//   	maximumRecordAgeInSeconds: jsii.Number(123),
//   	maximumRetryAttempts: jsii.Number(123),
//   	parallelizationFactor: jsii.Number(123),
//   }
//
type CfnFunction_DynamoDBEventProperty struct {
	// `CfnFunction.DynamoDBEventProperty.StartingPosition`.
	StartingPosition *string `field:"required" json:"startingPosition" yaml:"startingPosition"`
	// `CfnFunction.DynamoDBEventProperty.Stream`.
	Stream *string `field:"required" json:"stream" yaml:"stream"`
	// `CfnFunction.DynamoDBEventProperty.BatchSize`.
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// `CfnFunction.DynamoDBEventProperty.BisectBatchOnFunctionError`.
	BisectBatchOnFunctionError interface{} `field:"optional" json:"bisectBatchOnFunctionError" yaml:"bisectBatchOnFunctionError"`
	// `CfnFunction.DynamoDBEventProperty.DestinationConfig`.
	DestinationConfig interface{} `field:"optional" json:"destinationConfig" yaml:"destinationConfig"`
	// `CfnFunction.DynamoDBEventProperty.Enabled`.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// `CfnFunction.DynamoDBEventProperty.MaximumBatchingWindowInSeconds`.
	MaximumBatchingWindowInSeconds *float64 `field:"optional" json:"maximumBatchingWindowInSeconds" yaml:"maximumBatchingWindowInSeconds"`
	// `CfnFunction.DynamoDBEventProperty.MaximumRecordAgeInSeconds`.
	MaximumRecordAgeInSeconds *float64 `field:"optional" json:"maximumRecordAgeInSeconds" yaml:"maximumRecordAgeInSeconds"`
	// `CfnFunction.DynamoDBEventProperty.MaximumRetryAttempts`.
	MaximumRetryAttempts *float64 `field:"optional" json:"maximumRetryAttempts" yaml:"maximumRetryAttempts"`
	// `CfnFunction.DynamoDBEventProperty.ParallelizationFactor`.
	ParallelizationFactor *float64 `field:"optional" json:"parallelizationFactor" yaml:"parallelizationFactor"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   emptySAMPTProperty := &emptySAMPTProperty{
//   }
//
type CfnFunction_EmptySAMPTProperty struct {
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var pattern interface{}
//
//   eventBridgeRuleEventProperty := &eventBridgeRuleEventProperty{
//   	pattern: pattern,
//
//   	// the properties below are optional
//   	eventBusName: jsii.String("eventBusName"),
//   	input: jsii.String("input"),
//   	inputPath: jsii.String("inputPath"),
//   }
//
type CfnFunction_EventBridgeRuleEventProperty struct {
	// `CfnFunction.EventBridgeRuleEventProperty.Pattern`.
	Pattern interface{} `field:"required" json:"pattern" yaml:"pattern"`
	// `CfnFunction.EventBridgeRuleEventProperty.EventBusName`.
	EventBusName *string `field:"optional" json:"eventBusName" yaml:"eventBusName"`
	// `CfnFunction.EventBridgeRuleEventProperty.Input`.
	Input *string `field:"optional" json:"input" yaml:"input"`
	// `CfnFunction.EventBridgeRuleEventProperty.InputPath`.
	InputPath *string `field:"optional" json:"inputPath" yaml:"inputPath"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventInvokeConfigProperty := &eventInvokeConfigProperty{
//   	destinationConfig: &eventInvokeDestinationConfigProperty{
//   		onFailure: &destinationProperty{
//   			destination: jsii.String("destination"),
//
//   			// the properties below are optional
//   			type: jsii.String("type"),
//   		},
//   		onSuccess: &destinationProperty{
//   			destination: jsii.String("destination"),
//
//   			// the properties below are optional
//   			type: jsii.String("type"),
//   		},
//   	},
//   	maximumEventAgeInSeconds: jsii.Number(123),
//   	maximumRetryAttempts: jsii.Number(123),
//   }
//
type CfnFunction_EventInvokeConfigProperty struct {
	// `CfnFunction.EventInvokeConfigProperty.DestinationConfig`.
	DestinationConfig interface{} `field:"optional" json:"destinationConfig" yaml:"destinationConfig"`
	// `CfnFunction.EventInvokeConfigProperty.MaximumEventAgeInSeconds`.
	MaximumEventAgeInSeconds *float64 `field:"optional" json:"maximumEventAgeInSeconds" yaml:"maximumEventAgeInSeconds"`
	// `CfnFunction.EventInvokeConfigProperty.MaximumRetryAttempts`.
	MaximumRetryAttempts *float64 `field:"optional" json:"maximumRetryAttempts" yaml:"maximumRetryAttempts"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventInvokeDestinationConfigProperty := &eventInvokeDestinationConfigProperty{
//   	onFailure: &destinationProperty{
//   		destination: jsii.String("destination"),
//
//   		// the properties below are optional
//   		type: jsii.String("type"),
//   	},
//   	onSuccess: &destinationProperty{
//   		destination: jsii.String("destination"),
//
//   		// the properties below are optional
//   		type: jsii.String("type"),
//   	},
//   }
//
type CfnFunction_EventInvokeDestinationConfigProperty struct {
	// `CfnFunction.EventInvokeDestinationConfigProperty.OnFailure`.
	OnFailure interface{} `field:"required" json:"onFailure" yaml:"onFailure"`
	// `CfnFunction.EventInvokeDestinationConfigProperty.OnSuccess`.
	OnSuccess interface{} `field:"required" json:"onSuccess" yaml:"onSuccess"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventSourceProperty := &eventSourceProperty{
//   	properties: &s3EventProperty{
//   		variables: map[string]*string{
//   			"variablesKey": jsii.String("variables"),
//   		},
//   	},
//   	type: jsii.String("type"),
//   }
//
type CfnFunction_EventSourceProperty struct {
	// `CfnFunction.EventSourceProperty.Properties`.
	Properties interface{} `field:"required" json:"properties" yaml:"properties"`
	// `CfnFunction.EventSourceProperty.Type`.
	Type *string `field:"required" json:"type" yaml:"type"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fileSystemConfigProperty := &fileSystemConfigProperty{
//   	arn: jsii.String("arn"),
//   	localMountPath: jsii.String("localMountPath"),
//   }
//
type CfnFunction_FileSystemConfigProperty struct {
	// `CfnFunction.FileSystemConfigProperty.Arn`.
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// `CfnFunction.FileSystemConfigProperty.LocalMountPath`.
	LocalMountPath *string `field:"optional" json:"localMountPath" yaml:"localMountPath"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   functionEnvironmentProperty := &functionEnvironmentProperty{
//   	variables: map[string]*string{
//   		"variablesKey": jsii.String("variables"),
//   	},
//   }
//
type CfnFunction_FunctionEnvironmentProperty struct {
	// `CfnFunction.FunctionEnvironmentProperty.Variables`.
	Variables interface{} `field:"required" json:"variables" yaml:"variables"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   functionSAMPTProperty := &functionSAMPTProperty{
//   	functionName: jsii.String("functionName"),
//   }
//
type CfnFunction_FunctionSAMPTProperty struct {
	// `CfnFunction.FunctionSAMPTProperty.FunctionName`.
	FunctionName *string `field:"required" json:"functionName" yaml:"functionName"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hooksProperty := &hooksProperty{
//   	postTraffic: jsii.String("postTraffic"),
//   	preTraffic: jsii.String("preTraffic"),
//   }
//
type CfnFunction_HooksProperty struct {
	// `CfnFunction.HooksProperty.PostTraffic`.
	PostTraffic *string `field:"optional" json:"postTraffic" yaml:"postTraffic"`
	// `CfnFunction.HooksProperty.PreTraffic`.
	PreTraffic *string `field:"optional" json:"preTraffic" yaml:"preTraffic"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var statement interface{}
//
//   iAMPolicyDocumentProperty := map[string]interface{}{
//   	"statement": statement,
//   	"version": jsii.String("version"),
//   }
//
type CfnFunction_IAMPolicyDocumentProperty struct {
	// `CfnFunction.IAMPolicyDocumentProperty.Statement`.
	Statement interface{} `field:"required" json:"statement" yaml:"statement"`
	// `CfnFunction.IAMPolicyDocumentProperty.Version`.
	Version *string `field:"required" json:"version" yaml:"version"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   identitySAMPTProperty := &identitySAMPTProperty{
//   	identityName: jsii.String("identityName"),
//   }
//
type CfnFunction_IdentitySAMPTProperty struct {
	// `CfnFunction.IdentitySAMPTProperty.IdentityName`.
	IdentityName *string `field:"required" json:"identityName" yaml:"identityName"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   imageConfigProperty := &imageConfigProperty{
//   	command: []*string{
//   		jsii.String("command"),
//   	},
//   	entryPoint: []*string{
//   		jsii.String("entryPoint"),
//   	},
//   	workingDirectory: jsii.String("workingDirectory"),
//   }
//
type CfnFunction_ImageConfigProperty struct {
	// `CfnFunction.ImageConfigProperty.Command`.
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// `CfnFunction.ImageConfigProperty.EntryPoint`.
	EntryPoint *[]*string `field:"optional" json:"entryPoint" yaml:"entryPoint"`
	// `CfnFunction.ImageConfigProperty.WorkingDirectory`.
	WorkingDirectory *string `field:"optional" json:"workingDirectory" yaml:"workingDirectory"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ioTRuleEventProperty := &ioTRuleEventProperty{
//   	sql: jsii.String("sql"),
//
//   	// the properties below are optional
//   	awsIotSqlVersion: jsii.String("awsIotSqlVersion"),
//   }
//
type CfnFunction_IoTRuleEventProperty struct {
	// `CfnFunction.IoTRuleEventProperty.Sql`.
	Sql *string `field:"required" json:"sql" yaml:"sql"`
	// `CfnFunction.IoTRuleEventProperty.AwsIotSqlVersion`.
	AwsIotSqlVersion *string `field:"optional" json:"awsIotSqlVersion" yaml:"awsIotSqlVersion"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   keySAMPTProperty := &keySAMPTProperty{
//   	keyId: jsii.String("keyId"),
//   }
//
type CfnFunction_KeySAMPTProperty struct {
	// `CfnFunction.KeySAMPTProperty.KeyId`.
	KeyId *string `field:"required" json:"keyId" yaml:"keyId"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kinesisEventProperty := &kinesisEventProperty{
//   	startingPosition: jsii.String("startingPosition"),
//   	stream: jsii.String("stream"),
//
//   	// the properties below are optional
//   	batchSize: jsii.Number(123),
//   	enabled: jsii.Boolean(false),
//   }
//
type CfnFunction_KinesisEventProperty struct {
	// `CfnFunction.KinesisEventProperty.StartingPosition`.
	StartingPosition *string `field:"required" json:"startingPosition" yaml:"startingPosition"`
	// `CfnFunction.KinesisEventProperty.Stream`.
	Stream *string `field:"required" json:"stream" yaml:"stream"`
	// `CfnFunction.KinesisEventProperty.BatchSize`.
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// `CfnFunction.KinesisEventProperty.Enabled`.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   logGroupSAMPTProperty := &logGroupSAMPTProperty{
//   	logGroupName: jsii.String("logGroupName"),
//   }
//
type CfnFunction_LogGroupSAMPTProperty struct {
	// `CfnFunction.LogGroupSAMPTProperty.LogGroupName`.
	LogGroupName *string `field:"required" json:"logGroupName" yaml:"logGroupName"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parameterNameSAMPTProperty := &parameterNameSAMPTProperty{
//   	parameterName: jsii.String("parameterName"),
//   }
//
type CfnFunction_ParameterNameSAMPTProperty struct {
	// `CfnFunction.ParameterNameSAMPTProperty.ParameterName`.
	ParameterName *string `field:"required" json:"parameterName" yaml:"parameterName"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   provisionedConcurrencyConfigProperty := &provisionedConcurrencyConfigProperty{
//   	provisionedConcurrentExecutions: jsii.String("provisionedConcurrentExecutions"),
//   }
//
type CfnFunction_ProvisionedConcurrencyConfigProperty struct {
	// `CfnFunction.ProvisionedConcurrencyConfigProperty.ProvisionedConcurrentExecutions`.
	ProvisionedConcurrentExecutions *string `field:"required" json:"provisionedConcurrentExecutions" yaml:"provisionedConcurrentExecutions"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   queueSAMPTProperty := &queueSAMPTProperty{
//   	queueName: jsii.String("queueName"),
//   }
//
type CfnFunction_QueueSAMPTProperty struct {
	// `CfnFunction.QueueSAMPTProperty.QueueName`.
	QueueName *string `field:"required" json:"queueName" yaml:"queueName"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   requestModelProperty := &requestModelProperty{
//   	model: jsii.String("model"),
//
//   	// the properties below are optional
//   	required: jsii.Boolean(false),
//   	validateBody: jsii.Boolean(false),
//   	validateParameters: jsii.Boolean(false),
//   }
//
type CfnFunction_RequestModelProperty struct {
	// `CfnFunction.RequestModelProperty.Model`.
	Model *string `field:"required" json:"model" yaml:"model"`
	// `CfnFunction.RequestModelProperty.Required`.
	Required interface{} `field:"optional" json:"required" yaml:"required"`
	// `CfnFunction.RequestModelProperty.ValidateBody`.
	ValidateBody interface{} `field:"optional" json:"validateBody" yaml:"validateBody"`
	// `CfnFunction.RequestModelProperty.ValidateParameters`.
	ValidateParameters interface{} `field:"optional" json:"validateParameters" yaml:"validateParameters"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   requestParameterProperty := &requestParameterProperty{
//   	caching: jsii.Boolean(false),
//   	required: jsii.Boolean(false),
//   }
//
type CfnFunction_RequestParameterProperty struct {
	// `CfnFunction.RequestParameterProperty.Caching`.
	Caching interface{} `field:"optional" json:"caching" yaml:"caching"`
	// `CfnFunction.RequestParameterProperty.Required`.
	Required interface{} `field:"optional" json:"required" yaml:"required"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3EventProperty := &s3EventProperty{
//   	bucket: jsii.String("bucket"),
//   	events: jsii.String("events"),
//
//   	// the properties below are optional
//   	filter: &s3NotificationFilterProperty{
//   		s3Key: &s3KeyFilterProperty{
//   			rules: []interface{}{
//   				&s3KeyFilterRuleProperty{
//   					name: jsii.String("name"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnFunction_S3EventProperty struct {
	// `CfnFunction.S3EventProperty.Bucket`.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// `CfnFunction.S3EventProperty.Events`.
	Events interface{} `field:"required" json:"events" yaml:"events"`
	// `CfnFunction.S3EventProperty.Filter`.
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3KeyFilterProperty := &s3KeyFilterProperty{
//   	rules: []interface{}{
//   		&s3KeyFilterRuleProperty{
//   			name: jsii.String("name"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnFunction_S3KeyFilterProperty struct {
	// `CfnFunction.S3KeyFilterProperty.Rules`.
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3KeyFilterRuleProperty := &s3KeyFilterRuleProperty{
//   	name: jsii.String("name"),
//   	value: jsii.String("value"),
//   }
//
type CfnFunction_S3KeyFilterRuleProperty struct {
	// `CfnFunction.S3KeyFilterRuleProperty.Name`.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `CfnFunction.S3KeyFilterRuleProperty.Value`.
	Value *string `field:"required" json:"value" yaml:"value"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3LocationProperty := &s3LocationProperty{
//   	bucket: jsii.String("bucket"),
//   	key: jsii.String("key"),
//
//   	// the properties below are optional
//   	version: jsii.Number(123),
//   }
//
type CfnFunction_S3LocationProperty struct {
	// `CfnFunction.S3LocationProperty.Bucket`.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// `CfnFunction.S3LocationProperty.Key`.
	Key *string `field:"required" json:"key" yaml:"key"`
	// `CfnFunction.S3LocationProperty.Version`.
	Version *float64 `field:"optional" json:"version" yaml:"version"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3NotificationFilterProperty := &s3NotificationFilterProperty{
//   	s3Key: &s3KeyFilterProperty{
//   		rules: []interface{}{
//   			&s3KeyFilterRuleProperty{
//   				name: jsii.String("name"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   	},
//   }
//
type CfnFunction_S3NotificationFilterProperty struct {
	// `CfnFunction.S3NotificationFilterProperty.S3Key`.
	S3Key interface{} `field:"required" json:"s3Key" yaml:"s3Key"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sAMPolicyTemplateProperty := &sAMPolicyTemplateProperty{
//   	amiDescribePolicy: &emptySAMPTProperty{
//   	},
//   	awsSecretsManagerGetSecretValuePolicy: &secretArnSAMPTProperty{
//   		secretArn: jsii.String("secretArn"),
//   	},
//   	cloudFormationDescribeStacksPolicy: &emptySAMPTProperty{
//   	},
//   	cloudWatchPutMetricPolicy: &emptySAMPTProperty{
//   	},
//   	dynamoDbCrudPolicy: &tableSAMPTProperty{
//   		tableName: jsii.String("tableName"),
//   	},
//   	dynamoDbReadPolicy: &tableSAMPTProperty{
//   		tableName: jsii.String("tableName"),
//   	},
//   	dynamoDbStreamReadPolicy: &tableStreamSAMPTProperty{
//   		streamName: jsii.String("streamName"),
//   		tableName: jsii.String("tableName"),
//   	},
//   	dynamoDbWritePolicy: &tableSAMPTProperty{
//   		tableName: jsii.String("tableName"),
//   	},
//   	ec2DescribePolicy: &emptySAMPTProperty{
//   	},
//   	elasticsearchHttpPostPolicy: &domainSAMPTProperty{
//   		domainName: jsii.String("domainName"),
//   	},
//   	filterLogEventsPolicy: &logGroupSAMPTProperty{
//   		logGroupName: jsii.String("logGroupName"),
//   	},
//   	kinesisCrudPolicy: &streamSAMPTProperty{
//   		streamName: jsii.String("streamName"),
//   	},
//   	kinesisStreamReadPolicy: &streamSAMPTProperty{
//   		streamName: jsii.String("streamName"),
//   	},
//   	kmsDecryptPolicy: &keySAMPTProperty{
//   		keyId: jsii.String("keyId"),
//   	},
//   	lambdaInvokePolicy: &functionSAMPTProperty{
//   		functionName: jsii.String("functionName"),
//   	},
//   	rekognitionDetectOnlyPolicy: &emptySAMPTProperty{
//   	},
//   	rekognitionLabelsPolicy: &emptySAMPTProperty{
//   	},
//   	rekognitionNoDataAccessPolicy: &collectionSAMPTProperty{
//   		collectionId: jsii.String("collectionId"),
//   	},
//   	rekognitionReadPolicy: &collectionSAMPTProperty{
//   		collectionId: jsii.String("collectionId"),
//   	},
//   	rekognitionWriteOnlyAccessPolicy: &collectionSAMPTProperty{
//   		collectionId: jsii.String("collectionId"),
//   	},
//   	s3CrudPolicy: &bucketSAMPTProperty{
//   		bucketName: jsii.String("bucketName"),
//   	},
//   	s3ReadPolicy: &bucketSAMPTProperty{
//   		bucketName: jsii.String("bucketName"),
//   	},
//   	s3WritePolicy: &bucketSAMPTProperty{
//   		bucketName: jsii.String("bucketName"),
//   	},
//   	sesBulkTemplatedCrudPolicy: &identitySAMPTProperty{
//   		identityName: jsii.String("identityName"),
//   	},
//   	sesCrudPolicy: &identitySAMPTProperty{
//   		identityName: jsii.String("identityName"),
//   	},
//   	sesEmailTemplateCrudPolicy: &emptySAMPTProperty{
//   	},
//   	sesSendBouncePolicy: &identitySAMPTProperty{
//   		identityName: jsii.String("identityName"),
//   	},
//   	snsCrudPolicy: &topicSAMPTProperty{
//   		topicName: jsii.String("topicName"),
//   	},
//   	snsPublishMessagePolicy: &topicSAMPTProperty{
//   		topicName: jsii.String("topicName"),
//   	},
//   	sqsPollerPolicy: &queueSAMPTProperty{
//   		queueName: jsii.String("queueName"),
//   	},
//   	sqsSendMessagePolicy: &queueSAMPTProperty{
//   		queueName: jsii.String("queueName"),
//   	},
//   	ssmParameterReadPolicy: &parameterNameSAMPTProperty{
//   		parameterName: jsii.String("parameterName"),
//   	},
//   	stepFunctionsExecutionPolicy: &stateMachineSAMPTProperty{
//   		stateMachineName: jsii.String("stateMachineName"),
//   	},
//   	vpcAccessPolicy: &emptySAMPTProperty{
//   	},
//   }
//
type CfnFunction_SAMPolicyTemplateProperty struct {
	// `CfnFunction.SAMPolicyTemplateProperty.AMIDescribePolicy`.
	AmiDescribePolicy interface{} `field:"optional" json:"amiDescribePolicy" yaml:"amiDescribePolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.AWSSecretsManagerGetSecretValuePolicy`.
	AwsSecretsManagerGetSecretValuePolicy interface{} `field:"optional" json:"awsSecretsManagerGetSecretValuePolicy" yaml:"awsSecretsManagerGetSecretValuePolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.CloudFormationDescribeStacksPolicy`.
	CloudFormationDescribeStacksPolicy interface{} `field:"optional" json:"cloudFormationDescribeStacksPolicy" yaml:"cloudFormationDescribeStacksPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.CloudWatchPutMetricPolicy`.
	CloudWatchPutMetricPolicy interface{} `field:"optional" json:"cloudWatchPutMetricPolicy" yaml:"cloudWatchPutMetricPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.DynamoDBCrudPolicy`.
	DynamoDbCrudPolicy interface{} `field:"optional" json:"dynamoDbCrudPolicy" yaml:"dynamoDbCrudPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.DynamoDBReadPolicy`.
	DynamoDbReadPolicy interface{} `field:"optional" json:"dynamoDbReadPolicy" yaml:"dynamoDbReadPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.DynamoDBStreamReadPolicy`.
	DynamoDbStreamReadPolicy interface{} `field:"optional" json:"dynamoDbStreamReadPolicy" yaml:"dynamoDbStreamReadPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.DynamoDBWritePolicy`.
	DynamoDbWritePolicy interface{} `field:"optional" json:"dynamoDbWritePolicy" yaml:"dynamoDbWritePolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.EC2DescribePolicy`.
	Ec2DescribePolicy interface{} `field:"optional" json:"ec2DescribePolicy" yaml:"ec2DescribePolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.ElasticsearchHttpPostPolicy`.
	ElasticsearchHttpPostPolicy interface{} `field:"optional" json:"elasticsearchHttpPostPolicy" yaml:"elasticsearchHttpPostPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.FilterLogEventsPolicy`.
	FilterLogEventsPolicy interface{} `field:"optional" json:"filterLogEventsPolicy" yaml:"filterLogEventsPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.KinesisCrudPolicy`.
	KinesisCrudPolicy interface{} `field:"optional" json:"kinesisCrudPolicy" yaml:"kinesisCrudPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.KinesisStreamReadPolicy`.
	KinesisStreamReadPolicy interface{} `field:"optional" json:"kinesisStreamReadPolicy" yaml:"kinesisStreamReadPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.KMSDecryptPolicy`.
	KmsDecryptPolicy interface{} `field:"optional" json:"kmsDecryptPolicy" yaml:"kmsDecryptPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.LambdaInvokePolicy`.
	LambdaInvokePolicy interface{} `field:"optional" json:"lambdaInvokePolicy" yaml:"lambdaInvokePolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.RekognitionDetectOnlyPolicy`.
	RekognitionDetectOnlyPolicy interface{} `field:"optional" json:"rekognitionDetectOnlyPolicy" yaml:"rekognitionDetectOnlyPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.RekognitionLabelsPolicy`.
	RekognitionLabelsPolicy interface{} `field:"optional" json:"rekognitionLabelsPolicy" yaml:"rekognitionLabelsPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.RekognitionNoDataAccessPolicy`.
	RekognitionNoDataAccessPolicy interface{} `field:"optional" json:"rekognitionNoDataAccessPolicy" yaml:"rekognitionNoDataAccessPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.RekognitionReadPolicy`.
	RekognitionReadPolicy interface{} `field:"optional" json:"rekognitionReadPolicy" yaml:"rekognitionReadPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.RekognitionWriteOnlyAccessPolicy`.
	RekognitionWriteOnlyAccessPolicy interface{} `field:"optional" json:"rekognitionWriteOnlyAccessPolicy" yaml:"rekognitionWriteOnlyAccessPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.S3CrudPolicy`.
	S3CrudPolicy interface{} `field:"optional" json:"s3CrudPolicy" yaml:"s3CrudPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.S3ReadPolicy`.
	S3ReadPolicy interface{} `field:"optional" json:"s3ReadPolicy" yaml:"s3ReadPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.S3WritePolicy`.
	S3WritePolicy interface{} `field:"optional" json:"s3WritePolicy" yaml:"s3WritePolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.SESBulkTemplatedCrudPolicy`.
	SesBulkTemplatedCrudPolicy interface{} `field:"optional" json:"sesBulkTemplatedCrudPolicy" yaml:"sesBulkTemplatedCrudPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.SESCrudPolicy`.
	SesCrudPolicy interface{} `field:"optional" json:"sesCrudPolicy" yaml:"sesCrudPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.SESEmailTemplateCrudPolicy`.
	SesEmailTemplateCrudPolicy interface{} `field:"optional" json:"sesEmailTemplateCrudPolicy" yaml:"sesEmailTemplateCrudPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.SESSendBouncePolicy`.
	SesSendBouncePolicy interface{} `field:"optional" json:"sesSendBouncePolicy" yaml:"sesSendBouncePolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.SNSCrudPolicy`.
	SnsCrudPolicy interface{} `field:"optional" json:"snsCrudPolicy" yaml:"snsCrudPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.SNSPublishMessagePolicy`.
	SnsPublishMessagePolicy interface{} `field:"optional" json:"snsPublishMessagePolicy" yaml:"snsPublishMessagePolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.SQSPollerPolicy`.
	SqsPollerPolicy interface{} `field:"optional" json:"sqsPollerPolicy" yaml:"sqsPollerPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.SQSSendMessagePolicy`.
	SqsSendMessagePolicy interface{} `field:"optional" json:"sqsSendMessagePolicy" yaml:"sqsSendMessagePolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.SSMParameterReadPolicy`.
	SsmParameterReadPolicy interface{} `field:"optional" json:"ssmParameterReadPolicy" yaml:"ssmParameterReadPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.StepFunctionsExecutionPolicy`.
	StepFunctionsExecutionPolicy interface{} `field:"optional" json:"stepFunctionsExecutionPolicy" yaml:"stepFunctionsExecutionPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.VPCAccessPolicy`.
	VpcAccessPolicy interface{} `field:"optional" json:"vpcAccessPolicy" yaml:"vpcAccessPolicy"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sNSEventProperty := &sNSEventProperty{
//   	topic: jsii.String("topic"),
//   }
//
type CfnFunction_SNSEventProperty struct {
	// `CfnFunction.SNSEventProperty.Topic`.
	Topic *string `field:"required" json:"topic" yaml:"topic"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sQSEventProperty := &sQSEventProperty{
//   	queue: jsii.String("queue"),
//
//   	// the properties below are optional
//   	batchSize: jsii.Number(123),
//   	enabled: jsii.Boolean(false),
//   }
//
type CfnFunction_SQSEventProperty struct {
	// `CfnFunction.SQSEventProperty.Queue`.
	Queue *string `field:"required" json:"queue" yaml:"queue"`
	// `CfnFunction.SQSEventProperty.BatchSize`.
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// `CfnFunction.SQSEventProperty.Enabled`.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scheduleEventProperty := &scheduleEventProperty{
//   	schedule: jsii.String("schedule"),
//
//   	// the properties below are optional
//   	input: jsii.String("input"),
//   }
//
type CfnFunction_ScheduleEventProperty struct {
	// `CfnFunction.ScheduleEventProperty.Schedule`.
	Schedule *string `field:"required" json:"schedule" yaml:"schedule"`
	// `CfnFunction.ScheduleEventProperty.Input`.
	Input *string `field:"optional" json:"input" yaml:"input"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   secretArnSAMPTProperty := &secretArnSAMPTProperty{
//   	secretArn: jsii.String("secretArn"),
//   }
//
type CfnFunction_SecretArnSAMPTProperty struct {
	// `CfnFunction.SecretArnSAMPTProperty.SecretArn`.
	SecretArn *string `field:"required" json:"secretArn" yaml:"secretArn"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stateMachineSAMPTProperty := &stateMachineSAMPTProperty{
//   	stateMachineName: jsii.String("stateMachineName"),
//   }
//
type CfnFunction_StateMachineSAMPTProperty struct {
	// `CfnFunction.StateMachineSAMPTProperty.StateMachineName`.
	StateMachineName *string `field:"required" json:"stateMachineName" yaml:"stateMachineName"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   streamSAMPTProperty := &streamSAMPTProperty{
//   	streamName: jsii.String("streamName"),
//   }
//
type CfnFunction_StreamSAMPTProperty struct {
	// `CfnFunction.StreamSAMPTProperty.StreamName`.
	StreamName *string `field:"required" json:"streamName" yaml:"streamName"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tableSAMPTProperty := &tableSAMPTProperty{
//   	tableName: jsii.String("tableName"),
//   }
//
type CfnFunction_TableSAMPTProperty struct {
	// `CfnFunction.TableSAMPTProperty.TableName`.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tableStreamSAMPTProperty := &tableStreamSAMPTProperty{
//   	streamName: jsii.String("streamName"),
//   	tableName: jsii.String("tableName"),
//   }
//
type CfnFunction_TableStreamSAMPTProperty struct {
	// `CfnFunction.TableStreamSAMPTProperty.StreamName`.
	StreamName *string `field:"required" json:"streamName" yaml:"streamName"`
	// `CfnFunction.TableStreamSAMPTProperty.TableName`.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   topicSAMPTProperty := &topicSAMPTProperty{
//   	topicName: jsii.String("topicName"),
//   }
//
type CfnFunction_TopicSAMPTProperty struct {
	// `CfnFunction.TopicSAMPTProperty.TopicName`.
	TopicName *string `field:"required" json:"topicName" yaml:"topicName"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcConfigProperty := &vpcConfigProperty{
//   	securityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	subnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   }
//
type CfnFunction_VpcConfigProperty struct {
	// `CfnFunction.VpcConfigProperty.SecurityGroupIds`.
	SecurityGroupIds *[]*string `field:"required" json:"securityGroupIds" yaml:"securityGroupIds"`
	// `CfnFunction.VpcConfigProperty.SubnetIds`.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
}

// Properties for defining a `CfnFunction`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var assumeRolePolicyDocument interface{}
//
//   cfnFunctionProps := &cfnFunctionProps{
//   	architectures: []*string{
//   		jsii.String("architectures"),
//   	},
//   	assumeRolePolicyDocument: assumeRolePolicyDocument,
//   	autoPublishAlias: jsii.String("autoPublishAlias"),
//   	autoPublishCodeSha256: jsii.String("autoPublishCodeSha256"),
//   	codeSigningConfigArn: jsii.String("codeSigningConfigArn"),
//   	codeUri: jsii.String("codeUri"),
//   	deadLetterQueue: &deadLetterQueueProperty{
//   		targetArn: jsii.String("targetArn"),
//   		type: jsii.String("type"),
//   	},
//   	deploymentPreference: &deploymentPreferenceProperty{
//   		enabled: jsii.Boolean(false),
//   		type: jsii.String("type"),
//
//   		// the properties below are optional
//   		alarms: []*string{
//   			jsii.String("alarms"),
//   		},
//   		hooks: &hooksProperty{
//   			postTraffic: jsii.String("postTraffic"),
//   			preTraffic: jsii.String("preTraffic"),
//   		},
//   	},
//   	description: jsii.String("description"),
//   	environment: &functionEnvironmentProperty{
//   		variables: map[string]*string{
//   			"variablesKey": jsii.String("variables"),
//   		},
//   	},
//   	eventInvokeConfig: &eventInvokeConfigProperty{
//   		destinationConfig: &eventInvokeDestinationConfigProperty{
//   			onFailure: &destinationProperty{
//   				destination: jsii.String("destination"),
//
//   				// the properties below are optional
//   				type: jsii.String("type"),
//   			},
//   			onSuccess: &destinationProperty{
//   				destination: jsii.String("destination"),
//
//   				// the properties below are optional
//   				type: jsii.String("type"),
//   			},
//   		},
//   		maximumEventAgeInSeconds: jsii.Number(123),
//   		maximumRetryAttempts: jsii.Number(123),
//   	},
//   	events: map[string]interface{}{
//   		"eventsKey": &EventSourceProperty{
//   			"properties": &S3EventProperty{
//   				"variables": map[string]*string{
//   					"variablesKey": jsii.String("variables"),
//   				},
//   			},
//   			"type": jsii.String("type"),
//   		},
//   	},
//   	fileSystemConfigs: []interface{}{
//   		&fileSystemConfigProperty{
//   			arn: jsii.String("arn"),
//   			localMountPath: jsii.String("localMountPath"),
//   		},
//   	},
//   	functionName: jsii.String("functionName"),
//   	handler: jsii.String("handler"),
//   	imageConfig: &imageConfigProperty{
//   		command: []*string{
//   			jsii.String("command"),
//   		},
//   		entryPoint: []*string{
//   			jsii.String("entryPoint"),
//   		},
//   		workingDirectory: jsii.String("workingDirectory"),
//   	},
//   	imageUri: jsii.String("imageUri"),
//   	inlineCode: jsii.String("inlineCode"),
//   	kmsKeyArn: jsii.String("kmsKeyArn"),
//   	layers: []*string{
//   		jsii.String("layers"),
//   	},
//   	memorySize: jsii.Number(123),
//   	packageType: jsii.String("packageType"),
//   	permissionsBoundary: jsii.String("permissionsBoundary"),
//   	policies: jsii.String("policies"),
//   	provisionedConcurrencyConfig: &provisionedConcurrencyConfigProperty{
//   		provisionedConcurrentExecutions: jsii.String("provisionedConcurrentExecutions"),
//   	},
//   	reservedConcurrentExecutions: jsii.Number(123),
//   	role: jsii.String("role"),
//   	runtime: jsii.String("runtime"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	timeout: jsii.Number(123),
//   	tracing: jsii.String("tracing"),
//   	versionDescription: jsii.String("versionDescription"),
//   	vpcConfig: &vpcConfigProperty{
//   		securityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		subnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   }
//
type CfnFunctionProps struct {
	// `AWS::Serverless::Function.Architectures`.
	Architectures *[]*string `field:"optional" json:"architectures" yaml:"architectures"`
	// `AWS::Serverless::Function.AssumeRolePolicyDocument`.
	AssumeRolePolicyDocument interface{} `field:"optional" json:"assumeRolePolicyDocument" yaml:"assumeRolePolicyDocument"`
	// `AWS::Serverless::Function.AutoPublishAlias`.
	AutoPublishAlias *string `field:"optional" json:"autoPublishAlias" yaml:"autoPublishAlias"`
	// `AWS::Serverless::Function.AutoPublishCodeSha256`.
	AutoPublishCodeSha256 *string `field:"optional" json:"autoPublishCodeSha256" yaml:"autoPublishCodeSha256"`
	// `AWS::Serverless::Function.CodeSigningConfigArn`.
	CodeSigningConfigArn *string `field:"optional" json:"codeSigningConfigArn" yaml:"codeSigningConfigArn"`
	// `AWS::Serverless::Function.CodeUri`.
	CodeUri interface{} `field:"optional" json:"codeUri" yaml:"codeUri"`
	// `AWS::Serverless::Function.DeadLetterQueue`.
	DeadLetterQueue interface{} `field:"optional" json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// `AWS::Serverless::Function.DeploymentPreference`.
	DeploymentPreference interface{} `field:"optional" json:"deploymentPreference" yaml:"deploymentPreference"`
	// `AWS::Serverless::Function.Description`.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `AWS::Serverless::Function.Environment`.
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// `AWS::Serverless::Function.EventInvokeConfig`.
	EventInvokeConfig interface{} `field:"optional" json:"eventInvokeConfig" yaml:"eventInvokeConfig"`
	// `AWS::Serverless::Function.Events`.
	Events interface{} `field:"optional" json:"events" yaml:"events"`
	// `AWS::Serverless::Function.FileSystemConfigs`.
	FileSystemConfigs interface{} `field:"optional" json:"fileSystemConfigs" yaml:"fileSystemConfigs"`
	// `AWS::Serverless::Function.FunctionName`.
	FunctionName *string `field:"optional" json:"functionName" yaml:"functionName"`
	// `AWS::Serverless::Function.Handler`.
	Handler *string `field:"optional" json:"handler" yaml:"handler"`
	// `AWS::Serverless::Function.ImageConfig`.
	ImageConfig interface{} `field:"optional" json:"imageConfig" yaml:"imageConfig"`
	// `AWS::Serverless::Function.ImageUri`.
	ImageUri *string `field:"optional" json:"imageUri" yaml:"imageUri"`
	// `AWS::Serverless::Function.InlineCode`.
	InlineCode *string `field:"optional" json:"inlineCode" yaml:"inlineCode"`
	// `AWS::Serverless::Function.KmsKeyArn`.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// `AWS::Serverless::Function.Layers`.
	Layers *[]*string `field:"optional" json:"layers" yaml:"layers"`
	// `AWS::Serverless::Function.MemorySize`.
	MemorySize *float64 `field:"optional" json:"memorySize" yaml:"memorySize"`
	// `AWS::Serverless::Function.PackageType`.
	PackageType *string `field:"optional" json:"packageType" yaml:"packageType"`
	// `AWS::Serverless::Function.PermissionsBoundary`.
	PermissionsBoundary *string `field:"optional" json:"permissionsBoundary" yaml:"permissionsBoundary"`
	// `AWS::Serverless::Function.Policies`.
	Policies interface{} `field:"optional" json:"policies" yaml:"policies"`
	// `AWS::Serverless::Function.ProvisionedConcurrencyConfig`.
	ProvisionedConcurrencyConfig interface{} `field:"optional" json:"provisionedConcurrencyConfig" yaml:"provisionedConcurrencyConfig"`
	// `AWS::Serverless::Function.ReservedConcurrentExecutions`.
	ReservedConcurrentExecutions *float64 `field:"optional" json:"reservedConcurrentExecutions" yaml:"reservedConcurrentExecutions"`
	// `AWS::Serverless::Function.Role`.
	Role *string `field:"optional" json:"role" yaml:"role"`
	// `AWS::Serverless::Function.Runtime`.
	Runtime *string `field:"optional" json:"runtime" yaml:"runtime"`
	// `AWS::Serverless::Function.Tags`.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// `AWS::Serverless::Function.Timeout`.
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
	// `AWS::Serverless::Function.Tracing`.
	Tracing *string `field:"optional" json:"tracing" yaml:"tracing"`
	// `AWS::Serverless::Function.VersionDescription`.
	VersionDescription *string `field:"optional" json:"versionDescription" yaml:"versionDescription"`
	// `AWS::Serverless::Function.VpcConfig`.
	VpcConfig interface{} `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

// A CloudFormation `AWS::Serverless::HttpApi`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var authorizers interface{}
//   var definitionBody interface{}
//
//   cfnHttpApi := awscdk.Aws_sam.NewCfnHttpApi(this, jsii.String("MyCfnHttpApi"), &cfnHttpApiProps{
//   	accessLogSetting: &accessLogSettingProperty{
//   		destinationArn: jsii.String("destinationArn"),
//   		format: jsii.String("format"),
//   	},
//   	auth: &httpApiAuthProperty{
//   		authorizers: authorizers,
//   		defaultAuthorizer: jsii.String("defaultAuthorizer"),
//   	},
//   	corsConfiguration: jsii.Boolean(false),
//   	defaultRouteSettings: &routeSettingsProperty{
//   		dataTraceEnabled: jsii.Boolean(false),
//   		detailedMetricsEnabled: jsii.Boolean(false),
//   		loggingLevel: jsii.String("loggingLevel"),
//   		throttlingBurstLimit: jsii.Number(123),
//   		throttlingRateLimit: jsii.Number(123),
//   	},
//   	definitionBody: definitionBody,
//   	definitionUri: jsii.String("definitionUri"),
//   	description: jsii.String("description"),
//   	disableExecuteApiEndpoint: jsii.Boolean(false),
//   	domain: &httpApiDomainConfigurationProperty{
//   		certificateArn: jsii.String("certificateArn"),
//   		domainName: jsii.String("domainName"),
//
//   		// the properties below are optional
//   		basePath: jsii.String("basePath"),
//   		endpointConfiguration: jsii.String("endpointConfiguration"),
//   		mutualTlsAuthentication: &mutualTlsAuthenticationProperty{
//   			truststoreUri: jsii.String("truststoreUri"),
//   			truststoreVersion: jsii.Boolean(false),
//   		},
//   		route53: &route53ConfigurationProperty{
//   			distributedDomainName: jsii.String("distributedDomainName"),
//   			evaluateTargetHealth: jsii.Boolean(false),
//   			hostedZoneId: jsii.String("hostedZoneId"),
//   			hostedZoneName: jsii.String("hostedZoneName"),
//   			ipV6: jsii.Boolean(false),
//   		},
//   		securityPolicy: jsii.String("securityPolicy"),
//   	},
//   	failOnWarnings: jsii.Boolean(false),
//   	routeSettings: &routeSettingsProperty{
//   		dataTraceEnabled: jsii.Boolean(false),
//   		detailedMetricsEnabled: jsii.Boolean(false),
//   		loggingLevel: jsii.String("loggingLevel"),
//   		throttlingBurstLimit: jsii.Number(123),
//   		throttlingRateLimit: jsii.Number(123),
//   	},
//   	stageName: jsii.String("stageName"),
//   	stageVariables: map[string]*string{
//   		"stageVariablesKey": jsii.String("stageVariables"),
//   	},
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   })
//
type CfnHttpApi interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// `AWS::Serverless::HttpApi.AccessLogSetting`.
	AccessLogSetting() interface{}
	SetAccessLogSetting(val interface{})
	// `AWS::Serverless::HttpApi.Auth`.
	Auth() interface{}
	SetAuth(val interface{})
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// `AWS::Serverless::HttpApi.CorsConfiguration`.
	CorsConfiguration() interface{}
	SetCorsConfiguration(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// `AWS::Serverless::HttpApi.DefaultRouteSettings`.
	DefaultRouteSettings() interface{}
	SetDefaultRouteSettings(val interface{})
	// `AWS::Serverless::HttpApi.DefinitionBody`.
	DefinitionBody() interface{}
	SetDefinitionBody(val interface{})
	// `AWS::Serverless::HttpApi.DefinitionUri`.
	DefinitionUri() interface{}
	SetDefinitionUri(val interface{})
	// `AWS::Serverless::HttpApi.Description`.
	Description() *string
	SetDescription(val *string)
	// `AWS::Serverless::HttpApi.DisableExecuteApiEndpoint`.
	DisableExecuteApiEndpoint() interface{}
	SetDisableExecuteApiEndpoint(val interface{})
	// `AWS::Serverless::HttpApi.Domain`.
	Domain() interface{}
	SetDomain(val interface{})
	// `AWS::Serverless::HttpApi.FailOnWarnings`.
	FailOnWarnings() interface{}
	SetFailOnWarnings(val interface{})
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
	// `AWS::Serverless::HttpApi.RouteSettings`.
	RouteSettings() interface{}
	SetRouteSettings(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// `AWS::Serverless::HttpApi.StageName`.
	StageName() *string
	SetStageName(val *string)
	// `AWS::Serverless::HttpApi.StageVariables`.
	StageVariables() interface{}
	SetStageVariables(val interface{})
	// `AWS::Serverless::HttpApi.Tags`.
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

// The jsii proxy struct for CfnHttpApi
type jsiiProxy_CfnHttpApi struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnHttpApi) AccessLogSetting() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessLogSetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHttpApi) Auth() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"auth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHttpApi) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHttpApi) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHttpApi) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHttpApi) CorsConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"corsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHttpApi) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHttpApi) DefaultRouteSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultRouteSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHttpApi) DefinitionBody() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"definitionBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHttpApi) DefinitionUri() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"definitionUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHttpApi) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHttpApi) DisableExecuteApiEndpoint() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableExecuteApiEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHttpApi) Domain() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHttpApi) FailOnWarnings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failOnWarnings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHttpApi) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHttpApi) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHttpApi) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHttpApi) RouteSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"routeSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHttpApi) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHttpApi) StageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHttpApi) StageVariables() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stageVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHttpApi) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHttpApi) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Serverless::HttpApi`.
func NewCfnHttpApi(scope awscdk.Construct, id *string, props *CfnHttpApiProps) CfnHttpApi {
	_init_.Initialize()

	j := jsiiProxy_CfnHttpApi{}

	_jsii_.Create(
		"monocdk.aws_sam.CfnHttpApi",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Serverless::HttpApi`.
func NewCfnHttpApi_Override(c CfnHttpApi, scope awscdk.Construct, id *string, props *CfnHttpApiProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_sam.CfnHttpApi",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnHttpApi) SetAccessLogSetting(val interface{}) {
	_jsii_.Set(
		j,
		"accessLogSetting",
		val,
	)
}

func (j *jsiiProxy_CfnHttpApi) SetAuth(val interface{}) {
	_jsii_.Set(
		j,
		"auth",
		val,
	)
}

func (j *jsiiProxy_CfnHttpApi) SetCorsConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"corsConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnHttpApi) SetDefaultRouteSettings(val interface{}) {
	_jsii_.Set(
		j,
		"defaultRouteSettings",
		val,
	)
}

func (j *jsiiProxy_CfnHttpApi) SetDefinitionBody(val interface{}) {
	_jsii_.Set(
		j,
		"definitionBody",
		val,
	)
}

func (j *jsiiProxy_CfnHttpApi) SetDefinitionUri(val interface{}) {
	_jsii_.Set(
		j,
		"definitionUri",
		val,
	)
}

func (j *jsiiProxy_CfnHttpApi) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnHttpApi) SetDisableExecuteApiEndpoint(val interface{}) {
	_jsii_.Set(
		j,
		"disableExecuteApiEndpoint",
		val,
	)
}

func (j *jsiiProxy_CfnHttpApi) SetDomain(val interface{}) {
	_jsii_.Set(
		j,
		"domain",
		val,
	)
}

func (j *jsiiProxy_CfnHttpApi) SetFailOnWarnings(val interface{}) {
	_jsii_.Set(
		j,
		"failOnWarnings",
		val,
	)
}

func (j *jsiiProxy_CfnHttpApi) SetRouteSettings(val interface{}) {
	_jsii_.Set(
		j,
		"routeSettings",
		val,
	)
}

func (j *jsiiProxy_CfnHttpApi) SetStageName(val *string) {
	_jsii_.Set(
		j,
		"stageName",
		val,
	)
}

func (j *jsiiProxy_CfnHttpApi) SetStageVariables(val interface{}) {
	_jsii_.Set(
		j,
		"stageVariables",
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
func CfnHttpApi_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sam.CfnHttpApi",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnHttpApi_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sam.CfnHttpApi",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnHttpApi_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sam.CfnHttpApi",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnHttpApi_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_sam.CfnHttpApi",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func CfnHttpApi_REQUIRED_TRANSFORM() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_sam.CfnHttpApi",
		"REQUIRED_TRANSFORM",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnHttpApi) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnHttpApi) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnHttpApi) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnHttpApi) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnHttpApi) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnHttpApi) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnHttpApi) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnHttpApi) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnHttpApi) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnHttpApi) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnHttpApi) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnHttpApi) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnHttpApi) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnHttpApi) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnHttpApi) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnHttpApi) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnHttpApi) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnHttpApi) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnHttpApi) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnHttpApi) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnHttpApi) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessLogSettingProperty := &accessLogSettingProperty{
//   	destinationArn: jsii.String("destinationArn"),
//   	format: jsii.String("format"),
//   }
//
type CfnHttpApi_AccessLogSettingProperty struct {
	// `CfnHttpApi.AccessLogSettingProperty.DestinationArn`.
	DestinationArn *string `field:"optional" json:"destinationArn" yaml:"destinationArn"`
	// `CfnHttpApi.AccessLogSettingProperty.Format`.
	Format *string `field:"optional" json:"format" yaml:"format"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   corsConfigurationObjectProperty := &corsConfigurationObjectProperty{
//   	allowCredentials: jsii.Boolean(false),
//   	allowHeaders: jsii.String("allowHeaders"),
//   	allowMethods: jsii.String("allowMethods"),
//   	allowOrigin: jsii.String("allowOrigin"),
//   	exposeHeaders: []*string{
//   		jsii.String("exposeHeaders"),
//   	},
//   	maxAge: jsii.String("maxAge"),
//   }
//
type CfnHttpApi_CorsConfigurationObjectProperty struct {
	// `CfnHttpApi.CorsConfigurationObjectProperty.AllowCredentials`.
	AllowCredentials interface{} `field:"optional" json:"allowCredentials" yaml:"allowCredentials"`
	// `CfnHttpApi.CorsConfigurationObjectProperty.AllowHeaders`.
	AllowHeaders *string `field:"optional" json:"allowHeaders" yaml:"allowHeaders"`
	// `CfnHttpApi.CorsConfigurationObjectProperty.AllowMethods`.
	AllowMethods *string `field:"optional" json:"allowMethods" yaml:"allowMethods"`
	// `CfnHttpApi.CorsConfigurationObjectProperty.AllowOrigin`.
	AllowOrigin *string `field:"optional" json:"allowOrigin" yaml:"allowOrigin"`
	// `CfnHttpApi.CorsConfigurationObjectProperty.ExposeHeaders`.
	ExposeHeaders *[]*string `field:"optional" json:"exposeHeaders" yaml:"exposeHeaders"`
	// `CfnHttpApi.CorsConfigurationObjectProperty.MaxAge`.
	MaxAge *string `field:"optional" json:"maxAge" yaml:"maxAge"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var authorizers interface{}
//
//   httpApiAuthProperty := &httpApiAuthProperty{
//   	authorizers: authorizers,
//   	defaultAuthorizer: jsii.String("defaultAuthorizer"),
//   }
//
type CfnHttpApi_HttpApiAuthProperty struct {
	// `CfnHttpApi.HttpApiAuthProperty.Authorizers`.
	Authorizers interface{} `field:"optional" json:"authorizers" yaml:"authorizers"`
	// `CfnHttpApi.HttpApiAuthProperty.DefaultAuthorizer`.
	DefaultAuthorizer *string `field:"optional" json:"defaultAuthorizer" yaml:"defaultAuthorizer"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpApiDomainConfigurationProperty := &httpApiDomainConfigurationProperty{
//   	certificateArn: jsii.String("certificateArn"),
//   	domainName: jsii.String("domainName"),
//
//   	// the properties below are optional
//   	basePath: jsii.String("basePath"),
//   	endpointConfiguration: jsii.String("endpointConfiguration"),
//   	mutualTlsAuthentication: &mutualTlsAuthenticationProperty{
//   		truststoreUri: jsii.String("truststoreUri"),
//   		truststoreVersion: jsii.Boolean(false),
//   	},
//   	route53: &route53ConfigurationProperty{
//   		distributedDomainName: jsii.String("distributedDomainName"),
//   		evaluateTargetHealth: jsii.Boolean(false),
//   		hostedZoneId: jsii.String("hostedZoneId"),
//   		hostedZoneName: jsii.String("hostedZoneName"),
//   		ipV6: jsii.Boolean(false),
//   	},
//   	securityPolicy: jsii.String("securityPolicy"),
//   }
//
type CfnHttpApi_HttpApiDomainConfigurationProperty struct {
	// `CfnHttpApi.HttpApiDomainConfigurationProperty.CertificateArn`.
	CertificateArn *string `field:"required" json:"certificateArn" yaml:"certificateArn"`
	// `CfnHttpApi.HttpApiDomainConfigurationProperty.DomainName`.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// `CfnHttpApi.HttpApiDomainConfigurationProperty.BasePath`.
	BasePath *string `field:"optional" json:"basePath" yaml:"basePath"`
	// `CfnHttpApi.HttpApiDomainConfigurationProperty.EndpointConfiguration`.
	EndpointConfiguration *string `field:"optional" json:"endpointConfiguration" yaml:"endpointConfiguration"`
	// `CfnHttpApi.HttpApiDomainConfigurationProperty.MutualTlsAuthentication`.
	MutualTlsAuthentication interface{} `field:"optional" json:"mutualTlsAuthentication" yaml:"mutualTlsAuthentication"`
	// `CfnHttpApi.HttpApiDomainConfigurationProperty.Route53`.
	Route53 interface{} `field:"optional" json:"route53" yaml:"route53"`
	// `CfnHttpApi.HttpApiDomainConfigurationProperty.SecurityPolicy`.
	SecurityPolicy *string `field:"optional" json:"securityPolicy" yaml:"securityPolicy"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mutualTlsAuthenticationProperty := &mutualTlsAuthenticationProperty{
//   	truststoreUri: jsii.String("truststoreUri"),
//   	truststoreVersion: jsii.Boolean(false),
//   }
//
type CfnHttpApi_MutualTlsAuthenticationProperty struct {
	// `CfnHttpApi.MutualTlsAuthenticationProperty.TruststoreUri`.
	TruststoreUri *string `field:"optional" json:"truststoreUri" yaml:"truststoreUri"`
	// `CfnHttpApi.MutualTlsAuthenticationProperty.TruststoreVersion`.
	TruststoreVersion interface{} `field:"optional" json:"truststoreVersion" yaml:"truststoreVersion"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   route53ConfigurationProperty := &route53ConfigurationProperty{
//   	distributedDomainName: jsii.String("distributedDomainName"),
//   	evaluateTargetHealth: jsii.Boolean(false),
//   	hostedZoneId: jsii.String("hostedZoneId"),
//   	hostedZoneName: jsii.String("hostedZoneName"),
//   	ipV6: jsii.Boolean(false),
//   }
//
type CfnHttpApi_Route53ConfigurationProperty struct {
	// `CfnHttpApi.Route53ConfigurationProperty.DistributedDomainName`.
	DistributedDomainName *string `field:"optional" json:"distributedDomainName" yaml:"distributedDomainName"`
	// `CfnHttpApi.Route53ConfigurationProperty.EvaluateTargetHealth`.
	EvaluateTargetHealth interface{} `field:"optional" json:"evaluateTargetHealth" yaml:"evaluateTargetHealth"`
	// `CfnHttpApi.Route53ConfigurationProperty.HostedZoneId`.
	HostedZoneId *string `field:"optional" json:"hostedZoneId" yaml:"hostedZoneId"`
	// `CfnHttpApi.Route53ConfigurationProperty.HostedZoneName`.
	HostedZoneName *string `field:"optional" json:"hostedZoneName" yaml:"hostedZoneName"`
	// `CfnHttpApi.Route53ConfigurationProperty.IpV6`.
	IpV6 interface{} `field:"optional" json:"ipV6" yaml:"ipV6"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   routeSettingsProperty := &routeSettingsProperty{
//   	dataTraceEnabled: jsii.Boolean(false),
//   	detailedMetricsEnabled: jsii.Boolean(false),
//   	loggingLevel: jsii.String("loggingLevel"),
//   	throttlingBurstLimit: jsii.Number(123),
//   	throttlingRateLimit: jsii.Number(123),
//   }
//
type CfnHttpApi_RouteSettingsProperty struct {
	// `CfnHttpApi.RouteSettingsProperty.DataTraceEnabled`.
	DataTraceEnabled interface{} `field:"optional" json:"dataTraceEnabled" yaml:"dataTraceEnabled"`
	// `CfnHttpApi.RouteSettingsProperty.DetailedMetricsEnabled`.
	DetailedMetricsEnabled interface{} `field:"optional" json:"detailedMetricsEnabled" yaml:"detailedMetricsEnabled"`
	// `CfnHttpApi.RouteSettingsProperty.LoggingLevel`.
	LoggingLevel *string `field:"optional" json:"loggingLevel" yaml:"loggingLevel"`
	// `CfnHttpApi.RouteSettingsProperty.ThrottlingBurstLimit`.
	ThrottlingBurstLimit *float64 `field:"optional" json:"throttlingBurstLimit" yaml:"throttlingBurstLimit"`
	// `CfnHttpApi.RouteSettingsProperty.ThrottlingRateLimit`.
	ThrottlingRateLimit *float64 `field:"optional" json:"throttlingRateLimit" yaml:"throttlingRateLimit"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3LocationProperty := &s3LocationProperty{
//   	bucket: jsii.String("bucket"),
//   	key: jsii.String("key"),
//   	version: jsii.Number(123),
//   }
//
type CfnHttpApi_S3LocationProperty struct {
	// `CfnHttpApi.S3LocationProperty.Bucket`.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// `CfnHttpApi.S3LocationProperty.Key`.
	Key *string `field:"required" json:"key" yaml:"key"`
	// `CfnHttpApi.S3LocationProperty.Version`.
	Version *float64 `field:"required" json:"version" yaml:"version"`
}

// Properties for defining a `CfnHttpApi`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var authorizers interface{}
//   var definitionBody interface{}
//
//   cfnHttpApiProps := &cfnHttpApiProps{
//   	accessLogSetting: &accessLogSettingProperty{
//   		destinationArn: jsii.String("destinationArn"),
//   		format: jsii.String("format"),
//   	},
//   	auth: &httpApiAuthProperty{
//   		authorizers: authorizers,
//   		defaultAuthorizer: jsii.String("defaultAuthorizer"),
//   	},
//   	corsConfiguration: jsii.Boolean(false),
//   	defaultRouteSettings: &routeSettingsProperty{
//   		dataTraceEnabled: jsii.Boolean(false),
//   		detailedMetricsEnabled: jsii.Boolean(false),
//   		loggingLevel: jsii.String("loggingLevel"),
//   		throttlingBurstLimit: jsii.Number(123),
//   		throttlingRateLimit: jsii.Number(123),
//   	},
//   	definitionBody: definitionBody,
//   	definitionUri: jsii.String("definitionUri"),
//   	description: jsii.String("description"),
//   	disableExecuteApiEndpoint: jsii.Boolean(false),
//   	domain: &httpApiDomainConfigurationProperty{
//   		certificateArn: jsii.String("certificateArn"),
//   		domainName: jsii.String("domainName"),
//
//   		// the properties below are optional
//   		basePath: jsii.String("basePath"),
//   		endpointConfiguration: jsii.String("endpointConfiguration"),
//   		mutualTlsAuthentication: &mutualTlsAuthenticationProperty{
//   			truststoreUri: jsii.String("truststoreUri"),
//   			truststoreVersion: jsii.Boolean(false),
//   		},
//   		route53: &route53ConfigurationProperty{
//   			distributedDomainName: jsii.String("distributedDomainName"),
//   			evaluateTargetHealth: jsii.Boolean(false),
//   			hostedZoneId: jsii.String("hostedZoneId"),
//   			hostedZoneName: jsii.String("hostedZoneName"),
//   			ipV6: jsii.Boolean(false),
//   		},
//   		securityPolicy: jsii.String("securityPolicy"),
//   	},
//   	failOnWarnings: jsii.Boolean(false),
//   	routeSettings: &routeSettingsProperty{
//   		dataTraceEnabled: jsii.Boolean(false),
//   		detailedMetricsEnabled: jsii.Boolean(false),
//   		loggingLevel: jsii.String("loggingLevel"),
//   		throttlingBurstLimit: jsii.Number(123),
//   		throttlingRateLimit: jsii.Number(123),
//   	},
//   	stageName: jsii.String("stageName"),
//   	stageVariables: map[string]*string{
//   		"stageVariablesKey": jsii.String("stageVariables"),
//   	},
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnHttpApiProps struct {
	// `AWS::Serverless::HttpApi.AccessLogSetting`.
	AccessLogSetting interface{} `field:"optional" json:"accessLogSetting" yaml:"accessLogSetting"`
	// `AWS::Serverless::HttpApi.Auth`.
	Auth interface{} `field:"optional" json:"auth" yaml:"auth"`
	// `AWS::Serverless::HttpApi.CorsConfiguration`.
	CorsConfiguration interface{} `field:"optional" json:"corsConfiguration" yaml:"corsConfiguration"`
	// `AWS::Serverless::HttpApi.DefaultRouteSettings`.
	DefaultRouteSettings interface{} `field:"optional" json:"defaultRouteSettings" yaml:"defaultRouteSettings"`
	// `AWS::Serverless::HttpApi.DefinitionBody`.
	DefinitionBody interface{} `field:"optional" json:"definitionBody" yaml:"definitionBody"`
	// `AWS::Serverless::HttpApi.DefinitionUri`.
	DefinitionUri interface{} `field:"optional" json:"definitionUri" yaml:"definitionUri"`
	// `AWS::Serverless::HttpApi.Description`.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `AWS::Serverless::HttpApi.DisableExecuteApiEndpoint`.
	DisableExecuteApiEndpoint interface{} `field:"optional" json:"disableExecuteApiEndpoint" yaml:"disableExecuteApiEndpoint"`
	// `AWS::Serverless::HttpApi.Domain`.
	Domain interface{} `field:"optional" json:"domain" yaml:"domain"`
	// `AWS::Serverless::HttpApi.FailOnWarnings`.
	FailOnWarnings interface{} `field:"optional" json:"failOnWarnings" yaml:"failOnWarnings"`
	// `AWS::Serverless::HttpApi.RouteSettings`.
	RouteSettings interface{} `field:"optional" json:"routeSettings" yaml:"routeSettings"`
	// `AWS::Serverless::HttpApi.StageName`.
	StageName *string `field:"optional" json:"stageName" yaml:"stageName"`
	// `AWS::Serverless::HttpApi.StageVariables`.
	StageVariables interface{} `field:"optional" json:"stageVariables" yaml:"stageVariables"`
	// `AWS::Serverless::HttpApi.Tags`.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::Serverless::LayerVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLayerVersion := awscdk.Aws_sam.NewCfnLayerVersion(this, jsii.String("MyCfnLayerVersion"), &cfnLayerVersionProps{
//   	compatibleRuntimes: []*string{
//   		jsii.String("compatibleRuntimes"),
//   	},
//   	contentUri: jsii.String("contentUri"),
//   	description: jsii.String("description"),
//   	layerName: jsii.String("layerName"),
//   	licenseInfo: jsii.String("licenseInfo"),
//   	retentionPolicy: jsii.String("retentionPolicy"),
//   })
//
type CfnLayerVersion interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// `AWS::Serverless::LayerVersion.CompatibleRuntimes`.
	CompatibleRuntimes() *[]*string
	SetCompatibleRuntimes(val *[]*string)
	// `AWS::Serverless::LayerVersion.ContentUri`.
	ContentUri() interface{}
	SetContentUri(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// `AWS::Serverless::LayerVersion.Description`.
	Description() *string
	SetDescription(val *string)
	// `AWS::Serverless::LayerVersion.LayerName`.
	LayerName() *string
	SetLayerName(val *string)
	// `AWS::Serverless::LayerVersion.LicenseInfo`.
	LicenseInfo() *string
	SetLicenseInfo(val *string)
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
	// `AWS::Serverless::LayerVersion.RetentionPolicy`.
	RetentionPolicy() *string
	SetRetentionPolicy(val *string)
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

// The jsii proxy struct for CfnLayerVersion
type jsiiProxy_CfnLayerVersion struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnLayerVersion) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayerVersion) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayerVersion) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayerVersion) CompatibleRuntimes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"compatibleRuntimes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayerVersion) ContentUri() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contentUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayerVersion) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayerVersion) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayerVersion) LayerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"layerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayerVersion) LicenseInfo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayerVersion) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayerVersion) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayerVersion) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayerVersion) RetentionPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retentionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayerVersion) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayerVersion) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Serverless::LayerVersion`.
func NewCfnLayerVersion(scope awscdk.Construct, id *string, props *CfnLayerVersionProps) CfnLayerVersion {
	_init_.Initialize()

	j := jsiiProxy_CfnLayerVersion{}

	_jsii_.Create(
		"monocdk.aws_sam.CfnLayerVersion",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Serverless::LayerVersion`.
func NewCfnLayerVersion_Override(c CfnLayerVersion, scope awscdk.Construct, id *string, props *CfnLayerVersionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_sam.CfnLayerVersion",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnLayerVersion) SetCompatibleRuntimes(val *[]*string) {
	_jsii_.Set(
		j,
		"compatibleRuntimes",
		val,
	)
}

func (j *jsiiProxy_CfnLayerVersion) SetContentUri(val interface{}) {
	_jsii_.Set(
		j,
		"contentUri",
		val,
	)
}

func (j *jsiiProxy_CfnLayerVersion) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnLayerVersion) SetLayerName(val *string) {
	_jsii_.Set(
		j,
		"layerName",
		val,
	)
}

func (j *jsiiProxy_CfnLayerVersion) SetLicenseInfo(val *string) {
	_jsii_.Set(
		j,
		"licenseInfo",
		val,
	)
}

func (j *jsiiProxy_CfnLayerVersion) SetRetentionPolicy(val *string) {
	_jsii_.Set(
		j,
		"retentionPolicy",
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
func CfnLayerVersion_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sam.CfnLayerVersion",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnLayerVersion_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sam.CfnLayerVersion",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnLayerVersion_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sam.CfnLayerVersion",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLayerVersion_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_sam.CfnLayerVersion",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func CfnLayerVersion_REQUIRED_TRANSFORM() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_sam.CfnLayerVersion",
		"REQUIRED_TRANSFORM",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnLayerVersion) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnLayerVersion) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnLayerVersion) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnLayerVersion) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnLayerVersion) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnLayerVersion) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnLayerVersion) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnLayerVersion) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLayerVersion) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLayerVersion) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnLayerVersion) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnLayerVersion) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnLayerVersion) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLayerVersion) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnLayerVersion) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnLayerVersion) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLayerVersion) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLayerVersion) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnLayerVersion) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLayerVersion) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLayerVersion) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3LocationProperty := &s3LocationProperty{
//   	bucket: jsii.String("bucket"),
//   	key: jsii.String("key"),
//
//   	// the properties below are optional
//   	version: jsii.Number(123),
//   }
//
type CfnLayerVersion_S3LocationProperty struct {
	// `CfnLayerVersion.S3LocationProperty.Bucket`.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// `CfnLayerVersion.S3LocationProperty.Key`.
	Key *string `field:"required" json:"key" yaml:"key"`
	// `CfnLayerVersion.S3LocationProperty.Version`.
	Version *float64 `field:"optional" json:"version" yaml:"version"`
}

// Properties for defining a `CfnLayerVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLayerVersionProps := &cfnLayerVersionProps{
//   	compatibleRuntimes: []*string{
//   		jsii.String("compatibleRuntimes"),
//   	},
//   	contentUri: jsii.String("contentUri"),
//   	description: jsii.String("description"),
//   	layerName: jsii.String("layerName"),
//   	licenseInfo: jsii.String("licenseInfo"),
//   	retentionPolicy: jsii.String("retentionPolicy"),
//   }
//
type CfnLayerVersionProps struct {
	// `AWS::Serverless::LayerVersion.CompatibleRuntimes`.
	CompatibleRuntimes *[]*string `field:"optional" json:"compatibleRuntimes" yaml:"compatibleRuntimes"`
	// `AWS::Serverless::LayerVersion.ContentUri`.
	ContentUri interface{} `field:"optional" json:"contentUri" yaml:"contentUri"`
	// `AWS::Serverless::LayerVersion.Description`.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `AWS::Serverless::LayerVersion.LayerName`.
	LayerName *string `field:"optional" json:"layerName" yaml:"layerName"`
	// `AWS::Serverless::LayerVersion.LicenseInfo`.
	LicenseInfo *string `field:"optional" json:"licenseInfo" yaml:"licenseInfo"`
	// `AWS::Serverless::LayerVersion.RetentionPolicy`.
	RetentionPolicy *string `field:"optional" json:"retentionPolicy" yaml:"retentionPolicy"`
}

// A CloudFormation `AWS::Serverless::SimpleTable`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSimpleTable := awscdk.Aws_sam.NewCfnSimpleTable(this, jsii.String("MyCfnSimpleTable"), &cfnSimpleTableProps{
//   	primaryKey: &primaryKeyProperty{
//   		type: jsii.String("type"),
//
//   		// the properties below are optional
//   		name: jsii.String("name"),
//   	},
//   	provisionedThroughput: &provisionedThroughputProperty{
//   		writeCapacityUnits: jsii.Number(123),
//
//   		// the properties below are optional
//   		readCapacityUnits: jsii.Number(123),
//   	},
//   	sseSpecification: &sSESpecificationProperty{
//   		sseEnabled: jsii.Boolean(false),
//   	},
//   	tableName: jsii.String("tableName"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   })
//
type CfnSimpleTable interface {
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
	// `AWS::Serverless::SimpleTable.PrimaryKey`.
	PrimaryKey() interface{}
	SetPrimaryKey(val interface{})
	// `AWS::Serverless::SimpleTable.ProvisionedThroughput`.
	ProvisionedThroughput() interface{}
	SetProvisionedThroughput(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// `AWS::Serverless::SimpleTable.SSESpecification`.
	SseSpecification() interface{}
	SetSseSpecification(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// `AWS::Serverless::SimpleTable.TableName`.
	TableName() *string
	SetTableName(val *string)
	// `AWS::Serverless::SimpleTable.Tags`.
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

// The jsii proxy struct for CfnSimpleTable
type jsiiProxy_CfnSimpleTable struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnSimpleTable) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSimpleTable) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSimpleTable) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSimpleTable) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSimpleTable) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSimpleTable) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSimpleTable) PrimaryKey() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"primaryKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSimpleTable) ProvisionedThroughput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"provisionedThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSimpleTable) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSimpleTable) SseSpecification() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sseSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSimpleTable) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSimpleTable) TableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSimpleTable) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSimpleTable) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Serverless::SimpleTable`.
func NewCfnSimpleTable(scope awscdk.Construct, id *string, props *CfnSimpleTableProps) CfnSimpleTable {
	_init_.Initialize()

	j := jsiiProxy_CfnSimpleTable{}

	_jsii_.Create(
		"monocdk.aws_sam.CfnSimpleTable",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Serverless::SimpleTable`.
func NewCfnSimpleTable_Override(c CfnSimpleTable, scope awscdk.Construct, id *string, props *CfnSimpleTableProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_sam.CfnSimpleTable",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnSimpleTable) SetPrimaryKey(val interface{}) {
	_jsii_.Set(
		j,
		"primaryKey",
		val,
	)
}

func (j *jsiiProxy_CfnSimpleTable) SetProvisionedThroughput(val interface{}) {
	_jsii_.Set(
		j,
		"provisionedThroughput",
		val,
	)
}

func (j *jsiiProxy_CfnSimpleTable) SetSseSpecification(val interface{}) {
	_jsii_.Set(
		j,
		"sseSpecification",
		val,
	)
}

func (j *jsiiProxy_CfnSimpleTable) SetTableName(val *string) {
	_jsii_.Set(
		j,
		"tableName",
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
func CfnSimpleTable_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sam.CfnSimpleTable",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnSimpleTable_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sam.CfnSimpleTable",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnSimpleTable_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sam.CfnSimpleTable",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSimpleTable_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_sam.CfnSimpleTable",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func CfnSimpleTable_REQUIRED_TRANSFORM() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_sam.CfnSimpleTable",
		"REQUIRED_TRANSFORM",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSimpleTable) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnSimpleTable) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnSimpleTable) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnSimpleTable) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnSimpleTable) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnSimpleTable) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnSimpleTable) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnSimpleTable) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSimpleTable) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSimpleTable) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnSimpleTable) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnSimpleTable) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnSimpleTable) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSimpleTable) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnSimpleTable) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnSimpleTable) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSimpleTable) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSimpleTable) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnSimpleTable) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSimpleTable) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSimpleTable) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   primaryKeyProperty := &primaryKeyProperty{
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	name: jsii.String("name"),
//   }
//
type CfnSimpleTable_PrimaryKeyProperty struct {
	// `CfnSimpleTable.PrimaryKeyProperty.Type`.
	Type *string `field:"required" json:"type" yaml:"type"`
	// `CfnSimpleTable.PrimaryKeyProperty.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   provisionedThroughputProperty := &provisionedThroughputProperty{
//   	writeCapacityUnits: jsii.Number(123),
//
//   	// the properties below are optional
//   	readCapacityUnits: jsii.Number(123),
//   }
//
type CfnSimpleTable_ProvisionedThroughputProperty struct {
	// `CfnSimpleTable.ProvisionedThroughputProperty.WriteCapacityUnits`.
	WriteCapacityUnits *float64 `field:"required" json:"writeCapacityUnits" yaml:"writeCapacityUnits"`
	// `CfnSimpleTable.ProvisionedThroughputProperty.ReadCapacityUnits`.
	ReadCapacityUnits *float64 `field:"optional" json:"readCapacityUnits" yaml:"readCapacityUnits"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sSESpecificationProperty := &sSESpecificationProperty{
//   	sseEnabled: jsii.Boolean(false),
//   }
//
type CfnSimpleTable_SSESpecificationProperty struct {
	// `CfnSimpleTable.SSESpecificationProperty.SSEEnabled`.
	SseEnabled interface{} `field:"optional" json:"sseEnabled" yaml:"sseEnabled"`
}

// Properties for defining a `CfnSimpleTable`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSimpleTableProps := &cfnSimpleTableProps{
//   	primaryKey: &primaryKeyProperty{
//   		type: jsii.String("type"),
//
//   		// the properties below are optional
//   		name: jsii.String("name"),
//   	},
//   	provisionedThroughput: &provisionedThroughputProperty{
//   		writeCapacityUnits: jsii.Number(123),
//
//   		// the properties below are optional
//   		readCapacityUnits: jsii.Number(123),
//   	},
//   	sseSpecification: &sSESpecificationProperty{
//   		sseEnabled: jsii.Boolean(false),
//   	},
//   	tableName: jsii.String("tableName"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnSimpleTableProps struct {
	// `AWS::Serverless::SimpleTable.PrimaryKey`.
	PrimaryKey interface{} `field:"optional" json:"primaryKey" yaml:"primaryKey"`
	// `AWS::Serverless::SimpleTable.ProvisionedThroughput`.
	ProvisionedThroughput interface{} `field:"optional" json:"provisionedThroughput" yaml:"provisionedThroughput"`
	// `AWS::Serverless::SimpleTable.SSESpecification`.
	SseSpecification interface{} `field:"optional" json:"sseSpecification" yaml:"sseSpecification"`
	// `AWS::Serverless::SimpleTable.TableName`.
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
	// `AWS::Serverless::SimpleTable.Tags`.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::Serverless::StateMachine`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var definition interface{}
//
//   cfnStateMachine := awscdk.Aws_sam.NewCfnStateMachine(this, jsii.String("MyCfnStateMachine"), &cfnStateMachineProps{
//   	definition: definition,
//   	definitionSubstitutions: map[string]*string{
//   		"definitionSubstitutionsKey": jsii.String("definitionSubstitutions"),
//   	},
//   	definitionUri: jsii.String("definitionUri"),
//   	events: map[string]interface{}{
//   		"eventsKey": &EventSourceProperty{
//   			"properties": &CloudWatchEventEventProperty{
//   				"method": jsii.String("method"),
//   				"path": jsii.String("path"),
//
//   				// the properties below are optional
//   				"restApiId": jsii.String("restApiId"),
//   			},
//   			"type": jsii.String("type"),
//   		},
//   	},
//   	logging: &loggingConfigurationProperty{
//   		destinations: []interface{}{
//   			&logDestinationProperty{
//   				cloudWatchLogsLogGroup: &cloudWatchLogsLogGroupProperty{
//   					logGroupArn: jsii.String("logGroupArn"),
//   				},
//   			},
//   		},
//   		includeExecutionData: jsii.Boolean(false),
//   		level: jsii.String("level"),
//   	},
//   	name: jsii.String("name"),
//   	permissionsBoundaries: jsii.String("permissionsBoundaries"),
//   	policies: jsii.String("policies"),
//   	role: jsii.String("role"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	tracing: &tracingConfigurationProperty{
//   		enabled: jsii.Boolean(false),
//   	},
//   	type: jsii.String("type"),
//   })
//
type CfnStateMachine interface {
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
	// `AWS::Serverless::StateMachine.Definition`.
	Definition() interface{}
	SetDefinition(val interface{})
	// `AWS::Serverless::StateMachine.DefinitionSubstitutions`.
	DefinitionSubstitutions() interface{}
	SetDefinitionSubstitutions(val interface{})
	// `AWS::Serverless::StateMachine.DefinitionUri`.
	DefinitionUri() interface{}
	SetDefinitionUri(val interface{})
	// `AWS::Serverless::StateMachine.Events`.
	Events() interface{}
	SetEvents(val interface{})
	// `AWS::Serverless::StateMachine.Logging`.
	Logging() interface{}
	SetLogging(val interface{})
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
	// `AWS::Serverless::StateMachine.Name`.
	Name() *string
	SetName(val *string)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// `AWS::Serverless::StateMachine.PermissionsBoundaries`.
	PermissionsBoundaries() *string
	SetPermissionsBoundaries(val *string)
	// `AWS::Serverless::StateMachine.Policies`.
	Policies() interface{}
	SetPolicies(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// `AWS::Serverless::StateMachine.Role`.
	Role() *string
	SetRole(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// `AWS::Serverless::StateMachine.Tags`.
	Tags() awscdk.TagManager
	// `AWS::Serverless::StateMachine.Tracing`.
	Tracing() interface{}
	SetTracing(val interface{})
	// `AWS::Serverless::StateMachine.Type`.
	Type() *string
	SetType(val *string)
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

// The jsii proxy struct for CfnStateMachine
type jsiiProxy_CfnStateMachine struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnStateMachine) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStateMachine) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStateMachine) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStateMachine) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStateMachine) Definition() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"definition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStateMachine) DefinitionSubstitutions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"definitionSubstitutions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStateMachine) DefinitionUri() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"definitionUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStateMachine) Events() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"events",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStateMachine) Logging() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStateMachine) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStateMachine) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStateMachine) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStateMachine) PermissionsBoundaries() *string {
	var returns *string
	_jsii_.Get(
		j,
		"permissionsBoundaries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStateMachine) Policies() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"policies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStateMachine) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStateMachine) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStateMachine) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStateMachine) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStateMachine) Tracing() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tracing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStateMachine) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStateMachine) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Serverless::StateMachine`.
func NewCfnStateMachine(scope awscdk.Construct, id *string, props *CfnStateMachineProps) CfnStateMachine {
	_init_.Initialize()

	j := jsiiProxy_CfnStateMachine{}

	_jsii_.Create(
		"monocdk.aws_sam.CfnStateMachine",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Serverless::StateMachine`.
func NewCfnStateMachine_Override(c CfnStateMachine, scope awscdk.Construct, id *string, props *CfnStateMachineProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_sam.CfnStateMachine",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnStateMachine) SetDefinition(val interface{}) {
	_jsii_.Set(
		j,
		"definition",
		val,
	)
}

func (j *jsiiProxy_CfnStateMachine) SetDefinitionSubstitutions(val interface{}) {
	_jsii_.Set(
		j,
		"definitionSubstitutions",
		val,
	)
}

func (j *jsiiProxy_CfnStateMachine) SetDefinitionUri(val interface{}) {
	_jsii_.Set(
		j,
		"definitionUri",
		val,
	)
}

func (j *jsiiProxy_CfnStateMachine) SetEvents(val interface{}) {
	_jsii_.Set(
		j,
		"events",
		val,
	)
}

func (j *jsiiProxy_CfnStateMachine) SetLogging(val interface{}) {
	_jsii_.Set(
		j,
		"logging",
		val,
	)
}

func (j *jsiiProxy_CfnStateMachine) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnStateMachine) SetPermissionsBoundaries(val *string) {
	_jsii_.Set(
		j,
		"permissionsBoundaries",
		val,
	)
}

func (j *jsiiProxy_CfnStateMachine) SetPolicies(val interface{}) {
	_jsii_.Set(
		j,
		"policies",
		val,
	)
}

func (j *jsiiProxy_CfnStateMachine) SetRole(val *string) {
	_jsii_.Set(
		j,
		"role",
		val,
	)
}

func (j *jsiiProxy_CfnStateMachine) SetTracing(val interface{}) {
	_jsii_.Set(
		j,
		"tracing",
		val,
	)
}

func (j *jsiiProxy_CfnStateMachine) SetType(val *string) {
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
func CfnStateMachine_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sam.CfnStateMachine",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnStateMachine_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sam.CfnStateMachine",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnStateMachine_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sam.CfnStateMachine",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnStateMachine_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_sam.CfnStateMachine",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func CfnStateMachine_REQUIRED_TRANSFORM() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_sam.CfnStateMachine",
		"REQUIRED_TRANSFORM",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnStateMachine) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnStateMachine) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnStateMachine) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnStateMachine) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnStateMachine) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnStateMachine) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnStateMachine) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnStateMachine) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStateMachine) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStateMachine) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnStateMachine) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnStateMachine) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnStateMachine) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStateMachine) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnStateMachine) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnStateMachine) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStateMachine) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStateMachine) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnStateMachine) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStateMachine) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStateMachine) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   apiEventProperty := &apiEventProperty{
//   	method: jsii.String("method"),
//   	path: jsii.String("path"),
//
//   	// the properties below are optional
//   	restApiId: jsii.String("restApiId"),
//   }
//
type CfnStateMachine_ApiEventProperty struct {
	// `CfnStateMachine.ApiEventProperty.Method`.
	Method *string `field:"required" json:"method" yaml:"method"`
	// `CfnStateMachine.ApiEventProperty.Path`.
	Path *string `field:"required" json:"path" yaml:"path"`
	// `CfnStateMachine.ApiEventProperty.RestApiId`.
	RestApiId *string `field:"optional" json:"restApiId" yaml:"restApiId"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var pattern interface{}
//
//   cloudWatchEventEventProperty := &cloudWatchEventEventProperty{
//   	pattern: pattern,
//
//   	// the properties below are optional
//   	eventBusName: jsii.String("eventBusName"),
//   	input: jsii.String("input"),
//   	inputPath: jsii.String("inputPath"),
//   }
//
type CfnStateMachine_CloudWatchEventEventProperty struct {
	// `CfnStateMachine.CloudWatchEventEventProperty.Pattern`.
	Pattern interface{} `field:"required" json:"pattern" yaml:"pattern"`
	// `CfnStateMachine.CloudWatchEventEventProperty.EventBusName`.
	EventBusName *string `field:"optional" json:"eventBusName" yaml:"eventBusName"`
	// `CfnStateMachine.CloudWatchEventEventProperty.Input`.
	Input *string `field:"optional" json:"input" yaml:"input"`
	// `CfnStateMachine.CloudWatchEventEventProperty.InputPath`.
	InputPath *string `field:"optional" json:"inputPath" yaml:"inputPath"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudWatchLogsLogGroupProperty := &cloudWatchLogsLogGroupProperty{
//   	logGroupArn: jsii.String("logGroupArn"),
//   }
//
type CfnStateMachine_CloudWatchLogsLogGroupProperty struct {
	// `CfnStateMachine.CloudWatchLogsLogGroupProperty.LogGroupArn`.
	LogGroupArn *string `field:"required" json:"logGroupArn" yaml:"logGroupArn"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var pattern interface{}
//
//   eventBridgeRuleEventProperty := &eventBridgeRuleEventProperty{
//   	pattern: pattern,
//
//   	// the properties below are optional
//   	eventBusName: jsii.String("eventBusName"),
//   	input: jsii.String("input"),
//   	inputPath: jsii.String("inputPath"),
//   }
//
type CfnStateMachine_EventBridgeRuleEventProperty struct {
	// `CfnStateMachine.EventBridgeRuleEventProperty.Pattern`.
	Pattern interface{} `field:"required" json:"pattern" yaml:"pattern"`
	// `CfnStateMachine.EventBridgeRuleEventProperty.EventBusName`.
	EventBusName *string `field:"optional" json:"eventBusName" yaml:"eventBusName"`
	// `CfnStateMachine.EventBridgeRuleEventProperty.Input`.
	Input *string `field:"optional" json:"input" yaml:"input"`
	// `CfnStateMachine.EventBridgeRuleEventProperty.InputPath`.
	InputPath *string `field:"optional" json:"inputPath" yaml:"inputPath"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventSourceProperty := &eventSourceProperty{
//   	properties: &cloudWatchEventEventProperty{
//   		method: jsii.String("method"),
//   		path: jsii.String("path"),
//
//   		// the properties below are optional
//   		restApiId: jsii.String("restApiId"),
//   	},
//   	type: jsii.String("type"),
//   }
//
type CfnStateMachine_EventSourceProperty struct {
	// `CfnStateMachine.EventSourceProperty.Properties`.
	Properties interface{} `field:"required" json:"properties" yaml:"properties"`
	// `CfnStateMachine.EventSourceProperty.Type`.
	Type *string `field:"required" json:"type" yaml:"type"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   functionSAMPTProperty := &functionSAMPTProperty{
//   	functionName: jsii.String("functionName"),
//   }
//
type CfnStateMachine_FunctionSAMPTProperty struct {
	// `CfnStateMachine.FunctionSAMPTProperty.FunctionName`.
	FunctionName *string `field:"required" json:"functionName" yaml:"functionName"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var statement interface{}
//
//   iAMPolicyDocumentProperty := map[string]interface{}{
//   	"statement": statement,
//   	"version": jsii.String("version"),
//   }
//
type CfnStateMachine_IAMPolicyDocumentProperty struct {
	// `CfnStateMachine.IAMPolicyDocumentProperty.Statement`.
	Statement interface{} `field:"required" json:"statement" yaml:"statement"`
	// `CfnStateMachine.IAMPolicyDocumentProperty.Version`.
	Version *string `field:"required" json:"version" yaml:"version"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   logDestinationProperty := &logDestinationProperty{
//   	cloudWatchLogsLogGroup: &cloudWatchLogsLogGroupProperty{
//   		logGroupArn: jsii.String("logGroupArn"),
//   	},
//   }
//
type CfnStateMachine_LogDestinationProperty struct {
	// `CfnStateMachine.LogDestinationProperty.CloudWatchLogsLogGroup`.
	CloudWatchLogsLogGroup interface{} `field:"required" json:"cloudWatchLogsLogGroup" yaml:"cloudWatchLogsLogGroup"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loggingConfigurationProperty := &loggingConfigurationProperty{
//   	destinations: []interface{}{
//   		&logDestinationProperty{
//   			cloudWatchLogsLogGroup: &cloudWatchLogsLogGroupProperty{
//   				logGroupArn: jsii.String("logGroupArn"),
//   			},
//   		},
//   	},
//   	includeExecutionData: jsii.Boolean(false),
//   	level: jsii.String("level"),
//   }
//
type CfnStateMachine_LoggingConfigurationProperty struct {
	// `CfnStateMachine.LoggingConfigurationProperty.Destinations`.
	Destinations interface{} `field:"required" json:"destinations" yaml:"destinations"`
	// `CfnStateMachine.LoggingConfigurationProperty.IncludeExecutionData`.
	IncludeExecutionData interface{} `field:"required" json:"includeExecutionData" yaml:"includeExecutionData"`
	// `CfnStateMachine.LoggingConfigurationProperty.Level`.
	Level *string `field:"required" json:"level" yaml:"level"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3LocationProperty := &s3LocationProperty{
//   	bucket: jsii.String("bucket"),
//   	key: jsii.String("key"),
//
//   	// the properties below are optional
//   	version: jsii.Number(123),
//   }
//
type CfnStateMachine_S3LocationProperty struct {
	// `CfnStateMachine.S3LocationProperty.Bucket`.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// `CfnStateMachine.S3LocationProperty.Key`.
	Key *string `field:"required" json:"key" yaml:"key"`
	// `CfnStateMachine.S3LocationProperty.Version`.
	Version *float64 `field:"optional" json:"version" yaml:"version"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sAMPolicyTemplateProperty := &sAMPolicyTemplateProperty{
//   	lambdaInvokePolicy: &functionSAMPTProperty{
//   		functionName: jsii.String("functionName"),
//   	},
//   	stepFunctionsExecutionPolicy: &stateMachineSAMPTProperty{
//   		stateMachineName: jsii.String("stateMachineName"),
//   	},
//   }
//
type CfnStateMachine_SAMPolicyTemplateProperty struct {
	// `CfnStateMachine.SAMPolicyTemplateProperty.LambdaInvokePolicy`.
	LambdaInvokePolicy interface{} `field:"optional" json:"lambdaInvokePolicy" yaml:"lambdaInvokePolicy"`
	// `CfnStateMachine.SAMPolicyTemplateProperty.StepFunctionsExecutionPolicy`.
	StepFunctionsExecutionPolicy interface{} `field:"optional" json:"stepFunctionsExecutionPolicy" yaml:"stepFunctionsExecutionPolicy"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scheduleEventProperty := &scheduleEventProperty{
//   	schedule: jsii.String("schedule"),
//
//   	// the properties below are optional
//   	input: jsii.String("input"),
//   }
//
type CfnStateMachine_ScheduleEventProperty struct {
	// `CfnStateMachine.ScheduleEventProperty.Schedule`.
	Schedule *string `field:"required" json:"schedule" yaml:"schedule"`
	// `CfnStateMachine.ScheduleEventProperty.Input`.
	Input *string `field:"optional" json:"input" yaml:"input"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stateMachineSAMPTProperty := &stateMachineSAMPTProperty{
//   	stateMachineName: jsii.String("stateMachineName"),
//   }
//
type CfnStateMachine_StateMachineSAMPTProperty struct {
	// `CfnStateMachine.StateMachineSAMPTProperty.StateMachineName`.
	StateMachineName *string `field:"required" json:"stateMachineName" yaml:"stateMachineName"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tracingConfigurationProperty := &tracingConfigurationProperty{
//   	enabled: jsii.Boolean(false),
//   }
//
type CfnStateMachine_TracingConfigurationProperty struct {
	// `CfnStateMachine.TracingConfigurationProperty.Enabled`.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

// Properties for defining a `CfnStateMachine`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var definition interface{}
//
//   cfnStateMachineProps := &cfnStateMachineProps{
//   	definition: definition,
//   	definitionSubstitutions: map[string]*string{
//   		"definitionSubstitutionsKey": jsii.String("definitionSubstitutions"),
//   	},
//   	definitionUri: jsii.String("definitionUri"),
//   	events: map[string]interface{}{
//   		"eventsKey": &EventSourceProperty{
//   			"properties": &CloudWatchEventEventProperty{
//   				"method": jsii.String("method"),
//   				"path": jsii.String("path"),
//
//   				// the properties below are optional
//   				"restApiId": jsii.String("restApiId"),
//   			},
//   			"type": jsii.String("type"),
//   		},
//   	},
//   	logging: &loggingConfigurationProperty{
//   		destinations: []interface{}{
//   			&logDestinationProperty{
//   				cloudWatchLogsLogGroup: &cloudWatchLogsLogGroupProperty{
//   					logGroupArn: jsii.String("logGroupArn"),
//   				},
//   			},
//   		},
//   		includeExecutionData: jsii.Boolean(false),
//   		level: jsii.String("level"),
//   	},
//   	name: jsii.String("name"),
//   	permissionsBoundaries: jsii.String("permissionsBoundaries"),
//   	policies: jsii.String("policies"),
//   	role: jsii.String("role"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	tracing: &tracingConfigurationProperty{
//   		enabled: jsii.Boolean(false),
//   	},
//   	type: jsii.String("type"),
//   }
//
type CfnStateMachineProps struct {
	// `AWS::Serverless::StateMachine.Definition`.
	Definition interface{} `field:"optional" json:"definition" yaml:"definition"`
	// `AWS::Serverless::StateMachine.DefinitionSubstitutions`.
	DefinitionSubstitutions interface{} `field:"optional" json:"definitionSubstitutions" yaml:"definitionSubstitutions"`
	// `AWS::Serverless::StateMachine.DefinitionUri`.
	DefinitionUri interface{} `field:"optional" json:"definitionUri" yaml:"definitionUri"`
	// `AWS::Serverless::StateMachine.Events`.
	Events interface{} `field:"optional" json:"events" yaml:"events"`
	// `AWS::Serverless::StateMachine.Logging`.
	Logging interface{} `field:"optional" json:"logging" yaml:"logging"`
	// `AWS::Serverless::StateMachine.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// `AWS::Serverless::StateMachine.PermissionsBoundaries`.
	PermissionsBoundaries *string `field:"optional" json:"permissionsBoundaries" yaml:"permissionsBoundaries"`
	// `AWS::Serverless::StateMachine.Policies`.
	Policies interface{} `field:"optional" json:"policies" yaml:"policies"`
	// `AWS::Serverless::StateMachine.Role`.
	Role *string `field:"optional" json:"role" yaml:"role"`
	// `AWS::Serverless::StateMachine.Tags`.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// `AWS::Serverless::StateMachine.Tracing`.
	Tracing interface{} `field:"optional" json:"tracing" yaml:"tracing"`
	// `AWS::Serverless::StateMachine.Type`.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

