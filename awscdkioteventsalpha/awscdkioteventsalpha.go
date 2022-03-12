// The CDK Construct Library for AWS::IoTEvents
package awscdkioteventsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkioteventsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotevents"
	"github.com/aws/aws-cdk-go/awscdkioteventsalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Options when binding a Action to a detector model.
//
// TODO: EXAMPLE
//
// Experimental.
type ActionBindOptions struct {
	// The IAM role assumed by IoT Events to perform the action.
	// Experimental.
	Role awsiam.IRole `json:"role" yaml:"role"`
}

// Properties for a AWS IoT Events action.
//
// TODO: EXAMPLE
//
// Experimental.
type ActionConfig struct {
	// The configuration for this action.
	// Experimental.
	Configuration *awsiotevents.CfnDetectorModel_ActionProperty `json:"configuration" yaml:"configuration"`
}

// Defines an AWS IoT Events detector model in this stack.
//
// TODO: EXAMPLE
//
// Experimental.
type DetectorModel interface {
	awscdk.Resource
	IDetectorModel
	DetectorModelName() *string
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

// The jsii proxy struct for DetectorModel
type jsiiProxy_DetectorModel struct {
	internal.Type__awscdkResource
	jsiiProxy_IDetectorModel
}

func (j *jsiiProxy_DetectorModel) DetectorModelName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"detectorModelName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DetectorModel) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DetectorModel) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DetectorModel) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DetectorModel) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewDetectorModel(scope constructs.Construct, id *string, props *DetectorModelProps) DetectorModel {
	_init_.Initialize()

	j := jsiiProxy_DetectorModel{}

	_jsii_.Create(
		"@aws-cdk/aws-iotevents-alpha.DetectorModel",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewDetectorModel_Override(d DetectorModel, scope constructs.Construct, id *string, props *DetectorModelProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-iotevents-alpha.DetectorModel",
		[]interface{}{scope, id, props},
		d,
	)
}

// Import an existing detector model.
// Experimental.
func DetectorModel_FromDetectorModelName(scope constructs.Construct, id *string, detectorModelName *string) IDetectorModel {
	_init_.Initialize()

	var returns IDetectorModel

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-iotevents-alpha.DetectorModel",
		"fromDetectorModelName",
		[]interface{}{scope, id, detectorModelName},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func DetectorModel_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-iotevents-alpha.DetectorModel",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func DetectorModel_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-iotevents-alpha.DetectorModel",
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
func (d *jsiiProxy_DetectorModel) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		d,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (d *jsiiProxy_DetectorModel) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		d,
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
func (d *jsiiProxy_DetectorModel) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		d,
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
func (d *jsiiProxy_DetectorModel) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (d *jsiiProxy_DetectorModel) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for defining an AWS IoT Events detector model.
//
// TODO: EXAMPLE
//
// Experimental.
type DetectorModelProps struct {
	// The state that is entered at the creation of each detector.
	// Experimental.
	InitialState State `json:"initialState" yaml:"initialState"`
	// A brief description of the detector model.
	// Experimental.
	Description *string `json:"description" yaml:"description"`
	// The value used to identify a detector instance.
	//
	// When a device or system sends input, a new
	// detector instance with a unique key value is created. AWS IoT Events can continue to route
	// input to its corresponding detector instance based on this identifying information.
	//
	// This parameter uses a JSON-path expression to select the attribute-value pair in the message
	// payload that is used for identification. To route the message to the correct detector instance,
	// the device must send a message payload that contains the same attribute-value.
	// Experimental.
	DetectorKey *string `json:"detectorKey" yaml:"detectorKey"`
	// The name of the detector model.
	// Experimental.
	DetectorModelName *string `json:"detectorModelName" yaml:"detectorModelName"`
	// Information about the order in which events are evaluated and how actions are executed.
	//
	// When setting to SERIAL, variables are updated and event conditions are evaluated in the order
	// that the events are defined.
	// When setting to BATCH, variables within a state are updated and events within a state are
	// performed only after all event conditions are evaluated.
	// Experimental.
	EvaluationMethod EventEvaluation `json:"evaluationMethod" yaml:"evaluationMethod"`
	// The role that grants permission to AWS IoT Events to perform its operations.
	// Experimental.
	Role awsiam.IRole `json:"role" yaml:"role"`
}

// Specifies the actions to be performed when the condition evaluates to `true`.
//
// TODO: EXAMPLE
//
// Experimental.
type Event struct {
	// The name of the event.
	// Experimental.
	EventName *string `json:"eventName" yaml:"eventName"`
	// The actions to be performed.
	// Experimental.
	Actions *[]IAction `json:"actions" yaml:"actions"`
	// The Boolean expression that, when `true`, causes the actions to be performed.
	// Experimental.
	Condition Expression `json:"condition" yaml:"condition"`
}

// Information about the order in which events are evaluated and how actions are executed.
//
// TODO: EXAMPLE
//
// Experimental.
type EventEvaluation string

const (
	EventEvaluation_BATCH EventEvaluation = "BATCH"
	EventEvaluation_SERIAL EventEvaluation = "SERIAL"
)

// Expression for events in Detector Model state.
//
// TODO: EXAMPLE
//
// See: https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-expressions.html
//
// Experimental.
type Expression interface {
	Evaluate() *string
}

// The jsii proxy struct for Expression
type jsiiProxy_Expression struct {
	_ byte // padding
}

// Experimental.
func NewExpression_Override(e Expression) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-iotevents-alpha.Expression",
		nil, // no parameters
		e,
	)
}

// Create a expression for the AND operator.
// Experimental.
func Expression_And(left Expression, right Expression) Expression {
	_init_.Initialize()

	var returns Expression

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-iotevents-alpha.Expression",
		"and",
		[]interface{}{left, right},
		&returns,
	)

	return returns
}

// Create a expression for function `currentInput()`.
//
// It is evaluated to true if the specified input message was received.
// Experimental.
func Expression_CurrentInput(input IInput) Expression {
	_init_.Initialize()

	var returns Expression

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-iotevents-alpha.Expression",
		"currentInput",
		[]interface{}{input},
		&returns,
	)

	return returns
}

// Create a expression for the Equal operator.
// Experimental.
func Expression_Eq(left Expression, right Expression) Expression {
	_init_.Initialize()

	var returns Expression

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-iotevents-alpha.Expression",
		"eq",
		[]interface{}{left, right},
		&returns,
	)

	return returns
}

// Create a expression from the given string.
// Experimental.
func Expression_FromString(value *string) Expression {
	_init_.Initialize()

	var returns Expression

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-iotevents-alpha.Expression",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Create a expression for get an input attribute as `$input.TemperatureInput.temperatures[2]`.
// Experimental.
func Expression_InputAttribute(input IInput, path *string) Expression {
	_init_.Initialize()

	var returns Expression

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-iotevents-alpha.Expression",
		"inputAttribute",
		[]interface{}{input, path},
		&returns,
	)

	return returns
}

// This is called to evaluate the expression.
// Experimental.
func (e *jsiiProxy_Expression) Evaluate() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"evaluate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// An abstract action for DetectorModel.
// Experimental.
type IAction interface {
	// Returns the AWS IoT Events action specification.
	// Experimental.
	Bind(scope constructs.Construct, options *ActionBindOptions) *ActionConfig
}

// The jsii proxy for IAction
type jsiiProxy_IAction struct {
	_ byte // padding
}

func (i *jsiiProxy_IAction) Bind(scope constructs.Construct, options *ActionBindOptions) *ActionConfig {
	var returns *ActionConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope, options},
		&returns,
	)

	return returns
}

// Represents an AWS IoT Events detector model.
// Experimental.
type IDetectorModel interface {
	awscdk.IResource
	// The name of the detector model.
	// Experimental.
	DetectorModelName() *string
}

// The jsii proxy for IDetectorModel
type jsiiProxy_IDetectorModel struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IDetectorModel) DetectorModelName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"detectorModelName",
		&returns,
	)
	return returns
}

// Represents an AWS IoT Events input.
// Experimental.
type IInput interface {
	awscdk.IResource
	// Grant the indicated permissions on this input to the given IAM principal (Role/Group/User).
	// Experimental.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Grant write permissions on this input and its contents to an IAM principal (Role/Group/User).
	// Experimental.
	GrantWrite(grantee awsiam.IGrantable) awsiam.Grant
	// The ARN of the input.
	// Experimental.
	InputArn() *string
	// The name of the input.
	// Experimental.
	InputName() *string
}

// The jsii proxy for IInput
type jsiiProxy_IInput struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IInput) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
	args := []interface{}{grantee}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grant",
		args,
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IInput) GrantWrite(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantWrite",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IInput) InputArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputArn",
		&returns,
	)
	return returns
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
	InputArn() *string
	InputName() *string
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	GrantWrite(grantee awsiam.IGrantable) awsiam.Grant
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

func (j *jsiiProxy_Input) InputArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputArn",
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

// Grant the indicated permissions on this input to the given IAM principal (Role/Group/User).
// Experimental.
func (i *jsiiProxy_Input) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
	args := []interface{}{grantee}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grant",
		args,
		&returns,
	)

	return returns
}

// Grant write permissions on this input and its contents to an IAM principal (Role/Group/User).
// Experimental.
func (i *jsiiProxy_Input) GrantWrite(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantWrite",
		[]interface{}{grantee},
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
	AttributeJsonPaths *[]*string `json:"attributeJsonPaths" yaml:"attributeJsonPaths"`
	// The name of the input.
	// Experimental.
	InputName *string `json:"inputName" yaml:"inputName"`
}

// Defines a state of a detector.
//
// TODO: EXAMPLE
//
// Experimental.
type State interface {
	StateName() *string
	TransitionTo(targetState State, options *TransitionOptions)
}

// The jsii proxy struct for State
type jsiiProxy_State struct {
	_ byte // padding
}

func (j *jsiiProxy_State) StateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateName",
		&returns,
	)
	return returns
}


// Experimental.
func NewState(props *StateProps) State {
	_init_.Initialize()

	j := jsiiProxy_State{}

	_jsii_.Create(
		"@aws-cdk/aws-iotevents-alpha.State",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewState_Override(s State, props *StateProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-iotevents-alpha.State",
		[]interface{}{props},
		s,
	)
}

// Add a transition event to the state.
//
// The transition event will be triggered if condition is evaluated to `true`.
// Experimental.
func (s *jsiiProxy_State) TransitionTo(targetState State, options *TransitionOptions) {
	_jsii_.InvokeVoid(
		s,
		"transitionTo",
		[]interface{}{targetState, options},
	)
}

// Properties for defining a state of a detector.
//
// TODO: EXAMPLE
//
// Experimental.
type StateProps struct {
	// The name of the state.
	// Experimental.
	StateName *string `json:"stateName" yaml:"stateName"`
	// Specifies the events on enter.
	//
	// The conditions of the events will be evaluated when entering this state.
	// If the condition of the event evaluates to `true`, the actions of the event will be executed.
	// Experimental.
	OnEnter *[]*Event `json:"onEnter" yaml:"onEnter"`
	// Specifies the events on exit.
	//
	// The conditions of the events are evaluated when an exiting this state.
	// If the condition evaluates to `true`, the actions of the event will be executed.
	// Experimental.
	OnExit *[]*Event `json:"onExit" yaml:"onExit"`
	// Specifies the events on input.
	//
	// The conditions of the events will be evaluated when any input is received.
	// If the condition of the event evaluates to `true`, the actions of the event will be executed.
	// Experimental.
	OnInput *[]*Event `json:"onInput" yaml:"onInput"`
}

// Properties for options of state transition.
//
// TODO: EXAMPLE
//
// Experimental.
type TransitionOptions struct {
	// The condition that is used to determine to cause the state transition and the actions.
	//
	// When this was evaluated to `true`, the state transition and the actions are triggered.
	// Experimental.
	When Expression `json:"when" yaml:"when"`
	// The name of the event.
	// Experimental.
	EventName *string `json:"eventName" yaml:"eventName"`
	// The actions to be performed with the transition.
	// Experimental.
	Executing *[]IAction `json:"executing" yaml:"executing"`
}

