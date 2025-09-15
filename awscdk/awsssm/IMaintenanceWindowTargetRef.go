package awsssm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MaintenanceWindowTarget.
// Experimental.
type IMaintenanceWindowTargetRef interface {
	constructs.IConstruct
	// A reference to a MaintenanceWindowTarget resource.
	// Experimental.
	MaintenanceWindowTargetRef() *MaintenanceWindowTargetReference
}

// The jsii proxy for IMaintenanceWindowTargetRef
type jsiiProxy_IMaintenanceWindowTargetRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IMaintenanceWindowTargetRef) MaintenanceWindowTargetRef() *MaintenanceWindowTargetReference {
	var returns *MaintenanceWindowTargetReference
	_jsii_.Get(
		j,
		"maintenanceWindowTargetRef",
		&returns,
	)
	return returns
}

