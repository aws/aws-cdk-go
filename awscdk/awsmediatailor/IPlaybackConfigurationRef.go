package awsmediatailor

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmediatailor/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PlaybackConfiguration.
// Experimental.
type IPlaybackConfigurationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IPlaybackConfigurationRef
type jsiiProxy_IPlaybackConfigurationRef struct {
	internal.Type__constructsIConstruct
}

