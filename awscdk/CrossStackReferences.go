package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Ergonomic API for configuring cross-stack reference strength on a construct.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   crossStackReferences := cdk.CrossStackReferences_Of(this)
//
type CrossStackReferences interface {
	// Set the default reference strength used when this scope consumes references from other stacks.
	//
	// This controls the consuming side: sets the context key that determines how
	// incoming cross-stack references are resolved for this scope and its descendants.
	//
	// Equivalent to `scope.node.setContext(DEFAULT_CROSS_STACK_REFERENCES, strength)`.
	Consume(strength ReferenceStrength)
	// Set how this resource is referenced when consumed from another stack.
	//
	// This controls the producing side: any cross-stack reference pointing at
	// this resource will use the specified strength instead of the global default.
	//
	// Equivalent to `scope.applyCrossStackReferenceStrength(strength)`.
	Produce(strength ReferenceStrength)
}

// The jsii proxy struct for CrossStackReferences
type jsiiProxy_CrossStackReferences struct {
	_ byte // padding
}

// Returns a `CrossStackReferences` configurator for the given construct.
func CrossStackReferences_Of(scope constructs.IConstruct) CrossStackReferences {
	_init_.Initialize()

	if err := validateCrossStackReferences_OfParameters(scope); err != nil {
		panic(err)
	}
	var returns CrossStackReferences

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CrossStackReferences",
		"of",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CrossStackReferences) Consume(strength ReferenceStrength) {
	if err := c.validateConsumeParameters(strength); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"consume",
		[]interface{}{strength},
	)
}

func (c *jsiiProxy_CrossStackReferences) Produce(strength ReferenceStrength) {
	if err := c.validateProduceParameters(strength); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"produce",
		[]interface{}{strength},
	)
}

