package awsbackup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsbackup/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a BackupPlan.
// Experimental.
type IBackupPlanRef interface {
	constructs.IConstruct
	// A reference to a BackupPlan resource.
	// Experimental.
	BackupPlanRef() *BackupPlanReference
}

// The jsii proxy for IBackupPlanRef
type jsiiProxy_IBackupPlanRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IBackupPlanRef) BackupPlanRef() *BackupPlanReference {
	var returns *BackupPlanReference
	_jsii_.Get(
		j,
		"backupPlanRef",
		&returns,
	)
	return returns
}

