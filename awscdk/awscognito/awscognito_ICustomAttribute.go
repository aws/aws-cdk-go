package awscognito

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents a custom attribute type.
type ICustomAttribute interface {
	// Bind this custom attribute type to the values as expected by CloudFormation.
	Bind() *CustomAttributeConfig
}

// The jsii proxy for ICustomAttribute
type jsiiProxy_ICustomAttribute struct {
	_ byte // padding
}

func (i *jsiiProxy_ICustomAttribute) Bind() *CustomAttributeConfig {
	var returns *CustomAttributeConfig

	_jsii_.Invoke(
		i,
		"bind",
		nil, // no parameters
		&returns,
	)

	return returns
}

