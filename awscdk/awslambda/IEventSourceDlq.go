package awslambda

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A DLQ for an event source.
type IEventSourceDlq interface {
	// Returns the DLQ destination config of the DLQ.
	Bind(target IEventSourceMapping, targetHandler IFunction) *DlqDestinationConfig
}

// The jsii proxy for IEventSourceDlq
type jsiiProxy_IEventSourceDlq struct {
	_ byte // padding
}

func (i *jsiiProxy_IEventSourceDlq) Bind(target IEventSourceMapping, targetHandler IFunction) *DlqDestinationConfig {
	if err := i.validateBindParameters(target, targetHandler); err != nil {
		panic(err)
	}
	var returns *DlqDestinationConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{target, targetHandler},
		&returns,
	)

	return returns
}

