package awslambda

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A schema registry for an event source.
type ISchemaRegistry interface {
	// Returns the schema registry config of the event source.
	Bind(target IEventSourceMapping, targetHandler IFunction) *KafkaSchemaRegistryConfig
}

// The jsii proxy for ISchemaRegistry
type jsiiProxy_ISchemaRegistry struct {
	_ byte // padding
}

func (i *jsiiProxy_ISchemaRegistry) Bind(target IEventSourceMapping, targetHandler IFunction) *KafkaSchemaRegistryConfig {
	if err := i.validateBindParameters(target, targetHandler); err != nil {
		panic(err)
	}
	var returns *KafkaSchemaRegistryConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{target, targetHandler},
		&returns,
	)

	return returns
}

