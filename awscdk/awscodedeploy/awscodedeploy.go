package awscodedeploy

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsautoscaling"
	"github.com/aws/aws-cdk-go/awscdk/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/awscodedeploy/internal"
	"github.com/aws/aws-cdk-go/awscdk/awselasticloadbalancing"
	"github.com/aws/aws-cdk-go/awscdk/awselasticloadbalancingv2"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/constructs-go/constructs/v3"
)

// The configuration for automatically rolling back deployments in a given Deployment Group.
//
// TODO: EXAMPLE
//
// Experimental.
type AutoRollbackConfig struct {
	// Whether to automatically roll back a deployment during which one of the configured CloudWatch alarms for this Deployment Group went off.
	// Experimental.
	DeploymentInAlarm *bool `json:"deploymentInAlarm" yaml:"deploymentInAlarm"`
	// Whether to automatically roll back a deployment that fails.
	// Experimental.
	FailedDeployment *bool `json:"failedDeployment" yaml:"failedDeployment"`
	// Whether to automatically roll back a deployment that was manually stopped.
	// Experimental.
	StoppedDeployment *bool `json:"stoppedDeployment" yaml:"stoppedDeployment"`
}

// A CloudFormation `AWS::CodeDeploy::Application`.
//
// The `AWS::CodeDeploy::Application` resource creates an AWS CodeDeploy application. In CodeDeploy , an application is a name that functions as a container to ensure that the correct combination of revision, deployment configuration, and deployment group are referenced during a deployment. You can use the `AWS::CodeDeploy::DeploymentGroup` resource to associate the application with a CodeDeploy deployment group. For more information, see [CodeDeploy Deployments](https://docs.aws.amazon.com/codedeploy/latest/userguide/deployment-steps.html) in the *AWS CodeDeploy User Guide* .
//
// TODO: EXAMPLE
//
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
	Node() awscdk.ConstructNode
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

func (j *jsiiProxy_CfnApplication) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
func NewCfnApplication(scope awscdk.Construct, id *string, props *CfnApplicationProps) CfnApplication {
	_init_.Initialize()

	j := jsiiProxy_CfnApplication{}

	_jsii_.Create(
		"monocdk.aws_codedeploy.CfnApplication",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CodeDeploy::Application`.
func NewCfnApplication_Override(c CfnApplication, scope awscdk.Construct, id *string, props *CfnApplicationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_codedeploy.CfnApplication",
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
		"monocdk.aws_codedeploy.CfnApplication",
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
		"monocdk.aws_codedeploy.CfnApplication",
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
		"monocdk.aws_codedeploy.CfnApplication",
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
		"monocdk.aws_codedeploy.CfnApplication",
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
//
// The `value` argument to `addOverride` will not be processed or translated
// in any way. Pass raw JSON values in here with the correct capitalization
// for CloudFormation. If you pass CDK classes or structs, they will be
// rendered with lowercased key names, and CloudFormation will reject the
// template.
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

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnApplication) OnPrepare() {
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
func (c *jsiiProxy_CfnApplication) OnSynthesize(session constructs.ISynthesisSession) {
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

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnApplication) OverrideLogicalId(newLogicalId *string) {
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

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnApplication) Synthesize(session awscdk.ISynthesisSession) {
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

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
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

// Experimental.
func (c *jsiiProxy_CfnApplication) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnApplication`.
//
// TODO: EXAMPLE
//
type CfnApplicationProps struct {
	// A name for the application.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the application name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
	//
	// > Updates to `ApplicationName` are not supported.
	ApplicationName *string `json:"applicationName" yaml:"applicationName"`
	// The compute platform that CodeDeploy deploys the application to.
	ComputePlatform *string `json:"computePlatform" yaml:"computePlatform"`
	// The metadata that you apply to CodeDeploy applications to help you organize and categorize them.
	//
	// Each tag consists of a key and an optional value, both of which you define.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::CodeDeploy::DeploymentConfig`.
//
// The `AWS::CodeDeploy::DeploymentConfig` resource creates a set of deployment rules, deployment success conditions, and deployment failure conditions that AWS CodeDeploy uses during a deployment. The deployment configuration specifies, through the use of a `MinimumHealthyHosts` value, the number or percentage of instances that must remain available at any time during a deployment.
//
// TODO: EXAMPLE
//
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
	Node() awscdk.ConstructNode
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

func (j *jsiiProxy_CfnDeploymentConfig) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
func NewCfnDeploymentConfig(scope awscdk.Construct, id *string, props *CfnDeploymentConfigProps) CfnDeploymentConfig {
	_init_.Initialize()

	j := jsiiProxy_CfnDeploymentConfig{}

	_jsii_.Create(
		"monocdk.aws_codedeploy.CfnDeploymentConfig",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CodeDeploy::DeploymentConfig`.
func NewCfnDeploymentConfig_Override(c CfnDeploymentConfig, scope awscdk.Construct, id *string, props *CfnDeploymentConfigProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_codedeploy.CfnDeploymentConfig",
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
		"monocdk.aws_codedeploy.CfnDeploymentConfig",
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
		"monocdk.aws_codedeploy.CfnDeploymentConfig",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnDeploymentConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_codedeploy.CfnDeploymentConfig",
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
		"monocdk.aws_codedeploy.CfnDeploymentConfig",
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
//
// The `value` argument to `addOverride` will not be processed or translated
// in any way. Pass raw JSON values in here with the correct capitalization
// for CloudFormation. If you pass CDK classes or structs, they will be
// rendered with lowercased key names, and CloudFormation will reject the
// template.
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

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnDeploymentConfig) OnPrepare() {
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
func (c *jsiiProxy_CfnDeploymentConfig) OnSynthesize(session constructs.ISynthesisSession) {
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
func (c *jsiiProxy_CfnDeploymentConfig) OnValidate() *[]*string {
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
func (c *jsiiProxy_CfnDeploymentConfig) OverrideLogicalId(newLogicalId *string) {
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
func (c *jsiiProxy_CfnDeploymentConfig) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
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

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnDeploymentConfig) Synthesize(session awscdk.ISynthesisSession) {
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

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnDeploymentConfig) Validate() *[]*string {
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
func (c *jsiiProxy_CfnDeploymentConfig) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// `MinimumHealthyHosts` is a property of the [DeploymentConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentconfig.html) resource that defines how many instances must remain healthy during an AWS CodeDeploy deployment.
//
// TODO: EXAMPLE
//
type CfnDeploymentConfig_MinimumHealthyHostsProperty struct {
	// The minimum healthy instance type:.
	//
	// - HOST_COUNT: The minimum number of healthy instance as an absolute value.
	// - FLEET_PERCENT: The minimum number of healthy instance as a percentage of the total number of instance in the deployment.
	//
	// In an example of nine instance, if a HOST_COUNT of six is specified, deploy to up to three instances at a time. The deployment is successful if six or more instances are deployed to successfully. Otherwise, the deployment fails. If a FLEET_PERCENT of 40 is specified, deploy to up to five instance at a time. The deployment is successful if four or more instance are deployed to successfully. Otherwise, the deployment fails.
	//
	// > In a call to `GetDeploymentConfig` , CodeDeployDefault.OneAtATime returns a minimum healthy instance type of MOST_CONCURRENCY and a value of 1. This means a deployment to only one instance at a time. (You cannot set the type to MOST_CONCURRENCY, only to HOST_COUNT or FLEET_PERCENT.) In addition, with CodeDeployDefault.OneAtATime, AWS CodeDeploy attempts to ensure that all instances but one are kept in a healthy state during the deployment. Although this allows one instance at a time to be taken offline for a new deployment, it also means that if the deployment to the last instance fails, the overall deployment is still successful.
	//
	// For more information, see [AWS CodeDeploy Instance Health](https://docs.aws.amazon.com//codedeploy/latest/userguide/instances-health.html) in the *AWS CodeDeploy User Guide* .
	Type *string `json:"type" yaml:"type"`
	// The minimum healthy instance value.
	Value *float64 `json:"value" yaml:"value"`
}

// A configuration that shifts traffic from one version of a Lambda function or Amazon ECS task set to another in two increments.
//
// The original and target Lambda function versions or ECS task sets are specified in the deployment's AppSpec file.
//
// TODO: EXAMPLE
//
type CfnDeploymentConfig_TimeBasedCanaryProperty struct {
	// The number of minutes between the first and second traffic shifts of a `TimeBasedCanary` deployment.
	CanaryInterval *float64 `json:"canaryInterval" yaml:"canaryInterval"`
	// The percentage of traffic to shift in the first increment of a `TimeBasedCanary` deployment.
	CanaryPercentage *float64 `json:"canaryPercentage" yaml:"canaryPercentage"`
}

// A configuration that shifts traffic from one version of a Lambda function or ECS task set to another in equal increments, with an equal number of minutes between each increment.
//
// The original and target Lambda function versions or ECS task sets are specified in the deployment's AppSpec file.
//
// TODO: EXAMPLE
//
type CfnDeploymentConfig_TimeBasedLinearProperty struct {
	// The number of minutes between each incremental traffic shift of a `TimeBasedLinear` deployment.
	LinearInterval *float64 `json:"linearInterval" yaml:"linearInterval"`
	// The percentage of traffic that is shifted at the start of each increment of a `TimeBasedLinear` deployment.
	LinearPercentage *float64 `json:"linearPercentage" yaml:"linearPercentage"`
}

// The configuration that specifies how traffic is shifted from one version of a Lambda function to another version during an AWS Lambda deployment, or from one Amazon ECS task set to another during an Amazon ECS deployment.
//
// TODO: EXAMPLE
//
type CfnDeploymentConfig_TrafficRoutingConfigProperty struct {
	// The type of traffic shifting ( `TimeBasedCanary` or `TimeBasedLinear` ) used by a deployment configuration.
	Type *string `json:"type" yaml:"type"`
	// A configuration that shifts traffic from one version of a Lambda function or ECS task set to another in two increments.
	//
	// The original and target Lambda function versions or ECS task sets are specified in the deployment's AppSpec file.
	TimeBasedCanary interface{} `json:"timeBasedCanary" yaml:"timeBasedCanary"`
	// A configuration that shifts traffic from one version of a Lambda function or Amazon ECS task set to another in equal increments, with an equal number of minutes between each increment.
	//
	// The original and target Lambda function versions or Amazon ECS task sets are specified in the deployment's AppSpec file.
	TimeBasedLinear interface{} `json:"timeBasedLinear" yaml:"timeBasedLinear"`
}

// Properties for defining a `CfnDeploymentConfig`.
//
// TODO: EXAMPLE
//
type CfnDeploymentConfigProps struct {
	// The destination platform type for the deployment ( `Lambda` , `Server` , or `ECS` ).
	ComputePlatform *string `json:"computePlatform" yaml:"computePlatform"`
	// A name for the deployment configuration.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the deployment configuration name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
	//
	// > If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	DeploymentConfigName *string `json:"deploymentConfigName" yaml:"deploymentConfigName"`
	// The minimum number of healthy instances that should be available at any time during the deployment.
	//
	// There are two parameters expected in the input: type and value.
	//
	// The type parameter takes either of the following values:
	//
	// - HOST_COUNT: The value parameter represents the minimum number of healthy instances as an absolute value.
	// - FLEET_PERCENT: The value parameter represents the minimum number of healthy instances as a percentage of the total number of instances in the deployment. If you specify FLEET_PERCENT, at the start of the deployment, AWS CodeDeploy converts the percentage to the equivalent number of instance and rounds up fractional instances.
	//
	// The value parameter takes an integer.
	//
	// For example, to set a minimum of 95% healthy instance, specify a type of FLEET_PERCENT and a value of 95.
	//
	// For more information about instance health, see [CodeDeploy Instance Health](https://docs.aws.amazon.com/codedeploy/latest/userguide/instances-health.html) in the AWS CodeDeploy User Guide.
	MinimumHealthyHosts interface{} `json:"minimumHealthyHosts" yaml:"minimumHealthyHosts"`
	// The configuration that specifies how the deployment traffic is routed.
	TrafficRoutingConfig interface{} `json:"trafficRoutingConfig" yaml:"trafficRoutingConfig"`
}

// A CloudFormation `AWS::CodeDeploy::DeploymentGroup`.
//
// The `AWS::CodeDeploy::DeploymentGroup` resource creates an AWS CodeDeploy deployment group that specifies which instances your application revisions are deployed to, along with other deployment options. For more information, see [CreateDeploymentGroup](https://docs.aws.amazon.com/codedeploy/latest/APIReference/API_CreateDeploymentGroup.html) in the *CodeDeploy API Reference* .
//
// > Amazon ECS blue/green deployments through CodeDeploy do not use the `AWS::CodeDeploy::DeploymentGroup` resource. To perform Amazon ECS blue/green deployments, use the `AWS::CodeDeploy::BlueGreen` hook. See [Perform Amazon ECS blue/green deployments through CodeDeploy using AWS CloudFormation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/blue-green.html) for more information.
//
// TODO: EXAMPLE
//
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
	Node() awscdk.ConstructNode
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

func (j *jsiiProxy_CfnDeploymentGroup) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
func NewCfnDeploymentGroup(scope awscdk.Construct, id *string, props *CfnDeploymentGroupProps) CfnDeploymentGroup {
	_init_.Initialize()

	j := jsiiProxy_CfnDeploymentGroup{}

	_jsii_.Create(
		"monocdk.aws_codedeploy.CfnDeploymentGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CodeDeploy::DeploymentGroup`.
func NewCfnDeploymentGroup_Override(c CfnDeploymentGroup, scope awscdk.Construct, id *string, props *CfnDeploymentGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_codedeploy.CfnDeploymentGroup",
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
		"monocdk.aws_codedeploy.CfnDeploymentGroup",
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
		"monocdk.aws_codedeploy.CfnDeploymentGroup",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnDeploymentGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_codedeploy.CfnDeploymentGroup",
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
		"monocdk.aws_codedeploy.CfnDeploymentGroup",
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
//
// The `value` argument to `addOverride` will not be processed or translated
// in any way. Pass raw JSON values in here with the correct capitalization
// for CloudFormation. If you pass CDK classes or structs, they will be
// rendered with lowercased key names, and CloudFormation will reject the
// template.
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

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnDeploymentGroup) OnPrepare() {
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
func (c *jsiiProxy_CfnDeploymentGroup) OnSynthesize(session constructs.ISynthesisSession) {
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
func (c *jsiiProxy_CfnDeploymentGroup) OnValidate() *[]*string {
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
func (c *jsiiProxy_CfnDeploymentGroup) OverrideLogicalId(newLogicalId *string) {
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
func (c *jsiiProxy_CfnDeploymentGroup) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
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

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnDeploymentGroup) Synthesize(session awscdk.ISynthesisSession) {
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

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnDeploymentGroup) Validate() *[]*string {
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
func (c *jsiiProxy_CfnDeploymentGroup) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The `AlarmConfiguration` property type configures CloudWatch alarms for an AWS CodeDeploy deployment group.
//
// `AlarmConfiguration` is a property of the [DeploymentGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html) resource.
//
// TODO: EXAMPLE
//
type CfnDeploymentGroup_AlarmConfigurationProperty struct {
	// A list of alarms configured for the deployment group.
	//
	// A maximum of 10 alarms can be added to a deployment group.
	Alarms interface{} `json:"alarms" yaml:"alarms"`
	// Indicates whether the alarm configuration is enabled.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
	// Indicates whether a deployment should continue if information about the current state of alarms cannot be retrieved from Amazon CloudWatch .
	//
	// The default value is `false` .
	//
	// - `true` : The deployment proceeds even if alarm status information can't be retrieved from CloudWatch .
	// - `false` : The deployment stops if alarm status information can't be retrieved from CloudWatch .
	IgnorePollAlarmFailure interface{} `json:"ignorePollAlarmFailure" yaml:"ignorePollAlarmFailure"`
}

// The `Alarm` property type specifies a CloudWatch alarm to use for an AWS CodeDeploy deployment group.
//
// The `Alarm` property of the [CodeDeploy DeploymentGroup AlarmConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-alarmconfiguration.html) property contains a list of `Alarm` property types.
//
// TODO: EXAMPLE
//
type CfnDeploymentGroup_AlarmProperty struct {
	// The name of the alarm.
	//
	// Maximum length is 255 characters. Each alarm name can be used only once in a list of alarms.
	Name *string `json:"name" yaml:"name"`
}

// The `AutoRollbackConfiguration` property type configures automatic rollback for an AWS CodeDeploy deployment group when a deployment is not completed successfully.
//
// For more information, see [Automatic Rollbacks](https://docs.aws.amazon.com/codedeploy/latest/userguide/deployments-rollback-and-redeploy.html#deployments-rollback-and-redeploy-automatic-rollbacks) in the *AWS CodeDeploy User Guide* .
//
// `AutoRollbackConfiguration` is a property of the [DeploymentGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html) resource.
//
// TODO: EXAMPLE
//
type CfnDeploymentGroup_AutoRollbackConfigurationProperty struct {
	// Indicates whether a defined automatic rollback configuration is currently enabled.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
	// The event type or types that trigger a rollback.
	//
	// Valid values are `DEPLOYMENT_FAILURE` , `DEPLOYMENT_STOP_ON_ALARM` , or `DEPLOYMENT_STOP_ON_REQUEST` .
	Events *[]*string `json:"events" yaml:"events"`
}

// Information about blue/green deployment options for a deployment group.
//
// TODO: EXAMPLE
//
type CfnDeploymentGroup_BlueGreenDeploymentConfigurationProperty struct {
	// Information about the action to take when newly provisioned instances are ready to receive traffic in a blue/green deployment.
	DeploymentReadyOption interface{} `json:"deploymentReadyOption" yaml:"deploymentReadyOption"`
	// Information about how instances are provisioned for a replacement environment in a blue/green deployment.
	GreenFleetProvisioningOption interface{} `json:"greenFleetProvisioningOption" yaml:"greenFleetProvisioningOption"`
	// Information about whether to terminate instances in the original fleet during a blue/green deployment.
	TerminateBlueInstancesOnDeploymentSuccess interface{} `json:"terminateBlueInstancesOnDeploymentSuccess" yaml:"terminateBlueInstancesOnDeploymentSuccess"`
}

// Information about whether instances in the original environment are terminated when a blue/green deployment is successful.
//
// `BlueInstanceTerminationOption` does not apply to Lambda deployments.
//
// TODO: EXAMPLE
//
type CfnDeploymentGroup_BlueInstanceTerminationOptionProperty struct {
	// The action to take on instances in the original environment after a successful blue/green deployment.
	//
	// - `TERMINATE` : Instances are terminated after a specified wait time.
	// - `KEEP_ALIVE` : Instances are left running after they are deregistered from the load balancer and removed from the deployment group.
	Action *string `json:"action" yaml:"action"`
	// For an Amazon EC2 deployment, the number of minutes to wait after a successful blue/green deployment before terminating instances from the original environment.
	//
	// For an Amazon ECS deployment, the number of minutes before deleting the original (blue) task set. During an Amazon ECS deployment, CodeDeploy shifts traffic from the original (blue) task set to a replacement (green) task set.
	//
	// The maximum setting is 2880 minutes (2 days).
	TerminationWaitTimeInMinutes *float64 `json:"terminationWaitTimeInMinutes" yaml:"terminationWaitTimeInMinutes"`
}

// `Deployment` is a property of the [DeploymentGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html) resource that specifies an AWS CodeDeploy application revision to be deployed to instances in the deployment group. If you specify an application revision, your target revision is deployed as soon as the provisioning process is complete.
//
// TODO: EXAMPLE
//
type CfnDeploymentGroup_DeploymentProperty struct {
	// Information about the location of stored application artifacts and the service from which to retrieve them.
	Revision interface{} `json:"revision" yaml:"revision"`
	// A comment about the deployment.
	Description *string `json:"description" yaml:"description"`
	// If true, then if an `ApplicationStop` , `BeforeBlockTraffic` , or `AfterBlockTraffic` deployment lifecycle event to an instance fails, then the deployment continues to the next deployment lifecycle event.
	//
	// For example, if `ApplicationStop` fails, the deployment continues with DownloadBundle. If `BeforeBlockTraffic` fails, the deployment continues with `BlockTraffic` . If `AfterBlockTraffic` fails, the deployment continues with `ApplicationStop` .
	//
	// If false or not specified, then if a lifecycle event fails during a deployment to an instance, that deployment fails. If deployment to that instance is part of an overall deployment and the number of healthy hosts is not less than the minimum number of healthy hosts, then a deployment to the next instance is attempted.
	//
	// During a deployment, the AWS CodeDeploy agent runs the scripts specified for `ApplicationStop` , `BeforeBlockTraffic` , and `AfterBlockTraffic` in the AppSpec file from the previous successful deployment. (All other scripts are run from the AppSpec file in the current deployment.) If one of these scripts contains an error and does not run successfully, the deployment can fail.
	//
	// If the cause of the failure is a script from the last successful deployment that will never run successfully, create a new deployment and use `ignoreApplicationStopFailures` to specify that the `ApplicationStop` , `BeforeBlockTraffic` , and `AfterBlockTraffic` failures should be ignored.
	IgnoreApplicationStopFailures interface{} `json:"ignoreApplicationStopFailures" yaml:"ignoreApplicationStopFailures"`
}

// Information about how traffic is rerouted to instances in a replacement environment in a blue/green deployment.
//
// TODO: EXAMPLE
//
type CfnDeploymentGroup_DeploymentReadyOptionProperty struct {
	// Information about when to reroute traffic from an original environment to a replacement environment in a blue/green deployment.
	//
	// - CONTINUE_DEPLOYMENT: Register new instances with the load balancer immediately after the new application revision is installed on the instances in the replacement environment.
	// - STOP_DEPLOYMENT: Do not register new instances with a load balancer unless traffic rerouting is started using [ContinueDeployment](https://docs.aws.amazon.com/codedeploy/latest/APIReference/API_ContinueDeployment.html) . If traffic rerouting is not started before the end of the specified wait period, the deployment status is changed to Stopped.
	ActionOnTimeout *string `json:"actionOnTimeout" yaml:"actionOnTimeout"`
	// The number of minutes to wait before the status of a blue/green deployment is changed to Stopped if rerouting is not started manually.
	//
	// Applies only to the `STOP_DEPLOYMENT` option for `actionOnTimeout` .
	WaitTimeInMinutes *float64 `json:"waitTimeInMinutes" yaml:"waitTimeInMinutes"`
}

// Information about the type of deployment, either in-place or blue/green, you want to run and whether to route deployment traffic behind a load balancer.
//
// TODO: EXAMPLE
//
type CfnDeploymentGroup_DeploymentStyleProperty struct {
	// Indicates whether to route deployment traffic behind a load balancer.
	//
	// > An Amazon EC2 Application Load Balancer or Network Load Balancer is required for an Amazon ECS deployment.
	DeploymentOption *string `json:"deploymentOption" yaml:"deploymentOption"`
	// Indicates whether to run an in-place or blue/green deployment.
	DeploymentType *string `json:"deploymentType" yaml:"deploymentType"`
}

// Information about an Amazon EC2 tag filter.
//
// For more information about using tags and tag groups to help manage your Amazon EC2 instances and on-premises instances, see [Tagging Instances for Deployment Groups in AWS CodeDeploy](https://docs.aws.amazon.com/codedeploy/latest/userguide/instances-tagging.html) in the *AWS CodeDeploy User Guide* .
//
// TODO: EXAMPLE
//
type CfnDeploymentGroup_EC2TagFilterProperty struct {
	// The tag filter key.
	Key *string `json:"key" yaml:"key"`
	// The tag filter type:.
	//
	// - `KEY_ONLY` : Key only.
	// - `VALUE_ONLY` : Value only.
	// - `KEY_AND_VALUE` : Key and value.
	Type *string `json:"type" yaml:"type"`
	// The tag filter value.
	Value *string `json:"value" yaml:"value"`
}

// The `EC2TagSet` property type specifies information about groups of tags applied to Amazon EC2 instances.
//
// The deployment group includes only Amazon EC2 instances identified by all the tag groups. Cannot be used in the same template as EC2TagFilters.
//
// For more information about using tags and tag groups to help manage your Amazon EC2 instances and on-premises instances, see [Tagging Instances for Deployment Groups in AWS CodeDeploy](https://docs.aws.amazon.com/codedeploy/latest/userguide/instances-tagging.html) in the *AWS CodeDeploy User Guide* .
//
// `EC2TagSet` is a property of the [DeploymentGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html) resource type.
//
// TODO: EXAMPLE
//
type CfnDeploymentGroup_EC2TagSetListObjectProperty struct {
	// A list that contains other lists of Amazon EC2 instance tag groups.
	//
	// For an instance to be included in the deployment group, it must be identified by all of the tag groups in the list.
	Ec2TagGroup interface{} `json:"ec2TagGroup" yaml:"ec2TagGroup"`
}

// The `EC2TagSet` property type specifies information about groups of tags applied to Amazon EC2 instances.
//
// The deployment group includes only Amazon EC2 instances identified by all the tag groups. `EC2TagSet` cannot be used in the same template as `EC2TagFilter` .
//
// For information about using tags and tag groups to help manage your Amazon EC2 instances and on-premises instances, see [Tagging Instances for Deployment Groups in AWS CodeDeploy](https://docs.aws.amazon.com/codedeploy/latest/userguide/instances-tagging.html) .
//
// TODO: EXAMPLE
//
type CfnDeploymentGroup_EC2TagSetProperty struct {
	// The Amazon EC2 tags that are already applied to Amazon EC2 instances that you want to include in the deployment group.
	//
	// CodeDeploy includes all Amazon EC2 instances identified by any of the tags you specify in this deployment group.
	//
	// Duplicates are not allowed.
	Ec2TagSetList interface{} `json:"ec2TagSetList" yaml:"ec2TagSetList"`
}

// Contains the service and cluster names used to identify an Amazon ECS deployment's target.
//
// TODO: EXAMPLE
//
type CfnDeploymentGroup_ECSServiceProperty struct {
	// The name of the cluster that the Amazon ECS service is associated with.
	ClusterName *string `json:"clusterName" yaml:"clusterName"`
	// The name of the target Amazon ECS service.
	ServiceName *string `json:"serviceName" yaml:"serviceName"`
}

// The `ELBInfo` property type specifies information about the Elastic Load Balancing load balancer used for an CodeDeploy deployment group.
//
// If you specify the `ELBInfo` property, the `DeploymentStyle.DeploymentOption` property must be set to `WITH_TRAFFIC_CONTROL` for AWS CodeDeploy to route your traffic using the specified load balancers.
//
// `ELBInfo` is a property of the [AWS CodeDeploy DeploymentGroup LoadBalancerInfo](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-loadbalancerinfo.html) property type.
//
// TODO: EXAMPLE
//
type CfnDeploymentGroup_ELBInfoProperty struct {
	// For blue/green deployments, the name of the load balancer that is used to route traffic from original instances to replacement instances in a blue/green deployment.
	//
	// For in-place deployments, the name of the load balancer that instances are deregistered from so they are not serving traffic during a deployment, and then re-registered with after the deployment is complete.
	//
	// > AWS CloudFormation supports blue/green deployments on AWS Lambda compute platforms only.
	Name *string `json:"name" yaml:"name"`
}

// `GitHubLocation` is a property of the [CodeDeploy DeploymentGroup Revision](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-deployment-revision.html) property that specifies the location of an application revision that is stored in GitHub.
//
// TODO: EXAMPLE
//
type CfnDeploymentGroup_GitHubLocationProperty struct {
	// The SHA1 commit ID of the GitHub commit that represents the bundled artifacts for the application revision.
	CommitId *string `json:"commitId" yaml:"commitId"`
	// The GitHub account and repository pair that stores a reference to the commit that represents the bundled artifacts for the application revision.
	//
	// Specify the value as `account/repository` .
	Repository *string `json:"repository" yaml:"repository"`
}

// Information about the instances that belong to the replacement environment in a blue/green deployment.
//
// TODO: EXAMPLE
//
type CfnDeploymentGroup_GreenFleetProvisioningOptionProperty struct {
	// The method used to add instances to a replacement environment.
	//
	// - `DISCOVER_EXISTING` : Use instances that already exist or will be created manually.
	// - `COPY_AUTO_SCALING_GROUP` : Use settings from a specified Auto Scaling group to define and create instances in a new Auto Scaling group.
	Action *string `json:"action" yaml:"action"`
}

// The `LoadBalancerInfo` property type specifies information about the load balancer or target group used for an AWS CodeDeploy deployment group.
//
// For more information, see [Integrating CodeDeploy with Elastic Load Balancing](https://docs.aws.amazon.com/codedeploy/latest/userguide/integrations-aws-elastic-load-balancing.html) in the *AWS CodeDeploy User Guide* .
//
// For AWS CloudFormation to use the properties specified in `LoadBalancerInfo` , the `DeploymentStyle.DeploymentOption` property must be set to `WITH_TRAFFIC_CONTROL` . If `DeploymentStyle.DeploymentOption` is not set to `WITH_TRAFFIC_CONTROL` , AWS CloudFormation ignores any settings specified in `LoadBalancerInfo` .
//
// > AWS CloudFormation supports blue/green deployments on the AWS Lambda compute platform only.
//
// `LoadBalancerInfo` is a property of the [DeploymentGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html) resource.
//
// TODO: EXAMPLE
//
type CfnDeploymentGroup_LoadBalancerInfoProperty struct {
	// An array that contains information about the load balancer to use for load balancing in a deployment.
	//
	// In Elastic Load Balancing, load balancers are used with Classic Load Balancers.
	//
	// > Adding more than one load balancer to the array is not supported.
	ElbInfoList interface{} `json:"elbInfoList" yaml:"elbInfoList"`
	// An array that contains information about the target group to use for load balancing in a deployment.
	//
	// In Elastic Load Balancing , target groups are used with Application Load Balancers .
	//
	// > Adding more than one target group to the array is not supported.
	TargetGroupInfoList interface{} `json:"targetGroupInfoList" yaml:"targetGroupInfoList"`
}

// The `OnPremisesTagSetListObject` property type specifies lists of on-premises instance tag groups.
//
// In order for an instance to be included in the deployment group, it must be identified by all the tag groups in the list.
//
// `OnPremisesTagSetListObject` is a property of the [CodeDeploy DeploymentGroup OnPremisesTagSet](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-onpremisestagset.html) property type.
//
// TODO: EXAMPLE
//
type CfnDeploymentGroup_OnPremisesTagSetListObjectProperty struct {
	// Information about groups of on-premises instance tags.
	OnPremisesTagGroup interface{} `json:"onPremisesTagGroup" yaml:"onPremisesTagGroup"`
}

// The `OnPremisesTagSet` property type specifies a list containing other lists of on-premises instance tag groups.
//
// In order for an instance to be included in the deployment group, it must be identified by all the tag groups in the list.
//
// For more information about using tags and tag groups to help manage your Amazon EC2 instances and on-premises instances, see [Tagging Instances for Deployment Groups in AWS CodeDeploy](https://docs.aws.amazon.com/codedeploy/latest/userguide/instances-tagging.html) in the *AWS CodeDeploy User Guide* .
//
// `OnPremisesTagSet` is a property of the [DeploymentGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html) resource.
//
// TODO: EXAMPLE
//
type CfnDeploymentGroup_OnPremisesTagSetProperty struct {
	// A list that contains other lists of on-premises instance tag groups.
	//
	// For an instance to be included in the deployment group, it must be identified by all of the tag groups in the list.
	//
	// Duplicates are not allowed.
	OnPremisesTagSetList interface{} `json:"onPremisesTagSetList" yaml:"onPremisesTagSetList"`
}

// `RevisionLocation` is a property that defines the location of the CodeDeploy application revision to deploy.
//
// TODO: EXAMPLE
//
type CfnDeploymentGroup_RevisionLocationProperty struct {
	// Information about the location of application artifacts stored in GitHub.
	GitHubLocation interface{} `json:"gitHubLocation" yaml:"gitHubLocation"`
	// The type of application revision:.
	//
	// - S3: An application revision stored in Amazon S3.
	// - GitHub: An application revision stored in GitHub (EC2/On-premises deployments only).
	// - String: A YAML-formatted or JSON-formatted string ( AWS Lambda deployments only).
	// - AppSpecContent: An `AppSpecContent` object that contains the contents of an AppSpec file for an AWS Lambda or Amazon ECS deployment. The content is formatted as JSON or YAML stored as a RawString.
	RevisionType *string `json:"revisionType" yaml:"revisionType"`
	// Information about the location of a revision stored in Amazon S3.
	S3Location interface{} `json:"s3Location" yaml:"s3Location"`
}

// `S3Location` is a property of the [CodeDeploy DeploymentGroup Revision](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-deployment-revision.html) property that specifies the location of an application revision that is stored in Amazon Simple Storage Service ( Amazon S3 ).
//
// TODO: EXAMPLE
//
type CfnDeploymentGroup_S3LocationProperty struct {
	// The name of the Amazon S3 bucket where the application revision is stored.
	Bucket *string `json:"bucket" yaml:"bucket"`
	// The name of the Amazon S3 object that represents the bundled artifacts for the application revision.
	Key *string `json:"key" yaml:"key"`
	// The file type of the application revision. Must be one of the following:.
	//
	// - JSON
	// - tar: A tar archive file.
	// - tgz: A compressed tar archive file.
	// - YAML
	// - zip: A zip archive file.
	BundleType *string `json:"bundleType" yaml:"bundleType"`
	// The ETag of the Amazon S3 object that represents the bundled artifacts for the application revision.
	//
	// If the ETag is not specified as an input parameter, ETag validation of the object is skipped.
	ETag *string `json:"eTag" yaml:"eTag"`
	// A specific version of the Amazon S3 object that represents the bundled artifacts for the application revision.
	//
	// If the version is not specified, the system uses the most recent version by default.
	Version *string `json:"version" yaml:"version"`
}

// `TagFilter` is a property type of the [AWS::CodeDeploy::DeploymentGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html) resource that specifies which on-premises instances to associate with the deployment group. To register on-premise instances with AWS CodeDeploy , see [Configure Existing On-Premises Instances by Using AWS CodeDeploy](https://docs.aws.amazon.com/codedeploy/latest/userguide/instances-on-premises.html) in the *AWS CodeDeploy User Guide* .
//
// For more information about using tags and tag groups to help manage your Amazon EC2 instances and on-premises instances, see [Tagging Instances for Deployment Groups in AWS CodeDeploy](https://docs.aws.amazon.com/codedeploy/latest/userguide/instances-tagging.html) in the *AWS CodeDeploy User Guide* .
//
// TODO: EXAMPLE
//
type CfnDeploymentGroup_TagFilterProperty struct {
	// The on-premises instance tag filter key.
	Key *string `json:"key" yaml:"key"`
	// The on-premises instance tag filter type:.
	//
	// - KEY_ONLY: Key only.
	// - VALUE_ONLY: Value only.
	// - KEY_AND_VALUE: Key and value.
	Type *string `json:"type" yaml:"type"`
	// The on-premises instance tag filter value.
	Value *string `json:"value" yaml:"value"`
}

// The `TargetGroupInfo` property type specifies information about a target group in Elastic Load Balancing to use in a deployment.
//
// Instances are registered as targets in a target group, and traffic is routed to the target group. For more information, see [TargetGroupInfo](https://docs.aws.amazon.com/codedeploy/latest/APIReference/API_TargetGroupInfo.html) in the *AWS CodeDeploy API Reference*
//
// If you specify the `TargetGroupInfo` property, the `DeploymentStyle.DeploymentOption` property must be set to `WITH_TRAFFIC_CONTROL` for CodeDeploy to route your traffic using the specified target groups.
//
// `TargetGroupInfo` is a property of the [LoadBalancerInfo](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-loadbalancerinfo.html) property type.
//
// TODO: EXAMPLE
//
type CfnDeploymentGroup_TargetGroupInfoProperty struct {
	// For blue/green deployments, the name of the target group that instances in the original environment are deregistered from, and instances in the replacement environment registered with.
	//
	// For in-place deployments, the name of the target group that instances are deregistered from, so they are not serving traffic during a deployment, and then re-registered with after the deployment completes. No duplicates allowed.
	//
	// > AWS CloudFormation supports blue/green deployments on AWS Lambda compute platforms only.
	//
	// This value cannot exceed 32 characters, so you should use the `Name` property of the target group, or the `TargetGroupName` attribute with the `Fn::GetAtt` intrinsic function, as shown in the following example. Don't use the group's Amazon Resource Name (ARN) or `TargetGroupFullName` attribute.
	Name *string `json:"name" yaml:"name"`
}

// Information about notification triggers for the deployment group.
//
// TODO: EXAMPLE
//
type CfnDeploymentGroup_TriggerConfigProperty struct {
	// The event type or types that trigger notifications.
	TriggerEvents *[]*string `json:"triggerEvents" yaml:"triggerEvents"`
	// The name of the notification trigger.
	TriggerName *string `json:"triggerName" yaml:"triggerName"`
	// The Amazon Resource Name (ARN) of the Amazon Simple Notification Service topic through which notifications about deployment or instance events are sent.
	TriggerTargetArn *string `json:"triggerTargetArn" yaml:"triggerTargetArn"`
}

// Properties for defining a `CfnDeploymentGroup`.
//
// TODO: EXAMPLE
//
type CfnDeploymentGroupProps struct {
	// The name of an existing CodeDeploy application to associate this deployment group with.
	ApplicationName *string `json:"applicationName" yaml:"applicationName"`
	// A service role Amazon Resource Name (ARN) that grants CodeDeploy permission to make calls to AWS services on your behalf.
	//
	// For more information, see [Create a Service Role for AWS CodeDeploy](https://docs.aws.amazon.com/codedeploy/latest/userguide/getting-started-create-service-role.html) in the *AWS CodeDeploy User Guide* .
	//
	// > In some cases, you might need to add a dependency on the service role's policy. For more information, see IAM role policy in [DependsOn Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html) .
	ServiceRoleArn *string `json:"serviceRoleArn" yaml:"serviceRoleArn"`
	// Information about the Amazon CloudWatch alarms that are associated with the deployment group.
	AlarmConfiguration interface{} `json:"alarmConfiguration" yaml:"alarmConfiguration"`
	// Information about the automatic rollback configuration that is associated with the deployment group.
	//
	// If you specify this property, don't specify the `Deployment` property.
	AutoRollbackConfiguration interface{} `json:"autoRollbackConfiguration" yaml:"autoRollbackConfiguration"`
	// A list of associated Auto Scaling groups that CodeDeploy automatically deploys revisions to when new instances are created.
	//
	// Duplicates are not allowed.
	AutoScalingGroups *[]*string `json:"autoScalingGroups" yaml:"autoScalingGroups"`
	// Information about blue/green deployment options for a deployment group.
	BlueGreenDeploymentConfiguration interface{} `json:"blueGreenDeploymentConfiguration" yaml:"blueGreenDeploymentConfiguration"`
	// The application revision to deploy to this deployment group.
	//
	// If you specify this property, your target application revision is deployed as soon as the provisioning process is complete. If you specify this property, don't specify the `AutoRollbackConfiguration` property.
	Deployment interface{} `json:"deployment" yaml:"deployment"`
	// A deployment configuration name or a predefined configuration name.
	//
	// With predefined configurations, you can deploy application revisions to one instance at a time ( `CodeDeployDefault.OneAtATime` ), half of the instances at a time ( `CodeDeployDefault.HalfAtATime` ), or all the instances at once ( `CodeDeployDefault.AllAtOnce` ). For more information and valid values, see [Working with Deployment Configurations](https://docs.aws.amazon.com/codedeploy/latest/userguide/deployment-configurations.html) in the *AWS CodeDeploy User Guide* .
	DeploymentConfigName *string `json:"deploymentConfigName" yaml:"deploymentConfigName"`
	// A name for the deployment group.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the deployment group name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
	//
	// > If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	DeploymentGroupName *string `json:"deploymentGroupName" yaml:"deploymentGroupName"`
	// Attributes that determine the type of deployment to run and whether to route deployment traffic behind a load balancer.
	//
	// If you specify this property with a blue/green deployment type, don't specify the `AutoScalingGroups` , `LoadBalancerInfo` , or `Deployment` properties.
	//
	// > For blue/green deployments, AWS CloudFormation supports deployments on Lambda compute platforms only. You can perform Amazon ECS blue/green deployments using `AWS::CodeDeploy::BlueGreen` hook. See [Perform Amazon ECS blue/green deployments through CodeDeploy using AWS CloudFormation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/blue-green.html) for more information.
	DeploymentStyle interface{} `json:"deploymentStyle" yaml:"deploymentStyle"`
	// The Amazon EC2 tags that are already applied to Amazon EC2 instances that you want to include in the deployment group.
	//
	// CodeDeploy includes all Amazon EC2 instances identified by any of the tags you specify in this deployment group. Duplicates are not allowed.
	//
	// You can specify `EC2TagFilters` or `Ec2TagSet` , but not both.
	Ec2TagFilters interface{} `json:"ec2TagFilters" yaml:"ec2TagFilters"`
	// Information about groups of tags applied to Amazon EC2 instances.
	//
	// The deployment group includes only Amazon EC2 instances identified by all the tag groups. Cannot be used in the same call as `ec2TagFilter` .
	Ec2TagSet interface{} `json:"ec2TagSet" yaml:"ec2TagSet"`
	// The target Amazon ECS services in the deployment group.
	//
	// This applies only to deployment groups that use the Amazon ECS compute platform. A target Amazon ECS service is specified as an Amazon ECS cluster and service name pair using the format `<clustername>:<servicename>` .
	EcsServices interface{} `json:"ecsServices" yaml:"ecsServices"`
	// Information about the load balancer to use in a deployment.
	//
	// For more information, see [Integrating CodeDeploy with Elastic Load Balancing](https://docs.aws.amazon.com/codedeploy/latest/userguide/integrations-aws-elastic-load-balancing.html) in the *AWS CodeDeploy User Guide* .
	LoadBalancerInfo interface{} `json:"loadBalancerInfo" yaml:"loadBalancerInfo"`
	// The on-premises instance tags already applied to on-premises instances that you want to include in the deployment group.
	//
	// CodeDeploy includes all on-premises instances identified by any of the tags you specify in this deployment group. To register on-premises instances with CodeDeploy , see [Working with On-Premises Instances for CodeDeploy](https://docs.aws.amazon.com/codedeploy/latest/userguide/instances-on-premises.html) in the *AWS CodeDeploy User Guide* . Duplicates are not allowed.
	//
	// You can specify `OnPremisesInstanceTagFilters` or `OnPremisesInstanceTagSet` , but not both.
	OnPremisesInstanceTagFilters interface{} `json:"onPremisesInstanceTagFilters" yaml:"onPremisesInstanceTagFilters"`
	// Information about groups of tags applied to on-premises instances.
	//
	// The deployment group includes only on-premises instances identified by all the tag groups.
	//
	// You can specify `OnPremisesInstanceTagFilters` or `OnPremisesInstanceTagSet` , but not both.
	OnPremisesTagSet interface{} `json:"onPremisesTagSet" yaml:"onPremisesTagSet"`
	// Information about triggers associated with the deployment group.
	//
	// Duplicates are not allowed
	TriggerConfigurations interface{} `json:"triggerConfigurations" yaml:"triggerConfigurations"`
}

// A custom Deployment Configuration for a Lambda Deployment Group.
//
// TODO: EXAMPLE
//
// Experimental.
type CustomLambdaDeploymentConfig interface {
	awscdk.Resource
	ILambdaDeploymentConfig
	DeploymentConfigArn() *string
	DeploymentConfigName() *string
	Env() *awscdk.ResourceEnvironment
	Node() awscdk.ConstructNode
	PhysicalName() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
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

func (j *jsiiProxy_CustomLambdaDeploymentConfig) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
		"monocdk.aws_codedeploy.CustomLambdaDeploymentConfig",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewCustomLambdaDeploymentConfig_Override(c CustomLambdaDeploymentConfig, scope constructs.Construct, id *string, props *CustomLambdaDeploymentConfigProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_codedeploy.CustomLambdaDeploymentConfig",
		[]interface{}{scope, id, props},
		c,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func CustomLambdaDeploymentConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_codedeploy.CustomLambdaDeploymentConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func CustomLambdaDeploymentConfig_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_codedeploy.CustomLambdaDeploymentConfig",
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

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CustomLambdaDeploymentConfig) OnPrepare() {
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
func (c *jsiiProxy_CustomLambdaDeploymentConfig) OnSynthesize(session constructs.ISynthesisSession) {
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
func (c *jsiiProxy_CustomLambdaDeploymentConfig) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
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
func (c *jsiiProxy_CustomLambdaDeploymentConfig) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CustomLambdaDeploymentConfig) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
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

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CustomLambdaDeploymentConfig) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties of a reference to a CodeDeploy Lambda Deployment Configuration.
//
// TODO: EXAMPLE
//
// Experimental.
type CustomLambdaDeploymentConfigProps struct {
	// The interval, in number of minutes: - For LINEAR, how frequently additional traffic is shifted - For CANARY, how long to shift traffic before the full deployment.
	// Experimental.
	Interval awscdk.Duration `json:"interval" yaml:"interval"`
	// The integer percentage of traffic to shift: - For LINEAR, the percentage to shift every interval - For CANARY, the percentage to shift until the interval passes, before the full deployment.
	// Experimental.
	Percentage *float64 `json:"percentage" yaml:"percentage"`
	// The type of deployment config, either CANARY or LINEAR.
	// Experimental.
	Type CustomLambdaDeploymentConfigType `json:"type" yaml:"type"`
	// The verbatim name of the deployment config.
	//
	// Must be unique per account/region.
	// Other parameters cannot be updated if this name is provided.
	// Experimental.
	DeploymentConfigName *string `json:"deploymentConfigName" yaml:"deploymentConfigName"`
}

// Lambda Deployment config type.
//
// TODO: EXAMPLE
//
// Experimental.
type CustomLambdaDeploymentConfigType string

const (
	CustomLambdaDeploymentConfigType_CANARY CustomLambdaDeploymentConfigType = "CANARY"
	CustomLambdaDeploymentConfigType_LINEAR CustomLambdaDeploymentConfigType = "LINEAR"
)

// A CodeDeploy Application that deploys to an Amazon ECS service.
//
// TODO: EXAMPLE
//
// Experimental.
type EcsApplication interface {
	awscdk.Resource
	IEcsApplication
	ApplicationArn() *string
	ApplicationName() *string
	Env() *awscdk.ResourceEnvironment
	Node() awscdk.ConstructNode
	PhysicalName() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
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

func (j *jsiiProxy_EcsApplication) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
		"monocdk.aws_codedeploy.EcsApplication",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewEcsApplication_Override(e EcsApplication, scope constructs.Construct, id *string, props *EcsApplicationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_codedeploy.EcsApplication",
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
		"monocdk.aws_codedeploy.EcsApplication",
		"fromEcsApplicationName",
		[]interface{}{scope, id, ecsApplicationName},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func EcsApplication_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_codedeploy.EcsApplication",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func EcsApplication_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_codedeploy.EcsApplication",
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

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (e *jsiiProxy_EcsApplication) OnPrepare() {
	_jsii_.InvokeVoid(
		e,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (e *jsiiProxy_EcsApplication) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		e,
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
func (e *jsiiProxy_EcsApplication) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
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
func (e *jsiiProxy_EcsApplication) Prepare() {
	_jsii_.InvokeVoid(
		e,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (e *jsiiProxy_EcsApplication) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		e,
		"synthesize",
		[]interface{}{session},
	)
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

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (e *jsiiProxy_EcsApplication) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Construction properties for {@link EcsApplication}.
//
// TODO: EXAMPLE
//
// Experimental.
type EcsApplicationProps struct {
	// The physical, human-readable name of the CodeDeploy Application.
	// Experimental.
	ApplicationName *string `json:"applicationName" yaml:"applicationName"`
}

// A custom Deployment Configuration for an ECS Deployment Group.
//
// Note: This class currently stands as namespaced container of the default configurations
// until CloudFormation supports custom ECS Deployment Configs. Until then it is closed
// (private constructor) and does not extend {@link cdk.Construct}
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
		"monocdk.aws_codedeploy.EcsDeploymentConfig",
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
		"monocdk.aws_codedeploy.EcsDeploymentConfig",
		"ALL_AT_ONCE",
		&returns,
	)
	return returns
}

// Note: This class currently stands as a namespaced container for importing an ECS Deployment Group defined outside the CDK app until CloudFormation supports provisioning ECS Deployment Groups.
//
// Until then it is closed (private constructor) and does not
// extend {@link cdk.Construct}.
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
		"monocdk.aws_codedeploy.EcsDeploymentGroup",
		"fromEcsDeploymentGroupAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Properties of a reference to a CodeDeploy ECS Deployment Group.
//
// TODO: EXAMPLE
//
// See: EcsDeploymentGroup#fromEcsDeploymentGroupAttributes
//
// Experimental.
type EcsDeploymentGroupAttributes struct {
	// The reference to the CodeDeploy ECS Application that this Deployment Group belongs to.
	// Experimental.
	Application IEcsApplication `json:"application" yaml:"application"`
	// The physical, human-readable name of the CodeDeploy ECS Deployment Group that we are referencing.
	// Experimental.
	DeploymentGroupName *string `json:"deploymentGroupName" yaml:"deploymentGroupName"`
	// The Deployment Configuration this Deployment Group uses.
	// Experimental.
	DeploymentConfig IEcsDeploymentConfig `json:"deploymentConfig" yaml:"deploymentConfig"`
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
//
// TODO: EXAMPLE
//
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
		"monocdk.aws_codedeploy.InstanceTagSet",
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
		"monocdk.aws_codedeploy.InstanceTagSet",
		args,
		i,
	)
}

// A CodeDeploy Application that deploys to an AWS Lambda function.
//
// TODO: EXAMPLE
//
// Experimental.
type LambdaApplication interface {
	awscdk.Resource
	ILambdaApplication
	ApplicationArn() *string
	ApplicationName() *string
	Env() *awscdk.ResourceEnvironment
	Node() awscdk.ConstructNode
	PhysicalName() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
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

func (j *jsiiProxy_LambdaApplication) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
		"monocdk.aws_codedeploy.LambdaApplication",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewLambdaApplication_Override(l LambdaApplication, scope constructs.Construct, id *string, props *LambdaApplicationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_codedeploy.LambdaApplication",
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
		"monocdk.aws_codedeploy.LambdaApplication",
		"fromLambdaApplicationName",
		[]interface{}{scope, id, lambdaApplicationName},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func LambdaApplication_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_codedeploy.LambdaApplication",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func LambdaApplication_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_codedeploy.LambdaApplication",
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

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (l *jsiiProxy_LambdaApplication) OnPrepare() {
	_jsii_.InvokeVoid(
		l,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (l *jsiiProxy_LambdaApplication) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		l,
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
func (l *jsiiProxy_LambdaApplication) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
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
func (l *jsiiProxy_LambdaApplication) Prepare() {
	_jsii_.InvokeVoid(
		l,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (l *jsiiProxy_LambdaApplication) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		l,
		"synthesize",
		[]interface{}{session},
	)
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

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (l *jsiiProxy_LambdaApplication) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Construction properties for {@link LambdaApplication}.
//
// TODO: EXAMPLE
//
// Experimental.
type LambdaApplicationProps struct {
	// The physical, human-readable name of the CodeDeploy Application.
	// Experimental.
	ApplicationName *string `json:"applicationName" yaml:"applicationName"`
}

// A custom Deployment Configuration for a Lambda Deployment Group.
//
// Note: This class currently stands as namespaced container of the default configurations
// until CloudFormation supports custom Lambda Deployment Configs. Until then it is closed
// (private constructor) and does not extend {@link cdk.Construct}
//
// TODO: EXAMPLE
//
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
		"monocdk.aws_codedeploy.LambdaDeploymentConfig",
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
		"monocdk.aws_codedeploy.LambdaDeploymentConfig",
		"ALL_AT_ONCE",
		&returns,
	)
	return returns
}

func LambdaDeploymentConfig_CANARY_10PERCENT_10MINUTES() ILambdaDeploymentConfig {
	_init_.Initialize()
	var returns ILambdaDeploymentConfig
	_jsii_.StaticGet(
		"monocdk.aws_codedeploy.LambdaDeploymentConfig",
		"CANARY_10PERCENT_10MINUTES",
		&returns,
	)
	return returns
}

func LambdaDeploymentConfig_CANARY_10PERCENT_15MINUTES() ILambdaDeploymentConfig {
	_init_.Initialize()
	var returns ILambdaDeploymentConfig
	_jsii_.StaticGet(
		"monocdk.aws_codedeploy.LambdaDeploymentConfig",
		"CANARY_10PERCENT_15MINUTES",
		&returns,
	)
	return returns
}

func LambdaDeploymentConfig_CANARY_10PERCENT_30MINUTES() ILambdaDeploymentConfig {
	_init_.Initialize()
	var returns ILambdaDeploymentConfig
	_jsii_.StaticGet(
		"monocdk.aws_codedeploy.LambdaDeploymentConfig",
		"CANARY_10PERCENT_30MINUTES",
		&returns,
	)
	return returns
}

func LambdaDeploymentConfig_CANARY_10PERCENT_5MINUTES() ILambdaDeploymentConfig {
	_init_.Initialize()
	var returns ILambdaDeploymentConfig
	_jsii_.StaticGet(
		"monocdk.aws_codedeploy.LambdaDeploymentConfig",
		"CANARY_10PERCENT_5MINUTES",
		&returns,
	)
	return returns
}

func LambdaDeploymentConfig_LINEAR_10PERCENT_EVERY_10MINUTES() ILambdaDeploymentConfig {
	_init_.Initialize()
	var returns ILambdaDeploymentConfig
	_jsii_.StaticGet(
		"monocdk.aws_codedeploy.LambdaDeploymentConfig",
		"LINEAR_10PERCENT_EVERY_10MINUTES",
		&returns,
	)
	return returns
}

func LambdaDeploymentConfig_LINEAR_10PERCENT_EVERY_1MINUTE() ILambdaDeploymentConfig {
	_init_.Initialize()
	var returns ILambdaDeploymentConfig
	_jsii_.StaticGet(
		"monocdk.aws_codedeploy.LambdaDeploymentConfig",
		"LINEAR_10PERCENT_EVERY_1MINUTE",
		&returns,
	)
	return returns
}

func LambdaDeploymentConfig_LINEAR_10PERCENT_EVERY_2MINUTES() ILambdaDeploymentConfig {
	_init_.Initialize()
	var returns ILambdaDeploymentConfig
	_jsii_.StaticGet(
		"monocdk.aws_codedeploy.LambdaDeploymentConfig",
		"LINEAR_10PERCENT_EVERY_2MINUTES",
		&returns,
	)
	return returns
}

func LambdaDeploymentConfig_LINEAR_10PERCENT_EVERY_3MINUTES() ILambdaDeploymentConfig {
	_init_.Initialize()
	var returns ILambdaDeploymentConfig
	_jsii_.StaticGet(
		"monocdk.aws_codedeploy.LambdaDeploymentConfig",
		"LINEAR_10PERCENT_EVERY_3MINUTES",
		&returns,
	)
	return returns
}

// Properties of a reference to a CodeDeploy Lambda Deployment Configuration.
//
// TODO: EXAMPLE
//
// See: LambdaDeploymentConfig#import
//
// Experimental.
type LambdaDeploymentConfigImportProps struct {
	// The physical, human-readable name of the custom CodeDeploy Lambda Deployment Configuration that we are referencing.
	// Experimental.
	DeploymentConfigName *string `json:"deploymentConfigName" yaml:"deploymentConfigName"`
}

// TODO: EXAMPLE
//
// Experimental.
type LambdaDeploymentGroup interface {
	awscdk.Resource
	ILambdaDeploymentGroup
	Application() ILambdaApplication
	DeploymentConfig() ILambdaDeploymentConfig
	DeploymentGroupArn() *string
	DeploymentGroupName() *string
	Env() *awscdk.ResourceEnvironment
	Node() awscdk.ConstructNode
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
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
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

func (j *jsiiProxy_LambdaDeploymentGroup) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
		"monocdk.aws_codedeploy.LambdaDeploymentGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewLambdaDeploymentGroup_Override(l LambdaDeploymentGroup, scope constructs.Construct, id *string, props *LambdaDeploymentGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_codedeploy.LambdaDeploymentGroup",
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
		"monocdk.aws_codedeploy.LambdaDeploymentGroup",
		"fromLambdaDeploymentGroupAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func LambdaDeploymentGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_codedeploy.LambdaDeploymentGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func LambdaDeploymentGroup_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_codedeploy.LambdaDeploymentGroup",
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

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (l *jsiiProxy_LambdaDeploymentGroup) OnPrepare() {
	_jsii_.InvokeVoid(
		l,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (l *jsiiProxy_LambdaDeploymentGroup) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		l,
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
func (l *jsiiProxy_LambdaDeploymentGroup) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
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
func (l *jsiiProxy_LambdaDeploymentGroup) Prepare() {
	_jsii_.InvokeVoid(
		l,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (l *jsiiProxy_LambdaDeploymentGroup) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		l,
		"synthesize",
		[]interface{}{session},
	)
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

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (l *jsiiProxy_LambdaDeploymentGroup) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties of a reference to a CodeDeploy Lambda Deployment Group.
//
// TODO: EXAMPLE
//
// See: LambdaDeploymentGroup#fromLambdaDeploymentGroupAttributes
//
// Experimental.
type LambdaDeploymentGroupAttributes struct {
	// The reference to the CodeDeploy Lambda Application that this Deployment Group belongs to.
	// Experimental.
	Application ILambdaApplication `json:"application" yaml:"application"`
	// The physical, human-readable name of the CodeDeploy Lambda Deployment Group that we are referencing.
	// Experimental.
	DeploymentGroupName *string `json:"deploymentGroupName" yaml:"deploymentGroupName"`
	// The Deployment Configuration this Deployment Group uses.
	// Experimental.
	DeploymentConfig ILambdaDeploymentConfig `json:"deploymentConfig" yaml:"deploymentConfig"`
}

// Construction properties for {@link LambdaDeploymentGroup}.
//
// TODO: EXAMPLE
//
// Experimental.
type LambdaDeploymentGroupProps struct {
	// Lambda Alias to shift traffic. Updating the version of the alias will trigger a CodeDeploy deployment.
	//
	// [disable-awslint:ref-via-interface] since we need to modify the alias CFN resource update policy
	// Experimental.
	Alias awslambda.Alias `json:"alias" yaml:"alias"`
	// The CloudWatch alarms associated with this Deployment Group.
	//
	// CodeDeploy will stop (and optionally roll back)
	// a deployment if during it any of the alarms trigger.
	//
	// Alarms can also be added after the Deployment Group is created using the {@link #addAlarm} method.
	// See: https://docs.aws.amazon.com/codedeploy/latest/userguide/monitoring-create-alarms.html
	//
	// Experimental.
	Alarms *[]awscloudwatch.IAlarm `json:"alarms" yaml:"alarms"`
	// The reference to the CodeDeploy Lambda Application that this Deployment Group belongs to.
	// Experimental.
	Application ILambdaApplication `json:"application" yaml:"application"`
	// The auto-rollback configuration for this Deployment Group.
	// Experimental.
	AutoRollback *AutoRollbackConfig `json:"autoRollback" yaml:"autoRollback"`
	// The Deployment Configuration this Deployment Group uses.
	// Experimental.
	DeploymentConfig ILambdaDeploymentConfig `json:"deploymentConfig" yaml:"deploymentConfig"`
	// The physical, human-readable name of the CodeDeploy Deployment Group.
	// Experimental.
	DeploymentGroupName *string `json:"deploymentGroupName" yaml:"deploymentGroupName"`
	// Whether to continue a deployment even if fetching the alarm status from CloudWatch failed.
	// Experimental.
	IgnorePollAlarmsFailure *bool `json:"ignorePollAlarmsFailure" yaml:"ignorePollAlarmsFailure"`
	// The Lambda function to run after traffic routing starts.
	// Experimental.
	PostHook awslambda.IFunction `json:"postHook" yaml:"postHook"`
	// The Lambda function to run before traffic routing starts.
	// Experimental.
	PreHook awslambda.IFunction `json:"preHook" yaml:"preHook"`
	// The service Role of this Deployment Group.
	// Experimental.
	Role awsiam.IRole `json:"role" yaml:"role"`
}

// An interface of an abstract load balancer, as needed by CodeDeploy.
//
// Create instances using the static factory methods:
// {@link #classic}, {@link #application} and {@link #network}.
//
// TODO: EXAMPLE
//
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
		"monocdk.aws_codedeploy.LoadBalancer",
		nil, // no parameters
		l,
	)
}

// Creates a new CodeDeploy load balancer from an Application Load Balancer Target Group.
// Experimental.
func LoadBalancer_Application(albTargetGroup awselasticloadbalancingv2.IApplicationTargetGroup) LoadBalancer {
	_init_.Initialize()

	var returns LoadBalancer

	_jsii_.StaticInvoke(
		"monocdk.aws_codedeploy.LoadBalancer",
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
		"monocdk.aws_codedeploy.LoadBalancer",
		"classic",
		[]interface{}{loadBalancer},
		&returns,
	)

	return returns
}

// Creates a new CodeDeploy load balancer from a Network Load Balancer Target Group.
// Experimental.
func LoadBalancer_Network(nlbTargetGroup awselasticloadbalancingv2.INetworkTargetGroup) LoadBalancer {
	_init_.Initialize()

	var returns LoadBalancer

	_jsii_.StaticInvoke(
		"monocdk.aws_codedeploy.LoadBalancer",
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
//
// TODO: EXAMPLE
//
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
		"monocdk.aws_codedeploy.MinimumHealthyHosts",
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
		"monocdk.aws_codedeploy.MinimumHealthyHosts",
		"percentage",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// A CodeDeploy Application that deploys to EC2/on-premise instances.
//
// TODO: EXAMPLE
//
// Experimental.
type ServerApplication interface {
	awscdk.Resource
	IServerApplication
	ApplicationArn() *string
	ApplicationName() *string
	Env() *awscdk.ResourceEnvironment
	Node() awscdk.ConstructNode
	PhysicalName() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
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

func (j *jsiiProxy_ServerApplication) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
		"monocdk.aws_codedeploy.ServerApplication",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewServerApplication_Override(s ServerApplication, scope constructs.Construct, id *string, props *ServerApplicationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_codedeploy.ServerApplication",
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
		"monocdk.aws_codedeploy.ServerApplication",
		"fromServerApplicationName",
		[]interface{}{scope, id, serverApplicationName},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func ServerApplication_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_codedeploy.ServerApplication",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func ServerApplication_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_codedeploy.ServerApplication",
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

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (s *jsiiProxy_ServerApplication) OnPrepare() {
	_jsii_.InvokeVoid(
		s,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (s *jsiiProxy_ServerApplication) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
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
func (s *jsiiProxy_ServerApplication) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
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
func (s *jsiiProxy_ServerApplication) Prepare() {
	_jsii_.InvokeVoid(
		s,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (s *jsiiProxy_ServerApplication) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		[]interface{}{session},
	)
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

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (s *jsiiProxy_ServerApplication) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Construction properties for {@link ServerApplication}.
//
// TODO: EXAMPLE
//
// Experimental.
type ServerApplicationProps struct {
	// The physical, human-readable name of the CodeDeploy Application.
	// Experimental.
	ApplicationName *string `json:"applicationName" yaml:"applicationName"`
}

// A custom Deployment Configuration for an EC2/on-premise Deployment Group.
//
// TODO: EXAMPLE
//
// Experimental.
type ServerDeploymentConfig interface {
	awscdk.Resource
	IServerDeploymentConfig
	DeploymentConfigArn() *string
	DeploymentConfigName() *string
	Env() *awscdk.ResourceEnvironment
	Node() awscdk.ConstructNode
	PhysicalName() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
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

func (j *jsiiProxy_ServerDeploymentConfig) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
		"monocdk.aws_codedeploy.ServerDeploymentConfig",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewServerDeploymentConfig_Override(s ServerDeploymentConfig, scope constructs.Construct, id *string, props *ServerDeploymentConfigProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_codedeploy.ServerDeploymentConfig",
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
		"monocdk.aws_codedeploy.ServerDeploymentConfig",
		"fromServerDeploymentConfigName",
		[]interface{}{scope, id, serverDeploymentConfigName},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func ServerDeploymentConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_codedeploy.ServerDeploymentConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func ServerDeploymentConfig_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_codedeploy.ServerDeploymentConfig",
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
		"monocdk.aws_codedeploy.ServerDeploymentConfig",
		"ALL_AT_ONCE",
		&returns,
	)
	return returns
}

func ServerDeploymentConfig_HALF_AT_A_TIME() IServerDeploymentConfig {
	_init_.Initialize()
	var returns IServerDeploymentConfig
	_jsii_.StaticGet(
		"monocdk.aws_codedeploy.ServerDeploymentConfig",
		"HALF_AT_A_TIME",
		&returns,
	)
	return returns
}

func ServerDeploymentConfig_ONE_AT_A_TIME() IServerDeploymentConfig {
	_init_.Initialize()
	var returns IServerDeploymentConfig
	_jsii_.StaticGet(
		"monocdk.aws_codedeploy.ServerDeploymentConfig",
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

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (s *jsiiProxy_ServerDeploymentConfig) OnPrepare() {
	_jsii_.InvokeVoid(
		s,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (s *jsiiProxy_ServerDeploymentConfig) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
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
func (s *jsiiProxy_ServerDeploymentConfig) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
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
func (s *jsiiProxy_ServerDeploymentConfig) Prepare() {
	_jsii_.InvokeVoid(
		s,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (s *jsiiProxy_ServerDeploymentConfig) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		[]interface{}{session},
	)
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

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (s *jsiiProxy_ServerDeploymentConfig) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Construction properties of {@link ServerDeploymentConfig}.
//
// TODO: EXAMPLE
//
// Experimental.
type ServerDeploymentConfigProps struct {
	// Minimum number of healthy hosts.
	// Experimental.
	MinimumHealthyHosts MinimumHealthyHosts `json:"minimumHealthyHosts" yaml:"minimumHealthyHosts"`
	// The physical, human-readable name of the Deployment Configuration.
	// Experimental.
	DeploymentConfigName *string `json:"deploymentConfigName" yaml:"deploymentConfigName"`
}

// A CodeDeploy Deployment Group that deploys to EC2/on-premise instances.
//
// TODO: EXAMPLE
//
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
	Node() awscdk.ConstructNode
	PhysicalName() *string
	Role() awsiam.IRole
	Stack() awscdk.Stack
	AddAlarm(alarm awscloudwatch.IAlarm)
	AddAutoScalingGroup(asg awsautoscaling.AutoScalingGroup)
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
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

func (j *jsiiProxy_ServerDeploymentGroup) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
		"monocdk.aws_codedeploy.ServerDeploymentGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewServerDeploymentGroup_Override(s ServerDeploymentGroup, scope constructs.Construct, id *string, props *ServerDeploymentGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_codedeploy.ServerDeploymentGroup",
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
		"monocdk.aws_codedeploy.ServerDeploymentGroup",
		"fromServerDeploymentGroupAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func ServerDeploymentGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_codedeploy.ServerDeploymentGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func ServerDeploymentGroup_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_codedeploy.ServerDeploymentGroup",
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

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (s *jsiiProxy_ServerDeploymentGroup) OnPrepare() {
	_jsii_.InvokeVoid(
		s,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (s *jsiiProxy_ServerDeploymentGroup) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
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
func (s *jsiiProxy_ServerDeploymentGroup) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
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
func (s *jsiiProxy_ServerDeploymentGroup) Prepare() {
	_jsii_.InvokeVoid(
		s,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (s *jsiiProxy_ServerDeploymentGroup) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		[]interface{}{session},
	)
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

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (s *jsiiProxy_ServerDeploymentGroup) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties of a reference to a CodeDeploy EC2/on-premise Deployment Group.
//
// TODO: EXAMPLE
//
// See: ServerDeploymentGroup#import
//
// Experimental.
type ServerDeploymentGroupAttributes struct {
	// The reference to the CodeDeploy EC2/on-premise Application that this Deployment Group belongs to.
	// Experimental.
	Application IServerApplication `json:"application" yaml:"application"`
	// The physical, human-readable name of the CodeDeploy EC2/on-premise Deployment Group that we are referencing.
	// Experimental.
	DeploymentGroupName *string `json:"deploymentGroupName" yaml:"deploymentGroupName"`
	// The Deployment Configuration this Deployment Group uses.
	// Experimental.
	DeploymentConfig IServerDeploymentConfig `json:"deploymentConfig" yaml:"deploymentConfig"`
}

// Construction properties for {@link ServerDeploymentGroup}.
//
// TODO: EXAMPLE
//
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
	Alarms *[]awscloudwatch.IAlarm `json:"alarms" yaml:"alarms"`
	// The CodeDeploy EC2/on-premise Application this Deployment Group belongs to.
	// Experimental.
	Application IServerApplication `json:"application" yaml:"application"`
	// The auto-rollback configuration for this Deployment Group.
	// Experimental.
	AutoRollback *AutoRollbackConfig `json:"autoRollback" yaml:"autoRollback"`
	// The auto-scaling groups belonging to this Deployment Group.
	//
	// Auto-scaling groups can also be added after the Deployment Group is created
	// using the {@link #addAutoScalingGroup} method.
	//
	// [disable-awslint:ref-via-interface] is needed because we update userdata
	// for ASGs to install the codedeploy agent.
	// Experimental.
	AutoScalingGroups *[]awsautoscaling.IAutoScalingGroup `json:"autoScalingGroups" yaml:"autoScalingGroups"`
	// The EC2/on-premise Deployment Configuration to use for this Deployment Group.
	// Experimental.
	DeploymentConfig IServerDeploymentConfig `json:"deploymentConfig" yaml:"deploymentConfig"`
	// The physical, human-readable name of the CodeDeploy Deployment Group.
	// Experimental.
	DeploymentGroupName *string `json:"deploymentGroupName" yaml:"deploymentGroupName"`
	// All EC2 instances matching the given set of tags when a deployment occurs will be added to this Deployment Group.
	// Experimental.
	Ec2InstanceTags InstanceTagSet `json:"ec2InstanceTags" yaml:"ec2InstanceTags"`
	// Whether to continue a deployment even if fetching the alarm status from CloudWatch failed.
	// Experimental.
	IgnorePollAlarmsFailure *bool `json:"ignorePollAlarmsFailure" yaml:"ignorePollAlarmsFailure"`
	// If you've provided any auto-scaling groups with the {@link #autoScalingGroups} property, you can set this property to add User Data that installs the CodeDeploy agent on the instances.
	// See: https://docs.aws.amazon.com/codedeploy/latest/userguide/codedeploy-agent-operations-install.html
	//
	// Experimental.
	InstallAgent *bool `json:"installAgent" yaml:"installAgent"`
	// The load balancer to place in front of this Deployment Group.
	//
	// Can be created from either a classic Elastic Load Balancer,
	// or an Application Load Balancer / Network Load Balancer Target Group.
	// Experimental.
	LoadBalancer LoadBalancer `json:"loadBalancer" yaml:"loadBalancer"`
	// All on-premise instances matching the given set of tags when a deployment occurs will be added to this Deployment Group.
	// Experimental.
	OnPremiseInstanceTags InstanceTagSet `json:"onPremiseInstanceTags" yaml:"onPremiseInstanceTags"`
	// The service Role of this Deployment Group.
	// Experimental.
	Role awsiam.IRole `json:"role" yaml:"role"`
}

