package awsiotwireless

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotwireless/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Destination.
// Experimental.
type IDestinationRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Destination resource.
	// Experimental.
	DestinationRef() *DestinationReference
}

// The jsii proxy for IDestinationRef
type jsiiProxy_IDestinationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IDestinationRef) DestinationRef() *DestinationReference {
	var returns *DestinationReference
	_jsii_.Get(
		j,
		"destinationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDestinationRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDestinationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

