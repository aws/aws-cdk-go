package interfacesawsefs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsefs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MountTarget.
// Experimental.
type IMountTargetRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a MountTarget resource.
	// Experimental.
	MountTargetRef() *MountTargetReference
}

// The jsii proxy for IMountTargetRef
type jsiiProxy_IMountTargetRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IMountTargetRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IMountTargetRef) MountTargetRef() *MountTargetReference {
	var returns *MountTargetReference
	_jsii_.Get(
		j,
		"mountTargetRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMountTargetRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMountTargetRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

