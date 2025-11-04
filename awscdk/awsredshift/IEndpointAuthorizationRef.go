package awsredshift

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsredshift/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EndpointAuthorization.
// Experimental.
type IEndpointAuthorizationRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a EndpointAuthorization resource.
	// Experimental.
	EndpointAuthorizationRef() *EndpointAuthorizationReference
}

// The jsii proxy for IEndpointAuthorizationRef
type jsiiProxy_IEndpointAuthorizationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IEndpointAuthorizationRef) EndpointAuthorizationRef() *EndpointAuthorizationReference {
	var returns *EndpointAuthorizationReference
	_jsii_.Get(
		j,
		"endpointAuthorizationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEndpointAuthorizationRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEndpointAuthorizationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

