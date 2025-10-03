package awskinesis

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesis/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Stream.
// Experimental.
type IStreamRef interface {
	constructs.IConstruct
}

// The jsii proxy for IStreamRef
type jsiiProxy_IStreamRef struct {
	internal.Type__constructsIConstruct
}

