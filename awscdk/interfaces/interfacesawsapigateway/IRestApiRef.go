package interfacesawsapigateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsapigateway/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RestApi.
// Experimental.
type IRestApiRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a RestApi resource.
	// Experimental.
	RestApiRef() *RestApiReference
}

// The jsii proxy for IRestApiRef
type jsiiProxy_IRestApiRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IRestApiRef) RestApiRef() *RestApiReference {
	var returns *RestApiReference
	_jsii_.Get(
		j,
		"restApiRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRestApiRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRestApiRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

