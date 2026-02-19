package interfacesawsrolesanywhere

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsrolesanywhere/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CRL.
// Experimental.
type ICRLRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a CRL resource.
	// Experimental.
	CrlRef() *CRLReference
}

// The jsii proxy for ICRLRef
type jsiiProxy_ICRLRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ICRLRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ICRLRef) CrlRef() *CRLReference {
	var returns *CRLReference
	_jsii_.Get(
		j,
		"crlRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICRLRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICRLRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

