package interfacesawsiot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CACertificate.
// Experimental.
type ICACertificateRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a CACertificate resource.
	// Experimental.
	CaCertificateRef() *CACertificateReference
}

// The jsii proxy for ICACertificateRef
type jsiiProxy_ICACertificateRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ICACertificateRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ICACertificateRef) CaCertificateRef() *CACertificateReference {
	var returns *CACertificateReference
	_jsii_.Get(
		j,
		"caCertificateRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICACertificateRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICACertificateRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

