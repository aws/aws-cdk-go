package interfacesawsapigatewayv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsapigatewayv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ApiGatewayManagedOverrides.
// Experimental.
type IApiGatewayManagedOverridesRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ApiGatewayManagedOverrides resource.
	// Experimental.
	ApiGatewayManagedOverridesRef() *ApiGatewayManagedOverridesReference
}

// The jsii proxy for IApiGatewayManagedOverridesRef
type jsiiProxy_IApiGatewayManagedOverridesRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IApiGatewayManagedOverridesRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IApiGatewayManagedOverridesRef) ApiGatewayManagedOverridesRef() *ApiGatewayManagedOverridesReference {
	var returns *ApiGatewayManagedOverridesReference
	_jsii_.Get(
		j,
		"apiGatewayManagedOverridesRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApiGatewayManagedOverridesRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApiGatewayManagedOverridesRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

