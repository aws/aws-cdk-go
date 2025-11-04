package awsivs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsivs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PlaybackKeyPair.
// Experimental.
type IPlaybackKeyPairRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a PlaybackKeyPair resource.
	// Experimental.
	PlaybackKeyPairRef() *PlaybackKeyPairReference
}

// The jsii proxy for IPlaybackKeyPairRef
type jsiiProxy_IPlaybackKeyPairRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IPlaybackKeyPairRef) PlaybackKeyPairRef() *PlaybackKeyPairReference {
	var returns *PlaybackKeyPairReference
	_jsii_.Get(
		j,
		"playbackKeyPairRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPlaybackKeyPairRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPlaybackKeyPairRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

