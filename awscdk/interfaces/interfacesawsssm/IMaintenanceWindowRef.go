package interfacesawsssm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsssm/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MaintenanceWindow.
// Experimental.
type IMaintenanceWindowRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a MaintenanceWindow resource.
	// Experimental.
	MaintenanceWindowRef() *MaintenanceWindowReference
}

// The jsii proxy for IMaintenanceWindowRef
type jsiiProxy_IMaintenanceWindowRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IMaintenanceWindowRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IMaintenanceWindowRef) MaintenanceWindowRef() *MaintenanceWindowReference {
	var returns *MaintenanceWindowReference
	_jsii_.Get(
		j,
		"maintenanceWindowRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMaintenanceWindowRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMaintenanceWindowRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

