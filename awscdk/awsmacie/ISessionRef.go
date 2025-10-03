package awsmacie

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmacie/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Session.
// Experimental.
type ISessionRef interface {
	constructs.IConstruct
}

// The jsii proxy for ISessionRef
type jsiiProxy_ISessionRef struct {
	internal.Type__constructsIConstruct
}

