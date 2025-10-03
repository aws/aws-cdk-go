package awsmediatailor

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmediatailor/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LiveSource.
// Experimental.
type ILiveSourceRef interface {
	constructs.IConstruct
}

// The jsii proxy for ILiveSourceRef
type jsiiProxy_ILiveSourceRef struct {
	internal.Type__constructsIConstruct
}

