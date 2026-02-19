package interfacesawsapigatewayv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsapigatewayv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a IntegrationResponse.
// Experimental.
type IIntegrationResponseRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a IntegrationResponse resource.
	// Experimental.
	IntegrationResponseRef() *IntegrationResponseReference
}

// The jsii proxy for IIntegrationResponseRef
type jsiiProxy_IIntegrationResponseRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IIntegrationResponseRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IIntegrationResponseRef) IntegrationResponseRef() *IntegrationResponseReference {
	var returns *IntegrationResponseReference
	_jsii_.Get(
		j,
		"integrationResponseRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIntegrationResponseRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIntegrationResponseRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

