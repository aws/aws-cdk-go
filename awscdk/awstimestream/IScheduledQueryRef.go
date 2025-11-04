package awstimestream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awstimestream/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ScheduledQuery.
// Experimental.
type IScheduledQueryRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ScheduledQuery resource.
	// Experimental.
	ScheduledQueryRef() *ScheduledQueryReference
}

// The jsii proxy for IScheduledQueryRef
type jsiiProxy_IScheduledQueryRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IScheduledQueryRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IScheduledQueryRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

