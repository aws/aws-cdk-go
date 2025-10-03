package awsresourcegroups

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsresourcegroups/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TagSyncTask.
// Experimental.
type ITagSyncTaskRef interface {
	constructs.IConstruct
}

// The jsii proxy for ITagSyncTaskRef
type jsiiProxy_ITagSyncTaskRef struct {
	internal.Type__constructsIConstruct
}

