package interfacesawswaf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawswaf/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a XssMatchSet.
// Experimental.
type IXssMatchSetRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a XssMatchSet resource.
	// Experimental.
	XssMatchSetRef() *XssMatchSetReference
}

// The jsii proxy for IXssMatchSetRef
type jsiiProxy_IXssMatchSetRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IXssMatchSetRef) XssMatchSetRef() *XssMatchSetReference {
	var returns *XssMatchSetReference
	_jsii_.Get(
		j,
		"xssMatchSetRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IXssMatchSetRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IXssMatchSetRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

