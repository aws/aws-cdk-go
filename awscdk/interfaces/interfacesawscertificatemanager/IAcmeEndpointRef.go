package interfacesawscertificatemanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscertificatemanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AcmeEndpoint.
// Experimental.
type IAcmeEndpointRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a AcmeEndpoint resource.
	// Experimental.
	AcmeEndpointRef() *AcmeEndpointReference
}

// The jsii proxy for IAcmeEndpointRef
type jsiiProxy_IAcmeEndpointRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IAcmeEndpointRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IAcmeEndpointRef) AcmeEndpointRef() *AcmeEndpointReference {
	var returns *AcmeEndpointReference
	_jsii_.Get(
		j,
		"acmeEndpointRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAcmeEndpointRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAcmeEndpointRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

