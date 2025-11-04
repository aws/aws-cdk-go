package awsssm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MaintenanceWindowTask.
// Experimental.
type IMaintenanceWindowTaskRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a MaintenanceWindowTask resource.
	// Experimental.
	MaintenanceWindowTaskRef() *MaintenanceWindowTaskReference
}

// The jsii proxy for IMaintenanceWindowTaskRef
type jsiiProxy_IMaintenanceWindowTaskRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IMaintenanceWindowTaskRef) MaintenanceWindowTaskRef() *MaintenanceWindowTaskReference {
	var returns *MaintenanceWindowTaskReference
	_jsii_.Get(
		j,
		"maintenanceWindowTaskRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMaintenanceWindowTaskRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMaintenanceWindowTaskRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

