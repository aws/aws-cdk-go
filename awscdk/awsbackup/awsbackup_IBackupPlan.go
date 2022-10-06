package awsbackup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsbackup/internal"
)

// A backup plan.
// Experimental.
type IBackupPlan interface {
	awscdk.IResource
	// The identifier of the backup plan.
	// Experimental.
	BackupPlanId() *string
}

// The jsii proxy for IBackupPlan
type jsiiProxy_IBackupPlan struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IBackupPlan) BackupPlanId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupPlanId",
		&returns,
	)
	return returns
}

