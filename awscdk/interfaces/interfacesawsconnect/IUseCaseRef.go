package interfacesawsconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a UseCase.
// Experimental.
type IUseCaseRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a UseCase resource.
	// Experimental.
	UseCaseRef() *UseCaseReference
}

// The jsii proxy for IUseCaseRef
type jsiiProxy_IUseCaseRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IUseCaseRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IUseCaseRef) UseCaseRef() *UseCaseReference {
	var returns *UseCaseReference
	_jsii_.Get(
		j,
		"useCaseRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUseCaseRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUseCaseRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

