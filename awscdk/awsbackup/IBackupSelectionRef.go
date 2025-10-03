package awsbackup

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbackup/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a BackupSelection.
// Experimental.
type IBackupSelectionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IBackupSelectionRef
type jsiiProxy_IBackupSelectionRef struct {
	internal.Type__constructsIConstruct
}

