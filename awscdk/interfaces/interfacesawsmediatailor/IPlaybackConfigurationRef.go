package interfacesawsmediatailor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmediatailor/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PlaybackConfiguration.
// Experimental.
type IPlaybackConfigurationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a PlaybackConfiguration resource.
	// Experimental.
	PlaybackConfigurationRef() *PlaybackConfigurationReference
}

// The jsii proxy for IPlaybackConfigurationRef
type jsiiProxy_IPlaybackConfigurationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IPlaybackConfigurationRef) PlaybackConfigurationRef() *PlaybackConfigurationReference {
	var returns *PlaybackConfigurationReference
	_jsii_.Get(
		j,
		"playbackConfigurationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPlaybackConfigurationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPlaybackConfigurationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

