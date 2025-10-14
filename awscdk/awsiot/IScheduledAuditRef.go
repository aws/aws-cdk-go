package awsiot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ScheduledAudit.
// Experimental.
type IScheduledAuditRef interface {
	constructs.IConstruct
	// A reference to a ScheduledAudit resource.
	// Experimental.
	ScheduledAuditRef() *ScheduledAuditReference
}

// The jsii proxy for IScheduledAuditRef
type jsiiProxy_IScheduledAuditRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IScheduledAuditRef) ScheduledAuditRef() *ScheduledAuditReference {
	var returns *ScheduledAuditReference
	_jsii_.Get(
		j,
		"scheduledAuditRef",
		&returns,
	)
	return returns
}

