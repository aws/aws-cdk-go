package interfacesawsbackup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsbackup/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LogicallyAirGappedBackupVault.
// Experimental.
type ILogicallyAirGappedBackupVaultRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a LogicallyAirGappedBackupVault resource.
	// Experimental.
	LogicallyAirGappedBackupVaultRef() *LogicallyAirGappedBackupVaultReference
}

// The jsii proxy for ILogicallyAirGappedBackupVaultRef
type jsiiProxy_ILogicallyAirGappedBackupVaultRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
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

func (j *jsiiProxy_ILogicallyAirGappedBackupVaultRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILogicallyAirGappedBackupVaultRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

