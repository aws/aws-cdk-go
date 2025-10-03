package awsssm

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MaintenanceWindowTask.
// Experimental.
type IMaintenanceWindowTaskRef interface {
	constructs.IConstruct
}

// The jsii proxy for IMaintenanceWindowTaskRef
type jsiiProxy_IMaintenanceWindowTaskRef struct {
	internal.Type__constructsIConstruct
}

