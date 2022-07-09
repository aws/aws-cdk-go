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
// Example:
//   import autoscaling "github.com/aws/aws-cdk-go/awscdk"
//   import cloudwatch "github.com/aws/aws-cdk-go/awscdk"
//
//   var application serverApplication
//   var asg autoScalingGroup
//   var alarm alarm
//
//   deploymentGroup := codedeploy.NewServerDeploymentGroup(this, jsii.String("CodeDeployDeploymentGroup"), &serverDeploymentGroupProps{
//   	application: application,
//   	deploymentGroupName: jsii.String("MyDeploymentGroup"),
//   	autoScalingGroups: []iAutoScalingGroup{
//   		asg,
//   	},
//   	// adds User Data that installs the CodeDeploy agent on your auto-scaling groups hosts
//   	// default: true
//   	installAgent: jsii.Boolean(true),
//   	// adds EC2 instances matching tags
//   	ec2InstanceTags: codedeploy.NewInstanceTagSet(map[string][]*string{
//   		// any instance with tags satisfying
//   		// key1=v1 or key1=v2 or key2 (any value) or value v3 (any key)
//   		// will match this group
//   		"key1": []*string{
//   			jsii.String("v1"),
//   			jsii.String("v2"),
//   		},
//   		"key2": []*string{
//   		},
//   		"": []*string{
//   			jsii.String("v3"),
//   		},
//   	}),
//   	// adds on-premise instances matching tags
//   	onPremiseInstanceTags: codedeploy.NewInstanceTagSet(map[string][]*string{
//   		"key1": []*string{
//   			jsii.String("v1"),
//   			jsii.String("v2"),
//   		},
//   	}, map[string][]*string{
//   		"key2": []*string{
//   			jsii.String("v3"),
//   		},
//   	}),
//   	// CloudWatch alarms
//   	alarms: []iAlarm{
//   		alarm,
//   	},
//   	// whether to ignore failure to fetch the status of alarms from CloudWatch
//   	// default: false
//   	ignorePollAlarmsFailure: jsii.Boolean(false),
//   	// auto-rollback configuration
//   	autoRollback: &autoRollbackConfig{
//   		failedDeployment: jsii.Boolean(true),
//   		 // default: true
//   		stoppedDeployment: jsii.Boolean(true),
//   		 // default: false
//   		deploymentInAlarm: jsii.Boolean(true),
//   	},
//   })
//
// Experimental.
type AutoRollbackConfig struct {
	// Whether to automatically roll back a deployment during which one of the configured CloudWatch alarms for this Deployment Group went off.
	// Experimental.
	DeploymentInAlarm *bool `field:"optional" json:"deploymentInAlarm" yaml:"deploymentInAlarm"`
	// Whether to automatically roll back a deployment that fails.
	// Experimental.
	FailedDeployment *bool `field:"optional" json:"failedDeployment" yaml:"failedDeployment"`
	// Whether to automatically roll back a deployment that was manually stopped.
	// Experimental.
	StoppedDeployment *bool `field:"optional" json:"stoppedDeployment" yaml:"stoppedDeployment"`
}

// A CloudFormation `AWS::CodeDeploy::Application`.
//
// The `AWS::CodeDeploy::Application` resource creates an AWS CodeDeploy application. In CodeDeploy , an application is a name that functions as a container to ensure that the correct combination of revision, deployment configuration, and deployment group are referenced during a deployment. You can use the `AWS::CodeDeploy::DeploymentGroup` resource to associate the application with a CodeDeploy deployment group. For more information, see [CodeDeploy Deployments](https://docs.aws.amazon.com/codedeploy/latest/userguide/deployment-steps.html) in the *AWS CodeDeploy User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplication := awscdk.Aws_codedeploy.NewCfnApplication(this, jsii.String("MyCfnApplication"), &cfnApplicationProps{
//   	applicationName: jsii.String("applicationName"),
//   	computePlatform: jsii.String("computePlatform"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnApplication interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// A name for the application.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the application name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
	//
	// > Updates to `ApplicationName` are not supported.
	ApplicationName() *string
	SetApplicationName(val *string)
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// The compute platform that CodeDeploy deploys the application to.
	ComputePlatform() *string
	SetComputePlatform(val *string)
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
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// The metadata that you apply to CodeDeploy applications to help you organize and categorize them.
	//
	// Each tag consists of a key and an optional value, both of which you define.
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

// Properties for defining a `CfnApplication`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationProps := &cfnApplicationProps{
//   	applicationName: jsii.String("applicationName"),
//   	computePlatform: jsii.String("computePlatform"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnApplicationProps struct {
	// A name for the application.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the application name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
	//
	// > Updates to `ApplicationName` are not supported.
	ApplicationName *string `field:"optional" json:"applicationName" yaml:"applicationName"`
	// The compute platform that CodeDeploy deploys the application to.
	ComputePlatform *string `field:"optional" json:"computePlatform" yaml:"computePlatform"`
	// The metadata that you apply to CodeDeploy applications to help you organize and categorize them.
	//
	// Each tag consists of a key and an optional value, both of which you define.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::CodeDeploy::DeploymentConfig`.
//
// The `AWS::CodeDeploy::DeploymentConfig` resource creates a set of deployment rules, deployment success conditions, and deployment failure conditions that AWS CodeDeploy uses during a deployment. The deployment configuration specifies, through the use of a `MinimumHealthyHosts` value, the number or percentage of instances that must remain available at any time during a deployment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDeploymentConfig := awscdk.Aws_codedeploy.NewCfnDeploymentConfig(this, jsii.String("MyCfnDeploymentConfig"), &cfnDeploymentConfigProps{
//   	computePlatform: jsii.String("computePlatform"),
//   	deploymentConfigName: jsii.String("deploymentConfigName"),
//   	minimumHealthyHosts: &minimumHealthyHostsProperty{
//   		type: jsii.String("type"),
//   		value: jsii.Number(123),
//   	},
//   	trafficRoutingConfig: &trafficRoutingConfigProperty{
//   		type: jsii.String("type"),
//
//   		// the properties below are optional
//   		timeBasedCanary: &timeBasedCanaryProperty{
//   			canaryInterval: jsii.Number(123),
//   			canaryPercentage: jsii.Number(123),
//   		},
//   		timeBasedLinear: &timeBasedLinearProperty{
//   			linearInterval: jsii.Number(123),
//   			linearPercentage: jsii.Number(123),
//   		},
//   	},
//   })
//
type CfnDeploymentConfig interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// The destination platform type for the deployment ( `Lambda` , `Server` , or `ECS` ).
	ComputePlatform() *string
	SetComputePlatform(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// A name for the deployment configuration.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the deployment configuration name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
	//
	// > If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	DeploymentConfigName() *string
	SetDeploymentConfigName(val *string)
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
	MinimumHealthyHosts() interface{}
	SetMinimumHealthyHosts(val interface{})
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
	// The configuration that specifies how the deployment traffic is routed.
	TrafficRoutingConfig() interface{}
	SetTrafficRoutingConfig(val interface{})
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

func (c *jsiiProxy_CfnDeploymentConfig) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnDeploymentConfig) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDeploymentConfig) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnDeploymentConfig) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnDeploymentConfig) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnDeploymentConfig) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnDeploymentConfig) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

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

func (c *jsiiProxy_CfnDeploymentConfig) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnDeploymentConfig) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnDeploymentConfig) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_CfnDeploymentConfig) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

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

func (c *jsiiProxy_CfnDeploymentConfig) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_CfnDeploymentConfig) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// `MinimumHealthyHosts` is a property of the [DeploymentConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentconfig.html) resource that defines how many instances must remain healthy during an AWS CodeDeploy deployment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   minimumHealthyHostsProperty := &minimumHealthyHostsProperty{
//   	type: jsii.String("type"),
//   	value: jsii.Number(123),
//   }
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
	Type *string `field:"required" json:"type" yaml:"type"`
	// The minimum healthy instance value.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

// A configuration that shifts traffic from one version of a Lambda function or Amazon ECS task set to another in two increments.
//
// The original and target Lambda function versions or ECS task sets are specified in the deployment's AppSpec file.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   timeBasedCanaryProperty := &timeBasedCanaryProperty{
//   	canaryInterval: jsii.Number(123),
//   	canaryPercentage: jsii.Number(123),
//   }
//
type CfnDeploymentConfig_TimeBasedCanaryProperty struct {
	// The number of minutes between the first and second traffic shifts of a `TimeBasedCanary` deployment.
	CanaryInterval *float64 `field:"required" json:"canaryInterval" yaml:"canaryInterval"`
	// The percentage of traffic to shift in the first increment of a `TimeBasedCanary` deployment.
	CanaryPercentage *float64 `field:"required" json:"canaryPercentage" yaml:"canaryPercentage"`
}

// A configuration that shifts traffic from one version of a Lambda function or ECS task set to another in equal increments, with an equal number of minutes between each increment.
//
// The original and target Lambda function versions or ECS task sets are specified in the deployment's AppSpec file.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   timeBasedLinearProperty := &timeBasedLinearProperty{
//   	linearInterval: jsii.Number(123),
//   	linearPercentage: jsii.Number(123),
//   }
//
type CfnDeploymentConfig_TimeBasedLinearProperty struct {
	// The number of minutes between each incremental traffic shift of a `TimeBasedLinear` deployment.
	LinearInterval *float64 `field:"required" json:"linearInterval" yaml:"linearInterval"`
	// The percentage of traffic that is shifted at the start of each increment of a `TimeBasedLinear` deployment.
	LinearPercentage *float64 `field:"required" json:"linearPercentage" yaml:"linearPercentage"`
}

// The configuration that specifies how traffic is shifted from one version of a Lambda function to another version during an AWS Lambda deployment, or from one Amazon ECS task set to another during an Amazon ECS deployment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   trafficRoutingConfigProperty := &trafficRoutingConfigProperty{
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	timeBasedCanary: &timeBasedCanaryProperty{
//   		canaryInterval: jsii.Number(123),
//   		canaryPercentage: jsii.Number(123),
//   	},
//   	timeBasedLinear: &timeBasedLinearProperty{
//   		linearInterval: jsii.Number(123),
//   		linearPercentage: jsii.Number(123),
//   	},
//   }
//
type CfnDeploymentConfig_TrafficRoutingConfigProperty struct {
	// The type of traffic shifting ( `TimeBasedCanary` or `TimeBasedLinear` ) used by a deployment configuration.
	Type *string `field:"required" json:"type" yaml:"type"`
	// A configuration that shifts traffic from one version of a Lambda function or ECS task set to another in two increments.
	//
	// The original and target Lambda function versions or ECS task sets are specified in the deployment's AppSpec file.
	TimeBasedCanary interface{} `field:"optional" json:"timeBasedCanary" yaml:"timeBasedCanary"`
	// A configuration that shifts traffic from one version of a Lambda function or Amazon ECS task set to another in equal increments, with an equal number of minutes between each increment.
	//
	// The original and target Lambda function versions or Amazon ECS task sets are specified in the deployment's AppSpec file.
	TimeBasedLinear interface{} `field:"optional" json:"timeBasedLinear" yaml:"timeBasedLinear"`
}

// Properties for defining a `CfnDeploymentConfig`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDeploymentConfigProps := &cfnDeploymentConfigProps{
//   	computePlatform: jsii.String("computePlatform"),
//   	deploymentConfigName: jsii.String("deploymentConfigName"),
//   	minimumHealthyHosts: &minimumHealthyHostsProperty{
//   		type: jsii.String("type"),
//   		value: jsii.Number(123),
//   	},
//   	trafficRoutingConfig: &trafficRoutingConfigProperty{
//   		type: jsii.String("type"),
//
//   		// the properties below are optional
//   		timeBasedCanary: &timeBasedCanaryProperty{
//   			canaryInterval: jsii.Number(123),
//   			canaryPercentage: jsii.Number(123),
//   		},
//   		timeBasedLinear: &timeBasedLinearProperty{
//   			linearInterval: jsii.Number(123),
//   			linearPercentage: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnDeploymentConfigProps struct {
	// The destination platform type for the deployment ( `Lambda` , `Server` , or `ECS` ).
	ComputePlatform *string `field:"optional" json:"computePlatform" yaml:"computePlatform"`
	// A name for the deployment configuration.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the deployment configuration name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
	//
	// > If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	DeploymentConfigName *string `field:"optional" json:"deploymentConfigName" yaml:"deploymentConfigName"`
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
	MinimumHealthyHosts interface{} `field:"optional" json:"minimumHealthyHosts" yaml:"minimumHealthyHosts"`
	// The configuration that specifies how the deployment traffic is routed.
	TrafficRoutingConfig interface{} `field:"optional" json:"trafficRoutingConfig" yaml:"trafficRoutingConfig"`
}

// A CloudFormation `AWS::CodeDeploy::DeploymentGroup`.
//
// The `AWS::CodeDeploy::DeploymentGroup` resource creates an AWS CodeDeploy deployment group that specifies which instances your application revisions are deployed to, along with other deployment options. For more information, see [CreateDeploymentGroup](https://docs.aws.amazon.com/codedeploy/latest/APIReference/API_CreateDeploymentGroup.html) in the *CodeDeploy API Reference* .
//
// > Amazon ECS blue/green deployments through CodeDeploy do not use the `AWS::CodeDeploy::DeploymentGroup` resource. To perform Amazon ECS blue/green deployments, use the `AWS::CodeDeploy::BlueGreen` hook. See [Perform Amazon ECS blue/green deployments through CodeDeploy using AWS CloudFormation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/blue-green.html) for more information.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDeploymentGroup := awscdk.Aws_codedeploy.NewCfnDeploymentGroup(this, jsii.String("MyCfnDeploymentGroup"), &cfnDeploymentGroupProps{
//   	applicationName: jsii.String("applicationName"),
//   	serviceRoleArn: jsii.String("serviceRoleArn"),
//
//   	// the properties below are optional
//   	alarmConfiguration: &alarmConfigurationProperty{
//   		alarms: []interface{}{
//   			&alarmProperty{
//   				name: jsii.String("name"),
//   			},
//   		},
//   		enabled: jsii.Boolean(false),
//   		ignorePollAlarmFailure: jsii.Boolean(false),
//   	},
//   	autoRollbackConfiguration: &autoRollbackConfigurationProperty{
//   		enabled: jsii.Boolean(false),
//   		events: []*string{
//   			jsii.String("events"),
//   		},
//   	},
//   	autoScalingGroups: []*string{
//   		jsii.String("autoScalingGroups"),
//   	},
//   	blueGreenDeploymentConfiguration: &blueGreenDeploymentConfigurationProperty{
//   		deploymentReadyOption: &deploymentReadyOptionProperty{
//   			actionOnTimeout: jsii.String("actionOnTimeout"),
//   			waitTimeInMinutes: jsii.Number(123),
//   		},
//   		greenFleetProvisioningOption: &greenFleetProvisioningOptionProperty{
//   			action: jsii.String("action"),
//   		},
//   		terminateBlueInstancesOnDeploymentSuccess: &blueInstanceTerminationOptionProperty{
//   			action: jsii.String("action"),
//   			terminationWaitTimeInMinutes: jsii.Number(123),
//   		},
//   	},
//   	deployment: &deploymentProperty{
//   		revision: &revisionLocationProperty{
//   			gitHubLocation: &gitHubLocationProperty{
//   				commitId: jsii.String("commitId"),
//   				repository: jsii.String("repository"),
//   			},
//   			revisionType: jsii.String("revisionType"),
//   			s3Location: &s3LocationProperty{
//   				bucket: jsii.String("bucket"),
//   				key: jsii.String("key"),
//
//   				// the properties below are optional
//   				bundleType: jsii.String("bundleType"),
//   				eTag: jsii.String("eTag"),
//   				version: jsii.String("version"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		description: jsii.String("description"),
//   		ignoreApplicationStopFailures: jsii.Boolean(false),
//   	},
//   	deploymentConfigName: jsii.String("deploymentConfigName"),
//   	deploymentGroupName: jsii.String("deploymentGroupName"),
//   	deploymentStyle: &deploymentStyleProperty{
//   		deploymentOption: jsii.String("deploymentOption"),
//   		deploymentType: jsii.String("deploymentType"),
//   	},
//   	ec2TagFilters: []interface{}{
//   		&eC2TagFilterProperty{
//   			key: jsii.String("key"),
//   			type: jsii.String("type"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	ec2TagSet: &eC2TagSetProperty{
//   		ec2TagSetList: []interface{}{
//   			&eC2TagSetListObjectProperty{
//   				ec2TagGroup: []interface{}{
//   					&eC2TagFilterProperty{
//   						key: jsii.String("key"),
//   						type: jsii.String("type"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	ecsServices: []interface{}{
//   		&eCSServiceProperty{
//   			clusterName: jsii.String("clusterName"),
//   			serviceName: jsii.String("serviceName"),
//   		},
//   	},
//   	loadBalancerInfo: &loadBalancerInfoProperty{
//   		elbInfoList: []interface{}{
//   			&eLBInfoProperty{
//   				name: jsii.String("name"),
//   			},
//   		},
//   		targetGroupInfoList: []interface{}{
//   			&targetGroupInfoProperty{
//   				name: jsii.String("name"),
//   			},
//   		},
//   		targetGroupPairInfoList: []interface{}{
//   			&targetGroupPairInfoProperty{
//   				prodTrafficRoute: &trafficRouteProperty{
//   					listenerArns: []*string{
//   						jsii.String("listenerArns"),
//   					},
//   				},
//   				targetGroups: []interface{}{
//   					&targetGroupInfoProperty{
//   						name: jsii.String("name"),
//   					},
//   				},
//   				testTrafficRoute: &trafficRouteProperty{
//   					listenerArns: []*string{
//   						jsii.String("listenerArns"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	onPremisesInstanceTagFilters: []interface{}{
//   		&tagFilterProperty{
//   			key: jsii.String("key"),
//   			type: jsii.String("type"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	onPremisesTagSet: &onPremisesTagSetProperty{
//   		onPremisesTagSetList: []interface{}{
//   			&onPremisesTagSetListObjectProperty{
//   				onPremisesTagGroup: []interface{}{
//   					&tagFilterProperty{
//   						key: jsii.String("key"),
//   						type: jsii.String("type"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	outdatedInstancesStrategy: jsii.String("outdatedInstancesStrategy"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	triggerConfigurations: []interface{}{
//   		&triggerConfigProperty{
//   			triggerEvents: []*string{
//   				jsii.String("triggerEvents"),
//   			},
//   			triggerName: jsii.String("triggerName"),
//   			triggerTargetArn: jsii.String("triggerTargetArn"),
//   		},
//   	},
//   })
//
type CfnDeploymentGroup interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Information about the Amazon CloudWatch alarms that are associated with the deployment group.
	AlarmConfiguration() interface{}
	SetAlarmConfiguration(val interface{})
	// The name of an existing CodeDeploy application to associate this deployment group with.
	ApplicationName() *string
	SetApplicationName(val *string)
	// Information about the automatic rollback configuration that is associated with the deployment group.
	//
	// If you specify this property, don't specify the `Deployment` property.
	AutoRollbackConfiguration() interface{}
	SetAutoRollbackConfiguration(val interface{})
	// A list of associated Auto Scaling groups that CodeDeploy automatically deploys revisions to when new instances are created.
	//
	// Duplicates are not allowed.
	AutoScalingGroups() *[]*string
	SetAutoScalingGroups(val *[]*string)
	// Information about blue/green deployment options for a deployment group.
	BlueGreenDeploymentConfiguration() interface{}
	SetBlueGreenDeploymentConfiguration(val interface{})
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
	// The application revision to deploy to this deployment group.
	//
	// If you specify this property, your target application revision is deployed as soon as the provisioning process is complete. If you specify this property, don't specify the `AutoRollbackConfiguration` property.
	Deployment() interface{}
	SetDeployment(val interface{})
	// A deployment configuration name or a predefined configuration name.
	//
	// With predefined configurations, you can deploy application revisions to one instance at a time ( `CodeDeployDefault.OneAtATime` ), half of the instances at a time ( `CodeDeployDefault.HalfAtATime` ), or all the instances at once ( `CodeDeployDefault.AllAtOnce` ). For more information and valid values, see [Working with Deployment Configurations](https://docs.aws.amazon.com/codedeploy/latest/userguide/deployment-configurations.html) in the *AWS CodeDeploy User Guide* .
	DeploymentConfigName() *string
	SetDeploymentConfigName(val *string)
	// A name for the deployment group.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the deployment group name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
	//
	// > If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	DeploymentGroupName() *string
	SetDeploymentGroupName(val *string)
	// Attributes that determine the type of deployment to run and whether to route deployment traffic behind a load balancer.
	//
	// If you specify this property with a blue/green deployment type, don't specify the `AutoScalingGroups` , `LoadBalancerInfo` , or `Deployment` properties.
	//
	// > For blue/green deployments, AWS CloudFormation supports deployments on Lambda compute platforms only. You can perform Amazon ECS blue/green deployments using `AWS::CodeDeploy::BlueGreen` hook. See [Perform Amazon ECS blue/green deployments through CodeDeploy using AWS CloudFormation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/blue-green.html) for more information.
	DeploymentStyle() interface{}
	SetDeploymentStyle(val interface{})
	// The Amazon EC2 tags that are already applied to Amazon EC2 instances that you want to include in the deployment group.
	//
	// CodeDeploy includes all Amazon EC2 instances identified by any of the tags you specify in this deployment group. Duplicates are not allowed.
	//
	// You can specify `EC2TagFilters` or `Ec2TagSet` , but not both.
	Ec2TagFilters() interface{}
	SetEc2TagFilters(val interface{})
	// Information about groups of tags applied to Amazon EC2 instances.
	//
	// The deployment group includes only Amazon EC2 instances identified by all the tag groups. Cannot be used in the same call as `ec2TagFilter` .
	Ec2TagSet() interface{}
	SetEc2TagSet(val interface{})
	// The target Amazon ECS services in the deployment group.
	//
	// This applies only to deployment groups that use the Amazon ECS compute platform. A target Amazon ECS service is specified as an Amazon ECS cluster and service name pair using the format `<clustername>:<servicename>` .
	EcsServices() interface{}
	SetEcsServices(val interface{})
	// Information about the load balancer to use in a deployment.
	//
	// For more information, see [Integrating CodeDeploy with Elastic Load Balancing](https://docs.aws.amazon.com/codedeploy/latest/userguide/integrations-aws-elastic-load-balancing.html) in the *AWS CodeDeploy User Guide* .
	LoadBalancerInfo() interface{}
	SetLoadBalancerInfo(val interface{})
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
	// The on-premises instance tags already applied to on-premises instances that you want to include in the deployment group.
	//
	// CodeDeploy includes all on-premises instances identified by any of the tags you specify in this deployment group. To register on-premises instances with CodeDeploy , see [Working with On-Premises Instances for CodeDeploy](https://docs.aws.amazon.com/codedeploy/latest/userguide/instances-on-premises.html) in the *AWS CodeDeploy User Guide* . Duplicates are not allowed.
	//
	// You can specify `OnPremisesInstanceTagFilters` or `OnPremisesInstanceTagSet` , but not both.
	OnPremisesInstanceTagFilters() interface{}
	SetOnPremisesInstanceTagFilters(val interface{})
	// Information about groups of tags applied to on-premises instances.
	//
	// The deployment group includes only on-premises instances identified by all the tag groups.
	//
	// You can specify `OnPremisesInstanceTagFilters` or `OnPremisesInstanceTagSet` , but not both.
	OnPremisesTagSet() interface{}
	SetOnPremisesTagSet(val interface{})
	// Indicates what happens when new Amazon EC2 instances are launched mid-deployment and do not receive the deployed application revision.
	//
	// If this option is set to `UPDATE` or is unspecified, CodeDeploy initiates one or more 'auto-update outdated instances' deployments to apply the deployed application revision to the new Amazon EC2 instances.
	//
	// If this option is set to `IGNORE` , CodeDeploy does not initiate a deployment to update the new Amazon EC2 instances. This may result in instances having different revisions.
	OutdatedInstancesStrategy() *string
	SetOutdatedInstancesStrategy(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// A service role Amazon Resource Name (ARN) that grants CodeDeploy permission to make calls to AWS services on your behalf.
	//
	// For more information, see [Create a Service Role for AWS CodeDeploy](https://docs.aws.amazon.com/codedeploy/latest/userguide/getting-started-create-service-role.html) in the *AWS CodeDeploy User Guide* .
	//
	// > In some cases, you might need to add a dependency on the service role's policy. For more information, see IAM role policy in [DependsOn Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html) .
	ServiceRoleArn() *string
	SetServiceRoleArn(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// The metadata that you apply to CodeDeploy deployment groups to help you organize and categorize them.
	//
	// Each tag consists of a key and an optional value, both of which you define.
	Tags() awscdk.TagManager
	// Information about triggers associated with the deployment group.
	//
	// Duplicates are not allowed.
	TriggerConfigurations() interface{}
	SetTriggerConfigurations(val interface{})
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

func (j *jsiiProxy_CfnDeploymentGroup) OutdatedInstancesStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outdatedInstancesStrategy",
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

func (j *jsiiProxy_CfnDeploymentGroup) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
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

func (j *jsiiProxy_CfnDeploymentGroup) SetOutdatedInstancesStrategy(val *string) {
	_jsii_.Set(
		j,
		"outdatedInstancesStrategy",
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

func (c *jsiiProxy_CfnDeploymentGroup) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnDeploymentGroup) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDeploymentGroup) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnDeploymentGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnDeploymentGroup) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnDeploymentGroup) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnDeploymentGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

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

func (c *jsiiProxy_CfnDeploymentGroup) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnDeploymentGroup) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnDeploymentGroup) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_CfnDeploymentGroup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

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

func (c *jsiiProxy_CfnDeploymentGroup) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   alarmConfigurationProperty := &alarmConfigurationProperty{
//   	alarms: []interface{}{
//   		&alarmProperty{
//   			name: jsii.String("name"),
//   		},
//   	},
//   	enabled: jsii.Boolean(false),
//   	ignorePollAlarmFailure: jsii.Boolean(false),
//   }
//
type CfnDeploymentGroup_AlarmConfigurationProperty struct {
	// A list of alarms configured for the deployment group.
	//
	// A maximum of 10 alarms can be added to a deployment group.
	Alarms interface{} `field:"optional" json:"alarms" yaml:"alarms"`
	// Indicates whether the alarm configuration is enabled.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Indicates whether a deployment should continue if information about the current state of alarms cannot be retrieved from Amazon CloudWatch .
	//
	// The default value is `false` .
	//
	// - `true` : The deployment proceeds even if alarm status information can't be retrieved from CloudWatch .
	// - `false` : The deployment stops if alarm status information can't be retrieved from CloudWatch .
	IgnorePollAlarmFailure interface{} `field:"optional" json:"ignorePollAlarmFailure" yaml:"ignorePollAlarmFailure"`
}

// The `Alarm` property type specifies a CloudWatch alarm to use for an AWS CodeDeploy deployment group.
//
// The `Alarm` property of the [CodeDeploy DeploymentGroup AlarmConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-alarmconfiguration.html) property contains a list of `Alarm` property types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   alarmProperty := &alarmProperty{
//   	name: jsii.String("name"),
//   }
//
type CfnDeploymentGroup_AlarmProperty struct {
	// The name of the alarm.
	//
	// Maximum length is 255 characters. Each alarm name can be used only once in a list of alarms.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

// The `AutoRollbackConfiguration` property type configures automatic rollback for an AWS CodeDeploy deployment group when a deployment is not completed successfully.
//
// For more information, see [Automatic Rollbacks](https://docs.aws.amazon.com/codedeploy/latest/userguide/deployments-rollback-and-redeploy.html#deployments-rollback-and-redeploy-automatic-rollbacks) in the *AWS CodeDeploy User Guide* .
//
// `AutoRollbackConfiguration` is a property of the [DeploymentGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   autoRollbackConfigurationProperty := &autoRollbackConfigurationProperty{
//   	enabled: jsii.Boolean(false),
//   	events: []*string{
//   		jsii.String("events"),
//   	},
//   }
//
type CfnDeploymentGroup_AutoRollbackConfigurationProperty struct {
	// Indicates whether a defined automatic rollback configuration is currently enabled.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The event type or types that trigger a rollback.
	//
	// Valid values are `DEPLOYMENT_FAILURE` , `DEPLOYMENT_STOP_ON_ALARM` , or `DEPLOYMENT_STOP_ON_REQUEST` .
	Events *[]*string `field:"optional" json:"events" yaml:"events"`
}

// Information about blue/green deployment options for a deployment group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   blueGreenDeploymentConfigurationProperty := &blueGreenDeploymentConfigurationProperty{
//   	deploymentReadyOption: &deploymentReadyOptionProperty{
//   		actionOnTimeout: jsii.String("actionOnTimeout"),
//   		waitTimeInMinutes: jsii.Number(123),
//   	},
//   	greenFleetProvisioningOption: &greenFleetProvisioningOptionProperty{
//   		action: jsii.String("action"),
//   	},
//   	terminateBlueInstancesOnDeploymentSuccess: &blueInstanceTerminationOptionProperty{
//   		action: jsii.String("action"),
//   		terminationWaitTimeInMinutes: jsii.Number(123),
//   	},
//   }
//
type CfnDeploymentGroup_BlueGreenDeploymentConfigurationProperty struct {
	// Information about the action to take when newly provisioned instances are ready to receive traffic in a blue/green deployment.
	DeploymentReadyOption interface{} `field:"optional" json:"deploymentReadyOption" yaml:"deploymentReadyOption"`
	// Information about how instances are provisioned for a replacement environment in a blue/green deployment.
	GreenFleetProvisioningOption interface{} `field:"optional" json:"greenFleetProvisioningOption" yaml:"greenFleetProvisioningOption"`
	// Information about whether to terminate instances in the original fleet during a blue/green deployment.
	TerminateBlueInstancesOnDeploymentSuccess interface{} `field:"optional" json:"terminateBlueInstancesOnDeploymentSuccess" yaml:"terminateBlueInstancesOnDeploymentSuccess"`
}

// Information about whether instances in the original environment are terminated when a blue/green deployment is successful.
//
// `BlueInstanceTerminationOption` does not apply to Lambda deployments.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   blueInstanceTerminationOptionProperty := &blueInstanceTerminationOptionProperty{
//   	action: jsii.String("action"),
//   	terminationWaitTimeInMinutes: jsii.Number(123),
//   }
//
type CfnDeploymentGroup_BlueInstanceTerminationOptionProperty struct {
	// The action to take on instances in the original environment after a successful blue/green deployment.
	//
	// - `TERMINATE` : Instances are terminated after a specified wait time.
	// - `KEEP_ALIVE` : Instances are left running after they are deregistered from the load balancer and removed from the deployment group.
	Action *string `field:"optional" json:"action" yaml:"action"`
	// For an Amazon EC2 deployment, the number of minutes to wait after a successful blue/green deployment before terminating instances from the original environment.
	//
	// For an Amazon ECS deployment, the number of minutes before deleting the original (blue) task set. During an Amazon ECS deployment, CodeDeploy shifts traffic from the original (blue) task set to a replacement (green) task set.
	//
	// The maximum setting is 2880 minutes (2 days).
	TerminationWaitTimeInMinutes *float64 `field:"optional" json:"terminationWaitTimeInMinutes" yaml:"terminationWaitTimeInMinutes"`
}

// `Deployment` is a property of the [DeploymentGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html) resource that specifies an AWS CodeDeploy application revision to be deployed to instances in the deployment group. If you specify an application revision, your target revision is deployed as soon as the provisioning process is complete.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deploymentProperty := &deploymentProperty{
//   	revision: &revisionLocationProperty{
//   		gitHubLocation: &gitHubLocationProperty{
//   			commitId: jsii.String("commitId"),
//   			repository: jsii.String("repository"),
//   		},
//   		revisionType: jsii.String("revisionType"),
//   		s3Location: &s3LocationProperty{
//   			bucket: jsii.String("bucket"),
//   			key: jsii.String("key"),
//
//   			// the properties below are optional
//   			bundleType: jsii.String("bundleType"),
//   			eTag: jsii.String("eTag"),
//   			version: jsii.String("version"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	ignoreApplicationStopFailures: jsii.Boolean(false),
//   }
//
type CfnDeploymentGroup_DeploymentProperty struct {
	// Information about the location of stored application artifacts and the service from which to retrieve them.
	Revision interface{} `field:"required" json:"revision" yaml:"revision"`
	// A comment about the deployment.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// If true, then if an `ApplicationStop` , `BeforeBlockTraffic` , or `AfterBlockTraffic` deployment lifecycle event to an instance fails, then the deployment continues to the next deployment lifecycle event.
	//
	// For example, if `ApplicationStop` fails, the deployment continues with DownloadBundle. If `BeforeBlockTraffic` fails, the deployment continues with `BlockTraffic` . If `AfterBlockTraffic` fails, the deployment continues with `ApplicationStop` .
	//
	// If false or not specified, then if a lifecycle event fails during a deployment to an instance, that deployment fails. If deployment to that instance is part of an overall deployment and the number of healthy hosts is not less than the minimum number of healthy hosts, then a deployment to the next instance is attempted.
	//
	// During a deployment, the AWS CodeDeploy agent runs the scripts specified for `ApplicationStop` , `BeforeBlockTraffic` , and `AfterBlockTraffic` in the AppSpec file from the previous successful deployment. (All other scripts are run from the AppSpec file in the current deployment.) If one of these scripts contains an error and does not run successfully, the deployment can fail.
	//
	// If the cause of the failure is a script from the last successful deployment that will never run successfully, create a new deployment and use `ignoreApplicationStopFailures` to specify that the `ApplicationStop` , `BeforeBlockTraffic` , and `AfterBlockTraffic` failures should be ignored.
	IgnoreApplicationStopFailures interface{} `field:"optional" json:"ignoreApplicationStopFailures" yaml:"ignoreApplicationStopFailures"`
}

// Information about how traffic is rerouted to instances in a replacement environment in a blue/green deployment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deploymentReadyOptionProperty := &deploymentReadyOptionProperty{
//   	actionOnTimeout: jsii.String("actionOnTimeout"),
//   	waitTimeInMinutes: jsii.Number(123),
//   }
//
type CfnDeploymentGroup_DeploymentReadyOptionProperty struct {
	// Information about when to reroute traffic from an original environment to a replacement environment in a blue/green deployment.
	//
	// - CONTINUE_DEPLOYMENT: Register new instances with the load balancer immediately after the new application revision is installed on the instances in the replacement environment.
	// - STOP_DEPLOYMENT: Do not register new instances with a load balancer unless traffic rerouting is started using [ContinueDeployment](https://docs.aws.amazon.com/codedeploy/latest/APIReference/API_ContinueDeployment.html) . If traffic rerouting is not started before the end of the specified wait period, the deployment status is changed to Stopped.
	ActionOnTimeout *string `field:"optional" json:"actionOnTimeout" yaml:"actionOnTimeout"`
	// The number of minutes to wait before the status of a blue/green deployment is changed to Stopped if rerouting is not started manually.
	//
	// Applies only to the `STOP_DEPLOYMENT` option for `actionOnTimeout` .
	WaitTimeInMinutes *float64 `field:"optional" json:"waitTimeInMinutes" yaml:"waitTimeInMinutes"`
}

// Information about the type of deployment, either in-place or blue/green, you want to run and whether to route deployment traffic behind a load balancer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deploymentStyleProperty := &deploymentStyleProperty{
//   	deploymentOption: jsii.String("deploymentOption"),
//   	deploymentType: jsii.String("deploymentType"),
//   }
//
type CfnDeploymentGroup_DeploymentStyleProperty struct {
	// Indicates whether to route deployment traffic behind a load balancer.
	//
	// > An Amazon EC2 Application Load Balancer or Network Load Balancer is required for an Amazon ECS deployment.
	DeploymentOption *string `field:"optional" json:"deploymentOption" yaml:"deploymentOption"`
	// Indicates whether to run an in-place or blue/green deployment.
	DeploymentType *string `field:"optional" json:"deploymentType" yaml:"deploymentType"`
}

// Information about an Amazon EC2 tag filter.
//
// For more information about using tags and tag groups to help manage your Amazon EC2 instances and on-premises instances, see [Tagging Instances for Deployment Groups in AWS CodeDeploy](https://docs.aws.amazon.com/codedeploy/latest/userguide/instances-tagging.html) in the *AWS CodeDeploy User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eC2TagFilterProperty := &eC2TagFilterProperty{
//   	key: jsii.String("key"),
//   	type: jsii.String("type"),
//   	value: jsii.String("value"),
//   }
//
type CfnDeploymentGroup_EC2TagFilterProperty struct {
	// The tag filter key.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The tag filter type:.
	//
	// - `KEY_ONLY` : Key only.
	// - `VALUE_ONLY` : Value only.
	// - `KEY_AND_VALUE` : Key and value.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The tag filter value.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

// The `EC2TagSet` property type specifies information about groups of tags applied to Amazon EC2 instances.
//
// The deployment group includes only Amazon EC2 instances identified by all the tag groups. Cannot be used in the same template as EC2TagFilters.
//
// For more information about using tags and tag groups to help manage your Amazon EC2 instances and on-premises instances, see [Tagging Instances for Deployment Groups in AWS CodeDeploy](https://docs.aws.amazon.com/codedeploy/latest/userguide/instances-tagging.html) in the *AWS CodeDeploy User Guide* .
//
// `EC2TagSet` is a property of the [DeploymentGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html) resource type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eC2TagSetListObjectProperty := &eC2TagSetListObjectProperty{
//   	ec2TagGroup: []interface{}{
//   		&eC2TagFilterProperty{
//   			key: jsii.String("key"),
//   			type: jsii.String("type"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDeploymentGroup_EC2TagSetListObjectProperty struct {
	// A list that contains other lists of Amazon EC2 instance tag groups.
	//
	// For an instance to be included in the deployment group, it must be identified by all of the tag groups in the list.
	Ec2TagGroup interface{} `field:"optional" json:"ec2TagGroup" yaml:"ec2TagGroup"`
}

// The `EC2TagSet` property type specifies information about groups of tags applied to Amazon EC2 instances.
//
// The deployment group includes only Amazon EC2 instances identified by all the tag groups. `EC2TagSet` cannot be used in the same template as `EC2TagFilter` .
//
// For information about using tags and tag groups to help manage your Amazon EC2 instances and on-premises instances, see [Tagging Instances for Deployment Groups in AWS CodeDeploy](https://docs.aws.amazon.com/codedeploy/latest/userguide/instances-tagging.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eC2TagSetProperty := &eC2TagSetProperty{
//   	ec2TagSetList: []interface{}{
//   		&eC2TagSetListObjectProperty{
//   			ec2TagGroup: []interface{}{
//   				&eC2TagFilterProperty{
//   					key: jsii.String("key"),
//   					type: jsii.String("type"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnDeploymentGroup_EC2TagSetProperty struct {
	// The Amazon EC2 tags that are already applied to Amazon EC2 instances that you want to include in the deployment group.
	//
	// CodeDeploy includes all Amazon EC2 instances identified by any of the tags you specify in this deployment group.
	//
	// Duplicates are not allowed.
	Ec2TagSetList interface{} `field:"optional" json:"ec2TagSetList" yaml:"ec2TagSetList"`
}

// Contains the service and cluster names used to identify an Amazon ECS deployment's target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eCSServiceProperty := &eCSServiceProperty{
//   	clusterName: jsii.String("clusterName"),
//   	serviceName: jsii.String("serviceName"),
//   }
//
type CfnDeploymentGroup_ECSServiceProperty struct {
	// The name of the cluster that the Amazon ECS service is associated with.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// The name of the target Amazon ECS service.
	ServiceName *string `field:"required" json:"serviceName" yaml:"serviceName"`
}

// The `ELBInfo` property type specifies information about the Elastic Load Balancing load balancer used for an CodeDeploy deployment group.
//
// If you specify the `ELBInfo` property, the `DeploymentStyle.DeploymentOption` property must be set to `WITH_TRAFFIC_CONTROL` for AWS CodeDeploy to route your traffic using the specified load balancers.
//
// `ELBInfo` is a property of the [AWS CodeDeploy DeploymentGroup LoadBalancerInfo](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-loadbalancerinfo.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eLBInfoProperty := &eLBInfoProperty{
//   	name: jsii.String("name"),
//   }
//
type CfnDeploymentGroup_ELBInfoProperty struct {
	// For blue/green deployments, the name of the load balancer that is used to route traffic from original instances to replacement instances in a blue/green deployment.
	//
	// For in-place deployments, the name of the load balancer that instances are deregistered from so they are not serving traffic during a deployment, and then re-registered with after the deployment is complete.
	//
	// > AWS CloudFormation supports blue/green deployments on AWS Lambda compute platforms only.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

// `GitHubLocation` is a property of the [CodeDeploy DeploymentGroup Revision](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-deployment-revision.html) property that specifies the location of an application revision that is stored in GitHub.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gitHubLocationProperty := &gitHubLocationProperty{
//   	commitId: jsii.String("commitId"),
//   	repository: jsii.String("repository"),
//   }
//
type CfnDeploymentGroup_GitHubLocationProperty struct {
	// The SHA1 commit ID of the GitHub commit that represents the bundled artifacts for the application revision.
	CommitId *string `field:"required" json:"commitId" yaml:"commitId"`
	// The GitHub account and repository pair that stores a reference to the commit that represents the bundled artifacts for the application revision.
	//
	// Specify the value as `account/repository` .
	Repository *string `field:"required" json:"repository" yaml:"repository"`
}

// Information about the instances that belong to the replacement environment in a blue/green deployment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   greenFleetProvisioningOptionProperty := &greenFleetProvisioningOptionProperty{
//   	action: jsii.String("action"),
//   }
//
type CfnDeploymentGroup_GreenFleetProvisioningOptionProperty struct {
	// The method used to add instances to a replacement environment.
	//
	// - `DISCOVER_EXISTING` : Use instances that already exist or will be created manually.
	// - `COPY_AUTO_SCALING_GROUP` : Use settings from a specified Auto Scaling group to define and create instances in a new Auto Scaling group.
	Action *string `field:"optional" json:"action" yaml:"action"`
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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loadBalancerInfoProperty := &loadBalancerInfoProperty{
//   	elbInfoList: []interface{}{
//   		&eLBInfoProperty{
//   			name: jsii.String("name"),
//   		},
//   	},
//   	targetGroupInfoList: []interface{}{
//   		&targetGroupInfoProperty{
//   			name: jsii.String("name"),
//   		},
//   	},
//   	targetGroupPairInfoList: []interface{}{
//   		&targetGroupPairInfoProperty{
//   			prodTrafficRoute: &trafficRouteProperty{
//   				listenerArns: []*string{
//   					jsii.String("listenerArns"),
//   				},
//   			},
//   			targetGroups: []interface{}{
//   				&targetGroupInfoProperty{
//   					name: jsii.String("name"),
//   				},
//   			},
//   			testTrafficRoute: &trafficRouteProperty{
//   				listenerArns: []*string{
//   					jsii.String("listenerArns"),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnDeploymentGroup_LoadBalancerInfoProperty struct {
	// An array that contains information about the load balancer to use for load balancing in a deployment.
	//
	// In Elastic Load Balancing, load balancers are used with Classic Load Balancers.
	//
	// > Adding more than one load balancer to the array is not supported.
	ElbInfoList interface{} `field:"optional" json:"elbInfoList" yaml:"elbInfoList"`
	// An array that contains information about the target group to use for load balancing in a deployment.
	//
	// In Elastic Load Balancing , target groups are used with Application Load Balancers .
	//
	// > Adding more than one target group to the array is not supported.
	TargetGroupInfoList interface{} `field:"optional" json:"targetGroupInfoList" yaml:"targetGroupInfoList"`
	// The target group pair information.
	//
	// This is an array of `TargeGroupPairInfo` objects with a maximum size of one.
	TargetGroupPairInfoList interface{} `field:"optional" json:"targetGroupPairInfoList" yaml:"targetGroupPairInfoList"`
}

// The `OnPremisesTagSetListObject` property type specifies lists of on-premises instance tag groups.
//
// In order for an instance to be included in the deployment group, it must be identified by all the tag groups in the list.
//
// `OnPremisesTagSetListObject` is a property of the [CodeDeploy DeploymentGroup OnPremisesTagSet](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-onpremisestagset.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   onPremisesTagSetListObjectProperty := &onPremisesTagSetListObjectProperty{
//   	onPremisesTagGroup: []interface{}{
//   		&tagFilterProperty{
//   			key: jsii.String("key"),
//   			type: jsii.String("type"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDeploymentGroup_OnPremisesTagSetListObjectProperty struct {
	// Information about groups of on-premises instance tags.
	OnPremisesTagGroup interface{} `field:"optional" json:"onPremisesTagGroup" yaml:"onPremisesTagGroup"`
}

// The `OnPremisesTagSet` property type specifies a list containing other lists of on-premises instance tag groups.
//
// In order for an instance to be included in the deployment group, it must be identified by all the tag groups in the list.
//
// For more information about using tags and tag groups to help manage your Amazon EC2 instances and on-premises instances, see [Tagging Instances for Deployment Groups in AWS CodeDeploy](https://docs.aws.amazon.com/codedeploy/latest/userguide/instances-tagging.html) in the *AWS CodeDeploy User Guide* .
//
// `OnPremisesTagSet` is a property of the [DeploymentGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   onPremisesTagSetProperty := &onPremisesTagSetProperty{
//   	onPremisesTagSetList: []interface{}{
//   		&onPremisesTagSetListObjectProperty{
//   			onPremisesTagGroup: []interface{}{
//   				&tagFilterProperty{
//   					key: jsii.String("key"),
//   					type: jsii.String("type"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnDeploymentGroup_OnPremisesTagSetProperty struct {
	// A list that contains other lists of on-premises instance tag groups.
	//
	// For an instance to be included in the deployment group, it must be identified by all of the tag groups in the list.
	//
	// Duplicates are not allowed.
	OnPremisesTagSetList interface{} `field:"optional" json:"onPremisesTagSetList" yaml:"onPremisesTagSetList"`
}

// `RevisionLocation` is a property that defines the location of the CodeDeploy application revision to deploy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   revisionLocationProperty := &revisionLocationProperty{
//   	gitHubLocation: &gitHubLocationProperty{
//   		commitId: jsii.String("commitId"),
//   		repository: jsii.String("repository"),
//   	},
//   	revisionType: jsii.String("revisionType"),
//   	s3Location: &s3LocationProperty{
//   		bucket: jsii.String("bucket"),
//   		key: jsii.String("key"),
//
//   		// the properties below are optional
//   		bundleType: jsii.String("bundleType"),
//   		eTag: jsii.String("eTag"),
//   		version: jsii.String("version"),
//   	},
//   }
//
type CfnDeploymentGroup_RevisionLocationProperty struct {
	// Information about the location of application artifacts stored in GitHub.
	GitHubLocation interface{} `field:"optional" json:"gitHubLocation" yaml:"gitHubLocation"`
	// The type of application revision:.
	//
	// - S3: An application revision stored in Amazon S3.
	// - GitHub: An application revision stored in GitHub (EC2/On-premises deployments only).
	// - String: A YAML-formatted or JSON-formatted string ( AWS Lambda deployments only).
	// - AppSpecContent: An `AppSpecContent` object that contains the contents of an AppSpec file for an AWS Lambda or Amazon ECS deployment. The content is formatted as JSON or YAML stored as a RawString.
	RevisionType *string `field:"optional" json:"revisionType" yaml:"revisionType"`
	// Information about the location of a revision stored in Amazon S3.
	S3Location interface{} `field:"optional" json:"s3Location" yaml:"s3Location"`
}

// `S3Location` is a property of the [CodeDeploy DeploymentGroup Revision](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-deployment-revision.html) property that specifies the location of an application revision that is stored in Amazon Simple Storage Service ( Amazon S3 ).
//
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
//   	bundleType: jsii.String("bundleType"),
//   	eTag: jsii.String("eTag"),
//   	version: jsii.String("version"),
//   }
//
type CfnDeploymentGroup_S3LocationProperty struct {
	// The name of the Amazon S3 bucket where the application revision is stored.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The name of the Amazon S3 object that represents the bundled artifacts for the application revision.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The file type of the application revision. Must be one of the following:.
	//
	// - JSON
	// - tar: A tar archive file.
	// - tgz: A compressed tar archive file.
	// - YAML
	// - zip: A zip archive file.
	BundleType *string `field:"optional" json:"bundleType" yaml:"bundleType"`
	// The ETag of the Amazon S3 object that represents the bundled artifacts for the application revision.
	//
	// If the ETag is not specified as an input parameter, ETag validation of the object is skipped.
	ETag *string `field:"optional" json:"eTag" yaml:"eTag"`
	// A specific version of the Amazon S3 object that represents the bundled artifacts for the application revision.
	//
	// If the version is not specified, the system uses the most recent version by default.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

// `TagFilter` is a property type of the [AWS::CodeDeploy::DeploymentGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html) resource that specifies which on-premises instances to associate with the deployment group. To register on-premise instances with AWS CodeDeploy , see [Configure Existing On-Premises Instances by Using AWS CodeDeploy](https://docs.aws.amazon.com/codedeploy/latest/userguide/instances-on-premises.html) in the *AWS CodeDeploy User Guide* .
//
// For more information about using tags and tag groups to help manage your Amazon EC2 instances and on-premises instances, see [Tagging Instances for Deployment Groups in AWS CodeDeploy](https://docs.aws.amazon.com/codedeploy/latest/userguide/instances-tagging.html) in the *AWS CodeDeploy User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tagFilterProperty := &tagFilterProperty{
//   	key: jsii.String("key"),
//   	type: jsii.String("type"),
//   	value: jsii.String("value"),
//   }
//
type CfnDeploymentGroup_TagFilterProperty struct {
	// The on-premises instance tag filter key.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The on-premises instance tag filter type:.
	//
	// - KEY_ONLY: Key only.
	// - VALUE_ONLY: Value only.
	// - KEY_AND_VALUE: Key and value.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The on-premises instance tag filter value.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

// The `TargetGroupInfo` property type specifies information about a target group in Elastic Load Balancing to use in a deployment.
//
// Instances are registered as targets in a target group, and traffic is routed to the target group. For more information, see [TargetGroupInfo](https://docs.aws.amazon.com/codedeploy/latest/APIReference/API_TargetGroupInfo.html) in the *AWS CodeDeploy API Reference*
//
// If you specify the `TargetGroupInfo` property, the `DeploymentStyle.DeploymentOption` property must be set to `WITH_TRAFFIC_CONTROL` for CodeDeploy to route your traffic using the specified target groups.
//
// `TargetGroupInfo` is a property of the [LoadBalancerInfo](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-loadbalancerinfo.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetGroupInfoProperty := &targetGroupInfoProperty{
//   	name: jsii.String("name"),
//   }
//
type CfnDeploymentGroup_TargetGroupInfoProperty struct {
	// For blue/green deployments, the name of the target group that instances in the original environment are deregistered from, and instances in the replacement environment registered with.
	//
	// For in-place deployments, the name of the target group that instances are deregistered from, so they are not serving traffic during a deployment, and then re-registered with after the deployment completes. No duplicates allowed.
	//
	// > AWS CloudFormation supports blue/green deployments on AWS Lambda compute platforms only.
	//
	// This value cannot exceed 32 characters, so you should use the `Name` property of the target group, or the `TargetGroupName` attribute with the `Fn::GetAtt` intrinsic function, as shown in the following example. Don't use the group's Amazon Resource Name (ARN) or `TargetGroupFullName` attribute.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

// Information about two target groups and how traffic is routed during an Amazon ECS deployment.
//
// An optional test traffic route can be specified.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetGroupPairInfoProperty := &targetGroupPairInfoProperty{
//   	prodTrafficRoute: &trafficRouteProperty{
//   		listenerArns: []*string{
//   			jsii.String("listenerArns"),
//   		},
//   	},
//   	targetGroups: []interface{}{
//   		&targetGroupInfoProperty{
//   			name: jsii.String("name"),
//   		},
//   	},
//   	testTrafficRoute: &trafficRouteProperty{
//   		listenerArns: []*string{
//   			jsii.String("listenerArns"),
//   		},
//   	},
//   }
//
type CfnDeploymentGroup_TargetGroupPairInfoProperty struct {
	// The path used by a load balancer to route production traffic when an Amazon ECS deployment is complete.
	ProdTrafficRoute interface{} `field:"optional" json:"prodTrafficRoute" yaml:"prodTrafficRoute"`
	// One pair of target groups.
	//
	// One is associated with the original task set. The second is associated with the task set that serves traffic after the deployment is complete.
	TargetGroups interface{} `field:"optional" json:"targetGroups" yaml:"targetGroups"`
	// An optional path used by a load balancer to route test traffic after an Amazon ECS deployment.
	//
	// Validation can occur while test traffic is served during a deployment.
	TestTrafficRoute interface{} `field:"optional" json:"testTrafficRoute" yaml:"testTrafficRoute"`
}

// Information about a listener.
//
// The listener contains the path used to route traffic that is received from the load balancer to a target group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   trafficRouteProperty := &trafficRouteProperty{
//   	listenerArns: []*string{
//   		jsii.String("listenerArns"),
//   	},
//   }
//
type CfnDeploymentGroup_TrafficRouteProperty struct {
	// The Amazon Resource Name (ARN) of one listener.
	//
	// The listener identifies the route between a target group and a load balancer. This is an array of strings with a maximum size of one.
	ListenerArns *[]*string `field:"optional" json:"listenerArns" yaml:"listenerArns"`
}

// Information about notification triggers for the deployment group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   triggerConfigProperty := &triggerConfigProperty{
//   	triggerEvents: []*string{
//   		jsii.String("triggerEvents"),
//   	},
//   	triggerName: jsii.String("triggerName"),
//   	triggerTargetArn: jsii.String("triggerTargetArn"),
//   }
//
type CfnDeploymentGroup_TriggerConfigProperty struct {
	// The event type or types that trigger notifications.
	TriggerEvents *[]*string `field:"optional" json:"triggerEvents" yaml:"triggerEvents"`
	// The name of the notification trigger.
	TriggerName *string `field:"optional" json:"triggerName" yaml:"triggerName"`
	// The Amazon Resource Name (ARN) of the Amazon Simple Notification Service topic through which notifications about deployment or instance events are sent.
	TriggerTargetArn *string `field:"optional" json:"triggerTargetArn" yaml:"triggerTargetArn"`
}

// Properties for defining a `CfnDeploymentGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDeploymentGroupProps := &cfnDeploymentGroupProps{
//   	applicationName: jsii.String("applicationName"),
//   	serviceRoleArn: jsii.String("serviceRoleArn"),
//
//   	// the properties below are optional
//   	alarmConfiguration: &alarmConfigurationProperty{
//   		alarms: []interface{}{
//   			&alarmProperty{
//   				name: jsii.String("name"),
//   			},
//   		},
//   		enabled: jsii.Boolean(false),
//   		ignorePollAlarmFailure: jsii.Boolean(false),
//   	},
//   	autoRollbackConfiguration: &autoRollbackConfigurationProperty{
//   		enabled: jsii.Boolean(false),
//   		events: []*string{
//   			jsii.String("events"),
//   		},
//   	},
//   	autoScalingGroups: []*string{
//   		jsii.String("autoScalingGroups"),
//   	},
//   	blueGreenDeploymentConfiguration: &blueGreenDeploymentConfigurationProperty{
//   		deploymentReadyOption: &deploymentReadyOptionProperty{
//   			actionOnTimeout: jsii.String("actionOnTimeout"),
//   			waitTimeInMinutes: jsii.Number(123),
//   		},
//   		greenFleetProvisioningOption: &greenFleetProvisioningOptionProperty{
//   			action: jsii.String("action"),
//   		},
//   		terminateBlueInstancesOnDeploymentSuccess: &blueInstanceTerminationOptionProperty{
//   			action: jsii.String("action"),
//   			terminationWaitTimeInMinutes: jsii.Number(123),
//   		},
//   	},
//   	deployment: &deploymentProperty{
//   		revision: &revisionLocationProperty{
//   			gitHubLocation: &gitHubLocationProperty{
//   				commitId: jsii.String("commitId"),
//   				repository: jsii.String("repository"),
//   			},
//   			revisionType: jsii.String("revisionType"),
//   			s3Location: &s3LocationProperty{
//   				bucket: jsii.String("bucket"),
//   				key: jsii.String("key"),
//
//   				// the properties below are optional
//   				bundleType: jsii.String("bundleType"),
//   				eTag: jsii.String("eTag"),
//   				version: jsii.String("version"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		description: jsii.String("description"),
//   		ignoreApplicationStopFailures: jsii.Boolean(false),
//   	},
//   	deploymentConfigName: jsii.String("deploymentConfigName"),
//   	deploymentGroupName: jsii.String("deploymentGroupName"),
//   	deploymentStyle: &deploymentStyleProperty{
//   		deploymentOption: jsii.String("deploymentOption"),
//   		deploymentType: jsii.String("deploymentType"),
//   	},
//   	ec2TagFilters: []interface{}{
//   		&eC2TagFilterProperty{
//   			key: jsii.String("key"),
//   			type: jsii.String("type"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	ec2TagSet: &eC2TagSetProperty{
//   		ec2TagSetList: []interface{}{
//   			&eC2TagSetListObjectProperty{
//   				ec2TagGroup: []interface{}{
//   					&eC2TagFilterProperty{
//   						key: jsii.String("key"),
//   						type: jsii.String("type"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	ecsServices: []interface{}{
//   		&eCSServiceProperty{
//   			clusterName: jsii.String("clusterName"),
//   			serviceName: jsii.String("serviceName"),
//   		},
//   	},
//   	loadBalancerInfo: &loadBalancerInfoProperty{
//   		elbInfoList: []interface{}{
//   			&eLBInfoProperty{
//   				name: jsii.String("name"),
//   			},
//   		},
//   		targetGroupInfoList: []interface{}{
//   			&targetGroupInfoProperty{
//   				name: jsii.String("name"),
//   			},
//   		},
//   		targetGroupPairInfoList: []interface{}{
//   			&targetGroupPairInfoProperty{
//   				prodTrafficRoute: &trafficRouteProperty{
//   					listenerArns: []*string{
//   						jsii.String("listenerArns"),
//   					},
//   				},
//   				targetGroups: []interface{}{
//   					&targetGroupInfoProperty{
//   						name: jsii.String("name"),
//   					},
//   				},
//   				testTrafficRoute: &trafficRouteProperty{
//   					listenerArns: []*string{
//   						jsii.String("listenerArns"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	onPremisesInstanceTagFilters: []interface{}{
//   		&tagFilterProperty{
//   			key: jsii.String("key"),
//   			type: jsii.String("type"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	onPremisesTagSet: &onPremisesTagSetProperty{
//   		onPremisesTagSetList: []interface{}{
//   			&onPremisesTagSetListObjectProperty{
//   				onPremisesTagGroup: []interface{}{
//   					&tagFilterProperty{
//   						key: jsii.String("key"),
//   						type: jsii.String("type"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	outdatedInstancesStrategy: jsii.String("outdatedInstancesStrategy"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	triggerConfigurations: []interface{}{
//   		&triggerConfigProperty{
//   			triggerEvents: []*string{
//   				jsii.String("triggerEvents"),
//   			},
//   			triggerName: jsii.String("triggerName"),
//   			triggerTargetArn: jsii.String("triggerTargetArn"),
//   		},
//   	},
//   }
//
type CfnDeploymentGroupProps struct {
	// The name of an existing CodeDeploy application to associate this deployment group with.
	ApplicationName *string `field:"required" json:"applicationName" yaml:"applicationName"`
	// A service role Amazon Resource Name (ARN) that grants CodeDeploy permission to make calls to AWS services on your behalf.
	//
	// For more information, see [Create a Service Role for AWS CodeDeploy](https://docs.aws.amazon.com/codedeploy/latest/userguide/getting-started-create-service-role.html) in the *AWS CodeDeploy User Guide* .
	//
	// > In some cases, you might need to add a dependency on the service role's policy. For more information, see IAM role policy in [DependsOn Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html) .
	ServiceRoleArn *string `field:"required" json:"serviceRoleArn" yaml:"serviceRoleArn"`
	// Information about the Amazon CloudWatch alarms that are associated with the deployment group.
	AlarmConfiguration interface{} `field:"optional" json:"alarmConfiguration" yaml:"alarmConfiguration"`
	// Information about the automatic rollback configuration that is associated with the deployment group.
	//
	// If you specify this property, don't specify the `Deployment` property.
	AutoRollbackConfiguration interface{} `field:"optional" json:"autoRollbackConfiguration" yaml:"autoRollbackConfiguration"`
	// A list of associated Auto Scaling groups that CodeDeploy automatically deploys revisions to when new instances are created.
	//
	// Duplicates are not allowed.
	AutoScalingGroups *[]*string `field:"optional" json:"autoScalingGroups" yaml:"autoScalingGroups"`
	// Information about blue/green deployment options for a deployment group.
	BlueGreenDeploymentConfiguration interface{} `field:"optional" json:"blueGreenDeploymentConfiguration" yaml:"blueGreenDeploymentConfiguration"`
	// The application revision to deploy to this deployment group.
	//
	// If you specify this property, your target application revision is deployed as soon as the provisioning process is complete. If you specify this property, don't specify the `AutoRollbackConfiguration` property.
	Deployment interface{} `field:"optional" json:"deployment" yaml:"deployment"`
	// A deployment configuration name or a predefined configuration name.
	//
	// With predefined configurations, you can deploy application revisions to one instance at a time ( `CodeDeployDefault.OneAtATime` ), half of the instances at a time ( `CodeDeployDefault.HalfAtATime` ), or all the instances at once ( `CodeDeployDefault.AllAtOnce` ). For more information and valid values, see [Working with Deployment Configurations](https://docs.aws.amazon.com/codedeploy/latest/userguide/deployment-configurations.html) in the *AWS CodeDeploy User Guide* .
	DeploymentConfigName *string `field:"optional" json:"deploymentConfigName" yaml:"deploymentConfigName"`
	// A name for the deployment group.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the deployment group name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
	//
	// > If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	DeploymentGroupName *string `field:"optional" json:"deploymentGroupName" yaml:"deploymentGroupName"`
	// Attributes that determine the type of deployment to run and whether to route deployment traffic behind a load balancer.
	//
	// If you specify this property with a blue/green deployment type, don't specify the `AutoScalingGroups` , `LoadBalancerInfo` , or `Deployment` properties.
	//
	// > For blue/green deployments, AWS CloudFormation supports deployments on Lambda compute platforms only. You can perform Amazon ECS blue/green deployments using `AWS::CodeDeploy::BlueGreen` hook. See [Perform Amazon ECS blue/green deployments through CodeDeploy using AWS CloudFormation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/blue-green.html) for more information.
	DeploymentStyle interface{} `field:"optional" json:"deploymentStyle" yaml:"deploymentStyle"`
	// The Amazon EC2 tags that are already applied to Amazon EC2 instances that you want to include in the deployment group.
	//
	// CodeDeploy includes all Amazon EC2 instances identified by any of the tags you specify in this deployment group. Duplicates are not allowed.
	//
	// You can specify `EC2TagFilters` or `Ec2TagSet` , but not both.
	Ec2TagFilters interface{} `field:"optional" json:"ec2TagFilters" yaml:"ec2TagFilters"`
	// Information about groups of tags applied to Amazon EC2 instances.
	//
	// The deployment group includes only Amazon EC2 instances identified by all the tag groups. Cannot be used in the same call as `ec2TagFilter` .
	Ec2TagSet interface{} `field:"optional" json:"ec2TagSet" yaml:"ec2TagSet"`
	// The target Amazon ECS services in the deployment group.
	//
	// This applies only to deployment groups that use the Amazon ECS compute platform. A target Amazon ECS service is specified as an Amazon ECS cluster and service name pair using the format `<clustername>:<servicename>` .
	EcsServices interface{} `field:"optional" json:"ecsServices" yaml:"ecsServices"`
	// Information about the load balancer to use in a deployment.
	//
	// For more information, see [Integrating CodeDeploy with Elastic Load Balancing](https://docs.aws.amazon.com/codedeploy/latest/userguide/integrations-aws-elastic-load-balancing.html) in the *AWS CodeDeploy User Guide* .
	LoadBalancerInfo interface{} `field:"optional" json:"loadBalancerInfo" yaml:"loadBalancerInfo"`
	// The on-premises instance tags already applied to on-premises instances that you want to include in the deployment group.
	//
	// CodeDeploy includes all on-premises instances identified by any of the tags you specify in this deployment group. To register on-premises instances with CodeDeploy , see [Working with On-Premises Instances for CodeDeploy](https://docs.aws.amazon.com/codedeploy/latest/userguide/instances-on-premises.html) in the *AWS CodeDeploy User Guide* . Duplicates are not allowed.
	//
	// You can specify `OnPremisesInstanceTagFilters` or `OnPremisesInstanceTagSet` , but not both.
	OnPremisesInstanceTagFilters interface{} `field:"optional" json:"onPremisesInstanceTagFilters" yaml:"onPremisesInstanceTagFilters"`
	// Information about groups of tags applied to on-premises instances.
	//
	// The deployment group includes only on-premises instances identified by all the tag groups.
	//
	// You can specify `OnPremisesInstanceTagFilters` or `OnPremisesInstanceTagSet` , but not both.
	OnPremisesTagSet interface{} `field:"optional" json:"onPremisesTagSet" yaml:"onPremisesTagSet"`
	// Indicates what happens when new Amazon EC2 instances are launched mid-deployment and do not receive the deployed application revision.
	//
	// If this option is set to `UPDATE` or is unspecified, CodeDeploy initiates one or more 'auto-update outdated instances' deployments to apply the deployed application revision to the new Amazon EC2 instances.
	//
	// If this option is set to `IGNORE` , CodeDeploy does not initiate a deployment to update the new Amazon EC2 instances. This may result in instances having different revisions.
	OutdatedInstancesStrategy *string `field:"optional" json:"outdatedInstancesStrategy" yaml:"outdatedInstancesStrategy"`
	// The metadata that you apply to CodeDeploy deployment groups to help you organize and categorize them.
	//
	// Each tag consists of a key and an optional value, both of which you define.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Information about triggers associated with the deployment group.
	//
	// Duplicates are not allowed.
	TriggerConfigurations interface{} `field:"optional" json:"triggerConfigurations" yaml:"triggerConfigurations"`
}

// A custom Deployment Configuration for a Lambda Deployment Group.
//
// Example:
//   var application lambdaApplication
//   var alias alias
//   config := codedeploy.NewCustomLambdaDeploymentConfig(this, jsii.String("CustomConfig"), &customLambdaDeploymentConfigProps{
//   	type: codedeploy.customLambdaDeploymentConfigType_CANARY,
//   	interval: awscdk.Duration.minutes(jsii.Number(1)),
//   	percentage: jsii.Number(5),
//   })
//   deploymentGroup := codedeploy.NewLambdaDeploymentGroup(this, jsii.String("BlueGreenDeployment"), &lambdaDeploymentGroupProps{
//   	application: application,
//   	alias: alias,
//   	deploymentConfig: config,
//   })
//
// Experimental.
type CustomLambdaDeploymentConfig interface {
	awscdk.Resource
	ILambdaDeploymentConfig
	// The arn of the deployment config.
	// Experimental.
	DeploymentConfigArn() *string
	// The name of the deployment config.
	// Experimental.
	DeploymentConfigName() *string
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
	// Experimental.
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

func (c *jsiiProxy_CustomLambdaDeploymentConfig) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

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

func (c *jsiiProxy_CustomLambdaDeploymentConfig) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomLambdaDeploymentConfig) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_CustomLambdaDeploymentConfig) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomLambdaDeploymentConfig) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

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
// Example:
//   var application lambdaApplication
//   var alias alias
//   config := codedeploy.NewCustomLambdaDeploymentConfig(this, jsii.String("CustomConfig"), &customLambdaDeploymentConfigProps{
//   	type: codedeploy.customLambdaDeploymentConfigType_CANARY,
//   	interval: awscdk.Duration.minutes(jsii.Number(1)),
//   	percentage: jsii.Number(5),
//   })
//   deploymentGroup := codedeploy.NewLambdaDeploymentGroup(this, jsii.String("BlueGreenDeployment"), &lambdaDeploymentGroupProps{
//   	application: application,
//   	alias: alias,
//   	deploymentConfig: config,
//   })
//
// Experimental.
type CustomLambdaDeploymentConfigProps struct {
	// The interval, in number of minutes: - For LINEAR, how frequently additional traffic is shifted - For CANARY, how long to shift traffic before the full deployment.
	// Experimental.
	Interval awscdk.Duration `field:"required" json:"interval" yaml:"interval"`
	// The integer percentage of traffic to shift: - For LINEAR, the percentage to shift every interval - For CANARY, the percentage to shift until the interval passes, before the full deployment.
	// Experimental.
	Percentage *float64 `field:"required" json:"percentage" yaml:"percentage"`
	// The type of deployment config, either CANARY or LINEAR.
	// Experimental.
	Type CustomLambdaDeploymentConfigType `field:"required" json:"type" yaml:"type"`
	// The verbatim name of the deployment config.
	//
	// Must be unique per account/region.
	// Other parameters cannot be updated if this name is provided.
	// Experimental.
	DeploymentConfigName *string `field:"optional" json:"deploymentConfigName" yaml:"deploymentConfigName"`
}

// Lambda Deployment config type.
//
// Example:
//   var application lambdaApplication
//   var alias alias
//   config := codedeploy.NewCustomLambdaDeploymentConfig(this, jsii.String("CustomConfig"), &customLambdaDeploymentConfigProps{
//   	type: codedeploy.customLambdaDeploymentConfigType_CANARY,
//   	interval: awscdk.Duration.minutes(jsii.Number(1)),
//   	percentage: jsii.Number(5),
//   })
//   deploymentGroup := codedeploy.NewLambdaDeploymentGroup(this, jsii.String("BlueGreenDeployment"), &lambdaDeploymentGroupProps{
//   	application: application,
//   	alias: alias,
//   	deploymentConfig: config,
//   })
//
// Experimental.
type CustomLambdaDeploymentConfigType string

const (
	// Canary deployment type.
	// Experimental.
	CustomLambdaDeploymentConfigType_CANARY CustomLambdaDeploymentConfigType = "CANARY"
	// Linear deployment type.
	// Experimental.
	CustomLambdaDeploymentConfigType_LINEAR CustomLambdaDeploymentConfigType = "LINEAR"
)

// A CodeDeploy Application that deploys to an Amazon ECS service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ecsApplication := awscdk.Aws_codedeploy.NewEcsApplication(this, jsii.String("MyEcsApplication"), &ecsApplicationProps{
//   	applicationName: jsii.String("applicationName"),
//   })
//
// Experimental.
type EcsApplication interface {
	awscdk.Resource
	IEcsApplication
	// Experimental.
	ApplicationArn() *string
	// Experimental.
	ApplicationName() *string
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
	// Experimental.
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
// Returns: a Construct representing a reference to an existing Application.
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

func (e *jsiiProxy_EcsApplication) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		e,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

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

func (e *jsiiProxy_EcsApplication) OnPrepare() {
	_jsii_.InvokeVoid(
		e,
		"onPrepare",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsApplication) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		e,
		"onSynthesize",
		[]interface{}{session},
	)
}

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

func (e *jsiiProxy_EcsApplication) Prepare() {
	_jsii_.InvokeVoid(
		e,
		"prepare",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsApplication) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		e,
		"synthesize",
		[]interface{}{session},
	)
}

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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ecsApplicationProps := &ecsApplicationProps{
//   	applicationName: jsii.String("applicationName"),
//   }
//
// Experimental.
type EcsApplicationProps struct {
	// The physical, human-readable name of the CodeDeploy Application.
	// Experimental.
	ApplicationName *string `field:"optional" json:"applicationName" yaml:"applicationName"`
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
// Returns: a Construct representing a reference to an existing custom Deployment Configuration.
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
// Returns: a Construct representing a reference to an existing Deployment Group.
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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var ecsApplication ecsApplication
//   var ecsDeploymentConfig iEcsDeploymentConfig
//
//   ecsDeploymentGroupAttributes := &ecsDeploymentGroupAttributes{
//   	application: ecsApplication,
//   	deploymentGroupName: jsii.String("deploymentGroupName"),
//
//   	// the properties below are optional
//   	deploymentConfig: ecsDeploymentConfig,
//   }
//
// See: EcsDeploymentGroup#fromEcsDeploymentGroupAttributes.
//
// Experimental.
type EcsDeploymentGroupAttributes struct {
	// The reference to the CodeDeploy ECS Application that this Deployment Group belongs to.
	// Experimental.
	Application IEcsApplication `field:"required" json:"application" yaml:"application"`
	// The physical, human-readable name of the CodeDeploy ECS Deployment Group that we are referencing.
	// Experimental.
	DeploymentGroupName *string `field:"required" json:"deploymentGroupName" yaml:"deploymentGroupName"`
	// The Deployment Configuration this Deployment Group uses.
	// Experimental.
	DeploymentConfig IEcsDeploymentConfig `field:"optional" json:"deploymentConfig" yaml:"deploymentConfig"`
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
// Example:
//   import autoscaling "github.com/aws/aws-cdk-go/awscdk"
//   import cloudwatch "github.com/aws/aws-cdk-go/awscdk"
//
//   var application serverApplication
//   var asg autoScalingGroup
//   var alarm alarm
//
//   deploymentGroup := codedeploy.NewServerDeploymentGroup(this, jsii.String("CodeDeployDeploymentGroup"), &serverDeploymentGroupProps{
//   	application: application,
//   	deploymentGroupName: jsii.String("MyDeploymentGroup"),
//   	autoScalingGroups: []iAutoScalingGroup{
//   		asg,
//   	},
//   	// adds User Data that installs the CodeDeploy agent on your auto-scaling groups hosts
//   	// default: true
//   	installAgent: jsii.Boolean(true),
//   	// adds EC2 instances matching tags
//   	ec2InstanceTags: codedeploy.NewInstanceTagSet(map[string][]*string{
//   		// any instance with tags satisfying
//   		// key1=v1 or key1=v2 or key2 (any value) or value v3 (any key)
//   		// will match this group
//   		"key1": []*string{
//   			jsii.String("v1"),
//   			jsii.String("v2"),
//   		},
//   		"key2": []*string{
//   		},
//   		"": []*string{
//   			jsii.String("v3"),
//   		},
//   	}),
//   	// adds on-premise instances matching tags
//   	onPremiseInstanceTags: codedeploy.NewInstanceTagSet(map[string][]*string{
//   		"key1": []*string{
//   			jsii.String("v1"),
//   			jsii.String("v2"),
//   		},
//   	}, map[string][]*string{
//   		"key2": []*string{
//   			jsii.String("v3"),
//   		},
//   	}),
//   	// CloudWatch alarms
//   	alarms: []iAlarm{
//   		alarm,
//   	},
//   	// whether to ignore failure to fetch the status of alarms from CloudWatch
//   	// default: false
//   	ignorePollAlarmsFailure: jsii.Boolean(false),
//   	// auto-rollback configuration
//   	autoRollback: &autoRollbackConfig{
//   		failedDeployment: jsii.Boolean(true),
//   		 // default: true
//   		stoppedDeployment: jsii.Boolean(true),
//   		 // default: false
//   		deploymentInAlarm: jsii.Boolean(true),
//   	},
//   })
//
// Experimental.
type InstanceTagSet interface {
	// Experimental.
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
// Example:
//   application := codedeploy.NewLambdaApplication(this, jsii.String("CodeDeployApplication"), &lambdaApplicationProps{
//   	applicationName: jsii.String("MyApplication"),
//   })
//
// Experimental.
type LambdaApplication interface {
	awscdk.Resource
	ILambdaApplication
	// Experimental.
	ApplicationArn() *string
	// Experimental.
	ApplicationName() *string
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
	// Experimental.
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
// Returns: a Construct representing a reference to an existing Application.
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

func (l *jsiiProxy_LambdaApplication) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		l,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

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

func (l *jsiiProxy_LambdaApplication) OnPrepare() {
	_jsii_.InvokeVoid(
		l,
		"onPrepare",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaApplication) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		l,
		"onSynthesize",
		[]interface{}{session},
	)
}

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

func (l *jsiiProxy_LambdaApplication) Prepare() {
	_jsii_.InvokeVoid(
		l,
		"prepare",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaApplication) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		l,
		"synthesize",
		[]interface{}{session},
	)
}

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
// Example:
//   application := codedeploy.NewLambdaApplication(this, jsii.String("CodeDeployApplication"), &lambdaApplicationProps{
//   	applicationName: jsii.String("MyApplication"),
//   })
//
// Experimental.
type LambdaApplicationProps struct {
	// The physical, human-readable name of the CodeDeploy Application.
	// Experimental.
	ApplicationName *string `field:"optional" json:"applicationName" yaml:"applicationName"`
}

// A custom Deployment Configuration for a Lambda Deployment Group.
//
// Note: This class currently stands as namespaced container of the default configurations
// until CloudFormation supports custom Lambda Deployment Configs. Until then it is closed
// (private constructor) and does not extend {@link cdk.Construct}
//
// Example:
//   var myApplication lambdaApplication
//   var func function
//
//   version := func.currentVersion
//   version1Alias := lambda.NewAlias(this, jsii.String("alias"), &aliasProps{
//   	aliasName: jsii.String("prod"),
//   	version: version,
//   })
//
//   deploymentGroup := codedeploy.NewLambdaDeploymentGroup(this, jsii.String("BlueGreenDeployment"), &lambdaDeploymentGroupProps{
//   	application: myApplication,
//   	 // optional property: one will be created for you if not provided
//   	alias: version1Alias,
//   	deploymentConfig: codedeploy.lambdaDeploymentConfig_LINEAR_10PERCENT_EVERY_1MINUTE(),
//   })
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
// Returns: a Construct representing a reference to an existing custom Deployment Configuration.
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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaDeploymentConfigImportProps := &lambdaDeploymentConfigImportProps{
//   	deploymentConfigName: jsii.String("deploymentConfigName"),
//   }
//
// See: LambdaDeploymentConfig#import.
//
// Experimental.
type LambdaDeploymentConfigImportProps struct {
	// The physical, human-readable name of the custom CodeDeploy Lambda Deployment Configuration that we are referencing.
	// Experimental.
	DeploymentConfigName *string `field:"required" json:"deploymentConfigName" yaml:"deploymentConfigName"`
}

// Example:
//   var application lambdaApplication
//   var alias alias
//   config := codedeploy.NewCustomLambdaDeploymentConfig(this, jsii.String("CustomConfig"), &customLambdaDeploymentConfigProps{
//   	type: codedeploy.customLambdaDeploymentConfigType_CANARY,
//   	interval: awscdk.Duration.minutes(jsii.Number(1)),
//   	percentage: jsii.Number(5),
//   })
//   deploymentGroup := codedeploy.NewLambdaDeploymentGroup(this, jsii.String("BlueGreenDeployment"), &lambdaDeploymentGroupProps{
//   	application: application,
//   	alias: alias,
//   	deploymentConfig: config,
//   })
//
// Experimental.
type LambdaDeploymentGroup interface {
	awscdk.Resource
	ILambdaDeploymentGroup
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
	// Experimental.
	Role() awsiam.IRole
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// Associates an additional alarm with this Deployment Group.
	// Experimental.
	AddAlarm(alarm awscloudwatch.IAlarm)
	// Associate a function to run after deployment completes.
	// Experimental.
	AddPostHook(postHook awslambda.IFunction)
	// Associate a function to run before deployment begins.
	// Experimental.
	AddPreHook(preHook awslambda.IFunction)
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
	// Grant a principal permission to codedeploy:PutLifecycleEventHookExecutionStatus on this deployment group resource.
	// Experimental.
	GrantPutLifecycleEventHookExecutionStatus(grantee awsiam.IGrantable) awsiam.Grant
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
	// Experimental.
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
// Returns: a Construct representing a reference to an existing Deployment Group.
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

func (l *jsiiProxy_LambdaDeploymentGroup) AddAlarm(alarm awscloudwatch.IAlarm) {
	_jsii_.InvokeVoid(
		l,
		"addAlarm",
		[]interface{}{alarm},
	)
}

func (l *jsiiProxy_LambdaDeploymentGroup) AddPostHook(postHook awslambda.IFunction) {
	_jsii_.InvokeVoid(
		l,
		"addPostHook",
		[]interface{}{postHook},
	)
}

func (l *jsiiProxy_LambdaDeploymentGroup) AddPreHook(preHook awslambda.IFunction) {
	_jsii_.InvokeVoid(
		l,
		"addPreHook",
		[]interface{}{preHook},
	)
}

func (l *jsiiProxy_LambdaDeploymentGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		l,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

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

func (l *jsiiProxy_LambdaDeploymentGroup) OnPrepare() {
	_jsii_.InvokeVoid(
		l,
		"onPrepare",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaDeploymentGroup) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		l,
		"onSynthesize",
		[]interface{}{session},
	)
}

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

func (l *jsiiProxy_LambdaDeploymentGroup) Prepare() {
	_jsii_.InvokeVoid(
		l,
		"prepare",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaDeploymentGroup) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		l,
		"synthesize",
		[]interface{}{session},
	)
}

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
// Example:
//   var application lambdaApplication
//
//   deploymentGroup := codedeploy.lambdaDeploymentGroup.fromLambdaDeploymentGroupAttributes(this, jsii.String("ExistingCodeDeployDeploymentGroup"), &lambdaDeploymentGroupAttributes{
//   	application: application,
//   	deploymentGroupName: jsii.String("MyExistingDeploymentGroup"),
//   })
//
// See: LambdaDeploymentGroup#fromLambdaDeploymentGroupAttributes.
//
// Experimental.
type LambdaDeploymentGroupAttributes struct {
	// The reference to the CodeDeploy Lambda Application that this Deployment Group belongs to.
	// Experimental.
	Application ILambdaApplication `field:"required" json:"application" yaml:"application"`
	// The physical, human-readable name of the CodeDeploy Lambda Deployment Group that we are referencing.
	// Experimental.
	DeploymentGroupName *string `field:"required" json:"deploymentGroupName" yaml:"deploymentGroupName"`
	// The Deployment Configuration this Deployment Group uses.
	// Experimental.
	DeploymentConfig ILambdaDeploymentConfig `field:"optional" json:"deploymentConfig" yaml:"deploymentConfig"`
}

// Construction properties for {@link LambdaDeploymentGroup}.
//
// Example:
//   var application lambdaApplication
//   var alias alias
//   config := codedeploy.NewCustomLambdaDeploymentConfig(this, jsii.String("CustomConfig"), &customLambdaDeploymentConfigProps{
//   	type: codedeploy.customLambdaDeploymentConfigType_CANARY,
//   	interval: awscdk.Duration.minutes(jsii.Number(1)),
//   	percentage: jsii.Number(5),
//   })
//   deploymentGroup := codedeploy.NewLambdaDeploymentGroup(this, jsii.String("BlueGreenDeployment"), &lambdaDeploymentGroupProps{
//   	application: application,
//   	alias: alias,
//   	deploymentConfig: config,
//   })
//
// Experimental.
type LambdaDeploymentGroupProps struct {
	// Lambda Alias to shift traffic. Updating the version of the alias will trigger a CodeDeploy deployment.
	//
	// [disable-awslint:ref-via-interface] since we need to modify the alias CFN resource update policy.
	// Experimental.
	Alias awslambda.Alias `field:"required" json:"alias" yaml:"alias"`
	// The CloudWatch alarms associated with this Deployment Group.
	//
	// CodeDeploy will stop (and optionally roll back)
	// a deployment if during it any of the alarms trigger.
	//
	// Alarms can also be added after the Deployment Group is created using the {@link #addAlarm} method.
	// See: https://docs.aws.amazon.com/codedeploy/latest/userguide/monitoring-create-alarms.html
	//
	// Experimental.
	Alarms *[]awscloudwatch.IAlarm `field:"optional" json:"alarms" yaml:"alarms"`
	// The reference to the CodeDeploy Lambda Application that this Deployment Group belongs to.
	// Experimental.
	Application ILambdaApplication `field:"optional" json:"application" yaml:"application"`
	// The auto-rollback configuration for this Deployment Group.
	// Experimental.
	AutoRollback *AutoRollbackConfig `field:"optional" json:"autoRollback" yaml:"autoRollback"`
	// The Deployment Configuration this Deployment Group uses.
	// Experimental.
	DeploymentConfig ILambdaDeploymentConfig `field:"optional" json:"deploymentConfig" yaml:"deploymentConfig"`
	// The physical, human-readable name of the CodeDeploy Deployment Group.
	// Experimental.
	DeploymentGroupName *string `field:"optional" json:"deploymentGroupName" yaml:"deploymentGroupName"`
	// Whether to continue a deployment even if fetching the alarm status from CloudWatch failed.
	// Experimental.
	IgnorePollAlarmsFailure *bool `field:"optional" json:"ignorePollAlarmsFailure" yaml:"ignorePollAlarmsFailure"`
	// The Lambda function to run after traffic routing starts.
	// Experimental.
	PostHook awslambda.IFunction `field:"optional" json:"postHook" yaml:"postHook"`
	// The Lambda function to run before traffic routing starts.
	// Experimental.
	PreHook awslambda.IFunction `field:"optional" json:"preHook" yaml:"preHook"`
	// The service Role of this Deployment Group.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}

// An interface of an abstract load balancer, as needed by CodeDeploy.
//
// Create instances using the static factory methods:
// {@link #classic}, {@link #application} and {@link #network}.
//
// Example:
//   import elb "github.com/aws/aws-cdk-go/awscdk"
//
//   var lb loadBalancer
//
//   lb.addListener(&loadBalancerListener{
//   	externalPort: jsii.Number(80),
//   })
//
//   deploymentGroup := codedeploy.NewServerDeploymentGroup(this, jsii.String("DeploymentGroup"), &serverDeploymentGroupProps{
//   	loadBalancer: codedeploy.loadBalancer.classic(lb),
//   })
//
// Experimental.
type LoadBalancer interface {
	// Experimental.
	Generation() LoadBalancerGeneration
	// Experimental.
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
	// The first generation (ELB Classic).
	// Experimental.
	LoadBalancerGeneration_FIRST LoadBalancerGeneration = "FIRST"
	// The second generation (ALB and NLB).
	// Experimental.
	LoadBalancerGeneration_SECOND LoadBalancerGeneration = "SECOND"
)

// Minimum number of healthy hosts for a server deployment.
//
// Example:
//   deploymentConfig := codedeploy.NewServerDeploymentConfig(this, jsii.String("DeploymentConfiguration"), &serverDeploymentConfigProps{
//   	deploymentConfigName: jsii.String("MyDeploymentConfiguration"),
//   	 // optional property
//   	// one of these is required, but both cannot be specified at the same time
//   	minimumHealthyHosts: codedeploy.minimumHealthyHosts.count(jsii.Number(2)),
//   })
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
// Example:
//   application := codedeploy.NewServerApplication(this, jsii.String("CodeDeployApplication"), &serverApplicationProps{
//   	applicationName: jsii.String("MyApplication"),
//   })
//
// Experimental.
type ServerApplication interface {
	awscdk.Resource
	IServerApplication
	// Experimental.
	ApplicationArn() *string
	// Experimental.
	ApplicationName() *string
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
	// Experimental.
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
// Returns: a Construct representing a reference to an existing Application.
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

func (s *jsiiProxy_ServerApplication) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		s,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

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

func (s *jsiiProxy_ServerApplication) OnPrepare() {
	_jsii_.InvokeVoid(
		s,
		"onPrepare",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServerApplication) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"onSynthesize",
		[]interface{}{session},
	)
}

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

func (s *jsiiProxy_ServerApplication) Prepare() {
	_jsii_.InvokeVoid(
		s,
		"prepare",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServerApplication) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		[]interface{}{session},
	)
}

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
// Example:
//   application := codedeploy.NewServerApplication(this, jsii.String("CodeDeployApplication"), &serverApplicationProps{
//   	applicationName: jsii.String("MyApplication"),
//   })
//
// Experimental.
type ServerApplicationProps struct {
	// The physical, human-readable name of the CodeDeploy Application.
	// Experimental.
	ApplicationName *string `field:"optional" json:"applicationName" yaml:"applicationName"`
}

// A custom Deployment Configuration for an EC2/on-premise Deployment Group.
//
// Example:
//   deploymentGroup := codedeploy.NewServerDeploymentGroup(this, jsii.String("CodeDeployDeploymentGroup"), &serverDeploymentGroupProps{
//   	deploymentConfig: codedeploy.serverDeploymentConfig_ALL_AT_ONCE(),
//   })
//
// Experimental.
type ServerDeploymentConfig interface {
	awscdk.Resource
	IServerDeploymentConfig
	// Experimental.
	DeploymentConfigArn() *string
	// Experimental.
	DeploymentConfigName() *string
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
	// Experimental.
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
// Returns: a Construct representing a reference to an existing custom Deployment Configuration.
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

func (s *jsiiProxy_ServerDeploymentConfig) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		s,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

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

func (s *jsiiProxy_ServerDeploymentConfig) OnPrepare() {
	_jsii_.InvokeVoid(
		s,
		"onPrepare",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServerDeploymentConfig) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"onSynthesize",
		[]interface{}{session},
	)
}

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

func (s *jsiiProxy_ServerDeploymentConfig) Prepare() {
	_jsii_.InvokeVoid(
		s,
		"prepare",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServerDeploymentConfig) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		[]interface{}{session},
	)
}

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
// Example:
//   deploymentConfig := codedeploy.NewServerDeploymentConfig(this, jsii.String("DeploymentConfiguration"), &serverDeploymentConfigProps{
//   	deploymentConfigName: jsii.String("MyDeploymentConfiguration"),
//   	 // optional property
//   	// one of these is required, but both cannot be specified at the same time
//   	minimumHealthyHosts: codedeploy.minimumHealthyHosts.count(jsii.Number(2)),
//   })
//
// Experimental.
type ServerDeploymentConfigProps struct {
	// Minimum number of healthy hosts.
	// Experimental.
	MinimumHealthyHosts MinimumHealthyHosts `field:"required" json:"minimumHealthyHosts" yaml:"minimumHealthyHosts"`
	// The physical, human-readable name of the Deployment Configuration.
	// Experimental.
	DeploymentConfigName *string `field:"optional" json:"deploymentConfigName" yaml:"deploymentConfigName"`
}

// A CodeDeploy Deployment Group that deploys to EC2/on-premise instances.
//
// Example:
//   import elb "github.com/aws/aws-cdk-go/awscdk"
//
//   var lb loadBalancer
//
//   lb.addListener(&loadBalancerListener{
//   	externalPort: jsii.Number(80),
//   })
//
//   deploymentGroup := codedeploy.NewServerDeploymentGroup(this, jsii.String("DeploymentGroup"), &serverDeploymentGroupProps{
//   	loadBalancer: codedeploy.loadBalancer.classic(lb),
//   })
//
// Experimental.
type ServerDeploymentGroup interface {
	awscdk.Resource
	IServerDeploymentGroup
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
	// Experimental.
	Role() awsiam.IRole
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// Associates an additional alarm with this Deployment Group.
	// Experimental.
	AddAlarm(alarm awscloudwatch.IAlarm)
	// Adds an additional auto-scaling group to this Deployment Group.
	// Experimental.
	AddAutoScalingGroup(asg awsautoscaling.AutoScalingGroup)
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
	// Experimental.
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
// Returns: a Construct representing a reference to an existing Deployment Group.
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

func (s *jsiiProxy_ServerDeploymentGroup) AddAlarm(alarm awscloudwatch.IAlarm) {
	_jsii_.InvokeVoid(
		s,
		"addAlarm",
		[]interface{}{alarm},
	)
}

func (s *jsiiProxy_ServerDeploymentGroup) AddAutoScalingGroup(asg awsautoscaling.AutoScalingGroup) {
	_jsii_.InvokeVoid(
		s,
		"addAutoScalingGroup",
		[]interface{}{asg},
	)
}

func (s *jsiiProxy_ServerDeploymentGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		s,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

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

func (s *jsiiProxy_ServerDeploymentGroup) OnPrepare() {
	_jsii_.InvokeVoid(
		s,
		"onPrepare",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServerDeploymentGroup) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"onSynthesize",
		[]interface{}{session},
	)
}

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

func (s *jsiiProxy_ServerDeploymentGroup) Prepare() {
	_jsii_.InvokeVoid(
		s,
		"prepare",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServerDeploymentGroup) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		[]interface{}{session},
	)
}

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
// Example:
//   var application serverApplication
//
//   deploymentGroup := codedeploy.serverDeploymentGroup.fromServerDeploymentGroupAttributes(this, jsii.String("ExistingCodeDeployDeploymentGroup"), &serverDeploymentGroupAttributes{
//   	application: application,
//   	deploymentGroupName: jsii.String("MyExistingDeploymentGroup"),
//   })
//
// See: ServerDeploymentGroup#import.
//
// Experimental.
type ServerDeploymentGroupAttributes struct {
	// The reference to the CodeDeploy EC2/on-premise Application that this Deployment Group belongs to.
	// Experimental.
	Application IServerApplication `field:"required" json:"application" yaml:"application"`
	// The physical, human-readable name of the CodeDeploy EC2/on-premise Deployment Group that we are referencing.
	// Experimental.
	DeploymentGroupName *string `field:"required" json:"deploymentGroupName" yaml:"deploymentGroupName"`
	// The Deployment Configuration this Deployment Group uses.
	// Experimental.
	DeploymentConfig IServerDeploymentConfig `field:"optional" json:"deploymentConfig" yaml:"deploymentConfig"`
}

// Construction properties for {@link ServerDeploymentGroup}.
//
// Example:
//   import elbv2 "github.com/aws/aws-cdk-go/awscdk"
//
//   var alb applicationLoadBalancer
//
//   listener := alb.addListener(jsii.String("Listener"), &baseApplicationListenerProps{
//   	port: jsii.Number(80),
//   })
//   targetGroup := listener.addTargets(jsii.String("Fleet"), &addApplicationTargetsProps{
//   	port: jsii.Number(80),
//   })
//
//   deploymentGroup := codedeploy.NewServerDeploymentGroup(this, jsii.String("DeploymentGroup"), &serverDeploymentGroupProps{
//   	loadBalancer: codedeploy.loadBalancer.application(targetGroup),
//   })
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
	Alarms *[]awscloudwatch.IAlarm `field:"optional" json:"alarms" yaml:"alarms"`
	// The CodeDeploy EC2/on-premise Application this Deployment Group belongs to.
	// Experimental.
	Application IServerApplication `field:"optional" json:"application" yaml:"application"`
	// The auto-rollback configuration for this Deployment Group.
	// Experimental.
	AutoRollback *AutoRollbackConfig `field:"optional" json:"autoRollback" yaml:"autoRollback"`
	// The auto-scaling groups belonging to this Deployment Group.
	//
	// Auto-scaling groups can also be added after the Deployment Group is created
	// using the {@link #addAutoScalingGroup} method.
	//
	// [disable-awslint:ref-via-interface] is needed because we update userdata
	// for ASGs to install the codedeploy agent.
	// Experimental.
	AutoScalingGroups *[]awsautoscaling.IAutoScalingGroup `field:"optional" json:"autoScalingGroups" yaml:"autoScalingGroups"`
	// The EC2/on-premise Deployment Configuration to use for this Deployment Group.
	// Experimental.
	DeploymentConfig IServerDeploymentConfig `field:"optional" json:"deploymentConfig" yaml:"deploymentConfig"`
	// The physical, human-readable name of the CodeDeploy Deployment Group.
	// Experimental.
	DeploymentGroupName *string `field:"optional" json:"deploymentGroupName" yaml:"deploymentGroupName"`
	// All EC2 instances matching the given set of tags when a deployment occurs will be added to this Deployment Group.
	// Experimental.
	Ec2InstanceTags InstanceTagSet `field:"optional" json:"ec2InstanceTags" yaml:"ec2InstanceTags"`
	// Whether to continue a deployment even if fetching the alarm status from CloudWatch failed.
	// Experimental.
	IgnorePollAlarmsFailure *bool `field:"optional" json:"ignorePollAlarmsFailure" yaml:"ignorePollAlarmsFailure"`
	// If you've provided any auto-scaling groups with the {@link #autoScalingGroups} property, you can set this property to add User Data that installs the CodeDeploy agent on the instances.
	// See: https://docs.aws.amazon.com/codedeploy/latest/userguide/codedeploy-agent-operations-install.html
	//
	// Experimental.
	InstallAgent *bool `field:"optional" json:"installAgent" yaml:"installAgent"`
	// The load balancer to place in front of this Deployment Group.
	//
	// Can be created from either a classic Elastic Load Balancer,
	// or an Application Load Balancer / Network Load Balancer Target Group.
	// Experimental.
	LoadBalancer LoadBalancer `field:"optional" json:"loadBalancer" yaml:"loadBalancer"`
	// All on-premise instances matching the given set of tags when a deployment occurs will be added to this Deployment Group.
	// Experimental.
	OnPremiseInstanceTags InstanceTagSet `field:"optional" json:"onPremiseInstanceTags" yaml:"onPremiseInstanceTags"`
	// The service Role of this Deployment Group.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}

