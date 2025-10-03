package awspinpoint

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awspinpoint/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a BaiduChannel.
// Experimental.
type IBaiduChannelRef interface {
	constructs.IConstruct
}

// The jsii proxy for IBaiduChannelRef
type jsiiProxy_IBaiduChannelRef struct {
	internal.Type__constructsIConstruct
}

