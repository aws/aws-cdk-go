package awsnimblestudio

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsnimblestudio/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a StudioComponent.
// Experimental.
type IStudioComponentRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a StudioComponent resource.
	// Experimental.
	StudioComponentRef() *StudioComponentReference
}

// The jsii proxy for IStudioComponentRef
type jsiiProxy_IStudioComponentRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IStudioComponentRef) StudioComponentRef() *StudioComponentReference {
	var returns *StudioComponentReference
	_jsii_.Get(
		j,
		"studioComponentRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStudioComponentRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStudioComponentRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

