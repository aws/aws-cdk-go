package awsbackup

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbackup/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a BackupPlan.
// Experimental.
type IBackupPlanRef interface {
	constructs.IConstruct
}

// The jsii proxy for IBackupPlanRef
type jsiiProxy_IBackupPlanRef struct {
	internal.Type__constructsIConstruct
}

