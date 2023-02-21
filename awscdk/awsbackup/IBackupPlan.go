package awsbackup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbackup/internal"
)

// A backup plan.
type IBackupPlan interface {
	awscdk.IResource
	// The identifier of the backup plan.
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

