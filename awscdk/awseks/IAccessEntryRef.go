package awseks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awseks/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AccessEntry.
// Experimental.
type IAccessEntryRef interface {
	constructs.IConstruct
}

// The jsii proxy for IAccessEntryRef
type jsiiProxy_IAccessEntryRef struct {
	internal.Type__constructsIConstruct
}

