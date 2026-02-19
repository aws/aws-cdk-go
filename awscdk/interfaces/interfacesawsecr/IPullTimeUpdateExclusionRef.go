package interfacesawsecr

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsecr/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PullTimeUpdateExclusion.
// Experimental.
type IPullTimeUpdateExclusionRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a PullTimeUpdateExclusion resource.
	// Experimental.
	PullTimeUpdateExclusionRef() *PullTimeUpdateExclusionReference
}

// The jsii proxy for IPullTimeUpdateExclusionRef
type jsiiProxy_IPullTimeUpdateExclusionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IPullTimeUpdateExclusionRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IPullTimeUpdateExclusionRef) PullTimeUpdateExclusionRef() *PullTimeUpdateExclusionReference {
	var returns *PullTimeUpdateExclusionReference
	_jsii_.Get(
		j,
		"pullTimeUpdateExclusionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPullTimeUpdateExclusionRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPullTimeUpdateExclusionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

