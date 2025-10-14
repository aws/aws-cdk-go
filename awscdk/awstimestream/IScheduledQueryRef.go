package awstimestream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awstimestream/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ScheduledQuery.
// Experimental.
type IScheduledQueryRef interface {
	constructs.IConstruct
	// A reference to a ScheduledQuery resource.
	// Experimental.
	ScheduledQueryRef() *ScheduledQueryReference
}

// The jsii proxy for IScheduledQueryRef
type jsiiProxy_IScheduledQueryRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IScheduledQueryRef) ScheduledQueryRef() *ScheduledQueryReference {
	var returns *ScheduledQueryReference
	_jsii_.Get(
		j,
		"scheduledQueryRef",
		&returns,
	)
	return returns
}

