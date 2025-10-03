package awsiot

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ScheduledAudit.
// Experimental.
type IScheduledAuditRef interface {
	constructs.IConstruct
}

// The jsii proxy for IScheduledAuditRef
type jsiiProxy_IScheduledAuditRef struct {
	internal.Type__constructsIConstruct
}

