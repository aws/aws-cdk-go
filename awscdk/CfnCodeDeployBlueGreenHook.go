package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation Hook for CodeDeploy blue-green ECS deployments.
//
// Example:
//   var cfnTemplate cfnInclude
//
//   // mutating the hook
//   var myRole role
//
//   hook := cfnTemplate.GetHook(jsii.String("MyOutput"))
//   codeDeployHook := hook.(cfnCodeDeployBlueGreenHook)
//   codeDeployHook.serviceRole = myRole.RoleArn
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/blue-green.html#blue-green-template-reference
//
type CfnCodeDeployBlueGreenHook interface {
	CfnHook
	// Additional options for the blue/green deployment.
	// Default: - no additional options.
	//
	AdditionalOptions() *CfnCodeDeployBlueGreenAdditionalOptions
	SetAdditionalOptions(val *CfnCodeDeployBlueGreenAdditionalOptions)
	// Properties of the Amazon ECS applications being deployed.
	Applications() *[]*CfnCodeDeployBlueGreenApplication
	SetApplications(val *[]*CfnCodeDeployBlueGreenApplication)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Use lifecycle event hooks to specify a Lambda function that CodeDeploy can call to validate a deployment.
	//
	// You can use the same function or a different one for deployment lifecycle events.
	// Following completion of the validation tests,
	// the Lambda `CfnCodeDeployBlueGreenLifecycleEventHooks.afterAllowTraffic`
	// function calls back CodeDeploy and delivers a result of 'Succeeded' or 'Failed'.
	// Default: - no lifecycle event hooks.
	//
	LifecycleEventHooks() *CfnCodeDeployBlueGreenLifecycleEventHooks
	SetLifecycleEventHooks(val *CfnCodeDeployBlueGreenLifecycleEventHooks)
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
	// The IAM Role for CloudFormation to use to perform blue-green deployments.
	ServiceRole() *string
	SetServiceRole(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() Stack
	// Traffic routing configuration settings.
	// Default: - time-based canary traffic shifting, with a 15% step percentage and a five minute bake time.
	//
	TrafficRoutingConfig() *CfnTrafficRoutingConfig
	SetTrafficRoutingConfig(val *CfnTrafficRoutingConfig)
	// The type of the hook (for example, "AWS::CodeDeploy::BlueGreen").
	Type() *string
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(_props *map[string]interface{}) *map[string]interface{}
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for CfnCodeDeployBlueGreenHook
type jsiiProxy_CfnCodeDeployBlueGreenHook struct {
	jsiiProxy_CfnHook
}

func (j *jsiiProxy_CfnCodeDeployBlueGreenHook) AdditionalOptions() *CfnCodeDeployBlueGreenAdditionalOptions {
	var returns *CfnCodeDeployBlueGreenAdditionalOptions
	_jsii_.Get(
		j,
		"additionalOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCodeDeployBlueGreenHook) Applications() *[]*CfnCodeDeployBlueGreenApplication {
	var returns *[]*CfnCodeDeployBlueGreenApplication
	_jsii_.Get(
		j,
		"applications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCodeDeployBlueGreenHook) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCodeDeployBlueGreenHook) LifecycleEventHooks() *CfnCodeDeployBlueGreenLifecycleEventHooks {
	var returns *CfnCodeDeployBlueGreenLifecycleEventHooks
	_jsii_.Get(
		j,
		"lifecycleEventHooks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCodeDeployBlueGreenHook) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCodeDeployBlueGreenHook) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCodeDeployBlueGreenHook) ServiceRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCodeDeployBlueGreenHook) Stack() Stack {
	var returns Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCodeDeployBlueGreenHook) TrafficRoutingConfig() *CfnTrafficRoutingConfig {
	var returns *CfnTrafficRoutingConfig
	_jsii_.Get(
		j,
		"trafficRoutingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCodeDeployBlueGreenHook) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Creates a new CodeDeploy blue-green ECS Hook.
func NewCfnCodeDeployBlueGreenHook(scope constructs.Construct, id *string, props *CfnCodeDeployBlueGreenHookProps) CfnCodeDeployBlueGreenHook {
	_init_.Initialize()

	if err := validateNewCfnCodeDeployBlueGreenHookParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCodeDeployBlueGreenHook{}

	_jsii_.Create(
		"aws-cdk-lib.CfnCodeDeployBlueGreenHook",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Creates a new CodeDeploy blue-green ECS Hook.
func NewCfnCodeDeployBlueGreenHook_Override(c CfnCodeDeployBlueGreenHook, scope constructs.Construct, id *string, props *CfnCodeDeployBlueGreenHookProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.CfnCodeDeployBlueGreenHook",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnCodeDeployBlueGreenHook)SetAdditionalOptions(val *CfnCodeDeployBlueGreenAdditionalOptions) {
	if err := j.validateSetAdditionalOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalOptions",
		val,
	)
}

func (j *jsiiProxy_CfnCodeDeployBlueGreenHook)SetApplications(val *[]*CfnCodeDeployBlueGreenApplication) {
	if err := j.validateSetApplicationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applications",
		val,
	)
}

func (j *jsiiProxy_CfnCodeDeployBlueGreenHook)SetLifecycleEventHooks(val *CfnCodeDeployBlueGreenLifecycleEventHooks) {
	if err := j.validateSetLifecycleEventHooksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycleEventHooks",
		val,
	)
}

func (j *jsiiProxy_CfnCodeDeployBlueGreenHook)SetServiceRole(val *string) {
	if err := j.validateSetServiceRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceRole",
		val,
	)
}

func (j *jsiiProxy_CfnCodeDeployBlueGreenHook)SetTrafficRoutingConfig(val *CfnTrafficRoutingConfig) {
	if err := j.validateSetTrafficRoutingConfigParameters(val); err != nil {
		panic(err)
	}
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
func CfnCodeDeployBlueGreenHook_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCodeDeployBlueGreenHook_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnCodeDeployBlueGreenHook",
		"isCfnElement",
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
func CfnCodeDeployBlueGreenHook_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCodeDeployBlueGreenHook_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnCodeDeployBlueGreenHook",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCodeDeployBlueGreenHook) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnCodeDeployBlueGreenHook) RenderProperties(_props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{_props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCodeDeployBlueGreenHook) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

