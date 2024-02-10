package awsstepfunctions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// A collection of connected states.
//
// A StateGraph is used to keep track of all states that are connected (have
// transitions between them). It does not include the substatemachines in
// a Parallel's branches: those are their own StateGraphs, but the graphs
// themselves have a hierarchical relationship as well.
//
// By assigning states to a definitive StateGraph, we verify that no state
// machines are constructed. In particular:
//
// - Every state object can only ever be in 1 StateGraph, and not inadvertently
//   be used in two graphs.
// - Every stateId must be unique across all states in the entire state
//   machine.
//
// All policy statements in all states in all substatemachines are bubbled so
// that the top-level StateMachine instantiation can read them all and add
// them to the IAM Role.
//
// You do not need to instantiate this class; it is used internally.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var state state
//
//   stateGraph := awscdk.Aws_stepfunctions.NewStateGraph(state, jsii.String("graphDescription"))
//
type StateGraph interface {
	// The accumulated policy statements.
	PolicyStatements() *[]awsiam.PolicyStatement
	// state that gets executed when the state machine is launched.
	StartState() State
	// Set a timeout to render into the graph JSON.
	//
	// Read/write. Only makes sense on the top-level graph, subgraphs
	// do not support this feature.
	// Default: No timeout.
	//
	Timeout() awscdk.Duration
	SetTimeout(val awscdk.Duration)
	// Binds this StateGraph to the StateMachine it defines and updates state machine permissions.
	Bind(stateMachine StateMachine)
	// Register a Policy Statement used by states in this graph.
	RegisterPolicyStatement(statement awsiam.PolicyStatement)
	// Register a state as part of this graph.
	//
	// Called by State.bindToGraph().
	RegisterState(state State)
	// Register this graph as a child of the given graph.
	//
	// Resource changes will be bubbled up to the given graph.
	RegisterSuperGraph(graph StateGraph)
	// Return the Amazon States Language JSON for this graph.
	ToGraphJson() *map[string]interface{}
	// Return a string description of this graph.
	ToString() *string
}

// The jsii proxy struct for StateGraph
type jsiiProxy_StateGraph struct {
	_ byte // padding
}

func (j *jsiiProxy_StateGraph) PolicyStatements() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement
	_jsii_.Get(
		j,
		"policyStatements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StateGraph) StartState() State {
	var returns State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StateGraph) Timeout() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}


func NewStateGraph(startState State, graphDescription *string) StateGraph {
	_init_.Initialize()

	if err := validateNewStateGraphParameters(startState, graphDescription); err != nil {
		panic(err)
	}
	j := jsiiProxy_StateGraph{}

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.StateGraph",
		[]interface{}{startState, graphDescription},
		&j,
	)

	return &j
}

func NewStateGraph_Override(s StateGraph, startState State, graphDescription *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.StateGraph",
		[]interface{}{startState, graphDescription},
		s,
	)
}

func (j *jsiiProxy_StateGraph)SetTimeout(val awscdk.Duration) {
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (s *jsiiProxy_StateGraph) Bind(stateMachine StateMachine) {
	if err := s.validateBindParameters(stateMachine); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"bind",
		[]interface{}{stateMachine},
	)
}

func (s *jsiiProxy_StateGraph) RegisterPolicyStatement(statement awsiam.PolicyStatement) {
	if err := s.validateRegisterPolicyStatementParameters(statement); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"registerPolicyStatement",
		[]interface{}{statement},
	)
}

func (s *jsiiProxy_StateGraph) RegisterState(state State) {
	if err := s.validateRegisterStateParameters(state); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"registerState",
		[]interface{}{state},
	)
}

func (s *jsiiProxy_StateGraph) RegisterSuperGraph(graph StateGraph) {
	if err := s.validateRegisterSuperGraphParameters(graph); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"registerSuperGraph",
		[]interface{}{graph},
	)
}

func (s *jsiiProxy_StateGraph) ToGraphJson() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"toGraphJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StateGraph) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

