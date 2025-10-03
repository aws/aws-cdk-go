package awsamplify

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsamplify/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a App.
// Experimental.
type IAppRef interface {
	constructs.IConstruct
}

// The jsii proxy for IAppRef
type jsiiProxy_IAppRef struct {
	internal.Type__constructsIConstruct
}

