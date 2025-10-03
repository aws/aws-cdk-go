package awsivs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsivs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PlaybackKeyPair.
// Experimental.
type IPlaybackKeyPairRef interface {
	constructs.IConstruct
}

// The jsii proxy for IPlaybackKeyPairRef
type jsiiProxy_IPlaybackKeyPairRef struct {
	internal.Type__constructsIConstruct
}

