package interfacesawssam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssam/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a HttpApi.
// Experimental.
type IHttpApiRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a HttpApi resource.
	// Experimental.
	HttpApiRef() *HttpApiReference
}

// The jsii proxy for IHttpApiRef
type jsiiProxy_IHttpApiRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IHttpApiRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IHttpApiRef) HttpApiRef() *HttpApiReference {
	var returns *HttpApiReference
	_jsii_.Get(
		j,
		"httpApiRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IHttpApiRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IHttpApiRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

