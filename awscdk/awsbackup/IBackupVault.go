package awsbackup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbackup/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsbackup"
	"github.com/aws/constructs-go/constructs/v10"
)

// A backup vault.
type IBackupVault interface {
	interfacesawsbackup.IBackupVaultRef
	awscdk.IResource
	// Grant the actions defined in actions to the given grantee on this backup vault.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// The ARN of the backup vault.
	BackupVaultArn() *string
	// The name of a logical container where backups are stored.
	BackupVaultName() *string
}

// The jsii proxy for IBackupVault
type jsiiProxy_IBackupVault struct {
	internal.Type__interfacesawsbackupIBackupVaultRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IBackupVault) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
	if err := i.validateGrantParameters(grantee); err != nil {
		panic(err)
	}
	args := []interface{}{grantee}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grant",
		args,
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IBackupVault) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IBackupVault) BackupVaultArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupVaultArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBackupVault) BackupVaultName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupVaultName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBackupVault) BackupVaultRef() *interfacesawsbackup.BackupVaultReference {
	var returns *interfacesawsbackup.BackupVaultReference
	_jsii_.Get(
		j,
		"backupVaultRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBackupVault) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBackupVault) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBackupVault) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

