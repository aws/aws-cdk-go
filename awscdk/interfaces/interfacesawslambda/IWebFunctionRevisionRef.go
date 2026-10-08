package interfacesawslambda

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslambda/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a WebFunctionRevision.
// Experimental.
type IWebFunctionRevisionRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a WebFunctionRevision resource.
	// Experimental.
	WebFunctionRevisionRef() *WebFunctionRevisionReference
}

// The jsii proxy for IWebFunctionRevisionRef
type jsiiProxy_IWebFunctionRevisionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IWebFunctionRevisionRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IWebFunctionRevisionRef) WebFunctionRevisionRef() *WebFunctionRevisionReference {
	var returns *WebFunctionRevisionReference
	_jsii_.Get(
		j,
		"webFunctionRevisionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWebFunctionRevisionRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWebFunctionRevisionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

