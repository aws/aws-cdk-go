package awsivs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsivs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PlaybackRestrictionPolicy.
// Experimental.
type IPlaybackRestrictionPolicyRef interface {
	constructs.IConstruct
	// A reference to a PlaybackRestrictionPolicy resource.
	// Experimental.
	PlaybackRestrictionPolicyRef() *PlaybackRestrictionPolicyReference
}

// The jsii proxy for IPlaybackRestrictionPolicyRef
type jsiiProxy_IPlaybackRestrictionPolicyRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IPlaybackRestrictionPolicyRef) PlaybackRestrictionPolicyRef() *PlaybackRestrictionPolicyReference {
	var returns *PlaybackRestrictionPolicyReference
	_jsii_.Get(
		j,
		"playbackRestrictionPolicyRef",
		&returns,
	)
	return returns
}

