package interfacesawsbackup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsbackup/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a BackupPlan.
// Experimental.
type IBackupPlanRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a BackupPlan resource.
	// Experimental.
	BackupPlanRef() *BackupPlanReference
}

// The jsii proxy for IBackupPlanRef
type jsiiProxy_IBackupPlanRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
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

func (j *jsiiProxy_IBackupPlanRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBackupPlanRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

