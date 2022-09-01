package awsecs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// An extension for Task Definitions.
//
// Classes that want to make changes to a TaskDefinition (such as
// adding helper containers) can implement this interface, and can
// then be "added" to a TaskDefinition like so:
//
//     taskDefinition.addExtension(new MyExtension("some_parameter"));
type ITaskDefinitionExtension interface {
	// Apply the extension to the given TaskDefinition.
	Extend(taskDefinition TaskDefinition)
}

// The jsii proxy for ITaskDefinitionExtension
type jsiiProxy_ITaskDefinitionExtension struct {
	_ byte // padding
}

func (i *jsiiProxy_ITaskDefinitionExtension) Extend(taskDefinition TaskDefinition) {
	if err := i.validateExtendParameters(taskDefinition); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"extend",
		[]interface{}{taskDefinition},
	)
}

