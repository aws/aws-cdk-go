package awstimestream

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awstimestream/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ScheduledQuery.
// Experimental.
type IScheduledQueryRef interface {
	constructs.IConstruct
}

// The jsii proxy for IScheduledQueryRef
type jsiiProxy_IScheduledQueryRef struct {
	internal.Type__constructsIConstruct
}

