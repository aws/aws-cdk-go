package pipelines

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Any class that produces, or is itself, a `FileSet`.
//
// Steps implicitly produce a primary FileSet as an output.
type IFileSetProducer interface {
	// The `FileSet` produced by this file set producer.
	PrimaryOutput() FileSet
}

// The jsii proxy for IFileSetProducer
type jsiiProxy_IFileSetProducer struct {
	_ byte // padding
}

func (j *jsiiProxy_IFileSetProducer) PrimaryOutput() FileSet {
	var returns FileSet
	_jsii_.Get(
		j,
		"primaryOutput",
		&returns,
	)
	return returns
}

