package awsbatch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Properties for specifying multi-node properties for compute resources.
// Experimental.
type IMultiNodeProps interface {
	// The number of nodes associated with a multi-node parallel job.
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(c *float64)
	// Specifies the node index for the main node of a multi-node parallel job.
	//
	// This node index value must be fewer than the number of nodes.
	// Experimental.
	MainNode() *float64
	// Experimental.
	SetMainNode(m *float64)
	// A list of node ranges and their properties associated with a multi-node parallel job.
	// Experimental.
	RangeProps() *[]INodeRangeProps
	// Experimental.
	SetRangeProps(r *[]INodeRangeProps)
}

// The jsii proxy for IMultiNodeProps
type jsiiProxy_IMultiNodeProps struct {
	_ byte // padding
}

func (j *jsiiProxy_IMultiNodeProps) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMultiNodeProps)SetCount(val *float64) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_IMultiNodeProps) MainNode() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mainNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMultiNodeProps)SetMainNode(val *float64) {
	if err := j.validateSetMainNodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mainNode",
		val,
	)
}

func (j *jsiiProxy_IMultiNodeProps) RangeProps() *[]INodeRangeProps {
	var returns *[]INodeRangeProps
	_jsii_.Get(
		j,
		"rangeProps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMultiNodeProps)SetRangeProps(val *[]INodeRangeProps) {
	if err := j.validateSetRangePropsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rangeProps",
		val,
	)
}

