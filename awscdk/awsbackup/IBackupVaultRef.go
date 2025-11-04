package awsbackup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbackup/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a BackupVault.
// Experimental.
type IBackupVaultRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a BackupVault resource.
	// Experimental.
	BackupVaultRef() *BackupVaultReference
}

// The jsii proxy for IBackupVaultRef
type jsiiProxy_IBackupVaultRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IBackupVaultRef) BackupVaultRef() *BackupVaultReference {
	var returns *BackupVaultReference
	_jsii_.Get(
		j,
		"backupVaultRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBackupVaultRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBackupVaultRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

