package awsssm

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MaintenanceWindow.
// Experimental.
type IMaintenanceWindowRef interface {
	constructs.IConstruct
}

// The jsii proxy for IMaintenanceWindowRef
type jsiiProxy_IMaintenanceWindowRef struct {
	internal.Type__constructsIConstruct
}

