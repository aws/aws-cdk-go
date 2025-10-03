package awsbackup

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbackup/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a BackupVault.
// Experimental.
type IBackupVaultRef interface {
	constructs.IConstruct
}

// The jsii proxy for IBackupVaultRef
type jsiiProxy_IBackupVaultRef struct {
	internal.Type__constructsIConstruct
}

