package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface for merging arrays.
type IArrayMergeStrategy interface {
	// Merge source array into target array.
	Merge(target *[]interface{}, source *[]interface{}) *[]interface{}
}

// The jsii proxy for IArrayMergeStrategy
type jsiiProxy_IArrayMergeStrategy struct {
	_ byte // padding
}

func (i *jsiiProxy_IArrayMergeStrategy) Merge(target *[]interface{}, source *[]interface{}) *[]interface{} {
	if err := i.validateMergeParameters(target, source); err != nil {
		panic(err)
	}
	var returns *[]interface{}

	_jsii_.Invoke(
		i,
		"merge",
		[]interface{}{target, source},
		&returns,
	)

	return returns
}

