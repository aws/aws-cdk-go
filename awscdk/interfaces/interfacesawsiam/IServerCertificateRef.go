package interfacesawsiam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsiam/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ServerCertificate.
// Experimental.
type IServerCertificateRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ServerCertificate resource.
	// Experimental.
	ServerCertificateRef() *ServerCertificateReference
}

// The jsii proxy for IServerCertificateRef
type jsiiProxy_IServerCertificateRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IServerCertificateRef) ServerCertificateRef() *ServerCertificateReference {
	var returns *ServerCertificateReference
	_jsii_.Get(
		j,
		"serverCertificateRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServerCertificateRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServerCertificateRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

