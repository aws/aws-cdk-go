package awsevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Archive.
// Experimental.
type IArchiveRef interface {
	constructs.IConstruct
}

// The jsii proxy for IArchiveRef
type jsiiProxy_IArchiveRef struct {
	internal.Type__constructsIConstruct
}

