package interfacesawsnimblestudio

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsnimblestudio/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LaunchProfile.
// Experimental.
type ILaunchProfileRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a LaunchProfile resource.
	// Experimental.
	LaunchProfileRef() *LaunchProfileReference
}

// The jsii proxy for ILaunchProfileRef
type jsiiProxy_ILaunchProfileRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ILaunchProfileRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ILaunchProfileRef) LaunchProfileRef() *LaunchProfileReference {
	var returns *LaunchProfileReference
	_jsii_.Get(
		j,
		"launchProfileRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILaunchProfileRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILaunchProfileRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

