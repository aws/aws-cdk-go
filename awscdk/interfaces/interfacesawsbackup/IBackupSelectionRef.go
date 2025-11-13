package interfacesawsbackup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsbackup/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a BackupSelection.
// Experimental.
type IBackupSelectionRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a BackupSelection resource.
	// Experimental.
	BackupSelectionRef() *BackupSelectionReference
}

// The jsii proxy for IBackupSelectionRef
type jsiiProxy_IBackupSelectionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
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

func (j *jsiiProxy_IBackupSelectionRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBackupSelectionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

