package awscdkiotalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkiotalpha/v2/internal"
)

// Represents AWS IoT Scheduled Audit.
// Experimental.
type IScheduledAudit interface {
	awscdk.IResource
	// The ARN of the scheduled audit.
	// Experimental.
	ScheduledAuditArn() *string
	// The scheduled audit name.
	// Experimental.
	ScheduledAuditName() *string
}

// The jsii proxy for IScheduledAudit
type jsiiProxy_IScheduledAudit struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IScheduledAudit) ScheduledAuditArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduledAuditArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IScheduledAudit) ScheduledAuditName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduledAuditName",
		&returns,
	)
	return returns
}

