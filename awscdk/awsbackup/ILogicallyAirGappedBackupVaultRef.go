package awsbackup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsbackup/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LogicallyAirGappedBackupVault.
// Experimental.
type ILogicallyAirGappedBackupVaultRef interface {
	constructs.IConstruct
	// A reference to a LogicallyAirGappedBackupVault resource.
	// Experimental.
	LogicallyAirGappedBackupVaultRef() *LogicallyAirGappedBackupVaultReference
}

// The jsii proxy for ILogicallyAirGappedBackupVaultRef
type jsiiProxy_ILogicallyAirGappedBackupVaultRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ILogicallyAirGappedBackupVaultRef) LogicallyAirGappedBackupVaultRef() *LogicallyAirGappedBackupVaultReference {
	var returns *LogicallyAirGappedBackupVaultReference
	_jsii_.Get(
		j,
		"logicallyAirGappedBackupVaultRef",
		&returns,
	)
	return returns
}

