package interfacesawsssm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsssm/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MaintenanceWindowTask.
// Experimental.
type IMaintenanceWindowTaskRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a MaintenanceWindowTask resource.
	// Experimental.
	MaintenanceWindowTaskRef() *MaintenanceWindowTaskReference
}

// The jsii proxy for IMaintenanceWindowTaskRef
type jsiiProxy_IMaintenanceWindowTaskRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IMaintenanceWindowTaskRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
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

func (j *jsiiProxy_IMaintenanceWindowTaskRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
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

