package awsbackup

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbackup/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LogicallyAirGappedBackupVault.
// Experimental.
type ILogicallyAirGappedBackupVaultRef interface {
	constructs.IConstruct
}

// The jsii proxy for ILogicallyAirGappedBackupVaultRef
type jsiiProxy_ILogicallyAirGappedBackupVaultRef struct {
	internal.Type__constructsIConstruct
}

