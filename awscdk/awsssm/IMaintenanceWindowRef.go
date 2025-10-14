package awsssm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MaintenanceWindow.
// Experimental.
type IMaintenanceWindowRef interface {
	constructs.IConstruct
	// A reference to a MaintenanceWindow resource.
	// Experimental.
	MaintenanceWindowRef() *MaintenanceWindowReference
}

// The jsii proxy for IMaintenanceWindowRef
type jsiiProxy_IMaintenanceWindowRef struct {
	internal.Type__constructsIConstruct
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

