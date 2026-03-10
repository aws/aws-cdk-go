package mixins

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface for applying properties to a target using a specific strategy.
// Experimental.
type IMergeStrategy interface {
	// Apply properties from source to target for the given keys.
	// Experimental.
	Apply(target *map[string]interface{}, source *map[string]interface{}, allowedKeys *[]*string)
}

// The jsii proxy for IMergeStrategy
type jsiiProxy_IMergeStrategy struct {
	_ byte // padding
}

func (i *jsiiProxy_IMergeStrategy) Apply(target *map[string]interface{}, source *map[string]interface{}, allowedKeys *[]*string) {
	if err := i.validateApplyParameters(target, source, allowedKeys); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"apply",
		[]interface{}{target, source, allowedKeys},
	)
}

