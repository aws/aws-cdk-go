package awscodedeploy

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsautoscaling"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodedeploy/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancing"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
)

// The configuration for automatically rolling back deployments in a given Deployment Group.
// Experimental.
type AutoRollbackConfig struct {
	// Whether to automatically roll back a deployment during which one of the configured CloudWatch alarms for this Deployment Group went off.
	// Experimental.
	DeploymentInAlarm *bool `json:"deploymentInAlarm"`
	// Whether to automatically roll back a deployment that fails.
	// Experimental.
	FailedDeployment *bool `json:"failedDeployment"`
	// Whether to automatically roll back a deployment that was manually stopped.
	// Experimental.
	StoppedDeployment *bool `json:"stoppedDeployment"`
}

// A CloudFormation `AWS::CodeDeploy::Application`.
type CfnApplication interface {
	awscdk.CfnResource
	awscdk.IInspectable
	ApplicationName() *string
	SetApplicationName(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	ComputePlatform() *string
	SetComputePlatform(val *string)
	CreationStack() *[]*string
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	Stack() awscdk.Stack
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
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnApplication
type jsiiProxy_CfnApplication struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnApplication) ApplicationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationName",
		&returns,
	)
	return returns
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

func (j *jsiiProxy_CfnApplication) ComputePlatform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computePlatform",
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

func (j *jsiiProxy_CfnApplication) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
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

func (j *jsiiProxy_CfnApplication) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::CodeDeploy::Application`.
func NewCfnApplication(scope constructs.Construct, id *string, props *CfnApplicationProps) CfnApplication {
	_init_.Initialize()

	j := jsiiProxy_CfnApplication{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codedeploy.CfnApplication",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CodeDeploy::Application`.
func NewCfnApplication_Override(c CfnApplication, scope constructs.Construct, id *string, props *CfnApplicationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codedeploy.CfnApplication",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnApplication) SetApplicationName(val *string) {
	_jsii_.Set(
		j,
		"applicationName",
		val,
	)
}

func (j *jsiiProxy_CfnApplication) SetComputePlatform(val *string) {
	_jsii_.Set(
		j,
		"computePlatform",
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
		"aws-cdk-lib.aws_codedeploy.CfnApplication",
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
		"aws-cdk-lib.aws_codedeploy.CfnApplication",
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
func CfnApplication_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codedeploy.CfnApplication",
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
		"aws-cdk-lib.aws_codedeploy.CfnApplication",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnApplication) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnApplication) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnApplication) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnApplication) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnApplication) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnApplication) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnApplication) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
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

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnApplication) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnApplication) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
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

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
// Experimental.
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

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
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

// Experimental.
func (c *jsiiProxy_CfnApplication) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `AWS::CodeDeploy::Application`.
type CfnApplicationProps struct {
	// `AWS::CodeDeploy::Application.ApplicationName`.
	ApplicationName *string `json:"applicationName"`
	// `AWS::CodeDeploy::Application.ComputePlatform`.
	ComputePlatform *string `json:"computePlatform"`
	// `AWS::CodeDeploy::Application.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
}

// A CloudFormation `AWS::CodeDeploy::DeploymentConfig`.
type CfnDeploymentConfig interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	ComputePlatform() *string
	SetComputePlatform(val *string)
	CreationStack() *[]*string
	DeploymentConfigName() *string
	SetDeploymentConfigName(val *string)
	LogicalId() *string
	MinimumHealthyHosts() interface{}
	SetMinimumHealthyHosts(val interface{})
	Node() constructs.Node
	Ref() *string
	Stack() awscdk.Stack
	TrafficRoutingConfig() interface{}
	SetTrafficRoutingConfig(val interface{})
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

// The jsii proxy struct for CfnDeploymentConfig
type jsiiProxy_CfnDeploymentConfig struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDeploymentConfig) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentConfig) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentConfig) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentConfig) ComputePlatform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computePlatform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentConfig) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentConfig) DeploymentConfigName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentConfigName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentConfig) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentConfig) MinimumHealthyHosts() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"minimumHealthyHosts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentConfig) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentConfig) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentConfig) TrafficRoutingConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trafficRoutingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentConfig) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::CodeDeploy::DeploymentConfig`.
func NewCfnDeploymentConfig(scope constructs.Construct, id *string, props *CfnDeploymentConfigProps) CfnDeploymentConfig {
	_init_.Initialize()

	j := jsiiProxy_CfnDeploymentConfig{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codedeploy.CfnDeploymentConfig",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CodeDeploy::DeploymentConfig`.
func NewCfnDeploymentConfig_Override(c CfnDeploymentConfig, scope constructs.Construct, id *string, props *CfnDeploymentConfigProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codedeploy.CfnDeploymentConfig",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDeploymentConfig) SetComputePlatform(val *string) {
	_jsii_.Set(
		j,
		"computePlatform",
		val,
	)
}

func (j *jsiiProxy_CfnDeploymentConfig) SetDeploymentConfigName(val *string) {
	_jsii_.Set(
		j,
		"deploymentConfigName",
		val,
	)
}

func (j *jsiiProxy_CfnDeploymentConfig) SetMinimumHealthyHosts(val interface{}) {
	_jsii_.Set(
		j,
		"minimumHealthyHosts",
		val,
	)
}

func (j *jsiiProxy_CfnDeploymentConfig) SetTrafficRoutingConfig(val interface{}) {
	_jsii_.Set(
		j,
		"trafficRoutingConfig",
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
func CfnDeploymentConfig_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codedeploy.CfnDeploymentConfig",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnDeploymentConfig_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codedeploy.CfnDeploymentConfig",
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
func CfnDeploymentConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codedeploy.CfnDeploymentConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDeploymentConfig_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codedeploy.CfnDeploymentConfig",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnDeploymentConfig) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnDeploymentConfig) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnDeploymentConfig) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnDeploymentConfig) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnDeploymentConfig) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnDeploymentConfig) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnDeploymentConfig) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnDeploymentConfig) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnDeploymentConfig) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnDeploymentConfig) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnDeploymentConfig) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDeploymentConfig) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnDeploymentConfig) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnDeploymentConfig) ToString() *string {
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
func (c *jsiiProxy_CfnDeploymentConfig) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnDeploymentConfig_MinimumHealthyHostsProperty struct {
	// `CfnDeploymentConfig.MinimumHealthyHostsProperty.Type`.
	Type *string `json:"type"`
	// `CfnDeploymentConfig.MinimumHealthyHostsProperty.Value`.
	Value *float64 `json:"value"`
}

type CfnDeploymentConfig_TimeBasedCanaryProperty struct {
	// `CfnDeploymentConfig.TimeBasedCanaryProperty.CanaryInterval`.
	CanaryInterval *float64 `json:"canaryInterval"`
	// `CfnDeploymentConfig.TimeBasedCanaryProperty.CanaryPercentage`.
	CanaryPercentage *float64 `json:"canaryPercentage"`
}

type CfnDeploymentConfig_TimeBasedLinearProperty struct {
	// `CfnDeploymentConfig.TimeBasedLinearProperty.LinearInterval`.
	LinearInterval *float64 `json:"linearInterval"`
	// `CfnDeploymentConfig.TimeBasedLinearProperty.LinearPercentage`.
	LinearPercentage *float64 `json:"linearPercentage"`
}

type CfnDeploymentConfig_TrafficRoutingConfigProperty struct {
	// `CfnDeploymentConfig.TrafficRoutingConfigProperty.Type`.
	Type *string `json:"type"`
	// `CfnDeploymentConfig.TrafficRoutingConfigProperty.TimeBasedCanary`.
	TimeBasedCanary interface{} `json:"timeBasedCanary"`
	// `CfnDeploymentConfig.TrafficRoutingConfigProperty.TimeBasedLinear`.
	TimeBasedLinear interface{} `json:"timeBasedLinear"`
}

// Properties for defining a `AWS::CodeDeploy::DeploymentConfig`.
type CfnDeploymentConfigProps struct {
	// `AWS::CodeDeploy::DeploymentConfig.ComputePlatform`.
	ComputePlatform *string `json:"computePlatform"`
	// `AWS::CodeDeploy::DeploymentConfig.DeploymentConfigName`.
	DeploymentConfigName *string `json:"deploymentConfigName"`
	// `AWS::CodeDeploy::DeploymentConfig.MinimumHealthyHosts`.
	MinimumHealthyHosts interface{} `json:"minimumHealthyHosts"`
	// `AWS::CodeDeploy::DeploymentConfig.TrafficRoutingConfig`.
	TrafficRoutingConfig interface{} `json:"trafficRoutingConfig"`
}

// A CloudFormation `AWS::CodeDeploy::DeploymentGroup`.
type CfnDeploymentGroup interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AlarmConfiguration() interface{}
	SetAlarmConfiguration(val interface{})
	ApplicationName() *string
	SetApplicationName(val *string)
	AutoRollbackConfiguration() interface{}
	SetAutoRollbackConfiguration(val interface{})
	AutoScalingGroups() *[]*string
	SetAutoScalingGroups(val *[]*string)
	BlueGreenDeploymentConfiguration() interface{}
	SetBlueGreenDeploymentConfiguration(val interface{})
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	Deployment() interface{}
	SetDeployment(val interface{})
	DeploymentConfigName() *string
	SetDeploymentConfigName(val *string)
	DeploymentGroupName() *string
	SetDeploymentGroupName(val *string)
	DeploymentStyle() interface{}
	SetDeploymentStyle(val interface{})
	Ec2TagFilters() interface{}
	SetEc2TagFilters(val interface{})
	Ec2TagSet() interface{}
	SetEc2TagSet(val interface{})
	EcsServices() interface{}
	SetEcsServices(val interface{})
	LoadBalancerInfo() interface{}
	SetLoadBalancerInfo(val interface{})
	LogicalId() *string
	Node() constructs.Node
	OnPremisesInstanceTagFilters() interface{}
	SetOnPremisesInstanceTagFilters(val interface{})
	OnPremisesTagSet() interface{}
	SetOnPremisesTagSet(val interface{})
	Ref() *string
	ServiceRoleArn() *string
	SetServiceRoleArn(val *string)
	Stack() awscdk.Stack
	TriggerConfigurations() interface{}
	SetTriggerConfigurations(val interface{})
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

// The jsii proxy struct for CfnDeploymentGroup
type jsiiProxy_CfnDeploymentGroup struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDeploymentGroup) AlarmConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alarmConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentGroup) ApplicationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentGroup) AutoRollbackConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoRollbackConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentGroup) AutoScalingGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"autoScalingGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentGroup) BlueGreenDeploymentConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blueGreenDeploymentConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentGroup) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentGroup) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentGroup) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentGroup) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentGroup) Deployment() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentGroup) DeploymentConfigName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentConfigName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentGroup) DeploymentGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentGroup) DeploymentStyle() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deploymentStyle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentGroup) Ec2TagFilters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ec2TagFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentGroup) Ec2TagSet() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ec2TagSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentGroup) EcsServices() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ecsServices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentGroup) LoadBalancerInfo() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loadBalancerInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentGroup) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentGroup) OnPremisesInstanceTagFilters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onPremisesInstanceTagFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentGroup) OnPremisesTagSet() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onPremisesTagSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentGroup) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentGroup) ServiceRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentGroup) TriggerConfigurations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"triggerConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentGroup) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::CodeDeploy::DeploymentGroup`.
func NewCfnDeploymentGroup(scope constructs.Construct, id *string, props *CfnDeploymentGroupProps) CfnDeploymentGroup {
	_init_.Initialize()

	j := jsiiProxy_CfnDeploymentGroup{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codedeploy.CfnDeploymentGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CodeDeploy::DeploymentGroup`.
func NewCfnDeploymentGroup_Override(c CfnDeploymentGroup, scope constructs.Construct, id *string, props *CfnDeploymentGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codedeploy.CfnDeploymentGroup",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDeploymentGroup) SetAlarmConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"alarmConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnDeploymentGroup) SetApplicationName(val *string) {
	_jsii_.Set(
		j,
		"applicationName",
		val,
	)
}

func (j *jsiiProxy_CfnDeploymentGroup) SetAutoRollbackConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"autoRollbackConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnDeploymentGroup) SetAutoScalingGroups(val *[]*string) {
	_jsii_.Set(
		j,
		"autoScalingGroups",
		val,
	)
}

func (j *jsiiProxy_CfnDeploymentGroup) SetBlueGreenDeploymentConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"blueGreenDeploymentConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnDeploymentGroup) SetDeployment(val interface{}) {
	_jsii_.Set(
		j,
		"deployment",
		val,
	)
}

func (j *jsiiProxy_CfnDeploymentGroup) SetDeploymentConfigName(val *string) {
	_jsii_.Set(
		j,
		"deploymentConfigName",
		val,
	)
}

func (j *jsiiProxy_CfnDeploymentGroup) SetDeploymentGroupName(val *string) {
	_jsii_.Set(
		j,
		"deploymentGroupName",
		val,
	)
}

func (j *jsiiProxy_CfnDeploymentGroup) SetDeploymentStyle(val interface{}) {
	_jsii_.Set(
		j,
		"deploymentStyle",
		val,
	)
}

func (j *jsiiProxy_CfnDeploymentGroup) SetEc2TagFilters(val interface{}) {
	_jsii_.Set(
		j,
		"ec2TagFilters",
		val,
	)
}

func (j *jsiiProxy_CfnDeploymentGroup) SetEc2TagSet(val interface{}) {
	_jsii_.Set(
		j,
		"ec2TagSet",
		val,
	)
}

func (j *jsiiProxy_CfnDeploymentGroup) SetEcsServices(val interface{}) {
	_jsii_.Set(
		j,
		"ecsServices",
		val,
	)
}

func (j *jsiiProxy_CfnDeploymentGroup) SetLoadBalancerInfo(val interface{}) {
	_jsii_.Set(
		j,
		"loadBalancerInfo",
		val,
	)
}

func (j *jsiiProxy_CfnDeploymentGroup) SetOnPremisesInstanceTagFilters(val interface{}) {
	_jsii_.Set(
		j,
		"onPremisesInstanceTagFilters",
		val,
	)
}

func (j *jsiiProxy_CfnDeploymentGroup) SetOnPremisesTagSet(val interface{}) {
	_jsii_.Set(
		j,
		"onPremisesTagSet",
		val,
	)
}

func (j *jsiiProxy_CfnDeploymentGroup) SetServiceRoleArn(val *string) {
	_jsii_.Set(
		j,
		"serviceRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnDeploymentGroup) SetTriggerConfigurations(val interface{}) {
	_jsii_.Set(
		j,
		"triggerConfigurations",
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
func CfnDeploymentGroup_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codedeploy.CfnDeploymentGroup",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnDeploymentGroup_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codedeploy.CfnDeploymentGroup",
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
func CfnDeploymentGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codedeploy.CfnDeploymentGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDeploymentGroup_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codedeploy.CfnDeploymentGroup",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnDeploymentGroup) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnDeploymentGroup) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnDeploymentGroup) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnDeploymentGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnDeploymentGroup) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnDeploymentGroup) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnDeploymentGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnDeploymentGroup) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnDeploymentGroup) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnDeploymentGroup) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnDeploymentGroup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDeploymentGroup) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnDeploymentGroup) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnDeploymentGroup) ToString() *string {
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
func (c *jsiiProxy_CfnDeploymentGroup) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnDeploymentGroup_AlarmConfigurationProperty struct {
	// `CfnDeploymentGroup.AlarmConfigurationProperty.Alarms`.
	Alarms interface{} `json:"alarms"`
	// `CfnDeploymentGroup.AlarmConfigurationProperty.Enabled`.
	Enabled interface{} `json:"enabled"`
	// `CfnDeploymentGroup.AlarmConfigurationProperty.IgnorePollAlarmFailure`.
	IgnorePollAlarmFailure interface{} `json:"ignorePollAlarmFailure"`
}

type CfnDeploymentGroup_AlarmProperty struct {
	// `CfnDeploymentGroup.AlarmProperty.Name`.
	Name *string `json:"name"`
}

type CfnDeploymentGroup_AutoRollbackConfigurationProperty struct {
	// `CfnDeploymentGroup.AutoRollbackConfigurationProperty.Enabled`.
	Enabled interface{} `json:"enabled"`
	// `CfnDeploymentGroup.AutoRollbackConfigurationProperty.Events`.
	Events *[]*string `json:"events"`
}

type CfnDeploymentGroup_BlueGreenDeploymentConfigurationProperty struct {
	// `CfnDeploymentGroup.BlueGreenDeploymentConfigurationProperty.DeploymentReadyOption`.
	DeploymentReadyOption interface{} `json:"deploymentReadyOption"`
	// `CfnDeploymentGroup.BlueGreenDeploymentConfigurationProperty.GreenFleetProvisioningOption`.
	GreenFleetProvisioningOption interface{} `json:"greenFleetProvisioningOption"`
	// `CfnDeploymentGroup.BlueGreenDeploymentConfigurationProperty.TerminateBlueInstancesOnDeploymentSuccess`.
	TerminateBlueInstancesOnDeploymentSuccess interface{} `json:"terminateBlueInstancesOnDeploymentSuccess"`
}

type CfnDeploymentGroup_BlueInstanceTerminationOptionProperty struct {
	// `CfnDeploymentGroup.BlueInstanceTerminationOptionProperty.Action`.
	Action *string `json:"action"`
	// `CfnDeploymentGroup.BlueInstanceTerminationOptionProperty.TerminationWaitTimeInMinutes`.
	TerminationWaitTimeInMinutes *float64 `json:"terminationWaitTimeInMinutes"`
}

type CfnDeploymentGroup_DeploymentProperty struct {
	// `CfnDeploymentGroup.DeploymentProperty.Revision`.
	Revision interface{} `json:"revision"`
	// `CfnDeploymentGroup.DeploymentProperty.Description`.
	Description *string `json:"description"`
	// `CfnDeploymentGroup.DeploymentProperty.IgnoreApplicationStopFailures`.
	IgnoreApplicationStopFailures interface{} `json:"ignoreApplicationStopFailures"`
}

type CfnDeploymentGroup_DeploymentReadyOptionProperty struct {
	// `CfnDeploymentGroup.DeploymentReadyOptionProperty.ActionOnTimeout`.
	ActionOnTimeout *string `json:"actionOnTimeout"`
	// `CfnDeploymentGroup.DeploymentReadyOptionProperty.WaitTimeInMinutes`.
	WaitTimeInMinutes *float64 `json:"waitTimeInMinutes"`
}

type CfnDeploymentGroup_DeploymentStyleProperty struct {
	// `CfnDeploymentGroup.DeploymentStyleProperty.DeploymentOption`.
	DeploymentOption *string `json:"deploymentOption"`
	// `CfnDeploymentGroup.DeploymentStyleProperty.DeploymentType`.
	DeploymentType *string `json:"deploymentType"`
}

type CfnDeploymentGroup_EC2TagFilterProperty struct {
	// `CfnDeploymentGroup.EC2TagFilterProperty.Key`.
	Key *string `json:"key"`
	// `CfnDeploymentGroup.EC2TagFilterProperty.Type`.
	Type *string `json:"type"`
	// `CfnDeploymentGroup.EC2TagFilterProperty.Value`.
	Value *string `json:"value"`
}

type CfnDeploymentGroup_EC2TagSetListObjectProperty struct {
	// `CfnDeploymentGroup.EC2TagSetListObjectProperty.Ec2TagGroup`.
	Ec2TagGroup interface{} `json:"ec2TagGroup"`
}

type CfnDeploymentGroup_EC2TagSetProperty struct {
	// `CfnDeploymentGroup.EC2TagSetProperty.Ec2TagSetList`.
	Ec2TagSetList interface{} `json:"ec2TagSetList"`
}

type CfnDeploymentGroup_ECSServiceProperty struct {
	// `CfnDeploymentGroup.ECSServiceProperty.ClusterName`.
	ClusterName *string `json:"clusterName"`
	// `CfnDeploymentGroup.ECSServiceProperty.ServiceName`.
	ServiceName *string `json:"serviceName"`
}

type CfnDeploymentGroup_ELBInfoProperty struct {
	// `CfnDeploymentGroup.ELBInfoProperty.Name`.
	Name *string `json:"name"`
}

type CfnDeploymentGroup_GitHubLocationProperty struct {
	// `CfnDeploymentGroup.GitHubLocationProperty.CommitId`.
	CommitId *string `json:"commitId"`
	// `CfnDeploymentGroup.GitHubLocationProperty.Repository`.
	Repository *string `json:"repository"`
}

type CfnDeploymentGroup_GreenFleetProvisioningOptionProperty struct {
	// `CfnDeploymentGroup.GreenFleetProvisioningOptionProperty.Action`.
	Action *string `json:"action"`
}

type CfnDeploymentGroup_LoadBalancerInfoProperty struct {
	// `CfnDeploymentGroup.LoadBalancerInfoProperty.ElbInfoList`.
	ElbInfoList interface{} `json:"elbInfoList"`
	// `CfnDeploymentGroup.LoadBalancerInfoProperty.TargetGroupInfoList`.
	TargetGroupInfoList interface{} `json:"targetGroupInfoList"`
}

type CfnDeploymentGroup_OnPremisesTagSetListObjectProperty struct {
	// `CfnDeploymentGroup.OnPremisesTagSetListObjectProperty.OnPremisesTagGroup`.
	OnPremisesTagGroup interface{} `json:"onPremisesTagGroup"`
}

type CfnDeploymentGroup_OnPremisesTagSetProperty struct {
	// `CfnDeploymentGroup.OnPremisesTagSetProperty.OnPremisesTagSetList`.
	OnPremisesTagSetList interface{} `json:"onPremisesTagSetList"`
}

type CfnDeploymentGroup_RevisionLocationProperty struct {
	// `CfnDeploymentGroup.RevisionLocationProperty.GitHubLocation`.
	GitHubLocation interface{} `json:"gitHubLocation"`
	// `CfnDeploymentGroup.RevisionLocationProperty.RevisionType`.
	RevisionType *string `json:"revisionType"`
	// `CfnDeploymentGroup.RevisionLocationProperty.S3Location`.
	S3Location interface{} `json:"s3Location"`
}

type CfnDeploymentGroup_S3LocationProperty struct {
	// `CfnDeploymentGroup.S3LocationProperty.Bucket`.
	Bucket *string `json:"bucket"`
	// `CfnDeploymentGroup.S3LocationProperty.Key`.
	Key *string `json:"key"`
	// `CfnDeploymentGroup.S3LocationProperty.BundleType`.
	BundleType *string `json:"bundleType"`
	// `CfnDeploymentGroup.S3LocationProperty.ETag`.
	ETag *string `json:"eTag"`
	// `CfnDeploymentGroup.S3LocationProperty.Version`.
	Version *string `json:"version"`
}

type CfnDeploymentGroup_TagFilterProperty struct {
	// `CfnDeploymentGroup.TagFilterProperty.Key`.
	Key *string `json:"key"`
	// `CfnDeploymentGroup.TagFilterProperty.Type`.
	Type *string `json:"type"`
	// `CfnDeploymentGroup.TagFilterProperty.Value`.
	Value *string `json:"value"`
}

type CfnDeploymentGroup_TargetGroupInfoProperty struct {
	// `CfnDeploymentGroup.TargetGroupInfoProperty.Name`.
	Name *string `json:"name"`
}

type CfnDeploymentGroup_TriggerConfigProperty struct {
	// `CfnDeploymentGroup.TriggerConfigProperty.TriggerEvents`.
	TriggerEvents *[]*string `json:"triggerEvents"`
	// `CfnDeploymentGroup.TriggerConfigProperty.TriggerName`.
	TriggerName *string `json:"triggerName"`
	// `CfnDeploymentGroup.TriggerConfigProperty.TriggerTargetArn`.
	TriggerTargetArn *string `json:"triggerTargetArn"`
}

// Properties for defining a `AWS::CodeDeploy::DeploymentGroup`.
type CfnDeploymentGroupProps struct {
	// `AWS::CodeDeploy::DeploymentGroup.ApplicationName`.
	ApplicationName *string `json:"applicationName"`
	// `AWS::CodeDeploy::DeploymentGroup.ServiceRoleArn`.
	ServiceRoleArn *string `json:"serviceRoleArn"`
	// `AWS::CodeDeploy::DeploymentGroup.AlarmConfiguration`.
	AlarmConfiguration interface{} `json:"alarmConfiguration"`
	// `AWS::CodeDeploy::DeploymentGroup.AutoRollbackConfiguration`.
	AutoRollbackConfiguration interface{} `json:"autoRollbackConfiguration"`
	// `AWS::CodeDeploy::DeploymentGroup.AutoScalingGroups`.
	AutoScalingGroups *[]*string `json:"autoScalingGroups"`
	// `AWS::CodeDeploy::DeploymentGroup.BlueGreenDeploymentConfiguration`.
	BlueGreenDeploymentConfiguration interface{} `json:"blueGreenDeploymentConfiguration"`
	// `AWS::CodeDeploy::DeploymentGroup.Deployment`.
	Deployment interface{} `json:"deployment"`
	// `AWS::CodeDeploy::DeploymentGroup.DeploymentConfigName`.
	DeploymentConfigName *string `json:"deploymentConfigName"`
	// `AWS::CodeDeploy::DeploymentGroup.DeploymentGroupName`.
	DeploymentGroupName *string `json:"deploymentGroupName"`
	// `AWS::CodeDeploy::DeploymentGroup.DeploymentStyle`.
	DeploymentStyle interface{} `json:"deploymentStyle"`
	// `AWS::CodeDeploy::DeploymentGroup.Ec2TagFilters`.
	Ec2TagFilters interface{} `json:"ec2TagFilters"`
	// `AWS::CodeDeploy::DeploymentGroup.Ec2TagSet`.
	Ec2TagSet interface{} `json:"ec2TagSet"`
	// `AWS::CodeDeploy::DeploymentGroup.ECSServices`.
	EcsServices interface{} `json:"ecsServices"`
	// `AWS::CodeDeploy::DeploymentGroup.LoadBalancerInfo`.
	LoadBalancerInfo interface{} `json:"loadBalancerInfo"`
	// `AWS::CodeDeploy::DeploymentGroup.OnPremisesInstanceTagFilters`.
	OnPremisesInstanceTagFilters interface{} `json:"onPremisesInstanceTagFilters"`
	// `AWS::CodeDeploy::DeploymentGroup.OnPremisesTagSet`.
	OnPremisesTagSet interface{} `json:"onPremisesTagSet"`
	// `AWS::CodeDeploy::DeploymentGroup.TriggerConfigurations`.
	TriggerConfigurations interface{} `json:"triggerConfigurations"`
}

// A custom Deployment Configuration for a Lambda Deployment Group.
// Experimental.
type CustomLambdaDeploymentConfig interface {
	awscdk.Resource
	ILambdaDeploymentConfig
	DeploymentConfigArn() *string
	DeploymentConfigName() *string
	Env() *awscdk.ResourceEnvironment
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for CustomLambdaDeploymentConfig
type jsiiProxy_CustomLambdaDeploymentConfig struct {
	internal.Type__awscdkResource
	jsiiProxy_ILambdaDeploymentConfig
}

func (j *jsiiProxy_CustomLambdaDeploymentConfig) DeploymentConfigArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentConfigArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomLambdaDeploymentConfig) DeploymentConfigName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentConfigName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomLambdaDeploymentConfig) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomLambdaDeploymentConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomLambdaDeploymentConfig) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomLambdaDeploymentConfig) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewCustomLambdaDeploymentConfig(scope constructs.Construct, id *string, props *CustomLambdaDeploymentConfigProps) CustomLambdaDeploymentConfig {
	_init_.Initialize()

	j := jsiiProxy_CustomLambdaDeploymentConfig{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codedeploy.CustomLambdaDeploymentConfig",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewCustomLambdaDeploymentConfig_Override(c CustomLambdaDeploymentConfig, scope constructs.Construct, id *string, props *CustomLambdaDeploymentConfigProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codedeploy.CustomLambdaDeploymentConfig",
		[]interface{}{scope, id, props},
		c,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CustomLambdaDeploymentConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codedeploy.CustomLambdaDeploymentConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func CustomLambdaDeploymentConfig_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codedeploy.CustomLambdaDeploymentConfig",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

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
func (c *jsiiProxy_CustomLambdaDeploymentConfig) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (c *jsiiProxy_CustomLambdaDeploymentConfig) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
func (c *jsiiProxy_CustomLambdaDeploymentConfig) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
func (c *jsiiProxy_CustomLambdaDeploymentConfig) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (c *jsiiProxy_CustomLambdaDeploymentConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties of a reference to a CodeDeploy Lambda Deployment Configuration.
// Experimental.
type CustomLambdaDeploymentConfigProps struct {
	// The interval, in number of minutes: - For LINEAR, how frequently additional traffic is shifted - For CANARY, how long to shift traffic before the full deployment.
	// Experimental.
	Interval awscdk.Duration `json:"interval"`
	// The integer percentage of traffic to shift: - For LINEAR, the percentage to shift every interval - For CANARY, the percentage to shift until the interval passes, before the full deployment.
	// Experimental.
	Percentage *float64 `json:"percentage"`
	// The type of deployment config, either CANARY or LINEAR.
	// Experimental.
	Type CustomLambdaDeploymentConfigType `json:"type"`
	// The verbatim name of the deployment config.
	//
	// Must be unique per account/region.
	// Other parameters cannot be updated if this name is provided.
	// Experimental.
	DeploymentConfigName *string `json:"deploymentConfigName"`
}

// Lambda Deployment config type.
// Experimental.
type CustomLambdaDeploymentConfigType string

const (
	CustomLambdaDeploymentConfigType_CANARY CustomLambdaDeploymentConfigType = "CANARY"
	CustomLambdaDeploymentConfigType_LINEAR CustomLambdaDeploymentConfigType = "LINEAR"
)

// A CodeDeploy Application that deploys to an Amazon ECS service.
// Experimental.
type EcsApplication interface {
	awscdk.Resource
	IEcsApplication
	ApplicationArn() *string
	ApplicationName() *string
	Env() *awscdk.ResourceEnvironment
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for EcsApplication
type jsiiProxy_EcsApplication struct {
	internal.Type__awscdkResource
	jsiiProxy_IEcsApplication
}

func (j *jsiiProxy_EcsApplication) ApplicationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsApplication) ApplicationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsApplication) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsApplication) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsApplication) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsApplication) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewEcsApplication(scope constructs.Construct, id *string, props *EcsApplicationProps) EcsApplication {
	_init_.Initialize()

	j := jsiiProxy_EcsApplication{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codedeploy.EcsApplication",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewEcsApplication_Override(e EcsApplication, scope constructs.Construct, id *string, props *EcsApplicationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codedeploy.EcsApplication",
		[]interface{}{scope, id, props},
		e,
	)
}

// Import an Application defined either outside the CDK, or in a different CDK Stack.
//
// Returns: a Construct representing a reference to an existing Application
// Experimental.
func EcsApplication_FromEcsApplicationName(scope constructs.Construct, id *string, ecsApplicationName *string) IEcsApplication {
	_init_.Initialize()

	var returns IEcsApplication

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codedeploy.EcsApplication",
		"fromEcsApplicationName",
		[]interface{}{scope, id, ecsApplicationName},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func EcsApplication_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codedeploy.EcsApplication",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func EcsApplication_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codedeploy.EcsApplication",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

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
func (e *jsiiProxy_EcsApplication) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		e,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (e *jsiiProxy_EcsApplication) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
func (e *jsiiProxy_EcsApplication) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
func (e *jsiiProxy_EcsApplication) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (e *jsiiProxy_EcsApplication) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Construction properties for {@link EcsApplication}.
// Experimental.
type EcsApplicationProps struct {
	// The physical, human-readable name of the CodeDeploy Application.
	// Experimental.
	ApplicationName *string `json:"applicationName"`
}

// A custom Deployment Configuration for an ECS Deployment Group.
//
// Note: This class currently stands as namespaced container of the default configurations
// until CloudFormation supports custom ECS Deployment Configs. Until then it is closed
// (private constructor) and does not extend {@link Construct}
// Experimental.
type EcsDeploymentConfig interface {
}

// The jsii proxy struct for EcsDeploymentConfig
type jsiiProxy_EcsDeploymentConfig struct {
	_ byte // padding
}

// Import a custom Deployment Configuration for an ECS Deployment Group defined outside the CDK.
//
// Returns: a Construct representing a reference to an existing custom Deployment Configuration
// Experimental.
func EcsDeploymentConfig_FromEcsDeploymentConfigName(_scope constructs.Construct, _id *string, ecsDeploymentConfigName *string) IEcsDeploymentConfig {
	_init_.Initialize()

	var returns IEcsDeploymentConfig

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codedeploy.EcsDeploymentConfig",
		"fromEcsDeploymentConfigName",
		[]interface{}{_scope, _id, ecsDeploymentConfigName},
		&returns,
	)

	return returns
}

func EcsDeploymentConfig_ALL_AT_ONCE() IEcsDeploymentConfig {
	_init_.Initialize()
	var returns IEcsDeploymentConfig
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codedeploy.EcsDeploymentConfig",
		"ALL_AT_ONCE",
		&returns,
	)
	return returns
}

// Note: This class currently stands as a namespaced container for importing an ECS Deployment Group defined outside the CDK app until CloudFormation supports provisioning ECS Deployment Groups.
//
// Until then it is closed (private constructor) and does not
// extend {@link Construct}.
// Experimental.
type EcsDeploymentGroup interface {
}

// The jsii proxy struct for EcsDeploymentGroup
type jsiiProxy_EcsDeploymentGroup struct {
	_ byte // padding
}

// Import an ECS Deployment Group defined outside the CDK app.
//
// Returns: a Construct representing a reference to an existing Deployment Group
// Experimental.
func EcsDeploymentGroup_FromEcsDeploymentGroupAttributes(scope constructs.Construct, id *string, attrs *EcsDeploymentGroupAttributes) IEcsDeploymentGroup {
	_init_.Initialize()

	var returns IEcsDeploymentGroup

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codedeploy.EcsDeploymentGroup",
		"fromEcsDeploymentGroupAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Properties of a reference to a CodeDeploy ECS Deployment Group.
// See: EcsDeploymentGroup#fromEcsDeploymentGroupAttributes
//
// Experimental.
type EcsDeploymentGroupAttributes struct {
	// The reference to the CodeDeploy ECS Application that this Deployment Group belongs to.
	// Experimental.
	Application IEcsApplication `json:"application"`
	// The physical, human-readable name of the CodeDeploy ECS Deployment Group that we are referencing.
	// Experimental.
	DeploymentGroupName *string `json:"deploymentGroupName"`
	// The Deployment Configuration this Deployment Group uses.
	// Experimental.
	DeploymentConfig IEcsDeploymentConfig `json:"deploymentConfig"`
}

// Represents a reference to a CodeDeploy Application deploying to Amazon ECS.
//
// If you're managing the Application alongside the rest of your CDK resources,
// use the {@link EcsApplication} class.
//
// If you want to reference an already existing Application,
// or one defined in a different CDK Stack,
// use the {@link EcsApplication#fromEcsApplicationName} method.
// Experimental.
type IEcsApplication interface {
	awscdk.IResource
	// Experimental.
	ApplicationArn() *string
	// Experimental.
	ApplicationName() *string
}

// The jsii proxy for IEcsApplication
type jsiiProxy_IEcsApplication struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IEcsApplication) ApplicationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEcsApplication) ApplicationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationName",
		&returns,
	)
	return returns
}

// The Deployment Configuration of an ECS Deployment Group.
//
// The default, pre-defined Configurations are available as constants on the {@link EcsDeploymentConfig} class
// (for example, `EcsDeploymentConfig.AllAtOnce`).
//
// Note: CloudFormation does not currently support creating custom ECS configs outside
// of using a custom resource. You can import custom deployment config created outside the
// CDK or via a custom resource with {@link EcsDeploymentConfig#fromEcsDeploymentConfigName}.
// Experimental.
type IEcsDeploymentConfig interface {
	// Experimental.
	DeploymentConfigArn() *string
	// Experimental.
	DeploymentConfigName() *string
}

// The jsii proxy for IEcsDeploymentConfig
type jsiiProxy_IEcsDeploymentConfig struct {
	_ byte // padding
}

func (j *jsiiProxy_IEcsDeploymentConfig) DeploymentConfigArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentConfigArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEcsDeploymentConfig) DeploymentConfigName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentConfigName",
		&returns,
	)
	return returns
}

// Interface for an ECS deployment group.
// Experimental.
type IEcsDeploymentGroup interface {
	awscdk.IResource
	// The reference to the CodeDeploy ECS Application that this Deployment Group belongs to.
	// Experimental.
	Application() IEcsApplication
	// The Deployment Configuration this Group uses.
	// Experimental.
	DeploymentConfig() IEcsDeploymentConfig
	// The ARN of this Deployment Group.
	// Experimental.
	DeploymentGroupArn() *string
	// The physical name of the CodeDeploy Deployment Group.
	// Experimental.
	DeploymentGroupName() *string
}

// The jsii proxy for IEcsDeploymentGroup
type jsiiProxy_IEcsDeploymentGroup struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IEcsDeploymentGroup) Application() IEcsApplication {
	var returns IEcsApplication
	_jsii_.Get(
		j,
		"application",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEcsDeploymentGroup) DeploymentConfig() IEcsDeploymentConfig {
	var returns IEcsDeploymentConfig
	_jsii_.Get(
		j,
		"deploymentConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEcsDeploymentGroup) DeploymentGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEcsDeploymentGroup) DeploymentGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentGroupName",
		&returns,
	)
	return returns
}

// Represents a reference to a CodeDeploy Application deploying to AWS Lambda.
//
// If you're managing the Application alongside the rest of your CDK resources,
// use the {@link LambdaApplication} class.
//
// If you want to reference an already existing Application,
// or one defined in a different CDK Stack,
// use the {@link LambdaApplication#fromLambdaApplicationName} method.
// Experimental.
type ILambdaApplication interface {
	awscdk.IResource
	// Experimental.
	ApplicationArn() *string
	// Experimental.
	ApplicationName() *string
}

// The jsii proxy for ILambdaApplication
type jsiiProxy_ILambdaApplication struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_ILambdaApplication) ApplicationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILambdaApplication) ApplicationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationName",
		&returns,
	)
	return returns
}

// The Deployment Configuration of a Lambda Deployment Group.
//
// The default, pre-defined Configurations are available as constants on the {@link LambdaDeploymentConfig} class
// (`LambdaDeploymentConfig.AllAtOnce`, `LambdaDeploymentConfig.Canary10Percent30Minutes`, etc.).
//
// Note: CloudFormation does not currently support creating custom lambda configs outside
// of using a custom resource. You can import custom deployment config created outside the
// CDK or via a custom resource with {@link LambdaDeploymentConfig#import}.
// Experimental.
type ILambdaDeploymentConfig interface {
	// Experimental.
	DeploymentConfigArn() *string
	// Experimental.
	DeploymentConfigName() *string
}

// The jsii proxy for ILambdaDeploymentConfig
type jsiiProxy_ILambdaDeploymentConfig struct {
	_ byte // padding
}

func (j *jsiiProxy_ILambdaDeploymentConfig) DeploymentConfigArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentConfigArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILambdaDeploymentConfig) DeploymentConfigName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentConfigName",
		&returns,
	)
	return returns
}

// Interface for a Lambda deployment groups.
// Experimental.
type ILambdaDeploymentGroup interface {
	awscdk.IResource
	// The reference to the CodeDeploy Lambda Application that this Deployment Group belongs to.
	// Experimental.
	Application() ILambdaApplication
	// The Deployment Configuration this Group uses.
	// Experimental.
	DeploymentConfig() ILambdaDeploymentConfig
	// The ARN of this Deployment Group.
	// Experimental.
	DeploymentGroupArn() *string
	// The physical name of the CodeDeploy Deployment Group.
	// Experimental.
	DeploymentGroupName() *string
}

// The jsii proxy for ILambdaDeploymentGroup
type jsiiProxy_ILambdaDeploymentGroup struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_ILambdaDeploymentGroup) Application() ILambdaApplication {
	var returns ILambdaApplication
	_jsii_.Get(
		j,
		"application",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILambdaDeploymentGroup) DeploymentConfig() ILambdaDeploymentConfig {
	var returns ILambdaDeploymentConfig
	_jsii_.Get(
		j,
		"deploymentConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILambdaDeploymentGroup) DeploymentGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILambdaDeploymentGroup) DeploymentGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentGroupName",
		&returns,
	)
	return returns
}

// Represents a reference to a CodeDeploy Application deploying to EC2/on-premise instances.
//
// If you're managing the Application alongside the rest of your CDK resources,
// use the {@link ServerApplication} class.
//
// If you want to reference an already existing Application,
// or one defined in a different CDK Stack,
// use the {@link #fromServerApplicationName} method.
// Experimental.
type IServerApplication interface {
	awscdk.IResource
	// Experimental.
	ApplicationArn() *string
	// Experimental.
	ApplicationName() *string
}

// The jsii proxy for IServerApplication
type jsiiProxy_IServerApplication struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IServerApplication) ApplicationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServerApplication) ApplicationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationName",
		&returns,
	)
	return returns
}

// The Deployment Configuration of an EC2/on-premise Deployment Group.
//
// The default, pre-defined Configurations are available as constants on the {@link ServerDeploymentConfig} class
// (`ServerDeploymentConfig.HALF_AT_A_TIME`, `ServerDeploymentConfig.ALL_AT_ONCE`, etc.).
// To create a custom Deployment Configuration,
// instantiate the {@link ServerDeploymentConfig} Construct.
// Experimental.
type IServerDeploymentConfig interface {
	// Experimental.
	DeploymentConfigArn() *string
	// Experimental.
	DeploymentConfigName() *string
}

// The jsii proxy for IServerDeploymentConfig
type jsiiProxy_IServerDeploymentConfig struct {
	_ byte // padding
}

func (j *jsiiProxy_IServerDeploymentConfig) DeploymentConfigArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentConfigArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServerDeploymentConfig) DeploymentConfigName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentConfigName",
		&returns,
	)
	return returns
}

// Experimental.
type IServerDeploymentGroup interface {
	awscdk.IResource
	// Experimental.
	Application() IServerApplication
	// Experimental.
	AutoScalingGroups() *[]awsautoscaling.IAutoScalingGroup
	// Experimental.
	DeploymentConfig() IServerDeploymentConfig
	// Experimental.
	DeploymentGroupArn() *string
	// Experimental.
	DeploymentGroupName() *string
	// Experimental.
	Role() awsiam.IRole
}

// The jsii proxy for IServerDeploymentGroup
type jsiiProxy_IServerDeploymentGroup struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IServerDeploymentGroup) Application() IServerApplication {
	var returns IServerApplication
	_jsii_.Get(
		j,
		"application",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServerDeploymentGroup) AutoScalingGroups() *[]awsautoscaling.IAutoScalingGroup {
	var returns *[]awsautoscaling.IAutoScalingGroup
	_jsii_.Get(
		j,
		"autoScalingGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServerDeploymentGroup) DeploymentConfig() IServerDeploymentConfig {
	var returns IServerDeploymentConfig
	_jsii_.Get(
		j,
		"deploymentConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServerDeploymentGroup) DeploymentGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServerDeploymentGroup) DeploymentGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServerDeploymentGroup) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

// Represents a set of instance tag groups.
//
// An instance will match a set if it matches all of the groups in the set -
// in other words, sets follow 'and' semantics.
// You can have a maximum of 3 tag groups inside a set.
// Experimental.
type InstanceTagSet interface {
	InstanceTagGroups() *[]*map[string]*[]*string
}

// The jsii proxy struct for InstanceTagSet
type jsiiProxy_InstanceTagSet struct {
	_ byte // padding
}

func (j *jsiiProxy_InstanceTagSet) InstanceTagGroups() *[]*map[string]*[]*string {
	var returns *[]*map[string]*[]*string
	_jsii_.Get(
		j,
		"instanceTagGroups",
		&returns,
	)
	return returns
}


// Experimental.
func NewInstanceTagSet(instanceTagGroups ...*map[string]*[]*string) InstanceTagSet {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range instanceTagGroups {
		args = append(args, a)
	}

	j := jsiiProxy_InstanceTagSet{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codedeploy.InstanceTagSet",
		args,
		&j,
	)

	return &j
}

// Experimental.
func NewInstanceTagSet_Override(i InstanceTagSet, instanceTagGroups ...*map[string]*[]*string) {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range instanceTagGroups {
		args = append(args, a)
	}

	_jsii_.Create(
		"aws-cdk-lib.aws_codedeploy.InstanceTagSet",
		args,
		i,
	)
}

// A CodeDeploy Application that deploys to an AWS Lambda function.
// Experimental.
type LambdaApplication interface {
	awscdk.Resource
	ILambdaApplication
	ApplicationArn() *string
	ApplicationName() *string
	Env() *awscdk.ResourceEnvironment
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for LambdaApplication
type jsiiProxy_LambdaApplication struct {
	internal.Type__awscdkResource
	jsiiProxy_ILambdaApplication
}

func (j *jsiiProxy_LambdaApplication) ApplicationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaApplication) ApplicationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaApplication) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaApplication) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaApplication) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaApplication) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewLambdaApplication(scope constructs.Construct, id *string, props *LambdaApplicationProps) LambdaApplication {
	_init_.Initialize()

	j := jsiiProxy_LambdaApplication{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codedeploy.LambdaApplication",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewLambdaApplication_Override(l LambdaApplication, scope constructs.Construct, id *string, props *LambdaApplicationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codedeploy.LambdaApplication",
		[]interface{}{scope, id, props},
		l,
	)
}

// Import an Application defined either outside the CDK, or in a different CDK Stack.
//
// Returns: a Construct representing a reference to an existing Application
// Experimental.
func LambdaApplication_FromLambdaApplicationName(scope constructs.Construct, id *string, lambdaApplicationName *string) ILambdaApplication {
	_init_.Initialize()

	var returns ILambdaApplication

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codedeploy.LambdaApplication",
		"fromLambdaApplicationName",
		[]interface{}{scope, id, lambdaApplicationName},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func LambdaApplication_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codedeploy.LambdaApplication",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func LambdaApplication_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codedeploy.LambdaApplication",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

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
func (l *jsiiProxy_LambdaApplication) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		l,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (l *jsiiProxy_LambdaApplication) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
func (l *jsiiProxy_LambdaApplication) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
func (l *jsiiProxy_LambdaApplication) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (l *jsiiProxy_LambdaApplication) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Construction properties for {@link LambdaApplication}.
// Experimental.
type LambdaApplicationProps struct {
	// The physical, human-readable name of the CodeDeploy Application.
	// Experimental.
	ApplicationName *string `json:"applicationName"`
}

// A custom Deployment Configuration for a Lambda Deployment Group.
//
// Note: This class currently stands as namespaced container of the default configurations
// until CloudFormation supports custom Lambda Deployment Configs. Until then it is closed
// (private constructor) and does not extend {@link Construct}
// Experimental.
type LambdaDeploymentConfig interface {
}

// The jsii proxy struct for LambdaDeploymentConfig
type jsiiProxy_LambdaDeploymentConfig struct {
	_ byte // padding
}

// Import a custom Deployment Configuration for a Lambda Deployment Group defined outside the CDK.
//
// Returns: a Construct representing a reference to an existing custom Deployment Configuration
// Experimental.
func LambdaDeploymentConfig_Import(_scope constructs.Construct, _id *string, props *LambdaDeploymentConfigImportProps) ILambdaDeploymentConfig {
	_init_.Initialize()

	var returns ILambdaDeploymentConfig

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codedeploy.LambdaDeploymentConfig",
		"import",
		[]interface{}{_scope, _id, props},
		&returns,
	)

	return returns
}

func LambdaDeploymentConfig_ALL_AT_ONCE() ILambdaDeploymentConfig {
	_init_.Initialize()
	var returns ILambdaDeploymentConfig
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codedeploy.LambdaDeploymentConfig",
		"ALL_AT_ONCE",
		&returns,
	)
	return returns
}

func LambdaDeploymentConfig_CANARY_10PERCENT_10MINUTES() ILambdaDeploymentConfig {
	_init_.Initialize()
	var returns ILambdaDeploymentConfig
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codedeploy.LambdaDeploymentConfig",
		"CANARY_10PERCENT_10MINUTES",
		&returns,
	)
	return returns
}

func LambdaDeploymentConfig_CANARY_10PERCENT_15MINUTES() ILambdaDeploymentConfig {
	_init_.Initialize()
	var returns ILambdaDeploymentConfig
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codedeploy.LambdaDeploymentConfig",
		"CANARY_10PERCENT_15MINUTES",
		&returns,
	)
	return returns
}

func LambdaDeploymentConfig_CANARY_10PERCENT_30MINUTES() ILambdaDeploymentConfig {
	_init_.Initialize()
	var returns ILambdaDeploymentConfig
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codedeploy.LambdaDeploymentConfig",
		"CANARY_10PERCENT_30MINUTES",
		&returns,
	)
	return returns
}

func LambdaDeploymentConfig_CANARY_10PERCENT_5MINUTES() ILambdaDeploymentConfig {
	_init_.Initialize()
	var returns ILambdaDeploymentConfig
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codedeploy.LambdaDeploymentConfig",
		"CANARY_10PERCENT_5MINUTES",
		&returns,
	)
	return returns
}

func LambdaDeploymentConfig_LINEAR_10PERCENT_EVERY_10MINUTES() ILambdaDeploymentConfig {
	_init_.Initialize()
	var returns ILambdaDeploymentConfig
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codedeploy.LambdaDeploymentConfig",
		"LINEAR_10PERCENT_EVERY_10MINUTES",
		&returns,
	)
	return returns
}

func LambdaDeploymentConfig_LINEAR_10PERCENT_EVERY_1MINUTE() ILambdaDeploymentConfig {
	_init_.Initialize()
	var returns ILambdaDeploymentConfig
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codedeploy.LambdaDeploymentConfig",
		"LINEAR_10PERCENT_EVERY_1MINUTE",
		&returns,
	)
	return returns
}

func LambdaDeploymentConfig_LINEAR_10PERCENT_EVERY_2MINUTES() ILambdaDeploymentConfig {
	_init_.Initialize()
	var returns ILambdaDeploymentConfig
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codedeploy.LambdaDeploymentConfig",
		"LINEAR_10PERCENT_EVERY_2MINUTES",
		&returns,
	)
	return returns
}

func LambdaDeploymentConfig_LINEAR_10PERCENT_EVERY_3MINUTES() ILambdaDeploymentConfig {
	_init_.Initialize()
	var returns ILambdaDeploymentConfig
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codedeploy.LambdaDeploymentConfig",
		"LINEAR_10PERCENT_EVERY_3MINUTES",
		&returns,
	)
	return returns
}

// Properties of a reference to a CodeDeploy Lambda Deployment Configuration.
// See: LambdaDeploymentConfig#import
//
// Experimental.
type LambdaDeploymentConfigImportProps struct {
	// The physical, human-readable name of the custom CodeDeploy Lambda Deployment Configuration that we are referencing.
	// Experimental.
	DeploymentConfigName *string `json:"deploymentConfigName"`
}

// Experimental.
type LambdaDeploymentGroup interface {
	awscdk.Resource
	ILambdaDeploymentGroup
	Application() ILambdaApplication
	DeploymentConfig() ILambdaDeploymentConfig
	DeploymentGroupArn() *string
	DeploymentGroupName() *string
	Env() *awscdk.ResourceEnvironment
	Node() constructs.Node
	PhysicalName() *string
	Role() awsiam.IRole
	Stack() awscdk.Stack
	AddAlarm(alarm awscloudwatch.IAlarm)
	AddPostHook(postHook awslambda.IFunction)
	AddPreHook(preHook awslambda.IFunction)
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	GrantPutLifecycleEventHookExecutionStatus(grantee awsiam.IGrantable) awsiam.Grant
	ToString() *string
}

// The jsii proxy struct for LambdaDeploymentGroup
type jsiiProxy_LambdaDeploymentGroup struct {
	internal.Type__awscdkResource
	jsiiProxy_ILambdaDeploymentGroup
}

func (j *jsiiProxy_LambdaDeploymentGroup) Application() ILambdaApplication {
	var returns ILambdaApplication
	_jsii_.Get(
		j,
		"application",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaDeploymentGroup) DeploymentConfig() ILambdaDeploymentConfig {
	var returns ILambdaDeploymentConfig
	_jsii_.Get(
		j,
		"deploymentConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaDeploymentGroup) DeploymentGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaDeploymentGroup) DeploymentGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaDeploymentGroup) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaDeploymentGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaDeploymentGroup) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaDeploymentGroup) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaDeploymentGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewLambdaDeploymentGroup(scope constructs.Construct, id *string, props *LambdaDeploymentGroupProps) LambdaDeploymentGroup {
	_init_.Initialize()

	j := jsiiProxy_LambdaDeploymentGroup{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codedeploy.LambdaDeploymentGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewLambdaDeploymentGroup_Override(l LambdaDeploymentGroup, scope constructs.Construct, id *string, props *LambdaDeploymentGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codedeploy.LambdaDeploymentGroup",
		[]interface{}{scope, id, props},
		l,
	)
}

// Import an Lambda Deployment Group defined either outside the CDK app, or in a different AWS region.
//
// Returns: a Construct representing a reference to an existing Deployment Group
// Experimental.
func LambdaDeploymentGroup_FromLambdaDeploymentGroupAttributes(scope constructs.Construct, id *string, attrs *LambdaDeploymentGroupAttributes) ILambdaDeploymentGroup {
	_init_.Initialize()

	var returns ILambdaDeploymentGroup

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codedeploy.LambdaDeploymentGroup",
		"fromLambdaDeploymentGroupAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func LambdaDeploymentGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codedeploy.LambdaDeploymentGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func LambdaDeploymentGroup_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codedeploy.LambdaDeploymentGroup",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Associates an additional alarm with this Deployment Group.
// Experimental.
func (l *jsiiProxy_LambdaDeploymentGroup) AddAlarm(alarm awscloudwatch.IAlarm) {
	_jsii_.InvokeVoid(
		l,
		"addAlarm",
		[]interface{}{alarm},
	)
}

// Associate a function to run after deployment completes.
// Experimental.
func (l *jsiiProxy_LambdaDeploymentGroup) AddPostHook(postHook awslambda.IFunction) {
	_jsii_.InvokeVoid(
		l,
		"addPostHook",
		[]interface{}{postHook},
	)
}

// Associate a function to run before deployment begins.
// Experimental.
func (l *jsiiProxy_LambdaDeploymentGroup) AddPreHook(preHook awslambda.IFunction) {
	_jsii_.InvokeVoid(
		l,
		"addPreHook",
		[]interface{}{preHook},
	)
}

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
func (l *jsiiProxy_LambdaDeploymentGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		l,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (l *jsiiProxy_LambdaDeploymentGroup) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
func (l *jsiiProxy_LambdaDeploymentGroup) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
func (l *jsiiProxy_LambdaDeploymentGroup) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Grant a principal permission to codedeploy:PutLifecycleEventHookExecutionStatus on this deployment group resource.
// Experimental.
func (l *jsiiProxy_LambdaDeploymentGroup) GrantPutLifecycleEventHookExecutionStatus(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		l,
		"grantPutLifecycleEventHookExecutionStatus",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (l *jsiiProxy_LambdaDeploymentGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties of a reference to a CodeDeploy Lambda Deployment Group.
// See: LambdaDeploymentGroup#fromLambdaDeploymentGroupAttributes
//
// Experimental.
type LambdaDeploymentGroupAttributes struct {
	// The reference to the CodeDeploy Lambda Application that this Deployment Group belongs to.
	// Experimental.
	Application ILambdaApplication `json:"application"`
	// The physical, human-readable name of the CodeDeploy Lambda Deployment Group that we are referencing.
	// Experimental.
	DeploymentGroupName *string `json:"deploymentGroupName"`
	// The Deployment Configuration this Deployment Group uses.
	// Experimental.
	DeploymentConfig ILambdaDeploymentConfig `json:"deploymentConfig"`
}

// Construction properties for {@link LambdaDeploymentGroup}.
// Experimental.
type LambdaDeploymentGroupProps struct {
	// Lambda Alias to shift traffic. Updating the version of the alias will trigger a CodeDeploy deployment.
	//
	// [disable-awslint:ref-via-interface] since we need to modify the alias CFN resource update policy
	// Experimental.
	Alias awslambda.Alias `json:"alias"`
	// The CloudWatch alarms associated with this Deployment Group.
	//
	// CodeDeploy will stop (and optionally roll back)
	// a deployment if during it any of the alarms trigger.
	//
	// Alarms can also be added after the Deployment Group is created using the {@link #addAlarm} method.
	// See: https://docs.aws.amazon.com/codedeploy/latest/userguide/monitoring-create-alarms.html
	//
	// Experimental.
	Alarms *[]awscloudwatch.IAlarm `json:"alarms"`
	// The reference to the CodeDeploy Lambda Application that this Deployment Group belongs to.
	// Experimental.
	Application ILambdaApplication `json:"application"`
	// The auto-rollback configuration for this Deployment Group.
	// Experimental.
	AutoRollback *AutoRollbackConfig `json:"autoRollback"`
	// The Deployment Configuration this Deployment Group uses.
	// Experimental.
	DeploymentConfig ILambdaDeploymentConfig `json:"deploymentConfig"`
	// The physical, human-readable name of the CodeDeploy Deployment Group.
	// Experimental.
	DeploymentGroupName *string `json:"deploymentGroupName"`
	// Whether to continue a deployment even if fetching the alarm status from CloudWatch failed.
	// Experimental.
	IgnorePollAlarmsFailure *bool `json:"ignorePollAlarmsFailure"`
	// The Lambda function to run after traffic routing starts.
	// Experimental.
	PostHook awslambda.IFunction `json:"postHook"`
	// The Lambda function to run before traffic routing starts.
	// Experimental.
	PreHook awslambda.IFunction `json:"preHook"`
	// The service Role of this Deployment Group.
	// Experimental.
	Role awsiam.IRole `json:"role"`
}

// An interface of an abstract load balancer, as needed by CodeDeploy.
//
// Create instances using the static factory methods:
// {@link #classic}, {@link #application} and {@link #network}.
// Experimental.
type LoadBalancer interface {
	Generation() LoadBalancerGeneration
	Name() *string
}

// The jsii proxy struct for LoadBalancer
type jsiiProxy_LoadBalancer struct {
	_ byte // padding
}

func (j *jsiiProxy_LoadBalancer) Generation() LoadBalancerGeneration {
	var returns LoadBalancerGeneration
	_jsii_.Get(
		j,
		"generation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Experimental.
func NewLoadBalancer_Override(l LoadBalancer) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codedeploy.LoadBalancer",
		nil, // no parameters
		l,
	)
}

// Creates a new CodeDeploy load balancer from an Application Load Balancer Target Group.
// Experimental.
func LoadBalancer_Application(albTargetGroup awselasticloadbalancingv2.ApplicationTargetGroup) LoadBalancer {
	_init_.Initialize()

	var returns LoadBalancer

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codedeploy.LoadBalancer",
		"application",
		[]interface{}{albTargetGroup},
		&returns,
	)

	return returns
}

// Creates a new CodeDeploy load balancer from a Classic ELB Load Balancer.
// Experimental.
func LoadBalancer_Classic(loadBalancer awselasticloadbalancing.LoadBalancer) LoadBalancer {
	_init_.Initialize()

	var returns LoadBalancer

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codedeploy.LoadBalancer",
		"classic",
		[]interface{}{loadBalancer},
		&returns,
	)

	return returns
}

// Creates a new CodeDeploy load balancer from a Network Load Balancer Target Group.
// Experimental.
func LoadBalancer_Network(nlbTargetGroup awselasticloadbalancingv2.NetworkTargetGroup) LoadBalancer {
	_init_.Initialize()

	var returns LoadBalancer

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codedeploy.LoadBalancer",
		"network",
		[]interface{}{nlbTargetGroup},
		&returns,
	)

	return returns
}

// The generations of AWS load balancing solutions.
// Experimental.
type LoadBalancerGeneration string

const (
	LoadBalancerGeneration_FIRST LoadBalancerGeneration = "FIRST"
	LoadBalancerGeneration_SECOND LoadBalancerGeneration = "SECOND"
)

// Minimum number of healthy hosts for a server deployment.
// Experimental.
type MinimumHealthyHosts interface {
}

// The jsii proxy struct for MinimumHealthyHosts
type jsiiProxy_MinimumHealthyHosts struct {
	_ byte // padding
}

// The minimum healhty hosts threshold expressed as an absolute number.
// Experimental.
func MinimumHealthyHosts_Count(value *float64) MinimumHealthyHosts {
	_init_.Initialize()

	var returns MinimumHealthyHosts

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codedeploy.MinimumHealthyHosts",
		"count",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// The minmum healhty hosts threshold expressed as a percentage of the fleet.
// Experimental.
func MinimumHealthyHosts_Percentage(value *float64) MinimumHealthyHosts {
	_init_.Initialize()

	var returns MinimumHealthyHosts

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codedeploy.MinimumHealthyHosts",
		"percentage",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// A CodeDeploy Application that deploys to EC2/on-premise instances.
// Experimental.
type ServerApplication interface {
	awscdk.Resource
	IServerApplication
	ApplicationArn() *string
	ApplicationName() *string
	Env() *awscdk.ResourceEnvironment
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for ServerApplication
type jsiiProxy_ServerApplication struct {
	internal.Type__awscdkResource
	jsiiProxy_IServerApplication
}

func (j *jsiiProxy_ServerApplication) ApplicationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerApplication) ApplicationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerApplication) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerApplication) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerApplication) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerApplication) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewServerApplication(scope constructs.Construct, id *string, props *ServerApplicationProps) ServerApplication {
	_init_.Initialize()

	j := jsiiProxy_ServerApplication{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codedeploy.ServerApplication",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewServerApplication_Override(s ServerApplication, scope constructs.Construct, id *string, props *ServerApplicationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codedeploy.ServerApplication",
		[]interface{}{scope, id, props},
		s,
	)
}

// Import an Application defined either outside the CDK app, or in a different region.
//
// Returns: a Construct representing a reference to an existing Application
// Experimental.
func ServerApplication_FromServerApplicationName(scope constructs.Construct, id *string, serverApplicationName *string) IServerApplication {
	_init_.Initialize()

	var returns IServerApplication

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codedeploy.ServerApplication",
		"fromServerApplicationName",
		[]interface{}{scope, id, serverApplicationName},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ServerApplication_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codedeploy.ServerApplication",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func ServerApplication_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codedeploy.ServerApplication",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

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
func (s *jsiiProxy_ServerApplication) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		s,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (s *jsiiProxy_ServerApplication) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
func (s *jsiiProxy_ServerApplication) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
func (s *jsiiProxy_ServerApplication) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (s *jsiiProxy_ServerApplication) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Construction properties for {@link ServerApplication}.
// Experimental.
type ServerApplicationProps struct {
	// The physical, human-readable name of the CodeDeploy Application.
	// Experimental.
	ApplicationName *string `json:"applicationName"`
}

// A custom Deployment Configuration for an EC2/on-premise Deployment Group.
// Experimental.
type ServerDeploymentConfig interface {
	awscdk.Resource
	IServerDeploymentConfig
	DeploymentConfigArn() *string
	DeploymentConfigName() *string
	Env() *awscdk.ResourceEnvironment
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for ServerDeploymentConfig
type jsiiProxy_ServerDeploymentConfig struct {
	internal.Type__awscdkResource
	jsiiProxy_IServerDeploymentConfig
}

func (j *jsiiProxy_ServerDeploymentConfig) DeploymentConfigArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentConfigArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerDeploymentConfig) DeploymentConfigName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentConfigName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerDeploymentConfig) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerDeploymentConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerDeploymentConfig) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerDeploymentConfig) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewServerDeploymentConfig(scope constructs.Construct, id *string, props *ServerDeploymentConfigProps) ServerDeploymentConfig {
	_init_.Initialize()

	j := jsiiProxy_ServerDeploymentConfig{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codedeploy.ServerDeploymentConfig",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewServerDeploymentConfig_Override(s ServerDeploymentConfig, scope constructs.Construct, id *string, props *ServerDeploymentConfigProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codedeploy.ServerDeploymentConfig",
		[]interface{}{scope, id, props},
		s,
	)
}

// Import a custom Deployment Configuration for an EC2/on-premise Deployment Group defined either outside the CDK app, or in a different region.
//
// Returns: a Construct representing a reference to an existing custom Deployment Configuration
// Experimental.
func ServerDeploymentConfig_FromServerDeploymentConfigName(scope constructs.Construct, id *string, serverDeploymentConfigName *string) IServerDeploymentConfig {
	_init_.Initialize()

	var returns IServerDeploymentConfig

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codedeploy.ServerDeploymentConfig",
		"fromServerDeploymentConfigName",
		[]interface{}{scope, id, serverDeploymentConfigName},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ServerDeploymentConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codedeploy.ServerDeploymentConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func ServerDeploymentConfig_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codedeploy.ServerDeploymentConfig",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func ServerDeploymentConfig_ALL_AT_ONCE() IServerDeploymentConfig {
	_init_.Initialize()
	var returns IServerDeploymentConfig
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codedeploy.ServerDeploymentConfig",
		"ALL_AT_ONCE",
		&returns,
	)
	return returns
}

func ServerDeploymentConfig_HALF_AT_A_TIME() IServerDeploymentConfig {
	_init_.Initialize()
	var returns IServerDeploymentConfig
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codedeploy.ServerDeploymentConfig",
		"HALF_AT_A_TIME",
		&returns,
	)
	return returns
}

func ServerDeploymentConfig_ONE_AT_A_TIME() IServerDeploymentConfig {
	_init_.Initialize()
	var returns IServerDeploymentConfig
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codedeploy.ServerDeploymentConfig",
		"ONE_AT_A_TIME",
		&returns,
	)
	return returns
}

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
func (s *jsiiProxy_ServerDeploymentConfig) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		s,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (s *jsiiProxy_ServerDeploymentConfig) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
func (s *jsiiProxy_ServerDeploymentConfig) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
func (s *jsiiProxy_ServerDeploymentConfig) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (s *jsiiProxy_ServerDeploymentConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Construction properties of {@link ServerDeploymentConfig}.
// Experimental.
type ServerDeploymentConfigProps struct {
	// Minimum number of healthy hosts.
	// Experimental.
	MinimumHealthyHosts MinimumHealthyHosts `json:"minimumHealthyHosts"`
	// The physical, human-readable name of the Deployment Configuration.
	// Experimental.
	DeploymentConfigName *string `json:"deploymentConfigName"`
}

// A CodeDeploy Deployment Group that deploys to EC2/on-premise instances.
// Experimental.
type ServerDeploymentGroup interface {
	awscdk.Resource
	IServerDeploymentGroup
	Application() IServerApplication
	AutoScalingGroups() *[]awsautoscaling.IAutoScalingGroup
	DeploymentConfig() IServerDeploymentConfig
	DeploymentGroupArn() *string
	DeploymentGroupName() *string
	Env() *awscdk.ResourceEnvironment
	Node() constructs.Node
	PhysicalName() *string
	Role() awsiam.IRole
	Stack() awscdk.Stack
	AddAlarm(alarm awscloudwatch.IAlarm)
	AddAutoScalingGroup(asg awsautoscaling.AutoScalingGroup)
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for ServerDeploymentGroup
type jsiiProxy_ServerDeploymentGroup struct {
	internal.Type__awscdkResource
	jsiiProxy_IServerDeploymentGroup
}

func (j *jsiiProxy_ServerDeploymentGroup) Application() IServerApplication {
	var returns IServerApplication
	_jsii_.Get(
		j,
		"application",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerDeploymentGroup) AutoScalingGroups() *[]awsautoscaling.IAutoScalingGroup {
	var returns *[]awsautoscaling.IAutoScalingGroup
	_jsii_.Get(
		j,
		"autoScalingGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerDeploymentGroup) DeploymentConfig() IServerDeploymentConfig {
	var returns IServerDeploymentConfig
	_jsii_.Get(
		j,
		"deploymentConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerDeploymentGroup) DeploymentGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerDeploymentGroup) DeploymentGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerDeploymentGroup) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerDeploymentGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerDeploymentGroup) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerDeploymentGroup) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerDeploymentGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewServerDeploymentGroup(scope constructs.Construct, id *string, props *ServerDeploymentGroupProps) ServerDeploymentGroup {
	_init_.Initialize()

	j := jsiiProxy_ServerDeploymentGroup{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codedeploy.ServerDeploymentGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewServerDeploymentGroup_Override(s ServerDeploymentGroup, scope constructs.Construct, id *string, props *ServerDeploymentGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codedeploy.ServerDeploymentGroup",
		[]interface{}{scope, id, props},
		s,
	)
}

// Import an EC2/on-premise Deployment Group defined either outside the CDK app, or in a different region.
//
// Returns: a Construct representing a reference to an existing Deployment Group
// Experimental.
func ServerDeploymentGroup_FromServerDeploymentGroupAttributes(scope constructs.Construct, id *string, attrs *ServerDeploymentGroupAttributes) IServerDeploymentGroup {
	_init_.Initialize()

	var returns IServerDeploymentGroup

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codedeploy.ServerDeploymentGroup",
		"fromServerDeploymentGroupAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ServerDeploymentGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codedeploy.ServerDeploymentGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func ServerDeploymentGroup_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codedeploy.ServerDeploymentGroup",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Associates an additional alarm with this Deployment Group.
// Experimental.
func (s *jsiiProxy_ServerDeploymentGroup) AddAlarm(alarm awscloudwatch.IAlarm) {
	_jsii_.InvokeVoid(
		s,
		"addAlarm",
		[]interface{}{alarm},
	)
}

// Adds an additional auto-scaling group to this Deployment Group.
// Experimental.
func (s *jsiiProxy_ServerDeploymentGroup) AddAutoScalingGroup(asg awsautoscaling.AutoScalingGroup) {
	_jsii_.InvokeVoid(
		s,
		"addAutoScalingGroup",
		[]interface{}{asg},
	)
}

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
func (s *jsiiProxy_ServerDeploymentGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		s,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (s *jsiiProxy_ServerDeploymentGroup) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
func (s *jsiiProxy_ServerDeploymentGroup) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
func (s *jsiiProxy_ServerDeploymentGroup) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (s *jsiiProxy_ServerDeploymentGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties of a reference to a CodeDeploy EC2/on-premise Deployment Group.
// See: ServerDeploymentGroup#import
//
// Experimental.
type ServerDeploymentGroupAttributes struct {
	// The reference to the CodeDeploy EC2/on-premise Application that this Deployment Group belongs to.
	// Experimental.
	Application IServerApplication `json:"application"`
	// The physical, human-readable name of the CodeDeploy EC2/on-premise Deployment Group that we are referencing.
	// Experimental.
	DeploymentGroupName *string `json:"deploymentGroupName"`
	// The Deployment Configuration this Deployment Group uses.
	// Experimental.
	DeploymentConfig IServerDeploymentConfig `json:"deploymentConfig"`
}

// Construction properties for {@link ServerDeploymentGroup}.
// Experimental.
type ServerDeploymentGroupProps struct {
	// The CloudWatch alarms associated with this Deployment Group.
	//
	// CodeDeploy will stop (and optionally roll back)
	// a deployment if during it any of the alarms trigger.
	//
	// Alarms can also be added after the Deployment Group is created using the {@link #addAlarm} method.
	// See: https://docs.aws.amazon.com/codedeploy/latest/userguide/monitoring-create-alarms.html
	//
	// Experimental.
	Alarms *[]awscloudwatch.IAlarm `json:"alarms"`
	// The CodeDeploy EC2/on-premise Application this Deployment Group belongs to.
	// Experimental.
	Application IServerApplication `json:"application"`
	// The auto-rollback configuration for this Deployment Group.
	// Experimental.
	AutoRollback *AutoRollbackConfig `json:"autoRollback"`
	// The auto-scaling groups belonging to this Deployment Group.
	//
	// Auto-scaling groups can also be added after the Deployment Group is created
	// using the {@link #addAutoScalingGroup} method.
	//
	// [disable-awslint:ref-via-interface] is needed because we update userdata
	// for ASGs to install the codedeploy agent.
	// Experimental.
	AutoScalingGroups *[]awsautoscaling.IAutoScalingGroup `json:"autoScalingGroups"`
	// The EC2/on-premise Deployment Configuration to use for this Deployment Group.
	// Experimental.
	DeploymentConfig IServerDeploymentConfig `json:"deploymentConfig"`
	// The physical, human-readable name of the CodeDeploy Deployment Group.
	// Experimental.
	DeploymentGroupName *string `json:"deploymentGroupName"`
	// All EC2 instances matching the given set of tags when a deployment occurs will be added to this Deployment Group.
	// Experimental.
	Ec2InstanceTags InstanceTagSet `json:"ec2InstanceTags"`
	// Whether to continue a deployment even if fetching the alarm status from CloudWatch failed.
	// Experimental.
	IgnorePollAlarmsFailure *bool `json:"ignorePollAlarmsFailure"`
	// If you've provided any auto-scaling groups with the {@link #autoScalingGroups} property, you can set this property to add User Data that installs the CodeDeploy agent on the instances.
	// See: https://docs.aws.amazon.com/codedeploy/latest/userguide/codedeploy-agent-operations-install.html
	//
	// Experimental.
	InstallAgent *bool `json:"installAgent"`
	// The load balancer to place in front of this Deployment Group.
	//
	// Can be created from either a classic Elastic Load Balancer,
	// or an Application Load Balancer / Network Load Balancer Target Group.
	// Experimental.
	LoadBalancer LoadBalancer `json:"loadBalancer"`
	// All on-premise instances matching the given set of tags when a deployment occurs will be added to this Deployment Group.
	// Experimental.
	OnPremiseInstanceTags InstanceTagSet `json:"onPremiseInstanceTags"`
	// The service Role of this Deployment Group.
	// Experimental.
	Role awsiam.IRole `json:"role"`
}

