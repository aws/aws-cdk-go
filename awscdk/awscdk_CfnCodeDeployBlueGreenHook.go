// An experiment to bundle the entire CDK into a single module
package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v3"
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
// Experimental.
type CfnCodeDeployBlueGreenHook interface {
	CfnHook
	// Additional options for the blue/green deployment.
	// Experimental.
	AdditionalOptions() *CfnCodeDeployBlueGreenAdditionalOptions
	// Experimental.
	SetAdditionalOptions(val *CfnCodeDeployBlueGreenAdditionalOptions)
	// Properties of the Amazon ECS applications being deployed.
	// Experimental.
	Applications() *[]*CfnCodeDeployBlueGreenApplication
	// Experimental.
	SetApplications(val *[]*CfnCodeDeployBlueGreenApplication)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// Use lifecycle event hooks to specify a Lambda function that CodeDeploy can call to validate a deployment.
	//
	// You can use the same function or a different one for deployment lifecycle events.
	// Following completion of the validation tests,
	// the Lambda {@link CfnCodeDeployBlueGreenLifecycleEventHooks.afterAllowTraffic}
	// function calls back CodeDeploy and delivers a result of 'Succeeded' or 'Failed'.
	// Experimental.
	LifecycleEventHooks() *CfnCodeDeployBlueGreenLifecycleEventHooks
	// Experimental.
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
	// Experimental.
	LogicalId() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() ConstructNode
	// The IAM Role for CloudFormation to use to perform blue-green deployments.
	// Experimental.
	ServiceRole() *string
	// Experimental.
	SetServiceRole(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() Stack
	// Traffic routing configuration settings.
	// Experimental.
	TrafficRoutingConfig() *CfnTrafficRoutingConfig
	// Experimental.
	SetTrafficRoutingConfig(val *CfnTrafficRoutingConfig)
	// The type of the hook (for example, "AWS::CodeDeploy::BlueGreen").
	// Experimental.
	Type() *string
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
	// Experimental.
	RenderProperties(_props *map[string]interface{}) *map[string]interface{}
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session ISynthesisSession)
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

func (j *jsiiProxy_CfnCodeDeployBlueGreenHook) Node() ConstructNode {
	var returns ConstructNode
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
// Experimental.
func NewCfnCodeDeployBlueGreenHook(scope constructs.Construct, id *string, props *CfnCodeDeployBlueGreenHookProps) CfnCodeDeployBlueGreenHook {
	_init_.Initialize()

	if err := validateNewCfnCodeDeployBlueGreenHookParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCodeDeployBlueGreenHook{}

	_jsii_.Create(
		"monocdk.CfnCodeDeployBlueGreenHook",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Creates a new CodeDeploy blue-green ECS Hook.
// Experimental.
func NewCfnCodeDeployBlueGreenHook_Override(c CfnCodeDeployBlueGreenHook, scope constructs.Construct, id *string, props *CfnCodeDeployBlueGreenHookProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.CfnCodeDeployBlueGreenHook",
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
// Experimental.
func CfnCodeDeployBlueGreenHook_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCodeDeployBlueGreenHook_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.CfnCodeDeployBlueGreenHook",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnCodeDeployBlueGreenHook_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCodeDeployBlueGreenHook_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.CfnCodeDeployBlueGreenHook",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCodeDeployBlueGreenHook) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnCodeDeployBlueGreenHook) OnSynthesize(session constructs.ISynthesisSession) {
	if err := c.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnCodeDeployBlueGreenHook) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
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

func (c *jsiiProxy_CfnCodeDeployBlueGreenHook) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
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

func (c *jsiiProxy_CfnCodeDeployBlueGreenHook) Synthesize(session ISynthesisSession) {
	if err := c.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
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

func (c *jsiiProxy_CfnCodeDeployBlueGreenHook) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

