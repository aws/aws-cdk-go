package awseventschemas

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awseventschemas/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Registry.
// Experimental.
type IRegistryRef interface {
	constructs.IConstruct
	// A reference to a Registry resource.
	// Experimental.
	RegistryRef() *RegistryReference
}

// The jsii proxy for IRegistryRef
type jsiiProxy_IRegistryRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IRegistryRef) RegistryRef() *RegistryReference {
	var returns *RegistryReference
	_jsii_.Get(
		j,
		"registryRef",
		&returns,
	)
	return returns
}

