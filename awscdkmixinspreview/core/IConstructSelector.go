package core

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Selects constructs from a construct tree.
// Experimental.
type IConstructSelector interface {
	// Selects constructs from the given scope based on the selector's criteria.
	// Experimental.
	Select(scope constructs.IConstruct) *[]constructs.IConstruct
}

// The jsii proxy for IConstructSelector
type jsiiProxy_IConstructSelector struct {
	_ byte // padding
}

func (i *jsiiProxy_IConstructSelector) Select(scope constructs.IConstruct) *[]constructs.IConstruct {
	if err := i.validateSelectParameters(scope); err != nil {
		panic(err)
	}
	var returns *[]constructs.IConstruct

	_jsii_.Invoke(
		i,
		"select",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

