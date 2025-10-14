package awsbackup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsbackup/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a BackupSelection.
// Experimental.
type IBackupSelectionRef interface {
	constructs.IConstruct
	// A reference to a BackupSelection resource.
	// Experimental.
	BackupSelectionRef() *BackupSelectionReference
}

// The jsii proxy for IBackupSelectionRef
type jsiiProxy_IBackupSelectionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IBackupSelectionRef) BackupSelectionRef() *BackupSelectionReference {
	var returns *BackupSelectionReference
	_jsii_.Get(
		j,
		"backupSelectionRef",
		&returns,
	)
	return returns
}

