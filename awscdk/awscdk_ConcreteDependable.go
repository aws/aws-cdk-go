// An experiment to bundle the entire CDK into a single module
package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A set of constructs to be used as a dependable.
//
// This class can be used when a set of constructs which are disjoint in the
// construct tree needs to be combined to be used as a single dependable.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   concreteDependable := monocdk.NewConcreteDependable()
//
// Experimental.
type ConcreteDependable interface {
	IDependable
	// Add a construct to the dependency roots.
	// Experimental.
	Add(construct IConstruct)
}

// The jsii proxy struct for ConcreteDependable
type jsiiProxy_ConcreteDependable struct {
	jsiiProxy_IDependable
}

// Experimental.
func NewConcreteDependable() ConcreteDependable {
	_init_.Initialize()

	j := jsiiProxy_ConcreteDependable{}

	_jsii_.Create(
		"monocdk.ConcreteDependable",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewConcreteDependable_Override(c ConcreteDependable) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.ConcreteDependable",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_ConcreteDependable) Add(construct IConstruct) {
	if err := c.validateAddParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"add",
		[]interface{}{construct},
	)
}

