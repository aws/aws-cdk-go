package interfacesawsapigateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsapigateway/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RequestValidator.
// Experimental.
type IRequestValidatorRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a RequestValidator resource.
	// Experimental.
	RequestValidatorRef() *RequestValidatorReference
}

// The jsii proxy for IRequestValidatorRef
type jsiiProxy_IRequestValidatorRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IRequestValidatorRef) RequestValidatorRef() *RequestValidatorReference {
	var returns *RequestValidatorReference
	_jsii_.Get(
		j,
		"requestValidatorRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRequestValidatorRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRequestValidatorRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

