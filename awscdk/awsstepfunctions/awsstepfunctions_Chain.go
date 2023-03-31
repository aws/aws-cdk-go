package awsstepfunctions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A collection of states to chain onto.
//
// A Chain has a start and zero or more chainable ends. If there are
// zero ends, calling next() on the Chain will fail.
//
// Example:
//   // Define a state machine with one Pass state
//   child := sfn.NewStateMachine(this, jsii.String("ChildStateMachine"), &stateMachineProps{
//   	definition: sfn.chain.start(sfn.NewPass(this, jsii.String("PassState"))),
//   })
//
//   // Include the state machine in a Task state with callback pattern
//   task := tasks.NewStepFunctionsStartExecution(this, jsii.String("ChildTask"), &stepFunctionsStartExecutionProps{
//   	stateMachine: child,
//   	integrationPattern: sfn.integrationPattern_WAIT_FOR_TASK_TOKEN,
//   	input: sfn.taskInput.fromObject(map[string]interface{}{
//   		"token": sfn.JsonPath.taskToken,
//   		"foo": jsii.String("bar"),
//   	}),
//   	name: jsii.String("MyExecutionName"),
//   })
//
//   // Define a second state machine with the Task state above
//   // Define a second state machine with the Task state above
//   sfn.NewStateMachine(this, jsii.String("ParentStateMachine"), &stateMachineProps{
//   	definition: task,
//   })
//
// Experimental.
type Chain interface {
	IChainable
	// The chainable end state(s) of this chain.
	// Experimental.
	EndStates() *[]INextable
	// Identify this Chain.
	// Experimental.
	Id() *string
	// The start state of this chain.
	// Experimental.
	StartState() State
	// Continue normal execution with the given state.
	// Experimental.
	Next(next IChainable) Chain
	// Return a single state that encompasses all states in the chain.
	//
	// This can be used to add error handling to a sequence of states.
	//
	// Be aware that this changes the result of the inner state machine
	// to be an array with the result of the state machine in it. Adjust
	// your paths accordingly. For example, change 'outputPath' to
	// '$[0]'.
	// Experimental.
	ToSingleState(id *string, props *ParallelProps) Parallel
}

// The jsii proxy struct for Chain
type jsiiProxy_Chain struct {
	jsiiProxy_IChainable
}

func (j *jsiiProxy_Chain) EndStates() *[]INextable {
	var returns *[]INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Chain) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Chain) StartState() State {
	var returns State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}


// Make a Chain with specific start and end states, and a last-added Chainable.
// Experimental.
func Chain_Custom(startState State, endStates *[]INextable, lastAdded IChainable) Chain {
	_init_.Initialize()

	if err := validateChain_CustomParameters(startState, endStates, lastAdded); err != nil {
		panic(err)
	}
	var returns Chain

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Chain",
		"custom",
		[]interface{}{startState, endStates, lastAdded},
		&returns,
	)

	return returns
}

// Make a Chain with the start from one chain and the ends from another.
// Experimental.
func Chain_Sequence(start IChainable, next IChainable) Chain {
	_init_.Initialize()

	if err := validateChain_SequenceParameters(start, next); err != nil {
		panic(err)
	}
	var returns Chain

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Chain",
		"sequence",
		[]interface{}{start, next},
		&returns,
	)

	return returns
}

// Begin a new Chain from one chainable.
// Experimental.
func Chain_Start(state IChainable) Chain {
	_init_.Initialize()

	if err := validateChain_StartParameters(state); err != nil {
		panic(err)
	}
	var returns Chain

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Chain",
		"start",
		[]interface{}{state},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Chain) Next(next IChainable) Chain {
	if err := c.validateNextParameters(next); err != nil {
		panic(err)
	}
	var returns Chain

	_jsii_.Invoke(
		c,
		"next",
		[]interface{}{next},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Chain) ToSingleState(id *string, props *ParallelProps) Parallel {
	if err := c.validateToSingleStateParameters(id, props); err != nil {
		panic(err)
	}
	var returns Parallel

	_jsii_.Invoke(
		c,
		"toSingleState",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

