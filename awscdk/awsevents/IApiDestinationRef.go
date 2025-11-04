package awsevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ApiDestination.
// Experimental.
type IApiDestinationRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ApiDestination resource.
	// Experimental.
	ApiDestinationRef() *ApiDestinationReference
}

// The jsii proxy for IApiDestinationRef
type jsiiProxy_IApiDestinationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IApiDestinationRef) ApiDestinationRef() *ApiDestinationReference {
	var returns *ApiDestinationReference
	_jsii_.Get(
		j,
		"apiDestinationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApiDestinationRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApiDestinationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

