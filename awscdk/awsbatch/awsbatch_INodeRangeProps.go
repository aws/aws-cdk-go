package awsbatch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Properties for a multi-node batch job.
// Experimental.
type INodeRangeProps interface {
	// The container details for the node range.
	// Experimental.
	Container() *JobDefinitionContainer
	// Experimental.
	SetContainer(c *JobDefinitionContainer)
	// The minimum node index value to apply this container definition against.
	//
	// You may nest node ranges, for example 0:10 and 4:5, in which case the 4:5 range properties override the 0:10 properties.
	// Experimental.
	FromNodeIndex() *float64
	// Experimental.
	SetFromNodeIndex(f *float64)
	// The maximum node index value to apply this container definition against. If omitted, the highest value is used relative.
	//
	// to the number of nodes associated with the job. You may nest node ranges, for example 0:10 and 4:5,
	// in which case the 4:5 range properties override the 0:10 properties.
	// Experimental.
	ToNodeIndex() *float64
	// Experimental.
	SetToNodeIndex(t *float64)
}

// The jsii proxy for INodeRangeProps
type jsiiProxy_INodeRangeProps struct {
	_ byte // padding
}

func (j *jsiiProxy_INodeRangeProps) Container() *JobDefinitionContainer {
	var returns *JobDefinitionContainer
	_jsii_.Get(
		j,
		"container",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INodeRangeProps) SetContainer(val *JobDefinitionContainer) {
	_jsii_.Set(
		j,
		"container",
		val,
	)
}

func (j *jsiiProxy_INodeRangeProps) FromNodeIndex() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fromNodeIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INodeRangeProps) SetFromNodeIndex(val *float64) {
	_jsii_.Set(
		j,
		"fromNodeIndex",
		val,
	)
}

func (j *jsiiProxy_INodeRangeProps) ToNodeIndex() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"toNodeIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INodeRangeProps) SetToNodeIndex(val *float64) {
	_jsii_.Set(
		j,
		"toNodeIndex",
		val,
	)
}

