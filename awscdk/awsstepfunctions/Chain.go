package awsstepfunctions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A collection of states to chain onto.
//
// A Chain has a start and zero or more chainable ends. If there are
// zero ends, calling next() on the Chain will fail.
//
// Example:
//   map := sfn.NewMap(this, jsii.String("Map State"), &MapProps{
//   	MaxConcurrency: jsii.Number(1),
//   	ItemsPath: sfn.JsonPath_StringAt(jsii.String("$.inputForMap")),
//   	ItemSelector: map[string]interface{}{
//   		"item": sfn.JsonPath_*StringAt(jsii.String("$.Map.Item.Value")),
//   	},
//   	ResultPath: jsii.String("$.mapOutput"),
//   })
//
//   // The Map iterator can contain a IChainable, which can be an individual or multiple steps chained together.
//   // Below example is with a Choice and Pass step
//   choice := sfn.NewChoice(this, jsii.String("Choice"))
//   condition1 := sfn.Condition_StringEquals(jsii.String("$.item.status"), jsii.String("SUCCESS"))
//   step1 := sfn.NewPass(this, jsii.String("Step1"))
//   step2 := sfn.NewPass(this, jsii.String("Step2"))
//   finish := sfn.NewPass(this, jsii.String("Finish"))
//
//   definition := choice.When(condition1, step1).Otherwise(step2).Afterwards().Next(finish)
//
//   map.ItemProcessor(definition)
//
type Chain interface {
	IChainable
	// The chainable end state(s) of this chain.
	EndStates() *[]INextable
	// Identify this Chain.
	Id() *string
	// The start state of this chain.
	StartState() State
	// Continue normal execution with the given state.
	Next(next IChainable) Chain
	// Return a single state that encompasses all states in the chain.
	//
	// This can be used to add error handling to a sequence of states.
	//
	// Be aware that this changes the result of the inner state machine
	// to be an array with the result of the state machine in it. Adjust
	// your paths accordingly. For example, change 'outputPath' to
	// '$[0]'.
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
func Chain_Custom(startState State, endStates *[]INextable, lastAdded IChainable) Chain {
	_init_.Initialize()

	if err := validateChain_CustomParameters(startState, endStates, lastAdded); err != nil {
		panic(err)
	}
	var returns Chain

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Chain",
		"custom",
		[]interface{}{startState, endStates, lastAdded},
		&returns,
	)

	return returns
}

// Make a Chain with the start from one chain and the ends from another.
func Chain_Sequence(start IChainable, next IChainable) Chain {
	_init_.Initialize()

	if err := validateChain_SequenceParameters(start, next); err != nil {
		panic(err)
	}
	var returns Chain

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Chain",
		"sequence",
		[]interface{}{start, next},
		&returns,
	)

	return returns
}

// Begin a new Chain from one chainable.
func Chain_Start(state IChainable) Chain {
	_init_.Initialize()

	if err := validateChain_StartParameters(state); err != nil {
		panic(err)
	}
	var returns Chain

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Chain",
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

