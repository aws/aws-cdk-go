package awsbackup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsbackup/internal"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// A backup vault.
// Experimental.
type IBackupVault interface {
	awscdk.IResource
	// Grant the actions defined in actions to the given grantee on this backup vault.
	// Experimental.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// The ARN of the backup vault.
	// Experimental.
	BackupVaultArn() *string
	// The name of a logical container where backups are stored.
	// Experimental.
	BackupVaultName() *string
}

// The jsii proxy for IBackupVault
type jsiiProxy_IBackupVault struct {
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

