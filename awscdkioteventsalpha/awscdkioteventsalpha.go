// The CDK Construct Library for AWS::IoTEvents
package awscdkioteventsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkioteventsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkioteventsalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents an AWS IoT Events input.
// Experimental.
type IInput interface {
	awscdk.IResource
	// The name of the input.
	// Experimental.
	InputName() *string
}

// The jsii proxy for IInput
type jsiiProxy_IInput struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IInput) InputName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputName",
		&returns,
	)
	return returns
}

// Defines an AWS IoT Events input in this stack.
//
// TODO: EXAMPLE
//
// Experimental.
type Input interface {
	awscdk.Resource
	IInput
	Env() *awscdk.ResourceEnvironment
	InputName() *string
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for Input
type jsiiProxy_Input struct {
	internal.Type__awscdkResource
	jsiiProxy_IInput
}

func (j *jsiiProxy_Input) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Input) InputName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Input) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Input) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Input) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewInput(scope constructs.Construct, id *string, props *InputProps) Input {
	_init_.Initialize()

	j := jsiiProxy_Input{}

	_jsii_.Create(
		"@aws-cdk/aws-iotevents-alpha.Input",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewInput_Override(i Input, scope constructs.Construct, id *string, props *InputProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-iotevents-alpha.Input",
		[]interface{}{scope, id, props},
		i,
	)
}

// Import an existing input.
// Experimental.
func Input_FromInputName(scope constructs.Construct, id *string, inputName *string) IInput {
	_init_.Initialize()

	var returns IInput

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-iotevents-alpha.Input",
		"fromInputName",
		[]interface{}{scope, id, inputName},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func Input_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-iotevents-alpha.Input",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func Input_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-iotevents-alpha.Input",
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
func (i *jsiiProxy_Input) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (i *jsiiProxy_Input) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		i,
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
func (i *jsiiProxy_Input) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		i,
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
func (i *jsiiProxy_Input) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (i *jsiiProxy_Input) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for defining an AWS IoT Events input.
//
// TODO: EXAMPLE
//
// Experimental.
type InputProps struct {
	// An expression that specifies an attribute-value pair in a JSON structure.
	//
	// Use this to specify an attribute from the JSON payload that is made available
	// by the input. Inputs are derived from messages sent to AWS IoT Events (BatchPutMessage).
	// Each such message contains a JSON payload. The attribute (and its paired value)
	// specified here are available for use in the condition expressions used by detectors.
	// Experimental.
	AttributeJsonPaths *[]*string `json:"attributeJsonPaths"`
	// The name of the input.
	// Experimental.
	InputName *string `json:"inputName"`
}

