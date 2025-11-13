package interfacesawscognito

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscognito/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Terms.
// Experimental.
type ITermsRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Terms resource.
	// Experimental.
	TermsRef() *TermsReference
}

// The jsii proxy for ITermsRef
type jsiiProxy_ITermsRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_ITermsRef) TermsRef() *TermsReference {
	var returns *TermsReference
	_jsii_.Get(
		j,
		"termsRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITermsRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITermsRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

