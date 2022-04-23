package awsapprunner

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsapprunner/internal"
	"github.com/aws/aws-cdk-go/awscdk/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/awsecrassets"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/constructs-go/constructs/v3"
)

// Properties of the image repository for `Source.fromAsset()`.
//
// Example:
//   import assets "github.com/aws/aws-cdk-go/awscdk"
//
//   imageAsset := assets.NewDockerImageAsset(this, jsii.String("ImageAssets"), &dockerImageAssetProps{
//   	directory: path.join(__dirname, jsii.String("./docker.assets")),
//   })
//   apprunner.NewService(this, jsii.String("Service"), &serviceProps{
//   	source: apprunner.source.fromAsset(&assetProps{
//   		imageConfiguration: &imageConfiguration{
//   			port: jsii.Number(8000),
//   		},
//   		asset: imageAsset,
//   	}),
//   })
//
// Experimental.
type AssetProps struct {
	// Represents the docker image asset.
	// Experimental.
	Asset awsecrassets.DockerImageAsset `json:"asset" yaml:"asset"`
	// The image configuration for the image built from the asset.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-imageconfiguration.html#cfn-apprunner-service-imageconfiguration-port
	//
	// Experimental.
	ImageConfiguration *ImageConfiguration `json:"imageConfiguration" yaml:"imageConfiguration"`
}

// Represents the source from local assets.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import apprunner "github.com/aws/aws-cdk-go/awscdk/aws_apprunner"import awscdk "github.com/aws/aws-cdk-go/awscdk"import ecr_assets "github.com/aws/aws-cdk-go/awscdk/aws_ecr_assets"
//
//   var dockerImageAsset dockerImageAsset
//   assetSource := apprunner.NewAssetSource(&assetProps{
//   	asset: dockerImageAsset,
//
//   	// the properties below are optional
//   	imageConfiguration: &imageConfiguration{
//   		environment: map[string]*string{
//   			"environmentKey": jsii.String("environment"),
//   		},
//   		port: jsii.Number(123),
//   		startCommand: jsii.String("startCommand"),
//   	},
//   })
//
// Experimental.
type AssetSource interface {
	Source
	// Called when the Job is initialized to allow this object to bind.
	// Experimental.
	Bind(_scope constructs.Construct) *SourceConfig
}

// The jsii proxy struct for AssetSource
type jsiiProxy_AssetSource struct {
	jsiiProxy_Source
}

// Experimental.
func NewAssetSource(props *AssetProps) AssetSource {
	_init_.Initialize()

	j := jsiiProxy_AssetSource{}

	_jsii_.Create(
		"monocdk.aws_apprunner.AssetSource",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewAssetSource_Override(a AssetSource, props *AssetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_apprunner.AssetSource",
		[]interface{}{props},
		a,
	)
}

// Source from local assets.
// Experimental.
func AssetSource_FromAsset(props *AssetProps) AssetSource {
	_init_.Initialize()

	var returns AssetSource

	_jsii_.StaticInvoke(
		"monocdk.aws_apprunner.AssetSource",
		"fromAsset",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Source from the ECR repository.
// Experimental.
func AssetSource_FromEcr(props *EcrProps) EcrSource {
	_init_.Initialize()

	var returns EcrSource

	_jsii_.StaticInvoke(
		"monocdk.aws_apprunner.AssetSource",
		"fromEcr",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Source from the ECR Public repository.
// Experimental.
func AssetSource_FromEcrPublic(props *EcrPublicProps) EcrPublicSource {
	_init_.Initialize()

	var returns EcrPublicSource

	_jsii_.StaticInvoke(
		"monocdk.aws_apprunner.AssetSource",
		"fromEcrPublic",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Source from the GitHub repository.
// Experimental.
func AssetSource_FromGitHub(props *GithubRepositoryProps) GithubSource {
	_init_.Initialize()

	var returns GithubSource

	_jsii_.StaticInvoke(
		"monocdk.aws_apprunner.AssetSource",
		"fromGitHub",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssetSource) Bind(_scope constructs.Construct) *SourceConfig {
	var returns *SourceConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

// A CloudFormation `AWS::AppRunner::ObservabilityConfiguration`.
//
// Specify an AWS App Runner observability configuration by using the `AWS::AppRunner::ObservabilityConfiguration` resource in an AWS CloudFormation template.
//
// The `AWS::AppRunner::ObservabilityConfiguration` resource is an AWS App Runner resource type that specifies an App Runner observability configuration.
//
// App Runner requires this resource when you specify App Runner services and you want to enable non-default observability features. You can share an observability configuration across multiple services.
//
// Create multiple revisions of a configuration by specifying this resource multiple times using the same `ObservabilityConfigurationName` . App Runner creates multiple resources with incremental `ObservabilityConfigurationRevision` values. When you specify a service and configure an observability configuration resource, the service uses the latest active revision of the observability configuration by default. You can optionally configure the service to use a specific revision.
//
// The observability configuration resource is designed to configure multiple features (currently one feature, tracing). This resource takes optional parameters that describe the configuration of these features (currently one parameter, `TraceConfiguration` ). If you don't specify a feature parameter, App Runner doesn't enable the feature.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import apprunner "github.com/aws/aws-cdk-go/awscdk/aws_apprunner"
//   cfnObservabilityConfiguration := apprunner.NewCfnObservabilityConfiguration(this, jsii.String("MyCfnObservabilityConfiguration"), &cfnObservabilityConfigurationProps{
//   	observabilityConfigurationName: jsii.String("observabilityConfigurationName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	traceConfiguration: &traceConfigurationProperty{
//   		vendor: jsii.String("vendor"),
//   	},
//   })
//
type CfnObservabilityConfiguration interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// It's set to `true` for the configuration with the highest `Revision` among all configurations that share the same `ObservabilityConfigurationName` .
	//
	// It's set to `false` otherwise.
	AttrLatest() awscdk.IResolvable
	// The Amazon Resource Name (ARN) of this observability configuration.
	AttrObservabilityConfigurationArn() *string
	// The revision of this observability configuration.
	//
	// It's unique among all the active configurations ( `"Status": "ACTIVE"` ) that share the same `ObservabilityConfigurationName` .
	AttrObservabilityConfigurationRevision() *float64
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
	// A name for the observability configuration.
	//
	// When you use it for the first time in an AWS Region , App Runner creates revision number `1` of this name. When you use the same name in subsequent calls, App Runner creates incremental revisions of the configuration.
	//
	// > The name `DefaultConfiguration` is reserved. You can't use it to create a new observability configuration, and you can't create a revision of it.
	// >
	// > When you want to use your own observability configuration for your App Runner service, *create a configuration with a different name* , and then provide it when you create or update your service.
	//
	// If you don't specify a name, AWS CloudFormation generates a name for your observability configuration.
	ObservabilityConfigurationName() *string
	SetObservabilityConfigurationName(val *string)
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
	// A list of metadata items that you can associate with your observability configuration resource.
	//
	// A tag is a key-value pair.
	Tags() awscdk.TagManager
	// The configuration of the tracing feature within this observability configuration.
	//
	// If you don't specify it, App Runner doesn't enable tracing.
	TraceConfiguration() interface{}
	SetTraceConfiguration(val interface{})
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

// The jsii proxy struct for CfnObservabilityConfiguration
type jsiiProxy_CfnObservabilityConfiguration struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnObservabilityConfiguration) AttrLatest() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrLatest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObservabilityConfiguration) AttrObservabilityConfigurationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrObservabilityConfigurationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObservabilityConfiguration) AttrObservabilityConfigurationRevision() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrObservabilityConfigurationRevision",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObservabilityConfiguration) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObservabilityConfiguration) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObservabilityConfiguration) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObservabilityConfiguration) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObservabilityConfiguration) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObservabilityConfiguration) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObservabilityConfiguration) ObservabilityConfigurationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"observabilityConfigurationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObservabilityConfiguration) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObservabilityConfiguration) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObservabilityConfiguration) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObservabilityConfiguration) TraceConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"traceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObservabilityConfiguration) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::AppRunner::ObservabilityConfiguration`.
func NewCfnObservabilityConfiguration(scope awscdk.Construct, id *string, props *CfnObservabilityConfigurationProps) CfnObservabilityConfiguration {
	_init_.Initialize()

	j := jsiiProxy_CfnObservabilityConfiguration{}

	_jsii_.Create(
		"monocdk.aws_apprunner.CfnObservabilityConfiguration",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AppRunner::ObservabilityConfiguration`.
func NewCfnObservabilityConfiguration_Override(c CfnObservabilityConfiguration, scope awscdk.Construct, id *string, props *CfnObservabilityConfigurationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_apprunner.CfnObservabilityConfiguration",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnObservabilityConfiguration) SetObservabilityConfigurationName(val *string) {
	_jsii_.Set(
		j,
		"observabilityConfigurationName",
		val,
	)
}

func (j *jsiiProxy_CfnObservabilityConfiguration) SetTraceConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"traceConfiguration",
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
func CfnObservabilityConfiguration_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_apprunner.CfnObservabilityConfiguration",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnObservabilityConfiguration_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_apprunner.CfnObservabilityConfiguration",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnObservabilityConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_apprunner.CfnObservabilityConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnObservabilityConfiguration_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_apprunner.CfnObservabilityConfiguration",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnObservabilityConfiguration) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnObservabilityConfiguration) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnObservabilityConfiguration) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnObservabilityConfiguration) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnObservabilityConfiguration) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnObservabilityConfiguration) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnObservabilityConfiguration) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnObservabilityConfiguration) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnObservabilityConfiguration) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnObservabilityConfiguration) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnObservabilityConfiguration) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnObservabilityConfiguration) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnObservabilityConfiguration) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnObservabilityConfiguration) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnObservabilityConfiguration) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnObservabilityConfiguration) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnObservabilityConfiguration) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnObservabilityConfiguration) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnObservabilityConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnObservabilityConfiguration) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnObservabilityConfiguration) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Describes the configuration of the tracing feature within an AWS App Runner observability configuration.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import apprunner "github.com/aws/aws-cdk-go/awscdk/aws_apprunner"
//   traceConfigurationProperty := &traceConfigurationProperty{
//   	vendor: jsii.String("vendor"),
//   }
//
type CfnObservabilityConfiguration_TraceConfigurationProperty struct {
	// The implementation provider chosen for tracing App Runner services.
	Vendor *string `json:"vendor" yaml:"vendor"`
}

// Properties for defining a `CfnObservabilityConfiguration`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import apprunner "github.com/aws/aws-cdk-go/awscdk/aws_apprunner"
//   cfnObservabilityConfigurationProps := &cfnObservabilityConfigurationProps{
//   	observabilityConfigurationName: jsii.String("observabilityConfigurationName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	traceConfiguration: &traceConfigurationProperty{
//   		vendor: jsii.String("vendor"),
//   	},
//   }
//
type CfnObservabilityConfigurationProps struct {
	// A name for the observability configuration.
	//
	// When you use it for the first time in an AWS Region , App Runner creates revision number `1` of this name. When you use the same name in subsequent calls, App Runner creates incremental revisions of the configuration.
	//
	// > The name `DefaultConfiguration` is reserved. You can't use it to create a new observability configuration, and you can't create a revision of it.
	// >
	// > When you want to use your own observability configuration for your App Runner service, *create a configuration with a different name* , and then provide it when you create or update your service.
	//
	// If you don't specify a name, AWS CloudFormation generates a name for your observability configuration.
	ObservabilityConfigurationName *string `json:"observabilityConfigurationName" yaml:"observabilityConfigurationName"`
	// A list of metadata items that you can associate with your observability configuration resource.
	//
	// A tag is a key-value pair.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
	// The configuration of the tracing feature within this observability configuration.
	//
	// If you don't specify it, App Runner doesn't enable tracing.
	TraceConfiguration interface{} `json:"traceConfiguration" yaml:"traceConfiguration"`
}

// A CloudFormation `AWS::AppRunner::Service`.
//
// Specify an AWS App Runner service by using the `AWS::AppRunner::Service` resource in an AWS CloudFormation template.
//
// The `AWS::AppRunner::Service` resource is an AWS App Runner resource type that specifies an App Runner service.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import apprunner "github.com/aws/aws-cdk-go/awscdk/aws_apprunner"
//   cfnService := apprunner.NewCfnService(this, jsii.String("MyCfnService"), &cfnServiceProps{
//   	sourceConfiguration: &sourceConfigurationProperty{
//   		authenticationConfiguration: &authenticationConfigurationProperty{
//   			accessRoleArn: jsii.String("accessRoleArn"),
//   			connectionArn: jsii.String("connectionArn"),
//   		},
//   		autoDeploymentsEnabled: jsii.Boolean(false),
//   		codeRepository: &codeRepositoryProperty{
//   			repositoryUrl: jsii.String("repositoryUrl"),
//   			sourceCodeVersion: &sourceCodeVersionProperty{
//   				type: jsii.String("type"),
//   				value: jsii.String("value"),
//   			},
//
//   			// the properties below are optional
//   			codeConfiguration: &codeConfigurationProperty{
//   				configurationSource: jsii.String("configurationSource"),
//
//   				// the properties below are optional
//   				codeConfigurationValues: &codeConfigurationValuesProperty{
//   					runtime: jsii.String("runtime"),
//
//   					// the properties below are optional
//   					buildCommand: jsii.String("buildCommand"),
//   					port: jsii.String("port"),
//   					runtimeEnvironmentVariables: []interface{}{
//   						&keyValuePairProperty{
//   							name: jsii.String("name"),
//   							value: jsii.String("value"),
//   						},
//   					},
//   					startCommand: jsii.String("startCommand"),
//   				},
//   			},
//   		},
//   		imageRepository: &imageRepositoryProperty{
//   			imageIdentifier: jsii.String("imageIdentifier"),
//   			imageRepositoryType: jsii.String("imageRepositoryType"),
//
//   			// the properties below are optional
//   			imageConfiguration: &imageConfigurationProperty{
//   				port: jsii.String("port"),
//   				runtimeEnvironmentVariables: []interface{}{
//   					&keyValuePairProperty{
//   						name: jsii.String("name"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   				startCommand: jsii.String("startCommand"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	autoScalingConfigurationArn: jsii.String("autoScalingConfigurationArn"),
//   	encryptionConfiguration: &encryptionConfigurationProperty{
//   		kmsKey: jsii.String("kmsKey"),
//   	},
//   	healthCheckConfiguration: &healthCheckConfigurationProperty{
//   		healthyThreshold: jsii.Number(123),
//   		interval: jsii.Number(123),
//   		path: jsii.String("path"),
//   		protocol: jsii.String("protocol"),
//   		timeout: jsii.Number(123),
//   		unhealthyThreshold: jsii.Number(123),
//   	},
//   	instanceConfiguration: &instanceConfigurationProperty{
//   		cpu: jsii.String("cpu"),
//   		instanceRoleArn: jsii.String("instanceRoleArn"),
//   		memory: jsii.String("memory"),
//   	},
//   	networkConfiguration: &networkConfigurationProperty{
//   		egressConfiguration: &egressConfigurationProperty{
//   			egressType: jsii.String("egressType"),
//
//   			// the properties below are optional
//   			vpcConnectorArn: jsii.String("vpcConnectorArn"),
//   		},
//   	},
//   	observabilityConfiguration: &serviceObservabilityConfigurationProperty{
//   		observabilityEnabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		observabilityConfigurationArn: jsii.String("observabilityConfigurationArn"),
//   	},
//   	serviceName: jsii.String("serviceName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnService interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of this service.
	AttrServiceArn() *string
	// An ID that App Runner generated for this service.
	//
	// It's unique within the AWS Region .
	AttrServiceId() *string
	// A subdomain URL that App Runner generated for this service.
	//
	// You can use this URL to access your service web application.
	AttrServiceUrl() *string
	// The current state of the App Runner service. These particular values mean the following.
	//
	// - `CREATE_FAILED` – The service failed to create. To troubleshoot this failure, read the failure events and logs, change any parameters that need to be fixed, and retry the call to create the service.
	//
	// The failed service isn't usable, and still counts towards your service quota. When you're done analyzing the failure, delete the service.
	// - `DELETE_FAILED` – The service failed to delete and can't be successfully recovered. Retry the service deletion call to ensure that all related resources are removed.
	AttrStatus() *string
	// The Amazon Resource Name (ARN) of an App Runner automatic scaling configuration resource that you want to associate with your service.
	//
	// If not provided, App Runner associates the latest revision of a default auto scaling configuration.
	//
	// Specify an ARN with a name and a revision number to associate that revision. For example: `arn:aws:apprunner:us-east-1:123456789012:autoscalingconfiguration/high-availability/3`
	//
	// Specify just the name to associate the latest revision. For example: `arn:aws:apprunner:us-east-1:123456789012:autoscalingconfiguration/high-availability`
	AutoScalingConfigurationArn() *string
	SetAutoScalingConfigurationArn(val *string)
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
	// An optional custom encryption key that App Runner uses to encrypt the copy of your source repository that it maintains and your service logs.
	//
	// By default, App Runner uses an AWS managed key .
	EncryptionConfiguration() interface{}
	SetEncryptionConfiguration(val interface{})
	// The settings for the health check that AWS App Runner performs to monitor the health of the App Runner service.
	HealthCheckConfiguration() interface{}
	SetHealthCheckConfiguration(val interface{})
	// The runtime configuration of instances (scaling units) of your service.
	InstanceConfiguration() interface{}
	SetInstanceConfiguration(val interface{})
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
	// Configuration settings related to network traffic of the web application that the App Runner service runs.
	NetworkConfiguration() interface{}
	SetNetworkConfiguration(val interface{})
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// The observability configuration of your service.
	ObservabilityConfiguration() interface{}
	SetObservabilityConfiguration(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// A name for the App Runner service.
	//
	// It must be unique across all the running App Runner services in your AWS account in the AWS Region .
	//
	// If you don't specify a name, AWS CloudFormation generates a name for your service.
	ServiceName() *string
	SetServiceName(val *string)
	// The source to deploy to the App Runner service.
	//
	// It can be a code or an image repository.
	SourceConfiguration() interface{}
	SetSourceConfiguration(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// An optional list of metadata items that you can associate with the App Runner service resource.
	//
	// A tag is a key-value pair.
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

// The jsii proxy struct for CfnService
type jsiiProxy_CfnService struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnService) AttrServiceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrServiceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) AttrServiceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrServiceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) AttrServiceUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrServiceUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) AttrStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) AutoScalingConfigurationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoScalingConfigurationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) EncryptionConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) HealthCheckConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"healthCheckConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) InstanceConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instanceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) NetworkConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) ObservabilityConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"observabilityConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) ServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) SourceConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::AppRunner::Service`.
func NewCfnService(scope awscdk.Construct, id *string, props *CfnServiceProps) CfnService {
	_init_.Initialize()

	j := jsiiProxy_CfnService{}

	_jsii_.Create(
		"monocdk.aws_apprunner.CfnService",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AppRunner::Service`.
func NewCfnService_Override(c CfnService, scope awscdk.Construct, id *string, props *CfnServiceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_apprunner.CfnService",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnService) SetAutoScalingConfigurationArn(val *string) {
	_jsii_.Set(
		j,
		"autoScalingConfigurationArn",
		val,
	)
}

func (j *jsiiProxy_CfnService) SetEncryptionConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"encryptionConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnService) SetHealthCheckConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"healthCheckConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnService) SetInstanceConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"instanceConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnService) SetNetworkConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"networkConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnService) SetObservabilityConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"observabilityConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnService) SetServiceName(val *string) {
	_jsii_.Set(
		j,
		"serviceName",
		val,
	)
}

func (j *jsiiProxy_CfnService) SetSourceConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"sourceConfiguration",
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
func CfnService_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_apprunner.CfnService",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnService_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_apprunner.CfnService",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnService_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_apprunner.CfnService",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnService_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_apprunner.CfnService",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnService) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnService) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnService) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnService) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnService) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnService) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnService) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnService) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnService) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnService) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnService) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnService) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnService) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnService) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnService) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnService) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnService) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnService) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnService) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnService) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnService) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Describes resources needed to authenticate access to some source repositories.
//
// The specific resource depends on the repository provider.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import apprunner "github.com/aws/aws-cdk-go/awscdk/aws_apprunner"
//   authenticationConfigurationProperty := &authenticationConfigurationProperty{
//   	accessRoleArn: jsii.String("accessRoleArn"),
//   	connectionArn: jsii.String("connectionArn"),
//   }
//
type CfnService_AuthenticationConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of the IAM role that grants the App Runner service access to a source repository.
	//
	// It's required for ECR image repositories (but not for ECR Public repositories).
	AccessRoleArn *string `json:"accessRoleArn" yaml:"accessRoleArn"`
	// The Amazon Resource Name (ARN) of the App Runner connection that enables the App Runner service to connect to a source repository.
	//
	// It's required for GitHub code repositories.
	ConnectionArn *string `json:"connectionArn" yaml:"connectionArn"`
}

// Describes the configuration that AWS App Runner uses to build and run an App Runner service from a source code repository.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import apprunner "github.com/aws/aws-cdk-go/awscdk/aws_apprunner"
//   codeConfigurationProperty := &codeConfigurationProperty{
//   	configurationSource: jsii.String("configurationSource"),
//
//   	// the properties below are optional
//   	codeConfigurationValues: &codeConfigurationValuesProperty{
//   		runtime: jsii.String("runtime"),
//
//   		// the properties below are optional
//   		buildCommand: jsii.String("buildCommand"),
//   		port: jsii.String("port"),
//   		runtimeEnvironmentVariables: []interface{}{
//   			&keyValuePairProperty{
//   				name: jsii.String("name"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		startCommand: jsii.String("startCommand"),
//   	},
//   }
//
type CfnService_CodeConfigurationProperty struct {
	// The source of the App Runner configuration. Values are interpreted as follows:.
	//
	// - `REPOSITORY` – App Runner reads configuration values from the `apprunner.yaml` file in the source code repository and ignores `CodeConfigurationValues` .
	// - `API` – App Runner uses configuration values provided in `CodeConfigurationValues` and ignores the `apprunner.yaml` file in the source code repository.
	ConfigurationSource *string `json:"configurationSource" yaml:"configurationSource"`
	// The basic configuration for building and running the App Runner service.
	//
	// Use it to quickly launch an App Runner service without providing a `apprunner.yaml` file in the source code repository (or ignoring the file if it exists).
	CodeConfigurationValues interface{} `json:"codeConfigurationValues" yaml:"codeConfigurationValues"`
}

// Describes the basic configuration needed for building and running an AWS App Runner service.
//
// This type doesn't support the full set of possible configuration options. Fur full configuration capabilities, use a `apprunner.yaml` file in the source code repository.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import apprunner "github.com/aws/aws-cdk-go/awscdk/aws_apprunner"
//   codeConfigurationValuesProperty := &codeConfigurationValuesProperty{
//   	runtime: jsii.String("runtime"),
//
//   	// the properties below are optional
//   	buildCommand: jsii.String("buildCommand"),
//   	port: jsii.String("port"),
//   	runtimeEnvironmentVariables: []interface{}{
//   		&keyValuePairProperty{
//   			name: jsii.String("name"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	startCommand: jsii.String("startCommand"),
//   }
//
type CfnService_CodeConfigurationValuesProperty struct {
	// A runtime environment type for building and running an App Runner service.
	//
	// It represents a programming language runtime.
	Runtime *string `json:"runtime" yaml:"runtime"`
	// The command App Runner runs to build your application.
	BuildCommand *string `json:"buildCommand" yaml:"buildCommand"`
	// The port that your application listens to in the container.
	//
	// Default: `8080`.
	Port *string `json:"port" yaml:"port"`
	// The environment variables that are available to your running App Runner service.
	//
	// An array of key-value pairs. Keys with a prefix of `AWSAPPRUNNER` are reserved for system use and aren't valid.
	RuntimeEnvironmentVariables interface{} `json:"runtimeEnvironmentVariables" yaml:"runtimeEnvironmentVariables"`
	// The command App Runner runs to start your application.
	StartCommand *string `json:"startCommand" yaml:"startCommand"`
}

// Describes a source code repository.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import apprunner "github.com/aws/aws-cdk-go/awscdk/aws_apprunner"
//   codeRepositoryProperty := &codeRepositoryProperty{
//   	repositoryUrl: jsii.String("repositoryUrl"),
//   	sourceCodeVersion: &sourceCodeVersionProperty{
//   		type: jsii.String("type"),
//   		value: jsii.String("value"),
//   	},
//
//   	// the properties below are optional
//   	codeConfiguration: &codeConfigurationProperty{
//   		configurationSource: jsii.String("configurationSource"),
//
//   		// the properties below are optional
//   		codeConfigurationValues: &codeConfigurationValuesProperty{
//   			runtime: jsii.String("runtime"),
//
//   			// the properties below are optional
//   			buildCommand: jsii.String("buildCommand"),
//   			port: jsii.String("port"),
//   			runtimeEnvironmentVariables: []interface{}{
//   				&keyValuePairProperty{
//   					name: jsii.String("name"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   			startCommand: jsii.String("startCommand"),
//   		},
//   	},
//   }
//
type CfnService_CodeRepositoryProperty struct {
	// The location of the repository that contains the source code.
	RepositoryUrl *string `json:"repositoryUrl" yaml:"repositoryUrl"`
	// The version that should be used within the source code repository.
	SourceCodeVersion interface{} `json:"sourceCodeVersion" yaml:"sourceCodeVersion"`
	// Configuration for building and running the service from a source code repository.
	CodeConfiguration interface{} `json:"codeConfiguration" yaml:"codeConfiguration"`
}

// Describes configuration settings related to outbound network traffic of an AWS App Runner service.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import apprunner "github.com/aws/aws-cdk-go/awscdk/aws_apprunner"
//   egressConfigurationProperty := &egressConfigurationProperty{
//   	egressType: jsii.String("egressType"),
//
//   	// the properties below are optional
//   	vpcConnectorArn: jsii.String("vpcConnectorArn"),
//   }
//
type CfnService_EgressConfigurationProperty struct {
	// The type of egress configuration.
	//
	// Set to `DEFAULT` for access to resources hosted on public networks.
	//
	// Set to `VPC` to associate your service to a custom VPC specified by `VpcConnectorArn` .
	EgressType *string `json:"egressType" yaml:"egressType"`
	// The Amazon Resource Name (ARN) of the App Runner VPC connector that you want to associate with your App Runner service.
	//
	// Only valid when `EgressType = VPC` .
	VpcConnectorArn *string `json:"vpcConnectorArn" yaml:"vpcConnectorArn"`
}

// Describes a custom encryption key that AWS App Runner uses to encrypt copies of the source repository and service logs.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import apprunner "github.com/aws/aws-cdk-go/awscdk/aws_apprunner"
//   encryptionConfigurationProperty := &encryptionConfigurationProperty{
//   	kmsKey: jsii.String("kmsKey"),
//   }
//
type CfnService_EncryptionConfigurationProperty struct {
	// The ARN of the KMS key that's used for encryption.
	KmsKey *string `json:"kmsKey" yaml:"kmsKey"`
}

// Describes the settings for the health check that AWS App Runner performs to monitor the health of a service.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import apprunner "github.com/aws/aws-cdk-go/awscdk/aws_apprunner"
//   healthCheckConfigurationProperty := &healthCheckConfigurationProperty{
//   	healthyThreshold: jsii.Number(123),
//   	interval: jsii.Number(123),
//   	path: jsii.String("path"),
//   	protocol: jsii.String("protocol"),
//   	timeout: jsii.Number(123),
//   	unhealthyThreshold: jsii.Number(123),
//   }
//
type CfnService_HealthCheckConfigurationProperty struct {
	// The number of consecutive checks that must succeed before App Runner decides that the service is healthy.
	//
	// Default: `1`.
	HealthyThreshold *float64 `json:"healthyThreshold" yaml:"healthyThreshold"`
	// The time interval, in seconds, between health checks.
	//
	// Default: `5`.
	Interval *float64 `json:"interval" yaml:"interval"`
	// The URL that health check requests are sent to.
	//
	// `Path` is only applicable when you set `Protocol` to `HTTP` .
	//
	// Default: `"/"`.
	Path *string `json:"path" yaml:"path"`
	// The IP protocol that App Runner uses to perform health checks for your service.
	//
	// If you set `Protocol` to `HTTP` , App Runner sends health check requests to the HTTP path specified by `Path` .
	//
	// Default: `TCP`.
	Protocol *string `json:"protocol" yaml:"protocol"`
	// The time, in seconds, to wait for a health check response before deciding it failed.
	//
	// Default: `2`.
	Timeout *float64 `json:"timeout" yaml:"timeout"`
	// The number of consecutive checks that must fail before App Runner decides that the service is unhealthy.
	//
	// Default: `5`.
	UnhealthyThreshold *float64 `json:"unhealthyThreshold" yaml:"unhealthyThreshold"`
}

// Describes the configuration that AWS App Runner uses to run an App Runner service using an image pulled from a source image repository.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import apprunner "github.com/aws/aws-cdk-go/awscdk/aws_apprunner"
//   imageConfigurationProperty := &imageConfigurationProperty{
//   	port: jsii.String("port"),
//   	runtimeEnvironmentVariables: []interface{}{
//   		&keyValuePairProperty{
//   			name: jsii.String("name"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	startCommand: jsii.String("startCommand"),
//   }
//
type CfnService_ImageConfigurationProperty struct {
	// The port that your application listens to in the container.
	//
	// Default: `8080`.
	Port *string `json:"port" yaml:"port"`
	// Environment variables that are available to your running App Runner service.
	//
	// An array of key-value pairs. Keys with a prefix of `AWSAPPRUNNER` are reserved for system use and aren't valid.
	RuntimeEnvironmentVariables interface{} `json:"runtimeEnvironmentVariables" yaml:"runtimeEnvironmentVariables"`
	// An optional command that App Runner runs to start the application in the source image.
	//
	// If specified, this command overrides the Docker image’s default start command.
	StartCommand *string `json:"startCommand" yaml:"startCommand"`
}

// Describes a source image repository.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import apprunner "github.com/aws/aws-cdk-go/awscdk/aws_apprunner"
//   imageRepositoryProperty := &imageRepositoryProperty{
//   	imageIdentifier: jsii.String("imageIdentifier"),
//   	imageRepositoryType: jsii.String("imageRepositoryType"),
//
//   	// the properties below are optional
//   	imageConfiguration: &imageConfigurationProperty{
//   		port: jsii.String("port"),
//   		runtimeEnvironmentVariables: []interface{}{
//   			&keyValuePairProperty{
//   				name: jsii.String("name"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		startCommand: jsii.String("startCommand"),
//   	},
//   }
//
type CfnService_ImageRepositoryProperty struct {
	// The identifier of an image.
	//
	// For an image in Amazon Elastic Container Registry (Amazon ECR), this is an image name. For the image name format, see [Pulling an image](https://docs.aws.amazon.com/AmazonECR/latest/userguide/docker-pull-ecr-image.html) in the *Amazon ECR User Guide* .
	ImageIdentifier *string `json:"imageIdentifier" yaml:"imageIdentifier"`
	// The type of the image repository.
	//
	// This reflects the repository provider and whether the repository is private or public.
	ImageRepositoryType *string `json:"imageRepositoryType" yaml:"imageRepositoryType"`
	// Configuration for running the identified image.
	ImageConfiguration interface{} `json:"imageConfiguration" yaml:"imageConfiguration"`
}

// Describes the runtime configuration of an AWS App Runner service instance (scaling unit).
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import apprunner "github.com/aws/aws-cdk-go/awscdk/aws_apprunner"
//   instanceConfigurationProperty := &instanceConfigurationProperty{
//   	cpu: jsii.String("cpu"),
//   	instanceRoleArn: jsii.String("instanceRoleArn"),
//   	memory: jsii.String("memory"),
//   }
//
type CfnService_InstanceConfigurationProperty struct {
	// The number of CPU units reserved for each instance of your App Runner service.
	//
	// Default: `1 vCPU`.
	Cpu *string `json:"cpu" yaml:"cpu"`
	// The Amazon Resource Name (ARN) of an IAM role that provides permissions to your App Runner service.
	//
	// These are permissions that your code needs when it calls any AWS APIs.
	InstanceRoleArn *string `json:"instanceRoleArn" yaml:"instanceRoleArn"`
	// The amount of memory, in MB or GB, reserved for each instance of your App Runner service.
	//
	// Default: `2 GB`.
	Memory *string `json:"memory" yaml:"memory"`
}

// Describes a key-value pair, which is a string-to-string mapping.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import apprunner "github.com/aws/aws-cdk-go/awscdk/aws_apprunner"
//   keyValuePairProperty := &keyValuePairProperty{
//   	name: jsii.String("name"),
//   	value: jsii.String("value"),
//   }
//
type CfnService_KeyValuePairProperty struct {
	// The key name string to map to a value.
	Name *string `json:"name" yaml:"name"`
	// The value string to which the key name is mapped.
	Value *string `json:"value" yaml:"value"`
}

// Describes configuration settings related to network traffic of an AWS App Runner service.
//
// Consists of embedded objects for each configurable network feature.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import apprunner "github.com/aws/aws-cdk-go/awscdk/aws_apprunner"
//   networkConfigurationProperty := &networkConfigurationProperty{
//   	egressConfiguration: &egressConfigurationProperty{
//   		egressType: jsii.String("egressType"),
//
//   		// the properties below are optional
//   		vpcConnectorArn: jsii.String("vpcConnectorArn"),
//   	},
//   }
//
type CfnService_NetworkConfigurationProperty struct {
	// Network configuration settings for outbound message traffic.
	EgressConfiguration interface{} `json:"egressConfiguration" yaml:"egressConfiguration"`
}

// Describes the observability configuration of an AWS App Runner service.
//
// These are additional observability features, like tracing, that you choose to enable. They're configured in a separate resource that you associate with your service.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import apprunner "github.com/aws/aws-cdk-go/awscdk/aws_apprunner"
//   serviceObservabilityConfigurationProperty := &serviceObservabilityConfigurationProperty{
//   	observabilityEnabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	observabilityConfigurationArn: jsii.String("observabilityConfigurationArn"),
//   }
//
type CfnService_ServiceObservabilityConfigurationProperty struct {
	// When `true` , an observability configuration resource is associated with the service, and an `ObservabilityConfigurationArn` is specified.
	ObservabilityEnabled interface{} `json:"observabilityEnabled" yaml:"observabilityEnabled"`
	// The Amazon Resource Name (ARN) of the observability configuration that is associated with the service.
	//
	// Specified only when `ObservabilityEnabled` is `true` .
	//
	// Specify an ARN with a name and a revision number to associate that revision. For example: `arn:aws:apprunner:us-east-1:123456789012:observabilityconfiguration/xray-tracing/3`
	//
	// Specify just the name to associate the latest revision. For example: `arn:aws:apprunner:us-east-1:123456789012:observabilityconfiguration/xray-tracing`
	ObservabilityConfigurationArn *string `json:"observabilityConfigurationArn" yaml:"observabilityConfigurationArn"`
}

// Identifies a version of code that AWS App Runner refers to within a source code repository.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import apprunner "github.com/aws/aws-cdk-go/awscdk/aws_apprunner"
//   sourceCodeVersionProperty := &sourceCodeVersionProperty{
//   	type: jsii.String("type"),
//   	value: jsii.String("value"),
//   }
//
type CfnService_SourceCodeVersionProperty struct {
	// The type of version identifier.
	//
	// For a git-based repository, branches represent versions.
	Type *string `json:"type" yaml:"type"`
	// A source code version.
	//
	// For a git-based repository, a branch name maps to a specific version. App Runner uses the most recent commit to the branch.
	Value *string `json:"value" yaml:"value"`
}

// Describes the source deployed to an AWS App Runner service.
//
// It can be a code or an image repository.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import apprunner "github.com/aws/aws-cdk-go/awscdk/aws_apprunner"
//   sourceConfigurationProperty := &sourceConfigurationProperty{
//   	authenticationConfiguration: &authenticationConfigurationProperty{
//   		accessRoleArn: jsii.String("accessRoleArn"),
//   		connectionArn: jsii.String("connectionArn"),
//   	},
//   	autoDeploymentsEnabled: jsii.Boolean(false),
//   	codeRepository: &codeRepositoryProperty{
//   		repositoryUrl: jsii.String("repositoryUrl"),
//   		sourceCodeVersion: &sourceCodeVersionProperty{
//   			type: jsii.String("type"),
//   			value: jsii.String("value"),
//   		},
//
//   		// the properties below are optional
//   		codeConfiguration: &codeConfigurationProperty{
//   			configurationSource: jsii.String("configurationSource"),
//
//   			// the properties below are optional
//   			codeConfigurationValues: &codeConfigurationValuesProperty{
//   				runtime: jsii.String("runtime"),
//
//   				// the properties below are optional
//   				buildCommand: jsii.String("buildCommand"),
//   				port: jsii.String("port"),
//   				runtimeEnvironmentVariables: []interface{}{
//   					&keyValuePairProperty{
//   						name: jsii.String("name"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   				startCommand: jsii.String("startCommand"),
//   			},
//   		},
//   	},
//   	imageRepository: &imageRepositoryProperty{
//   		imageIdentifier: jsii.String("imageIdentifier"),
//   		imageRepositoryType: jsii.String("imageRepositoryType"),
//
//   		// the properties below are optional
//   		imageConfiguration: &imageConfigurationProperty{
//   			port: jsii.String("port"),
//   			runtimeEnvironmentVariables: []interface{}{
//   				&keyValuePairProperty{
//   					name: jsii.String("name"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   			startCommand: jsii.String("startCommand"),
//   		},
//   	},
//   }
//
type CfnService_SourceConfigurationProperty struct {
	// Describes the resources that are needed to authenticate access to some source repositories.
	AuthenticationConfiguration interface{} `json:"authenticationConfiguration" yaml:"authenticationConfiguration"`
	// If `true` , continuous integration from the source repository is enabled for the App Runner service.
	//
	// Each repository change (including any source code commit or new image version) starts a deployment.
	//
	// Default: App Runner sets to `false` for a source image that uses an ECR Public repository or an ECR repository that's in an AWS account other than the one that the service is in. App Runner sets to `true` in all other cases (which currently include a source code repository or a source image using a same-account ECR repository).
	AutoDeploymentsEnabled interface{} `json:"autoDeploymentsEnabled" yaml:"autoDeploymentsEnabled"`
	// The description of a source code repository.
	//
	// You must provide either this member or `ImageRepository` (but not both).
	CodeRepository interface{} `json:"codeRepository" yaml:"codeRepository"`
	// The description of a source image repository.
	//
	// You must provide either this member or `CodeRepository` (but not both).
	ImageRepository interface{} `json:"imageRepository" yaml:"imageRepository"`
}

// Properties for defining a `CfnService`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import apprunner "github.com/aws/aws-cdk-go/awscdk/aws_apprunner"
//   cfnServiceProps := &cfnServiceProps{
//   	sourceConfiguration: &sourceConfigurationProperty{
//   		authenticationConfiguration: &authenticationConfigurationProperty{
//   			accessRoleArn: jsii.String("accessRoleArn"),
//   			connectionArn: jsii.String("connectionArn"),
//   		},
//   		autoDeploymentsEnabled: jsii.Boolean(false),
//   		codeRepository: &codeRepositoryProperty{
//   			repositoryUrl: jsii.String("repositoryUrl"),
//   			sourceCodeVersion: &sourceCodeVersionProperty{
//   				type: jsii.String("type"),
//   				value: jsii.String("value"),
//   			},
//
//   			// the properties below are optional
//   			codeConfiguration: &codeConfigurationProperty{
//   				configurationSource: jsii.String("configurationSource"),
//
//   				// the properties below are optional
//   				codeConfigurationValues: &codeConfigurationValuesProperty{
//   					runtime: jsii.String("runtime"),
//
//   					// the properties below are optional
//   					buildCommand: jsii.String("buildCommand"),
//   					port: jsii.String("port"),
//   					runtimeEnvironmentVariables: []interface{}{
//   						&keyValuePairProperty{
//   							name: jsii.String("name"),
//   							value: jsii.String("value"),
//   						},
//   					},
//   					startCommand: jsii.String("startCommand"),
//   				},
//   			},
//   		},
//   		imageRepository: &imageRepositoryProperty{
//   			imageIdentifier: jsii.String("imageIdentifier"),
//   			imageRepositoryType: jsii.String("imageRepositoryType"),
//
//   			// the properties below are optional
//   			imageConfiguration: &imageConfigurationProperty{
//   				port: jsii.String("port"),
//   				runtimeEnvironmentVariables: []interface{}{
//   					&keyValuePairProperty{
//   						name: jsii.String("name"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   				startCommand: jsii.String("startCommand"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	autoScalingConfigurationArn: jsii.String("autoScalingConfigurationArn"),
//   	encryptionConfiguration: &encryptionConfigurationProperty{
//   		kmsKey: jsii.String("kmsKey"),
//   	},
//   	healthCheckConfiguration: &healthCheckConfigurationProperty{
//   		healthyThreshold: jsii.Number(123),
//   		interval: jsii.Number(123),
//   		path: jsii.String("path"),
//   		protocol: jsii.String("protocol"),
//   		timeout: jsii.Number(123),
//   		unhealthyThreshold: jsii.Number(123),
//   	},
//   	instanceConfiguration: &instanceConfigurationProperty{
//   		cpu: jsii.String("cpu"),
//   		instanceRoleArn: jsii.String("instanceRoleArn"),
//   		memory: jsii.String("memory"),
//   	},
//   	networkConfiguration: &networkConfigurationProperty{
//   		egressConfiguration: &egressConfigurationProperty{
//   			egressType: jsii.String("egressType"),
//
//   			// the properties below are optional
//   			vpcConnectorArn: jsii.String("vpcConnectorArn"),
//   		},
//   	},
//   	observabilityConfiguration: &serviceObservabilityConfigurationProperty{
//   		observabilityEnabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		observabilityConfigurationArn: jsii.String("observabilityConfigurationArn"),
//   	},
//   	serviceName: jsii.String("serviceName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnServiceProps struct {
	// The source to deploy to the App Runner service.
	//
	// It can be a code or an image repository.
	SourceConfiguration interface{} `json:"sourceConfiguration" yaml:"sourceConfiguration"`
	// The Amazon Resource Name (ARN) of an App Runner automatic scaling configuration resource that you want to associate with your service.
	//
	// If not provided, App Runner associates the latest revision of a default auto scaling configuration.
	//
	// Specify an ARN with a name and a revision number to associate that revision. For example: `arn:aws:apprunner:us-east-1:123456789012:autoscalingconfiguration/high-availability/3`
	//
	// Specify just the name to associate the latest revision. For example: `arn:aws:apprunner:us-east-1:123456789012:autoscalingconfiguration/high-availability`
	AutoScalingConfigurationArn *string `json:"autoScalingConfigurationArn" yaml:"autoScalingConfigurationArn"`
	// An optional custom encryption key that App Runner uses to encrypt the copy of your source repository that it maintains and your service logs.
	//
	// By default, App Runner uses an AWS managed key .
	EncryptionConfiguration interface{} `json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// The settings for the health check that AWS App Runner performs to monitor the health of the App Runner service.
	HealthCheckConfiguration interface{} `json:"healthCheckConfiguration" yaml:"healthCheckConfiguration"`
	// The runtime configuration of instances (scaling units) of your service.
	InstanceConfiguration interface{} `json:"instanceConfiguration" yaml:"instanceConfiguration"`
	// Configuration settings related to network traffic of the web application that the App Runner service runs.
	NetworkConfiguration interface{} `json:"networkConfiguration" yaml:"networkConfiguration"`
	// The observability configuration of your service.
	ObservabilityConfiguration interface{} `json:"observabilityConfiguration" yaml:"observabilityConfiguration"`
	// A name for the App Runner service.
	//
	// It must be unique across all the running App Runner services in your AWS account in the AWS Region .
	//
	// If you don't specify a name, AWS CloudFormation generates a name for your service.
	ServiceName *string `json:"serviceName" yaml:"serviceName"`
	// An optional list of metadata items that you can associate with the App Runner service resource.
	//
	// A tag is a key-value pair.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::AppRunner::VpcConnector`.
//
// Specify an AWS App Runner VPC connector by using the `AWS::AppRunner::VpcConnector` resource in an AWS CloudFormation template.
//
// The `AWS::AppRunner::VpcConnector` resource is an AWS App Runner resource type that specifies an App Runner VPC connector.
//
// App Runner requires this resource when you want to associate your App Runner service to a custom Amazon Virtual Private Cloud ( Amazon VPC ).
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import apprunner "github.com/aws/aws-cdk-go/awscdk/aws_apprunner"
//   cfnVpcConnector := apprunner.NewCfnVpcConnector(this, jsii.String("MyCfnVpcConnector"), &cfnVpcConnectorProps{
//   	subnets: []*string{
//   		jsii.String("subnets"),
//   	},
//
//   	// the properties below are optional
//   	securityGroups: []*string{
//   		jsii.String("securityGroups"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	vpcConnectorName: jsii.String("vpcConnectorName"),
//   })
//
type CfnVpcConnector interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of this VPC connector.
	AttrVpcConnectorArn() *string
	// The revision of this VPC connector.
	//
	// It's unique among all the active connectors ( `"Status": "ACTIVE"` ) that share the same `Name` .
	//
	// > At this time, App Runner supports only one revision per name.
	AttrVpcConnectorRevision() *float64
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
	// A list of IDs of security groups that App Runner should use for access to AWS resources under the specified subnets.
	//
	// If not specified, App Runner uses the default security group of the Amazon VPC. The default security group allows all outbound traffic.
	SecurityGroups() *[]*string
	SetSecurityGroups(val *[]*string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// A list of IDs of subnets that App Runner should use when it associates your service with a custom Amazon VPC.
	//
	// Specify IDs of subnets of a single Amazon VPC. App Runner determines the Amazon VPC from the subnets you specify.
	//
	// > App Runner currently only provides support for IPv4.
	Subnets() *[]*string
	SetSubnets(val *[]*string)
	// A list of metadata items that you can associate with your VPC connector resource.
	//
	// A tag is a key-value pair.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// A name for the VPC connector.
	//
	// If you don't specify a name, AWS CloudFormation generates a name for your VPC connector.
	VpcConnectorName() *string
	SetVpcConnectorName(val *string)
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

// The jsii proxy struct for CfnVpcConnector
type jsiiProxy_CfnVpcConnector struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnVpcConnector) AttrVpcConnectorArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrVpcConnectorArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcConnector) AttrVpcConnectorRevision() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrVpcConnectorRevision",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcConnector) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcConnector) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcConnector) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcConnector) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcConnector) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcConnector) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcConnector) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcConnector) SecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcConnector) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcConnector) Subnets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcConnector) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcConnector) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcConnector) VpcConnectorName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcConnectorName",
		&returns,
	)
	return returns
}


// Create a new `AWS::AppRunner::VpcConnector`.
func NewCfnVpcConnector(scope awscdk.Construct, id *string, props *CfnVpcConnectorProps) CfnVpcConnector {
	_init_.Initialize()

	j := jsiiProxy_CfnVpcConnector{}

	_jsii_.Create(
		"monocdk.aws_apprunner.CfnVpcConnector",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AppRunner::VpcConnector`.
func NewCfnVpcConnector_Override(c CfnVpcConnector, scope awscdk.Construct, id *string, props *CfnVpcConnectorProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_apprunner.CfnVpcConnector",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnVpcConnector) SetSecurityGroups(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroups",
		val,
	)
}

func (j *jsiiProxy_CfnVpcConnector) SetSubnets(val *[]*string) {
	_jsii_.Set(
		j,
		"subnets",
		val,
	)
}

func (j *jsiiProxy_CfnVpcConnector) SetVpcConnectorName(val *string) {
	_jsii_.Set(
		j,
		"vpcConnectorName",
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
func CfnVpcConnector_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_apprunner.CfnVpcConnector",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnVpcConnector_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_apprunner.CfnVpcConnector",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnVpcConnector_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_apprunner.CfnVpcConnector",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnVpcConnector_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_apprunner.CfnVpcConnector",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnVpcConnector) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnVpcConnector) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnVpcConnector) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnVpcConnector) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnVpcConnector) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnVpcConnector) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnVpcConnector) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnVpcConnector) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVpcConnector) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVpcConnector) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnVpcConnector) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnVpcConnector) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnVpcConnector) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVpcConnector) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnVpcConnector) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnVpcConnector) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVpcConnector) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVpcConnector) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnVpcConnector) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVpcConnector) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVpcConnector) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnVpcConnector`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import apprunner "github.com/aws/aws-cdk-go/awscdk/aws_apprunner"
//   cfnVpcConnectorProps := &cfnVpcConnectorProps{
//   	subnets: []*string{
//   		jsii.String("subnets"),
//   	},
//
//   	// the properties below are optional
//   	securityGroups: []*string{
//   		jsii.String("securityGroups"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	vpcConnectorName: jsii.String("vpcConnectorName"),
//   }
//
type CfnVpcConnectorProps struct {
	// A list of IDs of subnets that App Runner should use when it associates your service with a custom Amazon VPC.
	//
	// Specify IDs of subnets of a single Amazon VPC. App Runner determines the Amazon VPC from the subnets you specify.
	//
	// > App Runner currently only provides support for IPv4.
	Subnets *[]*string `json:"subnets" yaml:"subnets"`
	// A list of IDs of security groups that App Runner should use for access to AWS resources under the specified subnets.
	//
	// If not specified, App Runner uses the default security group of the Amazon VPC. The default security group allows all outbound traffic.
	SecurityGroups *[]*string `json:"securityGroups" yaml:"securityGroups"`
	// A list of metadata items that you can associate with your VPC connector resource.
	//
	// A tag is a key-value pair.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
	// A name for the VPC connector.
	//
	// If you don't specify a name, AWS CloudFormation generates a name for your VPC connector.
	VpcConnectorName *string `json:"vpcConnectorName" yaml:"vpcConnectorName"`
}

// Describes the configuration that AWS App Runner uses to build and run an App Runner service from a source code repository.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import apprunner "github.com/aws/aws-cdk-go/awscdk/aws_apprunner"
//
//   var runtime runtime
//   codeConfiguration := &codeConfiguration{
//   	configurationSource: apprunner.configurationSourceType_REPOSITORY,
//
//   	// the properties below are optional
//   	configurationValues: &codeConfigurationValues{
//   		runtime: runtime,
//
//   		// the properties below are optional
//   		buildCommand: jsii.String("buildCommand"),
//   		environment: map[string]*string{
//   			"environmentKey": jsii.String("environment"),
//   		},
//   		port: jsii.String("port"),
//   		startCommand: jsii.String("startCommand"),
//   	},
//   }
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-codeconfiguration.html
//
// Experimental.
type CodeConfiguration struct {
	// The source of the App Runner configuration.
	// Experimental.
	ConfigurationSource ConfigurationSourceType `json:"configurationSource" yaml:"configurationSource"`
	// The basic configuration for building and running the App Runner service.
	//
	// Use it to quickly launch an App Runner service without providing a apprunner.yaml file in the
	// source code repository (or ignoring the file if it exists).
	// Experimental.
	ConfigurationValues *CodeConfigurationValues `json:"configurationValues" yaml:"configurationValues"`
}

// Describes the basic configuration needed for building and running an AWS App Runner service.
//
// This type doesn't support the full set of possible configuration options. Fur full configuration capabilities,
// use a `apprunner.yaml` file in the source code repository.
//
// Example:
//   apprunner.NewService(this, jsii.String("Service"), &serviceProps{
//   	source: apprunner.source.fromGitHub(&githubRepositoryProps{
//   		repositoryUrl: jsii.String("https://github.com/aws-containers/hello-app-runner"),
//   		branch: jsii.String("main"),
//   		configurationSource: apprunner.configurationSourceType_API,
//   		codeConfigurationValues: &codeConfigurationValues{
//   			runtime: apprunner.runtime_PYTHON_3(),
//   			port: jsii.String("8000"),
//   			startCommand: jsii.String("python app.py"),
//   			buildCommand: jsii.String("yum install -y pycairo && pip install -r requirements.txt"),
//   		},
//   		connection: apprunner.gitHubConnection.fromConnectionArn(jsii.String("CONNECTION_ARN")),
//   	}),
//   })
//
// Experimental.
type CodeConfigurationValues struct {
	// A runtime environment type for building and running an App Runner service.
	//
	// It represents
	// a programming language runtime.
	// Experimental.
	Runtime Runtime `json:"runtime" yaml:"runtime"`
	// The command App Runner runs to build your application.
	// Experimental.
	BuildCommand *string `json:"buildCommand" yaml:"buildCommand"`
	// The environment variables that are available to your running App Runner service.
	// Experimental.
	Environment *map[string]*string `json:"environment" yaml:"environment"`
	// The port that your application listens to in the container.
	// Experimental.
	Port *string `json:"port" yaml:"port"`
	// The command App Runner runs to start your application.
	// Experimental.
	StartCommand *string `json:"startCommand" yaml:"startCommand"`
}

// Properties of the CodeRepository.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import apprunner "github.com/aws/aws-cdk-go/awscdk/aws_apprunner"
//
//   var gitHubConnection gitHubConnection
//   var runtime runtime
//   codeRepositoryProps := &codeRepositoryProps{
//   	codeConfiguration: &codeConfiguration{
//   		configurationSource: apprunner.configurationSourceType_REPOSITORY,
//
//   		// the properties below are optional
//   		configurationValues: &codeConfigurationValues{
//   			runtime: runtime,
//
//   			// the properties below are optional
//   			buildCommand: jsii.String("buildCommand"),
//   			environment: map[string]*string{
//   				"environmentKey": jsii.String("environment"),
//   			},
//   			port: jsii.String("port"),
//   			startCommand: jsii.String("startCommand"),
//   		},
//   	},
//   	connection: gitHubConnection,
//   	repositoryUrl: jsii.String("repositoryUrl"),
//   	sourceCodeVersion: &sourceCodeVersion{
//   		type: jsii.String("type"),
//   		value: jsii.String("value"),
//   	},
//   }
//
// Experimental.
type CodeRepositoryProps struct {
	// Configuration for building and running the service from a source code repository.
	// Experimental.
	CodeConfiguration *CodeConfiguration `json:"codeConfiguration" yaml:"codeConfiguration"`
	// The App Runner connection for GitHub.
	// Experimental.
	Connection GitHubConnection `json:"connection" yaml:"connection"`
	// The location of the repository that contains the source code.
	// Experimental.
	RepositoryUrl *string `json:"repositoryUrl" yaml:"repositoryUrl"`
	// The version that should be used within the source code repository.
	// Experimental.
	SourceCodeVersion *SourceCodeVersion `json:"sourceCodeVersion" yaml:"sourceCodeVersion"`
}

// The source of the App Runner configuration.
//
// Example:
//   apprunner.NewService(this, jsii.String("Service"), &serviceProps{
//   	source: apprunner.source.fromGitHub(&githubRepositoryProps{
//   		repositoryUrl: jsii.String("https://github.com/aws-containers/hello-app-runner"),
//   		branch: jsii.String("main"),
//   		configurationSource: apprunner.configurationSourceType_REPOSITORY,
//   		connection: apprunner.gitHubConnection.fromConnectionArn(jsii.String("CONNECTION_ARN")),
//   	}),
//   })
//
// Experimental.
type ConfigurationSourceType string

const (
	// App Runner reads configuration values from `the apprunner.yaml` file in the source code repository and ignores `configurationValues`.
	// Experimental.
	ConfigurationSourceType_REPOSITORY ConfigurationSourceType = "REPOSITORY"
	// App Runner uses configuration values provided in `configurationValues` and ignores the `apprunner.yaml` file in the source code repository.
	// Experimental.
	ConfigurationSourceType_API ConfigurationSourceType = "API"
)

// The number of CPU units reserved for each instance of your App Runner service.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import apprunner "github.com/aws/aws-cdk-go/awscdk/aws_apprunner"
//   cpu := apprunner.cpu.of(jsii.String("unit"))
//
// Experimental.
type Cpu interface {
	// The unit of CPU.
	// Experimental.
	Unit() *string
}

// The jsii proxy struct for Cpu
type jsiiProxy_Cpu struct {
	_ byte // padding
}

func (j *jsiiProxy_Cpu) Unit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unit",
		&returns,
	)
	return returns
}


// Custom CPU unit.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-instanceconfiguration.html#cfn-apprunner-service-instanceconfiguration-cpu
//
// Experimental.
func Cpu_Of(unit *string) Cpu {
	_init_.Initialize()

	var returns Cpu

	_jsii_.StaticInvoke(
		"monocdk.aws_apprunner.Cpu",
		"of",
		[]interface{}{unit},
		&returns,
	)

	return returns
}

func Cpu_ONE_VCPU() Cpu {
	_init_.Initialize()
	var returns Cpu
	_jsii_.StaticGet(
		"monocdk.aws_apprunner.Cpu",
		"ONE_VCPU",
		&returns,
	)
	return returns
}

func Cpu_TWO_VCPU() Cpu {
	_init_.Initialize()
	var returns Cpu
	_jsii_.StaticGet(
		"monocdk.aws_apprunner.Cpu",
		"TWO_VCPU",
		&returns,
	)
	return returns
}

// Properties of the image repository for `Source.fromEcr()`.
//
// Example:
//   import ecr "github.com/aws/aws-cdk-go/awscdk"
//
//   apprunner.NewService(this, jsii.String("Service"), &serviceProps{
//   	source: apprunner.source.fromEcr(&ecrProps{
//   		imageConfiguration: &imageConfiguration{
//   			port: jsii.Number(80),
//   		},
//   		repository: ecr.repository.fromRepositoryName(this, jsii.String("NginxRepository"), jsii.String("nginx")),
//   		tagOrDigest: jsii.String("latest"),
//   	}),
//   })
//
// Experimental.
type EcrProps struct {
	// Represents the ECR repository.
	// Experimental.
	Repository awsecr.IRepository `json:"repository" yaml:"repository"`
	// The image configuration for the image from ECR.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-imageconfiguration.html#cfn-apprunner-service-imageconfiguration-port
	//
	// Experimental.
	ImageConfiguration *ImageConfiguration `json:"imageConfiguration" yaml:"imageConfiguration"`
	// Image tag.
	// Deprecated: use `tagOrDigest`.
	Tag *string `json:"tag" yaml:"tag"`
	// Image tag or digest (digests must start with `sha256:`).
	// Experimental.
	TagOrDigest *string `json:"tagOrDigest" yaml:"tagOrDigest"`
}

// Properties of the image repository for `Source.fromEcrPublic()`.
//
// Example:
//   apprunner.NewService(this, jsii.String("Service"), &serviceProps{
//   	source: apprunner.source.fromEcrPublic(&ecrPublicProps{
//   		imageConfiguration: &imageConfiguration{
//   			port: jsii.Number(8000),
//   		},
//   		imageIdentifier: jsii.String("public.ecr.aws/aws-containers/hello-app-runner:latest"),
//   	}),
//   })
//
// Experimental.
type EcrPublicProps struct {
	// The ECR Public image URI.
	// Experimental.
	ImageIdentifier *string `json:"imageIdentifier" yaml:"imageIdentifier"`
	// The image configuration for the image from ECR Public.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-imageconfiguration.html#cfn-apprunner-service-imageconfiguration-port
	//
	// Experimental.
	ImageConfiguration *ImageConfiguration `json:"imageConfiguration" yaml:"imageConfiguration"`
}

// Represents the service source from ECR Public.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import apprunner "github.com/aws/aws-cdk-go/awscdk/aws_apprunner"
//   ecrPublicSource := apprunner.NewEcrPublicSource(&ecrPublicProps{
//   	imageIdentifier: jsii.String("imageIdentifier"),
//
//   	// the properties below are optional
//   	imageConfiguration: &imageConfiguration{
//   		environment: map[string]*string{
//   			"environmentKey": jsii.String("environment"),
//   		},
//   		port: jsii.Number(123),
//   		startCommand: jsii.String("startCommand"),
//   	},
//   })
//
// Experimental.
type EcrPublicSource interface {
	Source
	// Called when the Job is initialized to allow this object to bind.
	// Experimental.
	Bind(_scope constructs.Construct) *SourceConfig
}

// The jsii proxy struct for EcrPublicSource
type jsiiProxy_EcrPublicSource struct {
	jsiiProxy_Source
}

// Experimental.
func NewEcrPublicSource(props *EcrPublicProps) EcrPublicSource {
	_init_.Initialize()

	j := jsiiProxy_EcrPublicSource{}

	_jsii_.Create(
		"monocdk.aws_apprunner.EcrPublicSource",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewEcrPublicSource_Override(e EcrPublicSource, props *EcrPublicProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_apprunner.EcrPublicSource",
		[]interface{}{props},
		e,
	)
}

// Source from local assets.
// Experimental.
func EcrPublicSource_FromAsset(props *AssetProps) AssetSource {
	_init_.Initialize()

	var returns AssetSource

	_jsii_.StaticInvoke(
		"monocdk.aws_apprunner.EcrPublicSource",
		"fromAsset",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Source from the ECR repository.
// Experimental.
func EcrPublicSource_FromEcr(props *EcrProps) EcrSource {
	_init_.Initialize()

	var returns EcrSource

	_jsii_.StaticInvoke(
		"monocdk.aws_apprunner.EcrPublicSource",
		"fromEcr",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Source from the ECR Public repository.
// Experimental.
func EcrPublicSource_FromEcrPublic(props *EcrPublicProps) EcrPublicSource {
	_init_.Initialize()

	var returns EcrPublicSource

	_jsii_.StaticInvoke(
		"monocdk.aws_apprunner.EcrPublicSource",
		"fromEcrPublic",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Source from the GitHub repository.
// Experimental.
func EcrPublicSource_FromGitHub(props *GithubRepositoryProps) GithubSource {
	_init_.Initialize()

	var returns GithubSource

	_jsii_.StaticInvoke(
		"monocdk.aws_apprunner.EcrPublicSource",
		"fromGitHub",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrPublicSource) Bind(_scope constructs.Construct) *SourceConfig {
	var returns *SourceConfig

	_jsii_.Invoke(
		e,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

// Represents the service source from ECR.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import apprunner "github.com/aws/aws-cdk-go/awscdk/aws_apprunner"import awscdk "github.com/aws/aws-cdk-go/awscdk"import ecr "github.com/aws/aws-cdk-go/awscdk/aws_ecr"
//
//   var repository repository
//   ecrSource := apprunner.NewEcrSource(&ecrProps{
//   	repository: repository,
//
//   	// the properties below are optional
//   	imageConfiguration: &imageConfiguration{
//   		environment: map[string]*string{
//   			"environmentKey": jsii.String("environment"),
//   		},
//   		port: jsii.Number(123),
//   		startCommand: jsii.String("startCommand"),
//   	},
//   	tag: jsii.String("tag"),
//   	tagOrDigest: jsii.String("tagOrDigest"),
//   })
//
// Experimental.
type EcrSource interface {
	Source
	// Called when the Job is initialized to allow this object to bind.
	// Experimental.
	Bind(_scope constructs.Construct) *SourceConfig
}

// The jsii proxy struct for EcrSource
type jsiiProxy_EcrSource struct {
	jsiiProxy_Source
}

// Experimental.
func NewEcrSource(props *EcrProps) EcrSource {
	_init_.Initialize()

	j := jsiiProxy_EcrSource{}

	_jsii_.Create(
		"monocdk.aws_apprunner.EcrSource",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewEcrSource_Override(e EcrSource, props *EcrProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_apprunner.EcrSource",
		[]interface{}{props},
		e,
	)
}

// Source from local assets.
// Experimental.
func EcrSource_FromAsset(props *AssetProps) AssetSource {
	_init_.Initialize()

	var returns AssetSource

	_jsii_.StaticInvoke(
		"monocdk.aws_apprunner.EcrSource",
		"fromAsset",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Source from the ECR repository.
// Experimental.
func EcrSource_FromEcr(props *EcrProps) EcrSource {
	_init_.Initialize()

	var returns EcrSource

	_jsii_.StaticInvoke(
		"monocdk.aws_apprunner.EcrSource",
		"fromEcr",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Source from the ECR Public repository.
// Experimental.
func EcrSource_FromEcrPublic(props *EcrPublicProps) EcrPublicSource {
	_init_.Initialize()

	var returns EcrPublicSource

	_jsii_.StaticInvoke(
		"monocdk.aws_apprunner.EcrSource",
		"fromEcrPublic",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Source from the GitHub repository.
// Experimental.
func EcrSource_FromGitHub(props *GithubRepositoryProps) GithubSource {
	_init_.Initialize()

	var returns GithubSource

	_jsii_.StaticInvoke(
		"monocdk.aws_apprunner.EcrSource",
		"fromGitHub",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrSource) Bind(_scope constructs.Construct) *SourceConfig {
	var returns *SourceConfig

	_jsii_.Invoke(
		e,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

// Represents the App Runner connection that enables the App Runner service to connect to a source repository.
//
// It's required for GitHub code repositories.
//
// Example:
//   apprunner.NewService(this, jsii.String("Service"), &serviceProps{
//   	source: apprunner.source.fromGitHub(&githubRepositoryProps{
//   		repositoryUrl: jsii.String("https://github.com/aws-containers/hello-app-runner"),
//   		branch: jsii.String("main"),
//   		configurationSource: apprunner.configurationSourceType_REPOSITORY,
//   		connection: apprunner.gitHubConnection.fromConnectionArn(jsii.String("CONNECTION_ARN")),
//   	}),
//   })
//
// Experimental.
type GitHubConnection interface {
	// The ARN of the Connection for App Runner service to connect to the repository.
	// Experimental.
	ConnectionArn() *string
}

// The jsii proxy struct for GitHubConnection
type jsiiProxy_GitHubConnection struct {
	_ byte // padding
}

func (j *jsiiProxy_GitHubConnection) ConnectionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionArn",
		&returns,
	)
	return returns
}


// Experimental.
func NewGitHubConnection(arn *string) GitHubConnection {
	_init_.Initialize()

	j := jsiiProxy_GitHubConnection{}

	_jsii_.Create(
		"monocdk.aws_apprunner.GitHubConnection",
		[]interface{}{arn},
		&j,
	)

	return &j
}

// Experimental.
func NewGitHubConnection_Override(g GitHubConnection, arn *string) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_apprunner.GitHubConnection",
		[]interface{}{arn},
		g,
	)
}

// Using existing App Runner connection by specifying the connection ARN.
//
// Returns: Connection.
// Experimental.
func GitHubConnection_FromConnectionArn(arn *string) GitHubConnection {
	_init_.Initialize()

	var returns GitHubConnection

	_jsii_.StaticInvoke(
		"monocdk.aws_apprunner.GitHubConnection",
		"fromConnectionArn",
		[]interface{}{arn},
		&returns,
	)

	return returns
}

// Properties of the Github repository for `Source.fromGitHub()`.
//
// Example:
//   apprunner.NewService(this, jsii.String("Service"), &serviceProps{
//   	source: apprunner.source.fromGitHub(&githubRepositoryProps{
//   		repositoryUrl: jsii.String("https://github.com/aws-containers/hello-app-runner"),
//   		branch: jsii.String("main"),
//   		configurationSource: apprunner.configurationSourceType_REPOSITORY,
//   		connection: apprunner.gitHubConnection.fromConnectionArn(jsii.String("CONNECTION_ARN")),
//   	}),
//   })
//
// Experimental.
type GithubRepositoryProps struct {
	// The source of the App Runner configuration.
	// Experimental.
	ConfigurationSource ConfigurationSourceType `json:"configurationSource" yaml:"configurationSource"`
	// ARN of the connection to Github.
	//
	// Only required for Github source.
	// Experimental.
	Connection GitHubConnection `json:"connection" yaml:"connection"`
	// The location of the repository that contains the source code.
	// Experimental.
	RepositoryUrl *string `json:"repositoryUrl" yaml:"repositoryUrl"`
	// The branch name that represents a specific version for the repository.
	// Experimental.
	Branch *string `json:"branch" yaml:"branch"`
	// The code configuration values.
	//
	// Will be ignored if configurationSource is `REPOSITORY`.
	// Experimental.
	CodeConfigurationValues *CodeConfigurationValues `json:"codeConfigurationValues" yaml:"codeConfigurationValues"`
}

// Represents the service source from a Github repository.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import apprunner "github.com/aws/aws-cdk-go/awscdk/aws_apprunner"
//
//   var gitHubConnection gitHubConnection
//   var runtime runtime
//   githubSource := apprunner.NewGithubSource(&githubRepositoryProps{
//   	configurationSource: apprunner.configurationSourceType_REPOSITORY,
//   	connection: gitHubConnection,
//   	repositoryUrl: jsii.String("repositoryUrl"),
//
//   	// the properties below are optional
//   	branch: jsii.String("branch"),
//   	codeConfigurationValues: &codeConfigurationValues{
//   		runtime: runtime,
//
//   		// the properties below are optional
//   		buildCommand: jsii.String("buildCommand"),
//   		environment: map[string]*string{
//   			"environmentKey": jsii.String("environment"),
//   		},
//   		port: jsii.String("port"),
//   		startCommand: jsii.String("startCommand"),
//   	},
//   })
//
// Experimental.
type GithubSource interface {
	Source
	// Called when the Job is initialized to allow this object to bind.
	// Experimental.
	Bind(_scope constructs.Construct) *SourceConfig
}

// The jsii proxy struct for GithubSource
type jsiiProxy_GithubSource struct {
	jsiiProxy_Source
}

// Experimental.
func NewGithubSource(props *GithubRepositoryProps) GithubSource {
	_init_.Initialize()

	j := jsiiProxy_GithubSource{}

	_jsii_.Create(
		"monocdk.aws_apprunner.GithubSource",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewGithubSource_Override(g GithubSource, props *GithubRepositoryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_apprunner.GithubSource",
		[]interface{}{props},
		g,
	)
}

// Source from local assets.
// Experimental.
func GithubSource_FromAsset(props *AssetProps) AssetSource {
	_init_.Initialize()

	var returns AssetSource

	_jsii_.StaticInvoke(
		"monocdk.aws_apprunner.GithubSource",
		"fromAsset",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Source from the ECR repository.
// Experimental.
func GithubSource_FromEcr(props *EcrProps) EcrSource {
	_init_.Initialize()

	var returns EcrSource

	_jsii_.StaticInvoke(
		"monocdk.aws_apprunner.GithubSource",
		"fromEcr",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Source from the ECR Public repository.
// Experimental.
func GithubSource_FromEcrPublic(props *EcrPublicProps) EcrPublicSource {
	_init_.Initialize()

	var returns EcrPublicSource

	_jsii_.StaticInvoke(
		"monocdk.aws_apprunner.GithubSource",
		"fromEcrPublic",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Source from the GitHub repository.
// Experimental.
func GithubSource_FromGitHub(props *GithubRepositoryProps) GithubSource {
	_init_.Initialize()

	var returns GithubSource

	_jsii_.StaticInvoke(
		"monocdk.aws_apprunner.GithubSource",
		"fromGitHub",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GithubSource) Bind(_scope constructs.Construct) *SourceConfig {
	var returns *SourceConfig

	_jsii_.Invoke(
		g,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

// Represents the App Runner Service.
// Experimental.
type IService interface {
	awscdk.IResource
	// The ARN of the service.
	// Experimental.
	ServiceArn() *string
	// The Name of the service.
	// Experimental.
	ServiceName() *string
}

// The jsii proxy for IService
type jsiiProxy_IService struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IService) ServiceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IService) ServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceName",
		&returns,
	)
	return returns
}

// Describes the configuration that AWS App Runner uses to run an App Runner service using an image pulled from a source image repository.
//
// Example:
//   import ecr "github.com/aws/aws-cdk-go/awscdk"
//
//   apprunner.NewService(this, jsii.String("Service"), &serviceProps{
//   	source: apprunner.source.fromEcr(&ecrProps{
//   		imageConfiguration: &imageConfiguration{
//   			port: jsii.Number(80),
//   		},
//   		repository: ecr.repository.fromRepositoryName(this, jsii.String("NginxRepository"), jsii.String("nginx")),
//   		tagOrDigest: jsii.String("latest"),
//   	}),
//   })
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-imageconfiguration.html
//
// Experimental.
type ImageConfiguration struct {
	// Environment variables that are available to your running App Runner service.
	// Experimental.
	Environment *map[string]*string `json:"environment" yaml:"environment"`
	// The port that your application listens to in the container.
	// Experimental.
	Port *float64 `json:"port" yaml:"port"`
	// An optional command that App Runner runs to start the application in the source image.
	//
	// If specified, this command overrides the Docker image’s default start command.
	// Experimental.
	StartCommand *string `json:"startCommand" yaml:"startCommand"`
}

// Describes a source image repository.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import apprunner "github.com/aws/aws-cdk-go/awscdk/aws_apprunner"
//   imageRepository := &imageRepository{
//   	imageIdentifier: jsii.String("imageIdentifier"),
//   	imageRepositoryType: apprunner.imageRepositoryType_ECR_PUBLIC,
//
//   	// the properties below are optional
//   	imageConfiguration: &imageConfiguration{
//   		environment: map[string]*string{
//   			"environmentKey": jsii.String("environment"),
//   		},
//   		port: jsii.Number(123),
//   		startCommand: jsii.String("startCommand"),
//   	},
//   }
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-imagerepository.html
//
// Experimental.
type ImageRepository struct {
	// The identifier of the image.
	//
	// For `ECR_PUBLIC` imageRepositoryType, the identifier domain should
	// always be `public.ecr.aws`. For `ECR`, the pattern should be
	// `([0-9]{12}.dkr.ecr.[a-z\-]+-[0-9]{1}.amazonaws.com\/.*)`.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-imagerepository.html for more details.
	//
	// Experimental.
	ImageIdentifier *string `json:"imageIdentifier" yaml:"imageIdentifier"`
	// The type of the image repository.
	//
	// This reflects the repository provider and whether
	// the repository is private or public.
	// Experimental.
	ImageRepositoryType ImageRepositoryType `json:"imageRepositoryType" yaml:"imageRepositoryType"`
	// Configuration for running the identified image.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-imageconfiguration.html#cfn-apprunner-service-imageconfiguration-port
	//
	// Experimental.
	ImageConfiguration *ImageConfiguration `json:"imageConfiguration" yaml:"imageConfiguration"`
}

// The image repository types.
// Experimental.
type ImageRepositoryType string

const (
	// Amazon ECR Public.
	// Experimental.
	ImageRepositoryType_ECR_PUBLIC ImageRepositoryType = "ECR_PUBLIC"
	// Amazon ECR.
	// Experimental.
	ImageRepositoryType_ECR ImageRepositoryType = "ECR"
)

// The amount of memory reserved for each instance of your App Runner service.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import apprunner "github.com/aws/aws-cdk-go/awscdk/aws_apprunner"
//   memory := apprunner.memory_FOUR_GB()
//
// Experimental.
type Memory interface {
	// The unit of memory.
	// Experimental.
	Unit() *string
}

// The jsii proxy struct for Memory
type jsiiProxy_Memory struct {
	_ byte // padding
}

func (j *jsiiProxy_Memory) Unit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unit",
		&returns,
	)
	return returns
}


// Custom Memory unit.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-instanceconfiguration.html#cfn-apprunner-service-instanceconfiguration-memory
//
// Experimental.
func Memory_Of(unit *string) Memory {
	_init_.Initialize()

	var returns Memory

	_jsii_.StaticInvoke(
		"monocdk.aws_apprunner.Memory",
		"of",
		[]interface{}{unit},
		&returns,
	)

	return returns
}

func Memory_FOUR_GB() Memory {
	_init_.Initialize()
	var returns Memory
	_jsii_.StaticGet(
		"monocdk.aws_apprunner.Memory",
		"FOUR_GB",
		&returns,
	)
	return returns
}

func Memory_THREE_GB() Memory {
	_init_.Initialize()
	var returns Memory
	_jsii_.StaticGet(
		"monocdk.aws_apprunner.Memory",
		"THREE_GB",
		&returns,
	)
	return returns
}

func Memory_TWO_GB() Memory {
	_init_.Initialize()
	var returns Memory
	_jsii_.StaticGet(
		"monocdk.aws_apprunner.Memory",
		"TWO_GB",
		&returns,
	)
	return returns
}

// The code runtimes.
//
// Example:
//   apprunner.NewService(this, jsii.String("Service"), &serviceProps{
//   	source: apprunner.source.fromGitHub(&githubRepositoryProps{
//   		repositoryUrl: jsii.String("https://github.com/aws-containers/hello-app-runner"),
//   		branch: jsii.String("main"),
//   		configurationSource: apprunner.configurationSourceType_API,
//   		codeConfigurationValues: &codeConfigurationValues{
//   			runtime: apprunner.runtime_PYTHON_3(),
//   			port: jsii.String("8000"),
//   			startCommand: jsii.String("python app.py"),
//   			buildCommand: jsii.String("yum install -y pycairo && pip install -r requirements.txt"),
//   		},
//   		connection: apprunner.gitHubConnection.fromConnectionArn(jsii.String("CONNECTION_ARN")),
//   	}),
//   })
//
// Experimental.
type Runtime interface {
	// The runtime name.
	// Experimental.
	Name() *string
}

// The jsii proxy struct for Runtime
type jsiiProxy_Runtime struct {
	_ byte // padding
}

func (j *jsiiProxy_Runtime) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Other runtimes.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-codeconfigurationvalues.html#cfn-apprunner-service-codeconfigurationvalues-runtime for all available runtimes.
//
// Experimental.
func Runtime_Of(name *string) Runtime {
	_init_.Initialize()

	var returns Runtime

	_jsii_.StaticInvoke(
		"monocdk.aws_apprunner.Runtime",
		"of",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func Runtime_NODEJS_12() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"monocdk.aws_apprunner.Runtime",
		"NODEJS_12",
		&returns,
	)
	return returns
}

func Runtime_PYTHON_3() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"monocdk.aws_apprunner.Runtime",
		"PYTHON_3",
		&returns,
	)
	return returns
}

// The App Runner Service.
//
// Example:
//   apprunner.NewService(this, jsii.String("Service"), &serviceProps{
//   	source: apprunner.source.fromGitHub(&githubRepositoryProps{
//   		repositoryUrl: jsii.String("https://github.com/aws-containers/hello-app-runner"),
//   		branch: jsii.String("main"),
//   		configurationSource: apprunner.configurationSourceType_REPOSITORY,
//   		connection: apprunner.gitHubConnection.fromConnectionArn(jsii.String("CONNECTION_ARN")),
//   	}),
//   })
//
// Experimental.
type Service interface {
	awscdk.Resource
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	// Experimental.
	Env() *awscdk.ResourceEnvironment
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// The ARN of the Service.
	// Experimental.
	ServiceArn() *string
	// The ID of the Service.
	// Experimental.
	ServiceId() *string
	// The name of the service.
	// Experimental.
	ServiceName() *string
	// The status of the Service.
	// Experimental.
	ServiceStatus() *string
	// The URL of the Service.
	// Experimental.
	ServiceUrl() *string
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Experimental.
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	// Experimental.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	// Experimental.
	GetResourceNameAttribute(nameAttr *string) *string
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
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
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
}

// The jsii proxy struct for Service
type jsiiProxy_Service struct {
	internal.Type__awscdkResource
}

func (j *jsiiProxy_Service) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) ServiceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) ServiceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) ServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) ServiceStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) ServiceUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewService(scope constructs.Construct, id *string, props *ServiceProps) Service {
	_init_.Initialize()

	j := jsiiProxy_Service{}

	_jsii_.Create(
		"monocdk.aws_apprunner.Service",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewService_Override(s Service, scope constructs.Construct, id *string, props *ServiceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_apprunner.Service",
		[]interface{}{scope, id, props},
		s,
	)
}

// Import from service attributes.
// Experimental.
func Service_FromServiceAttributes(scope constructs.Construct, id *string, attrs *ServiceAttributes) IService {
	_init_.Initialize()

	var returns IService

	_jsii_.StaticInvoke(
		"monocdk.aws_apprunner.Service",
		"fromServiceAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Import from service name.
// Experimental.
func Service_FromServiceName(scope constructs.Construct, id *string, serviceName *string) IService {
	_init_.Initialize()

	var returns IService

	_jsii_.StaticInvoke(
		"monocdk.aws_apprunner.Service",
		"fromServiceName",
		[]interface{}{scope, id, serviceName},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func Service_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_apprunner.Service",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func Service_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_apprunner.Service",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Service) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		s,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (s *jsiiProxy_Service) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Service) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Service) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Service) OnPrepare() {
	_jsii_.InvokeVoid(
		s,
		"onPrepare",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Service) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (s *jsiiProxy_Service) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Service) Prepare() {
	_jsii_.InvokeVoid(
		s,
		"prepare",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Service) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		[]interface{}{session},
	)
}

func (s *jsiiProxy_Service) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Service) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Attributes for the App Runner Service.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import apprunner "github.com/aws/aws-cdk-go/awscdk/aws_apprunner"
//   serviceAttributes := &serviceAttributes{
//   	serviceArn: jsii.String("serviceArn"),
//   	serviceName: jsii.String("serviceName"),
//   	serviceStatus: jsii.String("serviceStatus"),
//   	serviceUrl: jsii.String("serviceUrl"),
//   }
//
// Experimental.
type ServiceAttributes struct {
	// The ARN of the service.
	// Experimental.
	ServiceArn *string `json:"serviceArn" yaml:"serviceArn"`
	// The name of the service.
	// Experimental.
	ServiceName *string `json:"serviceName" yaml:"serviceName"`
	// The status of the service.
	// Experimental.
	ServiceStatus *string `json:"serviceStatus" yaml:"serviceStatus"`
	// The URL of the service.
	// Experimental.
	ServiceUrl *string `json:"serviceUrl" yaml:"serviceUrl"`
}

// Properties of the AppRunner Service.
//
// Example:
//   apprunner.NewService(this, jsii.String("Service"), &serviceProps{
//   	source: apprunner.source.fromGitHub(&githubRepositoryProps{
//   		repositoryUrl: jsii.String("https://github.com/aws-containers/hello-app-runner"),
//   		branch: jsii.String("main"),
//   		configurationSource: apprunner.configurationSourceType_REPOSITORY,
//   		connection: apprunner.gitHubConnection.fromConnectionArn(jsii.String("CONNECTION_ARN")),
//   	}),
//   })
//
// Experimental.
type ServiceProps struct {
	// The source of the repository for the service.
	// Experimental.
	Source Source `json:"source" yaml:"source"`
	// The IAM role that grants the App Runner service access to a source repository.
	//
	// It's required for ECR image repositories (but not for ECR Public repositories).
	//
	// The role must be assumable by the 'build.apprunner.amazonaws.com' service principal.
	// See: https://docs.aws.amazon.com/apprunner/latest/dg/security_iam_service-with-iam.html#security_iam_service-with-iam-roles-service.access
	//
	// Experimental.
	AccessRole awsiam.IRole `json:"accessRole" yaml:"accessRole"`
	// The number of CPU units reserved for each instance of your App Runner service.
	// Experimental.
	Cpu Cpu `json:"cpu" yaml:"cpu"`
	// The IAM role that provides permissions to your App Runner service.
	//
	// These are permissions that your code needs when it calls any AWS APIs.
	//
	// The role must be assumable by the 'tasks.apprunner.amazonaws.com' service principal.
	// See: https://docs.aws.amazon.com/apprunner/latest/dg/security_iam_service-with-iam.html#security_iam_service-with-iam-roles-service.instance
	//
	// Experimental.
	InstanceRole awsiam.IRole `json:"instanceRole" yaml:"instanceRole"`
	// The amount of memory reserved for each instance of your App Runner service.
	// Experimental.
	Memory Memory `json:"memory" yaml:"memory"`
	// Name of the service.
	// Experimental.
	ServiceName *string `json:"serviceName" yaml:"serviceName"`
}

// Represents the App Runner service source.
//
// Example:
//   apprunner.NewService(this, jsii.String("Service"), &serviceProps{
//   	source: apprunner.source.fromGitHub(&githubRepositoryProps{
//   		repositoryUrl: jsii.String("https://github.com/aws-containers/hello-app-runner"),
//   		branch: jsii.String("main"),
//   		configurationSource: apprunner.configurationSourceType_REPOSITORY,
//   		connection: apprunner.gitHubConnection.fromConnectionArn(jsii.String("CONNECTION_ARN")),
//   	}),
//   })
//
// Experimental.
type Source interface {
	// Called when the Job is initialized to allow this object to bind.
	// Experimental.
	Bind(scope constructs.Construct) *SourceConfig
}

// The jsii proxy struct for Source
type jsiiProxy_Source struct {
	_ byte // padding
}

// Experimental.
func NewSource_Override(s Source) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_apprunner.Source",
		nil, // no parameters
		s,
	)
}

// Source from local assets.
// Experimental.
func Source_FromAsset(props *AssetProps) AssetSource {
	_init_.Initialize()

	var returns AssetSource

	_jsii_.StaticInvoke(
		"monocdk.aws_apprunner.Source",
		"fromAsset",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Source from the ECR repository.
// Experimental.
func Source_FromEcr(props *EcrProps) EcrSource {
	_init_.Initialize()

	var returns EcrSource

	_jsii_.StaticInvoke(
		"monocdk.aws_apprunner.Source",
		"fromEcr",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Source from the ECR Public repository.
// Experimental.
func Source_FromEcrPublic(props *EcrPublicProps) EcrPublicSource {
	_init_.Initialize()

	var returns EcrPublicSource

	_jsii_.StaticInvoke(
		"monocdk.aws_apprunner.Source",
		"fromEcrPublic",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Source from the GitHub repository.
// Experimental.
func Source_FromGitHub(props *GithubRepositoryProps) GithubSource {
	_init_.Initialize()

	var returns GithubSource

	_jsii_.StaticInvoke(
		"monocdk.aws_apprunner.Source",
		"fromGitHub",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Source) Bind(scope constructs.Construct) *SourceConfig {
	var returns *SourceConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

// Identifies a version of code that AWS App Runner refers to within a source code repository.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import apprunner "github.com/aws/aws-cdk-go/awscdk/aws_apprunner"
//   sourceCodeVersion := &sourceCodeVersion{
//   	type: jsii.String("type"),
//   	value: jsii.String("value"),
//   }
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-sourcecodeversion.html
//
// Experimental.
type SourceCodeVersion struct {
	// The type of version identifier.
	// Experimental.
	Type *string `json:"type" yaml:"type"`
	// A source code version.
	// Experimental.
	Value *string `json:"value" yaml:"value"`
}

// Result of binding `Source` into a `Service`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import apprunner "github.com/aws/aws-cdk-go/awscdk/aws_apprunner"import awscdk "github.com/aws/aws-cdk-go/awscdk"import ecr "github.com/aws/aws-cdk-go/awscdk/aws_ecr"
//
//   var gitHubConnection gitHubConnection
//   var repository repository
//   var runtime runtime
//   sourceConfig := &sourceConfig{
//   	codeRepository: &codeRepositoryProps{
//   		codeConfiguration: &codeConfiguration{
//   			configurationSource: apprunner.configurationSourceType_REPOSITORY,
//
//   			// the properties below are optional
//   			configurationValues: &codeConfigurationValues{
//   				runtime: runtime,
//
//   				// the properties below are optional
//   				buildCommand: jsii.String("buildCommand"),
//   				environment: map[string]*string{
//   					"environmentKey": jsii.String("environment"),
//   				},
//   				port: jsii.String("port"),
//   				startCommand: jsii.String("startCommand"),
//   			},
//   		},
//   		connection: gitHubConnection,
//   		repositoryUrl: jsii.String("repositoryUrl"),
//   		sourceCodeVersion: &sourceCodeVersion{
//   			type: jsii.String("type"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	ecrRepository: repository,
//   	imageRepository: &imageRepository{
//   		imageIdentifier: jsii.String("imageIdentifier"),
//   		imageRepositoryType: apprunner.imageRepositoryType_ECR_PUBLIC,
//
//   		// the properties below are optional
//   		imageConfiguration: &imageConfiguration{
//   			environment: map[string]*string{
//   				"environmentKey": jsii.String("environment"),
//   			},
//   			port: jsii.Number(123),
//   			startCommand: jsii.String("startCommand"),
//   		},
//   	},
//   }
//
// Experimental.
type SourceConfig struct {
	// The code repository configuration (mutually exclusive  with `imageRepository`).
	// Experimental.
	CodeRepository *CodeRepositoryProps `json:"codeRepository" yaml:"codeRepository"`
	// The ECR repository (required to grant the pull privileges for the iam role).
	// Experimental.
	EcrRepository awsecr.IRepository `json:"ecrRepository" yaml:"ecrRepository"`
	// The image repository configuration (mutually exclusive  with `codeRepository`).
	// Experimental.
	ImageRepository *ImageRepository `json:"imageRepository" yaml:"imageRepository"`
}

