package awsivs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsivs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PlaybackRestrictionPolicy.
// Experimental.
type IPlaybackRestrictionPolicyRef interface {
	constructs.IConstruct
}

// The jsii proxy for IPlaybackRestrictionPolicyRef
type jsiiProxy_IPlaybackRestrictionPolicyRef struct {
	internal.Type__constructsIConstruct
}

