package interfacesawsapigateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsapigateway/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ClientCertificate.
// Experimental.
type IClientCertificateRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ClientCertificate resource.
	// Experimental.
	ClientCertificateRef() *ClientCertificateReference
}

// The jsii proxy for IClientCertificateRef
type jsiiProxy_IClientCertificateRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IClientCertificateRef) ClientCertificateRef() *ClientCertificateReference {
	var returns *ClientCertificateReference
	_jsii_.Get(
		j,
		"clientCertificateRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IClientCertificateRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IClientCertificateRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

