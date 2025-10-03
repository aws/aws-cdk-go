package awscustomerprofiles

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscustomerprofiles/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EventStream.
// Experimental.
type IEventStreamRef interface {
	constructs.IConstruct
}

// The jsii proxy for IEventStreamRef
type jsiiProxy_IEventStreamRef struct {
	internal.Type__constructsIConstruct
}

